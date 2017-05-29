package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func layouts() {

	//Vertical Layout
	vboxWidget := widgets.NewQWidget(nil, 0)
	vboxWidget.SetWindowTitle("Vertical Layout")

	vbox := widgets.NewQVBoxLayout2(vboxWidget)
	vbox.AddWidget(widgets.NewQLabel2("one", nil, 0), 0, 0)
	vbox.AddWidget(widgets.NewQLabel2("two", nil, 0), 0, 0)
	vbox.AddWidget(widgets.NewQLabel2("three", nil, 0), 0, 0)
	addWidget(vboxWidget)

	//Horizontal Layout
	hboxWidget := widgets.NewQWidget(nil, 0)
	hboxWidget.SetWindowTitle("Horizontal Layout")

	hbox := widgets.NewQHBoxLayout2(hboxWidget)
	hbox.AddWidget(widgets.NewQLabel2("one", nil, 0), 0, 0)
	hbox.AddWidget(widgets.NewQLabel2("two", nil, 0), 0, 0)
	hbox.AddWidget(widgets.NewQLabel2("three", nil, 0), 0, 0)
	addWidget(hboxWidget)

	//Grid Layout
	gridWidget := widgets.NewQWidget(nil, 0)
	gridWidget.SetWindowTitle("Grid Layout")

	grid := widgets.NewQGridLayout(gridWidget)
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			grid.AddWidget(widgets.NewQLabel2(fmt.Sprintf("[%v:%v]", row, column), nil, 0), row, column, core.Qt__AlignCenter)
		}
	}
	addWidget(gridWidget)

	//Form layout
	formWidget := widgets.NewQWidget(nil, 0)
	formWidget.SetWindowTitle("Form Layout")

	form := widgets.NewQFormLayout(formWidget)
	for row := 0; row < 3; row++ {
		form.AddRow3(fmt.Sprintf("label %v", row), widgets.NewQLineEdit2(fmt.Sprintf("field %v", row), nil))
	}
	addWidget(formWidget)

}
