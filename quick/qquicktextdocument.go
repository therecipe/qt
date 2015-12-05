package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextDocumentFromPointer(ptr unsafe.Pointer) *QQuickTextDocument {
	var n = new(QQuickTextDocument)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickTextDocument_") {
		n.SetObjectName("QQuickTextDocument_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return ptr
}

func NewQQuickTextDocument(parent QQuickItem_ITF) *QQuickTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickTextDocument::QQuickTextDocument")
		}
	}()

	return NewQQuickTextDocumentFromPointer(C.QQuickTextDocument_NewQQuickTextDocument(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickTextDocument::textDocument")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QQuickTextDocument_TextDocument(ptr.Pointer()))
	}
	return nil
}
