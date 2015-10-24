package gui

//#include "qtextframeformat.h"
import "C"
import (
	"unsafe"
)

type QTextFrameFormat struct {
	QTextFormat
}

type QTextFrameFormatITF interface {
	QTextFormatITF
	QTextFrameFormatPTR() *QTextFrameFormat
}

func PointerFromQTextFrameFormat(ptr QTextFrameFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrameFormatPTR().Pointer()
	}
	return nil
}

func QTextFrameFormatFromPointer(ptr unsafe.Pointer) *QTextFrameFormat {
	var n = new(QTextFrameFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextFrameFormat) QTextFrameFormatPTR() *QTextFrameFormat {
	return ptr
}

//QTextFrameFormat::BorderStyle
type QTextFrameFormat__BorderStyle int

var (
	QTextFrameFormat__BorderStyle_None       = QTextFrameFormat__BorderStyle(0)
	QTextFrameFormat__BorderStyle_Dotted     = QTextFrameFormat__BorderStyle(1)
	QTextFrameFormat__BorderStyle_Dashed     = QTextFrameFormat__BorderStyle(2)
	QTextFrameFormat__BorderStyle_Solid      = QTextFrameFormat__BorderStyle(3)
	QTextFrameFormat__BorderStyle_Double     = QTextFrameFormat__BorderStyle(4)
	QTextFrameFormat__BorderStyle_DotDash    = QTextFrameFormat__BorderStyle(5)
	QTextFrameFormat__BorderStyle_DotDotDash = QTextFrameFormat__BorderStyle(6)
	QTextFrameFormat__BorderStyle_Groove     = QTextFrameFormat__BorderStyle(7)
	QTextFrameFormat__BorderStyle_Ridge      = QTextFrameFormat__BorderStyle(8)
	QTextFrameFormat__BorderStyle_Inset      = QTextFrameFormat__BorderStyle(9)
	QTextFrameFormat__BorderStyle_Outset     = QTextFrameFormat__BorderStyle(10)
)

//QTextFrameFormat::Position
type QTextFrameFormat__Position int

var (
	QTextFrameFormat__InFlow     = QTextFrameFormat__Position(0)
	QTextFrameFormat__FloatLeft  = QTextFrameFormat__Position(1)
	QTextFrameFormat__FloatRight = QTextFrameFormat__Position(2)
)

func NewQTextFrameFormat() *QTextFrameFormat {
	return QTextFrameFormatFromPointer(unsafe.Pointer(C.QTextFrameFormat_NewQTextFrameFormat()))
}

func (ptr *QTextFrameFormat) BorderStyle() QTextFrameFormat__BorderStyle {
	if ptr.Pointer() != nil {
		return QTextFrameFormat__BorderStyle(C.QTextFrameFormat_BorderStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFrameFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextFrameFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextFrameFormat) PageBreakPolicy() QTextFormat__PageBreakFlag {
	if ptr.Pointer() != nil {
		return QTextFormat__PageBreakFlag(C.QTextFrameFormat_PageBreakPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFrameFormat) Position() QTextFrameFormat__Position {
	if ptr.Pointer() != nil {
		return QTextFrameFormat__Position(C.QTextFrameFormat_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFrameFormat) SetBorderBrush(brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorderBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QTextFrameFormat) SetBorderStyle(style QTextFrameFormat__BorderStyle) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorderStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QTextFrameFormat) SetHeight(height QTextLengthITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetHeight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextLength(height)))
	}
}

func (ptr *QTextFrameFormat) SetPageBreakPolicy(policy QTextFormat__PageBreakFlag) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetPageBreakPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QTextFrameFormat) SetPosition(policy QTextFrameFormat__Position) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QTextFrameFormat) SetWidth(width QTextLengthITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetWidth(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextLength(width)))
	}
}
