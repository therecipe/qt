package core

//#include "qhistorystate.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QHistoryState struct {
	QAbstractState
}

type QHistoryState_ITF interface {
	QAbstractState_ITF
	QHistoryState_PTR() *QHistoryState
}

func PointerFromQHistoryState(ptr QHistoryState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHistoryState_PTR().Pointer()
	}
	return nil
}

func NewQHistoryStateFromPointer(ptr unsafe.Pointer) *QHistoryState {
	var n = new(QHistoryState)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QHistoryState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHistoryState) QHistoryState_PTR() *QHistoryState {
	return ptr
}

//QHistoryState::HistoryType
type QHistoryState__HistoryType int64

const (
	QHistoryState__ShallowHistory = QHistoryState__HistoryType(0)
	QHistoryState__DeepHistory    = QHistoryState__HistoryType(1)
)

func NewQHistoryState2(ty QHistoryState__HistoryType, parent QState_ITF) *QHistoryState {
	return NewQHistoryStateFromPointer(C.QHistoryState_NewQHistoryState2(C.int(ty), PointerFromQState(parent)))
}

func NewQHistoryState(parent QState_ITF) *QHistoryState {
	return NewQHistoryStateFromPointer(C.QHistoryState_NewQHistoryState(PointerFromQState(parent)))
}

func (ptr *QHistoryState) DefaultState() *QAbstractState {
	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QHistoryState_DefaultState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHistoryState) ConnectDefaultStateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectDefaultStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "defaultStateChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectDefaultStateChanged() {
	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectDefaultStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "defaultStateChanged")
	}
}

//export callbackQHistoryStateDefaultStateChanged
func callbackQHistoryStateDefaultStateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "defaultStateChanged").(func())()
}

func (ptr *QHistoryState) HistoryType() QHistoryState__HistoryType {
	if ptr.Pointer() != nil {
		return QHistoryState__HistoryType(C.QHistoryState_HistoryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHistoryState) ConnectHistoryTypeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectHistoryTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "historyTypeChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectHistoryTypeChanged() {
	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectHistoryTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "historyTypeChanged")
	}
}

//export callbackQHistoryStateHistoryTypeChanged
func callbackQHistoryStateHistoryTypeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "historyTypeChanged").(func())()
}

func (ptr *QHistoryState) SetDefaultState(state QAbstractState_ITF) {
	if ptr.Pointer() != nil {
		C.QHistoryState_SetDefaultState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QHistoryState) SetHistoryType(ty QHistoryState__HistoryType) {
	if ptr.Pointer() != nil {
		C.QHistoryState_SetHistoryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QHistoryState) DestroyQHistoryState() {
	if ptr.Pointer() != nil {
		C.QHistoryState_DestroyQHistoryState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
