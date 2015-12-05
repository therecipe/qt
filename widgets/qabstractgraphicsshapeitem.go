package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QAbstractGraphicsShapeItem struct {
	QGraphicsItem
}

type QAbstractGraphicsShapeItem_ITF interface {
	QGraphicsItem_ITF
	QAbstractGraphicsShapeItem_PTR() *QAbstractGraphicsShapeItem
}

func PointerFromQAbstractGraphicsShapeItem(ptr QAbstractGraphicsShapeItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractGraphicsShapeItem_PTR().Pointer()
	}
	return nil
}

func NewQAbstractGraphicsShapeItemFromPointer(ptr unsafe.Pointer) *QAbstractGraphicsShapeItem {
	var n = new(QAbstractGraphicsShapeItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractGraphicsShapeItem) QAbstractGraphicsShapeItem_PTR() *QAbstractGraphicsShapeItem {
	return ptr
}

func (ptr *QAbstractGraphicsShapeItem) Brush() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractGraphicsShapeItem::brush")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QAbstractGraphicsShapeItem_Brush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractGraphicsShapeItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractGraphicsShapeItem::isObscuredBy")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractGraphicsShapeItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QAbstractGraphicsShapeItem) SetBrush(brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractGraphicsShapeItem::setBrush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractGraphicsShapeItem) SetPen(pen gui.QPen_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractGraphicsShapeItem::setPen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractGraphicsShapeItem) DestroyQAbstractGraphicsShapeItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractGraphicsShapeItem::~QAbstractGraphicsShapeItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_DestroyQAbstractGraphicsShapeItem(ptr.Pointer())
	}
}
