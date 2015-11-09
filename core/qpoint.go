package core

//#include "qpoint.h"
import "C"
import (
	"unsafe"
)

type QPoint struct {
	ptr unsafe.Pointer
}

type QPoint_ITF interface {
	QPoint_PTR() *QPoint
}

func (p *QPoint) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPoint) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPoint(ptr QPoint_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPoint_PTR().Pointer()
	}
	return nil
}

func NewQPointFromPointer(ptr unsafe.Pointer) *QPoint {
	var n = new(QPoint)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPoint) QPoint_PTR() *QPoint {
	return ptr
}

func NewQPoint() *QPoint {
	return NewQPointFromPointer(C.QPoint_NewQPoint())
}

func NewQPoint2(xpos int, ypos int) *QPoint {
	return NewQPointFromPointer(C.QPoint_NewQPoint2(C.int(xpos), C.int(ypos)))
}

func QPoint_DotProduct(p1 QPoint_ITF, p2 QPoint_ITF) int {
	return int(C.QPoint_QPoint_DotProduct(PointerFromQPoint(p1), PointerFromQPoint(p2)))
}

func (ptr *QPoint) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPoint_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPoint) ManhattanLength() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_ManhattanLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPoint) Rx() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_Rx(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPoint) Ry() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_Ry(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPoint) SetX(x int) {
	if ptr.Pointer() != nil {
		C.QPoint_SetX(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QPoint) SetY(y int) {
	if ptr.Pointer() != nil {
		C.QPoint_SetY(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QPoint) X() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPoint) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_Y(ptr.Pointer()))
	}
	return 0
}
