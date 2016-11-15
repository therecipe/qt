// +build !minimal

package sensors

//#include <stdint.h>
//#include <stdlib.h>
//#include "sensors.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSensors_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int64

const (
	QAccelerometer__Combined = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     = QAccelerometer__AccelerationMode(2)
)

type QAccelerometer struct {
	QSensor
}

type QAccelerometer_ITF interface {
	QSensor_ITF
	QAccelerometer_PTR() *QAccelerometer
}

func (p *QAccelerometer) QAccelerometer_PTR() *QAccelerometer {
	return p
}

func (p *QAccelerometer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QAccelerometer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQAccelerometer(ptr QAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometer_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFromPointer(ptr unsafe.Pointer) *QAccelerometer {
	var n = new(QAccelerometer)
	n.SetPointer(ptr)
	return n
}
func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {
	if ptr.Pointer() != nil {
		return QAccelerometer__AccelerationMode(C.QAccelerometer_AccelerationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAccelerometerReadingFromPointer(C.QAccelerometer_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQAccelerometer(parent core.QObject_ITF) *QAccelerometer {
	var tmpValue = NewQAccelerometerFromPointer(C.QAccelerometer_NewQAccelerometer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAccelerometer_AccelerationModeChanged
func callbackQAccelerometer_AccelerationModeChanged(ptr unsafe.Pointer, accelerationMode C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::accelerationModeChanged"); signal != nil {
		signal.(func(QAccelerometer__AccelerationMode))(QAccelerometer__AccelerationMode(accelerationMode))
	}

}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectAccelerationModeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::accelerationModeChanged", f)
	}
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectAccelerationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::accelerationModeChanged")
	}
}

func (ptr *QAccelerometer) AccelerationModeChanged(accelerationMode QAccelerometer__AccelerationMode) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_AccelerationModeChanged(ptr.Pointer(), C.longlong(accelerationMode))
	}
}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_SetAccelerationMode(ptr.Pointer(), C.longlong(accelerationMode))
	}
}

//export callbackQAccelerometer_DestroyQAccelerometer
func callbackQAccelerometer_DestroyQAccelerometer(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::~QAccelerometer"); signal != nil {
		signal.(func())()
	} else {
		NewQAccelerometerFromPointer(ptr).DestroyQAccelerometerDefault()
	}
}

func (ptr *QAccelerometer) ConnectDestroyQAccelerometer(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::~QAccelerometer", f)
	}
}

func (ptr *QAccelerometer) DisconnectDestroyQAccelerometer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::~QAccelerometer")
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DestroyQAccelerometer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometerDefault() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DestroyQAccelerometerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QAccelerometer_Type() string {
	return cGoUnpackString(C.QAccelerometer_QAccelerometer_Type())
}

func (ptr *QAccelerometer) Type() string {
	return cGoUnpackString(C.QAccelerometer_QAccelerometer_Type())
}

//export callbackQAccelerometer_Start
func callbackQAccelerometer_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAccelerometerFromPointer(ptr).StartDefault())))
}

func (ptr *QAccelerometer) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::start", f)
	}
}

func (ptr *QAccelerometer) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::start")
	}
}

func (ptr *QAccelerometer) Start() bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometer_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccelerometer) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometer_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAccelerometer_Stop
func callbackQAccelerometer_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQAccelerometerFromPointer(ptr).StopDefault()
	}
}

func (ptr *QAccelerometer) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::stop", f)
	}
}

func (ptr *QAccelerometer) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::stop")
	}
}

func (ptr *QAccelerometer) Stop() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_Stop(ptr.Pointer())
	}
}

func (ptr *QAccelerometer) StopDefault() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_StopDefault(ptr.Pointer())
	}
}

//export callbackQAccelerometer_TimerEvent
func callbackQAccelerometer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::timerEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::timerEvent")
	}
}

func (ptr *QAccelerometer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAccelerometer_ChildEvent
func callbackQAccelerometer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::childEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::childEvent")
	}
}

func (ptr *QAccelerometer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAccelerometer_ConnectNotify
func callbackQAccelerometer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAccelerometerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAccelerometer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::connectNotify", f)
	}
}

func (ptr *QAccelerometer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::connectNotify")
	}
}

func (ptr *QAccelerometer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAccelerometer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAccelerometer_CustomEvent
func callbackQAccelerometer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::customEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::customEvent")
	}
}

func (ptr *QAccelerometer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAccelerometer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAccelerometer_DeleteLater
func callbackQAccelerometer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAccelerometerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAccelerometer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::deleteLater", f)
	}
}

func (ptr *QAccelerometer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::deleteLater")
	}
}

func (ptr *QAccelerometer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAccelerometer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAccelerometer_DisconnectNotify
func callbackQAccelerometer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAccelerometerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAccelerometer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::disconnectNotify", f)
	}
}

func (ptr *QAccelerometer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::disconnectNotify")
	}
}

func (ptr *QAccelerometer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAccelerometer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAccelerometer_Event
func callbackQAccelerometer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAccelerometerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAccelerometer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::event", f)
	}
}

func (ptr *QAccelerometer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::event")
	}
}

func (ptr *QAccelerometer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAccelerometer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAccelerometer_EventFilter
func callbackQAccelerometer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAccelerometerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAccelerometer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::eventFilter", f)
	}
}

func (ptr *QAccelerometer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::eventFilter")
	}
}

func (ptr *QAccelerometer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAccelerometer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAccelerometer_MetaObject
func callbackQAccelerometer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAccelerometerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAccelerometer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::metaObject", f)
	}
}

func (ptr *QAccelerometer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometer::metaObject")
	}
}

func (ptr *QAccelerometer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAccelerometer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccelerometer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAccelerometer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAccelerometerFilter struct {
	QSensorFilter
}

type QAccelerometerFilter_ITF interface {
	QSensorFilter_ITF
	QAccelerometerFilter_PTR() *QAccelerometerFilter
}

func (p *QAccelerometerFilter) QAccelerometerFilter_PTR() *QAccelerometerFilter {
	return p
}

func (p *QAccelerometerFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QAccelerometerFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQAccelerometerFilter(ptr QAccelerometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilter_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFilterFromPointer(ptr unsafe.Pointer) *QAccelerometerFilter {
	var n = new(QAccelerometerFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccelerometerFilter) DestroyQAccelerometerFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQAccelerometerFilter_Filter
func callbackQAccelerometerFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QAccelerometerReading) bool)(NewQAccelerometerReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAccelerometerFilter) ConnectFilter(f func(reading *QAccelerometerReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerFilter::filter", f)
	}
}

func (ptr *QAccelerometerFilter) DisconnectFilter(reading QAccelerometerReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerFilter::filter")
	}
}

func (ptr *QAccelerometerFilter) Filter(reading QAccelerometerReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerFilter_Filter(ptr.Pointer(), PointerFromQAccelerometerReading(reading)) != 0
	}
	return false
}

type QAccelerometerReading struct {
	QSensorReading
}

type QAccelerometerReading_ITF interface {
	QSensorReading_ITF
	QAccelerometerReading_PTR() *QAccelerometerReading
}

func (p *QAccelerometerReading) QAccelerometerReading_PTR() *QAccelerometerReading {
	return p
}

func (p *QAccelerometerReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QAccelerometerReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQAccelerometerReading(ptr QAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerReadingFromPointer(ptr unsafe.Pointer) *QAccelerometerReading {
	var n = new(QAccelerometerReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccelerometerReading) DestroyQAccelerometerReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QAccelerometerReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QAccelerometerReading) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QAccelerometerReading) SetZ(z float64) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

//export callbackQAccelerometerReading_TimerEvent
func callbackQAccelerometerReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::timerEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::timerEvent")
	}
}

func (ptr *QAccelerometerReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometerReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAccelerometerReading_ChildEvent
func callbackQAccelerometerReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::childEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::childEvent")
	}
}

func (ptr *QAccelerometerReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometerReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAccelerometerReading_ConnectNotify
func callbackQAccelerometerReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAccelerometerReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::connectNotify", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::connectNotify")
	}
}

func (ptr *QAccelerometerReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAccelerometerReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAccelerometerReading_CustomEvent
func callbackQAccelerometerReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::customEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::customEvent")
	}
}

func (ptr *QAccelerometerReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAccelerometerReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAccelerometerReading_DeleteLater
func callbackQAccelerometerReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAccelerometerReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAccelerometerReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::deleteLater", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::deleteLater")
	}
}

func (ptr *QAccelerometerReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAccelerometerReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAccelerometerReading_DisconnectNotify
func callbackQAccelerometerReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAccelerometerReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::disconnectNotify", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::disconnectNotify")
	}
}

func (ptr *QAccelerometerReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAccelerometerReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAccelerometerReading_Event
func callbackQAccelerometerReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAccelerometerReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAccelerometerReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::event", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::event")
	}
}

func (ptr *QAccelerometerReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAccelerometerReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAccelerometerReading_EventFilter
func callbackQAccelerometerReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAccelerometerReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAccelerometerReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::eventFilter", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::eventFilter")
	}
}

func (ptr *QAccelerometerReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAccelerometerReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAccelerometerReading_MetaObject
func callbackQAccelerometerReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAccelerometerReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAccelerometerReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAccelerometerReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::metaObject", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAccelerometerReading::metaObject")
	}
}

func (ptr *QAccelerometerReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAccelerometerReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccelerometerReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAccelerometerReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAltimeter struct {
	QSensor
}

type QAltimeter_ITF interface {
	QSensor_ITF
	QAltimeter_PTR() *QAltimeter
}

func (p *QAltimeter) QAltimeter_PTR() *QAltimeter {
	return p
}

func (p *QAltimeter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QAltimeter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQAltimeter(ptr QAltimeter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFromPointer(ptr unsafe.Pointer) *QAltimeter {
	var n = new(QAltimeter)
	n.SetPointer(ptr)
	return n
}
func (ptr *QAltimeter) Reading() *QAltimeterReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAltimeterReadingFromPointer(C.QAltimeter_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQAltimeter(parent core.QObject_ITF) *QAltimeter {
	var tmpValue = NewQAltimeterFromPointer(C.QAltimeter_NewQAltimeter(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAltimeter) DestroyQAltimeter() {
	if ptr.Pointer() != nil {
		C.QAltimeter_DestroyQAltimeter(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QAltimeter_Type() string {
	return cGoUnpackString(C.QAltimeter_QAltimeter_Type())
}

func (ptr *QAltimeter) Type() string {
	return cGoUnpackString(C.QAltimeter_QAltimeter_Type())
}

//export callbackQAltimeter_Start
func callbackQAltimeter_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAltimeterFromPointer(ptr).StartDefault())))
}

func (ptr *QAltimeter) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::start", f)
	}
}

func (ptr *QAltimeter) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::start")
	}
}

func (ptr *QAltimeter) Start() bool {
	if ptr.Pointer() != nil {
		return C.QAltimeter_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAltimeter) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAltimeter_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAltimeter_Stop
func callbackQAltimeter_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQAltimeterFromPointer(ptr).StopDefault()
	}
}

func (ptr *QAltimeter) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::stop", f)
	}
}

func (ptr *QAltimeter) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::stop")
	}
}

func (ptr *QAltimeter) Stop() {
	if ptr.Pointer() != nil {
		C.QAltimeter_Stop(ptr.Pointer())
	}
}

func (ptr *QAltimeter) StopDefault() {
	if ptr.Pointer() != nil {
		C.QAltimeter_StopDefault(ptr.Pointer())
	}
}

//export callbackQAltimeter_TimerEvent
func callbackQAltimeter_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAltimeter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::timerEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::timerEvent")
	}
}

func (ptr *QAltimeter) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeter) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAltimeter_ChildEvent
func callbackQAltimeter_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAltimeter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::childEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::childEvent")
	}
}

func (ptr *QAltimeter) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeter) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAltimeter_ConnectNotify
func callbackQAltimeter_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAltimeterFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAltimeter) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::connectNotify", f)
	}
}

func (ptr *QAltimeter) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::connectNotify")
	}
}

func (ptr *QAltimeter) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAltimeter) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAltimeter_CustomEvent
func callbackQAltimeter_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAltimeter) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::customEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::customEvent")
	}
}

func (ptr *QAltimeter) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAltimeter) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAltimeter_DeleteLater
func callbackQAltimeter_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAltimeterFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAltimeter) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::deleteLater", f)
	}
}

func (ptr *QAltimeter) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::deleteLater")
	}
}

func (ptr *QAltimeter) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAltimeter_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAltimeter) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAltimeter_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAltimeter_DisconnectNotify
func callbackQAltimeter_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAltimeterFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAltimeter) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::disconnectNotify", f)
	}
}

func (ptr *QAltimeter) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::disconnectNotify")
	}
}

func (ptr *QAltimeter) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAltimeter) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeter_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAltimeter_Event
func callbackQAltimeter_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAltimeterFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAltimeter) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::event", f)
	}
}

func (ptr *QAltimeter) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::event")
	}
}

func (ptr *QAltimeter) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeter_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAltimeter) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeter_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAltimeter_EventFilter
func callbackQAltimeter_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAltimeterFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAltimeter) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::eventFilter", f)
	}
}

func (ptr *QAltimeter) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::eventFilter")
	}
}

func (ptr *QAltimeter) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeter_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAltimeter) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeter_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAltimeter_MetaObject
func callbackQAltimeter_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeter::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAltimeterFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAltimeter) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::metaObject", f)
	}
}

func (ptr *QAltimeter) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeter::metaObject")
	}
}

func (ptr *QAltimeter) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAltimeter_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAltimeter) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAltimeter_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAltimeterFilter struct {
	QSensorFilter
}

type QAltimeterFilter_ITF interface {
	QSensorFilter_ITF
	QAltimeterFilter_PTR() *QAltimeterFilter
}

func (p *QAltimeterFilter) QAltimeterFilter_PTR() *QAltimeterFilter {
	return p
}

func (p *QAltimeterFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QAltimeterFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQAltimeterFilter(ptr QAltimeterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterFilter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFilterFromPointer(ptr unsafe.Pointer) *QAltimeterFilter {
	var n = new(QAltimeterFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAltimeterFilter) DestroyQAltimeterFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQAltimeterFilter_Filter
func callbackQAltimeterFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QAltimeterReading) bool)(NewQAltimeterReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAltimeterFilter) ConnectFilter(f func(reading *QAltimeterReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterFilter::filter", f)
	}
}

func (ptr *QAltimeterFilter) DisconnectFilter(reading QAltimeterReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterFilter::filter")
	}
}

func (ptr *QAltimeterFilter) Filter(reading QAltimeterReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeterFilter_Filter(ptr.Pointer(), PointerFromQAltimeterReading(reading)) != 0
	}
	return false
}

type QAltimeterReading struct {
	QSensorReading
}

type QAltimeterReading_ITF interface {
	QSensorReading_ITF
	QAltimeterReading_PTR() *QAltimeterReading
}

func (p *QAltimeterReading) QAltimeterReading_PTR() *QAltimeterReading {
	return p
}

func (p *QAltimeterReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QAltimeterReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQAltimeterReading(ptr QAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReading_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterReadingFromPointer(ptr unsafe.Pointer) *QAltimeterReading {
	var n = new(QAltimeterReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAltimeterReading) DestroyQAltimeterReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QAltimeterReading) Altitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAltimeterReading_Altitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAltimeterReading) SetAltitude(altitude float64) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_SetAltitude(ptr.Pointer(), C.double(altitude))
	}
}

//export callbackQAltimeterReading_TimerEvent
func callbackQAltimeterReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::timerEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::timerEvent")
	}
}

func (ptr *QAltimeterReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeterReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAltimeterReading_ChildEvent
func callbackQAltimeterReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::childEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::childEvent")
	}
}

func (ptr *QAltimeterReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeterReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAltimeterReading_ConnectNotify
func callbackQAltimeterReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAltimeterReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAltimeterReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::connectNotify", f)
	}
}

func (ptr *QAltimeterReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::connectNotify")
	}
}

func (ptr *QAltimeterReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAltimeterReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAltimeterReading_CustomEvent
func callbackQAltimeterReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::customEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::customEvent")
	}
}

func (ptr *QAltimeterReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAltimeterReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAltimeterReading_DeleteLater
func callbackQAltimeterReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAltimeterReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAltimeterReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::deleteLater", f)
	}
}

func (ptr *QAltimeterReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::deleteLater")
	}
}

func (ptr *QAltimeterReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAltimeterReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAltimeterReading_DisconnectNotify
func callbackQAltimeterReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAltimeterReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAltimeterReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::disconnectNotify", f)
	}
}

func (ptr *QAltimeterReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::disconnectNotify")
	}
}

func (ptr *QAltimeterReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAltimeterReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAltimeterReading_Event
func callbackQAltimeterReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAltimeterReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAltimeterReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::event", f)
	}
}

func (ptr *QAltimeterReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::event")
	}
}

func (ptr *QAltimeterReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeterReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAltimeterReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeterReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAltimeterReading_EventFilter
func callbackQAltimeterReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAltimeterReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAltimeterReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::eventFilter", f)
	}
}

func (ptr *QAltimeterReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::eventFilter")
	}
}

func (ptr *QAltimeterReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeterReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAltimeterReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeterReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAltimeterReading_MetaObject
func callbackQAltimeterReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAltimeterReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAltimeterReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAltimeterReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::metaObject", f)
	}
}

func (ptr *QAltimeterReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAltimeterReading::metaObject")
	}
}

func (ptr *QAltimeterReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAltimeterReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAltimeterReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAltimeterReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAmbientLightFilter struct {
	QSensorFilter
}

type QAmbientLightFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientLightFilter_PTR() *QAmbientLightFilter
}

func (p *QAmbientLightFilter) QAmbientLightFilter_PTR() *QAmbientLightFilter {
	return p
}

func (p *QAmbientLightFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QAmbientLightFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQAmbientLightFilter(ptr QAmbientLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightFilter_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightFilterFromPointer(ptr unsafe.Pointer) *QAmbientLightFilter {
	var n = new(QAmbientLightFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientLightFilter) DestroyQAmbientLightFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQAmbientLightFilter_Filter
func callbackQAmbientLightFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QAmbientLightReading) bool)(NewQAmbientLightReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAmbientLightFilter) ConnectFilter(f func(reading *QAmbientLightReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightFilter::filter", f)
	}
}

func (ptr *QAmbientLightFilter) DisconnectFilter(reading QAmbientLightReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightFilter::filter")
	}
}

func (ptr *QAmbientLightFilter) Filter(reading QAmbientLightReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightFilter_Filter(ptr.Pointer(), PointerFromQAmbientLightReading(reading)) != 0
	}
	return false
}

//QAmbientLightReading::LightLevel
type QAmbientLightReading__LightLevel int64

const (
	QAmbientLightReading__Undefined = QAmbientLightReading__LightLevel(0)
	QAmbientLightReading__Dark      = QAmbientLightReading__LightLevel(1)
	QAmbientLightReading__Twilight  = QAmbientLightReading__LightLevel(2)
	QAmbientLightReading__Light     = QAmbientLightReading__LightLevel(3)
	QAmbientLightReading__Bright    = QAmbientLightReading__LightLevel(4)
	QAmbientLightReading__Sunny     = QAmbientLightReading__LightLevel(5)
)

type QAmbientLightReading struct {
	QSensorReading
}

type QAmbientLightReading_ITF interface {
	QSensorReading_ITF
	QAmbientLightReading_PTR() *QAmbientLightReading
}

func (p *QAmbientLightReading) QAmbientLightReading_PTR() *QAmbientLightReading {
	return p
}

func (p *QAmbientLightReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QAmbientLightReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQAmbientLightReading(ptr QAmbientLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightReadingFromPointer(ptr unsafe.Pointer) *QAmbientLightReading {
	var n = new(QAmbientLightReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientLightReading) DestroyQAmbientLightReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {
	if ptr.Pointer() != nil {
		return QAmbientLightReading__LightLevel(C.QAmbientLightReading_LightLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_SetLightLevel(ptr.Pointer(), C.longlong(lightLevel))
	}
}

//export callbackQAmbientLightReading_TimerEvent
func callbackQAmbientLightReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::timerEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::timerEvent")
	}
}

func (ptr *QAmbientLightReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAmbientLightReading_ChildEvent
func callbackQAmbientLightReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::childEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::childEvent")
	}
}

func (ptr *QAmbientLightReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAmbientLightReading_ConnectNotify
func callbackQAmbientLightReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientLightReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::connectNotify", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::connectNotify")
	}
}

func (ptr *QAmbientLightReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientLightReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientLightReading_CustomEvent
func callbackQAmbientLightReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::customEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::customEvent")
	}
}

func (ptr *QAmbientLightReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientLightReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAmbientLightReading_DeleteLater
func callbackQAmbientLightReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientLightReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAmbientLightReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::deleteLater", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::deleteLater")
	}
}

func (ptr *QAmbientLightReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientLightReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAmbientLightReading_DisconnectNotify
func callbackQAmbientLightReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientLightReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::disconnectNotify", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::disconnectNotify")
	}
}

func (ptr *QAmbientLightReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientLightReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientLightReading_Event
func callbackQAmbientLightReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientLightReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAmbientLightReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::event", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::event")
	}
}

func (ptr *QAmbientLightReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAmbientLightReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAmbientLightReading_EventFilter
func callbackQAmbientLightReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientLightReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAmbientLightReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::eventFilter", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::eventFilter")
	}
}

func (ptr *QAmbientLightReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAmbientLightReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAmbientLightReading_MetaObject
func callbackQAmbientLightReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAmbientLightReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAmbientLightReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::metaObject", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightReading::metaObject")
	}
}

func (ptr *QAmbientLightReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientLightReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAmbientLightReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientLightReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
}

func (p *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return p
}

func (p *QAmbientLightSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QAmbientLightSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightSensorFromPointer(ptr unsafe.Pointer) *QAmbientLightSensor {
	var n = new(QAmbientLightSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAmbientLightReadingFromPointer(C.QAmbientLightSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {
	var tmpValue = NewQAmbientLightSensorFromPointer(C.QAmbientLightSensor_NewQAmbientLightSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAmbientLightSensor_DestroyQAmbientLightSensor
func callbackQAmbientLightSensor_DestroyQAmbientLightSensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::~QAmbientLightSensor"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientLightSensorFromPointer(ptr).DestroyQAmbientLightSensorDefault()
	}
}

func (ptr *QAmbientLightSensor) ConnectDestroyQAmbientLightSensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::~QAmbientLightSensor", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectDestroyQAmbientLightSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::~QAmbientLightSensor")
	}
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensorDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QAmbientLightSensor_Type() string {
	return cGoUnpackString(C.QAmbientLightSensor_QAmbientLightSensor_Type())
}

func (ptr *QAmbientLightSensor) Type() string {
	return cGoUnpackString(C.QAmbientLightSensor_QAmbientLightSensor_Type())
}

//export callbackQAmbientLightSensor_Start
func callbackQAmbientLightSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientLightSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QAmbientLightSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::start", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::start")
	}
}

