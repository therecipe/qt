package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QShortcutEvent struct {
	core.QEvent
}

type QShortcutEvent_ITF interface {
	core.QEvent_ITF
	QShortcutEvent_PTR() *QShortcutEvent
}

func PointerFromQShortcutEvent(ptr QShortcutEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShortcutEvent_PTR().Pointer()
	}
	return nil
}

func NewQShortcutEventFromPointer(ptr unsafe.Pointer) *QShortcutEvent {
	var n = new(QShortcutEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QShortcutEvent) QShortcutEvent_PTR() *QShortcutEvent {
	return ptr
}

func NewQShortcutEvent(key QKeySequence_ITF, id int, ambiguous bool) *QShortcutEvent {
	defer qt.Recovering("QShortcutEvent::QShortcutEvent")

	return NewQShortcutEventFromPointer(C.QShortcutEvent_NewQShortcutEvent(PointerFromQKeySequence(key), C.int(id), C.int(qt.GoBoolToInt(ambiguous))))
}

func (ptr *QShortcutEvent) IsAmbiguous() bool {
	defer qt.Recovering("QShortcutEvent::isAmbiguous")

	if ptr.Pointer() != nil {
		return C.QShortcutEvent_IsAmbiguous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QShortcutEvent) ShortcutId() int {
	defer qt.Recovering("QShortcutEvent::shortcutId")

	if ptr.Pointer() != nil {
		return int(C.QShortcutEvent_ShortcutId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QShortcutEvent) DestroyQShortcutEvent() {
	defer qt.Recovering("QShortcutEvent::~QShortcutEvent")

	if ptr.Pointer() != nil {
		C.QShortcutEvent_DestroyQShortcutEvent(ptr.Pointer())
	}
}
