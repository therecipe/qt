package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QByteArrayList struct {
	QList
}

type QByteArrayList_ITF interface {
	QList_ITF
	QByteArrayList_PTR() *QByteArrayList
}

func PointerFromQByteArrayList(ptr QByteArrayList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArrayList_PTR().Pointer()
	}
	return nil
}

func NewQByteArrayListFromPointer(ptr unsafe.Pointer) *QByteArrayList {
	var n = new(QByteArrayList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QByteArrayList) QByteArrayList_PTR() *QByteArrayList {
	return ptr
}

func (ptr *QByteArrayList) Join() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QByteArrayList::join")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArrayList_Join(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArrayList) Join3(separator string) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QByteArrayList::join")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArrayList_Join3(ptr.Pointer(), C.CString(separator)))
	}
	return nil
}

func (ptr *QByteArrayList) Join2(separator QByteArray_ITF) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QByteArrayList::join")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArrayList_Join2(ptr.Pointer(), PointerFromQByteArray(separator)))
	}
	return nil
}
