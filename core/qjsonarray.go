package core

//#include "qjsonarray.h"
import "C"
import (
	"unsafe"
)

type QJsonArray struct {
	ptr unsafe.Pointer
}

type QJsonArrayITF interface {
	QJsonArrayPTR() *QJsonArray
}

func (p *QJsonArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonArray(ptr QJsonArrayITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonArrayPTR().Pointer()
	}
	return nil
}

func QJsonArrayFromPointer(ptr unsafe.Pointer) *QJsonArray {
	var n = new(QJsonArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonArray) QJsonArrayPTR() *QJsonArray {
	return ptr
}

func (ptr *QJsonArray) Append(value QJsonValueITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Append(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonValue(value)))
	}
}

func (ptr *QJsonArray) Contains(value QJsonValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QJsonArray_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonValue(value))) != 0
	}
	return false
}

func (ptr *QJsonArray) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonArray_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonArray) Empty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonArray_Empty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonArray_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonArray) Pop_back() {
	if ptr.Pointer() != nil {
		C.QJsonArray_Pop_back(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QJsonArray) Pop_front() {
	if ptr.Pointer() != nil {
		C.QJsonArray_Pop_front(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QJsonArray) Prepend(value QJsonValueITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Prepend(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonValue(value)))
	}
}

func (ptr *QJsonArray) Push_back(value QJsonValueITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Push_back(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonValue(value)))
	}
}

func (ptr *QJsonArray) Push_front(value QJsonValueITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Push_front(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonValue(value)))
	}
}

func (ptr *QJsonArray) RemoveAt(i int) {
	if ptr.Pointer() != nil {
		C.QJsonArray_RemoveAt(C.QtObjectPtr(ptr.Pointer()), C.int(i))
	}
}

func (ptr *QJsonArray) RemoveFirst() {
	if ptr.Pointer() != nil {
		C.QJsonArray_RemoveFirst(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QJsonArray) RemoveLast() {
	if ptr.Pointer() != nil {
		C.QJsonArray_RemoveLast(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QJsonArray) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonArray_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonArray) DestroyQJsonArray() {
	if ptr.Pointer() != nil {
		C.QJsonArray_DestroyQJsonArray(C.QtObjectPtr(ptr.Pointer()))
	}
}
