package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QFontMetrics::QFontMetrics")

	return NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics(PointerFromQFont(font)))
}

func NewQFontMetrics2(font QFont_ITF, paintdevice QPaintDevice_ITF) *QFontMetrics {
	defer qt.Recovering("QFontMetrics::QFontMetrics")

	return NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics2(PointerFromQFont(font), PointerFromQPaintDevice(paintdevice)))
}

func NewQFontMetrics3(fm QFontMetrics_ITF) *QFontMetrics {
	defer qt.Recovering("QFontMetrics::QFontMetrics")

	return NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics3(PointerFromQFontMetrics(fm)))
}

func (ptr *QFontMetrics) Ascent() int {
	defer qt.Recovering("QFontMetrics::ascent")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) AverageCharWidth() int {
	defer qt.Recovering("QFontMetrics::averageCharWidth")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_AverageCharWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) Descent() int {
	defer qt.Recovering("QFontMetrics::descent")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) ElidedText(text string, mode core.Qt__TextElideMode, width int, flags int) string {
	defer qt.Recovering("QFontMetrics::elidedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFontMetrics_ElidedText(ptr.Pointer(), C.CString(text), C.int(mode), C.int(width), C.int(flags)))
	}
	return ""
}

func (ptr *QFontMetrics) Height() int {
	defer qt.Recovering("QFontMetrics::height")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) InFont(ch core.QChar_ITF) bool {
	defer qt.Recovering("QFontMetrics::inFont")

	if ptr.Pointer() != nil {
		return C.QFontMetrics_InFont(ptr.Pointer(), core.PointerFromQChar(ch)) != 0
	}
	return false
}

func (ptr *QFontMetrics) Leading() int {
	defer qt.Recovering("QFontMetrics::leading")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Leading(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) LeftBearing(ch core.QChar_ITF) int {
	defer qt.Recovering("QFontMetrics::leftBearing")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LeftBearing(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetrics) LineSpacing() int {
	defer qt.Recovering("QFontMetrics::lineSpacing")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LineSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) LineWidth() int {
	defer qt.Recovering("QFontMetrics::lineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) MaxWidth() int {
	defer qt.Recovering("QFontMetrics::maxWidth")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MaxWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) MinLeftBearing() int {
	defer qt.Recovering("QFontMetrics::minLeftBearing")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MinLeftBearing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) MinRightBearing() int {
	defer qt.Recovering("QFontMetrics::minRightBearing")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MinRightBearing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) OverlinePos() int {
	defer qt.Recovering("QFontMetrics::overlinePos")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_OverlinePos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) RightBearing(ch core.QChar_ITF) int {
	defer qt.Recovering("QFontMetrics::rightBearing")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_RightBearing(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetrics) StrikeOutPos() int {
	defer qt.Recovering("QFontMetrics::strikeOutPos")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_StrikeOutPos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) Swap(other QFontMetrics_ITF) {
	defer qt.Recovering("QFontMetrics::swap")

	if ptr.Pointer() != nil {
		C.QFontMetrics_Swap(ptr.Pointer(), PointerFromQFontMetrics(other))
	}
}

func (ptr *QFontMetrics) UnderlinePos() int {
	defer qt.Recovering("QFontMetrics::underlinePos")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_UnderlinePos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) Width3(ch core.QChar_ITF) int {
	defer qt.Recovering("QFontMetrics::width")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Width3(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetrics) Width(text string, len int) int {
	defer qt.Recovering("QFontMetrics::width")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Width(ptr.Pointer(), C.CString(text), C.int(len)))
	}
	return 0
}

func (ptr *QFontMetrics) XHeight() int {
	defer qt.Recovering("QFontMetrics::xHeight")

	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_XHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetrics) DestroyQFontMetrics() {
	defer qt.Recovering("QFontMetrics::~QFontMetrics")

	if ptr.Pointer() != nil {
		C.QFontMetrics_DestroyQFontMetrics(ptr.Pointer())
	}
}
