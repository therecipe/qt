package gui

//#include "qfontmetrics.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFontMetrics struct {
	ptr unsafe.Pointer
}

type QFontMetrics_ITF interface {
	QFontMetrics_PTR() *QFontMetrics
}

func (p *QFontMetrics) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontMetrics) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontMetrics(ptr QFontMetrics_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontMetrics_PTR().Pointer()
	}
	return nil
}

func NewQFontMetricsFromPointer(ptr unsafe.Pointer) *QFontMetrics {
	var n = new(QFontMetrics)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontMetrics) QFontMetrics_PTR() *QFontMetrics {
	return ptr
}

func NewQFontMetrics(font QFont_ITF) *QFontMetrics {
	return NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics(PointerFromQFont(font)))
}

func NewQFontMetrics2(font QFont_ITF, paintdevice QPaintDevice_ITF) *QFontMetrics {
	return NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics2(PointerFromQFont(font), PointerFromQPaintDevice(paintdevice)))
}

func NewQFontMetrics3(fm QFontMetrics_ITF) *QFontMetrics {
	return NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics3(PointerFromQFontMetrics(fm)))
}

func (ptr *QFontMetrics) Ascent() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) AverageCharWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_AverageCharWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) Descent() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) ElidedText(text string, mode core.Qt__TextElideMode, width int, flags int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontMetrics_ElidedText(ptr.Pointer(), C.CString(text), C.int(mode), C.int(width), C.int(flags)))
	}
	return ""
}

func (ptr *QFontMetrics) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) InFont(ch core.QChar_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QFontMetrics_InFont(ptr.Pointer(), core.PointerFromQChar(ch)) != 0
	}
	return false
}

func (ptr *QFontMetrics) Leading() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Leading(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) LeftBearing(ch core.QChar_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LeftBearing(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetrics) LineSpacing() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LineSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) LineWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) MaxWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MaxWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) MinLeftBearing() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MinLeftBearing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) MinRightBearing() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MinRightBearing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) OverlinePos() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_OverlinePos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) RightBearing(ch core.QChar_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_RightBearing(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetrics) StrikeOutPos() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_StrikeOutPos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) Swap(other QFontMetrics_ITF) {
	if ptr.Pointer() != nil {
		C.QFontMetrics_Swap(ptr.Pointer(), PointerFromQFontMetrics(other))
	}
}

func (ptr *QFontMetrics) UnderlinePos() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_UnderlinePos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) Width3(ch core.QChar_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Width3(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetrics) Width(text string, len int) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Width(ptr.Pointer(), C.CString(text), C.int(len)))
	}
	return 0
}

func (ptr *QFontMetrics) XHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_XHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) DestroyQFontMetrics() {
	if ptr.Pointer() != nil {
		C.QFontMetrics_DestroyQFontMetrics(ptr.Pointer())
	}
}
