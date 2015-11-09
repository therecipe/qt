package core

//#include "qstringlist.h"
import "C"
import (
	"unsafe"
)

type QStringList struct {
	QList
}

type QStringList_ITF interface {
	QList_ITF
	QStringList_PTR() *QStringList
}

func PointerFromQStringList(ptr QStringList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringList_PTR().Pointer()
	}
	return nil
}

func NewQStringListFromPointer(ptr unsafe.Pointer) *QStringList {
	var n = new(QStringList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStringList) QStringList_PTR() *QStringList {
	return ptr
}
