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

type QGraphicsPixmapItemITF interface {
	QGraphicsItemITF
	QGraphicsPixmapItemPTR() *QGraphicsPixmapItem
}

func PointerFromQGraphicsPixmapItem(ptr QGraphicsPixmapItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsPixmapItemPTR().Pointer()
	}
	return nil
}

func QGraphicsPixmapItemFromPointer(ptr unsafe.Pointer) *QGraphicsPixmapItem {
	var n = new(QGraphicsPixmapItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsPixmapItem) QGraphicsPixmapItemPTR() *QGraphicsPixmapItem {
	return ptr
}

//QGraphicsPixmapItem::ShapeMode
type QGraphicsPixmapItem__ShapeMode int

var (
	QGraphicsPixmapItem__MaskShape          = QGraphicsPixmapItem__ShapeMode(0)
	QGraphicsPixmapItem__BoundingRectShape  = QGraphicsPixmapItem__ShapeMode(1)
	QGraphicsPixmapItem__HeuristicMaskShape = QGraphicsPixmapItem__ShapeMode(2)
)

func NewQGraphicsPixmapItem(parent QGraphicsItemITF) *QGraphicsPixmapItem {
	return QGraphicsPixmapItemFromPointer(unsafe.Pointer(C.QGraphicsPixmapItem_NewQGraphicsPixmapItem(C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func NewQGraphicsPixmapItem2(pixmap gui.QPixmapITF, parent QGraphicsItemITF) *QGraphicsPixmapItem {
	return QGraphicsPixmapItemFromPointer(unsafe.Pointer(C.QGraphicsPixmapItem_NewQGraphicsPixmapItem2(C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)), C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func (ptr *QGraphicsPixmapItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPixmapItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsPixmapItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsPixmapItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsPixmapItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsPixmapItem) SetOffset(offset core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(offset)))
	}
}

func (ptr *QGraphicsPixmapItem) SetPixmap(pixmap gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QGraphicsPixmapItem) SetShapeMode(mode QGraphicsPixmapItem__ShapeMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetShapeMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsPixmapItem) SetTransformationMode(mode core.Qt__TransformationMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetTransformationMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsPixmapItem) ShapeMode() QGraphicsPixmapItem__ShapeMode {
	if ptr.Pointer() != nil {
		return QGraphicsPixmapItem__ShapeMode(C.QGraphicsPixmapItem_ShapeMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) TransformationMode() core.Qt__TransformationMode {
	if ptr.Pointer() != nil {
		return core.Qt__TransformationMode(C.QGraphicsPixmapItem_TransformationMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsPixmapItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) DestroyQGraphicsPixmapItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
