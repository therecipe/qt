package core

//#include "qlist.h"
import "C"
import (
	"unsafe"
)

type QList struct {
	ptr unsafe.Pointer
}

type QListITF interface {
	QListPTR() *QList
}

func (p *QList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQList(ptr QListITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListPTR().Pointer()
	}
	return nil
}

func QListFromPointer(ptr unsafe.Pointer) *QList {
	var n = new(QList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QList) QListPTR() *QList {
	return ptr
}
