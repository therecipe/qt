package widgets

//#include "qcommonstyle.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QCommonStyle struct {
	QStyle
}

type QCommonStyle_ITF interface {
	QStyle_ITF
	QCommonStyle_PTR() *QCommonStyle
}

func PointerFromQCommonStyle(ptr QCommonStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommonStyle_PTR().Pointer()
	}
	return nil
}

func NewQCommonStyleFromPointer(ptr unsafe.Pointer) *QCommonStyle {
	var n = new(QCommonStyle)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QCommonStyle_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCommonStyle) QCommonStyle_PTR() *QCommonStyle {
	return ptr
}

func (ptr *QCommonStyle) DrawControl(element QStyle__ControlElement, opt QStyleOption_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawControl(ptr.Pointer(), C.int(element), PointerFromQStyleOption(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DrawPrimitive(pe QStyle__PrimitiveElement, opt QStyleOption_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawPrimitive(ptr.Pointer(), C.int(pe), PointerFromQStyleOption(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DrawComplexControl(cc QStyle__ComplexControl, opt QStyleOptionComplex_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawComplexControl(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) HitTestComplexControl(cc QStyle__ComplexControl, opt QStyleOptionComplex_ITF, pt core.QPoint_ITF, widget QWidget_ITF) QStyle__SubControl {
	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QCommonStyle_HitTestComplexControl(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(opt), core.PointerFromQPoint(pt), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QCommonStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_LayoutSpacing(ptr.Pointer(), C.int(control1), C.int(control2), C.int(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QCommonStyle) PixelMetric(m QStyle__PixelMetric, opt QStyleOption_ITF, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_PixelMetric(ptr.Pointer(), C.int(m), PointerFromQStyleOption(opt), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QCommonStyle) Polish2(app QApplication_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish2(ptr.Pointer(), PointerFromQApplication(app))
	}
}

func (ptr *QCommonStyle) Polish(pal gui.QPalette_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish(ptr.Pointer(), gui.PointerFromQPalette(pal))
	}
}

func (ptr *QCommonStyle) Polish3(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish3(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) StyleHint(sh QStyle__StyleHint, opt QStyleOption_ITF, widget QWidget_ITF, hret QStyleHintReturn_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_StyleHint(ptr.Pointer(), C.int(sh), PointerFromQStyleOption(opt), PointerFromQWidget(widget), PointerFromQStyleHintReturn(hret)))
	}
	return 0
}

func (ptr *QCommonStyle) Unpolish2(application QApplication_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Unpolish2(ptr.Pointer(), PointerFromQApplication(application))
	}
}

func (ptr *QCommonStyle) Unpolish(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Unpolish(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DestroyQCommonStyle() {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DestroyQCommonStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
