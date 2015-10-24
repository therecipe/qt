package core

//#include "qlinef.h"
import "C"
import (
	"unsafe"
)

type QLineF struct {
	ptr unsafe.Pointer
}

type QLineFITF interface {
	QLineFPTR() *QLineF
}

func (p *QLineF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLineF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLineF(ptr QLineFITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineFPTR().Pointer()
	}
	return nil
}

func QLineFFromPointer(ptr unsafe.Pointer) *QLineF {
	var n = new(QLineF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLineF) QLineFPTR() *QLineF {
	return ptr
}

//QLineF::IntersectType
type QLineF__IntersectType int

var (
	QLineF__NoIntersection        = QLineF__IntersectType(0)
	QLineF__BoundedIntersection   = QLineF__IntersectType(1)
	QLineF__UnboundedIntersection = QLineF__IntersectType(2)
)

func (ptr *QLineF) Intersect(line QLineFITF, intersectionPoint QPointFITF) QLineF__IntersectType {
	if ptr.Pointer() != nil {
		return QLineF__IntersectType(C.QLineF_Intersect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLineF(line)), C.QtObjectPtr(PointerFromQPointF(intersectionPoint))))
	}
	return 0
}

func NewQLineF() *QLineF {
	return QLineFFromPointer(unsafe.Pointer(C.QLineF_NewQLineF()))
}

func NewQLineF4(line QLineITF) *QLineF {
	return QLineFFromPointer(unsafe.Pointer(C.QLineF_NewQLineF4(C.QtObjectPtr(PointerFromQLine(line)))))
}

func NewQLineF2(p1 QPointFITF, p2 QPointFITF) *QLineF {
	return QLineFFromPointer(unsafe.Pointer(C.QLineF_NewQLineF2(C.QtObjectPtr(PointerFromQPointF(p1)), C.QtObjectPtr(PointerFromQPointF(p2)))))
}

func (ptr *QLineF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QLineF_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLineF) SetP1(p1 QPointFITF) {
	if ptr.Pointer() != nil {
		C.QLineF_SetP1(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(p1)))
	}
}

func (ptr *QLineF) SetP2(p2 QPointFITF) {
	if ptr.Pointer() != nil {
		C.QLineF_SetP2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(p2)))
	}
}

func (ptr *QLineF) SetPoints(p1 QPointFITF, p2 QPointFITF) {
	if ptr.Pointer() != nil {
		C.QLineF_SetPoints(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(p1)), C.QtObjectPtr(PointerFromQPointF(p2)))
	}
}

func (ptr *QLineF) Translate(offset QPointFITF) {
	if ptr.Pointer() != nil {
		C.QLineF_Translate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPointF(offset)))
	}
}
