package gui

//#include "qradialgradient.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRadialGradient struct {
	QGradient
}

type QRadialGradientITF interface {
	QGradientITF
	QRadialGradientPTR() *QRadialGradient
}

func PointerFromQRadialGradient(ptr QRadialGradientITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadialGradientPTR().Pointer()
	}
	return nil
}

func QRadialGradientFromPointer(ptr unsafe.Pointer) *QRadialGradient {
	var n = new(QRadialGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRadialGradient) QRadialGradientPTR() *QRadialGradient {
	return ptr
}

func NewQRadialGradient() *QRadialGradient {
	return QRadialGradientFromPointer(unsafe.Pointer(C.QRadialGradient_NewQRadialGradient()))
}

func (ptr *QRadialGradient) SetCenter(center core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRadialGradient_SetCenter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(center)))
	}
}

func (ptr *QRadialGradient) SetFocalPoint(focalPoint core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QRadialGradient_SetFocalPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(focalPoint)))
	}
}
