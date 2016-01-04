package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStackedLayout struct {
	QLayout
}

type QStackedLayout_ITF interface {
	QLayout_ITF
	QStackedLayout_PTR() *QStackedLayout
}

func PointerFromQStackedLayout(ptr QStackedLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedLayout_PTR().Pointer()
	}
	return nil
}

func NewQStackedLayoutFromPointer(ptr unsafe.Pointer) *QStackedLayout {
	var n = new(QStackedLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStackedLayout_") {
		n.SetObjectName("QStackedLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QStackedLayout) QStackedLayout_PTR() *QStackedLayout {
	return ptr
}

//QStackedLayout::StackingMode
type QStackedLayout__StackingMode int64

const (
	QStackedLayout__StackOne = QStackedLayout__StackingMode(0)
	QStackedLayout__StackAll = QStackedLayout__StackingMode(1)
)

func (ptr *QStackedLayout) Count() int {
	defer qt.Recovering("QStackedLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) CurrentIndex() int {
	defer qt.Recovering("QStackedLayout::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) SetCurrentIndex(index int) {
	defer qt.Recovering("QStackedLayout::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedLayout) SetCurrentWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStackedLayout::setCurrentWidget")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStackedLayout) SetStackingMode(stackingMode QStackedLayout__StackingMode) {
	defer qt.Recovering("QStackedLayout::setStackingMode")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetStackingMode(ptr.Pointer(), C.int(stackingMode))
	}
}

func (ptr *QStackedLayout) StackingMode() QStackedLayout__StackingMode {
	defer qt.Recovering("QStackedLayout::stackingMode")

	if ptr.Pointer() != nil {
		return QStackedLayout__StackingMode(C.QStackedLayout_StackingMode(ptr.Pointer()))
	}
	return 0
}

func NewQStackedLayout() *QStackedLayout {
	defer qt.Recovering("QStackedLayout::QStackedLayout")

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout())
}

func NewQStackedLayout3(parentLayout QLayout_ITF) *QStackedLayout {
	defer qt.Recovering("QStackedLayout::QStackedLayout")

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout3(PointerFromQLayout(parentLayout)))
}

func NewQStackedLayout2(parent QWidget_ITF) *QStackedLayout {
	defer qt.Recovering("QStackedLayout::QStackedLayout")

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout2(PointerFromQWidget(parent)))
}

func (ptr *QStackedLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	defer qt.Recovering("connect QStackedLayout::addItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "addItem", f)
	}
}

func (ptr *QStackedLayout) DisconnectAddItem() {
	defer qt.Recovering("disconnect QStackedLayout::addItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "addItem")
	}
}

