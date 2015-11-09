package core

//#include "qcollator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCollator struct {
	ptr unsafe.Pointer
}

type QCollator_ITF interface {
	QCollator_PTR() *QCollator
}

func (p *QCollator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCollator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCollator(ptr QCollator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCollator_PTR().Pointer()
	}
	return nil
}

func NewQCollatorFromPointer(ptr unsafe.Pointer) *QCollator {
	var n = new(QCollator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCollator) QCollator_PTR() *QCollator {
	return ptr
}

func (ptr *QCollator) CaseSensitivity() Qt__CaseSensitivity {
	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QCollator_CaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCollator) IgnorePunctuation() bool {
	if ptr.Pointer() != nil {
		return C.QCollator_IgnorePunctuation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCollator) NumericMode() bool {
	if ptr.Pointer() != nil {
		return C.QCollator_NumericMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCollator) SetCaseSensitivity(sensitivity Qt__CaseSensitivity) {
	if ptr.Pointer() != nil {
		C.QCollator_SetCaseSensitivity(ptr.Pointer(), C.int(sensitivity))
	}
}

func (ptr *QCollator) SetIgnorePunctuation(on bool) {
	if ptr.Pointer() != nil {
		C.QCollator_SetIgnorePunctuation(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QCollator) SetNumericMode(on bool) {
	if ptr.Pointer() != nil {
		C.QCollator_SetNumericMode(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func NewQCollator3(other QCollator_ITF) *QCollator {
	return NewQCollatorFromPointer(C.QCollator_NewQCollator3(PointerFromQCollator(other)))
}

func NewQCollator2(other QCollator_ITF) *QCollator {
	return NewQCollatorFromPointer(C.QCollator_NewQCollator2(PointerFromQCollator(other)))
}

func NewQCollator(locale QLocale_ITF) *QCollator {
	return NewQCollatorFromPointer(C.QCollator_NewQCollator(PointerFromQLocale(locale)))
}

func (ptr *QCollator) SetLocale(locale QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QCollator_SetLocale(ptr.Pointer(), PointerFromQLocale(locale))
	}
}

func (ptr *QCollator) Swap(other QCollator_ITF) {
	if ptr.Pointer() != nil {
		C.QCollator_Swap(ptr.Pointer(), PointerFromQCollator(other))
	}
}

func (ptr *QCollator) DestroyQCollator() {
	if ptr.Pointer() != nil {
		C.QCollator_DestroyQCollator(ptr.Pointer())
	}
}

func (ptr *QCollator) Compare3(s1 QChar_ITF, len1 int, s2 QChar_ITF, len2 int) int {
	if ptr.Pointer() != nil {
		return int(C.QCollator_Compare3(ptr.Pointer(), PointerFromQChar(s1), C.int(len1), PointerFromQChar(s2), C.int(len2)))
	}
	return 0
}

func (ptr *QCollator) Compare(s1 string, s2 string) int {
	if ptr.Pointer() != nil {
		return int(C.QCollator_Compare(ptr.Pointer(), C.CString(s1), C.CString(s2)))
	}
	return 0
}

func (ptr *QCollator) Compare2(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCollator_Compare2(ptr.Pointer(), PointerFromQStringRef(s1), PointerFromQStringRef(s2)))
	}
	return 0
}
