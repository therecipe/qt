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
func callbackQFontDialogChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQFontDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QFontDialog) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QFontDialog) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQFontDialogDone(ptr unsafe.Pointer, ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QFontDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
		return true
	}
	return false

}

func (ptr *QFontDialog) Done(result int) {
	defer qt.Recovering("QFontDialog::done")

	if ptr.Pointer() != nil {
		C.QFontDialog_Done(ptr.Pointer(), C.int(result))
	}
}

func (ptr *QFontDialog) DoneDefault(result int) {
	defer qt.Recovering("QFontDialog::done")

	if ptr.Pointer() != nil {
		C.QFontDialog_DoneDefault(ptr.Pointer(), C.int(result))
	}
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
func callbackQFontDialogSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFontDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFontDialog) SetVisible(visible bool) {
	defer qt.Recovering("QFontDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFontDialog) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QFontDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQFontDialogAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QFontDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontDialog) Accept() {
	defer qt.Recovering("QFontDialog::accept")

	if ptr.Pointer() != nil {
		C.QFontDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QFontDialog) AcceptDefault() {
	defer qt.Recovering("QFontDialog::accept")

	if ptr.Pointer() != nil {
		C.QFontDialog_AcceptDefault(ptr.Pointer())
	}
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
func callbackQFontDialogCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQFontDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QFontDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFontDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QFontDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFontDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
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
func callbackQFontDialogContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQFontDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QFontDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFontDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QFontDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFontDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
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
func callbackQFontDialogKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQFontDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QFontDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QFontDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQFontDialogReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QFontDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QFontDialog) Reject() {
	defer qt.Recovering("QFontDialog::reject")

	if ptr.Pointer() != nil {
		C.QFontDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QFontDialog) RejectDefault() {
	defer qt.Recovering("QFontDialog::reject")

	if ptr.Pointer() != nil {
		C.QFontDialog_RejectDefault(ptr.Pointer())
	}
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
func callbackQFontDialogResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQFontDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QFontDialog) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFontDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QFontDialog) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFontDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
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
func callbackQFontDialogShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QFontDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QFontDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QFontDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QFontDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQFontDialogActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QFontDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFontDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QFontDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFontDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQFontDialogDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QFontDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFontDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QFontDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFontDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQFontDialogDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QFontDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFontDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QFontDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFontDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQFontDialogDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QFontDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFontDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QFontDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFontDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQFontDialogDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QFontDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFontDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QFontDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFontDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQFontDialogEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFontDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFontDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQFontDialogFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QFontDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFontDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQFontDialogFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QFontDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFontDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFontDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQFontDialogHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QFontDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QFontDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QFontDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QFontDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQFontDialogLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFontDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFontDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQFontDialogMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QFontDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFontDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QFontDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFontDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQFontDialogPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QFontDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFontDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QFontDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFontDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQFontDialogInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQFontDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QFontDialog) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFontDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QFontDialog_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QFontDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFontDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QFontDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQFontDialogInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QFontDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFontDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QFontDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFontDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQFontDialogKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QFontDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFontDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFontDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQFontDialogMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFontDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFontDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQFontDialogMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFontDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFontDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQFontDialogMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFontDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFontDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQFontDialogMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFontDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFontDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFontDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQFontDialogTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QFontDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFontDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QFontDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFontDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQFontDialogWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QFontDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFontDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QFontDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFontDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQFontDialogTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFontDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFontDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFontDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFontDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQFontDialogChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QFontDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFontDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFontDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFontDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQFontDialogCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFontDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFontDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFontDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFontDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFontDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QFontDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
