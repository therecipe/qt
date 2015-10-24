package core

//#include "qdynamicpropertychangeevent.h"
import "C"
import (
	"unsafe"
)

type QDynamicPropertyChangeEvent struct {
	QEvent
}

type QDynamicPropertyChangeEventITF interface {
	QEventITF
	QDynamicPropertyChangeEventPTR() *QDynamicPropertyChangeEvent
}

func PointerFromQDynamicPropertyChangeEvent(ptr QDynamicPropertyChangeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDynamicPropertyChangeEventPTR().Pointer()
	}
	return nil
}

func QDynamicPropertyChangeEventFromPointer(ptr unsafe.Pointer) *QDynamicPropertyChangeEvent {
	var n = new(QDynamicPropertyChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDynamicPropertyChangeEvent) QDynamicPropertyChangeEventPTR() *QDynamicPropertyChangeEvent {
	return ptr
}

func NewQDynamicPropertyChangeEvent(name QByteArrayITF) *QDynamicPropertyChangeEvent {
	return QDynamicPropertyChangeEventFromPointer(unsafe.Pointer(C.QDynamicPropertyChangeEvent_NewQDynamicPropertyChangeEvent(C.QtObjectPtr(PointerFromQByteArray(name)))))
}
