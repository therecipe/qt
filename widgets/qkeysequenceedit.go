package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QKeySequenceEdit struct {
	QWidget
}

type QKeySequenceEdit_ITF interface {
	QWidget_ITF
	QKeySequenceEdit_PTR() *QKeySequenceEdit
}

func PointerFromQKeySequenceEdit(ptr QKeySequenceEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequenceEdit_PTR().Pointer()
	}
	return nil
}

func NewQKeySequenceEditFromPointer(ptr unsafe.Pointer) *QKeySequenceEdit {
	var n = new(QKeySequenceEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QKeySequenceEdit_") {
		n.SetObjectName("QKeySequenceEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QKeySequenceEdit) QKeySequenceEdit_PTR() *QKeySequenceEdit {
	return ptr
}

func (ptr *QKeySequenceEdit) SetKeySequence(keySequence gui.QKeySequence_ITF) {
	defer qt.Recovering("QKeySequenceEdit::setKeySequence")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetKeySequence(ptr.Pointer(), gui.PointerFromQKeySequence(keySequence))
	}
}

func NewQKeySequenceEdit(parent QWidget_ITF) *QKeySequenceEdit {
	defer qt.Recovering("QKeySequenceEdit::QKeySequenceEdit")

	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit(PointerFromQWidget(parent)))
}

func NewQKeySequenceEdit2(keySequence gui.QKeySequence_ITF, parent QWidget_ITF) *QKeySequenceEdit {
	defer qt.Recovering("QKeySequenceEdit::QKeySequenceEdit")

	return NewQKeySequenceEditFromPointer(C.QKeySequenceEdit_NewQKeySequenceEdit2(gui.PointerFromQKeySequence(keySequence), PointerFromQWidget(parent)))
}

func (ptr *QKeySequenceEdit) Clear() {
	defer qt.Recovering("QKeySequenceEdit::clear")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QKeySequenceEdit) ConnectEditingFinished(f func()) {
	defer qt.Recovering("connect QKeySequenceEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectEditingFinished() {
	defer qt.Recovering("disconnect QKeySequenceEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQKeySequenceEditEditingFinished
func callbackQKeySequenceEditEditingFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QKeySequenceEdit::editingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "editingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QKeySequenceEdit) EditingFinished() {
	defer qt.Recovering("QKeySequenceEdit::editingFinished")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_EditingFinished(ptr.Pointer())
	}
}

func (ptr *QKeySequenceEdit) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QKeySequenceEdit::event")

	if ptr.Pointer() != nil {
		return C.QKeySequenceEdit_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QKeySequenceEdit) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQKeySequenceEditKeyPressEvent
func callbackQKeySequenceEditKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQKeySequenceEditFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QKeySequenceEdit) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QKeySequenceEdit) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QKeySequenceEdit) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQKeySequenceEditKeyReleaseEvent
func callbackQKeySequenceEditKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQKeySequenceEditFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QKeySequenceEdit) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QKeySequenceEdit) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QKeySequenceEdit) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQKeySequenceEditTimerEvent
func callbackQKeySequenceEditTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQKeySequenceEditFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QKeySequenceEdit) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QKeySequenceEdit) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QKeySequenceEdit) DestroyQKeySequenceEdit() {
	defer qt.Recovering("QKeySequenceEdit::~QKeySequenceEdit")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DestroyQKeySequenceEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QKeySequenceEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQKeySequenceEditActionEvent
func callbackQKeySequenceEditActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQKeySequenceEditDragEnterEvent
func callbackQKeySequenceEditDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QKeySequenceEdit) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQKeySequenceEditDragLeaveEvent
func callbackQKeySequenceEditDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QKeySequenceEdit) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQKeySequenceEditDragMoveEvent
func callbackQKeySequenceEditDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QKeySequenceEdit) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQKeySequenceEditDropEvent
func callbackQKeySequenceEditDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QKeySequenceEdit) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQKeySequenceEditEnterEvent
func callbackQKeySequenceEditEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQKeySequenceEditFocusInEvent
func callbackQKeySequenceEditFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QKeySequenceEdit) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQKeySequenceEditFocusOutEvent
func callbackQKeySequenceEditFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QKeySequenceEdit) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQKeySequenceEditHideEvent
func callbackQKeySequenceEditHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QKeySequenceEdit) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQKeySequenceEditLeaveEvent
func callbackQKeySequenceEditLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQKeySequenceEditMoveEvent
func callbackQKeySequenceEditMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QKeySequenceEdit) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQKeySequenceEditPaintEvent
func callbackQKeySequenceEditPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QKeySequenceEdit) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QKeySequenceEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QKeySequenceEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQKeySequenceEditSetVisible
func callbackQKeySequenceEditSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QKeySequenceEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QKeySequenceEdit) SetVisible(visible bool) {
	defer qt.Recovering("QKeySequenceEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QKeySequenceEdit) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QKeySequenceEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QKeySequenceEdit) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQKeySequenceEditShowEvent
func callbackQKeySequenceEditShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQKeySequenceEditChangeEvent
func callbackQKeySequenceEditChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQKeySequenceEditCloseEvent
func callbackQKeySequenceEditCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQKeySequenceEditContextMenuEvent
func callbackQKeySequenceEditContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QKeySequenceEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QKeySequenceEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQKeySequenceEditInitPainter
func callbackQKeySequenceEditInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQKeySequenceEditFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QKeySequenceEdit) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QKeySequenceEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QKeySequenceEdit) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QKeySequenceEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QKeySequenceEdit) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQKeySequenceEditInputMethodEvent
func callbackQKeySequenceEditInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QKeySequenceEdit) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQKeySequenceEditMouseDoubleClickEvent
func callbackQKeySequenceEditMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQKeySequenceEditMouseMoveEvent
func callbackQKeySequenceEditMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQKeySequenceEditMousePressEvent
func callbackQKeySequenceEditMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQKeySequenceEditMouseReleaseEvent
func callbackQKeySequenceEditMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQKeySequenceEditResizeEvent
func callbackQKeySequenceEditResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQKeySequenceEditTabletEvent
func callbackQKeySequenceEditTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QKeySequenceEdit) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQKeySequenceEditWheelEvent
func callbackQKeySequenceEditWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QKeySequenceEdit) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQKeySequenceEditChildEvent
func callbackQKeySequenceEditChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QKeySequenceEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QKeySequenceEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QKeySequenceEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QKeySequenceEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQKeySequenceEditCustomEvent
func callbackQKeySequenceEditCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeySequenceEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQKeySequenceEditFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QKeySequenceEdit) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeySequenceEdit) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QKeySequenceEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QKeySequenceEdit_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
