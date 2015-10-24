package core

//#include "qmap.h"
import "C"
import (
	"unsafe"
)

type QMap struct {
	ptr unsafe.Pointer
}

type QMapITF interface {
	QMapPTR() *QMap
}

func (p *QMap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMap(ptr QMapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMapPTR().Pointer()
	}
	return nil
}

func QMapFromPointer(ptr unsafe.Pointer) *QMap {
	var n = new(QMap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMap) QMapPTR() *QMap {
	return ptr
}
