package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStringMatcher struct {
	ptr unsafe.Pointer
}

type QStringMatcher_ITF interface {
	QStringMatcher_PTR() *QStringMatcher
}

func (p *QStringMatcher) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStringMatcher) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStringMatcher(ptr QStringMatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringMatcher_PTR().Pointer()
	}
	return nil
}

func NewQStringMatcherFromPointer(ptr unsafe.Pointer) *QStringMatcher {
	var n = new(QStringMatcher)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStringMatcher) QStringMatcher_PTR() *QStringMatcher {
	return ptr
}

func NewQStringMatcher3(uc QChar_ITF, length int, cs Qt__CaseSensitivity) *QStringMatcher {
	defer qt.Recovering("QStringMatcher::QStringMatcher")

	return NewQStringMatcherFromPointer(C.QStringMatcher_NewQStringMatcher3(PointerFromQChar(uc), C.int(length), C.int(cs)))
}

func (ptr *QStringMatcher) Pattern() string {
	defer qt.Recovering("QStringMatcher::pattern")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStringMatcher_Pattern(ptr.Pointer()))
	}
	return ""
}

func NewQStringMatcher() *QStringMatcher {
	defer qt.Recovering("QStringMatcher::QStringMatcher")

	return NewQStringMatcherFromPointer(C.QStringMatcher_NewQStringMatcher())
}

func NewQStringMatcher2(pattern string, cs Qt__CaseSensitivity) *QStringMatcher {
	defer qt.Recovering("QStringMatcher::QStringMatcher")

	return NewQStringMatcherFromPointer(C.QStringMatcher_NewQStringMatcher2(C.CString(pattern), C.int(cs)))
}

func NewQStringMatcher4(other QStringMatcher_ITF) *QStringMatcher {
	defer qt.Recovering("QStringMatcher::QStringMatcher")

	return NewQStringMatcherFromPointer(C.QStringMatcher_NewQStringMatcher4(PointerFromQStringMatcher(other)))
}

func (ptr *QStringMatcher) CaseSensitivity() Qt__CaseSensitivity {
	defer qt.Recovering("QStringMatcher::caseSensitivity")

	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QStringMatcher_CaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringMatcher) IndexIn2(str QChar_ITF, length int, from int) int {
	defer qt.Recovering("QStringMatcher::indexIn")

	if ptr.Pointer() != nil {
		return int(C.QStringMatcher_IndexIn2(ptr.Pointer(), PointerFromQChar(str), C.int(length), C.int(from)))
	}
	return 0
}

func (ptr *QStringMatcher) IndexIn(str string, from int) int {
	defer qt.Recovering("QStringMatcher::indexIn")

	if ptr.Pointer() != nil {
		return int(C.QStringMatcher_IndexIn(ptr.Pointer(), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QStringMatcher) SetCaseSensitivity(cs Qt__CaseSensitivity) {
	defer qt.Recovering("QStringMatcher::setCaseSensitivity")

	if ptr.Pointer() != nil {
		C.QStringMatcher_SetCaseSensitivity(ptr.Pointer(), C.int(cs))
	}
}

func (ptr *QStringMatcher) SetPattern(pattern string) {
	defer qt.Recovering("QStringMatcher::setPattern")

	if ptr.Pointer() != nil {
		C.QStringMatcher_SetPattern(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QStringMatcher) DestroyQStringMatcher() {
	defer qt.Recovering("QStringMatcher::~QStringMatcher")

	if ptr.Pointer() != nil {
		C.QStringMatcher_DestroyQStringMatcher(ptr.Pointer())
	}
}
