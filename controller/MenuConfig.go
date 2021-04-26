package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
	"os"
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

type ConfigInfo struct {
	Index   int
	Name    string
	Byte    string
	Time    string
	Url     string
	checked bool
}

type ConfigInfoModel struct {
	walk.TableModelBase
	//walk.SorterBase
	sortColumn int
	//sortOrder  walk.SortOrder
	items []*ConfigInfo
}

func NewConfigInfoModel() *ConfigInfoModel {
	m := new(ConfigInfoModel)
	m.items = make([]*ConfigInfo, 3)
	m.items[0] = &ConfigInfo{Index: 0, Name: "Pre", Byte: "20KB", Time: "4/20 03:00", Url: "id9.com"}
	m.items[1] = &ConfigInfo{Index: 0, Name: "Pre2", Byte: "25KB", Time: "4/20 03:00", Url: "id9.com1"}
	m.items[2] = &ConfigInfo{Index: 0, Name: "Pre", Byte: "20KB", Time: "4/20 03:00", Url: "id9.com"}
	return m
}

func (m *ConfigInfoModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *ConfigInfoModel) RowCount() int {
	return len(m.items)
}

//func (m *ConfigInfoModel) Sort(col int,order walk.SortOrder) error{
//	m.sortColumn,m.sortOrder = col,order
//	sort.Stable(m)
//	return m.SorterBase.Sort(col,order)
//}

//func (m *ConfigInfoModel) Len() int {
//	return len(m.items)
//}

//func (m *ConfigInfoModel) Less(i,j int) bool {
//	a,b := m.items[i],m.items[j]
//	c:=func(ls bool)bool{
//		if m.sortOrder == walk.SortAscending{
//			return ls
//		}
//		return !ls
//	}
//	switch m.sortColumn {
//	case 0:
//		return c(a.Index < b.Index)
//	case 1:
//		return c(a.Name<b.Name)
//	case 2:
//		return c(a.Byte<b.Byte)
//	case 3:
//		return c(a.Time<b.Time)
//	case 4:
//		return c(a.Url<b.Url)
//	}
//	panic("unreachable")
//}

//func (m *ConfigInfoModel) Swap(i,j int) {
//	m.items[i],m.items[j] = m.items[j],m.items[i]
//}

func (m *ConfigInfoModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Name
	case 1:
		return item.Byte
	case 2:
		return item.Time
	case 3:
		return item.Url
	}
	panic("unexpected col")
}

func MenuConfig() {
	var model = NewConfigInfoModel()
	var tv *walk.TableView
	var MenuConfig *walk.MainWindow
	err := MainWindow{
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
						AssignTo:         &tv,
						CheckBoxes:       false,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []TableViewColumn{
							{Title: "配置名称"},
							{Title: "文件大小"},
							{Title: "更新日期"},
							{Title: "订阅地址", Width: 295},
						},
						Model: model,
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
							err := open.Run("https://id9.cc")
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: "打开目录",
						OnClicked: func() {
							exPath, _ := os.Getwd()
							err := open.Run(exPath + `/Profile`)
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: "关闭窗口",
						OnClicked: func() {
							err := MenuConfig.Close()
							if err != nil {
								return
							}
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		return
	}
	StyleMenuRun(MenuConfig, 650, 250)
}

func AddConfig() {
	var AddMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	err := MainWindow{
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
							err := AddMenuConfig.Close()
							if err != nil {
								return
							}
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		return
	}
	StyleMenuRun(AddMenuConfig, 420, 120)
}
