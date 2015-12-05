package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

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
		n.SetObjectName("QIRProximityReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIRProximityReading) QIRProximityReading_PTR() *QIRProximityReading {
	return ptr
}

func (ptr *QIRProximityReading) Reflectance() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIRProximityReading::reflectance")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QIRProximityReading_Reflectance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIRProximityReading) SetReflectance(reflectance float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIRProximityReading::setReflectance")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIRProximityReading_SetReflectance(ptr.Pointer(), C.double(reflectance))
	}
}
