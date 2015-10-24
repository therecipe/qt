package core

//#include "qmetaproperty.h"
import "C"
import (
	"unsafe"
)

type QMetaProperty struct {
	ptr unsafe.Pointer
}

type QMetaPropertyITF interface {
	QMetaPropertyPTR() *QMetaProperty
}

func (p *QMetaProperty) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaProperty) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaProperty(ptr QMetaPropertyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaPropertyPTR().Pointer()
	}
	return nil
}

func QMetaPropertyFromPointer(ptr unsafe.Pointer) *QMetaProperty {
	var n = new(QMetaProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaProperty) QMetaPropertyPTR() *QMetaProperty {
	return ptr
}

func (ptr *QMetaProperty) HasNotifySignal() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_HasNotifySignal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsConstant() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsConstant(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsDesignable(object QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsDesignable(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsEnumType() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsEnumType(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsFinal() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsFinal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsFlagType() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsFlagType(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsReadable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsResettable() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsResettable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsScriptable(object QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsScriptable(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsStored(object QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsStored(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsUser(object QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsUser(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_IsWritable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) NotifySignalIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_NotifySignalIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaProperty) PropertyIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_PropertyIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaProperty) Read(object QObjectITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMetaProperty_Read(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))))
	}
	return ""
}

func (ptr *QMetaProperty) Reset(object QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_Reset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QMetaProperty) Revision() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_Revision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaProperty) UserType() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaProperty_UserType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaProperty) Write(object QObjectITF, value string) bool {
	if ptr.Pointer() != nil {
		return C.QMetaProperty_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)), C.CString(value)) != 0
	}
	return false
}
