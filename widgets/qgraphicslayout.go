package widgets

//#include "qgraphicslayout.h"
import "C"
import (
	"unsafe"
)

type QGraphicsLayout struct {
	QGraphicsLayoutItem
}

type QGraphicsLayout_ITF interface {
	QGraphicsLayoutItem_ITF
	QGraphicsLayout_PTR() *QGraphicsLayout
}

func PointerFromQGraphicsLayout(ptr QGraphicsLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsLayout {
	var n = new(QGraphicsLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLayout) QGraphicsLayout_PTR() *QGraphicsLayout {
	return ptr
}
