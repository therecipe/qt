package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraInfoControl struct {
	QMediaControl
}

type QCameraInfoControl_ITF interface {
	QMediaControl_ITF
	QCameraInfoControl_PTR() *QCameraInfoControl
}

func PointerFromQCameraInfoControl(ptr QCameraInfoControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfoControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraInfoControlFromPointer(ptr unsafe.Pointer) *QCameraInfoControl {
	var n = new(QCameraInfoControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraInfoControl_") {
		n.SetObjectName("QCameraInfoControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraInfoControl) QCameraInfoControl_PTR() *QCameraInfoControl {
	return ptr
}

func (ptr *QCameraInfoControl) CameraOrientation(deviceName string) int {
	defer qt.Recovering("QCameraInfoControl::cameraOrientation")

	if ptr.Pointer() != nil {
		return int(C.QCameraInfoControl_CameraOrientation(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) CameraPosition(deviceName string) QCamera__Position {
	defer qt.Recovering("QCameraInfoControl::cameraPosition")

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfoControl_CameraPosition(ptr.Pointer(), C.CString(deviceName)))
	}
	return 0
}

func (ptr *QCameraInfoControl) DestroyQCameraInfoControl() {
	defer qt.Recovering("QCameraInfoControl::~QCameraInfoControl")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_DestroyQCameraInfoControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraInfoControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraInfoControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraInfoControlTimerEvent
func callbackQCameraInfoControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraInfoControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraInfoControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraInfoControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraInfoControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraInfoControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraInfoControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraInfoControlChildEvent
func callbackQCameraInfoControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraInfoControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraInfoControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraInfoControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraInfoControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraInfoControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraInfoControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraInfoControlCustomEvent
func callbackQCameraInfoControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraInfoControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraInfoControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraInfoControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraInfoControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraInfoControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraInfoControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
