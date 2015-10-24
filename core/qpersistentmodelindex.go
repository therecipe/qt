package core

//#include "qpersistentmodelindex.h"
import "C"
import (
	"unsafe"
)

type QPersistentModelIndex struct {
	ptr unsafe.Pointer
}

type QPersistentModelIndexITF interface {
	QPersistentModelIndexPTR() *QPersistentModelIndex
}

func (p *QPersistentModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPersistentModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPersistentModelIndex(ptr QPersistentModelIndexITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPersistentModelIndexPTR().Pointer()
	}
	return nil
}

func QPersistentModelIndexFromPointer(ptr unsafe.Pointer) *QPersistentModelIndex {
	var n = new(QPersistentModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPersistentModelIndex) QPersistentModelIndexPTR() *QPersistentModelIndex {
	return ptr
}

func NewQPersistentModelIndex3(other QPersistentModelIndexITF) *QPersistentModelIndex {
	return QPersistentModelIndexFromPointer(unsafe.Pointer(C.QPersistentModelIndex_NewQPersistentModelIndex3(C.QtObjectPtr(PointerFromQPersistentModelIndex(other)))))
}

func (ptr *QPersistentModelIndex) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QPersistentModelIndex_Column(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPersistentModelIndex) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QPersistentModelIndex_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPersistentModelIndex) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QPersistentModelIndex_Row(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQPersistentModelIndex4(other QPersistentModelIndexITF) *QPersistentModelIndex {
	return QPersistentModelIndexFromPointer(unsafe.Pointer(C.QPersistentModelIndex_NewQPersistentModelIndex4(C.QtObjectPtr(PointerFromQPersistentModelIndex(other)))))
}

func NewQPersistentModelIndex(index QModelIndexITF) *QPersistentModelIndex {
	return QPersistentModelIndexFromPointer(unsafe.Pointer(C.QPersistentModelIndex_NewQPersistentModelIndex(C.QtObjectPtr(PointerFromQModelIndex(index)))))
}

func (ptr *QPersistentModelIndex) Child(row int, column int) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QPersistentModelIndex_Child(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Data(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPersistentModelIndex_Data(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QPersistentModelIndex) Flags() Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QPersistentModelIndex_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPersistentModelIndex) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return QAbstractItemModelFromPointer(unsafe.Pointer(C.QPersistentModelIndex_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QPersistentModelIndex_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QPersistentModelIndex_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QPersistentModelIndex) Swap(other QPersistentModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QPersistentModelIndex_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPersistentModelIndex(other)))
	}
}
