package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QLabel struct {
	QFrame
}

type QLabel_ITF interface {
	QFrame_ITF
	QLabel_PTR() *QLabel
}

func PointerFromQLabel(ptr QLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabel_PTR().Pointer()
	}
	return nil
}

func NewQLabelFromPointer(ptr unsafe.Pointer) *QLabel {
	var n = new(QLabel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLabel_") {
		n.SetObjectName("QLabel_" + qt.Identifier())
	}
	return n
}

func (ptr *QLabel) QLabel_PTR() *QLabel {
	return ptr
}

func (ptr *QLabel) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QLabel::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLabel_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) HasScaledContents() bool {
	defer qt.Recovering("QLabel::hasScaledContents")

	if ptr.Pointer() != nil {
		return C.QLabel_HasScaledContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) HasSelectedText() bool {
	defer qt.Recovering("QLabel::hasSelectedText")

	if ptr.Pointer() != nil {
		return C.QLabel_HasSelectedText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) Indent() int {
	defer qt.Recovering("QLabel::indent")

	if ptr.Pointer() != nil {
		return int(C.QLabel_Indent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) Margin() int {
	defer qt.Recovering("QLabel::margin")

	if ptr.Pointer() != nil {
		return int(C.QLabel_Margin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) OpenExternalLinks() bool {
	defer qt.Recovering("QLabel::openExternalLinks")

	if ptr.Pointer() != nil {
		return C.QLabel_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) Pixmap() *gui.QPixmap {
	defer qt.Recovering("QLabel::pixmap")

	if ptr.Pointer() != nil {
		return gui.NewQPixmapFromPointer(C.QLabel_Pixmap(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) SelectedText() string {
	defer qt.Recovering("QLabel::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) SetAlignment(v core.Qt__AlignmentFlag) {
	defer qt.Recovering("QLabel::setAlignment")

	if ptr.Pointer() != nil {
		C.QLabel_SetAlignment(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetIndent(v int) {
	defer qt.Recovering("QLabel::setIndent")

	if ptr.Pointer() != nil {
		C.QLabel_SetIndent(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetMargin(v int) {
	defer qt.Recovering("QLabel::setMargin")

	if ptr.Pointer() != nil {
		C.QLabel_SetMargin(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetOpenExternalLinks(open bool) {
	defer qt.Recovering("QLabel::setOpenExternalLinks")

	if ptr.Pointer() != nil {
		C.QLabel_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QLabel) SetPixmap(v gui.QPixmap_ITF) {
	defer qt.Recovering("QLabel::setPixmap")

	if ptr.Pointer() != nil {
		C.QLabel_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(v))
	}
}

func (ptr *QLabel) SetScaledContents(v bool) {
	defer qt.Recovering("QLabel::setScaledContents")

	if ptr.Pointer() != nil {
		C.QLabel_SetScaledContents(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLabel) SetText(v string) {
	defer qt.Recovering("QLabel::setText")

	if ptr.Pointer() != nil {
		C.QLabel_SetText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLabel) SetTextFormat(v core.Qt__TextFormat) {
	defer qt.Recovering("QLabel::setTextFormat")

	if ptr.Pointer() != nil {
		C.QLabel_SetTextFormat(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer qt.Recovering("QLabel::setTextInteractionFlags")

	if ptr.Pointer() != nil {
		C.QLabel_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QLabel) SetWordWrap(on bool) {
	defer qt.Recovering("QLabel::setWordWrap")

	if ptr.Pointer() != nil {
		C.QLabel_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QLabel) Text() string {
	defer qt.Recovering("QLabel::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) TextFormat() core.Qt__TextFormat {
	defer qt.Recovering("QLabel::textFormat")

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QLabel_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer qt.Recovering("QLabel::textInteractionFlags")

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QLabel_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) WordWrap() bool {
	defer qt.Recovering("QLabel::wordWrap")

	if ptr.Pointer() != nil {
		return C.QLabel_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQLabel(parent QWidget_ITF, f core.Qt__WindowType) *QLabel {
	defer qt.Recovering("QLabel::QLabel")

	return NewQLabelFromPointer(C.QLabel_NewQLabel(PointerFromQWidget(parent), C.int(f)))
}

func NewQLabel2(text string, parent QWidget_ITF, f core.Qt__WindowType) *QLabel {
	defer qt.Recovering("QLabel::QLabel")

	return NewQLabelFromPointer(C.QLabel_NewQLabel2(C.CString(text), PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QLabel) Buddy() *QWidget {
	defer qt.Recovering("QLabel::buddy")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLabel_Buddy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QLabel::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QLabel) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QLabel::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQLabelChangeEvent
func callbackQLabelChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QLabel) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QLabel::changeEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QLabel) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QLabel::changeEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QLabel) Clear() {
	defer qt.Recovering("QLabel::clear")

	if ptr.Pointer() != nil {
		C.QLabel_Clear(ptr.Pointer())
	}
}

func (ptr *QLabel) ConnectContextMenuEvent(f func(ev *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QLabel::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QLabel) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QLabel::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQLabelContextMenuEvent
func callbackQLabelContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(ev))
	}
}

func (ptr *QLabel) ContextMenuEvent(ev gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QLabel::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

func (ptr *QLabel) ContextMenuEventDefault(ev gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QLabel::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

func (ptr *QLabel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLabel::event")

	if ptr.Pointer() != nil {
		return C.QLabel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLabel) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QLabel::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QLabel) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QLabel::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQLabelFocusInEvent
func callbackQLabelFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QLabel) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLabel::focusInEvent")

	if ptr.Pointer() != nil {
		C.QLabel_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QLabel) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLabel::focusInEvent")

	if ptr.Pointer() != nil {
		C.QLabel_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QLabel) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QLabel::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QLabel_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QLabel) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QLabel::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QLabel) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QLabel::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQLabelFocusOutEvent
func callbackQLabelFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QLabel) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLabel::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QLabel_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QLabel) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLabel::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QLabel_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QLabel) HeightForWidth(w int) int {
	defer qt.Recovering("QLabel::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QLabel_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLabel) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QLabel::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QLabel) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QLabel::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQLabelKeyPressEvent
func callbackQLabelKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QLabel) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLabel::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QLabel_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QLabel) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLabel::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QLabel_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QLabel) ConnectLinkActivated(f func(link string)) {
	defer qt.Recovering("connect QLabel::linkActivated")

	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QLabel) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QLabel::linkActivated")

	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQLabelLinkActivated
func callbackQLabelLinkActivated(ptr unsafe.Pointer, ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QLabel::linkActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
		signal.(func(string))(C.GoString(link))
	}

}

func (ptr *QLabel) LinkActivated(link string) {
	defer qt.Recovering("QLabel::linkActivated")

	if ptr.Pointer() != nil {
		C.QLabel_LinkActivated(ptr.Pointer(), C.CString(link))
	}
}

func (ptr *QLabel) ConnectLinkHovered(f func(link string)) {
	defer qt.Recovering("connect QLabel::linkHovered")

	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QLabel) DisconnectLinkHovered() {
	defer qt.Recovering("disconnect QLabel::linkHovered")

	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQLabelLinkHovered
func callbackQLabelLinkHovered(ptr unsafe.Pointer, ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QLabel::linkHovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkHovered"); signal != nil {
		signal.(func(string))(C.GoString(link))
	}

}

func (ptr *QLabel) LinkHovered(link string) {
	defer qt.Recovering("QLabel::linkHovered")

	if ptr.Pointer() != nil {
		C.QLabel_LinkHovered(ptr.Pointer(), C.CString(link))
	}
}

func (ptr *QLabel) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QLabel::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLabel_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) ConnectMouseMoveEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLabel::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QLabel) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QLabel::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQLabelMouseMoveEvent
func callbackQLabelMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QLabel) MouseMoveEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QLabel) MouseMoveEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QLabel) ConnectMousePressEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLabel::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QLabel) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QLabel::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQLabelMousePressEvent
func callbackQLabelMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QLabel) MousePressEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QLabel) MousePressEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QLabel) ConnectMouseReleaseEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLabel::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QLabel) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QLabel::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQLabelMouseReleaseEvent
func callbackQLabelMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQLabelFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QLabel) MouseReleaseEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QLabel) MouseReleaseEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QLabel) Movie() *gui.QMovie {
	defer qt.Recovering("QLabel::movie")

	if ptr.Pointer() != nil {
		return gui.NewQMovieFromPointer(C.QLabel_Movie(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QLabel::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QLabel) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QLabel::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQLabelPaintEvent
func callbackQLabelPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQLabelFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QLabel) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QLabel::paintEvent")

	if ptr.Pointer() != nil {
		C.QLabel_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QLabel) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QLabel::paintEvent")

	if ptr.Pointer() != nil {
		C.QLabel_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QLabel) Picture() *gui.QPicture {
	defer qt.Recovering("QLabel::picture")

	if ptr.Pointer() != nil {
		return gui.NewQPictureFromPointer(C.QLabel_Picture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) SelectionStart() int {
	defer qt.Recovering("QLabel::selectionStart")

	if ptr.Pointer() != nil {
		return int(C.QLabel_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) SetBuddy(buddy QWidget_ITF) {
	defer qt.Recovering("QLabel::setBuddy")

	if ptr.Pointer() != nil {
		C.QLabel_SetBuddy(ptr.Pointer(), PointerFromQWidget(buddy))
	}
}

func (ptr *QLabel) SetMovie(movie gui.QMovie_ITF) {
	defer qt.Recovering("QLabel::setMovie")

	if ptr.Pointer() != nil {
		C.QLabel_SetMovie(ptr.Pointer(), gui.PointerFromQMovie(movie))
	}
}

func (ptr *QLabel) SetNum(num int) {
	defer qt.Recovering("QLabel::setNum")

	if ptr.Pointer() != nil {
		C.QLabel_SetNum(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QLabel) SetPicture(picture gui.QPicture_ITF) {
	defer qt.Recovering("QLabel::setPicture")

	if ptr.Pointer() != nil {
		C.QLabel_SetPicture(ptr.Pointer(), gui.PointerFromQPicture(picture))
	}
}

func (ptr *QLabel) SetSelection(start int, length int) {
	defer qt.Recovering("QLabel::setSelection")

	if ptr.Pointer() != nil {
		C.QLabel_SetSelection(ptr.Pointer(), C.int(start), C.int(length))
	}
}

func (ptr *QLabel) SizeHint() *core.QSize {
	defer qt.Recovering("QLabel::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLabel_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) DestroyQLabel() {
	defer qt.Recovering("QLabel::~QLabel")

	if ptr.Pointer() != nil {
		C.QLabel_DestroyQLabel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLabel) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QLabel::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QLabel) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QLabel::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQLabelActionEvent
func callbackQLabelActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QLabel) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QLabel::actionEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QLabel) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QLabel::actionEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QLabel) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QLabel::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QLabel) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QLabel::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQLabelDragEnterEvent
func callbackQLabelDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QLabel) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QLabel::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QLabel) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QLabel::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QLabel) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QLabel::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QLabel) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QLabel::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQLabelDragLeaveEvent
func callbackQLabelDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QLabel) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QLabel::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QLabel) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QLabel::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QLabel) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QLabel::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QLabel) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QLabel::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQLabelDragMoveEvent
func callbackQLabelDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QLabel) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QLabel::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QLabel) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QLabel::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QLabel) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QLabel::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QLabel) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QLabel::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQLabelDropEvent
func callbackQLabelDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QLabel) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QLabel::dropEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QLabel) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QLabel::dropEvent")

	if ptr.Pointer() != nil {
		C.QLabel_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QLabel) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLabel::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QLabel) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QLabel::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQLabelEnterEvent
func callbackQLabelEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLabel) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLabel::enterEvent")

	if ptr.Pointer() != nil {
		C.QLabel_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLabel) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLabel::enterEvent")

	if ptr.Pointer() != nil {
		C.QLabel_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLabel) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QLabel::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QLabel) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QLabel::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQLabelHideEvent
func callbackQLabelHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QLabel) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QLabel::hideEvent")

	if ptr.Pointer() != nil {
		C.QLabel_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QLabel) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QLabel::hideEvent")

	if ptr.Pointer() != nil {
		C.QLabel_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QLabel) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLabel::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QLabel) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QLabel::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQLabelLeaveEvent
func callbackQLabelLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLabel) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLabel::leaveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLabel) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLabel::leaveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLabel) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QLabel::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QLabel) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QLabel::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQLabelMoveEvent
func callbackQLabelMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QLabel) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QLabel::moveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QLabel) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QLabel::moveEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QLabel) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QLabel::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QLabel) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QLabel::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQLabelSetVisible
func callbackQLabelSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QLabel::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QLabel) SetVisible(visible bool) {
	defer qt.Recovering("QLabel::setVisible")

	if ptr.Pointer() != nil {
		C.QLabel_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QLabel) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QLabel::setVisible")

	if ptr.Pointer() != nil {
		C.QLabel_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QLabel) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QLabel::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QLabel) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QLabel::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQLabelShowEvent
func callbackQLabelShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QLabel) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QLabel::showEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QLabel) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QLabel::showEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QLabel) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QLabel::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QLabel) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QLabel::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQLabelCloseEvent
func callbackQLabelCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QLabel) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QLabel::closeEvent")

	if ptr.Pointer() != nil {
		C.QLabel_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QLabel) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QLabel::closeEvent")

	if ptr.Pointer() != nil {
		C.QLabel_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QLabel) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QLabel::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QLabel) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QLabel::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQLabelInitPainter
func callbackQLabelInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQLabelFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QLabel) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QLabel::initPainter")

	if ptr.Pointer() != nil {
		C.QLabel_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QLabel) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QLabel::initPainter")

	if ptr.Pointer() != nil {
		C.QLabel_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QLabel) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QLabel::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QLabel) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QLabel::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQLabelInputMethodEvent
func callbackQLabelInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QLabel) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QLabel::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QLabel_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QLabel) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QLabel::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QLabel_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QLabel) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QLabel::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QLabel) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QLabel::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQLabelKeyReleaseEvent
func callbackQLabelKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLabel) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLabel::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLabel_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLabel) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLabel::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLabel_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLabel) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLabel::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QLabel) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QLabel::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQLabelMouseDoubleClickEvent
func callbackQLabelMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QLabel) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QLabel) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLabel::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QLabel_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QLabel) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QLabel::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QLabel) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QLabel::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQLabelResizeEvent
func callbackQLabelResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QLabel) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QLabel::resizeEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QLabel) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QLabel::resizeEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QLabel) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QLabel::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QLabel) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QLabel::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQLabelTabletEvent
func callbackQLabelTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QLabel) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QLabel::tabletEvent")

	if ptr.Pointer() != nil {
		C.QLabel_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QLabel) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QLabel::tabletEvent")

	if ptr.Pointer() != nil {
		C.QLabel_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QLabel) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QLabel::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QLabel) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QLabel::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQLabelWheelEvent
func callbackQLabelWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QLabel) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QLabel::wheelEvent")

	if ptr.Pointer() != nil {
		C.QLabel_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QLabel) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QLabel::wheelEvent")

	if ptr.Pointer() != nil {
		C.QLabel_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QLabel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLabel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLabel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLabel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLabelTimerEvent
func callbackQLabelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLabel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLabel::timerEvent")

	if ptr.Pointer() != nil {
		C.QLabel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLabel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLabel::timerEvent")

	if ptr.Pointer() != nil {
		C.QLabel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLabel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLabel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLabel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLabel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLabelChildEvent
func callbackQLabelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLabel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLabel::childEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLabel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLabel::childEvent")

	if ptr.Pointer() != nil {
		C.QLabel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLabel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLabel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLabel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLabel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLabelCustomEvent
func callbackQLabelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLabel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLabelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLabel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLabel::customEvent")

	if ptr.Pointer() != nil {
		C.QLabel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLabel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLabel::customEvent")

	if ptr.Pointer() != nil {
		C.QLabel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
