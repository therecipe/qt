package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoSatelliteInfoSource struct {
	core.QObject
}

type QGeoSatelliteInfoSource_ITF interface {
	core.QObject_ITF
	QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource
}

func PointerFromQGeoSatelliteInfoSource(ptr QGeoSatelliteInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfoSource {
	var n = new(QGeoSatelliteInfoSource)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGeoSatelliteInfoSource_") {
		n.SetObjectName("QGeoSatelliteInfoSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoSatelliteInfoSource) QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource {
	return ptr
}

//QGeoSatelliteInfoSource::Error
type QGeoSatelliteInfoSource__Error int64

const (
	QGeoSatelliteInfoSource__AccessError        = QGeoSatelliteInfoSource__Error(0)
	QGeoSatelliteInfoSource__ClosedError        = QGeoSatelliteInfoSource__Error(1)
	QGeoSatelliteInfoSource__NoError            = QGeoSatelliteInfoSource__Error(2)
	QGeoSatelliteInfoSource__UnknownSourceError = QGeoSatelliteInfoSource__Error(-1)
)

func (ptr *QGeoSatelliteInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setUpdateInterval", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSetUpdateInterval() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setUpdateInterval")
	}
}

//export callbackQGeoSatelliteInfoSourceSetUpdateInterval
func callbackQGeoSatelliteInfoSourceSetUpdateInterval(ptrName *C.char, msec C.int) bool {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::setUpdateInterval")

	if signal := qt.GetSignal(C.GoString(ptrName), "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(msec))
		return true
	}
	return false

}

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	defer qt.Recovering("QGeoSatelliteInfoSource::updateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	defer qt.Recovering("QGeoSatelliteInfoSource::availableSources")

	return strings.Split(C.GoString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), ",,,")
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer qt.Recovering("QGeoSatelliteInfoSource::createDefaultSource")

	return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer qt.Recovering("QGeoSatelliteInfoSource::createSource")

	return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoSatelliteInfoSource) ConnectError2(f func(satelliteError QGeoSatelliteInfoSource__Error)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoSatelliteInfoSourceError2
func callbackQGeoSatelliteInfoSourceError2(ptrName *C.char, satelliteError C.int) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoSatelliteInfoSource__Error))(QGeoSatelliteInfoSource__Error(satelliteError))
	}

}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {
	defer qt.Recovering("QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSource__Error(C.QGeoSatelliteInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {
	defer qt.Recovering("QGeoSatelliteInfoSource::minimumUpdateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::requestTimeout")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectRequestTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestTimeout", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::requestTimeout")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectRequestTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestTimeout")
	}
}

//export callbackQGeoSatelliteInfoSourceRequestTimeout
func callbackQGeoSatelliteInfoSourceRequestTimeout(ptrName *C.char) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::requestTimeout")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	defer qt.Recovering("QGeoSatelliteInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestUpdate(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	defer qt.Recovering("QGeoSatelliteInfoSource::sourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoSatelliteInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {
	defer qt.Recovering("QGeoSatelliteInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {
	defer qt.Recovering("QGeoSatelliteInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {
	defer qt.Recovering("QGeoSatelliteInfoSource::~QGeoSatelliteInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoSatelliteInfoSourceTimerEvent
func callbackQGeoSatelliteInfoSourceTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGeoSatelliteInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoSatelliteInfoSourceChildEvent
func callbackQGeoSatelliteInfoSourceChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGeoSatelliteInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoSatelliteInfoSourceCustomEvent
func callbackQGeoSatelliteInfoSourceCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
