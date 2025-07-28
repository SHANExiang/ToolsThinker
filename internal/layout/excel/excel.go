package excel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"tools-thinker/internal/excel/feature/merge"
	"tools-thinker/internal/layout/right"
)

// @Author dx
// @Date 2025-07-22 22:18:00
// @Desc

func Content(window fyne.Window) {
	// 定义常量以提高可维护性
	const (
		MergeButtonText      = "合并"
		Operation2ButtonText = "操作2"
		BackButtonText       = "返回"
		FolderDialogTitle    = "合并excel"
		LogEnterMerge        = "enter merge"
		LogStartProcessing   = "开始处理文件："
		NotificationTitle    = "操作2"
		NotificationContent  = "你点击了操作2！"
		ReturnToInitialState = "你点击了返回，当前是初始界面"
	)

	// 点击左侧按钮后，让右侧显示新的按钮区
	btn1 := widget.NewButton(MergeButtonText, func() {
		dialog.NewFolderOpen(
			func(uri fyne.ListableURI, err error) {
				// 增加空指针保护
				if right.PrintLog != nil {
					right.PrintLog(LogEnterMerge)
				}

				if err != nil {
					// 用户取消选择时也应有提示
					dialog.ShowInformation("提示", "未选择文件夹", fyne.CurrentApp().NewWindow("提示"))
					return
				}

				if uri != nil {
					if right.PrintLog != nil {
						right.PrintLog(LogStartProcessing + uri.Path())
					}
					err = merge.Handle(uri.Path())
					if err != nil {
						dialog.ShowError(err, fyne.CurrentApp().NewWindow("错误"))
					}
				} else {
					dialog.ShowInformation("提示", "未选择文件夹", fyne.CurrentApp().NewWindow("提示"))
				}
			},
			window,
		).Show()
	})

	btn2 := widget.NewButton(Operation2ButtonText, func() {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   NotificationTitle,
			Content: NotificationContent,
		})
	})

	back := widget.NewButton(BackButtonText, func() {
		// 恢复初始状态，并增加空指针保护
		if right.RefreshContent != nil {
			right.RefreshContent.Objects = []fyne.CanvasObject{
				widget.NewLabel(ReturnToInitialState),
			}
			right.RefreshContent.Refresh()
		}
	})

	// 更新右边区域内容并增加空指针保护
	if right.RefreshContent != nil {
		right.RefreshContent.Objects = []fyne.CanvasObject{btn1, btn2, back}
		right.RefreshContent.Refresh()
	}
}
