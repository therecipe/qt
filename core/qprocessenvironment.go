package core

//#include "qprocessenvironment.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QProcessEnvironment struct {
	ptr unsafe.Pointer
}

type QProcessEnvironmentITF interface {
	QProcessEnvironmentPTR() *QProcessEnvironment
}

func (p *QProcessEnvironment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QProcessEnvironment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQProcessEnvironment(ptr QProcessEnvironmentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProcessEnvironmentPTR().Pointer()
	}
	return nil
}

func QProcessEnvironmentFromPointer(ptr unsafe.Pointer) *QProcessEnvironment {
	var n = new(QProcessEnvironment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QProcessEnvironment) QProcessEnvironmentPTR() *QProcessEnvironment {
	return ptr
}

func NewQProcessEnvironment() *QProcessEnvironment {
	return QProcessEnvironmentFromPointer(unsafe.Pointer(C.QProcessEnvironment_NewQProcessEnvironment()))
}

func NewQProcessEnvironment2(other QProcessEnvironmentITF) *QProcessEnvironment {
	return QProcessEnvironmentFromPointer(unsafe.Pointer(C.QProcessEnvironment_NewQProcessEnvironment2(C.QtObjectPtr(PointerFromQProcessEnvironment(other)))))
}

func (ptr *QProcessEnvironment) Clear() {
	if ptr.Pointer() != nil {
		C.QProcessEnvironment_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProcessEnvironment) Contains(name string) bool {
	if ptr.Pointer() != nil {
		return C.QProcessEnvironment_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QProcessEnvironment) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QProcessEnvironment_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProcessEnvironment) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcessEnvironment_Keys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QProcessEnvironment) Swap(other QProcessEnvironmentITF) {
	if ptr.Pointer() != nil {
		C.QProcessEnvironment_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQProcessEnvironment(other)))
	}
}

func (ptr *QProcessEnvironment) ToStringList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcessEnvironment_ToStringList(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QProcessEnvironment) Value(name string, defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProcessEnvironment_Value(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QProcessEnvironment) DestroyQProcessEnvironment() {
	if ptr.Pointer() != nil {
		C.QProcessEnvironment_DestroyQProcessEnvironment(C.QtObjectPtr(ptr.Pointer()))
	}
}
