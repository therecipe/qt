package gui

//#include "qpolygon.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPolygon struct {
	core.QVector
}

type QPolygon_ITF interface {
	core.QVector_ITF
	QPolygon_PTR() *QPolygon
}

func PointerFromQPolygon(ptr QPolygon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygon_PTR().Pointer()
	}
	return nil
}

func NewQPolygonFromPointer(ptr unsafe.Pointer) *QPolygon {
	var n = new(QPolygon)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPolygon) QPolygon_PTR() *QPolygon {
	return ptr
}

func NewQPolygon5(rectangle core.QRect_ITF, closed bool) *QPolygon {
	return NewQPolygonFromPointer(C.QPolygon_NewQPolygon5(core.PointerFromQRect(rectangle), C.int(qt.GoBoolToInt(closed))))
}

func (ptr *QPolygon) ContainsPoint(point core.QPoint_ITF, fillRule core.Qt__FillRule) bool {
	if ptr.Pointer() != nil {
		return C.QPolygon_ContainsPoint(ptr.Pointer(), core.PointerFromQPoint(point), C.int(fillRule)) != 0
	}
	return false
}

func (ptr *QPolygon) PutPoints3(index int, nPoints int, fromPolygon QPolygon_ITF, fromIndex int) {
	if ptr.Pointer() != nil {
		C.QPolygon_PutPoints3(ptr.Pointer(), C.int(index), C.int(nPoints), PointerFromQPolygon(fromPolygon), C.int(fromIndex))
	}
}

func NewQPolygon() *QPolygon {
	return NewQPolygonFromPointer(C.QPolygon_NewQPolygon())
}

func NewQPolygon3(polygon QPolygon_ITF) *QPolygon {
	return NewQPolygonFromPointer(C.QPolygon_NewQPolygon3(PointerFromQPolygon(polygon)))
}

func NewQPolygon2(size int) *QPolygon {
	return NewQPolygonFromPointer(C.QPolygon_NewQPolygon2(C.int(size)))
}

func (ptr *QPolygon) Point(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Point(ptr.Pointer(), C.int(index), C.int(x), C.int(y))
	}
}

func (ptr *QPolygon) SetPoint2(index int, point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint2(ptr.Pointer(), C.int(index), core.PointerFromQPoint(point))
	}
}

func (ptr *QPolygon) SetPoint(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint(ptr.Pointer(), C.int(index), C.int(x), C.int(y))
	}
}

func (ptr *QPolygon) SetPoints(nPoints int, points int) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoints(ptr.Pointer(), C.int(nPoints), C.int(points))
	}
}

func (ptr *QPolygon) Swap(other QPolygon_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_Swap(ptr.Pointer(), PointerFromQPolygon(other))
	}
}

func (ptr *QPolygon) Translate2(offset core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_Translate2(ptr.Pointer(), core.PointerFromQPoint(offset))
	}
}

func (ptr *QPolygon) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Translate(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QPolygon) DestroyQPolygon() {
	if ptr.Pointer() != nil {
		C.QPolygon_DestroyQPolygon(ptr.Pointer())
	}
}
