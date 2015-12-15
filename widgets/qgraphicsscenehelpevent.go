package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneHelpEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneHelpEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneHelpEvent_PTR() *QGraphicsSceneHelpEvent
}

func PointerFromQGraphicsSceneHelpEvent(ptr QGraphicsSceneHelpEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneHelpEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneHelpEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneHelpEvent {
	var n = new(QGraphicsSceneHelpEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneHelpEvent) QGraphicsSceneHelpEvent_PTR() *QGraphicsSceneHelpEvent {
	return ptr
}

func (ptr *QGraphicsSceneHelpEvent) ScreenPos() *core.QPoint {
	defer qt.Recovering("QGraphicsSceneHelpEvent::screenPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsSceneHelpEvent_ScreenPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneHelpEvent) DestroyQGraphicsSceneHelpEvent() {
	defer qt.Recovering("QGraphicsSceneHelpEvent::~QGraphicsSceneHelpEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSceneHelpEvent_DestroyQGraphicsSceneHelpEvent(ptr.Pointer())
	}
}
