package widgets

//#include "qgraphicslayout.h"
import "C"
import (
	"unsafe"
)

type QGraphicsLayout struct {
	QGraphicsLayoutItem
}

type QGraphicsLayoutITF interface {
	QGraphicsLayoutItemITF
	QGraphicsLayoutPTR() *QGraphicsLayout
}

func PointerFromQGraphicsLayout(ptr QGraphicsLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutPTR().Pointer()
	}
	return nil
}

func QGraphicsLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsLayout {
	var n = new(QGraphicsLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsLayout) QGraphicsLayoutPTR() *QGraphicsLayout {
	return ptr
}
