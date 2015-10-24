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

type QCompassReadingITF interface {
	QSensorReadingITF
	QCompassReadingPTR() *QCompassReading
}

func PointerFromQCompassReading(ptr QCompassReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassReadingPTR().Pointer()
	}
	return nil
}

func QCompassReadingFromPointer(ptr unsafe.Pointer) *QCompassReading {
	var n = new(QCompassReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCompassReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompassReading) QCompassReadingPTR() *QCompassReading {
	return ptr
}
