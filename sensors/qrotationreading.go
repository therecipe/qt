package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

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
		n.SetObjectName("QRotationReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRotationReading) QRotationReading_PTR() *QRotationReading {
	return ptr
}

func (ptr *QRotationReading) X() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationReading::x")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Y() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationReading::y")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Z() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationReading::z")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) SetFromEuler(x float64, y float64, z float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationReading::setFromEuler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRotationReading_SetFromEuler(ptr.Pointer(), C.double(x), C.double(y), C.double(z))
	}
}
