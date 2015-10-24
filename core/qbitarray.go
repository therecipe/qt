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

type QBitArrayITF interface {
	QBitArrayPTR() *QBitArray
}

func (p *QBitArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBitArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBitArray(ptr QBitArrayITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitArrayPTR().Pointer()
	}
	return nil
}

func QBitArrayFromPointer(ptr unsafe.Pointer) *QBitArray {
	var n = new(QBitArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBitArray) QBitArrayPTR() *QBitArray {
	return ptr
}

func NewQBitArray() *QBitArray {
	return QBitArrayFromPointer(unsafe.Pointer(C.QBitArray_NewQBitArray()))
}

func NewQBitArray4(other QBitArrayITF) *QBitArray {
	return QBitArrayFromPointer(unsafe.Pointer(C.QBitArray_NewQBitArray4(C.QtObjectPtr(PointerFromQBitArray(other)))))
}

func NewQBitArray3(other QBitArrayITF) *QBitArray {
	return QBitArrayFromPointer(unsafe.Pointer(C.QBitArray_NewQBitArray3(C.QtObjectPtr(PointerFromQBitArray(other)))))
}

func NewQBitArray2(size int, value bool) *QBitArray {
	return QBitArrayFromPointer(unsafe.Pointer(C.QBitArray_NewQBitArray2(C.int(size), C.int(qt.GoBoolToInt(value)))))
}

func (ptr *QBitArray) At(i int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_At(C.QtObjectPtr(ptr.Pointer()), C.int(i)) != 0
	}
	return false
}

func (ptr *QBitArray) Clear() {
	if ptr.Pointer() != nil {
		C.QBitArray_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBitArray) ClearBit(i int) {
	if ptr.Pointer() != nil {
		C.QBitArray_ClearBit(C.QtObjectPtr(ptr.Pointer()), C.int(i))
	}
}

func (ptr *QBitArray) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QBitArray_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBitArray) Count2(on bool) int {
	if ptr.Pointer() != nil {
		return int(C.QBitArray_Count2(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on))))
	}
	return 0
}

func (ptr *QBitArray) Fill(value bool, size int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_Fill(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(value)), C.int(size)) != 0
	}
	return false
}

func (ptr *QBitArray) Fill2(value bool, begin int, end int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Fill2(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(value)), C.int(begin), C.int(end))
	}
}

func (ptr *QBitArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBitArray) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBitArray) Resize(size int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Resize(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QBitArray) SetBit(i int) {
	if ptr.Pointer() != nil {
		C.QBitArray_SetBit(C.QtObjectPtr(ptr.Pointer()), C.int(i))
	}
}

func (ptr *QBitArray) SetBit2(i int, value bool) {
	if ptr.Pointer() != nil {
		C.QBitArray_SetBit2(C.QtObjectPtr(ptr.Pointer()), C.int(i), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QBitArray) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QBitArray_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBitArray) Swap(other QBitArrayITF) {
	if ptr.Pointer() != nil {
		C.QBitArray_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBitArray(other)))
	}
}

func (ptr *QBitArray) TestBit(i int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_TestBit(C.QtObjectPtr(ptr.Pointer()), C.int(i)) != 0
	}
	return false
}

func (ptr *QBitArray) ToggleBit(i int) bool {
	if ptr.Pointer() != nil {
		return C.QBitArray_ToggleBit(C.QtObjectPtr(ptr.Pointer()), C.int(i)) != 0
	}
	return false
}

func (ptr *QBitArray) Truncate(pos int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Truncate(C.QtObjectPtr(ptr.Pointer()), C.int(pos))
	}
}
