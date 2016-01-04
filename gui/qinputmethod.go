package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QInputMethod struct {
	core.QObject
}

type QInputMethod_ITF interface {
	core.QObject_ITF
	QInputMethod_PTR() *QInputMethod
}

func PointerFromQInputMethod(ptr QInputMethod_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethod_PTR().Pointer()
	}
	return nil
}

func NewQInputMethodFromPointer(ptr unsafe.Pointer) *QInputMethod {
	var n = new(QInputMethod)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QInputMethod_") {
		n.SetObjectName("QInputMethod_" + qt.Identifier())
	}
	return n
}

func (ptr *QInputMethod) QInputMethod_PTR() *QInputMethod {
	return ptr
}

//QInputMethod::Action
type QInputMethod__Action int64

const (
	QInputMethod__Click       = QInputMethod__Action(0)
	QInputMethod__ContextMenu = QInputMethod__Action(1)
)

func (ptr *QInputMethod) InputDirection() core.Qt__LayoutDirection {
	defer qt.Recovering("QInputMethod::inputDirection")

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QInputMethod_InputDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputMethod) IsAnimating() bool {
	defer qt.Recovering("QInputMethod::isAnimating")

	if ptr.Pointer() != nil {
		return C.QInputMethod_IsAnimating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QInputMethod) IsVisible() bool {
	defer qt.Recovering("QInputMethod::isVisible")

	if ptr.Pointer() != nil {
		return C.QInputMethod_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QInputMethod) ConnectAnimatingChanged(f func()) {
	defer qt.Recovering("connect QInputMethod::animatingChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectAnimatingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "animatingChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectAnimatingChanged() {
	defer qt.Recovering("disconnect QInputMethod::animatingChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectAnimatingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "animatingChanged")
	}
}

//export callbackQInputMethodAnimatingChanged
func callbackQInputMethodAnimatingChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInputMethod::animatingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "animatingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInputMethod) AnimatingChanged() {
	defer qt.Recovering("QInputMethod::animatingChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_AnimatingChanged(ptr.Pointer())
	}
}

func (ptr *QInputMethod) Commit() {
	defer qt.Recovering("QInputMethod::commit")

	if ptr.Pointer() != nil {
		C.QInputMethod_Commit(ptr.Pointer())
	}
}

func (ptr *QInputMethod) ConnectCursorRectangleChanged(f func()) {
	defer qt.Recovering("connect QInputMethod::cursorRectangleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectCursorRectangleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorRectangleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectCursorRectangleChanged() {
	defer qt.Recovering("disconnect QInputMethod::cursorRectangleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectCursorRectangleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorRectangleChanged")
	}
}

//export callbackQInputMethodCursorRectangleChanged
func callbackQInputMethodCursorRectangleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInputMethod::cursorRectangleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorRectangleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInputMethod) CursorRectangleChanged() {
	defer qt.Recovering("QInputMethod::cursorRectangleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_CursorRectangleChanged(ptr.Pointer())
	}
}

func (ptr *QInputMethod) Hide() {
	defer qt.Recovering("QInputMethod::hide")

	if ptr.Pointer() != nil {
		C.QInputMethod_Hide(ptr.Pointer())
	}
}

func (ptr *QInputMethod) ConnectInputDirectionChanged(f func(newDirection core.Qt__LayoutDirection)) {
	defer qt.Recovering("connect QInputMethod::inputDirectionChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectInputDirectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "inputDirectionChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectInputDirectionChanged() {
	defer qt.Recovering("disconnect QInputMethod::inputDirectionChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectInputDirectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "inputDirectionChanged")
	}
}

//export callbackQInputMethodInputDirectionChanged
func callbackQInputMethodInputDirectionChanged(ptr unsafe.Pointer, ptrName *C.char, newDirection C.int) {
	defer qt.Recovering("callback QInputMethod::inputDirectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputDirectionChanged"); signal != nil {
		signal.(func(core.Qt__LayoutDirection))(core.Qt__LayoutDirection(newDirection))
	}

}

func (ptr *QInputMethod) InputDirectionChanged(newDirection core.Qt__LayoutDirection) {
	defer qt.Recovering("QInputMethod::inputDirectionChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_InputDirectionChanged(ptr.Pointer(), C.int(newDirection))
	}
}

func (ptr *QInputMethod) InvokeAction(a QInputMethod__Action, cursorPosition int) {
	defer qt.Recovering("QInputMethod::invokeAction")

	if ptr.Pointer() != nil {
		C.QInputMethod_InvokeAction(ptr.Pointer(), C.int(a), C.int(cursorPosition))
	}
}

func (ptr *QInputMethod) ConnectKeyboardRectangleChanged(f func()) {
	defer qt.Recovering("connect QInputMethod::keyboardRectangleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectKeyboardRectangleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "keyboardRectangleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectKeyboardRectangleChanged() {
	defer qt.Recovering("disconnect QInputMethod::keyboardRectangleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectKeyboardRectangleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "keyboardRectangleChanged")
	}
}

//export callbackQInputMethodKeyboardRectangleChanged
func callbackQInputMethodKeyboardRectangleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInputMethod::keyboardRectangleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardRectangleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInputMethod) KeyboardRectangleChanged() {
	defer qt.Recovering("QInputMethod::keyboardRectangleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_KeyboardRectangleChanged(ptr.Pointer())
	}
}

func (ptr *QInputMethod) ConnectLocaleChanged(f func()) {
	defer qt.Recovering("connect QInputMethod::localeChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectLocaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "localeChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectLocaleChanged() {
	defer qt.Recovering("disconnect QInputMethod::localeChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectLocaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "localeChanged")
	}
}

//export callbackQInputMethodLocaleChanged
func callbackQInputMethodLocaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInputMethod::localeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "localeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInputMethod) LocaleChanged() {
	defer qt.Recovering("QInputMethod::localeChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_LocaleChanged(ptr.Pointer())
	}
}

func QInputMethod_QueryFocusObject(query core.Qt__InputMethodQuery, argument core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QInputMethod::queryFocusObject")

	return core.NewQVariantFromPointer(C.QInputMethod_QInputMethod_QueryFocusObject(C.int(query), core.PointerFromQVariant(argument)))
}

func (ptr *QInputMethod) Reset() {
	defer qt.Recovering("QInputMethod::reset")

	if ptr.Pointer() != nil {
		C.QInputMethod_Reset(ptr.Pointer())
	}
}

func (ptr *QInputMethod) SetInputItemRectangle(rect core.QRectF_ITF) {
	defer qt.Recovering("QInputMethod::setInputItemRectangle")

	if ptr.Pointer() != nil {
		C.QInputMethod_SetInputItemRectangle(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QInputMethod) SetInputItemTransform(transform QTransform_ITF) {
	defer qt.Recovering("QInputMethod::setInputItemTransform")

	if ptr.Pointer() != nil {
		C.QInputMethod_SetInputItemTransform(ptr.Pointer(), PointerFromQTransform(transform))
	}
}

func (ptr *QInputMethod) SetVisible(visible bool) {
	defer qt.Recovering("QInputMethod::setVisible")

	if ptr.Pointer() != nil {
		C.QInputMethod_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QInputMethod) Show() {
	defer qt.Recovering("QInputMethod::show")

	if ptr.Pointer() != nil {
		C.QInputMethod_Show(ptr.Pointer())
	}
}

func (ptr *QInputMethod) Update(queries core.Qt__InputMethodQuery) {
	defer qt.Recovering("QInputMethod::update")

	if ptr.Pointer() != nil {
		C.QInputMethod_Update(ptr.Pointer(), C.int(queries))
	}
}

func (ptr *QInputMethod) ConnectVisibleChanged(f func()) {
	defer qt.Recovering("connect QInputMethod::visibleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectVisibleChanged() {
	defer qt.Recovering("disconnect QInputMethod::visibleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQInputMethodVisibleChanged
func callbackQInputMethodVisibleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QInputMethod::visibleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "visibleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QInputMethod) VisibleChanged() {
	defer qt.Recovering("QInputMethod::visibleChanged")

	if ptr.Pointer() != nil {
		C.QInputMethod_VisibleChanged(ptr.Pointer())
	}
}

func (ptr *QInputMethod) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QInputMethod::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QInputMethod) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QInputMethod::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQInputMethodTimerEvent
func callbackQInputMethodTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputMethod::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQInputMethodFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QInputMethod) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInputMethod::timerEvent")

	if ptr.Pointer() != nil {
		C.QInputMethod_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInputMethod) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QInputMethod::timerEvent")

	if ptr.Pointer() != nil {
		C.QInputMethod_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QInputMethod) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QInputMethod::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QInputMethod) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QInputMethod::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQInputMethodChildEvent
func callbackQInputMethodChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputMethod::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQInputMethodFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QInputMethod) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInputMethod::childEvent")

	if ptr.Pointer() != nil {
		C.QInputMethod_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInputMethod) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QInputMethod::childEvent")

	if ptr.Pointer() != nil {
		C.QInputMethod_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QInputMethod) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QInputMethod::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QInputMethod) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QInputMethod::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQInputMethodCustomEvent
func callbackQInputMethodCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QInputMethod::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQInputMethodFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QInputMethod) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QInputMethod::customEvent")

	if ptr.Pointer() != nil {
		C.QInputMethod_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QInputMethod) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QInputMethod::customEvent")

	if ptr.Pointer() != nil {
		C.QInputMethod_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
