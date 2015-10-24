package gui

//#include "qshowevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QShowEvent struct {
	core.QEvent
}

type QShowEventITF interface {
	core.QEventITF
	QShowEventPTR() *QShowEvent
}

func PointerFromQShowEvent(ptr QShowEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShowEventPTR().Pointer()
	}
	return nil
}

func QShowEventFromPointer(ptr unsafe.Pointer) *QShowEvent {
	var n = new(QShowEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QShowEvent) QShowEventPTR() *QShowEvent {
	return ptr
}

func NewQShowEvent() *QShowEvent {
	return QShowEventFromPointer(unsafe.Pointer(C.QShowEvent_NewQShowEvent()))
}
