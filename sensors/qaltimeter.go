package sensors

//#include "qaltimeter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAltimeter struct {
	QSensor
}

type QAltimeterITF interface {
	QSensorITF
	QAltimeterPTR() *QAltimeter
}

func PointerFromQAltimeter(ptr QAltimeterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterPTR().Pointer()
	}
	return nil
}

func QAltimeterFromPointer(ptr unsafe.Pointer) *QAltimeter {
	var n = new(QAltimeter)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAltimeter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAltimeter) QAltimeterPTR() *QAltimeter {
	return ptr
}

func (ptr *QAltimeter) Reading() *QAltimeterReading {
	if ptr.Pointer() != nil {
		return QAltimeterReadingFromPointer(unsafe.Pointer(C.QAltimeter_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQAltimeter(parent core.QObjectITF) *QAltimeter {
	return QAltimeterFromPointer(unsafe.Pointer(C.QAltimeter_NewQAltimeter(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAltimeter) DestroyQAltimeter() {
	if ptr.Pointer() != nil {
		C.QAltimeter_DestroyQAltimeter(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
