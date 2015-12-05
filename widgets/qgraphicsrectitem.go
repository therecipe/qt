package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QGraphicsRectItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsRectItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsRectItem_PTR() *QGraphicsRectItem
}

func PointerFromQGraphicsRectItem(ptr QGraphicsRectItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsRectItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsRectItemFromPointer(ptr unsafe.Pointer) *QGraphicsRectItem {
	var n = new(QGraphicsRectItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsRectItem) QGraphicsRectItem_PTR() *QGraphicsRectItem {
	return ptr
}

func (ptr *QGraphicsRectItem) SetRect(rectangle core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::setRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_SetRect(ptr.Pointer(), core.PointerFromQRectF(rectangle))
	}
}

func (ptr *QGraphicsRectItem) Contains(point core.QPointF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsRectItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsRectItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::isObscuredBy")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsRectItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsRectItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsRectItem) SetRect2(x float64, y float64, width float64, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::setRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QGraphicsRectItem) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsRectItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsRectItem) DestroyQGraphicsRectItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsRectItem::~QGraphicsRectItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsRectItem_DestroyQGraphicsRectItem(ptr.Pointer())
	}
}
