package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensorReading struct {
	core.QObject
}

type QSensorReading_ITF interface {
	core.QObject_ITF
	QSensorReading_PTR() *QSensorReading
}

func PointerFromQSensorReading(ptr QSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func NewQSensorReadingFromPointer(ptr unsafe.Pointer) *QSensorReading {
	var n = new(QSensorReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorReading_") {
		n.SetObjectName("QSensorReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorReading) QSensorReading_PTR() *QSensorReading {
	return ptr
}

func (ptr *QSensorReading) Value(index int) *core.QVariant {
	defer qt.Recovering("QSensorReading::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSensorReading_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSensorReading) ValueCount() int {
	defer qt.Recovering("QSensorReading::valueCount")

	if ptr.Pointer() != nil {
		return int(C.QSensorReading_ValueCount(ptr.Pointer()))
	}
	return 0
}
