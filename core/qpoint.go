package core

//#include "qpoint.h"
import "C"
import (
	"unsafe"
)

type QPoint struct {
	ptr unsafe.Pointer
}

type QPointITF interface {
	QPointPTR() *QPoint
}

func (p *QPoint) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPoint) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPoint(ptr QPointITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPointPTR().Pointer()
	}
	return nil
}

func QPointFromPointer(ptr unsafe.Pointer) *QPoint {
	var n = new(QPoint)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPoint) QPointPTR() *QPoint {
	return ptr
}

func NewQPoint() *QPoint {
	return QPointFromPointer(unsafe.Pointer(C.QPoint_NewQPoint()))
}

func NewQPoint2(xpos int, ypos int) *QPoint {
	return QPointFromPointer(unsafe.Pointer(C.QPoint_NewQPoint2(C.int(xpos), C.int(ypos))))
}

func QPoint_DotProduct(p1 QPointITF, p2 QPointITF) int {
	return int(C.QPoint_QPoint_DotProduct(C.QtObjectPtr(PointerFromQPoint(p1)), C.QtObjectPtr(PointerFromQPoint(p2))))
}

func (ptr *QPoint) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPoint_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPoint) ManhattanLength() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_ManhattanLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPoint) Rx() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_Rx(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPoint) Ry() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_Ry(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPoint) SetX(x int) {
	if ptr.Pointer() != nil {
		C.QPoint_SetX(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QPoint) SetY(y int) {
	if ptr.Pointer() != nil {
		C.QPoint_SetY(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QPoint) X() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPoint) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QPoint_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
