package widgets

//#include "qopenglwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QOpenGLWidget struct {
	QWidget
}

type QOpenGLWidgetITF interface {
	QWidgetITF
	QOpenGLWidgetPTR() *QOpenGLWidget
}

func PointerFromQOpenGLWidget(ptr QOpenGLWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLWidgetPTR().Pointer()
	}
	return nil
}

func QOpenGLWidgetFromPointer(ptr unsafe.Pointer) *QOpenGLWidget {
	var n = new(QOpenGLWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLWidget) QOpenGLWidgetPTR() *QOpenGLWidget {
	return ptr
}

//QOpenGLWidget::UpdateBehavior
type QOpenGLWidget__UpdateBehavior int

var (
	QOpenGLWidget__NoPartialUpdate = QOpenGLWidget__UpdateBehavior(0)
	QOpenGLWidget__PartialUpdate   = QOpenGLWidget__UpdateBehavior(1)
)

func NewQOpenGLWidget(parent QWidgetITF, f core.Qt__WindowType) *QOpenGLWidget {
	return QOpenGLWidgetFromPointer(unsafe.Pointer(C.QOpenGLWidget_NewQOpenGLWidget(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func (ptr *QOpenGLWidget) ConnectAboutToCompose(f func()) {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_ConnectAboutToCompose(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToCompose", f)
	}
}

func (ptr *QOpenGLWidget) DisconnectAboutToCompose() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_DisconnectAboutToCompose(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToCompose")
	}
}

//export callbackQOpenGLWidgetAboutToCompose
func callbackQOpenGLWidgetAboutToCompose(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToCompose").(func())()
}

func (ptr *QOpenGLWidget) ConnectAboutToResize(f func()) {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_ConnectAboutToResize(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToResize", f)
	}
}

func (ptr *QOpenGLWidget) DisconnectAboutToResize() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_DisconnectAboutToResize(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToResize")
	}
}

//export callbackQOpenGLWidgetAboutToResize
func callbackQOpenGLWidgetAboutToResize(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToResize").(func())()
}

func (ptr *QOpenGLWidget) Context() *gui.QOpenGLContext {
	if ptr.Pointer() != nil {
		return gui.QOpenGLContextFromPointer(unsafe.Pointer(C.QOpenGLWidget_Context(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QOpenGLWidget) DoneCurrent() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_DoneCurrent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLWidget) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLWidget_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLWidget) MakeCurrent() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_MakeCurrent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLWidget) ConnectFrameSwapped(f func()) {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_ConnectFrameSwapped(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QOpenGLWidget) DisconnectFrameSwapped() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_DisconnectFrameSwapped(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "frameSwapped")
	}
}

//export callbackQOpenGLWidgetFrameSwapped
func callbackQOpenGLWidgetFrameSwapped(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "frameSwapped").(func())()
}

func (ptr *QOpenGLWidget) ConnectResized(f func()) {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_ConnectResized(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QOpenGLWidget) DisconnectResized() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_DisconnectResized(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQOpenGLWidgetResized
func callbackQOpenGLWidgetResized(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "resized").(func())()
}

func (ptr *QOpenGLWidget) SetFormat(format gui.QSurfaceFormatITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQSurfaceFormat(format)))
	}
}

func (ptr *QOpenGLWidget) SetUpdateBehavior(updateBehavior QOpenGLWidget__UpdateBehavior) {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_SetUpdateBehavior(C.QtObjectPtr(ptr.Pointer()), C.int(updateBehavior))
	}
}

func (ptr *QOpenGLWidget) UpdateBehavior() QOpenGLWidget__UpdateBehavior {
	if ptr.Pointer() != nil {
		return QOpenGLWidget__UpdateBehavior(C.QOpenGLWidget_UpdateBehavior(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLWidget) DestroyQOpenGLWidget() {
	if ptr.Pointer() != nil {
		C.QOpenGLWidget_DestroyQOpenGLWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
