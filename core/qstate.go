package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QState struct {
	QAbstractState
}

type QState_ITF interface {
	QAbstractState_ITF
	QState_PTR() *QState
}

func PointerFromQState(ptr QState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QState_PTR().Pointer()
	}
	return nil
}

func NewQStateFromPointer(ptr unsafe.Pointer) *QState {
	var n = new(QState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QState_") {
		n.SetObjectName("QState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QState) QState_PTR() *QState {
	return ptr
}

//QState::ChildMode
type QState__ChildMode int64

const (
	QState__ExclusiveStates = QState__ChildMode(0)
	QState__ParallelStates  = QState__ChildMode(1)
)

//QState::RestorePolicy
type QState__RestorePolicy int64

const (
	QState__DontRestoreProperties = QState__RestorePolicy(0)
	QState__RestoreProperties     = QState__RestorePolicy(1)
)

func NewQState2(childMode QState__ChildMode, parent QState_ITF) *QState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::QState")
		}
	}()

	return NewQStateFromPointer(C.QState_NewQState2(C.int(childMode), PointerFromQState(parent)))
}

func NewQState(parent QState_ITF) *QState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::QState")
		}
	}()

	return NewQStateFromPointer(C.QState_NewQState(PointerFromQState(parent)))
}

func (ptr *QState) AddTransition3(target QAbstractState_ITF) *QAbstractTransition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::addTransition")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractTransitionFromPointer(C.QState_AddTransition3(ptr.Pointer(), PointerFromQAbstractState(target)))
	}
	return nil
}

func (ptr *QState) AddTransition2(sender QObject_ITF, signal string, target QAbstractState_ITF) *QSignalTransition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::addTransition")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSignalTransitionFromPointer(C.QState_AddTransition2(ptr.Pointer(), PointerFromQObject(sender), C.CString(signal), PointerFromQAbstractState(target)))
	}
	return nil
}

func (ptr *QState) AddTransition(transition QAbstractTransition_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::addTransition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_AddTransition(ptr.Pointer(), PointerFromQAbstractTransition(transition))
	}
}

func (ptr *QState) AssignProperty(object QObject_ITF, name string, value QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::assignProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_AssignProperty(ptr.Pointer(), PointerFromQObject(object), C.CString(name), PointerFromQVariant(value))
	}
}

func (ptr *QState) ChildMode() QState__ChildMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::childMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QState__ChildMode(C.QState_ChildMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QState) ConnectChildModeChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::childModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_ConnectChildModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "childModeChanged", f)
	}
}

func (ptr *QState) DisconnectChildModeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::childModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_DisconnectChildModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "childModeChanged")
	}
}

//export callbackQStateChildModeChanged
func callbackQStateChildModeChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::childModeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "childModeChanged").(func())()
}

func (ptr *QState) ErrorState() *QAbstractState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::errorState")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QState_ErrorState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QState) ConnectErrorStateChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::errorStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_ConnectErrorStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "errorStateChanged", f)
	}
}

func (ptr *QState) DisconnectErrorStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::errorStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_DisconnectErrorStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "errorStateChanged")
	}
}

//export callbackQStateErrorStateChanged
func callbackQStateErrorStateChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::errorStateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "errorStateChanged").(func())()
}

func (ptr *QState) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QState) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQStateFinished
func callbackQStateFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QState) InitialState() *QAbstractState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::initialState")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QState_InitialState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QState) ConnectInitialStateChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::initialStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_ConnectInitialStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "initialStateChanged", f)
	}
}

func (ptr *QState) DisconnectInitialStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::initialStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_DisconnectInitialStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "initialStateChanged")
	}
}

//export callbackQStateInitialStateChanged
func callbackQStateInitialStateChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::initialStateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "initialStateChanged").(func())()
}

func (ptr *QState) ConnectPropertiesAssigned(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::propertiesAssigned")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_ConnectPropertiesAssigned(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "propertiesAssigned", f)
	}
}

func (ptr *QState) DisconnectPropertiesAssigned() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::propertiesAssigned")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_DisconnectPropertiesAssigned(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "propertiesAssigned")
	}
}

//export callbackQStatePropertiesAssigned
func callbackQStatePropertiesAssigned(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::propertiesAssigned")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "propertiesAssigned").(func())()
}

func (ptr *QState) RemoveTransition(transition QAbstractTransition_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::removeTransition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_RemoveTransition(ptr.Pointer(), PointerFromQAbstractTransition(transition))
	}
}

func (ptr *QState) SetChildMode(mode QState__ChildMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::setChildMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_SetChildMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QState) SetErrorState(state QAbstractState_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::setErrorState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_SetErrorState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QState) SetInitialState(state QAbstractState_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::setInitialState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_SetInitialState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QState) DestroyQState() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QState::~QState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QState_DestroyQState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
