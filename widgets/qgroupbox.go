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
func callbackQGroupBoxChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQGroupBoxChildEvent(ptrName *C.char, c unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(c))
		return true
	}
	return false

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
func callbackQGroupBoxClicked(ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QGroupBox::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

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
func callbackQGroupBoxFocusInEvent(ptrName *C.char, fe unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(fe))
		return true
	}
	return false

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
func callbackQGroupBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

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
func callbackQGroupBoxToggled(ptrName *C.char, on C.int) {
	defer qt.Recovering("callback QGroupBox::toggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "toggled"); signal != nil {
		signal.(func(bool))(int(on) != 0)
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
func callbackQGroupBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QGroupBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQGroupBoxShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQGroupBoxInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGroupBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
