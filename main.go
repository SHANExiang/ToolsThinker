package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ToolsThinker/myTheme"
	"image/color"
	"log"
)

//func init() {
//	//设置中文字体:解决中文乱码问题
//	fontPaths := findfont.List()
//	for _, path := range fontPaths {
//		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") ||
//			strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
//			os.Setenv("FYNE_FONT", path)
//			break
//		}
//	}
//}

func main() {
	mainApp := app.New()
	mainApp.Settings().SetTheme(myTheme.NewMyTheme())
	w := mainApp.NewWindow("ToolsThinker") // 初始化窗口对象
	w.Resize(fyne.NewSize(800, 600))       // 设置窗口尺寸

	// 右侧主内容
	resource, err := fyne.LoadResourceFromPath("default.jpg")
	if err != nil {
		log.Fatal(err)
	}
	img := canvas.NewImageFromResource(resource)
	img.FillMode = canvas.ImageFillContain // 等比例缩放并居中显示，但不会拉伸填满整个区域。
	img.SetMinSize(fyne.NewSize(400, 300)) // 可选：设置最小尺寸以撑开容器
	rightContent := container.NewMax(img)

	// 左侧固定宽度区域
	leftWidth := float32(150)
	leftContent := container.NewVBox(
		widget.NewLabel("菜单"),
		widget.NewButton("excel", func() {
			newContent := canvas.NewText("你打开了excel", color.RGBA{0, 0, 255, 255})
			newContent.TextStyle = fyne.TextStyle{Bold: true}
			rightContent.Objects = []fyne.CanvasObject{newContent}
			rightContent.Refresh()
		}),
		widget.NewButton("word", func() {}),
	)
	leftFixed := container.NewMax(leftContent)
	leftFixed.Resize(fyne.NewSize(leftWidth, 0)) // 设置初始宽度

	// 分隔线（细线）
	separator := canvas.NewRectangle(color.Gray{Y: 180}) // 浅灰色
	separator.SetMinSize(fyne.NewSize(1, 0))             // 宽度为1，高度自动

	// 左 + 分隔线 + 右 的水平排列
	content := container.NewHBox(
		leftFixed,
		separator,
		rightContent,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
