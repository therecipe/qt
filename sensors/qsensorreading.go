package sensors

//#include "qsensorreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensorReading struct {
	core.QObject
}

type QSensorReadingITF interface {
	core.QObjectITF
	QSensorReadingPTR() *QSensorReading
}

func PointerFromQSensorReading(ptr QSensorReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReadingPTR().Pointer()
	}
	return nil
}

func QSensorReadingFromPointer(ptr unsafe.Pointer) *QSensorReading {
	var n = new(QSensorReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSensorReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensorReading) QSensorReadingPTR() *QSensorReading {
	return ptr
}

func (ptr *QSensorReading) Value(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorReading_Value(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QSensorReading) ValueCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSensorReading_ValueCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
