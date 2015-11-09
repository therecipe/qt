package sensors

//#include "qorientationreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOrientationReading struct {
	QSensorReading
}

type QOrientationReading_ITF interface {
	QSensorReading_ITF
	QOrientationReading_PTR() *QOrientationReading
}

func PointerFromQOrientationReading(ptr QOrientationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationReading_PTR().Pointer()
	}
	return nil
}

func NewQOrientationReadingFromPointer(ptr unsafe.Pointer) *QOrientationReading {
	var n = new(QOrientationReading)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QOrientationReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOrientationReading) QOrientationReading_PTR() *QOrientationReading {
	return ptr
}

//QOrientationReading::Orientation
type QOrientationReading__Orientation int64

const (
	QOrientationReading__Undefined = QOrientationReading__Orientation(0)
	QOrientationReading__TopUp     = QOrientationReading__Orientation(1)
	QOrientationReading__TopDown   = QOrientationReading__Orientation(2)
	QOrientationReading__LeftUp    = QOrientationReading__Orientation(3)
	QOrientationReading__RightUp   = QOrientationReading__Orientation(4)
	QOrientationReading__FaceUp    = QOrientationReading__Orientation(5)
	QOrientationReading__FaceDown  = QOrientationReading__Orientation(6)
)

func (ptr *QOrientationReading) Orientation() QOrientationReading__Orientation {
	if ptr.Pointer() != nil {
		return QOrientationReading__Orientation(C.QOrientationReading_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}
