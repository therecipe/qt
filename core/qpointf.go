package core

//#include "qpointf.h"
import "C"
import (
	"unsafe"
)

type QPointF struct {
	ptr unsafe.Pointer
}

type QPointF_ITF interface {
	QPointF_PTR() *QPointF
}

func (p *QPointF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPointF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPointF(ptr QPointF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPointF_PTR().Pointer()
	}
	return nil
}

func NewQPointFFromPointer(ptr unsafe.Pointer) *QPointF {
	var n = new(QPointF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPointF) QPointF_PTR() *QPointF {
	return ptr
}

func NewQPointF() *QPointF {
	return NewQPointFFromPointer(C.QPointF_NewQPointF())
}

func NewQPointF2(point QPoint_ITF) *QPointF {
	return NewQPointFFromPointer(C.QPointF_NewQPointF2(PointerFromQPoint(point)))
}

func NewQPointF3(xpos float64, ypos float64) *QPointF {
	return NewQPointFFromPointer(C.QPointF_NewQPointF3(C.double(xpos), C.double(ypos)))
}

func QPointF_DotProduct(p1 QPointF_ITF, p2 QPointF_ITF) float64 {
	return float64(C.QPointF_QPointF_DotProduct(PointerFromQPointF(p1), PointerFromQPointF(p2)))
}

func (ptr *QPointF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPointF_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPointF) ManhattanLength() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_ManhattanLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPointF) Rx() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_Rx(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPointF) Ry() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_Ry(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPointF) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QPointF_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QPointF) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QPointF_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QPointF) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPointF) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_Y(ptr.Pointer()))
	}
	return 0
}
