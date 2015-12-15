package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSyntaxHighlighter struct {
	core.QObject
}

type QSyntaxHighlighter_ITF interface {
	core.QObject_ITF
	QSyntaxHighlighter_PTR() *QSyntaxHighlighter
}

func PointerFromQSyntaxHighlighter(ptr QSyntaxHighlighter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSyntaxHighlighter_PTR().Pointer()
	}
	return nil
}

func NewQSyntaxHighlighterFromPointer(ptr unsafe.Pointer) *QSyntaxHighlighter {
	var n = new(QSyntaxHighlighter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSyntaxHighlighter_") {
		n.SetObjectName("QSyntaxHighlighter_" + qt.Identifier())
	}
	return n
}

func (ptr *QSyntaxHighlighter) QSyntaxHighlighter_PTR() *QSyntaxHighlighter {
	return ptr
}

func (ptr *QSyntaxHighlighter) Document() *QTextDocument {
	defer qt.Recovering("QSyntaxHighlighter::document")

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QSyntaxHighlighter_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSyntaxHighlighter) Rehighlight() {
	defer qt.Recovering("QSyntaxHighlighter::rehighlight")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_Rehighlight(ptr.Pointer())
	}
}

func (ptr *QSyntaxHighlighter) RehighlightBlock(block QTextBlock_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::rehighlightBlock")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_RehighlightBlock(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QSyntaxHighlighter) SetDocument(doc QTextDocument_ITF) {
	defer qt.Recovering("QSyntaxHighlighter::setDocument")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_SetDocument(ptr.Pointer(), PointerFromQTextDocument(doc))
	}
}

func (ptr *QSyntaxHighlighter) DestroyQSyntaxHighlighter() {
	defer qt.Recovering("QSyntaxHighlighter::~QSyntaxHighlighter")

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_DestroyQSyntaxHighlighter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
