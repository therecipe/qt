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

type QPlainTextDocumentLayoutITF interface {
	gui.QAbstractTextDocumentLayoutITF
	QPlainTextDocumentLayoutPTR() *QPlainTextDocumentLayout
}

func PointerFromQPlainTextDocumentLayout(ptr QPlainTextDocumentLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlainTextDocumentLayoutPTR().Pointer()
	}
	return nil
}

func QPlainTextDocumentLayoutFromPointer(ptr unsafe.Pointer) *QPlainTextDocumentLayout {
	var n = new(QPlainTextDocumentLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlainTextDocumentLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlainTextDocumentLayout) QPlainTextDocumentLayoutPTR() *QPlainTextDocumentLayout {
	return ptr
}

func (ptr *QPlainTextDocumentLayout) CursorWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextDocumentLayout_CursorWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_SetCursorWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func NewQPlainTextDocumentLayout(document gui.QTextDocumentITF) *QPlainTextDocumentLayout {
	return QPlainTextDocumentLayoutFromPointer(unsafe.Pointer(C.QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(C.QtObjectPtr(gui.PointerFromQTextDocument(document)))))
}

func (ptr *QPlainTextDocumentLayout) EnsureBlockLayout(block gui.QTextBlockITF) {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_EnsureBlockLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextBlock(block)))
	}
}

func (ptr *QPlainTextDocumentLayout) PageCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPlainTextDocumentLayout_PageCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPlainTextDocumentLayout) RequestUpdate() {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_RequestUpdate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPlainTextDocumentLayout) DestroyQPlainTextDocumentLayout() {
	if ptr.Pointer() != nil {
		C.QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
