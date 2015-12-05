package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QHelpEvent struct {
	core.QEvent
}

type QHelpEvent_ITF interface {
	core.QEvent_ITF
	QHelpEvent_PTR() *QHelpEvent
}

func PointerFromQHelpEvent(ptr QHelpEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEvent_PTR().Pointer()
	}
	return nil
}

func NewQHelpEventFromPointer(ptr unsafe.Pointer) *QHelpEvent {
	var n = new(QHelpEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpEvent) QHelpEvent_PTR() *QHelpEvent {
	return ptr
}

func NewQHelpEvent(ty core.QEvent__Type, pos core.QPoint_ITF, globalPos core.QPoint_ITF) *QHelpEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpEvent::QHelpEvent")
		}
	}()

	return NewQHelpEventFromPointer(C.QHelpEvent_NewQHelpEvent(C.int(ty), core.PointerFromQPoint(pos), core.PointerFromQPoint(globalPos)))
}

func (ptr *QHelpEvent) GlobalX() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpEvent::globalX")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpEvent) GlobalY() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpEvent::globalY")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpEvent) X() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpEvent::x")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpEvent) Y() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpEvent::y")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_Y(ptr.Pointer()))
	}
	return 0
}
