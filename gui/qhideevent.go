package gui

//#include "qhideevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHideEvent struct {
	core.QEvent
}

type QHideEventITF interface {
	core.QEventITF
	QHideEventPTR() *QHideEvent
}

func PointerFromQHideEvent(ptr QHideEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHideEventPTR().Pointer()
	}
	return nil
}

func QHideEventFromPointer(ptr unsafe.Pointer) *QHideEvent {
	var n = new(QHideEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHideEvent) QHideEventPTR() *QHideEvent {
	return ptr
}

func NewQHideEvent() *QHideEvent {
	return QHideEventFromPointer(unsafe.Pointer(C.QHideEvent_NewQHideEvent()))
}
