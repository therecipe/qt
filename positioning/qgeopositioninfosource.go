package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoPositionInfoSource struct {
	core.QObject
}

type QGeoPositionInfoSource_ITF interface {
	core.QObject_ITF
	QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource
}

func PointerFromQGeoPositionInfoSource(ptr QGeoPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSource {
	var n = new(QGeoPositionInfoSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoPositionInfoSource_") {
		n.SetObjectName("QGeoPositionInfoSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoPositionInfoSource) QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource {
	return ptr
}

//QGeoPositionInfoSource::Error
type QGeoPositionInfoSource__Error int64

const (
	QGeoPositionInfoSource__AccessError        = QGeoPositionInfoSource__Error(0)
	QGeoPositionInfoSource__ClosedError        = QGeoPositionInfoSource__Error(1)
	QGeoPositionInfoSource__UnknownSourceError = QGeoPositionInfoSource__Error(2)
	QGeoPositionInfoSource__NoError            = QGeoPositionInfoSource__Error(3)
)

//QGeoPositionInfoSource::PositioningMethod
type QGeoPositionInfoSource__PositioningMethod int64

const (
	QGeoPositionInfoSource__NoPositioningMethods           = QGeoPositionInfoSource__PositioningMethod(0x00000000)
	QGeoPositionInfoSource__SatellitePositioningMethods    = QGeoPositionInfoSource__PositioningMethod(0x000000ff)
	QGeoPositionInfoSource__NonSatellitePositioningMethods = QGeoPositionInfoSource__PositioningMethod(0xffffff00)
	QGeoPositionInfoSource__AllPositioningMethods          = QGeoPositionInfoSource__PositioningMethod(0xffffffff)
)

func (ptr *QGeoPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setUpdateInterval", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetUpdateInterval() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setUpdateInterval")
	}
}

//export callbackQGeoPositionInfoSourceSetUpdateInterval
func callbackQGeoPositionInfoSourceSetUpdateInterval(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QGeoPositionInfoSource::setUpdateInterval")

	if signal := qt.GetSignal(C.GoString(ptrName), "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SetUpdateInterval(msec int) {
	defer qt.Recovering("QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SetUpdateIntervalDefault(msec int) {
	defer qt.Recovering("QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SourceName() string {
	defer qt.Recovering("QGeoPositionInfoSource::sourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoPositionInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	defer qt.Recovering("QGeoPositionInfoSource::updateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func QGeoPositionInfoSource_AvailableSources() []string {
	defer qt.Recovering("QGeoPositionInfoSource::availableSources")

	return strings.Split(C.GoString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), ",,,")
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoPositionInfoSource::createDefaultSource")

	return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoPositionInfoSource::createSource")

	return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoPositionInfoSource) ConnectError2(f func(positioningError QGeoPositionInfoSource__Error)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoPositionInfoSourceError2
func callbackQGeoPositionInfoSourceError2(ptr unsafe.Pointer, ptrName *C.char, positioningError C.int) {
	defer qt.Recovering("callback QGeoPositionInfoSource::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoPositionInfoSource__Error))(QGeoPositionInfoSource__Error(positioningError))
	}

}

func (ptr *QGeoPositionInfoSource) Error2(positioningError QGeoPositionInfoSource__Error) {
	defer qt.Recovering("QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_Error2(ptr.Pointer(), C.int(positioningError))
	}
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	defer qt.Recovering("QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QGeoPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) MinimumUpdateInterval() int {
	defer qt.Recovering("QGeoPositionInfoSource::minimumUpdateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer qt.Recovering("QGeoPositionInfoSource::preferredPositioningMethods")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_PreferredPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	defer qt.Recovering("QGeoPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods")
	}
}

//export callbackQGeoPositionInfoSourceSetPreferredPositioningMethods
func callbackQGeoPositionInfoSourceSetPreferredPositioningMethods(ptr unsafe.Pointer, ptrName *C.char, methods C.int) {
	defer qt.Recovering("callback QGeoPositionInfoSource::setPreferredPositioningMethods")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPreferredPositioningMethods"); signal != nil {
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethodsDefault(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {
	defer qt.Recovering("QGeoPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {
	defer qt.Recovering("QGeoPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer qt.Recovering("QGeoPositionInfoSource::supportedPositioningMethods")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	defer qt.Recovering("connect QGeoPositionInfoSource::updateTimeout")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectUpdateTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateTimeout", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::updateTimeout")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectUpdateTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateTimeout")
	}
}

//export callbackQGeoPositionInfoSourceUpdateTimeout
func callbackQGeoPositionInfoSourceUpdateTimeout(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGeoPositionInfoSource::updateTimeout")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) UpdateTimeout() {
	defer qt.Recovering("QGeoPositionInfoSource::updateTimeout")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_UpdateTimeout(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {
	defer qt.Recovering("QGeoPositionInfoSource::~QGeoPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoPositionInfoSourceTimerEvent
func callbackQGeoPositionInfoSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoPositionInfoSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoPositionInfoSourceChildEvent
func callbackQGeoPositionInfoSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoPositionInfoSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoPositionInfoSourceCustomEvent
func callbackQGeoPositionInfoSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoPositionInfoSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
