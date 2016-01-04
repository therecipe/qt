package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLayout struct {
	QLayoutItem
	core.QObject
}

type QLayout_ITF interface {
	QLayoutItem_ITF
	core.QObject_ITF
	QLayout_PTR() *QLayout
}

func (p *QLayout) Pointer() unsafe.Pointer {
	return p.QLayoutItem_PTR().Pointer()
}

func (p *QLayout) SetPointer(ptr unsafe.Pointer) {
	p.QLayoutItem_PTR().SetPointer(ptr)
	p.QObject_PTR().SetPointer(ptr)
}

func PointerFromQLayout(ptr QLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func NewQLayoutFromPointer(ptr unsafe.Pointer) *QLayout {
	var n = new(QLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLayout_") {
		n.SetObjectName("QLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QLayout) QLayout_PTR() *QLayout {
	return ptr
}

//QLayout::SizeConstraint
type QLayout__SizeConstraint int64

const (
	QLayout__SetDefaultConstraint = QLayout__SizeConstraint(0)
	QLayout__SetNoConstraint      = QLayout__SizeConstraint(1)
	QLayout__SetMinimumSize       = QLayout__SizeConstraint(2)
	QLayout__SetFixedSize         = QLayout__SizeConstraint(3)
	QLayout__SetMaximumSize       = QLayout__SizeConstraint(4)
	QLayout__SetMinAndMaxSize     = QLayout__SizeConstraint(5)
)

func (ptr *QLayout) SetSizeConstraint(v QLayout__SizeConstraint) {
	defer qt.Recovering("QLayout::setSizeConstraint")

	if ptr.Pointer() != nil {
		C.QLayout_SetSizeConstraint(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLayout) SetSpacing(v int) {
	defer qt.Recovering("QLayout::setSpacing")

	if ptr.Pointer() != nil {
		C.QLayout_SetSpacing(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLayout) SizeConstraint() QLayout__SizeConstraint {
	defer qt.Recovering("QLayout::sizeConstraint")

	if ptr.Pointer() != nil {
		return QLayout__SizeConstraint(C.QLayout_SizeConstraint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Spacing() int {
	defer qt.Recovering("QLayout::spacing")

	if ptr.Pointer() != nil {
		return int(C.QLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Activate() bool {
	defer qt.Recovering("QLayout::activate")

	if ptr.Pointer() != nil {
		return C.QLayout_Activate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) AddItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QLayout::addItem")

	if ptr.Pointer() != nil {
		C.QLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) AddWidget(w QWidget_ITF) {
	defer qt.Recovering("QLayout::addWidget")

	if ptr.Pointer() != nil {
		C.QLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	defer qt.Recovering("connect QLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLayoutChildEvent
func callbackQLayoutChildEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QLayout::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(e))
	} else {
		NewQLayoutFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QLayout) ChildEvent(e core.QChildEvent_ITF) {
	defer qt.Recovering("QLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QLayout_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QLayout) ChildEventDefault(e core.QChildEvent_ITF) {
	defer qt.Recovering("QLayout::childEvent")

	if ptr.Pointer() != nil {
		C.QLayout_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func QLayout_ClosestAcceptableSize(widget QWidget_ITF, size core.QSize_ITF) *core.QSize {
	defer qt.Recovering("QLayout::closestAcceptableSize")

	return core.NewQSizeFromPointer(C.QLayout_QLayout_ClosestAcceptableSize(PointerFromQWidget(widget), core.PointerFromQSize(size)))
}

func (ptr *QLayout) ContentsRect() *core.QRect {
	defer qt.Recovering("QLayout::contentsRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QLayout_ContentsRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) ControlTypes() QSizePolicy__ControlType {
	defer qt.Recovering("QLayout::controlTypes")

	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayout_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Count() int {
	defer qt.Recovering("QLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QLayout::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Geometry() *core.QRect {
	defer qt.Recovering("QLayout::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QLayout_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) GetContentsMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QLayout::getContentsMargins")

	if ptr.Pointer() != nil {
		C.QLayout_GetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) IndexOf(widget QWidget_ITF) int {
	defer qt.Recovering("QLayout::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QLayout_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQLayoutInvalidate
func callbackQLayoutInvalidate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QLayout) Invalidate() {
	defer qt.Recovering("QLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QLayout) InvalidateDefault() {
	defer qt.Recovering("QLayout::invalidate")

	if ptr.Pointer() != nil {
		C.QLayout_InvalidateDefault(ptr.Pointer())
	}
}

func (ptr *QLayout) IsEmpty() bool {
	defer qt.Recovering("QLayout::isEmpty")

	if ptr.Pointer() != nil {
		return C.QLayout_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) IsEnabled() bool {
	defer qt.Recovering("QLayout::isEnabled")

	if ptr.Pointer() != nil {
		return C.QLayout_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) ItemAt(index int) *QLayoutItem {
	defer qt.Recovering("QLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QLayout) Layout() *QLayout {
	defer qt.Recovering("QLayout::layout")

	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QLayout_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) MaximumSize() *core.QSize {
	defer qt.Recovering("QLayout::maximumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLayout_MaximumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) MenuBar() *QWidget {
	defer qt.Recovering("QLayout::menuBar")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayout_MenuBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) MinimumSize() *core.QSize {
	defer qt.Recovering("QLayout::minimumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLayout_MinimumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) ParentWidget() *QWidget {
	defer qt.Recovering("QLayout::parentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayout_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) RemoveItem(item QLayoutItem_ITF) {
	defer qt.Recovering("QLayout::removeItem")

	if ptr.Pointer() != nil {
		C.QLayout_RemoveItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) RemoveWidget(widget QWidget_ITF) {
	defer qt.Recovering("QLayout::removeWidget")

	if ptr.Pointer() != nil {
		C.QLayout_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QLayout) ReplaceWidget(from QWidget_ITF, to QWidget_ITF, options core.Qt__FindChildOption) *QLayoutItem {
	defer qt.Recovering("QLayout::replaceWidget")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ReplaceWidget(ptr.Pointer(), PointerFromQWidget(from), PointerFromQWidget(to), C.int(options)))
	}
	return nil
}

func (ptr *QLayout) SetAlignment2(l QLayout_ITF, alignment core.Qt__AlignmentFlag) bool {
	defer qt.Recovering("QLayout::setAlignment")

	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment2(ptr.Pointer(), PointerFromQLayout(l), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetAlignment(w QWidget_ITF, alignment core.Qt__AlignmentFlag) bool {
	defer qt.Recovering("QLayout::setAlignment")

	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment(ptr.Pointer(), PointerFromQWidget(w), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetContentsMargins2(margins core.QMargins_ITF) {
	defer qt.Recovering("QLayout::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QLayout::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) SetEnabled(enable bool) {
	defer qt.Recovering("QLayout::setEnabled")

	if ptr.Pointer() != nil {
		C.QLayout_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLayout) ConnectSetGeometry(f func(r *core.QRect)) {
	defer qt.Recovering("connect QLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setGeometry", f)
	}
}

func (ptr *QLayout) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setGeometry")
	}
}

//export callbackQLayoutSetGeometry
func callbackQLayoutSetGeometry(ptr unsafe.Pointer, ptrName *C.char, r unsafe.Pointer) {
	defer qt.Recovering("callback QLayout::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(r))
	} else {
		NewQLayoutFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QLayout) SetGeometry(r core.QRect_ITF) {
	defer qt.Recovering("QLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayout) SetGeometryDefault(r core.QRect_ITF) {
	defer qt.Recovering("QLayout::setGeometry")

	if ptr.Pointer() != nil {
		C.QLayout_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayout) SetMenuBar(widget QWidget_ITF) {
	defer qt.Recovering("QLayout::setMenuBar")

	if ptr.Pointer() != nil {
		C.QLayout_SetMenuBar(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QLayout) TakeAt(index int) *QLayoutItem {
	defer qt.Recovering("QLayout::takeAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QLayout) Update() {
	defer qt.Recovering("QLayout::update")

	if ptr.Pointer() != nil {
		C.QLayout_Update(ptr.Pointer())
	}
}

func (ptr *QLayout) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLayout) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLayoutTimerEvent
func callbackQLayoutTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLayout::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLayoutFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLayout) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QLayout_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLayout) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLayout::timerEvent")

	if ptr.Pointer() != nil {
		C.QLayout_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLayout) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLayout) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLayoutCustomEvent
func callbackQLayoutCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLayout::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLayoutFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLayout) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QLayout_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLayout) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLayout::customEvent")

	if ptr.Pointer() != nil {
		C.QLayout_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
