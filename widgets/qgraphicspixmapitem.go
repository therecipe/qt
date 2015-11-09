package widgets

//#include "qgraphicspixmapitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsPixmapItem struct {
	QGraphicsItem
}

type QGraphicsPixmapItem_ITF interface {
	QGraphicsItem_ITF
	QGraphicsPixmapItem_PTR() *QGraphicsPixmapItem
}

func PointerFromQGraphicsPixmapItem(ptr QGraphicsPixmapItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsPixmapItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsPixmapItemFromPointer(ptr unsafe.Pointer) *QGraphicsPixmapItem {
	var n = new(QGraphicsPixmapItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsPixmapItem) QGraphicsPixmapItem_PTR() *QGraphicsPixmapItem {
	return ptr
}

//QGraphicsPixmapItem::ShapeMode
type QGraphicsPixmapItem__ShapeMode int64

const (
	QGraphicsPixmapItem__MaskShape          = QGraphicsPixmapItem__ShapeMode(0)
	QGraphicsPixmapItem__BoundingRectShape  = QGraphicsPixmapItem__ShapeMode(1)
	QGraphicsPixmapItem__HeuristicMaskShape = QGraphicsPixmapItem__ShapeMode(2)
)

func NewQGraphicsPixmapItem(parent QGraphicsItem_ITF) *QGraphicsPixmapItem {
	return NewQGraphicsPixmapItemFromPointer(C.QGraphicsPixmapItem_NewQGraphicsPixmapItem(PointerFromQGraphicsItem(parent)))
}

func NewQGraphicsPixmapItem2(pixmap gui.QPixmap_ITF, parent QGraphicsItem_ITF) *QGraphicsPixmapItem {
	return NewQGraphicsPixmapItemFromPointer(C.QGraphicsPixmapItem_NewQGraphicsPixmapItem2(gui.PointerFromQPixmap(pixmap), PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsPixmapItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPixmapItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsPixmapItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPixmapItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsPixmapItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsPixmapItem) SetOffset(offset core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetOffset(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QGraphicsPixmapItem) SetOffset2(x float64, y float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetOffset2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QGraphicsPixmapItem) SetPixmap(pixmap gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QGraphicsPixmapItem) SetShapeMode(mode QGraphicsPixmapItem__ShapeMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetShapeMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsPixmapItem) SetTransformationMode(mode core.Qt__TransformationMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetTransformationMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsPixmapItem) ShapeMode() QGraphicsPixmapItem__ShapeMode {
	if ptr.Pointer() != nil {
		return QGraphicsPixmapItem__ShapeMode(C.QGraphicsPixmapItem_ShapeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) TransformationMode() core.Qt__TransformationMode {
	if ptr.Pointer() != nil {
		return core.Qt__TransformationMode(C.QGraphicsPixmapItem_TransformationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsPixmapItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) DestroyQGraphicsPixmapItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(ptr.Pointer())
	}
}
