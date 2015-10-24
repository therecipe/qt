package core

//#include "qvariant.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QVariant struct {
	ptr unsafe.Pointer
}

type QVariantITF interface {
	QVariantPTR() *QVariant
}

func (p *QVariant) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVariant) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVariant(ptr QVariantITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVariantPTR().Pointer()
	}
	return nil
}

func QVariantFromPointer(ptr unsafe.Pointer) *QVariant {
	var n = new(QVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVariant) QVariantPTR() *QVariant {
	return ptr
}

func (ptr *QVariant) ToStringList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QVariant_ToStringList(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QVariant) ToUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVariant_ToUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QVariant) DestroyQVariant() {
	if ptr.Pointer() != nil {
		C.QVariant_DestroyQVariant(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QVariant) Clear() {
	if ptr.Pointer() != nil {
		C.QVariant_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QVariant) Convert(targetTypeId int) bool {
	if ptr.Pointer() != nil {
		return C.QVariant_Convert(C.QtObjectPtr(ptr.Pointer()), C.int(targetTypeId)) != 0
	}
	return false
}

func (ptr *QVariant) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QVariant_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVariant) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QVariant_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVariant) ToBool() bool {
	if ptr.Pointer() != nil {
		return C.QVariant_ToBool(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVariant) ToInt(ok bool) int {
	if ptr.Pointer() != nil {
		return int(C.QVariant_ToInt(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QVariant) ToModelIndex() *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QVariant_ToModelIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QVariant) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVariant_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QVariant) UserType() int {
	if ptr.Pointer() != nil {
		return int(C.QVariant_UserType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
