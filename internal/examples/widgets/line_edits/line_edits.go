//source: https://doc.qt.io/qt-5/qtwidgets-widgets-lineedits-example.html

package main

import (
	"os"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		echoGroup    = widgets.NewQGroupBox2("Echo", nil)
		echoLabel    = widgets.NewQLabel2("Mode:", nil, 0)
		echoComboBox = widgets.NewQComboBox(nil)
		echoLineEdit = widgets.NewQLineEdit(nil)
	)
	echoComboBox.AddItems([]string{"Normal", "Password", "PasswordEchoOnEdit", "No Echo"})
	echoLineEdit.SetPlaceholderText("Placeholder Text")

	var (
		validatorGroup    = widgets.NewQGroupBox2("Validator", nil)
		validatorLabel    = widgets.NewQLabel2("Type:", nil, 0)
		validatorComboBox = widgets.NewQComboBox(nil)
		validatorLineEdit = widgets.NewQLineEdit(nil)
	)
	validatorComboBox.AddItems([]string{"No validator", "Integer validator", "Double validator"})
	validatorLineEdit.SetPlaceholderText("Placeholder Text")

	var (
		alignmentGroup    = widgets.NewQGroupBox2("Alignment", nil)
		alignmentLabel    = widgets.NewQLabel2("Type:", nil, 0)
		alignmentComboBox = widgets.NewQComboBox(nil)
		alignmentLineEdit = widgets.NewQLineEdit(nil)
	)
	alignmentComboBox.AddItems([]string{"Left", "Centered", "Right"})
	alignmentLineEdit.SetPlaceholderText("Placeholder Text")

	var (
		inputMaskGroup    = widgets.NewQGroupBox2("Input mask", nil)
		inputMaskLabel    = widgets.NewQLabel2("Type:", nil, 0)
		inputMaskComboBox = widgets.NewQComboBox(nil)
		inputMaskLineEdit = widgets.NewQLineEdit(nil)
	)
	inputMaskComboBox.AddItems([]string{"No mask", "Phone number", "ISO date", "License key"})
	inputMaskLineEdit.SetPlaceholderText("Placeholder Text")

	var (
		accessGroup    = widgets.NewQGroupBox2("Access", nil)
		accessLabel    = widgets.NewQLabel2("Read-only:", nil, 0)
		accessComboBox = widgets.NewQComboBox(nil)
		accessLineEdit = widgets.NewQLineEdit(nil)
	)
	accessComboBox.AddItems([]string{"False", "True"})
	accessLineEdit.SetPlaceholderText("Placeholder Text")

	echoComboBox.ConnectCurrentIndexChanged(func(index int) { echoChanged(echoLineEdit, index) })
	validatorComboBox.ConnectCurrentIndexChanged(func(index int) { validatorChanged(validatorLineEdit, index) })
	alignmentComboBox.ConnectCurrentIndexChanged(func(index int) { alignmentChanged(alignmentLineEdit, index) })
	inputMaskComboBox.ConnectCurrentIndexChanged(func(index int) { inputMaskChanged(inputMaskLineEdit, index) })
	accessComboBox.ConnectCurrentIndexChanged(func(index int) { accessChanged(accessLineEdit, index) })

	var echoLayout = widgets.NewQGridLayout2()
	echoLayout.AddWidget2(echoLabel, 0, 0, 0)
	echoLayout.AddWidget2(echoComboBox, 0, 1, 0)
	echoLayout.AddWidget3(echoLineEdit, 1, 0, 1, 2, 0)
	echoGroup.SetLayout(echoLayout)

	var validatorLayout = widgets.NewQGridLayout2()
	validatorLayout.AddWidget2(validatorLabel, 0, 0, 0)
	validatorLayout.AddWidget2(validatorComboBox, 0, 1, 0)
	validatorLayout.AddWidget3(validatorLineEdit, 1, 0, 1, 2, 0)
	validatorGroup.SetLayout(validatorLayout)

	var alignmentLayout = widgets.NewQGridLayout2()
	alignmentLayout.AddWidget2(alignmentLabel, 0, 0, 0)
	alignmentLayout.AddWidget2(alignmentComboBox, 0, 1, 0)
	alignmentLayout.AddWidget3(alignmentLineEdit, 1, 0, 1, 2, 0)
	alignmentGroup.SetLayout(alignmentLayout)

	var inputMaskLayout = widgets.NewQGridLayout2()
	inputMaskLayout.AddWidget2(inputMaskLabel, 0, 0, 0)
	inputMaskLayout.AddWidget2(inputMaskComboBox, 0, 1, 0)
	inputMaskLayout.AddWidget3(inputMaskLineEdit, 1, 0, 1, 2, 0)
	inputMaskGroup.SetLayout(inputMaskLayout)

	var accessLayout = widgets.NewQGridLayout2()
	accessLayout.AddWidget2(accessLabel, 0, 0, 0)
	accessLayout.AddWidget2(accessComboBox, 0, 1, 0)
	accessLayout.AddWidget3(accessLineEdit, 1, 0, 1, 2, 0)
	accessGroup.SetLayout(accessLayout)

	var layout = widgets.NewQGridLayout2()
	layout.AddWidget2(echoGroup, 0, 0, 0)
	layout.AddWidget2(validatorGroup, 1, 0, 0)
	layout.AddWidget2(alignmentGroup, 2, 0, 0)
	layout.AddWidget2(inputMaskGroup, 0, 1, 0)
	layout.AddWidget2(accessGroup, 1, 1, 0)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Line Edits")

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}