func (ptr *QAmbientLightSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAmbientLightSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAmbientLightSensor_Stop
func callbackQAmbientLightSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientLightSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QAmbientLightSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::stop", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::stop")
	}
}

func (ptr *QAmbientLightSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QAmbientLightSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQAmbientLightSensor_TimerEvent
func callbackQAmbientLightSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::timerEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::timerEvent")
	}
}

func (ptr *QAmbientLightSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAmbientLightSensor_ChildEvent
func callbackQAmbientLightSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::childEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::childEvent")
	}
}

func (ptr *QAmbientLightSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAmbientLightSensor_ConnectNotify
func callbackQAmbientLightSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientLightSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::connectNotify", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::connectNotify")
	}
}

func (ptr *QAmbientLightSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientLightSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientLightSensor_CustomEvent
func callbackQAmbientLightSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::customEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::customEvent")
	}
}

func (ptr *QAmbientLightSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientLightSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAmbientLightSensor_DeleteLater
func callbackQAmbientLightSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientLightSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAmbientLightSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::deleteLater", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::deleteLater")
	}
}

func (ptr *QAmbientLightSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientLightSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAmbientLightSensor_DisconnectNotify
func callbackQAmbientLightSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientLightSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::disconnectNotify", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::disconnectNotify")
	}
}

func (ptr *QAmbientLightSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientLightSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientLightSensor_Event
func callbackQAmbientLightSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientLightSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAmbientLightSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::event", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::event")
	}
}

func (ptr *QAmbientLightSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAmbientLightSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAmbientLightSensor_EventFilter
func callbackQAmbientLightSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientLightSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAmbientLightSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::eventFilter", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::eventFilter")
	}
}

func (ptr *QAmbientLightSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAmbientLightSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAmbientLightSensor_MetaObject
func callbackQAmbientLightSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientLightSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAmbientLightSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAmbientLightSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::metaObject", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientLightSensor::metaObject")
	}
}

func (ptr *QAmbientLightSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientLightSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAmbientLightSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientLightSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAmbientTemperatureFilter struct {
	QSensorFilter
}

type QAmbientTemperatureFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter
}

func (p *QAmbientTemperatureFilter) QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter {
	return p
}

func (p *QAmbientTemperatureFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QAmbientTemperatureFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQAmbientTemperatureFilter(ptr QAmbientTemperatureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureFilter_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureFilterFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureFilter {
	var n = new(QAmbientTemperatureFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientTemperatureFilter) DestroyQAmbientTemperatureFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQAmbientTemperatureFilter_Filter
func callbackQAmbientTemperatureFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QAmbientTemperatureReading) bool)(NewQAmbientTemperatureReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAmbientTemperatureFilter) ConnectFilter(f func(reading *QAmbientTemperatureReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureFilter::filter", f)
	}
}

func (ptr *QAmbientTemperatureFilter) DisconnectFilter(reading QAmbientTemperatureReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureFilter::filter")
	}
}

func (ptr *QAmbientTemperatureFilter) Filter(reading QAmbientTemperatureReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureFilter_Filter(ptr.Pointer(), PointerFromQAmbientTemperatureReading(reading)) != 0
	}
	return false
}

type QAmbientTemperatureReading struct {
	QSensorReading
}

type QAmbientTemperatureReading_ITF interface {
	QSensorReading_ITF
	QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading
}

func (p *QAmbientTemperatureReading) QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading {
	return p
}

func (p *QAmbientTemperatureReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QAmbientTemperatureReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureReading {
	var n = new(QAmbientTemperatureReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientTemperatureReading) DestroyQAmbientTemperatureReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QAmbientTemperatureReading) Temperature() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAmbientTemperatureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientTemperatureReading) SetTemperature(temperature float64) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

//export callbackQAmbientTemperatureReading_TimerEvent
func callbackQAmbientTemperatureReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::timerEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::timerEvent")
	}
}

func (ptr *QAmbientTemperatureReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAmbientTemperatureReading_ChildEvent
func callbackQAmbientTemperatureReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::childEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::childEvent")
	}
}

func (ptr *QAmbientTemperatureReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAmbientTemperatureReading_ConnectNotify
func callbackQAmbientTemperatureReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::connectNotify", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::connectNotify")
	}
}

func (ptr *QAmbientTemperatureReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientTemperatureReading_CustomEvent
func callbackQAmbientTemperatureReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::customEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::customEvent")
	}
}

func (ptr *QAmbientTemperatureReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAmbientTemperatureReading_DeleteLater
func callbackQAmbientTemperatureReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAmbientTemperatureReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::deleteLater", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::deleteLater")
	}
}

func (ptr *QAmbientTemperatureReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientTemperatureReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAmbientTemperatureReading_DisconnectNotify
func callbackQAmbientTemperatureReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::disconnectNotify", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::disconnectNotify")
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientTemperatureReading_Event
func callbackQAmbientTemperatureReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientTemperatureReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAmbientTemperatureReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::event", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::event")
	}
}

func (ptr *QAmbientTemperatureReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAmbientTemperatureReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAmbientTemperatureReading_EventFilter
func callbackQAmbientTemperatureReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientTemperatureReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAmbientTemperatureReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::eventFilter", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::eventFilter")
	}
}

func (ptr *QAmbientTemperatureReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAmbientTemperatureReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAmbientTemperatureReading_MetaObject
func callbackQAmbientTemperatureReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAmbientTemperatureReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAmbientTemperatureReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::metaObject", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureReading::metaObject")
	}
}

func (ptr *QAmbientTemperatureReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientTemperatureReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAmbientTemperatureReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientTemperatureReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAmbientTemperatureSensor struct {
	QSensor
}

type QAmbientTemperatureSensor_ITF interface {
	QSensor_ITF
	QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor
}

func (p *QAmbientTemperatureSensor) QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor {
	return p
}

func (p *QAmbientTemperatureSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QAmbientTemperatureSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQAmbientTemperatureSensor(ptr QAmbientTemperatureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureSensorFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureSensor {
	var n = new(QAmbientTemperatureSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QAmbientTemperatureSensor) Reading() *QAmbientTemperatureReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQAmbientTemperatureReadingFromPointer(C.QAmbientTemperatureSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQAmbientTemperatureSensor(parent core.QObject_ITF) *QAmbientTemperatureSensor {
	var tmpValue = NewQAmbientTemperatureSensorFromPointer(C.QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensor() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QAmbientTemperatureSensor_Type() string {
	return cGoUnpackString(C.QAmbientTemperatureSensor_QAmbientTemperatureSensor_Type())
}

func (ptr *QAmbientTemperatureSensor) Type() string {
	return cGoUnpackString(C.QAmbientTemperatureSensor_QAmbientTemperatureSensor_Type())
}

//export callbackQAmbientTemperatureSensor_Start
func callbackQAmbientTemperatureSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientTemperatureSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QAmbientTemperatureSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::start", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::start")
	}
}

func (ptr *QAmbientTemperatureSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAmbientTemperatureSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAmbientTemperatureSensor_Stop
func callbackQAmbientTemperatureSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::stop", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::stop")
	}
}

func (ptr *QAmbientTemperatureSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QAmbientTemperatureSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQAmbientTemperatureSensor_TimerEvent
func callbackQAmbientTemperatureSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::timerEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::timerEvent")
	}
}

func (ptr *QAmbientTemperatureSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAmbientTemperatureSensor_ChildEvent
func callbackQAmbientTemperatureSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::childEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::childEvent")
	}
}

func (ptr *QAmbientTemperatureSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAmbientTemperatureSensor_ConnectNotify
func callbackQAmbientTemperatureSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::connectNotify", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::connectNotify")
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientTemperatureSensor_CustomEvent
func callbackQAmbientTemperatureSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::customEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::customEvent")
	}
}

func (ptr *QAmbientTemperatureSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAmbientTemperatureSensor_DeleteLater
func callbackQAmbientTemperatureSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::deleteLater", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::deleteLater")
	}
}

func (ptr *QAmbientTemperatureSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientTemperatureSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAmbientTemperatureSensor_DisconnectNotify
func callbackQAmbientTemperatureSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::disconnectNotify", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::disconnectNotify")
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAmbientTemperatureSensor_Event
func callbackQAmbientTemperatureSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientTemperatureSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAmbientTemperatureSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::event", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::event")
	}
}

func (ptr *QAmbientTemperatureSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAmbientTemperatureSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAmbientTemperatureSensor_EventFilter
func callbackQAmbientTemperatureSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAmbientTemperatureSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAmbientTemperatureSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::eventFilter", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::eventFilter")
	}
}

func (ptr *QAmbientTemperatureSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAmbientTemperatureSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAmbientTemperatureSensor_MetaObject
func callbackQAmbientTemperatureSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAmbientTemperatureSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAmbientTemperatureSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAmbientTemperatureSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::metaObject", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAmbientTemperatureSensor::metaObject")
	}
}

func (ptr *QAmbientTemperatureSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientTemperatureSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAmbientTemperatureSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAmbientTemperatureSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
}

func (p *QCompass) QCompass_PTR() *QCompass {
	return p
}

func (p *QCompass) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QCompass) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQCompass(ptr QCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompass_PTR().Pointer()
	}
	return nil
}

func NewQCompassFromPointer(ptr unsafe.Pointer) *QCompass {
	var n = new(QCompass)
	n.SetPointer(ptr)
	return n
}
func (ptr *QCompass) Reading() *QCompassReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQCompassReadingFromPointer(C.QCompass_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQCompass(parent core.QObject_ITF) *QCompass {
	var tmpValue = NewQCompassFromPointer(C.QCompass_NewQCompass(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQCompass_DestroyQCompass
func callbackQCompass_DestroyQCompass(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::~QCompass"); signal != nil {
		signal.(func())()
	} else {
		NewQCompassFromPointer(ptr).DestroyQCompassDefault()
	}
}

func (ptr *QCompass) ConnectDestroyQCompass(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::~QCompass", f)
	}
}

func (ptr *QCompass) DisconnectDestroyQCompass() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::~QCompass")
	}
}

func (ptr *QCompass) DestroyQCompass() {
	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompass(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompass) DestroyQCompassDefault() {
	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompassDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QCompass_Type() string {
	return cGoUnpackString(C.QCompass_QCompass_Type())
}

func (ptr *QCompass) Type() string {
	return cGoUnpackString(C.QCompass_QCompass_Type())
}

//export callbackQCompass_Start
func callbackQCompass_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCompassFromPointer(ptr).StartDefault())))
}

func (ptr *QCompass) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::start", f)
	}
}

func (ptr *QCompass) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::start")
	}
}

func (ptr *QCompass) Start() bool {
	if ptr.Pointer() != nil {
		return C.QCompass_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCompass) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QCompass_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCompass_Stop
func callbackQCompass_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQCompassFromPointer(ptr).StopDefault()
	}
}

func (ptr *QCompass) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::stop", f)
	}
}

func (ptr *QCompass) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::stop")
	}
}

func (ptr *QCompass) Stop() {
	if ptr.Pointer() != nil {
		C.QCompass_Stop(ptr.Pointer())
	}
}

func (ptr *QCompass) StopDefault() {
	if ptr.Pointer() != nil {
		C.QCompass_StopDefault(ptr.Pointer())
	}
}

//export callbackQCompass_TimerEvent
func callbackQCompass_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCompass) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::timerEvent", f)
	}
}

func (ptr *QCompass) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::timerEvent")
	}
}

func (ptr *QCompass) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompass) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCompass_ChildEvent
func callbackQCompass_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCompass) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::childEvent", f)
	}
}

func (ptr *QCompass) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::childEvent")
	}
}

func (ptr *QCompass) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompass) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCompass_ConnectNotify
func callbackQCompass_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCompassFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCompass) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::connectNotify", f)
	}
}

func (ptr *QCompass) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::connectNotify")
	}
}

func (ptr *QCompass) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCompass) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCompass_CustomEvent
func callbackQCompass_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCompass) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::customEvent", f)
	}
}

func (ptr *QCompass) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::customEvent")
	}
}

func (ptr *QCompass) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCompass) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCompass_DeleteLater
func callbackQCompass_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCompassFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCompass) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::deleteLater", f)
	}
}

func (ptr *QCompass) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::deleteLater")
	}
}

func (ptr *QCompass) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QCompass_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompass) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QCompass_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQCompass_DisconnectNotify
func callbackQCompass_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCompassFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCompass) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::disconnectNotify", f)
	}
}

func (ptr *QCompass) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::disconnectNotify")
	}
}

func (ptr *QCompass) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCompass) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompass_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCompass_Event
func callbackQCompass_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCompassFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCompass) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::event", f)
	}
}

func (ptr *QCompass) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::event")
	}
}

func (ptr *QCompass) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompass_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCompass) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompass_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCompass_EventFilter
func callbackQCompass_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCompassFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCompass) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::eventFilter", f)
	}
}

func (ptr *QCompass) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::eventFilter")
	}
}

func (ptr *QCompass) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompass_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCompass) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompass_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCompass_MetaObject
func callbackQCompass_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompass::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCompassFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCompass) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::metaObject", f)
	}
}

func (ptr *QCompass) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompass::metaObject")
	}
}

func (ptr *QCompass) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCompass_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompass) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCompass_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QCompassFilter struct {
	QSensorFilter
}

type QCompassFilter_ITF interface {
	QSensorFilter_ITF
	QCompassFilter_PTR() *QCompassFilter
}

func (p *QCompassFilter) QCompassFilter_PTR() *QCompassFilter {
	return p
}

func (p *QCompassFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QCompassFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQCompassFilter(ptr QCompassFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassFilter_PTR().Pointer()
	}
	return nil
}

func NewQCompassFilterFromPointer(ptr unsafe.Pointer) *QCompassFilter {
	var n = new(QCompassFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCompassFilter) DestroyQCompassFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQCompassFilter_Filter
func callbackQCompassFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QCompassReading) bool)(NewQCompassReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QCompassFilter) ConnectFilter(f func(reading *QCompassReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassFilter::filter", f)
	}
}

func (ptr *QCompassFilter) DisconnectFilter(reading QCompassReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassFilter::filter")
	}
}

func (ptr *QCompassFilter) Filter(reading QCompassReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompassFilter_Filter(ptr.Pointer(), PointerFromQCompassReading(reading)) != 0
	}
	return false
}

type QCompassReading struct {
	QSensorReading
}

type QCompassReading_ITF interface {
	QSensorReading_ITF
	QCompassReading_PTR() *QCompassReading
}

func (p *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return p
}

func (p *QCompassReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QCompassReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQCompassReading(ptr QCompassReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassReading_PTR().Pointer()
	}
	return nil
}

func NewQCompassReadingFromPointer(ptr unsafe.Pointer) *QCompassReading {
	var n = new(QCompassReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCompassReading) DestroyQCompassReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QCompassReading) Azimuth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_Azimuth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) CalibrationLevel() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) SetAzimuth(azimuth float64) {
	if ptr.Pointer() != nil {
		C.QCompassReading_SetAzimuth(ptr.Pointer(), C.double(azimuth))
	}
}

func (ptr *QCompassReading) SetCalibrationLevel(calibrationLevel float64) {
	if ptr.Pointer() != nil {
		C.QCompassReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

//export callbackQCompassReading_TimerEvent
func callbackQCompassReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCompassReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::timerEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::timerEvent")
	}
}

func (ptr *QCompassReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompassReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCompassReading_ChildEvent
func callbackQCompassReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCompassReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::childEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::childEvent")
	}
}

func (ptr *QCompassReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompassReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCompassReading_ConnectNotify
func callbackQCompassReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCompassReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCompassReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::connectNotify", f)
	}
}

func (ptr *QCompassReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::connectNotify")
	}
}

func (ptr *QCompassReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCompassReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCompassReading_CustomEvent
func callbackQCompassReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCompassReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::customEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::customEvent")
	}
}

func (ptr *QCompassReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCompassReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCompassReading_DeleteLater
func callbackQCompassReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCompassReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCompassReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::deleteLater", f)
	}
}

func (ptr *QCompassReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::deleteLater")
	}
}

func (ptr *QCompassReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QCompassReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompassReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QCompassReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQCompassReading_DisconnectNotify
func callbackQCompassReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCompassReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCompassReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::disconnectNotify", f)
	}
}

func (ptr *QCompassReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::disconnectNotify")
	}
}

func (ptr *QCompassReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCompassReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCompassReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCompassReading_Event
func callbackQCompassReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCompassReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCompassReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::event", f)
	}
}

func (ptr *QCompassReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::event")
	}
}

func (ptr *QCompassReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompassReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCompassReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompassReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCompassReading_EventFilter
func callbackQCompassReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCompassReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCompassReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::eventFilter", f)
	}
}

func (ptr *QCompassReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::eventFilter")
	}
}

func (ptr *QCompassReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompassReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCompassReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompassReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCompassReading_MetaObject
func callbackQCompassReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCompassReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCompassReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCompassReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::metaObject", f)
	}
}

func (ptr *QCompassReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCompassReading::metaObject")
	}
}

func (ptr *QCompassReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCompassReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompassReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCompassReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDistanceFilter struct {
	QSensorFilter
}

type QDistanceFilter_ITF interface {
	QSensorFilter_ITF
	QDistanceFilter_PTR() *QDistanceFilter
}

func (p *QDistanceFilter) QDistanceFilter_PTR() *QDistanceFilter {
	return p
}

func (p *QDistanceFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QDistanceFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQDistanceFilter(ptr QDistanceFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceFilter_PTR().Pointer()
	}
	return nil
}

func NewQDistanceFilterFromPointer(ptr unsafe.Pointer) *QDistanceFilter {
	var n = new(QDistanceFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDistanceFilter) DestroyQDistanceFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQDistanceFilter_Filter
func callbackQDistanceFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QDistanceReading) bool)(NewQDistanceReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDistanceFilter) ConnectFilter(f func(reading *QDistanceReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceFilter::filter", f)
	}
}

func (ptr *QDistanceFilter) DisconnectFilter(reading QDistanceReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceFilter::filter")
	}
}

func (ptr *QDistanceFilter) Filter(reading QDistanceReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceFilter_Filter(ptr.Pointer(), PointerFromQDistanceReading(reading)) != 0
	}
	return false
}

type QDistanceReading struct {
	QSensorReading
}

type QDistanceReading_ITF interface {
	QSensorReading_ITF
	QDistanceReading_PTR() *QDistanceReading
}

func (p *QDistanceReading) QDistanceReading_PTR() *QDistanceReading {
	return p
}

func (p *QDistanceReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QDistanceReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQDistanceReading(ptr QDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReading_PTR().Pointer()
	}
	return nil
}

func NewQDistanceReadingFromPointer(ptr unsafe.Pointer) *QDistanceReading {
	var n = new(QDistanceReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDistanceReading) DestroyQDistanceReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QDistanceReading) Distance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QDistanceReading_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDistanceReading) SetDistance(distance float64) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

//export callbackQDistanceReading_TimerEvent
func callbackQDistanceReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::timerEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::timerEvent")
	}
}

func (ptr *QDistanceReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDistanceReading_ChildEvent
func callbackQDistanceReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::childEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::childEvent")
	}
}

func (ptr *QDistanceReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDistanceReading_ConnectNotify
func callbackQDistanceReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDistanceReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDistanceReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::connectNotify", f)
	}
}

func (ptr *QDistanceReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::connectNotify")
	}
}

func (ptr *QDistanceReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDistanceReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDistanceReading_CustomEvent
func callbackQDistanceReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::customEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::customEvent")
	}
}

func (ptr *QDistanceReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDistanceReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDistanceReading_DeleteLater
func callbackQDistanceReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDistanceReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDistanceReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::deleteLater", f)
	}
}

func (ptr *QDistanceReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::deleteLater")
	}
}

func (ptr *QDistanceReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDistanceReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDistanceReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDistanceReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDistanceReading_DisconnectNotify
func callbackQDistanceReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDistanceReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDistanceReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::disconnectNotify", f)
	}
}

func (ptr *QDistanceReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::disconnectNotify")
	}
}

func (ptr *QDistanceReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDistanceReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDistanceReading_Event
func callbackQDistanceReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDistanceReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDistanceReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::event", f)
	}
}

func (ptr *QDistanceReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::event")
	}
}

func (ptr *QDistanceReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDistanceReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDistanceReading_EventFilter
func callbackQDistanceReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDistanceReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDistanceReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::eventFilter", f)
	}
}

func (ptr *QDistanceReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::eventFilter")
	}
}

func (ptr *QDistanceReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDistanceReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDistanceReading_MetaObject
func callbackQDistanceReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDistanceReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDistanceReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::metaObject", f)
	}
}

func (ptr *QDistanceReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceReading::metaObject")
	}
}

func (ptr *QDistanceReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDistanceReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDistanceReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDistanceReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensor_ITF interface {
	QSensor_ITF
	QDistanceSensor_PTR() *QDistanceSensor
}

func (p *QDistanceSensor) QDistanceSensor_PTR() *QDistanceSensor {
	return p
}

func (p *QDistanceSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QDistanceSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQDistanceSensor(ptr QDistanceSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceSensor_PTR().Pointer()
	}
	return nil
}

func NewQDistanceSensorFromPointer(ptr unsafe.Pointer) *QDistanceSensor {
	var n = new(QDistanceSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QDistanceSensor) Reading() *QDistanceReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDistanceReadingFromPointer(C.QDistanceSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQDistanceSensor(parent core.QObject_ITF) *QDistanceSensor {
	var tmpValue = NewQDistanceSensorFromPointer(C.QDistanceSensor_NewQDistanceSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_DestroyQDistanceSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QDistanceSensor_Type() string {
	return cGoUnpackString(C.QDistanceSensor_QDistanceSensor_Type())
}

func (ptr *QDistanceSensor) Type() string {
	return cGoUnpackString(C.QDistanceSensor_QDistanceSensor_Type())
}

//export callbackQDistanceSensor_Start
func callbackQDistanceSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDistanceSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QDistanceSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::start", f)
	}
}

func (ptr *QDistanceSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::start")
	}
}

func (ptr *QDistanceSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QDistanceSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDistanceSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDistanceSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDistanceSensor_Stop
func callbackQDistanceSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQDistanceSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QDistanceSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::stop", f)
	}
}

func (ptr *QDistanceSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::stop")
	}
}

