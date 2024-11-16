package gui

import (
	"log"

	wd "github.com/lxn/walk/declarative"
)

func NewDataModel() *DataModel {
	model := new(DataModel)
	model.items = make([]*Data, 3)

	model.items[0] = &Data{Index: 0, Name: "a", Price: 20}
	model.items[1] = &Data{Index: 1, Name: "b", Price: 18}
	model.items[2] = &Data{Index: 2, Name: "c", Price: 19}

	return model
}

func Window() {
	mw := &MyDataMainWindow{model: NewDataModel()}
	window := wd.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Data展示",
		Icon:     "favicon.ico",
		Size:     wd.Size{Width: 800, Height: 600},
		Layout:   wd.VBox{},
		Children: []wd.Widget{
			wd.Composite{
				Layout: wd.HBox{MarginsZero: true},
				Children: []wd.Widget{
					wd.HSpacer{},
					wd.PushButton{Text: "Add", OnClicked: func() { addBtn(mw) }},
					wd.PushButton{Text: "Delete", OnClicked: func() { deleteBtn(mw) }},
					wd.PushButton{Text: "ExecChecked", OnClicked: func() { execChecked(mw) }},
					wd.PushButton{Text: "AddPriceChecked", OnClicked: func() { addPriceChecked(mw) }},
				},
			},
			wd.Composite{
				Layout: wd.VBox{},
				ContextMenuItems: []wd.MenuItem{
					wd.Action{Text: "I&nfo", OnTriggered: mw.tvItemActivated},
					wd.Action{Text: "E&xit", OnTriggered: func() { mw.Close() }},
				},
				Children: []wd.Widget{
					wd.TableView{
						AssignTo:         &mw.tableView,
						CheckBoxes:       true,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []wd.TableViewColumn{
							{Title: "编号"},
							{Title: "名称"},
							{Title: "价格"},
						},
						Model:                 mw.model,
						OnCurrentIndexChanged: func() { currentIndexChanged(mw) },
						OnItemActivated:       mw.tvItemActivated,
					},
				},
			},
		},
	}
	Win2Center(mw.MainWindow, SIZE_WIDTH, SIZE_HEIGHT)
	_, e := window.Run()
	if e != nil {
		log.Println(e)
	}
}

func addBtn(mw *MyDataMainWindow) {
	mw.model.items = append(mw.model.items, &Data{
		Index: mw.model.Len() + 1,
		Name:  "啥名字",
		Price: mw.model.Len() * 5,
	})
	mw.model.PublishRowsReset()
	mw.tableView.SetSelectedIndexes([]int{})
}

func deleteBtn(mw *MyDataMainWindow) {
	items := []*Data{}
	remove := mw.tableView.SelectedIndexes()
	for i, item := range mw.model.items {
		removeOk := false
		for _, j := range remove {
			if i == j {
				removeOk = true
				break
			}
		}
		if !removeOk {
			items = append(items, item)
		}
	}
	mw.model.items = items
	mw.model.PublishRowsReset()
	mw.tableView.SetSelectedIndexes([]int{})
}

func execChecked(mw *MyDataMainWindow) {
	for _, item := range mw.model.items {
		if item.checked {
			log.Printf("checked: %v\n", item)
		}
	}
}

func addPriceChecked(mw *MyDataMainWindow) {
	for i, item := range mw.model.items {
		if item.checked {
			item.Price++
			mw.model.PublishRowChanged(i)
		}
	}
}

func currentIndexChanged(mw *MyDataMainWindow) {
	i := mw.tableView.CurrentIndex()
	if 0 <= i {
		log.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
	}
}
