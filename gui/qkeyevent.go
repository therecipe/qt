package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QKeyEvent struct {
	QInputEvent
}

type QKeyEvent_ITF interface {
	QInputEvent_ITF
	QKeyEvent_PTR() *QKeyEvent
}

func PointerFromQKeyEvent(ptr QKeyEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEvent_PTR().Pointer()
	}
	return nil
}

func NewQKeyEventFromPointer(ptr unsafe.Pointer) *QKeyEvent {
	var n = new(QKeyEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QKeyEvent) QKeyEvent_PTR() *QKeyEvent {
	return ptr
}

func (ptr *QKeyEvent) Matches(key QKeySequence__StandardKey) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEvent::matches")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QKeyEvent_Matches(ptr.Pointer(), C.int(key)) != 0
	}
	return false
}

func (ptr *QKeyEvent) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEvent::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QKeyEvent_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEvent) IsAutoRepeat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEvent::isAutoRepeat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QKeyEvent_IsAutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QKeyEvent) Key() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEvent::key")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QKeyEvent_Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEvent) Modifiers() core.Qt__KeyboardModifier {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEvent::modifiers")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QKeyEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEvent) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QKeyEvent::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QKeyEvent_Text(ptr.Pointer()))
	}
	return ""
}
