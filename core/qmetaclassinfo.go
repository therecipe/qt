package core

//#include "qmetaclassinfo.h"
import "C"
import (
	"unsafe"
)

type QMetaClassInfo struct {
	ptr unsafe.Pointer
}

type QMetaClassInfoITF interface {
	QMetaClassInfoPTR() *QMetaClassInfo
}

func (p *QMetaClassInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaClassInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaClassInfo(ptr QMetaClassInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaClassInfoPTR().Pointer()
	}
	return nil
}

func QMetaClassInfoFromPointer(ptr unsafe.Pointer) *QMetaClassInfo {
	var n = new(QMetaClassInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaClassInfo) QMetaClassInfoPTR() *QMetaClassInfo {
	return ptr
}
