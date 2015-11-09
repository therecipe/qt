package gui

//#include "qaccessibletableinterface.h"
import "C"
import (
	"unsafe"
)

type QAccessibleTableInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleTableInterface_ITF interface {
	QAccessibleTableInterface_PTR() *QAccessibleTableInterface
}

func (p *QAccessibleTableInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleTableInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleTableInterface(ptr QAccessibleTableInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTableInterface_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTableInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleTableInterface {
	var n = new(QAccessibleTableInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTableInterface) QAccessibleTableInterface_PTR() *QAccessibleTableInterface {
	return ptr
}

func (ptr *QAccessibleTableInterface) Caption() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_Caption(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) CellAt(row int, column int) *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_CellAt(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) ColumnDescription(column int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_ColumnDescription(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) IsColumnSelected(column int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsColumnSelected(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) IsRowSelected(row int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsRowSelected(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) ModelChange(event QAccessibleTableModelChangeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_ModelChange(ptr.Pointer(), PointerFromQAccessibleTableModelChangeEvent(event))
	}
}

func (ptr *QAccessibleTableInterface) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) RowDescription(row int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_RowDescription(ptr.Pointer(), C.int(row)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) SelectColumn(column int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectColumn(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectRow(row int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectedCellCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedCellCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedRowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedRowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) Summary() *QAccessibleInterface {
	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_Summary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) UnselectColumn(column int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectColumn(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) UnselectRow(row int) bool {
	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) DestroyQAccessibleTableInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_DestroyQAccessibleTableInterface(ptr.Pointer())
	}
}
