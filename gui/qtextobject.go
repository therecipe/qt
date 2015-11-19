package gui

//#include "qtextobject.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextObject struct {
	core.QObject
}

type QTextObject_ITF interface {
	core.QObject_ITF
	QTextObject_PTR() *QTextObject
}

func PointerFromQTextObject(ptr QTextObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObject_PTR().Pointer()
	}
	return nil
}

func NewQTextObjectFromPointer(ptr unsafe.Pointer) *QTextObject {
	var n = new(QTextObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextObject) QTextObject_PTR() *QTextObject {
	return ptr
}

func (ptr *QTextObject) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		return NewQTextDocumentFromPointer(C.QTextObject_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextObject) FormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextObject_FormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextObject) ObjectIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextObject_ObjectIndex(ptr.Pointer()))
	}
	return 0
}
