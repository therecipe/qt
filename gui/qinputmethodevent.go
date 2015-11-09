package gui

//#include "qinputmethodevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QInputMethodEvent struct {
	core.QEvent
}

type QInputMethodEvent_ITF interface {
	core.QEvent_ITF
	QInputMethodEvent_PTR() *QInputMethodEvent
}

func PointerFromQInputMethodEvent(ptr QInputMethodEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethodEvent_PTR().Pointer()
	}
	return nil
}

func NewQInputMethodEventFromPointer(ptr unsafe.Pointer) *QInputMethodEvent {
	var n = new(QInputMethodEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInputMethodEvent) QInputMethodEvent_PTR() *QInputMethodEvent {
	return ptr
}

//QInputMethodEvent::AttributeType
type QInputMethodEvent__AttributeType int64

const (
	QInputMethodEvent__TextFormat = QInputMethodEvent__AttributeType(0)
	QInputMethodEvent__Cursor     = QInputMethodEvent__AttributeType(1)
	QInputMethodEvent__Language   = QInputMethodEvent__AttributeType(2)
	QInputMethodEvent__Ruby       = QInputMethodEvent__AttributeType(3)
	QInputMethodEvent__Selection  = QInputMethodEvent__AttributeType(4)
)

func NewQInputMethodEvent() *QInputMethodEvent {
	return NewQInputMethodEventFromPointer(C.QInputMethodEvent_NewQInputMethodEvent())
}

func NewQInputMethodEvent3(other QInputMethodEvent_ITF) *QInputMethodEvent {
	return NewQInputMethodEventFromPointer(C.QInputMethodEvent_NewQInputMethodEvent3(PointerFromQInputMethodEvent(other)))
}

func (ptr *QInputMethodEvent) CommitString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputMethodEvent_CommitString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInputMethodEvent) PreeditString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputMethodEvent_PreeditString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QInputMethodEvent) ReplacementLength() int {
	if ptr.Pointer() != nil {
		return int(C.QInputMethodEvent_ReplacementLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputMethodEvent) ReplacementStart() int {
	if ptr.Pointer() != nil {
		return int(C.QInputMethodEvent_ReplacementStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputMethodEvent) SetCommitString(commitString string, replaceFrom int, replaceLength int) {
	if ptr.Pointer() != nil {
		C.QInputMethodEvent_SetCommitString(ptr.Pointer(), C.CString(commitString), C.int(replaceFrom), C.int(replaceLength))
	}
}
