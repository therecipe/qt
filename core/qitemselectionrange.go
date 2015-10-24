package core

//#include "qitemselectionrange.h"
import "C"
import (
	"unsafe"
)

type QItemSelectionRange struct {
	ptr unsafe.Pointer
}

type QItemSelectionRangeITF interface {
	QItemSelectionRangePTR() *QItemSelectionRange
}

func (p *QItemSelectionRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemSelectionRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemSelectionRange(ptr QItemSelectionRangeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelectionRangePTR().Pointer()
	}
	return nil
}

func QItemSelectionRangeFromPointer(ptr unsafe.Pointer) *QItemSelectionRange {
	var n = new(QItemSelectionRange)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemSelectionRange) QItemSelectionRangePTR() *QItemSelectionRange {
	return ptr
}

func (ptr *QItemSelectionRange) Intersects(other QItemSelectionRangeITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_Intersects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQItemSelectionRange(other))) != 0
	}
	return false
}

func NewQItemSelectionRange() *QItemSelectionRange {
	return QItemSelectionRangeFromPointer(unsafe.Pointer(C.QItemSelectionRange_NewQItemSelectionRange()))
}

func NewQItemSelectionRange2(other QItemSelectionRangeITF) *QItemSelectionRange {
	return QItemSelectionRangeFromPointer(unsafe.Pointer(C.QItemSelectionRange_NewQItemSelectionRange2(C.QtObjectPtr(PointerFromQItemSelectionRange(other)))))
}

func NewQItemSelectionRange4(index QModelIndexITF) *QItemSelectionRange {
	return QItemSelectionRangeFromPointer(unsafe.Pointer(C.QItemSelectionRange_NewQItemSelectionRange4(C.QtObjectPtr(PointerFromQModelIndex(index)))))
}

func NewQItemSelectionRange3(topLeft QModelIndexITF, bottomRight QModelIndexITF) *QItemSelectionRange {
	return QItemSelectionRangeFromPointer(unsafe.Pointer(C.QItemSelectionRange_NewQItemSelectionRange3(C.QtObjectPtr(PointerFromQModelIndex(topLeft)), C.QtObjectPtr(PointerFromQModelIndex(bottomRight)))))
}

func (ptr *QItemSelectionRange) Bottom() int {
	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Bottom(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Contains(index QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Contains2(row int, column int, parentIndex QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_Contains2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parentIndex))) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Left() int {
	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Left(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return QAbstractItemModelFromPointer(unsafe.Pointer(C.QItemSelectionRange_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QItemSelectionRange) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QItemSelectionRange_Parent(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QItemSelectionRange) Right() int {
	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Right(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Top() int {
	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Top(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
