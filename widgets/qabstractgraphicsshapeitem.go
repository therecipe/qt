package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
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
	for len(n.ObjectNameAbs()) < len("QAbstractGraphicsShapeItem_") {
		n.SetObjectNameAbs("QAbstractGraphicsShapeItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractGraphicsShapeItem) QAbstractGraphicsShapeItem_PTR() *QAbstractGraphicsShapeItem {
	return ptr
}

func (ptr *QAbstractGraphicsShapeItem) Brush() *gui.QBrush {
	defer qt.Recovering("QAbstractGraphicsShapeItem::brush")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QAbstractGraphicsShapeItem_Brush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractGraphicsShapeItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QAbstractGraphicsShapeItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QAbstractGraphicsShapeItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QAbstractGraphicsShapeItem) SetBrush(brush gui.QBrush_ITF) {
	defer qt.Recovering("QAbstractGraphicsShapeItem::setBrush")

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QAbstractGraphicsShapeItem) SetPen(pen gui.QPen_ITF) {
	defer qt.Recovering("QAbstractGraphicsShapeItem::setPen")

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetPen(ptr.Pointer(), gui.PointerFromQPen(pen))
	}
}

func (ptr *QAbstractGraphicsShapeItem) DestroyQAbstractGraphicsShapeItem() {
	defer qt.Recovering("QAbstractGraphicsShapeItem::~QAbstractGraphicsShapeItem")

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_DestroyQAbstractGraphicsShapeItem(ptr.Pointer())
	}
}

func (ptr *QAbstractGraphicsShapeItem) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractGraphicsShapeItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractGraphicsShapeItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractGraphicsShapeItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractGraphicsShapeItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractGraphicsShapeItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