func (ptr *QDistanceSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QDistanceSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQDistanceSensor_TimerEvent
func callbackQDistanceSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::timerEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::timerEvent")
	}
}

func (ptr *QDistanceSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDistanceSensor_ChildEvent
func callbackQDistanceSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::childEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::childEvent")
	}
}

func (ptr *QDistanceSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDistanceSensor_ConnectNotify
func callbackQDistanceSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDistanceSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDistanceSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::connectNotify", f)
	}
}

func (ptr *QDistanceSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::connectNotify")
	}
}

func (ptr *QDistanceSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDistanceSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDistanceSensor_CustomEvent
func callbackQDistanceSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::customEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::customEvent")
	}
}

func (ptr *QDistanceSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDistanceSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDistanceSensor_DeleteLater
func callbackQDistanceSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDistanceSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDistanceSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::deleteLater", f)
	}
}

func (ptr *QDistanceSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::deleteLater")
	}
}

func (ptr *QDistanceSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDistanceSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDistanceSensor_DisconnectNotify
func callbackQDistanceSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDistanceSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDistanceSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::disconnectNotify", f)
	}
}

func (ptr *QDistanceSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::disconnectNotify")
	}
}

func (ptr *QDistanceSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDistanceSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDistanceSensor_Event
func callbackQDistanceSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDistanceSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDistanceSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::event", f)
	}
}

func (ptr *QDistanceSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::event")
	}
}

func (ptr *QDistanceSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDistanceSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDistanceSensor_EventFilter
func callbackQDistanceSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDistanceSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDistanceSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::eventFilter", f)
	}
}

func (ptr *QDistanceSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::eventFilter")
	}
}

func (ptr *QDistanceSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDistanceSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDistanceSensor_MetaObject
func callbackQDistanceSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDistanceSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDistanceSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDistanceSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::metaObject", f)
	}
}

func (ptr *QDistanceSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDistanceSensor::metaObject")
	}
}

func (ptr *QDistanceSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDistanceSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDistanceSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDistanceSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QGyroscope struct {
	QSensor
}

type QGyroscope_ITF interface {
	QSensor_ITF
	QGyroscope_PTR() *QGyroscope
}

func (p *QGyroscope) QGyroscope_PTR() *QGyroscope {
	return p
}

func (p *QGyroscope) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QGyroscope) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQGyroscope(ptr QGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscope_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeFromPointer(ptr unsafe.Pointer) *QGyroscope {
	var n = new(QGyroscope)
	n.SetPointer(ptr)
	return n
}
func (ptr *QGyroscope) Reading() *QGyroscopeReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGyroscopeReadingFromPointer(C.QGyroscope_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQGyroscope(parent core.QObject_ITF) *QGyroscope {
	var tmpValue = NewQGyroscopeFromPointer(C.QGyroscope_NewQGyroscope(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGyroscope_DestroyQGyroscope
func callbackQGyroscope_DestroyQGyroscope(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::~QGyroscope"); signal != nil {
		signal.(func())()
	} else {
		NewQGyroscopeFromPointer(ptr).DestroyQGyroscopeDefault()
	}
}

func (ptr *QGyroscope) ConnectDestroyQGyroscope(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::~QGyroscope", f)
	}
}

func (ptr *QGyroscope) DisconnectDestroyQGyroscope() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::~QGyroscope")
	}
}

func (ptr *QGyroscope) DestroyQGyroscope() {
	if ptr.Pointer() != nil {
		C.QGyroscope_DestroyQGyroscope(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGyroscope) DestroyQGyroscopeDefault() {
	if ptr.Pointer() != nil {
		C.QGyroscope_DestroyQGyroscopeDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QGyroscope_Type() string {
	return cGoUnpackString(C.QGyroscope_QGyroscope_Type())
}

func (ptr *QGyroscope) Type() string {
	return cGoUnpackString(C.QGyroscope_QGyroscope_Type())
}

//export callbackQGyroscope_Start
func callbackQGyroscope_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGyroscopeFromPointer(ptr).StartDefault())))
}

func (ptr *QGyroscope) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::start", f)
	}
}

func (ptr *QGyroscope) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::start")
	}
}

func (ptr *QGyroscope) Start() bool {
	if ptr.Pointer() != nil {
		return C.QGyroscope_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGyroscope) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QGyroscope_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQGyroscope_Stop
func callbackQGyroscope_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQGyroscopeFromPointer(ptr).StopDefault()
	}
}

func (ptr *QGyroscope) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::stop", f)
	}
}

func (ptr *QGyroscope) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::stop")
	}
}

func (ptr *QGyroscope) Stop() {
	if ptr.Pointer() != nil {
		C.QGyroscope_Stop(ptr.Pointer())
	}
}

func (ptr *QGyroscope) StopDefault() {
	if ptr.Pointer() != nil {
		C.QGyroscope_StopDefault(ptr.Pointer())
	}
}

//export callbackQGyroscope_TimerEvent
func callbackQGyroscope_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGyroscope) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::timerEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::timerEvent")
	}
}

func (ptr *QGyroscope) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscope) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGyroscope_ChildEvent
func callbackQGyroscope_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGyroscope) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::childEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::childEvent")
	}
}

func (ptr *QGyroscope) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscope) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGyroscope_ConnectNotify
func callbackQGyroscope_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGyroscopeFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGyroscope) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::connectNotify", f)
	}
}

func (ptr *QGyroscope) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::connectNotify")
	}
}

func (ptr *QGyroscope) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGyroscope) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGyroscope_CustomEvent
func callbackQGyroscope_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGyroscope) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::customEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::customEvent")
	}
}

func (ptr *QGyroscope) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGyroscope) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGyroscope_DeleteLater
func callbackQGyroscope_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGyroscopeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGyroscope) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::deleteLater", f)
	}
}

func (ptr *QGyroscope) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::deleteLater")
	}
}

func (ptr *QGyroscope) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QGyroscope_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGyroscope) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGyroscope_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGyroscope_DisconnectNotify
func callbackQGyroscope_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGyroscopeFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGyroscope) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::disconnectNotify", f)
	}
}

func (ptr *QGyroscope) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::disconnectNotify")
	}
}

func (ptr *QGyroscope) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGyroscope) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscope_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGyroscope_Event
func callbackQGyroscope_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGyroscopeFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGyroscope) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::event", f)
	}
}

func (ptr *QGyroscope) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::event")
	}
}

func (ptr *QGyroscope) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscope_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGyroscope) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscope_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGyroscope_EventFilter
func callbackQGyroscope_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGyroscopeFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGyroscope) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::eventFilter", f)
	}
}

func (ptr *QGyroscope) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::eventFilter")
	}
}

func (ptr *QGyroscope) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscope_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGyroscope) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscope_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGyroscope_MetaObject
func callbackQGyroscope_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscope::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGyroscopeFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGyroscope) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::metaObject", f)
	}
}

func (ptr *QGyroscope) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscope::metaObject")
	}
}

func (ptr *QGyroscope) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGyroscope_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGyroscope) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGyroscope_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QGyroscopeFilter struct {
	QSensorFilter
}

type QGyroscopeFilter_ITF interface {
	QSensorFilter_ITF
	QGyroscopeFilter_PTR() *QGyroscopeFilter
}

func (p *QGyroscopeFilter) QGyroscopeFilter_PTR() *QGyroscopeFilter {
	return p
}

func (p *QGyroscopeFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QGyroscopeFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQGyroscopeFilter(ptr QGyroscopeFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeFilter_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeFilterFromPointer(ptr unsafe.Pointer) *QGyroscopeFilter {
	var n = new(QGyroscopeFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGyroscopeFilter) DestroyQGyroscopeFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQGyroscopeFilter_Filter
func callbackQGyroscopeFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGyroscopeReading) bool)(NewQGyroscopeReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGyroscopeFilter) ConnectFilter(f func(reading *QGyroscopeReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeFilter::filter", f)
	}
}

func (ptr *QGyroscopeFilter) DisconnectFilter(reading QGyroscopeReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeFilter::filter")
	}
}

func (ptr *QGyroscopeFilter) Filter(reading QGyroscopeReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscopeFilter_Filter(ptr.Pointer(), PointerFromQGyroscopeReading(reading)) != 0
	}
	return false
}

type QGyroscopeReading struct {
	QSensorReading
}

type QGyroscopeReading_ITF interface {
	QSensorReading_ITF
	QGyroscopeReading_PTR() *QGyroscopeReading
}

func (p *QGyroscopeReading) QGyroscopeReading_PTR() *QGyroscopeReading {
	return p
}

func (p *QGyroscopeReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QGyroscopeReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQGyroscopeReading(ptr QGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeReadingFromPointer(ptr unsafe.Pointer) *QGyroscopeReading {
	var n = new(QGyroscopeReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGyroscopeReading) DestroyQGyroscopeReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QGyroscopeReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QGyroscopeReading) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QGyroscopeReading) SetZ(z float64) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

//export callbackQGyroscopeReading_TimerEvent
func callbackQGyroscopeReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGyroscopeReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::timerEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::timerEvent")
	}
}

func (ptr *QGyroscopeReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscopeReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGyroscopeReading_ChildEvent
func callbackQGyroscopeReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGyroscopeReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::childEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::childEvent")
	}
}

func (ptr *QGyroscopeReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscopeReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGyroscopeReading_ConnectNotify
func callbackQGyroscopeReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGyroscopeReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::connectNotify", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::connectNotify")
	}
}

func (ptr *QGyroscopeReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGyroscopeReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGyroscopeReading_CustomEvent
func callbackQGyroscopeReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGyroscopeReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::customEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::customEvent")
	}
}

func (ptr *QGyroscopeReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGyroscopeReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGyroscopeReading_DeleteLater
func callbackQGyroscopeReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGyroscopeReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGyroscopeReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::deleteLater", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::deleteLater")
	}
}

func (ptr *QGyroscopeReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGyroscopeReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGyroscopeReading_DisconnectNotify
func callbackQGyroscopeReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGyroscopeReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::disconnectNotify", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::disconnectNotify")
	}
}

func (ptr *QGyroscopeReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGyroscopeReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGyroscopeReading_Event
func callbackQGyroscopeReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGyroscopeReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGyroscopeReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::event", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::event")
	}
}

func (ptr *QGyroscopeReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscopeReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGyroscopeReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscopeReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGyroscopeReading_EventFilter
func callbackQGyroscopeReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGyroscopeReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGyroscopeReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::eventFilter", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::eventFilter")
	}
}

func (ptr *QGyroscopeReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscopeReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGyroscopeReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscopeReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGyroscopeReading_MetaObject
func callbackQGyroscopeReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGyroscopeReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGyroscopeReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGyroscopeReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::metaObject", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGyroscopeReading::metaObject")
	}
}

func (ptr *QGyroscopeReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGyroscopeReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGyroscopeReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGyroscopeReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHolsterFilter struct {
	QSensorFilter
}

type QHolsterFilter_ITF interface {
	QSensorFilter_ITF
	QHolsterFilter_PTR() *QHolsterFilter
}

func (p *QHolsterFilter) QHolsterFilter_PTR() *QHolsterFilter {
	return p
}

func (p *QHolsterFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QHolsterFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQHolsterFilter(ptr QHolsterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterFilter_PTR().Pointer()
	}
	return nil
}

func NewQHolsterFilterFromPointer(ptr unsafe.Pointer) *QHolsterFilter {
	var n = new(QHolsterFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHolsterFilter) DestroyQHolsterFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQHolsterFilter_Filter
func callbackQHolsterFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QHolsterReading) bool)(NewQHolsterReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QHolsterFilter) ConnectFilter(f func(reading *QHolsterReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterFilter::filter", f)
	}
}

func (ptr *QHolsterFilter) DisconnectFilter(reading QHolsterReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterFilter::filter")
	}
}

func (ptr *QHolsterFilter) Filter(reading QHolsterReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterFilter_Filter(ptr.Pointer(), PointerFromQHolsterReading(reading)) != 0
	}
	return false
}

type QHolsterReading struct {
	QSensorReading
}

type QHolsterReading_ITF interface {
	QSensorReading_ITF
	QHolsterReading_PTR() *QHolsterReading
}

func (p *QHolsterReading) QHolsterReading_PTR() *QHolsterReading {
	return p
}

func (p *QHolsterReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QHolsterReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQHolsterReading(ptr QHolsterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterReading_PTR().Pointer()
	}
	return nil
}

func NewQHolsterReadingFromPointer(ptr unsafe.Pointer) *QHolsterReading {
	var n = new(QHolsterReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHolsterReading) DestroyQHolsterReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QHolsterReading) Holstered() bool {
	if ptr.Pointer() != nil {
		return C.QHolsterReading_Holstered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHolsterReading) SetHolstered(holstered bool) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_SetHolstered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(holstered))))
	}
}

//export callbackQHolsterReading_TimerEvent
func callbackQHolsterReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHolsterReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHolsterReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::timerEvent", f)
	}
}

func (ptr *QHolsterReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::timerEvent")
	}
}

func (ptr *QHolsterReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHolsterReading_ChildEvent
func callbackQHolsterReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHolsterReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHolsterReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::childEvent", f)
	}
}

func (ptr *QHolsterReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::childEvent")
	}
}

func (ptr *QHolsterReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHolsterReading_ConnectNotify
func callbackQHolsterReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHolsterReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHolsterReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::connectNotify", f)
	}
}

func (ptr *QHolsterReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::connectNotify")
	}
}

func (ptr *QHolsterReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHolsterReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHolsterReading_CustomEvent
func callbackQHolsterReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHolsterReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHolsterReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::customEvent", f)
	}
}

func (ptr *QHolsterReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::customEvent")
	}
}

func (ptr *QHolsterReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHolsterReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHolsterReading_DeleteLater
func callbackQHolsterReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHolsterReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHolsterReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::deleteLater", f)
	}
}

func (ptr *QHolsterReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::deleteLater")
	}
}

func (ptr *QHolsterReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHolsterReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHolsterReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHolsterReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHolsterReading_DisconnectNotify
func callbackQHolsterReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHolsterReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHolsterReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::disconnectNotify", f)
	}
}

func (ptr *QHolsterReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::disconnectNotify")
	}
}

func (ptr *QHolsterReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHolsterReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHolsterReading_Event
func callbackQHolsterReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHolsterReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHolsterReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::event", f)
	}
}

func (ptr *QHolsterReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::event")
	}
}

func (ptr *QHolsterReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHolsterReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHolsterReading_EventFilter
func callbackQHolsterReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHolsterReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHolsterReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::eventFilter", f)
	}
}

func (ptr *QHolsterReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::eventFilter")
	}
}

func (ptr *QHolsterReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHolsterReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHolsterReading_MetaObject
func callbackQHolsterReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHolsterReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHolsterReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::metaObject", f)
	}
}

func (ptr *QHolsterReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterReading::metaObject")
	}
}

func (ptr *QHolsterReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHolsterReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHolsterReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHolsterReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHolsterSensor struct {
	QSensor
}

type QHolsterSensor_ITF interface {
	QSensor_ITF
	QHolsterSensor_PTR() *QHolsterSensor
}

func (p *QHolsterSensor) QHolsterSensor_PTR() *QHolsterSensor {
	return p
}

func (p *QHolsterSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QHolsterSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQHolsterSensor(ptr QHolsterSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterSensor_PTR().Pointer()
	}
	return nil
}

func NewQHolsterSensorFromPointer(ptr unsafe.Pointer) *QHolsterSensor {
	var n = new(QHolsterSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHolsterSensor) Reading() *QHolsterReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHolsterReadingFromPointer(C.QHolsterSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQHolsterSensor(parent core.QObject_ITF) *QHolsterSensor {
	var tmpValue = NewQHolsterSensorFromPointer(C.QHolsterSensor_NewQHolsterSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DestroyQHolsterSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QHolsterSensor_Type() string {
	return cGoUnpackString(C.QHolsterSensor_QHolsterSensor_Type())
}

func (ptr *QHolsterSensor) Type() string {
	return cGoUnpackString(C.QHolsterSensor_QHolsterSensor_Type())
}

//export callbackQHolsterSensor_Start
func callbackQHolsterSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHolsterSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QHolsterSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::start", f)
	}
}

func (ptr *QHolsterSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::start")
	}
}

func (ptr *QHolsterSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QHolsterSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHolsterSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHolsterSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHolsterSensor_Stop
func callbackQHolsterSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQHolsterSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QHolsterSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::stop", f)
	}
}

func (ptr *QHolsterSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::stop")
	}
}

func (ptr *QHolsterSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QHolsterSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQHolsterSensor_TimerEvent
func callbackQHolsterSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::timerEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::timerEvent")
	}
}

func (ptr *QHolsterSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHolsterSensor_ChildEvent
func callbackQHolsterSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::childEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::childEvent")
	}
}

func (ptr *QHolsterSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHolsterSensor_ConnectNotify
func callbackQHolsterSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHolsterSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHolsterSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::connectNotify", f)
	}
}

func (ptr *QHolsterSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::connectNotify")
	}
}

func (ptr *QHolsterSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHolsterSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHolsterSensor_CustomEvent
func callbackQHolsterSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::customEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::customEvent")
	}
}

func (ptr *QHolsterSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHolsterSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHolsterSensor_DeleteLater
func callbackQHolsterSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHolsterSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHolsterSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::deleteLater", f)
	}
}

func (ptr *QHolsterSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::deleteLater")
	}
}

func (ptr *QHolsterSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHolsterSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHolsterSensor_DisconnectNotify
func callbackQHolsterSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHolsterSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHolsterSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::disconnectNotify", f)
	}
}

func (ptr *QHolsterSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::disconnectNotify")
	}
}

func (ptr *QHolsterSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHolsterSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHolsterSensor_Event
func callbackQHolsterSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHolsterSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHolsterSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::event", f)
	}
}

func (ptr *QHolsterSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::event")
	}
}

func (ptr *QHolsterSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHolsterSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHolsterSensor_EventFilter
func callbackQHolsterSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHolsterSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHolsterSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::eventFilter", f)
	}
}

func (ptr *QHolsterSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::eventFilter")
	}
}

func (ptr *QHolsterSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHolsterSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHolsterSensor_MetaObject
func callbackQHolsterSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHolsterSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHolsterSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHolsterSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::metaObject", f)
	}
}

func (ptr *QHolsterSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHolsterSensor::metaObject")
	}
}

func (ptr *QHolsterSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHolsterSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHolsterSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHolsterSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QIRProximityFilter struct {
	QSensorFilter
}

type QIRProximityFilter_ITF interface {
	QSensorFilter_ITF
	QIRProximityFilter_PTR() *QIRProximityFilter
}

func (p *QIRProximityFilter) QIRProximityFilter_PTR() *QIRProximityFilter {
	return p
}

func (p *QIRProximityFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QIRProximityFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQIRProximityFilter(ptr QIRProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityFilter_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityFilterFromPointer(ptr unsafe.Pointer) *QIRProximityFilter {
	var n = new(QIRProximityFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIRProximityFilter) DestroyQIRProximityFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQIRProximityFilter_Filter
func callbackQIRProximityFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QIRProximityReading) bool)(NewQIRProximityReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QIRProximityFilter) ConnectFilter(f func(reading *QIRProximityReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityFilter::filter", f)
	}
}

func (ptr *QIRProximityFilter) DisconnectFilter(reading QIRProximityReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityFilter::filter")
	}
}

func (ptr *QIRProximityFilter) Filter(reading QIRProximityReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximityFilter_Filter(ptr.Pointer(), PointerFromQIRProximityReading(reading)) != 0
	}
	return false
}

type QIRProximityReading struct {
	QSensorReading
}

type QIRProximityReading_ITF interface {
	QSensorReading_ITF
	QIRProximityReading_PTR() *QIRProximityReading
}

func (p *QIRProximityReading) QIRProximityReading_PTR() *QIRProximityReading {
	return p
}

func (p *QIRProximityReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QIRProximityReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQIRProximityReading(ptr QIRProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityReadingFromPointer(ptr unsafe.Pointer) *QIRProximityReading {
	var n = new(QIRProximityReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIRProximityReading) DestroyQIRProximityReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QIRProximityReading) Reflectance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QIRProximityReading_Reflectance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIRProximityReading) SetReflectance(reflectance float64) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_SetReflectance(ptr.Pointer(), C.double(reflectance))
	}
}

//export callbackQIRProximityReading_TimerEvent
func callbackQIRProximityReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::timerEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::timerEvent")
	}
}

func (ptr *QIRProximityReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximityReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQIRProximityReading_ChildEvent
func callbackQIRProximityReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::childEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::childEvent")
	}
}

func (ptr *QIRProximityReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximityReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQIRProximityReading_ConnectNotify
func callbackQIRProximityReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQIRProximityReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QIRProximityReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::connectNotify", f)
	}
}

func (ptr *QIRProximityReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::connectNotify")
	}
}

func (ptr *QIRProximityReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QIRProximityReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQIRProximityReading_CustomEvent
func callbackQIRProximityReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::customEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::customEvent")
	}
}

func (ptr *QIRProximityReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIRProximityReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQIRProximityReading_DeleteLater
func callbackQIRProximityReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQIRProximityReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QIRProximityReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::deleteLater", f)
	}
}

func (ptr *QIRProximityReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::deleteLater")
	}
}

func (ptr *QIRProximityReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QIRProximityReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQIRProximityReading_DisconnectNotify
func callbackQIRProximityReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQIRProximityReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QIRProximityReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::disconnectNotify", f)
	}
}

func (ptr *QIRProximityReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::disconnectNotify")
	}
}

func (ptr *QIRProximityReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QIRProximityReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQIRProximityReading_Event
func callbackQIRProximityReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIRProximityReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QIRProximityReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::event", f)
	}
}

func (ptr *QIRProximityReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::event")
	}
}

func (ptr *QIRProximityReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximityReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QIRProximityReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximityReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQIRProximityReading_EventFilter
func callbackQIRProximityReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIRProximityReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QIRProximityReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::eventFilter", f)
	}
}

func (ptr *QIRProximityReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::eventFilter")
	}
}

func (ptr *QIRProximityReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximityReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QIRProximityReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximityReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQIRProximityReading_MetaObject
func callbackQIRProximityReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximityReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQIRProximityReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QIRProximityReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::metaObject", f)
	}
}

func (ptr *QIRProximityReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximityReading::metaObject")
	}
}

