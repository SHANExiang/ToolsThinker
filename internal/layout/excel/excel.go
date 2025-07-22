package excel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"tools-thinker/internal/layout/right"
)

// @Author dx
// @Date 2025-07-22 22:18:00
// @Desc

func Content() {
	// 点击左侧按钮后，让右侧显示新的按钮区
	btn1 := widget.NewButton("操作1", func() {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "操作1",
			Content: "你点击了操作1！",
		})
	})

	btn2 := widget.NewButton("操作2", func() {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "操作2",
			Content: "你点击了操作2！",
		})
	})

	back := widget.NewButton("返回", func() {
		// 恢复初始状态
		right.RefreshContent.Objects = []fyne.CanvasObject{
			widget.NewLabel("你点击了返回，当前是初始界面"),
		}
		right.RefreshContent.Refresh()
	})

	// 更新右边区域内容
	right.RefreshContent.Objects = []fyne.CanvasObject{btn1, btn2, back}
	right.RefreshContent.Refresh()
}
