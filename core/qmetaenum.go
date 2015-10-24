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

type QMetaEnumITF interface {
	QMetaEnumPTR() *QMetaEnum
}

func (p *QMetaEnum) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaEnum) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaEnum(ptr QMetaEnumITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaEnumPTR().Pointer()
	}
	return nil
}

func QMetaEnumFromPointer(ptr unsafe.Pointer) *QMetaEnum {
	var n = new(QMetaEnum)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaEnum) QMetaEnumPTR() *QMetaEnum {
	return ptr
}

func (ptr *QMetaEnum) IsFlag() bool {
	if ptr.Pointer() != nil {
		return C.QMetaEnum_IsFlag(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaEnum) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMetaEnum_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaEnum) KeyCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_KeyCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaEnum) KeyToValue(key string, ok bool) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_KeyToValue(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QMetaEnum) KeysToValue(keys string, ok bool) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_KeysToValue(C.QtObjectPtr(ptr.Pointer()), C.CString(keys), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QMetaEnum) Value(index int) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaEnum_Value(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return 0
}
