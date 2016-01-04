package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QTimeEdit struct {
	QDateTimeEdit
}

type QTimeEdit_ITF interface {
	QDateTimeEdit_ITF
	QTimeEdit_PTR() *QTimeEdit
}

func PointerFromQTimeEdit(ptr QTimeEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeEdit_PTR().Pointer()
	}
	return nil
}

func NewQTimeEditFromPointer(ptr unsafe.Pointer) *QTimeEdit {
	var n = new(QTimeEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTimeEdit_") {
		n.SetObjectName("QTimeEdit_" + qt.Identifier())
	}
	return n
}

func (ptr *QTimeEdit) QTimeEdit_PTR() *QTimeEdit {
	return ptr
}

func NewQTimeEdit(parent QWidget_ITF) *QTimeEdit {
	defer qt.Recovering("QTimeEdit::QTimeEdit")

	return NewQTimeEditFromPointer(C.QTimeEdit_NewQTimeEdit(PointerFromQWidget(parent)))
}

func NewQTimeEdit2(time core.QTime_ITF, parent QWidget_ITF) *QTimeEdit {
	defer qt.Recovering("QTimeEdit::QTimeEdit")

	return NewQTimeEditFromPointer(C.QTimeEdit_NewQTimeEdit2(core.PointerFromQTime(time), PointerFromQWidget(parent)))
}

func (ptr *QTimeEdit) DestroyQTimeEdit() {
	defer qt.Recovering("QTimeEdit::~QTimeEdit")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DestroyQTimeEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTimeEdit) ConnectClear(f func()) {
	defer qt.Recovering("connect QTimeEdit::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QTimeEdit) DisconnectClear() {
	defer qt.Recovering("disconnect QTimeEdit::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQTimeEditClear
func callbackQTimeEditClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTimeEdit::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQTimeEditFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QTimeEdit) Clear() {
	defer qt.Recovering("QTimeEdit::clear")

	if ptr.Pointer() != nil {
		C.QTimeEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QTimeEdit) ClearDefault() {
	defer qt.Recovering("QTimeEdit::clear")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QTimeEdit) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTimeEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTimeEdit::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTimeEditFocusInEvent
func callbackQTimeEditFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTimeEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTimeEdit) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTimeEdit::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTimeEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTimeEdit::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTimeEditKeyPressEvent
func callbackQTimeEditKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTimeEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTimeEdit) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTimeEdit::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTimeEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTimeEdit::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTimeEditMousePressEvent
func callbackQTimeEditMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTimeEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTimeEdit::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTimeEditPaintEvent
func callbackQTimeEditPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTimeEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QTimeEdit) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTimeEdit::paintEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectStepBy(f func(steps int)) {
	defer qt.Recovering("connect QTimeEdit::stepBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stepBy", f)
	}
}

func (ptr *QTimeEdit) DisconnectStepBy() {
	defer qt.Recovering("disconnect QTimeEdit::stepBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stepBy")
	}
}

//export callbackQTimeEditStepBy
func callbackQTimeEditStepBy(ptr unsafe.Pointer, ptrName *C.char, steps C.int) {
	defer qt.Recovering("callback QTimeEdit::stepBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "stepBy"); signal != nil {
		signal.(func(int))(int(steps))
	} else {
		NewQTimeEditFromPointer(ptr).StepByDefault(int(steps))
	}
}

