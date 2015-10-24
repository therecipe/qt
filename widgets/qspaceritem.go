package widgets

//#include "qspaceritem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSpacerItem struct {
	QLayoutItem
}

type QSpacerItemITF interface {
	QLayoutItemITF
	QSpacerItemPTR() *QSpacerItem
}

func PointerFromQSpacerItem(ptr QSpacerItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpacerItemPTR().Pointer()
	}
	return nil
}

func QSpacerItemFromPointer(ptr unsafe.Pointer) *QSpacerItem {
	var n = new(QSpacerItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSpacerItem) QSpacerItemPTR() *QSpacerItem {
	return ptr
}

func NewQSpacerItem(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) *QSpacerItem {
	return QSpacerItemFromPointer(unsafe.Pointer(C.QSpacerItem_NewQSpacerItem(C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy))))
}

func (ptr *QSpacerItem) ChangeSize(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QSpacerItem_ChangeSize(C.QtObjectPtr(ptr.Pointer()), C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy))
	}
}

func (ptr *QSpacerItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSpacerItem_ExpandingDirections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSpacerItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSpacerItem_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSpacerItem) SetGeometry(r core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QSpacerItem_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(r)))
	}
}

func (ptr *QSpacerItem) SpacerItem() *QSpacerItem {
	if ptr.Pointer() != nil {
		return QSpacerItemFromPointer(unsafe.Pointer(C.QSpacerItem_SpacerItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSpacerItem) DestroyQSpacerItem() {
	if ptr.Pointer() != nil {
		C.QSpacerItem_DestroyQSpacerItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
