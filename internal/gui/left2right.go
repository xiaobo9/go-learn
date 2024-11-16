package gui

import (
	"strings"

	"github.com/lxn/walk"
	wd "github.com/lxn/walk/declarative"
)

func Left2Right() {
	var inTE, outTE *walk.TextEdit

	wd.MainWindow{
		Title:   "SCREAMO",
		MinSize: wd.Size{Width: 600, Height: 400},
		Layout:  wd.VBox{},
		Children: []wd.Widget{
			wd.HSplitter{
				Children: []wd.Widget{
					wd.TextEdit{AssignTo: &inTE},
					wd.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			wd.PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
