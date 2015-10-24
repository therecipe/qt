package gui

//#include "qexposeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QExposeEvent struct {
	core.QEvent
}

type QExposeEventITF interface {
	core.QEventITF
	QExposeEventPTR() *QExposeEvent
}

func PointerFromQExposeEvent(ptr QExposeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExposeEventPTR().Pointer()
	}
	return nil
}

func QExposeEventFromPointer(ptr unsafe.Pointer) *QExposeEvent {
	var n = new(QExposeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QExposeEvent) QExposeEventPTR() *QExposeEvent {
	return ptr
}

func NewQExposeEvent(exposeRegion QRegionITF) *QExposeEvent {
	return QExposeEventFromPointer(unsafe.Pointer(C.QExposeEvent_NewQExposeEvent(C.QtObjectPtr(PointerFromQRegion(exposeRegion)))))
}
