package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QCompassReading_") {
		n.SetObjectName("QCompassReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return ptr
}

func (ptr *QCompassReading) Azimuth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompassReading::azimuth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_Azimuth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) CalibrationLevel() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompassReading::calibrationLevel")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) SetAzimuth(azimuth float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompassReading::setAzimuth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompassReading_SetAzimuth(ptr.Pointer(), C.double(azimuth))
	}
}

func (ptr *QCompassReading) SetCalibrationLevel(calibrationLevel float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompassReading::setCalibrationLevel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompassReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}
