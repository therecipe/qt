package sensors

//#include "qgyroscopereading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGyroscopeReading struct {
	QSensorReading
}

type QGyroscopeReadingITF interface {
	QSensorReadingITF
	QGyroscopeReadingPTR() *QGyroscopeReading
}

func PointerFromQGyroscopeReading(ptr QGyroscopeReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReadingPTR().Pointer()
	}
	return nil
}

func QGyroscopeReadingFromPointer(ptr unsafe.Pointer) *QGyroscopeReading {
	var n = new(QGyroscopeReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGyroscopeReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGyroscopeReading) QGyroscopeReadingPTR() *QGyroscopeReading {
	return ptr
}
