package gui

//#include "qscrollevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScrollEvent struct {
	core.QEvent
}

type QScrollEvent_ITF interface {
	core.QEvent_ITF
	QScrollEvent_PTR() *QScrollEvent
}

func PointerFromQScrollEvent(ptr QScrollEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollEvent_PTR().Pointer()
	}
	return nil
}

func NewQScrollEventFromPointer(ptr unsafe.Pointer) *QScrollEvent {
	var n = new(QScrollEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScrollEvent) QScrollEvent_PTR() *QScrollEvent {
	return ptr
}

//QScrollEvent::ScrollState
type QScrollEvent__ScrollState int64

const (
	QScrollEvent__ScrollStarted  = QScrollEvent__ScrollState(0)
	QScrollEvent__ScrollUpdated  = QScrollEvent__ScrollState(1)
	QScrollEvent__ScrollFinished = QScrollEvent__ScrollState(2)
)

func NewQScrollEvent(contentPos core.QPointF_ITF, overshootDistance core.QPointF_ITF, scrollState QScrollEvent__ScrollState) *QScrollEvent {
	return NewQScrollEventFromPointer(C.QScrollEvent_NewQScrollEvent(core.PointerFromQPointF(contentPos), core.PointerFromQPointF(overshootDistance), C.int(scrollState)))
}

func (ptr *QScrollEvent) ScrollState() QScrollEvent__ScrollState {
	if ptr.Pointer() != nil {
		return QScrollEvent__ScrollState(C.QScrollEvent_ScrollState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScrollEvent) DestroyQScrollEvent() {
	if ptr.Pointer() != nil {
		C.QScrollEvent_DestroyQScrollEvent(ptr.Pointer())
	}
}
