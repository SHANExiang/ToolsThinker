package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"tools-thinker/myTheme"
)

var App fyne.App
var Window fyne.Window

func InitApp() {
	App = app.New()
	App.Settings().SetTheme(myTheme.NewMyTheme())
	Window = App.NewWindow("tools-thinker") // 初始化窗口对象
	Window.Resize(fyne.NewSize(800, 600))   // 设置窗口尺寸
}

func setWindows() {

}
