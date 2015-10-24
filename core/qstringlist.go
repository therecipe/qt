package core

//#include "qstringlist.h"
import "C"
import (
	"unsafe"
)

type QStringList struct {
	QList
}

type QStringListITF interface {
	QListITF
	QStringListPTR() *QStringList
}

func PointerFromQStringList(ptr QStringListITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringListPTR().Pointer()
	}
	return nil
}

func QStringListFromPointer(ptr unsafe.Pointer) *QStringList {
	var n = new(QStringList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStringList) QStringListPTR() *QStringList {
	return ptr
}
