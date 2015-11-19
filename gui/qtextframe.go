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

type QTextFrame_ITF interface {
	QTextObject_ITF
	QTextFrame_PTR() *QTextFrame
}

func PointerFromQTextFrame(ptr QTextFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrame_PTR().Pointer()
	}
	return nil
}

func NewQTextFrameFromPointer(ptr unsafe.Pointer) *QTextFrame {
	var n = new(QTextFrame)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextFrame) QTextFrame_PTR() *QTextFrame {
	return ptr
}

func NewQTextFrame(document QTextDocument_ITF) *QTextFrame {
	return NewQTextFrameFromPointer(C.QTextFrame_NewQTextFrame(PointerFromQTextDocument(document)))
}

func (ptr *QTextFrame) FirstPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFrame_FirstPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrame) LastPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextFrame_LastPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrame) ParentFrame() *QTextFrame {
	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextFrame_ParentFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextFrame) SetFrameFormat(format QTextFrameFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFrame_SetFrameFormat(ptr.Pointer(), PointerFromQTextFrameFormat(format))
	}
}

func (ptr *QTextFrame) DestroyQTextFrame() {
	if ptr.Pointer() != nil {
		C.QTextFrame_DestroyQTextFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
