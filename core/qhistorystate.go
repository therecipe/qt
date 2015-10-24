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

type QHistoryStateITF interface {
	QAbstractStateITF
	QHistoryStatePTR() *QHistoryState
}

func PointerFromQHistoryState(ptr QHistoryStateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHistoryStatePTR().Pointer()
	}
	return nil
}

func QHistoryStateFromPointer(ptr unsafe.Pointer) *QHistoryState {
	var n = new(QHistoryState)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHistoryState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHistoryState) QHistoryStatePTR() *QHistoryState {
	return ptr
}

//QHistoryState::HistoryType
type QHistoryState__HistoryType int

var (
	QHistoryState__ShallowHistory = QHistoryState__HistoryType(0)
	QHistoryState__DeepHistory    = QHistoryState__HistoryType(1)
)

func NewQHistoryState2(ty QHistoryState__HistoryType, parent QStateITF) *QHistoryState {
	return QHistoryStateFromPointer(unsafe.Pointer(C.QHistoryState_NewQHistoryState2(C.int(ty), C.QtObjectPtr(PointerFromQState(parent)))))
}

func NewQHistoryState(parent QStateITF) *QHistoryState {
	return QHistoryStateFromPointer(unsafe.Pointer(C.QHistoryState_NewQHistoryState(C.QtObjectPtr(PointerFromQState(parent)))))
}

func (ptr *QHistoryState) DefaultState() *QAbstractState {
	if ptr.Pointer() != nil {
		return QAbstractStateFromPointer(unsafe.Pointer(C.QHistoryState_DefaultState(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHistoryState) ConnectDefaultStateChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectDefaultStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "defaultStateChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectDefaultStateChanged() {
	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectDefaultStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "defaultStateChanged")
	}
}

//export callbackQHistoryStateDefaultStateChanged
func callbackQHistoryStateDefaultStateChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "defaultStateChanged").(func())()
}

func (ptr *QHistoryState) HistoryType() QHistoryState__HistoryType {
	if ptr.Pointer() != nil {
		return QHistoryState__HistoryType(C.QHistoryState_HistoryType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHistoryState) ConnectHistoryTypeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectHistoryTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "historyTypeChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectHistoryTypeChanged() {
	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectHistoryTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "historyTypeChanged")
	}
}

//export callbackQHistoryStateHistoryTypeChanged
func callbackQHistoryStateHistoryTypeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "historyTypeChanged").(func())()
}

func (ptr *QHistoryState) SetDefaultState(state QAbstractStateITF) {
	if ptr.Pointer() != nil {
		C.QHistoryState_SetDefaultState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(state)))
	}
}

func (ptr *QHistoryState) SetHistoryType(ty QHistoryState__HistoryType) {
	if ptr.Pointer() != nil {
		C.QHistoryState_SetHistoryType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QHistoryState) DestroyQHistoryState() {
	if ptr.Pointer() != nil {
		C.QHistoryState_DestroyQHistoryState(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
