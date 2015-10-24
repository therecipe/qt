package core

//#include "qbytearraymatcher.h"
import "C"
import (
	"unsafe"
)

type QByteArrayMatcher struct {
	ptr unsafe.Pointer
}

type QByteArrayMatcherITF interface {
	QByteArrayMatcherPTR() *QByteArrayMatcher
}

func (p *QByteArrayMatcher) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QByteArrayMatcher) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQByteArrayMatcher(ptr QByteArrayMatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArrayMatcherPTR().Pointer()
	}
	return nil
}

func QByteArrayMatcherFromPointer(ptr unsafe.Pointer) *QByteArrayMatcher {
	var n = new(QByteArrayMatcher)
	n.SetPointer(ptr)
	return n
}

func (ptr *QByteArrayMatcher) QByteArrayMatcherPTR() *QByteArrayMatcher {
	return ptr
}

func NewQByteArrayMatcher() *QByteArrayMatcher {
	return QByteArrayMatcherFromPointer(unsafe.Pointer(C.QByteArrayMatcher_NewQByteArrayMatcher()))
}

func NewQByteArrayMatcher2(pattern QByteArrayITF) *QByteArrayMatcher {
	return QByteArrayMatcherFromPointer(unsafe.Pointer(C.QByteArrayMatcher_NewQByteArrayMatcher2(C.QtObjectPtr(PointerFromQByteArray(pattern)))))
}

func NewQByteArrayMatcher4(other QByteArrayMatcherITF) *QByteArrayMatcher {
	return QByteArrayMatcherFromPointer(unsafe.Pointer(C.QByteArrayMatcher_NewQByteArrayMatcher4(C.QtObjectPtr(PointerFromQByteArrayMatcher(other)))))
}

func NewQByteArrayMatcher3(pattern string, length int) *QByteArrayMatcher {
	return QByteArrayMatcherFromPointer(unsafe.Pointer(C.QByteArrayMatcher_NewQByteArrayMatcher3(C.CString(pattern), C.int(length))))
}

func (ptr *QByteArrayMatcher) IndexIn(ba QByteArrayITF, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArrayMatcher_IndexIn(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba)), C.int(from)))
	}
	return 0
}

func (ptr *QByteArrayMatcher) IndexIn2(str string, len int, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArrayMatcher_IndexIn2(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(len), C.int(from)))
	}
	return 0
}

func (ptr *QByteArrayMatcher) SetPattern(pattern QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QByteArrayMatcher_SetPattern(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(pattern)))
	}
}

func (ptr *QByteArrayMatcher) DestroyQByteArrayMatcher() {
	if ptr.Pointer() != nil {
		C.QByteArrayMatcher_DestroyQByteArrayMatcher(C.QtObjectPtr(ptr.Pointer()))
	}
}
