package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QBitmap struct {
	QPixmap
}

type QBitmap_ITF interface {
	QPixmap_ITF
	QBitmap_PTR() *QBitmap
}

func PointerFromQBitmap(ptr QBitmap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitmap_PTR().Pointer()
	}
	return nil
}

func NewQBitmapFromPointer(ptr unsafe.Pointer) *QBitmap {
	var n = new(QBitmap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBitmap) QBitmap_PTR() *QBitmap {
	return ptr
}

func (ptr *QBitmap) Clear() {
	defer qt.Recovering("QBitmap::clear")

	if ptr.Pointer() != nil {
		C.QBitmap_Clear(ptr.Pointer())
	}
}

func (ptr *QBitmap) Swap(other QBitmap_ITF) {
	defer qt.Recovering("QBitmap::swap")

	if ptr.Pointer() != nil {
		C.QBitmap_Swap(ptr.Pointer(), PointerFromQBitmap(other))
	}
}

func (ptr *QBitmap) DestroyQBitmap() {
	defer qt.Recovering("QBitmap::~QBitmap")

	if ptr.Pointer() != nil {
		C.QBitmap_DestroyQBitmap(ptr.Pointer())
	}
}
