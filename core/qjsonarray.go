package core

//#include "qjsonarray.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QJsonArray struct {
	ptr unsafe.Pointer
}

type QJsonArray_ITF interface {
	QJsonArray_PTR() *QJsonArray
}

func (p *QJsonArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonArray(ptr QJsonArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonArray_PTR().Pointer()
	}
	return nil
}

func NewQJsonArrayFromPointer(ptr unsafe.Pointer) *QJsonArray {
	var n = new(QJsonArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonArray) QJsonArray_PTR() *QJsonArray {
	return ptr
}

func (ptr *QJsonArray) Append(value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Append(ptr.Pointer(), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) Contains(value QJsonValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJsonArray_Contains(ptr.Pointer(), PointerFromQJsonValue(value)) != 0
	}
	return false
}

func (ptr *QJsonArray) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonArray_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonArray) Empty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonArray_Empty(ptr.Pointer()) != 0
	}
	return false
}

func QJsonArray_FromStringList(list []string) *QJsonArray {
	return NewQJsonArrayFromPointer(C.QJsonArray_QJsonArray_FromStringList(C.CString(strings.Join(list, "|"))))
}

func (ptr *QJsonArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonArray_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonArray) Pop_back() {
	if ptr.Pointer() != nil {
		C.QJsonArray_Pop_back(ptr.Pointer())
	}
}

func (ptr *QJsonArray) Pop_front() {
	if ptr.Pointer() != nil {
		C.QJsonArray_Pop_front(ptr.Pointer())
	}
}

func (ptr *QJsonArray) Prepend(value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Prepend(ptr.Pointer(), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) Push_back(value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Push_back(ptr.Pointer(), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) Push_front(value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Push_front(ptr.Pointer(), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) RemoveAt(i int) {
	if ptr.Pointer() != nil {
		C.QJsonArray_RemoveAt(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QJsonArray) RemoveFirst() {
	if ptr.Pointer() != nil {
		C.QJsonArray_RemoveFirst(ptr.Pointer())
	}
}

func (ptr *QJsonArray) RemoveLast() {
	if ptr.Pointer() != nil {
		C.QJsonArray_RemoveLast(ptr.Pointer())
	}
}

func (ptr *QJsonArray) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonArray_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonArray) DestroyQJsonArray() {
	if ptr.Pointer() != nil {
		C.QJsonArray_DestroyQJsonArray(ptr.Pointer())
	}
}
