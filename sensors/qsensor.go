package sensors

//#include "qsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensor struct {
	core.QObject
}

type QSensorITF interface {
	core.QObjectITF
	QSensorPTR() *QSensor
}

func PointerFromQSensor(ptr QSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorPTR().Pointer()
	}
	return nil
}

func QSensorFromPointer(ptr unsafe.Pointer) *QSensor {
	var n = new(QSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensor) QSensorPTR() *QSensor {
	return ptr
}

//QSensor::AxesOrientationMode
type QSensor__AxesOrientationMode int

var (
	QSensor__FixedOrientation     = QSensor__AxesOrientationMode(0)
	QSensor__AutomaticOrientation = QSensor__AxesOrientationMode(1)
	QSensor__UserOrientation      = QSensor__AxesOrientationMode(2)
)

//QSensor::Feature
type QSensor__Feature int

var (
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
	if ptr.Pointer() != nil {
		return QSensor__AxesOrientationMode(C.QSensor_AxesOrientationMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_BufferSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) CurrentOrientation() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_CurrentOrientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) DataRate() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_DataRate(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSensor_Description(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSensor) EfficientBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_EfficientBufferSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Error() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsAlwaysOn() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsAlwaysOn(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsBusy() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsBusy(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsConnectedToBackend() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsConnectedToBackend(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) MaxBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_MaxBufferSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) OutputRange() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_OutputRange(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		return QSensorReadingFromPointer(unsafe.Pointer(C.QSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSensor) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetActive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAlwaysOn(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(alwaysOn)))
	}
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAxesOrientationMode(C.QtObjectPtr(ptr.Pointer()), C.int(axesOrientationMode))
	}
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetBufferSize(C.QtObjectPtr(ptr.Pointer()), C.int(bufferSize))
	}
}

func (ptr *QSensor) SetDataRate(rate int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetDataRate(C.QtObjectPtr(ptr.Pointer()), C.int(rate))
	}
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QSensor_SetIdentifier(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(identifier)))
	}
}

func (ptr *QSensor) SetOutputRange(index int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetOutputRange(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetUserOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(userOrientation))
	}
}

func (ptr *QSensor) SkipDuplicates() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_SkipDuplicates(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) UserOrientation() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_UserOrientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQSensor(ty core.QByteArrayITF, parent core.QObjectITF) *QSensor {
	return QSensorFromPointer(unsafe.Pointer(C.QSensor_NewQSensor(C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSensor) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QSensor) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQSensorActiveChanged
func callbackQSensorActiveChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QSensor) AddFilter(filter QSensorFilterITF) {
	if ptr.Pointer() != nil {
		C.QSensor_AddFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSensorFilter(filter)))
	}
}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAlwaysOnChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "alwaysOnChanged", f)
	}
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAlwaysOnChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "alwaysOnChanged")
	}
}

//export callbackQSensorAlwaysOnChanged
func callbackQSensorAlwaysOnChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "alwaysOnChanged").(func())()
}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAvailableSensorsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "availableSensorsChanged", f)
	}
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAvailableSensorsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "availableSensorsChanged")
	}
}

//export callbackQSensorAvailableSensorsChanged
func callbackQSensorAvailableSensorsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "availableSensorsChanged").(func())()
}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAxesOrientationModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "axesOrientationModeChanged", f)
	}
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAxesOrientationModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "axesOrientationModeChanged")
	}
}

//export callbackQSensorAxesOrientationModeChanged
func callbackQSensorAxesOrientationModeChanged(ptrName *C.char, axesOrientationMode C.int) {
	qt.GetSignal(C.GoString(ptrName), "axesOrientationModeChanged").(func(QSensor__AxesOrientationMode))(QSensor__AxesOrientationMode(axesOrientationMode))
}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectBufferSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBufferSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferSizeChanged")
	}
}

//export callbackQSensorBufferSizeChanged
func callbackQSensorBufferSizeChanged(ptrName *C.char, bufferSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferSizeChanged").(func(int))(int(bufferSize))
}

func (ptr *QSensor) ConnectBusyChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectBusyChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "busyChanged", f)
	}
}

func (ptr *QSensor) DisconnectBusyChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBusyChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "busyChanged")
	}
}

//export callbackQSensorBusyChanged
func callbackQSensorBusyChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "busyChanged").(func())()
}

func (ptr *QSensor) ConnectToBackend() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_ConnectToBackend(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectCurrentOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectCurrentOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentOrientationChanged")
	}
}

//export callbackQSensorCurrentOrientationChanged
func callbackQSensorCurrentOrientationChanged(ptrName *C.char, currentOrientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentOrientationChanged").(func(int))(int(currentOrientation))
}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectDataRateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "dataRateChanged", f)
	}
}

func (ptr *QSensor) DisconnectDataRateChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectDataRateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "dataRateChanged")
	}
}

//export callbackQSensorDataRateChanged
func callbackQSensorDataRateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "dataRateChanged").(func())()
}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectEfficientBufferSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectEfficientBufferSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged")
	}
}

//export callbackQSensorEfficientBufferSizeChanged
func callbackQSensorEfficientBufferSizeChanged(ptrName *C.char, efficientBufferSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "efficientBufferSizeChanged").(func(int))(int(efficientBufferSize))
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsFeatureSupported(C.QtObjectPtr(ptr.Pointer()), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectMaxBufferSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "maxBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectMaxBufferSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "maxBufferSizeChanged")
	}
}

//export callbackQSensorMaxBufferSizeChanged
func callbackQSensorMaxBufferSizeChanged(ptrName *C.char, maxBufferSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "maxBufferSizeChanged").(func(int))(int(maxBufferSize))
}

func (ptr *QSensor) ConnectReadingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectReadingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readingChanged", f)
	}
}

func (ptr *QSensor) DisconnectReadingChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectReadingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readingChanged")
	}
}

//export callbackQSensorReadingChanged
func callbackQSensorReadingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readingChanged").(func())()
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilterITF) {
	if ptr.Pointer() != nil {
		C.QSensor_RemoveFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSensorFilter(filter)))
	}
}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectSensorError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sensorError", f)
	}
}

func (ptr *QSensor) DisconnectSensorError() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSensorError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sensorError")
	}
}

//export callbackQSensorSensorError
func callbackQSensorSensorError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "sensorError").(func(int))(int(error))
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetCurrentOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(currentOrientation))
	}
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetEfficientBufferSize(C.QtObjectPtr(ptr.Pointer()), C.int(efficientBufferSize))
	}
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetMaxBufferSize(C.QtObjectPtr(ptr.Pointer()), C.int(maxBufferSize))
	}
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetSkipDuplicates(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(skipDuplicates)))
	}
}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectSkipDuplicatesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "skipDuplicatesChanged", f)
	}
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSkipDuplicatesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "skipDuplicatesChanged")
	}
}

//export callbackQSensorSkipDuplicatesChanged
func callbackQSensorSkipDuplicatesChanged(ptrName *C.char, skipDuplicates C.int) {
	qt.GetSignal(C.GoString(ptrName), "skipDuplicatesChanged").(func(bool))(int(skipDuplicates) != 0)
}

func (ptr *QSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_Start(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QSensor_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectUserOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "userOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectUserOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "userOrientationChanged")
	}
}

//export callbackQSensorUserOrientationChanged
func callbackQSensorUserOrientationChanged(ptrName *C.char, userOrientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "userOrientationChanged").(func(int))(int(userOrientation))
}

func (ptr *QSensor) DestroyQSensor() {
	if ptr.Pointer() != nil {
		C.QSensor_DestroyQSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
