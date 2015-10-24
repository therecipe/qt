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

type QFontMetricsITF interface {
	QFontMetricsPTR() *QFontMetrics
}

func (p *QFontMetrics) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontMetrics) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontMetrics(ptr QFontMetricsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontMetricsPTR().Pointer()
	}
	return nil
}

func QFontMetricsFromPointer(ptr unsafe.Pointer) *QFontMetrics {
	var n = new(QFontMetrics)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontMetrics) QFontMetricsPTR() *QFontMetrics {
	return ptr
}

func NewQFontMetrics(font QFontITF) *QFontMetrics {
	return QFontMetricsFromPointer(unsafe.Pointer(C.QFontMetrics_NewQFontMetrics(C.QtObjectPtr(PointerFromQFont(font)))))
}

func NewQFontMetrics2(font QFontITF, paintdevice QPaintDeviceITF) *QFontMetrics {
	return QFontMetricsFromPointer(unsafe.Pointer(C.QFontMetrics_NewQFontMetrics2(C.QtObjectPtr(PointerFromQFont(font)), C.QtObjectPtr(PointerFromQPaintDevice(paintdevice)))))
}

func NewQFontMetrics3(fm QFontMetricsITF) *QFontMetrics {
	return QFontMetricsFromPointer(unsafe.Pointer(C.QFontMetrics_NewQFontMetrics3(C.QtObjectPtr(PointerFromQFontMetrics(fm)))))
}

func (ptr *QFontMetrics) Ascent() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Ascent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) AverageCharWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_AverageCharWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) Descent() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Descent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) ElidedText(text string, mode core.Qt__TextElideMode, width int, flags int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontMetrics_ElidedText(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.int(mode), C.int(width), C.int(flags)))
	}
	return ""
}

func (ptr *QFontMetrics) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) InFont(ch core.QCharITF) bool {
	if ptr.Pointer() != nil {
		return C.QFontMetrics_InFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQChar(ch))) != 0
	}
	return false
}

func (ptr *QFontMetrics) Leading() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Leading(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) LeftBearing(ch core.QCharITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LeftBearing(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQChar(ch))))
	}
	return 0
}

func (ptr *QFontMetrics) LineSpacing() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LineSpacing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) LineWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_LineWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) MaxWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MaxWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) MinLeftBearing() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MinLeftBearing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) MinRightBearing() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_MinRightBearing(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) OverlinePos() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_OverlinePos(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) RightBearing(ch core.QCharITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_RightBearing(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQChar(ch))))
	}
	return 0
}

func (ptr *QFontMetrics) StrikeOutPos() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_StrikeOutPos(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) Swap(other QFontMetricsITF) {
	if ptr.Pointer() != nil {
		C.QFontMetrics_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFontMetrics(other)))
	}
}

func (ptr *QFontMetrics) UnderlinePos() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_UnderlinePos(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) Width3(ch core.QCharITF) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Width3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQChar(ch))))
	}
	return 0
}

func (ptr *QFontMetrics) Width(text string, len int) int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_Width(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.int(len)))
	}
	return 0
}

func (ptr *QFontMetrics) XHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QFontMetrics_XHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) DestroyQFontMetrics() {
	if ptr.Pointer() != nil {
		C.QFontMetrics_DestroyQFontMetrics(C.QtObjectPtr(ptr.Pointer()))
	}
}
