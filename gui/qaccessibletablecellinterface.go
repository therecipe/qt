package gui

//#include "qaccessibletablecellinterface.h"
import "C"
import (
	"unsafe"
)

type QAccessibleTableCellInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleTableCellInterface_ITF interface {
	QAccessibleTableCellInterface_PTR() *QAccessibleTableCellInterface
}

func (p *QAccessibleTableCellInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleTableCellInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleTableCellInterface(ptr QAccessibleTableCellInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTableCellInterface_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTableCellInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleTableCellInterface {
	var n = new(QAccessibleTableCellInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTableCellInterface) QAccessibleTableCellInterface_PTR() *QAccessibleTableCellInterface {
	return ptr
}

func (ptr *QAccessibleTableCellInterface) ColumnExtent() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_ColumnExtent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) ColumnIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_ColumnIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) IsSelected() bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableCellInterface_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleTableCellInterface) RowExtent() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_RowExtent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) RowIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_RowIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) Table() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableCellInterface_Table(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableCellInterface) DestroyQAccessibleTableCellInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleTableCellInterface_DestroyQAccessibleTableCellInterface(ptr.Pointer())
	}
}
