package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QRubberBand struct {
	QWidget
}

type QRubberBand_ITF interface {
	QWidget_ITF
	QRubberBand_PTR() *QRubberBand
}

func PointerFromQRubberBand(ptr QRubberBand_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRubberBand_PTR().Pointer()
	}
	return nil
}

func NewQRubberBandFromPointer(ptr unsafe.Pointer) *QRubberBand {
	var n = new(QRubberBand)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRubberBand_") {
		n.SetObjectName("QRubberBand_" + qt.Identifier())
	}
	return n
}

func (ptr *QRubberBand) QRubberBand_PTR() *QRubberBand {
	return ptr
}

//QRubberBand::Shape
type QRubberBand__Shape int64

const (
	QRubberBand__Line      = QRubberBand__Shape(0)
	QRubberBand__Rectangle = QRubberBand__Shape(1)
)

func (ptr *QRubberBand) SetGeometry(rect core.QRect_ITF) {
	defer qt.Recovering("QRubberBand::setGeometry")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func NewQRubberBand(s QRubberBand__Shape, p QWidget_ITF) *QRubberBand {
	defer qt.Recovering("QRubberBand::QRubberBand")

	return NewQRubberBandFromPointer(C.QRubberBand_NewQRubberBand(C.int(s), PointerFromQWidget(p)))
}

func (ptr *QRubberBand) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QRubberBand::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQRubberBandChangeEvent
func callbackQRubberBandChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQRubberBandFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QRubberBand) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::changeEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QRubberBand) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::changeEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QRubberBand) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QRubberBand::event")

	if ptr.Pointer() != nil {
		return C.QRubberBand_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QRubberBand) Move2(p core.QPoint_ITF) {
	defer qt.Recovering("QRubberBand::move")

	if ptr.Pointer() != nil {
		C.QRubberBand_Move2(ptr.Pointer(), core.PointerFromQPoint(p))
	}
}

func (ptr *QRubberBand) Move(x int, y int) {
	defer qt.Recovering("QRubberBand::move")

	if ptr.Pointer() != nil {
		C.QRubberBand_Move(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QRubberBand) ConnectMoveEvent(f func(v *gui.QMoveEvent)) {
	defer qt.Recovering("connect QRubberBand::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQRubberBandMoveEvent
func callbackQRubberBandMoveEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(v))
	} else {
		NewQRubberBandFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(v))
	}
}

func (ptr *QRubberBand) MoveEvent(v gui.QMoveEvent_ITF) {
	defer qt.Recovering("QRubberBand::moveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(v))
	}
}

func (ptr *QRubberBand) MoveEventDefault(v gui.QMoveEvent_ITF) {
	defer qt.Recovering("QRubberBand::moveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(v))
	}
}

func (ptr *QRubberBand) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QRubberBand::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRubberBand::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRubberBandPaintEvent
func callbackQRubberBandPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQRubberBandFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QRubberBand) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QRubberBand::paintEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QRubberBand) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QRubberBand::paintEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QRubberBand) Resize2(size core.QSize_ITF) {
	defer qt.Recovering("QRubberBand::resize")

	if ptr.Pointer() != nil {
		C.QRubberBand_Resize2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QRubberBand) Resize(width int, height int) {
	defer qt.Recovering("QRubberBand::resize")

	if ptr.Pointer() != nil {
		C.QRubberBand_Resize(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQRubberBandResizeEvent
func callbackQRubberBandResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQRubberBandFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QRubberBand) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QRubberBand) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QRubberBand) SetGeometry2(x int, y int, width int, height int) {
	defer qt.Recovering("QRubberBand::setGeometry")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) Shape() QRubberBand__Shape {
	defer qt.Recovering("QRubberBand::shape")

	if ptr.Pointer() != nil {
		return QRubberBand__Shape(C.QRubberBand_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRubberBand) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QRubberBand::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QRubberBand::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQRubberBandShowEvent
func callbackQRubberBandShowEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
	} else {
		NewQRubberBandFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(e))
	}
}

func (ptr *QRubberBand) ShowEvent(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QRubberBand::showEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
}

func (ptr *QRubberBand) ShowEventDefault(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QRubberBand::showEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
}

