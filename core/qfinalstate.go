package core

//#include "qfinalstate.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFinalState struct {
	QAbstractState
}

type QFinalStateITF interface {
	QAbstractStateITF
	QFinalStatePTR() *QFinalState
}

func PointerFromQFinalState(ptr QFinalStateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFinalStatePTR().Pointer()
	}
	return nil
}

func QFinalStateFromPointer(ptr unsafe.Pointer) *QFinalState {
	var n = new(QFinalState)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFinalState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFinalState) QFinalStatePTR() *QFinalState {
	return ptr
}

func NewQFinalState(parent QStateITF) *QFinalState {
	return QFinalStateFromPointer(unsafe.Pointer(C.QFinalState_NewQFinalState(C.QtObjectPtr(PointerFromQState(parent)))))
}

func (ptr *QFinalState) DestroyQFinalState() {
	if ptr.Pointer() != nil {
		C.QFinalState_DestroyQFinalState(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
