package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QLinkedList struct {
	ptr unsafe.Pointer
}

type QLinkedList_ITF interface {
	QLinkedList_PTR() *QLinkedList
}

func (p *QLinkedList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLinkedList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLinkedList(ptr QLinkedList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinkedList_PTR().Pointer()
	}
	return nil
}

func NewQLinkedListFromPointer(ptr unsafe.Pointer) *QLinkedList {
	var n = new(QLinkedList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLinkedList) QLinkedList_PTR() *QLinkedList {
	return ptr
}
