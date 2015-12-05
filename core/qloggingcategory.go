package core

//#include "core.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::QLoggingCategory")
		}
	}()

	return NewQLoggingCategoryFromPointer(C.QLoggingCategory_NewQLoggingCategory(C.CString(category)))
}

func QLoggingCategory_DefaultCategory() *QLoggingCategory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::defaultCategory")
		}
	}()

	return NewQLoggingCategoryFromPointer(C.QLoggingCategory_QLoggingCategory_DefaultCategory())
}

func (ptr *QLoggingCategory) IsCriticalEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::isCriticalEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsCriticalEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsDebugEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::isDebugEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsDebugEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsInfoEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::isInfoEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsInfoEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLoggingCategory) IsWarningEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::isWarningEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLoggingCategory_IsWarningEnabled(ptr.Pointer()) != 0
	}
	return false
}

func QLoggingCategory_SetFilterRules(rules string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::setFilterRules")
		}
	}()

	C.QLoggingCategory_QLoggingCategory_SetFilterRules(C.CString(rules))
}

func (ptr *QLoggingCategory) DestroyQLoggingCategory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLoggingCategory::~QLoggingCategory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLoggingCategory_DestroyQLoggingCategory(ptr.Pointer())
	}
}
