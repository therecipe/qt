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

type QCollatorITF interface {
	QCollatorPTR() *QCollator
}

func (p *QCollator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCollator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCollator(ptr QCollatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCollatorPTR().Pointer()
	}
	return nil
}

func QCollatorFromPointer(ptr unsafe.Pointer) *QCollator {
	var n = new(QCollator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCollator) QCollatorPTR() *QCollator {
	return ptr
}

func (ptr *QCollator) CaseSensitivity() Qt__CaseSensitivity {
	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QCollator_CaseSensitivity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCollator) IgnorePunctuation() bool {
	if ptr.Pointer() != nil {
		return C.QCollator_IgnorePunctuation(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCollator) NumericMode() bool {
	if ptr.Pointer() != nil {
		return C.QCollator_NumericMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCollator) SetCaseSensitivity(sensitivity Qt__CaseSensitivity) {
	if ptr.Pointer() != nil {
		C.QCollator_SetCaseSensitivity(C.QtObjectPtr(ptr.Pointer()), C.int(sensitivity))
	}
}

func (ptr *QCollator) SetIgnorePunctuation(on bool) {
	if ptr.Pointer() != nil {
		C.QCollator_SetIgnorePunctuation(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QCollator) SetNumericMode(on bool) {
	if ptr.Pointer() != nil {
		C.QCollator_SetNumericMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func NewQCollator3(other QCollatorITF) *QCollator {
	return QCollatorFromPointer(unsafe.Pointer(C.QCollator_NewQCollator3(C.QtObjectPtr(PointerFromQCollator(other)))))
}

func NewQCollator2(other QCollatorITF) *QCollator {
	return QCollatorFromPointer(unsafe.Pointer(C.QCollator_NewQCollator2(C.QtObjectPtr(PointerFromQCollator(other)))))
}

func NewQCollator(locale QLocaleITF) *QCollator {
	return QCollatorFromPointer(unsafe.Pointer(C.QCollator_NewQCollator(C.QtObjectPtr(PointerFromQLocale(locale)))))
}

func (ptr *QCollator) SetLocale(locale QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QCollator_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLocale(locale)))
	}
}

func (ptr *QCollator) Swap(other QCollatorITF) {
	if ptr.Pointer() != nil {
		C.QCollator_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCollator(other)))
	}
}

func (ptr *QCollator) DestroyQCollator() {
	if ptr.Pointer() != nil {
		C.QCollator_DestroyQCollator(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCollator) Compare3(s1 QCharITF, len1 int, s2 QCharITF, len2 int) int {
	if ptr.Pointer() != nil {
		return int(C.QCollator_Compare3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(s1)), C.int(len1), C.QtObjectPtr(PointerFromQChar(s2)), C.int(len2)))
	}
	return 0
}

func (ptr *QCollator) Compare(s1 string, s2 string) int {
	if ptr.Pointer() != nil {
		return int(C.QCollator_Compare(C.QtObjectPtr(ptr.Pointer()), C.CString(s1), C.CString(s2)))
	}
	return 0
}

func (ptr *QCollator) Compare2(s1 QStringRefITF, s2 QStringRefITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCollator_Compare2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(s1)), C.QtObjectPtr(PointerFromQStringRef(s2))))
	}
	return 0
}
