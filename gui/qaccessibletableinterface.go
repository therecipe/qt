package gui

//#include "qaccessibletableinterface.h"
import "C"
import (
	"unsafe"
)

type QAccessibleTableInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleTableInterfaceITF interface {
	QAccessibleTableInterfacePTR() *QAccessibleTableInterface
}

func (p *QAccessibleTableInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleTableInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleTableInterface(ptr QAccessibleTableInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTableInterfacePTR().Pointer()
	}
	return nil
}

func QAccessibleTableInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleTableInterface {
	var n = new(QAccessibleTableInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTableInterface) QAccessibleTableInterfacePTR() *QAccessibleTableInterface {
	return ptr
}

func (ptr *QAccessibleTableInterface) Caption() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleTableInterface_Caption(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) CellAt(row int, column int) *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleTableInterface_CellAt(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) ColumnDescription(column int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_ColumnDescription(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) IsColumnSelected(column int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsColumnSelected(C.QtObjectPtr(ptr.Pointer()), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) IsRowSelected(row int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsRowSelected(C.QtObjectPtr(ptr.Pointer()), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) ModelChange(event QAccessibleTableModelChangeEventITF) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_ModelChange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAccessibleTableModelChangeEvent(event)))
	}
}

func (ptr *QAccessibleTableInterface) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) RowDescription(row int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_RowDescription(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) SelectColumn(column int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectRow(row int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectRow(C.QtObjectPtr(ptr.Pointer()), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectedCellCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedCellCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedRowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedRowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) Summary() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessibleTableInterface_Summary(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) UnselectColumn(column int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) UnselectRow(row int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectRow(C.QtObjectPtr(ptr.Pointer()), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) DestroyQAccessibleTableInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_DestroyQAccessibleTableInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
