package core

//#include "qeventtransition.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QEventTransition struct {
	QAbstractTransition
}

type QEventTransition_ITF interface {
	QAbstractTransition_ITF
	QEventTransition_PTR() *QEventTransition
}

func PointerFromQEventTransition(ptr QEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQEventTransitionFromPointer(ptr unsafe.Pointer) *QEventTransition {
	var n = new(QEventTransition)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QEventTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QEventTransition) QEventTransition_PTR() *QEventTransition {
	return ptr
}

func NewQEventTransition2(object QObject_ITF, ty QEvent__Type, sourceState QState_ITF) *QEventTransition {
	return NewQEventTransitionFromPointer(C.QEventTransition_NewQEventTransition2(PointerFromQObject(object), C.int(ty), PointerFromQState(sourceState)))
}

func NewQEventTransition(sourceState QState_ITF) *QEventTransition {
	return NewQEventTransitionFromPointer(C.QEventTransition_NewQEventTransition(PointerFromQState(sourceState)))
}

func (ptr *QEventTransition) EventSource() *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QEventTransition_EventSource(ptr.Pointer()))
	}
	return nil
}

func (ptr *QEventTransition) EventType() QEvent__Type {
	if ptr.Pointer() != nil {
		return QEvent__Type(C.QEventTransition_EventType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEventTransition) SetEventSource(object QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QEventTransition_SetEventSource(ptr.Pointer(), PointerFromQObject(object))
	}
}

func (ptr *QEventTransition) SetEventType(ty QEvent__Type) {
	if ptr.Pointer() != nil {
		C.QEventTransition_SetEventType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QEventTransition) DestroyQEventTransition() {
	if ptr.Pointer() != nil {
		C.QEventTransition_DestroyQEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
