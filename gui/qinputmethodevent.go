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

type QInputMethodEventITF interface {
	core.QEventITF
	QInputMethodEventPTR() *QInputMethodEvent
}

func PointerFromQInputMethodEvent(ptr QInputMethodEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethodEventPTR().Pointer()
	}
	return nil
}

func QInputMethodEventFromPointer(ptr unsafe.Pointer) *QInputMethodEvent {
	var n = new(QInputMethodEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInputMethodEvent) QInputMethodEventPTR() *QInputMethodEvent {
	return ptr
}

//QInputMethodEvent::AttributeType
type QInputMethodEvent__AttributeType int

var (
	QInputMethodEvent__TextFormat = QInputMethodEvent__AttributeType(0)
	QInputMethodEvent__Cursor     = QInputMethodEvent__AttributeType(1)
	QInputMethodEvent__Language   = QInputMethodEvent__AttributeType(2)
	QInputMethodEvent__Ruby       = QInputMethodEvent__AttributeType(3)
	QInputMethodEvent__Selection  = QInputMethodEvent__AttributeType(4)
)

func NewQInputMethodEvent() *QInputMethodEvent {
	return QInputMethodEventFromPointer(unsafe.Pointer(C.QInputMethodEvent_NewQInputMethodEvent()))
}

func NewQInputMethodEvent3(other QInputMethodEventITF) *QInputMethodEvent {
	return QInputMethodEventFromPointer(unsafe.Pointer(C.QInputMethodEvent_NewQInputMethodEvent3(C.QtObjectPtr(PointerFromQInputMethodEvent(other)))))
}

func (ptr *QInputMethodEvent) CommitString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputMethodEvent_CommitString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QInputMethodEvent) PreeditString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputMethodEvent_PreeditString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QInputMethodEvent) ReplacementLength() int {
	if ptr.Pointer() != nil {
		return int(C.QInputMethodEvent_ReplacementLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputMethodEvent) ReplacementStart() int {
	if ptr.Pointer() != nil {
		return int(C.QInputMethodEvent_ReplacementStart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputMethodEvent) SetCommitString(commitString string, replaceFrom int, replaceLength int) {
	if ptr.Pointer() != nil {
		C.QInputMethodEvent_SetCommitString(C.QtObjectPtr(ptr.Pointer()), C.CString(commitString), C.int(replaceFrom), C.int(replaceLength))
	}
}
