package core

//#include "qbytearraymatcher.h"
import "C"
import (
	"unsafe"
)

type QByteArrayMatcher struct {
	ptr unsafe.Pointer
}

type QByteArrayMatcher_ITF interface {
	QByteArrayMatcher_PTR() *QByteArrayMatcher
}

func (p *QByteArrayMatcher) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QByteArrayMatcher) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQByteArrayMatcher(ptr QByteArrayMatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArrayMatcher_PTR().Pointer()
	}
	return nil
}

func NewQByteArrayMatcherFromPointer(ptr unsafe.Pointer) *QByteArrayMatcher {
	var n = new(QByteArrayMatcher)
	n.SetPointer(ptr)
	return n
}

func (ptr *QByteArrayMatcher) QByteArrayMatcher_PTR() *QByteArrayMatcher {
	return ptr
}

func NewQByteArrayMatcher() *QByteArrayMatcher {
	return NewQByteArrayMatcherFromPointer(C.QByteArrayMatcher_NewQByteArrayMatcher())
}

func NewQByteArrayMatcher2(pattern QByteArray_ITF) *QByteArrayMatcher {
	return NewQByteArrayMatcherFromPointer(C.QByteArrayMatcher_NewQByteArrayMatcher2(PointerFromQByteArray(pattern)))
}

func NewQByteArrayMatcher4(other QByteArrayMatcher_ITF) *QByteArrayMatcher {
	return NewQByteArrayMatcherFromPointer(C.QByteArrayMatcher_NewQByteArrayMatcher4(PointerFromQByteArrayMatcher(other)))
}

func NewQByteArrayMatcher3(pattern string, length int) *QByteArrayMatcher {
	return NewQByteArrayMatcherFromPointer(C.QByteArrayMatcher_NewQByteArrayMatcher3(C.CString(pattern), C.int(length)))
}

func (ptr *QByteArrayMatcher) IndexIn(ba QByteArray_ITF, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArrayMatcher_IndexIn(ptr.Pointer(), PointerFromQByteArray(ba), C.int(from)))
	}
	return 0
}

func (ptr *QByteArrayMatcher) IndexIn2(str string, len int, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArrayMatcher_IndexIn2(ptr.Pointer(), C.CString(str), C.int(len), C.int(from)))
	}
	return 0
}

func (ptr *QByteArrayMatcher) Pattern() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArrayMatcher_Pattern(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArrayMatcher) SetPattern(pattern QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QByteArrayMatcher_SetPattern(ptr.Pointer(), PointerFromQByteArray(pattern))
	}
}

func (ptr *QByteArrayMatcher) DestroyQByteArrayMatcher() {
	if ptr.Pointer() != nil {
		C.QByteArrayMatcher_DestroyQByteArrayMatcher(ptr.Pointer())
	}
}