func (ptr *QIRProximityReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QIRProximityReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIRProximityReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QIRProximityReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensor_ITF interface {
	QSensor_ITF
	QIRProximitySensor_PTR() *QIRProximitySensor
}

func (p *QIRProximitySensor) QIRProximitySensor_PTR() *QIRProximitySensor {
	return p
}

func (p *QIRProximitySensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QIRProximitySensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQIRProximitySensor(ptr QIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQIRProximitySensorFromPointer(ptr unsafe.Pointer) *QIRProximitySensor {
	var n = new(QIRProximitySensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQIRProximityReadingFromPointer(C.QIRProximitySensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQIRProximitySensor(parent core.QObject_ITF) *QIRProximitySensor {
	var tmpValue = NewQIRProximitySensorFromPointer(C.QIRProximitySensor_NewQIRProximitySensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQIRProximitySensor_DestroyQIRProximitySensor
func callbackQIRProximitySensor_DestroyQIRProximitySensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::~QIRProximitySensor"); signal != nil {
		signal.(func())()
	} else {
		NewQIRProximitySensorFromPointer(ptr).DestroyQIRProximitySensorDefault()
	}
}

func (ptr *QIRProximitySensor) ConnectDestroyQIRProximitySensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::~QIRProximitySensor", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectDestroyQIRProximitySensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::~QIRProximitySensor")
	}
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DestroyQIRProximitySensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensorDefault() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DestroyQIRProximitySensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QIRProximitySensor_Type() string {
	return cGoUnpackString(C.QIRProximitySensor_QIRProximitySensor_Type())
}

func (ptr *QIRProximitySensor) Type() string {
	return cGoUnpackString(C.QIRProximitySensor_QIRProximitySensor_Type())
}

//export callbackQIRProximitySensor_Start
func callbackQIRProximitySensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIRProximitySensorFromPointer(ptr).StartDefault())))
}

func (ptr *QIRProximitySensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::start", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::start")
	}
}

func (ptr *QIRProximitySensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QIRProximitySensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIRProximitySensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QIRProximitySensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQIRProximitySensor_Stop
func callbackQIRProximitySensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQIRProximitySensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QIRProximitySensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::stop", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::stop")
	}
}

func (ptr *QIRProximitySensor) Stop() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_Stop(ptr.Pointer())
	}
}

func (ptr *QIRProximitySensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQIRProximitySensor_TimerEvent
func callbackQIRProximitySensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::timerEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::timerEvent")
	}
}

func (ptr *QIRProximitySensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximitySensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQIRProximitySensor_ChildEvent
func callbackQIRProximitySensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::childEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::childEvent")
	}
}

func (ptr *QIRProximitySensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximitySensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQIRProximitySensor_ConnectNotify
func callbackQIRProximitySensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQIRProximitySensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QIRProximitySensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::connectNotify", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::connectNotify")
	}
}

func (ptr *QIRProximitySensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QIRProximitySensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQIRProximitySensor_CustomEvent
func callbackQIRProximitySensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::customEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::customEvent")
	}
}

func (ptr *QIRProximitySensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIRProximitySensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQIRProximitySensor_DeleteLater
func callbackQIRProximitySensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQIRProximitySensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QIRProximitySensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::deleteLater", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::deleteLater")
	}
}

func (ptr *QIRProximitySensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QIRProximitySensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQIRProximitySensor_DisconnectNotify
func callbackQIRProximitySensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQIRProximitySensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QIRProximitySensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::disconnectNotify", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::disconnectNotify")
	}
}

func (ptr *QIRProximitySensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QIRProximitySensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQIRProximitySensor_Event
func callbackQIRProximitySensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIRProximitySensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QIRProximitySensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::event", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::event")
	}
}

func (ptr *QIRProximitySensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximitySensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QIRProximitySensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximitySensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQIRProximitySensor_EventFilter
func callbackQIRProximitySensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIRProximitySensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QIRProximitySensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::eventFilter", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::eventFilter")
	}
}

func (ptr *QIRProximitySensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximitySensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QIRProximitySensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximitySensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQIRProximitySensor_MetaObject
func callbackQIRProximitySensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QIRProximitySensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQIRProximitySensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QIRProximitySensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::metaObject", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QIRProximitySensor::metaObject")
	}
}

func (ptr *QIRProximitySensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QIRProximitySensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIRProximitySensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QIRProximitySensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QLightFilter struct {
	QSensorFilter
}

type QLightFilter_ITF interface {
	QSensorFilter_ITF
	QLightFilter_PTR() *QLightFilter
}

func (p *QLightFilter) QLightFilter_PTR() *QLightFilter {
	return p
}

func (p *QLightFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QLightFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQLightFilter(ptr QLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightFilter_PTR().Pointer()
	}
	return nil
}

func NewQLightFilterFromPointer(ptr unsafe.Pointer) *QLightFilter {
	var n = new(QLightFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLightFilter) DestroyQLightFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQLightFilter_Filter
func callbackQLightFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QLightReading) bool)(NewQLightReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QLightFilter) ConnectFilter(f func(reading *QLightReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightFilter::filter", f)
	}
}

func (ptr *QLightFilter) DisconnectFilter(reading QLightReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightFilter::filter")
	}
}

func (ptr *QLightFilter) Filter(reading QLightReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightFilter_Filter(ptr.Pointer(), PointerFromQLightReading(reading)) != 0
	}
	return false
}

type QLightReading struct {
	QSensorReading
}

type QLightReading_ITF interface {
	QSensorReading_ITF
	QLightReading_PTR() *QLightReading
}

func (p *QLightReading) QLightReading_PTR() *QLightReading {
	return p
}

func (p *QLightReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QLightReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQLightReading(ptr QLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReading_PTR().Pointer()
	}
	return nil
}

func NewQLightReadingFromPointer(ptr unsafe.Pointer) *QLightReading {
	var n = new(QLightReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLightReading) DestroyQLightReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QLightReading) Lux() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLightReading_Lux(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightReading) SetLux(lux float64) {
	if ptr.Pointer() != nil {
		C.QLightReading_SetLux(ptr.Pointer(), C.double(lux))
	}
}

//export callbackQLightReading_TimerEvent
func callbackQLightReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLightReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::timerEvent", f)
	}
}

func (ptr *QLightReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::timerEvent")
	}
}

func (ptr *QLightReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLightReading_ChildEvent
func callbackQLightReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLightReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::childEvent", f)
	}
}

func (ptr *QLightReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::childEvent")
	}
}

func (ptr *QLightReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLightReading_ConnectNotify
func callbackQLightReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLightReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLightReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::connectNotify", f)
	}
}

func (ptr *QLightReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::connectNotify")
	}
}

func (ptr *QLightReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLightReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLightReading_CustomEvent
func callbackQLightReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLightReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::customEvent", f)
	}
}

func (ptr *QLightReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::customEvent")
	}
}

func (ptr *QLightReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLightReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLightReading_DeleteLater
func callbackQLightReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLightReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLightReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::deleteLater", f)
	}
}

func (ptr *QLightReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::deleteLater")
	}
}

func (ptr *QLightReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QLightReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLightReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLightReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLightReading_DisconnectNotify
func callbackQLightReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLightReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLightReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::disconnectNotify", f)
	}
}

func (ptr *QLightReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::disconnectNotify")
	}
}

func (ptr *QLightReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLightReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLightReading_Event
func callbackQLightReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLightReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLightReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::event", f)
	}
}

func (ptr *QLightReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::event")
	}
}

func (ptr *QLightReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLightReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLightReading_EventFilter
func callbackQLightReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLightReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLightReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::eventFilter", f)
	}
}

func (ptr *QLightReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::eventFilter")
	}
}

func (ptr *QLightReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLightReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLightReading_MetaObject
func callbackQLightReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLightReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLightReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::metaObject", f)
	}
}

func (ptr *QLightReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightReading::metaObject")
	}
}

func (ptr *QLightReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLightReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLightReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLightReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QLightSensor struct {
	QSensor
}

type QLightSensor_ITF interface {
	QSensor_ITF
	QLightSensor_PTR() *QLightSensor
}

func (p *QLightSensor) QLightSensor_PTR() *QLightSensor {
	return p
}

func (p *QLightSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QLightSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQLightSensor(ptr QLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQLightSensorFromPointer(ptr unsafe.Pointer) *QLightSensor {
	var n = new(QLightSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QLightSensor) FieldOfView() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLightSensor_FieldOfView(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightSensor) Reading() *QLightReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLightReadingFromPointer(C.QLightSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQLightSensor(parent core.QObject_ITF) *QLightSensor {
	var tmpValue = NewQLightSensorFromPointer(C.QLightSensor_NewQLightSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQLightSensor_FieldOfViewChanged
func callbackQLightSensor_FieldOfViewChanged(ptr unsafe.Pointer, fieldOfView C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::fieldOfViewChanged"); signal != nil {
		signal.(func(float64))(float64(fieldOfView))
	}

}

func (ptr *QLightSensor) ConnectFieldOfViewChanged(f func(fieldOfView float64)) {
	if ptr.Pointer() != nil {
		C.QLightSensor_ConnectFieldOfViewChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::fieldOfViewChanged", f)
	}
}

func (ptr *QLightSensor) DisconnectFieldOfViewChanged() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DisconnectFieldOfViewChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::fieldOfViewChanged")
	}
}

func (ptr *QLightSensor) FieldOfViewChanged(fieldOfView float64) {
	if ptr.Pointer() != nil {
		C.QLightSensor_FieldOfViewChanged(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) SetFieldOfView(fieldOfView float64) {
	if ptr.Pointer() != nil {
		C.QLightSensor_SetFieldOfView(ptr.Pointer(), C.double(fieldOfView))
	}
}

//export callbackQLightSensor_DestroyQLightSensor
func callbackQLightSensor_DestroyQLightSensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::~QLightSensor"); signal != nil {
		signal.(func())()
	} else {
		NewQLightSensorFromPointer(ptr).DestroyQLightSensorDefault()
	}
}

func (ptr *QLightSensor) ConnectDestroyQLightSensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::~QLightSensor", f)
	}
}

func (ptr *QLightSensor) DisconnectDestroyQLightSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::~QLightSensor")
	}
}

func (ptr *QLightSensor) DestroyQLightSensor() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DestroyQLightSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLightSensor) DestroyQLightSensorDefault() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DestroyQLightSensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QLightSensor_Type() string {
	return cGoUnpackString(C.QLightSensor_QLightSensor_Type())
}

func (ptr *QLightSensor) Type() string {
	return cGoUnpackString(C.QLightSensor_QLightSensor_Type())
}

//export callbackQLightSensor_Start
func callbackQLightSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLightSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QLightSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::start", f)
	}
}

func (ptr *QLightSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::start")
	}
}

func (ptr *QLightSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QLightSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLightSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QLightSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLightSensor_Stop
func callbackQLightSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQLightSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QLightSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::stop", f)
	}
}

func (ptr *QLightSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::stop")
	}
}

func (ptr *QLightSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QLightSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QLightSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QLightSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQLightSensor_TimerEvent
func callbackQLightSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLightSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::timerEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::timerEvent")
	}
}

func (ptr *QLightSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLightSensor_ChildEvent
func callbackQLightSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLightSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::childEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::childEvent")
	}
}

func (ptr *QLightSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLightSensor_ConnectNotify
func callbackQLightSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLightSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLightSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::connectNotify", f)
	}
}

func (ptr *QLightSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::connectNotify")
	}
}

func (ptr *QLightSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLightSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLightSensor_CustomEvent
func callbackQLightSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLightSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::customEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::customEvent")
	}
}

func (ptr *QLightSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLightSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLightSensor_DeleteLater
func callbackQLightSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLightSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLightSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::deleteLater", f)
	}
}

func (ptr *QLightSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::deleteLater")
	}
}

func (ptr *QLightSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLightSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLightSensor_DisconnectNotify
func callbackQLightSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLightSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLightSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::disconnectNotify", f)
	}
}

func (ptr *QLightSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::disconnectNotify")
	}
}

func (ptr *QLightSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLightSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLightSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLightSensor_Event
func callbackQLightSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLightSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLightSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::event", f)
	}
}

func (ptr *QLightSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::event")
	}
}

func (ptr *QLightSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLightSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLightSensor_EventFilter
func callbackQLightSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLightSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLightSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::eventFilter", f)
	}
}

func (ptr *QLightSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::eventFilter")
	}
}

func (ptr *QLightSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLightSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLightSensor_MetaObject
func callbackQLightSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLightSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLightSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLightSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::metaObject", f)
	}
}

func (ptr *QLightSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLightSensor::metaObject")
	}
}

func (ptr *QLightSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLightSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLightSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLightSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QMagnetometer struct {
	QSensor
}

type QMagnetometer_ITF interface {
	QSensor_ITF
	QMagnetometer_PTR() *QMagnetometer
}

func (p *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return p
}

func (p *QMagnetometer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QMagnetometer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQMagnetometer(ptr QMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometer_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFromPointer(ptr unsafe.Pointer) *QMagnetometer {
	var n = new(QMagnetometer)
	n.SetPointer(ptr)
	return n
}
func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQMagnetometerReadingFromPointer(C.QMagnetometer_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_ReturnGeoValues(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnGeoValues))))
	}
}

func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {
	var tmpValue = NewQMagnetometerFromPointer(C.QMagnetometer_NewQMagnetometer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQMagnetometer_ReturnGeoValuesChanged
func callbackQMagnetometer_ReturnGeoValuesChanged(ptr unsafe.Pointer, returnGeoValues C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::returnGeoValuesChanged"); signal != nil {
		signal.(func(bool))(int8(returnGeoValues) != 0)
	}

}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectReturnGeoValuesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::returnGeoValuesChanged", f)
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::returnGeoValuesChanged")
	}
}

func (ptr *QMagnetometer) ReturnGeoValuesChanged(returnGeoValues bool) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ReturnGeoValuesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnGeoValues))))
	}
}

//export callbackQMagnetometer_DestroyQMagnetometer
func callbackQMagnetometer_DestroyQMagnetometer(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::~QMagnetometer"); signal != nil {
		signal.(func())()
	} else {
		NewQMagnetometerFromPointer(ptr).DestroyQMagnetometerDefault()
	}
}

func (ptr *QMagnetometer) ConnectDestroyQMagnetometer(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::~QMagnetometer", f)
	}
}

func (ptr *QMagnetometer) DisconnectDestroyQMagnetometer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::~QMagnetometer")
	}
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMagnetometer) DestroyQMagnetometerDefault() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QMagnetometer_Type() string {
	return cGoUnpackString(C.QMagnetometer_QMagnetometer_Type())
}

func (ptr *QMagnetometer) Type() string {
	return cGoUnpackString(C.QMagnetometer_QMagnetometer_Type())
}

//export callbackQMagnetometer_Start
func callbackQMagnetometer_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMagnetometerFromPointer(ptr).StartDefault())))
}

func (ptr *QMagnetometer) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::start", f)
	}
}

func (ptr *QMagnetometer) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::start")
	}
}

func (ptr *QMagnetometer) Start() bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMagnetometer) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQMagnetometer_Stop
func callbackQMagnetometer_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQMagnetometerFromPointer(ptr).StopDefault()
	}
}

func (ptr *QMagnetometer) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::stop", f)
	}
}

func (ptr *QMagnetometer) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::stop")
	}
}

func (ptr *QMagnetometer) Stop() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_Stop(ptr.Pointer())
	}
}

func (ptr *QMagnetometer) StopDefault() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_StopDefault(ptr.Pointer())
	}
}

//export callbackQMagnetometer_TimerEvent
func callbackQMagnetometer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::timerEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::timerEvent")
	}
}

func (ptr *QMagnetometer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMagnetometer_ChildEvent
func callbackQMagnetometer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::childEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::childEvent")
	}
}

func (ptr *QMagnetometer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMagnetometer_ConnectNotify
func callbackQMagnetometer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMagnetometerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMagnetometer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::connectNotify", f)
	}
}

func (ptr *QMagnetometer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::connectNotify")
	}
}

func (ptr *QMagnetometer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMagnetometer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMagnetometer_CustomEvent
func callbackQMagnetometer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::customEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::customEvent")
	}
}

func (ptr *QMagnetometer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMagnetometer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMagnetometer_DeleteLater
func callbackQMagnetometer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMagnetometerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMagnetometer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::deleteLater", f)
	}
}

func (ptr *QMagnetometer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::deleteLater")
	}
}

func (ptr *QMagnetometer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMagnetometer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMagnetometer_DisconnectNotify
func callbackQMagnetometer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMagnetometerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMagnetometer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::disconnectNotify", f)
	}
}

func (ptr *QMagnetometer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::disconnectNotify")
	}
}

func (ptr *QMagnetometer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMagnetometer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMagnetometer_Event
func callbackQMagnetometer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMagnetometerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMagnetometer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::event", f)
	}
}

func (ptr *QMagnetometer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::event")
	}
}

func (ptr *QMagnetometer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMagnetometer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMagnetometer_EventFilter
func callbackQMagnetometer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMagnetometerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMagnetometer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::eventFilter", f)
	}
}

func (ptr *QMagnetometer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::eventFilter")
	}
}

func (ptr *QMagnetometer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMagnetometer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMagnetometer_MetaObject
func callbackQMagnetometer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMagnetometerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMagnetometer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::metaObject", f)
	}
}

func (ptr *QMagnetometer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometer::metaObject")
	}
}

func (ptr *QMagnetometer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMagnetometer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMagnetometer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMagnetometer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QMagnetometerFilter struct {
	QSensorFilter
}

type QMagnetometerFilter_ITF interface {
	QSensorFilter_ITF
	QMagnetometerFilter_PTR() *QMagnetometerFilter
}

func (p *QMagnetometerFilter) QMagnetometerFilter_PTR() *QMagnetometerFilter {
	return p
}

func (p *QMagnetometerFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QMagnetometerFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQMagnetometerFilter(ptr QMagnetometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerFilter_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFilterFromPointer(ptr unsafe.Pointer) *QMagnetometerFilter {
	var n = new(QMagnetometerFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMagnetometerFilter) DestroyQMagnetometerFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQMagnetometerFilter_Filter
func callbackQMagnetometerFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QMagnetometerReading) bool)(NewQMagnetometerReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMagnetometerFilter) ConnectFilter(f func(reading *QMagnetometerReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerFilter::filter", f)
	}
}

func (ptr *QMagnetometerFilter) DisconnectFilter(reading QMagnetometerReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerFilter::filter")
	}
}

func (ptr *QMagnetometerFilter) Filter(reading QMagnetometerReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometerFilter_Filter(ptr.Pointer(), PointerFromQMagnetometerReading(reading)) != 0
	}
	return false
}

type QMagnetometerReading struct {
	QSensorReading
}

type QMagnetometerReading_ITF interface {
	QSensorReading_ITF
	QMagnetometerReading_PTR() *QMagnetometerReading
}

func (p *QMagnetometerReading) QMagnetometerReading_PTR() *QMagnetometerReading {
	return p
}

func (p *QMagnetometerReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QMagnetometerReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQMagnetometerReading(ptr QMagnetometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerReading_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerReadingFromPointer(ptr unsafe.Pointer) *QMagnetometerReading {
	var n = new(QMagnetometerReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMagnetometerReading) DestroyQMagnetometerReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QMagnetometerReading) CalibrationLevel() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) SetCalibrationLevel(calibrationLevel float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QMagnetometerReading) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QMagnetometerReading) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QMagnetometerReading) SetZ(z float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

//export callbackQMagnetometerReading_TimerEvent
func callbackQMagnetometerReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::timerEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::timerEvent")
	}
}

func (ptr *QMagnetometerReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometerReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMagnetometerReading_ChildEvent
func callbackQMagnetometerReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::childEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::childEvent")
	}
}

func (ptr *QMagnetometerReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometerReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMagnetometerReading_ConnectNotify
func callbackQMagnetometerReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMagnetometerReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::connectNotify", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::connectNotify")
	}
}

func (ptr *QMagnetometerReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMagnetometerReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMagnetometerReading_CustomEvent
func callbackQMagnetometerReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::customEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::customEvent")
	}
}

func (ptr *QMagnetometerReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMagnetometerReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMagnetometerReading_DeleteLater
func callbackQMagnetometerReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMagnetometerReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMagnetometerReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::deleteLater", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::deleteLater")
	}
}

func (ptr *QMagnetometerReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMagnetometerReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMagnetometerReading_DisconnectNotify
func callbackQMagnetometerReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMagnetometerReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::disconnectNotify", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::disconnectNotify")
	}
}

func (ptr *QMagnetometerReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMagnetometerReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMagnetometerReading_Event
func callbackQMagnetometerReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMagnetometerReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMagnetometerReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::event", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::event")
	}
}

func (ptr *QMagnetometerReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometerReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMagnetometerReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometerReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMagnetometerReading_EventFilter
func callbackQMagnetometerReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMagnetometerReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMagnetometerReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::eventFilter", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::eventFilter")
	}
}

func (ptr *QMagnetometerReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometerReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMagnetometerReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometerReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMagnetometerReading_MetaObject
func callbackQMagnetometerReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMagnetometerReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMagnetometerReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMagnetometerReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::metaObject", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMagnetometerReading::metaObject")
	}
}

func (ptr *QMagnetometerReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMagnetometerReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMagnetometerReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMagnetometerReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QOrientationFilter struct {
	QSensorFilter
}

type QOrientationFilter_ITF interface {
	QSensorFilter_ITF
	QOrientationFilter_PTR() *QOrientationFilter
}

func (p *QOrientationFilter) QOrientationFilter_PTR() *QOrientationFilter {
	return p
}

func (p *QOrientationFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QOrientationFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQOrientationFilter(ptr QOrientationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationFilter_PTR().Pointer()
	}
	return nil
}

func NewQOrientationFilterFromPointer(ptr unsafe.Pointer) *QOrientationFilter {
	var n = new(QOrientationFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOrientationFilter) DestroyQOrientationFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQOrientationFilter_Filter
func callbackQOrientationFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QOrientationReading) bool)(NewQOrientationReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QOrientationFilter) ConnectFilter(f func(reading *QOrientationReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationFilter::filter", f)
	}
}

func (ptr *QOrientationFilter) DisconnectFilter(reading QOrientationReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationFilter::filter")
	}
}

func (ptr *QOrientationFilter) Filter(reading QOrientationReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationFilter_Filter(ptr.Pointer(), PointerFromQOrientationReading(reading)) != 0
	}
	return false
}

//QOrientationReading::Orientation
type QOrientationReading__Orientation int64

