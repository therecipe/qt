package widgets

//#include "qdial.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDial struct {
	QAbstractSlider
}

type QDial_ITF interface {
	QAbstractSlider_ITF
	QDial_PTR() *QDial
}

func PointerFromQDial(ptr QDial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDial_PTR().Pointer()
	}
	return nil
}

func NewQDialFromPointer(ptr unsafe.Pointer) *QDial {
	var n = new(QDial)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QDial_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDial) QDial_PTR() *QDial {
	return ptr
}

func (ptr *QDial) NotchSize() int {
	if ptr.Pointer() != nil {
		return int(C.QDial_NotchSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchTarget() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QDial_NotchTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchesVisible() bool {
	if ptr.Pointer() != nil {
		return C.QDial_NotchesVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDial) SetNotchesVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QDial_SetNotchesVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) SetWrapping(on bool) {
	if ptr.Pointer() != nil {
		C.QDial_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QDial) Wrapping() bool {
	if ptr.Pointer() != nil {
		return C.QDial_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQDial(parent QWidget_ITF) *QDial {
	return NewQDialFromPointer(C.QDial_NewQDial(PointerFromQWidget(parent)))
}

func (ptr *QDial) DestroyQDial() {
	if ptr.Pointer() != nil {
		C.QDial_DestroyQDial(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
