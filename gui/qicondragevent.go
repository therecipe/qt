package gui

//#include "qicondragevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIconDragEvent struct {
	core.QEvent
}

type QIconDragEvent_ITF interface {
	core.QEvent_ITF
	QIconDragEvent_PTR() *QIconDragEvent
}

func PointerFromQIconDragEvent(ptr QIconDragEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconDragEvent_PTR().Pointer()
	}
	return nil
}

func NewQIconDragEventFromPointer(ptr unsafe.Pointer) *QIconDragEvent {
	var n = new(QIconDragEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIconDragEvent) QIconDragEvent_PTR() *QIconDragEvent {
	return ptr
}

func NewQIconDragEvent() *QIconDragEvent {
	return NewQIconDragEventFromPointer(C.QIconDragEvent_NewQIconDragEvent())
}
