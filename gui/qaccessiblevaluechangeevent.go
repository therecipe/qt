package gui

//#include "qaccessiblevaluechangeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleValueChangeEvent struct {
	QAccessibleEvent
}

type QAccessibleValueChangeEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleValueChangeEvent_PTR() *QAccessibleValueChangeEvent
}

func PointerFromQAccessibleValueChangeEvent(ptr QAccessibleValueChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleValueChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleValueChangeEventFromPointer(ptr unsafe.Pointer) *QAccessibleValueChangeEvent {
	var n = new(QAccessibleValueChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleValueChangeEvent) QAccessibleValueChangeEvent_PTR() *QAccessibleValueChangeEvent {
	return ptr
}

func NewQAccessibleValueChangeEvent2(iface QAccessibleInterface_ITF, val core.QVariant_ITF) *QAccessibleValueChangeEvent {
	return NewQAccessibleValueChangeEventFromPointer(C.QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent2(PointerFromQAccessibleInterface(iface), core.PointerFromQVariant(val)))
}

func NewQAccessibleValueChangeEvent(object core.QObject_ITF, value core.QVariant_ITF) *QAccessibleValueChangeEvent {
	return NewQAccessibleValueChangeEventFromPointer(C.QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent(core.PointerFromQObject(object), core.PointerFromQVariant(value)))
}

func (ptr *QAccessibleValueChangeEvent) SetValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAccessibleValueChangeEvent_SetValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QAccessibleValueChangeEvent) Value() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAccessibleValueChangeEvent_Value(ptr.Pointer()))
	}
	return nil
}
