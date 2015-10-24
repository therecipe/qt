package widgets

//#include "qgraphicsscenedragdropevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneDragDropEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneDragDropEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneDragDropEventPTR() *QGraphicsSceneDragDropEvent
}

func PointerFromQGraphicsSceneDragDropEvent(ptr QGraphicsSceneDragDropEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneDragDropEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneDragDropEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneDragDropEvent {
	var n = new(QGraphicsSceneDragDropEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneDragDropEvent) QGraphicsSceneDragDropEventPTR() *QGraphicsSceneDragDropEvent {
	return ptr
}

func (ptr *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneDragDropEvent_AcceptProposedAction(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsSceneDragDropEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneDragDropEvent_Buttons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) DropAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QGraphicsSceneDragDropEvent_DropAction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) MimeData() *core.QMimeData {
	if ptr.Pointer() != nil {
		return core.QMimeDataFromPointer(unsafe.Pointer(C.QGraphicsSceneDragDropEvent_MimeData(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsSceneDragDropEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneDragDropEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) PossibleActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QGraphicsSceneDragDropEvent_PossibleActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) ProposedAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QGraphicsSceneDragDropEvent_ProposedAction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) SetDropAction(action core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneDragDropEvent_SetDropAction(C.QtObjectPtr(ptr.Pointer()), C.int(action))
	}
}

func (ptr *QGraphicsSceneDragDropEvent) Source() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QGraphicsSceneDragDropEvent_Source(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsSceneDragDropEvent) DestroyQGraphicsSceneDragDropEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
