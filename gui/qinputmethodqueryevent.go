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

type QInputMethodQueryEventITF interface {
	core.QEventITF
	QInputMethodQueryEventPTR() *QInputMethodQueryEvent
}

func PointerFromQInputMethodQueryEvent(ptr QInputMethodQueryEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethodQueryEventPTR().Pointer()
	}
	return nil
}

func QInputMethodQueryEventFromPointer(ptr unsafe.Pointer) *QInputMethodQueryEvent {
	var n = new(QInputMethodQueryEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInputMethodQueryEvent) QInputMethodQueryEventPTR() *QInputMethodQueryEvent {
	return ptr
}

func NewQInputMethodQueryEvent(queries core.Qt__InputMethodQuery) *QInputMethodQueryEvent {
	return QInputMethodQueryEventFromPointer(unsafe.Pointer(C.QInputMethodQueryEvent_NewQInputMethodQueryEvent(C.int(queries))))
}

func (ptr *QInputMethodQueryEvent) Queries() core.Qt__InputMethodQuery {
	if ptr.Pointer() != nil {
		return core.Qt__InputMethodQuery(C.QInputMethodQueryEvent_Queries(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputMethodQueryEvent) SetValue(query core.Qt__InputMethodQuery, value string) {
	if ptr.Pointer() != nil {
		C.QInputMethodQueryEvent_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(query), C.CString(value))
	}
}

func (ptr *QInputMethodQueryEvent) Value(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputMethodQueryEvent_Value(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}
