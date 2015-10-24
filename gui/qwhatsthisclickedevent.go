package gui

//#include "qwhatsthisclickedevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWhatsThisClickedEvent struct {
	core.QEvent
}

type QWhatsThisClickedEventITF interface {
	core.QEventITF
	QWhatsThisClickedEventPTR() *QWhatsThisClickedEvent
}

func PointerFromQWhatsThisClickedEvent(ptr QWhatsThisClickedEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWhatsThisClickedEventPTR().Pointer()
	}
	return nil
}

func QWhatsThisClickedEventFromPointer(ptr unsafe.Pointer) *QWhatsThisClickedEvent {
	var n = new(QWhatsThisClickedEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWhatsThisClickedEvent) QWhatsThisClickedEventPTR() *QWhatsThisClickedEvent {
	return ptr
}

func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	return QWhatsThisClickedEventFromPointer(unsafe.Pointer(C.QWhatsThisClickedEvent_NewQWhatsThisClickedEvent(C.CString(href))))
}

func (ptr *QWhatsThisClickedEvent) Href() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWhatsThisClickedEvent_Href(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