func echoChanged(echoLineEdit *widgets.QLineEdit, index int) {
	switch index {
	case 0:
		{
			echoLineEdit.SetEchoMode(widgets.QLineEdit__Normal)
		}

	case 1:
		{
			echoLineEdit.SetEchoMode(widgets.QLineEdit__Password)
		}

	case 2:
		{
			echoLineEdit.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit)
		}

	case 3:
		{
			echoLineEdit.SetEchoMode(widgets.QLineEdit__NoEcho)
		}
	}
}

func validatorChanged(validatorLineEdit *widgets.QLineEdit, index int) {
	switch index {
	case 0:
		{
			validatorLineEdit.SetValidator(nil)
		}

	case 1:
		{
			validatorLineEdit.SetValidator(gui.NewQIntValidator(validatorLineEdit))
		}

	case 2:
		{
			validatorLineEdit.SetValidator(gui.NewQDoubleValidator2(-999.0, 999.0, 2, validatorLineEdit))
		}
	}

	validatorLineEdit.Clear()
}

func alignmentChanged(alignmentLineEdit *widgets.QLineEdit, index int) {
	switch index {
	case 0:
		{
			alignmentLineEdit.SetAlignment(core.Qt__AlignLeft)
		}

	case 1:
		{
			alignmentLineEdit.SetAlignment(core.Qt__AlignCenter)
		}

	case 2:
		{
			alignmentLineEdit.SetAlignment(core.Qt__AlignRight)
		}
	}
}

func inputMaskChanged(inputMaskLineEdit *widgets.QLineEdit, index int) {
	switch index {
	case 0:
		{
			inputMaskLineEdit.SetInputMask("")
		}

	case 1:
		{
			inputMaskLineEdit.SetInputMask("+99 99 99 99 99;_")
		}

	case 2:
		{
			inputMaskLineEdit.SetInputMask("0000-00-00")
			inputMaskLineEdit.SetText("00000000")
			inputMaskLineEdit.SetCursorPosition(0)
		}

	case 3:
		{
			inputMaskLineEdit.SetInputMask(">AAAAA-AAAAA-AAAAA-AAAAA-AAAAA;#")
		}
	}
}

func accessChanged(accessLineEdit *widgets.QLineEdit, index int) {
	switch index {
	case 0:
		{
			accessLineEdit.SetReadOnly(false)
		}

	case 1:
		{
			accessLineEdit.SetReadOnly(true)
		}
	}
}
