package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QProximityReading struct {
	QSensorReading
}

type QProximityReading_ITF interface {
	QSensorReading_ITF
	QProximityReading_PTR() *QProximityReading
}

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) *QProximityReading {
	var n = new(QProximityReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProximityReading_") {
		n.SetObjectName("QProximityReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Close() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProximityReading::close")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProximityReading_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProximityReading::setClose")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}
