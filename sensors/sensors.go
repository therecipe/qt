// +build !minimal

package sensors

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"strings"
	"unsafe"
)

type AndroidAccelerometer struct {
	internal.Internal
}

type AndroidAccelerometer_ITF interface {
	AndroidAccelerometer_PTR() *AndroidAccelerometer
}

func (ptr *AndroidAccelerometer) AndroidAccelerometer_PTR() *AndroidAccelerometer {
	return ptr
}

func (ptr *AndroidAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidAccelerometer(ptr AndroidAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *AndroidAccelerometer) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidAccelerometerFromPointer(ptr unsafe.Pointer) (n *AndroidAccelerometer) {
	n = new(AndroidAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidAccelerometer")
	return
}

func (ptr *AndroidAccelerometer) DestroyAndroidAccelerometer() {
}

type AndroidCompass struct {
	ThreadSafeSensorBackend
}

type AndroidCompass_ITF interface {
	ThreadSafeSensorBackend_ITF
	AndroidCompass_PTR() *AndroidCompass
}

func (ptr *AndroidCompass) AndroidCompass_PTR() *AndroidCompass {
	return ptr
}

func (ptr *AndroidCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ThreadSafeSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *AndroidCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ThreadSafeSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromAndroidCompass(ptr AndroidCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidCompass_PTR().Pointer()
	}
	return nil
}

func (n *AndroidCompass) InitFromInternal(ptr uintptr, name string) {
	n.ThreadSafeSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *AndroidCompass) ClassNameInternalF() string {
	return n.ThreadSafeSensorBackend_PTR().ClassNameInternalF()
}

func NewAndroidCompassFromPointer(ptr unsafe.Pointer) (n *AndroidCompass) {
	n = new(AndroidCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidCompass")
	return
}

func (ptr *AndroidCompass) DestroyAndroidCompass() {
}

type AndroidGyroscope struct {
	internal.Internal
}

type AndroidGyroscope_ITF interface {
	AndroidGyroscope_PTR() *AndroidGyroscope
}

func (ptr *AndroidGyroscope) AndroidGyroscope_PTR() *AndroidGyroscope {
	return ptr
}

func (ptr *AndroidGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidGyroscope(ptr AndroidGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *AndroidGyroscope) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidGyroscopeFromPointer(ptr unsafe.Pointer) (n *AndroidGyroscope) {
	n = new(AndroidGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidGyroscope")
	return
}

func (ptr *AndroidGyroscope) DestroyAndroidGyroscope() {
}

type AndroidLight struct {
	internal.Internal
}

type AndroidLight_ITF interface {
	AndroidLight_PTR() *AndroidLight
}

func (ptr *AndroidLight) AndroidLight_PTR() *AndroidLight {
	return ptr
}

func (ptr *AndroidLight) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidLight) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidLight(ptr AndroidLight_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidLight_PTR().Pointer()
	}
	return nil
}

func (n *AndroidLight) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidLightFromPointer(ptr unsafe.Pointer) (n *AndroidLight) {
	n = new(AndroidLight)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidLight")
	return
}

func (ptr *AndroidLight) DestroyAndroidLight() {
}

type AndroidMagnetometer struct {
	internal.Internal
}

type AndroidMagnetometer_ITF interface {
	AndroidMagnetometer_PTR() *AndroidMagnetometer
}

func (ptr *AndroidMagnetometer) AndroidMagnetometer_PTR() *AndroidMagnetometer {
	return ptr
}

func (ptr *AndroidMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidMagnetometer(ptr AndroidMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *AndroidMagnetometer) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidMagnetometerFromPointer(ptr unsafe.Pointer) (n *AndroidMagnetometer) {
	n = new(AndroidMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidMagnetometer")
	return
}

func (ptr *AndroidMagnetometer) DestroyAndroidMagnetometer() {
}

type AndroidPressure struct {
	internal.Internal
}

type AndroidPressure_ITF interface {
	AndroidPressure_PTR() *AndroidPressure
}

func (ptr *AndroidPressure) AndroidPressure_PTR() *AndroidPressure {
	return ptr
}

func (ptr *AndroidPressure) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidPressure) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidPressure(ptr AndroidPressure_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidPressure_PTR().Pointer()
	}
	return nil
}

func (n *AndroidPressure) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidPressureFromPointer(ptr unsafe.Pointer) (n *AndroidPressure) {
	n = new(AndroidPressure)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidPressure")
	return
}

func (ptr *AndroidPressure) DestroyAndroidPressure() {
}

type AndroidProximity struct {
	internal.Internal
}

type AndroidProximity_ITF interface {
	AndroidProximity_PTR() *AndroidProximity
}

func (ptr *AndroidProximity) AndroidProximity_PTR() *AndroidProximity {
	return ptr
}

func (ptr *AndroidProximity) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidProximity) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidProximity(ptr AndroidProximity_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidProximity_PTR().Pointer()
	}
	return nil
}

func (n *AndroidProximity) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidProximityFromPointer(ptr unsafe.Pointer) (n *AndroidProximity) {
	n = new(AndroidProximity)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidProximity")
	return
}

func (ptr *AndroidProximity) DestroyAndroidProximity() {
}

type AndroidRotation struct {
	internal.Internal
}

type AndroidRotation_ITF interface {
	AndroidRotation_PTR() *AndroidRotation
}

func (ptr *AndroidRotation) AndroidRotation_PTR() *AndroidRotation {
	return ptr
}

func (ptr *AndroidRotation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidRotation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidRotation(ptr AndroidRotation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidRotation_PTR().Pointer()
	}
	return nil
}

func (n *AndroidRotation) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidRotationFromPointer(ptr unsafe.Pointer) (n *AndroidRotation) {
	n = new(AndroidRotation)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidRotation")
	return
}

func (ptr *AndroidRotation) DestroyAndroidRotation() {
}

type AndroidTemperature struct {
	internal.Internal
}

type AndroidTemperature_ITF interface {
	AndroidTemperature_PTR() *AndroidTemperature
}

func (ptr *AndroidTemperature) AndroidTemperature_PTR() *AndroidTemperature {
	return ptr
}

func (ptr *AndroidTemperature) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AndroidTemperature) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAndroidTemperature(ptr AndroidTemperature_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AndroidTemperature_PTR().Pointer()
	}
	return nil
}

func (n *AndroidTemperature) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAndroidTemperatureFromPointer(ptr unsafe.Pointer) (n *AndroidTemperature) {
	n = new(AndroidTemperature)
	n.InitFromInternal(uintptr(ptr), "sensors.AndroidTemperature")
	return
}

func (ptr *AndroidTemperature) DestroyAndroidTemperature() {
}

type FunctionEvent struct {
	core.QEvent
}

type FunctionEvent_ITF interface {
	core.QEvent_ITF
	FunctionEvent_PTR() *FunctionEvent
}

func (ptr *FunctionEvent) FunctionEvent_PTR() *FunctionEvent {
	return ptr
}

func (ptr *FunctionEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *FunctionEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromFunctionEvent(ptr FunctionEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.FunctionEvent_PTR().Pointer()
	}
	return nil
}

func (n *FunctionEvent) InitFromInternal(ptr uintptr, name string) {
	n.QEvent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *FunctionEvent) ClassNameInternalF() string {
	return n.QEvent_PTR().ClassNameInternalF()
}

func NewFunctionEventFromPointer(ptr unsafe.Pointer) (n *FunctionEvent) {
	n = new(FunctionEvent)
	n.InitFromInternal(uintptr(ptr), "sensors.FunctionEvent")
	return
}

func (ptr *FunctionEvent) DestroyFunctionEvent() {
}

type GenericTiltSensor struct {
	QAccelerometerFilter
	QSensorBackend
}

type GenericTiltSensor_ITF interface {
	QAccelerometerFilter_ITF
	QSensorBackend_ITF
	GenericTiltSensor_PTR() *GenericTiltSensor
}

func (ptr *GenericTiltSensor) GenericTiltSensor_PTR() *GenericTiltSensor {
	return ptr
}

func (ptr *GenericTiltSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *GenericTiltSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAccelerometerFilter_PTR().SetPointer(p)
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromGenericTiltSensor(ptr GenericTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.GenericTiltSensor_PTR().Pointer()
	}
	return nil
}

func (n *GenericTiltSensor) InitFromInternal(ptr uintptr, name string) {
	n.QAccelerometerFilter_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *GenericTiltSensor) ClassNameInternalF() string {
	return n.QAccelerometerFilter_PTR().ClassNameInternalF()
}

func NewGenericTiltSensorFromPointer(ptr unsafe.Pointer) (n *GenericTiltSensor) {
	n = new(GenericTiltSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.GenericTiltSensor")
	return
}

func (ptr *GenericTiltSensor) DestroyGenericTiltSensor() {
}

type IIOSensorProxyCompass struct {
	IIOSensorProxySensorBase
}

type IIOSensorProxyCompass_ITF interface {
	IIOSensorProxySensorBase_ITF
	IIOSensorProxyCompass_PTR() *IIOSensorProxyCompass
}

func (ptr *IIOSensorProxyCompass) IIOSensorProxyCompass_PTR() *IIOSensorProxyCompass {
	return ptr
}

func (ptr *IIOSensorProxyCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxySensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *IIOSensorProxyCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.IIOSensorProxySensorBase_PTR().SetPointer(p)
	}
}

func PointerFromIIOSensorProxyCompass(ptr IIOSensorProxyCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxyCompass_PTR().Pointer()
	}
	return nil
}

func (n *IIOSensorProxyCompass) InitFromInternal(ptr uintptr, name string) {
	n.IIOSensorProxySensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IIOSensorProxyCompass) ClassNameInternalF() string {
	return n.IIOSensorProxySensorBase_PTR().ClassNameInternalF()
}

func NewIIOSensorProxyCompassFromPointer(ptr unsafe.Pointer) (n *IIOSensorProxyCompass) {
	n = new(IIOSensorProxyCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.IIOSensorProxyCompass")
	return
}

func (ptr *IIOSensorProxyCompass) DestroyIIOSensorProxyCompass() {
}

type IIOSensorProxyLightSensor struct {
	IIOSensorProxySensorBase
}

type IIOSensorProxyLightSensor_ITF interface {
	IIOSensorProxySensorBase_ITF
	IIOSensorProxyLightSensor_PTR() *IIOSensorProxyLightSensor
}

func (ptr *IIOSensorProxyLightSensor) IIOSensorProxyLightSensor_PTR() *IIOSensorProxyLightSensor {
	return ptr
}

func (ptr *IIOSensorProxyLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxySensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *IIOSensorProxyLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.IIOSensorProxySensorBase_PTR().SetPointer(p)
	}
}

func PointerFromIIOSensorProxyLightSensor(ptr IIOSensorProxyLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxyLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *IIOSensorProxyLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.IIOSensorProxySensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IIOSensorProxyLightSensor) ClassNameInternalF() string {
	return n.IIOSensorProxySensorBase_PTR().ClassNameInternalF()
}

func NewIIOSensorProxyLightSensorFromPointer(ptr unsafe.Pointer) (n *IIOSensorProxyLightSensor) {
	n = new(IIOSensorProxyLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.IIOSensorProxyLightSensor")
	return
}

func (ptr *IIOSensorProxyLightSensor) DestroyIIOSensorProxyLightSensor() {
}

type IIOSensorProxyOrientationSensor struct {
	IIOSensorProxySensorBase
}

type IIOSensorProxyOrientationSensor_ITF interface {
	IIOSensorProxySensorBase_ITF
	IIOSensorProxyOrientationSensor_PTR() *IIOSensorProxyOrientationSensor
}

func (ptr *IIOSensorProxyOrientationSensor) IIOSensorProxyOrientationSensor_PTR() *IIOSensorProxyOrientationSensor {
	return ptr
}

func (ptr *IIOSensorProxyOrientationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxySensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *IIOSensorProxyOrientationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.IIOSensorProxySensorBase_PTR().SetPointer(p)
	}
}

func PointerFromIIOSensorProxyOrientationSensor(ptr IIOSensorProxyOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxyOrientationSensor_PTR().Pointer()
	}
	return nil
}

func (n *IIOSensorProxyOrientationSensor) InitFromInternal(ptr uintptr, name string) {
	n.IIOSensorProxySensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IIOSensorProxyOrientationSensor) ClassNameInternalF() string {
	return n.IIOSensorProxySensorBase_PTR().ClassNameInternalF()
}

func NewIIOSensorProxyOrientationSensorFromPointer(ptr unsafe.Pointer) (n *IIOSensorProxyOrientationSensor) {
	n = new(IIOSensorProxyOrientationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.IIOSensorProxyOrientationSensor")
	return
}

func (ptr *IIOSensorProxyOrientationSensor) DestroyIIOSensorProxyOrientationSensor() {
}

type IIOSensorProxySensorBase struct {
	QSensorBackend
}

type IIOSensorProxySensorBase_ITF interface {
	QSensorBackend_ITF
	IIOSensorProxySensorBase_PTR() *IIOSensorProxySensorBase
}

func (ptr *IIOSensorProxySensorBase) IIOSensorProxySensorBase_PTR() *IIOSensorProxySensorBase {
	return ptr
}

func (ptr *IIOSensorProxySensorBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *IIOSensorProxySensorBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromIIOSensorProxySensorBase(ptr IIOSensorProxySensorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IIOSensorProxySensorBase_PTR().Pointer()
	}
	return nil
}

func (n *IIOSensorProxySensorBase) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IIOSensorProxySensorBase) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewIIOSensorProxySensorBaseFromPointer(ptr unsafe.Pointer) (n *IIOSensorProxySensorBase) {
	n = new(IIOSensorProxySensorBase)
	n.InitFromInternal(uintptr(ptr), "sensors.IIOSensorProxySensorBase")
	return
}

func (ptr *IIOSensorProxySensorBase) DestroyIIOSensorProxySensorBase() {
}

type IOSAccelerometer struct {
	QSensorBackend
}

type IOSAccelerometer_ITF interface {
	QSensorBackend_ITF
	IOSAccelerometer_PTR() *IOSAccelerometer
}

func (ptr *IOSAccelerometer) IOSAccelerometer_PTR() *IOSAccelerometer {
	return ptr
}

