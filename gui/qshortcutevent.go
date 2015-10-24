package gui

//#include "qshortcutevent.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QShortcutEvent struct {
	core.QEvent
}

type QShortcutEventITF interface {
	core.QEventITF
	QShortcutEventPTR() *QShortcutEvent
}

func PointerFromQShortcutEvent(ptr QShortcutEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShortcutEventPTR().Pointer()
	}
	return nil
}

func QShortcutEventFromPointer(ptr unsafe.Pointer) *QShortcutEvent {
	var n = new(QShortcutEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QShortcutEvent) QShortcutEventPTR() *QShortcutEvent {
	return ptr
}

func NewQShortcutEvent(key QKeySequenceITF, id int, ambiguous bool) *QShortcutEvent {
	return QShortcutEventFromPointer(unsafe.Pointer(C.QShortcutEvent_NewQShortcutEvent(C.QtObjectPtr(PointerFromQKeySequence(key)), C.int(id), C.int(qt.GoBoolToInt(ambiguous)))))
}

func (ptr *QShortcutEvent) IsAmbiguous() bool {
	if ptr.Pointer() != nil {
		return C.QShortcutEvent_IsAmbiguous(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QShortcutEvent) ShortcutId() int {
	if ptr.Pointer() != nil {
		return int(C.QShortcutEvent_ShortcutId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QShortcutEvent) DestroyQShortcutEvent() {
	if ptr.Pointer() != nil {
		C.QShortcutEvent_DestroyQShortcutEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
