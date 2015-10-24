package gui

//#include "qaccessibleevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleEvent struct {
	ptr unsafe.Pointer
}

type QAccessibleEventITF interface {
	QAccessibleEventPTR() *QAccessibleEvent
}

func (p *QAccessibleEvent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleEvent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleEvent(ptr QAccessibleEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleEventPTR().Pointer()
	}
	return nil
}

func QAccessibleEventFromPointer(ptr unsafe.Pointer) *QAccessibleEvent {
	var n = new(QAccessibleEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleEvent) QAccessibleEventPTR() *QAccessibleEvent {
	return ptr
}

func NewQAccessibleEvent2(interfa QAccessibleInterfaceITF, ty QAccessible__Event) *QAccessibleEvent {
	return QAccessibleEventFromPointer(unsafe.Pointer(C.QAccessibleEvent_NewQAccessibleEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(interfa)), C.int(ty))))
}

func NewQAccessibleEvent(object core.QObjectITF, ty QAccessible__Event) *QAccessibleEvent {
	return QAccessibleEventFromPointer(unsafe.Pointer(C.QAccessibleEvent_NewQAccessibleEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(ty))))
}

func (ptr *QAccessibleEvent) AccessibleInterface() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleEvent_AccessibleInterface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleEvent) Child() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleEvent_Child(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleEvent) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QAccessibleEvent_Object(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleEvent) SetChild(child int) {
	if ptr.Pointer() != nil {
		C.QAccessibleEvent_SetChild(C.QtObjectPtr(ptr.Pointer()), C.int(child))
	}
}

func (ptr *QAccessibleEvent) Type() QAccessible__Event {
	if ptr.Pointer() != nil {
		return QAccessible__Event(C.QAccessibleEvent_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleEvent) DestroyQAccessibleEvent() {
	if ptr.Pointer() != nil {
		C.QAccessibleEvent_DestroyQAccessibleEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
