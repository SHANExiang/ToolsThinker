package left

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"tools-thinker/internal/layout/excel"
)

// @Author dx
// @Date 2025-07-22 22:34:00
// @Desc

var Content *fyne.Container

func Init(window fyne.Window) {
	Content = container.NewVBox(
		widget.NewLabel("📋 Menu"),
		widget.NewButton("Excel", func() {
			excel.Content(window)
		}),
		widget.NewButton("Word", func() {
			//newContent := canvas.NewText("你打开了word", color.RGBA{0, 0, 255, 255})
			//newContent.TextStyle = fyne.TextStyle{Bold: true}
			//rightContent.Objects = []fyne.CanvasObject{newContent}
			//rightContent.Refresh()
		}),
	)

}
