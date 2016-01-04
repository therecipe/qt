package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNmeaPositionInfoSource struct {
	QGeoPositionInfoSource
}

type QNmeaPositionInfoSource_ITF interface {
	QGeoPositionInfoSource_ITF
	QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource
}

func PointerFromQNmeaPositionInfoSource(ptr QNmeaPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNmeaPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QNmeaPositionInfoSource {
	var n = new(QNmeaPositionInfoSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNmeaPositionInfoSource_") {
		n.SetObjectName("QNmeaPositionInfoSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QNmeaPositionInfoSource) QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource {
	return ptr
}

//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int64

const (
	QNmeaPositionInfoSource__RealTimeMode   = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode = QNmeaPositionInfoSource__UpdateMode(2)
)

func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObject_ITF) *QNmeaPositionInfoSource {
	defer qt.Recovering("QNmeaPositionInfoSource::QNmeaPositionInfoSource")

	return NewQNmeaPositionInfoSourceFromPointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.int(updateMode), core.PointerFromQObject(parent)))
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	defer qt.Recovering("QNmeaPositionInfoSource::device")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNmeaPositionInfoSource_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	defer qt.Recovering("QNmeaPositionInfoSource::error")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QNmeaPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateInterval() int {
	defer qt.Recovering("QNmeaPositionInfoSource::minimumUpdateInterval")

	if ptr.Pointer() != nil {
		return int(C.QNmeaPositionInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data string, size int, posInfo QGeoPositionInfo_ITF, hasFix bool) bool {
	defer qt.Recovering("QNmeaPositionInfoSource::parsePosInfoFromNmeaData")

	if ptr.Pointer() != nil {
		return C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr.Pointer(), C.CString(data), C.int(size), PointerFromQGeoPositionInfo(posInfo), C.int(qt.GoBoolToInt(hasFix))) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) ConnectRequestUpdate(f func(msec int)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestUpdate", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectRequestUpdate() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestUpdate")
	}
}

//export callbackQNmeaPositionInfoSourceRequestUpdate
func callbackQNmeaPositionInfoSourceRequestUpdate(ptr unsafe.Pointer, ptrName *C.char, msec C.int) bool {
	defer qt.Recovering("callback QNmeaPositionInfoSource::requestUpdate")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestUpdate"); signal != nil {
		signal.(func(int))(int(msec))
		return true
	}
	return false

}

func (ptr *QNmeaPositionInfoSource) RequestUpdate(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) RequestUpdateDefault(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdateDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::setDevice")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setUpdateInterval", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSetUpdateInterval() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setUpdateInterval")
	}
}

//export callbackQNmeaPositionInfoSourceSetUpdateInterval
func callbackQNmeaPositionInfoSourceSetUpdateInterval(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::setUpdateInterval")

	if signal := qt.GetSignal(C.GoString(ptrName), "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateInterval(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateIntervalDefault(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStartUpdates(f func()) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startUpdates", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStartUpdates() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startUpdates")
	}
}

//export callbackQNmeaPositionInfoSourceStartUpdates
func callbackQNmeaPositionInfoSourceStartUpdates(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QNmeaPositionInfoSource::startUpdates")

	if signal := qt.GetSignal(C.GoString(ptrName), "startUpdates"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QNmeaPositionInfoSource) StartUpdates() {
	defer qt.Recovering("QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StartUpdatesDefault() {
	defer qt.Recovering("QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdatesDefault(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStopUpdates(f func()) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stopUpdates", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStopUpdates() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stopUpdates")
	}
}

//export callbackQNmeaPositionInfoSourceStopUpdates
func callbackQNmeaPositionInfoSourceStopUpdates(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QNmeaPositionInfoSource::stopUpdates")

	if signal := qt.GetSignal(C.GoString(ptrName), "stopUpdates"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QNmeaPositionInfoSource) StopUpdates() {
	defer qt.Recovering("QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StopUpdatesDefault() {
	defer qt.Recovering("QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdatesDefault(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer qt.Recovering("QNmeaPositionInfoSource::supportedPositioningMethods")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QNmeaPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) UpdateMode() QNmeaPositionInfoSource__UpdateMode {
	defer qt.Recovering("QNmeaPositionInfoSource::updateMode")

	if ptr.Pointer() != nil {
		return QNmeaPositionInfoSource__UpdateMode(C.QNmeaPositionInfoSource_UpdateMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	defer qt.Recovering("QNmeaPositionInfoSource::~QNmeaPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods")
	}
}

//export callbackQNmeaPositionInfoSourceSetPreferredPositioningMethods
func callbackQNmeaPositionInfoSourceSetPreferredPositioningMethods(ptr unsafe.Pointer, ptrName *C.char, methods C.int) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPreferredPositioningMethods"); signal != nil {
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetPreferredPositioningMethodsDefault(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNmeaPositionInfoSourceTimerEvent
func callbackQNmeaPositionInfoSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNmeaPositionInfoSourceChildEvent
func callbackQNmeaPositionInfoSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNmeaPositionInfoSourceCustomEvent
func callbackQNmeaPositionInfoSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
