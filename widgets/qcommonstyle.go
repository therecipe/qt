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

type QCommonStyleITF interface {
	QStyleITF
	QCommonStylePTR() *QCommonStyle
}

func PointerFromQCommonStyle(ptr QCommonStyleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommonStylePTR().Pointer()
	}
	return nil
}

func QCommonStyleFromPointer(ptr unsafe.Pointer) *QCommonStyle {
	var n = new(QCommonStyle)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCommonStyle_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCommonStyle) QCommonStylePTR() *QCommonStyle {
	return ptr
}

func (ptr *QCommonStyle) DrawControl(element QStyle__ControlElement, opt QStyleOptionITF, p gui.QPainterITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawControl(C.QtObjectPtr(ptr.Pointer()), C.int(element), C.QtObjectPtr(PointerFromQStyleOption(opt)), C.QtObjectPtr(gui.PointerFromQPainter(p)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QCommonStyle) DrawPrimitive(pe QStyle__PrimitiveElement, opt QStyleOptionITF, p gui.QPainterITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawPrimitive(C.QtObjectPtr(ptr.Pointer()), C.int(pe), C.QtObjectPtr(PointerFromQStyleOption(opt)), C.QtObjectPtr(gui.PointerFromQPainter(p)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QCommonStyle) DrawComplexControl(cc QStyle__ComplexControl, opt QStyleOptionComplexITF, p gui.QPainterITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawComplexControl(C.QtObjectPtr(ptr.Pointer()), C.int(cc), C.QtObjectPtr(PointerFromQStyleOptionComplex(opt)), C.QtObjectPtr(gui.PointerFromQPainter(p)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QCommonStyle) HitTestComplexControl(cc QStyle__ComplexControl, opt QStyleOptionComplexITF, pt core.QPointITF, widget QWidgetITF) QStyle__SubControl {
	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QCommonStyle_HitTestComplexControl(C.QtObjectPtr(ptr.Pointer()), C.int(cc), C.QtObjectPtr(PointerFromQStyleOptionComplex(opt)), C.QtObjectPtr(core.PointerFromQPoint(pt)), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QCommonStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOptionITF, widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_LayoutSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(control1), C.int(control2), C.int(orientation), C.QtObjectPtr(PointerFromQStyleOption(option)), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QCommonStyle) PixelMetric(m QStyle__PixelMetric, opt QStyleOptionITF, widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_PixelMetric(C.QtObjectPtr(ptr.Pointer()), C.int(m), C.QtObjectPtr(PointerFromQStyleOption(opt)), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QCommonStyle) Polish2(app QApplicationITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQApplication(app)))
	}
}

func (ptr *QCommonStyle) Polish(pal gui.QPaletteITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPalette(pal)))
	}
}

func (ptr *QCommonStyle) Polish3(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QCommonStyle) StyleHint(sh QStyle__StyleHint, opt QStyleOptionITF, widget QWidgetITF, hret QStyleHintReturnITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_StyleHint(C.QtObjectPtr(ptr.Pointer()), C.int(sh), C.QtObjectPtr(PointerFromQStyleOption(opt)), C.QtObjectPtr(PointerFromQWidget(widget)), C.QtObjectPtr(PointerFromQStyleHintReturn(hret))))
	}
	return 0
}

func (ptr *QCommonStyle) Unpolish2(application QApplicationITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Unpolish2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQApplication(application)))
	}
}

func (ptr *QCommonStyle) Unpolish(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QCommonStyle_Unpolish(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QCommonStyle) DestroyQCommonStyle() {
	if ptr.Pointer() != nil {
		C.QCommonStyle_DestroyQCommonStyle(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
