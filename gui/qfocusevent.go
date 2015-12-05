package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QFocusEvent struct {
	core.QEvent
}

type QFocusEvent_ITF interface {
	core.QEvent_ITF
	QFocusEvent_PTR() *QFocusEvent
}

func PointerFromQFocusEvent(ptr QFocusEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusEvent_PTR().Pointer()
	}
	return nil
}

func NewQFocusEventFromPointer(ptr unsafe.Pointer) *QFocusEvent {
	var n = new(QFocusEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFocusEvent) QFocusEvent_PTR() *QFocusEvent {
	return ptr
}

func NewQFocusEvent(ty core.QEvent__Type, reason core.Qt__FocusReason) *QFocusEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusEvent::QFocusEvent")
		}
	}()

	return NewQFocusEventFromPointer(C.QFocusEvent_NewQFocusEvent(C.int(ty), C.int(reason)))
}

func (ptr *QFocusEvent) GotFocus() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusEvent::gotFocus")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFocusEvent_GotFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFocusEvent) LostFocus() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusEvent::lostFocus")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFocusEvent_LostFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFocusEvent) Reason() core.Qt__FocusReason {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusEvent::reason")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__FocusReason(C.QFocusEvent_Reason(ptr.Pointer()))
	}
	return 0
}
