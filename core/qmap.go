package core

//#include "qmap.h"
import "C"
import (
	"unsafe"
)

type QMap struct {
	ptr unsafe.Pointer
}

type QMap_ITF interface {
	QMap_PTR() *QMap
}

func (p *QMap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMap(ptr QMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMap_PTR().Pointer()
	}
	return nil
}

func NewQMapFromPointer(ptr unsafe.Pointer) *QMap {
	var n = new(QMap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMap) QMap_PTR() *QMap {
	return ptr
}
