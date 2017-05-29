package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func inputWidgets() {

	//Combo Box
	comboBox := widgets.NewQComboBox(nil)
	comboBox.SetWindowTitle("Combo Box")

	items := []string{"Some", "Combo", "Box", "Items"}
	comboBox.AddItems(items)

	comboBox.ConnectCurrentIndexChanged2(func(text string) {
		println("Combo Box Index Changed To:", text)
	})
	addWidget(comboBox)

	//Combo Box + List Model
	comboBoxL := widgets.NewQComboBox(nil)
	comboBoxL.SetWindowTitle("Combo Box + List Model")

	listModel := core.NewQAbstractListModel(nil)
	listModel.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(items)
	})
	listModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return core.NewQVariant14(items[index.Row()])
	})

	comboBoxL.SetModel(listModel)

	comboBoxL.ConnectCurrentIndexChanged2(func(text string) {
		println("Combo Box Index Changed To:", text)
	})
	addWidget(comboBoxL)

	//Combo Box + String List Model
	comboBoxSL := widgets.NewQComboBox(nil)
	comboBoxSL.SetWindowTitle("Combo Box + String List Model")
	comboBoxSL.SetModel(core.NewQStringListModel2(items, nil))
	comboBoxSL.ConnectCurrentIndexChanged2(func(text string) {
		println("Combo Box Index Changed To:", text)
	})
	addWidget(comboBoxSL)

	//Font Combo Box
	fontComboBox := widgets.NewQFontComboBox(nil)
	fontComboBox.SetWindowTitle("Font Combo Box")

	fontComboBox.ConnectCurrentFontChanged(func(font *gui.QFont) {
		println("Font Combo Box Index Changed To:", font.ToString())
	})
	addWidget(fontComboBox)

	//Line Edit
	lineEdit := widgets.NewQLineEdit(nil)
	lineEdit.SetWindowTitle("Line Edit")

	lineEdit.ConnectTextChanged(func(text string) {
		println("Line Edit Text Changed To:", text)
	})
	addWidget(lineEdit)

	//Text Edit
	textEdit := widgets.NewQTextEdit(nil)
	textEdit.SetWindowTitle("Text Edit")

	textEdit.ConnectTextChanged(func() {
		println("Text Edit Text Changed To:", textEdit.ToPlainText())
	})
	addWidget(textEdit)

	//Plain Text Edit
	plainTextEdit := widgets.NewQPlainTextEdit(nil)
	plainTextEdit.SetWindowTitle("Plain Text Edit")

	plainTextEdit.ConnectTextChanged(func() {
		println("Plain Text Edit Text Changed To:", plainTextEdit.ToPlainText())
	})
	addWidget(plainTextEdit)

	//Spin Box
	spinBox := widgets.NewQSpinBox(nil)
	spinBox.SetWindowTitle("Spin Box")

	spinBox.SetMinimum(0)
	spinBox.SetMaximum(1000)
	spinBox.SetValue(spinBox.Maximum() / 2)

	spinBox.ConnectValueChanged(func(i int) {
		println("Spin Box Value Changed To:", i)
	})
	addWidget(spinBox)

	//Double Spin Box
	doubleSpinBox := widgets.NewQDoubleSpinBox(nil)
	doubleSpinBox.SetWindowTitle("Double Spin Box")

	doubleSpinBox.SetMinimum(0)
	doubleSpinBox.SetMaximum(1000)
	doubleSpinBox.SetValue(doubleSpinBox.Maximum() / 2)

	doubleSpinBox.ConnectValueChanged(func(d float64) {
		println("Double Spin Box Value Changed To:", d)
	})
	addWidget(doubleSpinBox)

	//Time Edit
	timeEdit := widgets.NewQTimeEdit(nil)
	timeEdit.SetWindowTitle("Time Edit")

	timeEdit.SetTime(core.QTime_CurrentTime())

	timeEdit.ConnectTimeChanged(func(time *core.QTime) {
		println("Time Edit Time Changed To:", time.ToString2(core.Qt__TextDate))
	})
	addWidget(timeEdit)

	//Date Edit
	dateEdit := widgets.NewQDateEdit(nil)
	dateEdit.SetWindowTitle("Date Edit")

	dateEdit.SetDate(core.QDate_CurrentDate())

	dateEdit.ConnectDateChanged(func(date *core.QDate) {
		println("Date Edit Date Changed To:", date.ToString2(core.Qt__TextDate))
	})
	addWidget(dateEdit)

	//Date/Time Edit
	dateTimeEdit := widgets.NewQDateTimeEdit(nil)
	dateTimeEdit.SetWindowTitle("Date Time Edit")

	dateTimeEdit.SetDate(core.QDate_CurrentDate())
	dateTimeEdit.SetTime(core.QTime_CurrentTime())

	dateTimeEdit.ConnectDateTimeChanged(func(datetime *core.QDateTime) {
		println("Date Time Edit Changed To:", datetime.ToString2(core.Qt__TextDate))
	})
	addWidget(dateTimeEdit)

	//Dial
	dial := widgets.NewQDial(nil)
	dial.SetWindowTitle("Dial")

	dial.SetMinimum(0)
	dial.SetMaximum(1000)
	dial.SetValue(dial.Maximum() / 2)

	dial.ConnectValueChanged(func(value int) {
		println("Dial Value Changed To:", value)
	})
	addWidget(dial)

	//Horizontal Scroll Bar
	hScrollBar := widgets.NewQScrollBar2(core.Qt__Horizontal, nil)
	hScrollBar.SetWindowTitle("Horizontal Scroll Bar")

	hScrollBar.SetMinimum(0)
	hScrollBar.SetMaximum(1000)
	hScrollBar.SetValue(hScrollBar.Maximum() / 2)

	hScrollBar.ConnectValueChanged(func(value int) {
		println("Horizontal Scroll Bar Value Changed To:", value)
	})
	addWidget(hScrollBar)

	//Vertical Scroll Bar
	vScrollBar := widgets.NewQScrollBar2(core.Qt__Vertical, nil)
	vScrollBar.SetWindowTitle("Vertical Scroll Bar")

	vScrollBar.SetMinimum(0)
	vScrollBar.SetMaximum(1000)
	vScrollBar.SetValue(vScrollBar.Maximum() / 2)

	vScrollBar.ConnectValueChanged(func(value int) {
		println("Vertical Scroll Bar Value Changed To:", value)
	})
	addWidget(vScrollBar)

	//Horizontal Slider
	hSlider := widgets.NewQSlider2(core.Qt__Horizontal, nil)
	hSlider.SetWindowTitle("Horizontal Slider")

	hSlider.SetMinimum(0)
	hSlider.SetMaximum(1000)
	hSlider.SetValue(hSlider.Maximum() / 2)

	hSlider.ConnectValueChanged(func(value int) {
		println("Horizontal Slider Value Changed To:", value)
	})
	addWidget(hSlider)

	//Vertical Slider
	vSlider := widgets.NewQSlider2(core.Qt__Vertical, nil)
	vSlider.SetWindowTitle("Vertical Slider")

	vSlider.SetMinimum(0)
	vSlider.SetMaximum(1000)
	vSlider.SetValue(vSlider.Maximum() / 2)

	vSlider.ConnectValueChanged(func(value int) {
		println("Vertical Slider Value Changed To:", value)
	})
	addWidget(vSlider)

	//Key Sequence Edit
	keySequenceEdit := widgets.NewQKeySequenceEdit(nil)
	keySequenceEdit.SetWindowTitle("Key Sequence Edit")

	keySequenceEdit.ConnectKeySequenceChanged(func(keySequence *gui.QKeySequence) {
		println("Key Sequence Edit Key Sequence Changed To:", keySequence.ToString(gui.QKeySequence__NativeText))
	})
	addWidget(keySequenceEdit)

}
