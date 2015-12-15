package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QLoggingCategory struct {
	ptr unsafe.Pointer
}

type QLoggingCategory_ITF interface {
	QLoggingCategory_PTR() *QLoggingCategory
}

func (p *QLoggingCategory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLoggingCategory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLoggingCategory(ptr QLoggingCategory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLoggingCategory_PTR().Pointer()
	}
	return nil
}

func NewQLoggingCategoryFromPointer(ptr unsafe.Pointer) *QLoggingCategory {
	var n = new(QLoggingCategory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLoggingCategory) QLoggingCategory_PTR() *QLoggingCategory {
	return ptr
}

func NewQLoggingCategory(category string) *QLoggingCategory {
	defer qt.Recovering("QLoggingCategory::QLoggingCategory")

	return NewQLoggingCategoryFromPointer(C.QLoggingCategory_NewQLoggingCategory(C.CString(category)))
}

func QLoggingCategory_DefaultCategory() *QLoggingCategory {
	defer qt.Recovering("QLoggingCategory::defaultCategory")

	return NewQLoggingCategoryFromPointer(C.QLoggingCategory_QLoggingCategory_DefaultCategory())
}

func (ptr *QLoggingCategory) IsCriticalEnabled() bool {
	defer qt.Recovering("QLoggingCategory::isCriticalEnabled")

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsCriticalEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsDebugEnabled() bool {
	defer qt.Recovering("QLoggingCategory::isDebugEnabled")

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsDebugEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsInfoEnabled() bool {
	defer qt.Recovering("QLoggingCategory::isInfoEnabled")

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsInfoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsWarningEnabled() bool {
	defer qt.Recovering("QLoggingCategory::isWarningEnabled")

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsWarningEnabled(ptr.Pointer()) != 0
	}
	return false
}

func QLoggingCategory_SetFilterRules(rules string) {
	defer qt.Recovering("QLoggingCategory::setFilterRules")

	C.QLoggingCategory_QLoggingCategory_SetFilterRules(C.CString(rules))
}

func (ptr *QLoggingCategory) DestroyQLoggingCategory() {
	defer qt.Recovering("QLoggingCategory::~QLoggingCategory")

	if ptr.Pointer() != nil {
		C.QLoggingCategory_DestroyQLoggingCategory(ptr.Pointer())
	}
}
