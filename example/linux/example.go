//original: http://qt-project.org/doc/qt-5/qtwidgets-widgets-lineedits-example.html

package main

import "github.com/therecipe/qt"

func main() {
	qt.NewQApplication_Int_String(0, "")

	var (
		echoGroup    = qt.NewQGroupBox_String_QWidget("Echo", nil)
		echoLabel    = qt.NewQLabel_String_QWidget_WindowType("Mode", nil, 0)
		echoComboBox = qt.NewQComboBox_QWidget(nil)
		echoLineEdit = qt.NewQLineEdit_QWidget(nil)
	)
	echoComboBox.AddItems_QStringList([]string{"Normal", "Password", "PasswordEchoOnEdit", "No Echo"})
	echoLineEdit.SetPlaceholderText_String("Placeholder Text")

	var (
		validatorGroup    = qt.NewQGroupBox_String_QWidget("Validator", nil)
		validatorLabel    = qt.NewQLabel_String_QWidget_WindowType("Type:", nil, 0)
		validatorComboBox = qt.NewQComboBox_QWidget(nil)
		validatorLineEdit = qt.NewQLineEdit_QWidget(nil)
	)
	validatorComboBox.AddItems_QStringList([]string{"No validator", "Integer validator", "Double validator"})
	validatorLineEdit.SetPlaceholderText_String("Placeholder Text")
	validatorLineEdit.ConnectSlotClear()

	var (
		alignmentGroup    = qt.NewQGroupBox_String_QWidget("Alignment", nil)
		alignmentLabel    = qt.NewQLabel_String_QWidget_WindowType("Type:", nil, 0)
		alignmentComboBox = qt.NewQComboBox_QWidget(nil)
		alignmentLineEdit = qt.NewQLineEdit_QWidget(nil)
	)
	alignmentComboBox.AddItems_QStringList([]string{"Left", "Centered", "Right"})
	alignmentLineEdit.SetPlaceholderText_String("Placeholder Text")

	var (
		inputMaskGroup    = qt.NewQGroupBox_String_QWidget("Input mask", nil)
		inputMaskLabel    = qt.NewQLabel_String_QWidget_WindowType("Type:", nil, 0)
		inputMaskComboBox = qt.NewQComboBox_QWidget(nil)
		inputMaskLineEdit = qt.NewQLineEdit_QWidget(nil)
	)
	inputMaskComboBox.AddItems_QStringList([]string{"No mask", "Phone number", "ISO date", "License key"})
	inputMaskLineEdit.SetPlaceholderText_String("Placeholder Text")
	inputMaskLineEdit.ConnectSlotSetText()

	var (
		accessGroup    = qt.NewQGroupBox_String_QWidget("Access", nil)
		accessLabel    = qt.NewQLabel_String_QWidget_WindowType("Read-only:", nil, 0)
		accessComboBox = qt.NewQComboBox_QWidget(nil)
		accessLineEdit = qt.NewQLineEdit_QWidget(nil)
	)
	accessComboBox.AddItems_QStringList([]string{"False", "True"})
	accessLineEdit.SetPlaceholderText_String("Placeholder Text")

	echoComboBox.ConnectSignalCurrentTextChanged(func() { echoChanged(echoLineEdit, echoComboBox.CurrentIndex()) })
	validatorComboBox.ConnectSignalCurrentTextChanged(func() { validatorChanged(validatorLineEdit, validatorComboBox.CurrentIndex()) })
	alignmentComboBox.ConnectSignalCurrentTextChanged(func() { alignmentChanged(alignmentLineEdit, alignmentComboBox.CurrentIndex()) })
	inputMaskComboBox.ConnectSignalCurrentTextChanged(func() { inputMaskChanged(inputMaskLineEdit, inputMaskComboBox.CurrentIndex()) })
	accessComboBox.ConnectSignalCurrentTextChanged(func() { accessChanged(accessLineEdit, accessComboBox.CurrentIndex()) })

	var echoLayout = qt.NewQGridLayout()
	echoLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(echoLabel, 0, 0, 0)
	echoLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(echoComboBox, 0, 1, 0)
	echoLayout.AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(echoLineEdit, 1, 0, 1, 2, 0)
	echoGroup.SetLayout_QLayout(echoLayout)

	var validatorLayout = qt.NewQGridLayout()
	validatorLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(validatorLabel, 0, 0, 0)
	validatorLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(validatorComboBox, 0, 1, 0)
	validatorLayout.AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(validatorLineEdit, 1, 0, 1, 2, 0)
	validatorGroup.SetLayout_QLayout(validatorLayout)

	var alignmentLayout = qt.NewQGridLayout()
	alignmentLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(alignmentLabel, 0, 0, 0)
	alignmentLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(alignmentComboBox, 0, 1, 0)
	alignmentLayout.AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(alignmentLineEdit, 1, 0, 1, 2, 0)
	alignmentGroup.SetLayout_QLayout(alignmentLayout)

	var inputMaskLayout = qt.NewQGridLayout()
	inputMaskLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(inputMaskLabel, 0, 0, 0)
	inputMaskLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(inputMaskComboBox, 0, 1, 0)
	inputMaskLayout.AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(inputMaskLineEdit, 1, 0, 1, 2, 0)
	inputMaskGroup.SetLayout_QLayout(inputMaskLayout)

	var accessLayout = qt.NewQGridLayout()
	accessLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(accessLabel, 0, 0, 0)
	accessLayout.AddWidget_QWidget_Int_Int_AlignmentFlag(accessComboBox, 0, 1, 0)
	accessLayout.AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(accessLineEdit, 1, 0, 1, 2, 0)
	accessGroup.SetLayout_QLayout(accessLayout)

	var layout = qt.NewQGridLayout()
	layout.AddWidget_QWidget_Int_Int_AlignmentFlag(echoGroup, 0, 0, 0)
	layout.AddWidget_QWidget_Int_Int_AlignmentFlag(validatorGroup, 1, 0, 0)
	layout.AddWidget_QWidget_Int_Int_AlignmentFlag(alignmentGroup, 2, 0, 0)
	layout.AddWidget_QWidget_Int_Int_AlignmentFlag(inputMaskGroup, 0, 1, 0)
	layout.AddWidget_QWidget_Int_Int_AlignmentFlag(accessGroup, 1, 1, 0)

	var window = qt.NewQMainWindow_QWidget_WindowType(nil, 0)
	window.Layout().Destroy()
	window.SetLayout_QLayout(layout)

	window.ConnectSlotSetWindowTitle()
	window.SlotSetWindowTitle_String("Line Edits")

	window.ConnectSlotShow()
	window.SlotShow()

	qt.QApplication_Exec()
}

