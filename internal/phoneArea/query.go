package phoneArea

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lxn/walk"
	wd "github.com/lxn/walk/declarative"
)

type AreaLine struct {
	Index    int
	PhoneNo  string
	Area     string
	Province string
	checked  bool
}

type AreaModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*AreaLine
}

type PhoneAddress struct {
	RetCode string           `json:"retCode"`
	RetMsg  string           `json:"retMsg"`
	Data    PhoneAddressData `json:"data"`
}

type PhoneAddressData struct {
	ProvCd   string `json:"prov_cd"`
	IdAreaCd string `json:"id_area_cd"`
	IdNameCd string `json:"id_name_cd"`
	NumType  string `json:"num_type"`
}

func (m *AreaModel) RowCount() int {
	return len(m.items)
}

func (m *AreaModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.PhoneNo
	case 2:
		return item.Area
	case 3:
		return item.Province
	}
	panic("unexpected col")
}

func (m *AreaModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *AreaModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

func (m *AreaModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *AreaModel) Len() int {
	return len(m.items)
}

func (m *AreaModel) Less(i, j int) bool {
	a, b := m.items[i], m.items[j]

	c := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.sortColumn {
	case 0:
		return c(a.Index < b.Index)
	case 1:
		return c(a.PhoneNo < b.PhoneNo)
	case 2:
		return c(a.Area < b.Area)
	case 3:
		return c(a.Province < b.Province)

	}

	panic("unreachable")
}

func (m *AreaModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func NewAreaModel() *AreaModel {
	m := new(AreaModel)
	m.items = make([]*AreaLine, 0)
	return m
}

type AreaMainWindow struct {
	*walk.MainWindow
	model     *AreaModel
	tableView *walk.TableView
}

var inTE *walk.TextEdit

func DemoMain() {
	mainWindow := &AreaMainWindow{model: NewAreaModel()}

	tableView := wd.TableView{
		AssignTo:         &mainWindow.tableView,
		CheckBoxes:       false,
		ColumnsOrderable: true,
		MultiSelection:   true,
		Columns: []wd.TableViewColumn{
			{Title: "编号"},
			{Title: "手机号"},
			{Title: "地区"},
			{Title: "省份"},
		},
		Model: mainWindow.model,
		OnCurrentIndexChanged: func() {
			i := mainWindow.tableView.CurrentIndex()
			if 0 <= i {
				log.Printf("OnCurrentIndexChanged: %v\n", mainWindow.model.items[i].PhoneNo)
			}
		},
		OnItemActivated: mainWindow.tableViewItemActivated,
	}

	splitter := wd.HSplitter{
		Children: []wd.Widget{
			wd.Composite{
				Layout: wd.HBox{MarginsZero: true},
				Children: []wd.Widget{
					wd.HSpacer{},
					wd.PushButton{
						Text:      "查归属地",
						OnClicked: mainWindow.query,
					},
					wd.TextEdit{AssignTo: &inTE},
				},
			},
			wd.Composite{
				Layout: wd.VBox{},
				ContextMenuItems: []wd.MenuItem{
					wd.Action{
						Text:        "I&nfo",
						OnTriggered: mainWindow.tableViewItemActivated,
					},
				},
				Children: []wd.Widget{
					tableView,
				},
			},
		},
	}
	window := wd.MainWindow{
		AssignTo: &mainWindow.MainWindow,
		Title:    "查移动手机号归属地",
		Size:     wd.Size{Width: 900, Height: 600},
		Layout:   wd.VBox{},
		Children: []wd.Widget{
			splitter,
		},
	}
	window.Run()
}

func (mw *AreaMainWindow) query() {
	phones := strings.Split(inTE.Text(), "\n")
	// 清空列表
	mw.model.items = nil
	for _, phone := range phones {
		phoneAddress := queryByPhone(phone)
		area := GetArea(phoneAddress.Data.IdAreaCd)
		mw.model.items = append(mw.model.items, &AreaLine{
			Index:    mw.model.Len() + 1,
			PhoneNo:  phone,
			Area:     phoneAddress.Data.IdNameCd,
			Province: area.Province,
		})
		mw.model.PublishRowsReset()
		mw.tableView.SetSelectedIndexes([]int{})

		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

func (mw *AreaMainWindow) tableViewItemActivated() {
	for _, i := range mw.tableView.SelectedIndexes() {
		it := mw.model.items[i]
		text := it.Province + " " + it.Area
		walk.Clipboard().SetText(text)
	}
}

/**
 * 根据电话号码查询 归属地
 */
func queryByPhone(phone string) *PhoneAddress {
	phoneAddress := &PhoneAddress{}
	log.Printf("text: %v\n", phone)
	timeStamp := int(time.Now().Unix())
	url := "https://shop.10086.cn/i/v1/res/numarea/" + phone + "?_=" + strconv.Itoa(timeStamp)
	log.Printf("url: %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("err: %v\n", err)
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err := json.Unmarshal(body, phoneAddress); err != nil {
		log.Printf("err: %v\n", err)
	}
	return phoneAddress
}
