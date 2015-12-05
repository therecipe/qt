package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

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
		n.SetObjectName("QSensor_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::axesOrientationMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QSensor__AxesOrientationMode(C.QSensor_AxesOrientationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) BufferSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::bufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) CurrentOrientation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::currentOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_CurrentOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) DataRate() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::dataRate")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_DataRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Description() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::description")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensor_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensor) EfficientBufferSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::efficientBufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_EfficientBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Error() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::error")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Identifier() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::identifier")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSensor_Identifier(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsAlwaysOn() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::isAlwaysOn")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_IsAlwaysOn(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsBusy() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::isBusy")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_IsBusy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) IsConnectedToBackend() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::isConnectedToBackend")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_IsConnectedToBackend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) MaxBufferSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::maxBufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_MaxBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) OutputRange() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::outputRange")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_OutputRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Reading() *QSensorReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSensorReadingFromPointer(C.QSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) SetActive(active bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setActive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setAlwaysOn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetAlwaysOn(ptr.Pointer(), C.int(qt.GoBoolToInt(alwaysOn)))
	}
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setAxesOrientationMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetAxesOrientationMode(ptr.Pointer(), C.int(axesOrientationMode))
	}
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setBufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetBufferSize(ptr.Pointer(), C.int(bufferSize))
	}
}

func (ptr *QSensor) SetDataRate(rate int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setDataRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetDataRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setIdentifier")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetIdentifier(ptr.Pointer(), core.PointerFromQByteArray(identifier))
	}
}

func (ptr *QSensor) SetOutputRange(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setOutputRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetOutputRange(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setUserOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetUserOrientation(ptr.Pointer(), C.int(userOrientation))
	}
}

func (ptr *QSensor) SkipDuplicates() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::skipDuplicates")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_SkipDuplicates(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) Type() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::type")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSensor_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) UserOrientation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::userOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSensor_UserOrientation(ptr.Pointer()))
	}
	return 0
}

func NewQSensor(ty core.QByteArray_ITF, parent core.QObject_ITF) *QSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::QSensor")
		}
	}()

	return NewQSensorFromPointer(C.QSensor_NewQSensor(core.PointerFromQByteArray(ty), core.PointerFromQObject(parent)))
}

func (ptr *QSensor) ConnectActiveChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QSensor) DisconnectActiveChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQSensorActiveChanged
func callbackQSensorActiveChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::activeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QSensor) AddFilter(filter QSensorFilter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::addFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_AddFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::alwaysOnChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectAlwaysOnChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alwaysOnChanged", f)
	}
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::alwaysOnChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAlwaysOnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alwaysOnChanged")
	}
}

//export callbackQSensorAlwaysOnChanged
func callbackQSensorAlwaysOnChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::alwaysOnChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "alwaysOnChanged").(func())()
}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::availableSensorsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectAvailableSensorsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableSensorsChanged", f)
	}
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::availableSensorsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAvailableSensorsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableSensorsChanged")
	}
}

//export callbackQSensorAvailableSensorsChanged
func callbackQSensorAvailableSensorsChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::availableSensorsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "availableSensorsChanged").(func())()
}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::axesOrientationModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectAxesOrientationModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "axesOrientationModeChanged", f)
	}
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::axesOrientationModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAxesOrientationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "axesOrientationModeChanged")
	}
}

//export callbackQSensorAxesOrientationModeChanged
func callbackQSensorAxesOrientationModeChanged(ptrName *C.char, axesOrientationMode C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::axesOrientationModeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "axesOrientationModeChanged").(func(QSensor__AxesOrientationMode))(QSensor__AxesOrientationMode(axesOrientationMode))
}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::bufferSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::bufferSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferSizeChanged")
	}
}

//export callbackQSensorBufferSizeChanged
func callbackQSensorBufferSizeChanged(ptrName *C.char, bufferSize C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::bufferSizeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bufferSizeChanged").(func(int))(int(bufferSize))
}

func (ptr *QSensor) ConnectBusyChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::busyChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectBusyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "busyChanged", f)
	}
}

