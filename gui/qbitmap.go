package gui

//#include "gui.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBitmap::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBitmap_Clear(ptr.Pointer())
	}
}

func (ptr *QBitmap) Swap(other QBitmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBitmap::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBitmap_Swap(ptr.Pointer(), PointerFromQBitmap(other))
	}
}

func (ptr *QBitmap) DestroyQBitmap() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBitmap::~QBitmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBitmap_DestroyQBitmap(ptr.Pointer())
	}
}