//export callbackQStackedLayoutAddItem
func callbackQStackedLayoutAddItem(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QStackedLayout::addItem")

	if signal := qt.GetSignal(C.GoString(ptrName), "addItem"); signal != nil {
		signal.(func(*QLayoutItem))(NewQLayoutItemFromPointer(item))
	} else {
		NewQStackedLayoutFromPointer(ptr).AddItemDefault(NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QStackedLayout) AddItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QStackedLayout::addItem")

	if ptr.Pointer() != nil {
		C.QStackedLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QStackedLayout) AddItemDefault(item QLayoutItem_ITF) {
	defer qt.Recovering("QStackedLayout::addItem")

	if ptr.Pointer() != nil {
		C.QStackedLayout_AddItemDefault(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QStackedLayout) AddWidget(widget QWidget_ITF) int {
	defer qt.Recovering("QStackedLayout::addWidget")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ConnectCurrentChanged(f func(index int)) {
	defer qt.Recovering("connect QStackedLayout::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedLayout) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QStackedLayout::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedLayoutCurrentChanged
func callbackQStackedLayoutCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QStackedLayout::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QStackedLayout) CurrentChanged(index int) {
	defer qt.Recovering("QStackedLayout::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedLayout_CurrentChanged(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedLayout) CurrentWidget() *QWidget {
	defer qt.Recovering("QStackedLayout::currentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedLayout) HasHeightForWidth() bool {
	defer qt.Recovering("QStackedLayout::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QStackedLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStackedLayout) HeightForWidth(width int) int {
	defer qt.Recovering("QStackedLayout::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QStackedLayout) InsertWidget(index int, widget QWidget_ITF) int {
	defer qt.Recovering("QStackedLayout::insertWidget")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ItemAt(index int) *QLayoutItem {
	defer qt.Recovering("QStackedLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) MinimumSize() *core.QSize {
	defer qt.Recovering("QStackedLayout::minimumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QStackedLayout_MinimumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedLayout) ConnectSetGeometry(f func(rect *core.QRect)) {
	defer qt.Recovering("connect QStackedLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setGeometry", f)
	}
}

func (ptr *QStackedLayout) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QStackedLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setGeometry")
	}
}

//export callbackQStackedLayoutSetGeometry
func callbackQStackedLayoutSetGeometry(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer) {
	defer qt.Recovering("callback QStackedLayout::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(rect))
	} else {
		NewQStackedLayoutFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(rect))
	}
}

func (ptr *QStackedLayout) SetGeometry(rect core.QRect_ITF) {
	defer qt.Recovering("QStackedLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QStackedLayout) SetGeometryDefault(rect core.QRect_ITF) {
	defer qt.Recovering("QStackedLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QStackedLayout) SizeHint() *core.QSize {
	defer qt.Recovering("QStackedLayout::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QStackedLayout_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedLayout) TakeAt(index int) *QLayoutItem {
	defer qt.Recovering("QStackedLayout::takeAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) Widget(index int) *QWidget {
	defer qt.Recovering("QStackedLayout::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) ConnectWidgetRemoved(f func(index int)) {
	defer qt.Recovering("connect QStackedLayout::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedLayout) DisconnectWidgetRemoved() {
	defer qt.Recovering("disconnect QStackedLayout::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedLayoutWidgetRemoved
func callbackQStackedLayoutWidgetRemoved(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QStackedLayout::widgetRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "widgetRemoved"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QStackedLayout) WidgetRemoved(index int) {
	defer qt.Recovering("QStackedLayout::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedLayout_WidgetRemoved(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedLayout) DestroyQStackedLayout() {
	defer qt.Recovering("QStackedLayout::~QStackedLayout")

	if ptr.Pointer() != nil {
		C.QStackedLayout_DestroyQStackedLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStackedLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	defer qt.Recovering("connect QStackedLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStackedLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStackedLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStackedLayoutChildEvent
func callbackQStackedLayoutChildEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QStackedLayout::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(e))
	} else {
		NewQStackedLayoutFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QStackedLayout) ChildEvent(e core.QChildEvent_ITF) {
	defer qt.Recovering("QStackedLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QStackedLayout_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QStackedLayout) ChildEventDefault(e core.QChildEvent_ITF) {
	defer qt.Recovering("QStackedLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QStackedLayout_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QStackedLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QStackedLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QStackedLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QStackedLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQStackedLayoutInvalidate
func callbackQStackedLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QStackedLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQStackedLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QStackedLayout) Invalidate() {
	defer qt.Recovering("QStackedLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QStackedLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QStackedLayout) InvalidateDefault() {
	defer qt.Recovering("QStackedLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QStackedLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QStackedLayout) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStackedLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStackedLayout) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStackedLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStackedLayoutTimerEvent
func callbackQStackedLayoutTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedLayout::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQStackedLayoutFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStackedLayout) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStackedLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QStackedLayout_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStackedLayout) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStackedLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QStackedLayout_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStackedLayout) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStackedLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStackedLayout) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStackedLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStackedLayoutCustomEvent
func callbackQStackedLayoutCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedLayout::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStackedLayoutFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStackedLayout) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QStackedLayout_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStackedLayout) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QStackedLayout_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
