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

type QAccessibleTableModelChangeEventITF interface {
	QAccessibleEventITF
	QAccessibleTableModelChangeEventPTR() *QAccessibleTableModelChangeEvent
}

func PointerFromQAccessibleTableModelChangeEvent(ptr QAccessibleTableModelChangeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTableModelChangeEventPTR().Pointer()
	}
	return nil
}

func QAccessibleTableModelChangeEventFromPointer(ptr unsafe.Pointer) *QAccessibleTableModelChangeEvent {
	var n = new(QAccessibleTableModelChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTableModelChangeEvent) QAccessibleTableModelChangeEventPTR() *QAccessibleTableModelChangeEvent {
	return ptr
}

//QAccessibleTableModelChangeEvent::ModelChangeType
type QAccessibleTableModelChangeEvent__ModelChangeType int

var (
	QAccessibleTableModelChangeEvent__ModelReset      = QAccessibleTableModelChangeEvent__ModelChangeType(0)
	QAccessibleTableModelChangeEvent__DataChanged     = QAccessibleTableModelChangeEvent__ModelChangeType(1)
	QAccessibleTableModelChangeEvent__RowsInserted    = QAccessibleTableModelChangeEvent__ModelChangeType(2)
	QAccessibleTableModelChangeEvent__ColumnsInserted = QAccessibleTableModelChangeEvent__ModelChangeType(3)
	QAccessibleTableModelChangeEvent__RowsRemoved     = QAccessibleTableModelChangeEvent__ModelChangeType(4)
	QAccessibleTableModelChangeEvent__ColumnsRemoved  = QAccessibleTableModelChangeEvent__ModelChangeType(5)
)

func NewQAccessibleTableModelChangeEvent2(iface QAccessibleInterfaceITF, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	return QAccessibleTableModelChangeEventFromPointer(unsafe.Pointer(C.QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.int(changeType))))
}

func NewQAccessibleTableModelChangeEvent(object core.QObjectITF, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	return QAccessibleTableModelChangeEventFromPointer(unsafe.Pointer(C.QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(changeType))))
}

func (ptr *QAccessibleTableModelChangeEvent) FirstColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_FirstColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) FirstRow() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_FirstRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) LastColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_LastColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) LastRow() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableModelChangeEvent_LastRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) ModelChangeType() QAccessibleTableModelChangeEvent__ModelChangeType {
	if ptr.Pointer() != nil {
		return QAccessibleTableModelChangeEvent__ModelChangeType(C.QAccessibleTableModelChangeEvent_ModelChangeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableModelChangeEvent) SetFirstColumn(column int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetFirstColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetFirstRow(row int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetFirstRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetLastColumn(column int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetLastColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetLastRow(row int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetLastRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QAccessibleTableModelChangeEvent) SetModelChangeType(changeType QAccessibleTableModelChangeEvent__ModelChangeType) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableModelChangeEvent_SetModelChangeType(C.QtObjectPtr(ptr.Pointer()), C.int(changeType))
	}
}
