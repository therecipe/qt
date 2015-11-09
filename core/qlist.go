package core

//#include "qlist.h"
import "C"
import (
	"unsafe"
)

type QList struct {
	ptr unsafe.Pointer
}

type QList_ITF interface {
	QList_PTR() *QList
}

func (p *QList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQList(ptr QList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QList_PTR().Pointer()
	}
	return nil
}

func NewQListFromPointer(ptr unsafe.Pointer) *QList {
	var n = new(QList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QList) QList_PTR() *QList {
	return ptr
}
