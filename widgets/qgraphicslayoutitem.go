package widgets

//#include "qgraphicslayoutitem.h"
import "C"
import (
	"unsafe"
)

type QGraphicsLayoutItem struct {
	ptr unsafe.Pointer
}

type QGraphicsLayoutItem_ITF interface {
	QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem
}

func (p *QGraphicsLayoutItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGraphicsLayoutItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGraphicsLayoutItem(ptr QGraphicsLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLayoutItemFromPointer(ptr unsafe.Pointer) *QGraphicsLayoutItem {
	var n = new(QGraphicsLayoutItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLayoutItem) QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem {
	return ptr
}
