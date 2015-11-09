package core

//#include "qfuturewatcher.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFutureWatcher struct {
	QObject
}

type QFutureWatcher_ITF interface {
	QObject_ITF
	QFutureWatcher_PTR() *QFutureWatcher
}

func PointerFromQFutureWatcher(ptr QFutureWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFutureWatcher_PTR().Pointer()
	}
	return nil
}

func NewQFutureWatcherFromPointer(ptr unsafe.Pointer) *QFutureWatcher {
	var n = new(QFutureWatcher)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QFutureWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFutureWatcher) QFutureWatcher_PTR() *QFutureWatcher {
	return ptr
}
