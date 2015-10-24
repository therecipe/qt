package gui

//#include "qopenglwindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOpenGLWindow struct {
	QPaintDeviceWindow
}

type QOpenGLWindowITF interface {
	QPaintDeviceWindowITF
	QOpenGLWindowPTR() *QOpenGLWindow
}

func PointerFromQOpenGLWindow(ptr QOpenGLWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLWindowPTR().Pointer()
	}
	return nil
}

func QOpenGLWindowFromPointer(ptr unsafe.Pointer) *QOpenGLWindow {
	var n = new(QOpenGLWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLWindow) QOpenGLWindowPTR() *QOpenGLWindow {
	return ptr
}

//QOpenGLWindow::UpdateBehavior
type QOpenGLWindow__UpdateBehavior int

var (
	QOpenGLWindow__NoPartialUpdate    = QOpenGLWindow__UpdateBehavior(0)
	QOpenGLWindow__PartialUpdateBlit  = QOpenGLWindow__UpdateBehavior(1)
	QOpenGLWindow__PartialUpdateBlend = QOpenGLWindow__UpdateBehavior(2)
)

func NewQOpenGLWindow2(shareContext QOpenGLContextITF, updateBehavior QOpenGLWindow__UpdateBehavior, parent QWindowITF) *QOpenGLWindow {
	return QOpenGLWindowFromPointer(unsafe.Pointer(C.QOpenGLWindow_NewQOpenGLWindow2(C.QtObjectPtr(PointerFromQOpenGLContext(shareContext)), C.int(updateBehavior), C.QtObjectPtr(PointerFromQWindow(parent)))))
}

func NewQOpenGLWindow(updateBehavior QOpenGLWindow__UpdateBehavior, parent QWindowITF) *QOpenGLWindow {
	return QOpenGLWindowFromPointer(unsafe.Pointer(C.QOpenGLWindow_NewQOpenGLWindow(C.int(updateBehavior), C.QtObjectPtr(PointerFromQWindow(parent)))))
}

func (ptr *QOpenGLWindow) Context() *QOpenGLContext {
	if ptr.Pointer() != nil {
		return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLWindow_Context(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLWindow) DoneCurrent() {
	if ptr.Pointer() != nil {
		C.QOpenGLWindow_DoneCurrent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLWindow) ConnectFrameSwapped(f func()) {
	if ptr.Pointer() != nil {
		C.QOpenGLWindow_ConnectFrameSwapped(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QOpenGLWindow) DisconnectFrameSwapped() {
	if ptr.Pointer() != nil {
		C.QOpenGLWindow_DisconnectFrameSwapped(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "frameSwapped")
	}
}

//export callbackQOpenGLWindowFrameSwapped
func callbackQOpenGLWindowFrameSwapped(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "frameSwapped").(func())()
}

func (ptr *QOpenGLWindow) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLWindow_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLWindow) MakeCurrent() {
	if ptr.Pointer() != nil {
		C.QOpenGLWindow_MakeCurrent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLWindow) ShareContext() *QOpenGLContext {
	if ptr.Pointer() != nil {
		return QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLWindow_ShareContext(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLWindow) UpdateBehavior() QOpenGLWindow__UpdateBehavior {
	if ptr.Pointer() != nil {
		return QOpenGLWindow__UpdateBehavior(C.QOpenGLWindow_UpdateBehavior(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLWindow) DestroyQOpenGLWindow() {
	if ptr.Pointer() != nil {
		C.QOpenGLWindow_DestroyQOpenGLWindow(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
