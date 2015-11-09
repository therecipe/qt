package core

//#include "qbitarray.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QBitArray struct {
	ptr unsafe.Pointer
}

type QBitArray_ITF interface {
	QBitArray_PTR() *QBitArray
}

func (p *QBitArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBitArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBitArray(ptr QBitArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitArray_PTR().Pointer()
	}
	return nil
}

func NewQBitArrayFromPointer(ptr unsafe.Pointer) *QBitArray {
	var n = new(QBitArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBitArray) QBitArray_PTR() *QBitArray {
	return ptr
}

func NewQBitArray() *QBitArray {
	return NewQBitArrayFromPointer(C.QBitArray_NewQBitArray())
}

func NewQBitArray4(other QBitArray_ITF) *QBitArray {
	return NewQBitArrayFromPointer(C.QBitArray_NewQBitArray4(PointerFromQBitArray(other)))
}

func NewQBitArray3(other QBitArray_ITF) *QBitArray {
	return NewQBitArrayFromPointer(C.QBitArray_NewQBitArray3(PointerFromQBitArray(other)))
}

func NewQBitArray2(size int, value bool) *QBitArray {
	return NewQBitArrayFromPointer(C.QBitArray_NewQBitArray2(C.int(size), C.int(qt.GoBoolToInt(value))))
}

func (ptr *QBitArray) At(i int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_At(ptr.Pointer(), C.int(i)) != 0
	}
	return false
}

func (ptr *QBitArray) Clear() {
	if ptr.Pointer() != nil {
		C.QBitArray_Clear(ptr.Pointer())
	}
}

func (ptr *QBitArray) ClearBit(i int) {
	if ptr.Pointer() != nil {
		C.QBitArray_ClearBit(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QBitArray) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QBitArray_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBitArray) Count2(on bool) int {
	if ptr.Pointer() != nil {
		return int(C.QBitArray_Count2(ptr.Pointer(), C.int(qt.GoBoolToInt(on))))
	}
	return 0
}

func (ptr *QBitArray) Fill(value bool, size int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_Fill(ptr.Pointer(), C.int(qt.GoBoolToInt(value)), C.int(size)) != 0
	}
	return false
}

func (ptr *QBitArray) Fill2(value bool, begin int, end int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Fill2(ptr.Pointer(), C.int(qt.GoBoolToInt(value)), C.int(begin), C.int(end))
	}
}

func (ptr *QBitArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBitArray) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBitArray) Resize(size int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Resize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBitArray) SetBit(i int) {
	if ptr.Pointer() != nil {
		C.QBitArray_SetBit(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QBitArray) SetBit2(i int, value bool) {
	if ptr.Pointer() != nil {
		C.QBitArray_SetBit2(ptr.Pointer(), C.int(i), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QBitArray) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QBitArray_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBitArray) Swap(other QBitArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBitArray_Swap(ptr.Pointer(), PointerFromQBitArray(other))
	}
}

func (ptr *QBitArray) TestBit(i int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_TestBit(ptr.Pointer(), C.int(i)) != 0
	}
	return false
}

func (ptr *QBitArray) ToggleBit(i int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_ToggleBit(ptr.Pointer(), C.int(i)) != 0
	}
	return false
}

func (ptr *QBitArray) Truncate(pos int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Truncate(ptr.Pointer(), C.int(pos))
	}
}
