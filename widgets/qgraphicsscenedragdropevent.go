package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGraphicsSceneDragDropEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneDragDropEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneDragDropEvent_PTR() *QGraphicsSceneDragDropEvent
}

func PointerFromQGraphicsSceneDragDropEvent(ptr QGraphicsSceneDragDropEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneDragDropEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneDragDropEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneDragDropEvent {
	var n = new(QGraphicsSceneDragDropEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneDragDropEvent) QGraphicsSceneDragDropEvent_PTR() *QGraphicsSceneDragDropEvent {
	return ptr
}

func (ptr *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::acceptProposedAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSceneDragDropEvent_AcceptProposedAction(ptr.Pointer())
	}
}

func (ptr *QGraphicsSceneDragDropEvent) Buttons() core.Qt__MouseButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::buttons")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneDragDropEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) DropAction() core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::dropAction")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QGraphicsSceneDragDropEvent_DropAction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) MimeData() *core.QMimeData {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::mimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QGraphicsSceneDragDropEvent_MimeData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneDragDropEvent) Modifiers() core.Qt__KeyboardModifier {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::modifiers")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneDragDropEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) PossibleActions() core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::possibleActions")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QGraphicsSceneDragDropEvent_PossibleActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) ProposedAction() core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::proposedAction")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QGraphicsSceneDragDropEvent_ProposedAction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneDragDropEvent) SetDropAction(action core.Qt__DropAction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::setDropAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSceneDragDropEvent_SetDropAction(ptr.Pointer(), C.int(action))
	}
}

func (ptr *QGraphicsSceneDragDropEvent) Source() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::source")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QGraphicsSceneDragDropEvent_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneDragDropEvent) DestroyQGraphicsSceneDragDropEvent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneDragDropEvent::~QGraphicsSceneDragDropEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(ptr.Pointer())
	}
}
