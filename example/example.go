//original: http://doc.qt.io/qt-5/qtwidgets-widgets-lineedits-example.html

package main

import "github.com/therecipe/qt"

func main() {
	qt.NewQApplication(0, "")

	var (
		echoGroup    = qt.NewQGroupBox2("Echo", nil)
		echoLabel    = qt.NewQLabel2("Mode", nil, 0)
		echoComboBox = qt.NewQComboBox(nil)
		echoLineEdit = qt.NewQLineEdit1(nil)
	)
	echoComboBox.AddItems([]string{"Normal", "Password", "PasswordEchoOnEdit", "No Echo"})
	echoLineEdit.SetPlaceholderText("Placeholder Text")

	var (
		validatorGroup    = qt.NewQGroupBox2("Validator", nil)
		validatorLabel    = qt.NewQLabel2("Type:", nil, 0)
		validatorComboBox = qt.NewQComboBox(nil)
		validatorLineEdit = qt.NewQLineEdit1(nil)
	)
	validatorComboBox.AddItems([]string{"No validator", "Integer validator", "Double validator"})
	validatorLineEdit.SetPlaceholderText("Placeholder Text")
	validatorLineEdit.ConnectSlotClear()

	var (
		alignmentGroup    = qt.NewQGroupBox2("Alignment", nil)
		alignmentLabel    = qt.NewQLabel2("Type:", nil, 0)
		alignmentComboBox = qt.NewQComboBox(nil)
		alignmentLineEdit = qt.NewQLineEdit1(nil)
	)
	alignmentComboBox.AddItems([]string{"Left", "Centered", "Right"})
	alignmentLineEdit.SetPlaceholderText("Placeholder Text")

	var (
		inputMaskGroup    = qt.NewQGroupBox2("Input mask", nil)
		inputMaskLabel    = qt.NewQLabel2("Type:", nil, 0)
		inputMaskComboBox = qt.NewQComboBox(nil)
		inputMaskLineEdit = qt.NewQLineEdit1(nil)
	)
	inputMaskComboBox.AddItems([]string{"No mask", "Phone number", "ISO date", "License key"})
	inputMaskLineEdit.SetPlaceholderText("Placeholder Text")
	inputMaskLineEdit.ConnectSlotSetText()

	var (
		accessGroup    = qt.NewQGroupBox2("Access", nil)
		accessLabel    = qt.NewQLabel2("Read-only:", nil, 0)
		accessComboBox = qt.NewQComboBox(nil)
		accessLineEdit = qt.NewQLineEdit1(nil)
	)
	accessComboBox.AddItems([]string{"False", "True"})
	accessLineEdit.SetPlaceholderText("Placeholder Text")

	echoComboBox.ConnectSignalCurrentTextChanged(func() { echoChanged(echoLineEdit, echoComboBox.CurrentIndex()) })
	validatorComboBox.ConnectSignalCurrentTextChanged(func() { validatorChanged(validatorLineEdit, validatorComboBox.CurrentIndex()) })
	alignmentComboBox.ConnectSignalCurrentTextChanged(func() { alignmentChanged(alignmentLineEdit, alignmentComboBox.CurrentIndex()) })
	inputMaskComboBox.ConnectSignalCurrentTextChanged(func() { inputMaskChanged(inputMaskLineEdit, inputMaskComboBox.CurrentIndex()) })
	accessComboBox.ConnectSignalCurrentTextChanged(func() { accessChanged(accessLineEdit, accessComboBox.CurrentIndex()) })

	var echoLayout = qt.NewQGridLayout2()
	echoLayout.AddWidget1(echoLabel, 0, 0, 0)
	echoLayout.AddWidget1(echoComboBox, 0, 1, 0)
	echoLayout.AddWidget2(echoLineEdit, 1, 0, 1, 2, 0)
	echoGroup.SetLayout(echoLayout)

	var validatorLayout = qt.NewQGridLayout2()
	validatorLayout.AddWidget1(validatorLabel, 0, 0, 0)
	validatorLayout.AddWidget1(validatorComboBox, 0, 1, 0)
	validatorLayout.AddWidget2(validatorLineEdit, 1, 0, 1, 2, 0)
	validatorGroup.SetLayout(validatorLayout)

	var alignmentLayout = qt.NewQGridLayout2()
	alignmentLayout.AddWidget1(alignmentLabel, 0, 0, 0)
	alignmentLayout.AddWidget1(alignmentComboBox, 0, 1, 0)
	alignmentLayout.AddWidget2(alignmentLineEdit, 1, 0, 1, 2, 0)
	alignmentGroup.SetLayout(alignmentLayout)

	var inputMaskLayout = qt.NewQGridLayout2()
	inputMaskLayout.AddWidget1(inputMaskLabel, 0, 0, 0)
	inputMaskLayout.AddWidget1(inputMaskComboBox, 0, 1, 0)
	inputMaskLayout.AddWidget2(inputMaskLineEdit, 1, 0, 1, 2, 0)
	inputMaskGroup.SetLayout(inputMaskLayout)

	var accessLayout = qt.NewQGridLayout2()
	accessLayout.AddWidget1(accessLabel, 0, 0, 0)
	accessLayout.AddWidget1(accessComboBox, 0, 1, 0)
	accessLayout.AddWidget2(accessLineEdit, 1, 0, 1, 2, 0)
	accessGroup.SetLayout(accessLayout)

	var layout = qt.NewQGridLayout2()
	layout.AddWidget1(echoGroup, 0, 0, 0)
	layout.AddWidget1(validatorGroup, 1, 0, 0)
	layout.AddWidget1(alignmentGroup, 2, 0, 0)
	layout.AddWidget1(inputMaskGroup, 0, 1, 0)
	layout.AddWidget1(accessGroup, 1, 1, 0)

	var window = qt.NewQMainWindow(nil, 0)
	window.Layout().Destroy()
	window.SetLayout(layout)

	window.ConnectSlotSetWindowTitle()
	window.SlotSetWindowTitle("Line Edits")

	window.ConnectSlotShow()
	window.SlotShow()

	qt.QApplication_Exec()
}

