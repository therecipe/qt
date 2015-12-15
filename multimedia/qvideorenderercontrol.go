package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QVideoRendererControl_") {
		n.SetObjectName("QVideoRendererControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoRendererControl) QVideoRendererControl_PTR() *QVideoRendererControl {
	return ptr
}

func (ptr *QVideoRendererControl) SetSurface(surface QAbstractVideoSurface_ITF) {
	defer qt.Recovering("QVideoRendererControl::setSurface")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_SetSurface(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QVideoRendererControl) Surface() *QAbstractVideoSurface {
	defer qt.Recovering("QVideoRendererControl::surface")

	if ptr.Pointer() != nil {
		return NewQAbstractVideoSurfaceFromPointer(C.QVideoRendererControl_Surface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoRendererControl) DestroyQVideoRendererControl() {
	defer qt.Recovering("QVideoRendererControl::~QVideoRendererControl")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_DestroyQVideoRendererControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
