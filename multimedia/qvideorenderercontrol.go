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

type QVideoRendererControl_ITF interface {
	QMediaControl_ITF
	QVideoRendererControl_PTR() *QVideoRendererControl
}

func PointerFromQVideoRendererControl(ptr QVideoRendererControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoRendererControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoRendererControlFromPointer(ptr unsafe.Pointer) *QVideoRendererControl {
	var n = new(QVideoRendererControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QVideoRendererControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoRendererControl) QVideoRendererControl_PTR() *QVideoRendererControl {
	return ptr
}

func (ptr *QVideoRendererControl) SetSurface(surface QAbstractVideoSurface_ITF) {
	if ptr.Pointer() != nil {
		C.QVideoRendererControl_SetSurface(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QVideoRendererControl) Surface() *QAbstractVideoSurface {
	if ptr.Pointer() != nil {
		return NewQAbstractVideoSurfaceFromPointer(C.QVideoRendererControl_Surface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoRendererControl) DestroyQVideoRendererControl() {
	if ptr.Pointer() != nil {
		C.QVideoRendererControl_DestroyQVideoRendererControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
