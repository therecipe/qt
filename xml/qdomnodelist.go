package xml

//#include "qdomnodelist.h"
import "C"
import (
	"unsafe"
)

type QDomNodeList struct {
	ptr unsafe.Pointer
}

type QDomNodeList_ITF interface {
	QDomNodeList_PTR() *QDomNodeList
}

func (p *QDomNodeList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNodeList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNodeList(ptr QDomNodeList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNodeList_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeListFromPointer(ptr unsafe.Pointer) *QDomNodeList {
	var n = new(QDomNodeList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNodeList) QDomNodeList_PTR() *QDomNodeList {
	return ptr
}

func NewQDomNodeList() *QDomNodeList {
	return NewQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList())
}

func NewQDomNodeList2(n QDomNodeList_ITF) *QDomNodeList {
	return NewQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList2(PointerFromQDomNodeList(n)))
}

func (ptr *QDomNodeList) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QDomNodeList_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNodeList) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) DestroyQDomNodeList() {
	if ptr.Pointer() != nil {
		C.QDomNodeList_DestroyQDomNodeList(ptr.Pointer())
	}
}
