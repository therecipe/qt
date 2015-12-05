package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QResizeEvent struct {
	core.QEvent
}

type QResizeEvent_ITF interface {
	core.QEvent_ITF
	QResizeEvent_PTR() *QResizeEvent
}

func PointerFromQResizeEvent(ptr QResizeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QResizeEvent_PTR().Pointer()
	}
	return nil
}

func NewQResizeEventFromPointer(ptr unsafe.Pointer) *QResizeEvent {
	var n = new(QResizeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QResizeEvent) QResizeEvent_PTR() *QResizeEvent {
	return ptr
}

func NewQResizeEvent(size core.QSize_ITF, oldSize core.QSize_ITF) *QResizeEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResizeEvent::QResizeEvent")
		}
	}()

	return NewQResizeEventFromPointer(C.QResizeEvent_NewQResizeEvent(core.PointerFromQSize(size), core.PointerFromQSize(oldSize)))
}
