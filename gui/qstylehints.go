package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QStyleHints_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStyleHints) QStyleHints_PTR() *QStyleHints {
	return ptr
}

func (ptr *QStyleHints) CursorFlashTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::cursorFlashTime")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_CursorFlashTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) FontSmoothingGamma() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::fontSmoothingGamma")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QStyleHints_FontSmoothingGamma(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) KeyboardAutoRepeatRate() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::keyboardAutoRepeatRate")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_KeyboardAutoRepeatRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) KeyboardInputInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::keyboardInputInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_KeyboardInputInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) MouseDoubleClickInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::mouseDoubleClickInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_MouseDoubleClickInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) MousePressAndHoldInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::mousePressAndHoldInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_MousePressAndHoldInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) PasswordMaskDelay() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::passwordMaskDelay")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_PasswordMaskDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) SetFocusOnTouchRelease() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::setFocusOnTouchRelease")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStyleHints_SetFocusOnTouchRelease(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) ShowIsFullScreen() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::showIsFullScreen")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStyleHints_ShowIsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) SingleClickActivation() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::singleClickActivation")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStyleHints_SingleClickActivation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) StartDragDistance() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragDistance")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragDistance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) StartDragTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragTime")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) StartDragVelocity() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragVelocity")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStyleHints_StartDragVelocity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) TabFocusBehavior() core.Qt__TabFocusBehavior {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::tabFocusBehavior")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TabFocusBehavior(C.QStyleHints_TabFocusBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleHints) UseRtlExtensions() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::useRtlExtensions")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStyleHints_UseRtlExtensions(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStyleHints) ConnectCursorFlashTimeChanged(f func(cursorFlashTime int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::cursorFlashTimeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectCursorFlashTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorFlashTimeChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectCursorFlashTimeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::cursorFlashTimeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectCursorFlashTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorFlashTimeChanged")
	}
}

//export callbackQStyleHintsCursorFlashTimeChanged
func callbackQStyleHintsCursorFlashTimeChanged(ptrName *C.char, cursorFlashTime C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::cursorFlashTimeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cursorFlashTimeChanged").(func(int))(int(cursorFlashTime))
}

func (ptr *QStyleHints) ConnectKeyboardInputIntervalChanged(f func(keyboardInputInterval int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::keyboardInputIntervalChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectKeyboardInputIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "keyboardInputIntervalChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectKeyboardInputIntervalChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::keyboardInputIntervalChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectKeyboardInputIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "keyboardInputIntervalChanged")
	}
}

//export callbackQStyleHintsKeyboardInputIntervalChanged
func callbackQStyleHintsKeyboardInputIntervalChanged(ptrName *C.char, keyboardInputInterval C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::keyboardInputIntervalChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "keyboardInputIntervalChanged").(func(int))(int(keyboardInputInterval))
}

func (ptr *QStyleHints) ConnectMouseDoubleClickIntervalChanged(f func(mouseDoubleClickInterval int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::mouseDoubleClickIntervalChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectMouseDoubleClickIntervalChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickIntervalChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectMouseDoubleClickIntervalChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::mouseDoubleClickIntervalChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectMouseDoubleClickIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickIntervalChanged")
	}
}

//export callbackQStyleHintsMouseDoubleClickIntervalChanged
func callbackQStyleHintsMouseDoubleClickIntervalChanged(ptrName *C.char, mouseDoubleClickInterval C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::mouseDoubleClickIntervalChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickIntervalChanged").(func(int))(int(mouseDoubleClickInterval))
}

func (ptr *QStyleHints) ConnectStartDragDistanceChanged(f func(startDragDistance int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragDistanceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectStartDragDistanceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "startDragDistanceChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectStartDragDistanceChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragDistanceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectStartDragDistanceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "startDragDistanceChanged")
	}
}

//export callbackQStyleHintsStartDragDistanceChanged
func callbackQStyleHintsStartDragDistanceChanged(ptrName *C.char, startDragDistance C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragDistanceChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "startDragDistanceChanged").(func(int))(int(startDragDistance))
}

func (ptr *QStyleHints) ConnectStartDragTimeChanged(f func(startDragTime int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragTimeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_ConnectStartDragTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "startDragTimeChanged", f)
	}
}

func (ptr *QStyleHints) DisconnectStartDragTimeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragTimeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleHints_DisconnectStartDragTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "startDragTimeChanged")
	}
}

//export callbackQStyleHintsStartDragTimeChanged
func callbackQStyleHintsStartDragTimeChanged(ptrName *C.char, startDragTime C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleHints::startDragTimeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "startDragTimeChanged").(func(int))(int(startDragTime))
}
