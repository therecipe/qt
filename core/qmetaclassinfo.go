package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMetaClassInfo struct {
	ptr unsafe.Pointer
}

type QMetaClassInfo_ITF interface {
	QMetaClassInfo_PTR() *QMetaClassInfo
}

func (p *QMetaClassInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaClassInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaClassInfo(ptr QMetaClassInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaClassInfo_PTR().Pointer()
	}
	return nil
}

func NewQMetaClassInfoFromPointer(ptr unsafe.Pointer) *QMetaClassInfo {
	var n = new(QMetaClassInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaClassInfo) QMetaClassInfo_PTR() *QMetaClassInfo {
	return ptr
}
