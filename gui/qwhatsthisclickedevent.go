package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QWhatsThisClickedEvent struct {
	core.QEvent
}

type QWhatsThisClickedEvent_ITF interface {
	core.QEvent_ITF
	QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent
}

func PointerFromQWhatsThisClickedEvent(ptr QWhatsThisClickedEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWhatsThisClickedEvent_PTR().Pointer()
	}
	return nil
}

func NewQWhatsThisClickedEventFromPointer(ptr unsafe.Pointer) *QWhatsThisClickedEvent {
	var n = new(QWhatsThisClickedEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWhatsThisClickedEvent) QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent {
	return ptr
}

func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWhatsThisClickedEvent::QWhatsThisClickedEvent")
		}
	}()

	return NewQWhatsThisClickedEventFromPointer(C.QWhatsThisClickedEvent_NewQWhatsThisClickedEvent(C.CString(href)))
}

func (ptr *QWhatsThisClickedEvent) Href() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWhatsThisClickedEvent::href")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWhatsThisClickedEvent_Href(ptr.Pointer()))
	}
	return ""
}
