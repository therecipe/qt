package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMultiMap struct {
	QMap
}

type QMultiMap_ITF interface {
	QMap_ITF
	QMultiMap_PTR() *QMultiMap
}

func PointerFromQMultiMap(ptr QMultiMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMultiMap_PTR().Pointer()
	}
	return nil
}

func NewQMultiMapFromPointer(ptr unsafe.Pointer) *QMultiMap {
	var n = new(QMultiMap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMultiMap) QMultiMap_PTR() *QMultiMap {
	return ptr
}
