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

type QTextDocumentFragmentITF interface {
	QTextDocumentFragmentPTR() *QTextDocumentFragment
}

func (p *QTextDocumentFragment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextDocumentFragment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextDocumentFragment(ptr QTextDocumentFragmentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocumentFragmentPTR().Pointer()
	}
	return nil
}

func QTextDocumentFragmentFromPointer(ptr unsafe.Pointer) *QTextDocumentFragment {
	var n = new(QTextDocumentFragment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextDocumentFragment) QTextDocumentFragmentPTR() *QTextDocumentFragment {
	return ptr
}

func NewQTextDocumentFragment4(other QTextDocumentFragmentITF) *QTextDocumentFragment {
	return QTextDocumentFragmentFromPointer(unsafe.Pointer(C.QTextDocumentFragment_NewQTextDocumentFragment4(C.QtObjectPtr(PointerFromQTextDocumentFragment(other)))))
}

func NewQTextDocumentFragment() *QTextDocumentFragment {
	return QTextDocumentFragmentFromPointer(unsafe.Pointer(C.QTextDocumentFragment_NewQTextDocumentFragment()))
}

func NewQTextDocumentFragment3(cursor QTextCursorITF) *QTextDocumentFragment {
	return QTextDocumentFragmentFromPointer(unsafe.Pointer(C.QTextDocumentFragment_NewQTextDocumentFragment3(C.QtObjectPtr(PointerFromQTextCursor(cursor)))))
}

func NewQTextDocumentFragment2(document QTextDocumentITF) *QTextDocumentFragment {
	return QTextDocumentFragmentFromPointer(unsafe.Pointer(C.QTextDocumentFragment_NewQTextDocumentFragment2(C.QtObjectPtr(PointerFromQTextDocument(document)))))
}

func (ptr *QTextDocumentFragment) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QTextDocumentFragment_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocumentFragment) ToHtml(encoding core.QByteArrayITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentFragment_ToHtml(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(encoding))))
	}
	return ""
}

func (ptr *QTextDocumentFragment) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentFragment_ToPlainText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextDocumentFragment) DestroyQTextDocumentFragment() {
	if ptr.Pointer() != nil {
		C.QTextDocumentFragment_DestroyQTextDocumentFragment(C.QtObjectPtr(ptr.Pointer()))
	}
}
