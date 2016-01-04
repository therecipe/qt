package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGroupBox struct {
	QWidget
}

type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func PointerFromQGroupBox(ptr QGroupBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGroupBox_PTR().Pointer()
	}
	return nil
}

func NewQGroupBoxFromPointer(ptr unsafe.Pointer) *QGroupBox {
	var n = new(QGroupBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGroupBox_") {
		n.SetObjectName("QGroupBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox {
	return ptr
}

func (ptr *QGroupBox) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QGroupBox::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGroupBox_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGroupBox) IsCheckable() bool {
	defer qt.Recovering("QGroupBox::isCheckable")

	if ptr.Pointer() != nil {
		return C.QGroupBox_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) IsChecked() bool {
	defer qt.Recovering("QGroupBox::isChecked")

	if ptr.Pointer() != nil {
		return C.QGroupBox_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) IsFlat() bool {
	defer qt.Recovering("QGroupBox::isFlat")

	if ptr.Pointer() != nil {
		return C.QGroupBox_IsFlat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) SetAlignment(alignment int) {
	defer qt.Recovering("QGroupBox::setAlignment")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QGroupBox) SetCheckable(checkable bool) {
	defer qt.Recovering("QGroupBox::setCheckable")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QGroupBox) SetChecked(checked bool) {
	defer qt.Recovering("QGroupBox::setChecked")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QGroupBox) SetFlat(flat bool) {
	defer qt.Recovering("QGroupBox::setFlat")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetFlat(ptr.Pointer(), C.int(qt.GoBoolToInt(flat)))
	}
}

func (ptr *QGroupBox) SetTitle(title string) {
	defer qt.Recovering("QGroupBox::setTitle")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QGroupBox) Title() string {
	defer qt.Recovering("QGroupBox::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGroupBox_Title(ptr.Pointer()))
	}
	return ""
}

func NewQGroupBox(parent QWidget_ITF) *QGroupBox {
	defer qt.Recovering("QGroupBox::QGroupBox")

	return NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox(PointerFromQWidget(parent)))
}

func NewQGroupBox2(title string, parent QWidget_ITF) *QGroupBox {
	defer qt.Recovering("QGroupBox::QGroupBox")

	return NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox2(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QGroupBox) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QGroupBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QGroupBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQGroupBoxChangeEvent
func callbackQGroupBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQGroupBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QGroupBox) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QGroupBox) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QGroupBox) ConnectChildEvent(f func(c *core.QChildEvent)) {
	defer qt.Recovering("connect QGroupBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGroupBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGroupBoxChildEvent
func callbackQGroupBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, c unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(c))
	} else {
		NewQGroupBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(c))
	}
}

func (ptr *QGroupBox) ChildEvent(c core.QChildEvent_ITF) {
	defer qt.Recovering("QGroupBox::childEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(c))
	}
}

func (ptr *QGroupBox) ChildEventDefault(c core.QChildEvent_ITF) {
	defer qt.Recovering("QGroupBox::childEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(c))
	}
}

func (ptr *QGroupBox) ConnectClicked(f func(checked bool)) {
	defer qt.Recovering("connect QGroupBox::clicked")

	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QGroupBox) DisconnectClicked() {
	defer qt.Recovering("disconnect QGroupBox::clicked")

	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQGroupBoxClicked
func callbackQGroupBoxClicked(ptr unsafe.Pointer, ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QGroupBox::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QGroupBox) Clicked(checked bool) {
	defer qt.Recovering("QGroupBox::clicked")

	if ptr.Pointer() != nil {
		C.QGroupBox_Clicked(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QGroupBox) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QGroupBox::event")

	if ptr.Pointer() != nil {
		return C.QGroupBox_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGroupBox) ConnectFocusInEvent(f func(fe *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGroupBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGroupBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGroupBoxFocusInEvent
func callbackQGroupBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, fe unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(fe))
	} else {
		NewQGroupBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(fe))
	}
}

func (ptr *QGroupBox) FocusInEvent(fe gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGroupBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(fe))
	}
}

func (ptr *QGroupBox) FocusInEventDefault(fe gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGroupBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(fe))
	}
}

