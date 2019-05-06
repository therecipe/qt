package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	NewCalculatorForm().Show()

	widgets.QApplication_Exec()
}

func NewCalculatorForm() *widgets.QWidget {

	file := core.NewQFile2(":/qml/calculatorform.ui")
	file.Open(core.QIODevice__ReadOnly)
	formWidget := uitools.NewQUiLoader(nil).Load(file, nil)
	file.Close()

	formWidget.SetWindowTitle("Calculator Builder")

	var (
		ui_inputSpinBox1 = widgets.NewQSpinBoxFromPointer(formWidget.FindChild("inputSpinBox1", core.Qt__FindChildrenRecursively).Pointer())
		ui_inputSpinBox2 = widgets.NewQSpinBoxFromPointer(formWidget.FindChild("inputSpinBox2", core.Qt__FindChildrenRecursively).Pointer())
		ui_outputWidget  = widgets.NewQLabelFromPointer(formWidget.FindChild("outputWidget", core.Qt__FindChildrenRecursively).Pointer())
	)

	ui_inputSpinBox1.ConnectValueChanged(func(value int) {
		ui_outputWidget.SetText(fmt.Sprint(value + ui_inputSpinBox2.Value()))
	})

	ui_inputSpinBox2.ConnectValueChanged(func(value int) {
		ui_outputWidget.SetText(fmt.Sprint(value + ui_inputSpinBox1.Value()))
	})

	return formWidget
}