func echoChanged(echoLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		echoLineEdit.SetEchoMode(qt.NORMAL)
	case 1:
		echoLineEdit.SetEchoMode(qt.PASSWORD)
	case 2:
		echoLineEdit.SetEchoMode(qt.PASSWORDECHOONEDIT)
	case 3:
		echoLineEdit.SetEchoMode(qt.NOECHO)
	}
}

func validatorChanged(validatorLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		validatorLineEdit.SetValidator(nil)
	case 1:
		validatorLineEdit.SetValidator(qt.NewQIntValidator1(nil))
	case 2:
		validatorLineEdit.SetValidator(qt.NewQDoubleValidator(nil))
	}

	go validatorLineEdit.SlotClear()
}

func alignmentChanged(alignmentLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		alignmentLineEdit.SetAlignment(qt.ALIGNLEFT)
	case 1:
		alignmentLineEdit.SetAlignment(qt.ALIGNCENTER)
	case 2:
		alignmentLineEdit.SetAlignment(qt.ALIGNRIGHT)
	}
}

func inputMaskChanged(inputMaskLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		inputMaskLineEdit.SetInputMask("")
	case 1:
		inputMaskLineEdit.SetInputMask("+99 99 99 99 99;_")
	case 2:
		inputMaskLineEdit.SetInputMask("0000-00-00")
		inputMaskLineEdit.SlotSetText("00000000")
		inputMaskLineEdit.SetCursorPosition(0)
	case 3:
		inputMaskLineEdit.SetInputMask(">AAAAA-AAAAA-AAAAA-AAAAA-AAAAA;#")
	}
}

func accessChanged(accessLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		accessLineEdit.SetReadOnly(false)
	case 1:
		accessLineEdit.SetReadOnly(true)
	}
}
