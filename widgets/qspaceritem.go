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

type QSpacerItem_ITF interface {
	QLayoutItem_ITF
	QSpacerItem_PTR() *QSpacerItem
}

func PointerFromQSpacerItem(ptr QSpacerItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpacerItem_PTR().Pointer()
	}
	return nil
}

func NewQSpacerItemFromPointer(ptr unsafe.Pointer) *QSpacerItem {
	var n = new(QSpacerItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSpacerItem) QSpacerItem_PTR() *QSpacerItem {
	return ptr
}

func NewQSpacerItem(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) *QSpacerItem {
	return NewQSpacerItemFromPointer(C.QSpacerItem_NewQSpacerItem(C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy)))
}

func (ptr *QSpacerItem) ChangeSize(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QSpacerItem_ChangeSize(ptr.Pointer(), C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy))
	}
}

func (ptr *QSpacerItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSpacerItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpacerItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSpacerItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSpacerItem) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSpacerItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QSpacerItem) SpacerItem() *QSpacerItem {
	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QSpacerItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) DestroyQSpacerItem() {
	if ptr.Pointer() != nil {
		C.QSpacerItem_DestroyQSpacerItem(ptr.Pointer())
	}
}
