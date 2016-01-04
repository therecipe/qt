package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
func callbackQCameraZoomControlCurrentDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::currentDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) CurrentDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::currentDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CurrentDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
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
func callbackQCameraZoomControlCurrentOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::currentOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) CurrentOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::currentOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CurrentOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
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
func callbackQCameraZoomControlMaximumDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::maximumDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) MaximumDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_MaximumDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
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
func callbackQCameraZoomControlMaximumOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::maximumOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maximumOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) MaximumOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_MaximumOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
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
func callbackQCameraZoomControlRequestedDigitalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::requestedDigitalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedDigitalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) RequestedDigitalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::requestedDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_RequestedDigitalZoomChanged(ptr.Pointer(), C.double(zoom))
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
func callbackQCameraZoomControlRequestedOpticalZoomChanged(ptr unsafe.Pointer, ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraZoomControl::requestedOpticalZoomChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedOpticalZoomChanged"); signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraZoomControl) RequestedOpticalZoomChanged(zoom float64) {
	defer qt.Recovering("QCameraZoomControl::requestedOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_RequestedOpticalZoomChanged(ptr.Pointer(), C.double(zoom))
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

func (ptr *QCameraZoomControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraZoomControlTimerEvent
func callbackQCameraZoomControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraZoomControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraZoomControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraZoomControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraZoomControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraZoomControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraZoomControlChildEvent
func callbackQCameraZoomControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraZoomControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraZoomControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraZoomControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraZoomControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraZoomControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraZoomControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraZoomControlCustomEvent
func callbackQCameraZoomControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraZoomControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraZoomControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraZoomControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraZoomControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraZoomControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraZoomControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
