package gui

import (
	"log"
	"sort"

	"github.com/lxn/walk"
)

type Data struct {
	Index   int
	Name    string
	Price   int
	checked bool
}

type DataModel struct {
	walk.TableModelBase
	walk.SorterBase
	items []*Data
}

type MyDataMainWindow struct {
	*walk.MainWindow
	model     *DataModel
	tableView *walk.TableView
}

func (m *DataModel) RowCount() int {
	return len(m.items)
}

func (m *DataModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.Name
	case 2:
		return item.Price
	}
	panic("unexpected col")
}

func (m *DataModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *DataModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

func (m *DataModel) Sort(col int, order walk.SortOrder) error {
	sort.Stable(m)
	return m.SorterBase.Sort(col, order)
}

func (m *DataModel) Len() int {
	return len(m.items)
}

func (m *DataModel) Less(i, j int) bool {
	log.Println("Less")
	a, b := m.items[i], m.items[j]

	c := func(ls bool) bool {
		if m.SortOrder() == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.SortedColumn() {
	case 0:
		return c(a.Index < b.Index)
	case 1:
		return c(a.Name < b.Name)
	case 2:
		return c(a.Price < b.Price)
	}

	panic("unreachable")
}

func (m *DataModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func (mw *MyDataMainWindow) tvItemActivated() {
	msg := ``
	for _, i := range mw.tableView.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].Name
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}
