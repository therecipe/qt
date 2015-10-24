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

type QProxyStyleITF interface {
	QCommonStyleITF
	QProxyStylePTR() *QProxyStyle
}

func PointerFromQProxyStyle(ptr QProxyStyleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProxyStylePTR().Pointer()
	}
	return nil
}

func QProxyStyleFromPointer(ptr unsafe.Pointer) *QProxyStyle {
	var n = new(QProxyStyle)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProxyStyle_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProxyStyle) QProxyStylePTR() *QProxyStyle {
	return ptr
}

func (ptr *QProxyStyle) BaseStyle() *QStyle {
	if ptr.Pointer() != nil {
		return QStyleFromPointer(unsafe.Pointer(C.QProxyStyle_BaseStyle(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QProxyStyle) DrawComplexControl(control QStyle__ComplexControl, option QStyleOptionComplexITF, painter gui.QPainterITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawComplexControl(C.QtObjectPtr(ptr.Pointer()), C.int(control), C.QtObjectPtr(PointerFromQStyleOptionComplex(option)), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QProxyStyle) DrawControl(element QStyle__ControlElement, option QStyleOptionITF, painter gui.QPainterITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawControl(C.QtObjectPtr(ptr.Pointer()), C.int(element), C.QtObjectPtr(PointerFromQStyleOption(option)), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QProxyStyle) DrawItemPixmap(painter gui.QPainterITF, rect core.QRectITF, alignment int, pixmap gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawItemPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(alignment), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QProxyStyle) DrawItemText(painter gui.QPainterITF, rect core.QRectITF, flags int, pal gui.QPaletteITF, enabled bool, text string, textRole gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawItemText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(flags), C.QtObjectPtr(gui.PointerFromQPalette(pal)), C.int(qt.GoBoolToInt(enabled)), C.CString(text), C.int(textRole))
	}
}

func (ptr *QProxyStyle) DrawPrimitive(element QStyle__PrimitiveElement, option QStyleOptionITF, painter gui.QPainterITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DrawPrimitive(C.QtObjectPtr(ptr.Pointer()), C.int(element), C.QtObjectPtr(PointerFromQStyleOption(option)), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QProxyStyle) HitTestComplexControl(control QStyle__ComplexControl, option QStyleOptionComplexITF, pos core.QPointITF, widget QWidgetITF) QStyle__SubControl {
	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QProxyStyle_HitTestComplexControl(C.QtObjectPtr(ptr.Pointer()), C.int(control), C.QtObjectPtr(PointerFromQStyleOptionComplex(option)), C.QtObjectPtr(core.PointerFromQPoint(pos)), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QProxyStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOptionITF, widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_LayoutSpacing(C.QtObjectPtr(ptr.Pointer()), C.int(control1), C.int(control2), C.int(orientation), C.QtObjectPtr(PointerFromQStyleOption(option)), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QProxyStyle) PixelMetric(metric QStyle__PixelMetric, option QStyleOptionITF, widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_PixelMetric(C.QtObjectPtr(ptr.Pointer()), C.int(metric), C.QtObjectPtr(PointerFromQStyleOption(option)), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QProxyStyle) Polish3(app QApplicationITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Polish3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQApplication(app)))
	}
}

func (ptr *QProxyStyle) Polish2(pal gui.QPaletteITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Polish2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPalette(pal)))
	}
}

func (ptr *QProxyStyle) Polish(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Polish(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QProxyStyle) SetBaseStyle(style QStyleITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_SetBaseStyle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStyle(style)))
	}
}

func (ptr *QProxyStyle) StyleHint(hint QStyle__StyleHint, option QStyleOptionITF, widget QWidgetITF, returnData QStyleHintReturnITF) int {
	if ptr.Pointer() != nil {
		return int(C.QProxyStyle_StyleHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint), C.QtObjectPtr(PointerFromQStyleOption(option)), C.QtObjectPtr(PointerFromQWidget(widget)), C.QtObjectPtr(PointerFromQStyleHintReturn(returnData))))
	}
	return 0
}

func (ptr *QProxyStyle) Unpolish2(app QApplicationITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Unpolish2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQApplication(app)))
	}
}

func (ptr *QProxyStyle) Unpolish(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QProxyStyle_Unpolish(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QProxyStyle) DestroyQProxyStyle() {
	if ptr.Pointer() != nil {
		C.QProxyStyle_DestroyQProxyStyle(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