const (
	QOrientationReading__Undefined = QOrientationReading__Orientation(0)
	QOrientationReading__TopUp     = QOrientationReading__Orientation(1)
	QOrientationReading__TopDown   = QOrientationReading__Orientation(2)
	QOrientationReading__LeftUp    = QOrientationReading__Orientation(3)
	QOrientationReading__RightUp   = QOrientationReading__Orientation(4)
	QOrientationReading__FaceUp    = QOrientationReading__Orientation(5)
	QOrientationReading__FaceDown  = QOrientationReading__Orientation(6)
)

type QOrientationReading struct {
	QSensorReading
}

type QOrientationReading_ITF interface {
	QSensorReading_ITF
	QOrientationReading_PTR() *QOrientationReading
}

func (p *QOrientationReading) QOrientationReading_PTR() *QOrientationReading {
	return p
}

func (p *QOrientationReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QOrientationReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQOrientationReading(ptr QOrientationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationReading_PTR().Pointer()
	}
	return nil
}

func NewQOrientationReadingFromPointer(ptr unsafe.Pointer) *QOrientationReading {
	var n = new(QOrientationReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOrientationReading) DestroyQOrientationReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QOrientationReading) Orientation() QOrientationReading__Orientation {
	if ptr.Pointer() != nil {
		return QOrientationReading__Orientation(C.QOrientationReading_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

//export callbackQOrientationReading_TimerEvent
func callbackQOrientationReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::timerEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::timerEvent")
	}
}

func (ptr *QOrientationReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQOrientationReading_ChildEvent
func callbackQOrientationReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::childEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::childEvent")
	}
}

func (ptr *QOrientationReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQOrientationReading_ConnectNotify
func callbackQOrientationReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQOrientationReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QOrientationReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::connectNotify", f)
	}
}

func (ptr *QOrientationReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::connectNotify")
	}
}

func (ptr *QOrientationReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QOrientationReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQOrientationReading_CustomEvent
func callbackQOrientationReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::customEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::customEvent")
	}
}

func (ptr *QOrientationReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QOrientationReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQOrientationReading_DeleteLater
func callbackQOrientationReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQOrientationReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QOrientationReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::deleteLater", f)
	}
}

func (ptr *QOrientationReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::deleteLater")
	}
}

func (ptr *QOrientationReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QOrientationReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QOrientationReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QOrientationReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQOrientationReading_DisconnectNotify
func callbackQOrientationReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQOrientationReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QOrientationReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::disconnectNotify", f)
	}
}

func (ptr *QOrientationReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::disconnectNotify")
	}
}

func (ptr *QOrientationReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QOrientationReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQOrientationReading_Event
func callbackQOrientationReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQOrientationReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QOrientationReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::event", f)
	}
}

func (ptr *QOrientationReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::event")
	}
}

func (ptr *QOrientationReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QOrientationReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQOrientationReading_EventFilter
func callbackQOrientationReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQOrientationReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QOrientationReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::eventFilter", f)
	}
}

func (ptr *QOrientationReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::eventFilter")
	}
}

func (ptr *QOrientationReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QOrientationReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQOrientationReading_MetaObject
func callbackQOrientationReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQOrientationReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QOrientationReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::metaObject", f)
	}
}

func (ptr *QOrientationReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationReading::metaObject")
	}
}

func (ptr *QOrientationReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QOrientationReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QOrientationReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QOrientationReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensor_ITF interface {
	QSensor_ITF
	QOrientationSensor_PTR() *QOrientationSensor
}

func (p *QOrientationSensor) QOrientationSensor_PTR() *QOrientationSensor {
	return p
}

func (p *QOrientationSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QOrientationSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQOrientationSensor(ptr QOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationSensor_PTR().Pointer()
	}
	return nil
}

func NewQOrientationSensorFromPointer(ptr unsafe.Pointer) *QOrientationSensor {
	var n = new(QOrientationSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QOrientationSensor) Reading() *QOrientationReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQOrientationReadingFromPointer(C.QOrientationSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQOrientationSensor(parent core.QObject_ITF) *QOrientationSensor {
	var tmpValue = NewQOrientationSensorFromPointer(C.QOrientationSensor_NewQOrientationSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQOrientationSensor_DestroyQOrientationSensor
func callbackQOrientationSensor_DestroyQOrientationSensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::~QOrientationSensor"); signal != nil {
		signal.(func())()
	} else {
		NewQOrientationSensorFromPointer(ptr).DestroyQOrientationSensorDefault()
	}
}

func (ptr *QOrientationSensor) ConnectDestroyQOrientationSensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::~QOrientationSensor", f)
	}
}

func (ptr *QOrientationSensor) DisconnectDestroyQOrientationSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::~QOrientationSensor")
	}
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DestroyQOrientationSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QOrientationSensor) DestroyQOrientationSensorDefault() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DestroyQOrientationSensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QOrientationSensor_Type() string {
	return cGoUnpackString(C.QOrientationSensor_QOrientationSensor_Type())
}

func (ptr *QOrientationSensor) Type() string {
	return cGoUnpackString(C.QOrientationSensor_QOrientationSensor_Type())
}

//export callbackQOrientationSensor_Start
func callbackQOrientationSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQOrientationSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QOrientationSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::start", f)
	}
}

func (ptr *QOrientationSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::start")
	}
}

func (ptr *QOrientationSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QOrientationSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QOrientationSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QOrientationSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQOrientationSensor_Stop
func callbackQOrientationSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQOrientationSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QOrientationSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::stop", f)
	}
}

func (ptr *QOrientationSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::stop")
	}
}

func (ptr *QOrientationSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QOrientationSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQOrientationSensor_TimerEvent
func callbackQOrientationSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::timerEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::timerEvent")
	}
}

func (ptr *QOrientationSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQOrientationSensor_ChildEvent
func callbackQOrientationSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::childEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::childEvent")
	}
}

func (ptr *QOrientationSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQOrientationSensor_ConnectNotify
func callbackQOrientationSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQOrientationSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QOrientationSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::connectNotify", f)
	}
}

func (ptr *QOrientationSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::connectNotify")
	}
}

func (ptr *QOrientationSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QOrientationSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQOrientationSensor_CustomEvent
func callbackQOrientationSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::customEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::customEvent")
	}
}

func (ptr *QOrientationSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QOrientationSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQOrientationSensor_DeleteLater
func callbackQOrientationSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQOrientationSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QOrientationSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::deleteLater", f)
	}
}

func (ptr *QOrientationSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::deleteLater")
	}
}

func (ptr *QOrientationSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QOrientationSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQOrientationSensor_DisconnectNotify
func callbackQOrientationSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQOrientationSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QOrientationSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::disconnectNotify", f)
	}
}

func (ptr *QOrientationSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::disconnectNotify")
	}
}

func (ptr *QOrientationSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QOrientationSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQOrientationSensor_Event
func callbackQOrientationSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQOrientationSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QOrientationSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::event", f)
	}
}

func (ptr *QOrientationSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::event")
	}
}

func (ptr *QOrientationSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QOrientationSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQOrientationSensor_EventFilter
func callbackQOrientationSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQOrientationSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QOrientationSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::eventFilter", f)
	}
}

func (ptr *QOrientationSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::eventFilter")
	}
}

func (ptr *QOrientationSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QOrientationSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQOrientationSensor_MetaObject
func callbackQOrientationSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QOrientationSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQOrientationSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QOrientationSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::metaObject", f)
	}
}

func (ptr *QOrientationSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QOrientationSensor::metaObject")
	}
}

func (ptr *QOrientationSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QOrientationSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QOrientationSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QOrientationSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QPressureFilter struct {
	QSensorFilter
}

type QPressureFilter_ITF interface {
	QSensorFilter_ITF
	QPressureFilter_PTR() *QPressureFilter
}

func (p *QPressureFilter) QPressureFilter_PTR() *QPressureFilter {
	return p
}

func (p *QPressureFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QPressureFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQPressureFilter(ptr QPressureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureFilter_PTR().Pointer()
	}
	return nil
}

func NewQPressureFilterFromPointer(ptr unsafe.Pointer) *QPressureFilter {
	var n = new(QPressureFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPressureFilter) DestroyQPressureFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQPressureFilter_Filter
func callbackQPressureFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QPressureReading) bool)(NewQPressureReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPressureFilter) ConnectFilter(f func(reading *QPressureReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureFilter::filter", f)
	}
}

func (ptr *QPressureFilter) DisconnectFilter(reading QPressureReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureFilter::filter")
	}
}

func (ptr *QPressureFilter) Filter(reading QPressureReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureFilter_Filter(ptr.Pointer(), PointerFromQPressureReading(reading)) != 0
	}
	return false
}

type QPressureReading struct {
	QSensorReading
}

type QPressureReading_ITF interface {
	QSensorReading_ITF
	QPressureReading_PTR() *QPressureReading
}

func (p *QPressureReading) QPressureReading_PTR() *QPressureReading {
	return p
}

func (p *QPressureReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QPressureReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQPressureReading(ptr QPressureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureReading_PTR().Pointer()
	}
	return nil
}

func NewQPressureReadingFromPointer(ptr unsafe.Pointer) *QPressureReading {
	var n = new(QPressureReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPressureReading) DestroyQPressureReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QPressureReading) Pressure() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Pressure(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) Temperature() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) SetPressure(pressure float64) {
	if ptr.Pointer() != nil {
		C.QPressureReading_SetPressure(ptr.Pointer(), C.double(pressure))
	}
}

func (ptr *QPressureReading) SetTemperature(temperature float64) {
	if ptr.Pointer() != nil {
		C.QPressureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

//export callbackQPressureReading_TimerEvent
func callbackQPressureReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPressureReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::timerEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::timerEvent")
	}
}

func (ptr *QPressureReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQPressureReading_ChildEvent
func callbackQPressureReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPressureReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::childEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::childEvent")
	}
}

func (ptr *QPressureReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPressureReading_ConnectNotify
func callbackQPressureReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPressureReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPressureReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::connectNotify", f)
	}
}

func (ptr *QPressureReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::connectNotify")
	}
}

func (ptr *QPressureReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPressureReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPressureReading_CustomEvent
func callbackQPressureReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPressureReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::customEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::customEvent")
	}
}

func (ptr *QPressureReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPressureReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPressureReading_DeleteLater
func callbackQPressureReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPressureReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPressureReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::deleteLater", f)
	}
}

func (ptr *QPressureReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::deleteLater")
	}
}

func (ptr *QPressureReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QPressureReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPressureReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QPressureReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPressureReading_DisconnectNotify
func callbackQPressureReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPressureReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPressureReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::disconnectNotify", f)
	}
}

func (ptr *QPressureReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::disconnectNotify")
	}
}

func (ptr *QPressureReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPressureReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPressureReading_Event
func callbackQPressureReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPressureReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QPressureReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::event", f)
	}
}

func (ptr *QPressureReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::event")
	}
}

func (ptr *QPressureReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QPressureReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQPressureReading_EventFilter
func callbackQPressureReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPressureReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QPressureReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::eventFilter", f)
	}
}

func (ptr *QPressureReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::eventFilter")
	}
}

func (ptr *QPressureReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QPressureReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQPressureReading_MetaObject
func callbackQPressureReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPressureReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPressureReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::metaObject", f)
	}
}

func (ptr *QPressureReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureReading::metaObject")
	}
}

func (ptr *QPressureReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPressureReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPressureReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPressureReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QPressureSensor struct {
	QSensor
}

type QPressureSensor_ITF interface {
	QSensor_ITF
	QPressureSensor_PTR() *QPressureSensor
}

func (p *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return p
}

func (p *QPressureSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QPressureSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQPressureSensor(ptr QPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureSensor_PTR().Pointer()
	}
	return nil
}

func NewQPressureSensorFromPointer(ptr unsafe.Pointer) *QPressureSensor {
	var n = new(QPressureSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QPressureSensor) Reading() *QPressureReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQPressureReadingFromPointer(C.QPressureSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {
	var tmpValue = NewQPressureSensorFromPointer(C.QPressureSensor_NewQPressureSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DestroyQPressureSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QPressureSensor_Type() string {
	return cGoUnpackString(C.QPressureSensor_QPressureSensor_Type())
}

func (ptr *QPressureSensor) Type() string {
	return cGoUnpackString(C.QPressureSensor_QPressureSensor_Type())
}

//export callbackQPressureSensor_Start
func callbackQPressureSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPressureSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QPressureSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::start", f)
	}
}

func (ptr *QPressureSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::start")
	}
}

func (ptr *QPressureSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QPressureSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPressureSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QPressureSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQPressureSensor_Stop
func callbackQPressureSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQPressureSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QPressureSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::stop", f)
	}
}

func (ptr *QPressureSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::stop")
	}
}

func (ptr *QPressureSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QPressureSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQPressureSensor_TimerEvent
func callbackQPressureSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::timerEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::timerEvent")
	}
}

func (ptr *QPressureSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQPressureSensor_ChildEvent
func callbackQPressureSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::childEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::childEvent")
	}
}

func (ptr *QPressureSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQPressureSensor_ConnectNotify
func callbackQPressureSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPressureSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPressureSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::connectNotify", f)
	}
}

func (ptr *QPressureSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::connectNotify")
	}
}

func (ptr *QPressureSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPressureSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPressureSensor_CustomEvent
func callbackQPressureSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::customEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::customEvent")
	}
}

func (ptr *QPressureSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPressureSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQPressureSensor_DeleteLater
func callbackQPressureSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQPressureSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QPressureSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::deleteLater", f)
	}
}

func (ptr *QPressureSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::deleteLater")
	}
}

func (ptr *QPressureSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QPressureSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQPressureSensor_DisconnectNotify
func callbackQPressureSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQPressureSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QPressureSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::disconnectNotify", f)
	}
}

func (ptr *QPressureSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::disconnectNotify")
	}
}

func (ptr *QPressureSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QPressureSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQPressureSensor_Event
func callbackQPressureSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPressureSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QPressureSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::event", f)
	}
}

func (ptr *QPressureSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::event")
	}
}

func (ptr *QPressureSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QPressureSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQPressureSensor_EventFilter
func callbackQPressureSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPressureSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QPressureSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::eventFilter", f)
	}
}

func (ptr *QPressureSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::eventFilter")
	}
}

func (ptr *QPressureSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QPressureSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQPressureSensor_MetaObject
func callbackQPressureSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QPressureSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQPressureSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QPressureSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::metaObject", f)
	}
}

func (ptr *QPressureSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QPressureSensor::metaObject")
	}
}

func (ptr *QPressureSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPressureSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPressureSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QPressureSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QProximityFilter struct {
	QSensorFilter
}

type QProximityFilter_ITF interface {
	QSensorFilter_ITF
	QProximityFilter_PTR() *QProximityFilter
}

func (p *QProximityFilter) QProximityFilter_PTR() *QProximityFilter {
	return p
}

func (p *QProximityFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QProximityFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQProximityFilter(ptr QProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityFilter_PTR().Pointer()
	}
	return nil
}

func NewQProximityFilterFromPointer(ptr unsafe.Pointer) *QProximityFilter {
	var n = new(QProximityFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QProximityFilter) DestroyQProximityFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQProximityFilter_Filter
func callbackQProximityFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QProximityReading) bool)(NewQProximityReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QProximityFilter) ConnectFilter(f func(reading *QProximityReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityFilter::filter", f)
	}
}

func (ptr *QProximityFilter) DisconnectFilter(reading QProximityReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityFilter::filter")
	}
}

func (ptr *QProximityFilter) Filter(reading QProximityReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityFilter_Filter(ptr.Pointer(), PointerFromQProximityReading(reading)) != 0
	}
	return false
}

type QProximityReading struct {
	QSensorReading
}

type QProximityReading_ITF interface {
	QSensorReading_ITF
	QProximityReading_PTR() *QProximityReading
}

func (p *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return p
}

func (p *QProximityReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QProximityReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) *QProximityReading {
	var n = new(QProximityReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QProximityReading) DestroyQProximityReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QProximityReading) Close() bool {
	if ptr.Pointer() != nil {
		return C.QProximityReading_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(close))))
	}
}

//export callbackQProximityReading_TimerEvent
func callbackQProximityReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProximityReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::timerEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::timerEvent")
	}
}

func (ptr *QProximityReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximityReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQProximityReading_ChildEvent
func callbackQProximityReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProximityReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::childEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::childEvent")
	}
}

func (ptr *QProximityReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximityReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQProximityReading_ConnectNotify
func callbackQProximityReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQProximityReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QProximityReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::connectNotify", f)
	}
}

func (ptr *QProximityReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::connectNotify")
	}
}

func (ptr *QProximityReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QProximityReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQProximityReading_CustomEvent
func callbackQProximityReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProximityReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::customEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::customEvent")
	}
}

func (ptr *QProximityReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProximityReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQProximityReading_DeleteLater
func callbackQProximityReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQProximityReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QProximityReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::deleteLater", f)
	}
}

func (ptr *QProximityReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::deleteLater")
	}
}

func (ptr *QProximityReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QProximityReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QProximityReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QProximityReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQProximityReading_DisconnectNotify
func callbackQProximityReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQProximityReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QProximityReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::disconnectNotify", f)
	}
}

func (ptr *QProximityReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::disconnectNotify")
	}
}

func (ptr *QProximityReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QProximityReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximityReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQProximityReading_Event
func callbackQProximityReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQProximityReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QProximityReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::event", f)
	}
}

func (ptr *QProximityReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::event")
	}
}

func (ptr *QProximityReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QProximityReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQProximityReading_EventFilter
func callbackQProximityReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQProximityReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QProximityReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::eventFilter", f)
	}
}

func (ptr *QProximityReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::eventFilter")
	}
}

func (ptr *QProximityReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QProximityReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQProximityReading_MetaObject
func callbackQProximityReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximityReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQProximityReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QProximityReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::metaObject", f)
	}
}

func (ptr *QProximityReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximityReading::metaObject")
	}
}

func (ptr *QProximityReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QProximityReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProximityReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QProximityReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QProximitySensor struct {
	QSensor
}

type QProximitySensor_ITF interface {
	QSensor_ITF
	QProximitySensor_PTR() *QProximitySensor
}

func (p *QProximitySensor) QProximitySensor_PTR() *QProximitySensor {
	return p
}

func (p *QProximitySensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QProximitySensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQProximitySensor(ptr QProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQProximitySensorFromPointer(ptr unsafe.Pointer) *QProximitySensor {
	var n = new(QProximitySensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QProximitySensor) Reading() *QProximityReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQProximityReadingFromPointer(C.QProximitySensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQProximitySensor(parent core.QObject_ITF) *QProximitySensor {
	var tmpValue = NewQProximitySensorFromPointer(C.QProximitySensor_NewQProximitySensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQProximitySensor_DestroyQProximitySensor
func callbackQProximitySensor_DestroyQProximitySensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::~QProximitySensor"); signal != nil {
		signal.(func())()
	} else {
		NewQProximitySensorFromPointer(ptr).DestroyQProximitySensorDefault()
	}
}

func (ptr *QProximitySensor) ConnectDestroyQProximitySensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::~QProximitySensor", f)
	}
}

func (ptr *QProximitySensor) DisconnectDestroyQProximitySensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::~QProximitySensor")
	}
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DestroyQProximitySensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QProximitySensor) DestroyQProximitySensorDefault() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DestroyQProximitySensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QProximitySensor_Type() string {
	return cGoUnpackString(C.QProximitySensor_QProximitySensor_Type())
}

func (ptr *QProximitySensor) Type() string {
	return cGoUnpackString(C.QProximitySensor_QProximitySensor_Type())
}

//export callbackQProximitySensor_Start
func callbackQProximitySensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQProximitySensorFromPointer(ptr).StartDefault())))
}

func (ptr *QProximitySensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::start", f)
	}
}

func (ptr *QProximitySensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::start")
	}
}

func (ptr *QProximitySensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QProximitySensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProximitySensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QProximitySensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQProximitySensor_Stop
func callbackQProximitySensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQProximitySensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QProximitySensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::stop", f)
	}
}

func (ptr *QProximitySensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::stop")
	}
}

func (ptr *QProximitySensor) Stop() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_Stop(ptr.Pointer())
	}
}

func (ptr *QProximitySensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQProximitySensor_TimerEvent
func callbackQProximitySensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProximitySensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProximitySensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::timerEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::timerEvent")
	}
}

func (ptr *QProximitySensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximitySensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQProximitySensor_ChildEvent
func callbackQProximitySensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProximitySensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProximitySensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::childEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::childEvent")
	}
}

func (ptr *QProximitySensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximitySensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQProximitySensor_ConnectNotify
func callbackQProximitySensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQProximitySensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QProximitySensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::connectNotify", f)
	}
}

func (ptr *QProximitySensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::connectNotify")
	}
}

func (ptr *QProximitySensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QProximitySensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQProximitySensor_CustomEvent
func callbackQProximitySensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProximitySensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProximitySensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::customEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::customEvent")
	}
}

func (ptr *QProximitySensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProximitySensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQProximitySensor_DeleteLater
func callbackQProximitySensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQProximitySensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QProximitySensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::deleteLater", f)
	}
}

func (ptr *QProximitySensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::deleteLater")
	}
}

func (ptr *QProximitySensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QProximitySensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQProximitySensor_DisconnectNotify
func callbackQProximitySensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQProximitySensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QProximitySensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::disconnectNotify", f)
	}
}

func (ptr *QProximitySensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::disconnectNotify")
	}
}

func (ptr *QProximitySensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QProximitySensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQProximitySensor_Event
func callbackQProximitySensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQProximitySensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QProximitySensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::event", f)
	}
}

func (ptr *QProximitySensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::event")
	}
}

func (ptr *QProximitySensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximitySensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QProximitySensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximitySensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQProximitySensor_EventFilter
func callbackQProximitySensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQProximitySensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QProximitySensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::eventFilter", f)
	}
}

func (ptr *QProximitySensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::eventFilter")
	}
}

func (ptr *QProximitySensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximitySensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QProximitySensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximitySensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQProximitySensor_MetaObject
func callbackQProximitySensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QProximitySensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQProximitySensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QProximitySensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::metaObject", f)
	}
}

func (ptr *QProximitySensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QProximitySensor::metaObject")
	}
}

func (ptr *QProximitySensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QProximitySensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProximitySensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QProximitySensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QRotationFilter struct {
	QSensorFilter
}

type QRotationFilter_ITF interface {
	QSensorFilter_ITF
	QRotationFilter_PTR() *QRotationFilter
}

func (p *QRotationFilter) QRotationFilter_PTR() *QRotationFilter {
	return p
}

func (p *QRotationFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QRotationFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQRotationFilter(ptr QRotationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationFilter_PTR().Pointer()
	}
	return nil
}

func NewQRotationFilterFromPointer(ptr unsafe.Pointer) *QRotationFilter {
	var n = new(QRotationFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRotationFilter) DestroyQRotationFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQRotationFilter_Filter
func callbackQRotationFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QRotationReading) bool)(NewQRotationReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QRotationFilter) ConnectFilter(f func(reading *QRotationReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationFilter::filter", f)
	}
}