func (ptr *IOSAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *IOSAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromIOSAccelerometer(ptr IOSAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IOSAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *IOSAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IOSAccelerometer) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewIOSAccelerometerFromPointer(ptr unsafe.Pointer) (n *IOSAccelerometer) {
	n = new(IOSAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.IOSAccelerometer")
	return
}

func (ptr *IOSAccelerometer) DestroyIOSAccelerometer() {
}

type IOSCompass struct {
	QSensorBackend
}

type IOSCompass_ITF interface {
	QSensorBackend_ITF
	IOSCompass_PTR() *IOSCompass
}

func (ptr *IOSCompass) IOSCompass_PTR() *IOSCompass {
	return ptr
}

func (ptr *IOSCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *IOSCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromIOSCompass(ptr IOSCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IOSCompass_PTR().Pointer()
	}
	return nil
}

func (n *IOSCompass) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IOSCompass) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewIOSCompassFromPointer(ptr unsafe.Pointer) (n *IOSCompass) {
	n = new(IOSCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.IOSCompass")
	return
}

func (ptr *IOSCompass) DestroyIOSCompass() {
}

type IOSGyroscope struct {
	QSensorBackend
}

type IOSGyroscope_ITF interface {
	QSensorBackend_ITF
	IOSGyroscope_PTR() *IOSGyroscope
}

func (ptr *IOSGyroscope) IOSGyroscope_PTR() *IOSGyroscope {
	return ptr
}

func (ptr *IOSGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *IOSGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromIOSGyroscope(ptr IOSGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IOSGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *IOSGyroscope) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IOSGyroscope) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewIOSGyroscopeFromPointer(ptr unsafe.Pointer) (n *IOSGyroscope) {
	n = new(IOSGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.IOSGyroscope")
	return
}

func (ptr *IOSGyroscope) DestroyIOSGyroscope() {
}

type IOSMagnetometer struct {
	QSensorBackend
}

type IOSMagnetometer_ITF interface {
	QSensorBackend_ITF
	IOSMagnetometer_PTR() *IOSMagnetometer
}

func (ptr *IOSMagnetometer) IOSMagnetometer_PTR() *IOSMagnetometer {
	return ptr
}

func (ptr *IOSMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *IOSMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromIOSMagnetometer(ptr IOSMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IOSMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *IOSMagnetometer) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IOSMagnetometer) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewIOSMagnetometerFromPointer(ptr unsafe.Pointer) (n *IOSMagnetometer) {
	n = new(IOSMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.IOSMagnetometer")
	return
}

func (ptr *IOSMagnetometer) DestroyIOSMagnetometer() {
}

type IOSProximitySensor struct {
	QSensorBackend
}

type IOSProximitySensor_ITF interface {
	QSensorBackend_ITF
	IOSProximitySensor_PTR() *IOSProximitySensor
}

func (ptr *IOSProximitySensor) IOSProximitySensor_PTR() *IOSProximitySensor {
	return ptr
}

func (ptr *IOSProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *IOSProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromIOSProximitySensor(ptr IOSProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.IOSProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *IOSProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *IOSProximitySensor) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewIOSProximitySensorFromPointer(ptr unsafe.Pointer) (n *IOSProximitySensor) {
	n = new(IOSProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.IOSProximitySensor")
	return
}

func (ptr *IOSProximitySensor) DestroyIOSProximitySensor() {
}

type LinuxSysAccelerometer struct {
	QSensorBackend
}

type LinuxSysAccelerometer_ITF interface {
	QSensorBackend_ITF
	LinuxSysAccelerometer_PTR() *LinuxSysAccelerometer
}

func (ptr *LinuxSysAccelerometer) LinuxSysAccelerometer_PTR() *LinuxSysAccelerometer {
	return ptr
}

func (ptr *LinuxSysAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *LinuxSysAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromLinuxSysAccelerometer(ptr LinuxSysAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.LinuxSysAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *LinuxSysAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *LinuxSysAccelerometer) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewLinuxSysAccelerometerFromPointer(ptr unsafe.Pointer) (n *LinuxSysAccelerometer) {
	n = new(LinuxSysAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.LinuxSysAccelerometer")
	return
}

func (ptr *LinuxSysAccelerometer) DestroyLinuxSysAccelerometer() {
}

type QAccelerometer struct {
	QSensor
}

type QAccelerometer_ITF interface {
	QSensor_ITF
	QAccelerometer_PTR() *QAccelerometer
}

func (ptr *QAccelerometer) QAccelerometer_PTR() *QAccelerometer {
	return ptr
}

func (ptr *QAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAccelerometer(ptr QAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *QAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAccelerometer) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQAccelerometerFromPointer(ptr unsafe.Pointer) (n *QAccelerometer) {
	n = new(QAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.QAccelerometer")
	return
}

//go:generate stringer -type=QAccelerometer__AccelerationMode
//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int64

const (
	QAccelerometer__Combined QAccelerometer__AccelerationMode = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  QAccelerometer__AccelerationMode = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     QAccelerometer__AccelerationMode = QAccelerometer__AccelerationMode(2)
)

func NewQAccelerometer(parent core.QObject_ITF) *QAccelerometer {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQAccelerometer", "", parent}).(*QAccelerometer)
}

func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {

	return QAccelerometer__AccelerationMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AccelerationMode"}).(float64))
}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAccelerationModeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAccelerationModeChanged"})
}

func (ptr *QAccelerometer) AccelerationModeChanged(accelerationMode QAccelerometer__AccelerationMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AccelerationModeChanged", accelerationMode})
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QAccelerometerReading)
}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAccelerationMode", accelerationMode})
}

func (ptr *QAccelerometer) ConnectDestroyQAccelerometer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAccelerometer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAccelerometer) DisconnectDestroyQAccelerometer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAccelerometer"})
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAccelerometer"})
}

func (ptr *QAccelerometer) DestroyQAccelerometerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAccelerometerDefault"})
}

type QAccelerometerFilter struct {
	QSensorFilter
}

type QAccelerometerFilter_ITF interface {
	QSensorFilter_ITF
	QAccelerometerFilter_PTR() *QAccelerometerFilter
}

func (ptr *QAccelerometerFilter) QAccelerometerFilter_PTR() *QAccelerometerFilter {
	return ptr
}

func (ptr *QAccelerometerFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAccelerometerFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAccelerometerFilter(ptr QAccelerometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilter_PTR().Pointer()
	}
	return nil
}

func (n *QAccelerometerFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAccelerometerFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQAccelerometerFilterFromPointer(ptr unsafe.Pointer) (n *QAccelerometerFilter) {
	n = new(QAccelerometerFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QAccelerometerFilter")
	return
}

func (ptr *QAccelerometerFilter) DestroyQAccelerometerFilter() {
}

func (ptr *QAccelerometerFilter) ConnectFilter(f func(reading *QAccelerometerReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAccelerometerFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QAccelerometerFilter) Filter(reading QAccelerometerReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QAccelerometerReading struct {
	QSensorReading
}

type QAccelerometerReading_ITF interface {
	QSensorReading_ITF
	QAccelerometerReading_PTR() *QAccelerometerReading
}

func (ptr *QAccelerometerReading) QAccelerometerReading_PTR() *QAccelerometerReading {
	return ptr
}

func (ptr *QAccelerometerReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAccelerometerReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAccelerometerReading(ptr QAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func (n *QAccelerometerReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAccelerometerReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQAccelerometerReadingFromPointer(ptr unsafe.Pointer) (n *QAccelerometerReading) {
	n = new(QAccelerometerReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QAccelerometerReading")
	return
}

func (ptr *QAccelerometerReading) DestroyQAccelerometerReading() {
}

func (ptr *QAccelerometerReading) SetX(x float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetX", x})
}

func (ptr *QAccelerometerReading) SetY(y float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetY", y})
}

func (ptr *QAccelerometerReading) SetZ(z float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZ", z})
}

func (ptr *QAccelerometerReading) X() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "X"}).(float64)
}

func (ptr *QAccelerometerReading) Y() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Y"}).(float64)
}

func (ptr *QAccelerometerReading) Z() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Z"}).(float64)
}

type QAltimeter struct {
	QSensor
}

type QAltimeter_ITF interface {
	QSensor_ITF
	QAltimeter_PTR() *QAltimeter
}

func (ptr *QAltimeter) QAltimeter_PTR() *QAltimeter {
	return ptr
}

func (ptr *QAltimeter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAltimeter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAltimeter(ptr QAltimeter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeter_PTR().Pointer()
	}
	return nil
}

func (n *QAltimeter) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAltimeter) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQAltimeterFromPointer(ptr unsafe.Pointer) (n *QAltimeter) {
	n = new(QAltimeter)
	n.InitFromInternal(uintptr(ptr), "sensors.QAltimeter")
	return
}
func NewQAltimeter(parent core.QObject_ITF) *QAltimeter {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQAltimeter", "", parent}).(*QAltimeter)
}

func (ptr *QAltimeter) Reading() *QAltimeterReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QAltimeterReading)
}

func (ptr *QAltimeter) ConnectDestroyQAltimeter(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAltimeter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAltimeter) DisconnectDestroyQAltimeter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAltimeter"})
}

func (ptr *QAltimeter) DestroyQAltimeter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAltimeter"})
}

func (ptr *QAltimeter) DestroyQAltimeterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAltimeterDefault"})
}

type QAltimeterFilter struct {
	QSensorFilter
}

type QAltimeterFilter_ITF interface {
	QSensorFilter_ITF
	QAltimeterFilter_PTR() *QAltimeterFilter
}

func (ptr *QAltimeterFilter) QAltimeterFilter_PTR() *QAltimeterFilter {
	return ptr
}

func (ptr *QAltimeterFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAltimeterFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAltimeterFilter(ptr QAltimeterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterFilter_PTR().Pointer()
	}
	return nil
}

func (n *QAltimeterFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAltimeterFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQAltimeterFilterFromPointer(ptr unsafe.Pointer) (n *QAltimeterFilter) {
	n = new(QAltimeterFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QAltimeterFilter")
	return
}

func (ptr *QAltimeterFilter) DestroyQAltimeterFilter() {
}

func (ptr *QAltimeterFilter) ConnectFilter(f func(reading *QAltimeterReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAltimeterFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QAltimeterFilter) Filter(reading QAltimeterReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QAltimeterReading struct {
	QSensorReading
}

type QAltimeterReading_ITF interface {
	QSensorReading_ITF
	QAltimeterReading_PTR() *QAltimeterReading
}

func (ptr *QAltimeterReading) QAltimeterReading_PTR() *QAltimeterReading {
	return ptr
}

func (ptr *QAltimeterReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAltimeterReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAltimeterReading(ptr QAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReading_PTR().Pointer()
	}
	return nil
}

func (n *QAltimeterReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAltimeterReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQAltimeterReadingFromPointer(ptr unsafe.Pointer) (n *QAltimeterReading) {
	n = new(QAltimeterReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QAltimeterReading")
	return
}

func (ptr *QAltimeterReading) DestroyQAltimeterReading() {
}

func (ptr *QAltimeterReading) Altitude() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Altitude"}).(float64)
}

func (ptr *QAltimeterReading) SetAltitude(altitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAltitude", altitude})
}

type QAmbientLightFilter struct {
	QSensorFilter
}

type QAmbientLightFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientLightFilter_PTR() *QAmbientLightFilter
}

func (ptr *QAmbientLightFilter) QAmbientLightFilter_PTR() *QAmbientLightFilter {
	return ptr
}

func (ptr *QAmbientLightFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientLightFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientLightFilter(ptr QAmbientLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightFilter_PTR().Pointer()
	}
	return nil
}

func (n *QAmbientLightFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAmbientLightFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQAmbientLightFilterFromPointer(ptr unsafe.Pointer) (n *QAmbientLightFilter) {
	n = new(QAmbientLightFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QAmbientLightFilter")
	return
}

func (ptr *QAmbientLightFilter) DestroyQAmbientLightFilter() {
}

func (ptr *QAmbientLightFilter) ConnectFilter(f func(reading *QAmbientLightReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAmbientLightFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QAmbientLightFilter) Filter(reading QAmbientLightReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QAmbientLightReading struct {
	QSensorReading
}

type QAmbientLightReading_ITF interface {
	QSensorReading_ITF
	QAmbientLightReading_PTR() *QAmbientLightReading
}

func (ptr *QAmbientLightReading) QAmbientLightReading_PTR() *QAmbientLightReading {
	return ptr
}

func (ptr *QAmbientLightReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientLightReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientLightReading(ptr QAmbientLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightReading_PTR().Pointer()
	}
	return nil
}

func (n *QAmbientLightReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAmbientLightReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQAmbientLightReadingFromPointer(ptr unsafe.Pointer) (n *QAmbientLightReading) {
	n = new(QAmbientLightReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QAmbientLightReading")
	return
}

func (ptr *QAmbientLightReading) DestroyQAmbientLightReading() {
}

//go:generate stringer -type=QAmbientLightReading__LightLevel
//QAmbientLightReading::LightLevel
type QAmbientLightReading__LightLevel int64

const (
	QAmbientLightReading__Undefined QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(0)
	QAmbientLightReading__Dark      QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(1)
	QAmbientLightReading__Twilight  QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(2)
	QAmbientLightReading__Light     QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(3)
	QAmbientLightReading__Bright    QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(4)
	QAmbientLightReading__Sunny     QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(5)
)

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {

	return QAmbientLightReading__LightLevel(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LightLevel"}).(float64))
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLightLevel", lightLevel})
}

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
}

func (ptr *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *QAmbientLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAmbientLightSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQAmbientLightSensorFromPointer(ptr unsafe.Pointer) (n *QAmbientLightSensor) {
	n = new(QAmbientLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QAmbientLightSensor")
	return
}
func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQAmbientLightSensor", "", parent}).(*QAmbientLightSensor)
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QAmbientLightReading)
}

func (ptr *QAmbientLightSensor) ConnectDestroyQAmbientLightSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAmbientLightSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAmbientLightSensor) DisconnectDestroyQAmbientLightSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAmbientLightSensor"})
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAmbientLightSensor"})
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAmbientLightSensorDefault"})
}

type QAmbientTemperatureFilter struct {
	QSensorFilter
}

type QAmbientTemperatureFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter
}

func (ptr *QAmbientTemperatureFilter) QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter {
	return ptr
}

func (ptr *QAmbientTemperatureFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientTemperatureFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientTemperatureFilter(ptr QAmbientTemperatureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureFilter_PTR().Pointer()
	}
	return nil
}

func (n *QAmbientTemperatureFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAmbientTemperatureFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQAmbientTemperatureFilterFromPointer(ptr unsafe.Pointer) (n *QAmbientTemperatureFilter) {
	n = new(QAmbientTemperatureFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QAmbientTemperatureFilter")
	return
}

func (ptr *QAmbientTemperatureFilter) DestroyQAmbientTemperatureFilter() {
}

func (ptr *QAmbientTemperatureFilter) ConnectFilter(f func(reading *QAmbientTemperatureReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAmbientTemperatureFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QAmbientTemperatureFilter) Filter(reading QAmbientTemperatureReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QAmbientTemperatureReading struct {
	QSensorReading
}

type QAmbientTemperatureReading_ITF interface {
	QSensorReading_ITF
	QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading
}

func (ptr *QAmbientTemperatureReading) QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading {
	return ptr
}

func (ptr *QAmbientTemperatureReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientTemperatureReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func (n *QAmbientTemperatureReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAmbientTemperatureReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) (n *QAmbientTemperatureReading) {
	n = new(QAmbientTemperatureReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QAmbientTemperatureReading")
	return
}

func (ptr *QAmbientTemperatureReading) DestroyQAmbientTemperatureReading() {
}

func (ptr *QAmbientTemperatureReading) SetTemperature(temperature float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTemperature", temperature})
}

func (ptr *QAmbientTemperatureReading) Temperature() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Temperature"}).(float64)
}

type QAmbientTemperatureSensor struct {
	QSensor
}

type QAmbientTemperatureSensor_ITF interface {
	QSensor_ITF
	QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor
}

func (ptr *QAmbientTemperatureSensor) QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor {
	return ptr
}

func (ptr *QAmbientTemperatureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientTemperatureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientTemperatureSensor(ptr QAmbientTemperatureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureSensor_PTR().Pointer()
	}
	return nil
}

func (n *QAmbientTemperatureSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAmbientTemperatureSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQAmbientTemperatureSensorFromPointer(ptr unsafe.Pointer) (n *QAmbientTemperatureSensor) {
	n = new(QAmbientTemperatureSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QAmbientTemperatureSensor")
	return
}
func NewQAmbientTemperatureSensor(parent core.QObject_ITF) *QAmbientTemperatureSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQAmbientTemperatureSensor", "", parent}).(*QAmbientTemperatureSensor)
}

func (ptr *QAmbientTemperatureSensor) Reading() *QAmbientTemperatureReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QAmbientTemperatureReading)
}

func (ptr *QAmbientTemperatureSensor) ConnectDestroyQAmbientTemperatureSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAmbientTemperatureSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAmbientTemperatureSensor) DisconnectDestroyQAmbientTemperatureSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAmbientTemperatureSensor"})
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAmbientTemperatureSensor"})
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAmbientTemperatureSensorDefault"})
}

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
}

func (ptr *QCompass) QCompass_PTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQCompass(ptr QCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompass_PTR().Pointer()
	}
	return nil
}

func (n *QCompass) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCompass) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQCompassFromPointer(ptr unsafe.Pointer) (n *QCompass) {
	n = new(QCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.QCompass")
	return
}
func NewQCompass(parent core.QObject_ITF) *QCompass {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQCompass", "", parent}).(*QCompass)
}

func (ptr *QCompass) Reading() *QCompassReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QCompassReading)
}

func (ptr *QCompass) ConnectDestroyQCompass(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQCompass", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCompass) DisconnectDestroyQCompass() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQCompass"})
}

func (ptr *QCompass) DestroyQCompass() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCompass"})
}

func (ptr *QCompass) DestroyQCompassDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCompassDefault"})
}

type QCompassFilter struct {
	QSensorFilter
}

type QCompassFilter_ITF interface {
	QSensorFilter_ITF
	QCompassFilter_PTR() *QCompassFilter
}

func (ptr *QCompassFilter) QCompassFilter_PTR() *QCompassFilter {
	return ptr
}

func (ptr *QCompassFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QCompassFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQCompassFilter(ptr QCompassFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassFilter_PTR().Pointer()
	}
	return nil
}

func (n *QCompassFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCompassFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQCompassFilterFromPointer(ptr unsafe.Pointer) (n *QCompassFilter) {
	n = new(QCompassFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QCompassFilter")
	return
}

func (ptr *QCompassFilter) DestroyQCompassFilter() {
}

func (ptr *QCompassFilter) ConnectFilter(f func(reading *QCompassReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCompassFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QCompassFilter) Filter(reading QCompassReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QCompassReading struct {
	QSensorReading
}

type QCompassReading_ITF interface {
	QSensorReading_ITF
	QCompassReading_PTR() *QCompassReading
}

func (ptr *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return ptr
}

func (ptr *QCompassReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QCompassReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQCompassReading(ptr QCompassReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassReading_PTR().Pointer()
	}
	return nil
}

func (n *QCompassReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCompassReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQCompassReadingFromPointer(ptr unsafe.Pointer) (n *QCompassReading) {
	n = new(QCompassReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QCompassReading")
	return
}

func (ptr *QCompassReading) DestroyQCompassReading() {
}

func (ptr *QCompassReading) Azimuth() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Azimuth"}).(float64)
}

func (ptr *QCompassReading) CalibrationLevel() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CalibrationLevel"}).(float64)
}

func (ptr *QCompassReading) SetAzimuth(azimuth float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAzimuth", azimuth})
}

func (ptr *QCompassReading) SetCalibrationLevel(calibrationLevel float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCalibrationLevel", calibrationLevel})
}

type QDistanceFilter struct {
	QSensorFilter
}

type QDistanceFilter_ITF interface {
	QSensorFilter_ITF
	QDistanceFilter_PTR() *QDistanceFilter
}

func (ptr *QDistanceFilter) QDistanceFilter_PTR() *QDistanceFilter {
	return ptr
}

func (ptr *QDistanceFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QDistanceFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQDistanceFilter(ptr QDistanceFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceFilter_PTR().Pointer()
	}
	return nil
}

func (n *QDistanceFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDistanceFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQDistanceFilterFromPointer(ptr unsafe.Pointer) (n *QDistanceFilter) {
	n = new(QDistanceFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QDistanceFilter")
	return
}

func (ptr *QDistanceFilter) DestroyQDistanceFilter() {
}

func (ptr *QDistanceFilter) ConnectFilter(f func(reading *QDistanceReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDistanceFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QDistanceFilter) Filter(reading QDistanceReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QDistanceReading struct {
	QSensorReading
}

type QDistanceReading_ITF interface {
	QSensorReading_ITF
	QDistanceReading_PTR() *QDistanceReading
}

func (ptr *QDistanceReading) QDistanceReading_PTR() *QDistanceReading {
	return ptr
}

func (ptr *QDistanceReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QDistanceReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQDistanceReading(ptr QDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReading_PTR().Pointer()
	}
	return nil
}

func (n *QDistanceReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDistanceReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQDistanceReadingFromPointer(ptr unsafe.Pointer) (n *QDistanceReading) {
	n = new(QDistanceReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QDistanceReading")
	return
}

func (ptr *QDistanceReading) DestroyQDistanceReading() {
}

func (ptr *QDistanceReading) Distance() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Distance"}).(float64)
}

func (ptr *QDistanceReading) SetDistance(distance float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDistance", distance})
}

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensor_ITF interface {
	QSensor_ITF
	QDistanceSensor_PTR() *QDistanceSensor
}

func (ptr *QDistanceSensor) QDistanceSensor_PTR() *QDistanceSensor {
	return ptr
}

func (ptr *QDistanceSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QDistanceSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQDistanceSensor(ptr QDistanceSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceSensor_PTR().Pointer()
	}
	return nil
}

func (n *QDistanceSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDistanceSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQDistanceSensorFromPointer(ptr unsafe.Pointer) (n *QDistanceSensor) {
	n = new(QDistanceSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QDistanceSensor")
	return
}
func NewQDistanceSensor(parent core.QObject_ITF) *QDistanceSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQDistanceSensor", "", parent}).(*QDistanceSensor)
}

func (ptr *QDistanceSensor) Reading() *QDistanceReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QDistanceReading)
}

func (ptr *QDistanceSensor) ConnectDestroyQDistanceSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDistanceSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDistanceSensor) DisconnectDestroyQDistanceSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDistanceSensor"})
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDistanceSensor"})
}

func (ptr *QDistanceSensor) DestroyQDistanceSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDistanceSensorDefault"})
}

type QGyroscope struct {
	QSensor
}

type QGyroscope_ITF interface {
	QSensor_ITF
	QGyroscope_PTR() *QGyroscope
}

func (ptr *QGyroscope) QGyroscope_PTR() *QGyroscope {
	return ptr
}

func (ptr *QGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQGyroscope(ptr QGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *QGyroscope) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGyroscope) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQGyroscopeFromPointer(ptr unsafe.Pointer) (n *QGyroscope) {
	n = new(QGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.QGyroscope")
	return
}
func NewQGyroscope(parent core.QObject_ITF) *QGyroscope {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQGyroscope", "", parent}).(*QGyroscope)
}

func (ptr *QGyroscope) Reading() *QGyroscopeReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QGyroscopeReading)
}

func (ptr *QGyroscope) ConnectDestroyQGyroscope(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGyroscope", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGyroscope) DisconnectDestroyQGyroscope() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGyroscope"})
}

func (ptr *QGyroscope) DestroyQGyroscope() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGyroscope"})
}

func (ptr *QGyroscope) DestroyQGyroscopeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGyroscopeDefault"})
}

type QGyroscopeFilter struct {
	QSensorFilter
}

type QGyroscopeFilter_ITF interface {
	QSensorFilter_ITF
	QGyroscopeFilter_PTR() *QGyroscopeFilter
}

func (ptr *QGyroscopeFilter) QGyroscopeFilter_PTR() *QGyroscopeFilter {
	return ptr
}

func (ptr *QGyroscopeFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QGyroscopeFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQGyroscopeFilter(ptr QGyroscopeFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeFilter_PTR().Pointer()
	}
	return nil
}

func (n *QGyroscopeFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGyroscopeFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQGyroscopeFilterFromPointer(ptr unsafe.Pointer) (n *QGyroscopeFilter) {
	n = new(QGyroscopeFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QGyroscopeFilter")
	return
}

func (ptr *QGyroscopeFilter) DestroyQGyroscopeFilter() {
}

func (ptr *QGyroscopeFilter) ConnectFilter(f func(reading *QGyroscopeReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGyroscopeFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QGyroscopeFilter) Filter(reading QGyroscopeReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QGyroscopeReading struct {
	QSensorReading
}

type QGyroscopeReading_ITF interface {
	QSensorReading_ITF
	QGyroscopeReading_PTR() *QGyroscopeReading
}

func (ptr *QGyroscopeReading) QGyroscopeReading_PTR() *QGyroscopeReading {
	return ptr
}

func (ptr *QGyroscopeReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QGyroscopeReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQGyroscopeReading(ptr QGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func (n *QGyroscopeReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGyroscopeReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQGyroscopeReadingFromPointer(ptr unsafe.Pointer) (n *QGyroscopeReading) {
	n = new(QGyroscopeReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QGyroscopeReading")
	return
}

func (ptr *QGyroscopeReading) DestroyQGyroscopeReading() {
}

func (ptr *QGyroscopeReading) SetX(x float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetX", x})
}

func (ptr *QGyroscopeReading) SetY(y float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetY", y})
}

func (ptr *QGyroscopeReading) SetZ(z float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZ", z})
}

func (ptr *QGyroscopeReading) X() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "X"}).(float64)
}

func (ptr *QGyroscopeReading) Y() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Y"}).(float64)
}

func (ptr *QGyroscopeReading) Z() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Z"}).(float64)
}

type QHolsterFilter struct {
	QSensorFilter
}

type QHolsterFilter_ITF interface {
	QSensorFilter_ITF
	QHolsterFilter_PTR() *QHolsterFilter
}

func (ptr *QHolsterFilter) QHolsterFilter_PTR() *QHolsterFilter {
	return ptr
}

func (ptr *QHolsterFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QHolsterFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQHolsterFilter(ptr QHolsterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterFilter_PTR().Pointer()
	}
	return nil
}

func (n *QHolsterFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHolsterFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQHolsterFilterFromPointer(ptr unsafe.Pointer) (n *QHolsterFilter) {
	n = new(QHolsterFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QHolsterFilter")
	return
}

func (ptr *QHolsterFilter) DestroyQHolsterFilter() {
}

func (ptr *QHolsterFilter) ConnectFilter(f func(reading *QHolsterReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHolsterFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QHolsterFilter) Filter(reading QHolsterReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QHolsterReading struct {
	QSensorReading
}

type QHolsterReading_ITF interface {
	QSensorReading_ITF
	QHolsterReading_PTR() *QHolsterReading
}

func (ptr *QHolsterReading) QHolsterReading_PTR() *QHolsterReading {
	return ptr
}

func (ptr *QHolsterReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QHolsterReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQHolsterReading(ptr QHolsterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterReading_PTR().Pointer()
	}
	return nil
}

func (n *QHolsterReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHolsterReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQHolsterReadingFromPointer(ptr unsafe.Pointer) (n *QHolsterReading) {
	n = new(QHolsterReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QHolsterReading")
	return
}

func (ptr *QHolsterReading) DestroyQHolsterReading() {
}

func (ptr *QHolsterReading) Holstered() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Holstered"}).(bool)
}

func (ptr *QHolsterReading) SetHolstered(holstered bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHolstered", holstered})
}

type QHolsterSensor struct {
	QSensor
}

type QHolsterSensor_ITF interface {
	QSensor_ITF
	QHolsterSensor_PTR() *QHolsterSensor
}

func (ptr *QHolsterSensor) QHolsterSensor_PTR() *QHolsterSensor {
	return ptr
}

func (ptr *QHolsterSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QHolsterSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQHolsterSensor(ptr QHolsterSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterSensor_PTR().Pointer()
	}
	return nil
}

func (n *QHolsterSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHolsterSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQHolsterSensorFromPointer(ptr unsafe.Pointer) (n *QHolsterSensor) {
	n = new(QHolsterSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QHolsterSensor")
	return
}
func NewQHolsterSensor(parent core.QObject_ITF) *QHolsterSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQHolsterSensor", "", parent}).(*QHolsterSensor)
}

func (ptr *QHolsterSensor) Reading() *QHolsterReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QHolsterReading)
}

func (ptr *QHolsterSensor) ConnectDestroyQHolsterSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHolsterSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHolsterSensor) DisconnectDestroyQHolsterSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHolsterSensor"})
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHolsterSensor"})
}

func (ptr *QHolsterSensor) DestroyQHolsterSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHolsterSensorDefault"})
}

type QHumidityFilter struct {
	QSensorFilter
}

type QHumidityFilter_ITF interface {
	QSensorFilter_ITF
	QHumidityFilter_PTR() *QHumidityFilter
}

func (ptr *QHumidityFilter) QHumidityFilter_PTR() *QHumidityFilter {
	return ptr
}

func (ptr *QHumidityFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QHumidityFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQHumidityFilter(ptr QHumidityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHumidityFilter_PTR().Pointer()
	}
	return nil
}

func (n *QHumidityFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHumidityFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQHumidityFilterFromPointer(ptr unsafe.Pointer) (n *QHumidityFilter) {
	n = new(QHumidityFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QHumidityFilter")
	return
}

func (ptr *QHumidityFilter) DestroyQHumidityFilter() {
}

func (ptr *QHumidityFilter) ConnectFilter(f func(reading *QHumidityReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHumidityFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QHumidityFilter) Filter(reading QHumidityReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QHumidityReading struct {
	QSensorReading
}

type QHumidityReading_ITF interface {
	QSensorReading_ITF
	QHumidityReading_PTR() *QHumidityReading
}

func (ptr *QHumidityReading) QHumidityReading_PTR() *QHumidityReading {
	return ptr
}

func (ptr *QHumidityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QHumidityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQHumidityReading(ptr QHumidityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHumidityReading_PTR().Pointer()
	}
	return nil
}

func (n *QHumidityReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHumidityReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQHumidityReadingFromPointer(ptr unsafe.Pointer) (n *QHumidityReading) {
	n = new(QHumidityReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QHumidityReading")
	return
}

func (ptr *QHumidityReading) DestroyQHumidityReading() {
}

func (ptr *QHumidityReading) AbsoluteHumidity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AbsoluteHumidity"}).(float64)
}

func (ptr *QHumidityReading) RelativeHumidity() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RelativeHumidity"}).(float64)
}

func (ptr *QHumidityReading) SetAbsoluteHumidity(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAbsoluteHumidity", value})
}

func (ptr *QHumidityReading) SetRelativeHumidity(humidity float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRelativeHumidity", humidity})
}

type QHumiditySensor struct {
	QSensor
}

type QHumiditySensor_ITF interface {
	QSensor_ITF
	QHumiditySensor_PTR() *QHumiditySensor
}

func (ptr *QHumiditySensor) QHumiditySensor_PTR() *QHumiditySensor {
	return ptr
}

func (ptr *QHumiditySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QHumiditySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQHumiditySensor(ptr QHumiditySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHumiditySensor_PTR().Pointer()
	}
	return nil
}

func (n *QHumiditySensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHumiditySensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQHumiditySensorFromPointer(ptr unsafe.Pointer) (n *QHumiditySensor) {
	n = new(QHumiditySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QHumiditySensor")
	return
}
func NewQHumiditySensor(parent core.QObject_ITF) *QHumiditySensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQHumiditySensor", "", parent}).(*QHumiditySensor)
}

func (ptr *QHumiditySensor) Reading() *QHumidityReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QHumidityReading)
}

func (ptr *QHumiditySensor) ConnectDestroyQHumiditySensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHumiditySensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHumiditySensor) DisconnectDestroyQHumiditySensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHumiditySensor"})
}

func (ptr *QHumiditySensor) DestroyQHumiditySensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHumiditySensor"})
}

func (ptr *QHumiditySensor) DestroyQHumiditySensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHumiditySensorDefault"})
}

type QIRProximityFilter struct {
	QSensorFilter
}

type QIRProximityFilter_ITF interface {
	QSensorFilter_ITF
	QIRProximityFilter_PTR() *QIRProximityFilter
}

func (ptr *QIRProximityFilter) QIRProximityFilter_PTR() *QIRProximityFilter {
	return ptr
}

func (ptr *QIRProximityFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QIRProximityFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQIRProximityFilter(ptr QIRProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityFilter_PTR().Pointer()
	}
	return nil
}

func (n *QIRProximityFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QIRProximityFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQIRProximityFilterFromPointer(ptr unsafe.Pointer) (n *QIRProximityFilter) {
	n = new(QIRProximityFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QIRProximityFilter")
	return
}

func (ptr *QIRProximityFilter) DestroyQIRProximityFilter() {
}

func (ptr *QIRProximityFilter) ConnectFilter(f func(reading *QIRProximityReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QIRProximityFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QIRProximityFilter) Filter(reading QIRProximityReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QIRProximityReading struct {
	QSensorReading
}

type QIRProximityReading_ITF interface {
	QSensorReading_ITF
	QIRProximityReading_PTR() *QIRProximityReading
}

func (ptr *QIRProximityReading) QIRProximityReading_PTR() *QIRProximityReading {
	return ptr
}

func (ptr *QIRProximityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QIRProximityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQIRProximityReading(ptr QIRProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityReading_PTR().Pointer()
	}
	return nil
}

func (n *QIRProximityReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QIRProximityReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQIRProximityReadingFromPointer(ptr unsafe.Pointer) (n *QIRProximityReading) {
	n = new(QIRProximityReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QIRProximityReading")
	return
}

func (ptr *QIRProximityReading) DestroyQIRProximityReading() {
}

func (ptr *QIRProximityReading) Reflectance() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reflectance"}).(float64)
}

func (ptr *QIRProximityReading) SetReflectance(reflectance float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReflectance", reflectance})
}

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensor_ITF interface {
	QSensor_ITF
	QIRProximitySensor_PTR() *QIRProximitySensor
}

func (ptr *QIRProximitySensor) QIRProximitySensor_PTR() *QIRProximitySensor {
	return ptr
}

func (ptr *QIRProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QIRProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQIRProximitySensor(ptr QIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *QIRProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QIRProximitySensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQIRProximitySensorFromPointer(ptr unsafe.Pointer) (n *QIRProximitySensor) {
	n = new(QIRProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QIRProximitySensor")
	return
}
func NewQIRProximitySensor(parent core.QObject_ITF) *QIRProximitySensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQIRProximitySensor", "", parent}).(*QIRProximitySensor)
}

func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QIRProximityReading)
}

func (ptr *QIRProximitySensor) ConnectDestroyQIRProximitySensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQIRProximitySensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QIRProximitySensor) DisconnectDestroyQIRProximitySensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQIRProximitySensor"})
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQIRProximitySensor"})
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQIRProximitySensorDefault"})
}

type QLidFilter struct {
	QSensorFilter
}

type QLidFilter_ITF interface {
	QSensorFilter_ITF
	QLidFilter_PTR() *QLidFilter
}

func (ptr *QLidFilter) QLidFilter_PTR() *QLidFilter {
	return ptr
}

func (ptr *QLidFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QLidFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQLidFilter(ptr QLidFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLidFilter_PTR().Pointer()
	}
	return nil
}

func (n *QLidFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLidFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQLidFilterFromPointer(ptr unsafe.Pointer) (n *QLidFilter) {
	n = new(QLidFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QLidFilter")
	return
}

func (ptr *QLidFilter) DestroyQLidFilter() {
}

func (ptr *QLidFilter) ConnectFilter(f func(reading *QLidReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLidFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QLidFilter) Filter(reading QLidReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QLidReading struct {
	QSensorReading
}

type QLidReading_ITF interface {
	QSensorReading_ITF
	QLidReading_PTR() *QLidReading
}

func (ptr *QLidReading) QLidReading_PTR() *QLidReading {
	return ptr
}

func (ptr *QLidReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QLidReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQLidReading(ptr QLidReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLidReading_PTR().Pointer()
	}
	return nil
}

func (n *QLidReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLidReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQLidReadingFromPointer(ptr unsafe.Pointer) (n *QLidReading) {
	n = new(QLidReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QLidReading")
	return
}

func (ptr *QLidReading) DestroyQLidReading() {
}

func (ptr *QLidReading) BackLidClosed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackLidClosed"}).(bool)
}

func (ptr *QLidReading) FrontLidClosed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FrontLidClosed"}).(bool)
}

func (ptr *QLidReading) SetBackLidClosed(closed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackLidClosed", closed})
}

func (ptr *QLidReading) SetFrontLidClosed(closed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFrontLidClosed", closed})
}

type QLidSensor struct {
	QSensor
}

type QLidSensor_ITF interface {
	QSensor_ITF
	QLidSensor_PTR() *QLidSensor
}

func (ptr *QLidSensor) QLidSensor_PTR() *QLidSensor {
	return ptr
}

func (ptr *QLidSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QLidSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQLidSensor(ptr QLidSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLidSensor_PTR().Pointer()
	}
	return nil
}

func (n *QLidSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLidSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQLidSensorFromPointer(ptr unsafe.Pointer) (n *QLidSensor) {
	n = new(QLidSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QLidSensor")
	return
}
func NewQLidSensor(parent core.QObject_ITF) *QLidSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQLidSensor", "", parent}).(*QLidSensor)
}

func (ptr *QLidSensor) Reading() *QLidReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QLidReading)
}

func (ptr *QLidSensor) ConnectDestroyQLidSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLidSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLidSensor) DisconnectDestroyQLidSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLidSensor"})
}

func (ptr *QLidSensor) DestroyQLidSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLidSensor"})
}

func (ptr *QLidSensor) DestroyQLidSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLidSensorDefault"})
}

type QLightFilter struct {
	QSensorFilter
}

type QLightFilter_ITF interface {
	QSensorFilter_ITF
	QLightFilter_PTR() *QLightFilter
}

func (ptr *QLightFilter) QLightFilter_PTR() *QLightFilter {
	return ptr
}

func (ptr *QLightFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QLightFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQLightFilter(ptr QLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightFilter_PTR().Pointer()
	}
	return nil
}

func (n *QLightFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLightFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQLightFilterFromPointer(ptr unsafe.Pointer) (n *QLightFilter) {
	n = new(QLightFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QLightFilter")
	return
}

func (ptr *QLightFilter) DestroyQLightFilter() {
}

func (ptr *QLightFilter) ConnectFilter(f func(reading *QLightReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLightFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QLightFilter) Filter(reading QLightReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QLightReading struct {
	QSensorReading
}

type QLightReading_ITF interface {
	QSensorReading_ITF
	QLightReading_PTR() *QLightReading
}

func (ptr *QLightReading) QLightReading_PTR() *QLightReading {
	return ptr
}

func (ptr *QLightReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QLightReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQLightReading(ptr QLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReading_PTR().Pointer()
	}
	return nil
}

func (n *QLightReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLightReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQLightReadingFromPointer(ptr unsafe.Pointer) (n *QLightReading) {
	n = new(QLightReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QLightReading")
	return
}

func (ptr *QLightReading) DestroyQLightReading() {
}

func (ptr *QLightReading) Lux() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Lux"}).(float64)
}

func (ptr *QLightReading) SetLux(lux float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLux", lux})
}

type QLightSensor struct {
	QSensor
}

type QLightSensor_ITF interface {
	QSensor_ITF
	QLightSensor_PTR() *QLightSensor
}

func (ptr *QLightSensor) QLightSensor_PTR() *QLightSensor {
	return ptr
}

func (ptr *QLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQLightSensor(ptr QLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *QLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLightSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQLightSensorFromPointer(ptr unsafe.Pointer) (n *QLightSensor) {
	n = new(QLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QLightSensor")
	return
}
func NewQLightSensor(parent core.QObject_ITF) *QLightSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQLightSensor", "", parent}).(*QLightSensor)
}

func (ptr *QLightSensor) FieldOfView() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FieldOfView"}).(float64)
}

func (ptr *QLightSensor) ConnectFieldOfViewChanged(f func(fieldOfView float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFieldOfViewChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLightSensor) DisconnectFieldOfViewChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFieldOfViewChanged"})
}

func (ptr *QLightSensor) FieldOfViewChanged(fieldOfView float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FieldOfViewChanged", fieldOfView})
}

func (ptr *QLightSensor) Reading() *QLightReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QLightReading)
}

func (ptr *QLightSensor) SetFieldOfView(fieldOfView float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFieldOfView", fieldOfView})
}

func (ptr *QLightSensor) ConnectDestroyQLightSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLightSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLightSensor) DisconnectDestroyQLightSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLightSensor"})
}

func (ptr *QLightSensor) DestroyQLightSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLightSensor"})
}

func (ptr *QLightSensor) DestroyQLightSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLightSensorDefault"})
}

type QMagnetometer struct {
	QSensor
}

type QMagnetometer_ITF interface {
	QSensor_ITF
	QMagnetometer_PTR() *QMagnetometer
}

func (ptr *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQMagnetometer(ptr QMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *QMagnetometer) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QMagnetometer) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQMagnetometerFromPointer(ptr unsafe.Pointer) (n *QMagnetometer) {
	n = new(QMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.QMagnetometer")
	return
}
func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQMagnetometer", "", parent}).(*QMagnetometer)
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QMagnetometerReading)
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReturnGeoValues"}).(bool)
}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReturnGeoValuesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReturnGeoValuesChanged"})
}

func (ptr *QMagnetometer) ReturnGeoValuesChanged(returnGeoValues bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReturnGeoValuesChanged", returnGeoValues})
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReturnGeoValues", returnGeoValues})
}

func (ptr *QMagnetometer) ConnectDestroyQMagnetometer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQMagnetometer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMagnetometer) DisconnectDestroyQMagnetometer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQMagnetometer"})
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQMagnetometer"})
}

func (ptr *QMagnetometer) DestroyQMagnetometerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQMagnetometerDefault"})
}

type QMagnetometerFilter struct {
	QSensorFilter
}

type QMagnetometerFilter_ITF interface {
	QSensorFilter_ITF
	QMagnetometerFilter_PTR() *QMagnetometerFilter
}

func (ptr *QMagnetometerFilter) QMagnetometerFilter_PTR() *QMagnetometerFilter {
	return ptr
}

func (ptr *QMagnetometerFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QMagnetometerFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQMagnetometerFilter(ptr QMagnetometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerFilter_PTR().Pointer()
	}
	return nil
}

func (n *QMagnetometerFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QMagnetometerFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQMagnetometerFilterFromPointer(ptr unsafe.Pointer) (n *QMagnetometerFilter) {
	n = new(QMagnetometerFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QMagnetometerFilter")
	return
}

func (ptr *QMagnetometerFilter) DestroyQMagnetometerFilter() {
}

func (ptr *QMagnetometerFilter) ConnectFilter(f func(reading *QMagnetometerReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMagnetometerFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QMagnetometerFilter) Filter(reading QMagnetometerReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QMagnetometerReading struct {
	QSensorReading
}

type QMagnetometerReading_ITF interface {
	QSensorReading_ITF
	QMagnetometerReading_PTR() *QMagnetometerReading
}

func (ptr *QMagnetometerReading) QMagnetometerReading_PTR() *QMagnetometerReading {
	return ptr
}

func (ptr *QMagnetometerReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QMagnetometerReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQMagnetometerReading(ptr QMagnetometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerReading_PTR().Pointer()
	}
	return nil
}

func (n *QMagnetometerReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QMagnetometerReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQMagnetometerReadingFromPointer(ptr unsafe.Pointer) (n *QMagnetometerReading) {
	n = new(QMagnetometerReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QMagnetometerReading")
	return
}

func (ptr *QMagnetometerReading) DestroyQMagnetometerReading() {
}

func (ptr *QMagnetometerReading) CalibrationLevel() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CalibrationLevel"}).(float64)
}

func (ptr *QMagnetometerReading) SetCalibrationLevel(calibrationLevel float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCalibrationLevel", calibrationLevel})
}

func (ptr *QMagnetometerReading) SetX(x float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetX", x})
}

func (ptr *QMagnetometerReading) SetY(y float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetY", y})
}

func (ptr *QMagnetometerReading) SetZ(z float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetZ", z})
}

func (ptr *QMagnetometerReading) X() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "X"}).(float64)
}

func (ptr *QMagnetometerReading) Y() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Y"}).(float64)
}

func (ptr *QMagnetometerReading) Z() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Z"}).(float64)
}

type QOrientationFilter struct {
	QSensorFilter
}

type QOrientationFilter_ITF interface {
	QSensorFilter_ITF
	QOrientationFilter_PTR() *QOrientationFilter
}

func (ptr *QOrientationFilter) QOrientationFilter_PTR() *QOrientationFilter {
	return ptr
}

func (ptr *QOrientationFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QOrientationFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQOrientationFilter(ptr QOrientationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationFilter_PTR().Pointer()
	}
	return nil
}

func (n *QOrientationFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QOrientationFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQOrientationFilterFromPointer(ptr unsafe.Pointer) (n *QOrientationFilter) {
	n = new(QOrientationFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QOrientationFilter")
	return
}

func (ptr *QOrientationFilter) DestroyQOrientationFilter() {
}

func (ptr *QOrientationFilter) ConnectFilter(f func(reading *QOrientationReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QOrientationFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QOrientationFilter) Filter(reading QOrientationReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QOrientationReading struct {
	QSensorReading
}

type QOrientationReading_ITF interface {
	QSensorReading_ITF
	QOrientationReading_PTR() *QOrientationReading
}

func (ptr *QOrientationReading) QOrientationReading_PTR() *QOrientationReading {
	return ptr
}

func (ptr *QOrientationReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QOrientationReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQOrientationReading(ptr QOrientationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationReading_PTR().Pointer()
	}
	return nil
}

func (n *QOrientationReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QOrientationReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQOrientationReadingFromPointer(ptr unsafe.Pointer) (n *QOrientationReading) {
	n = new(QOrientationReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QOrientationReading")
	return
}

func (ptr *QOrientationReading) DestroyQOrientationReading() {
}

//go:generate stringer -type=QOrientationReading__Orientation
//QOrientationReading::Orientation
type QOrientationReading__Orientation int64

const (
	QOrientationReading__Undefined QOrientationReading__Orientation = QOrientationReading__Orientation(0)
	QOrientationReading__TopUp     QOrientationReading__Orientation = QOrientationReading__Orientation(1)
	QOrientationReading__TopDown   QOrientationReading__Orientation = QOrientationReading__Orientation(2)
	QOrientationReading__LeftUp    QOrientationReading__Orientation = QOrientationReading__Orientation(3)
	QOrientationReading__RightUp   QOrientationReading__Orientation = QOrientationReading__Orientation(4)
	QOrientationReading__FaceUp    QOrientationReading__Orientation = QOrientationReading__Orientation(5)
	QOrientationReading__FaceDown  QOrientationReading__Orientation = QOrientationReading__Orientation(6)
)

func (ptr *QOrientationReading) Orientation() QOrientationReading__Orientation {

	return QOrientationReading__Orientation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Orientation"}).(float64))
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOrientation", orientation})
}

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensor_ITF interface {
	QSensor_ITF
	QOrientationSensor_PTR() *QOrientationSensor
}

func (ptr *QOrientationSensor) QOrientationSensor_PTR() *QOrientationSensor {
	return ptr
}

func (ptr *QOrientationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QOrientationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQOrientationSensor(ptr QOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationSensor_PTR().Pointer()
	}
	return nil
}

func (n *QOrientationSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QOrientationSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQOrientationSensorFromPointer(ptr unsafe.Pointer) (n *QOrientationSensor) {
	n = new(QOrientationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QOrientationSensor")
	return
}
func NewQOrientationSensor(parent core.QObject_ITF) *QOrientationSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQOrientationSensor", "", parent}).(*QOrientationSensor)
}

func (ptr *QOrientationSensor) Reading() *QOrientationReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QOrientationReading)
}

func (ptr *QOrientationSensor) ConnectDestroyQOrientationSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQOrientationSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QOrientationSensor) DisconnectDestroyQOrientationSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQOrientationSensor"})
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQOrientationSensor"})
}

func (ptr *QOrientationSensor) DestroyQOrientationSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQOrientationSensorDefault"})
}

type QPressureFilter struct {
	QSensorFilter
}

type QPressureFilter_ITF interface {
	QSensorFilter_ITF
	QPressureFilter_PTR() *QPressureFilter
}

func (ptr *QPressureFilter) QPressureFilter_PTR() *QPressureFilter {
	return ptr
}

func (ptr *QPressureFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QPressureFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQPressureFilter(ptr QPressureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureFilter_PTR().Pointer()
	}
	return nil
}

func (n *QPressureFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPressureFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQPressureFilterFromPointer(ptr unsafe.Pointer) (n *QPressureFilter) {
	n = new(QPressureFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QPressureFilter")
	return
}

func (ptr *QPressureFilter) DestroyQPressureFilter() {
}

func (ptr *QPressureFilter) ConnectFilter(f func(reading *QPressureReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPressureFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QPressureFilter) Filter(reading QPressureReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QPressureReading struct {
	QSensorReading
}

type QPressureReading_ITF interface {
	QSensorReading_ITF
	QPressureReading_PTR() *QPressureReading
}

func (ptr *QPressureReading) QPressureReading_PTR() *QPressureReading {
	return ptr
}

func (ptr *QPressureReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QPressureReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQPressureReading(ptr QPressureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureReading_PTR().Pointer()
	}
	return nil
}

func (n *QPressureReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPressureReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQPressureReadingFromPointer(ptr unsafe.Pointer) (n *QPressureReading) {
	n = new(QPressureReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QPressureReading")
	return
}

func (ptr *QPressureReading) DestroyQPressureReading() {
}

func (ptr *QPressureReading) Pressure() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pressure"}).(float64)
}

func (ptr *QPressureReading) SetPressure(pressure float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPressure", pressure})
}

func (ptr *QPressureReading) SetTemperature(temperature float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTemperature", temperature})
}

func (ptr *QPressureReading) Temperature() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Temperature"}).(float64)
}

type QPressureSensor struct {
	QSensor
}

type QPressureSensor_ITF interface {
	QSensor_ITF
	QPressureSensor_PTR() *QPressureSensor
}

func (ptr *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QPressureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQPressureSensor(ptr QPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureSensor_PTR().Pointer()
	}
	return nil
}

func (n *QPressureSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPressureSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQPressureSensorFromPointer(ptr unsafe.Pointer) (n *QPressureSensor) {
	n = new(QPressureSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QPressureSensor")
	return
}
func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQPressureSensor", "", parent}).(*QPressureSensor)
}

func (ptr *QPressureSensor) Reading() *QPressureReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QPressureReading)
}

func (ptr *QPressureSensor) ConnectDestroyQPressureSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQPressureSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QPressureSensor) DisconnectDestroyQPressureSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQPressureSensor"})
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPressureSensor"})
}

func (ptr *QPressureSensor) DestroyQPressureSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQPressureSensorDefault"})
}

type QProximityFilter struct {
	QSensorFilter
}

type QProximityFilter_ITF interface {
	QSensorFilter_ITF
	QProximityFilter_PTR() *QProximityFilter
}

func (ptr *QProximityFilter) QProximityFilter_PTR() *QProximityFilter {
	return ptr
}

func (ptr *QProximityFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QProximityFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQProximityFilter(ptr QProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityFilter_PTR().Pointer()
	}
	return nil
}

func (n *QProximityFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QProximityFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQProximityFilterFromPointer(ptr unsafe.Pointer) (n *QProximityFilter) {
	n = new(QProximityFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QProximityFilter")
	return
}

func (ptr *QProximityFilter) DestroyQProximityFilter() {
}

func (ptr *QProximityFilter) ConnectFilter(f func(reading *QProximityReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QProximityFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QProximityFilter) Filter(reading QProximityReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QProximityReading struct {
	QSensorReading
}

type QProximityReading_ITF interface {
	QSensorReading_ITF
	QProximityReading_PTR() *QProximityReading
}

func (ptr *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QProximityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func (n *QProximityReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QProximityReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) (n *QProximityReading) {
	n = new(QProximityReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QProximityReading")
	return
}

func (ptr *QProximityReading) DestroyQProximityReading() {
}

func (ptr *QProximityReading) Close() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"}).(bool)
}

func (ptr *QProximityReading) SetClose(close bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetClose", close})
}

type QProximitySensor struct {
	QSensor
}

type QProximitySensor_ITF interface {
	QSensor_ITF
	QProximitySensor_PTR() *QProximitySensor
}

func (ptr *QProximitySensor) QProximitySensor_PTR() *QProximitySensor {
	return ptr
}

func (ptr *QProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQProximitySensor(ptr QProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *QProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QProximitySensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQProximitySensorFromPointer(ptr unsafe.Pointer) (n *QProximitySensor) {
	n = new(QProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QProximitySensor")
	return
}
func NewQProximitySensor(parent core.QObject_ITF) *QProximitySensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQProximitySensor", "", parent}).(*QProximitySensor)
}

func (ptr *QProximitySensor) Reading() *QProximityReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QProximityReading)
}

func (ptr *QProximitySensor) ConnectDestroyQProximitySensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQProximitySensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QProximitySensor) DisconnectDestroyQProximitySensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQProximitySensor"})
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQProximitySensor"})
}

func (ptr *QProximitySensor) DestroyQProximitySensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQProximitySensorDefault"})
}

type QRotationFilter struct {
	QSensorFilter
}

type QRotationFilter_ITF interface {
	QSensorFilter_ITF
	QRotationFilter_PTR() *QRotationFilter
}

func (ptr *QRotationFilter) QRotationFilter_PTR() *QRotationFilter {
	return ptr
}

func (ptr *QRotationFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QRotationFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQRotationFilter(ptr QRotationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationFilter_PTR().Pointer()
	}
	return nil
}

func (n *QRotationFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRotationFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQRotationFilterFromPointer(ptr unsafe.Pointer) (n *QRotationFilter) {
	n = new(QRotationFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QRotationFilter")
	return
}

func (ptr *QRotationFilter) DestroyQRotationFilter() {
}

func (ptr *QRotationFilter) ConnectFilter(f func(reading *QRotationReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRotationFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QRotationFilter) Filter(reading QRotationReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QRotationReading struct {
	QSensorReading
}

type QRotationReading_ITF interface {
	QSensorReading_ITF
	QRotationReading_PTR() *QRotationReading
}

func (ptr *QRotationReading) QRotationReading_PTR() *QRotationReading {
	return ptr
}

func (ptr *QRotationReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QRotationReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQRotationReading(ptr QRotationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationReading_PTR().Pointer()
	}
	return nil
}

func (n *QRotationReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRotationReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQRotationReadingFromPointer(ptr unsafe.Pointer) (n *QRotationReading) {
	n = new(QRotationReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QRotationReading")
	return
}

func (ptr *QRotationReading) DestroyQRotationReading() {
}

func (ptr *QRotationReading) SetFromEuler(x float64, y float64, z float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFromEuler", x, y, z})
}

func (ptr *QRotationReading) X() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "X"}).(float64)
}

func (ptr *QRotationReading) Y() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Y"}).(float64)
}

func (ptr *QRotationReading) Z() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Z"}).(float64)
}

type QRotationSensor struct {
	QSensor
}

type QRotationSensor_ITF interface {
	QSensor_ITF
	QRotationSensor_PTR() *QRotationSensor
}

func (ptr *QRotationSensor) QRotationSensor_PTR() *QRotationSensor {
	return ptr
}

func (ptr *QRotationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QRotationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQRotationSensor(ptr QRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationSensor_PTR().Pointer()
	}
	return nil
}

func (n *QRotationSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QRotationSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQRotationSensorFromPointer(ptr unsafe.Pointer) (n *QRotationSensor) {
	n = new(QRotationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QRotationSensor")
	return
}
func NewQRotationSensor(parent core.QObject_ITF) *QRotationSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQRotationSensor", "", parent}).(*QRotationSensor)
}

func (ptr *QRotationSensor) HasZ() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasZ"}).(bool)
}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasZChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasZChanged"})
}

func (ptr *QRotationSensor) HasZChanged(hasZ bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasZChanged", hasZ})
}

func (ptr *QRotationSensor) Reading() *QRotationReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QRotationReading)
}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHasZ", hasZ})
}

func (ptr *QRotationSensor) ConnectDestroyQRotationSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQRotationSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QRotationSensor) DisconnectDestroyQRotationSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQRotationSensor"})
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQRotationSensor"})
}

func (ptr *QRotationSensor) DestroyQRotationSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQRotationSensorDefault"})
}

