package gui

//#include "qinputmethodqueryevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QInputMethodQueryEvent struct {
	core.QEvent
}

type QInputMethodQueryEvent_ITF interface {
	core.QEvent_ITF
	QInputMethodQueryEvent_PTR() *QInputMethodQueryEvent
}

func PointerFromQInputMethodQueryEvent(ptr QInputMethodQueryEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethodQueryEvent_PTR().Pointer()
	}
	return nil
}

func NewQInputMethodQueryEventFromPointer(ptr unsafe.Pointer) *QInputMethodQueryEvent {
	var n = new(QInputMethodQueryEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInputMethodQueryEvent) QInputMethodQueryEvent_PTR() *QInputMethodQueryEvent {
	return ptr
}

func NewQInputMethodQueryEvent(queries core.Qt__InputMethodQuery) *QInputMethodQueryEvent {
	return NewQInputMethodQueryEventFromPointer(C.QInputMethodQueryEvent_NewQInputMethodQueryEvent(C.int(queries)))
}

func (ptr *QInputMethodQueryEvent) Queries() core.Qt__InputMethodQuery {
	if ptr.Pointer() != nil {
		return core.Qt__InputMethodQuery(C.QInputMethodQueryEvent_Queries(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputMethodQueryEvent) SetValue(query core.Qt__InputMethodQuery, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QInputMethodQueryEvent_SetValue(ptr.Pointer(), C.int(query), core.PointerFromQVariant(value))
	}
}

func (ptr *QInputMethodQueryEvent) Value(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QInputMethodQueryEvent_Value(ptr.Pointer(), C.int(query)))
	}
	return nil
}
