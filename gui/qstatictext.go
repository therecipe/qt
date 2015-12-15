package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStaticText struct {
	ptr unsafe.Pointer
}

type QStaticText_ITF interface {
	QStaticText_PTR() *QStaticText
}

func (p *QStaticText) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStaticText) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStaticText(ptr QStaticText_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStaticText_PTR().Pointer()
	}
	return nil
}

func NewQStaticTextFromPointer(ptr unsafe.Pointer) *QStaticText {
	var n = new(QStaticText)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStaticText) QStaticText_PTR() *QStaticText {
	return ptr
}

//QStaticText::PerformanceHint
type QStaticText__PerformanceHint int64

const (
	QStaticText__ModerateCaching   = QStaticText__PerformanceHint(0)
	QStaticText__AggressiveCaching = QStaticText__PerformanceHint(1)
)

func NewQStaticText() *QStaticText {
	defer qt.Recovering("QStaticText::QStaticText")

	return NewQStaticTextFromPointer(C.QStaticText_NewQStaticText())
}

func NewQStaticText3(other QStaticText_ITF) *QStaticText {
	defer qt.Recovering("QStaticText::QStaticText")

	return NewQStaticTextFromPointer(C.QStaticText_NewQStaticText3(PointerFromQStaticText(other)))
}

func NewQStaticText2(text string) *QStaticText {
	defer qt.Recovering("QStaticText::QStaticText")

	return NewQStaticTextFromPointer(C.QStaticText_NewQStaticText2(C.CString(text)))
}

func (ptr *QStaticText) PerformanceHint() QStaticText__PerformanceHint {
	defer qt.Recovering("QStaticText::performanceHint")

	if ptr.Pointer() != nil {
		return QStaticText__PerformanceHint(C.QStaticText_PerformanceHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStaticText) Prepare(matrix QTransform_ITF, font QFont_ITF) {
	defer qt.Recovering("QStaticText::prepare")

	if ptr.Pointer() != nil {
		C.QStaticText_Prepare(ptr.Pointer(), PointerFromQTransform(matrix), PointerFromQFont(font))
	}
}

func (ptr *QStaticText) SetPerformanceHint(performanceHint QStaticText__PerformanceHint) {
	defer qt.Recovering("QStaticText::setPerformanceHint")

	if ptr.Pointer() != nil {
		C.QStaticText_SetPerformanceHint(ptr.Pointer(), C.int(performanceHint))
	}
}

func (ptr *QStaticText) SetText(text string) {
	defer qt.Recovering("QStaticText::setText")

	if ptr.Pointer() != nil {
		C.QStaticText_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QStaticText) SetTextFormat(textFormat core.Qt__TextFormat) {
	defer qt.Recovering("QStaticText::setTextFormat")

	if ptr.Pointer() != nil {
		C.QStaticText_SetTextFormat(ptr.Pointer(), C.int(textFormat))
	}
}

func (ptr *QStaticText) SetTextOption(textOption QTextOption_ITF) {
	defer qt.Recovering("QStaticText::setTextOption")

	if ptr.Pointer() != nil {
		C.QStaticText_SetTextOption(ptr.Pointer(), PointerFromQTextOption(textOption))
	}
}

func (ptr *QStaticText) SetTextWidth(textWidth float64) {
	defer qt.Recovering("QStaticText::setTextWidth")

	if ptr.Pointer() != nil {
		C.QStaticText_SetTextWidth(ptr.Pointer(), C.double(textWidth))
	}
}

func (ptr *QStaticText) Swap(other QStaticText_ITF) {
	defer qt.Recovering("QStaticText::swap")

	if ptr.Pointer() != nil {
		C.QStaticText_Swap(ptr.Pointer(), PointerFromQStaticText(other))
	}
}

func (ptr *QStaticText) Text() string {
	defer qt.Recovering("QStaticText::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStaticText_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStaticText) TextFormat() core.Qt__TextFormat {
	defer qt.Recovering("QStaticText::textFormat")

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QStaticText_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStaticText) TextWidth() float64 {
	defer qt.Recovering("QStaticText::textWidth")

	if ptr.Pointer() != nil {
		return float64(C.QStaticText_TextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStaticText) DestroyQStaticText() {
	defer qt.Recovering("QStaticText::~QStaticText")

	if ptr.Pointer() != nil {
		C.QStaticText_DestroyQStaticText(ptr.Pointer())
	}
}