type QSensor struct {
	core.QObject
}

type QSensor_ITF interface {
	core.QObject_ITF
	QSensor_PTR() *QSensor
}

func (ptr *QSensor) QSensor_PTR() *QSensor {
	return ptr
}

func (ptr *QSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensor(ptr QSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (n *QSensor) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSensor) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSensorFromPointer(ptr unsafe.Pointer) (n *QSensor) {
	n = new(QSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensor")
	return
}

//go:generate stringer -type=QSensor__Feature
//QSensor::Feature
type QSensor__Feature int64

const (
	QSensor__Buffering                 QSensor__Feature = QSensor__Feature(0)
	QSensor__AlwaysOn                  QSensor__Feature = QSensor__Feature(1)
	QSensor__GeoValues                 QSensor__Feature = QSensor__Feature(2)
	QSensor__FieldOfView               QSensor__Feature = QSensor__Feature(3)
	QSensor__AccelerationMode          QSensor__Feature = QSensor__Feature(4)
	QSensor__SkipDuplicates            QSensor__Feature = QSensor__Feature(5)
	QSensor__AxesOrientation           QSensor__Feature = QSensor__Feature(6)
	QSensor__PressureSensorTemperature QSensor__Feature = QSensor__Feature(7)
	QSensor__Reserved                  QSensor__Feature = QSensor__Feature(257)
)

//go:generate stringer -type=QSensor__AxesOrientationMode
//QSensor::AxesOrientationMode
type QSensor__AxesOrientationMode int64

const (
	QSensor__FixedOrientation     QSensor__AxesOrientationMode = QSensor__AxesOrientationMode(0)
	QSensor__AutomaticOrientation QSensor__AxesOrientationMode = QSensor__AxesOrientationMode(1)
	QSensor__UserOrientation      QSensor__AxesOrientationMode = QSensor__AxesOrientationMode(2)
)

func NewQSensor(ty core.QByteArray_ITF, parent core.QObject_ITF) *QSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQSensor", "", ty, parent}).(*QSensor)
}

func (ptr *QSensor) ConnectActiveChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectActiveChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveChanged"})
}

