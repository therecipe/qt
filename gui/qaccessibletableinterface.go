package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QAccessibleTableInterface_") {
		n.SetObjectNameAbs("QAccessibleTableInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccessibleTableInterface) QAccessibleTableInterface_PTR() *QAccessibleTableInterface {
	return ptr
}

func (ptr *QAccessibleTableInterface) Caption() *QAccessibleInterface {
	defer qt.Recovering("QAccessibleTableInterface::caption")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_Caption(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) CellAt(row int, column int) *QAccessibleInterface {
	defer qt.Recovering("QAccessibleTableInterface::cellAt")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_CellAt(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) ColumnCount() int {
	defer qt.Recovering("QAccessibleTableInterface::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) ColumnDescription(column int) string {
	defer qt.Recovering("QAccessibleTableInterface::columnDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_ColumnDescription(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) IsColumnSelected(column int) bool {
	defer qt.Recovering("QAccessibleTableInterface::isColumnSelected")

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsColumnSelected(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) IsRowSelected(row int) bool {
	defer qt.Recovering("QAccessibleTableInterface::isRowSelected")

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsRowSelected(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) ModelChange(event QAccessibleTableModelChangeEvent_ITF) {
	defer qt.Recovering("QAccessibleTableInterface::modelChange")

	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_ModelChange(ptr.Pointer(), PointerFromQAccessibleTableModelChangeEvent(event))
	}
}

func (ptr *QAccessibleTableInterface) RowCount() int {
	defer qt.Recovering("QAccessibleTableInterface::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) RowDescription(row int) string {
	defer qt.Recovering("QAccessibleTableInterface::rowDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_RowDescription(ptr.Pointer(), C.int(row)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) SelectColumn(column int) bool {
	defer qt.Recovering("QAccessibleTableInterface::selectColumn")

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectColumn(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectRow(row int) bool {
	defer qt.Recovering("QAccessibleTableInterface::selectRow")

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectedCellCount() int {
	defer qt.Recovering("QAccessibleTableInterface::selectedCellCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedCellCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedColumnCount() int {
	defer qt.Recovering("QAccessibleTableInterface::selectedColumnCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedRowCount() int {
	defer qt.Recovering("QAccessibleTableInterface::selectedRowCount")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedRowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) Summary() *QAccessibleInterface {
	defer qt.Recovering("QAccessibleTableInterface::summary")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_Summary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) UnselectColumn(column int) bool {
	defer qt.Recovering("QAccessibleTableInterface::unselectColumn")

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectColumn(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) UnselectRow(row int) bool {
	defer qt.Recovering("QAccessibleTableInterface::unselectRow")

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) DestroyQAccessibleTableInterface() {
	defer qt.Recovering("QAccessibleTableInterface::~QAccessibleTableInterface")

	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_DestroyQAccessibleTableInterface(ptr.Pointer())
	}
}

func (ptr *QAccessibleTableInterface) ObjectNameAbs() string {
	defer qt.Recovering("QAccessibleTableInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAccessibleTableInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
