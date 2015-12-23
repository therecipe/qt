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

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "linkHovered"); signal != nil {
		signal.(func(string))(C.GoString(link))
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
func callbackQLabelMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
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
func callbackQLabelActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QLabel::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQLabelShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQLabelInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQLabelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLabel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
