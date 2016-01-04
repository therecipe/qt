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

func (ptr *QFontComboBox) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QFontComboBox::event")

	if ptr.Pointer() != nil {
		return C.QFontComboBox_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
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
func callbackQFontComboBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QFontComboBox) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQFontComboBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFontComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QFontComboBox) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFontComboBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
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
func callbackQFontComboBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontComboBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QFontComboBox) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontComboBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQFontComboBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontComboBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QFontComboBox) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontComboBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQFontComboBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) HideEvent(e gui.QHideEvent_ITF) {
	defer qt.Recovering("QFontComboBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(e))
	}
}

func (ptr *QFontComboBox) HideEventDefault(e gui.QHideEvent_ITF) {
	defer qt.Recovering("QFontComboBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(e))
	}
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
func callbackQFontComboBoxHidePopup(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QFontComboBox::hidePopup")

	if signal := qt.GetSignal(C.GoString(ptrName), "hidePopup"); signal != nil {
		signal.(func())()
	} else {
		NewQFontComboBoxFromPointer(ptr).HidePopupDefault()
	}
}

func (ptr *QFontComboBox) HidePopup() {
	defer qt.Recovering("QFontComboBox::hidePopup")

	if ptr.Pointer() != nil {
		C.QFontComboBox_HidePopup(ptr.Pointer())
	}
}

func (ptr *QFontComboBox) HidePopupDefault() {
	defer qt.Recovering("QFontComboBox::hidePopup")

	if ptr.Pointer() != nil {
		C.QFontComboBox_HidePopupDefault(ptr.Pointer())
	}
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
func callbackQFontComboBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFontComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QFontComboBox) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFontComboBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
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
func callbackQFontComboBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontComboBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QFontComboBox) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontComboBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQFontComboBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QFontComboBox) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontComboBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQFontComboBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QFontComboBox) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQFontComboBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QFontComboBox) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQFontComboBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFontComboBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QFontComboBox) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFontComboBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
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
func callbackQFontComboBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFontComboBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QFontComboBox) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFontComboBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
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
func callbackQFontComboBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) ShowEvent(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QFontComboBox::showEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
}

func (ptr *QFontComboBox) ShowEventDefault(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QFontComboBox::showEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
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
func callbackQFontComboBoxShowPopup(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QFontComboBox::showPopup")

	if signal := qt.GetSignal(C.GoString(ptrName), "showPopup"); signal != nil {
		signal.(func())()
	} else {
		NewQFontComboBoxFromPointer(ptr).ShowPopupDefault()
	}
}

func (ptr *QFontComboBox) ShowPopup() {
	defer qt.Recovering("QFontComboBox::showPopup")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ShowPopup(ptr.Pointer())
	}
}

func (ptr *QFontComboBox) ShowPopupDefault() {
	defer qt.Recovering("QFontComboBox::showPopup")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ShowPopupDefault(ptr.Pointer())
	}
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
func callbackQFontComboBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQFontComboBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QFontComboBox) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFontComboBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QFontComboBox) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFontComboBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
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
func callbackQFontComboBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFontComboBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QFontComboBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFontComboBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQFontComboBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QFontComboBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQFontComboBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QFontComboBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQFontComboBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QFontComboBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQFontComboBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QFontComboBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFontComboBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQFontComboBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFontComboBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQFontComboBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFontComboBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQFontComboBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFontComboBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QFontComboBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFontComboBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQFontComboBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFontComboBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFontComboBox) SetVisible(visible bool) {
	defer qt.Recovering("QFontComboBox::setVisible")

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFontComboBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QFontComboBox::setVisible")

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQFontComboBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QFontComboBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQFontComboBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQFontComboBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QFontComboBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFontComboBox::initPainter")

	if ptr.Pointer() != nil {
		C.QFontComboBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QFontComboBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFontComboBox::initPainter")

	if ptr.Pointer() != nil {
		C.QFontComboBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQFontComboBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFontComboBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQFontComboBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFontComboBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontComboBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQFontComboBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFontComboBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QFontComboBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFontComboBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQFontComboBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFontComboBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFontComboBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFontComboBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQFontComboBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFontComboBox::childEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFontComboBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFontComboBox::childEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQFontComboBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontComboBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFontComboBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFontComboBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::customEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFontComboBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFontComboBox::customEvent")

	if ptr.Pointer() != nil {
		C.QFontComboBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