func (ptr *QRotationFilter) DisconnectFilter(reading QRotationReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationFilter::filter")
	}
}

func (ptr *QRotationFilter) Filter(reading QRotationReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationFilter_Filter(ptr.Pointer(), PointerFromQRotationReading(reading)) != 0
	}
	return false
}

type QRotationReading struct {
	QSensorReading
}

type QRotationReading_ITF interface {
	QSensorReading_ITF
	QRotationReading_PTR() *QRotationReading
}

func (p *QRotationReading) QRotationReading_PTR() *QRotationReading {
	return p
}

func (p *QRotationReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QRotationReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQRotationReading(ptr QRotationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationReading_PTR().Pointer()
	}
	return nil
}

func NewQRotationReadingFromPointer(ptr unsafe.Pointer) *QRotationReading {
	var n = new(QRotationReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRotationReading) DestroyQRotationReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QRotationReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) SetFromEuler(x float64, y float64, z float64) {
	if ptr.Pointer() != nil {
		C.QRotationReading_SetFromEuler(ptr.Pointer(), C.double(x), C.double(y), C.double(z))
	}
}

//export callbackQRotationReading_TimerEvent
func callbackQRotationReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRotationReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::timerEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::timerEvent")
	}
}

func (ptr *QRotationReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQRotationReading_ChildEvent
func callbackQRotationReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRotationReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::childEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::childEvent")
	}
}

func (ptr *QRotationReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRotationReading_ConnectNotify
func callbackQRotationReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRotationReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRotationReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::connectNotify", f)
	}
}

func (ptr *QRotationReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::connectNotify")
	}
}

func (ptr *QRotationReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QRotationReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRotationReading_CustomEvent
func callbackQRotationReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRotationReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::customEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::customEvent")
	}
}

func (ptr *QRotationReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRotationReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRotationReading_DeleteLater
func callbackQRotationReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRotationReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRotationReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::deleteLater", f)
	}
}

func (ptr *QRotationReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::deleteLater")
	}
}

func (ptr *QRotationReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QRotationReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QRotationReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRotationReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQRotationReading_DisconnectNotify
func callbackQRotationReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRotationReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRotationReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::disconnectNotify", f)
	}
}

func (ptr *QRotationReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::disconnectNotify")
	}
}

func (ptr *QRotationReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QRotationReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRotationReading_Event
func callbackQRotationReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRotationReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRotationReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::event", f)
	}
}

func (ptr *QRotationReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::event")
	}
}

func (ptr *QRotationReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QRotationReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQRotationReading_EventFilter
func callbackQRotationReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRotationReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRotationReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::eventFilter", f)
	}
}

func (ptr *QRotationReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::eventFilter")
	}
}

func (ptr *QRotationReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QRotationReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQRotationReading_MetaObject
func callbackQRotationReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQRotationReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRotationReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::metaObject", f)
	}
}

func (ptr *QRotationReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationReading::metaObject")
	}
}

func (ptr *QRotationReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRotationReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRotationReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRotationReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QRotationSensor struct {
	QSensor
}

type QRotationSensor_ITF interface {
	QSensor_ITF
	QRotationSensor_PTR() *QRotationSensor
}

func (p *QRotationSensor) QRotationSensor_PTR() *QRotationSensor {
	return p
}

func (p *QRotationSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QRotationSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQRotationSensor(ptr QRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationSensor_PTR().Pointer()
	}
	return nil
}

func NewQRotationSensorFromPointer(ptr unsafe.Pointer) *QRotationSensor {
	var n = new(QRotationSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QRotationSensor) HasZ() bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_HasZ(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRotationSensor) Reading() *QRotationReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQRotationReadingFromPointer(C.QRotationSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQRotationSensor(parent core.QObject_ITF) *QRotationSensor {
	var tmpValue = NewQRotationSensorFromPointer(C.QRotationSensor_NewQRotationSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQRotationSensor_HasZChanged
func callbackQRotationSensor_HasZChanged(ptr unsafe.Pointer, hasZ C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::hasZChanged"); signal != nil {
		signal.(func(bool))(int8(hasZ) != 0)
	}

}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectHasZChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::hasZChanged", f)
	}
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectHasZChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::hasZChanged")
	}
}

func (ptr *QRotationSensor) HasZChanged(hasZ bool) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_HasZChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hasZ))))
	}
}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_SetHasZ(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hasZ))))
	}
}

//export callbackQRotationSensor_DestroyQRotationSensor
func callbackQRotationSensor_DestroyQRotationSensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::~QRotationSensor"); signal != nil {
		signal.(func())()
	} else {
		NewQRotationSensorFromPointer(ptr).DestroyQRotationSensorDefault()
	}
}

func (ptr *QRotationSensor) ConnectDestroyQRotationSensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::~QRotationSensor", f)
	}
}

func (ptr *QRotationSensor) DisconnectDestroyQRotationSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::~QRotationSensor")
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DestroyQRotationSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensorDefault() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DestroyQRotationSensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QRotationSensor_Type() string {
	return cGoUnpackString(C.QRotationSensor_QRotationSensor_Type())
}

func (ptr *QRotationSensor) Type() string {
	return cGoUnpackString(C.QRotationSensor_QRotationSensor_Type())
}

//export callbackQRotationSensor_Start
func callbackQRotationSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRotationSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QRotationSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::start", f)
	}
}

func (ptr *QRotationSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::start")
	}
}

func (ptr *QRotationSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRotationSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQRotationSensor_Stop
func callbackQRotationSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQRotationSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QRotationSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::stop", f)
	}
}

func (ptr *QRotationSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::stop")
	}
}

func (ptr *QRotationSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QRotationSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQRotationSensor_TimerEvent
func callbackQRotationSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRotationSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRotationSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::timerEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::timerEvent")
	}
}

func (ptr *QRotationSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQRotationSensor_ChildEvent
func callbackQRotationSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRotationSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRotationSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::childEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::childEvent")
	}
}

func (ptr *QRotationSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRotationSensor_ConnectNotify
func callbackQRotationSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRotationSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRotationSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::connectNotify", f)
	}
}

func (ptr *QRotationSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::connectNotify")
	}
}

func (ptr *QRotationSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QRotationSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRotationSensor_CustomEvent
func callbackQRotationSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRotationSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRotationSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::customEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::customEvent")
	}
}

func (ptr *QRotationSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRotationSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRotationSensor_DeleteLater
func callbackQRotationSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRotationSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRotationSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::deleteLater", f)
	}
}

func (ptr *QRotationSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::deleteLater")
	}
}

func (ptr *QRotationSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QRotationSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQRotationSensor_DisconnectNotify
func callbackQRotationSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRotationSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRotationSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::disconnectNotify", f)
	}
}

func (ptr *QRotationSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::disconnectNotify")
	}
}

func (ptr *QRotationSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QRotationSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRotationSensor_Event
func callbackQRotationSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRotationSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRotationSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::event", f)
	}
}

func (ptr *QRotationSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::event")
	}
}

func (ptr *QRotationSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QRotationSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQRotationSensor_EventFilter
func callbackQRotationSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRotationSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRotationSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::eventFilter", f)
	}
}

func (ptr *QRotationSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::eventFilter")
	}
}

func (ptr *QRotationSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QRotationSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQRotationSensor_MetaObject
func callbackQRotationSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QRotationSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQRotationSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRotationSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::metaObject", f)
	}
}

func (ptr *QRotationSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QRotationSensor::metaObject")
	}
}

func (ptr *QRotationSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRotationSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRotationSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRotationSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSensor::AxesOrientationMode
type QSensor__AxesOrientationMode int64

const (
	QSensor__FixedOrientation     = QSensor__AxesOrientationMode(0)
	QSensor__AutomaticOrientation = QSensor__AxesOrientationMode(1)
	QSensor__UserOrientation      = QSensor__AxesOrientationMode(2)
)

//QSensor::Feature
type QSensor__Feature int64

const (
	QSensor__Buffering                 = QSensor__Feature(0)
	QSensor__AlwaysOn                  = QSensor__Feature(1)
	QSensor__GeoValues                 = QSensor__Feature(2)
	QSensor__FieldOfView               = QSensor__Feature(3)
	QSensor__AccelerationMode          = QSensor__Feature(4)
	QSensor__SkipDuplicates            = QSensor__Feature(5)
	QSensor__AxesOrientation           = QSensor__Feature(6)
	QSensor__PressureSensorTemperature = QSensor__Feature(7)
	QSensor__Reserved                  = QSensor__Feature(257)
)

type QSensor struct {
	core.QObject
}

type QSensor_ITF interface {
	core.QObject_ITF
	QSensor_PTR() *QSensor
}

func (p *QSensor) QSensor_PTR() *QSensor {
	return p
}

func (p *QSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSensor(ptr QSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func NewQSensorFromPointer(ptr unsafe.Pointer) *QSensor {
	var n = new(QSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSensor) AxesOrientationMode() QSensor__AxesOrientationMode {
	if ptr.Pointer() != nil {
		return QSensor__AxesOrientationMode(C.QSensor_AxesOrientationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_BufferSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) CurrentOrientation() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_CurrentOrientation(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) DataRate() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_DataRate(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSensor_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensor) EfficientBufferSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_EfficientBufferSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Error() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_Error(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Identifier() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSensor_Identifier(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsAlwaysOn() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsAlwaysOn(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsBusy() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsBusy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsConnectedToBackend() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsConnectedToBackend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) MaxBufferSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_MaxBufferSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) OutputRange() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_OutputRange(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSensorReadingFromPointer(C.QSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAlwaysOn(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(alwaysOn))))
	}
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAxesOrientationMode(ptr.Pointer(), C.longlong(axesOrientationMode))
	}
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetBufferSize(ptr.Pointer(), C.int(int32(bufferSize)))
	}
}

func (ptr *QSensor) SetDataRate(rate int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetDataRate(ptr.Pointer(), C.int(int32(rate)))
	}
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_SetIdentifier(ptr.Pointer(), core.PointerFromQByteArray(identifier))
	}
}

func (ptr *QSensor) SetOutputRange(index int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetOutputRange(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetUserOrientation(ptr.Pointer(), C.int(int32(userOrientation)))
	}
}

func (ptr *QSensor) SkipDuplicates() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_SkipDuplicates(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) Type() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSensor_Type(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) UserOrientation() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_UserOrientation(ptr.Pointer())))
	}
	return 0
}

func NewQSensor(ty core.QByteArray_ITF, parent core.QObject_ITF) *QSensor {
	var tmpValue = NewQSensorFromPointer(C.QSensor_NewQSensor(core.PointerFromQByteArray(ty), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensor_ActiveChanged
func callbackQSensor_ActiveChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::activeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::activeChanged", f)
	}
}

func (ptr *QSensor) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::activeChanged")
	}
}

func (ptr *QSensor) ActiveChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_ActiveChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) AddFilter(filter QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_AddFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

//export callbackQSensor_AlwaysOnChanged
func callbackQSensor_AlwaysOnChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::alwaysOnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAlwaysOnChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::alwaysOnChanged", f)
	}
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAlwaysOnChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::alwaysOnChanged")
	}
}

func (ptr *QSensor) AlwaysOnChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_AlwaysOnChanged(ptr.Pointer())
	}
}

//export callbackQSensor_AvailableSensorsChanged
func callbackQSensor_AvailableSensorsChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::availableSensorsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAvailableSensorsChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::availableSensorsChanged", f)
	}
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAvailableSensorsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::availableSensorsChanged")
	}
}

func (ptr *QSensor) AvailableSensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_AvailableSensorsChanged(ptr.Pointer())
	}
}

//export callbackQSensor_AxesOrientationModeChanged
func callbackQSensor_AxesOrientationModeChanged(ptr unsafe.Pointer, axesOrientationMode C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::axesOrientationModeChanged"); signal != nil {
		signal.(func(QSensor__AxesOrientationMode))(QSensor__AxesOrientationMode(axesOrientationMode))
	}

}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAxesOrientationModeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::axesOrientationModeChanged", f)
	}
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAxesOrientationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::axesOrientationModeChanged")
	}
}

func (ptr *QSensor) AxesOrientationModeChanged(axesOrientationMode QSensor__AxesOrientationMode) {
	if ptr.Pointer() != nil {
		C.QSensor_AxesOrientationModeChanged(ptr.Pointer(), C.longlong(axesOrientationMode))
	}
}

//export callbackQSensor_BufferSizeChanged
func callbackQSensor_BufferSizeChanged(ptr unsafe.Pointer, bufferSize C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::bufferSizeChanged"); signal != nil {
		signal.(func(int))(int(int32(bufferSize)))
	}

}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::bufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::bufferSizeChanged")
	}
}

func (ptr *QSensor) BufferSizeChanged(bufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_BufferSizeChanged(ptr.Pointer(), C.int(int32(bufferSize)))
	}
}

//export callbackQSensor_BusyChanged
func callbackQSensor_BusyChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::busyChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectBusyChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectBusyChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::busyChanged", f)
	}
}

func (ptr *QSensor) DisconnectBusyChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBusyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::busyChanged")
	}
}

func (ptr *QSensor) BusyChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_BusyChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectToBackend() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_ConnectToBackend(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSensor_CurrentOrientationChanged
func callbackQSensor_CurrentOrientationChanged(ptr unsafe.Pointer, currentOrientation C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::currentOrientationChanged"); signal != nil {
		signal.(func(int))(int(int32(currentOrientation)))
	}

}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectCurrentOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::currentOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectCurrentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::currentOrientationChanged")
	}
}

func (ptr *QSensor) CurrentOrientationChanged(currentOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_CurrentOrientationChanged(ptr.Pointer(), C.int(int32(currentOrientation)))
	}
}

//export callbackQSensor_DataRateChanged
func callbackQSensor_DataRateChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::dataRateChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectDataRateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::dataRateChanged", f)
	}
}

func (ptr *QSensor) DisconnectDataRateChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectDataRateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::dataRateChanged")
	}
}

func (ptr *QSensor) DataRateChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DataRateChanged(ptr.Pointer())
	}
}

func QSensor_DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	var tmpValue = core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
	runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QSensor) DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	var tmpValue = core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
	runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

//export callbackQSensor_EfficientBufferSizeChanged
func callbackQSensor_EfficientBufferSizeChanged(ptr unsafe.Pointer, efficientBufferSize C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::efficientBufferSizeChanged"); signal != nil {
		signal.(func(int))(int(int32(efficientBufferSize)))
	}

}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::efficientBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::efficientBufferSizeChanged")
	}
}

func (ptr *QSensor) EfficientBufferSizeChanged(efficientBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_EfficientBufferSizeChanged(ptr.Pointer(), C.int(int32(efficientBufferSize)))
	}
}

func (ptr *QSensor) Filters() []*QSensorFilter {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSensors_PackedList) []*QSensorFilter {
			var out = make([]*QSensorFilter, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQSensorFromPointer(l.data).filters_atList(i)
			}
			return out
		}(C.QSensor_Filters(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsFeatureSupported(ptr.Pointer(), C.longlong(feature)) != 0
	}
	return false
}

//export callbackQSensor_MaxBufferSizeChanged
func callbackQSensor_MaxBufferSizeChanged(ptr unsafe.Pointer, maxBufferSize C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::maxBufferSizeChanged"); signal != nil {
		signal.(func(int))(int(int32(maxBufferSize)))
	}

}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectMaxBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::maxBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectMaxBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::maxBufferSizeChanged")
	}
}

func (ptr *QSensor) MaxBufferSizeChanged(maxBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_MaxBufferSizeChanged(ptr.Pointer(), C.int(int32(maxBufferSize)))
	}
}

//export callbackQSensor_ReadingChanged
func callbackQSensor_ReadingChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::readingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectReadingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectReadingChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::readingChanged", f)
	}
}

func (ptr *QSensor) DisconnectReadingChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectReadingChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::readingChanged")
	}
}

func (ptr *QSensor) ReadingChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_ReadingChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_RemoveFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

//export callbackQSensor_SensorError
func callbackQSensor_SensorError(ptr unsafe.Pointer, error C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::sensorError"); signal != nil {
		signal.(func(int))(int(int32(error)))
	}

}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectSensorError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::sensorError", f)
	}
}

func (ptr *QSensor) DisconnectSensorError() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSensorError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::sensorError")
	}
}

func (ptr *QSensor) SensorError(error int) {
	if ptr.Pointer() != nil {
		C.QSensor_SensorError(ptr.Pointer(), C.int(int32(error)))
	}
}

func QSensor_SensorTypes() []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		var out = make([]*core.QByteArray, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSensorFromPointer(l.data).sensorTypes_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorTypes())
}

func (ptr *QSensor) SensorTypes() []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		var out = make([]*core.QByteArray, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSensorFromPointer(l.data).sensorTypes_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorTypes())
}

func QSensor_SensorsForType(ty core.QByteArray_ITF) []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		var out = make([]*core.QByteArray, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSensorFromPointer(l.data).sensorsForType_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorsForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) SensorsForType(ty core.QByteArray_ITF) []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		var out = make([]*core.QByteArray, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSensorFromPointer(l.data).sensorsForType_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorsForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetCurrentOrientation(ptr.Pointer(), C.int(int32(currentOrientation)))
	}
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetEfficientBufferSize(ptr.Pointer(), C.int(int32(efficientBufferSize)))
	}
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetMaxBufferSize(ptr.Pointer(), C.int(int32(maxBufferSize)))
	}
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetSkipDuplicates(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(skipDuplicates))))
	}
}

//export callbackQSensor_SkipDuplicatesChanged
func callbackQSensor_SkipDuplicatesChanged(ptr unsafe.Pointer, skipDuplicates C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::skipDuplicatesChanged"); signal != nil {
		signal.(func(bool))(int8(skipDuplicates) != 0)
	}

}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectSkipDuplicatesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::skipDuplicatesChanged", f)
	}
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSkipDuplicatesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::skipDuplicatesChanged")
	}
}

func (ptr *QSensor) SkipDuplicatesChanged(skipDuplicates bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SkipDuplicatesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(skipDuplicates))))
	}
}

//export callbackQSensor_Start
func callbackQSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::start", f)
	}
}

func (ptr *QSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::start")
	}
}

func (ptr *QSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSensor_Stop
func callbackQSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::stop", f)
	}
}

func (ptr *QSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::stop")
	}
}

func (ptr *QSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QSensor_Stop(ptr.Pointer())
	}
}

//export callbackQSensor_UserOrientationChanged
func callbackQSensor_UserOrientationChanged(ptr unsafe.Pointer, userOrientation C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::userOrientationChanged"); signal != nil {
		signal.(func(int))(int(int32(userOrientation)))
	}

}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectUserOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::userOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectUserOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::userOrientationChanged")
	}
}

func (ptr *QSensor) UserOrientationChanged(userOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_UserOrientationChanged(ptr.Pointer(), C.int(int32(userOrientation)))
	}
}

//export callbackQSensor_DestroyQSensor
func callbackQSensor_DestroyQSensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::~QSensor"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorFromPointer(ptr).DestroyQSensorDefault()
	}
}

func (ptr *QSensor) ConnectDestroyQSensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::~QSensor", f)
	}
}

func (ptr *QSensor) DisconnectDestroyQSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::~QSensor")
	}
}

func (ptr *QSensor) DestroyQSensor() {
	if ptr.Pointer() != nil {
		C.QSensor_DestroyQSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensor) DestroyQSensorDefault() {
	if ptr.Pointer() != nil {
		C.QSensor_DestroyQSensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensor) filters_atList(i int) *QSensorFilter {
	if ptr.Pointer() != nil {
		return NewQSensorFilterFromPointer(C.QSensor_filters_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QSensor) sensorTypes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSensor_sensorTypes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) sensorsForType_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSensor_sensorsForType_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQSensor_TimerEvent
func callbackQSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::timerEvent", f)
	}
}

func (ptr *QSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::timerEvent")
	}
}

func (ptr *QSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSensor_ChildEvent
func callbackQSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::childEvent", f)
	}
}

func (ptr *QSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::childEvent")
	}
}

func (ptr *QSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensor_ConnectNotify
func callbackQSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::connectNotify", f)
	}
}

func (ptr *QSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::connectNotify")
	}
}

func (ptr *QSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensor_CustomEvent
func callbackQSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::customEvent", f)
	}
}

func (ptr *QSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::customEvent")
	}
}

func (ptr *QSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensor_DeleteLater
func callbackQSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::deleteLater", f)
	}
}

func (ptr *QSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::deleteLater")
	}
}

func (ptr *QSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensor_DisconnectNotify
func callbackQSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::disconnectNotify", f)
	}
}

func (ptr *QSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::disconnectNotify")
	}
}

func (ptr *QSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensor_Event
func callbackQSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::event", f)
	}
}

func (ptr *QSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::event")
	}
}

func (ptr *QSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSensor_EventFilter
func callbackQSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::eventFilter", f)
	}
}

func (ptr *QSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::eventFilter")
	}
}

func (ptr *QSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSensor_MetaObject
func callbackQSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::metaObject", f)
	}
}

func (ptr *QSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensor::metaObject")
	}
}

func (ptr *QSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSensorBackend struct {
	core.QObject
}

type QSensorBackend_ITF interface {
	core.QObject_ITF
	QSensorBackend_PTR() *QSensorBackend
}

func (p *QSensorBackend) QSensorBackend_PTR() *QSensorBackend {
	return p
}

func (p *QSensorBackend) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSensorBackend) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSensorBackend(ptr QSensorBackend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func NewQSensorBackendFromPointer(ptr unsafe.Pointer) *QSensorBackend {
	var n = new(QSensorBackend)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorBackend) DestroyQSensorBackend() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QSensorBackend) AddDataRate(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_AddDataRate(ptr.Pointer(), C.double(min), C.double(max))
	}
}

//export callbackQSensorBackend_IsFeatureSupported
func callbackQSensorBackend_IsFeatureSupported(ptr unsafe.Pointer, feature C.longlong) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::isFeatureSupported"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QSensor__Feature) bool)(QSensor__Feature(feature)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorBackendFromPointer(ptr).IsFeatureSupportedDefault(QSensor__Feature(feature)))))
}

func (ptr *QSensorBackend) ConnectIsFeatureSupported(f func(feature QSensor__Feature) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::isFeatureSupported", f)
	}
}

func (ptr *QSensorBackend) DisconnectIsFeatureSupported() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::isFeatureSupported")
	}
}

