package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QHistoryState_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::QHistoryState")
		}
	}()

	return NewQHistoryStateFromPointer(C.QHistoryState_NewQHistoryState2(C.int(ty), PointerFromQState(parent)))
}

func NewQHistoryState(parent QState_ITF) *QHistoryState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::QHistoryState")
		}
	}()

	return NewQHistoryStateFromPointer(C.QHistoryState_NewQHistoryState(PointerFromQState(parent)))
}

func (ptr *QHistoryState) DefaultState() *QAbstractState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::defaultState")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QHistoryState_DefaultState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHistoryState) ConnectDefaultStateChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::defaultStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectDefaultStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "defaultStateChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectDefaultStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::defaultStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectDefaultStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "defaultStateChanged")
	}
}

//export callbackQHistoryStateDefaultStateChanged
func callbackQHistoryStateDefaultStateChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::defaultStateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "defaultStateChanged").(func())()
}

func (ptr *QHistoryState) HistoryType() QHistoryState__HistoryType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::historyType")
		}
	}()

	if ptr.Pointer() != nil {
		return QHistoryState__HistoryType(C.QHistoryState_HistoryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHistoryState) ConnectHistoryTypeChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::historyTypeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectHistoryTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "historyTypeChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectHistoryTypeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::historyTypeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectHistoryTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "historyTypeChanged")
	}
}

//export callbackQHistoryStateHistoryTypeChanged
func callbackQHistoryStateHistoryTypeChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::historyTypeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "historyTypeChanged").(func())()
}

func (ptr *QHistoryState) SetDefaultState(state QAbstractState_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::setDefaultState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_SetDefaultState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QHistoryState) SetHistoryType(ty QHistoryState__HistoryType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::setHistoryType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_SetHistoryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QHistoryState) DestroyQHistoryState() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHistoryState::~QHistoryState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHistoryState_DestroyQHistoryState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
