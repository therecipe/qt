package sensors

//#include "qholsterreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QHolsterReading struct {
	QSensorReading
}

type QHolsterReadingITF interface {
	QSensorReadingITF
	QHolsterReadingPTR() *QHolsterReading
}

func PointerFromQHolsterReading(ptr QHolsterReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterReadingPTR().Pointer()
	}
	return nil
}

func QHolsterReadingFromPointer(ptr unsafe.Pointer) *QHolsterReading {
	var n = new(QHolsterReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHolsterReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHolsterReading) QHolsterReadingPTR() *QHolsterReading {
	return ptr
}

func (ptr *QHolsterReading) Holstered() bool {
	if ptr.Pointer() != nil {
		return C.QHolsterReading_Holstered(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHolsterReading) SetHolstered(holstered bool) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_SetHolstered(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(holstered)))
	}
}
