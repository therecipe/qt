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

type QEventTransitionITF interface {
	QAbstractTransitionITF
	QEventTransitionPTR() *QEventTransition
}

func PointerFromQEventTransition(ptr QEventTransitionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventTransitionPTR().Pointer()
	}
	return nil
}

func QEventTransitionFromPointer(ptr unsafe.Pointer) *QEventTransition {
	var n = new(QEventTransition)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QEventTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QEventTransition) QEventTransitionPTR() *QEventTransition {
	return ptr
}

func NewQEventTransition2(object QObjectITF, ty QEvent__Type, sourceState QStateITF) *QEventTransition {
	return QEventTransitionFromPointer(unsafe.Pointer(C.QEventTransition_NewQEventTransition2(C.QtObjectPtr(PointerFromQObject(object)), C.int(ty), C.QtObjectPtr(PointerFromQState(sourceState)))))
}

func NewQEventTransition(sourceState QStateITF) *QEventTransition {
	return QEventTransitionFromPointer(unsafe.Pointer(C.QEventTransition_NewQEventTransition(C.QtObjectPtr(PointerFromQState(sourceState)))))
}

func (ptr *QEventTransition) EventSource() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QEventTransition_EventSource(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QEventTransition) EventType() QEvent__Type {
	if ptr.Pointer() != nil {
		return QEvent__Type(C.QEventTransition_EventType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QEventTransition) SetEventSource(object QObjectITF) {
	if ptr.Pointer() != nil {
		C.QEventTransition_SetEventSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)))
	}
}

func (ptr *QEventTransition) SetEventType(ty QEvent__Type) {
	if ptr.Pointer() != nil {
		C.QEventTransition_SetEventType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QEventTransition) DestroyQEventTransition() {
	if ptr.Pointer() != nil {
		C.QEventTransition_DestroyQEventTransition(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
