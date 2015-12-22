package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QCommonStyle_") {
		n.SetObjectName("QCommonStyle_" + qt.Identifier())
	}
	return n
}

func (ptr *QCommonStyle) QCommonStyle_PTR() *QCommonStyle {
	return ptr
}

func (ptr *QCommonStyle) ConnectDrawControl(f func(element QStyle__ControlElement, opt *QStyleOption, p *gui.QPainter, widget *QWidget)) {
	defer qt.Recovering("connect QCommonStyle::drawControl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawControl", f)
	}
}

func (ptr *QCommonStyle) DisconnectDrawControl() {
	defer qt.Recovering("disconnect QCommonStyle::drawControl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawControl")
	}
}

//export callbackQCommonStyleDrawControl
func callbackQCommonStyleDrawControl(ptrName *C.char, element C.int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommonStyle::drawControl")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawControl"); signal != nil {
		signal.(func(QStyle__ControlElement, *QStyleOption, *gui.QPainter, *QWidget))(QStyle__ControlElement(element), NewQStyleOptionFromPointer(opt), gui.NewQPainterFromPointer(p), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QCommonStyle) ConnectDrawPrimitive(f func(pe QStyle__PrimitiveElement, opt *QStyleOption, p *gui.QPainter, widget *QWidget)) {
	defer qt.Recovering("connect QCommonStyle::drawPrimitive")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawPrimitive", f)
	}
}

func (ptr *QCommonStyle) DisconnectDrawPrimitive() {
	defer qt.Recovering("disconnect QCommonStyle::drawPrimitive")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawPrimitive")
	}
}

//export callbackQCommonStyleDrawPrimitive
func callbackQCommonStyleDrawPrimitive(ptrName *C.char, pe C.int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommonStyle::drawPrimitive")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawPrimitive"); signal != nil {
		signal.(func(QStyle__PrimitiveElement, *QStyleOption, *gui.QPainter, *QWidget))(QStyle__PrimitiveElement(pe), NewQStyleOptionFromPointer(opt), gui.NewQPainterFromPointer(p), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QCommonStyle) ConnectDrawComplexControl(f func(cc QStyle__ComplexControl, opt *QStyleOptionComplex, p *gui.QPainter, widget *QWidget)) {
	defer qt.Recovering("connect QCommonStyle::drawComplexControl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawComplexControl", f)
	}
}

func (ptr *QCommonStyle) DisconnectDrawComplexControl() {
	defer qt.Recovering("disconnect QCommonStyle::drawComplexControl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawComplexControl")
	}
}

//export callbackQCommonStyleDrawComplexControl
func callbackQCommonStyleDrawComplexControl(ptrName *C.char, cc C.int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommonStyle::drawComplexControl")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawComplexControl"); signal != nil {
		signal.(func(QStyle__ComplexControl, *QStyleOptionComplex, *gui.QPainter, *QWidget))(QStyle__ComplexControl(cc), NewQStyleOptionComplexFromPointer(opt), gui.NewQPainterFromPointer(p), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QCommonStyle) HitTestComplexControl(cc QStyle__ComplexControl, opt QStyleOptionComplex_ITF, pt core.QPoint_ITF, widget QWidget_ITF) QStyle__SubControl {
	defer qt.Recovering("QCommonStyle::hitTestComplexControl")

	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QCommonStyle_HitTestComplexControl(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(opt), core.PointerFromQPoint(pt), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QCommonStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QCommonStyle::layoutSpacing")

	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_LayoutSpacing(ptr.Pointer(), C.int(control1), C.int(control2), C.int(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QCommonStyle) PixelMetric(m QStyle__PixelMetric, opt QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QCommonStyle::pixelMetric")

	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_PixelMetric(ptr.Pointer(), C.int(m), PointerFromQStyleOption(opt), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QCommonStyle) SizeFromContents(ct QStyle__ContentsType, opt QStyleOption_ITF, csz core.QSize_ITF, widget QWidget_ITF) *core.QSize {
	defer qt.Recovering("QCommonStyle::sizeFromContents")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCommonStyle_SizeFromContents(ptr.Pointer(), C.int(ct), PointerFromQStyleOption(opt), core.PointerFromQSize(csz), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QCommonStyle) StyleHint(sh QStyle__StyleHint, opt QStyleOption_ITF, widget QWidget_ITF, hret QStyleHintReturn_ITF) int {
	defer qt.Recovering("QCommonStyle::styleHint")

	if ptr.Pointer() != nil {
		return int(C.QCommonStyle_StyleHint(ptr.Pointer(), C.int(sh), PointerFromQStyleOption(opt), PointerFromQWidget(widget), PointerFromQStyleHintReturn(hret)))
	}
	return 0
}

func (ptr *QCommonStyle) SubControlRect(cc QStyle__ComplexControl, opt QStyleOptionComplex_ITF, sc QStyle__SubControl, widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QCommonStyle::subControlRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QCommonStyle_SubControlRect(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(opt), C.int(sc), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QCommonStyle) SubElementRect(sr QStyle__SubElement, opt QStyleOption_ITF, widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QCommonStyle::subElementRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QCommonStyle_SubElementRect(ptr.Pointer(), C.int(sr), PointerFromQStyleOption(opt), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QCommonStyle) ConnectUnpolish(f func(widget *QWidget)) {
	defer qt.Recovering("connect QCommonStyle::unpolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unpolish", f)
	}
}

func (ptr *QCommonStyle) DisconnectUnpolish() {
	defer qt.Recovering("disconnect QCommonStyle::unpolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unpolish")
	}
}

//export callbackQCommonStyleUnpolish
func callbackQCommonStyleUnpolish(ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommonStyle::unpolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "unpolish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QCommonStyle) DestroyQCommonStyle() {
	defer qt.Recovering("QCommonStyle::~QCommonStyle")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DestroyQCommonStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
