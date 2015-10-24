package widgets

//#include "qstylepainter.h"
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

type QStylePainterITF interface {
	gui.QPainterITF
	QStylePainterPTR() *QStylePainter
}

func PointerFromQStylePainter(ptr QStylePainterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStylePainterPTR().Pointer()
	}
	return nil
}

func QStylePainterFromPointer(ptr unsafe.Pointer) *QStylePainter {
	var n = new(QStylePainter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStylePainter) QStylePainterPTR() *QStylePainter {
	return ptr
}

func NewQStylePainter() *QStylePainter {
	return QStylePainterFromPointer(unsafe.Pointer(C.QStylePainter_NewQStylePainter()))
}

func NewQStylePainter3(pd gui.QPaintDeviceITF, widget QWidgetITF) *QStylePainter {
	return QStylePainterFromPointer(unsafe.Pointer(C.QStylePainter_NewQStylePainter3(C.QtObjectPtr(gui.PointerFromQPaintDevice(pd)), C.QtObjectPtr(PointerFromQWidget(widget)))))
}

func NewQStylePainter2(widget QWidgetITF) *QStylePainter {
	return QStylePainterFromPointer(unsafe.Pointer(C.QStylePainter_NewQStylePainter2(C.QtObjectPtr(PointerFromQWidget(widget)))))
}

func (ptr *QStylePainter) Begin2(pd gui.QPaintDeviceITF, widget QWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QStylePainter_Begin2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPaintDevice(pd)), C.QtObjectPtr(PointerFromQWidget(widget))) != 0
	}
	return false
}

func (ptr *QStylePainter) Begin(widget QWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QStylePainter_Begin(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))) != 0
	}
	return false
}

func (ptr *QStylePainter) DrawComplexControl(cc QStyle__ComplexControl, option QStyleOptionComplexITF) {
	if ptr.Pointer() != nil {
		C.QStylePainter_DrawComplexControl(C.QtObjectPtr(ptr.Pointer()), C.int(cc), C.QtObjectPtr(PointerFromQStyleOptionComplex(option)))
	}
}

func (ptr *QStylePainter) DrawControl(ce QStyle__ControlElement, option QStyleOptionITF) {
	if ptr.Pointer() != nil {
		C.QStylePainter_DrawControl(C.QtObjectPtr(ptr.Pointer()), C.int(ce), C.QtObjectPtr(PointerFromQStyleOption(option)))
	}
}

func (ptr *QStylePainter) DrawItemPixmap(rect core.QRectITF, flags int, pixmap gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QStylePainter_DrawItemPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(flags), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QStylePainter) DrawItemText(rect core.QRectITF, flags int, pal gui.QPaletteITF, enabled bool, text string, textRole gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QStylePainter_DrawItemText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(flags), C.QtObjectPtr(gui.PointerFromQPalette(pal)), C.int(qt.GoBoolToInt(enabled)), C.CString(text), C.int(textRole))
	}
}

func (ptr *QStylePainter) DrawPrimitive(pe QStyle__PrimitiveElement, option QStyleOptionITF) {
	if ptr.Pointer() != nil {
		C.QStylePainter_DrawPrimitive(C.QtObjectPtr(ptr.Pointer()), C.int(pe), C.QtObjectPtr(PointerFromQStyleOption(option)))
	}
}

func (ptr *QStylePainter) Style() *QStyle {
	if ptr.Pointer() != nil {
		return QStyleFromPointer(unsafe.Pointer(C.QStylePainter_Style(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
