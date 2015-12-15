package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLinearGradient struct {
	QGradient
}

type QLinearGradient_ITF interface {
	QGradient_ITF
	QLinearGradient_PTR() *QLinearGradient
}

func PointerFromQLinearGradient(ptr QLinearGradient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinearGradient_PTR().Pointer()
	}
	return nil
}

func NewQLinearGradientFromPointer(ptr unsafe.Pointer) *QLinearGradient {
	var n = new(QLinearGradient)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLinearGradient) QLinearGradient_PTR() *QLinearGradient {
	return ptr
}

func NewQLinearGradient3(x1 float64, y1 float64, x2 float64, y2 float64) *QLinearGradient {
	defer qt.Recovering("QLinearGradient::QLinearGradient")

	return NewQLinearGradientFromPointer(C.QLinearGradient_NewQLinearGradient3(C.double(x1), C.double(y1), C.double(x2), C.double(y2)))
}

func NewQLinearGradient() *QLinearGradient {
	defer qt.Recovering("QLinearGradient::QLinearGradient")

	return NewQLinearGradientFromPointer(C.QLinearGradient_NewQLinearGradient())
}

func NewQLinearGradient2(start core.QPointF_ITF, finalStop core.QPointF_ITF) *QLinearGradient {
	defer qt.Recovering("QLinearGradient::QLinearGradient")

	return NewQLinearGradientFromPointer(C.QLinearGradient_NewQLinearGradient2(core.PointerFromQPointF(start), core.PointerFromQPointF(finalStop)))
}

func (ptr *QLinearGradient) SetFinalStop(stop core.QPointF_ITF) {
	defer qt.Recovering("QLinearGradient::setFinalStop")

	if ptr.Pointer() != nil {
		C.QLinearGradient_SetFinalStop(ptr.Pointer(), core.PointerFromQPointF(stop))
	}
}

func (ptr *QLinearGradient) SetFinalStop2(x float64, y float64) {
	defer qt.Recovering("QLinearGradient::setFinalStop")

	if ptr.Pointer() != nil {
		C.QLinearGradient_SetFinalStop2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QLinearGradient) SetStart(start core.QPointF_ITF) {
	defer qt.Recovering("QLinearGradient::setStart")

	if ptr.Pointer() != nil {
		C.QLinearGradient_SetStart(ptr.Pointer(), core.PointerFromQPointF(start))
	}
}

func (ptr *QLinearGradient) SetStart2(x float64, y float64) {
	defer qt.Recovering("QLinearGradient::setStart")

	if ptr.Pointer() != nil {
		C.QLinearGradient_SetStart2(ptr.Pointer(), C.double(x), C.double(y))
	}
}
