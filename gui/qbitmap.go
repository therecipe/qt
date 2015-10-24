package gui

//#include "qbitmap.h"
import "C"
import (
	"unsafe"
)

type QBitmap struct {
	QPixmap
}

type QBitmapITF interface {
	QPixmapITF
	QBitmapPTR() *QBitmap
}

func PointerFromQBitmap(ptr QBitmapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitmapPTR().Pointer()
	}
	return nil
}

func QBitmapFromPointer(ptr unsafe.Pointer) *QBitmap {
	var n = new(QBitmap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBitmap) QBitmapPTR() *QBitmap {
	return ptr
}

func (ptr *QBitmap) Clear() {
	if ptr.Pointer() != nil {
		C.QBitmap_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBitmap) Swap(other QBitmapITF) {
	if ptr.Pointer() != nil {
		C.QBitmap_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBitmap(other)))
	}
}

func (ptr *QBitmap) DestroyQBitmap() {
	if ptr.Pointer() != nil {
		C.QBitmap_DestroyQBitmap(C.QtObjectPtr(ptr.Pointer()))
	}
}
