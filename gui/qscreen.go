package gui

//#include "qscreen.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScreen struct {
	core.QObject
}

type QScreenITF interface {
	core.QObjectITF
	QScreenPTR() *QScreen
}

func PointerFromQScreen(ptr QScreenITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScreenPTR().Pointer()
	}
	return nil
}

func QScreenFromPointer(ptr unsafe.Pointer) *QScreen {
	var n = new(QScreen)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScreen_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScreen) QScreenPTR() *QScreen {
	return ptr
}

func (ptr *QScreen) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QScreen_Depth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScreen) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScreen_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScreen) NativeOrientation() core.Qt__ScreenOrientation {
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_NativeOrientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScreen) Orientation() core.Qt__ScreenOrientation {
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScreen) PrimaryOrientation() core.Qt__ScreenOrientation {
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_PrimaryOrientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScreen) AngleBetween(a core.Qt__ScreenOrientation, b core.Qt__ScreenOrientation) int {
	if ptr.Pointer() != nil {
		return int(C.QScreen_AngleBetween(C.QtObjectPtr(ptr.Pointer()), C.int(a), C.int(b)))
	}
	return 0
}

func (ptr *QScreen) IsLandscape(o core.Qt__ScreenOrientation) bool {
	if ptr.Pointer() != nil {
		return C.QScreen_IsLandscape(C.QtObjectPtr(ptr.Pointer()), C.int(o)) != 0
	}
	return false
}

func (ptr *QScreen) IsPortrait(o core.Qt__ScreenOrientation) bool {
	if ptr.Pointer() != nil {
		return C.QScreen_IsPortrait(C.QtObjectPtr(ptr.Pointer()), C.int(o)) != 0
	}
	return false
}

func (ptr *QScreen) ConnectOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	if ptr.Pointer() != nil {
		C.QScreen_ConnectOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "orientationChanged", f)
	}
}

func (ptr *QScreen) DisconnectOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QScreen_DisconnectOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "orientationChanged")
	}
}

//export callbackQScreenOrientationChanged
func callbackQScreenOrientationChanged(ptrName *C.char, orientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "orientationChanged").(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
}

func (ptr *QScreen) OrientationUpdateMask() core.Qt__ScreenOrientation {
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_OrientationUpdateMask(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScreen) ConnectPrimaryOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	if ptr.Pointer() != nil {
		C.QScreen_ConnectPrimaryOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "primaryOrientationChanged", f)
	}
}

func (ptr *QScreen) DisconnectPrimaryOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QScreen_DisconnectPrimaryOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "primaryOrientationChanged")
	}
}

//export callbackQScreenPrimaryOrientationChanged
func callbackQScreenPrimaryOrientationChanged(ptrName *C.char, orientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "primaryOrientationChanged").(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
}

func (ptr *QScreen) SetOrientationUpdateMask(mask core.Qt__ScreenOrientation) {
	if ptr.Pointer() != nil {
		C.QScreen_SetOrientationUpdateMask(C.QtObjectPtr(ptr.Pointer()), C.int(mask))
	}
}

func (ptr *QScreen) DestroyQScreen() {
	if ptr.Pointer() != nil {
		C.QScreen_DestroyQScreen(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
