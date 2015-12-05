package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::QTextDocumentFragment")
		}
	}()

	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment4(PointerFromQTextDocumentFragment(other)))
}

func NewQTextDocumentFragment() *QTextDocumentFragment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::QTextDocumentFragment")
		}
	}()

	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment())
}

func NewQTextDocumentFragment3(cursor QTextCursor_ITF) *QTextDocumentFragment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::QTextDocumentFragment")
		}
	}()

	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment3(PointerFromQTextCursor(cursor)))
}

func NewQTextDocumentFragment2(document QTextDocument_ITF) *QTextDocumentFragment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::QTextDocumentFragment")
		}
	}()

	return NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment2(PointerFromQTextDocument(document)))
}

func (ptr *QTextDocumentFragment) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextDocumentFragment_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextDocumentFragment) ToHtml(encoding core.QByteArray_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::toHtml")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentFragment_ToHtml(ptr.Pointer(), core.PointerFromQByteArray(encoding)))
	}
	return ""
}

func (ptr *QTextDocumentFragment) ToPlainText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::toPlainText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentFragment_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocumentFragment) DestroyQTextDocumentFragment() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextDocumentFragment::~QTextDocumentFragment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextDocumentFragment_DestroyQTextDocumentFragment(ptr.Pointer())
	}
}
