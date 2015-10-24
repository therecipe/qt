package core

//#include "qmimetype.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QMimeType struct {
	ptr unsafe.Pointer
}

type QMimeTypeITF interface {
	QMimeTypePTR() *QMimeType
}

func (p *QMimeType) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMimeType) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMimeType(ptr QMimeTypeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeTypePTR().Pointer()
	}
	return nil
}

func QMimeTypeFromPointer(ptr unsafe.Pointer) *QMimeType {
	var n = new(QMimeType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMimeType) QMimeTypePTR() *QMimeType {
	return ptr
}

func NewQMimeType() *QMimeType {
	return QMimeTypeFromPointer(unsafe.Pointer(C.QMimeType_NewQMimeType()))
}

func NewQMimeType2(other QMimeTypeITF) *QMimeType {
	return QMimeTypeFromPointer(unsafe.Pointer(C.QMimeType_NewQMimeType2(C.QtObjectPtr(PointerFromQMimeType(other)))))
}

func (ptr *QMimeType) FilterString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_FilterString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeType) GenericIconName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_GenericIconName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeType) GlobPatterns() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_GlobPatterns(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) IconName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_IconName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeType) Inherits(mimeTypeName string) bool {
	if ptr.Pointer() != nil {
		return C.QMimeType_Inherits(C.QtObjectPtr(ptr.Pointer()), C.CString(mimeTypeName)) != 0
	}
	return false
}

func (ptr *QMimeType) IsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QMimeType_IsDefault(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeType) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMimeType_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeType) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeType) DestroyQMimeType() {
	if ptr.Pointer() != nil {
		C.QMimeType_DestroyQMimeType(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMimeType) Aliases() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_Aliases(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) AllAncestors() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_AllAncestors(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) Comment() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_Comment(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeType) ParentMimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_ParentMimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) PreferredSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_PreferredSuffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeType) Suffixes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_Suffixes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) Swap(other QMimeTypeITF) {
	if ptr.Pointer() != nil {
		C.QMimeType_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeType(other)))
	}
}
