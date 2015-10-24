package gui

//#include "qtextline.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextLine struct {
	ptr unsafe.Pointer
}

type QTextLineITF interface {
	QTextLinePTR() *QTextLine
}

func (p *QTextLine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextLine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextLine(ptr QTextLineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLinePTR().Pointer()
	}
	return nil
}

func QTextLineFromPointer(ptr unsafe.Pointer) *QTextLine {
	var n = new(QTextLine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextLine) QTextLinePTR() *QTextLine {
	return ptr
}

//QTextLine::CursorPosition
type QTextLine__CursorPosition int

var (
	QTextLine__CursorBetweenCharacters = QTextLine__CursorPosition(0)
	QTextLine__CursorOnCharacter       = QTextLine__CursorPosition(1)
)

//QTextLine::Edge
type QTextLine__Edge int

var (
	QTextLine__Leading  = QTextLine__Edge(0)
	QTextLine__Trailing = QTextLine__Edge(1)
)

func NewQTextLine() *QTextLine {
	return QTextLineFromPointer(unsafe.Pointer(C.QTextLine_NewQTextLine()))
}

func (ptr *QTextLine) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextLine_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextLine) LeadingIncluded() bool {
	if ptr.Pointer() != nil {
		return C.QTextLine_LeadingIncluded(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextLine) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QTextLine_LineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextLine) SetLeadingIncluded(included bool) {
	if ptr.Pointer() != nil {
		C.QTextLine_SetLeadingIncluded(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(included)))
	}
}

func (ptr *QTextLine) SetNumColumns(numColumns int) {
	if ptr.Pointer() != nil {
		C.QTextLine_SetNumColumns(C.QtObjectPtr(ptr.Pointer()), C.int(numColumns))
	}
}

func (ptr *QTextLine) SetPosition(pos core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QTextLine_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)))
	}
}

func (ptr *QTextLine) TextLength() int {
	if ptr.Pointer() != nil {
		return int(C.QTextLine_TextLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextLine) TextStart() int {
	if ptr.Pointer() != nil {
		return int(C.QTextLine_TextStart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
