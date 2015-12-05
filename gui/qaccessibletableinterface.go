package gui

//#include "gui.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::caption")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_Caption(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) CellAt(row int, column int) *QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::cellAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_CellAt(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) ColumnDescription(column int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::columnDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_ColumnDescription(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) IsColumnSelected(column int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::isColumnSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsColumnSelected(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) IsRowSelected(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::isRowSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_IsRowSelected(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) ModelChange(event QAccessibleTableModelChangeEvent_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::modelChange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_ModelChange(ptr.Pointer(), PointerFromQAccessibleTableModelChangeEvent(event))
	}
}

func (ptr *QAccessibleTableInterface) RowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) RowDescription(row int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::rowDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTableInterface_RowDescription(ptr.Pointer(), C.int(row)))
	}
	return ""
}

func (ptr *QAccessibleTableInterface) SelectColumn(column int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::selectColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectColumn(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectRow(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::selectRow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_SelectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) SelectedCellCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::selectedCellCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedCellCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::selectedColumnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) SelectedRowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::selectedRowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableInterface_SelectedRowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableInterface) Summary() *QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::summary")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableInterface_Summary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableInterface) UnselectColumn(column int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::unselectColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectColumn(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) UnselectRow(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::unselectRow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableInterface_UnselectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QAccessibleTableInterface) DestroyQAccessibleTableInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableInterface::~QAccessibleTableInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleTableInterface_DestroyQAccessibleTableInterface(ptr.Pointer())
	}
}
