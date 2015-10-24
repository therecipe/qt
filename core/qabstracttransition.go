package core

//#include "qabstracttransition.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractTransition struct {
	QObject
}

type QAbstractTransitionITF interface {
	QObjectITF
	QAbstractTransitionPTR() *QAbstractTransition
}

func PointerFromQAbstractTransition(ptr QAbstractTransitionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTransitionPTR().Pointer()
	}
	return nil
}

func QAbstractTransitionFromPointer(ptr unsafe.Pointer) *QAbstractTransition {
	var n = new(QAbstractTransition)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractTransition) QAbstractTransitionPTR() *QAbstractTransition {
	return ptr
}

//QAbstractTransition::TransitionType
type QAbstractTransition__TransitionType int

var (
	QAbstractTransition__ExternalTransition = QAbstractTransition__TransitionType(0)
	QAbstractTransition__InternalTransition = QAbstractTransition__TransitionType(1)
)

func (ptr *QAbstractTransition) AddAnimation(animation QAbstractAnimationITF) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_AddAnimation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractAnimation(animation)))
	}
}

func (ptr *QAbstractTransition) Machine() *QStateMachine {
	if ptr.Pointer() != nil {
		return QStateMachineFromPointer(unsafe.Pointer(C.QAbstractTransition_Machine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractTransition) RemoveAnimation(animation QAbstractAnimationITF) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_RemoveAnimation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractAnimation(animation)))
	}
}

func (ptr *QAbstractTransition) SetTargetState(target QAbstractStateITF) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_SetTargetState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(target)))
	}
}

func (ptr *QAbstractTransition) SetTransitionType(ty QAbstractTransition__TransitionType) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_SetTransitionType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QAbstractTransition) SourceState() *QState {
	if ptr.Pointer() != nil {
		return QStateFromPointer(unsafe.Pointer(C.QAbstractTransition_SourceState(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractTransition) TargetState() *QAbstractState {
	if ptr.Pointer() != nil {
		return QAbstractStateFromPointer(unsafe.Pointer(C.QAbstractTransition_TargetState(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractTransition) ConnectTargetStateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_ConnectTargetStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "targetStateChanged", f)
	}
}

func (ptr *QAbstractTransition) DisconnectTargetStateChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_DisconnectTargetStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "targetStateChanged")
	}
}

//export callbackQAbstractTransitionTargetStateChanged
func callbackQAbstractTransitionTargetStateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "targetStateChanged").(func())()
}

func (ptr *QAbstractTransition) ConnectTargetStatesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_ConnectTargetStatesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "targetStatesChanged", f)
	}
}

func (ptr *QAbstractTransition) DisconnectTargetStatesChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_DisconnectTargetStatesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "targetStatesChanged")
	}
}

//export callbackQAbstractTransitionTargetStatesChanged
func callbackQAbstractTransitionTargetStatesChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "targetStatesChanged").(func())()
}

func (ptr *QAbstractTransition) TransitionType() QAbstractTransition__TransitionType {
	if ptr.Pointer() != nil {
		return QAbstractTransition__TransitionType(C.QAbstractTransition_TransitionType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractTransition) ConnectTriggered(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_ConnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QAbstractTransition) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_DisconnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQAbstractTransitionTriggered
func callbackQAbstractTransitionTriggered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func())()
}

func (ptr *QAbstractTransition) DestroyQAbstractTransition() {
	if ptr.Pointer() != nil {
		C.QAbstractTransition_DestroyQAbstractTransition(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
