package gui

//#include "qregion.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRegion struct {
	ptr unsafe.Pointer
}

type QRegion_ITF interface {
	QRegion_PTR() *QRegion
}

func (p *QRegion) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegion) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegion(ptr QRegion_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegion_PTR().Pointer()
	}
	return nil
}

func NewQRegionFromPointer(ptr unsafe.Pointer) *QRegion {
	var n = new(QRegion)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegion) QRegion_PTR() *QRegion {
	return ptr
}

//QRegion::RegionType
type QRegion__RegionType int64

const (
	QRegion__Rectangle = QRegion__RegionType(0)
	QRegion__Ellipse   = QRegion__RegionType(1)
)

func NewQRegion() *QRegion {
	return NewQRegionFromPointer(C.QRegion_NewQRegion())
}

func NewQRegion5(bm QBitmap_ITF) *QRegion {
	return NewQRegionFromPointer(C.QRegion_NewQRegion5(PointerFromQBitmap(bm)))
}

func NewQRegion3(a QPolygon_ITF, fillRule core.Qt__FillRule) *QRegion {
	return NewQRegionFromPointer(C.QRegion_NewQRegion3(PointerFromQPolygon(a), C.int(fillRule)))
}

func NewQRegion6(r core.QRect_ITF, t QRegion__RegionType) *QRegion {
	return NewQRegionFromPointer(C.QRegion_NewQRegion6(core.PointerFromQRect(r), C.int(t)))
}

func NewQRegion4(r QRegion_ITF) *QRegion {
	return NewQRegionFromPointer(C.QRegion_NewQRegion4(PointerFromQRegion(r)))
}

func (ptr *QRegion) Contains(p core.QPoint_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Contains(ptr.Pointer(), core.PointerFromQPoint(p)) != 0
	}
	return false
}

func (ptr *QRegion) Contains2(r core.QRect_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Contains2(ptr.Pointer(), core.PointerFromQRect(r)) != 0
	}
	return false
}

func (ptr *QRegion) Intersected2(rect core.QRect_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_Intersected2(ptr.Pointer(), core.PointerFromQRect(rect)))
	}
	return nil
}

func (ptr *QRegion) Intersected(r QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_Intersected(ptr.Pointer(), PointerFromQRegion(r)))
	}
	return nil
}

func (ptr *QRegion) Intersects2(rect core.QRect_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Intersects2(ptr.Pointer(), core.PointerFromQRect(rect)) != 0
	}
	return false
}

func (ptr *QRegion) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QRegion_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegion) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QRegion_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegion) RectCount() int {
	if ptr.Pointer() != nil {
		return int(C.QRegion_RectCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegion) SetRects(rects core.QRect_ITF, number int) {
	if ptr.Pointer() != nil {
		C.QRegion_SetRects(ptr.Pointer(), core.PointerFromQRect(rects), C.int(number))
	}
}

func (ptr *QRegion) Subtracted(r QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_Subtracted(ptr.Pointer(), PointerFromQRegion(r)))
	}
	return nil
}

func (ptr *QRegion) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QRegion_Translate(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QRegion) United2(rect core.QRect_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_United2(ptr.Pointer(), core.PointerFromQRect(rect)))
	}
	return nil
}

func (ptr *QRegion) United(r QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_United(ptr.Pointer(), PointerFromQRegion(r)))
	}
	return nil
}

func (ptr *QRegion) Xored(r QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_Xored(ptr.Pointer(), PointerFromQRegion(r)))
	}
	return nil
}

func NewQRegion2(x int, y int, w int, h int, t QRegion__RegionType) *QRegion {
	return NewQRegionFromPointer(C.QRegion_NewQRegion2(C.int(x), C.int(y), C.int(w), C.int(h), C.int(t)))
}

func (ptr *QRegion) Intersects(region QRegion_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Intersects(ptr.Pointer(), PointerFromQRegion(region)) != 0
	}
	return false
}

func (ptr *QRegion) Swap(other QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QRegion_Swap(ptr.Pointer(), PointerFromQRegion(other))
	}
}

func (ptr *QRegion) Translate2(point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QRegion_Translate2(ptr.Pointer(), core.PointerFromQPoint(point))
	}
}

func (ptr *QRegion) Translated2(p core.QPoint_ITF) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_Translated2(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QRegion) Translated(dx int, dy int) *QRegion {
	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QRegion_Translated(ptr.Pointer(), C.int(dx), C.int(dy)))
	}
	return nil
}
