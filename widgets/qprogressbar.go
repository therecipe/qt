package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QProgressBar struct {
	QWidget
}

type QProgressBar_ITF interface {
	QWidget_ITF
	QProgressBar_PTR() *QProgressBar
}

func PointerFromQProgressBar(ptr QProgressBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressBar_PTR().Pointer()
	}
	return nil
}

func NewQProgressBarFromPointer(ptr unsafe.Pointer) *QProgressBar {
	var n = new(QProgressBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProgressBar_") {
		n.SetObjectName("QProgressBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QProgressBar) QProgressBar_PTR() *QProgressBar {
	return ptr
}

//QProgressBar::Direction
type QProgressBar__Direction int64

const (
	QProgressBar__TopToBottom = QProgressBar__Direction(0)
	QProgressBar__BottomToTop = QProgressBar__Direction(1)
)

func (ptr *QProgressBar) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QProgressBar::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QProgressBar_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Format() string {
	defer qt.Recovering("QProgressBar::format")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Format(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) InvertedAppearance() bool {
	defer qt.Recovering("QProgressBar::invertedAppearance")

	if ptr.Pointer() != nil {
		return C.QProgressBar_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) IsTextVisible() bool {
	defer qt.Recovering("QProgressBar::isTextVisible")

	if ptr.Pointer() != nil {
		return C.QProgressBar_IsTextVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressBar) Maximum() int {
	defer qt.Recovering("QProgressBar::maximum")

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Minimum() int {
	defer qt.Recovering("QProgressBar::minimum")

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QProgressBar::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QProgressBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) ResetFormat() {
	defer qt.Recovering("QProgressBar::resetFormat")

	if ptr.Pointer() != nil {
		C.QProgressBar_ResetFormat(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QProgressBar::setAlignment")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QProgressBar) SetFormat(format string) {
	defer qt.Recovering("QProgressBar::setFormat")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QProgressBar) SetInvertedAppearance(invert bool) {
	defer qt.Recovering("QProgressBar::setInvertedAppearance")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(invert)))
	}
}

func (ptr *QProgressBar) SetMaximum(maximum int) {
	defer qt.Recovering("QProgressBar::setMaximum")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressBar) SetMinimum(minimum int) {
	defer qt.Recovering("QProgressBar::setMinimum")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressBar) SetOrientation(v core.Qt__Orientation) {
	defer qt.Recovering("QProgressBar::setOrientation")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	defer qt.Recovering("QProgressBar::setTextDirection")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextDirection(ptr.Pointer(), C.int(textDirection))
	}
}

func (ptr *QProgressBar) SetTextVisible(visible bool) {
	defer qt.Recovering("QProgressBar::setTextVisible")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetTextVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) SetValue(value int) {
	defer qt.Recovering("QProgressBar::setValue")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetValue(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QProgressBar) Text() string {
	defer qt.Recovering("QProgressBar::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressBar_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressBar) TextDirection() QProgressBar__Direction {
	defer qt.Recovering("QProgressBar::textDirection")

	if ptr.Pointer() != nil {
		return QProgressBar__Direction(C.QProgressBar_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressBar) Value() int {
	defer qt.Recovering("QProgressBar::value")

	if ptr.Pointer() != nil {
		return int(C.QProgressBar_Value(ptr.Pointer()))
	}
	return 0
}

func NewQProgressBar(parent QWidget_ITF) *QProgressBar {
	defer qt.Recovering("QProgressBar::QProgressBar")

	return NewQProgressBarFromPointer(C.QProgressBar_NewQProgressBar(PointerFromQWidget(parent)))
}

func (ptr *QProgressBar) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QProgressBar::event")

	if ptr.Pointer() != nil {
		return C.QProgressBar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QProgressBar) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QProgressBar::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProgressBar_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProgressBar) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QProgressBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QProgressBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQProgressBarPaintEvent
func callbackQProgressBarPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQProgressBarFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QProgressBar) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QProgressBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QProgressBar) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QProgressBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QProgressBar) Reset() {
	defer qt.Recovering("QProgressBar::reset")

	if ptr.Pointer() != nil {
		C.QProgressBar_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressBar) SetRange(minimum int, maximum int) {
	defer qt.Recovering("QProgressBar::setRange")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressBar) SizeHint() *core.QSize {
	defer qt.Recovering("QProgressBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QProgressBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProgressBar) ConnectValueChanged(f func(value int)) {
	defer qt.Recovering("connect QProgressBar::valueChanged")

	if ptr.Pointer() != nil {
		C.QProgressBar_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QProgressBar) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QProgressBar::valueChanged")

	if ptr.Pointer() != nil {
		C.QProgressBar_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQProgressBarValueChanged
func callbackQProgressBarValueChanged(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QProgressBar::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QProgressBar) ValueChanged(value int) {
	defer qt.Recovering("QProgressBar::valueChanged")

	if ptr.Pointer() != nil {
		C.QProgressBar_ValueChanged(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QProgressBar) DestroyQProgressBar() {
	defer qt.Recovering("QProgressBar::~QProgressBar")

	if ptr.Pointer() != nil {
		C.QProgressBar_DestroyQProgressBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QProgressBar) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QProgressBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QProgressBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQProgressBarActionEvent
func callbackQProgressBarActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QProgressBar) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QProgressBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QProgressBar) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QProgressBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QProgressBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QProgressBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QProgressBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQProgressBarDragEnterEvent
func callbackQProgressBarDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QProgressBar) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QProgressBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QProgressBar) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QProgressBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QProgressBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QProgressBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QProgressBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQProgressBarDragLeaveEvent
func callbackQProgressBarDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QProgressBar) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QProgressBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QProgressBar) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QProgressBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QProgressBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QProgressBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QProgressBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQProgressBarDragMoveEvent
func callbackQProgressBarDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QProgressBar) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QProgressBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QProgressBar) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QProgressBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QProgressBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QProgressBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QProgressBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQProgressBarDropEvent
func callbackQProgressBarDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QProgressBar) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QProgressBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QProgressBar) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QProgressBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QProgressBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QProgressBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQProgressBarEnterEvent
func callbackQProgressBarEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressBar) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QProgressBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QProgressBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQProgressBarFocusInEvent
func callbackQProgressBarFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QProgressBar) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QProgressBar) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QProgressBar) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QProgressBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QProgressBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQProgressBarFocusOutEvent
func callbackQProgressBarFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QProgressBar) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QProgressBar) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QProgressBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QProgressBar) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QProgressBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QProgressBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQProgressBarHideEvent
func callbackQProgressBarHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QProgressBar) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QProgressBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QProgressBar) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QProgressBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QProgressBar) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QProgressBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQProgressBarLeaveEvent
func callbackQProgressBarLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressBar) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QProgressBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QProgressBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQProgressBarMoveEvent
func callbackQProgressBarMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QProgressBar) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QProgressBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QProgressBar) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QProgressBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QProgressBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QProgressBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QProgressBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QProgressBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQProgressBarSetVisible
func callbackQProgressBarSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QProgressBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QProgressBar) SetVisible(visible bool) {
	defer qt.Recovering("QProgressBar::setVisible")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QProgressBar::setVisible")

	if ptr.Pointer() != nil {
		C.QProgressBar_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QProgressBar) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QProgressBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QProgressBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQProgressBarShowEvent
func callbackQProgressBarShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QProgressBar) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QProgressBar::showEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QProgressBar) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QProgressBar::showEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QProgressBar) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QProgressBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQProgressBarChangeEvent
func callbackQProgressBarChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressBar) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QProgressBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QProgressBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQProgressBarCloseEvent
func callbackQProgressBarCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QProgressBar) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QProgressBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QProgressBar) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QProgressBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QProgressBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QProgressBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QProgressBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQProgressBarContextMenuEvent
func callbackQProgressBarContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QProgressBar) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QProgressBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QProgressBar) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QProgressBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QProgressBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QProgressBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QProgressBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QProgressBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQProgressBarInitPainter
func callbackQProgressBarInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQProgressBarFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QProgressBar) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QProgressBar::initPainter")

	if ptr.Pointer() != nil {
		C.QProgressBar_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QProgressBar) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QProgressBar::initPainter")

	if ptr.Pointer() != nil {
		C.QProgressBar_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QProgressBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QProgressBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QProgressBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQProgressBarInputMethodEvent
func callbackQProgressBarInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QProgressBar) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QProgressBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QProgressBar) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QProgressBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QProgressBar) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QProgressBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QProgressBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQProgressBarKeyPressEvent
func callbackQProgressBarKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QProgressBar) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QProgressBar) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QProgressBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QProgressBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QProgressBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQProgressBarKeyReleaseEvent
func callbackQProgressBarKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QProgressBar) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QProgressBar) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QProgressBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QProgressBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QProgressBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQProgressBarMouseDoubleClickEvent
func callbackQProgressBarMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressBar) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QProgressBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQProgressBarMouseMoveEvent
func callbackQProgressBarMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressBar) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QProgressBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQProgressBarMousePressEvent
func callbackQProgressBarMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressBar) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QProgressBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QProgressBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQProgressBarMouseReleaseEvent
func callbackQProgressBarMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QProgressBar) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QProgressBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QProgressBar) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QProgressBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QProgressBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQProgressBarResizeEvent
func callbackQProgressBarResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QProgressBar) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QProgressBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QProgressBar) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QProgressBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QProgressBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QProgressBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QProgressBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQProgressBarTabletEvent
func callbackQProgressBarTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QProgressBar) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QProgressBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QProgressBar) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QProgressBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QProgressBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QProgressBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QProgressBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQProgressBarWheelEvent
func callbackQProgressBarWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QProgressBar) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QProgressBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QProgressBar) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QProgressBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QProgressBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QProgressBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProgressBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProgressBarTimerEvent
func callbackQProgressBarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProgressBar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProgressBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProgressBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProgressBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProgressBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QProgressBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProgressBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProgressBarChildEvent
func callbackQProgressBarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProgressBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProgressBar::childEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProgressBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProgressBar::childEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProgressBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProgressBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProgressBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProgressBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProgressBarCustomEvent
func callbackQProgressBarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProgressBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProgressBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProgressBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::customEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProgressBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProgressBar::customEvent")

	if ptr.Pointer() != nil {
		C.QProgressBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
