package gui

//#include "qtextblockformat.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextBlockFormat struct {
	QTextFormat
}

type QTextBlockFormatITF interface {
	QTextFormatITF
	QTextBlockFormatPTR() *QTextBlockFormat
}

func PointerFromQTextBlockFormat(ptr QTextBlockFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockFormatPTR().Pointer()
	}
	return nil
}

func QTextBlockFormatFromPointer(ptr unsafe.Pointer) *QTextBlockFormat {
	var n = new(QTextBlockFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBlockFormat) QTextBlockFormatPTR() *QTextBlockFormat {
	return ptr
}

//QTextBlockFormat::LineHeightTypes
type QTextBlockFormat__LineHeightTypes int

var (
	QTextBlockFormat__SingleHeight       = QTextBlockFormat__LineHeightTypes(0)
	QTextBlockFormat__ProportionalHeight = QTextBlockFormat__LineHeightTypes(1)
	QTextBlockFormat__FixedHeight        = QTextBlockFormat__LineHeightTypes(2)
	QTextBlockFormat__MinimumHeight      = QTextBlockFormat__LineHeightTypes(3)
	QTextBlockFormat__LineDistanceHeight = QTextBlockFormat__LineHeightTypes(4)
)

func NewQTextBlockFormat() *QTextBlockFormat {
	return QTextBlockFormatFromPointer(unsafe.Pointer(C.QTextBlockFormat_NewQTextBlockFormat()))
}

func (ptr *QTextBlockFormat) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextBlockFormat_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlockFormat) Indent() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlockFormat_Indent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlockFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextBlockFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBlockFormat) LineHeightType() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBlockFormat_LineHeightType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlockFormat) NonBreakableLines() bool {
	if ptr.Pointer() != nil {
		return C.QTextBlockFormat_NonBreakableLines(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBlockFormat) PageBreakPolicy() QTextFormat__PageBreakFlag {
	if ptr.Pointer() != nil {
		return QTextFormat__PageBreakFlag(C.QTextBlockFormat_PageBreakPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlockFormat) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QTextBlockFormat) SetIndent(indentation int) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetIndent(C.QtObjectPtr(ptr.Pointer()), C.int(indentation))
	}
}

func (ptr *QTextBlockFormat) SetNonBreakableLines(b bool) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetNonBreakableLines(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextBlockFormat) SetPageBreakPolicy(policy QTextFormat__PageBreakFlag) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetPageBreakPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}
