package sensors

//#include "qcompassreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

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
	if n.ObjectName() == "" {
		n.SetObjectName("QCompassReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return ptr
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
