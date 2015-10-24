package gui

//#include "qstatustipevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStatusTipEvent struct {
	core.QEvent
}

type QStatusTipEventITF interface {
	core.QEventITF
	QStatusTipEventPTR() *QStatusTipEvent
}

func PointerFromQStatusTipEvent(ptr QStatusTipEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusTipEventPTR().Pointer()
	}
	return nil
}

func QStatusTipEventFromPointer(ptr unsafe.Pointer) *QStatusTipEvent {
	var n = new(QStatusTipEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStatusTipEvent) QStatusTipEventPTR() *QStatusTipEvent {
	return ptr
}

func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	return QStatusTipEventFromPointer(unsafe.Pointer(C.QStatusTipEvent_NewQStatusTipEvent(C.CString(tip))))
}

func (ptr *QStatusTipEvent) Tip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusTipEvent_Tip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
