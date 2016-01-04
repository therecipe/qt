package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoAreaMonitorSource struct {
	core.QObject
}

type QGeoAreaMonitorSource_ITF interface {
	core.QObject_ITF
	QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource
}

func PointerFromQGeoAreaMonitorSource(ptr QGeoAreaMonitorSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorSource {
	var n = new(QGeoAreaMonitorSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoAreaMonitorSource_") {
		n.SetObjectName("QGeoAreaMonitorSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoAreaMonitorSource) QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource {
	return ptr
}

//QGeoAreaMonitorSource::AreaMonitorFeature
type QGeoAreaMonitorSource__AreaMonitorFeature int64

const (
	QGeoAreaMonitorSource__PersistentAreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0x00000001)
	QGeoAreaMonitorSource__AnyAreaMonitorFeature        = QGeoAreaMonitorSource__AreaMonitorFeature(0xffffffff)
)

//QGeoAreaMonitorSource::Error
type QGeoAreaMonitorSource__Error int64

const (
	QGeoAreaMonitorSource__AccessError              = QGeoAreaMonitorSource__Error(0)
	QGeoAreaMonitorSource__InsufficientPositionInfo = QGeoAreaMonitorSource__Error(1)
	QGeoAreaMonitorSource__UnknownSourceError       = QGeoAreaMonitorSource__Error(2)
	QGeoAreaMonitorSource__NoError                  = QGeoAreaMonitorSource__Error(3)
)

func QGeoAreaMonitorSource_AvailableSources() []string {
	defer qt.Recovering("QGeoAreaMonitorSource::availableSources")

	return strings.Split(C.GoString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), ",,,")
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer qt.Recovering("QGeoAreaMonitorSource::createDefaultSource")

	return NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer qt.Recovering("QGeoAreaMonitorSource::createSource")

	return NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoAreaMonitorSource) ConnectError2(f func(areaMonitoringError QGeoAreaMonitorSource__Error)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoAreaMonitorSourceError2
func callbackQGeoAreaMonitorSourceError2(ptr unsafe.Pointer, ptrName *C.char, areaMonitoringError C.int) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoAreaMonitorSource__Error))(QGeoAreaMonitorSource__Error(areaMonitoringError))
	}

}

func (ptr *QGeoAreaMonitorSource) Error2(areaMonitoringError QGeoAreaMonitorSource__Error) {
	defer qt.Recovering("QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_Error2(ptr.Pointer(), C.int(areaMonitoringError))
	}
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {
	defer qt.Recovering("QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__Error(C.QGeoAreaMonitorSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoAreaMonitorSource::positionInfoSource")

	if ptr.Pointer() != nil {
		return NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSource(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, signal string) bool {
	defer qt.Recovering("QGeoAreaMonitorSource::requestUpdate")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_RequestUpdate(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), C.CString(signal)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) ConnectSetPositionInfoSource(f func(newSource *QGeoPositionInfoSource)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPositionInfoSource", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectSetPositionInfoSource() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPositionInfoSource")
	}
}

//export callbackQGeoAreaMonitorSourceSetPositionInfoSource
func callbackQGeoAreaMonitorSourceSetPositionInfoSource(ptr unsafe.Pointer, ptrName *C.char, newSource unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::setPositionInfoSource")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPositionInfoSource"); signal != nil {
		signal.(func(*QGeoPositionInfoSource))(NewQGeoPositionInfoSourceFromPointer(newSource))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).SetPositionInfoSourceDefault(NewQGeoPositionInfoSourceFromPointer(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSource(newSource QGeoPositionInfoSource_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSource(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSourceDefault(newSource QGeoPositionInfoSource_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSourceDefault(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {
	defer qt.Recovering("QGeoAreaMonitorSource::sourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	defer qt.Recovering("QGeoAreaMonitorSource::startMonitoring")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StartMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	defer qt.Recovering("QGeoAreaMonitorSource::stopMonitoring")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StopMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() QGeoAreaMonitorSource__AreaMonitorFeature {
	defer qt.Recovering("QGeoAreaMonitorSource::supportedAreaMonitorFeatures")

	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__AreaMonitorFeature(C.QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {
	defer qt.Recovering("QGeoAreaMonitorSource::~QGeoAreaMonitorSource")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoAreaMonitorSourceTimerEvent
func callbackQGeoAreaMonitorSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoAreaMonitorSourceChildEvent
func callbackQGeoAreaMonitorSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoAreaMonitorSourceCustomEvent
func callbackQGeoAreaMonitorSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
