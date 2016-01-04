package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraExposureControl struct {
	QMediaControl
}

type QCameraExposureControl_ITF interface {
	QMediaControl_ITF
	QCameraExposureControl_PTR() *QCameraExposureControl
}

func PointerFromQCameraExposureControl(ptr QCameraExposureControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraExposureControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraExposureControlFromPointer(ptr unsafe.Pointer) *QCameraExposureControl {
	var n = new(QCameraExposureControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraExposureControl_") {
		n.SetObjectName("QCameraExposureControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraExposureControl) QCameraExposureControl_PTR() *QCameraExposureControl {
	return ptr
}

//QCameraExposureControl::ExposureParameter
type QCameraExposureControl__ExposureParameter int64

const (
	QCameraExposureControl__ISO                       = QCameraExposureControl__ExposureParameter(0)
	QCameraExposureControl__Aperture                  = QCameraExposureControl__ExposureParameter(1)
	QCameraExposureControl__ShutterSpeed              = QCameraExposureControl__ExposureParameter(2)
	QCameraExposureControl__ExposureCompensation      = QCameraExposureControl__ExposureParameter(3)
	QCameraExposureControl__FlashPower                = QCameraExposureControl__ExposureParameter(4)
	QCameraExposureControl__FlashCompensation         = QCameraExposureControl__ExposureParameter(5)
	QCameraExposureControl__TorchPower                = QCameraExposureControl__ExposureParameter(6)
	QCameraExposureControl__SpotMeteringPoint         = QCameraExposureControl__ExposureParameter(7)
	QCameraExposureControl__ExposureMode              = QCameraExposureControl__ExposureParameter(8)
	QCameraExposureControl__MeteringMode              = QCameraExposureControl__ExposureParameter(9)
	QCameraExposureControl__ExtendedExposureParameter = QCameraExposureControl__ExposureParameter(1000)
)

func (ptr *QCameraExposureControl) ActualValue(parameter QCameraExposureControl__ExposureParameter) *core.QVariant {
	defer qt.Recovering("QCameraExposureControl::actualValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraExposureControl_ActualValue(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraExposureControl) ConnectActualValueChanged(f func(parameter int)) {
	defer qt.Recovering("connect QCameraExposureControl::actualValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectActualValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actualValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectActualValueChanged() {
	defer qt.Recovering("disconnect QCameraExposureControl::actualValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectActualValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actualValueChanged")
	}
}

//export callbackQCameraExposureControlActualValueChanged
func callbackQCameraExposureControlActualValueChanged(ptr unsafe.Pointer, ptrName *C.char, parameter C.int) {
	defer qt.Recovering("callback QCameraExposureControl::actualValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actualValueChanged"); signal != nil {
		signal.(func(int))(int(parameter))
	}

}

func (ptr *QCameraExposureControl) ActualValueChanged(parameter int) {
	defer qt.Recovering("QCameraExposureControl::actualValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ActualValueChanged(ptr.Pointer(), C.int(parameter))
	}
}

func (ptr *QCameraExposureControl) IsParameterSupported(parameter QCameraExposureControl__ExposureParameter) bool {
	defer qt.Recovering("QCameraExposureControl::isParameterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_IsParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) ConnectParameterRangeChanged(f func(parameter int)) {
	defer qt.Recovering("connect QCameraExposureControl::parameterRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectParameterRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parameterRangeChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectParameterRangeChanged() {
	defer qt.Recovering("disconnect QCameraExposureControl::parameterRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectParameterRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parameterRangeChanged")
	}
}

//export callbackQCameraExposureControlParameterRangeChanged
func callbackQCameraExposureControlParameterRangeChanged(ptr unsafe.Pointer, ptrName *C.char, parameter C.int) {
	defer qt.Recovering("callback QCameraExposureControl::parameterRangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "parameterRangeChanged"); signal != nil {
		signal.(func(int))(int(parameter))
	}

}

func (ptr *QCameraExposureControl) ParameterRangeChanged(parameter int) {
	defer qt.Recovering("QCameraExposureControl::parameterRangeChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ParameterRangeChanged(ptr.Pointer(), C.int(parameter))
	}
}

func (ptr *QCameraExposureControl) RequestedValue(parameter QCameraExposureControl__ExposureParameter) *core.QVariant {
	defer qt.Recovering("QCameraExposureControl::requestedValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraExposureControl_RequestedValue(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraExposureControl) ConnectRequestedValueChanged(f func(parameter int)) {
	defer qt.Recovering("connect QCameraExposureControl::requestedValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectRequestedValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectRequestedValueChanged() {
	defer qt.Recovering("disconnect QCameraExposureControl::requestedValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectRequestedValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedValueChanged")
	}
}

//export callbackQCameraExposureControlRequestedValueChanged
func callbackQCameraExposureControlRequestedValueChanged(ptr unsafe.Pointer, ptrName *C.char, parameter C.int) {
	defer qt.Recovering("callback QCameraExposureControl::requestedValueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestedValueChanged"); signal != nil {
		signal.(func(int))(int(parameter))
	}

}

func (ptr *QCameraExposureControl) RequestedValueChanged(parameter int) {
	defer qt.Recovering("QCameraExposureControl::requestedValueChanged")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_RequestedValueChanged(ptr.Pointer(), C.int(parameter))
	}
}

func (ptr *QCameraExposureControl) SetValue(parameter QCameraExposureControl__ExposureParameter, value core.QVariant_ITF) bool {
	defer qt.Recovering("QCameraExposureControl::setValue")

	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_SetValue(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) DestroyQCameraExposureControl() {
	defer qt.Recovering("QCameraExposureControl::~QCameraExposureControl")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DestroyQCameraExposureControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraExposureControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraExposureControlTimerEvent
func callbackQCameraExposureControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposureControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraExposureControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraExposureControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraExposureControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraExposureControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraExposureControlChildEvent
func callbackQCameraExposureControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposureControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraExposureControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraExposureControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraExposureControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraExposureControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraExposureControlCustomEvent
func callbackQCameraExposureControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraExposureControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraExposureControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraExposureControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraExposureControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraExposureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
