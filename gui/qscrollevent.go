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

type QScrollEventITF interface {
	core.QEventITF
	QScrollEventPTR() *QScrollEvent
}

func PointerFromQScrollEvent(ptr QScrollEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollEventPTR().Pointer()
	}
	return nil
}

func QScrollEventFromPointer(ptr unsafe.Pointer) *QScrollEvent {
	var n = new(QScrollEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScrollEvent) QScrollEventPTR() *QScrollEvent {
	return ptr
}

//QScrollEvent::ScrollState
type QScrollEvent__ScrollState int

var (
	QScrollEvent__ScrollStarted  = QScrollEvent__ScrollState(0)
	QScrollEvent__ScrollUpdated  = QScrollEvent__ScrollState(1)
	QScrollEvent__ScrollFinished = QScrollEvent__ScrollState(2)
)

func NewQScrollEvent(contentPos core.QPointFITF, overshootDistance core.QPointFITF, scrollState QScrollEvent__ScrollState) *QScrollEvent {
	return QScrollEventFromPointer(unsafe.Pointer(C.QScrollEvent_NewQScrollEvent(C.QtObjectPtr(core.PointerFromQPointF(contentPos)), C.QtObjectPtr(core.PointerFromQPointF(overshootDistance)), C.int(scrollState))))
}

func (ptr *QScrollEvent) ScrollState() QScrollEvent__ScrollState {
	if ptr.Pointer() != nil {
		return QScrollEvent__ScrollState(C.QScrollEvent_ScrollState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScrollEvent) DestroyQScrollEvent() {
	if ptr.Pointer() != nil {
		C.QScrollEvent_DestroyQScrollEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
