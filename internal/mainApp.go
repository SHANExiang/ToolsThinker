package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"tools-thinker/internal/layout/left"
	"tools-thinker/internal/layout/right"
	"tools-thinker/myTheme"
)

var App fyne.App
var Window fyne.Window

func InitApp() {
	App = app.New()
	App.Settings().SetTheme(myTheme.NewMyTheme())
	Window = App.NewWindow("tools-thinker") // 初始化窗口对象

	left.Init()  // 左侧区域初始化
	right.Init() // 右侧区域初始化

	// ----- Main layout: Left menu + Right panel -----
	mainLayout := container.NewHSplit(
		left.Content,  // left
		right.Content, // right: dynamically fills space / grows with window
	)

	mainLayout.Offset = 0.22 // Left panel = ~22%, Right panel = 78%
	Window.SetContent(mainLayout)

	Window.Resize(fyne.NewSize(800, 600)) // 设置窗口尺寸
}

func setWindows() {

}
