package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QContextMenuEvent struct {
	QInputEvent
}

type QContextMenuEvent_ITF interface {
	QInputEvent_ITF
	QContextMenuEvent_PTR() *QContextMenuEvent
}

func PointerFromQContextMenuEvent(ptr QContextMenuEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QContextMenuEvent_PTR().Pointer()
	}
	return nil
}

func NewQContextMenuEventFromPointer(ptr unsafe.Pointer) *QContextMenuEvent {
	var n = new(QContextMenuEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QContextMenuEvent) QContextMenuEvent_PTR() *QContextMenuEvent {
	return ptr
}

//QContextMenuEvent::Reason
type QContextMenuEvent__Reason int64

const (
	QContextMenuEvent__Mouse    = QContextMenuEvent__Reason(0)
	QContextMenuEvent__Keyboard = QContextMenuEvent__Reason(1)
	QContextMenuEvent__Other    = QContextMenuEvent__Reason(2)
)

func NewQContextMenuEvent3(reason QContextMenuEvent__Reason, pos core.QPoint_ITF) *QContextMenuEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::QContextMenuEvent")
		}
	}()

	return NewQContextMenuEventFromPointer(C.QContextMenuEvent_NewQContextMenuEvent3(C.int(reason), core.PointerFromQPoint(pos)))
}

func NewQContextMenuEvent2(reason QContextMenuEvent__Reason, pos core.QPoint_ITF, globalPos core.QPoint_ITF) *QContextMenuEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::QContextMenuEvent")
		}
	}()

	return NewQContextMenuEventFromPointer(C.QContextMenuEvent_NewQContextMenuEvent2(C.int(reason), core.PointerFromQPoint(pos), core.PointerFromQPoint(globalPos)))
}

func NewQContextMenuEvent(reason QContextMenuEvent__Reason, pos core.QPoint_ITF, globalPos core.QPoint_ITF, modifiers core.Qt__KeyboardModifier) *QContextMenuEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::QContextMenuEvent")
		}
	}()

	return NewQContextMenuEventFromPointer(C.QContextMenuEvent_NewQContextMenuEvent(C.int(reason), core.PointerFromQPoint(pos), core.PointerFromQPoint(globalPos), C.int(modifiers)))
}

func (ptr *QContextMenuEvent) GlobalX() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::globalX")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QContextMenuEvent) GlobalY() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::globalY")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QContextMenuEvent) Reason() QContextMenuEvent__Reason {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::reason")
		}
	}()

	if ptr.Pointer() != nil {
		return QContextMenuEvent__Reason(C.QContextMenuEvent_Reason(ptr.Pointer()))
	}
	return 0
}

func (ptr *QContextMenuEvent) X() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::x")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QContextMenuEvent) Y() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QContextMenuEvent::y")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_Y(ptr.Pointer()))
	}
	return 0
}
