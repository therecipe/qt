package widgets

//#include "qgraphicslayoutitem.h"
import "C"
import (
	"unsafe"
)

type QGraphicsLayoutItem struct {
	ptr unsafe.Pointer
}

type QGraphicsLayoutItemITF interface {
	QGraphicsLayoutItemPTR() *QGraphicsLayoutItem
}

func (p *QGraphicsLayoutItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGraphicsLayoutItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGraphicsLayoutItem(ptr QGraphicsLayoutItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutItemPTR().Pointer()
	}
	return nil
}

func QGraphicsLayoutItemFromPointer(ptr unsafe.Pointer) *QGraphicsLayoutItem {
	var n = new(QGraphicsLayoutItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLayoutItem) QGraphicsLayoutItemPTR() *QGraphicsLayoutItem {
	return ptr
}
