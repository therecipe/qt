package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

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
