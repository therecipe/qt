package gui

//#include "qtextframe.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextFrame struct {
	QTextObject
}

type QTextFrameITF interface {
	QTextObjectITF
	QTextFramePTR() *QTextFrame
}

func PointerFromQTextFrame(ptr QTextFrameITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFramePTR().Pointer()
	}
	return nil
}

func QTextFrameFromPointer(ptr unsafe.Pointer) *QTextFrame {
	var n = new(QTextFrame)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextFrame) QTextFramePTR() *QTextFrame {
	return ptr
}

func NewQTextFrame(document QTextDocumentITF) *QTextFrame {
	return QTextFrameFromPointer(unsafe.Pointer(C.QTextFrame_NewQTextFrame(C.QtObjectPtr(PointerFromQTextDocument(document)))))
}

func (ptr *QTextFrame) FirstPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFrame_FirstPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFrame) LastPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFrame_LastPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFrame) ParentFrame() *QTextFrame {
	if ptr.Pointer() != nil {
		return QTextFrameFromPointer(unsafe.Pointer(C.QTextFrame_ParentFrame(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextFrame) SetFrameFormat(format QTextFrameFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextFrame_SetFrameFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextFrameFormat(format)))
	}
}

func (ptr *QTextFrame) DestroyQTextFrame() {
	if ptr.Pointer() != nil {
		C.QTextFrame_DestroyQTextFrame(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
