package gui

//#include "qdropevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDropEvent struct {
	core.QEvent
}

type QDropEvent_ITF interface {
	core.QEvent_ITF
	QDropEvent_PTR() *QDropEvent
}

func PointerFromQDropEvent(ptr QDropEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDropEvent_PTR().Pointer()
	}
	return nil
}

func NewQDropEventFromPointer(ptr unsafe.Pointer) *QDropEvent {
	var n = new(QDropEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDropEvent) QDropEvent_PTR() *QDropEvent {
	return ptr
}

func (ptr *QDropEvent) SetDropAction(action core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QDropEvent_SetDropAction(ptr.Pointer(), C.int(action))
	}
}

func NewQDropEvent(pos core.QPointF_ITF, actions core.Qt__DropAction, data core.QMimeData_ITF, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, ty core.QEvent__Type) *QDropEvent {
	return NewQDropEventFromPointer(C.QDropEvent_NewQDropEvent(core.PointerFromQPointF(pos), C.int(actions), core.PointerFromQMimeData(data), C.int(buttons), C.int(modifiers), C.int(ty)))
}

func (ptr *QDropEvent) AcceptProposedAction() {
	if ptr.Pointer() != nil {
		C.QDropEvent_AcceptProposedAction(ptr.Pointer())
	}
}

func (ptr *QDropEvent) DropAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDropEvent_DropAction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDropEvent) KeyboardModifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QDropEvent_KeyboardModifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDropEvent) MimeData() *core.QMimeData {
	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QDropEvent_MimeData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDropEvent) MouseButtons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QDropEvent_MouseButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDropEvent) PossibleActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDropEvent_PossibleActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDropEvent) ProposedAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDropEvent_ProposedAction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDropEvent) Source() *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDropEvent_Source(ptr.Pointer()))
	}
	return nil
}
