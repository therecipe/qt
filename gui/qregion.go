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

type QRegionITF interface {
	QRegionPTR() *QRegion
}

func (p *QRegion) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegion) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegion(ptr QRegionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegionPTR().Pointer()
	}
	return nil
}

func QRegionFromPointer(ptr unsafe.Pointer) *QRegion {
	var n = new(QRegion)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegion) QRegionPTR() *QRegion {
	return ptr
}

//QRegion::RegionType
type QRegion__RegionType int

var (
	QRegion__Rectangle = QRegion__RegionType(0)
	QRegion__Ellipse   = QRegion__RegionType(1)
)

func NewQRegion() *QRegion {
	return QRegionFromPointer(unsafe.Pointer(C.QRegion_NewQRegion()))
}

func NewQRegion5(bm QBitmapITF) *QRegion {
	return QRegionFromPointer(unsafe.Pointer(C.QRegion_NewQRegion5(C.QtObjectPtr(PointerFromQBitmap(bm)))))
}

func NewQRegion3(a QPolygonITF, fillRule core.Qt__FillRule) *QRegion {
	return QRegionFromPointer(unsafe.Pointer(C.QRegion_NewQRegion3(C.QtObjectPtr(PointerFromQPolygon(a)), C.int(fillRule))))
}

func NewQRegion6(r core.QRectITF, t QRegion__RegionType) *QRegion {
	return QRegionFromPointer(unsafe.Pointer(C.QRegion_NewQRegion6(C.QtObjectPtr(core.PointerFromQRect(r)), C.int(t))))
}

func NewQRegion4(r QRegionITF) *QRegion {
	return QRegionFromPointer(unsafe.Pointer(C.QRegion_NewQRegion4(C.QtObjectPtr(PointerFromQRegion(r)))))
}

func (ptr *QRegion) Contains(p core.QPointITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p))) != 0
	}
	return false
}

func (ptr *QRegion) Contains2(r core.QRectITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Contains2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(r))) != 0
	}
	return false
}

func (ptr *QRegion) Intersects2(rect core.QRectITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Intersects2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect))) != 0
	}
	return false
}

func (ptr *QRegion) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QRegion_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegion) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QRegion_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegion) RectCount() int {
	if ptr.Pointer() != nil {
		return int(C.QRegion_RectCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRegion) SetRects(rects core.QRectITF, number int) {
	if ptr.Pointer() != nil {
		C.QRegion_SetRects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rects)), C.int(number))
	}
}

func (ptr *QRegion) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QRegion_Translate(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy))
	}
}

func NewQRegion2(x int, y int, w int, h int, t QRegion__RegionType) *QRegion {
	return QRegionFromPointer(unsafe.Pointer(C.QRegion_NewQRegion2(C.int(x), C.int(y), C.int(w), C.int(h), C.int(t))))
}

func (ptr *QRegion) Intersects(region QRegionITF) bool {
	if ptr.Pointer() != nil {
		return C.QRegion_Intersects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region))) != 0
	}
	return false
}

func (ptr *QRegion) Swap(other QRegionITF) {
	if ptr.Pointer() != nil {
		C.QRegion_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(other)))
	}
}

func (ptr *QRegion) Translate2(point core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QRegion_Translate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)))
	}
}
