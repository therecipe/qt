package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
	for len(n.ObjectName()) < len("QDesktopWidget_") {
		n.SetObjectName("QDesktopWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QDesktopWidget) QDesktopWidget_PTR() *QDesktopWidget {
	return ptr
}

func (ptr *QDesktopWidget) IsVirtualDesktop() bool {
	defer qt.Recovering("QDesktopWidget::isVirtualDesktop")

	if ptr.Pointer() != nil {
		return C.QDesktopWidget_IsVirtualDesktop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesktopWidget) PrimaryScreen() int {
	defer qt.Recovering("QDesktopWidget::primaryScreen")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_PrimaryScreen(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDesktopWidgetResizeEvent
func callbackQDesktopWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) Screen(screen int) *QWidget {
	defer qt.Recovering("QDesktopWidget::screen")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDesktopWidget_Screen(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenNumber2(point core.QPoint_ITF) int {
	defer qt.Recovering("QDesktopWidget::screenNumber")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber2(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return 0
}

func (ptr *QDesktopWidget) ScreenNumber(widget QWidget_ITF) int {
	defer qt.Recovering("QDesktopWidget::screenNumber")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResized(f func(screen int)) {
	defer qt.Recovering("connect QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResized() {
	defer qt.Recovering("disconnect QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQDesktopWidgetResized
func callbackQDesktopWidgetResized(ptrName *C.char, screen C.int) {
	defer qt.Recovering("callback QDesktopWidget::resized")

	var signal = qt.GetSignal(C.GoString(ptrName), "resized")
	if signal != nil {
		signal.(func(int))(int(screen))
	}

}

func (ptr *QDesktopWidget) ScreenCount() int {
	defer qt.Recovering("QDesktopWidget::screenCount")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectScreenCountChanged(f func(newCount int)) {
	defer qt.Recovering("connect QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectScreenCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenCountChanged", f)
	}
}

func (ptr *QDesktopWidget) DisconnectScreenCountChanged() {
	defer qt.Recovering("disconnect QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectScreenCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenCountChanged")
	}
}

//export callbackQDesktopWidgetScreenCountChanged
func callbackQDesktopWidgetScreenCountChanged(ptrName *C.char, newCount C.int) {
	defer qt.Recovering("callback QDesktopWidget::screenCountChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "screenCountChanged")
	if signal != nil {
		signal.(func(int))(int(newCount))
	}

}

func (ptr *QDesktopWidget) ConnectWorkAreaResized(f func(screen int)) {
	defer qt.Recovering("connect QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectWorkAreaResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "workAreaResized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWorkAreaResized() {
	defer qt.Recovering("disconnect QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectWorkAreaResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "workAreaResized")
	}
}

//export callbackQDesktopWidgetWorkAreaResized
func callbackQDesktopWidgetWorkAreaResized(ptrName *C.char, screen C.int) {
	defer qt.Recovering("callback QDesktopWidget::workAreaResized")

	var signal = qt.GetSignal(C.GoString(ptrName), "workAreaResized")
	if signal != nil {
		signal.(func(int))(int(screen))
	}

}
