package multimediawidgets

//#include "multimediawidgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QCameraViewfinder struct {
	QVideoWidget
}

type QCameraViewfinder_ITF interface {
	QVideoWidget_ITF
	QCameraViewfinder_PTR() *QCameraViewfinder
}

func PointerFromQCameraViewfinder(ptr QCameraViewfinder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinder_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderFromPointer(ptr unsafe.Pointer) *QCameraViewfinder {
	var n = new(QCameraViewfinder)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinder_") {
		n.SetObjectName("QCameraViewfinder_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraViewfinder) QCameraViewfinder_PTR() *QCameraViewfinder {
	return ptr
}

func NewQCameraViewfinder(parent widgets.QWidget_ITF) *QCameraViewfinder {
	defer qt.Recovering("QCameraViewfinder::QCameraViewfinder")

	return NewQCameraViewfinderFromPointer(C.QCameraViewfinder_NewQCameraViewfinder(widgets.PointerFromQWidget(parent)))
}

func (ptr *QCameraViewfinder) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QCameraViewfinder::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QCameraViewfinder_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) SetMediaObject(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_SetMediaObject(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) DestroyQCameraViewfinder() {
	defer qt.Recovering("QCameraViewfinder::~QCameraViewfinder")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DestroyQCameraViewfinder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinder) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQCameraViewfinderHideEvent
func callbackQCameraViewfinderHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCameraViewfinder) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQCameraViewfinderMoveEvent
func callbackQCameraViewfinderMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCameraViewfinder) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCameraViewfinderPaintEvent
func callbackQCameraViewfinderPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QCameraViewfinder) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQCameraViewfinderResizeEvent
func callbackQCameraViewfinderResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCameraViewfinder) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQCameraViewfinderShowEvent
func callbackQCameraViewfinderShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCameraViewfinder) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQCameraViewfinderActionEvent
func callbackQCameraViewfinderActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCameraViewfinder) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQCameraViewfinderDragEnterEvent
func callbackQCameraViewfinderDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCameraViewfinder) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQCameraViewfinderDragLeaveEvent
func callbackQCameraViewfinderDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCameraViewfinder) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQCameraViewfinderDragMoveEvent
func callbackQCameraViewfinderDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCameraViewfinder) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQCameraViewfinderDropEvent
func callbackQCameraViewfinderDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCameraViewfinder) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQCameraViewfinderEnterEvent
func callbackQCameraViewfinderEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQCameraViewfinderFocusInEvent
func callbackQCameraViewfinderFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCameraViewfinder) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQCameraViewfinderFocusOutEvent
func callbackQCameraViewfinderFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCameraViewfinder) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQCameraViewfinderLeaveEvent
func callbackQCameraViewfinderLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQCameraViewfinderSetVisible
func callbackQCameraViewfinderSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QCameraViewfinder::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) SetVisible(visible bool) {
	defer qt.Recovering("QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCameraViewfinder) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCameraViewfinder) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQCameraViewfinderChangeEvent
func callbackQCameraViewfinderChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQCameraViewfinderCloseEvent
func callbackQCameraViewfinderCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCameraViewfinder) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQCameraViewfinderContextMenuEvent
func callbackQCameraViewfinderContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCameraViewfinder) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQCameraViewfinderInitPainter
func callbackQCameraViewfinderInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQCameraViewfinderFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QCameraViewfinder) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCameraViewfinder) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCameraViewfinder) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQCameraViewfinderInputMethodEvent
func callbackQCameraViewfinderInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCameraViewfinder) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQCameraViewfinderKeyPressEvent
func callbackQCameraViewfinderKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCameraViewfinder) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQCameraViewfinderKeyReleaseEvent
func callbackQCameraViewfinderKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCameraViewfinder) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQCameraViewfinderMouseDoubleClickEvent
func callbackQCameraViewfinderMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCameraViewfinderMouseMoveEvent
func callbackQCameraViewfinderMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQCameraViewfinderMousePressEvent
func callbackQCameraViewfinderMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQCameraViewfinderMouseReleaseEvent
func callbackQCameraViewfinderMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQCameraViewfinderTabletEvent
func callbackQCameraViewfinderTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCameraViewfinder) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQCameraViewfinderWheelEvent
func callbackQCameraViewfinderWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCameraViewfinder) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraViewfinderTimerEvent
func callbackQCameraViewfinderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinder) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraViewfinderChildEvent
func callbackQCameraViewfinderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinder) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraViewfinderCustomEvent
func callbackQCameraViewfinderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QGraphicsVideoItem struct {
	multimedia.QMediaBindableInterface
	widgets.QGraphicsObject
}

