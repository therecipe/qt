package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QAccelerometer struct {
	QSensor
}

type QAccelerometer_ITF interface {
	QSensor_ITF
	QAccelerometer_PTR() *QAccelerometer
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
	for len(n.ObjectName()) < len("QAccelerometer_") {
		n.SetObjectName("QAccelerometer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccelerometer) QAccelerometer_PTR() *QAccelerometer {
	return ptr
}

//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int64

const (
	QAccelerometer__Combined = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     = QAccelerometer__AccelerationMode(2)
)

func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {
	defer qt.Recovering("QAccelerometer::accelerationMode")

	if ptr.Pointer() != nil {
		return QAccelerometer__AccelerationMode(C.QAccelerometer_AccelerationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {
	defer qt.Recovering("QAccelerometer::reading")

	if ptr.Pointer() != nil {
		return NewQAccelerometerReadingFromPointer(C.QAccelerometer_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAccelerometer(parent core.QObject_ITF) *QAccelerometer {
	defer qt.Recovering("QAccelerometer::QAccelerometer")

	return NewQAccelerometerFromPointer(C.QAccelerometer_NewQAccelerometer(core.PointerFromQObject(parent)))
}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {
	defer qt.Recovering("connect QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectAccelerationModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accelerationModeChanged", f)
	}
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {
	defer qt.Recovering("disconnect QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectAccelerationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accelerationModeChanged")
	}
}

//export callbackQAccelerometerAccelerationModeChanged
func callbackQAccelerometerAccelerationModeChanged(ptr unsafe.Pointer, ptrName *C.char, accelerationMode C.int) {
	defer qt.Recovering("callback QAccelerometer::accelerationModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "accelerationModeChanged"); signal != nil {
		signal.(func(QAccelerometer__AccelerationMode))(QAccelerometer__AccelerationMode(accelerationMode))
	}

}

func (ptr *QAccelerometer) AccelerationModeChanged(accelerationMode QAccelerometer__AccelerationMode) {
	defer qt.Recovering("QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_AccelerationModeChanged(ptr.Pointer(), C.int(accelerationMode))
	}
}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {
	defer qt.Recovering("QAccelerometer::setAccelerationMode")

	if ptr.Pointer() != nil {
		C.QAccelerometer_SetAccelerationMode(ptr.Pointer(), C.int(accelerationMode))
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {
	defer qt.Recovering("QAccelerometer::~QAccelerometer")

	if ptr.Pointer() != nil {
		C.QAccelerometer_DestroyQAccelerometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAccelerometer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAccelerometerTimerEvent
func callbackQAccelerometerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAccelerometer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAccelerometer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAccelerometerChildEvent
func callbackQAccelerometerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometer::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometer::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAccelerometer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAccelerometer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAccelerometerCustomEvent
func callbackQAccelerometerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometer::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAccelerometer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometer::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAccelerometerFilter struct {
	QSensorFilter
}

type QAccelerometerFilter_ITF interface {
	QSensorFilter_ITF
	QAccelerometerFilter_PTR() *QAccelerometerFilter
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

func (ptr *QAccelerometerFilter) QAccelerometerFilter_PTR() *QAccelerometerFilter {
	return ptr
}

func (ptr *QAccelerometerFilter) Filter(reading QAccelerometerReading_ITF) bool {
	defer qt.Recovering("QAccelerometerFilter::filter")

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

func PointerFromQAccelerometerReading(ptr QAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerReadingFromPointer(ptr unsafe.Pointer) *QAccelerometerReading {
	var n = new(QAccelerometerReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAccelerometerReading_") {
		n.SetObjectName("QAccelerometerReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccelerometerReading) QAccelerometerReading_PTR() *QAccelerometerReading {
	return ptr
}

func (ptr *QAccelerometerReading) X() float64 {
	defer qt.Recovering("QAccelerometerReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Y() float64 {
	defer qt.Recovering("QAccelerometerReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Z() float64 {
	defer qt.Recovering("QAccelerometerReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) SetX(x float64) {
	defer qt.Recovering("QAccelerometerReading::setX")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QAccelerometerReading) SetY(y float64) {
	defer qt.Recovering("QAccelerometerReading::setY")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QAccelerometerReading) SetZ(z float64) {
	defer qt.Recovering("QAccelerometerReading::setZ")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QAccelerometerReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAccelerometerReadingTimerEvent
func callbackQAccelerometerReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometerReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometerReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometerReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAccelerometerReadingChildEvent
func callbackQAccelerometerReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometerReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometerReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometerReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAccelerometerReadingCustomEvent
func callbackQAccelerometerReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometerReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAccelerometerReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAltimeter struct {
	QSensor
}

type QAltimeter_ITF interface {
	QSensor_ITF
	QAltimeter_PTR() *QAltimeter
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
	for len(n.ObjectName()) < len("QAltimeter_") {
		n.SetObjectName("QAltimeter_" + qt.Identifier())
	}
	return n
}

func (ptr *QAltimeter) QAltimeter_PTR() *QAltimeter {
	return ptr
}

func (ptr *QAltimeter) Reading() *QAltimeterReading {
	defer qt.Recovering("QAltimeter::reading")

	if ptr.Pointer() != nil {
		return NewQAltimeterReadingFromPointer(C.QAltimeter_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAltimeter(parent core.QObject_ITF) *QAltimeter {
	defer qt.Recovering("QAltimeter::QAltimeter")

	return NewQAltimeterFromPointer(C.QAltimeter_NewQAltimeter(core.PointerFromQObject(parent)))
}

func (ptr *QAltimeter) DestroyQAltimeter() {
	defer qt.Recovering("QAltimeter::~QAltimeter")

	if ptr.Pointer() != nil {
		C.QAltimeter_DestroyQAltimeter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAltimeter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAltimeter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAltimeter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAltimeterTimerEvent
func callbackQAltimeterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAltimeter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAltimeter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAltimeter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAltimeterChildEvent
func callbackQAltimeterChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAltimeter) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeter::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeter) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeter::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAltimeter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAltimeter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAltimeterCustomEvent
func callbackQAltimeterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAltimeter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeter::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAltimeter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeter::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAltimeterFilter struct {
	QSensorFilter
}

type QAltimeterFilter_ITF interface {
	QSensorFilter_ITF
	QAltimeterFilter_PTR() *QAltimeterFilter
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

func (ptr *QAltimeterFilter) QAltimeterFilter_PTR() *QAltimeterFilter {
	return ptr
}

func (ptr *QAltimeterFilter) Filter(reading QAltimeterReading_ITF) bool {
	defer qt.Recovering("QAltimeterFilter::filter")

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

func PointerFromQAltimeterReading(ptr QAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReading_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterReadingFromPointer(ptr unsafe.Pointer) *QAltimeterReading {
	var n = new(QAltimeterReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAltimeterReading_") {
		n.SetObjectName("QAltimeterReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAltimeterReading) QAltimeterReading_PTR() *QAltimeterReading {
	return ptr
}

func (ptr *QAltimeterReading) Altitude() float64 {
	defer qt.Recovering("QAltimeterReading::altitude")

	if ptr.Pointer() != nil {
		return float64(C.QAltimeterReading_Altitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAltimeterReading) SetAltitude(altitude float64) {
	defer qt.Recovering("QAltimeterReading::setAltitude")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_SetAltitude(ptr.Pointer(), C.double(altitude))
	}
}

func (ptr *QAltimeterReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAltimeterReadingTimerEvent
func callbackQAltimeterReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeterReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeterReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeterReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAltimeterReadingChildEvent
func callbackQAltimeterReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeterReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeterReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeterReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAltimeterReadingCustomEvent
func callbackQAltimeterReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeterReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAltimeterReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAmbientLightFilter struct {
	QSensorFilter
}

type QAmbientLightFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientLightFilter_PTR() *QAmbientLightFilter
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

func (ptr *QAmbientLightFilter) QAmbientLightFilter_PTR() *QAmbientLightFilter {
	return ptr
}

func (ptr *QAmbientLightFilter) Filter(reading QAmbientLightReading_ITF) bool {
	defer qt.Recovering("QAmbientLightFilter::filter")

	if ptr.Pointer() != nil {
		return C.QAmbientLightFilter_Filter(ptr.Pointer(), PointerFromQAmbientLightReading(reading)) != 0
	}
	return false
}

type QAmbientLightReading struct {
	QSensorReading
}

type QAmbientLightReading_ITF interface {
	QSensorReading_ITF
	QAmbientLightReading_PTR() *QAmbientLightReading
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
	for len(n.ObjectName()) < len("QAmbientLightReading_") {
		n.SetObjectName("QAmbientLightReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientLightReading) QAmbientLightReading_PTR() *QAmbientLightReading {
	return ptr
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

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {
	defer qt.Recovering("QAmbientLightReading::lightLevel")

	if ptr.Pointer() != nil {
		return QAmbientLightReading__LightLevel(C.QAmbientLightReading_LightLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {
	defer qt.Recovering("QAmbientLightReading::setLightLevel")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_SetLightLevel(ptr.Pointer(), C.int(lightLevel))
	}
}

func (ptr *QAmbientLightReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientLightReadingTimerEvent
func callbackQAmbientLightReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientLightReadingChildEvent
func callbackQAmbientLightReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientLightReadingCustomEvent
func callbackQAmbientLightReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientLightReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
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
	for len(n.ObjectName()) < len("QAmbientLightSensor_") {
		n.SetObjectName("QAmbientLightSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	defer qt.Recovering("QAmbientLightSensor::reading")

	if ptr.Pointer() != nil {
		return NewQAmbientLightReadingFromPointer(C.QAmbientLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {
	defer qt.Recovering("QAmbientLightSensor::QAmbientLightSensor")

	return NewQAmbientLightSensorFromPointer(C.QAmbientLightSensor_NewQAmbientLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	defer qt.Recovering("QAmbientLightSensor::~QAmbientLightSensor")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientLightSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientLightSensorTimerEvent
func callbackQAmbientLightSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientLightSensorChildEvent
func callbackQAmbientLightSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientLightSensorCustomEvent
func callbackQAmbientLightSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientLightSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAmbientTemperatureFilter struct {
	QSensorFilter
}

type QAmbientTemperatureFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter
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

func (ptr *QAmbientTemperatureFilter) QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter {
	return ptr
}

func (ptr *QAmbientTemperatureFilter) Filter(reading QAmbientTemperatureReading_ITF) bool {
	defer qt.Recovering("QAmbientTemperatureFilter::filter")

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

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureReading {
	var n = new(QAmbientTemperatureReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientTemperatureReading_") {
		n.SetObjectName("QAmbientTemperatureReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientTemperatureReading) QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading {
	return ptr
}

func (ptr *QAmbientTemperatureReading) Temperature() float64 {
	defer qt.Recovering("QAmbientTemperatureReading::temperature")

	if ptr.Pointer() != nil {
		return float64(C.QAmbientTemperatureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientTemperatureReading) SetTemperature(temperature float64) {
	defer qt.Recovering("QAmbientTemperatureReading::setTemperature")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientTemperatureReadingTimerEvent
func callbackQAmbientTemperatureReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientTemperatureReadingChildEvent
func callbackQAmbientTemperatureReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientTemperatureReadingCustomEvent
func callbackQAmbientTemperatureReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QAmbientTemperatureSensor struct {
	QSensor
}

type QAmbientTemperatureSensor_ITF interface {
	QSensor_ITF
	QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor
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
	for len(n.ObjectName()) < len("QAmbientTemperatureSensor_") {
		n.SetObjectName("QAmbientTemperatureSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientTemperatureSensor) QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor {
	return ptr
}

func (ptr *QAmbientTemperatureSensor) Reading() *QAmbientTemperatureReading {
	defer qt.Recovering("QAmbientTemperatureSensor::reading")

	if ptr.Pointer() != nil {
		return NewQAmbientTemperatureReadingFromPointer(C.QAmbientTemperatureSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAmbientTemperatureSensor(parent core.QObject_ITF) *QAmbientTemperatureSensor {
	defer qt.Recovering("QAmbientTemperatureSensor::QAmbientTemperatureSensor")

	return NewQAmbientTemperatureSensorFromPointer(C.QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(core.PointerFromQObject(parent)))
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensor() {
	defer qt.Recovering("QAmbientTemperatureSensor::~QAmbientTemperatureSensor")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientTemperatureSensorTimerEvent
func callbackQAmbientTemperatureSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientTemperatureSensorChildEvent
func callbackQAmbientTemperatureSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientTemperatureSensorCustomEvent
func callbackQAmbientTemperatureSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
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
	for len(n.ObjectName()) < len("QCompass_") {
		n.SetObjectName("QCompass_" + qt.Identifier())
	}
	return n
}

func (ptr *QCompass) QCompass_PTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Reading() *QCompassReading {
	defer qt.Recovering("QCompass::reading")

	if ptr.Pointer() != nil {
		return NewQCompassReadingFromPointer(C.QCompass_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQCompass(parent core.QObject_ITF) *QCompass {
	defer qt.Recovering("QCompass::QCompass")

	return NewQCompassFromPointer(C.QCompass_NewQCompass(core.PointerFromQObject(parent)))
}

func (ptr *QCompass) DestroyQCompass() {
	defer qt.Recovering("QCompass::~QCompass")

	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompass) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCompass::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCompass) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCompass::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCompassTimerEvent
func callbackQCompassTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompass::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCompass) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompass::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompass_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompass) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompass::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompass_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompass) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCompass::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCompass) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCompass::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCompassChildEvent
func callbackQCompassChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompass::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCompass) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompass::childEvent")

	if ptr.Pointer() != nil {
		C.QCompass_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompass) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompass::childEvent")

	if ptr.Pointer() != nil {
		C.QCompass_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompass) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCompass::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCompass) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCompass::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCompassCustomEvent
func callbackQCompassCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompass::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCompass) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCompass::customEvent")

	if ptr.Pointer() != nil {
		C.QCompass_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCompass) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCompass::customEvent")

	if ptr.Pointer() != nil {
		C.QCompass_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QCompassFilter struct {
	QSensorFilter
}

type QCompassFilter_ITF interface {
	QSensorFilter_ITF
	QCompassFilter_PTR() *QCompassFilter
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

func (ptr *QCompassFilter) QCompassFilter_PTR() *QCompassFilter {
	return ptr
}

func (ptr *QCompassFilter) Filter(reading QCompassReading_ITF) bool {
	defer qt.Recovering("QCompassFilter::filter")

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

func PointerFromQCompassReading(ptr QCompassReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassReading_PTR().Pointer()
	}
	return nil
}

func NewQCompassReadingFromPointer(ptr unsafe.Pointer) *QCompassReading {
	var n = new(QCompassReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompassReading_") {
		n.SetObjectName("QCompassReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return ptr
}

func (ptr *QCompassReading) Azimuth() float64 {
	defer qt.Recovering("QCompassReading::azimuth")

	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_Azimuth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) CalibrationLevel() float64 {
	defer qt.Recovering("QCompassReading::calibrationLevel")

	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) SetAzimuth(azimuth float64) {
	defer qt.Recovering("QCompassReading::setAzimuth")

	if ptr.Pointer() != nil {
		C.QCompassReading_SetAzimuth(ptr.Pointer(), C.double(azimuth))
	}
}

func (ptr *QCompassReading) SetCalibrationLevel(calibrationLevel float64) {
	defer qt.Recovering("QCompassReading::setCalibrationLevel")

	if ptr.Pointer() != nil {
		C.QCompassReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QCompassReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCompassReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCompassReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCompassReadingTimerEvent
func callbackQCompassReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompassReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCompassReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompassReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompassReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompassReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompassReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCompassReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCompassReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCompassReadingChildEvent
func callbackQCompassReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompassReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCompassReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompassReading::childEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompassReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompassReading::childEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompassReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCompassReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCompassReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCompassReadingCustomEvent
func callbackQCompassReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompassReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCompassReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCompassReading::customEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCompassReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCompassReading::customEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDistanceFilter struct {
	QSensorFilter
}

type QDistanceFilter_ITF interface {
	QSensorFilter_ITF
	QDistanceFilter_PTR() *QDistanceFilter
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

func (ptr *QDistanceFilter) QDistanceFilter_PTR() *QDistanceFilter {
	return ptr
}

func (ptr *QDistanceFilter) Filter(reading QDistanceReading_ITF) bool {
	defer qt.Recovering("QDistanceFilter::filter")

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

func PointerFromQDistanceReading(ptr QDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReading_PTR().Pointer()
	}
	return nil
}

func NewQDistanceReadingFromPointer(ptr unsafe.Pointer) *QDistanceReading {
	var n = new(QDistanceReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDistanceReading_") {
		n.SetObjectName("QDistanceReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QDistanceReading) QDistanceReading_PTR() *QDistanceReading {
	return ptr
}

func (ptr *QDistanceReading) Distance() float64 {
	defer qt.Recovering("QDistanceReading::distance")

	if ptr.Pointer() != nil {
		return float64(C.QDistanceReading_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDistanceReading) SetDistance(distance float64) {
	defer qt.Recovering("QDistanceReading::setDistance")

	if ptr.Pointer() != nil {
		C.QDistanceReading_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QDistanceReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDistanceReadingTimerEvent
func callbackQDistanceReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDistanceReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDistanceReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDistanceReadingChildEvent
func callbackQDistanceReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceReading::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceReading::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDistanceReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDistanceReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDistanceReadingCustomEvent
func callbackQDistanceReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceReading::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDistanceReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceReading::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensor_ITF interface {
	QSensor_ITF
	QDistanceSensor_PTR() *QDistanceSensor
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
	for len(n.ObjectName()) < len("QDistanceSensor_") {
		n.SetObjectName("QDistanceSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QDistanceSensor) QDistanceSensor_PTR() *QDistanceSensor {
	return ptr
}

func (ptr *QDistanceSensor) Reading() *QDistanceReading {
	defer qt.Recovering("QDistanceSensor::reading")

	if ptr.Pointer() != nil {
		return NewQDistanceReadingFromPointer(C.QDistanceSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQDistanceSensor(parent core.QObject_ITF) *QDistanceSensor {
	defer qt.Recovering("QDistanceSensor::QDistanceSensor")

	return NewQDistanceSensorFromPointer(C.QDistanceSensor_NewQDistanceSensor(core.PointerFromQObject(parent)))
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {
	defer qt.Recovering("QDistanceSensor::~QDistanceSensor")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_DestroyQDistanceSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDistanceSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDistanceSensorTimerEvent
func callbackQDistanceSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDistanceSensorChildEvent
func callbackQDistanceSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDistanceSensorCustomEvent
func callbackQDistanceSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDistanceSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QGyroscope struct {
	QSensor
}

type QGyroscope_ITF interface {
	QSensor_ITF
	QGyroscope_PTR() *QGyroscope
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
	for len(n.ObjectName()) < len("QGyroscope_") {
		n.SetObjectName("QGyroscope_" + qt.Identifier())
	}
	return n
}

func (ptr *QGyroscope) QGyroscope_PTR() *QGyroscope {
	return ptr
}

func (ptr *QGyroscope) Reading() *QGyroscopeReading {
	defer qt.Recovering("QGyroscope::reading")

	if ptr.Pointer() != nil {
		return NewQGyroscopeReadingFromPointer(C.QGyroscope_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQGyroscope(parent core.QObject_ITF) *QGyroscope {
	defer qt.Recovering("QGyroscope::QGyroscope")

	return NewQGyroscopeFromPointer(C.QGyroscope_NewQGyroscope(core.PointerFromQObject(parent)))
}

func (ptr *QGyroscope) DestroyQGyroscope() {
	defer qt.Recovering("QGyroscope::~QGyroscope")

	if ptr.Pointer() != nil {
		C.QGyroscope_DestroyQGyroscope(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGyroscope) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGyroscope::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGyroscope::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGyroscopeTimerEvent
func callbackQGyroscopeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscope::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGyroscope) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGyroscope::timerEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscope) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGyroscope::timerEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscope) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGyroscope::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGyroscope::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGyroscopeChildEvent
func callbackQGyroscopeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscope::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGyroscope) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGyroscope::childEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscope) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGyroscope::childEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscope) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGyroscope::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGyroscope::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGyroscopeCustomEvent
func callbackQGyroscopeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscope::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGyroscope) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGyroscope::customEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGyroscope) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGyroscope::customEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QGyroscopeFilter struct {
	QSensorFilter
}

type QGyroscopeFilter_ITF interface {
	QSensorFilter_ITF
	QGyroscopeFilter_PTR() *QGyroscopeFilter
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

func (ptr *QGyroscopeFilter) QGyroscopeFilter_PTR() *QGyroscopeFilter {
	return ptr
}

func (ptr *QGyroscopeFilter) Filter(reading QGyroscopeReading_ITF) bool {
	defer qt.Recovering("QGyroscopeFilter::filter")

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

func PointerFromQGyroscopeReading(ptr QGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeReadingFromPointer(ptr unsafe.Pointer) *QGyroscopeReading {
	var n = new(QGyroscopeReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGyroscopeReading_") {
		n.SetObjectName("QGyroscopeReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QGyroscopeReading) QGyroscopeReading_PTR() *QGyroscopeReading {
	return ptr
}

func (ptr *QGyroscopeReading) X() float64 {
	defer qt.Recovering("QGyroscopeReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Y() float64 {
	defer qt.Recovering("QGyroscopeReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Z() float64 {
	defer qt.Recovering("QGyroscopeReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) SetX(x float64) {
	defer qt.Recovering("QGyroscopeReading::setX")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QGyroscopeReading) SetY(y float64) {
	defer qt.Recovering("QGyroscopeReading::setY")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QGyroscopeReading) SetZ(z float64) {
	defer qt.Recovering("QGyroscopeReading::setZ")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QGyroscopeReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGyroscopeReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGyroscopeReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGyroscopeReadingTimerEvent
func callbackQGyroscopeReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscopeReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGyroscopeReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGyroscopeReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscopeReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGyroscopeReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscopeReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGyroscopeReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGyroscopeReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGyroscopeReadingChildEvent
func callbackQGyroscopeReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscopeReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGyroscopeReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGyroscopeReading::childEvent")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscopeReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGyroscopeReading::childEvent")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscopeReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGyroscopeReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGyroscopeReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGyroscopeReadingCustomEvent
func callbackQGyroscopeReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscopeReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGyroscopeReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGyroscopeReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGyroscopeReading::customEvent")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGyroscopeReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGyroscopeReading::customEvent")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHolsterFilter struct {
	QSensorFilter
}

type QHolsterFilter_ITF interface {
	QSensorFilter_ITF
	QHolsterFilter_PTR() *QHolsterFilter
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

func (ptr *QHolsterFilter) QHolsterFilter_PTR() *QHolsterFilter {
	return ptr
}

func (ptr *QHolsterFilter) Filter(reading QHolsterReading_ITF) bool {
	defer qt.Recovering("QHolsterFilter::filter")

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

func PointerFromQHolsterReading(ptr QHolsterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterReading_PTR().Pointer()
	}
	return nil
}

func NewQHolsterReadingFromPointer(ptr unsafe.Pointer) *QHolsterReading {
	var n = new(QHolsterReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHolsterReading_") {
		n.SetObjectName("QHolsterReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QHolsterReading) QHolsterReading_PTR() *QHolsterReading {
	return ptr
}

func (ptr *QHolsterReading) Holstered() bool {
	defer qt.Recovering("QHolsterReading::holstered")

	if ptr.Pointer() != nil {
		return C.QHolsterReading_Holstered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHolsterReading) SetHolstered(holstered bool) {
	defer qt.Recovering("QHolsterReading::setHolstered")

	if ptr.Pointer() != nil {
		C.QHolsterReading_SetHolstered(ptr.Pointer(), C.int(qt.GoBoolToInt(holstered)))
	}
}

func (ptr *QHolsterReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHolsterReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHolsterReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHolsterReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHolsterReadingTimerEvent
func callbackQHolsterReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHolsterReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHolsterReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHolsterReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QHolsterReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHolsterReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QHolsterReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHolsterReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHolsterReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHolsterReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHolsterReadingChildEvent
func callbackQHolsterReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHolsterReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHolsterReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHolsterReading::childEvent")

	if ptr.Pointer() != nil {
		C.QHolsterReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHolsterReading::childEvent")

	if ptr.Pointer() != nil {
		C.QHolsterReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHolsterReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHolsterReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHolsterReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHolsterReadingCustomEvent
func callbackQHolsterReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHolsterReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHolsterReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHolsterReading::customEvent")

	if ptr.Pointer() != nil {
		C.QHolsterReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHolsterReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHolsterReading::customEvent")

	if ptr.Pointer() != nil {
		C.QHolsterReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHolsterSensor struct {
	QSensor
}

type QHolsterSensor_ITF interface {
	QSensor_ITF
	QHolsterSensor_PTR() *QHolsterSensor
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
	for len(n.ObjectName()) < len("QHolsterSensor_") {
		n.SetObjectName("QHolsterSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QHolsterSensor) QHolsterSensor_PTR() *QHolsterSensor {
	return ptr
}

func (ptr *QHolsterSensor) Reading() *QHolsterReading {
	defer qt.Recovering("QHolsterSensor::reading")

	if ptr.Pointer() != nil {
		return NewQHolsterReadingFromPointer(C.QHolsterSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQHolsterSensor(parent core.QObject_ITF) *QHolsterSensor {
	defer qt.Recovering("QHolsterSensor::QHolsterSensor")

	return NewQHolsterSensorFromPointer(C.QHolsterSensor_NewQHolsterSensor(core.PointerFromQObject(parent)))
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {
	defer qt.Recovering("QHolsterSensor::~QHolsterSensor")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_DestroyQHolsterSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHolsterSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHolsterSensorTimerEvent
func callbackQHolsterSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHolsterSensorChildEvent
func callbackQHolsterSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHolsterSensorCustomEvent
func callbackQHolsterSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHolsterSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QIRProximityFilter struct {
	QSensorFilter
}

type QIRProximityFilter_ITF interface {
	QSensorFilter_ITF
	QIRProximityFilter_PTR() *QIRProximityFilter
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

func (ptr *QIRProximityFilter) QIRProximityFilter_PTR() *QIRProximityFilter {
	return ptr
}

func (ptr *QIRProximityFilter) Filter(reading QIRProximityReading_ITF) bool {
	defer qt.Recovering("QIRProximityFilter::filter")

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

func PointerFromQIRProximityReading(ptr QIRProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityReadingFromPointer(ptr unsafe.Pointer) *QIRProximityReading {
	var n = new(QIRProximityReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIRProximityReading_") {
		n.SetObjectName("QIRProximityReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QIRProximityReading) QIRProximityReading_PTR() *QIRProximityReading {
	return ptr
}

func (ptr *QIRProximityReading) Reflectance() float64 {
	defer qt.Recovering("QIRProximityReading::reflectance")

	if ptr.Pointer() != nil {
		return float64(C.QIRProximityReading_Reflectance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIRProximityReading) SetReflectance(reflectance float64) {
	defer qt.Recovering("QIRProximityReading::setReflectance")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_SetReflectance(ptr.Pointer(), C.double(reflectance))
	}
}

func (ptr *QIRProximityReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQIRProximityReadingTimerEvent
func callbackQIRProximityReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximityReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximityReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximityReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQIRProximityReadingChildEvent
func callbackQIRProximityReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximityReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximityReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximityReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQIRProximityReadingCustomEvent
func callbackQIRProximityReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximityReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIRProximityReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensor_ITF interface {
	QSensor_ITF
	QIRProximitySensor_PTR() *QIRProximitySensor
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
	for len(n.ObjectName()) < len("QIRProximitySensor_") {
		n.SetObjectName("QIRProximitySensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QIRProximitySensor) QIRProximitySensor_PTR() *QIRProximitySensor {
	return ptr
}

func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {
	defer qt.Recovering("QIRProximitySensor::reading")

	if ptr.Pointer() != nil {
		return NewQIRProximityReadingFromPointer(C.QIRProximitySensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQIRProximitySensor(parent core.QObject_ITF) *QIRProximitySensor {
	defer qt.Recovering("QIRProximitySensor::QIRProximitySensor")

	return NewQIRProximitySensorFromPointer(C.QIRProximitySensor_NewQIRProximitySensor(core.PointerFromQObject(parent)))
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {
	defer qt.Recovering("QIRProximitySensor::~QIRProximitySensor")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DestroyQIRProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIRProximitySensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQIRProximitySensorTimerEvent
func callbackQIRProximitySensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximitySensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximitySensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximitySensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQIRProximitySensorChildEvent
func callbackQIRProximitySensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximitySensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximitySensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximitySensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQIRProximitySensorCustomEvent
func callbackQIRProximitySensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximitySensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIRProximitySensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QLightFilter struct {
	QSensorFilter
}

type QLightFilter_ITF interface {
	QSensorFilter_ITF
	QLightFilter_PTR() *QLightFilter
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

func (ptr *QLightFilter) QLightFilter_PTR() *QLightFilter {
	return ptr
}

func (ptr *QLightFilter) Filter(reading QLightReading_ITF) bool {
	defer qt.Recovering("QLightFilter::filter")

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

func PointerFromQLightReading(ptr QLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReading_PTR().Pointer()
	}
	return nil
}

func NewQLightReadingFromPointer(ptr unsafe.Pointer) *QLightReading {
	var n = new(QLightReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLightReading_") {
		n.SetObjectName("QLightReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QLightReading) QLightReading_PTR() *QLightReading {
	return ptr
}

func (ptr *QLightReading) Lux() float64 {
	defer qt.Recovering("QLightReading::lux")

	if ptr.Pointer() != nil {
		return float64(C.QLightReading_Lux(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightReading) SetLux(lux float64) {
	defer qt.Recovering("QLightReading::setLux")

	if ptr.Pointer() != nil {
		C.QLightReading_SetLux(ptr.Pointer(), C.double(lux))
	}
}

func (ptr *QLightReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLightReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLightReadingTimerEvent
func callbackQLightReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLightReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLightReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLightReadingChildEvent
func callbackQLightReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLightReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLightReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLightReadingCustomEvent
func callbackQLightReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLightReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLightReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QLightSensor struct {
	QSensor
}

type QLightSensor_ITF interface {
	QSensor_ITF
	QLightSensor_PTR() *QLightSensor
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
	for len(n.ObjectName()) < len("QLightSensor_") {
		n.SetObjectName("QLightSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QLightSensor) QLightSensor_PTR() *QLightSensor {
	return ptr
}

func (ptr *QLightSensor) FieldOfView() float64 {
	defer qt.Recovering("QLightSensor::fieldOfView")

	if ptr.Pointer() != nil {
		return float64(C.QLightSensor_FieldOfView(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightSensor) Reading() *QLightReading {
	defer qt.Recovering("QLightSensor::reading")

	if ptr.Pointer() != nil {
		return NewQLightReadingFromPointer(C.QLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQLightSensor(parent core.QObject_ITF) *QLightSensor {
	defer qt.Recovering("QLightSensor::QLightSensor")

	return NewQLightSensorFromPointer(C.QLightSensor_NewQLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QLightSensor) ConnectFieldOfViewChanged(f func(fieldOfView float64)) {
	defer qt.Recovering("connect QLightSensor::fieldOfViewChanged")

	if ptr.Pointer() != nil {
		C.QLightSensor_ConnectFieldOfViewChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fieldOfViewChanged", f)
	}
}

func (ptr *QLightSensor) DisconnectFieldOfViewChanged() {
	defer qt.Recovering("disconnect QLightSensor::fieldOfViewChanged")

	if ptr.Pointer() != nil {
		C.QLightSensor_DisconnectFieldOfViewChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fieldOfViewChanged")
	}
}

//export callbackQLightSensorFieldOfViewChanged
func callbackQLightSensorFieldOfViewChanged(ptr unsafe.Pointer, ptrName *C.char, fieldOfView C.double) {
	defer qt.Recovering("callback QLightSensor::fieldOfViewChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fieldOfViewChanged"); signal != nil {
		signal.(func(float64))(float64(fieldOfView))
	}

}

func (ptr *QLightSensor) FieldOfViewChanged(fieldOfView float64) {
	defer qt.Recovering("QLightSensor::fieldOfViewChanged")

	if ptr.Pointer() != nil {
		C.QLightSensor_FieldOfViewChanged(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) SetFieldOfView(fieldOfView float64) {
	defer qt.Recovering("QLightSensor::setFieldOfView")

	if ptr.Pointer() != nil {
		C.QLightSensor_SetFieldOfView(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) DestroyQLightSensor() {
	defer qt.Recovering("QLightSensor::~QLightSensor")

	if ptr.Pointer() != nil {
		C.QLightSensor_DestroyQLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLightSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLightSensorTimerEvent
func callbackQLightSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLightSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLightSensorChildEvent
func callbackQLightSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLightSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLightSensorCustomEvent
func callbackQLightSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLightSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLightSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMagnetometer struct {
	QSensor
}

type QMagnetometer_ITF interface {
	QSensor_ITF
	QMagnetometer_PTR() *QMagnetometer
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
	for len(n.ObjectName()) < len("QMagnetometer_") {
		n.SetObjectName("QMagnetometer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	defer qt.Recovering("QMagnetometer::reading")

	if ptr.Pointer() != nil {
		return NewQMagnetometerReadingFromPointer(C.QMagnetometer_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	defer qt.Recovering("QMagnetometer::returnGeoValues")

	if ptr.Pointer() != nil {
		return C.QMagnetometer_ReturnGeoValues(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	defer qt.Recovering("QMagnetometer::setReturnGeoValues")

	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(ptr.Pointer(), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {
	defer qt.Recovering("QMagnetometer::QMagnetometer")

	return NewQMagnetometerFromPointer(C.QMagnetometer_NewQMagnetometer(core.PointerFromQObject(parent)))
}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	defer qt.Recovering("connect QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectReturnGeoValuesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnGeoValuesChanged", f)
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	defer qt.Recovering("disconnect QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnGeoValuesChanged")
	}
}

//export callbackQMagnetometerReturnGeoValuesChanged
func callbackQMagnetometerReturnGeoValuesChanged(ptr unsafe.Pointer, ptrName *C.char, returnGeoValues C.int) {
	defer qt.Recovering("callback QMagnetometer::returnGeoValuesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "returnGeoValuesChanged"); signal != nil {
		signal.(func(bool))(int(returnGeoValues) != 0)
	}

}

func (ptr *QMagnetometer) ReturnGeoValuesChanged(returnGeoValues bool) {
	defer qt.Recovering("QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ReturnGeoValuesChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	defer qt.Recovering("QMagnetometer::~QMagnetometer")

	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMagnetometer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMagnetometerTimerEvent
func callbackQMagnetometerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMagnetometer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMagnetometer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMagnetometerChildEvent
func callbackQMagnetometerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometer::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometer::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMagnetometer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMagnetometer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMagnetometerCustomEvent
func callbackQMagnetometerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometer::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMagnetometer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometer::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QMagnetometerFilter struct {
	QSensorFilter
}

type QMagnetometerFilter_ITF interface {
	QSensorFilter_ITF
	QMagnetometerFilter_PTR() *QMagnetometerFilter
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

func (ptr *QMagnetometerFilter) QMagnetometerFilter_PTR() *QMagnetometerFilter {
	return ptr
}

func (ptr *QMagnetometerFilter) Filter(reading QMagnetometerReading_ITF) bool {
	defer qt.Recovering("QMagnetometerFilter::filter")

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

func PointerFromQMagnetometerReading(ptr QMagnetometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerReading_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerReadingFromPointer(ptr unsafe.Pointer) *QMagnetometerReading {
	var n = new(QMagnetometerReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMagnetometerReading_") {
		n.SetObjectName("QMagnetometerReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QMagnetometerReading) QMagnetometerReading_PTR() *QMagnetometerReading {
	return ptr
}

func (ptr *QMagnetometerReading) CalibrationLevel() float64 {
	defer qt.Recovering("QMagnetometerReading::calibrationLevel")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) X() float64 {
	defer qt.Recovering("QMagnetometerReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Y() float64 {
	defer qt.Recovering("QMagnetometerReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Z() float64 {
	defer qt.Recovering("QMagnetometerReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) SetCalibrationLevel(calibrationLevel float64) {
	defer qt.Recovering("QMagnetometerReading::setCalibrationLevel")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QMagnetometerReading) SetX(x float64) {
	defer qt.Recovering("QMagnetometerReading::setX")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QMagnetometerReading) SetY(y float64) {
	defer qt.Recovering("QMagnetometerReading::setY")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QMagnetometerReading) SetZ(z float64) {
	defer qt.Recovering("QMagnetometerReading::setZ")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QMagnetometerReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMagnetometerReadingTimerEvent
func callbackQMagnetometerReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometerReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometerReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometerReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMagnetometerReadingChildEvent
func callbackQMagnetometerReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometerReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometerReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometerReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMagnetometerReadingCustomEvent
func callbackQMagnetometerReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometerReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMagnetometerReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QOrientationFilter struct {
	QSensorFilter
}

type QOrientationFilter_ITF interface {
	QSensorFilter_ITF
	QOrientationFilter_PTR() *QOrientationFilter
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

func (ptr *QOrientationFilter) QOrientationFilter_PTR() *QOrientationFilter {
	return ptr
}

func (ptr *QOrientationFilter) Filter(reading QOrientationReading_ITF) bool {
	defer qt.Recovering("QOrientationFilter::filter")

	if ptr.Pointer() != nil {
		return C.QOrientationFilter_Filter(ptr.Pointer(), PointerFromQOrientationReading(reading)) != 0
	}
	return false
}

type QOrientationReading struct {
	QSensorReading
}

type QOrientationReading_ITF interface {
	QSensorReading_ITF
	QOrientationReading_PTR() *QOrientationReading
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
	for len(n.ObjectName()) < len("QOrientationReading_") {
		n.SetObjectName("QOrientationReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QOrientationReading) QOrientationReading_PTR() *QOrientationReading {
	return ptr
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

func (ptr *QOrientationReading) Orientation() QOrientationReading__Orientation {
	defer qt.Recovering("QOrientationReading::orientation")

	if ptr.Pointer() != nil {
		return QOrientationReading__Orientation(C.QOrientationReading_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {
	defer qt.Recovering("QOrientationReading::setOrientation")

	if ptr.Pointer() != nil {
		C.QOrientationReading_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QOrientationReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQOrientationReadingTimerEvent
func callbackQOrientationReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QOrientationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QOrientationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQOrientationReadingChildEvent
func callbackQOrientationReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QOrientationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QOrientationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQOrientationReadingCustomEvent
func callbackQOrientationReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QOrientationReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensor_ITF interface {
	QSensor_ITF
	QOrientationSensor_PTR() *QOrientationSensor
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
	for len(n.ObjectName()) < len("QOrientationSensor_") {
		n.SetObjectName("QOrientationSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QOrientationSensor) QOrientationSensor_PTR() *QOrientationSensor {
	return ptr
}

func (ptr *QOrientationSensor) Reading() *QOrientationReading {
	defer qt.Recovering("QOrientationSensor::reading")

	if ptr.Pointer() != nil {
		return NewQOrientationReadingFromPointer(C.QOrientationSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQOrientationSensor(parent core.QObject_ITF) *QOrientationSensor {
	defer qt.Recovering("QOrientationSensor::QOrientationSensor")

	return NewQOrientationSensorFromPointer(C.QOrientationSensor_NewQOrientationSensor(core.PointerFromQObject(parent)))
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {
	defer qt.Recovering("QOrientationSensor::~QOrientationSensor")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_DestroyQOrientationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QOrientationSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQOrientationSensorTimerEvent
func callbackQOrientationSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQOrientationSensorChildEvent
func callbackQOrientationSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQOrientationSensorCustomEvent
func callbackQOrientationSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QOrientationSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QPressureFilter struct {
	QSensorFilter
}

type QPressureFilter_ITF interface {
	QSensorFilter_ITF
	QPressureFilter_PTR() *QPressureFilter
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

func (ptr *QPressureFilter) QPressureFilter_PTR() *QPressureFilter {
	return ptr
}

func (ptr *QPressureFilter) Filter(reading QPressureReading_ITF) bool {
	defer qt.Recovering("QPressureFilter::filter")

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

func PointerFromQPressureReading(ptr QPressureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureReading_PTR().Pointer()
	}
	return nil
}

func NewQPressureReadingFromPointer(ptr unsafe.Pointer) *QPressureReading {
	var n = new(QPressureReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPressureReading_") {
		n.SetObjectName("QPressureReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QPressureReading) QPressureReading_PTR() *QPressureReading {
	return ptr
}

func (ptr *QPressureReading) Pressure() float64 {
	defer qt.Recovering("QPressureReading::pressure")

	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Pressure(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) Temperature() float64 {
	defer qt.Recovering("QPressureReading::temperature")

	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) SetPressure(pressure float64) {
	defer qt.Recovering("QPressureReading::setPressure")

	if ptr.Pointer() != nil {
		C.QPressureReading_SetPressure(ptr.Pointer(), C.double(pressure))
	}
}

func (ptr *QPressureReading) SetTemperature(temperature float64) {
	defer qt.Recovering("QPressureReading::setTemperature")

	if ptr.Pointer() != nil {
		C.QPressureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

func (ptr *QPressureReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPressureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPressureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPressureReadingTimerEvent
func callbackQPressureReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPressureReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPressureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPressureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPressureReadingChildEvent
func callbackQPressureReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPressureReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPressureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPressureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPressureReadingCustomEvent
func callbackQPressureReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPressureReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPressureReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QPressureSensor struct {
	QSensor
}

type QPressureSensor_ITF interface {
	QSensor_ITF
	QPressureSensor_PTR() *QPressureSensor
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
	for len(n.ObjectName()) < len("QPressureSensor_") {
		n.SetObjectName("QPressureSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Reading() *QPressureReading {
	defer qt.Recovering("QPressureSensor::reading")

	if ptr.Pointer() != nil {
		return NewQPressureReadingFromPointer(C.QPressureSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {
	defer qt.Recovering("QPressureSensor::QPressureSensor")

	return NewQPressureSensorFromPointer(C.QPressureSensor_NewQPressureSensor(core.PointerFromQObject(parent)))
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	defer qt.Recovering("QPressureSensor::~QPressureSensor")

	if ptr.Pointer() != nil {
		C.QPressureSensor_DestroyQPressureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPressureSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPressureSensorTimerEvent
func callbackQPressureSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPressureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPressureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPressureSensorChildEvent
func callbackQPressureSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPressureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPressureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPressureSensorCustomEvent
func callbackQPressureSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPressureSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QProximityFilter struct {
	QSensorFilter
}

type QProximityFilter_ITF interface {
	QSensorFilter_ITF
	QProximityFilter_PTR() *QProximityFilter
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

func (ptr *QProximityFilter) QProximityFilter_PTR() *QProximityFilter {
	return ptr
}

func (ptr *QProximityFilter) Filter(reading QProximityReading_ITF) bool {
	defer qt.Recovering("QProximityFilter::filter")

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

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) *QProximityReading {
	var n = new(QProximityReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProximityReading_") {
		n.SetObjectName("QProximityReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Close() bool {
	defer qt.Recovering("QProximityReading::close")

	if ptr.Pointer() != nil {
		return C.QProximityReading_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	defer qt.Recovering("QProximityReading::setClose")

	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}

func (ptr *QProximityReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProximityReadingTimerEvent
func callbackQProximityReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximityReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProximityReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximityReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximityReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProximityReadingChildEvent
func callbackQProximityReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximityReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProximityReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximityReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximityReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProximityReadingCustomEvent
func callbackQProximityReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximityReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProximityReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProximityReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QProximitySensor struct {
	QSensor
}

type QProximitySensor_ITF interface {
	QSensor_ITF
	QProximitySensor_PTR() *QProximitySensor
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
	for len(n.ObjectName()) < len("QProximitySensor_") {
		n.SetObjectName("QProximitySensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QProximitySensor) QProximitySensor_PTR() *QProximitySensor {
	return ptr
}

func (ptr *QProximitySensor) Reading() *QProximityReading {
	defer qt.Recovering("QProximitySensor::reading")

	if ptr.Pointer() != nil {
		return NewQProximityReadingFromPointer(C.QProximitySensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQProximitySensor(parent core.QObject_ITF) *QProximitySensor {
	defer qt.Recovering("QProximitySensor::QProximitySensor")

	return NewQProximitySensorFromPointer(C.QProximitySensor_NewQProximitySensor(core.PointerFromQObject(parent)))
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {
	defer qt.Recovering("QProximitySensor::~QProximitySensor")

	if ptr.Pointer() != nil {
		C.QProximitySensor_DestroyQProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QProximitySensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProximitySensorTimerEvent
func callbackQProximitySensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximitySensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProximitySensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProximitySensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProximitySensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QProximitySensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximitySensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProximitySensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QProximitySensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximitySensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProximitySensorChildEvent
func callbackQProximitySensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximitySensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProximitySensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProximitySensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProximitySensor::childEvent")

	if ptr.Pointer() != nil {
		C.QProximitySensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximitySensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProximitySensor::childEvent")

	if ptr.Pointer() != nil {
		C.QProximitySensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximitySensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProximitySensorCustomEvent
func callbackQProximitySensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximitySensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProximitySensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProximitySensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProximitySensor::customEvent")

	if ptr.Pointer() != nil {
		C.QProximitySensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProximitySensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProximitySensor::customEvent")

	if ptr.Pointer() != nil {
		C.QProximitySensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QRotationFilter struct {
	QSensorFilter
}

type QRotationFilter_ITF interface {
	QSensorFilter_ITF
	QRotationFilter_PTR() *QRotationFilter
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

func (ptr *QRotationFilter) QRotationFilter_PTR() *QRotationFilter {
	return ptr
}

func (ptr *QRotationFilter) Filter(reading QRotationReading_ITF) bool {
	defer qt.Recovering("QRotationFilter::filter")

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

func PointerFromQRotationReading(ptr QRotationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationReading_PTR().Pointer()
	}
	return nil
}

func NewQRotationReadingFromPointer(ptr unsafe.Pointer) *QRotationReading {
	var n = new(QRotationReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRotationReading_") {
		n.SetObjectName("QRotationReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QRotationReading) QRotationReading_PTR() *QRotationReading {
	return ptr
}

func (ptr *QRotationReading) X() float64 {
	defer qt.Recovering("QRotationReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Y() float64 {
	defer qt.Recovering("QRotationReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Z() float64 {
	defer qt.Recovering("QRotationReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) SetFromEuler(x float64, y float64, z float64) {
	defer qt.Recovering("QRotationReading::setFromEuler")

	if ptr.Pointer() != nil {
		C.QRotationReading_SetFromEuler(ptr.Pointer(), C.double(x), C.double(y), C.double(z))
	}
}

func (ptr *QRotationReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRotationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRotationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRotationReadingTimerEvent
func callbackQRotationReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRotationReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRotationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRotationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRotationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRotationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRotationReadingChildEvent
func callbackQRotationReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRotationReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRotationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRotationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRotationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRotationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRotationReadingCustomEvent
func callbackQRotationReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRotationReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRotationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRotationReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRotationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QRotationSensor struct {
	QSensor
}

type QRotationSensor_ITF interface {
	QSensor_ITF
	QRotationSensor_PTR() *QRotationSensor
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
	for len(n.ObjectName()) < len("QRotationSensor_") {
		n.SetObjectName("QRotationSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QRotationSensor) QRotationSensor_PTR() *QRotationSensor {
	return ptr
}

func (ptr *QRotationSensor) HasZ() bool {
	defer qt.Recovering("QRotationSensor::hasZ")

	if ptr.Pointer() != nil {
		return C.QRotationSensor_HasZ(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRotationSensor) Reading() *QRotationReading {
	defer qt.Recovering("QRotationSensor::reading")

	if ptr.Pointer() != nil {
		return NewQRotationReadingFromPointer(C.QRotationSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQRotationSensor(parent core.QObject_ITF) *QRotationSensor {
	defer qt.Recovering("QRotationSensor::QRotationSensor")

	return NewQRotationSensorFromPointer(C.QRotationSensor_NewQRotationSensor(core.PointerFromQObject(parent)))
}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {
	defer qt.Recovering("connect QRotationSensor::hasZChanged")

	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectHasZChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hasZChanged", f)
	}
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {
	defer qt.Recovering("disconnect QRotationSensor::hasZChanged")

	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectHasZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hasZChanged")
	}
}

//export callbackQRotationSensorHasZChanged
func callbackQRotationSensorHasZChanged(ptr unsafe.Pointer, ptrName *C.char, hasZ C.int) {
	defer qt.Recovering("callback QRotationSensor::hasZChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasZChanged"); signal != nil {
		signal.(func(bool))(int(hasZ) != 0)
	}

}

func (ptr *QRotationSensor) HasZChanged(hasZ bool) {
	defer qt.Recovering("QRotationSensor::hasZChanged")

	if ptr.Pointer() != nil {
		C.QRotationSensor_HasZChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(hasZ)))
	}
}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {
	defer qt.Recovering("QRotationSensor::setHasZ")

	if ptr.Pointer() != nil {
		C.QRotationSensor_SetHasZ(ptr.Pointer(), C.int(qt.GoBoolToInt(hasZ)))
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {
	defer qt.Recovering("QRotationSensor::~QRotationSensor")

	if ptr.Pointer() != nil {
		C.QRotationSensor_DestroyQRotationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRotationSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRotationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRotationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRotationSensorTimerEvent
func callbackQRotationSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRotationSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRotationSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRotationSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QRotationSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRotationSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QRotationSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRotationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRotationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRotationSensorChildEvent
func callbackQRotationSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRotationSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRotationSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRotationSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QRotationSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRotationSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QRotationSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRotationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRotationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRotationSensorCustomEvent
func callbackQRotationSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRotationSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRotationSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRotationSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QRotationSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRotationSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRotationSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QRotationSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSensor struct {
	core.QObject
}

type QSensor_ITF interface {
	core.QObject_ITF
	QSensor_PTR() *QSensor
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
	for len(n.ObjectName()) < len("QSensor_") {
		n.SetObjectName("QSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensor) QSensor_PTR() *QSensor {
	return ptr
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

func (ptr *QSensor) AxesOrientationMode() QSensor__AxesOrientationMode {
	defer qt.Recovering("QSensor::axesOrientationMode")

	if ptr.Pointer() != nil {
		return QSensor__AxesOrientationMode(C.QSensor_AxesOrientationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) BufferSize() int {
	defer qt.Recovering("QSensor::bufferSize")

	if ptr.Pointer() != nil {
		return int(C.QSensor_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) CurrentOrientation() int {
	defer qt.Recovering("QSensor::currentOrientation")

	if ptr.Pointer() != nil {
		return int(C.QSensor_CurrentOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) DataRate() int {
	defer qt.Recovering("QSensor::dataRate")

	if ptr.Pointer() != nil {
		return int(C.QSensor_DataRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Description() string {
	defer qt.Recovering("QSensor::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensor_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensor) EfficientBufferSize() int {
	defer qt.Recovering("QSensor::efficientBufferSize")

	if ptr.Pointer() != nil {
		return int(C.QSensor_EfficientBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Error() int {
	defer qt.Recovering("QSensor::error")

	if ptr.Pointer() != nil {
		return int(C.QSensor_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Identifier() *core.QByteArray {
	defer qt.Recovering("QSensor::identifier")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSensor_Identifier(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) IsActive() bool {
	defer qt.Recovering("QSensor::isActive")

	if ptr.Pointer() != nil {
		return C.QSensor_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsAlwaysOn() bool {
	defer qt.Recovering("QSensor::isAlwaysOn")

	if ptr.Pointer() != nil {
		return C.QSensor_IsAlwaysOn(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsBusy() bool {
	defer qt.Recovering("QSensor::isBusy")

	if ptr.Pointer() != nil {
		return C.QSensor_IsBusy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsConnectedToBackend() bool {
	defer qt.Recovering("QSensor::isConnectedToBackend")

	if ptr.Pointer() != nil {
		return C.QSensor_IsConnectedToBackend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) MaxBufferSize() int {
	defer qt.Recovering("QSensor::maxBufferSize")

	if ptr.Pointer() != nil {
		return int(C.QSensor_MaxBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) OutputRange() int {
	defer qt.Recovering("QSensor::outputRange")

	if ptr.Pointer() != nil {
		return int(C.QSensor_OutputRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Reading() *QSensorReading {
	defer qt.Recovering("QSensor::reading")

	if ptr.Pointer() != nil {
		return NewQSensorReadingFromPointer(C.QSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) SetActive(active bool) {
	defer qt.Recovering("QSensor::setActive")

	if ptr.Pointer() != nil {
		C.QSensor_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {
	defer qt.Recovering("QSensor::setAlwaysOn")

	if ptr.Pointer() != nil {
		C.QSensor_SetAlwaysOn(ptr.Pointer(), C.int(qt.GoBoolToInt(alwaysOn)))
	}
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {
	defer qt.Recovering("QSensor::setAxesOrientationMode")

	if ptr.Pointer() != nil {
		C.QSensor_SetAxesOrientationMode(ptr.Pointer(), C.int(axesOrientationMode))
	}
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {
	defer qt.Recovering("QSensor::setBufferSize")

	if ptr.Pointer() != nil {
		C.QSensor_SetBufferSize(ptr.Pointer(), C.int(bufferSize))
	}
}

func (ptr *QSensor) SetDataRate(rate int) {
	defer qt.Recovering("QSensor::setDataRate")

	if ptr.Pointer() != nil {
		C.QSensor_SetDataRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArray_ITF) {
	defer qt.Recovering("QSensor::setIdentifier")

	if ptr.Pointer() != nil {
		C.QSensor_SetIdentifier(ptr.Pointer(), core.PointerFromQByteArray(identifier))
	}
}

func (ptr *QSensor) SetOutputRange(index int) {
	defer qt.Recovering("QSensor::setOutputRange")

	if ptr.Pointer() != nil {
		C.QSensor_SetOutputRange(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {
	defer qt.Recovering("QSensor::setUserOrientation")

	if ptr.Pointer() != nil {
		C.QSensor_SetUserOrientation(ptr.Pointer(), C.int(userOrientation))
	}
}

func (ptr *QSensor) SkipDuplicates() bool {
	defer qt.Recovering("QSensor::skipDuplicates")

	if ptr.Pointer() != nil {
		return C.QSensor_SkipDuplicates(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) Type() *core.QByteArray {
	defer qt.Recovering("QSensor::type")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSensor_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) UserOrientation() int {
	defer qt.Recovering("QSensor::userOrientation")

	if ptr.Pointer() != nil {
		return int(C.QSensor_UserOrientation(ptr.Pointer()))
	}
	return 0
}

func NewQSensor(ty core.QByteArray_ITF, parent core.QObject_ITF) *QSensor {
	defer qt.Recovering("QSensor::QSensor")

	return NewQSensorFromPointer(C.QSensor_NewQSensor(core.PointerFromQByteArray(ty), core.PointerFromQObject(parent)))
}

func (ptr *QSensor) ConnectActiveChanged(f func()) {
	defer qt.Recovering("connect QSensor::activeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QSensor) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QSensor::activeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQSensorActiveChanged
func callbackQSensorActiveChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensor::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ActiveChanged() {
	defer qt.Recovering("QSensor::activeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ActiveChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) AddFilter(filter QSensorFilter_ITF) {
	defer qt.Recovering("QSensor::addFilter")

	if ptr.Pointer() != nil {
		C.QSensor_AddFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {
	defer qt.Recovering("connect QSensor::alwaysOnChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectAlwaysOnChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alwaysOnChanged", f)
	}
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {
	defer qt.Recovering("disconnect QSensor::alwaysOnChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAlwaysOnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alwaysOnChanged")
	}
}

//export callbackQSensorAlwaysOnChanged
func callbackQSensorAlwaysOnChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensor::alwaysOnChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "alwaysOnChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) AlwaysOnChanged() {
	defer qt.Recovering("QSensor::alwaysOnChanged")

	if ptr.Pointer() != nil {
		C.QSensor_AlwaysOnChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {
	defer qt.Recovering("connect QSensor::availableSensorsChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectAvailableSensorsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableSensorsChanged", f)
	}
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {
	defer qt.Recovering("disconnect QSensor::availableSensorsChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAvailableSensorsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableSensorsChanged")
	}
}

//export callbackQSensorAvailableSensorsChanged
func callbackQSensorAvailableSensorsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensor::availableSensorsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableSensorsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) AvailableSensorsChanged() {
	defer qt.Recovering("QSensor::availableSensorsChanged")

	if ptr.Pointer() != nil {
		C.QSensor_AvailableSensorsChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {
	defer qt.Recovering("connect QSensor::axesOrientationModeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectAxesOrientationModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "axesOrientationModeChanged", f)
	}
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {
	defer qt.Recovering("disconnect QSensor::axesOrientationModeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAxesOrientationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "axesOrientationModeChanged")
	}
}

//export callbackQSensorAxesOrientationModeChanged
func callbackQSensorAxesOrientationModeChanged(ptr unsafe.Pointer, ptrName *C.char, axesOrientationMode C.int) {
	defer qt.Recovering("callback QSensor::axesOrientationModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "axesOrientationModeChanged"); signal != nil {
		signal.(func(QSensor__AxesOrientationMode))(QSensor__AxesOrientationMode(axesOrientationMode))
	}

}

func (ptr *QSensor) AxesOrientationModeChanged(axesOrientationMode QSensor__AxesOrientationMode) {
	defer qt.Recovering("QSensor::axesOrientationModeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_AxesOrientationModeChanged(ptr.Pointer(), C.int(axesOrientationMode))
	}
}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {
	defer qt.Recovering("connect QSensor::bufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {
	defer qt.Recovering("disconnect QSensor::bufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferSizeChanged")
	}
}

//export callbackQSensorBufferSizeChanged
func callbackQSensorBufferSizeChanged(ptr unsafe.Pointer, ptrName *C.char, bufferSize C.int) {
	defer qt.Recovering("callback QSensor::bufferSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferSizeChanged"); signal != nil {
		signal.(func(int))(int(bufferSize))
	}

}

func (ptr *QSensor) BufferSizeChanged(bufferSize int) {
	defer qt.Recovering("QSensor::bufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_BufferSizeChanged(ptr.Pointer(), C.int(bufferSize))
	}
}

func (ptr *QSensor) ConnectBusyChanged(f func()) {
	defer qt.Recovering("connect QSensor::busyChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectBusyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "busyChanged", f)
	}
}

func (ptr *QSensor) DisconnectBusyChanged() {
	defer qt.Recovering("disconnect QSensor::busyChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBusyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "busyChanged")
	}
}

//export callbackQSensorBusyChanged
func callbackQSensorBusyChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensor::busyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "busyChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) BusyChanged() {
	defer qt.Recovering("QSensor::busyChanged")

	if ptr.Pointer() != nil {
		C.QSensor_BusyChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectToBackend() bool {
	defer qt.Recovering("QSensor::connectToBackend")

	if ptr.Pointer() != nil {
		return C.QSensor_ConnectToBackend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {
	defer qt.Recovering("connect QSensor::currentOrientationChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectCurrentOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {
	defer qt.Recovering("disconnect QSensor::currentOrientationChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectCurrentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentOrientationChanged")
	}
}

//export callbackQSensorCurrentOrientationChanged
func callbackQSensorCurrentOrientationChanged(ptr unsafe.Pointer, ptrName *C.char, currentOrientation C.int) {
	defer qt.Recovering("callback QSensor::currentOrientationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentOrientationChanged"); signal != nil {
		signal.(func(int))(int(currentOrientation))
	}

}

func (ptr *QSensor) CurrentOrientationChanged(currentOrientation int) {
	defer qt.Recovering("QSensor::currentOrientationChanged")

	if ptr.Pointer() != nil {
		C.QSensor_CurrentOrientationChanged(ptr.Pointer(), C.int(currentOrientation))
	}
}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {
	defer qt.Recovering("connect QSensor::dataRateChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectDataRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataRateChanged", f)
	}
}

func (ptr *QSensor) DisconnectDataRateChanged() {
	defer qt.Recovering("disconnect QSensor::dataRateChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectDataRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataRateChanged")
	}
}

//export callbackQSensorDataRateChanged
func callbackQSensorDataRateChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensor::dataRateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "dataRateChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) DataRateChanged() {
	defer qt.Recovering("QSensor::dataRateChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DataRateChanged(ptr.Pointer())
	}
}

func QSensor_DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QSensor::defaultSensorForType")

	return core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {
	defer qt.Recovering("connect QSensor::efficientBufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {
	defer qt.Recovering("disconnect QSensor::efficientBufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged")
	}
}

//export callbackQSensorEfficientBufferSizeChanged
func callbackQSensorEfficientBufferSizeChanged(ptr unsafe.Pointer, ptrName *C.char, efficientBufferSize C.int) {
	defer qt.Recovering("callback QSensor::efficientBufferSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "efficientBufferSizeChanged"); signal != nil {
		signal.(func(int))(int(efficientBufferSize))
	}

}

func (ptr *QSensor) EfficientBufferSizeChanged(efficientBufferSize int) {
	defer qt.Recovering("QSensor::efficientBufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_EfficientBufferSizeChanged(ptr.Pointer(), C.int(efficientBufferSize))
	}
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {
	defer qt.Recovering("QSensor::isFeatureSupported")

	if ptr.Pointer() != nil {
		return C.QSensor_IsFeatureSupported(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {
	defer qt.Recovering("connect QSensor::maxBufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectMaxBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maxBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {
	defer qt.Recovering("disconnect QSensor::maxBufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectMaxBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maxBufferSizeChanged")
	}
}

//export callbackQSensorMaxBufferSizeChanged
func callbackQSensorMaxBufferSizeChanged(ptr unsafe.Pointer, ptrName *C.char, maxBufferSize C.int) {
	defer qt.Recovering("callback QSensor::maxBufferSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "maxBufferSizeChanged"); signal != nil {
		signal.(func(int))(int(maxBufferSize))
	}

}

func (ptr *QSensor) MaxBufferSizeChanged(maxBufferSize int) {
	defer qt.Recovering("QSensor::maxBufferSizeChanged")

	if ptr.Pointer() != nil {
		C.QSensor_MaxBufferSizeChanged(ptr.Pointer(), C.int(maxBufferSize))
	}
}

func (ptr *QSensor) ConnectReadingChanged(f func()) {
	defer qt.Recovering("connect QSensor::readingChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectReadingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readingChanged", f)
	}
}

func (ptr *QSensor) DisconnectReadingChanged() {
	defer qt.Recovering("disconnect QSensor::readingChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectReadingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readingChanged")
	}
}

//export callbackQSensorReadingChanged
func callbackQSensorReadingChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensor::readingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "readingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensor) ReadingChanged() {
	defer qt.Recovering("QSensor::readingChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ReadingChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilter_ITF) {
	defer qt.Recovering("QSensor::removeFilter")

	if ptr.Pointer() != nil {
		C.QSensor_RemoveFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {
	defer qt.Recovering("connect QSensor::sensorError")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectSensorError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sensorError", f)
	}
}

func (ptr *QSensor) DisconnectSensorError() {
	defer qt.Recovering("disconnect QSensor::sensorError")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSensorError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sensorError")
	}
}

//export callbackQSensorSensorError
func callbackQSensorSensorError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QSensor::sensorError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sensorError"); signal != nil {
		signal.(func(int))(int(error))
	}

}

func (ptr *QSensor) SensorError(error int) {
	defer qt.Recovering("QSensor::sensorError")

	if ptr.Pointer() != nil {
		C.QSensor_SensorError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {
	defer qt.Recovering("QSensor::setCurrentOrientation")

	if ptr.Pointer() != nil {
		C.QSensor_SetCurrentOrientation(ptr.Pointer(), C.int(currentOrientation))
	}
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {
	defer qt.Recovering("QSensor::setEfficientBufferSize")

	if ptr.Pointer() != nil {
		C.QSensor_SetEfficientBufferSize(ptr.Pointer(), C.int(efficientBufferSize))
	}
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {
	defer qt.Recovering("QSensor::setMaxBufferSize")

	if ptr.Pointer() != nil {
		C.QSensor_SetMaxBufferSize(ptr.Pointer(), C.int(maxBufferSize))
	}
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {
	defer qt.Recovering("QSensor::setSkipDuplicates")

	if ptr.Pointer() != nil {
		C.QSensor_SetSkipDuplicates(ptr.Pointer(), C.int(qt.GoBoolToInt(skipDuplicates)))
	}
}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {
	defer qt.Recovering("connect QSensor::skipDuplicatesChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectSkipDuplicatesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "skipDuplicatesChanged", f)
	}
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {
	defer qt.Recovering("disconnect QSensor::skipDuplicatesChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSkipDuplicatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "skipDuplicatesChanged")
	}
}

//export callbackQSensorSkipDuplicatesChanged
func callbackQSensorSkipDuplicatesChanged(ptr unsafe.Pointer, ptrName *C.char, skipDuplicates C.int) {
	defer qt.Recovering("callback QSensor::skipDuplicatesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "skipDuplicatesChanged"); signal != nil {
		signal.(func(bool))(int(skipDuplicates) != 0)
	}

}

func (ptr *QSensor) SkipDuplicatesChanged(skipDuplicates bool) {
	defer qt.Recovering("QSensor::skipDuplicatesChanged")

	if ptr.Pointer() != nil {
		C.QSensor_SkipDuplicatesChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(skipDuplicates)))
	}
}

func (ptr *QSensor) Start() bool {
	defer qt.Recovering("QSensor::start")

	if ptr.Pointer() != nil {
		return C.QSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) Stop() {
	defer qt.Recovering("QSensor::stop")

	if ptr.Pointer() != nil {
		C.QSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {
	defer qt.Recovering("connect QSensor::userOrientationChanged")

	if ptr.Pointer() != nil {
		C.QSensor_ConnectUserOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "userOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {
	defer qt.Recovering("disconnect QSensor::userOrientationChanged")

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectUserOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "userOrientationChanged")
	}
}

//export callbackQSensorUserOrientationChanged
func callbackQSensorUserOrientationChanged(ptr unsafe.Pointer, ptrName *C.char, userOrientation C.int) {
	defer qt.Recovering("callback QSensor::userOrientationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "userOrientationChanged"); signal != nil {
		signal.(func(int))(int(userOrientation))
	}

}

func (ptr *QSensor) UserOrientationChanged(userOrientation int) {
	defer qt.Recovering("QSensor::userOrientationChanged")

	if ptr.Pointer() != nil {
		C.QSensor_UserOrientationChanged(ptr.Pointer(), C.int(userOrientation))
	}
}

func (ptr *QSensor) DestroyQSensor() {
	defer qt.Recovering("QSensor::~QSensor")

	if ptr.Pointer() != nil {
		C.QSensor_DestroyQSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorTimerEvent
func callbackQSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorChildEvent
func callbackQSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorCustomEvent
func callbackQSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSensorBackend struct {
	core.QObject
}

type QSensorBackend_ITF interface {
	core.QObject_ITF
	QSensorBackend_PTR() *QSensorBackend
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
	for len(n.ObjectName()) < len("QSensorBackend_") {
		n.SetObjectName("QSensorBackend_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorBackend) QSensorBackend_PTR() *QSensorBackend {
	return ptr
}

func (ptr *QSensorBackend) AddDataRate(min float64, max float64) {
	defer qt.Recovering("QSensorBackend::addDataRate")

	if ptr.Pointer() != nil {
		C.QSensorBackend_AddDataRate(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QSensorBackend) IsFeatureSupported(feature QSensor__Feature) bool {
	defer qt.Recovering("QSensorBackend::isFeatureSupported")

	if ptr.Pointer() != nil {
		return C.QSensorBackend_IsFeatureSupported(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensorBackend) SensorBusy() {
	defer qt.Recovering("QSensorBackend::sensorBusy")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorBusy(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SensorError(error int) {
	defer qt.Recovering("QSensorBackend::sensorError")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QSensorBackend) AddOutputRange(min float64, max float64, accuracy float64) {
	defer qt.Recovering("QSensorBackend::addOutputRange")

	if ptr.Pointer() != nil {
		C.QSensorBackend_AddOutputRange(ptr.Pointer(), C.double(min), C.double(max), C.double(accuracy))
	}
}

func (ptr *QSensorBackend) NewReadingAvailable() {
	defer qt.Recovering("QSensorBackend::newReadingAvailable")

	if ptr.Pointer() != nil {
		C.QSensorBackend_NewReadingAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) Reading() *QSensorReading {
	defer qt.Recovering("QSensorBackend::reading")

	if ptr.Pointer() != nil {
		return NewQSensorReadingFromPointer(C.QSensorBackend_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorBackend) Sensor() *QSensor {
	defer qt.Recovering("QSensorBackend::sensor")

	if ptr.Pointer() != nil {
		return NewQSensorFromPointer(C.QSensorBackend_Sensor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorBackend) SensorStopped() {
	defer qt.Recovering("QSensorBackend::sensorStopped")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorStopped(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SetDataRates(otherSensor QSensor_ITF) {
	defer qt.Recovering("QSensorBackend::setDataRates")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDataRates(ptr.Pointer(), PointerFromQSensor(otherSensor))
	}
}

func (ptr *QSensorBackend) SetDescription(description string) {
	defer qt.Recovering("QSensorBackend::setDescription")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QSensorBackend) Start() {
	defer qt.Recovering("QSensorBackend::start")

	if ptr.Pointer() != nil {
		C.QSensorBackend_Start(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) Stop() {
	defer qt.Recovering("QSensorBackend::stop")

	if ptr.Pointer() != nil {
		C.QSensorBackend_Stop(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorBackend::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorBackend::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorBackendTimerEvent
func callbackQSensorBackendTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorBackend::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorBackend::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorBackend_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorBackend) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorBackend::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorBackend_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorBackend) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorBackend::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorBackend::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorBackendChildEvent
func callbackQSensorBackendChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorBackend::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorBackend::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorBackend_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorBackend) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorBackend::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorBackend_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorBackend) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorBackend::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorBackend::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorBackendCustomEvent
func callbackQSensorBackendCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorBackend::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorBackend::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorBackend_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorBackend) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorBackend::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorBackend_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSensorBackendFactory struct {
	ptr unsafe.Pointer
}

type QSensorBackendFactory_ITF interface {
	QSensorBackendFactory_PTR() *QSensorBackendFactory
}

func (p *QSensorBackendFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorBackendFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QSensorBackendFactory) QSensorBackendFactory_PTR() *QSensorBackendFactory {
	return ptr
}

func (ptr *QSensorBackendFactory) CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	defer qt.Recovering("QSensorBackendFactory::createBackend")

	if ptr.Pointer() != nil {
		return NewQSensorBackendFromPointer(C.QSensorBackendFactory_CreateBackend(ptr.Pointer(), PointerFromQSensor(sensor)))
	}
	return nil
}

type QSensorChangesInterface struct {
	ptr unsafe.Pointer
}

type QSensorChangesInterface_ITF interface {
	QSensorChangesInterface_PTR() *QSensorChangesInterface
}

func (p *QSensorChangesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorChangesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QSensorChangesInterface) QSensorChangesInterface_PTR() *QSensorChangesInterface {
	return ptr
}

func (ptr *QSensorChangesInterface) SensorsChanged() {
	defer qt.Recovering("QSensorChangesInterface::sensorsChanged")

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

func (p *QSensorFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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
	for len(n.ObjectNameAbs()) < len("QSensorFilter_") {
		n.SetObjectNameAbs("QSensorFilter_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorFilter) QSensorFilter_PTR() *QSensorFilter {
	return ptr
}

func (ptr *QSensorFilter) Filter(reading QSensorReading_ITF) bool {
	defer qt.Recovering("QSensorFilter::filter")

	if ptr.Pointer() != nil {
		return C.QSensorFilter_Filter(ptr.Pointer(), PointerFromQSensorReading(reading)) != 0
	}
	return false
}

func (ptr *QSensorFilter) DestroyQSensorFilter() {
	defer qt.Recovering("QSensorFilter::~QSensorFilter")

	if ptr.Pointer() != nil {
		C.QSensorFilter_DestroyQSensorFilter(ptr.Pointer())
	}
}

func (ptr *QSensorFilter) ObjectNameAbs() string {
	defer qt.Recovering("QSensorFilter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorFilter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorFilter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSensorFilter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSensorFilter_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSensorGesture struct {
	core.QObject
}

type QSensorGesture_ITF interface {
	core.QObject_ITF
	QSensorGesture_PTR() *QSensorGesture
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
	for len(n.ObjectName()) < len("QSensorGesture_") {
		n.SetObjectName("QSensorGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return ptr
}

func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {
	defer qt.Recovering("QSensorGesture::QSensorGesture")

	return NewQSensorGestureFromPointer(C.QSensorGesture_NewQSensorGesture(C.CString(strings.Join(ids, "|")), core.PointerFromQObject(parent)))
}

func (ptr *QSensorGesture) GestureSignals() []string {
	defer qt.Recovering("QSensorGesture::gestureSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_GestureSignals(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	defer qt.Recovering("QSensorGesture::invalidIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_InvalidIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	defer qt.Recovering("QSensorGesture::isActive")

	if ptr.Pointer() != nil {
		return C.QSensorGesture_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	defer qt.Recovering("QSensorGesture::startDetection")

	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) StopDetection() {
	defer qt.Recovering("QSensorGesture::stopDetection")

	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	defer qt.Recovering("QSensorGesture::validIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_ValidIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	defer qt.Recovering("QSensorGesture::~QSensorGesture")

	if ptr.Pointer() != nil {
		C.QSensorGesture_DestroyQSensorGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorGestureTimerEvent
func callbackQSensorGestureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorGestureChildEvent
func callbackQSensorGestureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorGestureCustomEvent
func callbackQSensorGestureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGesture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSensorGestureManager struct {
	core.QObject
}

type QSensorGestureManager_ITF interface {
	core.QObject_ITF
	QSensorGestureManager_PTR() *QSensorGestureManager
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
	for len(n.ObjectName()) < len("QSensorGestureManager_") {
		n.SetObjectName("QSensorGestureManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGestureManager) QSensorGestureManager_PTR() *QSensorGestureManager {
	return ptr
}

func NewQSensorGestureManager(parent core.QObject_ITF) *QSensorGestureManager {
	defer qt.Recovering("QSensorGestureManager::QSensorGestureManager")

	return NewQSensorGestureManagerFromPointer(C.QSensorGestureManager_NewQSensorGestureManager(core.PointerFromQObject(parent)))
}

func (ptr *QSensorGestureManager) GestureIds() []string {
	defer qt.Recovering("QSensorGestureManager::gestureIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureManager_GestureIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) ConnectNewSensorGestureAvailable(f func()) {
	defer qt.Recovering("connect QSensorGestureManager::newSensorGestureAvailable")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNewSensorGestureAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newSensorGestureAvailable", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectNewSensorGestureAvailable() {
	defer qt.Recovering("disconnect QSensorGestureManager::newSensorGestureAvailable")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNewSensorGestureAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newSensorGestureAvailable")
	}
}

//export callbackQSensorGestureManagerNewSensorGestureAvailable
func callbackQSensorGestureManagerNewSensorGestureAvailable(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSensorGestureManager::newSensorGestureAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "newSensorGestureAvailable"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorGestureManager) NewSensorGestureAvailable() {
	defer qt.Recovering("QSensorGestureManager::newSensorGestureAvailable")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_NewSensorGestureAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorGestureManager) RecognizerSignals(gestureId string) []string {
	defer qt.Recovering("QSensorGestureManager::recognizerSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureManager_RecognizerSignals(ptr.Pointer(), C.CString(gestureId))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) RegisterSensorGestureRecognizer(recognizer QSensorGestureRecognizer_ITF) bool {
	defer qt.Recovering("QSensorGestureManager::registerSensorGestureRecognizer")

	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_RegisterSensorGestureRecognizer(ptr.Pointer(), PointerFromQSensorGestureRecognizer(recognizer)) != 0
	}
	return false
}

func QSensorGestureManager_SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	defer qt.Recovering("QSensorGestureManager::sensorGestureRecognizer")

	return NewQSensorGestureRecognizerFromPointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(C.CString(id)))
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManager() {
	defer qt.Recovering("QSensorGestureManager::~QSensorGestureManager")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DestroyQSensorGestureManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorGestureManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorGestureManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorGestureManagerTimerEvent
func callbackQSensorGestureManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGestureManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGestureManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorGestureManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorGestureManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorGestureManagerChildEvent
func callbackQSensorGestureManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGestureManager::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGestureManager::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorGestureManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorGestureManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorGestureManagerCustomEvent
func callbackQSensorGestureManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGestureManager::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGestureManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGestureManager::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSensorGesturePluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorGesturePluginInterface_ITF interface {
	QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface
}

func (p *QSensorGesturePluginInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorGesturePluginInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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
	for len(n.ObjectNameAbs()) < len("QSensorGesturePluginInterface_") {
		n.SetObjectNameAbs("QSensorGesturePluginInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGesturePluginInterface) QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface {
	return ptr
}

func (ptr *QSensorGesturePluginInterface) Name() string {
	defer qt.Recovering("QSensorGesturePluginInterface::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGesturePluginInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {
	defer qt.Recovering("QSensorGesturePluginInterface::supportedIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesturePluginInterface_SupportedIds(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {
	defer qt.Recovering("QSensorGesturePluginInterface::~QSensorGesturePluginInterface")

	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr.Pointer())
	}
}

func (ptr *QSensorGesturePluginInterface) ObjectNameAbs() string {
	defer qt.Recovering("QSensorGesturePluginInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGesturePluginInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGesturePluginInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSensorGesturePluginInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSensorGestureRecognizer struct {
	core.QObject
}

type QSensorGestureRecognizer_ITF interface {
	core.QObject_ITF
	QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer
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
	for len(n.ObjectName()) < len("QSensorGestureRecognizer_") {
		n.SetObjectName("QSensorGestureRecognizer_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGestureRecognizer) QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer {
	return ptr
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::createBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CreateBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {
	defer qt.Recovering("QSensorGestureRecognizer::gestureSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureRecognizer_GestureSignals(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureRecognizer) Id() string {
	defer qt.Recovering("QSensorGestureRecognizer::id")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGestureRecognizer_Id(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {
	defer qt.Recovering("QSensorGestureRecognizer::isActive")

	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StartBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::startBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StartBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) StopBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::stopBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StopBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {
	defer qt.Recovering("QSensorGestureRecognizer::~QSensorGestureRecognizer")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureRecognizer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorGestureRecognizerTimerEvent
func callbackQSensorGestureRecognizerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureRecognizer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorGestureRecognizerChildEvent
func callbackQSensorGestureRecognizerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureRecognizer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorGestureRecognizerCustomEvent
func callbackQSensorGestureRecognizerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureRecognizer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSensorManager struct {
	ptr unsafe.Pointer
}

type QSensorManager_ITF interface {
	QSensorManager_PTR() *QSensorManager
}

func (p *QSensorManager) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorManager) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QSensorManager) QSensorManager_PTR() *QSensorManager {
	return ptr
}

func QSensorManager_CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	defer qt.Recovering("QSensorManager::createBackend")

	return NewQSensorBackendFromPointer(C.QSensorManager_QSensorManager_CreateBackend(PointerFromQSensor(sensor)))
}

func QSensorManager_IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {
	defer qt.Recovering("QSensorManager::isBackendRegistered")

	return C.QSensorManager_QSensorManager_IsBackendRegistered(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier)) != 0
}

func QSensorManager_RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {
	defer qt.Recovering("QSensorManager::registerBackend")

	C.QSensorManager_QSensorManager_RegisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier), PointerFromQSensorBackendFactory(factory))
}

func QSensorManager_SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	defer qt.Recovering("QSensorManager::setDefaultBackend")

	C.QSensorManager_QSensorManager_SetDefaultBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func QSensorManager_UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	defer qt.Recovering("QSensorManager::unregisterBackend")

	C.QSensorManager_QSensorManager_UnregisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

type QSensorPluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorPluginInterface_ITF interface {
	QSensorPluginInterface_PTR() *QSensorPluginInterface
}

func (p *QSensorPluginInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorPluginInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QSensorPluginInterface) QSensorPluginInterface_PTR() *QSensorPluginInterface {
	return ptr
}

func (ptr *QSensorPluginInterface) RegisterSensors() {
	defer qt.Recovering("QSensorPluginInterface::registerSensors")

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

func PointerFromQSensorReading(ptr QSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func NewQSensorReadingFromPointer(ptr unsafe.Pointer) *QSensorReading {
	var n = new(QSensorReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorReading_") {
		n.SetObjectName("QSensorReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorReading) QSensorReading_PTR() *QSensorReading {
	return ptr
}

func (ptr *QSensorReading) Value(index int) *core.QVariant {
	defer qt.Recovering("QSensorReading::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSensorReading_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSensorReading) ValueCount() int {
	defer qt.Recovering("QSensorReading::valueCount")

	if ptr.Pointer() != nil {
		return int(C.QSensorReading_ValueCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensorReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorReadingTimerEvent
func callbackQSensorReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorReadingChildEvent
func callbackQSensorReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorReading::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorReading::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorReadingCustomEvent
func callbackQSensorReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorReading::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorReading::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QTapFilter struct {
	QSensorFilter
}

type QTapFilter_ITF interface {
	QSensorFilter_ITF
	QTapFilter_PTR() *QTapFilter
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

func (ptr *QTapFilter) QTapFilter_PTR() *QTapFilter {
	return ptr
}

func (ptr *QTapFilter) Filter(reading QTapReading_ITF) bool {
	defer qt.Recovering("QTapFilter::filter")

	if ptr.Pointer() != nil {
		return C.QTapFilter_Filter(ptr.Pointer(), PointerFromQTapReading(reading)) != 0
	}
	return false
}

type QTapReading struct {
	QSensorReading
}

type QTapReading_ITF interface {
	QSensorReading_ITF
	QTapReading_PTR() *QTapReading
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
	for len(n.ObjectName()) < len("QTapReading_") {
		n.SetObjectName("QTapReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapReading) QTapReading_PTR() *QTapReading {
	return ptr
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

func (ptr *QTapReading) IsDoubleTap() bool {
	defer qt.Recovering("QTapReading::isDoubleTap")

	if ptr.Pointer() != nil {
		return C.QTapReading_IsDoubleTap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {
	defer qt.Recovering("QTapReading::tapDirection")

	if ptr.Pointer() != nil {
		return QTapReading__TapDirection(C.QTapReading_TapDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {
	defer qt.Recovering("QTapReading::setDoubleTap")

	if ptr.Pointer() != nil {
		C.QTapReading_SetDoubleTap(ptr.Pointer(), C.int(qt.GoBoolToInt(doubleTap)))
	}
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {
	defer qt.Recovering("QTapReading::setTapDirection")

	if ptr.Pointer() != nil {
		C.QTapReading_SetTapDirection(ptr.Pointer(), C.int(tapDirection))
	}
}

func (ptr *QTapReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTapReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTapReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTapReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTapReadingTimerEvent
func callbackQTapReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTapReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTapReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTapReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTapReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTapReadingChildEvent
func callbackQTapReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTapReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTapReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTapReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTapReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTapReadingCustomEvent
func callbackQTapReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTapReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTapReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTapReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTapReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QTapSensor struct {
	QSensor
}

type QTapSensor_ITF interface {
	QSensor_ITF
	QTapSensor_PTR() *QTapSensor
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
	for len(n.ObjectName()) < len("QTapSensor_") {
		n.SetObjectName("QTapSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapSensor) QTapSensor_PTR() *QTapSensor {
	return ptr
}

func (ptr *QTapSensor) Reading() *QTapReading {
	defer qt.Recovering("QTapSensor::reading")

	if ptr.Pointer() != nil {
		return NewQTapReadingFromPointer(C.QTapSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {
	defer qt.Recovering("QTapSensor::returnDoubleTapEvents")

	if ptr.Pointer() != nil {
		return C.QTapSensor_ReturnDoubleTapEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {
	defer qt.Recovering("QTapSensor::setReturnDoubleTapEvents")

	if ptr.Pointer() != nil {
		C.QTapSensor_SetReturnDoubleTapEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(returnDoubleTapEvents)))
	}
}

func NewQTapSensor(parent core.QObject_ITF) *QTapSensor {
	defer qt.Recovering("QTapSensor::QTapSensor")

	return NewQTapSensorFromPointer(C.QTapSensor_NewQTapSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {
	defer qt.Recovering("connect QTapSensor::returnDoubleTapEventsChanged")

	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged", f)
	}
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {
	defer qt.Recovering("disconnect QTapSensor::returnDoubleTapEventsChanged")

	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged")
	}
}

//export callbackQTapSensorReturnDoubleTapEventsChanged
func callbackQTapSensorReturnDoubleTapEventsChanged(ptr unsafe.Pointer, ptrName *C.char, returnDoubleTapEvents C.int) {
	defer qt.Recovering("callback QTapSensor::returnDoubleTapEventsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "returnDoubleTapEventsChanged"); signal != nil {
		signal.(func(bool))(int(returnDoubleTapEvents) != 0)
	}

}

func (ptr *QTapSensor) ReturnDoubleTapEventsChanged(returnDoubleTapEvents bool) {
	defer qt.Recovering("QTapSensor::returnDoubleTapEventsChanged")

	if ptr.Pointer() != nil {
		C.QTapSensor_ReturnDoubleTapEventsChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(returnDoubleTapEvents)))
	}
}

func (ptr *QTapSensor) DestroyQTapSensor() {
	defer qt.Recovering("QTapSensor::~QTapSensor")

	if ptr.Pointer() != nil {
		C.QTapSensor_DestroyQTapSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTapSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTapSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTapSensorTimerEvent
func callbackQTapSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTapSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTapSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTapSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTapSensorChildEvent
func callbackQTapSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTapSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTapSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTapSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTapSensorCustomEvent
func callbackQTapSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTapSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTapSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTapSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTapSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QTiltFilter struct {
	QSensorFilter
}

type QTiltFilter_ITF interface {
	QSensorFilter_ITF
	QTiltFilter_PTR() *QTiltFilter
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

func (ptr *QTiltFilter) QTiltFilter_PTR() *QTiltFilter {
	return ptr
}

func (ptr *QTiltFilter) Filter(reading QTiltReading_ITF) bool {
	defer qt.Recovering("QTiltFilter::filter")

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

func PointerFromQTiltReading(ptr QTiltReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReading_PTR().Pointer()
	}
	return nil
}

func NewQTiltReadingFromPointer(ptr unsafe.Pointer) *QTiltReading {
	var n = new(QTiltReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTiltReading_") {
		n.SetObjectName("QTiltReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QTiltReading) QTiltReading_PTR() *QTiltReading {
	return ptr
}

func (ptr *QTiltReading) XRotation() float64 {
	defer qt.Recovering("QTiltReading::xRotation")

	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_XRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) YRotation() float64 {
	defer qt.Recovering("QTiltReading::yRotation")

	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_YRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) SetXRotation(x float64) {
	defer qt.Recovering("QTiltReading::setXRotation")

	if ptr.Pointer() != nil {
		C.QTiltReading_SetXRotation(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QTiltReading) SetYRotation(y float64) {
	defer qt.Recovering("QTiltReading::setYRotation")

	if ptr.Pointer() != nil {
		C.QTiltReading_SetYRotation(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QTiltReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTiltReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTiltReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTiltReadingTimerEvent
func callbackQTiltReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTiltReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTiltReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTiltReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTiltReadingChildEvent
func callbackQTiltReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTiltReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTiltReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTiltReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTiltReadingCustomEvent
func callbackQTiltReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTiltReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTiltReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
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
	for len(n.ObjectName()) < len("QTiltSensor_") {
		n.SetObjectName("QTiltSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return ptr
}

func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {
	defer qt.Recovering("QTiltSensor::QTiltSensor")

	return NewQTiltSensorFromPointer(C.QTiltSensor_NewQTiltSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	defer qt.Recovering("QTiltSensor::reading")

	if ptr.Pointer() != nil {
		return NewQTiltReadingFromPointer(C.QTiltSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	defer qt.Recovering("QTiltSensor::~QTiltSensor")

	if ptr.Pointer() != nil {
		C.QTiltSensor_DestroyQTiltSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) Calibrate() {
	defer qt.Recovering("QTiltSensor::calibrate")

	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(ptr.Pointer())
	}
}

func (ptr *QTiltSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTiltSensorTimerEvent
func callbackQTiltSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTiltSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTiltSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTiltSensorChildEvent
func callbackQTiltSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTiltSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTiltSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTiltSensorCustomEvent
func callbackQTiltSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTiltSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
