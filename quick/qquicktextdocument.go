package quick

//#include "qquicktextdocument.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocumentITF interface {
	core.QObjectITF
	QQuickTextDocumentPTR() *QQuickTextDocument
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocumentPTR().Pointer()
	}
	return nil
}

func QQuickTextDocumentFromPointer(ptr unsafe.Pointer) *QQuickTextDocument {
	var n = new(QQuickTextDocument)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickTextDocument_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickTextDocument) QQuickTextDocumentPTR() *QQuickTextDocument {
	return ptr
}

func NewQQuickTextDocument(parent QQuickItemITF) *QQuickTextDocument {
	return QQuickTextDocumentFromPointer(unsafe.Pointer(C.QQuickTextDocument_NewQQuickTextDocument(C.QtObjectPtr(PointerFromQQuickItem(parent)))))
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.QTextDocumentFromPointer(unsafe.Pointer(C.QQuickTextDocument_TextDocument(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
