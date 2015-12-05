package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPaintEvent struct {
	core.QEvent
}

type QPaintEvent_ITF interface {
	core.QEvent_ITF
	QPaintEvent_PTR() *QPaintEvent
}

func PointerFromQPaintEvent(ptr QPaintEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEvent_PTR().Pointer()
	}
	return nil
}

func NewQPaintEventFromPointer(ptr unsafe.Pointer) *QPaintEvent {
	var n = new(QPaintEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintEvent) QPaintEvent_PTR() *QPaintEvent {
	return ptr
}

func NewQPaintEvent2(paintRect core.QRect_ITF) *QPaintEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEvent::QPaintEvent")
		}
	}()

	return NewQPaintEventFromPointer(C.QPaintEvent_NewQPaintEvent2(core.PointerFromQRect(paintRect)))
}

func NewQPaintEvent(paintRegion QRegion_ITF) *QPaintEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEvent::QPaintEvent")
		}
	}()

	return NewQPaintEventFromPointer(C.QPaintEvent_NewQPaintEvent(PointerFromQRegion(paintRegion)))
}

func (ptr *QPaintEvent) Region() *QRegion {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPaintEvent::region")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QPaintEvent_Region(ptr.Pointer()))
	}
	return nil
}
