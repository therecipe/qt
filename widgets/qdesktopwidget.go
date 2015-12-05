package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QDesktopWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDesktopWidget) QDesktopWidget_PTR() *QDesktopWidget {
	return ptr
}

func (ptr *QDesktopWidget) IsVirtualDesktop() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::isVirtualDesktop")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDesktopWidget_IsVirtualDesktop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesktopWidget) PrimaryScreen() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::primaryScreen")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_PrimaryScreen(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) Screen(screen int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screen")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDesktopWidget_Screen(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenNumber2(point core.QPoint_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screenNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber2(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return 0
}

func (ptr *QDesktopWidget) ScreenNumber(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screenNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResized(f func(screen int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::resized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResized() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::resized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQDesktopWidgetResized
func callbackQDesktopWidgetResized(ptrName *C.char, screen C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::resized")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "resized").(func(int))(int(screen))
}

func (ptr *QDesktopWidget) ScreenCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screenCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectScreenCountChanged(f func(newCount int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screenCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectScreenCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenCountChanged", f)
	}
}

func (ptr *QDesktopWidget) DisconnectScreenCountChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screenCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectScreenCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenCountChanged")
	}
}

//export callbackQDesktopWidgetScreenCountChanged
func callbackQDesktopWidgetScreenCountChanged(ptrName *C.char, newCount C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::screenCountChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "screenCountChanged").(func(int))(int(newCount))
}

func (ptr *QDesktopWidget) ConnectWorkAreaResized(f func(screen int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::workAreaResized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectWorkAreaResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "workAreaResized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWorkAreaResized() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::workAreaResized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectWorkAreaResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "workAreaResized")
	}
}

//export callbackQDesktopWidgetWorkAreaResized
func callbackQDesktopWidgetWorkAreaResized(ptrName *C.char, screen C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDesktopWidget::workAreaResized")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "workAreaResized").(func(int))(int(screen))
}
