package core

//#include "qqueue.h"
import "C"
import (
	"unsafe"
)

type QQueue struct {
	QList
}

type QQueueITF interface {
	QListITF
	QQueuePTR() *QQueue
}

func PointerFromQQueue(ptr QQueueITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQueuePTR().Pointer()
	}
	return nil
}

func QQueueFromPointer(ptr unsafe.Pointer) *QQueue {
	var n = new(QQueue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQueue) QQueuePTR() *QQueue {
	return ptr
}
