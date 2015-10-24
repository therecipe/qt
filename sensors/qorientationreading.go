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

type QOrientationReadingITF interface {
	QSensorReadingITF
	QOrientationReadingPTR() *QOrientationReading
}

func PointerFromQOrientationReading(ptr QOrientationReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationReadingPTR().Pointer()
	}
	return nil
}

func QOrientationReadingFromPointer(ptr unsafe.Pointer) *QOrientationReading {
	var n = new(QOrientationReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOrientationReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOrientationReading) QOrientationReadingPTR() *QOrientationReading {
	return ptr
}

//QOrientationReading::Orientation
type QOrientationReading__Orientation int

var (
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
		return QOrientationReading__Orientation(C.QOrientationReading_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}
