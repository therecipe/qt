package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QItemSelectionRange struct {
	ptr unsafe.Pointer
}

type QItemSelectionRange_ITF interface {
	QItemSelectionRange_PTR() *QItemSelectionRange
}

func (p *QItemSelectionRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemSelectionRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemSelectionRange(ptr QItemSelectionRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelectionRange_PTR().Pointer()
	}
	return nil
}

func NewQItemSelectionRangeFromPointer(ptr unsafe.Pointer) *QItemSelectionRange {
	var n = new(QItemSelectionRange)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemSelectionRange) QItemSelectionRange_PTR() *QItemSelectionRange {
	return ptr
}

func (ptr *QItemSelectionRange) Intersects(other QItemSelectionRange_ITF) bool {
	defer qt.Recovering("QItemSelectionRange::intersects")

	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_Intersects(ptr.Pointer(), PointerFromQItemSelectionRange(other)) != 0
	}
	return false
}

func NewQItemSelectionRange() *QItemSelectionRange {
	defer qt.Recovering("QItemSelectionRange::QItemSelectionRange")

	return NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange())
}

func NewQItemSelectionRange2(other QItemSelectionRange_ITF) *QItemSelectionRange {
	defer qt.Recovering("QItemSelectionRange::QItemSelectionRange")

	return NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange2(PointerFromQItemSelectionRange(other)))
}

func NewQItemSelectionRange4(index QModelIndex_ITF) *QItemSelectionRange {
	defer qt.Recovering("QItemSelectionRange::QItemSelectionRange")

	return NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange4(PointerFromQModelIndex(index)))
}

func NewQItemSelectionRange3(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) *QItemSelectionRange {
	defer qt.Recovering("QItemSelectionRange::QItemSelectionRange")

	return NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange3(PointerFromQModelIndex(topLeft), PointerFromQModelIndex(bottomRight)))
}

func (ptr *QItemSelectionRange) Bottom() int {
	defer qt.Recovering("QItemSelectionRange::bottom")

	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemSelectionRange) Contains(index QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionRange::contains")

	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_Contains(ptr.Pointer(), PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Contains2(row int, column int, parentIndex QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionRange::contains")

	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_Contains2(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parentIndex)) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Height() int {
	defer qt.Recovering("QItemSelectionRange::height")

	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemSelectionRange) IsEmpty() bool {
	defer qt.Recovering("QItemSelectionRange::isEmpty")

	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) IsValid() bool {
	defer qt.Recovering("QItemSelectionRange::isValid")

	if ptr.Pointer() != nil {
		return C.QItemSelectionRange_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Left() int {
	defer qt.Recovering("QItemSelectionRange::left")

	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Left(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemSelectionRange) Model() *QAbstractItemModel {
	defer qt.Recovering("QItemSelectionRange::model")

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionRange_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionRange) Parent() *QModelIndex {
	defer qt.Recovering("QItemSelectionRange::parent")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QItemSelectionRange_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionRange) Right() int {
	defer qt.Recovering("QItemSelectionRange::right")

	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Right(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemSelectionRange) Top() int {
	defer qt.Recovering("QItemSelectionRange::top")

	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QItemSelectionRange) Width() int {
	defer qt.Recovering("QItemSelectionRange::width")

	if ptr.Pointer() != nil {
		return int(C.QItemSelectionRange_Width(ptr.Pointer()))
	}
	return 0
}
