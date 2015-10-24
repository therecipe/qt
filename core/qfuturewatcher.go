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

type QFutureWatcherITF interface {
	QObjectITF
	QFutureWatcherPTR() *QFutureWatcher
}

func PointerFromQFutureWatcher(ptr QFutureWatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFutureWatcherPTR().Pointer()
	}
	return nil
}

func QFutureWatcherFromPointer(ptr unsafe.Pointer) *QFutureWatcher {
	var n = new(QFutureWatcher)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFutureWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFutureWatcher) QFutureWatcherPTR() *QFutureWatcher {
	return ptr
}
