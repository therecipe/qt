package gui

//#include "gui.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::columnExtent")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_ColumnExtent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) ColumnIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::columnIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_ColumnIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) IsSelected() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::isSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleTableCellInterface_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleTableCellInterface) RowExtent() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::rowExtent")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_RowExtent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) RowIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::rowIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTableCellInterface_RowIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTableCellInterface) Table() *QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::table")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleTableCellInterface_Table(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleTableCellInterface) DestroyQAccessibleTableCellInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTableCellInterface::~QAccessibleTableCellInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleTableCellInterface_DestroyQAccessibleTableCellInterface(ptr.Pointer())
	}
}
