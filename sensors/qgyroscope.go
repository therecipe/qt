package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGyroscope struct {
	QSensor
}

type QGyroscope_ITF interface {
	QSensor_ITF
	QGyroscope_PTR() *QGyroscope
}

func PointerFromQGyroscope(ptr QGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscope_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeFromPointer(ptr unsafe.Pointer) *QGyroscope {
	var n = new(QGyroscope)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGyroscope_") {
		n.SetObjectName("QGyroscope_" + qt.Identifier())
	}
	return n
}

func (ptr *QGyroscope) QGyroscope_PTR() *QGyroscope {
	return ptr
}

func (ptr *QGyroscope) Reading() *QGyroscopeReading {
	defer qt.Recovering("QGyroscope::reading")

	if ptr.Pointer() != nil {
		return NewQGyroscopeReadingFromPointer(C.QGyroscope_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQGyroscope(parent core.QObject_ITF) *QGyroscope {
	defer qt.Recovering("QGyroscope::QGyroscope")

	return NewQGyroscopeFromPointer(C.QGyroscope_NewQGyroscope(core.PointerFromQObject(parent)))
}

func (ptr *QGyroscope) DestroyQGyroscope() {
	defer qt.Recovering("QGyroscope::~QGyroscope")

	if ptr.Pointer() != nil {
		C.QGyroscope_DestroyQGyroscope(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
