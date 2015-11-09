package core

//#include "qpersistentmodelindex.h"
import "C"
import (
	"unsafe"
)

type QPersistentModelIndex struct {
	ptr unsafe.Pointer
}

type QPersistentModelIndex_ITF interface {
	QPersistentModelIndex_PTR() *QPersistentModelIndex
}

func (p *QPersistentModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPersistentModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPersistentModelIndex(ptr QPersistentModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPersistentModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQPersistentModelIndexFromPointer(ptr unsafe.Pointer) *QPersistentModelIndex {
	var n = new(QPersistentModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPersistentModelIndex) QPersistentModelIndex_PTR() *QPersistentModelIndex {
	return ptr
}

func NewQPersistentModelIndex3(other QPersistentModelIndex_ITF) *QPersistentModelIndex {
	return NewQPersistentModelIndexFromPointer(C.QPersistentModelIndex_NewQPersistentModelIndex3(PointerFromQPersistentModelIndex(other)))
}

func (ptr *QPersistentModelIndex) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QPersistentModelIndex_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPersistentModelIndex) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QPersistentModelIndex_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPersistentModelIndex) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QPersistentModelIndex_Row(ptr.Pointer()))
	}
	return 0
}

func NewQPersistentModelIndex4(other QPersistentModelIndex_ITF) *QPersistentModelIndex {
	return NewQPersistentModelIndexFromPointer(C.QPersistentModelIndex_NewQPersistentModelIndex4(PointerFromQPersistentModelIndex(other)))
}

func NewQPersistentModelIndex(index QModelIndex_ITF) *QPersistentModelIndex {
	return NewQPersistentModelIndexFromPointer(C.QPersistentModelIndex_NewQPersistentModelIndex(PointerFromQModelIndex(index)))
}

func (ptr *QPersistentModelIndex) Child(row int, column int) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QPersistentModelIndex_Child(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Data(role int) *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QPersistentModelIndex_Data(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Flags() Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QPersistentModelIndex_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPersistentModelIndex) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QPersistentModelIndex_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QPersistentModelIndex_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QPersistentModelIndex_Sibling(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Swap(other QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QPersistentModelIndex_Swap(ptr.Pointer(), PointerFromQPersistentModelIndex(other))
	}
}
