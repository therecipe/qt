package core

//#include "qstate.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QState struct {
	QAbstractState
}

type QStateITF interface {
	QAbstractStateITF
	QStatePTR() *QState
}

func PointerFromQState(ptr QStateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatePTR().Pointer()
	}
	return nil
}

func QStateFromPointer(ptr unsafe.Pointer) *QState {
	var n = new(QState)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QState) QStatePTR() *QState {
	return ptr
}

//QState::ChildMode
type QState__ChildMode int

var (
	QState__ExclusiveStates = QState__ChildMode(0)
	QState__ParallelStates  = QState__ChildMode(1)
)

//QState::RestorePolicy
type QState__RestorePolicy int

var (
	QState__DontRestoreProperties = QState__RestorePolicy(0)
	QState__RestoreProperties     = QState__RestorePolicy(1)
)

func NewQState2(childMode QState__ChildMode, parent QStateITF) *QState {
	return QStateFromPointer(unsafe.Pointer(C.QState_NewQState2(C.int(childMode), C.QtObjectPtr(PointerFromQState(parent)))))
}

func NewQState(parent QStateITF) *QState {
	return QStateFromPointer(unsafe.Pointer(C.QState_NewQState(C.QtObjectPtr(PointerFromQState(parent)))))
}

func (ptr *QState) AddTransition3(target QAbstractStateITF) *QAbstractTransition {
	if ptr.Pointer() != nil {
		return QAbstractTransitionFromPointer(unsafe.Pointer(C.QState_AddTransition3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(target)))))
	}
	return nil
}

func (ptr *QState) AddTransition2(sender QObjectITF, signal string, target QAbstractStateITF) *QSignalTransition {
	if ptr.Pointer() != nil {
		return QSignalTransitionFromPointer(unsafe.Pointer(C.QState_AddTransition2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)), C.CString(signal), C.QtObjectPtr(PointerFromQAbstractState(target)))))
	}
	return nil
}

func (ptr *QState) AddTransition(transition QAbstractTransitionITF) {
	if ptr.Pointer() != nil {
		C.QState_AddTransition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractTransition(transition)))
	}
}

func (ptr *QState) AssignProperty(object QObjectITF, name string, value string) {
	if ptr.Pointer() != nil {
		C.QState_AssignProperty(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)), C.CString(name), C.CString(value))
	}
}

func (ptr *QState) ChildMode() QState__ChildMode {
	if ptr.Pointer() != nil {
		return QState__ChildMode(C.QState_ChildMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QState) ConnectChildModeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QState_ConnectChildModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "childModeChanged", f)
	}
}

func (ptr *QState) DisconnectChildModeChanged() {
	if ptr.Pointer() != nil {
		C.QState_DisconnectChildModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "childModeChanged")
	}
}

//export callbackQStateChildModeChanged
func callbackQStateChildModeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "childModeChanged").(func())()
}

func (ptr *QState) ErrorState() *QAbstractState {
	if ptr.Pointer() != nil {
		return QAbstractStateFromPointer(unsafe.Pointer(C.QState_ErrorState(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QState) ConnectErrorStateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QState_ConnectErrorStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "errorStateChanged", f)
	}
}

func (ptr *QState) DisconnectErrorStateChanged() {
	if ptr.Pointer() != nil {
		C.QState_DisconnectErrorStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "errorStateChanged")
	}
}

//export callbackQStateErrorStateChanged
func callbackQStateErrorStateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "errorStateChanged").(func())()
}

func (ptr *QState) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QState_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QState) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QState_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQStateFinished
func callbackQStateFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QState) InitialState() *QAbstractState {
	if ptr.Pointer() != nil {
		return QAbstractStateFromPointer(unsafe.Pointer(C.QState_InitialState(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QState) ConnectInitialStateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QState_ConnectInitialStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "initialStateChanged", f)
	}
}

func (ptr *QState) DisconnectInitialStateChanged() {
	if ptr.Pointer() != nil {
		C.QState_DisconnectInitialStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "initialStateChanged")
	}
}

//export callbackQStateInitialStateChanged
func callbackQStateInitialStateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "initialStateChanged").(func())()
}

func (ptr *QState) ConnectPropertiesAssigned(f func()) {
	if ptr.Pointer() != nil {
		C.QState_ConnectPropertiesAssigned(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "propertiesAssigned", f)
	}
}

func (ptr *QState) DisconnectPropertiesAssigned() {
	if ptr.Pointer() != nil {
		C.QState_DisconnectPropertiesAssigned(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "propertiesAssigned")
	}
}

//export callbackQStatePropertiesAssigned
func callbackQStatePropertiesAssigned(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "propertiesAssigned").(func())()
}

func (ptr *QState) RemoveTransition(transition QAbstractTransitionITF) {
	if ptr.Pointer() != nil {
		C.QState_RemoveTransition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractTransition(transition)))
	}
}

func (ptr *QState) SetChildMode(mode QState__ChildMode) {
	if ptr.Pointer() != nil {
		C.QState_SetChildMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QState) SetErrorState(state QAbstractStateITF) {
	if ptr.Pointer() != nil {
		C.QState_SetErrorState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(state)))
	}
}

func (ptr *QState) SetInitialState(state QAbstractStateITF) {
	if ptr.Pointer() != nil {
		C.QState_SetInitialState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(state)))
	}
}

func (ptr *QState) DestroyQState() {
	if ptr.Pointer() != nil {
		C.QState_DestroyQState(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
