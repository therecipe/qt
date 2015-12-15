package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleEvent struct {
	ptr unsafe.Pointer
}

type QAccessibleEvent_ITF interface {
	QAccessibleEvent_PTR() *QAccessibleEvent
}

func (p *QAccessibleEvent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleEvent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleEvent(ptr QAccessibleEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleEventFromPointer(ptr unsafe.Pointer) *QAccessibleEvent {
	var n = new(QAccessibleEvent)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAccessibleEvent_") {
		n.SetObjectNameAbs("QAccessibleEvent_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccessibleEvent) QAccessibleEvent_PTR() *QAccessibleEvent {
	return ptr
}

func NewQAccessibleEvent2(interfa QAccessibleInterface_ITF, ty QAccessible__Event) *QAccessibleEvent {
	defer qt.Recovering("QAccessibleEvent::QAccessibleEvent")

	return NewQAccessibleEventFromPointer(C.QAccessibleEvent_NewQAccessibleEvent2(PointerFromQAccessibleInterface(interfa), C.int(ty)))
}

func NewQAccessibleEvent(object core.QObject_ITF, ty QAccessible__Event) *QAccessibleEvent {
	defer qt.Recovering("QAccessibleEvent::QAccessibleEvent")

	return NewQAccessibleEventFromPointer(C.QAccessibleEvent_NewQAccessibleEvent(core.PointerFromQObject(object), C.int(ty)))
}

func (ptr *QAccessibleEvent) AccessibleInterface() *QAccessibleInterface {
	defer qt.Recovering("QAccessibleEvent::accessibleInterface")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleEvent_AccessibleInterface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleEvent) Child() int {
	defer qt.Recovering("QAccessibleEvent::child")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleEvent_Child(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleEvent) Object() *core.QObject {
	defer qt.Recovering("QAccessibleEvent::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QAccessibleEvent_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleEvent) SetChild(child int) {
	defer qt.Recovering("QAccessibleEvent::setChild")

	if ptr.Pointer() != nil {
		C.QAccessibleEvent_SetChild(ptr.Pointer(), C.int(child))
	}
}

func (ptr *QAccessibleEvent) Type() QAccessible__Event {
	defer qt.Recovering("QAccessibleEvent::type")

	if ptr.Pointer() != nil {
		return QAccessible__Event(C.QAccessibleEvent_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleEvent) DestroyQAccessibleEvent() {
	defer qt.Recovering("QAccessibleEvent::~QAccessibleEvent")

	if ptr.Pointer() != nil {
		C.QAccessibleEvent_DestroyQAccessibleEvent(ptr.Pointer())
	}
}

func (ptr *QAccessibleEvent) ObjectNameAbs() string {
	defer qt.Recovering("QAccessibleEvent::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleEvent_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleEvent) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAccessibleEvent::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAccessibleEvent_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
