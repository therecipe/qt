package core

//#include "qmodelindex.h"
import "C"
import (
	"unsafe"
)

type QModelIndex struct {
	ptr unsafe.Pointer
}

type QModelIndexITF interface {
	QModelIndexPTR() *QModelIndex
}

func (p *QModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQModelIndex(ptr QModelIndexITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModelIndexPTR().Pointer()
	}
	return nil
}

func QModelIndexFromPointer(ptr unsafe.Pointer) *QModelIndex {
	var n = new(QModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QModelIndex) QModelIndexPTR() *QModelIndex {
	return ptr
}

func NewQModelIndex() *QModelIndex {
	return QModelIndexFromPointer(unsafe.Pointer(C.QModelIndex_NewQModelIndex()))
}

func (ptr *QModelIndex) Child(row int, column int) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QModelIndex_Child(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QModelIndex) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QModelIndex_Column(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModelIndex) Data(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QModelIndex_Data(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QModelIndex) Flags() Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QModelIndex_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModelIndex) InternalPointer() {
	if ptr.Pointer() != nil {
		C.QModelIndex_InternalPointer(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QModelIndex) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QModelIndex_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QModelIndex) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return QAbstractItemModelFromPointer(unsafe.Pointer(C.QModelIndex_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QModelIndex) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QModelIndex_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QModelIndex) Row() int {
	if ptr.Pointer() != nil {
		return int(C.QModelIndex_Row(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModelIndex) Sibling(row int, column int) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QModelIndex_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}
