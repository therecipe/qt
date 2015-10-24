package help

//#include "qhelpcontentitem.h"
import "C"
import (
	"unsafe"
)

type QHelpContentItem struct {
	ptr unsafe.Pointer
}

type QHelpContentItemITF interface {
	QHelpContentItemPTR() *QHelpContentItem
}

func (p *QHelpContentItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHelpContentItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHelpContentItem(ptr QHelpContentItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItemPTR().Pointer()
	}
	return nil
}

func QHelpContentItemFromPointer(ptr unsafe.Pointer) *QHelpContentItem {
	var n = new(QHelpContentItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpContentItem) QHelpContentItemPTR() *QHelpContentItem {
	return ptr
}

func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return QHelpContentItemFromPointer(unsafe.Pointer(C.QHelpContentItem_Child(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QHelpContentItem) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_ChildCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_ChildPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHelpContentItem(child))))
	}
	return 0
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	if ptr.Pointer() != nil {
		return QHelpContentItemFromPointer(unsafe.Pointer(C.QHelpContentItem_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpContentItem) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_Row(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpContentItem_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHelpContentItem) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpContentItem_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	if ptr.Pointer() != nil {
		C.QHelpContentItem_DestroyQHelpContentItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
