package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGraphicsAnchorLayout struct {
	QGraphicsLayout
}

type QGraphicsAnchorLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsAnchorLayout_PTR() *QGraphicsAnchorLayout
}

func PointerFromQGraphicsAnchorLayout(ptr QGraphicsAnchorLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsAnchorLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsAnchorLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsAnchorLayout {
	var n = new(QGraphicsAnchorLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsAnchorLayout) QGraphicsAnchorLayout_PTR() *QGraphicsAnchorLayout {
	return ptr
}

func NewQGraphicsAnchorLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsAnchorLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::QGraphicsAnchorLayout")
		}
	}()

	return NewQGraphicsAnchorLayoutFromPointer(C.QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsAnchorLayout) AddAnchor(firstItem QGraphicsLayoutItem_ITF, firstEdge core.Qt__AnchorPoint, secondItem QGraphicsLayoutItem_ITF, secondEdge core.Qt__AnchorPoint) *QGraphicsAnchor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::addAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsAnchorFromPointer(C.QGraphicsAnchorLayout_AddAnchor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), C.int(firstEdge), PointerFromQGraphicsLayoutItem(secondItem), C.int(secondEdge)))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) AddAnchors(firstItem QGraphicsLayoutItem_ITF, secondItem QGraphicsLayoutItem_ITF, orientations core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::addAnchors")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_AddAnchors(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), PointerFromQGraphicsLayoutItem(secondItem), C.int(orientations))
	}
}

func (ptr *QGraphicsAnchorLayout) AddCornerAnchors(firstItem QGraphicsLayoutItem_ITF, firstCorner core.Qt__Corner, secondItem QGraphicsLayoutItem_ITF, secondCorner core.Qt__Corner) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::addCornerAnchors")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_AddCornerAnchors(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), C.int(firstCorner), PointerFromQGraphicsLayoutItem(secondItem), C.int(secondCorner))
	}
}

func (ptr *QGraphicsAnchorLayout) Anchor(firstItem QGraphicsLayoutItem_ITF, firstEdge core.Qt__AnchorPoint, secondItem QGraphicsLayoutItem_ITF, secondEdge core.Qt__AnchorPoint) *QGraphicsAnchor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::anchor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsAnchorFromPointer(C.QGraphicsAnchorLayout_Anchor(ptr.Pointer(), PointerFromQGraphicsLayoutItem(firstItem), C.int(firstEdge), PointerFromQGraphicsLayoutItem(secondItem), C.int(secondEdge)))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsAnchorLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::horizontalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsAnchorLayout_HorizontalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsAnchorLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGraphicsAnchorLayout) RemoveAt(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::removeAt")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_RemoveAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsAnchorLayout) SetGeometry(geom core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(geom))
	}
}

func (ptr *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::setHorizontalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetHorizontalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::setVerticalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_SetVerticalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::verticalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsAnchorLayout_VerticalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchorLayout) DestroyQGraphicsAnchorLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchorLayout::~QGraphicsAnchorLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(ptr.Pointer())
	}
}
