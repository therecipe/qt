package gui

//#include "qsyntaxhighlighter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSyntaxHighlighter struct {
	core.QObject
}

type QSyntaxHighlighterITF interface {
	core.QObjectITF
	QSyntaxHighlighterPTR() *QSyntaxHighlighter
}

func PointerFromQSyntaxHighlighter(ptr QSyntaxHighlighterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSyntaxHighlighterPTR().Pointer()
	}
	return nil
}

func QSyntaxHighlighterFromPointer(ptr unsafe.Pointer) *QSyntaxHighlighter {
	var n = new(QSyntaxHighlighter)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSyntaxHighlighter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSyntaxHighlighter) QSyntaxHighlighterPTR() *QSyntaxHighlighter {
	return ptr
}

func (ptr *QSyntaxHighlighter) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		return QTextDocumentFromPointer(unsafe.Pointer(C.QSyntaxHighlighter_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSyntaxHighlighter) Rehighlight() {
	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_Rehighlight(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSyntaxHighlighter) RehighlightBlock(block QTextBlockITF) {
	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_RehighlightBlock(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextBlock(block)))
	}
}

func (ptr *QSyntaxHighlighter) SetDocument(doc QTextDocumentITF) {
	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_SetDocument(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextDocument(doc)))
	}
}

func (ptr *QSyntaxHighlighter) DestroyQSyntaxHighlighter() {
	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_DestroyQSyntaxHighlighter(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