type QGraphicsVideoItem_ITF interface {
	multimedia.QMediaBindableInterface_ITF
	widgets.QGraphicsObject_ITF
	QGraphicsVideoItem_PTR() *QGraphicsVideoItem
}

func (p *QGraphicsVideoItem) Pointer() unsafe.Pointer {
	return p.QMediaBindableInterface_PTR().Pointer()
}

func (p *QGraphicsVideoItem) SetPointer(ptr unsafe.Pointer) {
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
	p.QGraphicsObject_PTR().SetPointer(ptr)
}

func PointerFromQGraphicsVideoItem(ptr QGraphicsVideoItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsVideoItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsVideoItemFromPointer(ptr unsafe.Pointer) *QGraphicsVideoItem {
	var n = new(QGraphicsVideoItem)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsVideoItem_") {
		n.SetObjectName("QGraphicsVideoItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsVideoItem) QGraphicsVideoItem_PTR() *QGraphicsVideoItem {
	return ptr
}

func NewQGraphicsVideoItem(parent widgets.QGraphicsItem_ITF) *QGraphicsVideoItem {
	defer qt.Recovering("QGraphicsVideoItem::QGraphicsVideoItem")

	return NewQGraphicsVideoItemFromPointer(C.QGraphicsVideoItem_NewQGraphicsVideoItem(widgets.PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsVideoItem) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QGraphicsVideoItem::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QGraphicsVideoItem_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsVideoItem) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QGraphicsVideoItem::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QGraphicsVideoItem_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsVideoItemPaint
func callbackQGraphicsVideoItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsVideoItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsVideoItem) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsVideoItem) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsVideoItem::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsVideoItem) SetOffset(offset core.QPointF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetOffset(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QGraphicsVideoItem) SetSize(size core.QSizeF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::setSize")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsVideoItem) DestroyQGraphicsVideoItem() {
	defer qt.Recovering("QGraphicsVideoItem::~QGraphicsVideoItem")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DestroyQGraphicsVideoItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsVideoItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsVideoItemTimerEvent
func callbackQGraphicsVideoItemTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsVideoItemChildEvent
func callbackQGraphicsVideoItemChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsVideoItemCustomEvent
func callbackQGraphicsVideoItemCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoWidget struct {
	multimedia.QMediaBindableInterface
	widgets.QWidget
}

type QVideoWidget_ITF interface {
	multimedia.QMediaBindableInterface_ITF
	widgets.QWidget_ITF
	QVideoWidget_PTR() *QVideoWidget
}

func (p *QVideoWidget) Pointer() unsafe.Pointer {
	return p.QMediaBindableInterface_PTR().Pointer()
}

func (p *QVideoWidget) SetPointer(ptr unsafe.Pointer) {
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
	p.QWidget_PTR().SetPointer(ptr)
}

func PointerFromQVideoWidget(ptr QVideoWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidget_PTR().Pointer()
	}
	return nil
}

func NewQVideoWidgetFromPointer(ptr unsafe.Pointer) *QVideoWidget {
	var n = new(QVideoWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidget_") {
		n.SetObjectName("QVideoWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoWidget) QVideoWidget_PTR() *QVideoWidget {
	return ptr
}

func (ptr *QVideoWidget) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QVideoWidget::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidget_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Brightness() int {
	defer qt.Recovering("QVideoWidget::brightness")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Contrast() int {
	defer qt.Recovering("QVideoWidget::contrast")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Hue() int {
	defer qt.Recovering("QVideoWidget::hue")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QVideoWidget::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QVideoWidget_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) Saturation() int {
	defer qt.Recovering("QVideoWidget::saturation")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWidget::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWidget) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWidget::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidget) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWidget::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidget) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWidget::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidget) SetHue(hue int) {
	defer qt.Recovering("QVideoWidget::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidget) SetSaturation(saturation int) {
	defer qt.Recovering("QVideoWidget::setSaturation")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func NewQVideoWidget(parent widgets.QWidget_ITF) *QVideoWidget {
	defer qt.Recovering("QVideoWidget::QVideoWidget")

	return NewQVideoWidgetFromPointer(C.QVideoWidget_NewQVideoWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QVideoWidget) ConnectBrightnessChanged(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectBrightnessChanged() {
	defer qt.Recovering("disconnect QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetBrightnessChanged
func callbackQVideoWidgetBrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidget::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidget) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidget) ConnectContrastChanged(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectContrastChanged() {
	defer qt.Recovering("disconnect QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetContrastChanged
func callbackQVideoWidgetContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidget::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidget) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidget) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidget::event")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QVideoWidget) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectFullScreenChanged() {
	defer qt.Recovering("disconnect QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetFullScreenChanged
func callbackQVideoWidgetFullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidget::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidget) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQVideoWidgetHideEvent
func callbackQVideoWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QVideoWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectHueChanged(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidget::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectHueChanged() {
	defer qt.Recovering("disconnect QVideoWidget::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetHueChanged
func callbackQVideoWidgetHueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidget::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidget) HueChanged(hue int) {
	defer qt.Recovering("QVideoWidget::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HueChanged(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidget) IsFullScreen() bool {
	defer qt.Recovering("QVideoWidget::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQVideoWidgetMoveEvent
func callbackQVideoWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QVideoWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQVideoWidgetPaintEvent
func callbackQVideoWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QVideoWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQVideoWidgetResizeEvent
func callbackQVideoWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QVideoWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectSaturationChanged(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectSaturationChanged() {
	defer qt.Recovering("disconnect QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetSaturationChanged
func callbackQVideoWidgetSaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidget::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidget) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QVideoWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QVideoWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQVideoWidgetShowEvent
func callbackQVideoWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QVideoWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QVideoWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QVideoWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QVideoWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QVideoWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) DestroyQVideoWidget() {
	defer qt.Recovering("QVideoWidget::~QVideoWidget")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DestroyQVideoWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQVideoWidgetActionEvent
func callbackQVideoWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QVideoWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQVideoWidgetDragEnterEvent
func callbackQVideoWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QVideoWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQVideoWidgetDragLeaveEvent
func callbackQVideoWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QVideoWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQVideoWidgetDragMoveEvent
func callbackQVideoWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QVideoWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQVideoWidgetDropEvent
func callbackQVideoWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QVideoWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQVideoWidgetEnterEvent
func callbackQVideoWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQVideoWidgetFocusInEvent
func callbackQVideoWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQVideoWidgetFocusOutEvent
func callbackQVideoWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQVideoWidgetLeaveEvent
func callbackQVideoWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QVideoWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QVideoWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQVideoWidgetSetVisible
func callbackQVideoWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QVideoWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QVideoWidget) SetVisible(visible bool) {
	defer qt.Recovering("QVideoWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QVideoWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QVideoWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QVideoWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQVideoWidgetChangeEvent
func callbackQVideoWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQVideoWidgetCloseEvent
func callbackQVideoWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QVideoWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQVideoWidgetContextMenuEvent
func callbackQVideoWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QVideoWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QVideoWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QVideoWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QVideoWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQVideoWidgetInitPainter
func callbackQVideoWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQVideoWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QVideoWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QVideoWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QVideoWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QVideoWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QVideoWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQVideoWidgetInputMethodEvent
func callbackQVideoWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QVideoWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQVideoWidgetKeyPressEvent
func callbackQVideoWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQVideoWidgetKeyReleaseEvent
func callbackQVideoWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQVideoWidgetMouseDoubleClickEvent
func callbackQVideoWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQVideoWidgetMouseMoveEvent
func callbackQVideoWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQVideoWidgetMousePressEvent
func callbackQVideoWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQVideoWidgetMouseReleaseEvent
func callbackQVideoWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQVideoWidgetTabletEvent
func callbackQVideoWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QVideoWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQVideoWidgetWheelEvent
func callbackQVideoWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QVideoWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoWidgetTimerEvent
func callbackQVideoWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoWidgetChildEvent
func callbackQVideoWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoWidgetCustomEvent
func callbackQVideoWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QVideoWidgetControl struct {
	multimedia.QMediaControl
}

type QVideoWidgetControl_ITF interface {
	multimedia.QMediaControl_ITF
	QVideoWidgetControl_PTR() *QVideoWidgetControl
}

func PointerFromQVideoWidgetControl(ptr QVideoWidgetControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidgetControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoWidgetControlFromPointer(ptr unsafe.Pointer) *QVideoWidgetControl {
	var n = new(QVideoWidgetControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidgetControl_") {
		n.SetObjectName("QVideoWidgetControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoWidgetControl) QVideoWidgetControl_PTR() *QVideoWidgetControl {
	return ptr
}

func (ptr *QVideoWidgetControl) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QVideoWidgetControl::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidgetControl_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) Brightness() int {
	defer qt.Recovering("QVideoWidgetControl::brightness")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectBrightnessChanged(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectBrightnessChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetControlBrightnessChanged
func callbackQVideoWidgetControlBrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidgetControl) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidgetControl) Contrast() int {
	defer qt.Recovering("QVideoWidgetControl::contrast")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectContrastChanged(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectContrastChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetControlContrastChanged
func callbackQVideoWidgetControlContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidgetControl) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidgetControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectFullScreenChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetControlFullScreenChanged
func callbackQVideoWidgetControlFullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidgetControl) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidgetControl) Hue() int {
	defer qt.Recovering("QVideoWidgetControl::hue")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectHueChanged(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectHueChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetControlHueChanged
func callbackQVideoWidgetControlHueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidgetControl) HueChanged(hue int) {
	defer qt.Recovering("QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_HueChanged(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidgetControl) IsFullScreen() bool {
	defer qt.Recovering("QVideoWidgetControl::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidgetControl) Saturation() int {
	defer qt.Recovering("QVideoWidgetControl::saturation")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectSaturationChanged(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSaturationChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetControlSaturationChanged
func callbackQVideoWidgetControlSaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidgetControl) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidgetControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWidgetControl::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWidgetControl) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWidgetControl::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidgetControl) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWidgetControl::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidgetControl) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWidgetControl::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidgetControl) SetHue(hue int) {
	defer qt.Recovering("QVideoWidgetControl::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidgetControl) SetSaturation(saturation int) {
	defer qt.Recovering("QVideoWidgetControl::setSaturation")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidgetControl) VideoWidget() *widgets.QWidget {
	defer qt.Recovering("QVideoWidgetControl::videoWidget")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QVideoWidgetControl_VideoWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidgetControl) DestroyQVideoWidgetControl() {
	defer qt.Recovering("QVideoWidgetControl::~QVideoWidgetControl")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DestroyQVideoWidgetControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWidgetControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoWidgetControlTimerEvent
func callbackQVideoWidgetControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidgetControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoWidgetControlChildEvent
func callbackQVideoWidgetControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoWidgetControlCustomEvent
func callbackQVideoWidgetControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidgetControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
