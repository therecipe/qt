package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStyleHints struct {
	core.QObject
}

type QStyleHints_ITF interface {
	core.QObject_ITF
	QStyleHints_PTR() *QStyleHints
}

func PointerFromQStyleHints(ptr QStyleHints_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHints_PTR().Pointer()
	}
	return nil
}

func NewQStyleHintsFromPointer(ptr unsafe.Pointer) *QStyleHints {
	var n = new(QStyleHints)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStyleHints_") {
		n.SetObjectName("QStyleHints_" + qt.Identifier())
	}
	return n
}

func (ptr *QStyleHints) QStyleHints_PTR() *QStyleHints {
	return ptr
}

func (ptr *QStyleHints) CursorFlashTime() int {
	defer qt.Recovering("QStyleHints::cursorFlashTime")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_CursorFlashTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) FontSmoothingGamma() float64 {
	defer qt.Recovering("QStyleHints::fontSmoothingGamma")

	if ptr.Pointer() != nil {
		return float64(C.QStyleHints_FontSmoothingGamma(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) KeyboardAutoRepeatRate() int {
	defer qt.Recovering("QStyleHints::keyboardAutoRepeatRate")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_KeyboardAutoRepeatRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) KeyboardInputInterval() int {
	defer qt.Recovering("QStyleHints::keyboardInputInterval")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_KeyboardInputInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) MouseDoubleClickInterval() int {
	defer qt.Recovering("QStyleHints::mouseDoubleClickInterval")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_MouseDoubleClickInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) MousePressAndHoldInterval() int {
	defer qt.Recovering("QStyleHints::mousePressAndHoldInterval")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_MousePressAndHoldInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) PasswordMaskDelay() int {
	defer qt.Recovering("QStyleHints::passwordMaskDelay")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_PasswordMaskDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) SetFocusOnTouchRelease() bool {
	defer qt.Recovering("QStyleHints::setFocusOnTouchRelease")

	if ptr.Pointer() != nil {
		return C.QStyleHints_SetFocusOnTouchRelease(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) ShowIsFullScreen() bool {
	defer qt.Recovering("QStyleHints::showIsFullScreen")

	if ptr.Pointer() != nil {
		return C.QStyleHints_ShowIsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) SingleClickActivation() bool {
	defer qt.Recovering("QStyleHints::singleClickActivation")

	if ptr.Pointer() != nil {
		return C.QStyleHints_SingleClickActivation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) StartDragDistance() int {
	defer qt.Recovering("QStyleHints::startDragDistance")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragDistance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) StartDragTime() int {
	defer qt.Recovering("QStyleHints::startDragTime")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) StartDragVelocity() int {
	defer qt.Recovering("QStyleHints::startDragVelocity")

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragVelocity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) TabFocusBehavior() core.Qt__TabFocusBehavior {
	defer qt.Recovering("QStyleHints::tabFocusBehavior")

	if ptr.Pointer() != nil {
		return core.Qt__TabFocusBehavior(C.QStyleHints_TabFocusBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) UseRtlExtensions() bool {
	defer qt.Recovering("QStyleHints::useRtlExtensions")

	if ptr.Pointer() != nil {
		return C.QStyleHints_UseRtlExtensions(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) ConnectCursorFlashTimeChanged(f func(cursorFlashTime int)) {
	defer qt.Recovering("connect QStyleHints::cursorFlashTimeChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectCursorFlashTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorFlashTimeChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectCursorFlashTimeChanged() {
	defer qt.Recovering("disconnect QStyleHints::cursorFlashTimeChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectCursorFlashTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorFlashTimeChanged")
	}
}

//export callbackQStyleHintsCursorFlashTimeChanged
func callbackQStyleHintsCursorFlashTimeChanged(ptrName *C.char, cursorFlashTime C.int) {
	defer qt.Recovering("callback QStyleHints::cursorFlashTimeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cursorFlashTimeChanged"); signal != nil {
		signal.(func(int))(int(cursorFlashTime))
	}

}

func (ptr *QStyleHints) ConnectKeyboardInputIntervalChanged(f func(keyboardInputInterval int)) {
	defer qt.Recovering("connect QStyleHints::keyboardInputIntervalChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectKeyboardInputIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "keyboardInputIntervalChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectKeyboardInputIntervalChanged() {
	defer qt.Recovering("disconnect QStyleHints::keyboardInputIntervalChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectKeyboardInputIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "keyboardInputIntervalChanged")
	}
}

//export callbackQStyleHintsKeyboardInputIntervalChanged
func callbackQStyleHintsKeyboardInputIntervalChanged(ptrName *C.char, keyboardInputInterval C.int) {
	defer qt.Recovering("callback QStyleHints::keyboardInputIntervalChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardInputIntervalChanged"); signal != nil {
		signal.(func(int))(int(keyboardInputInterval))
	}

}

func (ptr *QStyleHints) ConnectMouseDoubleClickIntervalChanged(f func(mouseDoubleClickInterval int)) {
	defer qt.Recovering("connect QStyleHints::mouseDoubleClickIntervalChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectMouseDoubleClickIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickIntervalChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectMouseDoubleClickIntervalChanged() {
	defer qt.Recovering("disconnect QStyleHints::mouseDoubleClickIntervalChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectMouseDoubleClickIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickIntervalChanged")
	}
}

//export callbackQStyleHintsMouseDoubleClickIntervalChanged
func callbackQStyleHintsMouseDoubleClickIntervalChanged(ptrName *C.char, mouseDoubleClickInterval C.int) {
	defer qt.Recovering("callback QStyleHints::mouseDoubleClickIntervalChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickIntervalChanged"); signal != nil {
		signal.(func(int))(int(mouseDoubleClickInterval))
	}

}

func (ptr *QStyleHints) ConnectStartDragDistanceChanged(f func(startDragDistance int)) {
	defer qt.Recovering("connect QStyleHints::startDragDistanceChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectStartDragDistanceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "startDragDistanceChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectStartDragDistanceChanged() {
	defer qt.Recovering("disconnect QStyleHints::startDragDistanceChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectStartDragDistanceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "startDragDistanceChanged")
	}
}

//export callbackQStyleHintsStartDragDistanceChanged
func callbackQStyleHintsStartDragDistanceChanged(ptrName *C.char, startDragDistance C.int) {
	defer qt.Recovering("callback QStyleHints::startDragDistanceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDragDistanceChanged"); signal != nil {
		signal.(func(int))(int(startDragDistance))
	}

}

func (ptr *QStyleHints) ConnectStartDragTimeChanged(f func(startDragTime int)) {
	defer qt.Recovering("connect QStyleHints::startDragTimeChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectStartDragTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "startDragTimeChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectStartDragTimeChanged() {
	defer qt.Recovering("disconnect QStyleHints::startDragTimeChanged")

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectStartDragTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "startDragTimeChanged")
	}
}

//export callbackQStyleHintsStartDragTimeChanged
func callbackQStyleHintsStartDragTimeChanged(ptrName *C.char, startDragTime C.int) {
	defer qt.Recovering("callback QStyleHints::startDragTimeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDragTimeChanged"); signal != nil {
		signal.(func(int))(int(startDragTime))
	}

}

func (ptr *QStyleHints) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStyleHints::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStyleHints) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStyleHints::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStyleHintsTimerEvent
func callbackQStyleHintsTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyleHints::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStyleHints) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStyleHints::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStyleHints) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStyleHints::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStyleHintsChildEvent
func callbackQStyleHintsChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyleHints::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStyleHints) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStyleHints::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStyleHints) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStyleHints::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStyleHintsCustomEvent
func callbackQStyleHintsCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyleHints::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
