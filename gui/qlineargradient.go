package gui

//#include "qlineargradient.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLinearGradient struct {
	QGradient
}

type QLinearGradientITF interface {
	QGradientITF
	QLinearGradientPTR() *QLinearGradient
}

func PointerFromQLinearGradient(ptr QLinearGradientITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinearGradientPTR().Pointer()
	}
	return nil
}

func QLinearGradientFromPointer(ptr unsafe.Pointer) *QLinearGradient {
	var n = new(QLinearGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLinearGradient) QLinearGradientPTR() *QLinearGradient {
	return ptr
}

func NewQLinearGradient() *QLinearGradient {
	return QLinearGradientFromPointer(unsafe.Pointer(C.QLinearGradient_NewQLinearGradient()))
}

func NewQLinearGradient2(start core.QPointFITF, finalStop core.QPointFITF) *QLinearGradient {
	return QLinearGradientFromPointer(unsafe.Pointer(C.QLinearGradient_NewQLinearGradient2(C.QtObjectPtr(core.PointerFromQPointF(start)), C.QtObjectPtr(core.PointerFromQPointF(finalStop)))))
}

func (ptr *QLinearGradient) SetFinalStop(stop core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QLinearGradient_SetFinalStop(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(stop)))
	}
}

func (ptr *QLinearGradient) SetStart(start core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QLinearGradient_SetStart(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(start)))
	}
}