func echoChanged(echoLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		echoLineEdit.SetEchoMode_EchoMode(qt.NORMAL)
	case 1:
		echoLineEdit.SetEchoMode_EchoMode(qt.PASSWORD)
	case 2:
		echoLineEdit.SetEchoMode_EchoMode(qt.PASSWORDECHOONEDIT)
	case 3:
		echoLineEdit.SetEchoMode_EchoMode(qt.NOECHO)
	}
}

func validatorChanged(validatorLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		validatorLineEdit.SetValidator_QValidator(nil)
	case 1:
		validatorLineEdit.SetValidator_QValidator(qt.NewQIntValidator_QObject(nil))
	case 2:
		validatorLineEdit.SetValidator_QValidator(qt.NewQDoubleValidator_QObject(nil))
	}

	go validatorLineEdit.SlotClear()
}

func alignmentChanged(alignmentLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		alignmentLineEdit.SetAlignment_AlignmentFlag(qt.ALIGNLEFT)
	case 1:
		alignmentLineEdit.SetAlignment_AlignmentFlag(qt.ALIGNCENTER)
	case 2:
		alignmentLineEdit.SetAlignment_AlignmentFlag(qt.ALIGNRIGHT)
	}
}

func inputMaskChanged(inputMaskLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		inputMaskLineEdit.SetInputMask_String("")
	case 1:
		inputMaskLineEdit.SetInputMask_String("+99 99 99 99 99;_")
	case 2:
		inputMaskLineEdit.SetInputMask_String("0000-00-00")
		inputMaskLineEdit.SlotSetText_String("00000000")
		inputMaskLineEdit.SetCursorPosition_Int(0)
	case 3:
		inputMaskLineEdit.SetInputMask_String(">AAAAA-AAAAA-AAAAA-AAAAA-AAAAA;#")
	}
}

func accessChanged(accessLineEdit qt.QLineEdit, index int) {
	switch index {
	case 0:
		accessLineEdit.SetReadOnly_Bool(false)
	case 1:
		accessLineEdit.SetReadOnly_Bool(true)
	}
}
