package gui

//#include "qconicalgradient.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QConicalGradient struct {
	QGradient
}

type QConicalGradientITF interface {
	QGradientITF
	QConicalGradientPTR() *QConicalGradient
}

func PointerFromQConicalGradient(ptr QConicalGradientITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QConicalGradientPTR().Pointer()
	}
	return nil
}

func QConicalGradientFromPointer(ptr unsafe.Pointer) *QConicalGradient {
	var n = new(QConicalGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QConicalGradient) QConicalGradientPTR() *QConicalGradient {
	return ptr
}

func NewQConicalGradient() *QConicalGradient {
	return QConicalGradientFromPointer(unsafe.Pointer(C.QConicalGradient_NewQConicalGradient()))
}

func (ptr *QConicalGradient) SetCenter(center core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QConicalGradient_SetCenter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(center)))
	}
}
