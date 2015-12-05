package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QFinalState struct {
	QAbstractState
}

type QFinalState_ITF interface {
	QAbstractState_ITF
	QFinalState_PTR() *QFinalState
}

func PointerFromQFinalState(ptr QFinalState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFinalState_PTR().Pointer()
	}
	return nil
}

func NewQFinalStateFromPointer(ptr unsafe.Pointer) *QFinalState {
	var n = new(QFinalState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFinalState_") {
		n.SetObjectName("QFinalState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFinalState) QFinalState_PTR() *QFinalState {
	return ptr
}

func NewQFinalState(parent QState_ITF) *QFinalState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFinalState::QFinalState")
		}
	}()

	return NewQFinalStateFromPointer(C.QFinalState_NewQFinalState(PointerFromQState(parent)))
}

func (ptr *QFinalState) DestroyQFinalState() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFinalState::~QFinalState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFinalState_DestroyQFinalState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
