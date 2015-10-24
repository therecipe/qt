package core

//#include "qmultimap.h"
import "C"
import (
	"unsafe"
)

type QMultiMap struct {
	QMap
}

type QMultiMapITF interface {
	QMapITF
	QMultiMapPTR() *QMultiMap
}

func PointerFromQMultiMap(ptr QMultiMapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMultiMapPTR().Pointer()
	}
	return nil
}

func QMultiMapFromPointer(ptr unsafe.Pointer) *QMultiMap {
	var n = new(QMultiMap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMultiMap) QMultiMapPTR() *QMultiMap {
	return ptr
}
