package gui

//#include "qactionevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QActionEvent struct {
	core.QEvent
}

type QActionEventITF interface {
	core.QEventITF
	QActionEventPTR() *QActionEvent
}

func PointerFromQActionEvent(ptr QActionEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QActionEventPTR().Pointer()
	}
	return nil
}

func QActionEventFromPointer(ptr unsafe.Pointer) *QActionEvent {
	var n = new(QActionEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QActionEvent) QActionEventPTR() *QActionEvent {
	return ptr
}
