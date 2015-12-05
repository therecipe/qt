package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QTextFrame_") {
		n.SetObjectName("QTextFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextFrame) QTextFrame_PTR() *QTextFrame {
	return ptr
}

func NewQTextFrame(document QTextDocument_ITF) *QTextFrame {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFrame::QTextFrame")
		}
	}()

	return NewQTextFrameFromPointer(C.QTextFrame_NewQTextFrame(PointerFromQTextDocument(document)))
}

func (ptr *QTextFrame) FirstPosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFrame::firstPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFrame_FirstPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrame) LastPosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFrame::lastPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFrame_LastPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrame) ParentFrame() *QTextFrame {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFrame::parentFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextFrameFromPointer(C.QTextFrame_ParentFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextFrame) SetFrameFormat(format QTextFrameFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFrame::setFrameFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFrame_SetFrameFormat(ptr.Pointer(), PointerFromQTextFrameFormat(format))
	}
}

func (ptr *QTextFrame) DestroyQTextFrame() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFrame::~QTextFrame")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFrame_DestroyQTextFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
