package common

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
)

func OpenFileButton(window fyne.Window) *widget.Button {
	openButton := widget.NewButton("打开文件", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader == nil {
				return // 用户取消了文件选择
			}
			defer reader.Close()

			// 读取文件内容（可选）
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			// 打印文件内容或名称
			log.Println("文件名:", reader.URI().Name())
			log.Println("文件内容:")
			log.Println(string(data))

		}, window)
	})
	return openButton
}
