package help

//#include "qhelpcontentitem.h"
import "C"
import (
	"unsafe"
)

type QHelpContentItem struct {
	ptr unsafe.Pointer
}

type QHelpContentItem_ITF interface {
	QHelpContentItem_PTR() *QHelpContentItem
}

func (p *QHelpContentItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHelpContentItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHelpContentItem(ptr QHelpContentItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItem_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentItemFromPointer(ptr unsafe.Pointer) *QHelpContentItem {
	var n = new(QHelpContentItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpContentItem) QHelpContentItem_PTR() *QHelpContentItem {
	return ptr
}

func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Child(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QHelpContentItem) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_ChildPosition(ptr.Pointer(), PointerFromQHelpContentItem(child)))
	}
	return 0
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentItem) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpContentItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	if ptr.Pointer() != nil {
		C.QHelpContentItem_DestroyQHelpContentItem(ptr.Pointer())
	}
}
