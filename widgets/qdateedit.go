package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDateEdit struct {
	QDateTimeEdit
}

type QDateEdit_ITF interface {
	QDateTimeEdit_ITF
	QDateEdit_PTR() *QDateEdit
}

func PointerFromQDateEdit(ptr QDateEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateEdit_PTR().Pointer()
	}
	return nil
}

func NewQDateEditFromPointer(ptr unsafe.Pointer) *QDateEdit {
	var n = new(QDateEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDateEdit_") {
		n.SetObjectName("QDateEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QDateEdit) QDateEdit_PTR() *QDateEdit {
	return ptr
}

func NewQDateEdit(parent QWidget_ITF) *QDateEdit {
	defer qt.Recovering("QDateEdit::QDateEdit")

	return NewQDateEditFromPointer(C.QDateEdit_NewQDateEdit(PointerFromQWidget(parent)))
}

func NewQDateEdit2(date core.QDate_ITF, parent QWidget_ITF) *QDateEdit {
	defer qt.Recovering("QDateEdit::QDateEdit")

	return NewQDateEditFromPointer(C.QDateEdit_NewQDateEdit2(core.PointerFromQDate(date), PointerFromQWidget(parent)))
}

func (ptr *QDateEdit) DestroyQDateEdit() {
	defer qt.Recovering("QDateEdit::~QDateEdit")

	if ptr.Pointer() != nil {
		C.QDateEdit_DestroyQDateEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDateEdit) ConnectClear(f func()) {
	defer qt.Recovering("connect QDateEdit::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QDateEdit) DisconnectClear() {
	defer qt.Recovering("disconnect QDateEdit::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQDateEditClear
func callbackQDateEditClear(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QDateEdit::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QDateEdit) Clear() {
	defer qt.Recovering("QDateEdit::clear")

	if ptr.Pointer() != nil {
		C.QDateEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QDateEdit) ClearDefault() {
	defer qt.Recovering("QDateEdit::clear")

	if ptr.Pointer() != nil {
		C.QDateEdit_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QDateEdit) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDateEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDateEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDateEditFocusInEvent
func callbackQDateEditFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDateEdit) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDateEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDateEdit) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDateEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDateEdit) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDateEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDateEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDateEditKeyPressEvent
func callbackQDateEditKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDateEdit) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDateEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDateEdit) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDateEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDateEdit) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDateEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDateEditMousePressEvent
func callbackQDateEditMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDateEdit) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDateEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDateEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDateEditPaintEvent
func callbackQDateEditPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDateEdit) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDateEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDateEdit) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDateEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDateEdit) ConnectStepBy(f func(steps int)) {
	defer qt.Recovering("connect QDateEdit::stepBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stepBy", f)
	}
}

func (ptr *QDateEdit) DisconnectStepBy() {
	defer qt.Recovering("disconnect QDateEdit::stepBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stepBy")
	}
}

//export callbackQDateEditStepBy
func callbackQDateEditStepBy(ptr unsafe.Pointer, ptrName *C.char, steps C.int) {
	defer qt.Recovering("callback QDateEdit::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
	} else {
		NewQDateEditFromPointer(ptr).StepByDefault(int(steps))
	}
}

func (ptr *QDateEdit) StepBy(steps int) {
	defer qt.Recovering("QDateEdit::stepBy")

	if ptr.Pointer() != nil {
		C.QDateEdit_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QDateEdit) StepByDefault(steps int) {
	defer qt.Recovering("QDateEdit::stepBy")

	if ptr.Pointer() != nil {
		C.QDateEdit_StepByDefault(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QDateEdit) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDateEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDateEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDateEditWheelEvent
func callbackQDateEditWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDateEdit) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDateEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDateEdit) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDateEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDateEdit) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDateEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDateEditChangeEvent
func callbackQDateEditChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDateEdit) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDateEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDateEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDateEditCloseEvent
func callbackQDateEditCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDateEdit) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDateEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDateEdit) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDateEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDateEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDateEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDateEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDateEditContextMenuEvent
func callbackQDateEditContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDateEdit) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDateEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDateEdit) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDateEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDateEdit) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDateEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDateEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDateEditFocusOutEvent
func callbackQDateEditFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDateEdit) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDateEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDateEdit) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDateEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDateEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDateEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDateEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDateEditHideEvent
func callbackQDateEditHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDateEdit) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDateEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDateEdit) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDateEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDateEdit) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDateEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDateEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDateEditKeyReleaseEvent
func callbackQDateEditKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDateEdit) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDateEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDateEdit) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDateEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDateEdit) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDateEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDateEditMouseMoveEvent
func callbackQDateEditMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDateEdit) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDateEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDateEditMouseReleaseEvent
func callbackQDateEditMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDateEdit) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDateEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDateEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDateEditResizeEvent
func callbackQDateEditResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDateEdit) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDateEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDateEdit) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDateEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDateEdit) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDateEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDateEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDateEditShowEvent
func callbackQDateEditShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDateEdit) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDateEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDateEdit) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDateEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDateEdit) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDateEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDateEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDateEditTimerEvent
func callbackQDateEditTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDateEdit) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDateEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDateEdit) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDateEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDateEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDateEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDateEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDateEditActionEvent
func callbackQDateEditActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDateEdit) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDateEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDateEdit) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDateEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDateEdit) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDateEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDateEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDateEditDragEnterEvent
func callbackQDateEditDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDateEdit) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDateEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDateEdit) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDateEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDateEdit) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDateEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDateEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDateEditDragLeaveEvent
func callbackQDateEditDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDateEdit) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDateEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDateEdit) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDateEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDateEdit) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDateEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDateEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDateEditDragMoveEvent
func callbackQDateEditDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDateEdit) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDateEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDateEdit) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDateEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDateEdit) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDateEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDateEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDateEditDropEvent
func callbackQDateEditDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDateEdit) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDateEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDateEdit) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDateEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDateEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDateEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDateEditEnterEvent
func callbackQDateEditEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDateEdit) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDateEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDateEditLeaveEvent
func callbackQDateEditLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDateEdit) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDateEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDateEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDateEditMoveEvent
func callbackQDateEditMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDateEdit) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDateEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDateEdit) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDateEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDateEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDateEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDateEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDateEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDateEditSetVisible
func callbackQDateEditSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDateEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDateEdit) SetVisible(visible bool) {
	defer qt.Recovering("QDateEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QDateEdit_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDateEdit) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDateEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QDateEdit_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDateEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDateEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDateEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDateEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDateEditInitPainter
func callbackQDateEditInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQDateEditFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QDateEdit) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDateEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QDateEdit_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDateEdit) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDateEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QDateEdit_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDateEdit) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDateEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDateEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDateEditInputMethodEvent
func callbackQDateEditInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDateEdit) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDateEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDateEdit) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDateEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDateEdit) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDateEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDateEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDateEditMouseDoubleClickEvent
func callbackQDateEditMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDateEdit) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDateEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDateEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDateEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDateEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDateEditTabletEvent
func callbackQDateEditTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDateEdit) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDateEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDateEdit) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDateEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDateEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDateEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDateEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDateEditChildEvent
func callbackQDateEditChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDateEdit) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDateEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDateEdit) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDateEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDateEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDateEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDateEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDateEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDateEditCustomEvent
func callbackQDateEditCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDateEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDateEditFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDateEdit) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDateEdit) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDateEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QDateEdit_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
