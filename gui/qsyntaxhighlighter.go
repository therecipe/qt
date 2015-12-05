package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QSyntaxHighlighter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSyntaxHighlighter) QSyntaxHighlighter_PTR() *QSyntaxHighlighter {
	return ptr
}

func (ptr *QSyntaxHighlighter) Document() *QTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSyntaxHighlighter::document")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QSyntaxHighlighter_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSyntaxHighlighter) Rehighlight() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSyntaxHighlighter::rehighlight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_Rehighlight(ptr.Pointer())
	}
}

func (ptr *QSyntaxHighlighter) RehighlightBlock(block QTextBlock_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSyntaxHighlighter::rehighlightBlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_RehighlightBlock(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QSyntaxHighlighter) SetDocument(doc QTextDocument_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSyntaxHighlighter::setDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_SetDocument(ptr.Pointer(), PointerFromQTextDocument(doc))
	}
}

func (ptr *QSyntaxHighlighter) DestroyQSyntaxHighlighter() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSyntaxHighlighter::~QSyntaxHighlighter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSyntaxHighlighter_DestroyQSyntaxHighlighter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
