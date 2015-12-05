package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMdiSubWindow struct {
	QWidget
}

type QMdiSubWindow_ITF interface {
	QWidget_ITF
	QMdiSubWindow_PTR() *QMdiSubWindow
}

func PointerFromQMdiSubWindow(ptr QMdiSubWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiSubWindow_PTR().Pointer()
	}
	return nil
}

func NewQMdiSubWindowFromPointer(ptr unsafe.Pointer) *QMdiSubWindow {
	var n = new(QMdiSubWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMdiSubWindow_") {
		n.SetObjectName("QMdiSubWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMdiSubWindow) QMdiSubWindow_PTR() *QMdiSubWindow {
	return ptr
}

//QMdiSubWindow::SubWindowOption
type QMdiSubWindow__SubWindowOption int64

const (
	QMdiSubWindow__AllowOutsideAreaHorizontally = QMdiSubWindow__SubWindowOption(0x1)
	QMdiSubWindow__AllowOutsideAreaVertically   = QMdiSubWindow__SubWindowOption(0x2)
	QMdiSubWindow__RubberBandResize             = QMdiSubWindow__SubWindowOption(0x4)
	QMdiSubWindow__RubberBandMove               = QMdiSubWindow__SubWindowOption(0x8)
)

func (ptr *QMdiSubWindow) KeyboardPageStep() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::keyboardPageStep")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardPageStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiSubWindow) KeyboardSingleStep() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::keyboardSingleStep")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardSingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiSubWindow) SetKeyboardPageStep(step int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::setKeyboardPageStep")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardPageStep(ptr.Pointer(), C.int(step))
	}
}

func (ptr *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::setKeyboardSingleStep")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardSingleStep(ptr.Pointer(), C.int(step))
	}
}

func NewQMdiSubWindow(parent QWidget_ITF, flags core.Qt__WindowType) *QMdiSubWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::QMdiSubWindow")
		}
	}()

	return NewQMdiSubWindowFromPointer(C.QMdiSubWindow_NewQMdiSubWindow(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QMdiSubWindow) ConnectAboutToActivate(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::aboutToActivate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectAboutToActivate(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToActivate", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectAboutToActivate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::aboutToActivate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectAboutToActivate(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToActivate")
	}
}

//export callbackQMdiSubWindowAboutToActivate
func callbackQMdiSubWindowAboutToActivate(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::aboutToActivate")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "aboutToActivate").(func())()
}

func (ptr *QMdiSubWindow) IsShaded() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::isShaded")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_IsShaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) MdiArea() *QMdiArea {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::mdiArea")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMdiAreaFromPointer(C.QMdiSubWindow_MdiArea(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) SetOption(option QMdiSubWindow__SubWindowOption, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::setOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiSubWindow) SetSystemMenu(systemMenu QMenu_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::setSystemMenu")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetSystemMenu(ptr.Pointer(), PointerFromQMenu(systemMenu))
	}
}

func (ptr *QMdiSubWindow) SetWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMdiSubWindow) ShowShaded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::showShaded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowShaded(ptr.Pointer())
	}
}

func (ptr *QMdiSubWindow) ShowSystemMenu() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::showSystemMenu")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowSystemMenu(ptr.Pointer())
	}
}

func (ptr *QMdiSubWindow) SystemMenu() *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::systemMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMdiSubWindow_SystemMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) TestOption(option QMdiSubWindow__SubWindowOption) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::testOption")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMdiSubWindow_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) ConnectWindowStateChanged(f func(oldState core.Qt__WindowState, newState core.Qt__WindowState)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::windowStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectWindowStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectWindowStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::windowStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectWindowStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQMdiSubWindowWindowStateChanged
func callbackQMdiSubWindowWindowStateChanged(ptrName *C.char, oldState C.int, newState C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::windowStateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "windowStateChanged").(func(core.Qt__WindowState, core.Qt__WindowState))(core.Qt__WindowState(oldState), core.Qt__WindowState(newState))
}

func (ptr *QMdiSubWindow) DestroyQMdiSubWindow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiSubWindow::~QMdiSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DestroyQMdiSubWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