func (ptr *QSensor) ActiveChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveChanged"})
}

func (ptr *QSensor) AddFilter(filter QSensorFilter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddFilter", filter})
}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAlwaysOnChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAlwaysOnChanged"})
}

func (ptr *QSensor) AlwaysOnChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AlwaysOnChanged"})
}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAvailableSensorsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAvailableSensorsChanged"})
}

func (ptr *QSensor) AvailableSensorsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableSensorsChanged"})
}

func (ptr *QSensor) AxesOrientationMode() QSensor__AxesOrientationMode {

	return QSensor__AxesOrientationMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxesOrientationMode"}).(float64))
}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxesOrientationModeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxesOrientationModeChanged"})
}

func (ptr *QSensor) AxesOrientationModeChanged(axesOrientationMode QSensor__AxesOrientationMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxesOrientationModeChanged", axesOrientationMode})
}

func (ptr *QSensor) BufferSize() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BufferSize"}).(float64))
}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBufferSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBufferSizeChanged"})
}

func (ptr *QSensor) BufferSizeChanged(bufferSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BufferSizeChanged", bufferSize})
}

func (ptr *QSensor) ConnectBusyChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBusyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectBusyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBusyChanged"})
}

func (ptr *QSensor) BusyChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BusyChanged"})
}

func (ptr *QSensor) ConnectToBackend() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToBackend"}).(bool)
}

func (ptr *QSensor) CurrentOrientation() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentOrientation"}).(float64))
}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCurrentOrientationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCurrentOrientationChanged"})
}

func (ptr *QSensor) CurrentOrientationChanged(currentOrientation int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentOrientationChanged", currentOrientation})
}

func (ptr *QSensor) DataRate() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataRate"}).(float64))
}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataRateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectDataRateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataRateChanged"})
}

func (ptr *QSensor) DataRateChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataRateChanged"})
}

func QSensor_DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensor_DefaultSensorForType", "", ty}).(*core.QByteArray)
}

func (ptr *QSensor) DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensor_DefaultSensorForType", "", ty}).(*core.QByteArray)
}

func (ptr *QSensor) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QSensor) EfficientBufferSize() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EfficientBufferSize"}).(float64))
}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEfficientBufferSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEfficientBufferSizeChanged"})
}

func (ptr *QSensor) EfficientBufferSizeChanged(efficientBufferSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EfficientBufferSizeChanged", efficientBufferSize})
}

func (ptr *QSensor) Error() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QSensor) Filters() []*QSensorFilter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filters"}).([]*QSensorFilter)
}

func (ptr *QSensor) Identifier() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Identifier"}).(*core.QByteArray)
}

func (ptr *QSensor) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QSensor) IsAlwaysOn() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAlwaysOn"}).(bool)
}

func (ptr *QSensor) IsBusy() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBusy"}).(bool)
}

func (ptr *QSensor) IsConnectedToBackend() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsConnectedToBackend"}).(bool)
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFeatureSupported", feature}).(bool)
}

func (ptr *QSensor) MaxBufferSize() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxBufferSize"}).(float64))
}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMaxBufferSizeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMaxBufferSizeChanged"})
}

func (ptr *QSensor) MaxBufferSizeChanged(maxBufferSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxBufferSizeChanged", maxBufferSize})
}

func (ptr *QSensor) OutputRange() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OutputRange"}).(float64))
}

func (ptr *QSensor) Reading() *QSensorReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QSensorReading)
}

func (ptr *QSensor) ConnectReadingChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadingChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectReadingChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadingChanged"})
}

func (ptr *QSensor) ReadingChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadingChanged"})
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveFilter", filter})
}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSensorError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectSensorError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSensorError"})
}

func (ptr *QSensor) SensorError(error int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SensorError", error})
}

func QSensor_SensorTypes() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensor_SensorTypes", ""}).([]*core.QByteArray)
}

func (ptr *QSensor) SensorTypes() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensor_SensorTypes", ""}).([]*core.QByteArray)
}

func QSensor_SensorsForType(ty core.QByteArray_ITF) []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensor_SensorsForType", "", ty}).([]*core.QByteArray)
}

func (ptr *QSensor) SensorsForType(ty core.QByteArray_ITF) []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensor_SensorsForType", "", ty}).([]*core.QByteArray)
}

func (ptr *QSensor) SetActive(active bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActive", active})
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAlwaysOn", alwaysOn})
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAxesOrientationMode", axesOrientationMode})
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBufferSize", bufferSize})
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentOrientation", currentOrientation})
}

func (ptr *QSensor) SetDataRate(rate int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataRate", rate})
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEfficientBufferSize", efficientBufferSize})
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIdentifier", identifier})
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxBufferSize", maxBufferSize})
}

func (ptr *QSensor) SetOutputRange(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOutputRange", index})
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSkipDuplicates", skipDuplicates})
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUserOrientation", userOrientation})
}

func (ptr *QSensor) SkipDuplicates() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SkipDuplicates"}).(bool)
}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSkipDuplicatesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSkipDuplicatesChanged"})
}

func (ptr *QSensor) SkipDuplicatesChanged(skipDuplicates bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SkipDuplicatesChanged", skipDuplicates})
}

func (ptr *QSensor) ConnectStart(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QSensor) Start() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start"}).(bool)
}

func (ptr *QSensor) StartDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDefault"}).(bool)
}

func (ptr *QSensor) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QSensor) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QSensor) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func (ptr *QSensor) Type() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(*core.QByteArray)
}

func (ptr *QSensor) UserOrientation() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UserOrientation"}).(float64))
}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUserOrientationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUserOrientationChanged"})
}

func (ptr *QSensor) UserOrientationChanged(userOrientation int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UserOrientationChanged", userOrientation})
}

func (ptr *QSensor) ConnectDestroyQSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensor) DisconnectDestroyQSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSensor"})
}

func (ptr *QSensor) DestroyQSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensor"})
}

func (ptr *QSensor) DestroyQSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorDefault"})
}

func (ptr *QSensor) __filters_atList(i int) *QSensorFilter {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__filters_atList", i}).(*QSensorFilter)
}

func (ptr *QSensor) __filters_setList(i QSensorFilter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__filters_setList", i})
}

func (ptr *QSensor) __filters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__filters_newList"}).(unsafe.Pointer)
}

func (ptr *QSensor) __sensorTypes_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sensorTypes_atList", i}).(*core.QByteArray)
}

func (ptr *QSensor) __sensorTypes_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sensorTypes_setList", i})
}

func (ptr *QSensor) __sensorTypes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sensorTypes_newList"}).(unsafe.Pointer)
}

func (ptr *QSensor) __sensorsForType_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sensorsForType_atList", i}).(*core.QByteArray)
}

func (ptr *QSensor) __sensorsForType_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sensorsForType_setList", i})
}

func (ptr *QSensor) __sensorsForType_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sensorsForType_newList"}).(unsafe.Pointer)
}

func (ptr *QSensor) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSensor) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSensor) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSensor) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSensor) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSensor) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSensor) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSensor) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSensor) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSensor) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSensor) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSensor) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSensor) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSensor) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSensor) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSensor) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSensor) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSensor) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSensorBackend struct {
	core.QObject
}

type QSensorBackend_ITF interface {
	core.QObject_ITF
	QSensorBackend_PTR() *QSensorBackend
}

func (ptr *QSensorBackend) QSensorBackend_PTR() *QSensorBackend {
	return ptr
}

func (ptr *QSensorBackend) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorBackend) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorBackend(ptr QSensorBackend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (n *QSensorBackend) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSensorBackend) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSensorBackendFromPointer(ptr unsafe.Pointer) (n *QSensorBackend) {
	n = new(QSensorBackend)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorBackend")
	return
}

func (ptr *QSensorBackend) DestroyQSensorBackend() {
}

func (ptr *QSensorBackend) AddDataRate(min float64, max float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddDataRate", min, max})
}

func (ptr *QSensorBackend) AddOutputRange(min float64, max float64, accuracy float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddOutputRange", min, max, accuracy})
}

func (ptr *QSensorBackend) ConnectIsFeatureSupported(f func(feature QSensor__Feature) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsFeatureSupported", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorBackend) DisconnectIsFeatureSupported() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsFeatureSupported"})
}

func (ptr *QSensorBackend) IsFeatureSupported(feature QSensor__Feature) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFeatureSupported", feature}).(bool)
}

func (ptr *QSensorBackend) IsFeatureSupportedDefault(feature QSensor__Feature) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFeatureSupportedDefault", feature}).(bool)
}

func (ptr *QSensorBackend) NewReadingAvailable() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewReadingAvailable"})
}

func (ptr *QSensorBackend) Reading() *QSensorReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QSensorReading)
}

func (ptr *QSensorBackend) Sensor() *QSensor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Sensor"}).(*QSensor)
}

func (ptr *QSensorBackend) SensorBusy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SensorBusy"})
}

func (ptr *QSensorBackend) SensorError(error int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SensorError", error})
}

func (ptr *QSensorBackend) SensorStopped() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SensorStopped"})
}

func (ptr *QSensorBackend) SetDataRates(otherSensor QSensor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataRates", otherSensor})
}

func (ptr *QSensorBackend) SetDescription(description string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDescription", description})
}

func (ptr *QSensorBackend) ConnectStart(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorBackend) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QSensorBackend) Start() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start"})
}

func (ptr *QSensorBackend) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorBackend) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QSensorBackend) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QSensorBackend) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSensorBackend) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSensorBackend) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorBackend) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSensorBackend) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSensorBackend) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorBackend) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSensorBackend) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSensorBackend) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorBackend) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSensorBackend) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSensorBackend) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSensorBackend) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSensorBackend) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSensorBackend) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSensorBackend) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSensorBackend) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSensorBackend) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSensorBackend) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSensorBackend) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSensorBackend) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSensorBackendFactory struct {
	internal.Internal
}

type QSensorBackendFactory_ITF interface {
	QSensorBackendFactory_PTR() *QSensorBackendFactory
}

func (ptr *QSensorBackendFactory) QSensorBackendFactory_PTR() *QSensorBackendFactory {
	return ptr
}

func (ptr *QSensorBackendFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSensorBackendFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSensorBackendFactory(ptr QSensorBackendFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackendFactory_PTR().Pointer()
	}
	return nil
}

func (n *QSensorBackendFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSensorBackendFactoryFromPointer(ptr unsafe.Pointer) (n *QSensorBackendFactory) {
	n = new(QSensorBackendFactory)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorBackendFactory")
	return
}

func (ptr *QSensorBackendFactory) DestroyQSensorBackendFactory() {
}

func (ptr *QSensorBackendFactory) ConnectCreateBackend(f func(sensor *QSensor) *QSensorBackend) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateBackend", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorBackendFactory) DisconnectCreateBackend() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateBackend"})
}

func (ptr *QSensorBackendFactory) CreateBackend(sensor QSensor_ITF) *QSensorBackend {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateBackend", sensor}).(*QSensorBackend)
}

type QSensorChangesInterface struct {
	internal.Internal
}

type QSensorChangesInterface_ITF interface {
	QSensorChangesInterface_PTR() *QSensorChangesInterface
}

func (ptr *QSensorChangesInterface) QSensorChangesInterface_PTR() *QSensorChangesInterface {
	return ptr
}

func (ptr *QSensorChangesInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSensorChangesInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSensorChangesInterface(ptr QSensorChangesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorChangesInterface_PTR().Pointer()
	}
	return nil
}

func (n *QSensorChangesInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSensorChangesInterfaceFromPointer(ptr unsafe.Pointer) (n *QSensorChangesInterface) {
	n = new(QSensorChangesInterface)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorChangesInterface")
	return
}

func (ptr *QSensorChangesInterface) DestroyQSensorChangesInterface() {
}

func (ptr *QSensorChangesInterface) ConnectSensorsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSensorsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorChangesInterface) DisconnectSensorsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSensorsChanged"})
}

func (ptr *QSensorChangesInterface) SensorsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SensorsChanged"})
}

type QSensorFilter struct {
	internal.Internal
}

type QSensorFilter_ITF interface {
	QSensorFilter_PTR() *QSensorFilter
}

func (ptr *QSensorFilter) QSensorFilter_PTR() *QSensorFilter {
	return ptr
}

func (ptr *QSensorFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSensorFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSensorFilter(ptr QSensorFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (n *QSensorFilter) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSensorFilterFromPointer(ptr unsafe.Pointer) (n *QSensorFilter) {
	n = new(QSensorFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorFilter")
	return
}
func (ptr *QSensorFilter) ConnectFilter(f func(reading *QSensorReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QSensorFilter) Filter(reading QSensorReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

func (ptr *QSensorFilter) ConnectDestroyQSensorFilter(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSensorFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorFilter) DisconnectDestroyQSensorFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSensorFilter"})
}

func (ptr *QSensorFilter) DestroyQSensorFilter() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorFilter"})
}

func (ptr *QSensorFilter) DestroyQSensorFilterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorFilterDefault"})
}

type QSensorGesture struct {
	core.QObject
}

type QSensorGesture_ITF interface {
	core.QObject_ITF
	QSensorGesture_PTR() *QSensorGesture
}

func (ptr *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return ptr
}

func (ptr *QSensorGesture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorGesture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorGesture(ptr QSensorGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesture_PTR().Pointer()
	}
	return nil
}

func (n *QSensorGesture) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSensorGesture) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSensorGestureFromPointer(ptr unsafe.Pointer) (n *QSensorGesture) {
	n = new(QSensorGesture)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorGesture")
	return
}
func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQSensorGesture", "", ids, parent}).(*QSensorGesture)
}

func (ptr *QSensorGesture) ConnectDetected(f func(vqs string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDetected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGesture) DisconnectDetected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDetected"})
}

func (ptr *QSensorGesture) Detected(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Detected", vqs})
}

func (ptr *QSensorGesture) GestureSignals() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GestureSignals"}).([]string)
}

func (ptr *QSensorGesture) InvalidIds() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InvalidIds"}).([]string)
}

func (ptr *QSensorGesture) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QSensorGesture) StartDetection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDetection"})
}

func (ptr *QSensorGesture) StopDetection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDetection"})
}

func (ptr *QSensorGesture) ValidIds() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValidIds"}).([]string)
}

func (ptr *QSensorGesture) ConnectDestroyQSensorGesture(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSensorGesture", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGesture) DisconnectDestroyQSensorGesture() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSensorGesture"})
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGesture"})
}

func (ptr *QSensorGesture) DestroyQSensorGestureDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGestureDefault"})
}

func (ptr *QSensorGesture) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSensorGesture) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSensorGesture) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGesture) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSensorGesture) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSensorGesture) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGesture) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSensorGesture) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSensorGesture) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGesture) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSensorGesture) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSensorGesture) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSensorGesture) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSensorGesture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSensorGesture) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSensorGesture) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSensorGesture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSensorGesture) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSensorGesture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSensorGesture) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSensorGesture) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSensorGestureManager struct {
	core.QObject
}

type QSensorGestureManager_ITF interface {
	core.QObject_ITF
	QSensorGestureManager_PTR() *QSensorGestureManager
}

func (ptr *QSensorGestureManager) QSensorGestureManager_PTR() *QSensorGestureManager {
	return ptr
}

func (ptr *QSensorGestureManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorGestureManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorGestureManager(ptr QSensorGestureManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureManager_PTR().Pointer()
	}
	return nil
}

func (n *QSensorGestureManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSensorGestureManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSensorGestureManagerFromPointer(ptr unsafe.Pointer) (n *QSensorGestureManager) {
	n = new(QSensorGestureManager)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorGestureManager")
	return
}
func NewQSensorGestureManager(parent core.QObject_ITF) *QSensorGestureManager {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQSensorGestureManager", "", parent}).(*QSensorGestureManager)
}

func (ptr *QSensorGestureManager) GestureIds() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GestureIds"}).([]string)
}

func (ptr *QSensorGestureManager) ConnectNewSensorGestureAvailable(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewSensorGestureAvailable", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureManager) DisconnectNewSensorGestureAvailable() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewSensorGestureAvailable"})
}

func (ptr *QSensorGestureManager) NewSensorGestureAvailable() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewSensorGestureAvailable"})
}

func (ptr *QSensorGestureManager) RecognizerSignals(gestureId string) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecognizerSignals", gestureId}).([]string)
}

func (ptr *QSensorGestureManager) RegisterSensorGestureRecognizer(recognizer QSensorGestureRecognizer_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterSensorGestureRecognizer", recognizer}).(bool)
}

func QSensorGestureManager_SensorGestureRecognizer(id string) *QSensorGestureRecognizer {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorGestureManager_SensorGestureRecognizer", "", id}).(*QSensorGestureRecognizer)
}

func (ptr *QSensorGestureManager) SensorGestureRecognizer(id string) *QSensorGestureRecognizer {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorGestureManager_SensorGestureRecognizer", "", id}).(*QSensorGestureRecognizer)
}

func (ptr *QSensorGestureManager) ConnectDestroyQSensorGestureManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSensorGestureManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureManager) DisconnectDestroyQSensorGestureManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSensorGestureManager"})
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGestureManager"})
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGestureManagerDefault"})
}

func (ptr *QSensorGestureManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSensorGestureManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSensorGestureManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSensorGestureManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSensorGestureManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSensorGestureManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSensorGestureManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSensorGestureManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSensorGestureManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSensorGestureManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSensorGestureManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSensorGestureManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSensorGestureManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSensorGestureManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSensorGestureManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSensorGestureManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSensorGestureManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSensorGesturePluginInterface struct {
	internal.Internal
}

type QSensorGesturePluginInterface_ITF interface {
	QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface
}

func (ptr *QSensorGesturePluginInterface) QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface {
	return ptr
}

func (ptr *QSensorGesturePluginInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSensorGesturePluginInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSensorGesturePluginInterface(ptr QSensorGesturePluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesturePluginInterface_PTR().Pointer()
	}
	return nil
}

func (n *QSensorGesturePluginInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSensorGesturePluginInterfaceFromPointer(ptr unsafe.Pointer) (n *QSensorGesturePluginInterface) {
	n = new(QSensorGesturePluginInterface)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorGesturePluginInterface")
	return
}
func NewQSensorGesturePluginInterface() *QSensorGesturePluginInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQSensorGesturePluginInterface", ""}).(*QSensorGesturePluginInterface)
}

func (ptr *QSensorGesturePluginInterface) ConnectCreateRecognizers(f func() []*QSensorGestureRecognizer) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateRecognizers", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGesturePluginInterface) DisconnectCreateRecognizers() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateRecognizers"})
}

func (ptr *QSensorGesturePluginInterface) CreateRecognizers() []*QSensorGestureRecognizer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateRecognizers"}).([]*QSensorGestureRecognizer)
}

func (ptr *QSensorGesturePluginInterface) ConnectName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGesturePluginInterface) DisconnectName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectName"})
}

func (ptr *QSensorGesturePluginInterface) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QSensorGesturePluginInterface) ConnectSupportedIds(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportedIds", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGesturePluginInterface) DisconnectSupportedIds() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportedIds"})
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedIds"}).([]string)
}

func (ptr *QSensorGesturePluginInterface) ConnectDestroyQSensorGesturePluginInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSensorGesturePluginInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGesturePluginInterface) DisconnectDestroyQSensorGesturePluginInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSensorGesturePluginInterface"})
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGesturePluginInterface"})
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGesturePluginInterfaceDefault"})
}

func (ptr *QSensorGesturePluginInterface) __createRecognizers_atList(i int) *QSensorGestureRecognizer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRecognizers_atList", i}).(*QSensorGestureRecognizer)
}

func (ptr *QSensorGesturePluginInterface) __createRecognizers_setList(i QSensorGestureRecognizer_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRecognizers_setList", i})
}

func (ptr *QSensorGesturePluginInterface) __createRecognizers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRecognizers_newList"}).(unsafe.Pointer)
}

