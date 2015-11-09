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

type QHideEvent_ITF interface {
	core.QEvent_ITF
	QHideEvent_PTR() *QHideEvent
}

func PointerFromQHideEvent(ptr QHideEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHideEvent_PTR().Pointer()
	}
	return nil
}

func NewQHideEventFromPointer(ptr unsafe.Pointer) *QHideEvent {
	var n = new(QHideEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHideEvent) QHideEvent_PTR() *QHideEvent {
	return ptr
}

func NewQHideEvent() *QHideEvent {
	return NewQHideEventFromPointer(C.QHideEvent_NewQHideEvent())
}
