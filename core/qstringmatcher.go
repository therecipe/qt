package core

//#include "qstringmatcher.h"
import "C"
import (
	"unsafe"
)

type QStringMatcher struct {
	ptr unsafe.Pointer
}

type QStringMatcherITF interface {
	QStringMatcherPTR() *QStringMatcher
}

func (p *QStringMatcher) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStringMatcher) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStringMatcher(ptr QStringMatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringMatcherPTR().Pointer()
	}
	return nil
}

func QStringMatcherFromPointer(ptr unsafe.Pointer) *QStringMatcher {
	var n = new(QStringMatcher)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStringMatcher) QStringMatcherPTR() *QStringMatcher {
	return ptr
}

func NewQStringMatcher3(uc QCharITF, length int, cs Qt__CaseSensitivity) *QStringMatcher {
	return QStringMatcherFromPointer(unsafe.Pointer(C.QStringMatcher_NewQStringMatcher3(C.QtObjectPtr(PointerFromQChar(uc)), C.int(length), C.int(cs))))
}

func (ptr *QStringMatcher) Pattern() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStringMatcher_Pattern(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQStringMatcher() *QStringMatcher {
	return QStringMatcherFromPointer(unsafe.Pointer(C.QStringMatcher_NewQStringMatcher()))
}

func NewQStringMatcher2(pattern string, cs Qt__CaseSensitivity) *QStringMatcher {
	return QStringMatcherFromPointer(unsafe.Pointer(C.QStringMatcher_NewQStringMatcher2(C.CString(pattern), C.int(cs))))
}

func NewQStringMatcher4(other QStringMatcherITF) *QStringMatcher {
	return QStringMatcherFromPointer(unsafe.Pointer(C.QStringMatcher_NewQStringMatcher4(C.QtObjectPtr(PointerFromQStringMatcher(other)))))
}

func (ptr *QStringMatcher) CaseSensitivity() Qt__CaseSensitivity {
	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QStringMatcher_CaseSensitivity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringMatcher) IndexIn2(str QCharITF, length int, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QStringMatcher_IndexIn2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(str)), C.int(length), C.int(from)))
	}
	return 0
}

func (ptr *QStringMatcher) IndexIn(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QStringMatcher_IndexIn(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QStringMatcher) SetCaseSensitivity(cs Qt__CaseSensitivity) {
	if ptr.Pointer() != nil {
		C.QStringMatcher_SetCaseSensitivity(C.QtObjectPtr(ptr.Pointer()), C.int(cs))
	}
}

func (ptr *QStringMatcher) SetPattern(pattern string) {
	if ptr.Pointer() != nil {
		C.QStringMatcher_SetPattern(C.QtObjectPtr(ptr.Pointer()), C.CString(pattern))
	}
}

func (ptr *QStringMatcher) DestroyQStringMatcher() {
	if ptr.Pointer() != nil {
		C.QStringMatcher_DestroyQStringMatcher(C.QtObjectPtr(ptr.Pointer()))
	}
}
