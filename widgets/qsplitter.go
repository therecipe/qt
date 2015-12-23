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
func callbackQSplitterChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSplitterChildEvent(ptrName *C.char, c unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(c))
		return true
	}
	return false

}

func (ptr *QSplitter) Count() int {
	defer qt.Recovering("QSplitter::count")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_Count(ptr.Pointer()))
	}
	return 0
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
func callbackQSplitterResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

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
func callbackQSplitterSplitterMoved(ptrName *C.char, pos C.int, index C.int) {
	defer qt.Recovering("callback QSplitter::splitterMoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "splitterMoved"); signal != nil {
		signal.(func(int, int))(int(pos), int(index))
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
func callbackQSplitterPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

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
func callbackQSplitterActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSplitter::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQSplitterShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQSplitterInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
