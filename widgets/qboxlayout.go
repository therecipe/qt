package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBoxLayout struct {
	QLayout
}

type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func PointerFromQBoxLayout(ptr QBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQBoxLayoutFromPointer(ptr unsafe.Pointer) *QBoxLayout {
	var n = new(QBoxLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBoxLayout_") {
		n.SetObjectName("QBoxLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout {
	return ptr
}

//QBoxLayout::Direction
type QBoxLayout__Direction int64

const (
	QBoxLayout__LeftToRight = QBoxLayout__Direction(0)
	QBoxLayout__RightToLeft = QBoxLayout__Direction(1)
	QBoxLayout__TopToBottom = QBoxLayout__Direction(2)
	QBoxLayout__BottomToTop = QBoxLayout__Direction(3)
	QBoxLayout__Down        = QBoxLayout__Direction(QBoxLayout__TopToBottom)
	QBoxLayout__Up          = QBoxLayout__Direction(QBoxLayout__BottomToTop)
)

func (ptr *QBoxLayout) Direction() QBoxLayout__Direction {
	defer qt.Recovering("QBoxLayout::direction")

	if ptr.Pointer() != nil {
		return QBoxLayout__Direction(C.QBoxLayout_Direction(ptr.Pointer()))
	}
	return 0
}

func NewQBoxLayout(dir QBoxLayout__Direction, parent QWidget_ITF) *QBoxLayout {
	defer qt.Recovering("QBoxLayout::QBoxLayout")

	return NewQBoxLayoutFromPointer(C.QBoxLayout_NewQBoxLayout(C.int(dir), PointerFromQWidget(parent)))
}

func (ptr *QBoxLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	defer qt.Recovering("connect QBoxLayout::addItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "addItem", f)
	}
}

func (ptr *QBoxLayout) DisconnectAddItem() {
	defer qt.Recovering("disconnect QBoxLayout::addItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "addItem")
	}
}

//export callbackQBoxLayoutAddItem
func callbackQBoxLayoutAddItem(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QBoxLayout::addItem")

	if signal := qt.GetSignal(C.GoString(ptrName), "addItem"); signal != nil {
		signal.(func(*QLayoutItem))(NewQLayoutItemFromPointer(item))
	} else {
		NewQBoxLayoutFromPointer(ptr).AddItemDefault(NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QBoxLayout) AddItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QBoxLayout::addItem")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddItemDefault(item QLayoutItem_ITF) {
	defer qt.Recovering("QBoxLayout::addItem")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItemDefault(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddLayout(layout QLayout_ITF, stretch int) {
	defer qt.Recovering("QBoxLayout::addLayout")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddLayout(ptr.Pointer(), PointerFromQLayout(layout), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddSpacerItem(spacerItem QSpacerItem_ITF) {
	defer qt.Recovering("QBoxLayout::addSpacerItem")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacerItem(ptr.Pointer(), PointerFromQSpacerItem(spacerItem))
	}
}

func (ptr *QBoxLayout) AddSpacing(size int) {
	defer qt.Recovering("QBoxLayout::addSpacing")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddSpacing(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBoxLayout) AddStretch(stretch int) {
	defer qt.Recovering("QBoxLayout::addStretch")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStretch(ptr.Pointer(), C.int(stretch))
	}
}

func (ptr *QBoxLayout) AddStrut(size int) {
	defer qt.Recovering("QBoxLayout::addStrut")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddStrut(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QBoxLayout) AddWidget(widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QBoxLayout::addWidget")

	if ptr.Pointer() != nil {
		C.QBoxLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) Count() int {
	defer qt.Recovering("QBoxLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QBoxLayout::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QBoxLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) HasHeightForWidth() bool {
	defer qt.Recovering("QBoxLayout::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QBoxLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBoxLayout) HeightForWidth(w int) int {
	defer qt.Recovering("QBoxLayout::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) InsertItem(index int, item QLayoutItem_ITF) {
	defer qt.Recovering("QBoxLayout::insertItem")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertItem(ptr.Pointer(), C.int(index), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) InsertLayout(index int, layout QLayout_ITF, stretch int) {
	defer qt.Recovering("QBoxLayout::insertLayout")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertLayout(ptr.Pointer(), C.int(index), PointerFromQLayout(layout), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertSpacerItem(index int, spacerItem QSpacerItem_ITF) {
	defer qt.Recovering("QBoxLayout::insertSpacerItem")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacerItem(ptr.Pointer(), C.int(index), PointerFromQSpacerItem(spacerItem))
	}
}

func (ptr *QBoxLayout) InsertSpacing(index int, size int) {
	defer qt.Recovering("QBoxLayout::insertSpacing")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertSpacing(ptr.Pointer(), C.int(index), C.int(size))
	}
}

func (ptr *QBoxLayout) InsertStretch(index int, stretch int) {
	defer qt.Recovering("QBoxLayout::insertStretch")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) InsertWidget(index int, widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QBoxLayout::insertWidget")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch), C.int(alignment))
	}
}

func (ptr *QBoxLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QBoxLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QBoxLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QBoxLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQBoxLayoutInvalidate
func callbackQBoxLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBoxLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQBoxLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QBoxLayout) Invalidate() {
	defer qt.Recovering("QBoxLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QBoxLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QBoxLayout) InvalidateDefault() {
	defer qt.Recovering("QBoxLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QBoxLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QBoxLayout) ItemAt(index int) *QLayoutItem {
	defer qt.Recovering("QBoxLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QBoxLayout) MaximumSize() *core.QSize {
	defer qt.Recovering("QBoxLayout::maximumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QBoxLayout_MaximumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxLayout) MinimumHeightForWidth(w int) int {
	defer qt.Recovering("QBoxLayout::minimumHeightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QBoxLayout) MinimumSize() *core.QSize {
	defer qt.Recovering("QBoxLayout::minimumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QBoxLayout_MinimumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxLayout) SetDirection(direction QBoxLayout__Direction) {
	defer qt.Recovering("QBoxLayout::setDirection")

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QBoxLayout) ConnectSetGeometry(f func(r *core.QRect)) {
	defer qt.Recovering("connect QBoxLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setGeometry", f)
	}
}

func (ptr *QBoxLayout) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QBoxLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setGeometry")
	}
}

//export callbackQBoxLayoutSetGeometry
func callbackQBoxLayoutSetGeometry(ptr unsafe.Pointer, ptrName *C.char, r unsafe.Pointer) {
	defer qt.Recovering("callback QBoxLayout::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(r))
	} else {
		NewQBoxLayoutFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QBoxLayout) SetGeometry(r core.QRect_ITF) {
	defer qt.Recovering("QBoxLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QBoxLayout) SetGeometryDefault(r core.QRect_ITF) {
	defer qt.Recovering("QBoxLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QBoxLayout) SetSpacing(spacing int) {
	defer qt.Recovering("QBoxLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QBoxLayout) SetStretch(index int, stretch int) {
	defer qt.Recovering("QBoxLayout::setStretch")

	if ptr.Pointer() != nil {
		C.QBoxLayout_SetStretch(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QBoxLayout) SetStretchFactor2(layout QLayout_ITF, stretch int) bool {
	defer qt.Recovering("QBoxLayout::setStretchFactor")

	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor2(ptr.Pointer(), PointerFromQLayout(layout), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) SetStretchFactor(widget QWidget_ITF, stretch int) bool {
	defer qt.Recovering("QBoxLayout::setStretchFactor")

	if ptr.Pointer() != nil {
		return C.QBoxLayout_SetStretchFactor(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch)) != 0
	}
	return false
}

func (ptr *QBoxLayout) SizeHint() *core.QSize {
	defer qt.Recovering("QBoxLayout::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QBoxLayout_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBoxLayout) Spacing() int {
	defer qt.Recovering("QBoxLayout::spacing")

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBoxLayout) Stretch(index int) int {
	defer qt.Recovering("QBoxLayout::stretch")

	if ptr.Pointer() != nil {
		return int(C.QBoxLayout_Stretch(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QBoxLayout) TakeAt(index int) *QLayoutItem {
	defer qt.Recovering("QBoxLayout::takeAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QBoxLayout) DestroyQBoxLayout() {
	defer qt.Recovering("QBoxLayout::~QBoxLayout")

	if ptr.Pointer() != nil {
		C.QBoxLayout_DestroyQBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	defer qt.Recovering("connect QBoxLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBoxLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBoxLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBoxLayoutChildEvent
func callbackQBoxLayoutChildEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QBoxLayout::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(e))
	} else {
		NewQBoxLayoutFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QBoxLayout) ChildEvent(e core.QChildEvent_ITF) {
	defer qt.Recovering("QBoxLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QBoxLayout_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QBoxLayout) ChildEventDefault(e core.QChildEvent_ITF) {
	defer qt.Recovering("QBoxLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QBoxLayout_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QBoxLayout) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBoxLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBoxLayout) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBoxLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBoxLayoutTimerEvent
func callbackQBoxLayoutTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBoxLayout::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBoxLayoutFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBoxLayout) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBoxLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QBoxLayout_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBoxLayout) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBoxLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QBoxLayout_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBoxLayout) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBoxLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBoxLayout) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBoxLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBoxLayoutCustomEvent
func callbackQBoxLayoutCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBoxLayout::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBoxLayoutFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBoxLayout) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBoxLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QBoxLayout_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBoxLayout) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBoxLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QBoxLayout_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
