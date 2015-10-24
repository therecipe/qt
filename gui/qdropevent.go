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

type QDropEventITF interface {
	core.QEventITF
	QDropEventPTR() *QDropEvent
}

func PointerFromQDropEvent(ptr QDropEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDropEventPTR().Pointer()
	}
	return nil
}

func QDropEventFromPointer(ptr unsafe.Pointer) *QDropEvent {
	var n = new(QDropEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDropEvent) QDropEventPTR() *QDropEvent {
	return ptr
}

func (ptr *QDropEvent) SetDropAction(action core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QDropEvent_SetDropAction(C.QtObjectPtr(ptr.Pointer()), C.int(action))
	}
}

func NewQDropEvent(pos core.QPointFITF, actions core.Qt__DropAction, data core.QMimeDataITF, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, ty core.QEvent__Type) *QDropEvent {
	return QDropEventFromPointer(unsafe.Pointer(C.QDropEvent_NewQDropEvent(C.QtObjectPtr(core.PointerFromQPointF(pos)), C.int(actions), C.QtObjectPtr(core.PointerFromQMimeData(data)), C.int(buttons), C.int(modifiers), C.int(ty))))
}

func (ptr *QDropEvent) AcceptProposedAction() {
	if ptr.Pointer() != nil {
		C.QDropEvent_AcceptProposedAction(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDropEvent) DropAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDropEvent_DropAction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDropEvent) KeyboardModifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QDropEvent_KeyboardModifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDropEvent) MimeData() *core.QMimeData {
	if ptr.Pointer() != nil {
		return core.QMimeDataFromPointer(unsafe.Pointer(C.QDropEvent_MimeData(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDropEvent) MouseButtons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QDropEvent_MouseButtons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDropEvent) PossibleActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDropEvent_PossibleActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDropEvent) ProposedAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDropEvent_ProposedAction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDropEvent) Source() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QDropEvent_Source(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
