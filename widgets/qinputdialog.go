package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
func callbackQInputDialogDone(ptr unsafe.Pointer, ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QInputDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
		return true
	}
	return false

}

func (ptr *QInputDialog) Done(result int) {
	defer qt.Recovering("QInputDialog::done")

	if ptr.Pointer() != nil {
		C.QInputDialog_Done(ptr.Pointer(), C.int(result))
	}
}

func (ptr *QInputDialog) DoneDefault(result int) {
	defer qt.Recovering("QInputDialog::done")

	if ptr.Pointer() != nil {
		C.QInputDialog_DoneDefault(ptr.Pointer(), C.int(result))
	}
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
func callbackQInputDialogIntValueChanged(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QInputDialog::intValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "intValueChanged"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QInputDialog) IntValueChanged(value int) {
	defer qt.Recovering("QInputDialog::intValueChanged")

	if ptr.Pointer() != nil {
		C.QInputDialog_IntValueChanged(ptr.Pointer(), C.int(value))
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
func callbackQInputDialogIntValueSelected(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QInputDialog::intValueSelected")

	if signal := qt.GetSignal(C.GoString(ptrName), "intValueSelected"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QInputDialog) IntValueSelected(value int) {
	defer qt.Recovering("QInputDialog::intValueSelected")

	if ptr.Pointer() != nil {
		C.QInputDialog_IntValueSelected(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QInputDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QInputDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QInputDialog_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
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
func callbackQInputDialogSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QInputDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QInputDialog) SetVisible(visible bool) {
	defer qt.Recovering("QInputDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QInputDialog) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QInputDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QInputDialog_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QInputDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QInputDialog::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QInputDialog_SizeHint(ptr.Pointer()))
	}
	return nil
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
func callbackQInputDialogTextValueChanged(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QInputDialog::textValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textValueChanged"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QInputDialog) TextValueChanged(text string) {
	defer qt.Recovering("QInputDialog::textValueChanged")

	if ptr.Pointer() != nil {
		C.QInputDialog_TextValueChanged(ptr.Pointer(), C.CString(text))
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
func callbackQInputDialogTextValueSelected(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QInputDialog::textValueSelected")

	if signal := qt.GetSignal(C.GoString(ptrName), "textValueSelected"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QInputDialog) TextValueSelected(text string) {
	defer qt.Recovering("QInputDialog::textValueSelected")

	if ptr.Pointer() != nil {
		C.QInputDialog_TextValueSelected(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QInputDialog) DestroyQInputDialog() {
	defer qt.Recovering("QInputDialog::~QInputDialog")

	if ptr.Pointer() != nil {
		C.QInputDialog_DestroyQInputDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QInputDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QInputDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QInputDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QInputDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQInputDialogAccept
func callbackQInputDialogAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QInputDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QInputDialog) Accept() {
	defer qt.Recovering("QInputDialog::accept")

	if ptr.Pointer() != nil {
		C.QInputDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QInputDialog) AcceptDefault() {
	defer qt.Recovering("QInputDialog::accept")

	if ptr.Pointer() != nil {
		C.QInputDialog_AcceptDefault(ptr.Pointer())
	}
}

func (ptr *QInputDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QInputDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QInputDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQInputDialogCloseEvent
func callbackQInputDialogCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQInputDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QInputDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QInputDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QInputDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QInputDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QInputDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QInputDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QInputDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQInputDialogContextMenuEvent
func callbackQInputDialogContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQInputDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QInputDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QInputDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QInputDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QInputDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QInputDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QInputDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QInputDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQInputDialogKeyPressEvent
func callbackQInputDialogKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQInputDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QInputDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QInputDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QInputDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QInputDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QInputDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QInputDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QInputDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QInputDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQInputDialogReject
func callbackQInputDialogReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QInputDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QInputDialog) Reject() {
	defer qt.Recovering("QInputDialog::reject")

	if ptr.Pointer() != nil {
		C.QInputDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QInputDialog) RejectDefault() {
	defer qt.Recovering("QInputDialog::reject")

	if ptr.Pointer() != nil {
		C.QInputDialog_RejectDefault(ptr.Pointer())
	}
}

func (ptr *QInputDialog) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QInputDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QInputDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQInputDialogResizeEvent
func callbackQInputDialogResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQInputDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QInputDialog) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QInputDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QInputDialog) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QInputDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QInputDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QInputDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QInputDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQInputDialogShowEvent
func callbackQInputDialogShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QInputDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QInputDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QInputDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QInputDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QInputDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QInputDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QInputDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQInputDialogActionEvent
func callbackQInputDialogActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QInputDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QInputDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QInputDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QInputDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QInputDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QInputDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QInputDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQInputDialogDragEnterEvent
func callbackQInputDialogDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QInputDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QInputDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QInputDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QInputDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QInputDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QInputDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QInputDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQInputDialogDragLeaveEvent
func callbackQInputDialogDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QInputDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QInputDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QInputDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QInputDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QInputDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QInputDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QInputDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQInputDialogDragMoveEvent
func callbackQInputDialogDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QInputDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QInputDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QInputDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QInputDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QInputDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QInputDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QInputDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQInputDialogDropEvent
func callbackQInputDialogDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QInputDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QInputDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QInputDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QInputDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QInputDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInputDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QInputDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQInputDialogEnterEvent
func callbackQInputDialogEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInputDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QInputDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QInputDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQInputDialogFocusInEvent
func callbackQInputDialogFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QInputDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QInputDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QInputDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QInputDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QInputDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QInputDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QInputDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQInputDialogFocusOutEvent
func callbackQInputDialogFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QInputDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QInputDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QInputDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QInputDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QInputDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QInputDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QInputDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQInputDialogHideEvent
func callbackQInputDialogHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QInputDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QInputDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QInputDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QInputDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QInputDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInputDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QInputDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQInputDialogLeaveEvent
func callbackQInputDialogLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInputDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QInputDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QInputDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQInputDialogMoveEvent
func callbackQInputDialogMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QInputDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QInputDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QInputDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QInputDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QInputDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QInputDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QInputDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQInputDialogPaintEvent
func callbackQInputDialogPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QInputDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QInputDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QInputDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QInputDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QInputDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInputDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QInputDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQInputDialogChangeEvent
func callbackQInputDialogChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInputDialog) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QInputDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QInputDialog) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QInputDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQInputDialogInitPainter
func callbackQInputDialogInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQInputDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QInputDialog) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QInputDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QInputDialog_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QInputDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QInputDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QInputDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QInputDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QInputDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QInputDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQInputDialogInputMethodEvent
func callbackQInputDialogInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QInputDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QInputDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QInputDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QInputDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QInputDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QInputDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QInputDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQInputDialogKeyReleaseEvent
func callbackQInputDialogKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QInputDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QInputDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QInputDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QInputDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QInputDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QInputDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QInputDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQInputDialogMouseDoubleClickEvent
func callbackQInputDialogMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QInputDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QInputDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QInputDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQInputDialogMouseMoveEvent
func callbackQInputDialogMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QInputDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QInputDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QInputDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQInputDialogMousePressEvent
func callbackQInputDialogMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QInputDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QInputDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QInputDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQInputDialogMouseReleaseEvent
func callbackQInputDialogMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QInputDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QInputDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QInputDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QInputDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QInputDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQInputDialogTabletEvent
func callbackQInputDialogTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QInputDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QInputDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QInputDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QInputDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QInputDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QInputDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QInputDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQInputDialogWheelEvent
func callbackQInputDialogWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QInputDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QInputDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QInputDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QInputDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QInputDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QInputDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QInputDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQInputDialogTimerEvent
func callbackQInputDialogTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInputDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInputDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInputDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInputDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInputDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QInputDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QInputDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQInputDialogChildEvent
func callbackQInputDialogChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInputDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInputDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInputDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInputDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInputDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInputDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QInputDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QInputDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQInputDialogCustomEvent
func callbackQInputDialogCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInputDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInputDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInputDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QInputDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
