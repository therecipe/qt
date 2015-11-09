package widgets

//#include "qplaintextdocumentlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QPlainTextDocumentLayout struct {
	gui.QAbstractTextDocumentLayout
}

type QPlainTextDocumentLayout_ITF interface {
	gui.QAbstractTextDocumentLayout_ITF
	QPlainTextDocumentLayout_PTR() *QPlainTextDocumentLayout
}

func PointerFromQPlainTextDocumentLayout(ptr QPlainTextDocumentLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlainTextDocumentLayout_PTR().Pointer()
	}
	return nil
}

func NewQPlainTextDocumentLayoutFromPointer(ptr unsafe.Pointer) *QPlainTextDocumentLayout {
	var n = new(QPlainTextDocumentLayout)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPlainTextDocumentLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlainTextDocumentLayout) QPlainTextDocumentLayout_PTR() *QPlainTextDocumentLayout {
	return ptr
}

func (ptr *QPlainTextDocumentLayout) CursorWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextDocumentLayout_CursorWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_SetCursorWidth(ptr.Pointer(), C.int(width))
	}
}

func NewQPlainTextDocumentLayout(document gui.QTextDocument_ITF) *QPlainTextDocumentLayout {
	return NewQPlainTextDocumentLayoutFromPointer(C.QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(gui.PointerFromQTextDocument(document)))
}

func (ptr *QPlainTextDocumentLayout) EnsureBlockLayout(block gui.QTextBlock_ITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_EnsureBlockLayout(ptr.Pointer(), gui.PointerFromQTextBlock(block))
	}
}

func (ptr *QPlainTextDocumentLayout) PageCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextDocumentLayout_PageCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPlainTextDocumentLayout) RequestUpdate() {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QPlainTextDocumentLayout) DestroyQPlainTextDocumentLayout() {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
