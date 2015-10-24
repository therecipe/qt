package core

//#include "qloggingcategory.h"
import "C"
import (
	"unsafe"
)

type QLoggingCategory struct {
	ptr unsafe.Pointer
}

type QLoggingCategoryITF interface {
	QLoggingCategoryPTR() *QLoggingCategory
}

func (p *QLoggingCategory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLoggingCategory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLoggingCategory(ptr QLoggingCategoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLoggingCategoryPTR().Pointer()
	}
	return nil
}

func QLoggingCategoryFromPointer(ptr unsafe.Pointer) *QLoggingCategory {
	var n = new(QLoggingCategory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLoggingCategory) QLoggingCategoryPTR() *QLoggingCategory {
	return ptr
}

func NewQLoggingCategory(category string) *QLoggingCategory {
	return QLoggingCategoryFromPointer(unsafe.Pointer(C.QLoggingCategory_NewQLoggingCategory(C.CString(category))))
}

func QLoggingCategory_DefaultCategory() *QLoggingCategory {
	return QLoggingCategoryFromPointer(unsafe.Pointer(C.QLoggingCategory_QLoggingCategory_DefaultCategory()))
}

func (ptr *QLoggingCategory) IsCriticalEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsCriticalEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsDebugEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsDebugEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsInfoEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsInfoEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsWarningEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsWarningEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QLoggingCategory_SetFilterRules(rules string) {
	C.QLoggingCategory_QLoggingCategory_SetFilterRules(C.CString(rules))
}

func (ptr *QLoggingCategory) DestroyQLoggingCategory() {
	if ptr.Pointer() != nil {
		C.QLoggingCategory_DestroyQLoggingCategory(C.QtObjectPtr(ptr.Pointer()))
	}
}
