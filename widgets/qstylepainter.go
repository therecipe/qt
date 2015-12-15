package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QStylePainter struct {
	gui.QPainter
}

type QStylePainter_ITF interface {
	gui.QPainter_ITF
	QStylePainter_PTR() *QStylePainter
}

func PointerFromQStylePainter(ptr QStylePainter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStylePainter_PTR().Pointer()
	}
	return nil
}

func NewQStylePainterFromPointer(ptr unsafe.Pointer) *QStylePainter {
	var n = new(QStylePainter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStylePainter) QStylePainter_PTR() *QStylePainter {
	return ptr
}

func NewQStylePainter() *QStylePainter {
	defer qt.Recovering("QStylePainter::QStylePainter")

	return NewQStylePainterFromPointer(C.QStylePainter_NewQStylePainter())
}

func NewQStylePainter3(pd gui.QPaintDevice_ITF, widget QWidget_ITF) *QStylePainter {
	defer qt.Recovering("QStylePainter::QStylePainter")

	return NewQStylePainterFromPointer(C.QStylePainter_NewQStylePainter3(gui.PointerFromQPaintDevice(pd), PointerFromQWidget(widget)))
}

func NewQStylePainter2(widget QWidget_ITF) *QStylePainter {
	defer qt.Recovering("QStylePainter::QStylePainter")

	return NewQStylePainterFromPointer(C.QStylePainter_NewQStylePainter2(PointerFromQWidget(widget)))
}

func (ptr *QStylePainter) Begin2(pd gui.QPaintDevice_ITF, widget QWidget_ITF) bool {
	defer qt.Recovering("QStylePainter::begin")

	if ptr.Pointer() != nil {
		return C.QStylePainter_Begin2(ptr.Pointer(), gui.PointerFromQPaintDevice(pd), PointerFromQWidget(widget)) != 0
	}
	return false
}

func (ptr *QStylePainter) Begin(widget QWidget_ITF) bool {
	defer qt.Recovering("QStylePainter::begin")

	if ptr.Pointer() != nil {
		return C.QStylePainter_Begin(ptr.Pointer(), PointerFromQWidget(widget)) != 0
	}
	return false
}

func (ptr *QStylePainter) DrawComplexControl(cc QStyle__ComplexControl, option QStyleOptionComplex_ITF) {
	defer qt.Recovering("QStylePainter::drawComplexControl")

	if ptr.Pointer() != nil {
		C.QStylePainter_DrawComplexControl(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(option))
	}
}

func (ptr *QStylePainter) DrawControl(ce QStyle__ControlElement, option QStyleOption_ITF) {
	defer qt.Recovering("QStylePainter::drawControl")

	if ptr.Pointer() != nil {
		C.QStylePainter_DrawControl(ptr.Pointer(), C.int(ce), PointerFromQStyleOption(option))
	}
}

func (ptr *QStylePainter) DrawItemPixmap(rect core.QRect_ITF, flags int, pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QStylePainter::drawItemPixmap")

	if ptr.Pointer() != nil {
		C.QStylePainter_DrawItemPixmap(ptr.Pointer(), core.PointerFromQRect(rect), C.int(flags), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QStylePainter) DrawItemText(rect core.QRect_ITF, flags int, pal gui.QPalette_ITF, enabled bool, text string, textRole gui.QPalette__ColorRole) {
	defer qt.Recovering("QStylePainter::drawItemText")

	if ptr.Pointer() != nil {
		C.QStylePainter_DrawItemText(ptr.Pointer(), core.PointerFromQRect(rect), C.int(flags), gui.PointerFromQPalette(pal), C.int(qt.GoBoolToInt(enabled)), C.CString(text), C.int(textRole))
	}
}

func (ptr *QStylePainter) DrawPrimitive(pe QStyle__PrimitiveElement, option QStyleOption_ITF) {
	defer qt.Recovering("QStylePainter::drawPrimitive")

	if ptr.Pointer() != nil {
		C.QStylePainter_DrawPrimitive(ptr.Pointer(), C.int(pe), PointerFromQStyleOption(option))
	}
}

func (ptr *QStylePainter) Style() *QStyle {
	defer qt.Recovering("QStylePainter::style")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QStylePainter_Style(ptr.Pointer()))
	}
	return nil
}
