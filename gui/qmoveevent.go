package gui

//#include "qmoveevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMoveEvent struct {
	core.QEvent
}

type QMoveEventITF interface {
	core.QEventITF
	QMoveEventPTR() *QMoveEvent
}

func PointerFromQMoveEvent(ptr QMoveEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMoveEventPTR().Pointer()
	}
	return nil
}

func QMoveEventFromPointer(ptr unsafe.Pointer) *QMoveEvent {
	var n = new(QMoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMoveEvent) QMoveEventPTR() *QMoveEvent {
	return ptr
}

func NewQMoveEvent(pos core.QPointITF, oldPos core.QPointITF) *QMoveEvent {
	return QMoveEventFromPointer(unsafe.Pointer(C.QMoveEvent_NewQMoveEvent(C.QtObjectPtr(core.PointerFromQPoint(pos)), C.QtObjectPtr(core.PointerFromQPoint(oldPos)))))
}
