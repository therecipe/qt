package core

//#include "qchildevent.h"
import "C"
import (
	"unsafe"
)

type QChildEvent struct {
	QEvent
}

type QChildEventITF interface {
	QEventITF
	QChildEventPTR() *QChildEvent
}

func PointerFromQChildEvent(ptr QChildEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChildEventPTR().Pointer()
	}
	return nil
}

func QChildEventFromPointer(ptr unsafe.Pointer) *QChildEvent {
	var n = new(QChildEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QChildEvent) QChildEventPTR() *QChildEvent {
	return ptr
}

func NewQChildEvent(ty QEvent__Type, child QObjectITF) *QChildEvent {
	return QChildEventFromPointer(unsafe.Pointer(C.QChildEvent_NewQChildEvent(C.int(ty), C.QtObjectPtr(PointerFromQObject(child)))))
}

func (ptr *QChildEvent) Added() bool {
	if ptr.Pointer() != nil {
		return C.QChildEvent_Added(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChildEvent) Child() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QChildEvent_Child(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QChildEvent) Polished() bool {
	if ptr.Pointer() != nil {
		return C.QChildEvent_Polished(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QChildEvent) Removed() bool {
	if ptr.Pointer() != nil {
		return C.QChildEvent_Removed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
