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

type QStaticTextITF interface {
	QStaticTextPTR() *QStaticText
}

func (p *QStaticText) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStaticText) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStaticText(ptr QStaticTextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStaticTextPTR().Pointer()
	}
	return nil
}

func QStaticTextFromPointer(ptr unsafe.Pointer) *QStaticText {
	var n = new(QStaticText)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStaticText) QStaticTextPTR() *QStaticText {
	return ptr
}

//QStaticText::PerformanceHint
type QStaticText__PerformanceHint int

var (
	QStaticText__ModerateCaching   = QStaticText__PerformanceHint(0)
	QStaticText__AggressiveCaching = QStaticText__PerformanceHint(1)
)

func NewQStaticText() *QStaticText {
	return QStaticTextFromPointer(unsafe.Pointer(C.QStaticText_NewQStaticText()))
}

func NewQStaticText3(other QStaticTextITF) *QStaticText {
	return QStaticTextFromPointer(unsafe.Pointer(C.QStaticText_NewQStaticText3(C.QtObjectPtr(PointerFromQStaticText(other)))))
}

func NewQStaticText2(text string) *QStaticText {
	return QStaticTextFromPointer(unsafe.Pointer(C.QStaticText_NewQStaticText2(C.CString(text))))
}

func (ptr *QStaticText) PerformanceHint() QStaticText__PerformanceHint {
	if ptr.Pointer() != nil {
		return QStaticText__PerformanceHint(C.QStaticText_PerformanceHint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStaticText) Prepare(matrix QTransformITF, font QFontITF) {
	if ptr.Pointer() != nil {
		C.QStaticText_Prepare(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTransform(matrix)), C.QtObjectPtr(PointerFromQFont(font)))
	}
}

func (ptr *QStaticText) SetPerformanceHint(performanceHint QStaticText__PerformanceHint) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetPerformanceHint(C.QtObjectPtr(ptr.Pointer()), C.int(performanceHint))
	}
}

func (ptr *QStaticText) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QStaticText) SetTextFormat(textFormat core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetTextFormat(C.QtObjectPtr(ptr.Pointer()), C.int(textFormat))
	}
}

func (ptr *QStaticText) SetTextOption(textOption QTextOptionITF) {
	if ptr.Pointer() != nil {
		C.QStaticText_SetTextOption(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextOption(textOption)))
	}
}

func (ptr *QStaticText) Swap(other QStaticTextITF) {
	if ptr.Pointer() != nil {
		C.QStaticText_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStaticText(other)))
	}
}

func (ptr *QStaticText) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStaticText_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStaticText) TextFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QStaticText_TextFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStaticText) DestroyQStaticText() {
	if ptr.Pointer() != nil {
		C.QStaticText_DestroyQStaticText(C.QtObjectPtr(ptr.Pointer()))
	}
}
