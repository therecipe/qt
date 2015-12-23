package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFontComboBox struct {
	QComboBox
}

type QFontComboBox_ITF interface {
	QComboBox_ITF
	QFontComboBox_PTR() *QFontComboBox
}

func PointerFromQFontComboBox(ptr QFontComboBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontComboBox_PTR().Pointer()
	}
	return nil
}

func NewQFontComboBoxFromPointer(ptr unsafe.Pointer) *QFontComboBox {
	var n = new(QFontComboBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFontComboBox_") {
		n.SetObjectName("QFontComboBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QFontComboBox) QFontComboBox_PTR() *QFontComboBox {
	return ptr
}

//QFontComboBox::FontFilter
type QFontComboBox__FontFilter int64

const (
	QFontComboBox__AllFonts          = QFontComboBox__FontFilter(0)
	QFontComboBox__ScalableFonts     = QFontComboBox__FontFilter(0x1)
	QFontComboBox__NonScalableFonts  = QFontComboBox__FontFilter(0x2)
	QFontComboBox__MonospacedFonts   = QFontComboBox__FontFilter(0x4)
	QFontComboBox__ProportionalFonts = QFontComboBox__FontFilter(0x8)
)

func (ptr *QFontComboBox) FontFilters() QFontComboBox__FontFilter {
	defer qt.Recovering("QFontComboBox::fontFilters")

	if ptr.Pointer() != nil {
		return QFontComboBox__FontFilter(C.QFontComboBox_FontFilters(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontComboBox) SetCurrentFont(font gui.QFont_ITF) {
	defer qt.Recovering("QFontComboBox::setCurrentFont")

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QFontComboBox) SetFontFilters(filters QFontComboBox__FontFilter) {
	defer qt.Recovering("QFontComboBox::setFontFilters")

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetFontFilters(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFontComboBox) SetWritingSystem(script gui.QFontDatabase__WritingSystem) {
	defer qt.Recovering("QFontComboBox::setWritingSystem")

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetWritingSystem(ptr.Pointer(), C.int(script))
	}
}

func (ptr *QFontComboBox) WritingSystem() gui.QFontDatabase__WritingSystem {
	defer qt.Recovering("QFontComboBox::writingSystem")

	if ptr.Pointer() != nil {
		return gui.QFontDatabase__WritingSystem(C.QFontComboBox_WritingSystem(ptr.Pointer()))
	}
	return 0
}

func NewQFontComboBox(parent QWidget_ITF) *QFontComboBox {
	defer qt.Recovering("QFontComboBox::QFontComboBox")

	return NewQFontComboBoxFromPointer(C.QFontComboBox_NewQFontComboBox(PointerFromQWidget(parent)))
}

func (ptr *QFontComboBox) SizeHint() *core.QSize {
	defer qt.Recovering("QFontComboBox::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QFontComboBox_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFontComboBox) DestroyQFontComboBox() {
	defer qt.Recovering("QFontComboBox::~QFontComboBox")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DestroyQFontComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFontComboBox) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QFontComboBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFontComboBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFontComboBoxChangeEvent
func callbackQFontComboBoxChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QFontComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QFontComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQFontComboBoxContextMenuEvent
func callbackQFontComboBoxContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFontComboBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QFontComboBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQFontComboBoxFocusInEvent
func callbackQFontComboBoxFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFontComboBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QFontComboBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQFontComboBoxFocusOutEvent
func callbackQFontComboBoxFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectHideEvent(f func(e *gui.QHideEvent)) {
	defer qt.Recovering("connect QFontComboBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QFontComboBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQFontComboBoxHideEvent
func callbackQFontComboBoxHideEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectHidePopup(f func()) {
	defer qt.Recovering("connect QFontComboBox::hidePopup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hidePopup", f)
	}
}

func (ptr *QFontComboBox) DisconnectHidePopup() {
	defer qt.Recovering("disconnect QFontComboBox::hidePopup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hidePopup")
	}
}

//export callbackQFontComboBoxHidePopup
func callbackQFontComboBoxHidePopup(ptrName *C.char) bool {
	defer qt.Recovering("callback QFontComboBox::hidePopup")

	if signal := qt.GetSignal(C.GoString(ptrName), "hidePopup"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QFontComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QFontComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQFontComboBoxInputMethodEvent
func callbackQFontComboBoxInputMethodEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFontComboBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QFontComboBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQFontComboBoxKeyPressEvent
func callbackQFontComboBoxKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFontComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QFontComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQFontComboBoxKeyReleaseEvent
func callbackQFontComboBoxKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontComboBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QFontComboBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQFontComboBoxMousePressEvent
func callbackQFontComboBoxMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QFontComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQFontComboBoxMouseReleaseEvent
func callbackQFontComboBoxMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFontComboBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFontComboBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFontComboBoxPaintEvent
func callbackQFontComboBoxPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QFontComboBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QFontComboBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQFontComboBoxResizeEvent
func callbackQFontComboBoxResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QFontComboBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QFontComboBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQFontComboBoxShowEvent
func callbackQFontComboBoxShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectShowPopup(f func()) {
	defer qt.Recovering("connect QFontComboBox::showPopup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showPopup", f)
	}
}

func (ptr *QFontComboBox) DisconnectShowPopup() {
	defer qt.Recovering("disconnect QFontComboBox::showPopup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showPopup")
	}
}

//export callbackQFontComboBoxShowPopup
func callbackQFontComboBoxShowPopup(ptrName *C.char) bool {
	defer qt.Recovering("callback QFontComboBox::showPopup")

	if signal := qt.GetSignal(C.GoString(ptrName), "showPopup"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QFontComboBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QFontComboBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQFontComboBoxWheelEvent
func callbackQFontComboBoxWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QFontComboBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QFontComboBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQFontComboBoxActionEvent
func callbackQFontComboBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QFontComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QFontComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQFontComboBoxDragEnterEvent
func callbackQFontComboBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QFontComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QFontComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQFontComboBoxDragLeaveEvent
func callbackQFontComboBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QFontComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QFontComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQFontComboBoxDragMoveEvent
func callbackQFontComboBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QFontComboBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QFontComboBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQFontComboBoxDropEvent
func callbackQFontComboBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFontComboBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QFontComboBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQFontComboBoxEnterEvent
func callbackQFontComboBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFontComboBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QFontComboBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQFontComboBoxLeaveEvent
func callbackQFontComboBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QFontComboBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QFontComboBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQFontComboBoxMoveEvent
func callbackQFontComboBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFontComboBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFontComboBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFontComboBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFontComboBoxSetVisible
func callbackQFontComboBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFontComboBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QFontComboBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QFontComboBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQFontComboBoxCloseEvent
func callbackQFontComboBoxCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QFontComboBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QFontComboBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QFontComboBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQFontComboBoxInitPainter
func callbackQFontComboBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QFontComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQFontComboBoxMouseDoubleClickEvent
func callbackQFontComboBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QFontComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQFontComboBoxMouseMoveEvent
func callbackQFontComboBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QFontComboBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QFontComboBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQFontComboBoxTabletEvent
func callbackQFontComboBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFontComboBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFontComboBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFontComboBoxTimerEvent
func callbackQFontComboBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFontComboBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFontComboBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFontComboBoxChildEvent
func callbackQFontComboBoxChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontComboBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFontComboBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFontComboBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFontComboBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFontComboBoxCustomEvent
func callbackQFontComboBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontComboBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
