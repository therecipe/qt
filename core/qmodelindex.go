package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QModelIndex struct {
	ptr unsafe.Pointer
}

type QModelIndex_ITF interface {
	QModelIndex_PTR() *QModelIndex
}

func (p *QModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQModelIndex(ptr QModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQModelIndexFromPointer(ptr unsafe.Pointer) *QModelIndex {
	var n = new(QModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QModelIndex) QModelIndex_PTR() *QModelIndex {
	return ptr
}

func NewQModelIndex() *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::QModelIndex")
		}
	}()

	return NewQModelIndexFromPointer(C.QModelIndex_NewQModelIndex())
}

func (ptr *QModelIndex) Child(row int, column int) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::child")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QModelIndex_Child(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QModelIndex) Column() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::column")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QModelIndex_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModelIndex) Data(role int) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::data")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QModelIndex_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QModelIndex) Flags() Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QModelIndex_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModelIndex) InternalPointer() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::internalPointer")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QModelIndex_InternalPointer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModelIndex) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QModelIndex_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModelIndex) Model() *QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::model")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QModelIndex_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModelIndex) Parent() *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QModelIndex_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModelIndex) Row() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::row")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QModelIndex_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModelIndex) Sibling(row int, column int) *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QModelIndex::sibling")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QModelIndex_Sibling(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}
