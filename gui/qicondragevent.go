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

type QIconDragEventITF interface {
	core.QEventITF
	QIconDragEventPTR() *QIconDragEvent
}

func PointerFromQIconDragEvent(ptr QIconDragEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconDragEventPTR().Pointer()
	}
	return nil
}

func QIconDragEventFromPointer(ptr unsafe.Pointer) *QIconDragEvent {
	var n = new(QIconDragEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIconDragEvent) QIconDragEventPTR() *QIconDragEvent {
	return ptr
}

func NewQIconDragEvent() *QIconDragEvent {
	return QIconDragEventFromPointer(unsafe.Pointer(C.QIconDragEvent_NewQIconDragEvent()))
}
