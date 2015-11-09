package gui

//#include "qtextdocumentfragment.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextDocumentFragment struct {
	ptr unsafe.Pointer
}

type QTextDocumentFragment_ITF interface {
	QTextDocumentFragment_PTR() *QTextDocumentFragment
}

func (p *QTextDocumentFragment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextDocumentFragment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextDocumentFragment(ptr QTextDocumentFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocumentFragment_PTR().Pointer()
	}
	return nil
}

func NewQTextDocumentFragmentFromPointer(ptr unsafe.Pointer) *QTextDocumentFragment {
	var n = new(QTextDocumentFragment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextDocumentFragment) QTextDocumentFragment_PTR() *QTextDocumentFragment {
	return ptr
}

func NewQTextDocumentFragment4(other QTextDocumentFragment_ITF) *QTextDocumentFragment {
	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment4(PointerFromQTextDocumentFragment(other)))
}

func NewQTextDocumentFragment() *QTextDocumentFragment {
	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment())
}

func NewQTextDocumentFragment3(cursor QTextCursor_ITF) *QTextDocumentFragment {
	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment3(PointerFromQTextCursor(cursor)))
}

func NewQTextDocumentFragment2(document QTextDocument_ITF) *QTextDocumentFragment {
	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment2(PointerFromQTextDocument(document)))
}

func (ptr *QTextDocumentFragment) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocumentFragment_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocumentFragment) ToHtml(encoding core.QByteArray_ITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentFragment_ToHtml(ptr.Pointer(), core.PointerFromQByteArray(encoding)))
	}
	return ""
}

func (ptr *QTextDocumentFragment) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentFragment_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocumentFragment) DestroyQTextDocumentFragment() {
	if ptr.Pointer() != nil {
		C.QTextDocumentFragment_DestroyQTextDocumentFragment(ptr.Pointer())
	}
}
