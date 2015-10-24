package core

//#include "qwineventnotifier.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QWinEventNotifier struct {
	QObject
}

type QWinEventNotifierITF interface {
	QObjectITF
	QWinEventNotifierPTR() *QWinEventNotifier
}

func PointerFromQWinEventNotifier(ptr QWinEventNotifierITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinEventNotifierPTR().Pointer()
	}
	return nil
}

func QWinEventNotifierFromPointer(ptr unsafe.Pointer) *QWinEventNotifier {
	var n = new(QWinEventNotifier)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWinEventNotifier_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWinEventNotifier) QWinEventNotifierPTR() *QWinEventNotifier {
	return ptr
}
