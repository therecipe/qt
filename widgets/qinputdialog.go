package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QInputDialog struct {
	QDialog
}

type QInputDialog_ITF interface {
	QDialog_ITF
	QInputDialog_PTR() *QInputDialog
}

func PointerFromQInputDialog(ptr QInputDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputDialog_PTR().Pointer()
	}
	return nil
}

func NewQInputDialogFromPointer(ptr unsafe.Pointer) *QInputDialog {
	var n = new(QInputDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QInputDialog_") {
		n.SetObjectName("QInputDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QInputDialog) QInputDialog_PTR() *QInputDialog {
	return ptr
}

//QInputDialog::InputDialogOption
type QInputDialog__InputDialogOption int64

const (
	QInputDialog__NoButtons                    = QInputDialog__InputDialogOption(0x00000001)
	QInputDialog__UseListViewForComboBoxItems  = QInputDialog__InputDialogOption(0x00000002)
	QInputDialog__UsePlainTextEditForTextInput = QInputDialog__InputDialogOption(0x00000004)
)

//QInputDialog::InputMode
type QInputDialog__InputMode int64

const (
	QInputDialog__TextInput   = QInputDialog__InputMode(0)
	QInputDialog__IntInput    = QInputDialog__InputMode(1)
	QInputDialog__DoubleInput = QInputDialog__InputMode(2)
)

func (ptr *QInputDialog) CancelButtonText() string {
	defer qt.Recovering("QInputDialog::cancelButtonText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_CancelButtonText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInputDialog) ComboBoxItems() []string {
	defer qt.Recovering("QInputDialog::comboBoxItems")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QInputDialog_ComboBoxItems(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QInputDialog) DoubleDecimals() int {
	defer qt.Recovering("QInputDialog::doubleDecimals")

	if ptr.Pointer() != nil {
		return int(C.QInputDialog_DoubleDecimals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) InputMode() QInputDialog__InputMode {
	defer qt.Recovering("QInputDialog::inputMode")

	if ptr.Pointer() != nil {
		return QInputDialog__InputMode(C.QInputDialog_InputMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) IntMaximum() int {
	defer qt.Recovering("QInputDialog::intMaximum")

	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntMaximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) IntMinimum() int {
	defer qt.Recovering("QInputDialog::intMinimum")

	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntMinimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) IntStep() int {
	defer qt.Recovering("QInputDialog::intStep")

	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) IntValue() int {
	defer qt.Recovering("QInputDialog::intValue")

	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) IsComboBoxEditable() bool {
	defer qt.Recovering("QInputDialog::isComboBoxEditable")

	if ptr.Pointer() != nil {
		return C.QInputDialog_IsComboBoxEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QInputDialog) LabelText() string {
	defer qt.Recovering("QInputDialog::labelText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_LabelText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInputDialog) OkButtonText() string {
	defer qt.Recovering("QInputDialog::okButtonText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_OkButtonText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInputDialog) Options() QInputDialog__InputDialogOption {
	defer qt.Recovering("QInputDialog::options")

	if ptr.Pointer() != nil {
		return QInputDialog__InputDialogOption(C.QInputDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) SetCancelButtonText(text string) {
	defer qt.Recovering("QInputDialog::setCancelButtonText")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetCancelButtonText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QInputDialog) SetComboBoxEditable(editable bool) {
	defer qt.Recovering("QInputDialog::setComboBoxEditable")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetComboBoxEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QInputDialog) SetComboBoxItems(items []string) {
	defer qt.Recovering("QInputDialog::setComboBoxItems")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetComboBoxItems(ptr.Pointer(), C.CString(strings.Join(items, ",,,")))
	}
}

func (ptr *QInputDialog) SetDoubleDecimals(decimals int) {
	defer qt.Recovering("QInputDialog::setDoubleDecimals")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetDoubleDecimals(ptr.Pointer(), C.int(decimals))
	}
}

func (ptr *QInputDialog) SetInputMode(mode QInputDialog__InputMode) {
	defer qt.Recovering("QInputDialog::setInputMode")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetInputMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QInputDialog) SetIntMaximum(max int) {
	defer qt.Recovering("QInputDialog::setIntMaximum")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntMaximum(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QInputDialog) SetIntMinimum(min int) {
	defer qt.Recovering("QInputDialog::setIntMinimum")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntMinimum(ptr.Pointer(), C.int(min))
	}
}

func (ptr *QInputDialog) SetIntStep(step int) {
	defer qt.Recovering("QInputDialog::setIntStep")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntStep(ptr.Pointer(), C.int(step))
	}
}

func (ptr *QInputDialog) SetIntValue(value int) {
	defer qt.Recovering("QInputDialog::setIntValue")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntValue(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QInputDialog) SetLabelText(text string) {
	defer qt.Recovering("QInputDialog::setLabelText")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetLabelText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QInputDialog) SetOkButtonText(text string) {
	defer qt.Recovering("QInputDialog::setOkButtonText")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetOkButtonText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QInputDialog) SetOptions(options QInputDialog__InputDialogOption) {
	defer qt.Recovering("QInputDialog::setOptions")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QInputDialog) SetTextEchoMode(mode QLineEdit__EchoMode) {
	defer qt.Recovering("QInputDialog::setTextEchoMode")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetTextEchoMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QInputDialog) SetTextValue(text string) {
	defer qt.Recovering("QInputDialog::setTextValue")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetTextValue(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QInputDialog) TextEchoMode() QLineEdit__EchoMode {
	defer qt.Recovering("QInputDialog::textEchoMode")

	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QInputDialog_TextEchoMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputDialog) TextValue() string {
	defer qt.Recovering("QInputDialog::textValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_TextValue(ptr.Pointer()))
	}
	return ""
}

func NewQInputDialog(parent QWidget_ITF, flags core.Qt__WindowType) *QInputDialog {
	defer qt.Recovering("QInputDialog::QInputDialog")

	return NewQInputDialogFromPointer(C.QInputDialog_NewQInputDialog(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QInputDialog) ConnectDone(f func(result int)) {
	defer qt.Recovering("connect QInputDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QInputDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QInputDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQInputDialogDone
func callbackQInputDialogDone(ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QInputDialog::done")

	var signal = qt.GetSignal(C.GoString(ptrName), "done")
	if signal != nil {
		defer signal.(func(int))(int(result))
		return true
	}
	return false

}

func QInputDialog_GetInt(parent QWidget_ITF, title string, label string, value int, min int, max int, step int, ok bool, flags core.Qt__WindowType) int {
	defer qt.Recovering("QInputDialog::getInt")

	return int(C.QInputDialog_QInputDialog_GetInt(PointerFromQWidget(parent), C.CString(title), C.CString(label), C.int(value), C.int(min), C.int(max), C.int(step), C.int(qt.GoBoolToInt(ok)), C.int(flags)))
}

func QInputDialog_GetItem(parent QWidget_ITF, title string, label string, items []string, current int, editable bool, ok bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	defer qt.Recovering("QInputDialog::getItem")

	return C.GoString(C.QInputDialog_QInputDialog_GetItem(PointerFromQWidget(parent), C.CString(title), C.CString(label), C.CString(strings.Join(items, ",,,")), C.int(current), C.int(qt.GoBoolToInt(editable)), C.int(qt.GoBoolToInt(ok)), C.int(flags), C.int(inputMethodHints)))
}

func QInputDialog_GetMultiLineText(parent QWidget_ITF, title string, label string, text string, ok bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	defer qt.Recovering("QInputDialog::getMultiLineText")

	return C.GoString(C.QInputDialog_QInputDialog_GetMultiLineText(PointerFromQWidget(parent), C.CString(title), C.CString(label), C.CString(text), C.int(qt.GoBoolToInt(ok)), C.int(flags), C.int(inputMethodHints)))
}

func QInputDialog_GetText(parent QWidget_ITF, title string, label string, mode QLineEdit__EchoMode, text string, ok bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	defer qt.Recovering("QInputDialog::getText")

	return C.GoString(C.QInputDialog_QInputDialog_GetText(PointerFromQWidget(parent), C.CString(title), C.CString(label), C.int(mode), C.CString(text), C.int(qt.GoBoolToInt(ok)), C.int(flags), C.int(inputMethodHints)))
}

func (ptr *QInputDialog) ConnectIntValueChanged(f func(value int)) {
	defer qt.Recovering("connect QInputDialog::intValueChanged")

	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectIntValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "intValueChanged", f)
	}
}

func (ptr *QInputDialog) DisconnectIntValueChanged() {
	defer qt.Recovering("disconnect QInputDialog::intValueChanged")

	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectIntValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "intValueChanged")
	}
}

//export callbackQInputDialogIntValueChanged
func callbackQInputDialogIntValueChanged(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QInputDialog::intValueChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "intValueChanged")
	if signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QInputDialog) ConnectIntValueSelected(f func(value int)) {
	defer qt.Recovering("connect QInputDialog::intValueSelected")

	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectIntValueSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "intValueSelected", f)
	}
}

func (ptr *QInputDialog) DisconnectIntValueSelected() {
	defer qt.Recovering("disconnect QInputDialog::intValueSelected")

	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectIntValueSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "intValueSelected")
	}
}

//export callbackQInputDialogIntValueSelected
func callbackQInputDialogIntValueSelected(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QInputDialog::intValueSelected")

	var signal = qt.GetSignal(C.GoString(ptrName), "intValueSelected")
	if signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QInputDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QInputDialog::open")

	if ptr.Pointer() != nil {
		C.QInputDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QInputDialog) SetIntRange(min int, max int) {
	defer qt.Recovering("QInputDialog::setIntRange")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntRange(ptr.Pointer(), C.int(min), C.int(max))
	}
}

func (ptr *QInputDialog) SetOption(option QInputDialog__InputDialogOption, on bool) {
	defer qt.Recovering("QInputDialog::setOption")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QInputDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QInputDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QInputDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QInputDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQInputDialogSetVisible
func callbackQInputDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QInputDialog::setVisible")

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QInputDialog) TestOption(option QInputDialog__InputDialogOption) bool {
	defer qt.Recovering("QInputDialog::testOption")

	if ptr.Pointer() != nil {
		return C.QInputDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QInputDialog) ConnectTextValueChanged(f func(text string)) {
	defer qt.Recovering("connect QInputDialog::textValueChanged")

	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectTextValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textValueChanged", f)
	}
}

func (ptr *QInputDialog) DisconnectTextValueChanged() {
	defer qt.Recovering("disconnect QInputDialog::textValueChanged")

	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectTextValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textValueChanged")
	}
}

//export callbackQInputDialogTextValueChanged
func callbackQInputDialogTextValueChanged(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QInputDialog::textValueChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "textValueChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QInputDialog) ConnectTextValueSelected(f func(text string)) {
	defer qt.Recovering("connect QInputDialog::textValueSelected")

	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectTextValueSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textValueSelected", f)
	}
}

func (ptr *QInputDialog) DisconnectTextValueSelected() {
	defer qt.Recovering("disconnect QInputDialog::textValueSelected")

	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectTextValueSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textValueSelected")
	}
}

//export callbackQInputDialogTextValueSelected
func callbackQInputDialogTextValueSelected(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QInputDialog::textValueSelected")

	var signal = qt.GetSignal(C.GoString(ptrName), "textValueSelected")
	if signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QInputDialog) DestroyQInputDialog() {
	defer qt.Recovering("QInputDialog::~QInputDialog")

	if ptr.Pointer() != nil {
		C.QInputDialog_DestroyQInputDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
