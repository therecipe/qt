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
func callbackQCommonStyleDrawControl(ptr unsafe.Pointer, ptrName *C.char, element C.int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::drawControl")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawControl"); signal != nil {
		signal.(func(QStyle__ControlElement, *QStyleOption, *gui.QPainter, *QWidget))(QStyle__ControlElement(element), NewQStyleOptionFromPointer(opt), gui.NewQPainterFromPointer(p), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QCommonStyle) DrawControl(element QStyle__ControlElement, opt QStyleOption_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::drawControl")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawControl(ptr.Pointer(), C.int(element), PointerFromQStyleOption(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DrawControlDefault(element QStyle__ControlElement, opt QStyleOption_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::drawControl")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawControlDefault(ptr.Pointer(), C.int(element), PointerFromQStyleOption(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
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
func callbackQCommonStyleDrawPrimitive(ptr unsafe.Pointer, ptrName *C.char, pe C.int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::drawPrimitive")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawPrimitive"); signal != nil {
		signal.(func(QStyle__PrimitiveElement, *QStyleOption, *gui.QPainter, *QWidget))(QStyle__PrimitiveElement(pe), NewQStyleOptionFromPointer(opt), gui.NewQPainterFromPointer(p), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QCommonStyle) DrawPrimitive(pe QStyle__PrimitiveElement, opt QStyleOption_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::drawPrimitive")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawPrimitive(ptr.Pointer(), C.int(pe), PointerFromQStyleOption(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DrawPrimitiveDefault(pe QStyle__PrimitiveElement, opt QStyleOption_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::drawPrimitive")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawPrimitiveDefault(ptr.Pointer(), C.int(pe), PointerFromQStyleOption(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
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
func callbackQCommonStyleDrawComplexControl(ptr unsafe.Pointer, ptrName *C.char, cc C.int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::drawComplexControl")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawComplexControl"); signal != nil {
		signal.(func(QStyle__ComplexControl, *QStyleOptionComplex, *gui.QPainter, *QWidget))(QStyle__ComplexControl(cc), NewQStyleOptionComplexFromPointer(opt), gui.NewQPainterFromPointer(p), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QCommonStyle) DrawComplexControl(cc QStyle__ComplexControl, opt QStyleOptionComplex_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::drawComplexControl")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawComplexControl(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DrawComplexControlDefault(cc QStyle__ComplexControl, opt QStyleOptionComplex_ITF, p gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::drawComplexControl")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DrawComplexControlDefault(ptr.Pointer(), C.int(cc), PointerFromQStyleOptionComplex(opt), gui.PointerFromQPainter(p), PointerFromQWidget(widget))
	}
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
func callbackQCommonStyleUnpolish(ptr unsafe.Pointer, ptrName *C.char, widget unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::unpolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "unpolish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
	} else {
		NewQCommonStyleFromPointer(ptr).UnpolishDefault(NewQWidgetFromPointer(widget))
	}
}

func (ptr *QCommonStyle) Unpolish(widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::unpolish")

	if ptr.Pointer() != nil {
		C.QCommonStyle_Unpolish(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) UnpolishDefault(widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::unpolish")

	if ptr.Pointer() != nil {
		C.QCommonStyle_UnpolishDefault(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) DestroyQCommonStyle() {
	defer qt.Recovering("QCommonStyle::~QCommonStyle")

	if ptr.Pointer() != nil {
		C.QCommonStyle_DestroyQCommonStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCommonStyle) ConnectPolish(f func(widget *QWidget)) {
	defer qt.Recovering("connect QCommonStyle::polish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "polish", f)
	}
}

func (ptr *QCommonStyle) DisconnectPolish() {
	defer qt.Recovering("disconnect QCommonStyle::polish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "polish")
	}
}

//export callbackQCommonStylePolish
func callbackQCommonStylePolish(ptr unsafe.Pointer, ptrName *C.char, widget unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::polish")

	if signal := qt.GetSignal(C.GoString(ptrName), "polish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
	} else {
		NewQCommonStyleFromPointer(ptr).PolishDefault(NewQWidgetFromPointer(widget))
	}
}

func (ptr *QCommonStyle) Polish(widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::polish")

	if ptr.Pointer() != nil {
		C.QCommonStyle_Polish(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) PolishDefault(widget QWidget_ITF) {
	defer qt.Recovering("QCommonStyle::polish")

	if ptr.Pointer() != nil {
		C.QCommonStyle_PolishDefault(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCommonStyle) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCommonStyle::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCommonStyle) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCommonStyle::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCommonStyleTimerEvent
func callbackQCommonStyleTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCommonStyleFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCommonStyle) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCommonStyle::timerEvent")

	if ptr.Pointer() != nil {
		C.QCommonStyle_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCommonStyle) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCommonStyle::timerEvent")

	if ptr.Pointer() != nil {
		C.QCommonStyle_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCommonStyle) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCommonStyle::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCommonStyle) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCommonStyle::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCommonStyleChildEvent
func callbackQCommonStyleChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCommonStyleFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCommonStyle) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCommonStyle::childEvent")

	if ptr.Pointer() != nil {
		C.QCommonStyle_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCommonStyle) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCommonStyle::childEvent")

	if ptr.Pointer() != nil {
		C.QCommonStyle_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCommonStyle) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommonStyle::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCommonStyle) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCommonStyle::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCommonStyleCustomEvent
func callbackQCommonStyleCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCommonStyle::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCommonStyleFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCommonStyle) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCommonStyle::customEvent")

	if ptr.Pointer() != nil {
		C.QCommonStyle_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCommonStyle) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCommonStyle::customEvent")

	if ptr.Pointer() != nil {
		C.QCommonStyle_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