type QSensorGestureRecognizer struct {
	core.QObject
}

type QSensorGestureRecognizer_ITF interface {
	core.QObject_ITF
	QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer
}

func (ptr *QSensorGestureRecognizer) QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer {
	return ptr
}

func (ptr *QSensorGestureRecognizer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorGestureRecognizer(ptr QSensorGestureRecognizer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureRecognizer_PTR().Pointer()
	}
	return nil
}

func (n *QSensorGestureRecognizer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSensorGestureRecognizer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSensorGestureRecognizerFromPointer(ptr unsafe.Pointer) (n *QSensorGestureRecognizer) {
	n = new(QSensorGestureRecognizer)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorGestureRecognizer")
	return
}
func NewQSensorGestureRecognizer(parent core.QObject_ITF) *QSensorGestureRecognizer {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQSensorGestureRecognizer", "", parent}).(*QSensorGestureRecognizer)
}

func (ptr *QSensorGestureRecognizer) ConnectCreate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectCreate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreate"})
}

func (ptr *QSensorGestureRecognizer) Create() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Create"})
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateBackend"})
}

func (ptr *QSensorGestureRecognizer) ConnectDetected(f func(vqs string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDetected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectDetected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDetected"})
}

func (ptr *QSensorGestureRecognizer) Detected(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Detected", vqs})
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GestureSignals"}).([]string)
}

func (ptr *QSensorGestureRecognizer) ConnectId(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectId", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectId() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectId"})
}

func (ptr *QSensorGestureRecognizer) Id() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Id"}).(string)
}

func (ptr *QSensorGestureRecognizer) ConnectIsActive(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsActive", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectIsActive() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsActive"})
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QSensorGestureRecognizer) ConnectStart(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QSensorGestureRecognizer) Start() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start"}).(bool)
}

func (ptr *QSensorGestureRecognizer) StartBackend() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartBackend"})
}

func (ptr *QSensorGestureRecognizer) ConnectStop(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QSensorGestureRecognizer) Stop() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"}).(bool)
}

func (ptr *QSensorGestureRecognizer) StopBackend() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopBackend"})
}

func (ptr *QSensorGestureRecognizer) ConnectDestroyQSensorGestureRecognizer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSensorGestureRecognizer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorGestureRecognizer) DisconnectDestroyQSensorGestureRecognizer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSensorGestureRecognizer"})
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGestureRecognizer"})
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSensorGestureRecognizerDefault"})
}

func (ptr *QSensorGestureRecognizer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSensorGestureRecognizer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSensorGestureRecognizer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureRecognizer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSensorGestureRecognizer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSensorGestureRecognizer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureRecognizer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSensorGestureRecognizer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSensorGestureRecognizer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureRecognizer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSensorGestureRecognizer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSensorGestureRecognizer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSensorGestureRecognizer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSensorGestureRecognizer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSensorGestureRecognizer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSensorGestureRecognizer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSensorGestureRecognizer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSensorGestureRecognizer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSensorGestureRecognizer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSensorGestureRecognizer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSensorGestureRecognizer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSensorManager struct {
	internal.Internal
}

type QSensorManager_ITF interface {
	QSensorManager_PTR() *QSensorManager
}

func (ptr *QSensorManager) QSensorManager_PTR() *QSensorManager {
	return ptr
}

func (ptr *QSensorManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSensorManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSensorManager(ptr QSensorManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorManager_PTR().Pointer()
	}
	return nil
}

func (n *QSensorManager) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSensorManagerFromPointer(ptr unsafe.Pointer) (n *QSensorManager) {
	n = new(QSensorManager)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorManager")
	return
}

func (ptr *QSensorManager) DestroyQSensorManager() {
}

func QSensorManager_CreateBackend(sensor QSensor_ITF) *QSensorBackend {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_CreateBackend", "", sensor}).(*QSensorBackend)
}

func (ptr *QSensorManager) CreateBackend(sensor QSensor_ITF) *QSensorBackend {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_CreateBackend", "", sensor}).(*QSensorBackend)
}

func QSensorManager_IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_IsBackendRegistered", "", ty, identifier}).(bool)
}

func (ptr *QSensorManager) IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_IsBackendRegistered", "", ty, identifier}).(bool)
}

func QSensorManager_RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_RegisterBackend", "", ty, identifier, factory})
}

func (ptr *QSensorManager) RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_RegisterBackend", "", ty, identifier, factory})
}

func QSensorManager_SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_SetDefaultBackend", "", ty, identifier})
}

func (ptr *QSensorManager) SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_SetDefaultBackend", "", ty, identifier})
}

func QSensorManager_UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_UnregisterBackend", "", ty, identifier})
}

func (ptr *QSensorManager) UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sensors.QSensorManager_UnregisterBackend", "", ty, identifier})
}

type QSensorPluginInterface struct {
	internal.Internal
}

type QSensorPluginInterface_ITF interface {
	QSensorPluginInterface_PTR() *QSensorPluginInterface
}

func (ptr *QSensorPluginInterface) QSensorPluginInterface_PTR() *QSensorPluginInterface {
	return ptr
}

func (ptr *QSensorPluginInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSensorPluginInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSensorPluginInterface(ptr QSensorPluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorPluginInterface_PTR().Pointer()
	}
	return nil
}

func (n *QSensorPluginInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSensorPluginInterfaceFromPointer(ptr unsafe.Pointer) (n *QSensorPluginInterface) {
	n = new(QSensorPluginInterface)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorPluginInterface")
	return
}

func (ptr *QSensorPluginInterface) DestroyQSensorPluginInterface() {
}

func (ptr *QSensorPluginInterface) ConnectRegisterSensors(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRegisterSensors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSensorPluginInterface) DisconnectRegisterSensors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRegisterSensors"})
}

func (ptr *QSensorPluginInterface) RegisterSensors() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterSensors"})
}

type QSensorReading struct {
	core.QObject
}

type QSensorReading_ITF interface {
	core.QObject_ITF
	QSensorReading_PTR() *QSensorReading
}

func (ptr *QSensorReading) QSensorReading_PTR() *QSensorReading {
	return ptr
}

func (ptr *QSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorReading(ptr QSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSensorReading) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSensorReadingFromPointer(ptr unsafe.Pointer) (n *QSensorReading) {
	n = new(QSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QSensorReading")
	return
}

func (ptr *QSensorReading) DestroyQSensorReading() {
}

func (ptr *QSensorReading) SetTimestamp(timestamp uint64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimestamp", timestamp})
}

func (ptr *QSensorReading) Timestamp() uint64 {

	return uint64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timestamp"}).(float64))
}

func (ptr *QSensorReading) Value(index int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", index}).(*core.QVariant)
}

func (ptr *QSensorReading) ValueCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueCount"}).(float64))
}

func (ptr *QSensorReading) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSensorReading) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSensorReading) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorReading) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSensorReading) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSensorReading) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorReading) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSensorReading) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSensorReading) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSensorReading) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSensorReading) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSensorReading) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSensorReading) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSensorReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSensorReading) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSensorReading) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSensorReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSensorReading) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSensorReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSensorReading) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSensorReading) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QTapFilter struct {
	QSensorFilter
}

type QTapFilter_ITF interface {
	QSensorFilter_ITF
	QTapFilter_PTR() *QTapFilter
}

func (ptr *QTapFilter) QTapFilter_PTR() *QTapFilter {
	return ptr
}

func (ptr *QTapFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QTapFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQTapFilter(ptr QTapFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapFilter_PTR().Pointer()
	}
	return nil
}

func (n *QTapFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTapFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQTapFilterFromPointer(ptr unsafe.Pointer) (n *QTapFilter) {
	n = new(QTapFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QTapFilter")
	return
}

func (ptr *QTapFilter) DestroyQTapFilter() {
}

func (ptr *QTapFilter) ConnectFilter(f func(reading *QTapReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTapFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QTapFilter) Filter(reading QTapReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QTapReading struct {
	QSensorReading
}

type QTapReading_ITF interface {
	QSensorReading_ITF
	QTapReading_PTR() *QTapReading
}

func (ptr *QTapReading) QTapReading_PTR() *QTapReading {
	return ptr
}

func (ptr *QTapReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QTapReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQTapReading(ptr QTapReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapReading_PTR().Pointer()
	}
	return nil
}

func (n *QTapReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTapReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQTapReadingFromPointer(ptr unsafe.Pointer) (n *QTapReading) {
	n = new(QTapReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QTapReading")
	return
}

func (ptr *QTapReading) DestroyQTapReading() {
}

//go:generate stringer -type=QTapReading__TapDirection
//QTapReading::TapDirection
type QTapReading__TapDirection int64

const (
	QTapReading__Undefined QTapReading__TapDirection = QTapReading__TapDirection(0)
	QTapReading__X         QTapReading__TapDirection = QTapReading__TapDirection(0x0001)
	QTapReading__Y         QTapReading__TapDirection = QTapReading__TapDirection(0x0002)
	QTapReading__Z         QTapReading__TapDirection = QTapReading__TapDirection(0x0004)
	QTapReading__X_Pos     QTapReading__TapDirection = QTapReading__TapDirection(0x0011)
	QTapReading__Y_Pos     QTapReading__TapDirection = QTapReading__TapDirection(0x0022)
	QTapReading__Z_Pos     QTapReading__TapDirection = QTapReading__TapDirection(0x0044)
	QTapReading__X_Neg     QTapReading__TapDirection = QTapReading__TapDirection(0x0101)
	QTapReading__Y_Neg     QTapReading__TapDirection = QTapReading__TapDirection(0x0202)
	QTapReading__Z_Neg     QTapReading__TapDirection = QTapReading__TapDirection(0x0404)
	QTapReading__X_Both    QTapReading__TapDirection = QTapReading__TapDirection(0x0111)
	QTapReading__Y_Both    QTapReading__TapDirection = QTapReading__TapDirection(0x0222)
	QTapReading__Z_Both    QTapReading__TapDirection = QTapReading__TapDirection(0x0444)
)

func (ptr *QTapReading) IsDoubleTap() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDoubleTap"}).(bool)
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDoubleTap", doubleTap})
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTapDirection", tapDirection})
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {

	return QTapReading__TapDirection(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TapDirection"}).(float64))
}

type QTapSensor struct {
	QSensor
}

type QTapSensor_ITF interface {
	QSensor_ITF
	QTapSensor_PTR() *QTapSensor
}

func (ptr *QTapSensor) QTapSensor_PTR() *QTapSensor {
	return ptr
}

func (ptr *QTapSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QTapSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQTapSensor(ptr QTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapSensor_PTR().Pointer()
	}
	return nil
}

func (n *QTapSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTapSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQTapSensorFromPointer(ptr unsafe.Pointer) (n *QTapSensor) {
	n = new(QTapSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QTapSensor")
	return
}
func NewQTapSensor(parent core.QObject_ITF) *QTapSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQTapSensor", "", parent}).(*QTapSensor)
}

func (ptr *QTapSensor) Reading() *QTapReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QTapReading)
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReturnDoubleTapEvents"}).(bool)
}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReturnDoubleTapEventsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReturnDoubleTapEventsChanged"})
}

func (ptr *QTapSensor) ReturnDoubleTapEventsChanged(returnDoubleTapEvents bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReturnDoubleTapEventsChanged", returnDoubleTapEvents})
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReturnDoubleTapEvents", returnDoubleTapEvents})
}

func (ptr *QTapSensor) ConnectDestroyQTapSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQTapSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTapSensor) DisconnectDestroyQTapSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQTapSensor"})
}

func (ptr *QTapSensor) DestroyQTapSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTapSensor"})
}

func (ptr *QTapSensor) DestroyQTapSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTapSensorDefault"})
}

type QTiltFilter struct {
	QSensorFilter
}

type QTiltFilter_ITF interface {
	QSensorFilter_ITF
	QTiltFilter_PTR() *QTiltFilter
}

func (ptr *QTiltFilter) QTiltFilter_PTR() *QTiltFilter {
	return ptr
}

func (ptr *QTiltFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QTiltFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQTiltFilter(ptr QTiltFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltFilter_PTR().Pointer()
	}
	return nil
}

func (n *QTiltFilter) InitFromInternal(ptr uintptr, name string) {
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTiltFilter) ClassNameInternalF() string {
	return n.QSensorFilter_PTR().ClassNameInternalF()
}

func NewQTiltFilterFromPointer(ptr unsafe.Pointer) (n *QTiltFilter) {
	n = new(QTiltFilter)
	n.InitFromInternal(uintptr(ptr), "sensors.QTiltFilter")
	return
}

func (ptr *QTiltFilter) DestroyQTiltFilter() {
}

func (ptr *QTiltFilter) ConnectFilter(f func(reading *QTiltReading) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTiltFilter) DisconnectFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFilter"})
}

func (ptr *QTiltFilter) Filter(reading QTiltReading_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter", reading}).(bool)
}

type QTiltReading struct {
	QSensorReading
}

type QTiltReading_ITF interface {
	QSensorReading_ITF
	QTiltReading_PTR() *QTiltReading
}

func (ptr *QTiltReading) QTiltReading_PTR() *QTiltReading {
	return ptr
}

func (ptr *QTiltReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QTiltReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQTiltReading(ptr QTiltReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReading_PTR().Pointer()
	}
	return nil
}

func (n *QTiltReading) InitFromInternal(ptr uintptr, name string) {
	n.QSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTiltReading) ClassNameInternalF() string {
	return n.QSensorReading_PTR().ClassNameInternalF()
}

func NewQTiltReadingFromPointer(ptr unsafe.Pointer) (n *QTiltReading) {
	n = new(QTiltReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QTiltReading")
	return
}

func (ptr *QTiltReading) DestroyQTiltReading() {
}

func (ptr *QTiltReading) SetXRotation(x float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetXRotation", x})
}

func (ptr *QTiltReading) SetYRotation(y float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetYRotation", y})
}

func (ptr *QTiltReading) XRotation() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "XRotation"}).(float64)
}

func (ptr *QTiltReading) YRotation() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "YRotation"}).(float64)
}

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
}

func (ptr *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return ptr
}

func (ptr *QTiltSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QTiltSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQTiltSensor(ptr QTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensor_PTR().Pointer()
	}
	return nil
}

func (n *QTiltSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTiltSensor) ClassNameInternalF() string {
	return n.QSensor_PTR().ClassNameInternalF()
}

func NewQTiltSensorFromPointer(ptr unsafe.Pointer) (n *QTiltSensor) {
	n = new(QTiltSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QTiltSensor")
	return
}
func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {

	return internal.CallLocalFunction([]interface{}{"", "", "sensors.NewQTiltSensor", "", parent}).(*QTiltSensor)
}

func (ptr *QTiltSensor) Calibrate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Calibrate"})
}

func (ptr *QTiltSensor) Reading() *QTiltReading {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reading"}).(*QTiltReading)
}

func (ptr *QTiltSensor) ConnectDestroyQTiltSensor(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQTiltSensor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTiltSensor) DisconnectDestroyQTiltSensor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQTiltSensor"})
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTiltSensor"})
}

func (ptr *QTiltSensor) DestroyQTiltSensorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTiltSensorDefault"})
}

type QmlAccelerometer struct {
	QmlSensor
}

type QmlAccelerometer_ITF interface {
	QmlSensor_ITF
	QmlAccelerometer_PTR() *QmlAccelerometer
}

func (ptr *QmlAccelerometer) QmlAccelerometer_PTR() *QmlAccelerometer {
	return ptr
}

