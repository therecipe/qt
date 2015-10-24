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

type QTextObjectITF interface {
	core.QObjectITF
	QTextObjectPTR() *QTextObject
}

func PointerFromQTextObject(ptr QTextObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObjectPTR().Pointer()
	}
	return nil
}

func QTextObjectFromPointer(ptr unsafe.Pointer) *QTextObject {
	var n = new(QTextObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextObject) QTextObjectPTR() *QTextObject {
	return ptr
}

func (ptr *QTextObject) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		return QTextDocumentFromPointer(unsafe.Pointer(C.QTextObject_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextObject) FormatIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextObject_FormatIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextObject) ObjectIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTextObject_ObjectIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
