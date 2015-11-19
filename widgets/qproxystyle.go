package widgets

//#include "qproxystyle.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QProxyStyle struct {
	QCommonStyle
}

type QProxyStyle_ITF interface {
	QCommonStyle_ITF
	QProxyStyle_PTR() *QProxyStyle
}

func PointerFromQProxyStyle(ptr QProxyStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProxyStyle_PTR().Pointer()
	}
	return nil
}

func NewQProxyStyleFromPointer(ptr unsafe.Pointer) *QProxyStyle {
	var n = new(QProxyStyle)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProxyStyle_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProxyStyle) QProxyStyle_PTR() *QProxyStyle {
	return ptr
}

func (ptr *QProxyStyle) BaseStyle() *QStyle {
	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QProxyStyle_BaseStyle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProxyStyle) DrawComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawComplexControl(ptr.Pointer(), C.int(control), PointerFromQStyleOptionComplex(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

func (ptr *QProxyStyle) DrawControl(element QStyle__ControlElement, option QStyleOption_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawControl(ptr.Pointer(), C.int(element), PointerFromQStyleOption(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

func (ptr *QProxyStyle) DrawItemPixmap(painter gui.QPainter_ITF, rect core.QRect_ITF, alignment int, pixmap gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawItemPixmap(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(alignment), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QProxyStyle) DrawItemText(painter gui.QPainter_ITF, rect core.QRect_ITF, flags int, pal gui.QPalette_ITF, enabled bool, text string, textRole gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawItemText(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(flags), gui.PointerFromQPalette(pal), C.int(qt.GoBoolToInt(enabled)), C.CString(text), C.int(textRole))
	}
}

func (ptr *QProxyStyle) DrawPrimitive(element QStyle__PrimitiveElement, option QStyleOption_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawPrimitive(ptr.Pointer(), C.int(element), PointerFromQStyleOption(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

func (ptr *QProxyStyle) HitTestComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, pos core.QPoint_ITF, widget QWidget_ITF) QStyle__SubControl {
	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QProxyStyle_HitTestComplexControl(ptr.Pointer(), C.int(control), PointerFromQStyleOptionComplex(option), core.PointerFromQPoint(pos), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QProxyStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_LayoutSpacing(ptr.Pointer(), C.int(control1), C.int(control2), C.int(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QProxyStyle) PixelMetric(metric QStyle__PixelMetric, option QStyleOption_ITF, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_PixelMetric(ptr.Pointer(), C.int(metric), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QProxyStyle) Polish3(app QApplication_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Polish3(ptr.Pointer(), PointerFromQApplication(app))
	}
}

func (ptr *QProxyStyle) Polish2(pal gui.QPalette_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Polish2(ptr.Pointer(), gui.PointerFromQPalette(pal))
	}
}

func (ptr *QProxyStyle) Polish(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Polish(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QProxyStyle) SetBaseStyle(style QStyle_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_SetBaseStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func (ptr *QProxyStyle) StyleHint(hint QStyle__StyleHint, option QStyleOption_ITF, widget QWidget_ITF, returnData QStyleHintReturn_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_StyleHint(ptr.Pointer(), C.int(hint), PointerFromQStyleOption(option), PointerFromQWidget(widget), PointerFromQStyleHintReturn(returnData)))
	}
	return 0
}

func (ptr *QProxyStyle) Unpolish2(app QApplication_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Unpolish2(ptr.Pointer(), PointerFromQApplication(app))
	}
}

func (ptr *QProxyStyle) Unpolish(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Unpolish(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QProxyStyle) DestroyQProxyStyle() {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DestroyQProxyStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