func (ptr *QSensor) DisconnectBusyChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::busyChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBusyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "busyChanged")
	}
}

//export callbackQSensorBusyChanged
func callbackQSensorBusyChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::busyChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "busyChanged").(func())()
}

func (ptr *QSensor) ConnectToBackend() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::connectToBackend")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_ConnectToBackend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::currentOrientationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectCurrentOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::currentOrientationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectCurrentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentOrientationChanged")
	}
}

//export callbackQSensorCurrentOrientationChanged
func callbackQSensorCurrentOrientationChanged(ptrName *C.char, currentOrientation C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::currentOrientationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentOrientationChanged").(func(int))(int(currentOrientation))
}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::dataRateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectDataRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataRateChanged", f)
	}
}

func (ptr *QSensor) DisconnectDataRateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::dataRateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectDataRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataRateChanged")
	}
}

//export callbackQSensorDataRateChanged
func callbackQSensorDataRateChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::dataRateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "dataRateChanged").(func())()
}

func QSensor_DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::defaultSensorForType")
		}
	}()

	return core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::efficientBufferSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::efficientBufferSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged")
	}
}

//export callbackQSensorEfficientBufferSizeChanged
func callbackQSensorEfficientBufferSizeChanged(ptrName *C.char, efficientBufferSize C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::efficientBufferSizeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "efficientBufferSizeChanged").(func(int))(int(efficientBufferSize))
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::isFeatureSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_IsFeatureSupported(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::maxBufferSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectMaxBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maxBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::maxBufferSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectMaxBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maxBufferSizeChanged")
	}
}

//export callbackQSensorMaxBufferSizeChanged
func callbackQSensorMaxBufferSizeChanged(ptrName *C.char, maxBufferSize C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::maxBufferSizeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "maxBufferSizeChanged").(func(int))(int(maxBufferSize))
}

func (ptr *QSensor) ConnectReadingChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::readingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectReadingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readingChanged", f)
	}
}

func (ptr *QSensor) DisconnectReadingChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::readingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectReadingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readingChanged")
	}
}

//export callbackQSensorReadingChanged
func callbackQSensorReadingChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::readingChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "readingChanged").(func())()
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::removeFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_RemoveFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::sensorError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectSensorError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sensorError", f)
	}
}

func (ptr *QSensor) DisconnectSensorError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::sensorError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSensorError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sensorError")
	}
}

//export callbackQSensorSensorError
func callbackQSensorSensorError(ptrName *C.char, error C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::sensorError")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sensorError").(func(int))(int(error))
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setCurrentOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetCurrentOrientation(ptr.Pointer(), C.int(currentOrientation))
	}
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setEfficientBufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetEfficientBufferSize(ptr.Pointer(), C.int(efficientBufferSize))
	}
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setMaxBufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetMaxBufferSize(ptr.Pointer(), C.int(maxBufferSize))
	}
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::setSkipDuplicates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_SetSkipDuplicates(ptr.Pointer(), C.int(qt.GoBoolToInt(skipDuplicates)))
	}
}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::skipDuplicatesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectSkipDuplicatesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "skipDuplicatesChanged", f)
	}
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::skipDuplicatesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSkipDuplicatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "skipDuplicatesChanged")
	}
}

//export callbackQSensorSkipDuplicatesChanged
func callbackQSensorSkipDuplicatesChanged(ptrName *C.char, skipDuplicates C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::skipDuplicatesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "skipDuplicatesChanged").(func(bool))(int(skipDuplicates) != 0)
}

func (ptr *QSensor) Start() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::start")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::userOrientationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_ConnectUserOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "userOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::userOrientationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DisconnectUserOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "userOrientationChanged")
	}
}

//export callbackQSensorUserOrientationChanged
func callbackQSensorUserOrientationChanged(ptrName *C.char, userOrientation C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::userOrientationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "userOrientationChanged").(func(int))(int(userOrientation))
}

func (ptr *QSensor) DestroyQSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensor::~QSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensor_DestroyQSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
