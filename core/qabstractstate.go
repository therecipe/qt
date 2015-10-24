package core

//#include "qabstractstate.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractState struct {
	QObject
}

type QAbstractStateITF interface {
	QObjectITF
	QAbstractStatePTR() *QAbstractState
}

func PointerFromQAbstractState(ptr QAbstractStateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractStatePTR().Pointer()
	}
	return nil
}

func QAbstractStateFromPointer(ptr unsafe.Pointer) *QAbstractState {
	var n = new(QAbstractState)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractState) QAbstractStatePTR() *QAbstractState {
	return ptr
}

func (ptr *QAbstractState) Active() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractState_Active(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractState) ConnectActiveChanged(f func(active bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractState) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractStateActiveChanged
func callbackQAbstractStateActiveChanged(ptrName *C.char, active C.int) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func(bool))(int(active) != 0)
}

func (ptr *QAbstractState) ConnectEntered(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractState) DisconnectEntered() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractStateEntered
func callbackQAbstractStateEntered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "entered").(func())()
}

func (ptr *QAbstractState) ConnectExited(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectExited(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "exited", f)
	}
}

func (ptr *QAbstractState) DisconnectExited() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectExited(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "exited")
	}
}

//export callbackQAbstractStateExited
func callbackQAbstractStateExited(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "exited").(func())()
}

func (ptr *QAbstractState) Machine() *QStateMachine {
	if ptr.Pointer() != nil {
		return QStateMachineFromPointer(unsafe.Pointer(C.QAbstractState_Machine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractState) ParentState() *QState {
	if ptr.Pointer() != nil {
		return QStateFromPointer(unsafe.Pointer(C.QAbstractState_ParentState(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractState) DestroyQAbstractState() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DestroyQAbstractState(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
