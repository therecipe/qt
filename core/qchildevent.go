package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QChildEvent struct {
	QEvent
}

type QChildEvent_ITF interface {
	QEvent_ITF
	QChildEvent_PTR() *QChildEvent
}

func PointerFromQChildEvent(ptr QChildEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChildEvent_PTR().Pointer()
	}
	return nil
}

func NewQChildEventFromPointer(ptr unsafe.Pointer) *QChildEvent {
	var n = new(QChildEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QChildEvent) QChildEvent_PTR() *QChildEvent {
	return ptr
}

func NewQChildEvent(ty QEvent__Type, child QObject_ITF) *QChildEvent {
	defer qt.Recovering("QChildEvent::QChildEvent")

	return NewQChildEventFromPointer(C.QChildEvent_NewQChildEvent(C.int(ty), PointerFromQObject(child)))
}

func (ptr *QChildEvent) Added() bool {
	defer qt.Recovering("QChildEvent::added")

	if ptr.Pointer() != nil {
		return C.QChildEvent_Added(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChildEvent) Child() *QObject {
	defer qt.Recovering("QChildEvent::child")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QChildEvent_Child(ptr.Pointer()))
	}
	return nil
}

func (ptr *QChildEvent) Polished() bool {
	defer qt.Recovering("QChildEvent::polished")

	if ptr.Pointer() != nil {
		return C.QChildEvent_Polished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChildEvent) Removed() bool {
	defer qt.Recovering("QChildEvent::removed")

	if ptr.Pointer() != nil {
		return C.QChildEvent_Removed(ptr.Pointer()) != 0
	}
	return false
}
