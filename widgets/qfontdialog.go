package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFontDialog struct {
	QDialog
}

type QFontDialog_ITF interface {
	QDialog_ITF
	QFontDialog_PTR() *QFontDialog
}

func PointerFromQFontDialog(ptr QFontDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontDialog_PTR().Pointer()
	}
	return nil
}

func NewQFontDialogFromPointer(ptr unsafe.Pointer) *QFontDialog {
	var n = new(QFontDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFontDialog_") {
		n.SetObjectName("QFontDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QFontDialog) QFontDialog_PTR() *QFontDialog {
	return ptr
}

//QFontDialog::FontDialogOption
type QFontDialog__FontDialogOption int64

const (
	QFontDialog__NoButtons           = QFontDialog__FontDialogOption(0x00000001)
	QFontDialog__DontUseNativeDialog = QFontDialog__FontDialogOption(0x00000002)
	QFontDialog__ScalableFonts       = QFontDialog__FontDialogOption(0x00000004)
	QFontDialog__NonScalableFonts    = QFontDialog__FontDialogOption(0x00000008)
	QFontDialog__MonospacedFonts     = QFontDialog__FontDialogOption(0x00000010)
	QFontDialog__ProportionalFonts   = QFontDialog__FontDialogOption(0x00000020)
)

func (ptr *QFontDialog) Options() QFontDialog__FontDialogOption {
	defer qt.Recovering("QFontDialog::options")

	if ptr.Pointer() != nil {
		return QFontDialog__FontDialogOption(C.QFontDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontDialog) SetOptions(options QFontDialog__FontDialogOption) {
	defer qt.Recovering("QFontDialog::setOptions")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQFontDialog(parent QWidget_ITF) *QFontDialog {
	defer qt.Recovering("QFontDialog::QFontDialog")

	return NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog(PointerFromQWidget(parent)))
}

func NewQFontDialog2(initial gui.QFont_ITF, parent QWidget_ITF) *QFontDialog {
	defer qt.Recovering("QFontDialog::QFontDialog")

	return NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog2(gui.PointerFromQFont(initial), PointerFromQWidget(parent)))
}

func (ptr *QFontDialog) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QFontDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFontDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFontDialogChangeEvent
func callbackQFontDialogChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectDone(f func(result int)) {
	defer qt.Recovering("connect QFontDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QFontDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QFontDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQFontDialogDone
func callbackQFontDialogDone(ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QFontDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
		return true
	}
	return false

}

func (ptr *QFontDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QFontDialog::open")

	if ptr.Pointer() != nil {
		C.QFontDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QFontDialog) SetCurrentFont(font gui.QFont_ITF) {
	defer qt.Recovering("QFontDialog::setCurrentFont")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QFontDialog) SetOption(option QFontDialog__FontDialogOption, on bool) {
	defer qt.Recovering("QFontDialog::setOption")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFontDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFontDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFontDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFontDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFontDialogSetVisible
func callbackQFontDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFontDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFontDialog) TestOption(option QFontDialog__FontDialogOption) bool {
	defer qt.Recovering("QFontDialog::testOption")

	if ptr.Pointer() != nil {
		return C.QFontDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QFontDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QFontDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QFontDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QFontDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQFontDialogAccept
func callbackQFontDialogAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QFontDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QFontDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QFontDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQFontDialogCloseEvent
func callbackQFontDialogCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QFontDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QFontDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQFontDialogContextMenuEvent
func callbackQFontDialogContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFontDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QFontDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQFontDialogKeyPressEvent
func callbackQFontDialogKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QFontDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QFontDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QFontDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQFontDialogOpen
func callbackQFontDialogOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QFontDialog::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QFontDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QFontDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QFontDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQFontDialogReject
func callbackQFontDialogReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QFontDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QFontDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QFontDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQFontDialogResizeEvent
func callbackQFontDialogResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QFontDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QFontDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQFontDialogShowEvent
func callbackQFontDialogShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QFontDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QFontDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQFontDialogActionEvent
func callbackQFontDialogActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QFontDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QFontDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQFontDialogDragEnterEvent
func callbackQFontDialogDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QFontDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QFontDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQFontDialogDragLeaveEvent
func callbackQFontDialogDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QFontDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QFontDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQFontDialogDragMoveEvent
func callbackQFontDialogDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QFontDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QFontDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQFontDialogDropEvent
func callbackQFontDialogDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFontDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QFontDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQFontDialogEnterEvent
func callbackQFontDialogEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFontDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QFontDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQFontDialogFocusInEvent
func callbackQFontDialogFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFontDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QFontDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQFontDialogFocusOutEvent
func callbackQFontDialogFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QFontDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QFontDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQFontDialogHideEvent
func callbackQFontDialogHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFontDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QFontDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQFontDialogLeaveEvent
func callbackQFontDialogLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QFontDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QFontDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQFontDialogMoveEvent
func callbackQFontDialogMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFontDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFontDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFontDialogPaintEvent
func callbackQFontDialogPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QFontDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QFontDialog) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QFontDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQFontDialogInitPainter
func callbackQFontDialogInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QFontDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QFontDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQFontDialogInputMethodEvent
func callbackQFontDialogInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFontDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QFontDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQFontDialogKeyReleaseEvent
func callbackQFontDialogKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QFontDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQFontDialogMouseDoubleClickEvent
func callbackQFontDialogMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QFontDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQFontDialogMouseMoveEvent
func callbackQFontDialogMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QFontDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQFontDialogMousePressEvent
func callbackQFontDialogMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFontDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QFontDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQFontDialogMouseReleaseEvent
func callbackQFontDialogMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QFontDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QFontDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQFontDialogTabletEvent
func callbackQFontDialogTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QFontDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QFontDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQFontDialogWheelEvent
func callbackQFontDialogWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFontDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFontDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFontDialogTimerEvent
func callbackQFontDialogTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFontDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFontDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFontDialogChildEvent
func callbackQFontDialogChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFontDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFontDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFontDialogCustomEvent
func callbackQFontDialogCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
