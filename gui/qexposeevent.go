package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QExposeEvent struct {
	core.QEvent
}

type QExposeEvent_ITF interface {
	core.QEvent_ITF
	QExposeEvent_PTR() *QExposeEvent
}

func PointerFromQExposeEvent(ptr QExposeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExposeEvent_PTR().Pointer()
	}
	return nil
}

func NewQExposeEventFromPointer(ptr unsafe.Pointer) *QExposeEvent {
	var n = new(QExposeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QExposeEvent) QExposeEvent_PTR() *QExposeEvent {
	return ptr
}

func NewQExposeEvent(exposeRegion QRegion_ITF) *QExposeEvent {
	defer qt.Recovering("QExposeEvent::QExposeEvent")

	return NewQExposeEventFromPointer(C.QExposeEvent_NewQExposeEvent(PointerFromQRegion(exposeRegion)))
}

func (ptr *QExposeEvent) Region() *QRegion {
	defer qt.Recovering("QExposeEvent::region")

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QExposeEvent_Region(ptr.Pointer()))
	}
	return nil
}