func (ptr *QGroupBox) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QGroupBox::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QGroupBox_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGroupBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGroupBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQGroupBoxMouseMoveEvent
func callbackQGroupBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGroupBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGroupBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQGroupBoxMousePressEvent
func callbackQGroupBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGroupBox) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGroupBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQGroupBoxMouseReleaseEvent
func callbackQGroupBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGroupBox) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QGroupBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QGroupBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQGroupBoxPaintEvent
func callbackQGroupBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QGroupBox) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QGroupBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QGroupBox) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QGroupBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QGroupBox) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QGroupBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QGroupBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQGroupBoxResizeEvent
func callbackQGroupBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQGroupBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QGroupBox) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QGroupBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QGroupBox) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QGroupBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QGroupBox) ConnectToggled(f func(on bool)) {
	defer qt.Recovering("connect QGroupBox::toggled")

	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QGroupBox) DisconnectToggled() {
	defer qt.Recovering("disconnect QGroupBox::toggled")

	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQGroupBoxToggled
func callbackQGroupBoxToggled(ptr unsafe.Pointer, ptrName *C.char, on C.int) {
	defer qt.Recovering("callback QGroupBox::toggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "toggled"); signal != nil {
		signal.(func(bool))(int(on) != 0)
	}

}

func (ptr *QGroupBox) Toggled(on bool) {
	defer qt.Recovering("QGroupBox::toggled")

	if ptr.Pointer() != nil {
		C.QGroupBox_Toggled(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QGroupBox) DestroyQGroupBox() {
	defer qt.Recovering("QGroupBox::~QGroupBox")

	if ptr.Pointer() != nil {
		C.QGroupBox_DestroyQGroupBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGroupBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QGroupBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QGroupBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQGroupBoxActionEvent
func callbackQGroupBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QGroupBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QGroupBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QGroupBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QGroupBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QGroupBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QGroupBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGroupBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQGroupBoxDragEnterEvent
func callbackQGroupBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QGroupBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QGroupBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QGroupBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QGroupBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QGroupBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QGroupBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGroupBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQGroupBoxDragLeaveEvent
func callbackQGroupBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QGroupBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QGroupBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QGroupBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QGroupBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QGroupBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QGroupBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGroupBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQGroupBoxDragMoveEvent
func callbackQGroupBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QGroupBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QGroupBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QGroupBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QGroupBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QGroupBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QGroupBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGroupBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQGroupBoxDropEvent
func callbackQGroupBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QGroupBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QGroupBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QGroupBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QGroupBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QGroupBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGroupBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QGroupBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQGroupBoxEnterEvent
func callbackQGroupBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGroupBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGroupBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGroupBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGroupBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGroupBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQGroupBoxFocusOutEvent
func callbackQGroupBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGroupBox) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGroupBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGroupBox) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGroupBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGroupBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QGroupBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QGroupBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQGroupBoxHideEvent
func callbackQGroupBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGroupBox) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGroupBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGroupBox) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGroupBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGroupBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGroupBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QGroupBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQGroupBoxLeaveEvent
func callbackQGroupBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGroupBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGroupBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGroupBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QGroupBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QGroupBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQGroupBoxMoveEvent
func callbackQGroupBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QGroupBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QGroupBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QGroupBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QGroupBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QGroupBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QGroupBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QGroupBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QGroupBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQGroupBoxSetVisible
func callbackQGroupBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QGroupBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QGroupBox) SetVisible(visible bool) {
	defer qt.Recovering("QGroupBox::setVisible")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QGroupBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QGroupBox::setVisible")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QGroupBox) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QGroupBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QGroupBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQGroupBoxShowEvent
func callbackQGroupBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGroupBox) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGroupBox::showEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGroupBox) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGroupBox::showEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGroupBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QGroupBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QGroupBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQGroupBoxCloseEvent
func callbackQGroupBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGroupBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGroupBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGroupBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGroupBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGroupBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QGroupBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGroupBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQGroupBoxContextMenuEvent
func callbackQGroupBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QGroupBox) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QGroupBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QGroupBox) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QGroupBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QGroupBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QGroupBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QGroupBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QGroupBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQGroupBoxInitPainter
func callbackQGroupBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQGroupBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QGroupBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QGroupBox::initPainter")

	if ptr.Pointer() != nil {
		C.QGroupBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QGroupBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QGroupBox::initPainter")

	if ptr.Pointer() != nil {
		C.QGroupBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QGroupBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGroupBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGroupBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQGroupBoxInputMethodEvent
func callbackQGroupBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGroupBox) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGroupBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGroupBox) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGroupBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGroupBox) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGroupBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGroupBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQGroupBoxKeyPressEvent
func callbackQGroupBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGroupBox) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGroupBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGroupBox) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGroupBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGroupBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGroupBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGroupBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQGroupBoxKeyReleaseEvent
func callbackQGroupBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGroupBox) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGroupBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGroupBox) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGroupBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGroupBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGroupBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQGroupBoxMouseDoubleClickEvent
func callbackQGroupBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGroupBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGroupBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGroupBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QGroupBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QGroupBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQGroupBoxTabletEvent
func callbackQGroupBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QGroupBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QGroupBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QGroupBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QGroupBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QGroupBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QGroupBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGroupBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQGroupBoxWheelEvent
func callbackQGroupBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QGroupBox) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QGroupBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QGroupBox) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QGroupBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QGroupBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGroupBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGroupBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGroupBoxTimerEvent
func callbackQGroupBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGroupBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGroupBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGroupBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGroupBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGroupBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGroupBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGroupBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGroupBoxCustomEvent
func callbackQGroupBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGroupBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGroupBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGroupBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::customEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGroupBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGroupBox::customEvent")

	if ptr.Pointer() != nil {
		C.QGroupBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
