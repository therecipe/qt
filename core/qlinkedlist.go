package core

//#include "qlinkedlist.h"
import "C"
import (
	"unsafe"
)

type QLinkedList struct {
	ptr unsafe.Pointer
}

type QLinkedListITF interface {
	QLinkedListPTR() *QLinkedList
}

func (p *QLinkedList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLinkedList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLinkedList(ptr QLinkedListITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinkedListPTR().Pointer()
	}
	return nil
}

func QLinkedListFromPointer(ptr unsafe.Pointer) *QLinkedList {
	var n = new(QLinkedList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLinkedList) QLinkedListPTR() *QLinkedList {
	return ptr
}
