package gui

//#include "qaccessibletablecellinterface.h"
import "C"
import (
	"unsafe"
)

type QAccessibleTableCellInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleTableCellInterfaceITF interface {
	QAccessibleTableCellInterfacePTR() *QAccessibleTableCellInterface
}

func (p *QAccessibleTableCellInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleTableCellInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleTableCellInterface(ptr QAccessibleTableCellInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTableCellInterfacePTR().Pointer()
	}
	return nil
}

func QAccessibleTableCellInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleTableCellInterface {
	var n = new(QAccessibleTableCellInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTableCellInterface) QAccessibleTableCellInterfacePTR() *QAccessibleTableCellInterface {
	return ptr
}

func (ptr *QAccessibleTableCellInterface) ColumnExtent() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_ColumnExtent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) ColumnIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_ColumnIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) IsSelected() bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableCellInterface_IsSelected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAccessibleTableCellInterface) RowExtent() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_RowExtent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) RowIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_RowIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) Table() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleTableCellInterface_Table(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleTableCellInterface) DestroyQAccessibleTableCellInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleTableCellInterface_DestroyQAccessibleTableCellInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
