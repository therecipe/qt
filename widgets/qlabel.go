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
func callbackQLabelChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLabelContextMenuEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::contextMenuEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
		return true
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
func callbackQLabelFocusInEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::focusInEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
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
func callbackQLabelFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::focusOutEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLabelKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLabelLinkActivated(ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QLabel::linkActivated")

	var signal = qt.GetSignal(C.GoString(ptrName), "linkActivated")
	if signal != nil {
		signal.(func(string))(C.GoString(link))
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
func callbackQLabelLinkHovered(ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QLabel::linkHovered")

	var signal = qt.GetSignal(C.GoString(ptrName), "linkHovered")
	if signal != nil {
		signal.(func(string))(C.GoString(link))
	}

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
func callbackQLabelMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLabelMousePressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLabelMouseReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQLabelPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

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

func (ptr *QLabel) DestroyQLabel() {
	defer qt.Recovering("QLabel::~QLabel")

	if ptr.Pointer() != nil {
		C.QLabel_DestroyQLabel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
