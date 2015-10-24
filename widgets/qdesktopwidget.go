package widgets

//#include "qdesktopwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDesktopWidget struct {
	QWidget
}

type QDesktopWidgetITF interface {
	QWidgetITF
	QDesktopWidgetPTR() *QDesktopWidget
}

func PointerFromQDesktopWidget(ptr QDesktopWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesktopWidgetPTR().Pointer()
	}
	return nil
}

func QDesktopWidgetFromPointer(ptr unsafe.Pointer) *QDesktopWidget {
	var n = new(QDesktopWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDesktopWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDesktopWidget) QDesktopWidgetPTR() *QDesktopWidget {
	return ptr
}

func (ptr *QDesktopWidget) IsVirtualDesktop() bool {
	if ptr.Pointer() != nil {
		return C.QDesktopWidget_IsVirtualDesktop(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDesktopWidget) PrimaryScreen() int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_PrimaryScreen(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDesktopWidget) Screen(screen int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QDesktopWidget_Screen(C.QtObjectPtr(ptr.Pointer()), C.int(screen))))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenNumber2(point core.QPointITF) int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point))))
	}
	return 0
}

func (ptr *QDesktopWidget) ScreenNumber(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResized(f func(screen int)) {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectResized(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResized() {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectResized(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQDesktopWidgetResized
func callbackQDesktopWidgetResized(ptrName *C.char, screen C.int) {
	qt.GetSignal(C.GoString(ptrName), "resized").(func(int))(int(screen))
}

func (ptr *QDesktopWidget) ScreenCount() int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectScreenCountChanged(f func(newCount int)) {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectScreenCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "screenCountChanged", f)
	}
}

func (ptr *QDesktopWidget) DisconnectScreenCountChanged() {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectScreenCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "screenCountChanged")
	}
}

//export callbackQDesktopWidgetScreenCountChanged
func callbackQDesktopWidgetScreenCountChanged(ptrName *C.char, newCount C.int) {
	qt.GetSignal(C.GoString(ptrName), "screenCountChanged").(func(int))(int(newCount))
}

func (ptr *QDesktopWidget) ConnectWorkAreaResized(f func(screen int)) {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectWorkAreaResized(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "workAreaResized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWorkAreaResized() {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectWorkAreaResized(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "workAreaResized")
	}
}

//export callbackQDesktopWidgetWorkAreaResized
func callbackQDesktopWidgetWorkAreaResized(ptrName *C.char, screen C.int) {
	qt.GetSignal(C.GoString(ptrName), "workAreaResized").(func(int))(int(screen))
}
