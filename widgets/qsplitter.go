package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSplitter struct {
	QFrame
}

type QSplitter_ITF interface {
	QFrame_ITF
	QSplitter_PTR() *QSplitter
}

func PointerFromQSplitter(ptr QSplitter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitter_PTR().Pointer()
	}
	return nil
}

func NewQSplitterFromPointer(ptr unsafe.Pointer) *QSplitter {
	var n = new(QSplitter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplitter_") {
		n.SetObjectName("QSplitter_" + qt.Identifier())
	}
	return n
}

func (ptr *QSplitter) QSplitter_PTR() *QSplitter {
	return ptr
}

func (ptr *QSplitter) ChildrenCollapsible() bool {
	defer qt.Recovering("QSplitter::childrenCollapsible")

	if ptr.Pointer() != nil {
		return C.QSplitter_ChildrenCollapsible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitter) HandleWidth() int {
	defer qt.Recovering("QSplitter::handleWidth")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_HandleWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) IndexOf(widget QWidget_ITF) int {
	defer qt.Recovering("QSplitter::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QSplitter) OpaqueResize() bool {
	defer qt.Recovering("QSplitter::opaqueResize")

	if ptr.Pointer() != nil {
		return C.QSplitter_OpaqueResize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitter) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QSplitter::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitter_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) SetChildrenCollapsible(v bool) {
	defer qt.Recovering("QSplitter::setChildrenCollapsible")

	if ptr.Pointer() != nil {
		C.QSplitter_SetChildrenCollapsible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QSplitter) SetHandleWidth(v int) {
	defer qt.Recovering("QSplitter::setHandleWidth")

	if ptr.Pointer() != nil {
		C.QSplitter_SetHandleWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QSplitter) SetOpaqueResize(opaque bool) {
	defer qt.Recovering("QSplitter::setOpaqueResize")

	if ptr.Pointer() != nil {
		C.QSplitter_SetOpaqueResize(ptr.Pointer(), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QSplitter) SetOrientation(v core.Qt__Orientation) {
	defer qt.Recovering("QSplitter::setOrientation")

	if ptr.Pointer() != nil {
		C.QSplitter_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func NewQSplitter(parent QWidget_ITF) *QSplitter {
	defer qt.Recovering("QSplitter::QSplitter")

	return NewQSplitterFromPointer(C.QSplitter_NewQSplitter(PointerFromQWidget(parent)))
}

func NewQSplitter2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSplitter {
	defer qt.Recovering("QSplitter::QSplitter")

	return NewQSplitterFromPointer(C.QSplitter_NewQSplitter2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QSplitter) AddWidget(widget QWidget_ITF) {
	defer qt.Recovering("QSplitter::addWidget")

	if ptr.Pointer() != nil {
		C.QSplitter_AddWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QSplitter::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSplitter) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSplitter::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSplitterChangeEvent
func callbackQSplitterChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQSplitterFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QSplitter) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::changeEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QSplitter) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::changeEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QSplitter) ConnectChildEvent(f func(c *core.QChildEvent)) {
	defer qt.Recovering("connect QSplitter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSplitter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSplitter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSplitterChildEvent
func callbackQSplitterChildEvent(ptr unsafe.Pointer, ptrName *C.char, c unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(c))
	} else {
		NewQSplitterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(c))
	}
}

func (ptr *QSplitter) ChildEvent(c core.QChildEvent_ITF) {
	defer qt.Recovering("QSplitter::childEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(c))
	}
}

func (ptr *QSplitter) ChildEventDefault(c core.QChildEvent_ITF) {
	defer qt.Recovering("QSplitter::childEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(c))
	}
}

func (ptr *QSplitter) Count() int {
	defer qt.Recovering("QSplitter::count")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) CreateHandle() *QSplitterHandle {
	defer qt.Recovering("QSplitter::createHandle")

	if ptr.Pointer() != nil {
		return NewQSplitterHandleFromPointer(C.QSplitter_CreateHandle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSplitter::event")

	if ptr.Pointer() != nil {
		return C.QSplitter_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSplitter) GetRange(index int, min int, max int) {
	defer qt.Recovering("QSplitter::getRange")

	if ptr.Pointer() != nil {
		C.QSplitter_GetRange(ptr.Pointer(), C.int(index), C.int(min), C.int(max))
	}
}

func (ptr *QSplitter) Handle(index int) *QSplitterHandle {
	defer qt.Recovering("QSplitter::handle")

	if ptr.Pointer() != nil {
		return NewQSplitterHandleFromPointer(C.QSplitter_Handle(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSplitter) InsertWidget(index int, widget QWidget_ITF) {
	defer qt.Recovering("QSplitter::insertWidget")

	if ptr.Pointer() != nil {
		C.QSplitter_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) IsCollapsible(index int) bool {
	defer qt.Recovering("QSplitter::isCollapsible")

	if ptr.Pointer() != nil {
		return C.QSplitter_IsCollapsible(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSplitter) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QSplitter::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSplitter_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) Refresh() {
	defer qt.Recovering("QSplitter::refresh")

	if ptr.Pointer() != nil {
		C.QSplitter_Refresh(ptr.Pointer())
	}
}

func (ptr *QSplitter) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSplitter::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSplitter) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSplitter::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSplitterResizeEvent
func callbackQSplitterResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQSplitterFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QSplitter) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSplitter::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QSplitter) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSplitter::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QSplitter) RestoreState(state core.QByteArray_ITF) bool {
	defer qt.Recovering("QSplitter::restoreState")

	if ptr.Pointer() != nil {
		return C.QSplitter_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QSplitter) SaveState() *core.QByteArray {
	defer qt.Recovering("QSplitter::saveState")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSplitter_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) SetCollapsible(index int, collapse bool) {
	defer qt.Recovering("QSplitter::setCollapsible")

	if ptr.Pointer() != nil {
		C.QSplitter_SetCollapsible(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QSplitter) SetStretchFactor(index int, stretch int) {
	defer qt.Recovering("QSplitter::setStretchFactor")

	if ptr.Pointer() != nil {
		C.QSplitter_SetStretchFactor(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QSplitter) SizeHint() *core.QSize {
	defer qt.Recovering("QSplitter::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSplitter_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) ConnectSplitterMoved(f func(pos int, index int)) {
	defer qt.Recovering("connect QSplitter::splitterMoved")

	if ptr.Pointer() != nil {
		C.QSplitter_ConnectSplitterMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "splitterMoved", f)
	}
}

func (ptr *QSplitter) DisconnectSplitterMoved() {
	defer qt.Recovering("disconnect QSplitter::splitterMoved")

	if ptr.Pointer() != nil {
		C.QSplitter_DisconnectSplitterMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "splitterMoved")
	}
}

//export callbackQSplitterSplitterMoved
func callbackQSplitterSplitterMoved(ptr unsafe.Pointer, ptrName *C.char, pos C.int, index C.int) {
	defer qt.Recovering("callback QSplitter::splitterMoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "splitterMoved"); signal != nil {
		signal.(func(int, int))(int(pos), int(index))
	}

}

func (ptr *QSplitter) SplitterMoved(pos int, index int) {
	defer qt.Recovering("QSplitter::splitterMoved")

	if ptr.Pointer() != nil {
		C.QSplitter_SplitterMoved(ptr.Pointer(), C.int(pos), C.int(index))
	}
}

func (ptr *QSplitter) Widget(index int) *QWidget {
	defer qt.Recovering("QSplitter::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QSplitter_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSplitter) DestroyQSplitter() {
	defer qt.Recovering("QSplitter::~QSplitter")

	if ptr.Pointer() != nil {
		C.QSplitter_DestroyQSplitter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSplitter) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSplitter::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSplitter) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSplitter::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSplitterPaintEvent
func callbackQSplitterPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQSplitterFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QSplitter) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSplitter::paintEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QSplitter) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSplitter::paintEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QSplitter) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSplitter::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSplitter) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSplitter::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSplitterActionEvent
func callbackQSplitterActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSplitter) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSplitter::actionEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSplitter) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSplitter::actionEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSplitter) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSplitter::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSplitter) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSplitter::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSplitterDragEnterEvent
func callbackQSplitterDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSplitter) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSplitter::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSplitter) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSplitter::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSplitter) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSplitter::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSplitter) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSplitter::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSplitterDragLeaveEvent
func callbackQSplitterDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSplitter) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSplitter::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSplitter) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSplitter::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSplitter) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSplitter::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSplitter) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSplitter::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSplitterDragMoveEvent
func callbackQSplitterDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSplitter) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSplitter::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSplitter) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSplitter::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSplitter) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSplitter::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSplitter) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSplitter::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSplitterDropEvent
func callbackQSplitterDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSplitter) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSplitter::dropEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSplitter) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSplitter::dropEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSplitter) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitter::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSplitter) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSplitter::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSplitterEnterEvent
func callbackQSplitterEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitter) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::enterEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitter) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::enterEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitter) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSplitter::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSplitter) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSplitter::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSplitterFocusInEvent
func callbackQSplitterFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSplitter) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitter::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitter) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitter::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitter) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSplitter::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSplitter) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSplitter::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSplitterFocusOutEvent
func callbackQSplitterFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSplitter) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitter::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitter) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitter::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitter) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSplitter::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSplitter) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSplitter::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSplitterHideEvent
func callbackQSplitterHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSplitter) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSplitter::hideEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSplitter) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSplitter::hideEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSplitter) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitter::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSplitter) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSplitter::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSplitterLeaveEvent
func callbackQSplitterLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitter) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitter) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitter) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSplitter::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSplitter) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSplitter::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSplitterMoveEvent
func callbackQSplitterMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSplitter) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSplitter::moveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSplitter) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSplitter::moveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSplitter) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSplitter::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSplitter) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSplitter::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSplitterSetVisible
func callbackQSplitterSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSplitter::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSplitter) SetVisible(visible bool) {
	defer qt.Recovering("QSplitter::setVisible")

	if ptr.Pointer() != nil {
		C.QSplitter_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSplitter) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSplitter::setVisible")

	if ptr.Pointer() != nil {
		C.QSplitter_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSplitter) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSplitter::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSplitter) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSplitter::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSplitterShowEvent
func callbackQSplitterShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSplitter) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSplitter::showEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSplitter) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSplitter::showEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSplitter) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSplitter::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSplitter) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSplitter::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSplitterCloseEvent
func callbackQSplitterCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSplitter) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSplitter::closeEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSplitter) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSplitter::closeEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSplitter) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSplitter::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSplitter) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSplitter::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSplitterContextMenuEvent
func callbackQSplitterContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSplitter) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSplitter::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSplitter) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSplitter::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSplitter) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSplitter::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSplitter) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSplitter::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSplitterInitPainter
func callbackQSplitterInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSplitterFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSplitter) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplitter::initPainter")

	if ptr.Pointer() != nil {
		C.QSplitter_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplitter) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplitter::initPainter")

	if ptr.Pointer() != nil {
		C.QSplitter_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplitter) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSplitter::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSplitter) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSplitter::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSplitterInputMethodEvent
func callbackQSplitterInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSplitter) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSplitter::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSplitter) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSplitter::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSplitter) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSplitter::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSplitter) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSplitter::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSplitterKeyPressEvent
func callbackQSplitterKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSplitter) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitter::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitter) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitter::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitter) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSplitter::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSplitter) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSplitter::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSplitterKeyReleaseEvent
func callbackQSplitterKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSplitter) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitter::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitter) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitter::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitter) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitter::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSplitter) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSplitter::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSplitterMouseDoubleClickEvent
func callbackQSplitterMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplitter) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitter::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSplitter) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSplitter::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSplitterMouseMoveEvent
func callbackQSplitterMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplitter) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitter::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSplitter) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSplitter::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSplitterMousePressEvent
func callbackQSplitterMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplitter) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitter::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSplitter) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSplitter::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSplitterMouseReleaseEvent
func callbackQSplitterMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplitter) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitter::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitter) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSplitter::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSplitter) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSplitter::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSplitterTabletEvent
func callbackQSplitterTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSplitter) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSplitter::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSplitter) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSplitter::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSplitter) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSplitter::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSplitter) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSplitter::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSplitterWheelEvent
func callbackQSplitterWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSplitter) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSplitter::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSplitter) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSplitter::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSplitter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSplitter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSplitter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSplitter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSplitterTimerEvent
func callbackQSplitterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSplitter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSplitter::timerEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSplitter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSplitter::timerEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSplitter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSplitter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSplitter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSplitterCustomEvent
func callbackQSplitterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::customEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitter::customEvent")

	if ptr.Pointer() != nil {
		C.QSplitter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
