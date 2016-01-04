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
