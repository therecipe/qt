package widgets

//#include "qgraphicssceneevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneEvent struct {
	core.QEvent
}

type QGraphicsSceneEventITF interface {
	core.QEventITF
	QGraphicsSceneEventPTR() *QGraphicsSceneEvent
}

func PointerFromQGraphicsSceneEvent(ptr QGraphicsSceneEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneEvent {
	var n = new(QGraphicsSceneEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneEvent) QGraphicsSceneEventPTR() *QGraphicsSceneEvent {
	return ptr
}

func (ptr *QGraphicsSceneEvent) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QGraphicsSceneEvent_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsSceneEvent) DestroyQGraphicsSceneEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
