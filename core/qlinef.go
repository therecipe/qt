package core

//#include "qlinef.h"
import "C"
import (
	"unsafe"
)

type QLineF struct {
	ptr unsafe.Pointer
}

type QLineF_ITF interface {
	QLineF_PTR() *QLineF
}

func (p *QLineF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLineF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLineF(ptr QLineF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineF_PTR().Pointer()
	}
	return nil
}

func NewQLineFFromPointer(ptr unsafe.Pointer) *QLineF {
	var n = new(QLineF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLineF) QLineF_PTR() *QLineF {
	return ptr
}

//QLineF::IntersectType
type QLineF__IntersectType int64

const (
	QLineF__NoIntersection        = QLineF__IntersectType(0)
	QLineF__BoundedIntersection   = QLineF__IntersectType(1)
	QLineF__UnboundedIntersection = QLineF__IntersectType(2)
)

func (ptr *QLineF) AngleTo(line QLineF_ITF) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_AngleTo(ptr.Pointer(), PointerFromQLineF(line)))
	}
	return 0
}

func (ptr *QLineF) Intersect(line QLineF_ITF, intersectionPoint QPointF_ITF) QLineF__IntersectType {
	if ptr.Pointer() != nil {
		return QLineF__IntersectType(C.QLineF_Intersect(ptr.Pointer(), PointerFromQLineF(line), PointerFromQPointF(intersectionPoint)))
	}
	return 0
}

func NewQLineF() *QLineF {
	return NewQLineFFromPointer(C.QLineF_NewQLineF())
}

func NewQLineF4(line QLine_ITF) *QLineF {
	return NewQLineFFromPointer(C.QLineF_NewQLineF4(PointerFromQLine(line)))
}

func NewQLineF2(p1 QPointF_ITF, p2 QPointF_ITF) *QLineF {
	return NewQLineFFromPointer(C.QLineF_NewQLineF2(PointerFromQPointF(p1), PointerFromQPointF(p2)))
}

func NewQLineF3(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	return NewQLineFFromPointer(C.QLineF_NewQLineF3(C.double(x1), C.double(y1), C.double(x2), C.double(y2)))
}

func (ptr *QLineF) Angle() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Angle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) Dx() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Dx(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) Dy() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Dy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QLineF_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLineF) Length() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) SetAngle(angle float64) {
	if ptr.Pointer() != nil {
		C.QLineF_SetAngle(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QLineF) SetLength(length float64) {
	if ptr.Pointer() != nil {
		C.QLineF_SetLength(ptr.Pointer(), C.double(length))
	}
}

func (ptr *QLineF) SetLine(x1 float64, y1 float64, x2 float64, y2 float64) {
	if ptr.Pointer() != nil {
		C.QLineF_SetLine(ptr.Pointer(), C.double(x1), C.double(y1), C.double(x2), C.double(y2))
	}
}

func (ptr *QLineF) SetP1(p1 QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QLineF_SetP1(ptr.Pointer(), PointerFromQPointF(p1))
	}
}

func (ptr *QLineF) SetP2(p2 QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QLineF_SetP2(ptr.Pointer(), PointerFromQPointF(p2))
	}
}

func (ptr *QLineF) SetPoints(p1 QPointF_ITF, p2 QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QLineF_SetPoints(ptr.Pointer(), PointerFromQPointF(p1), PointerFromQPointF(p2))
	}
}

func (ptr *QLineF) Translate(offset QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QLineF_Translate(ptr.Pointer(), PointerFromQPointF(offset))
	}
}

func (ptr *QLineF) Translate2(dx float64, dy float64) {
	if ptr.Pointer() != nil {
		C.QLineF_Translate2(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QLineF) X1() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_X1(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) X2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_X2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) Y1() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Y1(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLineF) Y2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Y2(ptr.Pointer()))
	}
	return 0
}