func (ptr *QRubberBand) DestroyQRubberBand() {
	defer qt.Recovering("QRubberBand::~QRubberBand")

	if ptr.Pointer() != nil {
		C.QRubberBand_DestroyQRubberBand(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRubberBand) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QRubberBand::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QRubberBand::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQRubberBandActionEvent
func callbackQRubberBandActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QRubberBand) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QRubberBand::actionEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QRubberBand) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QRubberBand::actionEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QRubberBand) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QRubberBand::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QRubberBand::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQRubberBandDragEnterEvent
func callbackQRubberBandDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QRubberBand) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QRubberBand::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QRubberBand) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QRubberBand::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QRubberBand) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QRubberBand::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QRubberBand::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQRubberBandDragLeaveEvent
func callbackQRubberBandDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QRubberBand) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QRubberBand::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QRubberBand) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QRubberBand::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QRubberBand) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QRubberBand::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQRubberBandDragMoveEvent
func callbackQRubberBandDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QRubberBand) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QRubberBand::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QRubberBand) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QRubberBand::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QRubberBand) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QRubberBand::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QRubberBand::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQRubberBandDropEvent
func callbackQRubberBandDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QRubberBand) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QRubberBand::dropEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QRubberBand) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QRubberBand::dropEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QRubberBand) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QRubberBand::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQRubberBandEnterEvent
func callbackQRubberBandEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRubberBand) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::enterEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRubberBand) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::enterEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRubberBand) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRubberBand::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QRubberBand::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQRubberBandFocusInEvent
func callbackQRubberBandFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QRubberBand) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRubberBand::focusInEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QRubberBand) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRubberBand::focusInEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QRubberBand) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRubberBand::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QRubberBand::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQRubberBandFocusOutEvent
func callbackQRubberBandFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QRubberBand) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRubberBand::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QRubberBand) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QRubberBand::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QRubberBand) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QRubberBand::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QRubberBand::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQRubberBandHideEvent
func callbackQRubberBandHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QRubberBand) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QRubberBand::hideEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QRubberBand) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QRubberBand::hideEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QRubberBand) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QRubberBand::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQRubberBandLeaveEvent
func callbackQRubberBandLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRubberBand) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::leaveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRubberBand) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::leaveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRubberBand) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QRubberBand::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QRubberBand) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QRubberBand::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQRubberBandSetVisible
func callbackQRubberBandSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QRubberBand::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QRubberBand) SetVisible(visible bool) {
	defer qt.Recovering("QRubberBand::setVisible")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QRubberBand) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QRubberBand::setVisible")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QRubberBand) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QRubberBand::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QRubberBand::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQRubberBandCloseEvent
func callbackQRubberBandCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QRubberBand) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QRubberBand::closeEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QRubberBand) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QRubberBand::closeEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QRubberBand) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QRubberBand::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QRubberBand::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQRubberBandContextMenuEvent
func callbackQRubberBandContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QRubberBand) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QRubberBand::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QRubberBand) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QRubberBand::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QRubberBand) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QRubberBand::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QRubberBand) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QRubberBand::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQRubberBandInitPainter
func callbackQRubberBandInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQRubberBandFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QRubberBand) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QRubberBand::initPainter")

	if ptr.Pointer() != nil {
		C.QRubberBand_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QRubberBand) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QRubberBand::initPainter")

	if ptr.Pointer() != nil {
		C.QRubberBand_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QRubberBand) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QRubberBand::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QRubberBand::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQRubberBandInputMethodEvent
func callbackQRubberBandInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QRubberBand) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QRubberBand::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QRubberBand) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QRubberBand::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QRubberBand) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRubberBand::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QRubberBand::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQRubberBandKeyPressEvent
func callbackQRubberBandKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QRubberBand) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRubberBand::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QRubberBand) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRubberBand::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QRubberBand) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRubberBand::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QRubberBand::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQRubberBandKeyReleaseEvent
func callbackQRubberBandKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QRubberBand) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRubberBand::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QRubberBand) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QRubberBand::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QRubberBand) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QRubberBand::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQRubberBandMouseDoubleClickEvent
func callbackQRubberBandMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QRubberBand) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQRubberBandMouseMoveEvent
func callbackQRubberBandMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QRubberBand) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QRubberBand::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQRubberBandMousePressEvent
func callbackQRubberBandMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QRubberBand) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QRubberBand::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQRubberBandMouseReleaseEvent
func callbackQRubberBandMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QRubberBand) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QRubberBand::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QRubberBand) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QRubberBand::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QRubberBand::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQRubberBandTabletEvent
func callbackQRubberBandTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QRubberBand) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QRubberBand::tabletEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QRubberBand) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QRubberBand::tabletEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QRubberBand) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QRubberBand::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QRubberBand::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQRubberBandWheelEvent
func callbackQRubberBandWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QRubberBand) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QRubberBand::wheelEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QRubberBand) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QRubberBand::wheelEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QRubberBand) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRubberBand::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRubberBand::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRubberBandTimerEvent
func callbackQRubberBandTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRubberBand) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRubberBand::timerEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRubberBand) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRubberBand::timerEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRubberBand) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRubberBand::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRubberBand::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRubberBandChildEvent
func callbackQRubberBandChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRubberBand) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRubberBand::childEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRubberBand) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRubberBand::childEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRubberBand) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRubberBand::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRubberBandCustomEvent
func callbackQRubberBandCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRubberBand::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRubberBandFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRubberBand) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::customEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRubberBand) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRubberBand::customEvent")

	if ptr.Pointer() != nil {
		C.QRubberBand_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
