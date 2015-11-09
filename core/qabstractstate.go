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

type QAbstractState_ITF interface {
	QObject_ITF
	QAbstractState_PTR() *QAbstractState
}

func PointerFromQAbstractState(ptr QAbstractState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractState_PTR().Pointer()
	}
	return nil
}

func NewQAbstractStateFromPointer(ptr unsafe.Pointer) *QAbstractState {
	var n = new(QAbstractState)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAbstractState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractState) QAbstractState_PTR() *QAbstractState {
	return ptr
}

func (ptr *QAbstractState) Active() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractState_Active(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractState) ConnectActiveChanged(f func(active bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractState) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractStateActiveChanged
func callbackQAbstractStateActiveChanged(ptrName *C.char, active C.int) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func(bool))(int(active) != 0)
}

func (ptr *QAbstractState) ConnectEntered(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractState) DisconnectEntered() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractStateEntered
func callbackQAbstractStateEntered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "entered").(func())()
}

func (ptr *QAbstractState) ConnectExited(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectExited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "exited", f)
	}
}

func (ptr *QAbstractState) DisconnectExited() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectExited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "exited")
	}
}

//export callbackQAbstractStateExited
func callbackQAbstractStateExited(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "exited").(func())()
}

func (ptr *QAbstractState) Machine() *QStateMachine {
	if ptr.Pointer() != nil {
		return NewQStateMachineFromPointer(C.QAbstractState_Machine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) ParentState() *QState {
	if ptr.Pointer() != nil {
		return NewQStateFromPointer(C.QAbstractState_ParentState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) DestroyQAbstractState() {
	if ptr.Pointer() != nil {
		C.QAbstractState_DestroyQAbstractState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
