package core

//#include "qmetaenum.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMetaEnum struct {
	ptr unsafe.Pointer
}

type QMetaEnum_ITF interface {
	QMetaEnum_PTR() *QMetaEnum
}

func (p *QMetaEnum) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaEnum) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaEnum(ptr QMetaEnum_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaEnum_PTR().Pointer()
	}
	return nil
}

func NewQMetaEnumFromPointer(ptr unsafe.Pointer) *QMetaEnum {
	var n = new(QMetaEnum)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaEnum) QMetaEnum_PTR() *QMetaEnum {
	return ptr
}

func (ptr *QMetaEnum) IsFlag() bool {
	if ptr.Pointer() != nil {
		return C.QMetaEnum_IsFlag(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaEnum) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMetaEnum_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaEnum) KeyCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_KeyCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaEnum) KeyToValue(key string, ok bool) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_KeyToValue(ptr.Pointer(), C.CString(key), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QMetaEnum) KeysToValue(keys string, ok bool) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_KeysToValue(ptr.Pointer(), C.CString(keys), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QMetaEnum) Value(index int) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_Value(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QMetaEnum) ValueToKeys(value int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMetaEnum_ValueToKeys(ptr.Pointer(), C.int(value)))
	}
	return nil
}
