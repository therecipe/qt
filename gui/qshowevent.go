package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QShowEvent struct {
	core.QEvent
}

type QShowEvent_ITF interface {
	core.QEvent_ITF
	QShowEvent_PTR() *QShowEvent
}

func PointerFromQShowEvent(ptr QShowEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShowEvent_PTR().Pointer()
	}
	return nil
}

func NewQShowEventFromPointer(ptr unsafe.Pointer) *QShowEvent {
	var n = new(QShowEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QShowEvent) QShowEvent_PTR() *QShowEvent {
	return ptr
}

func NewQShowEvent() *QShowEvent {
	defer qt.Recovering("QShowEvent::QShowEvent")

	return NewQShowEventFromPointer(C.QShowEvent_NewQShowEvent())
}
