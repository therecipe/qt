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
	if len(n.ObjectName()) == 0 {
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
	if ptr.Pointer() != nil {
		return QSensor__AxesOrientationMode(C.QSensor_AxesOrientationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) CurrentOrientation() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_CurrentOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) DataRate() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_DataRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSensor_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensor) EfficientBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_EfficientBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Error() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Identifier() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSensor_Identifier(ptr.Pointer()))
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
		return int(C.QSensor_MaxBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) OutputRange() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_OutputRange(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensor) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		return NewQSensorReadingFromPointer(C.QSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAlwaysOn(ptr.Pointer(), C.int(qt.GoBoolToInt(alwaysOn)))
	}
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAxesOrientationMode(ptr.Pointer(), C.int(axesOrientationMode))
	}
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetBufferSize(ptr.Pointer(), C.int(bufferSize))
	}
}

func (ptr *QSensor) SetDataRate(rate int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetDataRate(ptr.Pointer(), C.int(rate))
	}
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_SetIdentifier(ptr.Pointer(), core.PointerFromQByteArray(identifier))
	}
}

func (ptr *QSensor) SetOutputRange(index int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetOutputRange(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetUserOrientation(ptr.Pointer(), C.int(userOrientation))
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
		return core.NewQByteArrayFromPointer(C.QSensor_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensor) UserOrientation() int {
	if ptr.Pointer() != nil {
		return int(C.QSensor_UserOrientation(ptr.Pointer()))
	}
	return 0
}

func NewQSensor(ty core.QByteArray_ITF, parent core.QObject_ITF) *QSensor {
	return NewQSensorFromPointer(C.QSensor_NewQSensor(core.PointerFromQByteArray(ty), core.PointerFromQObject(parent)))
}

func (ptr *QSensor) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QSensor) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQSensorActiveChanged
func callbackQSensorActiveChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QSensor) AddFilter(filter QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_AddFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAlwaysOnChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alwaysOnChanged", f)
	}
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAlwaysOnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alwaysOnChanged")
	}
}

//export callbackQSensorAlwaysOnChanged
func callbackQSensorAlwaysOnChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "alwaysOnChanged").(func())()
}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAvailableSensorsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableSensorsChanged", f)
	}
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAvailableSensorsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableSensorsChanged")
	}
}

//export callbackQSensorAvailableSensorsChanged
func callbackQSensorAvailableSensorsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "availableSensorsChanged").(func())()
}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectAxesOrientationModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "axesOrientationModeChanged", f)
	}
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAxesOrientationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "axesOrientationModeChanged")
	}
}

//export callbackQSensorAxesOrientationModeChanged
func callbackQSensorAxesOrientationModeChanged(ptrName *C.char, axesOrientationMode C.int) {
	qt.GetSignal(C.GoString(ptrName), "axesOrientationModeChanged").(func(QSensor__AxesOrientationMode))(QSensor__AxesOrientationMode(axesOrientationMode))
}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferSizeChanged")
	}
}

//export callbackQSensorBufferSizeChanged
func callbackQSensorBufferSizeChanged(ptrName *C.char, bufferSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferSizeChanged").(func(int))(int(bufferSize))
}

func (ptr *QSensor) ConnectBusyChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectBusyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "busyChanged", f)
	}
}

func (ptr *QSensor) DisconnectBusyChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBusyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "busyChanged")
	}
}

//export callbackQSensorBusyChanged
func callbackQSensorBusyChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "busyChanged").(func())()
}

func (ptr *QSensor) ConnectToBackend() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_ConnectToBackend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectCurrentOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectCurrentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentOrientationChanged")
	}
}

//export callbackQSensorCurrentOrientationChanged
func callbackQSensorCurrentOrientationChanged(ptrName *C.char, currentOrientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentOrientationChanged").(func(int))(int(currentOrientation))
}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectDataRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataRateChanged", f)
	}
}

func (ptr *QSensor) DisconnectDataRateChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectDataRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataRateChanged")
	}
}

//export callbackQSensorDataRateChanged
func callbackQSensorDataRateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "dataRateChanged").(func())()
}

func QSensor_DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	return core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "efficientBufferSizeChanged")
	}
}

//export callbackQSensorEfficientBufferSizeChanged
func callbackQSensorEfficientBufferSizeChanged(ptrName *C.char, efficientBufferSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "efficientBufferSizeChanged").(func(int))(int(efficientBufferSize))
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return C.QSensor_IsFeatureSupported(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectMaxBufferSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maxBufferSizeChanged", f)
	}
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectMaxBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maxBufferSizeChanged")
	}
}

//export callbackQSensorMaxBufferSizeChanged
func callbackQSensorMaxBufferSizeChanged(ptrName *C.char, maxBufferSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "maxBufferSizeChanged").(func(int))(int(maxBufferSize))
}

func (ptr *QSensor) ConnectReadingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectReadingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readingChanged", f)
	}
}

func (ptr *QSensor) DisconnectReadingChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectReadingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readingChanged")
	}
}

//export callbackQSensorReadingChanged
func callbackQSensorReadingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readingChanged").(func())()
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_RemoveFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectSensorError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sensorError", f)
	}
}

func (ptr *QSensor) DisconnectSensorError() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSensorError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sensorError")
	}
}

//export callbackQSensorSensorError
func callbackQSensorSensorError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "sensorError").(func(int))(int(error))
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetCurrentOrientation(ptr.Pointer(), C.int(currentOrientation))
	}
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetEfficientBufferSize(ptr.Pointer(), C.int(efficientBufferSize))
	}
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetMaxBufferSize(ptr.Pointer(), C.int(maxBufferSize))
	}
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetSkipDuplicates(ptr.Pointer(), C.int(qt.GoBoolToInt(skipDuplicates)))
	}
}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectSkipDuplicatesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "skipDuplicatesChanged", f)
	}
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSkipDuplicatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "skipDuplicatesChanged")
	}
}

//export callbackQSensorSkipDuplicatesChanged
func callbackQSensorSkipDuplicatesChanged(ptrName *C.char, skipDuplicates C.int) {
	qt.GetSignal(C.GoString(ptrName), "skipDuplicatesChanged").(func(bool))(int(skipDuplicates) != 0)
}

func (ptr *QSensor) Start() bool {
	if ptr.Pointer() != nil {
		return C.QSensor_Start(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectUserOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "userOrientationChanged", f)
	}
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectUserOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "userOrientationChanged")
	}
}

//export callbackQSensorUserOrientationChanged
func callbackQSensorUserOrientationChanged(ptrName *C.char, userOrientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "userOrientationChanged").(func(int))(int(userOrientation))
}

func (ptr *QSensor) DestroyQSensor() {
	if ptr.Pointer() != nil {
		C.QSensor_DestroyQSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