func (ptr *QmlAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlAccelerometer(ptr QmlAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *QmlAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAccelerometer) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlAccelerometerFromPointer(ptr unsafe.Pointer) (n *QmlAccelerometer) {
	n = new(QmlAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAccelerometer")
	return
}

func (ptr *QmlAccelerometer) DestroyQmlAccelerometer() {
}

type QmlAccelerometerReading struct {
	QmlSensorReading
}

type QmlAccelerometerReading_ITF interface {
	QmlSensorReading_ITF
	QmlAccelerometerReading_PTR() *QmlAccelerometerReading
}

func (ptr *QmlAccelerometerReading) QmlAccelerometerReading_PTR() *QmlAccelerometerReading {
	return ptr
}

func (ptr *QmlAccelerometerReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAccelerometerReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlAccelerometerReading(ptr QmlAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlAccelerometerReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAccelerometerReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlAccelerometerReadingFromPointer(ptr unsafe.Pointer) (n *QmlAccelerometerReading) {
	n = new(QmlAccelerometerReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAccelerometerReading")
	return
}

func (ptr *QmlAccelerometerReading) DestroyQmlAccelerometerReading() {
}

type QmlAltimeter struct {
	QmlSensor
}

type QmlAltimeter_ITF interface {
	QmlSensor_ITF
	QmlAltimeter_PTR() *QmlAltimeter
}

func (ptr *QmlAltimeter) QmlAltimeter_PTR() *QmlAltimeter {
	return ptr
}

func (ptr *QmlAltimeter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAltimeter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlAltimeter(ptr QmlAltimeter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAltimeter_PTR().Pointer()
	}
	return nil
}

func (n *QmlAltimeter) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAltimeter) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlAltimeterFromPointer(ptr unsafe.Pointer) (n *QmlAltimeter) {
	n = new(QmlAltimeter)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAltimeter")
	return
}

func (ptr *QmlAltimeter) DestroyQmlAltimeter() {
}

type QmlAltimeterReading struct {
	QmlSensorReading
}

type QmlAltimeterReading_ITF interface {
	QmlSensorReading_ITF
	QmlAltimeterReading_PTR() *QmlAltimeterReading
}

func (ptr *QmlAltimeterReading) QmlAltimeterReading_PTR() *QmlAltimeterReading {
	return ptr
}

func (ptr *QmlAltimeterReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAltimeterReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlAltimeterReading(ptr QmlAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAltimeterReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlAltimeterReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAltimeterReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlAltimeterReadingFromPointer(ptr unsafe.Pointer) (n *QmlAltimeterReading) {
	n = new(QmlAltimeterReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAltimeterReading")
	return
}

func (ptr *QmlAltimeterReading) DestroyQmlAltimeterReading() {
}

type QmlAmbientLightSensor struct {
	QmlSensor
}

type QmlAmbientLightSensor_ITF interface {
	QmlSensor_ITF
	QmlAmbientLightSensor_PTR() *QmlAmbientLightSensor
}

func (ptr *QmlAmbientLightSensor) QmlAmbientLightSensor_PTR() *QmlAmbientLightSensor {
	return ptr
}

func (ptr *QmlAmbientLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAmbientLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlAmbientLightSensor(ptr QmlAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlAmbientLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAmbientLightSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlAmbientLightSensorFromPointer(ptr unsafe.Pointer) (n *QmlAmbientLightSensor) {
	n = new(QmlAmbientLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAmbientLightSensor")
	return
}

func (ptr *QmlAmbientLightSensor) DestroyQmlAmbientLightSensor() {
}

type QmlAmbientLightSensorReading struct {
	QmlSensorReading
}

type QmlAmbientLightSensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlAmbientLightSensorReading_PTR() *QmlAmbientLightSensorReading
}

func (ptr *QmlAmbientLightSensorReading) QmlAmbientLightSensorReading_PTR() *QmlAmbientLightSensorReading {
	return ptr
}

func (ptr *QmlAmbientLightSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAmbientLightSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlAmbientLightSensorReading(ptr QmlAmbientLightSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAmbientLightSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlAmbientLightSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAmbientLightSensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlAmbientLightSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlAmbientLightSensorReading) {
	n = new(QmlAmbientLightSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAmbientLightSensorReading")
	return
}

func (ptr *QmlAmbientLightSensorReading) DestroyQmlAmbientLightSensorReading() {
}

type QmlAmbientTemperatureReading struct {
	QmlSensorReading
}

type QmlAmbientTemperatureReading_ITF interface {
	QmlSensorReading_ITF
	QmlAmbientTemperatureReading_PTR() *QmlAmbientTemperatureReading
}

func (ptr *QmlAmbientTemperatureReading) QmlAmbientTemperatureReading_PTR() *QmlAmbientTemperatureReading {
	return ptr
}

func (ptr *QmlAmbientTemperatureReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAmbientTemperatureReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlAmbientTemperatureReading(ptr QmlAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlAmbientTemperatureReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAmbientTemperatureReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) (n *QmlAmbientTemperatureReading) {
	n = new(QmlAmbientTemperatureReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAmbientTemperatureReading")
	return
}

func (ptr *QmlAmbientTemperatureReading) DestroyQmlAmbientTemperatureReading() {
}

type QmlAmbientTemperatureSensor struct {
	QmlSensor
}

type QmlAmbientTemperatureSensor_ITF interface {
	QmlSensor_ITF
	QmlAmbientTemperatureSensor_PTR() *QmlAmbientTemperatureSensor
}

func (ptr *QmlAmbientTemperatureSensor) QmlAmbientTemperatureSensor_PTR() *QmlAmbientTemperatureSensor {
	return ptr
}

func (ptr *QmlAmbientTemperatureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlAmbientTemperatureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlAmbientTemperatureSensor(ptr QmlAmbientTemperatureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlAmbientTemperatureSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlAmbientTemperatureSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlAmbientTemperatureSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlAmbientTemperatureSensorFromPointer(ptr unsafe.Pointer) (n *QmlAmbientTemperatureSensor) {
	n = new(QmlAmbientTemperatureSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlAmbientTemperatureSensor")
	return
}

func (ptr *QmlAmbientTemperatureSensor) DestroyQmlAmbientTemperatureSensor() {
}

type QmlCompass struct {
	QmlSensor
}

type QmlCompass_ITF interface {
	QmlSensor_ITF
	QmlCompass_PTR() *QmlCompass
}

func (ptr *QmlCompass) QmlCompass_PTR() *QmlCompass {
	return ptr
}

func (ptr *QmlCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlCompass(ptr QmlCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlCompass_PTR().Pointer()
	}
	return nil
}

func (n *QmlCompass) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlCompass) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlCompassFromPointer(ptr unsafe.Pointer) (n *QmlCompass) {
	n = new(QmlCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlCompass")
	return
}

func (ptr *QmlCompass) DestroyQmlCompass() {
}

type QmlCompassReading struct {
	QmlSensorReading
}

type QmlCompassReading_ITF interface {
	QmlSensorReading_ITF
	QmlCompassReading_PTR() *QmlCompassReading
}

func (ptr *QmlCompassReading) QmlCompassReading_PTR() *QmlCompassReading {
	return ptr
}

func (ptr *QmlCompassReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlCompassReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlCompassReading(ptr QmlCompassReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlCompassReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlCompassReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlCompassReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlCompassReadingFromPointer(ptr unsafe.Pointer) (n *QmlCompassReading) {
	n = new(QmlCompassReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlCompassReading")
	return
}

func (ptr *QmlCompassReading) DestroyQmlCompassReading() {
}

type QmlDistanceReading struct {
	QmlSensorReading
}

type QmlDistanceReading_ITF interface {
	QmlSensorReading_ITF
	QmlDistanceReading_PTR() *QmlDistanceReading
}

func (ptr *QmlDistanceReading) QmlDistanceReading_PTR() *QmlDistanceReading {
	return ptr
}

func (ptr *QmlDistanceReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlDistanceReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlDistanceReading(ptr QmlDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlDistanceReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlDistanceReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlDistanceReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlDistanceReadingFromPointer(ptr unsafe.Pointer) (n *QmlDistanceReading) {
	n = new(QmlDistanceReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlDistanceReading")
	return
}

func (ptr *QmlDistanceReading) DestroyQmlDistanceReading() {
}

type QmlDistanceSensor struct {
	QmlSensor
}

type QmlDistanceSensor_ITF interface {
	QmlSensor_ITF
	QmlDistanceSensor_PTR() *QmlDistanceSensor
}

func (ptr *QmlDistanceSensor) QmlDistanceSensor_PTR() *QmlDistanceSensor {
	return ptr
}

func (ptr *QmlDistanceSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlDistanceSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlDistanceSensor(ptr QmlDistanceSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlDistanceSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlDistanceSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlDistanceSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlDistanceSensorFromPointer(ptr unsafe.Pointer) (n *QmlDistanceSensor) {
	n = new(QmlDistanceSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlDistanceSensor")
	return
}

func (ptr *QmlDistanceSensor) DestroyQmlDistanceSensor() {
}

type QmlGyroscope struct {
	QmlSensor
}

type QmlGyroscope_ITF interface {
	QmlSensor_ITF
	QmlGyroscope_PTR() *QmlGyroscope
}

func (ptr *QmlGyroscope) QmlGyroscope_PTR() *QmlGyroscope {
	return ptr
}

func (ptr *QmlGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlGyroscope(ptr QmlGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *QmlGyroscope) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlGyroscope) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlGyroscopeFromPointer(ptr unsafe.Pointer) (n *QmlGyroscope) {
	n = new(QmlGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlGyroscope")
	return
}

func (ptr *QmlGyroscope) DestroyQmlGyroscope() {
}

type QmlGyroscopeReading struct {
	QmlSensorReading
}

type QmlGyroscopeReading_ITF interface {
	QmlSensorReading_ITF
	QmlGyroscopeReading_PTR() *QmlGyroscopeReading
}

func (ptr *QmlGyroscopeReading) QmlGyroscopeReading_PTR() *QmlGyroscopeReading {
	return ptr
}

func (ptr *QmlGyroscopeReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlGyroscopeReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlGyroscopeReading(ptr QmlGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlGyroscopeReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlGyroscopeReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlGyroscopeReadingFromPointer(ptr unsafe.Pointer) (n *QmlGyroscopeReading) {
	n = new(QmlGyroscopeReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlGyroscopeReading")
	return
}

func (ptr *QmlGyroscopeReading) DestroyQmlGyroscopeReading() {
}

type QmlHolsterReading struct {
	QmlSensorReading
}

type QmlHolsterReading_ITF interface {
	QmlSensorReading_ITF
	QmlHolsterReading_PTR() *QmlHolsterReading
}

func (ptr *QmlHolsterReading) QmlHolsterReading_PTR() *QmlHolsterReading {
	return ptr
}

func (ptr *QmlHolsterReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlHolsterReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlHolsterReading(ptr QmlHolsterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlHolsterReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlHolsterReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlHolsterReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlHolsterReadingFromPointer(ptr unsafe.Pointer) (n *QmlHolsterReading) {
	n = new(QmlHolsterReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlHolsterReading")
	return
}

func (ptr *QmlHolsterReading) DestroyQmlHolsterReading() {
}

type QmlHolsterSensor struct {
	QmlSensor
}

type QmlHolsterSensor_ITF interface {
	QmlSensor_ITF
	QmlHolsterSensor_PTR() *QmlHolsterSensor
}

func (ptr *QmlHolsterSensor) QmlHolsterSensor_PTR() *QmlHolsterSensor {
	return ptr
}

func (ptr *QmlHolsterSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlHolsterSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlHolsterSensor(ptr QmlHolsterSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlHolsterSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlHolsterSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlHolsterSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlHolsterSensorFromPointer(ptr unsafe.Pointer) (n *QmlHolsterSensor) {
	n = new(QmlHolsterSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlHolsterSensor")
	return
}

func (ptr *QmlHolsterSensor) DestroyQmlHolsterSensor() {
}

type QmlHumidityReading struct {
	QmlSensorReading
}

type QmlHumidityReading_ITF interface {
	QmlSensorReading_ITF
	QmlHumidityReading_PTR() *QmlHumidityReading
}

func (ptr *QmlHumidityReading) QmlHumidityReading_PTR() *QmlHumidityReading {
	return ptr
}

func (ptr *QmlHumidityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlHumidityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlHumidityReading(ptr QmlHumidityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlHumidityReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlHumidityReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlHumidityReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlHumidityReadingFromPointer(ptr unsafe.Pointer) (n *QmlHumidityReading) {
	n = new(QmlHumidityReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlHumidityReading")
	return
}

func (ptr *QmlHumidityReading) DestroyQmlHumidityReading() {
}

type QmlHumiditySensor struct {
	QmlSensor
}

type QmlHumiditySensor_ITF interface {
	QmlSensor_ITF
	QmlHumiditySensor_PTR() *QmlHumiditySensor
}

func (ptr *QmlHumiditySensor) QmlHumiditySensor_PTR() *QmlHumiditySensor {
	return ptr
}

func (ptr *QmlHumiditySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlHumiditySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlHumiditySensor(ptr QmlHumiditySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlHumiditySensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlHumiditySensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlHumiditySensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlHumiditySensorFromPointer(ptr unsafe.Pointer) (n *QmlHumiditySensor) {
	n = new(QmlHumiditySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlHumiditySensor")
	return
}

func (ptr *QmlHumiditySensor) DestroyQmlHumiditySensor() {
}

type QmlIRProximitySensor struct {
	QmlSensor
}

type QmlIRProximitySensor_ITF interface {
	QmlSensor_ITF
	QmlIRProximitySensor_PTR() *QmlIRProximitySensor
}

func (ptr *QmlIRProximitySensor) QmlIRProximitySensor_PTR() *QmlIRProximitySensor {
	return ptr
}

func (ptr *QmlIRProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlIRProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlIRProximitySensor(ptr QmlIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlIRProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlIRProximitySensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlIRProximitySensorFromPointer(ptr unsafe.Pointer) (n *QmlIRProximitySensor) {
	n = new(QmlIRProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlIRProximitySensor")
	return
}

func (ptr *QmlIRProximitySensor) DestroyQmlIRProximitySensor() {
}

type QmlIRProximitySensorReading struct {
	QmlSensorReading
}

type QmlIRProximitySensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlIRProximitySensorReading_PTR() *QmlIRProximitySensorReading
}

func (ptr *QmlIRProximitySensorReading) QmlIRProximitySensorReading_PTR() *QmlIRProximitySensorReading {
	return ptr
}

func (ptr *QmlIRProximitySensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlIRProximitySensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlIRProximitySensorReading(ptr QmlIRProximitySensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlIRProximitySensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlIRProximitySensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlIRProximitySensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlIRProximitySensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlIRProximitySensorReading) {
	n = new(QmlIRProximitySensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlIRProximitySensorReading")
	return
}

func (ptr *QmlIRProximitySensorReading) DestroyQmlIRProximitySensorReading() {
}

type QmlLidReading struct {
	QmlSensorReading
}

type QmlLidReading_ITF interface {
	QmlSensorReading_ITF
	QmlLidReading_PTR() *QmlLidReading
}

func (ptr *QmlLidReading) QmlLidReading_PTR() *QmlLidReading {
	return ptr
}

func (ptr *QmlLidReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlLidReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlLidReading(ptr QmlLidReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlLidReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlLidReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlLidReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlLidReadingFromPointer(ptr unsafe.Pointer) (n *QmlLidReading) {
	n = new(QmlLidReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlLidReading")
	return
}

func (ptr *QmlLidReading) DestroyQmlLidReading() {
}

type QmlLidSensor struct {
	QmlSensor
}

type QmlLidSensor_ITF interface {
	QmlSensor_ITF
	QmlLidSensor_PTR() *QmlLidSensor
}

func (ptr *QmlLidSensor) QmlLidSensor_PTR() *QmlLidSensor {
	return ptr
}

func (ptr *QmlLidSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlLidSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlLidSensor(ptr QmlLidSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlLidSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlLidSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlLidSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlLidSensorFromPointer(ptr unsafe.Pointer) (n *QmlLidSensor) {
	n = new(QmlLidSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlLidSensor")
	return
}

func (ptr *QmlLidSensor) DestroyQmlLidSensor() {
}

type QmlLightSensor struct {
	QmlSensor
}

type QmlLightSensor_ITF interface {
	QmlSensor_ITF
	QmlLightSensor_PTR() *QmlLightSensor
}

func (ptr *QmlLightSensor) QmlLightSensor_PTR() *QmlLightSensor {
	return ptr
}

func (ptr *QmlLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlLightSensor(ptr QmlLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlLightSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlLightSensorFromPointer(ptr unsafe.Pointer) (n *QmlLightSensor) {
	n = new(QmlLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlLightSensor")
	return
}

func (ptr *QmlLightSensor) DestroyQmlLightSensor() {
}

type QmlLightSensorReading struct {
	QmlSensorReading
}

type QmlLightSensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlLightSensorReading_PTR() *QmlLightSensorReading
}

func (ptr *QmlLightSensorReading) QmlLightSensorReading_PTR() *QmlLightSensorReading {
	return ptr
}

func (ptr *QmlLightSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlLightSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlLightSensorReading(ptr QmlLightSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlLightSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlLightSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlLightSensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlLightSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlLightSensorReading) {
	n = new(QmlLightSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlLightSensorReading")
	return
}

func (ptr *QmlLightSensorReading) DestroyQmlLightSensorReading() {
}

type QmlMagnetometer struct {
	QmlSensor
}

type QmlMagnetometer_ITF interface {
	QmlSensor_ITF
	QmlMagnetometer_PTR() *QmlMagnetometer
}

func (ptr *QmlMagnetometer) QmlMagnetometer_PTR() *QmlMagnetometer {
	return ptr
}

func (ptr *QmlMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlMagnetometer(ptr QmlMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *QmlMagnetometer) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlMagnetometer) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlMagnetometerFromPointer(ptr unsafe.Pointer) (n *QmlMagnetometer) {
	n = new(QmlMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlMagnetometer")
	return
}

func (ptr *QmlMagnetometer) DestroyQmlMagnetometer() {
}

type QmlMagnetometerReading struct {
	QmlSensorReading
}

type QmlMagnetometerReading_ITF interface {
	QmlSensorReading_ITF
	QmlMagnetometerReading_PTR() *QmlMagnetometerReading
}

func (ptr *QmlMagnetometerReading) QmlMagnetometerReading_PTR() *QmlMagnetometerReading {
	return ptr
}

func (ptr *QmlMagnetometerReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlMagnetometerReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlMagnetometerReading(ptr QmlMagnetometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlMagnetometerReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlMagnetometerReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlMagnetometerReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlMagnetometerReadingFromPointer(ptr unsafe.Pointer) (n *QmlMagnetometerReading) {
	n = new(QmlMagnetometerReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlMagnetometerReading")
	return
}

func (ptr *QmlMagnetometerReading) DestroyQmlMagnetometerReading() {
}

type QmlOrientationSensor struct {
	QmlSensor
}

type QmlOrientationSensor_ITF interface {
	QmlSensor_ITF
	QmlOrientationSensor_PTR() *QmlOrientationSensor
}

func (ptr *QmlOrientationSensor) QmlOrientationSensor_PTR() *QmlOrientationSensor {
	return ptr
}

func (ptr *QmlOrientationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlOrientationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlOrientationSensor(ptr QmlOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlOrientationSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlOrientationSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlOrientationSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlOrientationSensorFromPointer(ptr unsafe.Pointer) (n *QmlOrientationSensor) {
	n = new(QmlOrientationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlOrientationSensor")
	return
}

func (ptr *QmlOrientationSensor) DestroyQmlOrientationSensor() {
}

type QmlOrientationSensorReading struct {
	QmlSensorReading
}

type QmlOrientationSensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlOrientationSensorReading_PTR() *QmlOrientationSensorReading
}

func (ptr *QmlOrientationSensorReading) QmlOrientationSensorReading_PTR() *QmlOrientationSensorReading {
	return ptr
}

func (ptr *QmlOrientationSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlOrientationSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlOrientationSensorReading(ptr QmlOrientationSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlOrientationSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlOrientationSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlOrientationSensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlOrientationSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlOrientationSensorReading) {
	n = new(QmlOrientationSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlOrientationSensorReading")
	return
}

func (ptr *QmlOrientationSensorReading) DestroyQmlOrientationSensorReading() {
}

type QmlPressureReading struct {
	QmlSensorReading
}

type QmlPressureReading_ITF interface {
	QmlSensorReading_ITF
	QmlPressureReading_PTR() *QmlPressureReading
}

func (ptr *QmlPressureReading) QmlPressureReading_PTR() *QmlPressureReading {
	return ptr
}

func (ptr *QmlPressureReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlPressureReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlPressureReading(ptr QmlPressureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlPressureReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlPressureReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlPressureReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlPressureReadingFromPointer(ptr unsafe.Pointer) (n *QmlPressureReading) {
	n = new(QmlPressureReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlPressureReading")
	return
}

func (ptr *QmlPressureReading) DestroyQmlPressureReading() {
}

type QmlPressureSensor struct {
	QmlSensor
}

type QmlPressureSensor_ITF interface {
	QmlSensor_ITF
	QmlPressureSensor_PTR() *QmlPressureSensor
}

func (ptr *QmlPressureSensor) QmlPressureSensor_PTR() *QmlPressureSensor {
	return ptr
}

func (ptr *QmlPressureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlPressureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlPressureSensor(ptr QmlPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlPressureSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlPressureSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlPressureSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlPressureSensorFromPointer(ptr unsafe.Pointer) (n *QmlPressureSensor) {
	n = new(QmlPressureSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlPressureSensor")
	return
}

func (ptr *QmlPressureSensor) DestroyQmlPressureSensor() {
}

type QmlProximitySensor struct {
	QmlSensor
}

type QmlProximitySensor_ITF interface {
	QmlSensor_ITF
	QmlProximitySensor_PTR() *QmlProximitySensor
}

func (ptr *QmlProximitySensor) QmlProximitySensor_PTR() *QmlProximitySensor {
	return ptr
}

func (ptr *QmlProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlProximitySensor(ptr QmlProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlProximitySensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlProximitySensorFromPointer(ptr unsafe.Pointer) (n *QmlProximitySensor) {
	n = new(QmlProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlProximitySensor")
	return
}

func (ptr *QmlProximitySensor) DestroyQmlProximitySensor() {
}

type QmlProximitySensorReading struct {
	QmlSensorReading
}

type QmlProximitySensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlProximitySensorReading_PTR() *QmlProximitySensorReading
}

func (ptr *QmlProximitySensorReading) QmlProximitySensorReading_PTR() *QmlProximitySensorReading {
	return ptr
}

func (ptr *QmlProximitySensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlProximitySensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlProximitySensorReading(ptr QmlProximitySensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlProximitySensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlProximitySensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlProximitySensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlProximitySensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlProximitySensorReading) {
	n = new(QmlProximitySensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlProximitySensorReading")
	return
}

func (ptr *QmlProximitySensorReading) DestroyQmlProximitySensorReading() {
}

type QmlRotationSensor struct {
	QmlSensor
}

type QmlRotationSensor_ITF interface {
	QmlSensor_ITF
	QmlRotationSensor_PTR() *QmlRotationSensor
}

func (ptr *QmlRotationSensor) QmlRotationSensor_PTR() *QmlRotationSensor {
	return ptr
}

func (ptr *QmlRotationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlRotationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlRotationSensor(ptr QmlRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlRotationSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlRotationSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlRotationSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlRotationSensorFromPointer(ptr unsafe.Pointer) (n *QmlRotationSensor) {
	n = new(QmlRotationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlRotationSensor")
	return
}

func (ptr *QmlRotationSensor) DestroyQmlRotationSensor() {
}

type QmlRotationSensorReading struct {
	QmlSensorReading
}

type QmlRotationSensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlRotationSensorReading_PTR() *QmlRotationSensorReading
}

func (ptr *QmlRotationSensorReading) QmlRotationSensorReading_PTR() *QmlRotationSensorReading {
	return ptr
}

func (ptr *QmlRotationSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlRotationSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlRotationSensorReading(ptr QmlRotationSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlRotationSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlRotationSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlRotationSensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlRotationSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlRotationSensorReading) {
	n = new(QmlRotationSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlRotationSensorReading")
	return
}

func (ptr *QmlRotationSensorReading) DestroyQmlRotationSensorReading() {
}

type QmlSensor struct {
	core.QObject
}

type QmlSensor_ITF interface {
	core.QObject_ITF
	QmlSensor_PTR() *QmlSensor
}

func (ptr *QmlSensor) QmlSensor_PTR() *QmlSensor {
	return ptr
}

func (ptr *QmlSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlSensor(ptr QmlSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlSensor) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlSensor) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQmlSensorFromPointer(ptr unsafe.Pointer) (n *QmlSensor) {
	n = new(QmlSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlSensor")
	return
}

func (ptr *QmlSensor) DestroyQmlSensor() {
}

type QmlSensorGesture struct {
	core.QObject
}

type QmlSensorGesture_ITF interface {
	core.QObject_ITF
	QmlSensorGesture_PTR() *QmlSensorGesture
}

func (ptr *QmlSensorGesture) QmlSensorGesture_PTR() *QmlSensorGesture {
	return ptr
}

func (ptr *QmlSensorGesture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlSensorGesture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlSensorGesture(ptr QmlSensorGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorGesture_PTR().Pointer()
	}
	return nil
}

func (n *QmlSensorGesture) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlSensorGesture) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQmlSensorGestureFromPointer(ptr unsafe.Pointer) (n *QmlSensorGesture) {
	n = new(QmlSensorGesture)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlSensorGesture")
	return
}

func (ptr *QmlSensorGesture) DestroyQmlSensorGesture() {
}

type QmlSensorGlobal struct {
	core.QObject
}

type QmlSensorGlobal_ITF interface {
	core.QObject_ITF
	QmlSensorGlobal_PTR() *QmlSensorGlobal
}

func (ptr *QmlSensorGlobal) QmlSensorGlobal_PTR() *QmlSensorGlobal {
	return ptr
}

func (ptr *QmlSensorGlobal) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlSensorGlobal) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlSensorGlobal(ptr QmlSensorGlobal_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorGlobal_PTR().Pointer()
	}
	return nil
}

func (n *QmlSensorGlobal) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlSensorGlobal) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQmlSensorGlobalFromPointer(ptr unsafe.Pointer) (n *QmlSensorGlobal) {
	n = new(QmlSensorGlobal)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlSensorGlobal")
	return
}

func (ptr *QmlSensorGlobal) DestroyQmlSensorGlobal() {
}

type QmlSensorOutputRange struct {
	core.QObject
}

type QmlSensorOutputRange_ITF interface {
	core.QObject_ITF
	QmlSensorOutputRange_PTR() *QmlSensorOutputRange
}

func (ptr *QmlSensorOutputRange) QmlSensorOutputRange_PTR() *QmlSensorOutputRange {
	return ptr
}

func (ptr *QmlSensorOutputRange) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlSensorOutputRange) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlSensorOutputRange(ptr QmlSensorOutputRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorOutputRange_PTR().Pointer()
	}
	return nil
}

func (n *QmlSensorOutputRange) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlSensorOutputRange) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQmlSensorOutputRangeFromPointer(ptr unsafe.Pointer) (n *QmlSensorOutputRange) {
	n = new(QmlSensorOutputRange)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlSensorOutputRange")
	return
}

func (ptr *QmlSensorOutputRange) DestroyQmlSensorOutputRange() {
}

type QmlSensorRange struct {
	core.QObject
}

type QmlSensorRange_ITF interface {
	core.QObject_ITF
	QmlSensorRange_PTR() *QmlSensorRange
}

func (ptr *QmlSensorRange) QmlSensorRange_PTR() *QmlSensorRange {
	return ptr
}

func (ptr *QmlSensorRange) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlSensorRange) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlSensorRange(ptr QmlSensorRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorRange_PTR().Pointer()
	}
	return nil
}

func (n *QmlSensorRange) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlSensorRange) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQmlSensorRangeFromPointer(ptr unsafe.Pointer) (n *QmlSensorRange) {
	n = new(QmlSensorRange)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlSensorRange")
	return
}

func (ptr *QmlSensorRange) DestroyQmlSensorRange() {
}

type QmlSensorReading struct {
	core.QObject
}

type QmlSensorReading_ITF interface {
	core.QObject_ITF
	QmlSensorReading_PTR() *QmlSensorReading
}

func (ptr *QmlSensorReading) QmlSensorReading_PTR() *QmlSensorReading {
	return ptr
}

func (ptr *QmlSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlSensorReading(ptr QmlSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlSensorReading) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQmlSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlSensorReading) {
	n = new(QmlSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlSensorReading")
	return
}

func (ptr *QmlSensorReading) DestroyQmlSensorReading() {
}

type QmlTapSensor struct {
	QmlSensor
}

type QmlTapSensor_ITF interface {
	QmlSensor_ITF
	QmlTapSensor_PTR() *QmlTapSensor
}

func (ptr *QmlTapSensor) QmlTapSensor_PTR() *QmlTapSensor {
	return ptr
}

func (ptr *QmlTapSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlTapSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlTapSensor(ptr QmlTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlTapSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlTapSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlTapSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlTapSensorFromPointer(ptr unsafe.Pointer) (n *QmlTapSensor) {
	n = new(QmlTapSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlTapSensor")
	return
}

func (ptr *QmlTapSensor) DestroyQmlTapSensor() {
}

type QmlTapSensorReading struct {
	QmlSensorReading
}

type QmlTapSensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlTapSensorReading_PTR() *QmlTapSensorReading
}

func (ptr *QmlTapSensorReading) QmlTapSensorReading_PTR() *QmlTapSensorReading {
	return ptr
}

func (ptr *QmlTapSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlTapSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlTapSensorReading(ptr QmlTapSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlTapSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlTapSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlTapSensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlTapSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlTapSensorReading) {
	n = new(QmlTapSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlTapSensorReading")
	return
}

func (ptr *QmlTapSensorReading) DestroyQmlTapSensorReading() {
}

type QmlTiltSensor struct {
	QmlSensor
}

type QmlTiltSensor_ITF interface {
	QmlSensor_ITF
	QmlTiltSensor_PTR() *QmlTiltSensor
}

func (ptr *QmlTiltSensor) QmlTiltSensor_PTR() *QmlTiltSensor {
	return ptr
}

func (ptr *QmlTiltSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlTiltSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensor_PTR().SetPointer(p)
	}
}

func PointerFromQmlTiltSensor(ptr QmlTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlTiltSensor_PTR().Pointer()
	}
	return nil
}

func (n *QmlTiltSensor) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensor_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlTiltSensor) ClassNameInternalF() string {
	return n.QmlSensor_PTR().ClassNameInternalF()
}

func NewQmlTiltSensorFromPointer(ptr unsafe.Pointer) (n *QmlTiltSensor) {
	n = new(QmlTiltSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlTiltSensor")
	return
}

func (ptr *QmlTiltSensor) DestroyQmlTiltSensor() {
}

type QmlTiltSensorReading struct {
	QmlSensorReading
}

type QmlTiltSensorReading_ITF interface {
	QmlSensorReading_ITF
	QmlTiltSensorReading_PTR() *QmlTiltSensorReading
}

func (ptr *QmlTiltSensorReading) QmlTiltSensorReading_PTR() *QmlTiltSensorReading {
	return ptr
}

func (ptr *QmlTiltSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlTiltSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QmlSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQmlTiltSensorReading(ptr QmlTiltSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlTiltSensorReading_PTR().Pointer()
	}
	return nil
}

func (n *QmlTiltSensorReading) InitFromInternal(ptr uintptr, name string) {
	n.QmlSensorReading_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QmlTiltSensorReading) ClassNameInternalF() string {
	return n.QmlSensorReading_PTR().ClassNameInternalF()
}

func NewQmlTiltSensorReadingFromPointer(ptr unsafe.Pointer) (n *QmlTiltSensorReading) {
	n = new(QmlTiltSensorReading)
	n.InitFromInternal(uintptr(ptr), "sensors.QmlTiltSensorReading")
	return
}

func (ptr *QmlTiltSensorReading) DestroyQmlTiltSensorReading() {
}

type SensorEventQueue struct {
	ThreadSafeSensorBackend
}

type SensorEventQueue_ITF interface {
	ThreadSafeSensorBackend_ITF
	SensorEventQueue_PTR() *SensorEventQueue
}

func (ptr *SensorEventQueue) SensorEventQueue_PTR() *SensorEventQueue {
	return ptr
}

func (ptr *SensorEventQueue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ThreadSafeSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorEventQueue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ThreadSafeSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromSensorEventQueue(ptr SensorEventQueue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorEventQueue_PTR().Pointer()
	}
	return nil
}

func (n *SensorEventQueue) InitFromInternal(ptr uintptr, name string) {
	n.ThreadSafeSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorEventQueue) ClassNameInternalF() string {
	return n.ThreadSafeSensorBackend_PTR().ClassNameInternalF()
}

func NewSensorEventQueueFromPointer(ptr unsafe.Pointer) (n *SensorEventQueue) {
	n = new(SensorEventQueue)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorEventQueue")
	return
}

func (ptr *SensorEventQueue) DestroySensorEventQueue() {
}

type SensorManager struct {
	core.QThread
}

type SensorManager_ITF interface {
	core.QThread_ITF
	SensorManager_PTR() *SensorManager
}

func (ptr *SensorManager) SensorManager_PTR() *SensorManager {
	return ptr
}

func (ptr *SensorManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QThread_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QThread_PTR().SetPointer(p)
	}
}

func PointerFromSensorManager(ptr SensorManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorManager_PTR().Pointer()
	}
	return nil
}

func (n *SensorManager) InitFromInternal(ptr uintptr, name string) {
	n.QThread_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorManager) ClassNameInternalF() string {
	return n.QThread_PTR().ClassNameInternalF()
}

func NewSensorManagerFromPointer(ptr unsafe.Pointer) (n *SensorManager) {
	n = new(SensorManager)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorManager")
	return
}

func (ptr *SensorManager) DestroySensorManager() {
}

type SensorTagAccelerometer struct {
	SensorTagBase
}

type SensorTagAccelerometer_ITF interface {
	SensorTagBase_ITF
	SensorTagAccelerometer_PTR() *SensorTagAccelerometer
}

func (ptr *SensorTagAccelerometer) SensorTagAccelerometer_PTR() *SensorTagAccelerometer {
	return ptr
}

func (ptr *SensorTagAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagAccelerometer(ptr SensorTagAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagAccelerometer) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagAccelerometerFromPointer(ptr unsafe.Pointer) (n *SensorTagAccelerometer) {
	n = new(SensorTagAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagAccelerometer")
	return
}

func (ptr *SensorTagAccelerometer) DestroySensorTagAccelerometer() {
}

type SensorTagAls struct {
	SensorTagBase
}

type SensorTagAls_ITF interface {
	SensorTagBase_ITF
	SensorTagAls_PTR() *SensorTagAls
}

func (ptr *SensorTagAls) SensorTagAls_PTR() *SensorTagAls {
	return ptr
}

func (ptr *SensorTagAls) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagAls) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagAls(ptr SensorTagAls_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagAls_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagAls) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagAls) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagAlsFromPointer(ptr unsafe.Pointer) (n *SensorTagAls) {
	n = new(SensorTagAls)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagAls")
	return
}

func (ptr *SensorTagAls) DestroySensorTagAls() {
}

type SensorTagBase struct {
	QSensorBackend
}

type SensorTagBase_ITF interface {
	QSensorBackend_ITF
	SensorTagBase_PTR() *SensorTagBase
}

func (ptr *SensorTagBase) SensorTagBase_PTR() *SensorTagBase {
	return ptr
}

func (ptr *SensorTagBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagBase(ptr SensorTagBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagBase) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagBase) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewSensorTagBaseFromPointer(ptr unsafe.Pointer) (n *SensorTagBase) {
	n = new(SensorTagBase)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagBase")
	return
}

func (ptr *SensorTagBase) DestroySensorTagBase() {
}

type SensorTagGyroscope struct {
	SensorTagBase
}

type SensorTagGyroscope_ITF interface {
	SensorTagBase_ITF
	SensorTagGyroscope_PTR() *SensorTagGyroscope
}

func (ptr *SensorTagGyroscope) SensorTagGyroscope_PTR() *SensorTagGyroscope {
	return ptr
}

func (ptr *SensorTagGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagGyroscope(ptr SensorTagGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagGyroscope) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagGyroscope) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagGyroscopeFromPointer(ptr unsafe.Pointer) (n *SensorTagGyroscope) {
	n = new(SensorTagGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagGyroscope")
	return
}

func (ptr *SensorTagGyroscope) DestroySensorTagGyroscope() {
}

type SensorTagHumiditySensor struct {
	SensorTagBase
}

type SensorTagHumiditySensor_ITF interface {
	SensorTagBase_ITF
	SensorTagHumiditySensor_PTR() *SensorTagHumiditySensor
}

func (ptr *SensorTagHumiditySensor) SensorTagHumiditySensor_PTR() *SensorTagHumiditySensor {
	return ptr
}

func (ptr *SensorTagHumiditySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagHumiditySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagHumiditySensor(ptr SensorTagHumiditySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagHumiditySensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagHumiditySensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagHumiditySensor) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagHumiditySensorFromPointer(ptr unsafe.Pointer) (n *SensorTagHumiditySensor) {
	n = new(SensorTagHumiditySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagHumiditySensor")
	return
}

func (ptr *SensorTagHumiditySensor) DestroySensorTagHumiditySensor() {
}

type SensorTagLightSensor struct {
	SensorTagBase
}

type SensorTagLightSensor_ITF interface {
	SensorTagBase_ITF
	SensorTagLightSensor_PTR() *SensorTagLightSensor
}

func (ptr *SensorTagLightSensor) SensorTagLightSensor_PTR() *SensorTagLightSensor {
	return ptr
}

func (ptr *SensorTagLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagLightSensor(ptr SensorTagLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagLightSensor) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagLightSensorFromPointer(ptr unsafe.Pointer) (n *SensorTagLightSensor) {
	n = new(SensorTagLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagLightSensor")
	return
}

func (ptr *SensorTagLightSensor) DestroySensorTagLightSensor() {
}

type SensorTagMagnetometer struct {
	SensorTagBase
}

type SensorTagMagnetometer_ITF interface {
	SensorTagBase_ITF
	SensorTagMagnetometer_PTR() *SensorTagMagnetometer
}

func (ptr *SensorTagMagnetometer) SensorTagMagnetometer_PTR() *SensorTagMagnetometer {
	return ptr
}

func (ptr *SensorTagMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagMagnetometer(ptr SensorTagMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagMagnetometer) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagMagnetometer) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagMagnetometerFromPointer(ptr unsafe.Pointer) (n *SensorTagMagnetometer) {
	n = new(SensorTagMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagMagnetometer")
	return
}

func (ptr *SensorTagMagnetometer) DestroySensorTagMagnetometer() {
}

type SensorTagPressureSensor struct {
	SensorTagBase
}

type SensorTagPressureSensor_ITF interface {
	SensorTagBase_ITF
	SensorTagPressureSensor_PTR() *SensorTagPressureSensor
}

func (ptr *SensorTagPressureSensor) SensorTagPressureSensor_PTR() *SensorTagPressureSensor {
	return ptr
}

func (ptr *SensorTagPressureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagPressureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagPressureSensor(ptr SensorTagPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagPressureSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagPressureSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagPressureSensor) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagPressureSensorFromPointer(ptr unsafe.Pointer) (n *SensorTagPressureSensor) {
	n = new(SensorTagPressureSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagPressureSensor")
	return
}

func (ptr *SensorTagPressureSensor) DestroySensorTagPressureSensor() {
}

type SensorTagTemperatureSensor struct {
	SensorTagBase
}

type SensorTagTemperatureSensor_ITF interface {
	SensorTagBase_ITF
	SensorTagTemperatureSensor_PTR() *SensorTagTemperatureSensor
}

func (ptr *SensorTagTemperatureSensor) SensorTagTemperatureSensor_PTR() *SensorTagTemperatureSensor {
	return ptr
}

func (ptr *SensorTagTemperatureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorTagTemperatureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorTagBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorTagTemperatureSensor(ptr SensorTagTemperatureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorTagTemperatureSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorTagTemperatureSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorTagBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorTagTemperatureSensor) ClassNameInternalF() string {
	return n.SensorTagBase_PTR().ClassNameInternalF()
}

func NewSensorTagTemperatureSensorFromPointer(ptr unsafe.Pointer) (n *SensorTagTemperatureSensor) {
	n = new(SensorTagTemperatureSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorTagTemperatureSensor")
	return
}

func (ptr *SensorTagTemperatureSensor) DestroySensorTagTemperatureSensor() {
}

type SensorfwCompass struct {
	SensorfwSensorBase
}

type SensorfwCompass_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwCompass_PTR() *SensorfwCompass
}

func (ptr *SensorfwCompass) SensorfwCompass_PTR() *SensorfwCompass {
	return ptr
}

func (ptr *SensorfwCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwCompass(ptr SensorfwCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwCompass_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwCompass) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwCompass) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwCompassFromPointer(ptr unsafe.Pointer) (n *SensorfwCompass) {
	n = new(SensorfwCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwCompass")
	return
}

func (ptr *SensorfwCompass) DestroySensorfwCompass() {
}

type SensorfwGyroscope struct {
	SensorfwSensorBase
}

type SensorfwGyroscope_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwGyroscope_PTR() *SensorfwGyroscope
}

func (ptr *SensorfwGyroscope) SensorfwGyroscope_PTR() *SensorfwGyroscope {
	return ptr
}

func (ptr *SensorfwGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwGyroscope(ptr SensorfwGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwGyroscope) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwGyroscope) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwGyroscopeFromPointer(ptr unsafe.Pointer) (n *SensorfwGyroscope) {
	n = new(SensorfwGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwGyroscope")
	return
}

func (ptr *SensorfwGyroscope) DestroySensorfwGyroscope() {
}

type SensorfwIrProximitySensor struct {
	SensorfwSensorBase
}

type SensorfwIrProximitySensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwIrProximitySensor_PTR() *SensorfwIrProximitySensor
}

func (ptr *SensorfwIrProximitySensor) SensorfwIrProximitySensor_PTR() *SensorfwIrProximitySensor {
	return ptr
}

func (ptr *SensorfwIrProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwIrProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwIrProximitySensor(ptr SensorfwIrProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwIrProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwIrProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwIrProximitySensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwIrProximitySensorFromPointer(ptr unsafe.Pointer) (n *SensorfwIrProximitySensor) {
	n = new(SensorfwIrProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwIrProximitySensor")
	return
}

func (ptr *SensorfwIrProximitySensor) DestroySensorfwIrProximitySensor() {
}

type SensorfwLidSensor struct {
	SensorfwSensorBase
}

type SensorfwLidSensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwLidSensor_PTR() *SensorfwLidSensor
}

func (ptr *SensorfwLidSensor) SensorfwLidSensor_PTR() *SensorfwLidSensor {
	return ptr
}

func (ptr *SensorfwLidSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwLidSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwLidSensor(ptr SensorfwLidSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwLidSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwLidSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwLidSensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwLidSensorFromPointer(ptr unsafe.Pointer) (n *SensorfwLidSensor) {
	n = new(SensorfwLidSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwLidSensor")
	return
}

func (ptr *SensorfwLidSensor) DestroySensorfwLidSensor() {
}

type SensorfwLightSensor struct {
	SensorfwSensorBase
}

type SensorfwLightSensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwLightSensor_PTR() *SensorfwLightSensor
}

func (ptr *SensorfwLightSensor) SensorfwLightSensor_PTR() *SensorfwLightSensor {
	return ptr
}

func (ptr *SensorfwLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwLightSensor(ptr SensorfwLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwLightSensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwLightSensorFromPointer(ptr unsafe.Pointer) (n *SensorfwLightSensor) {
	n = new(SensorfwLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwLightSensor")
	return
}

func (ptr *SensorfwLightSensor) DestroySensorfwLightSensor() {
}

type SensorfwMagnetometer struct {
	SensorfwSensorBase
}

type SensorfwMagnetometer_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwMagnetometer_PTR() *SensorfwMagnetometer
}

func (ptr *SensorfwMagnetometer) SensorfwMagnetometer_PTR() *SensorfwMagnetometer {
	return ptr
}

func (ptr *SensorfwMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwMagnetometer(ptr SensorfwMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwMagnetometer) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwMagnetometer) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwMagnetometerFromPointer(ptr unsafe.Pointer) (n *SensorfwMagnetometer) {
	n = new(SensorfwMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwMagnetometer")
	return
}

func (ptr *SensorfwMagnetometer) DestroySensorfwMagnetometer() {
}

type SensorfwOrientationSensor struct {
	SensorfwSensorBase
}

type SensorfwOrientationSensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwOrientationSensor_PTR() *SensorfwOrientationSensor
}

func (ptr *SensorfwOrientationSensor) SensorfwOrientationSensor_PTR() *SensorfwOrientationSensor {
	return ptr
}

func (ptr *SensorfwOrientationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwOrientationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwOrientationSensor(ptr SensorfwOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwOrientationSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwOrientationSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwOrientationSensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwOrientationSensorFromPointer(ptr unsafe.Pointer) (n *SensorfwOrientationSensor) {
	n = new(SensorfwOrientationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwOrientationSensor")
	return
}

func (ptr *SensorfwOrientationSensor) DestroySensorfwOrientationSensor() {
}

type SensorfwProximitySensor struct {
	SensorfwSensorBase
}

type SensorfwProximitySensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwProximitySensor_PTR() *SensorfwProximitySensor
}

func (ptr *SensorfwProximitySensor) SensorfwProximitySensor_PTR() *SensorfwProximitySensor {
	return ptr
}

func (ptr *SensorfwProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwProximitySensor(ptr SensorfwProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwProximitySensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwProximitySensorFromPointer(ptr unsafe.Pointer) (n *SensorfwProximitySensor) {
	n = new(SensorfwProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwProximitySensor")
	return
}

func (ptr *SensorfwProximitySensor) DestroySensorfwProximitySensor() {
}

type SensorfwRotationSensor struct {
	SensorfwSensorBase
}

type SensorfwRotationSensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwRotationSensor_PTR() *SensorfwRotationSensor
}

func (ptr *SensorfwRotationSensor) SensorfwRotationSensor_PTR() *SensorfwRotationSensor {
	return ptr
}

func (ptr *SensorfwRotationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwRotationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwRotationSensor(ptr SensorfwRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwRotationSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwRotationSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwRotationSensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwRotationSensorFromPointer(ptr unsafe.Pointer) (n *SensorfwRotationSensor) {
	n = new(SensorfwRotationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwRotationSensor")
	return
}

func (ptr *SensorfwRotationSensor) DestroySensorfwRotationSensor() {
}

type SensorfwSensorBase struct {
	QSensorBackend
}

type SensorfwSensorBase_ITF interface {
	QSensorBackend_ITF
	SensorfwSensorBase_PTR() *SensorfwSensorBase
}

func (ptr *SensorfwSensorBase) SensorfwSensorBase_PTR() *SensorfwSensorBase {
	return ptr
}

func (ptr *SensorfwSensorBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwSensorBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwSensorBase(ptr SensorfwSensorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwSensorBase) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwSensorBase) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewSensorfwSensorBaseFromPointer(ptr unsafe.Pointer) (n *SensorfwSensorBase) {
	n = new(SensorfwSensorBase)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwSensorBase")
	return
}

func (ptr *SensorfwSensorBase) DestroySensorfwSensorBase() {
}

type SensorfwTapSensor struct {
	SensorfwSensorBase
}

type SensorfwTapSensor_ITF interface {
	SensorfwSensorBase_ITF
	SensorfwTapSensor_PTR() *SensorfwTapSensor
}

func (ptr *SensorfwTapSensor) SensorfwTapSensor_PTR() *SensorfwTapSensor {
	return ptr
}

func (ptr *SensorfwTapSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorfwTapSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwTapSensor(ptr SensorfwTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwTapSensor_PTR().Pointer()
	}
	return nil
}

func (n *SensorfwTapSensor) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorfwTapSensor) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwTapSensorFromPointer(ptr unsafe.Pointer) (n *SensorfwTapSensor) {
	n = new(SensorfwTapSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorfwTapSensor")
	return
}

func (ptr *SensorfwTapSensor) DestroySensorfwTapSensor() {
}

type Sensorfwals struct {
	SensorfwSensorBase
}

type Sensorfwals_ITF interface {
	SensorfwSensorBase_ITF
	Sensorfwals_PTR() *Sensorfwals
}

func (ptr *Sensorfwals) Sensorfwals_PTR() *Sensorfwals {
	return ptr
}

func (ptr *Sensorfwals) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *Sensorfwals) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwals(ptr Sensorfwals_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Sensorfwals_PTR().Pointer()
	}
	return nil
}

func (n *Sensorfwals) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *Sensorfwals) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwalsFromPointer(ptr unsafe.Pointer) (n *Sensorfwals) {
	n = new(Sensorfwals)
	n.InitFromInternal(uintptr(ptr), "sensors.Sensorfwals")
	return
}

func (ptr *Sensorfwals) DestroySensorfwals() {
}

type SensorsConnection struct {
	core.QObject
}

type SensorsConnection_ITF interface {
	core.QObject_ITF
	SensorsConnection_PTR() *SensorsConnection
}

func (ptr *SensorsConnection) SensorsConnection_PTR() *SensorsConnection {
	return ptr
}

func (ptr *SensorsConnection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *SensorsConnection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromSensorsConnection(ptr SensorsConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorsConnection_PTR().Pointer()
	}
	return nil
}

func (n *SensorsConnection) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SensorsConnection) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewSensorsConnectionFromPointer(ptr unsafe.Pointer) (n *SensorsConnection) {
	n = new(SensorsConnection)
	n.InitFromInternal(uintptr(ptr), "sensors.SensorsConnection")
	return
}

func (ptr *SensorsConnection) DestroySensorsConnection() {
}

type SimulatorAccelerometer struct {
	SimulatorCommon
}

type SimulatorAccelerometer_ITF interface {
	SimulatorCommon_ITF
	SimulatorAccelerometer_PTR() *SimulatorAccelerometer
}

func (ptr *SimulatorAccelerometer) SimulatorAccelerometer_PTR() *SimulatorAccelerometer {
	return ptr
}

func (ptr *SimulatorAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorAccelerometer(ptr SimulatorAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorAccelerometer) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorAccelerometerFromPointer(ptr unsafe.Pointer) (n *SimulatorAccelerometer) {
	n = new(SimulatorAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorAccelerometer")
	return
}

func (ptr *SimulatorAccelerometer) DestroySimulatorAccelerometer() {
}

type SimulatorAmbientLightSensor struct {
	SimulatorCommon
}

type SimulatorAmbientLightSensor_ITF interface {
	SimulatorCommon_ITF
	SimulatorAmbientLightSensor_PTR() *SimulatorAmbientLightSensor
}

func (ptr *SimulatorAmbientLightSensor) SimulatorAmbientLightSensor_PTR() *SimulatorAmbientLightSensor {
	return ptr
}

func (ptr *SimulatorAmbientLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorAmbientLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorAmbientLightSensor(ptr SimulatorAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorAmbientLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorAmbientLightSensor) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorAmbientLightSensorFromPointer(ptr unsafe.Pointer) (n *SimulatorAmbientLightSensor) {
	n = new(SimulatorAmbientLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorAmbientLightSensor")
	return
}

func (ptr *SimulatorAmbientLightSensor) DestroySimulatorAmbientLightSensor() {
}

type SimulatorCommon struct {
	QSensorBackend
}

type SimulatorCommon_ITF interface {
	QSensorBackend_ITF
	SimulatorCommon_PTR() *SimulatorCommon
}

func (ptr *SimulatorCommon) SimulatorCommon_PTR() *SimulatorCommon {
	return ptr
}

func (ptr *SimulatorCommon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorCommon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorCommon(ptr SimulatorCommon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorCommon) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorCommon) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewSimulatorCommonFromPointer(ptr unsafe.Pointer) (n *SimulatorCommon) {
	n = new(SimulatorCommon)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorCommon")
	return
}

func (ptr *SimulatorCommon) DestroySimulatorCommon() {
}

type SimulatorCompass struct {
	SimulatorCommon
}

type SimulatorCompass_ITF interface {
	SimulatorCommon_ITF
	SimulatorCompass_PTR() *SimulatorCompass
}

func (ptr *SimulatorCompass) SimulatorCompass_PTR() *SimulatorCompass {
	return ptr
}

func (ptr *SimulatorCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorCompass(ptr SimulatorCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCompass_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorCompass) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorCompass) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorCompassFromPointer(ptr unsafe.Pointer) (n *SimulatorCompass) {
	n = new(SimulatorCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorCompass")
	return
}

func (ptr *SimulatorCompass) DestroySimulatorCompass() {
}

type SimulatorIRProximitySensor struct {
	SimulatorCommon
}

type SimulatorIRProximitySensor_ITF interface {
	SimulatorCommon_ITF
	SimulatorIRProximitySensor_PTR() *SimulatorIRProximitySensor
}

func (ptr *SimulatorIRProximitySensor) SimulatorIRProximitySensor_PTR() *SimulatorIRProximitySensor {
	return ptr
}

func (ptr *SimulatorIRProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorIRProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorIRProximitySensor(ptr SimulatorIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorIRProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorIRProximitySensor) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorIRProximitySensorFromPointer(ptr unsafe.Pointer) (n *SimulatorIRProximitySensor) {
	n = new(SimulatorIRProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorIRProximitySensor")
	return
}

func (ptr *SimulatorIRProximitySensor) DestroySimulatorIRProximitySensor() {
}

type SimulatorLightSensor struct {
	SimulatorCommon
}

type SimulatorLightSensor_ITF interface {
	SimulatorCommon_ITF
	SimulatorLightSensor_PTR() *SimulatorLightSensor
}

func (ptr *SimulatorLightSensor) SimulatorLightSensor_PTR() *SimulatorLightSensor {
	return ptr
}

func (ptr *SimulatorLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorLightSensor(ptr SimulatorLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorLightSensor) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorLightSensorFromPointer(ptr unsafe.Pointer) (n *SimulatorLightSensor) {
	n = new(SimulatorLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorLightSensor")
	return
}

func (ptr *SimulatorLightSensor) DestroySimulatorLightSensor() {
}

type SimulatorMagnetometer struct {
	SimulatorCommon
}

type SimulatorMagnetometer_ITF interface {
	SimulatorCommon_ITF
	SimulatorMagnetometer_PTR() *SimulatorMagnetometer
}

func (ptr *SimulatorMagnetometer) SimulatorMagnetometer_PTR() *SimulatorMagnetometer {
	return ptr
}

func (ptr *SimulatorMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorMagnetometer(ptr SimulatorMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorMagnetometer_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorMagnetometer) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorMagnetometer) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorMagnetometerFromPointer(ptr unsafe.Pointer) (n *SimulatorMagnetometer) {
	n = new(SimulatorMagnetometer)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorMagnetometer")
	return
}

func (ptr *SimulatorMagnetometer) DestroySimulatorMagnetometer() {
}

type SimulatorProximitySensor struct {
	SimulatorCommon
}

type SimulatorProximitySensor_ITF interface {
	SimulatorCommon_ITF
	SimulatorProximitySensor_PTR() *SimulatorProximitySensor
}

func (ptr *SimulatorProximitySensor) SimulatorProximitySensor_PTR() *SimulatorProximitySensor {
	return ptr
}

func (ptr *SimulatorProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorCommon_PTR().Pointer()
	}
	return nil
}

func (ptr *SimulatorProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SimulatorCommon_PTR().SetPointer(p)
	}
}

func PointerFromSimulatorProximitySensor(ptr SimulatorProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SimulatorProximitySensor_PTR().Pointer()
	}
	return nil
}

func (n *SimulatorProximitySensor) InitFromInternal(ptr uintptr, name string) {
	n.SimulatorCommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SimulatorProximitySensor) ClassNameInternalF() string {
	return n.SimulatorCommon_PTR().ClassNameInternalF()
}

func NewSimulatorProximitySensorFromPointer(ptr unsafe.Pointer) (n *SimulatorProximitySensor) {
	n = new(SimulatorProximitySensor)
	n.InitFromInternal(uintptr(ptr), "sensors.SimulatorProximitySensor")
	return
}

func (ptr *SimulatorProximitySensor) DestroySimulatorProximitySensor() {
}

type ThreadSafeSensorBackend struct {
	QSensorBackend
}

type ThreadSafeSensorBackend_ITF interface {
	QSensorBackend_ITF
	ThreadSafeSensorBackend_PTR() *ThreadSafeSensorBackend
}

func (ptr *ThreadSafeSensorBackend) ThreadSafeSensorBackend_PTR() *ThreadSafeSensorBackend {
	return ptr
}

func (ptr *ThreadSafeSensorBackend) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *ThreadSafeSensorBackend) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromThreadSafeSensorBackend(ptr ThreadSafeSensorBackend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ThreadSafeSensorBackend_PTR().Pointer()
	}
	return nil
}

func (n *ThreadSafeSensorBackend) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *ThreadSafeSensorBackend) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewThreadSafeSensorBackendFromPointer(ptr unsafe.Pointer) (n *ThreadSafeSensorBackend) {
	n = new(ThreadSafeSensorBackend)
	n.InitFromInternal(uintptr(ptr), "sensors.ThreadSafeSensorBackend")
	return
}

func (ptr *ThreadSafeSensorBackend) DestroyThreadSafeSensorBackend() {
}

type WinRtAccelerometer struct {
	QSensorBackend
}

type WinRtAccelerometer_ITF interface {
	QSensorBackend_ITF
	WinRtAccelerometer_PTR() *WinRtAccelerometer
}

func (ptr *WinRtAccelerometer) WinRtAccelerometer_PTR() *WinRtAccelerometer {
	return ptr
}

func (ptr *WinRtAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *WinRtAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromWinRtAccelerometer(ptr WinRtAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WinRtAccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *WinRtAccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *WinRtAccelerometer) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewWinRtAccelerometerFromPointer(ptr unsafe.Pointer) (n *WinRtAccelerometer) {
	n = new(WinRtAccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.WinRtAccelerometer")
	return
}

func (ptr *WinRtAccelerometer) DestroyWinRtAccelerometer() {
}

type WinRtAmbientLightSensor struct {
	QSensorBackend
}

type WinRtAmbientLightSensor_ITF interface {
	QSensorBackend_ITF
	WinRtAmbientLightSensor_PTR() *WinRtAmbientLightSensor
}

func (ptr *WinRtAmbientLightSensor) WinRtAmbientLightSensor_PTR() *WinRtAmbientLightSensor {
	return ptr
}

func (ptr *WinRtAmbientLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *WinRtAmbientLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromWinRtAmbientLightSensor(ptr WinRtAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WinRtAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func (n *WinRtAmbientLightSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *WinRtAmbientLightSensor) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewWinRtAmbientLightSensorFromPointer(ptr unsafe.Pointer) (n *WinRtAmbientLightSensor) {
	n = new(WinRtAmbientLightSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.WinRtAmbientLightSensor")
	return
}

func (ptr *WinRtAmbientLightSensor) DestroyWinRtAmbientLightSensor() {
}

type WinRtCompass struct {
	QSensorBackend
}

type WinRtCompass_ITF interface {
	QSensorBackend_ITF
	WinRtCompass_PTR() *WinRtCompass
}

func (ptr *WinRtCompass) WinRtCompass_PTR() *WinRtCompass {
	return ptr
}

func (ptr *WinRtCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *WinRtCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromWinRtCompass(ptr WinRtCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WinRtCompass_PTR().Pointer()
	}
	return nil
}

func (n *WinRtCompass) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *WinRtCompass) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewWinRtCompassFromPointer(ptr unsafe.Pointer) (n *WinRtCompass) {
	n = new(WinRtCompass)
	n.InitFromInternal(uintptr(ptr), "sensors.WinRtCompass")
	return
}

func (ptr *WinRtCompass) DestroyWinRtCompass() {
}

type WinRtGyroscope struct {
	QSensorBackend
}

type WinRtGyroscope_ITF interface {
	QSensorBackend_ITF
	WinRtGyroscope_PTR() *WinRtGyroscope
}

func (ptr *WinRtGyroscope) WinRtGyroscope_PTR() *WinRtGyroscope {
	return ptr
}

func (ptr *WinRtGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *WinRtGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromWinRtGyroscope(ptr WinRtGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WinRtGyroscope_PTR().Pointer()
	}
	return nil
}

func (n *WinRtGyroscope) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *WinRtGyroscope) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewWinRtGyroscopeFromPointer(ptr unsafe.Pointer) (n *WinRtGyroscope) {
	n = new(WinRtGyroscope)
	n.InitFromInternal(uintptr(ptr), "sensors.WinRtGyroscope")
	return
}

func (ptr *WinRtGyroscope) DestroyWinRtGyroscope() {
}

type WinRtOrientationSensor struct {
	QSensorBackend
}

type WinRtOrientationSensor_ITF interface {
	QSensorBackend_ITF
	WinRtOrientationSensor_PTR() *WinRtOrientationSensor
}

func (ptr *WinRtOrientationSensor) WinRtOrientationSensor_PTR() *WinRtOrientationSensor {
	return ptr
}

func (ptr *WinRtOrientationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *WinRtOrientationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromWinRtOrientationSensor(ptr WinRtOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WinRtOrientationSensor_PTR().Pointer()
	}
	return nil
}

func (n *WinRtOrientationSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *WinRtOrientationSensor) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewWinRtOrientationSensorFromPointer(ptr unsafe.Pointer) (n *WinRtOrientationSensor) {
	n = new(WinRtOrientationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.WinRtOrientationSensor")
	return
}

func (ptr *WinRtOrientationSensor) DestroyWinRtOrientationSensor() {
}

type WinRtRotationSensor struct {
	QSensorBackend
}

type WinRtRotationSensor_ITF interface {
	QSensorBackend_ITF
	WinRtRotationSensor_PTR() *WinRtRotationSensor
}

func (ptr *WinRtRotationSensor) WinRtRotationSensor_PTR() *WinRtRotationSensor {
	return ptr
}

func (ptr *WinRtRotationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *WinRtRotationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromWinRtRotationSensor(ptr WinRtRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WinRtRotationSensor_PTR().Pointer()
	}
	return nil
}

func (n *WinRtRotationSensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *WinRtRotationSensor) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewWinRtRotationSensorFromPointer(ptr unsafe.Pointer) (n *WinRtRotationSensor) {
	n = new(WinRtRotationSensor)
	n.InitFromInternal(uintptr(ptr), "sensors.WinRtRotationSensor")
	return
}

func (ptr *WinRtRotationSensor) DestroyWinRtRotationSensor() {
}

type dummyaccelerometer struct {
	dummycommon
}

type dummyaccelerometer_ITF interface {
	dummycommon_ITF
	dummyaccelerometer_PTR() *dummyaccelerometer
}

func (ptr *dummyaccelerometer) dummyaccelerometer_PTR() *dummyaccelerometer {
	return ptr
}

func (ptr *dummyaccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.dummycommon_PTR().Pointer()
	}
	return nil
}

func (ptr *dummyaccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.dummycommon_PTR().SetPointer(p)
	}
}

func PointerFromDummyaccelerometer(ptr dummyaccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.dummyaccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *dummyaccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.dummycommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *dummyaccelerometer) ClassNameInternalF() string {
	return n.dummycommon_PTR().ClassNameInternalF()
}

func NewDummyaccelerometerFromPointer(ptr unsafe.Pointer) (n *dummyaccelerometer) {
	n = new(dummyaccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.dummyaccelerometer")
	return
}

func (ptr *dummyaccelerometer) DestroyDummyaccelerometer() {
}

type dummycommon struct {
	QSensorBackend
}

type dummycommon_ITF interface {
	QSensorBackend_ITF
	dummycommon_PTR() *dummycommon
}

func (ptr *dummycommon) dummycommon_PTR() *dummycommon {
	return ptr
}

func (ptr *dummycommon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *dummycommon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromDummycommon(ptr dummycommon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.dummycommon_PTR().Pointer()
	}
	return nil
}

func (n *dummycommon) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *dummycommon) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewDummycommonFromPointer(ptr unsafe.Pointer) (n *dummycommon) {
	n = new(dummycommon)
	n.InitFromInternal(uintptr(ptr), "sensors.dummycommon")
	return
}

func (ptr *dummycommon) DestroyDummycommon() {
}

type dummylightsensor struct {
	dummycommon
}

type dummylightsensor_ITF interface {
	dummycommon_ITF
	dummylightsensor_PTR() *dummylightsensor
}

func (ptr *dummylightsensor) dummylightsensor_PTR() *dummylightsensor {
	return ptr
}

func (ptr *dummylightsensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.dummycommon_PTR().Pointer()
	}
	return nil
}

func (ptr *dummylightsensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.dummycommon_PTR().SetPointer(p)
	}
}

func PointerFromDummylightsensor(ptr dummylightsensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.dummylightsensor_PTR().Pointer()
	}
	return nil
}

func (n *dummylightsensor) InitFromInternal(ptr uintptr, name string) {
	n.dummycommon_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *dummylightsensor) ClassNameInternalF() string {
	return n.dummycommon_PTR().ClassNameInternalF()
}

func NewDummylightsensorFromPointer(ptr unsafe.Pointer) (n *dummylightsensor) {
	n = new(dummylightsensor)
	n.InitFromInternal(uintptr(ptr), "sensors.dummylightsensor")
	return
}

func (ptr *dummylightsensor) DestroyDummylightsensor() {
}

type genericalssensor struct {
	QLightFilter
	QSensorBackend
}

type genericalssensor_ITF interface {
	QLightFilter_ITF
	QSensorBackend_ITF
	genericalssensor_PTR() *genericalssensor
}

func (ptr *genericalssensor) genericalssensor_PTR() *genericalssensor {
	return ptr
}

func (ptr *genericalssensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *genericalssensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLightFilter_PTR().SetPointer(p)
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromGenericalssensor(ptr genericalssensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.genericalssensor_PTR().Pointer()
	}
	return nil
}

func (n *genericalssensor) InitFromInternal(ptr uintptr, name string) {
	n.QLightFilter_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *genericalssensor) ClassNameInternalF() string {
	return n.QLightFilter_PTR().ClassNameInternalF()
}

func NewGenericalssensorFromPointer(ptr unsafe.Pointer) (n *genericalssensor) {
	n = new(genericalssensor)
	n.InitFromInternal(uintptr(ptr), "sensors.genericalssensor")
	return
}

func (ptr *genericalssensor) DestroyGenericalssensor() {
}

type genericorientationsensor struct {
	QAccelerometerFilter
	QSensorBackend
}

type genericorientationsensor_ITF interface {
	QAccelerometerFilter_ITF
	QSensorBackend_ITF
	genericorientationsensor_PTR() *genericorientationsensor
}

func (ptr *genericorientationsensor) genericorientationsensor_PTR() *genericorientationsensor {
	return ptr
}

func (ptr *genericorientationsensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *genericorientationsensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAccelerometerFilter_PTR().SetPointer(p)
		ptr.QSensorBackend_PTR().SetPointer(p)
	}
}

func PointerFromGenericorientationsensor(ptr genericorientationsensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.genericorientationsensor_PTR().Pointer()
	}
	return nil
}

func (n *genericorientationsensor) InitFromInternal(ptr uintptr, name string) {
	n.QAccelerometerFilter_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *genericorientationsensor) ClassNameInternalF() string {
	return n.QAccelerometerFilter_PTR().ClassNameInternalF()
}

func NewGenericorientationsensorFromPointer(ptr unsafe.Pointer) (n *genericorientationsensor) {
	n = new(genericorientationsensor)
	n.InitFromInternal(uintptr(ptr), "sensors.genericorientationsensor")
	return
}

func (ptr *genericorientationsensor) DestroyGenericorientationsensor() {
}

type genericrotationsensor struct {
	QSensorBackend
	QSensorFilter
}

type genericrotationsensor_ITF interface {
	QSensorBackend_ITF
	QSensorFilter_ITF
	genericrotationsensor_PTR() *genericrotationsensor
}

func (ptr *genericrotationsensor) genericrotationsensor_PTR() *genericrotationsensor {
	return ptr
}

func (ptr *genericrotationsensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func (ptr *genericrotationsensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorBackend_PTR().SetPointer(p)
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromGenericrotationsensor(ptr genericrotationsensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.genericrotationsensor_PTR().Pointer()
	}
	return nil
}

func (n *genericrotationsensor) InitFromInternal(ptr uintptr, name string) {
	n.QSensorBackend_PTR().InitFromInternal(uintptr(ptr), name)
	n.QSensorFilter_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *genericrotationsensor) ClassNameInternalF() string {
	return n.QSensorBackend_PTR().ClassNameInternalF()
}

func NewGenericrotationsensorFromPointer(ptr unsafe.Pointer) (n *genericrotationsensor) {
	n = new(genericrotationsensor)
	n.InitFromInternal(uintptr(ptr), "sensors.genericrotationsensor")
	return
}

func (ptr *genericrotationsensor) DestroyGenericrotationsensor() {
}

type sensorfwaccelerometer struct {
	SensorfwSensorBase
}

type sensorfwaccelerometer_ITF interface {
	SensorfwSensorBase_ITF
	sensorfwaccelerometer_PTR() *sensorfwaccelerometer
}

func (ptr *sensorfwaccelerometer) sensorfwaccelerometer_PTR() *sensorfwaccelerometer {
	return ptr
}

func (ptr *sensorfwaccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.SensorfwSensorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *sensorfwaccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.SensorfwSensorBase_PTR().SetPointer(p)
	}
}

func PointerFromSensorfwaccelerometer(ptr sensorfwaccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.sensorfwaccelerometer_PTR().Pointer()
	}
	return nil
}

func (n *sensorfwaccelerometer) InitFromInternal(ptr uintptr, name string) {
	n.SensorfwSensorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *sensorfwaccelerometer) ClassNameInternalF() string {
	return n.SensorfwSensorBase_PTR().ClassNameInternalF()
}

func NewSensorfwaccelerometerFromPointer(ptr unsafe.Pointer) (n *sensorfwaccelerometer) {
	n = new(sensorfwaccelerometer)
	n.InitFromInternal(uintptr(ptr), "sensors.sensorfwaccelerometer")
	return
}

func (ptr *sensorfwaccelerometer) DestroySensorfwaccelerometer() {
}

func init() {
	internal.ConstructorTable["sensors.QAccelerometer"] = NewQAccelerometerFromPointer
	internal.ConstructorTable["sensors.QAccelerometerFilter"] = NewQAccelerometerFilterFromPointer
	internal.ConstructorTable["sensors.QAccelerometerReading"] = NewQAccelerometerReadingFromPointer
	internal.ConstructorTable["sensors.QAltimeter"] = NewQAltimeterFromPointer
	internal.ConstructorTable["sensors.QAltimeterFilter"] = NewQAltimeterFilterFromPointer
	internal.ConstructorTable["sensors.QAltimeterReading"] = NewQAltimeterReadingFromPointer
	internal.ConstructorTable["sensors.QAmbientLightFilter"] = NewQAmbientLightFilterFromPointer
	internal.ConstructorTable["sensors.QAmbientLightReading"] = NewQAmbientLightReadingFromPointer
	internal.ConstructorTable["sensors.QAmbientLightSensor"] = NewQAmbientLightSensorFromPointer
	internal.ConstructorTable["sensors.QAmbientTemperatureFilter"] = NewQAmbientTemperatureFilterFromPointer
	internal.ConstructorTable["sensors.QAmbientTemperatureReading"] = NewQAmbientTemperatureReadingFromPointer
	internal.ConstructorTable["sensors.QAmbientTemperatureSensor"] = NewQAmbientTemperatureSensorFromPointer
	internal.ConstructorTable["sensors.QCompass"] = NewQCompassFromPointer
	internal.ConstructorTable["sensors.QCompassFilter"] = NewQCompassFilterFromPointer
	internal.ConstructorTable["sensors.QCompassReading"] = NewQCompassReadingFromPointer
	internal.ConstructorTable["sensors.QDistanceFilter"] = NewQDistanceFilterFromPointer
	internal.ConstructorTable["sensors.QDistanceReading"] = NewQDistanceReadingFromPointer
	internal.ConstructorTable["sensors.QDistanceSensor"] = NewQDistanceSensorFromPointer
	internal.ConstructorTable["sensors.QGyroscope"] = NewQGyroscopeFromPointer
	internal.ConstructorTable["sensors.QGyroscopeFilter"] = NewQGyroscopeFilterFromPointer
	internal.ConstructorTable["sensors.QGyroscopeReading"] = NewQGyroscopeReadingFromPointer
	internal.ConstructorTable["sensors.QHolsterFilter"] = NewQHolsterFilterFromPointer
	internal.ConstructorTable["sensors.QHolsterReading"] = NewQHolsterReadingFromPointer
	internal.ConstructorTable["sensors.QHolsterSensor"] = NewQHolsterSensorFromPointer
	internal.ConstructorTable["sensors.QHumidityFilter"] = NewQHumidityFilterFromPointer
	internal.ConstructorTable["sensors.QHumidityReading"] = NewQHumidityReadingFromPointer
	internal.ConstructorTable["sensors.QHumiditySensor"] = NewQHumiditySensorFromPointer
	internal.ConstructorTable["sensors.QIRProximityFilter"] = NewQIRProximityFilterFromPointer
	internal.ConstructorTable["sensors.QIRProximityReading"] = NewQIRProximityReadingFromPointer
	internal.ConstructorTable["sensors.QIRProximitySensor"] = NewQIRProximitySensorFromPointer
	internal.ConstructorTable["sensors.QLidFilter"] = NewQLidFilterFromPointer
	internal.ConstructorTable["sensors.QLidReading"] = NewQLidReadingFromPointer
	internal.ConstructorTable["sensors.QLidSensor"] = NewQLidSensorFromPointer
	internal.ConstructorTable["sensors.QLightFilter"] = NewQLightFilterFromPointer
	internal.ConstructorTable["sensors.QLightReading"] = NewQLightReadingFromPointer
	internal.ConstructorTable["sensors.QLightSensor"] = NewQLightSensorFromPointer
	internal.ConstructorTable["sensors.QMagnetometer"] = NewQMagnetometerFromPointer
	internal.ConstructorTable["sensors.QMagnetometerFilter"] = NewQMagnetometerFilterFromPointer
	internal.ConstructorTable["sensors.QMagnetometerReading"] = NewQMagnetometerReadingFromPointer
	internal.ConstructorTable["sensors.QOrientationFilter"] = NewQOrientationFilterFromPointer
	internal.ConstructorTable["sensors.QOrientationReading"] = NewQOrientationReadingFromPointer
	internal.ConstructorTable["sensors.QOrientationSensor"] = NewQOrientationSensorFromPointer
	internal.ConstructorTable["sensors.QPressureFilter"] = NewQPressureFilterFromPointer
	internal.ConstructorTable["sensors.QPressureReading"] = NewQPressureReadingFromPointer
	internal.ConstructorTable["sensors.QPressureSensor"] = NewQPressureSensorFromPointer
	internal.ConstructorTable["sensors.QProximityFilter"] = NewQProximityFilterFromPointer
	internal.ConstructorTable["sensors.QProximityReading"] = NewQProximityReadingFromPointer
	internal.ConstructorTable["sensors.QProximitySensor"] = NewQProximitySensorFromPointer
	internal.ConstructorTable["sensors.QRotationFilter"] = NewQRotationFilterFromPointer
	internal.ConstructorTable["sensors.QRotationReading"] = NewQRotationReadingFromPointer
	internal.ConstructorTable["sensors.QRotationSensor"] = NewQRotationSensorFromPointer
	internal.ConstructorTable["sensors.QSensor"] = NewQSensorFromPointer
	internal.ConstructorTable["sensors.QSensorBackend"] = NewQSensorBackendFromPointer
	internal.ConstructorTable["sensors.QSensorBackendFactory"] = NewQSensorBackendFactoryFromPointer
	internal.ConstructorTable["sensors.QSensorChangesInterface"] = NewQSensorChangesInterfaceFromPointer
	internal.ConstructorTable["sensors.QSensorFilter"] = NewQSensorFilterFromPointer
	internal.ConstructorTable["sensors.QSensorGesture"] = NewQSensorGestureFromPointer
	internal.ConstructorTable["sensors.QSensorGestureManager"] = NewQSensorGestureManagerFromPointer
	internal.ConstructorTable["sensors.QSensorGesturePluginInterface"] = NewQSensorGesturePluginInterfaceFromPointer
	internal.ConstructorTable["sensors.QSensorGestureRecognizer"] = NewQSensorGestureRecognizerFromPointer
	internal.ConstructorTable["sensors.QSensorManager"] = NewQSensorManagerFromPointer
	internal.ConstructorTable["sensors.QSensorPluginInterface"] = NewQSensorPluginInterfaceFromPointer
	internal.ConstructorTable["sensors.QSensorReading"] = NewQSensorReadingFromPointer
	internal.ConstructorTable["sensors.QTapFilter"] = NewQTapFilterFromPointer
	internal.ConstructorTable["sensors.QTapReading"] = NewQTapReadingFromPointer
	internal.ConstructorTable["sensors.QTapSensor"] = NewQTapSensorFromPointer
	internal.ConstructorTable["sensors.QTiltFilter"] = NewQTiltFilterFromPointer
	internal.ConstructorTable["sensors.QTiltReading"] = NewQTiltReadingFromPointer
	internal.ConstructorTable["sensors.QTiltSensor"] = NewQTiltSensorFromPointer
}