func (ptr *QSensorBackend) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_IsFeatureSupported(ptr.Pointer(), C.longlong(feature)) != 0
	}
	return false
}

func (ptr *QSensorBackend) IsFeatureSupportedDefault(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_IsFeatureSupportedDefault(ptr.Pointer(), C.longlong(feature)) != 0
	}
	return false
}

func (ptr *QSensorBackend) SensorBusy() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorBusy(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SensorError(error int) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorError(ptr.Pointer(), C.int(int32(error)))
	}
}

func (ptr *QSensorBackend) AddOutputRange(min float64, max float64, accuracy float64) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_AddOutputRange(ptr.Pointer(), C.double(min), C.double(max), C.double(accuracy))
	}
}

func (ptr *QSensorBackend) NewReadingAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_NewReadingAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSensorReadingFromPointer(C.QSensorBackend_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) Sensor() *QSensor {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSensorFromPointer(C.QSensorBackend_Sensor(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) SensorStopped() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorStopped(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SetDataRates(otherSensor QSensor_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDataRates(ptr.Pointer(), PointerFromQSensor(otherSensor))
	}
}

func (ptr *QSensorBackend) SetDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
		C.QSensorBackend_SetDescription(ptr.Pointer(), descriptionC)
	}
}

//export callbackQSensorBackend_Start
func callbackQSensorBackend_Start(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::start"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorBackend) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::start", f)
	}
}

func (ptr *QSensorBackend) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::start")
	}
}

func (ptr *QSensorBackend) Start() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_Start(ptr.Pointer())
	}
}

//export callbackQSensorBackend_Stop
func callbackQSensorBackend_Stop(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorBackend) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::stop", f)
	}
}

func (ptr *QSensorBackend) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::stop")
	}
}

func (ptr *QSensorBackend) Stop() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_Stop(ptr.Pointer())
	}
}

//export callbackQSensorBackend_TimerEvent
func callbackQSensorBackend_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::timerEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::timerEvent")
	}
}

func (ptr *QSensorBackend) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorBackend) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSensorBackend_ChildEvent
func callbackQSensorBackend_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::childEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::childEvent")
	}
}

func (ptr *QSensorBackend) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorBackend) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorBackend_ConnectNotify
func callbackQSensorBackend_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorBackendFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorBackend) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::connectNotify", f)
	}
}

func (ptr *QSensorBackend) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::connectNotify")
	}
}

func (ptr *QSensorBackend) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorBackend) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorBackend_CustomEvent
func callbackQSensorBackend_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::customEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::customEvent")
	}
}

func (ptr *QSensorBackend) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorBackend) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorBackend_DeleteLater
func callbackQSensorBackend_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorBackendFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorBackend) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::deleteLater", f)
	}
}

func (ptr *QSensorBackend) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::deleteLater")
	}
}

func (ptr *QSensorBackend) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorBackend) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorBackend_DisconnectNotify
func callbackQSensorBackend_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorBackendFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorBackend) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::disconnectNotify", f)
	}
}

func (ptr *QSensorBackend) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::disconnectNotify")
	}
}

func (ptr *QSensorBackend) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorBackend) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorBackend_Event
func callbackQSensorBackend_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorBackendFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorBackend) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::event", f)
	}
}

func (ptr *QSensorBackend) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::event")
	}
}

func (ptr *QSensorBackend) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSensorBackend) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSensorBackend_EventFilter
func callbackQSensorBackend_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorBackendFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorBackend) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::eventFilter", f)
	}
}

func (ptr *QSensorBackend) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::eventFilter")
	}
}

func (ptr *QSensorBackend) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSensorBackend) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSensorBackend_MetaObject
func callbackQSensorBackend_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackend::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSensorBackendFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorBackend) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::metaObject", f)
	}
}

func (ptr *QSensorBackend) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackend::metaObject")
	}
}

func (ptr *QSensorBackend) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorBackend_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorBackend) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorBackend_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSensorBackendFactory struct {
	ptr unsafe.Pointer
}

type QSensorBackendFactory_ITF interface {
	QSensorBackendFactory_PTR() *QSensorBackendFactory
}

func (p *QSensorBackendFactory) QSensorBackendFactory_PTR() *QSensorBackendFactory {
	return p
}

func (p *QSensorBackendFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSensorBackendFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSensorBackendFactory(ptr QSensorBackendFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackendFactory_PTR().Pointer()
	}
	return nil
}

func NewQSensorBackendFactoryFromPointer(ptr unsafe.Pointer) *QSensorBackendFactory {
	var n = new(QSensorBackendFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorBackendFactory) DestroyQSensorBackendFactory() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSensorBackendFactory_CreateBackend
func callbackQSensorBackendFactory_CreateBackend(ptr unsafe.Pointer, sensor unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorBackendFactory::createBackend"); signal != nil {
		return PointerFromQSensorBackend(signal.(func(*QSensor) *QSensorBackend)(NewQSensorFromPointer(sensor)))
	}

	return PointerFromQSensorBackend(nil)
}

func (ptr *QSensorBackendFactory) ConnectCreateBackend(f func(sensor *QSensor) *QSensorBackend) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackendFactory::createBackend", f)
	}
}

func (ptr *QSensorBackendFactory) DisconnectCreateBackend(sensor QSensor_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorBackendFactory::createBackend")
	}
}

func (ptr *QSensorBackendFactory) CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSensorBackendFromPointer(C.QSensorBackendFactory_CreateBackend(ptr.Pointer(), PointerFromQSensor(sensor)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QSensorChangesInterface struct {
	ptr unsafe.Pointer
}

type QSensorChangesInterface_ITF interface {
	QSensorChangesInterface_PTR() *QSensorChangesInterface
}

func (p *QSensorChangesInterface) QSensorChangesInterface_PTR() *QSensorChangesInterface {
	return p
}

func (p *QSensorChangesInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSensorChangesInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSensorChangesInterface(ptr QSensorChangesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorChangesInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorChangesInterfaceFromPointer(ptr unsafe.Pointer) *QSensorChangesInterface {
	var n = new(QSensorChangesInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorChangesInterface) DestroyQSensorChangesInterface() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSensorChangesInterface_SensorsChanged
func callbackQSensorChangesInterface_SensorsChanged(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorChangesInterface::sensorsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorChangesInterface) ConnectSensorsChanged(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorChangesInterface::sensorsChanged", f)
	}
}

func (ptr *QSensorChangesInterface) DisconnectSensorsChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorChangesInterface::sensorsChanged")
	}
}

func (ptr *QSensorChangesInterface) SensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensorChangesInterface_SensorsChanged(ptr.Pointer())
	}
}

type QSensorFilter struct {
	ptr unsafe.Pointer
}

type QSensorFilter_ITF interface {
	QSensorFilter_PTR() *QSensorFilter
}

func (p *QSensorFilter) QSensorFilter_PTR() *QSensorFilter {
	return p
}

func (p *QSensorFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSensorFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSensorFilter(ptr QSensorFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func NewQSensorFilterFromPointer(ptr unsafe.Pointer) *QSensorFilter {
	var n = new(QSensorFilter)
	n.SetPointer(ptr)
	return n
}

//export callbackQSensorFilter_Filter
func callbackQSensorFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QSensorReading) bool)(NewQSensorReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorFilter) ConnectFilter(f func(reading *QSensorReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorFilter::filter", f)
	}
}

func (ptr *QSensorFilter) DisconnectFilter(reading QSensorReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorFilter::filter")
	}
}

func (ptr *QSensorFilter) Filter(reading QSensorReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorFilter_Filter(ptr.Pointer(), PointerFromQSensorReading(reading)) != 0
	}
	return false
}

//export callbackQSensorFilter_DestroyQSensorFilter
func callbackQSensorFilter_DestroyQSensorFilter(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorFilter::~QSensorFilter"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorFilterFromPointer(ptr).DestroyQSensorFilterDefault()
	}
}

func (ptr *QSensorFilter) ConnectDestroyQSensorFilter(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorFilter::~QSensorFilter", f)
	}
}

func (ptr *QSensorFilter) DisconnectDestroyQSensorFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorFilter::~QSensorFilter")
	}
}

func (ptr *QSensorFilter) DestroyQSensorFilter() {
	if ptr.Pointer() != nil {
		C.QSensorFilter_DestroyQSensorFilter(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorFilter) DestroyQSensorFilterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorFilter_DestroyQSensorFilterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorFilter) M_sensor() *QSensor {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSensorFromPointer(C.QSensorFilter_M_sensor(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorFilter) SetM_sensor(vqs QSensor_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorFilter_SetM_sensor(ptr.Pointer(), PointerFromQSensor(vqs))
	}
}

type QSensorGesture struct {
	core.QObject
}

type QSensorGesture_ITF interface {
	core.QObject_ITF
	QSensorGesture_PTR() *QSensorGesture
}

func (p *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return p
}

func (p *QSensorGesture) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSensorGesture) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSensorGesture(ptr QSensorGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesture_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureFromPointer(ptr unsafe.Pointer) *QSensorGesture {
	var n = new(QSensorGesture)
	n.SetPointer(ptr)
	return n
}
func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {
	var idsC = C.CString(strings.Join(ids, "|"))
	defer C.free(unsafe.Pointer(idsC))
	var tmpValue = NewQSensorGestureFromPointer(C.QSensorGesture_NewQSensorGesture(idsC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensorGesture_Detected
func callbackQSensorGesture_Detected(ptr unsafe.Pointer, gestureId C.struct_QtSensors_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::detected"); signal != nil {
		signal.(func(string))(cGoUnpackString(gestureId))
	}

}

func (ptr *QSensorGesture) ConnectDetected(f func(gestureId string)) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ConnectDetected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::detected", f)
	}
}

func (ptr *QSensorGesture) DisconnectDetected() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DisconnectDetected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::detected")
	}
}

func (ptr *QSensorGesture) Detected(gestureId string) {
	if ptr.Pointer() != nil {
		var gestureIdC = C.CString(gestureId)
		defer C.free(unsafe.Pointer(gestureIdC))
		C.QSensorGesture_Detected(ptr.Pointer(), gestureIdC)
	}
}

func (ptr *QSensorGesture) GestureSignals() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSensorGesture_GestureSignals(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSensorGesture_InvalidIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSensorGesture_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) StopDetection() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSensorGesture_ValidIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DestroyQSensorGesture(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorGesture_TimerEvent
func callbackQSensorGesture_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::timerEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::timerEvent")
	}
}

func (ptr *QSensorGesture) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSensorGesture_ChildEvent
func callbackQSensorGesture_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::childEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::childEvent")
	}
}

func (ptr *QSensorGesture) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorGesture_ConnectNotify
func callbackQSensorGesture_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGesture) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::connectNotify", f)
	}
}

func (ptr *QSensorGesture) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::connectNotify")
	}
}

func (ptr *QSensorGesture) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorGesture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGesture_CustomEvent
func callbackQSensorGesture_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::customEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::customEvent")
	}
}

func (ptr *QSensorGesture) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGesture) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorGesture_DeleteLater
func callbackQSensorGesture_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorGestureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorGesture) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::deleteLater", f)
	}
}

func (ptr *QSensorGesture) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::deleteLater")
	}
}

func (ptr *QSensorGesture) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesture) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorGesture_DisconnectNotify
func callbackQSensorGesture_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGesture) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::disconnectNotify", f)
	}
}

func (ptr *QSensorGesture) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::disconnectNotify")
	}
}

func (ptr *QSensorGesture) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorGesture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGesture_Event
func callbackQSensorGesture_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorGesture) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::event", f)
	}
}

func (ptr *QSensorGesture) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::event")
	}
}

func (ptr *QSensorGesture) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGesture_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSensorGesture) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGesture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSensorGesture_EventFilter
func callbackQSensorGesture_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorGesture) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::eventFilter", f)
	}
}

func (ptr *QSensorGesture) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::eventFilter")
	}
}

func (ptr *QSensorGesture) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGesture_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSensorGesture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGesture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSensorGesture_MetaObject
func callbackQSensorGesture_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesture::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSensorGestureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorGesture) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::metaObject", f)
	}
}

func (ptr *QSensorGesture) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesture::metaObject")
	}
}

func (ptr *QSensorGesture) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGesture_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorGesture) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGesture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSensorGestureManager struct {
	core.QObject
}

type QSensorGestureManager_ITF interface {
	core.QObject_ITF
	QSensorGestureManager_PTR() *QSensorGestureManager
}

func (p *QSensorGestureManager) QSensorGestureManager_PTR() *QSensorGestureManager {
	return p
}

func (p *QSensorGestureManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSensorGestureManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSensorGestureManager(ptr QSensorGestureManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureManager_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureManagerFromPointer(ptr unsafe.Pointer) *QSensorGestureManager {
	var n = new(QSensorGestureManager)
	n.SetPointer(ptr)
	return n
}
func NewQSensorGestureManager(parent core.QObject_ITF) *QSensorGestureManager {
	var tmpValue = NewQSensorGestureManagerFromPointer(C.QSensorGestureManager_NewQSensorGestureManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorGestureManager) GestureIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSensorGestureManager_GestureIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSensorGestureManager_NewSensorGestureAvailable
func callbackQSensorGestureManager_NewSensorGestureAvailable(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::newSensorGestureAvailable"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorGestureManager) ConnectNewSensorGestureAvailable(f func()) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNewSensorGestureAvailable(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::newSensorGestureAvailable", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectNewSensorGestureAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNewSensorGestureAvailable(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::newSensorGestureAvailable")
	}
}

func (ptr *QSensorGestureManager) NewSensorGestureAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_NewSensorGestureAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorGestureManager) RecognizerSignals(gestureId string) []string {
	if ptr.Pointer() != nil {
		var gestureIdC = C.CString(gestureId)
		defer C.free(unsafe.Pointer(gestureIdC))
		return strings.Split(cGoUnpackString(C.QSensorGestureManager_RecognizerSignals(ptr.Pointer(), gestureIdC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) RegisterSensorGestureRecognizer(recognizer QSensorGestureRecognizer_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_RegisterSensorGestureRecognizer(ptr.Pointer(), PointerFromQSensorGestureRecognizer(recognizer)) != 0
	}
	return false
}

func QSensorGestureManager_SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	var idC = C.CString(id)
	defer C.free(unsafe.Pointer(idC))
	var tmpValue = NewQSensorGestureRecognizerFromPointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(idC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorGestureManager) SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	var idC = C.CString(id)
	defer C.free(unsafe.Pointer(idC))
	var tmpValue = NewQSensorGestureRecognizerFromPointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(idC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManager() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DestroyQSensorGestureManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorGestureManager_TimerEvent
func callbackQSensorGestureManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::timerEvent", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::timerEvent")
	}
}

func (ptr *QSensorGestureManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSensorGestureManager_ChildEvent
func callbackQSensorGestureManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::childEvent", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::childEvent")
	}
}

func (ptr *QSensorGestureManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorGestureManager_ConnectNotify
func callbackQSensorGestureManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::connectNotify", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::connectNotify")
	}
}

func (ptr *QSensorGestureManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorGestureManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureManager_CustomEvent
func callbackQSensorGestureManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::customEvent", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::customEvent")
	}
}

func (ptr *QSensorGestureManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGestureManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorGestureManager_DeleteLater
func callbackQSensorGestureManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorGestureManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorGestureManager) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::deleteLater", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::deleteLater")
	}
}

func (ptr *QSensorGestureManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorGestureManager_DisconnectNotify
func callbackQSensorGestureManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::disconnectNotify", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::disconnectNotify")
	}
}

func (ptr *QSensorGestureManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorGestureManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureManager_Event
func callbackQSensorGestureManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorGestureManager) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::event", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::event")
	}
}

func (ptr *QSensorGestureManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSensorGestureManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSensorGestureManager_EventFilter
func callbackQSensorGestureManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorGestureManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::eventFilter", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::eventFilter")
	}
}

func (ptr *QSensorGestureManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSensorGestureManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSensorGestureManager_MetaObject
func callbackQSensorGestureManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSensorGestureManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorGestureManager) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::metaObject", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureManager::metaObject")
	}
}

func (ptr *QSensorGestureManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGestureManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorGestureManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGestureManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSensorGesturePluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorGesturePluginInterface_ITF interface {
	QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface
}

func (p *QSensorGesturePluginInterface) QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface {
	return p
}

func (p *QSensorGesturePluginInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSensorGesturePluginInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSensorGesturePluginInterface(ptr QSensorGesturePluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesturePluginInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorGesturePluginInterfaceFromPointer(ptr unsafe.Pointer) *QSensorGesturePluginInterface {
	var n = new(QSensorGesturePluginInterface)
	n.SetPointer(ptr)
	return n
}

//export callbackQSensorGesturePluginInterface_Name
func callbackQSensorGesturePluginInterface_Name(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesturePluginInterface::name"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QSensorGesturePluginInterface) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesturePluginInterface::name", f)
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesturePluginInterface::name")
	}
}

func (ptr *QSensorGesturePluginInterface) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSensorGesturePluginInterface_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQSensorGesturePluginInterface_SupportedIds
func callbackQSensorGesturePluginInterface_SupportedIds(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesturePluginInterface::supportedIds"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(make([]string, 0), "|"))
}

func (ptr *QSensorGesturePluginInterface) ConnectSupportedIds(f func() []string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesturePluginInterface::supportedIds", f)
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectSupportedIds() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesturePluginInterface::supportedIds")
	}
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSensorGesturePluginInterface_SupportedIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface
func callbackQSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGesturePluginInterface::~QSensorGesturePluginInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorGesturePluginInterfaceFromPointer(ptr).DestroyQSensorGesturePluginInterfaceDefault()
	}
}

func (ptr *QSensorGesturePluginInterface) ConnectDestroyQSensorGesturePluginInterface(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesturePluginInterface::~QSensorGesturePluginInterface", f)
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectDestroyQSensorGesturePluginInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGesturePluginInterface::~QSensorGesturePluginInterface")
	}
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {
	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterfaceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesturePluginInterface) createRecognizers_atList(i int) *QSensorGestureRecognizer {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSensorGestureRecognizerFromPointer(C.QSensorGesturePluginInterface_createRecognizers_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QSensorGestureRecognizer struct {
	core.QObject
}

type QSensorGestureRecognizer_ITF interface {
	core.QObject_ITF
	QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer
}

func (p *QSensorGestureRecognizer) QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer {
	return p
}

func (p *QSensorGestureRecognizer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSensorGestureRecognizer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSensorGestureRecognizer(ptr QSensorGestureRecognizer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureRecognizer_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureRecognizerFromPointer(ptr unsafe.Pointer) *QSensorGestureRecognizer {
	var n = new(QSensorGestureRecognizer)
	n.SetPointer(ptr)
	return n
}
func NewQSensorGestureRecognizer(parent core.QObject_ITF) *QSensorGestureRecognizer {
	var tmpValue = NewQSensorGestureRecognizerFromPointer(C.QSensorGestureRecognizer_NewQSensorGestureRecognizer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensorGestureRecognizer_Create
func callbackQSensorGestureRecognizer_Create(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::create"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorGestureRecognizer) ConnectCreate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::create", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::create")
	}
}

func (ptr *QSensorGestureRecognizer) Create() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_Create(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CreateBackend(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_Detected
func callbackQSensorGestureRecognizer_Detected(ptr unsafe.Pointer, gestureId C.struct_QtSensors_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::detected"); signal != nil {
		signal.(func(string))(cGoUnpackString(gestureId))
	}

}

func (ptr *QSensorGestureRecognizer) ConnectDetected(f func(gestureId string)) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ConnectDetected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::detected", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectDetected() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DisconnectDetected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::detected")
	}
}

func (ptr *QSensorGestureRecognizer) Detected(gestureId string) {
	if ptr.Pointer() != nil {
		var gestureIdC = C.CString(gestureId)
		defer C.free(unsafe.Pointer(gestureIdC))
		C.QSensorGestureRecognizer_Detected(ptr.Pointer(), gestureIdC)
	}
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSensorGestureRecognizer_GestureSignals(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSensorGestureRecognizer_Id
func callbackQSensorGestureRecognizer_Id(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::id"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QSensorGestureRecognizer) ConnectId(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::id", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::id")
	}
}

func (ptr *QSensorGestureRecognizer) Id() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSensorGestureRecognizer_Id(ptr.Pointer()))
	}
	return ""
}

//export callbackQSensorGestureRecognizer_IsActive
func callbackQSensorGestureRecognizer_IsActive(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::isActive"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorGestureRecognizer) ConnectIsActive(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::isActive", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectIsActive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::isActive")
	}
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSensorGestureRecognizer_Start
func callbackQSensorGestureRecognizer_Start(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorGestureRecognizer) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::start", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::start")
	}
}

func (ptr *QSensorGestureRecognizer) Start() bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StartBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StartBackend(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_Stop
func callbackQSensorGestureRecognizer_Stop(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::stop"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorGestureRecognizer) ConnectStop(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::stop", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::stop")
	}
}

func (ptr *QSensorGestureRecognizer) Stop() bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_Stop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StopBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StopBackend(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_DestroyQSensorGestureRecognizer
func callbackQSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::~QSensorGestureRecognizer"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).DestroyQSensorGestureRecognizerDefault()
	}
}

func (ptr *QSensorGestureRecognizer) ConnectDestroyQSensorGestureRecognizer(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::~QSensorGestureRecognizer", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectDestroyQSensorGestureRecognizer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::~QSensorGestureRecognizer")
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizerDefault() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorGestureRecognizer_TimerEvent
func callbackQSensorGestureRecognizer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::timerEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::timerEvent")
	}
}

func (ptr *QSensorGestureRecognizer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSensorGestureRecognizer_ChildEvent
func callbackQSensorGestureRecognizer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::childEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::childEvent")
	}
}

func (ptr *QSensorGestureRecognizer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorGestureRecognizer_ConnectNotify
func callbackQSensorGestureRecognizer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::connectNotify", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::connectNotify")
	}
}

func (ptr *QSensorGestureRecognizer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureRecognizer_CustomEvent
func callbackQSensorGestureRecognizer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::customEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::customEvent")
	}
}

func (ptr *QSensorGestureRecognizer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorGestureRecognizer_DeleteLater
func callbackQSensorGestureRecognizer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorGestureRecognizer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::deleteLater", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::deleteLater")
	}
}

func (ptr *QSensorGestureRecognizer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureRecognizer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorGestureRecognizer_DisconnectNotify
func callbackQSensorGestureRecognizer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::disconnectNotify", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::disconnectNotify")
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureRecognizer_Event
func callbackQSensorGestureRecognizer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureRecognizerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorGestureRecognizer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::event", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::event")
	}
}

func (ptr *QSensorGestureRecognizer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSensorGestureRecognizer_EventFilter
func callbackQSensorGestureRecognizer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureRecognizerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorGestureRecognizer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::eventFilter", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::eventFilter")
	}
}

