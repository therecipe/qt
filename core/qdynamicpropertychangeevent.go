package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDynamicPropertyChangeEvent struct {
	QEvent
}

type QDynamicPropertyChangeEvent_ITF interface {
	QEvent_ITF
	QDynamicPropertyChangeEvent_PTR() *QDynamicPropertyChangeEvent
}

func PointerFromQDynamicPropertyChangeEvent(ptr QDynamicPropertyChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDynamicPropertyChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQDynamicPropertyChangeEventFromPointer(ptr unsafe.Pointer) *QDynamicPropertyChangeEvent {
	var n = new(QDynamicPropertyChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDynamicPropertyChangeEvent) QDynamicPropertyChangeEvent_PTR() *QDynamicPropertyChangeEvent {
	return ptr
}

func NewQDynamicPropertyChangeEvent(name QByteArray_ITF) *QDynamicPropertyChangeEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDynamicPropertyChangeEvent::QDynamicPropertyChangeEvent")
		}
	}()

	return NewQDynamicPropertyChangeEventFromPointer(C.QDynamicPropertyChangeEvent_NewQDynamicPropertyChangeEvent(PointerFromQByteArray(name)))
}

func (ptr *QDynamicPropertyChangeEvent) PropertyName() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDynamicPropertyChangeEvent::propertyName")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QDynamicPropertyChangeEvent_PropertyName(ptr.Pointer()))
	}
	return nil
}