func (ptr *QTimeEdit) StepBy(steps int) {
	defer qt.Recovering("QTimeEdit::stepBy")

	if ptr.Pointer() != nil {
		C.QTimeEdit_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QTimeEdit) StepByDefault(steps int) {
	defer qt.Recovering("QTimeEdit::stepBy")

	if ptr.Pointer() != nil {
		C.QTimeEdit_StepByDefault(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QTimeEdit) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTimeEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTimeEdit::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTimeEditWheelEvent
func callbackQTimeEditWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTimeEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QTimeEdit) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTimeEdit::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTimeEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTimeEdit::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTimeEditChangeEvent
func callbackQTimeEditChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::changeEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTimeEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTimeEdit::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTimeEditCloseEvent
func callbackQTimeEditCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTimeEdit) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::closeEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTimeEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTimeEdit::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTimeEditContextMenuEvent
func callbackQTimeEditContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTimeEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTimeEdit) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTimeEdit::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTimeEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTimeEdit::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTimeEditFocusOutEvent
func callbackQTimeEditFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTimeEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTimeEdit) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTimeEdit::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTimeEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTimeEdit::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTimeEditHideEvent
func callbackQTimeEditHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTimeEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTimeEdit) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTimeEdit::hideEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTimeEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTimeEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTimeEditKeyReleaseEvent
func callbackQTimeEditKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTimeEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTimeEdit) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTimeEdit::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTimeEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTimeEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTimeEditMouseMoveEvent
func callbackQTimeEditMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTimeEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTimeEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTimeEditMouseReleaseEvent
func callbackQTimeEditMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTimeEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTimeEdit::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTimeEditResizeEvent
func callbackQTimeEditResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTimeEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QTimeEdit) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTimeEdit::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QTimeEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTimeEdit::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTimeEditShowEvent
func callbackQTimeEditShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QTimeEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QTimeEdit) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QTimeEdit::showEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTimeEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTimeEdit::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTimeEditTimerEvent
func callbackQTimeEditTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTimeEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTimeEdit) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTimeEdit::timerEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTimeEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTimeEdit::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTimeEditActionEvent
func callbackQTimeEditActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTimeEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTimeEdit) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTimeEdit::actionEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTimeEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTimeEdit::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTimeEditDragEnterEvent
func callbackQTimeEditDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QTimeEdit) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTimeEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTimeEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTimeEditDragLeaveEvent
func callbackQTimeEditDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QTimeEdit) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTimeEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTimeEdit::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTimeEditDragMoveEvent
func callbackQTimeEditDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QTimeEdit) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QTimeEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTimeEdit::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTimeEditDropEvent
func callbackQTimeEditDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QTimeEdit) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QTimeEdit::dropEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTimeEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTimeEdit::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTimeEditEnterEvent
func callbackQTimeEditEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::enterEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTimeEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTimeEdit::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTimeEditLeaveEvent
func callbackQTimeEditLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTimeEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTimeEdit::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTimeEditMoveEvent
func callbackQTimeEditMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTimeEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTimeEdit) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTimeEdit::moveEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTimeEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTimeEdit) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTimeEdit::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTimeEditSetVisible
func callbackQTimeEditSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTimeEdit::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTimeEdit) SetVisible(visible bool) {
	defer qt.Recovering("QTimeEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QTimeEdit_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTimeEdit) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QTimeEdit::setVisible")

	if ptr.Pointer() != nil {
		C.QTimeEdit_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTimeEdit) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTimeEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTimeEdit) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTimeEdit::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTimeEditInitPainter
func callbackQTimeEditInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQTimeEditFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QTimeEdit) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTimeEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QTimeEdit_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTimeEdit) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTimeEdit::initPainter")

	if ptr.Pointer() != nil {
		C.QTimeEdit_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTimeEdit) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTimeEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTimeEdit::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTimeEditInputMethodEvent
func callbackQTimeEditInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTimeEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QTimeEdit) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTimeEdit::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTimeEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTimeEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTimeEditMouseDoubleClickEvent
func callbackQTimeEditMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTimeEdit::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTimeEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTimeEdit::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTimeEditTabletEvent
func callbackQTimeEditTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTimeEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTimeEdit) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTimeEdit::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTimeEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTimeEdit::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTimeEditChildEvent
func callbackQTimeEditChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTimeEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTimeEdit) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTimeEdit::childEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTimeEdit) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTimeEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTimeEdit) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTimeEdit::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTimeEditCustomEvent
func callbackQTimeEditCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTimeEdit::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTimeEditFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTimeEdit) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTimeEdit) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTimeEdit::customEvent")

	if ptr.Pointer() != nil {
		C.QTimeEdit_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
