package core

//#include "qregularexpressionmatch.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QRegularExpressionMatch struct {
	ptr unsafe.Pointer
}

type QRegularExpressionMatchITF interface {
	QRegularExpressionMatchPTR() *QRegularExpressionMatch
}

func (p *QRegularExpressionMatch) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegularExpressionMatch) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegularExpressionMatch(ptr QRegularExpressionMatchITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionMatchPTR().Pointer()
	}
	return nil
}

func QRegularExpressionMatchFromPointer(ptr unsafe.Pointer) *QRegularExpressionMatch {
	var n = new(QRegularExpressionMatch)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegularExpressionMatch) QRegularExpressionMatchPTR() *QRegularExpressionMatch {
	return ptr
}

func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	return QRegularExpressionMatchFromPointer(unsafe.Pointer(C.QRegularExpressionMatch_NewQRegularExpressionMatch()))
}

func NewQRegularExpressionMatch2(match QRegularExpressionMatchITF) *QRegularExpressionMatch {
	return QRegularExpressionMatchFromPointer(unsafe.Pointer(C.QRegularExpressionMatch_NewQRegularExpressionMatch2(C.QtObjectPtr(PointerFromQRegularExpressionMatch(match)))))
}

func (ptr *QRegularExpressionMatch) Captured2(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRegularExpressionMatch_Captured2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QRegularExpressionMatch) Captured(nth int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRegularExpressionMatch_Captured(C.QtObjectPtr(ptr.Pointer()), C.int(nth)))
	}
	return ""
}

func (ptr *QRegularExpressionMatch) CapturedEnd2(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedEnd2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedEnd(nth int) int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedEnd(C.QtObjectPtr(ptr.Pointer()), C.int(nth)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedLength2(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedLength2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedLength(nth int) int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedLength(C.QtObjectPtr(ptr.Pointer()), C.int(nth)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedStart2(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedStart2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedStart(nth int) int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedStart(C.QtObjectPtr(ptr.Pointer()), C.int(nth)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedTexts() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QRegularExpressionMatch_CapturedTexts(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QRegularExpressionMatch) HasMatch() bool {
	if ptr.Pointer() != nil {
		return C.QRegularExpressionMatch_HasMatch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) HasPartialMatch() bool {
	if ptr.Pointer() != nil {
		return C.QRegularExpressionMatch_HasPartialMatch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QRegularExpressionMatch_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) LastCapturedIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_LastCapturedIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) MatchOptions() QRegularExpression__MatchOption {
	if ptr.Pointer() != nil {
		return QRegularExpression__MatchOption(C.QRegularExpressionMatch_MatchOptions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) MatchType() QRegularExpression__MatchType {
	if ptr.Pointer() != nil {
		return QRegularExpression__MatchType(C.QRegularExpressionMatch_MatchType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) Swap(other QRegularExpressionMatchITF) {
	if ptr.Pointer() != nil {
		C.QRegularExpressionMatch_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegularExpressionMatch(other)))
	}
}

func (ptr *QRegularExpressionMatch) DestroyQRegularExpressionMatch() {
	if ptr.Pointer() != nil {
		C.QRegularExpressionMatch_DestroyQRegularExpressionMatch(C.QtObjectPtr(ptr.Pointer()))
	}
}
