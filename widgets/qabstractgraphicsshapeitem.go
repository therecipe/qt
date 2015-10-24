package widgets

//#include "qabstractgraphicsshapeitem.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractGraphicsShapeItem struct {
	QGraphicsItem
}

type QAbstractGraphicsShapeItemITF interface {
	QGraphicsItemITF
	QAbstractGraphicsShapeItemPTR() *QAbstractGraphicsShapeItem
}

func PointerFromQAbstractGraphicsShapeItem(ptr QAbstractGraphicsShapeItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractGraphicsShapeItemPTR().Pointer()
	}
	return nil
}

func QAbstractGraphicsShapeItemFromPointer(ptr unsafe.Pointer) *QAbstractGraphicsShapeItem {
	var n = new(QAbstractGraphicsShapeItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractGraphicsShapeItem) QAbstractGraphicsShapeItemPTR() *QAbstractGraphicsShapeItem {
	return ptr
}

func (ptr *QAbstractGraphicsShapeItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractGraphicsShapeItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QAbstractGraphicsShapeItem) SetBrush(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QAbstractGraphicsShapeItem) SetPen(pen gui.QPenITF) {
	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetPen(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPen(pen)))
	}
}

func (ptr *QAbstractGraphicsShapeItem) DestroyQAbstractGraphicsShapeItem() {
	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_DestroyQAbstractGraphicsShapeItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
