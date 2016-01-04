package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QErrorMessage struct {
	QDialog
}

type QErrorMessage_ITF interface {
	QDialog_ITF
	QErrorMessage_PTR() *QErrorMessage
}

func PointerFromQErrorMessage(ptr QErrorMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QErrorMessage_PTR().Pointer()
	}
	return nil
}

func NewQErrorMessageFromPointer(ptr unsafe.Pointer) *QErrorMessage {
	var n = new(QErrorMessage)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QErrorMessage_") {
		n.SetObjectName("QErrorMessage_" + qt.Identifier())
	}
	return n
}

func (ptr *QErrorMessage) QErrorMessage_PTR() *QErrorMessage {
	return ptr
}

func NewQErrorMessage(parent QWidget_ITF) *QErrorMessage {
	defer qt.Recovering("QErrorMessage::QErrorMessage")

	return NewQErrorMessageFromPointer(C.QErrorMessage_NewQErrorMessage(PointerFromQWidget(parent)))
}

func (ptr *QErrorMessage) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQErrorMessageChangeEvent
func callbackQErrorMessageChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQErrorMessageFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QErrorMessage) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QErrorMessage) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QErrorMessage) ConnectDone(f func(a int)) {
	defer qt.Recovering("connect QErrorMessage::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QErrorMessage) DisconnectDone() {
	defer qt.Recovering("disconnect QErrorMessage::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQErrorMessageDone
func callbackQErrorMessageDone(ptr unsafe.Pointer, ptrName *C.char, a C.int) bool {
	defer qt.Recovering("callback QErrorMessage::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(a))
		return true
	}
	return false

}

func (ptr *QErrorMessage) Done(a int) {
	defer qt.Recovering("QErrorMessage::done")

	if ptr.Pointer() != nil {
		C.QErrorMessage_Done(ptr.Pointer(), C.int(a))
	}
}

func (ptr *QErrorMessage) DoneDefault(a int) {
	defer qt.Recovering("QErrorMessage::done")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DoneDefault(ptr.Pointer(), C.int(a))
	}
}

func QErrorMessage_QtHandler() *QErrorMessage {
	defer qt.Recovering("QErrorMessage::qtHandler")

	return NewQErrorMessageFromPointer(C.QErrorMessage_QErrorMessage_QtHandler())
}

func (ptr *QErrorMessage) ShowMessage(message string) {
	defer qt.Recovering("QErrorMessage::showMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QErrorMessage) ShowMessage2(message string, ty string) {
	defer qt.Recovering("QErrorMessage::showMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage2(ptr.Pointer(), C.CString(message), C.CString(ty))
	}
}

func (ptr *QErrorMessage) DestroyQErrorMessage() {
	defer qt.Recovering("QErrorMessage::~QErrorMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DestroyQErrorMessage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QErrorMessage) ConnectAccept(f func()) {
	defer qt.Recovering("connect QErrorMessage::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QErrorMessage) DisconnectAccept() {
	defer qt.Recovering("disconnect QErrorMessage::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQErrorMessageAccept
func callbackQErrorMessageAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QErrorMessage::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QErrorMessage) Accept() {
	defer qt.Recovering("QErrorMessage::accept")

	if ptr.Pointer() != nil {
		C.QErrorMessage_Accept(ptr.Pointer())
	}
}

func (ptr *QErrorMessage) AcceptDefault() {
	defer qt.Recovering("QErrorMessage::accept")

	if ptr.Pointer() != nil {
		C.QErrorMessage_AcceptDefault(ptr.Pointer())
	}
}

func (ptr *QErrorMessage) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QErrorMessage::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QErrorMessage::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQErrorMessageCloseEvent
func callbackQErrorMessageCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQErrorMessageFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QErrorMessage) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::closeEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QErrorMessage) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::closeEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QErrorMessage) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QErrorMessage::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QErrorMessage::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQErrorMessageContextMenuEvent
func callbackQErrorMessageContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQErrorMessageFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QErrorMessage) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QErrorMessage::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QErrorMessage) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QErrorMessage::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QErrorMessage) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QErrorMessage::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QErrorMessage::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQErrorMessageKeyPressEvent
func callbackQErrorMessageKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQErrorMessageFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QErrorMessage) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QErrorMessage::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QErrorMessage) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QErrorMessage::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QErrorMessage) ConnectOpen(f func()) {
	defer qt.Recovering("connect QErrorMessage::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QErrorMessage) DisconnectOpen() {
	defer qt.Recovering("disconnect QErrorMessage::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQErrorMessageOpen
func callbackQErrorMessageOpen(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QErrorMessage::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QErrorMessage) Open() {
	defer qt.Recovering("QErrorMessage::open")

	if ptr.Pointer() != nil {
		C.QErrorMessage_Open(ptr.Pointer())
	}
}

func (ptr *QErrorMessage) OpenDefault() {
	defer qt.Recovering("QErrorMessage::open")

	if ptr.Pointer() != nil {
		C.QErrorMessage_OpenDefault(ptr.Pointer())
	}
}

func (ptr *QErrorMessage) ConnectReject(f func()) {
	defer qt.Recovering("connect QErrorMessage::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QErrorMessage) DisconnectReject() {
	defer qt.Recovering("disconnect QErrorMessage::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQErrorMessageReject
func callbackQErrorMessageReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QErrorMessage::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QErrorMessage) Reject() {
	defer qt.Recovering("QErrorMessage::reject")

	if ptr.Pointer() != nil {
		C.QErrorMessage_Reject(ptr.Pointer())
	}
}

func (ptr *QErrorMessage) RejectDefault() {
	defer qt.Recovering("QErrorMessage::reject")

	if ptr.Pointer() != nil {
		C.QErrorMessage_RejectDefault(ptr.Pointer())
	}
}

func (ptr *QErrorMessage) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QErrorMessage::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QErrorMessage::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQErrorMessageResizeEvent
func callbackQErrorMessageResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQErrorMessageFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QErrorMessage) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QErrorMessage::resizeEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QErrorMessage) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QErrorMessage::resizeEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QErrorMessage) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QErrorMessage::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QErrorMessage) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QErrorMessage::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQErrorMessageSetVisible
func callbackQErrorMessageSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QErrorMessage::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QErrorMessage) SetVisible(visible bool) {
	defer qt.Recovering("QErrorMessage::setVisible")

	if ptr.Pointer() != nil {
		C.QErrorMessage_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QErrorMessage) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QErrorMessage::setVisible")

	if ptr.Pointer() != nil {
		C.QErrorMessage_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QErrorMessage) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QErrorMessage::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QErrorMessage::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQErrorMessageShowEvent
func callbackQErrorMessageShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QErrorMessage::showEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QErrorMessage) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QErrorMessage::showEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QErrorMessage::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QErrorMessage::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQErrorMessageActionEvent
func callbackQErrorMessageActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QErrorMessage::actionEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QErrorMessage) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QErrorMessage::actionEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QErrorMessage::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQErrorMessageDragEnterEvent
func callbackQErrorMessageDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QErrorMessage) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QErrorMessage::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQErrorMessageDragLeaveEvent
func callbackQErrorMessageDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QErrorMessage) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QErrorMessage::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQErrorMessageDragMoveEvent
func callbackQErrorMessageDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QErrorMessage) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QErrorMessage::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQErrorMessageDropEvent
func callbackQErrorMessageDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dropEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QErrorMessage) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QErrorMessage::dropEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QErrorMessage::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQErrorMessageEnterEvent
func callbackQErrorMessageEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::enterEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QErrorMessage) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::enterEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QErrorMessage::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QErrorMessage::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQErrorMessageFocusInEvent
func callbackQErrorMessageFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QErrorMessage::focusInEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QErrorMessage) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QErrorMessage::focusInEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QErrorMessage::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QErrorMessage::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQErrorMessageFocusOutEvent
func callbackQErrorMessageFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QErrorMessage::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QErrorMessage) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QErrorMessage::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QErrorMessage::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QErrorMessage::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQErrorMessageHideEvent
func callbackQErrorMessageHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QErrorMessage::hideEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QErrorMessage) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QErrorMessage::hideEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQErrorMessageLeaveEvent
func callbackQErrorMessageLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::leaveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QErrorMessage) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::leaveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QErrorMessage::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQErrorMessageMoveEvent
func callbackQErrorMessageMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QErrorMessage::moveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QErrorMessage) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QErrorMessage::moveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QErrorMessage::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QErrorMessage::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQErrorMessagePaintEvent
func callbackQErrorMessagePaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QErrorMessage::paintEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QErrorMessage) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QErrorMessage::paintEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QErrorMessage::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QErrorMessage) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QErrorMessage::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQErrorMessageInitPainter
func callbackQErrorMessageInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQErrorMessageFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QErrorMessage) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QErrorMessage::initPainter")

	if ptr.Pointer() != nil {
		C.QErrorMessage_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QErrorMessage) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QErrorMessage::initPainter")

	if ptr.Pointer() != nil {
		C.QErrorMessage_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QErrorMessage) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QErrorMessage::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QErrorMessage::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQErrorMessageInputMethodEvent
func callbackQErrorMessageInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QErrorMessage::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QErrorMessage) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QErrorMessage::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QErrorMessage::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QErrorMessage::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQErrorMessageKeyReleaseEvent
func callbackQErrorMessageKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QErrorMessage::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QErrorMessage) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QErrorMessage::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQErrorMessageMouseDoubleClickEvent
func callbackQErrorMessageMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQErrorMessageMouseMoveEvent
func callbackQErrorMessageMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQErrorMessageMousePressEvent
func callbackQErrorMessageMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQErrorMessageMouseReleaseEvent
func callbackQErrorMessageMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QErrorMessage::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QErrorMessage::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QErrorMessage::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQErrorMessageTabletEvent
func callbackQErrorMessageTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QErrorMessage::tabletEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QErrorMessage) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QErrorMessage::tabletEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QErrorMessage::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QErrorMessage::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQErrorMessageWheelEvent
func callbackQErrorMessageWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QErrorMessage::wheelEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QErrorMessage) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QErrorMessage::wheelEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QErrorMessage::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QErrorMessage::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQErrorMessageTimerEvent
func callbackQErrorMessageTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QErrorMessage::timerEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QErrorMessage) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QErrorMessage::timerEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QErrorMessage::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QErrorMessage::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQErrorMessageChildEvent
func callbackQErrorMessageChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QErrorMessage::childEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QErrorMessage) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QErrorMessage::childEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QErrorMessage) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QErrorMessage::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQErrorMessageCustomEvent
func callbackQErrorMessageCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QErrorMessage::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQErrorMessageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QErrorMessage) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::customEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QErrorMessage) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QErrorMessage::customEvent")

	if ptr.Pointer() != nil {
		C.QErrorMessage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
