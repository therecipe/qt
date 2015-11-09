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

type QMimeType_ITF interface {
	QMimeType_PTR() *QMimeType
}

func (p *QMimeType) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMimeType) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMimeType(ptr QMimeType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeType_PTR().Pointer()
	}
	return nil
}

func NewQMimeTypeFromPointer(ptr unsafe.Pointer) *QMimeType {
	var n = new(QMimeType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMimeType) QMimeType_PTR() *QMimeType {
	return ptr
}

func NewQMimeType() *QMimeType {
	return NewQMimeTypeFromPointer(C.QMimeType_NewQMimeType())
}

func NewQMimeType2(other QMimeType_ITF) *QMimeType {
	return NewQMimeTypeFromPointer(C.QMimeType_NewQMimeType2(PointerFromQMimeType(other)))
}

func (ptr *QMimeType) FilterString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_FilterString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeType) GenericIconName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_GenericIconName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeType) GlobPatterns() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_GlobPatterns(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) IconName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_IconName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeType) Inherits(mimeTypeName string) bool {
	if ptr.Pointer() != nil {
		return C.QMimeType_Inherits(ptr.Pointer(), C.CString(mimeTypeName)) != 0
	}
	return false
}

func (ptr *QMimeType) IsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QMimeType_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeType) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMimeType_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeType) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeType) DestroyQMimeType() {
	if ptr.Pointer() != nil {
		C.QMimeType_DestroyQMimeType(ptr.Pointer())
	}
}

func (ptr *QMimeType) Aliases() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_Aliases(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) AllAncestors() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_AllAncestors(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) Comment() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_Comment(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeType) ParentMimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_ParentMimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) PreferredSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeType_PreferredSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeType) Suffixes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeType_Suffixes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeType) Swap(other QMimeType_ITF) {
	if ptr.Pointer() != nil {
		C.QMimeType_Swap(ptr.Pointer(), PointerFromQMimeType(other))
	}
}
