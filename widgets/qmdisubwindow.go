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

type QMdiSubWindowITF interface {
	QWidgetITF
	QMdiSubWindowPTR() *QMdiSubWindow
}

func PointerFromQMdiSubWindow(ptr QMdiSubWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiSubWindowPTR().Pointer()
	}
	return nil
}

func QMdiSubWindowFromPointer(ptr unsafe.Pointer) *QMdiSubWindow {
	var n = new(QMdiSubWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMdiSubWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMdiSubWindow) QMdiSubWindowPTR() *QMdiSubWindow {
	return ptr
}

//QMdiSubWindow::SubWindowOption
type QMdiSubWindow__SubWindowOption int

var (
	QMdiSubWindow__AllowOutsideAreaHorizontally = QMdiSubWindow__SubWindowOption(0x1)
	QMdiSubWindow__AllowOutsideAreaVertically   = QMdiSubWindow__SubWindowOption(0x2)
	QMdiSubWindow__RubberBandResize             = QMdiSubWindow__SubWindowOption(0x4)
	QMdiSubWindow__RubberBandMove               = QMdiSubWindow__SubWindowOption(0x8)
)

func (ptr *QMdiSubWindow) KeyboardPageStep() int {
	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardPageStep(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMdiSubWindow) KeyboardSingleStep() int {
	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardSingleStep(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMdiSubWindow) SetKeyboardPageStep(step int) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardPageStep(C.QtObjectPtr(ptr.Pointer()), C.int(step))
	}
}

func (ptr *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardSingleStep(C.QtObjectPtr(ptr.Pointer()), C.int(step))
	}
}

func NewQMdiSubWindow(parent QWidgetITF, flags core.Qt__WindowType) *QMdiSubWindow {
	return QMdiSubWindowFromPointer(unsafe.Pointer(C.QMdiSubWindow_NewQMdiSubWindow(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func (ptr *QMdiSubWindow) ConnectAboutToActivate(f func()) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectAboutToActivate(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToActivate", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectAboutToActivate() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectAboutToActivate(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToActivate")
	}
}

//export callbackQMdiSubWindowAboutToActivate
func callbackQMdiSubWindowAboutToActivate(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToActivate").(func())()
}

func (ptr *QMdiSubWindow) IsShaded() bool {
	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_IsShaded(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) MdiArea() *QMdiArea {
	if ptr.Pointer() != nil {
		return QMdiAreaFromPointer(unsafe.Pointer(C.QMdiSubWindow_MdiArea(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMdiSubWindow) SetOption(option QMdiSubWindow__SubWindowOption, on bool) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiSubWindow) SetSystemMenu(systemMenu QMenuITF) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetSystemMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(systemMenu)))
	}
}

func (ptr *QMdiSubWindow) SetWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QMdiSubWindow) ShowShaded() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowShaded(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiSubWindow) ShowSystemMenu() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowSystemMenu(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiSubWindow) SystemMenu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QMdiSubWindow_SystemMenu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMdiSubWindow) TestOption(option QMdiSubWindow__SubWindowOption) bool {
	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QMdiSubWindow_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMdiSubWindow) ConnectWindowStateChanged(f func(oldState core.Qt__WindowState, newState core.Qt__WindowState)) {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectWindowStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectWindowStateChanged() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectWindowStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQMdiSubWindowWindowStateChanged
func callbackQMdiSubWindowWindowStateChanged(ptrName *C.char, oldState C.int, newState C.int) {
	qt.GetSignal(C.GoString(ptrName), "windowStateChanged").(func(core.Qt__WindowState, core.Qt__WindowState))(core.Qt__WindowState(oldState), core.Qt__WindowState(newState))
}

func (ptr *QMdiSubWindow) DestroyQMdiSubWindow() {
	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DestroyQMdiSubWindow(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
