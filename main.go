package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func init() {
	//设置中文字体:解决中文乱码问题
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	mainApp := app.New()
	mainApp.Settings().SetTheme(theme.LightTheme())
	w := mainApp.NewWindow("ExcelThinker") // 初始化窗口对象
	w.Resize(fyne.NewSize(800, 600))       // 设置窗口尺寸

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
