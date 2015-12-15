package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFontMetricsF struct {
	ptr unsafe.Pointer
}

type QFontMetricsF_ITF interface {
	QFontMetricsF_PTR() *QFontMetricsF
}

func (p *QFontMetricsF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontMetricsF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontMetricsF(ptr QFontMetricsF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontMetricsF_PTR().Pointer()
	}
	return nil
}

func NewQFontMetricsFFromPointer(ptr unsafe.Pointer) *QFontMetricsF {
	var n = new(QFontMetricsF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontMetricsF) QFontMetricsF_PTR() *QFontMetricsF {
	return ptr
}

func NewQFontMetricsF(font QFont_ITF) *QFontMetricsF {
	defer qt.Recovering("QFontMetricsF::QFontMetricsF")

	return NewQFontMetricsFFromPointer(C.QFontMetricsF_NewQFontMetricsF(PointerFromQFont(font)))
}

func NewQFontMetricsF2(font QFont_ITF, paintdevice QPaintDevice_ITF) *QFontMetricsF {
	defer qt.Recovering("QFontMetricsF::QFontMetricsF")

	return NewQFontMetricsFFromPointer(C.QFontMetricsF_NewQFontMetricsF2(PointerFromQFont(font), PointerFromQPaintDevice(paintdevice)))
}

func NewQFontMetricsF3(fontMetrics QFontMetrics_ITF) *QFontMetricsF {
	defer qt.Recovering("QFontMetricsF::QFontMetricsF")

	return NewQFontMetricsFFromPointer(C.QFontMetricsF_NewQFontMetricsF3(PointerFromQFontMetrics(fontMetrics)))
}

func NewQFontMetricsF4(fm QFontMetricsF_ITF) *QFontMetricsF {
	defer qt.Recovering("QFontMetricsF::QFontMetricsF")

	return NewQFontMetricsFFromPointer(C.QFontMetricsF_NewQFontMetricsF4(PointerFromQFontMetricsF(fm)))
}

func (ptr *QFontMetricsF) Ascent() float64 {
	defer qt.Recovering("QFontMetricsF::ascent")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) AverageCharWidth() float64 {
	defer qt.Recovering("QFontMetricsF::averageCharWidth")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_AverageCharWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) Descent() float64 {
	defer qt.Recovering("QFontMetricsF::descent")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) ElidedText(text string, mode core.Qt__TextElideMode, width float64, flags int) string {
	defer qt.Recovering("QFontMetricsF::elidedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFontMetricsF_ElidedText(ptr.Pointer(), C.CString(text), C.int(mode), C.double(width), C.int(flags)))
	}
	return ""
}

func (ptr *QFontMetricsF) Height() float64 {
	defer qt.Recovering("QFontMetricsF::height")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) InFont(ch core.QChar_ITF) bool {
	defer qt.Recovering("QFontMetricsF::inFont")

	if ptr.Pointer() != nil {
		return C.QFontMetricsF_InFont(ptr.Pointer(), core.PointerFromQChar(ch)) != 0
	}
	return false
}

func (ptr *QFontMetricsF) Leading() float64 {
	defer qt.Recovering("QFontMetricsF::leading")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_Leading(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) LeftBearing(ch core.QChar_ITF) float64 {
	defer qt.Recovering("QFontMetricsF::leftBearing")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_LeftBearing(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetricsF) LineSpacing() float64 {
	defer qt.Recovering("QFontMetricsF::lineSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_LineSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) LineWidth() float64 {
	defer qt.Recovering("QFontMetricsF::lineWidth")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) MaxWidth() float64 {
	defer qt.Recovering("QFontMetricsF::maxWidth")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_MaxWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) MinLeftBearing() float64 {
	defer qt.Recovering("QFontMetricsF::minLeftBearing")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_MinLeftBearing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) MinRightBearing() float64 {
	defer qt.Recovering("QFontMetricsF::minRightBearing")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_MinRightBearing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) OverlinePos() float64 {
	defer qt.Recovering("QFontMetricsF::overlinePos")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_OverlinePos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) RightBearing(ch core.QChar_ITF) float64 {
	defer qt.Recovering("QFontMetricsF::rightBearing")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_RightBearing(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetricsF) StrikeOutPos() float64 {
	defer qt.Recovering("QFontMetricsF::strikeOutPos")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_StrikeOutPos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) Swap(other QFontMetricsF_ITF) {
	defer qt.Recovering("QFontMetricsF::swap")

	if ptr.Pointer() != nil {
		C.QFontMetricsF_Swap(ptr.Pointer(), PointerFromQFontMetricsF(other))
	}
}

func (ptr *QFontMetricsF) UnderlinePos() float64 {
	defer qt.Recovering("QFontMetricsF::underlinePos")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_UnderlinePos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) Width2(ch core.QChar_ITF) float64 {
	defer qt.Recovering("QFontMetricsF::width")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_Width2(ptr.Pointer(), core.PointerFromQChar(ch)))
	}
	return 0
}

func (ptr *QFontMetricsF) Width(text string) float64 {
	defer qt.Recovering("QFontMetricsF::width")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_Width(ptr.Pointer(), C.CString(text)))
	}
	return 0
}

func (ptr *QFontMetricsF) XHeight() float64 {
	defer qt.Recovering("QFontMetricsF::xHeight")

	if ptr.Pointer() != nil {
		return float64(C.QFontMetricsF_XHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontMetricsF) DestroyQFontMetricsF() {
	defer qt.Recovering("QFontMetricsF::~QFontMetricsF")

	if ptr.Pointer() != nil {
		C.QFontMetricsF_DestroyQFontMetricsF(ptr.Pointer())
	}
}
