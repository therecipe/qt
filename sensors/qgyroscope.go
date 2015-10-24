package sensors

//#include "qgyroscope.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGyroscope struct {
	QSensor
}

type QGyroscopeITF interface {
	QSensorITF
	QGyroscopePTR() *QGyroscope
}

func PointerFromQGyroscope(ptr QGyroscopeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopePTR().Pointer()
	}
	return nil
}

func QGyroscopeFromPointer(ptr unsafe.Pointer) *QGyroscope {
	var n = new(QGyroscope)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGyroscope_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGyroscope) QGyroscopePTR() *QGyroscope {
	return ptr
}

func (ptr *QGyroscope) Reading() *QGyroscopeReading {
	if ptr.Pointer() != nil {
		return QGyroscopeReadingFromPointer(unsafe.Pointer(C.QGyroscope_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQGyroscope(parent core.QObjectITF) *QGyroscope {
	return QGyroscopeFromPointer(unsafe.Pointer(C.QGyroscope_NewQGyroscope(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGyroscope) DestroyQGyroscope() {
	if ptr.Pointer() != nil {
		C.QGyroscope_DestroyQGyroscope(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
