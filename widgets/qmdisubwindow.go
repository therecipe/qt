package widgets

//#include "qmdisubwindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
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
	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardPageStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiSubWindow) KeyboardSingleStep() int {
	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardSingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiSubWindow) SetKeyboardPageStep(step int) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardPageStep(ptr.Pointer(), C.int(step))
	}
}

func (ptr *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardSingleStep(ptr.Pointer(), C.int(step))
	}
}

func NewQMdiSubWindow(parent QWidget_ITF, flags core.Qt__WindowType) *QMdiSubWindow {
	return NewQMdiSubWindowFromPointer(C.QMdiSubWindow_NewQMdiSubWindow(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QMdiSubWindow) ConnectAboutToActivate(f func()) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectAboutToActivate(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToActivate", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectAboutToActivate() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectAboutToActivate(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToActivate")
	}
}

//export callbackQMdiSubWindowAboutToActivate
func callbackQMdiSubWindowAboutToActivate(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToActivate").(func())()
}

func (ptr *QMdiSubWindow) IsShaded() bool {
	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_IsShaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) MdiArea() *QMdiArea {
	if ptr.Pointer() != nil {
		return NewQMdiAreaFromPointer(C.QMdiSubWindow_MdiArea(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) SetOption(option QMdiSubWindow__SubWindowOption, on bool) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiSubWindow) SetSystemMenu(systemMenu QMenu_ITF) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetSystemMenu(ptr.Pointer(), PointerFromQMenu(systemMenu))
	}
}

func (ptr *QMdiSubWindow) SetWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMdiSubWindow) ShowShaded() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowShaded(ptr.Pointer())
	}
}

func (ptr *QMdiSubWindow) ShowSystemMenu() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowSystemMenu(ptr.Pointer())
	}
}

func (ptr *QMdiSubWindow) SystemMenu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMdiSubWindow_SystemMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) TestOption(option QMdiSubWindow__SubWindowOption) bool {
	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMdiSubWindow_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) ConnectWindowStateChanged(f func(oldState core.Qt__WindowState, newState core.Qt__WindowState)) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectWindowStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectWindowStateChanged() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectWindowStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQMdiSubWindowWindowStateChanged
func callbackQMdiSubWindowWindowStateChanged(ptrName *C.char, oldState C.int, newState C.int) {
	qt.GetSignal(C.GoString(ptrName), "windowStateChanged").(func(core.Qt__WindowState, core.Qt__WindowState))(core.Qt__WindowState(oldState), core.Qt__WindowState(newState))
}

func (ptr *QMdiSubWindow) DestroyQMdiSubWindow() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DestroyQMdiSubWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
