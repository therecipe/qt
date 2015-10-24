package multimedia

//#include "qvideorenderercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVideoRendererControl struct {
	QMediaControl
}

type QVideoRendererControlITF interface {
	QMediaControlITF
	QVideoRendererControlPTR() *QVideoRendererControl
}

func PointerFromQVideoRendererControl(ptr QVideoRendererControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoRendererControlPTR().Pointer()
	}
	return nil
}

func QVideoRendererControlFromPointer(ptr unsafe.Pointer) *QVideoRendererControl {
	var n = new(QVideoRendererControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVideoRendererControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoRendererControl) QVideoRendererControlPTR() *QVideoRendererControl {
	return ptr
}

func (ptr *QVideoRendererControl) SetSurface(surface QAbstractVideoSurfaceITF) {
	if ptr.Pointer() != nil {
		C.QVideoRendererControl_SetSurface(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractVideoSurface(surface)))
	}
}

func (ptr *QVideoRendererControl) Surface() *QAbstractVideoSurface {
	if ptr.Pointer() != nil {
		return QAbstractVideoSurfaceFromPointer(unsafe.Pointer(C.QVideoRendererControl_Surface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QVideoRendererControl) DestroyQVideoRendererControl() {
	if ptr.Pointer() != nil {
		C.QVideoRendererControl_DestroyQVideoRendererControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
