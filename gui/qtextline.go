package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTextLine struct {
	ptr unsafe.Pointer
}

type QTextLine_ITF interface {
	QTextLine_PTR() *QTextLine
}

func (p *QTextLine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextLine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextLine(ptr QTextLine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLine_PTR().Pointer()
	}
	return nil
}

func NewQTextLineFromPointer(ptr unsafe.Pointer) *QTextLine {
	var n = new(QTextLine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextLine) QTextLine_PTR() *QTextLine {
	return ptr
}

//QTextLine::CursorPosition
type QTextLine__CursorPosition int64

const (
	QTextLine__CursorBetweenCharacters = QTextLine__CursorPosition(0)
	QTextLine__CursorOnCharacter       = QTextLine__CursorPosition(1)
)

//QTextLine::Edge
type QTextLine__Edge int64

const (
	QTextLine__Leading  = QTextLine__Edge(0)
	QTextLine__Trailing = QTextLine__Edge(1)
)

func (ptr *QTextLine) XToCursor(x float64, cpos QTextLine__CursorPosition) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::xToCursor")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLine_XToCursor(ptr.Pointer(), C.double(x), C.int(cpos)))
	}
	return 0
}

func NewQTextLine() *QTextLine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::QTextLine")
		}
	}()

	return NewQTextLineFromPointer(C.QTextLine_NewQTextLine())
}

func (ptr *QTextLine) Ascent() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::ascent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) CursorToX(cursorPos int, edge QTextLine__Edge) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::cursorToX")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_CursorToX(ptr.Pointer(), C.int(cursorPos), C.int(edge)))
	}
	return 0
}

func (ptr *QTextLine) CursorToX2(cursorPos int, edge QTextLine__Edge) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::cursorToX")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_CursorToX2(ptr.Pointer(), C.int(cursorPos), C.int(edge)))
	}
	return 0
}

func (ptr *QTextLine) Descent() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::descent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) Height() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::height")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) HorizontalAdvance() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::horizontalAdvance")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_HorizontalAdvance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextLine_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextLine) Leading() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::leading")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_Leading(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) LeadingIncluded() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::leadingIncluded")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextLine_LeadingIncluded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextLine) LineNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::lineNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLine_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) NaturalTextWidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::naturalTextWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_NaturalTextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) SetLeadingIncluded(included bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::setLeadingIncluded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLine_SetLeadingIncluded(ptr.Pointer(), C.int(qt.GoBoolToInt(included)))
	}
}

func (ptr *QTextLine) SetLineWidth(width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::setLineWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLine_SetLineWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextLine) SetNumColumns(numColumns int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::setNumColumns")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLine_SetNumColumns(ptr.Pointer(), C.int(numColumns))
	}
}

func (ptr *QTextLine) SetNumColumns2(numColumns int, alignmentWidth float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::setNumColumns")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLine_SetNumColumns2(ptr.Pointer(), C.int(numColumns), C.double(alignmentWidth))
	}
}

func (ptr *QTextLine) SetPosition(pos core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::setPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLine_SetPosition(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QTextLine) TextLength() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::textLength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLine_TextLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) TextStart() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::textStart")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLine_TextStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) Width() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::width")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) X() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::x")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLine) Y() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLine::y")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLine_Y(ptr.Pointer()))
	}
	return 0
}
