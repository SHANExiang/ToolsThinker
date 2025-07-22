package main

import (
	"tools-thinker/internal"
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

func init() {
	internal.InitApp() // 主窗口初始化
}

func main() {
	internal.Window.ShowAndRun()
}