func (ptr *QSensorGestureRecognizer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSensorGestureRecognizer_MetaObject
func callbackQSensorGestureRecognizer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorGestureRecognizer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSensorGestureRecognizerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorGestureRecognizer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::metaObject", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorGestureRecognizer::metaObject")
	}
}

func (ptr *QSensorGestureRecognizer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGestureRecognizer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGestureRecognizer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSensorManager struct {
	ptr unsafe.Pointer
}

type QSensorManager_ITF interface {
	QSensorManager_PTR() *QSensorManager
}

func (p *QSensorManager) QSensorManager_PTR() *QSensorManager {
	return p
}

func (p *QSensorManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSensorManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSensorManager(ptr QSensorManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorManager_PTR().Pointer()
	}
	return nil
}

func NewQSensorManagerFromPointer(ptr unsafe.Pointer) *QSensorManager {
	var n = new(QSensorManager)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorManager) DestroyQSensorManager() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func QSensorManager_CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	var tmpValue = NewQSensorBackendFromPointer(C.QSensorManager_QSensorManager_CreateBackend(PointerFromQSensor(sensor)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorManager) CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	var tmpValue = NewQSensorBackendFromPointer(C.QSensorManager_QSensorManager_CreateBackend(PointerFromQSensor(sensor)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSensorManager_IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {
	return C.QSensorManager_QSensorManager_IsBackendRegistered(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier)) != 0
}

func (ptr *QSensorManager) IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {
	return C.QSensorManager_QSensorManager_IsBackendRegistered(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier)) != 0
}

func QSensorManager_RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {
	C.QSensorManager_QSensorManager_RegisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier), PointerFromQSensorBackendFactory(factory))
}

func (ptr *QSensorManager) RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {
	C.QSensorManager_QSensorManager_RegisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier), PointerFromQSensorBackendFactory(factory))
}

func QSensorManager_SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_SetDefaultBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func (ptr *QSensorManager) SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_SetDefaultBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func QSensorManager_UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_UnregisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func (ptr *QSensorManager) UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_UnregisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

type QSensorPluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorPluginInterface_ITF interface {
	QSensorPluginInterface_PTR() *QSensorPluginInterface
}

func (p *QSensorPluginInterface) QSensorPluginInterface_PTR() *QSensorPluginInterface {
	return p
}

func (p *QSensorPluginInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSensorPluginInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSensorPluginInterface(ptr QSensorPluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorPluginInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorPluginInterfaceFromPointer(ptr unsafe.Pointer) *QSensorPluginInterface {
	var n = new(QSensorPluginInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorPluginInterface) DestroyQSensorPluginInterface() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQSensorPluginInterface_RegisterSensors
func callbackQSensorPluginInterface_RegisterSensors(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorPluginInterface::registerSensors"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorPluginInterface) ConnectRegisterSensors(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorPluginInterface::registerSensors", f)
	}
}

func (ptr *QSensorPluginInterface) DisconnectRegisterSensors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorPluginInterface::registerSensors")
	}
}

func (ptr *QSensorPluginInterface) RegisterSensors() {
	if ptr.Pointer() != nil {
		C.QSensorPluginInterface_RegisterSensors(ptr.Pointer())
	}
}

type QSensorReading struct {
	core.QObject
}

type QSensorReading_ITF interface {
	core.QObject_ITF
	QSensorReading_PTR() *QSensorReading
}

func (p *QSensorReading) QSensorReading_PTR() *QSensorReading {
	return p
}

func (p *QSensorReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSensorReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSensorReading(ptr QSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func NewQSensorReadingFromPointer(ptr unsafe.Pointer) *QSensorReading {
	var n = new(QSensorReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorReading) DestroyQSensorReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QSensorReading) SetTimestamp(timestamp uint64) {
	if ptr.Pointer() != nil {
		C.QSensorReading_SetTimestamp(ptr.Pointer(), C.ulonglong(timestamp))
	}
}

func (ptr *QSensorReading) Timestamp() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QSensorReading_Timestamp(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensorReading) Value(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSensorReading_Value(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorReading) ValueCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensorReading_ValueCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensorReading_TimerEvent
func callbackQSensorReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::timerEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::timerEvent")
	}
}

func (ptr *QSensorReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSensorReading_ChildEvent
func callbackQSensorReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::childEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::childEvent")
	}
}

func (ptr *QSensorReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorReading_ConnectNotify
func callbackQSensorReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::connectNotify", f)
	}
}

func (ptr *QSensorReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::connectNotify")
	}
}

func (ptr *QSensorReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorReading_CustomEvent
func callbackQSensorReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::customEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::customEvent")
	}
}

func (ptr *QSensorReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorReading_DeleteLater
func callbackQSensorReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSensorReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::deleteLater", f)
	}
}

func (ptr *QSensorReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::deleteLater")
	}
}

func (ptr *QSensorReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QSensorReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorReading_DisconnectNotify
func callbackQSensorReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::disconnectNotify", f)
	}
}

func (ptr *QSensorReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::disconnectNotify")
	}
}

func (ptr *QSensorReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSensorReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorReading_Event
func callbackQSensorReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::event", f)
	}
}

func (ptr *QSensorReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::event")
	}
}

func (ptr *QSensorReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSensorReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSensorReading_EventFilter
func callbackQSensorReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::eventFilter", f)
	}
}

func (ptr *QSensorReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::eventFilter")
	}
}

func (ptr *QSensorReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSensorReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSensorReading_MetaObject
func callbackQSensorReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSensorReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSensorReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::metaObject", f)
	}
}

func (ptr *QSensorReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSensorReading::metaObject")
	}
}

func (ptr *QSensorReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTapFilter struct {
	QSensorFilter
}

type QTapFilter_ITF interface {
	QSensorFilter_ITF
	QTapFilter_PTR() *QTapFilter
}

func (p *QTapFilter) QTapFilter_PTR() *QTapFilter {
	return p
}

func (p *QTapFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QTapFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQTapFilter(ptr QTapFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapFilter_PTR().Pointer()
	}
	return nil
}

func NewQTapFilterFromPointer(ptr unsafe.Pointer) *QTapFilter {
	var n = new(QTapFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTapFilter) DestroyQTapFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQTapFilter_Filter
func callbackQTapFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QTapReading) bool)(NewQTapReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTapFilter) ConnectFilter(f func(reading *QTapReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapFilter::filter", f)
	}
}

func (ptr *QTapFilter) DisconnectFilter(reading QTapReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapFilter::filter")
	}
}

func (ptr *QTapFilter) Filter(reading QTapReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapFilter_Filter(ptr.Pointer(), PointerFromQTapReading(reading)) != 0
	}
	return false
}

//QTapReading::TapDirection
type QTapReading__TapDirection int64

const (
	QTapReading__Undefined = QTapReading__TapDirection(0)
	QTapReading__X         = QTapReading__TapDirection(0x0001)
	QTapReading__Y         = QTapReading__TapDirection(0x0002)
	QTapReading__Z         = QTapReading__TapDirection(0x0004)
	QTapReading__X_Pos     = QTapReading__TapDirection(0x0011)
	QTapReading__Y_Pos     = QTapReading__TapDirection(0x0022)
	QTapReading__Z_Pos     = QTapReading__TapDirection(0x0044)
	QTapReading__X_Neg     = QTapReading__TapDirection(0x0101)
	QTapReading__Y_Neg     = QTapReading__TapDirection(0x0202)
	QTapReading__Z_Neg     = QTapReading__TapDirection(0x0404)
	QTapReading__X_Both    = QTapReading__TapDirection(0x0111)
	QTapReading__Y_Both    = QTapReading__TapDirection(0x0222)
	QTapReading__Z_Both    = QTapReading__TapDirection(0x0444)
)

type QTapReading struct {
	QSensorReading
}

type QTapReading_ITF interface {
	QSensorReading_ITF
	QTapReading_PTR() *QTapReading
}

func (p *QTapReading) QTapReading_PTR() *QTapReading {
	return p
}

func (p *QTapReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QTapReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQTapReading(ptr QTapReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapReading_PTR().Pointer()
	}
	return nil
}

func NewQTapReadingFromPointer(ptr unsafe.Pointer) *QTapReading {
	var n = new(QTapReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTapReading) DestroyQTapReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QTapReading) IsDoubleTap() bool {
	if ptr.Pointer() != nil {
		return C.QTapReading_IsDoubleTap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {
	if ptr.Pointer() != nil {
		return QTapReading__TapDirection(C.QTapReading_TapDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {
	if ptr.Pointer() != nil {
		C.QTapReading_SetDoubleTap(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(doubleTap))))
	}
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {
	if ptr.Pointer() != nil {
		C.QTapReading_SetTapDirection(ptr.Pointer(), C.longlong(tapDirection))
	}
}

//export callbackQTapReading_TimerEvent
func callbackQTapReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTapReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::timerEvent", f)
	}
}

func (ptr *QTapReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::timerEvent")
	}
}

func (ptr *QTapReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTapReading_ChildEvent
func callbackQTapReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTapReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::childEvent", f)
	}
}

func (ptr *QTapReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::childEvent")
	}
}

func (ptr *QTapReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTapReading_ConnectNotify
func callbackQTapReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTapReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTapReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::connectNotify", f)
	}
}

func (ptr *QTapReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::connectNotify")
	}
}

func (ptr *QTapReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTapReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTapReading_CustomEvent
func callbackQTapReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTapReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::customEvent", f)
	}
}

func (ptr *QTapReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::customEvent")
	}
}

func (ptr *QTapReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTapReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTapReading_DeleteLater
func callbackQTapReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTapReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTapReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::deleteLater", f)
	}
}

func (ptr *QTapReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::deleteLater")
	}
}

func (ptr *QTapReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QTapReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTapReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTapReading_DisconnectNotify
func callbackQTapReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTapReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTapReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::disconnectNotify", f)
	}
}

func (ptr *QTapReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::disconnectNotify")
	}
}

func (ptr *QTapReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTapReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTapReading_Event
func callbackQTapReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTapReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTapReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::event", f)
	}
}

func (ptr *QTapReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::event")
	}
}

func (ptr *QTapReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTapReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTapReading_EventFilter
func callbackQTapReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTapReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTapReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::eventFilter", f)
	}
}

func (ptr *QTapReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::eventFilter")
	}
}

func (ptr *QTapReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTapReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTapReading_MetaObject
func callbackQTapReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTapReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTapReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::metaObject", f)
	}
}

func (ptr *QTapReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapReading::metaObject")
	}
}

func (ptr *QTapReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTapReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTapReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTapReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTapSensor struct {
	QSensor
}

type QTapSensor_ITF interface {
	QSensor_ITF
	QTapSensor_PTR() *QTapSensor
}

func (p *QTapSensor) QTapSensor_PTR() *QTapSensor {
	return p
}

func (p *QTapSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QTapSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQTapSensor(ptr QTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapSensor_PTR().Pointer()
	}
	return nil
}

func NewQTapSensorFromPointer(ptr unsafe.Pointer) *QTapSensor {
	var n = new(QTapSensor)
	n.SetPointer(ptr)
	return n
}
func (ptr *QTapSensor) Reading() *QTapReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQTapReadingFromPointer(C.QTapSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_ReturnDoubleTapEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {
	if ptr.Pointer() != nil {
		C.QTapSensor_SetReturnDoubleTapEvents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnDoubleTapEvents))))
	}
}

func NewQTapSensor(parent core.QObject_ITF) *QTapSensor {
	var tmpValue = NewQTapSensorFromPointer(C.QTapSensor_NewQTapSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTapSensor_ReturnDoubleTapEventsChanged
func callbackQTapSensor_ReturnDoubleTapEventsChanged(ptr unsafe.Pointer, returnDoubleTapEvents C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::returnDoubleTapEventsChanged"); signal != nil {
		signal.(func(bool))(int8(returnDoubleTapEvents) != 0)
	}

}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::returnDoubleTapEventsChanged", f)
	}
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::returnDoubleTapEventsChanged")
	}
}

func (ptr *QTapSensor) ReturnDoubleTapEventsChanged(returnDoubleTapEvents bool) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ReturnDoubleTapEventsChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnDoubleTapEvents))))
	}
}

//export callbackQTapSensor_DestroyQTapSensor
func callbackQTapSensor_DestroyQTapSensor(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::~QTapSensor"); signal != nil {
		signal.(func())()
	} else {
		NewQTapSensorFromPointer(ptr).DestroyQTapSensorDefault()
	}
}

func (ptr *QTapSensor) ConnectDestroyQTapSensor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::~QTapSensor", f)
	}
}

func (ptr *QTapSensor) DisconnectDestroyQTapSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::~QTapSensor")
	}
}

func (ptr *QTapSensor) DestroyQTapSensor() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DestroyQTapSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapSensor) DestroyQTapSensorDefault() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DestroyQTapSensorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QTapSensor_Type() string {
	return cGoUnpackString(C.QTapSensor_QTapSensor_Type())
}

func (ptr *QTapSensor) Type() string {
	return cGoUnpackString(C.QTapSensor_QTapSensor_Type())
}

//export callbackQTapSensor_Start
func callbackQTapSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTapSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QTapSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::start", f)
	}
}

func (ptr *QTapSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::start")
	}
}

func (ptr *QTapSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTapSensor_Stop
func callbackQTapSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQTapSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QTapSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::stop", f)
	}
}

func (ptr *QTapSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::stop")
	}
}

func (ptr *QTapSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QTapSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QTapSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QTapSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQTapSensor_TimerEvent
func callbackQTapSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTapSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::timerEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::timerEvent")
	}
}

func (ptr *QTapSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTapSensor_ChildEvent
func callbackQTapSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTapSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::childEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::childEvent")
	}
}

func (ptr *QTapSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTapSensor_ConnectNotify
func callbackQTapSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTapSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTapSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::connectNotify", f)
	}
}

func (ptr *QTapSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::connectNotify")
	}
}

func (ptr *QTapSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTapSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTapSensor_CustomEvent
func callbackQTapSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTapSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::customEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::customEvent")
	}
}

func (ptr *QTapSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTapSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTapSensor_DeleteLater
func callbackQTapSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTapSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTapSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::deleteLater", f)
	}
}

func (ptr *QTapSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::deleteLater")
	}
}

func (ptr *QTapSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTapSensor_DisconnectNotify
func callbackQTapSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTapSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTapSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::disconnectNotify", f)
	}
}

func (ptr *QTapSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::disconnectNotify")
	}
}

func (ptr *QTapSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTapSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTapSensor_Event
func callbackQTapSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTapSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTapSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::event", f)
	}
}

func (ptr *QTapSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::event")
	}
}

func (ptr *QTapSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTapSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTapSensor_EventFilter
func callbackQTapSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTapSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTapSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::eventFilter", f)
	}
}

func (ptr *QTapSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::eventFilter")
	}
}

func (ptr *QTapSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTapSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTapSensor_MetaObject
func callbackQTapSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTapSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTapSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTapSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::metaObject", f)
	}
}

func (ptr *QTapSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTapSensor::metaObject")
	}
}

func (ptr *QTapSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTapSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTapSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTapSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTiltFilter struct {
	QSensorFilter
}

type QTiltFilter_ITF interface {
	QSensorFilter_ITF
	QTiltFilter_PTR() *QTiltFilter
}

func (p *QTiltFilter) QTiltFilter_PTR() *QTiltFilter {
	return p
}

func (p *QTiltFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (p *QTiltFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorFilter_PTR().SetPointer(ptr)
	}
}

func PointerFromQTiltFilter(ptr QTiltFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltFilter_PTR().Pointer()
	}
	return nil
}

func NewQTiltFilterFromPointer(ptr unsafe.Pointer) *QTiltFilter {
	var n = new(QTiltFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTiltFilter) DestroyQTiltFilter() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQTiltFilter_Filter
func callbackQTiltFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltFilter::filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QTiltReading) bool)(NewQTiltReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTiltFilter) ConnectFilter(f func(reading *QTiltReading) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltFilter::filter", f)
	}
}

func (ptr *QTiltFilter) DisconnectFilter(reading QTiltReading_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltFilter::filter")
	}
}

func (ptr *QTiltFilter) Filter(reading QTiltReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltFilter_Filter(ptr.Pointer(), PointerFromQTiltReading(reading)) != 0
	}
	return false
}

type QTiltReading struct {
	QSensorReading
}

type QTiltReading_ITF interface {
	QSensorReading_ITF
	QTiltReading_PTR() *QTiltReading
}

func (p *QTiltReading) QTiltReading_PTR() *QTiltReading {
	return p
}

func (p *QTiltReading) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (p *QTiltReading) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensorReading_PTR().SetPointer(ptr)
	}
}

func PointerFromQTiltReading(ptr QTiltReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReading_PTR().Pointer()
	}
	return nil
}

func NewQTiltReadingFromPointer(ptr unsafe.Pointer) *QTiltReading {
	var n = new(QTiltReading)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTiltReading) DestroyQTiltReading() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QTiltReading) XRotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_XRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) YRotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_YRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) SetXRotation(x float64) {
	if ptr.Pointer() != nil {
		C.QTiltReading_SetXRotation(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QTiltReading) SetYRotation(y float64) {
	if ptr.Pointer() != nil {
		C.QTiltReading_SetYRotation(ptr.Pointer(), C.double(y))
	}
}

//export callbackQTiltReading_TimerEvent
func callbackQTiltReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTiltReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::timerEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::timerEvent")
	}
}

func (ptr *QTiltReading) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTiltReading_ChildEvent
func callbackQTiltReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTiltReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::childEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::childEvent")
	}
}

func (ptr *QTiltReading) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTiltReading_ConnectNotify
func callbackQTiltReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTiltReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTiltReading) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::connectNotify", f)
	}
}

func (ptr *QTiltReading) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::connectNotify")
	}
}

func (ptr *QTiltReading) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTiltReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTiltReading_CustomEvent
func callbackQTiltReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTiltReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::customEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::customEvent")
	}
}

func (ptr *QTiltReading) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTiltReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTiltReading_DeleteLater
func callbackQTiltReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTiltReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTiltReading) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::deleteLater", f)
	}
}

func (ptr *QTiltReading) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::deleteLater")
	}
}

func (ptr *QTiltReading) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QTiltReading_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTiltReading_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTiltReading_DisconnectNotify
func callbackQTiltReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTiltReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTiltReading) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::disconnectNotify", f)
	}
}

func (ptr *QTiltReading) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::disconnectNotify")
	}
}

func (ptr *QTiltReading) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTiltReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTiltReading_Event
func callbackQTiltReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTiltReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTiltReading) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::event", f)
	}
}

func (ptr *QTiltReading) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::event")
	}
}

func (ptr *QTiltReading) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltReading_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTiltReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTiltReading_EventFilter
func callbackQTiltReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTiltReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTiltReading) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::eventFilter", f)
	}
}

func (ptr *QTiltReading) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::eventFilter")
	}
}

func (ptr *QTiltReading) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltReading_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTiltReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTiltReading_MetaObject
func callbackQTiltReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltReading::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTiltReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTiltReading) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::metaObject", f)
	}
}

func (ptr *QTiltReading) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltReading::metaObject")
	}
}

func (ptr *QTiltReading) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTiltReading_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTiltReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTiltReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
}

func (p *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return p
}

func (p *QTiltSensor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSensor_PTR().Pointer()
	}
	return nil
}

func (p *QTiltSensor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSensor_PTR().SetPointer(ptr)
	}
}

func PointerFromQTiltSensor(ptr QTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensor_PTR().Pointer()
	}
	return nil
}

func NewQTiltSensorFromPointer(ptr unsafe.Pointer) *QTiltSensor {
	var n = new(QTiltSensor)
	n.SetPointer(ptr)
	return n
}
func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {
	var tmpValue = NewQTiltSensorFromPointer(C.QTiltSensor_NewQTiltSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	if ptr.Pointer() != nil {
		var tmpValue = NewQTiltReadingFromPointer(C.QTiltSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DestroyQTiltSensor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) Calibrate() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(ptr.Pointer())
	}
}

func QTiltSensor_Type() string {
	return cGoUnpackString(C.QTiltSensor_QTiltSensor_Type())
}

func (ptr *QTiltSensor) Type() string {
	return cGoUnpackString(C.QTiltSensor_QTiltSensor_Type())
}

//export callbackQTiltSensor_Start
func callbackQTiltSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTiltSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QTiltSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::start", f)
	}
}

func (ptr *QTiltSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::start")
	}
}

func (ptr *QTiltSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QTiltSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTiltSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return C.QTiltSensor_StartDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTiltSensor_Stop
func callbackQTiltSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQTiltSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QTiltSensor) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::stop", f)
	}
}

func (ptr *QTiltSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::stop")
	}
}

func (ptr *QTiltSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QTiltSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_StopDefault(ptr.Pointer())
	}
}

//export callbackQTiltSensor_TimerEvent
func callbackQTiltSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::timerEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::timerEvent")
	}
}

func (ptr *QTiltSensor) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTiltSensor_ChildEvent
func callbackQTiltSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::childEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::childEvent")
	}
}

func (ptr *QTiltSensor) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTiltSensor_ConnectNotify
func callbackQTiltSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTiltSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTiltSensor) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::connectNotify", f)
	}
}

func (ptr *QTiltSensor) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::connectNotify")
	}
}

func (ptr *QTiltSensor) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTiltSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTiltSensor_CustomEvent
func callbackQTiltSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::customEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::customEvent")
	}
}

func (ptr *QTiltSensor) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTiltSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTiltSensor_DeleteLater
func callbackQTiltSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTiltSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTiltSensor) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::deleteLater", f)
	}
}

func (ptr *QTiltSensor) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::deleteLater")
	}
}

func (ptr *QTiltSensor) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTiltSensor_DisconnectNotify
func callbackQTiltSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTiltSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTiltSensor) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::disconnectNotify", f)
	}
}

func (ptr *QTiltSensor) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::disconnectNotify")
	}
}

func (ptr *QTiltSensor) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTiltSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTiltSensor_Event
func callbackQTiltSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTiltSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTiltSensor) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::event", f)
	}
}

func (ptr *QTiltSensor) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::event")
	}
}

func (ptr *QTiltSensor) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltSensor_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTiltSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTiltSensor_EventFilter
func callbackQTiltSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTiltSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTiltSensor) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::eventFilter", f)
	}
}

func (ptr *QTiltSensor) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::eventFilter")
	}
}

func (ptr *QTiltSensor) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltSensor_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTiltSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTiltSensor_MetaObject
func callbackQTiltSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTiltSensor::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTiltSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTiltSensor) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::metaObject", f)
	}
}

func (ptr *QTiltSensor) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTiltSensor::metaObject")
	}
}

func (ptr *QTiltSensor) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTiltSensor_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTiltSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTiltSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
