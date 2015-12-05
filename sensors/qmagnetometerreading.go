package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

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
		n.SetObjectName("QMagnetometerReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMagnetometerReading) QMagnetometerReading_PTR() *QMagnetometerReading {
	return ptr
}

func (ptr *QMagnetometerReading) CalibrationLevel() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::calibrationLevel")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) X() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::x")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Y() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::y")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Z() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::z")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) SetCalibrationLevel(calibrationLevel float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::setCalibrationLevel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QMagnetometerReading) SetX(x float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::setX")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QMagnetometerReading) SetY(y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::setY")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QMagnetometerReading) SetZ(z float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometerReading::setZ")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}
