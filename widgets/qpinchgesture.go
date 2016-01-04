package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPinchGesture struct {
	QGesture
}

type QPinchGesture_ITF interface {
	QGesture_ITF
	QPinchGesture_PTR() *QPinchGesture
}

func PointerFromQPinchGesture(ptr QPinchGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPinchGesture_PTR().Pointer()
	}
	return nil
}

func NewQPinchGestureFromPointer(ptr unsafe.Pointer) *QPinchGesture {
	var n = new(QPinchGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPinchGesture_") {
		n.SetObjectName("QPinchGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QPinchGesture) QPinchGesture_PTR() *QPinchGesture {
	return ptr
}

//QPinchGesture::ChangeFlag
type QPinchGesture__ChangeFlag int64

const (
	QPinchGesture__ScaleFactorChanged   = QPinchGesture__ChangeFlag(0x1)
	QPinchGesture__RotationAngleChanged = QPinchGesture__ChangeFlag(0x2)
	QPinchGesture__CenterPointChanged   = QPinchGesture__ChangeFlag(0x4)
)

func (ptr *QPinchGesture) ChangeFlags() QPinchGesture__ChangeFlag {
	defer qt.Recovering("QPinchGesture::changeFlags")

	if ptr.Pointer() != nil {
		return QPinchGesture__ChangeFlag(C.QPinchGesture_ChangeFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) LastRotationAngle() float64 {
	defer qt.Recovering("QPinchGesture::lastRotationAngle")

	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_LastRotationAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) LastScaleFactor() float64 {
	defer qt.Recovering("QPinchGesture::lastScaleFactor")

	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_LastScaleFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) RotationAngle() float64 {
	defer qt.Recovering("QPinchGesture::rotationAngle")

	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_RotationAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) ScaleFactor() float64 {
	defer qt.Recovering("QPinchGesture::scaleFactor")

	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_ScaleFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) SetCenterPoint(value core.QPointF_ITF) {
	defer qt.Recovering("QPinchGesture::setCenterPoint")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetCenterPoint(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPinchGesture) SetChangeFlags(value QPinchGesture__ChangeFlag) {
	defer qt.Recovering("QPinchGesture::setChangeFlags")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetChangeFlags(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QPinchGesture) SetLastCenterPoint(value core.QPointF_ITF) {
	defer qt.Recovering("QPinchGesture::setLastCenterPoint")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastCenterPoint(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPinchGesture) SetLastRotationAngle(value float64) {
	defer qt.Recovering("QPinchGesture::setLastRotationAngle")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastRotationAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetLastScaleFactor(value float64) {
	defer qt.Recovering("QPinchGesture::setLastScaleFactor")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetLastScaleFactor(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetRotationAngle(value float64) {
	defer qt.Recovering("QPinchGesture::setRotationAngle")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetRotationAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetScaleFactor(value float64) {
	defer qt.Recovering("QPinchGesture::setScaleFactor")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetScaleFactor(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetStartCenterPoint(value core.QPointF_ITF) {
	defer qt.Recovering("QPinchGesture::setStartCenterPoint")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetStartCenterPoint(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPinchGesture) SetTotalChangeFlags(value QPinchGesture__ChangeFlag) {
	defer qt.Recovering("QPinchGesture::setTotalChangeFlags")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalChangeFlags(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QPinchGesture) SetTotalRotationAngle(value float64) {
	defer qt.Recovering("QPinchGesture::setTotalRotationAngle")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalRotationAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) SetTotalScaleFactor(value float64) {
	defer qt.Recovering("QPinchGesture::setTotalScaleFactor")

	if ptr.Pointer() != nil {
		C.QPinchGesture_SetTotalScaleFactor(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPinchGesture) TotalChangeFlags() QPinchGesture__ChangeFlag {
	defer qt.Recovering("QPinchGesture::totalChangeFlags")

	if ptr.Pointer() != nil {
		return QPinchGesture__ChangeFlag(C.QPinchGesture_TotalChangeFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) TotalRotationAngle() float64 {
	defer qt.Recovering("QPinchGesture::totalRotationAngle")

	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_TotalRotationAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) TotalScaleFactor() float64 {
	defer qt.Recovering("QPinchGesture::totalScaleFactor")

	if ptr.Pointer() != nil {
		return float64(C.QPinchGesture_TotalScaleFactor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPinchGesture) DestroyQPinchGesture() {
	defer qt.Recovering("QPinchGesture::~QPinchGesture")

	if ptr.Pointer() != nil {
		C.QPinchGesture_DestroyQPinchGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPinchGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPinchGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPinchGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPinchGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPinchGestureTimerEvent
func callbackQPinchGestureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPinchGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPinchGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPinchGesture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPinchGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QPinchGesture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPinchGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPinchGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QPinchGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPinchGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPinchGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPinchGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPinchGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPinchGestureChildEvent
func callbackQPinchGestureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPinchGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPinchGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPinchGesture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPinchGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QPinchGesture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPinchGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPinchGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QPinchGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPinchGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPinchGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPinchGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPinchGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPinchGestureCustomEvent
func callbackQPinchGestureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPinchGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPinchGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPinchGesture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPinchGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QPinchGesture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPinchGesture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPinchGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QPinchGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
