package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QAccessibleTableModelChangeEvent::QAccessibleTableModelChangeEvent")

	return NewQAccessibleTableModelChangeEventFromPointer(C.QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(PointerFromQAccessibleInterface(iface), C.int(changeType)))
}

func NewQAccessibleTableModelChangeEvent(object core.QObject_ITF, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::QAccessibleTableModelChangeEvent")

	return NewQAccessibleTableModelChangeEventFromPointer(C.QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(core.PointerFromQObject(object), C.int(changeType)))
}

func (ptr *QAccessibleTableModelChangeEvent) FirstColumn() int {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::firstColumn")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_FirstColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) FirstRow() int {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::firstRow")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_FirstRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) LastColumn() int {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::lastColumn")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_LastColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) LastRow() int {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::lastRow")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_LastRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) ModelChangeType() QAccessibleTableModelChangeEvent__ModelChangeType {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::modelChangeType")

	if ptr.Pointer() != nil {
		return QAccessibleTableModelChangeEvent__ModelChangeType(C.QAccessibleTableModelChangeEvent_ModelChangeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) SetFirstColumn(column int) {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::setFirstColumn")

	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetFirstColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetFirstRow(row int) {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::setFirstRow")

	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetFirstRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetLastColumn(column int) {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::setLastColumn")

	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetLastColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetLastRow(row int) {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::setLastRow")

	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetLastRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetModelChangeType(changeType QAccessibleTableModelChangeEvent__ModelChangeType) {
	defer qt.Recovering("QAccessibleTableModelChangeEvent::setModelChangeType")

	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetModelChangeType(ptr.Pointer(), C.int(changeType))
	}
}
