package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneEvent struct {
	core.QEvent
}

type QGraphicsSceneEvent_ITF interface {
	core.QEvent_ITF
	QGraphicsSceneEvent_PTR() *QGraphicsSceneEvent
}

func PointerFromQGraphicsSceneEvent(ptr QGraphicsSceneEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneEvent {
	var n = new(QGraphicsSceneEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneEvent) QGraphicsSceneEvent_PTR() *QGraphicsSceneEvent {
	return ptr
}

func (ptr *QGraphicsSceneEvent) Widget() *QWidget {
	defer qt.Recovering("QGraphicsSceneEvent::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QGraphicsSceneEvent_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneEvent) DestroyQGraphicsSceneEvent() {
	defer qt.Recovering("QGraphicsSceneEvent::~QGraphicsSceneEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(ptr.Pointer())
	}
}
