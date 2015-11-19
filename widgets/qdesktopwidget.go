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

type QDesktopWidget_ITF interface {
	QWidget_ITF
	QDesktopWidget_PTR() *QDesktopWidget
}

func PointerFromQDesktopWidget(ptr QDesktopWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesktopWidget_PTR().Pointer()
	}
	return nil
}

func NewQDesktopWidgetFromPointer(ptr unsafe.Pointer) *QDesktopWidget {
	var n = new(QDesktopWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDesktopWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDesktopWidget) QDesktopWidget_PTR() *QDesktopWidget {
	return ptr
}

func (ptr *QDesktopWidget) IsVirtualDesktop() bool {
	if ptr.Pointer() != nil {
		return C.QDesktopWidget_IsVirtualDesktop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesktopWidget) PrimaryScreen() int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_PrimaryScreen(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) Screen(screen int) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDesktopWidget_Screen(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenNumber2(point core.QPoint_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber2(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return 0
}

func (ptr *QDesktopWidget) ScreenNumber(widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResized(f func(screen int)) {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResized() {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQDesktopWidgetResized
func callbackQDesktopWidgetResized(ptrName *C.char, screen C.int) {
	qt.GetSignal(C.GoString(ptrName), "resized").(func(int))(int(screen))
}

func (ptr *QDesktopWidget) ScreenCount() int {
	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectScreenCountChanged(f func(newCount int)) {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectScreenCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenCountChanged", f)
	}
}

func (ptr *QDesktopWidget) DisconnectScreenCountChanged() {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectScreenCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenCountChanged")
	}
}

//export callbackQDesktopWidgetScreenCountChanged
func callbackQDesktopWidgetScreenCountChanged(ptrName *C.char, newCount C.int) {
	qt.GetSignal(C.GoString(ptrName), "screenCountChanged").(func(int))(int(newCount))
}

func (ptr *QDesktopWidget) ConnectWorkAreaResized(f func(screen int)) {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectWorkAreaResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "workAreaResized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWorkAreaResized() {
	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectWorkAreaResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "workAreaResized")
	}
}

//export callbackQDesktopWidgetWorkAreaResized
func callbackQDesktopWidgetWorkAreaResized(ptrName *C.char, screen C.int) {
	qt.GetSignal(C.GoString(ptrName), "workAreaResized").(func(int))(int(screen))
}
