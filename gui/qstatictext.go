package gui

//#include "qstatictext.h"
import "C"
import (
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
	return NewQStaticTextFromPointer(C.QStaticText_NewQStaticText())
}

func NewQStaticText3(other QStaticText_ITF) *QStaticText {
	return NewQStaticTextFromPointer(C.QStaticText_NewQStaticText3(PointerFromQStaticText(other)))
}

func NewQStaticText2(text string) *QStaticText {
	return NewQStaticTextFromPointer(C.QStaticText_NewQStaticText2(C.CString(text)))
}

func (ptr *QStaticText) PerformanceHint() QStaticText__PerformanceHint {
	if ptr.Pointer() != nil {
		return QStaticText__PerformanceHint(C.QStaticText_PerformanceHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStaticText) Prepare(matrix QTransform_ITF, font QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QStaticText_Prepare(ptr.Pointer(), PointerFromQTransform(matrix), PointerFromQFont(font))
	}
}

func (ptr *QStaticText) SetPerformanceHint(performanceHint QStaticText__PerformanceHint) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetPerformanceHint(ptr.Pointer(), C.int(performanceHint))
	}
}

func (ptr *QStaticText) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QStaticText) SetTextFormat(textFormat core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetTextFormat(ptr.Pointer(), C.int(textFormat))
	}
}

func (ptr *QStaticText) SetTextOption(textOption QTextOption_ITF) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetTextOption(ptr.Pointer(), PointerFromQTextOption(textOption))
	}
}

func (ptr *QStaticText) SetTextWidth(textWidth float64) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetTextWidth(ptr.Pointer(), C.double(textWidth))
	}
}

func (ptr *QStaticText) Swap(other QStaticText_ITF) {
	if ptr.Pointer() != nil {
		C.QStaticText_Swap(ptr.Pointer(), PointerFromQStaticText(other))
	}
}

func (ptr *QStaticText) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStaticText_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStaticText) TextFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QStaticText_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStaticText) TextWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QStaticText_TextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStaticText) DestroyQStaticText() {
	if ptr.Pointer() != nil {
		C.QStaticText_DestroyQStaticText(ptr.Pointer())
	}
}
