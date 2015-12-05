package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTextBlockFormat struct {
	QTextFormat
}

type QTextBlockFormat_ITF interface {
	QTextFormat_ITF
	QTextBlockFormat_PTR() *QTextBlockFormat
}

func PointerFromQTextBlockFormat(ptr QTextBlockFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockFormatFromPointer(ptr unsafe.Pointer) *QTextBlockFormat {
	var n = new(QTextBlockFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBlockFormat) QTextBlockFormat_PTR() *QTextBlockFormat {
	return ptr
}

//QTextBlockFormat::LineHeightTypes
type QTextBlockFormat__LineHeightTypes int64

const (
	QTextBlockFormat__SingleHeight       = QTextBlockFormat__LineHeightTypes(0)
	QTextBlockFormat__ProportionalHeight = QTextBlockFormat__LineHeightTypes(1)
	QTextBlockFormat__FixedHeight        = QTextBlockFormat__LineHeightTypes(2)
	QTextBlockFormat__MinimumHeight      = QTextBlockFormat__LineHeightTypes(3)
	QTextBlockFormat__LineDistanceHeight = QTextBlockFormat__LineHeightTypes(4)
)

func NewQTextBlockFormat() *QTextBlockFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::QTextBlockFormat")
		}
	}()

	return NewQTextBlockFormatFromPointer(C.QTextBlockFormat_NewQTextBlockFormat())
}

func (ptr *QTextBlockFormat) Alignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextBlockFormat_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) BottomMargin() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::bottomMargin")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_BottomMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) Indent() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::indent")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextBlockFormat_Indent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextBlockFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBlockFormat) LeftMargin() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::leftMargin")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_LeftMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) LineHeight2() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::lineHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_LineHeight2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) LineHeight(scriptLineHeight float64, scaling float64) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::lineHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_LineHeight(ptr.Pointer(), C.double(scriptLineHeight), C.double(scaling)))
	}
	return 0
}

func (ptr *QTextBlockFormat) LineHeightType() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::lineHeightType")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextBlockFormat_LineHeightType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) NonBreakableLines() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::nonBreakableLines")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextBlockFormat_NonBreakableLines(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBlockFormat) PageBreakPolicy() QTextFormat__PageBreakFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::pageBreakPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QTextFormat__PageBreakFlag(C.QTextBlockFormat_PageBreakPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) RightMargin() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::rightMargin")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_RightMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QTextBlockFormat) SetBottomMargin(margin float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setBottomMargin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetBottomMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextBlockFormat) SetIndent(indentation int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setIndent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetIndent(ptr.Pointer(), C.int(indentation))
	}
}

func (ptr *QTextBlockFormat) SetLeftMargin(margin float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setLeftMargin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetLeftMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextBlockFormat) SetLineHeight(height float64, heightType int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setLineHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetLineHeight(ptr.Pointer(), C.double(height), C.int(heightType))
	}
}

func (ptr *QTextBlockFormat) SetNonBreakableLines(b bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setNonBreakableLines")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetNonBreakableLines(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTextBlockFormat) SetPageBreakPolicy(policy QTextFormat__PageBreakFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setPageBreakPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetPageBreakPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTextBlockFormat) SetRightMargin(margin float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setRightMargin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetRightMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextBlockFormat) SetTextIndent(indent float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setTextIndent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetTextIndent(ptr.Pointer(), C.double(indent))
	}
}

func (ptr *QTextBlockFormat) SetTopMargin(margin float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::setTopMargin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetTopMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextBlockFormat) TextIndent() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::textIndent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_TextIndent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) TopMargin() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextBlockFormat::topMargin")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_TopMargin(ptr.Pointer()))
	}
	return 0
}
