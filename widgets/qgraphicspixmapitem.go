package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QGraphicsPixmapItem_") {
		n.SetObjectNameAbs("QGraphicsPixmapItem_" + qt.Identifier())
	}
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
	defer qt.Recovering("QGraphicsPixmapItem::QGraphicsPixmapItem")

	return NewQGraphicsPixmapItemFromPointer(C.QGraphicsPixmapItem_NewQGraphicsPixmapItem(PointerFromQGraphicsItem(parent)))
}

func NewQGraphicsPixmapItem2(pixmap gui.QPixmap_ITF, parent QGraphicsItem_ITF) *QGraphicsPixmapItem {
	defer qt.Recovering("QGraphicsPixmapItem::QGraphicsPixmapItem")

	return NewQGraphicsPixmapItemFromPointer(C.QGraphicsPixmapItem_NewQGraphicsPixmapItem2(gui.PointerFromQPixmap(pixmap), PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsPixmapItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsPixmapItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsPixmapItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsPixmapItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsPixmapItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsPixmapItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsPixmapItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsPixmapItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "paint", f)
	}
}

func (ptr *QGraphicsPixmapItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsPixmapItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "paint")
	}
}

//export callbackQGraphicsPixmapItemPaint
func callbackQGraphicsPixmapItemPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsPixmapItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsPixmapItem) SetOffset(offset core.QPointF_ITF) {
	defer qt.Recovering("QGraphicsPixmapItem::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetOffset(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QGraphicsPixmapItem) SetOffset2(x float64, y float64) {
	defer qt.Recovering("QGraphicsPixmapItem::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetOffset2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QGraphicsPixmapItem) SetPixmap(pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QGraphicsPixmapItem::setPixmap")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QGraphicsPixmapItem) SetShapeMode(mode QGraphicsPixmapItem__ShapeMode) {
	defer qt.Recovering("QGraphicsPixmapItem::setShapeMode")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetShapeMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsPixmapItem) SetTransformationMode(mode core.Qt__TransformationMode) {
	defer qt.Recovering("QGraphicsPixmapItem::setTransformationMode")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetTransformationMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsPixmapItem) ShapeMode() QGraphicsPixmapItem__ShapeMode {
	defer qt.Recovering("QGraphicsPixmapItem::shapeMode")

	if ptr.Pointer() != nil {
		return QGraphicsPixmapItem__ShapeMode(C.QGraphicsPixmapItem_ShapeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) TransformationMode() core.Qt__TransformationMode {
	defer qt.Recovering("QGraphicsPixmapItem::transformationMode")

	if ptr.Pointer() != nil {
		return core.Qt__TransformationMode(C.QGraphicsPixmapItem_TransformationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) Type() int {
	defer qt.Recovering("QGraphicsPixmapItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsPixmapItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsPixmapItem) DestroyQGraphicsPixmapItem() {
	defer qt.Recovering("QGraphicsPixmapItem::~QGraphicsPixmapItem")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(ptr.Pointer())
	}
}

func (ptr *QGraphicsPixmapItem) ObjectNameAbs() string {
	defer qt.Recovering("QGraphicsPixmapItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsPixmapItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsPixmapItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGraphicsPixmapItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGraphicsPixmapItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
