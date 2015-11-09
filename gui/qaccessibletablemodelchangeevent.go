package gui

//#include "qaccessibletablemodelchangeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTableModelChangeEvent struct {
	QAccessibleEvent
}

type QAccessibleTableModelChangeEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleTableModelChangeEvent_PTR() *QAccessibleTableModelChangeEvent
}

func PointerFromQAccessibleTableModelChangeEvent(ptr QAccessibleTableModelChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTableModelChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTableModelChangeEventFromPointer(ptr unsafe.Pointer) *QAccessibleTableModelChangeEvent {
	var n = new(QAccessibleTableModelChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTableModelChangeEvent) QAccessibleTableModelChangeEvent_PTR() *QAccessibleTableModelChangeEvent {
	return ptr
}

//QAccessibleTableModelChangeEvent::ModelChangeType
type QAccessibleTableModelChangeEvent__ModelChangeType int64

const (
	QAccessibleTableModelChangeEvent__ModelReset      = QAccessibleTableModelChangeEvent__ModelChangeType(0)
	QAccessibleTableModelChangeEvent__DataChanged     = QAccessibleTableModelChangeEvent__ModelChangeType(1)
	QAccessibleTableModelChangeEvent__RowsInserted    = QAccessibleTableModelChangeEvent__ModelChangeType(2)
	QAccessibleTableModelChangeEvent__ColumnsInserted = QAccessibleTableModelChangeEvent__ModelChangeType(3)
	QAccessibleTableModelChangeEvent__RowsRemoved     = QAccessibleTableModelChangeEvent__ModelChangeType(4)
	QAccessibleTableModelChangeEvent__ColumnsRemoved  = QAccessibleTableModelChangeEvent__ModelChangeType(5)
)

func NewQAccessibleTableModelChangeEvent2(iface QAccessibleInterface_ITF, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	return NewQAccessibleTableModelChangeEventFromPointer(C.QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(PointerFromQAccessibleInterface(iface), C.int(changeType)))
}

func NewQAccessibleTableModelChangeEvent(object core.QObject_ITF, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	return NewQAccessibleTableModelChangeEventFromPointer(C.QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(core.PointerFromQObject(object), C.int(changeType)))
}

func (ptr *QAccessibleTableModelChangeEvent) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_FirstColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_FirstRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) LastColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_LastColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) LastRow() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_LastRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) ModelChangeType() QAccessibleTableModelChangeEvent__ModelChangeType {
	if ptr.Pointer() != nil {
		return QAccessibleTableModelChangeEvent__ModelChangeType(C.QAccessibleTableModelChangeEvent_ModelChangeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) SetFirstColumn(column int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetFirstColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetFirstRow(row int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetFirstRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetLastColumn(column int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetLastColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetLastRow(row int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetLastRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetModelChangeType(changeType QAccessibleTableModelChangeEvent__ModelChangeType) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetModelChangeType(ptr.Pointer(), C.int(changeType))
	}
}
