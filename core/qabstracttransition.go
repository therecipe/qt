package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QAbstractTransition struct {
	QObject
}

type QAbstractTransition_ITF interface {
	QObject_ITF
	QAbstractTransition_PTR() *QAbstractTransition
}

func PointerFromQAbstractTransition(ptr QAbstractTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTransition_PTR().Pointer()
	}
	return nil
}

func NewQAbstractTransitionFromPointer(ptr unsafe.Pointer) *QAbstractTransition {
	var n = new(QAbstractTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractTransition_") {
		n.SetObjectName("QAbstractTransition_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractTransition) QAbstractTransition_PTR() *QAbstractTransition {
	return ptr
}

//QAbstractTransition::TransitionType
type QAbstractTransition__TransitionType int64

const (
	QAbstractTransition__ExternalTransition = QAbstractTransition__TransitionType(0)
	QAbstractTransition__InternalTransition = QAbstractTransition__TransitionType(1)
)

func (ptr *QAbstractTransition) AddAnimation(animation QAbstractAnimation_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::addAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_AddAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QAbstractTransition) Machine() *QStateMachine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::machine")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStateMachineFromPointer(C.QAbstractTransition_Machine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractTransition) RemoveAnimation(animation QAbstractAnimation_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::removeAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_RemoveAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QAbstractTransition) SetTargetState(target QAbstractState_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::setTargetState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_SetTargetState(ptr.Pointer(), PointerFromQAbstractState(target))
	}
}

func (ptr *QAbstractTransition) SetTransitionType(ty QAbstractTransition__TransitionType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::setTransitionType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_SetTransitionType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QAbstractTransition) SourceState() *QState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::sourceState")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStateFromPointer(C.QAbstractTransition_SourceState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractTransition) TargetState() *QAbstractState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetState")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QAbstractTransition_TargetState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractTransition) ConnectTargetStateChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_ConnectTargetStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetStateChanged", f)
	}
}

func (ptr *QAbstractTransition) DisconnectTargetStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_DisconnectTargetStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetStateChanged")
	}
}

//export callbackQAbstractTransitionTargetStateChanged
func callbackQAbstractTransitionTargetStateChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetStateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "targetStateChanged").(func())()
}

func (ptr *QAbstractTransition) ConnectTargetStatesChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetStatesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_ConnectTargetStatesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetStatesChanged", f)
	}
}

func (ptr *QAbstractTransition) DisconnectTargetStatesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetStatesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_DisconnectTargetStatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetStatesChanged")
	}
}

//export callbackQAbstractTransitionTargetStatesChanged
func callbackQAbstractTransitionTargetStatesChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::targetStatesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "targetStatesChanged").(func())()
}

func (ptr *QAbstractTransition) TransitionType() QAbstractTransition__TransitionType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::transitionType")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractTransition__TransitionType(C.QAbstractTransition_TransitionType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractTransition) ConnectTriggered(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::triggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QAbstractTransition) DisconnectTriggered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::triggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQAbstractTransitionTriggered
func callbackQAbstractTransitionTriggered(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::triggered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "triggered").(func())()
}

func (ptr *QAbstractTransition) DestroyQAbstractTransition() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractTransition::~QAbstractTransition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractTransition_DestroyQAbstractTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
