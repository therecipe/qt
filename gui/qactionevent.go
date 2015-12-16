package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QActionEvent struct {
	core.QEvent
}

type QActionEvent_ITF interface {
	core.QEvent_ITF
	QActionEvent_PTR() *QActionEvent
}

func PointerFromQActionEvent(ptr QActionEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QActionEvent_PTR().Pointer()
	}
	return nil
}

func NewQActionEventFromPointer(ptr unsafe.Pointer) *QActionEvent {
	var n = new(QActionEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QActionEvent) QActionEvent_PTR() *QActionEvent {
	return ptr
}

func NewQActionEvent(ty int, action unsafe.Pointer, before unsafe.Pointer) *QActionEvent {
	defer qt.Recovering("QActionEvent::QActionEvent")

	return NewQActionEventFromPointer(C.QActionEvent_NewQActionEvent(C.int(ty), action, before))
}

func (ptr *QActionEvent) Action() unsafe.Pointer {
	defer qt.Recovering("QActionEvent::action")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QActionEvent_Action(ptr.Pointer()))
	}
	return nil
}

func (ptr *QActionEvent) Before() unsafe.Pointer {
	defer qt.Recovering("QActionEvent::before")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QActionEvent_Before(ptr.Pointer()))
	}
	return nil
}
