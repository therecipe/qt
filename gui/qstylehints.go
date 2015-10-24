package gui

//#include "qstylehints.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStyleHints struct {
	core.QObject
}

type QStyleHintsITF interface {
	core.QObjectITF
	QStyleHintsPTR() *QStyleHints
}

func PointerFromQStyleHints(ptr QStyleHintsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintsPTR().Pointer()
	}
	return nil
}

func QStyleHintsFromPointer(ptr unsafe.Pointer) *QStyleHints {
	var n = new(QStyleHints)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStyleHints_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStyleHints) QStyleHintsPTR() *QStyleHints {
	return ptr
}

func (ptr *QStyleHints) CursorFlashTime() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_CursorFlashTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) KeyboardAutoRepeatRate() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_KeyboardAutoRepeatRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) KeyboardInputInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_KeyboardInputInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) MouseDoubleClickInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_MouseDoubleClickInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) MousePressAndHoldInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_MousePressAndHoldInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) PasswordMaskDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_PasswordMaskDelay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) SetFocusOnTouchRelease() bool {
	if ptr.Pointer() != nil {
		return C.QStyleHints_SetFocusOnTouchRelease(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStyleHints) ShowIsFullScreen() bool {
	if ptr.Pointer() != nil {
		return C.QStyleHints_ShowIsFullScreen(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStyleHints) SingleClickActivation() bool {
	if ptr.Pointer() != nil {
		return C.QStyleHints_SingleClickActivation(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStyleHints) StartDragDistance() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragDistance(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) StartDragTime() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) StartDragVelocity() int {
	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragVelocity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) TabFocusBehavior() core.Qt__TabFocusBehavior {
	if ptr.Pointer() != nil {
		return core.Qt__TabFocusBehavior(C.QStyleHints_TabFocusBehavior(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHints) UseRtlExtensions() bool {
	if ptr.Pointer() != nil {
		return C.QStyleHints_UseRtlExtensions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStyleHints) ConnectCursorFlashTimeChanged(f func(cursorFlashTime int)) {
	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectCursorFlashTimeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cursorFlashTimeChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectCursorFlashTimeChanged() {
	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectCursorFlashTimeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cursorFlashTimeChanged")
	}
}

//export callbackQStyleHintsCursorFlashTimeChanged
func callbackQStyleHintsCursorFlashTimeChanged(ptrName *C.char, cursorFlashTime C.int) {
	qt.GetSignal(C.GoString(ptrName), "cursorFlashTimeChanged").(func(int))(int(cursorFlashTime))
}

func (ptr *QStyleHints) ConnectKeyboardInputIntervalChanged(f func(keyboardInputInterval int)) {
	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectKeyboardInputIntervalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "keyboardInputIntervalChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectKeyboardInputIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectKeyboardInputIntervalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "keyboardInputIntervalChanged")
	}
}

//export callbackQStyleHintsKeyboardInputIntervalChanged
func callbackQStyleHintsKeyboardInputIntervalChanged(ptrName *C.char, keyboardInputInterval C.int) {
	qt.GetSignal(C.GoString(ptrName), "keyboardInputIntervalChanged").(func(int))(int(keyboardInputInterval))
}

func (ptr *QStyleHints) ConnectMouseDoubleClickIntervalChanged(f func(mouseDoubleClickInterval int)) {
	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectMouseDoubleClickIntervalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickIntervalChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectMouseDoubleClickIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectMouseDoubleClickIntervalChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickIntervalChanged")
	}
}

//export callbackQStyleHintsMouseDoubleClickIntervalChanged
func callbackQStyleHintsMouseDoubleClickIntervalChanged(ptrName *C.char, mouseDoubleClickInterval C.int) {
	qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickIntervalChanged").(func(int))(int(mouseDoubleClickInterval))
}

func (ptr *QStyleHints) ConnectStartDragDistanceChanged(f func(startDragDistance int)) {
	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectStartDragDistanceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "startDragDistanceChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectStartDragDistanceChanged() {
	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectStartDragDistanceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "startDragDistanceChanged")
	}
}

//export callbackQStyleHintsStartDragDistanceChanged
func callbackQStyleHintsStartDragDistanceChanged(ptrName *C.char, startDragDistance C.int) {
	qt.GetSignal(C.GoString(ptrName), "startDragDistanceChanged").(func(int))(int(startDragDistance))
}

func (ptr *QStyleHints) ConnectStartDragTimeChanged(f func(startDragTime int)) {
	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectStartDragTimeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "startDragTimeChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectStartDragTimeChanged() {
	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectStartDragTimeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "startDragTimeChanged")
	}
}

//export callbackQStyleHintsStartDragTimeChanged
func callbackQStyleHintsStartDragTimeChanged(ptrName *C.char, startDragTime C.int) {
	qt.GetSignal(C.GoString(ptrName), "startDragTimeChanged").(func(int))(int(startDragTime))
}
