package gui

//#include "qcloseevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCloseEvent struct {
	core.QEvent
}

type QCloseEventITF interface {
	core.QEventITF
	QCloseEventPTR() *QCloseEvent
}

func PointerFromQCloseEvent(ptr QCloseEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCloseEventPTR().Pointer()
	}
	return nil
}

func QCloseEventFromPointer(ptr unsafe.Pointer) *QCloseEvent {
	var n = new(QCloseEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCloseEvent) QCloseEventPTR() *QCloseEvent {
	return ptr
}

func NewQCloseEvent() *QCloseEvent {
	return QCloseEventFromPointer(unsafe.Pointer(C.QCloseEvent_NewQCloseEvent()))
}
