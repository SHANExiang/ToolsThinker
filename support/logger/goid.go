package logger

// 官方http2 使用的goid方法 https://github.com/golang/net/blob/master/http2/gotrack.go#L51
import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var DebugGoroutines = os.Getenv("DEBUG_HTTP2_GOROUTINES") == "1"

type goroutineLock uint64

func newGoroutineLock() goroutineLock {
	if !DebugGoroutines {
		return 0
	}
	return goroutineLock(curGoroutineID())
}

func (g goroutineLock) check() {
	if !DebugGoroutines {
		return
	}
	if curGoroutineID() != uint64(g) {
		panic("running on the wrong goroutine")
	}
}

func (g goroutineLock) checkNotOn() {
	if !DebugGoroutines {
		return
	}
	if curGoroutineID() == uint64(g) {
		panic("running on the wrong goroutine")
	}
}

var goroutineSpace = []byte("goroutine ")

func CurGoroutineID() uint64 {
	return curGoroutineID()
}

func curGoroutineID() uint64 {
	bp := littleBuf.Get().(*[]byte)
	defer littleBuf.Put(bp)
	b := *bp
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in \"%s\"", b))
	}
	b = b[:i]
	n, err := parseUintBytes(b, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of \"%s\": %v", b, err))
	}
	return n
}

var littleBuf = sync.Pool{
	New: func() any {
		buf := make([]byte, 64)
		return &buf
	},
}

// like strconv.ParseUint, but using a []byte.
func parseUintBytes(s []byte, base int, bitSize int) (n uint64, err error) {
	var cutoff, maxVal uint64

	if bitSize == 0 {
		bitSize = int(strconv.IntSize)
	}

	s0 := s
	switch {
	case len(s) < 1:
		err = strconv.ErrSyntax
		goto Error

	case 2 <= base && base <= 36:
		// valid base; nothing to do

	case base == 0:
		// Look for octal, hex prefix.
		if s[0] == '0' && len(s) > 1 && (s[1] == 'x' || s[1] == 'X') {
			base = 16
			s = s[2:]
			if len(s) < 1 {
				err = strconv.ErrSyntax
				goto Error
			}
		} else if s[0] == '0' {
			base = 8
		} else {
			base = 10
		}

	default:
		err = errors.New("invalid base " + strconv.Itoa(base))
		goto Error
	}

	n = 0
	cutoff = cutoff64(base)
	maxVal = 1<<uint(bitSize) - 1

	for i := 0; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			n = 0
			err = strconv.ErrSyntax
			goto Error
		}
		if int(v) >= base {
			n = 0
			err = strconv.ErrSyntax
			goto Error
		}

		if n >= cutoff {
			// n*base overflows
			n = 1<<64 - 1
			err = strconv.ErrRange
			goto Error
		}
		n *= uint64(base)

		n1 := n + uint64(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = 1<<64 - 1
			err = strconv.ErrRange
			goto Error
		}
		n = n1
	}

	return n, nil

Error:
	return n, &strconv.NumError{Func: "ParseUint", Num: string(s0), Err: err}
}

// Return the first number n such that n*base >= 1<<64.
func cutoff64(base int) uint64 {
	if base < 2 {
		return 0
	}
	return (1<<64-1)/uint64(base) + 1
}
