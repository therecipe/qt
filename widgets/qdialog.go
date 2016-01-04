package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDialog struct {
	QWidget
}

type QDialog_ITF interface {
	QWidget_ITF
	QDialog_PTR() *QDialog
}

func PointerFromQDialog(ptr QDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func NewQDialogFromPointer(ptr unsafe.Pointer) *QDialog {
	var n = new(QDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDialog_") {
		n.SetObjectName("QDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QDialog) QDialog_PTR() *QDialog {
	return ptr
}

//QDialog::DialogCode
type QDialog__DialogCode int64

const (
	QDialog__Rejected = QDialog__DialogCode(0)
	QDialog__Accepted = QDialog__DialogCode(1)
)

func (ptr *QDialog) IsSizeGripEnabled() bool {
	defer qt.Recovering("QDialog::isSizeGripEnabled")

	if ptr.Pointer() != nil {
		return C.QDialog_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDialog) SetModal(modal bool) {
	defer qt.Recovering("QDialog::setModal")

	if ptr.Pointer() != nil {
		C.QDialog_SetModal(ptr.Pointer(), C.int(qt.GoBoolToInt(modal)))
	}
}

func (ptr *QDialog) SetResult(i int) {
	defer qt.Recovering("QDialog::setResult")

	if ptr.Pointer() != nil {
		C.QDialog_SetResult(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QDialog) SetSizeGripEnabled(v bool) {
	defer qt.Recovering("QDialog::setSizeGripEnabled")

	if ptr.Pointer() != nil {
		C.QDialog_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQDialog(parent QWidget_ITF, f core.Qt__WindowType) *QDialog {
	defer qt.Recovering("QDialog::QDialog")

	return NewQDialogFromPointer(C.QDialog_NewQDialog(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQDialogAccept
func callbackQDialogAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDialog) Accept() {
	defer qt.Recovering("QDialog::accept")

	if ptr.Pointer() != nil {
		C.QDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QDialog) AcceptDefault() {
	defer qt.Recovering("QDialog::accept")

	if ptr.Pointer() != nil {
		C.QDialog_AcceptDefault(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectAccepted(f func()) {
	defer qt.Recovering("connect QDialog::accepted")

	if ptr.Pointer() != nil {
		C.QDialog_ConnectAccepted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accepted", f)
	}
}

func (ptr *QDialog) DisconnectAccepted() {
	defer qt.Recovering("disconnect QDialog::accepted")

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accepted")
	}
}

//export callbackQDialogAccepted
func callbackQDialogAccepted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDialog::accepted")

	if signal := qt.GetSignal(C.GoString(ptrName), "accepted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialog) Accepted() {
	defer qt.Recovering("QDialog::accepted")

	if ptr.Pointer() != nil {
		C.QDialog_Accepted(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDialogCloseEvent
func callbackQDialogCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDialogContextMenuEvent
func callbackQDialogContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QDialog) ConnectDone(f func(r int)) {
	defer qt.Recovering("connect QDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQDialogDone
func callbackQDialogDone(ptr unsafe.Pointer, ptrName *C.char, r C.int) bool {
	defer qt.Recovering("callback QDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(r))
		return true
	}
	return false

}

func (ptr *QDialog) Done(r int) {
	defer qt.Recovering("QDialog::done")

	if ptr.Pointer() != nil {
		C.QDialog_Done(ptr.Pointer(), C.int(r))
	}
}

func (ptr *QDialog) DoneDefault(r int) {
	defer qt.Recovering("QDialog::done")

	if ptr.Pointer() != nil {
		C.QDialog_DoneDefault(ptr.Pointer(), C.int(r))
	}
}

func (ptr *QDialog) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDialog::event")

	if ptr.Pointer() != nil {
		return C.QDialog_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDialog) EventFilter(o core.QObject_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QDialog::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDialog_EventFilter(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDialog) Exec() int {
	defer qt.Recovering("QDialog::exec")

	if ptr.Pointer() != nil {
		return int(C.QDialog_Exec(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialog) ConnectFinished(f func(result int)) {
	defer qt.Recovering("connect QDialog::finished")

	if ptr.Pointer() != nil {
		C.QDialog_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDialog) DisconnectFinished() {
	defer qt.Recovering("disconnect QDialog::finished")

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDialogFinished
func callbackQDialogFinished(ptr unsafe.Pointer, ptrName *C.char, result C.int) {
	defer qt.Recovering("callback QDialog::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(int))(int(result))
	}

}

func (ptr *QDialog) Finished(result int) {
	defer qt.Recovering("QDialog::finished")

	if ptr.Pointer() != nil {
		C.QDialog_Finished(ptr.Pointer(), C.int(result))
	}
}

func (ptr *QDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDialogKeyPressEvent
func callbackQDialogKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QDialog) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDialog::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDialog_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQDialogOpen
func callbackQDialogOpen(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QDialog::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDialog) Open() {
	defer qt.Recovering("QDialog::open")

	if ptr.Pointer() != nil {
		C.QDialog_Open(ptr.Pointer())
	}
}

func (ptr *QDialog) OpenDefault() {
	defer qt.Recovering("QDialog::open")

	if ptr.Pointer() != nil {
		C.QDialog_OpenDefault(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQDialogReject
func callbackQDialogReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDialog) Reject() {
	defer qt.Recovering("QDialog::reject")

	if ptr.Pointer() != nil {
		C.QDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QDialog) RejectDefault() {
	defer qt.Recovering("QDialog::reject")

	if ptr.Pointer() != nil {
		C.QDialog_RejectDefault(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectRejected(f func()) {
	defer qt.Recovering("connect QDialog::rejected")

	if ptr.Pointer() != nil {
		C.QDialog_ConnectRejected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rejected", f)
	}
}

func (ptr *QDialog) DisconnectRejected() {
	defer qt.Recovering("disconnect QDialog::rejected")

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectRejected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rejected")
	}
}

//export callbackQDialogRejected
func callbackQDialogRejected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDialog::rejected")

	if signal := qt.GetSignal(C.GoString(ptrName), "rejected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialog) Rejected() {
	defer qt.Recovering("QDialog::rejected")

	if ptr.Pointer() != nil {
		C.QDialog_Rejected(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDialogResizeEvent
func callbackQDialogResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QDialog) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QDialog) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QDialog) Result() int {
	defer qt.Recovering("QDialog::result")

	if ptr.Pointer() != nil {
		return int(C.QDialog_Result(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDialogSetVisible
func callbackQDialogSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDialog) SetVisible(visible bool) {
	defer qt.Recovering("QDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDialog) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QDialog_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDialogShowEvent
func callbackQDialogShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDialog) SizeHint() *core.QSize {
	defer qt.Recovering("QDialog::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDialog_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDialog) DestroyQDialog() {
	defer qt.Recovering("QDialog::~QDialog")

	if ptr.Pointer() != nil {
		C.QDialog_DestroyQDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDialogActionEvent
func callbackQDialogActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDialogDragEnterEvent
func callbackQDialogDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDialogDragLeaveEvent
func callbackQDialogDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDialogDragMoveEvent
func callbackQDialogDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDialogDropEvent
func callbackQDialogDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDialogEnterEvent
func callbackQDialogEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDialogFocusInEvent
func callbackQDialogFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDialogFocusOutEvent
func callbackQDialogFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDialogHideEvent
func callbackQDialogHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDialogLeaveEvent
func callbackQDialogLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDialogMoveEvent
func callbackQDialogMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDialogPaintEvent
func callbackQDialogPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDialog) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDialogChangeEvent
func callbackQDialogChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialog) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDialog) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDialogInitPainter
func callbackQDialogInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QDialog) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QDialog_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDialogInputMethodEvent
func callbackQDialogInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDialogKeyReleaseEvent
func callbackQDialogKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDialogMouseDoubleClickEvent
func callbackQDialogMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDialogMouseMoveEvent
func callbackQDialogMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDialogMousePressEvent
func callbackQDialogMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDialogMouseReleaseEvent
func callbackQDialogMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDialogTabletEvent
func callbackQDialogTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDialogWheelEvent
func callbackQDialogWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDialogTimerEvent
func callbackQDialogTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDialogChildEvent
func callbackQDialogChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDialogCustomEvent
func callbackQDialogCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
