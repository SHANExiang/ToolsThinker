package myTheme

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed msyh.ttf
var msyhTTF []byte // 将微软雅黑字体嵌入到变量中

type MyTheme struct {
	fyne.Theme
	fontRes fyne.Resource
}

func (m *MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return m.fontRes
}

func NewMyTheme() *MyTheme {
	// 创建嵌入的字体资源
	yaheiFont := fyne.NewStaticResource("msyh.ttf", msyhTTF)

	// 设置主题
	return &MyTheme{
		Theme:   theme.LightTheme(),
		fontRes: yaheiFont,
	}
}
