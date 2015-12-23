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
func callbackQInputDialogDone(ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QInputDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "intValueChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "intValueSelected"); signal != nil {
		signal.(func(int))(int(value))
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
func callbackQInputDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QInputDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQInputDialogTextValueChanged(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QInputDialog::textValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textValueChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "textValueSelected"); signal != nil {
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
func callbackQInputDialogAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QInputDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQInputDialogCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQInputDialogContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

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
func callbackQInputDialogKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QInputDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QInputDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QInputDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QInputDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQInputDialogOpen
func callbackQInputDialogOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QInputDialog::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQInputDialogReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QInputDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQInputDialogResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

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
func callbackQInputDialogShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQInputDialogInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQInputDialogCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QInputDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
