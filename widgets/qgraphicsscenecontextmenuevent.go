package widgets

//#include "qgraphicsscenecontextmenuevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneContextMenuEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneContextMenuEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneContextMenuEventPTR() *QGraphicsSceneContextMenuEvent
}

func PointerFromQGraphicsSceneContextMenuEvent(ptr QGraphicsSceneContextMenuEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneContextMenuEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneContextMenuEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneContextMenuEvent {
	var n = new(QGraphicsSceneContextMenuEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneContextMenuEvent) QGraphicsSceneContextMenuEventPTR() *QGraphicsSceneContextMenuEvent {
	return ptr
}

//QGraphicsSceneContextMenuEvent::Reason
type QGraphicsSceneContextMenuEvent__Reason int

var (
	QGraphicsSceneContextMenuEvent__Mouse    = QGraphicsSceneContextMenuEvent__Reason(0)
	QGraphicsSceneContextMenuEvent__Keyboard = QGraphicsSceneContextMenuEvent__Reason(1)
	QGraphicsSceneContextMenuEvent__Other    = QGraphicsSceneContextMenuEvent__Reason(2)
)

func (ptr *QGraphicsSceneContextMenuEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneContextMenuEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneContextMenuEvent) Reason() QGraphicsSceneContextMenuEvent__Reason {
	if ptr.Pointer() != nil {
		return QGraphicsSceneContextMenuEvent__Reason(C.QGraphicsSceneContextMenuEvent_Reason(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneContextMenuEvent) DestroyQGraphicsSceneContextMenuEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
