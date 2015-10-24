package core

//#include "qbytearraylist.h"
import "C"
import (
	"unsafe"
)

type QByteArrayList struct {
	QList
}

type QByteArrayListITF interface {
	QListITF
	QByteArrayListPTR() *QByteArrayList
}

func PointerFromQByteArrayList(ptr QByteArrayListITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArrayListPTR().Pointer()
	}
	return nil
}

func QByteArrayListFromPointer(ptr unsafe.Pointer) *QByteArrayList {
	var n = new(QByteArrayList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QByteArrayList) QByteArrayListPTR() *QByteArrayList {
	return ptr
}
