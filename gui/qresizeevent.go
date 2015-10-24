package gui

//#include "qresizeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QResizeEvent struct {
	core.QEvent
}

type QResizeEventITF interface {
	core.QEventITF
	QResizeEventPTR() *QResizeEvent
}

func PointerFromQResizeEvent(ptr QResizeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QResizeEventPTR().Pointer()
	}
	return nil
}

func QResizeEventFromPointer(ptr unsafe.Pointer) *QResizeEvent {
	var n = new(QResizeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QResizeEvent) QResizeEventPTR() *QResizeEvent {
	return ptr
}

func NewQResizeEvent(size core.QSizeITF, oldSize core.QSizeITF) *QResizeEvent {
	return QResizeEventFromPointer(unsafe.Pointer(C.QResizeEvent_NewQResizeEvent(C.QtObjectPtr(core.PointerFromQSize(size)), C.QtObjectPtr(core.PointerFromQSize(oldSize)))))
}
