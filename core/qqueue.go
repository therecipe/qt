package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QQueue struct {
	QList
}

type QQueue_ITF interface {
	QList_ITF
	QQueue_PTR() *QQueue
}

func PointerFromQQueue(ptr QQueue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQueue_PTR().Pointer()
	}
	return nil
}

func NewQQueueFromPointer(ptr unsafe.Pointer) *QQueue {
	var n = new(QQueue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQueue) QQueue_PTR() *QQueue {
	return ptr
}
