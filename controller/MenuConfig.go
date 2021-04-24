package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
)

func StyleMenuRun(w *walk.MainWindow, SizeW int32, SizeH int32) {
	currStyle := win.GetWindowLong(w.Handle(), win.GWL_STYLE)
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX) //removes default styling
	hMenu := win.GetSystemMenu(w.Handle(), false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(w.Handle(), 0, (xScreen-SizeW)/2, (yScreen-SizeH)/2, SizeW, SizeH, win.SWP_FRAMECHANGED)
	win.ShowWindow(w.Handle(), win.SW_SHOW)
	w.Run()

}

func MenuConfig() {
	var MenuConfig *walk.MainWindow
	MainWindow{
		Visible:  false,
		AssignTo: &MenuConfig,
		Name:     "ok",
		Title:    "配置管理 - Clash.Tray",
		Icon:     "./icon/Clash.Tray.png",
		Layout:   VBox{}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				Children: []Widget{
					TableView{
						CheckBoxes:       false,
						ColumnsOrderable: false,
						MultiSelection:   false,
						Columns: []TableViewColumn{
							{Title: "配置名称"},
							{Title: "文件大小"},
							{Title: "更新日期"},
							{Title: "订阅地址", Width: 295},
						},
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "添加配置",
						OnClicked: func() {
							AddConfig()
						},
					},
					PushButton{
						Text: "启用配置",
					},
					PushButton{
						Text: "订阅转换",
						OnClicked: func() {
							open.Run("https://id9.cc")
						},
					},
					PushButton{
						Text: "打开目录",
					},
					PushButton{
						Text: "关闭窗口",
						OnClicked: func() {
							MenuConfig.Close()
						},
					},
				},
			},
		},
	}.Create()
	StyleMenuRun(MenuConfig, 650, 250)
}

func AddConfig() {
	var AddMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	MainWindow{
		Visible:  true,
		AssignTo: &AddMenuConfig,
		Title:    "添加配置 - Clash.Tray",
		Icon:     "./icon/Clash.Tray.png",
		Layout:   VBox{}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "请输入订阅链接:",
					},
					LineEdit{
						AssignTo: &oUrl,
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "添加",
						OnClicked: func() {
							//待添加
						},
					},
					PushButton{
						Text: "取消",
						OnClicked: func() {
							AddMenuConfig.Close()
						},
					},
				},
			},
		},
	}.Create()
	StyleMenuRun(AddMenuConfig, 420, 120)
}
