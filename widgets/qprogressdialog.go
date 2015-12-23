package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QProgressDialog struct {
	QDialog
}

type QProgressDialog_ITF interface {
	QDialog_ITF
	QProgressDialog_PTR() *QProgressDialog
}

func PointerFromQProgressDialog(ptr QProgressDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressDialog_PTR().Pointer()
	}
	return nil
}

func NewQProgressDialogFromPointer(ptr unsafe.Pointer) *QProgressDialog {
	var n = new(QProgressDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProgressDialog_") {
		n.SetObjectName("QProgressDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QProgressDialog) QProgressDialog_PTR() *QProgressDialog {
	return ptr
}

func (ptr *QProgressDialog) AutoClose() bool {
	defer qt.Recovering("QProgressDialog::autoClose")

	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoClose(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressDialog) AutoReset() bool {
	defer qt.Recovering("QProgressDialog::autoReset")

	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoReset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressDialog) LabelText() string {
	defer qt.Recovering("QProgressDialog::labelText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressDialog_LabelText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressDialog) Maximum() int {
	defer qt.Recovering("QProgressDialog::maximum")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) Minimum() int {
	defer qt.Recovering("QProgressDialog::minimum")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) MinimumDuration() int {
	defer qt.Recovering("QProgressDialog::minimumDuration")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_MinimumDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) SetAutoClose(close bool) {
	defer qt.Recovering("QProgressDialog::setAutoClose")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}

func (ptr *QProgressDialog) SetAutoReset(reset bool) {
	defer qt.Recovering("QProgressDialog::setAutoReset")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoReset(ptr.Pointer(), C.int(qt.GoBoolToInt(reset)))
	}
}

func (ptr *QProgressDialog) SetLabelText(text string) {
	defer qt.Recovering("QProgressDialog::setLabelText")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabelText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QProgressDialog) SetMaximum(maximum int) {
	defer qt.Recovering("QProgressDialog::setMaximum")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressDialog) SetMinimum(minimum int) {
	defer qt.Recovering("QProgressDialog::setMinimum")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressDialog) SetMinimumDuration(ms int) {
	defer qt.Recovering("QProgressDialog::setMinimumDuration")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimumDuration(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QProgressDialog) SetValue(progress int) {
	defer qt.Recovering("QProgressDialog::setValue")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetValue(ptr.Pointer(), C.int(progress))
	}
}

func (ptr *QProgressDialog) Value() int {
	defer qt.Recovering("QProgressDialog::value")

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Value(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) WasCanceled() bool {
	defer qt.Recovering("QProgressDialog::wasCanceled")

	if ptr.Pointer() != nil {
		return C.QProgressDialog_WasCanceled(ptr.Pointer()) != 0
	}
	return false
}

func NewQProgressDialog(parent QWidget_ITF, f core.Qt__WindowType) *QProgressDialog {
	defer qt.Recovering("QProgressDialog::QProgressDialog")

	return NewQProgressDialogFromPointer(C.QProgressDialog_NewQProgressDialog(PointerFromQWidget(parent), C.int(f)))
}

func NewQProgressDialog2(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF, f core.Qt__WindowType) *QProgressDialog {
	defer qt.Recovering("QProgressDialog::QProgressDialog")

	return NewQProgressDialogFromPointer(C.QProgressDialog_NewQProgressDialog2(C.CString(labelText), C.CString(cancelButtonText), C.int(minimum), C.int(maximum), PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QProgressDialog) Cancel() {
	defer qt.Recovering("QProgressDialog::cancel")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Cancel(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QProgressDialog::canceled")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QProgressDialog) DisconnectCanceled() {
	defer qt.Recovering("disconnect QProgressDialog::canceled")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQProgressDialogCanceled
func callbackQProgressDialogCanceled(ptrName *C.char) {
	defer qt.Recovering("callback QProgressDialog::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QProgressDialog) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QProgressDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QProgressDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQProgressDialogChangeEvent
func callbackQProgressDialogChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QProgressDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QProgressDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQProgressDialogCloseEvent
func callbackQProgressDialogCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QProgressDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QProgressDialog::open")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QProgressDialog) Reset() {
	defer qt.Recovering("QProgressDialog::reset")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QProgressDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QProgressDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQProgressDialogResizeEvent
func callbackQProgressDialogResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) SetBar(bar QProgressBar_ITF) {
	defer qt.Recovering("QProgressDialog::setBar")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetBar(ptr.Pointer(), PointerFromQProgressBar(bar))
	}
}

func (ptr *QProgressDialog) SetCancelButton(cancelButton QPushButton_ITF) {
	defer qt.Recovering("QProgressDialog::setCancelButton")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButton(ptr.Pointer(), PointerFromQPushButton(cancelButton))
	}
}

func (ptr *QProgressDialog) SetCancelButtonText(cancelButtonText string) {
	defer qt.Recovering("QProgressDialog::setCancelButtonText")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButtonText(ptr.Pointer(), C.CString(cancelButtonText))
	}
}

func (ptr *QProgressDialog) SetLabel(label QLabel_ITF) {
	defer qt.Recovering("QProgressDialog::setLabel")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabel(ptr.Pointer(), PointerFromQLabel(label))
	}
}

func (ptr *QProgressDialog) SetRange(minimum int, maximum int) {
	defer qt.Recovering("QProgressDialog::setRange")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressDialog) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QProgressDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QProgressDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQProgressDialogShowEvent
func callbackQProgressDialogShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QProgressDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QProgressDialog::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProgressDialog_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProgressDialog) DestroyQProgressDialog() {
	defer qt.Recovering("QProgressDialog::~QProgressDialog")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DestroyQProgressDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QProgressDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QProgressDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QProgressDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QProgressDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQProgressDialogAccept
func callbackQProgressDialogAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QProgressDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QProgressDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QProgressDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQProgressDialogContextMenuEvent
func callbackQProgressDialogContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectDone(f func(r int)) {
	defer qt.Recovering("connect QProgressDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QProgressDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QProgressDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQProgressDialogDone
func callbackQProgressDialogDone(ptrName *C.char, r C.int) bool {
	defer qt.Recovering("callback QProgressDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(r))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QProgressDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QProgressDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQProgressDialogKeyPressEvent
func callbackQProgressDialogKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QProgressDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QProgressDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QProgressDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQProgressDialogOpen
func callbackQProgressDialogOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QProgressDialog::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QProgressDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QProgressDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QProgressDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQProgressDialogReject
func callbackQProgressDialogReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QProgressDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QProgressDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QProgressDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QProgressDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQProgressDialogSetVisible
func callbackQProgressDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QProgressDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QProgressDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QProgressDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQProgressDialogActionEvent
func callbackQProgressDialogActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QProgressDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QProgressDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQProgressDialogDragEnterEvent
func callbackQProgressDialogDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QProgressDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QProgressDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQProgressDialogDragLeaveEvent
func callbackQProgressDialogDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QProgressDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QProgressDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQProgressDialogDragMoveEvent
func callbackQProgressDialogDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QProgressDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QProgressDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQProgressDialogDropEvent
func callbackQProgressDialogDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QProgressDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQProgressDialogEnterEvent
func callbackQProgressDialogEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QProgressDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QProgressDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQProgressDialogFocusInEvent
func callbackQProgressDialogFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QProgressDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QProgressDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQProgressDialogFocusOutEvent
func callbackQProgressDialogFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QProgressDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QProgressDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQProgressDialogHideEvent
func callbackQProgressDialogHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QProgressDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQProgressDialogLeaveEvent
func callbackQProgressDialogLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QProgressDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QProgressDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQProgressDialogMoveEvent
func callbackQProgressDialogMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QProgressDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QProgressDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQProgressDialogPaintEvent
func callbackQProgressDialogPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QProgressDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QProgressDialog) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QProgressDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQProgressDialogInitPainter
func callbackQProgressDialogInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QProgressDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QProgressDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQProgressDialogInputMethodEvent
func callbackQProgressDialogInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QProgressDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QProgressDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQProgressDialogKeyReleaseEvent
func callbackQProgressDialogKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QProgressDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQProgressDialogMouseDoubleClickEvent
func callbackQProgressDialogMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QProgressDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQProgressDialogMouseMoveEvent
func callbackQProgressDialogMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QProgressDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQProgressDialogMousePressEvent
func callbackQProgressDialogMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QProgressDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQProgressDialogMouseReleaseEvent
func callbackQProgressDialogMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QProgressDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QProgressDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQProgressDialogTabletEvent
func callbackQProgressDialogTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QProgressDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QProgressDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQProgressDialogWheelEvent
func callbackQProgressDialogWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QProgressDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProgressDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProgressDialogTimerEvent
func callbackQProgressDialogTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QProgressDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProgressDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProgressDialogChildEvent
func callbackQProgressDialogChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProgressDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProgressDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProgressDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProgressDialogCustomEvent
func callbackQProgressDialogCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProgressDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
