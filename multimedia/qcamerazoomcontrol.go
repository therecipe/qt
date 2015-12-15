package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraZoomControl struct {
	QMediaControl
}

type QCameraZoomControl_ITF interface {
	QMediaControl_ITF
	QCameraZoomControl_PTR() *QCameraZoomControl
}

func PointerFromQCameraZoomControl(ptr QCameraZoomControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraZoomControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraZoomControlFromPointer(ptr unsafe.Pointer) *QCameraZoomControl {
	var n = new(QCameraZoomControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraZoomControl_") {
		n.SetObjectName("QCameraZoomControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraZoomControl) QCameraZoomControl_PTR() *QCameraZoomControl {
	return ptr
}

func (ptr *QCameraZoomControl) CurrentDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::currentDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_CurrentDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) CurrentOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::currentOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_CurrentOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) MaximumDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::maximumDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_MaximumDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) MaximumOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::maximumOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_MaximumOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) RequestedDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::requestedDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_RequestedDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) RequestedOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::requestedOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_RequestedOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ZoomTo(optical float64, digital float64) {
	defer qt.Recovering("QCameraZoomControl::zoomTo")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ZoomTo(ptr.Pointer(), C.double(optical), C.double(digital))
	}
}

func (ptr *QCameraZoomControl) DestroyQCameraZoomControl() {
	defer qt.Recovering("QCameraZoomControl::~QCameraZoomControl")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DestroyQCameraZoomControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
