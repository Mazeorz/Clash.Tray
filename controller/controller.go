package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
)

//type MyMainWindow struct {
//	*walk.MainWindow
//}

func ConfigMain() {
	const (
		SIZE_W = 650
		SIZE_H = 250
	)
	//w := new(MyMainWindow)
	var w *walk.MainWindow
	MainWindow{
		Visible:  false,
		AssignTo: &w,
		Name:     "ok",
		Title:    "配置管理 - Clash.Tray",
		Icon:     "./icon/Clash.Tray.png",
		Layout:   VBox{}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				ContextMenuItems: []MenuItem{
					Action{
						Text: "i&nfo",
					},
					Action{
						Text: "E&xit",
						OnTriggered: func() {
							w.Close()
						},
					},
				},
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
					},
					PushButton{
						Text: "启用配置",
					},
					PushButton{
						Text: "订阅转换",
						OnClicked: func() {
							open.Run("http://id9.cc")
						},
					},
					PushButton{
						Text: "打开目录",
					},
					PushButton{
						Text: "关闭窗口",
						OnClicked: func() {
							w.Close()
						},
					},
				},
			},
		},
	}.Create()
	//win.SetWindowLong(w.Handle(), win.GWL_STYLE, win.WS_BORDER)
	currStyle := win.GetWindowLong(w.Handle(), win.GWL_STYLE)
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX) //removes default styling
	hMenu := win.GetSystemMenu(w.Handle(), false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		w.Handle(),
		0,
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(w.Handle(), win.SW_SHOW)
	w.Run()
}
