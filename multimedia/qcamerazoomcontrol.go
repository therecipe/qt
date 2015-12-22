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

func (ptr *QCameraZoomControl) ConnectCurrentDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::currentDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectCurrentDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentDigitalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectCurrentDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::currentDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectCurrentDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentDigitalZoomChanged")
	}
}

//export callbackQCameraZoomControlCurrentDigitalZoomChanged
func callbackQCameraZoomControlCurrentDigitalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::currentDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) CurrentOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::currentOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_CurrentOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectCurrentOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::currentOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectCurrentOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentOpticalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectCurrentOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::currentOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectCurrentOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentOpticalZoomChanged")
	}
}

//export callbackQCameraZoomControlCurrentOpticalZoomChanged
func callbackQCameraZoomControlCurrentOpticalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::currentOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) MaximumDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::maximumDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_MaximumDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectMaximumDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectMaximumDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged")
	}
}

//export callbackQCameraZoomControlMaximumDigitalZoomChanged
func callbackQCameraZoomControlMaximumDigitalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::maximumDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) MaximumOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::maximumOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_MaximumOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectMaximumOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectMaximumOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged")
	}
}

//export callbackQCameraZoomControlMaximumOpticalZoomChanged
func callbackQCameraZoomControlMaximumOpticalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::maximumOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) RequestedDigitalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::requestedDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_RequestedDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectRequestedDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::requestedDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectRequestedDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedDigitalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectRequestedDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::requestedDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectRequestedDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedDigitalZoomChanged")
	}
}

//export callbackQCameraZoomControlRequestedDigitalZoomChanged
func callbackQCameraZoomControlRequestedDigitalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::requestedDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) RequestedOpticalZoom() float64 {
	defer qt.Recovering("QCameraZoomControl::requestedOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraZoomControl_RequestedOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraZoomControl) ConnectRequestedOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraZoomControl::requestedOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ConnectRequestedOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedOpticalZoomChanged", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectRequestedOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraZoomControl::requestedOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_DisconnectRequestedOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedOpticalZoomChanged")
	}
}

//export callbackQCameraZoomControlRequestedOpticalZoomChanged
func callbackQCameraZoomControlRequestedOpticalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::requestedOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

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
