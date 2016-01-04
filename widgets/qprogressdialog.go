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
func callbackQProgressDialogCanceled(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QProgressDialog::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QProgressDialog) Canceled() {
	defer qt.Recovering("QProgressDialog::canceled")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Canceled(ptr.Pointer())
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
func callbackQProgressDialogChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQProgressDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QProgressDialog) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QProgressDialog) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
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
func callbackQProgressDialogCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQProgressDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QProgressDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QProgressDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
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
func callbackQProgressDialogResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QProgressDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QProgressDialog) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QProgressDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQProgressDialogShowEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
	} else {
		NewQProgressDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(e))
	}
}

func (ptr *QProgressDialog) ShowEvent(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QProgressDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
}

func (ptr *QProgressDialog) ShowEventDefault(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QProgressDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
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
func callbackQProgressDialogAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QProgressDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProgressDialog) Accept() {
	defer qt.Recovering("QProgressDialog::accept")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) AcceptDefault() {
	defer qt.Recovering("QProgressDialog::accept")

	if ptr.Pointer() != nil {
		C.QProgressDialog_AcceptDefault(ptr.Pointer())
	}
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
func callbackQProgressDialogContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQProgressDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QProgressDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QProgressDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QProgressDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QProgressDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
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
func callbackQProgressDialogDone(ptr unsafe.Pointer, ptrName *C.char, r C.int) bool {
	defer qt.Recovering("callback QProgressDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(r))
		return true
	}
	return false

}

func (ptr *QProgressDialog) Done(r int) {
	defer qt.Recovering("QProgressDialog::done")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Done(ptr.Pointer(), C.int(r))
	}
}

func (ptr *QProgressDialog) DoneDefault(r int) {
	defer qt.Recovering("QProgressDialog::done")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DoneDefault(ptr.Pointer(), C.int(r))
	}
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
func callbackQProgressDialogKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQProgressDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QProgressDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QProgressDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQProgressDialogReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QProgressDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProgressDialog) Reject() {
	defer qt.Recovering("QProgressDialog::reject")

	if ptr.Pointer() != nil {
		C.QProgressDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) RejectDefault() {
	defer qt.Recovering("QProgressDialog::reject")

	if ptr.Pointer() != nil {
		C.QProgressDialog_RejectDefault(ptr.Pointer())
	}
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
func callbackQProgressDialogSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QProgressDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QProgressDialog) SetVisible(visible bool) {
	defer qt.Recovering("QProgressDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressDialog) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QProgressDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQProgressDialogActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QProgressDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QProgressDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QProgressDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQProgressDialogDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QProgressDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQProgressDialogDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QProgressDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQProgressDialogDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QProgressDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQProgressDialogDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QProgressDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QProgressDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQProgressDialogEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQProgressDialogFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QProgressDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQProgressDialogFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QProgressDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQProgressDialogHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QProgressDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QProgressDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QProgressDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQProgressDialogLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQProgressDialogMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QProgressDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QProgressDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QProgressDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQProgressDialogPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QProgressDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QProgressDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QProgressDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQProgressDialogInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQProgressDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QProgressDialog) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QProgressDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QProgressDialog_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QProgressDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QProgressDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QProgressDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQProgressDialogInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QProgressDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QProgressDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QProgressDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQProgressDialogKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QProgressDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQProgressDialogMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQProgressDialogMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQProgressDialogMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQProgressDialogMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQProgressDialogTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QProgressDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QProgressDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QProgressDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQProgressDialogWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QProgressDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QProgressDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QProgressDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQProgressDialogTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProgressDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProgressDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProgressDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQProgressDialogChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProgressDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProgressDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProgressDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQProgressDialogCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QProgressDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
