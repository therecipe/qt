package widgets

//#include "qgraphicsanchorlayout.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsAnchorLayout struct {
	QGraphicsLayout
}

type QGraphicsAnchorLayoutITF interface {
	QGraphicsLayoutITF
	QGraphicsAnchorLayoutPTR() *QGraphicsAnchorLayout
}

func PointerFromQGraphicsAnchorLayout(ptr QGraphicsAnchorLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsAnchorLayoutPTR().Pointer()
	}
	return nil
}

func QGraphicsAnchorLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsAnchorLayout {
	var n = new(QGraphicsAnchorLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsAnchorLayout) QGraphicsAnchorLayoutPTR() *QGraphicsAnchorLayout {
	return ptr
}

func NewQGraphicsAnchorLayout(parent QGraphicsLayoutItemITF) *QGraphicsAnchorLayout {
	return QGraphicsAnchorLayoutFromPointer(unsafe.Pointer(C.QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(C.QtObjectPtr(PointerFromQGraphicsLayoutItem(parent)))))
}

func (ptr *QGraphicsAnchorLayout) AddAnchor(firstItem QGraphicsLayoutItemITF, firstEdge core.Qt__AnchorPoint, secondItem QGraphicsLayoutItemITF, secondEdge core.Qt__AnchorPoint) *QGraphicsAnchor {
	if ptr.Pointer() != nil {
		return QGraphicsAnchorFromPointer(unsafe.Pointer(C.QGraphicsAnchorLayout_AddAnchor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(firstItem)), C.int(firstEdge), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(secondItem)), C.int(secondEdge))))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) AddAnchors(firstItem QGraphicsLayoutItemITF, secondItem QGraphicsLayoutItemITF, orientations core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_AddAnchors(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(firstItem)), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(secondItem)), C.int(orientations))
	}
}

func (ptr *QGraphicsAnchorLayout) AddCornerAnchors(firstItem QGraphicsLayoutItemITF, firstCorner core.Qt__Corner, secondItem QGraphicsLayoutItemITF, secondCorner core.Qt__Corner) {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_AddCornerAnchors(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(firstItem)), C.int(firstCorner), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(secondItem)), C.int(secondCorner))
	}
}

func (ptr *QGraphicsAnchorLayout) Anchor(firstItem QGraphicsLayoutItemITF, firstEdge core.Qt__AnchorPoint, secondItem QGraphicsLayoutItemITF, secondEdge core.Qt__AnchorPoint) *QGraphicsAnchor {
	if ptr.Pointer() != nil {
		return QGraphicsAnchorFromPointer(unsafe.Pointer(C.QGraphicsAnchorLayout_Anchor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(firstItem)), C.int(firstEdge), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(secondItem)), C.int(secondEdge))))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsAnchorLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return QGraphicsLayoutItemFromPointer(unsafe.Pointer(C.QGraphicsAnchorLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) RemoveAt(index int) {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_RemoveAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QGraphicsAnchorLayout) SetGeometry(geom core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(geom)))
	}
}

func (ptr *QGraphicsAnchorLayout) DestroyQGraphicsAnchorLayout() {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}
