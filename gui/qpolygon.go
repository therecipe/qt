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

type QPolygonITF interface {
	core.QVectorITF
	QPolygonPTR() *QPolygon
}

func PointerFromQPolygon(ptr QPolygonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygonPTR().Pointer()
	}
	return nil
}

func QPolygonFromPointer(ptr unsafe.Pointer) *QPolygon {
	var n = new(QPolygon)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPolygon) QPolygonPTR() *QPolygon {
	return ptr
}

func NewQPolygon5(rectangle core.QRectITF, closed bool) *QPolygon {
	return QPolygonFromPointer(unsafe.Pointer(C.QPolygon_NewQPolygon5(C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(qt.GoBoolToInt(closed)))))
}

func (ptr *QPolygon) ContainsPoint(point core.QPointITF, fillRule core.Qt__FillRule) bool {
	if ptr.Pointer() != nil {
		return C.QPolygon_ContainsPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)), C.int(fillRule)) != 0
	}
	return false
}

func (ptr *QPolygon) PutPoints3(index int, nPoints int, fromPolygon QPolygonITF, fromIndex int) {
	if ptr.Pointer() != nil {
		C.QPolygon_PutPoints3(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(nPoints), C.QtObjectPtr(PointerFromQPolygon(fromPolygon)), C.int(fromIndex))
	}
}

func NewQPolygon() *QPolygon {
	return QPolygonFromPointer(unsafe.Pointer(C.QPolygon_NewQPolygon()))
}

func NewQPolygon3(polygon QPolygonITF) *QPolygon {
	return QPolygonFromPointer(unsafe.Pointer(C.QPolygon_NewQPolygon3(C.QtObjectPtr(PointerFromQPolygon(polygon)))))
}

func NewQPolygon2(size int) *QPolygon {
	return QPolygonFromPointer(unsafe.Pointer(C.QPolygon_NewQPolygon2(C.int(size))))
}

func (ptr *QPolygon) Point(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Point(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(x), C.int(y))
	}
}

func (ptr *QPolygon) SetPoint2(index int, point core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint2(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(core.PointerFromQPoint(point)))
	}
}

func (ptr *QPolygon) SetPoint(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(x), C.int(y))
	}
}

func (ptr *QPolygon) SetPoints(nPoints int, points int) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoints(C.QtObjectPtr(ptr.Pointer()), C.int(nPoints), C.int(points))
	}
}

func (ptr *QPolygon) Swap(other QPolygonITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygon(other)))
	}
}

func (ptr *QPolygon) Translate2(offset core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_Translate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(offset)))
	}
}

func (ptr *QPolygon) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Translate(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy))
	}
}

func (ptr *QPolygon) DestroyQPolygon() {
	if ptr.Pointer() != nil {
		C.QPolygon_DestroyQPolygon(C.QtObjectPtr(ptr.Pointer()))
	}
}
