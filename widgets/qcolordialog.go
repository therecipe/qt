package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QColorDialog struct {
	QDialog
}

type QColorDialog_ITF interface {
	QDialog_ITF
	QColorDialog_PTR() *QColorDialog
}

func PointerFromQColorDialog(ptr QColorDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColorDialog_PTR().Pointer()
	}
	return nil
}

func NewQColorDialogFromPointer(ptr unsafe.Pointer) *QColorDialog {
	var n = new(QColorDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QColorDialog_") {
		n.SetObjectName("QColorDialog_" + qt.Identifier())
	}
	return n
}

func (ptr *QColorDialog) QColorDialog_PTR() *QColorDialog {
	return ptr
}

//QColorDialog::ColorDialogOption
type QColorDialog__ColorDialogOption int64

const (
	QColorDialog__ShowAlphaChannel    = QColorDialog__ColorDialogOption(0x00000001)
	QColorDialog__NoButtons           = QColorDialog__ColorDialogOption(0x00000002)
	QColorDialog__DontUseNativeDialog = QColorDialog__ColorDialogOption(0x00000004)
)

func (ptr *QColorDialog) CurrentColor() *gui.QColor {
	defer qt.Recovering("QColorDialog::currentColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QColorDialog_CurrentColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColorDialog) Options() QColorDialog__ColorDialogOption {
	defer qt.Recovering("QColorDialog::options")

	if ptr.Pointer() != nil {
		return QColorDialog__ColorDialogOption(C.QColorDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColorDialog) SetCurrentColor(color gui.QColor_ITF) {
	defer qt.Recovering("QColorDialog::setCurrentColor")

	if ptr.Pointer() != nil {
		C.QColorDialog_SetCurrentColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QColorDialog) SetOptions(options QColorDialog__ColorDialogOption) {
	defer qt.Recovering("QColorDialog::setOptions")

	if ptr.Pointer() != nil {
		C.QColorDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQColorDialog(parent QWidget_ITF) *QColorDialog {
	defer qt.Recovering("QColorDialog::QColorDialog")

	return NewQColorDialogFromPointer(C.QColorDialog_NewQColorDialog(PointerFromQWidget(parent)))
}

func NewQColorDialog2(initial gui.QColor_ITF, parent QWidget_ITF) *QColorDialog {
	defer qt.Recovering("QColorDialog::QColorDialog")

	return NewQColorDialogFromPointer(C.QColorDialog_NewQColorDialog2(gui.PointerFromQColor(initial), PointerFromQWidget(parent)))
}

func (ptr *QColorDialog) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QColorDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QColorDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQColorDialogChangeEvent
func callbackQColorDialogChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectColorSelected(f func(color *gui.QColor)) {
	defer qt.Recovering("connect QColorDialog::colorSelected")

	if ptr.Pointer() != nil {
		C.QColorDialog_ConnectColorSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorSelected", f)
	}
}

func (ptr *QColorDialog) DisconnectColorSelected() {
	defer qt.Recovering("disconnect QColorDialog::colorSelected")

	if ptr.Pointer() != nil {
		C.QColorDialog_DisconnectColorSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorSelected")
	}
}

//export callbackQColorDialogColorSelected
func callbackQColorDialogColorSelected(ptrName *C.char, color unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::colorSelected")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorSelected"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QColorDialog) ConnectCurrentColorChanged(f func(color *gui.QColor)) {
	defer qt.Recovering("connect QColorDialog::currentColorChanged")

	if ptr.Pointer() != nil {
		C.QColorDialog_ConnectCurrentColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentColorChanged", f)
	}
}

func (ptr *QColorDialog) DisconnectCurrentColorChanged() {
	defer qt.Recovering("disconnect QColorDialog::currentColorChanged")

	if ptr.Pointer() != nil {
		C.QColorDialog_DisconnectCurrentColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentColorChanged")
	}
}

//export callbackQColorDialogCurrentColorChanged
func callbackQColorDialogCurrentColorChanged(ptrName *C.char, color unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::currentColorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func QColorDialog_CustomColor(index int) *gui.QColor {
	defer qt.Recovering("QColorDialog::customColor")

	return gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_CustomColor(C.int(index)))
}

func QColorDialog_CustomCount() int {
	defer qt.Recovering("QColorDialog::customCount")

	return int(C.QColorDialog_QColorDialog_CustomCount())
}

func (ptr *QColorDialog) ConnectDone(f func(result int)) {
	defer qt.Recovering("connect QColorDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QColorDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QColorDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQColorDialogDone
func callbackQColorDialogDone(ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QColorDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
		return true
	}
	return false

}

func QColorDialog_GetColor(initial gui.QColor_ITF, parent QWidget_ITF, title string, options QColorDialog__ColorDialogOption) *gui.QColor {
	defer qt.Recovering("QColorDialog::getColor")

	return gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_GetColor(gui.PointerFromQColor(initial), PointerFromQWidget(parent), C.CString(title), C.int(options)))
}

func (ptr *QColorDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QColorDialog::open")

	if ptr.Pointer() != nil {
		C.QColorDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QColorDialog) SelectedColor() *gui.QColor {
	defer qt.Recovering("QColorDialog::selectedColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QColorDialog_SelectedColor(ptr.Pointer()))
	}
	return nil
}

func QColorDialog_SetCustomColor(index int, color gui.QColor_ITF) {
	defer qt.Recovering("QColorDialog::setCustomColor")

	C.QColorDialog_QColorDialog_SetCustomColor(C.int(index), gui.PointerFromQColor(color))
}

func (ptr *QColorDialog) SetOption(option QColorDialog__ColorDialogOption, on bool) {
	defer qt.Recovering("QColorDialog::setOption")

	if ptr.Pointer() != nil {
		C.QColorDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func QColorDialog_SetStandardColor(index int, color gui.QColor_ITF) {
	defer qt.Recovering("QColorDialog::setStandardColor")

	C.QColorDialog_QColorDialog_SetStandardColor(C.int(index), gui.PointerFromQColor(color))
}

func (ptr *QColorDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QColorDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QColorDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QColorDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQColorDialogSetVisible
func callbackQColorDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QColorDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func QColorDialog_StandardColor(index int) *gui.QColor {
	defer qt.Recovering("QColorDialog::standardColor")

	return gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_StandardColor(C.int(index)))
}

func (ptr *QColorDialog) TestOption(option QColorDialog__ColorDialogOption) bool {
	defer qt.Recovering("QColorDialog::testOption")

	if ptr.Pointer() != nil {
		return C.QColorDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QColorDialog) DestroyQColorDialog() {
	defer qt.Recovering("QColorDialog::~QColorDialog")

	if ptr.Pointer() != nil {
		C.QColorDialog_DestroyQColorDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QColorDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QColorDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QColorDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QColorDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQColorDialogAccept
func callbackQColorDialogAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QColorDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QColorDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QColorDialog::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQColorDialogCloseEvent
func callbackQColorDialogCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QColorDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QColorDialog::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQColorDialogContextMenuEvent
func callbackQColorDialogContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QColorDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QColorDialog::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQColorDialogKeyPressEvent
func callbackQColorDialogKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectOpen(f func()) {
	defer qt.Recovering("connect QColorDialog::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QColorDialog) DisconnectOpen() {
	defer qt.Recovering("disconnect QColorDialog::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQColorDialogOpen
func callbackQColorDialogOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QColorDialog::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectReject(f func()) {
	defer qt.Recovering("connect QColorDialog::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QColorDialog) DisconnectReject() {
	defer qt.Recovering("disconnect QColorDialog::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQColorDialogReject
func callbackQColorDialogReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QColorDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QColorDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QColorDialog::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQColorDialogResizeEvent
func callbackQColorDialogResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QColorDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QColorDialog::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQColorDialogShowEvent
func callbackQColorDialogShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QColorDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QColorDialog::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQColorDialogActionEvent
func callbackQColorDialogActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QColorDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QColorDialog::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQColorDialogDragEnterEvent
func callbackQColorDialogDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QColorDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QColorDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQColorDialogDragLeaveEvent
func callbackQColorDialogDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QColorDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QColorDialog::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQColorDialogDragMoveEvent
func callbackQColorDialogDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QColorDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QColorDialog::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQColorDialogDropEvent
func callbackQColorDialogDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QColorDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QColorDialog::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQColorDialogEnterEvent
func callbackQColorDialogEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QColorDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QColorDialog::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQColorDialogFocusInEvent
func callbackQColorDialogFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QColorDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QColorDialog::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQColorDialogFocusOutEvent
func callbackQColorDialogFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QColorDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QColorDialog::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQColorDialogHideEvent
func callbackQColorDialogHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QColorDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QColorDialog::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQColorDialogLeaveEvent
func callbackQColorDialogLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QColorDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QColorDialog::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQColorDialogMoveEvent
func callbackQColorDialogMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QColorDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QColorDialog::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQColorDialogPaintEvent
func callbackQColorDialogPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QColorDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QColorDialog) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QColorDialog::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQColorDialogInitPainter
func callbackQColorDialogInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QColorDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QColorDialog::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQColorDialogInputMethodEvent
func callbackQColorDialogInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QColorDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QColorDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQColorDialogKeyReleaseEvent
func callbackQColorDialogKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColorDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QColorDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQColorDialogMouseDoubleClickEvent
func callbackQColorDialogMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColorDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QColorDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQColorDialogMouseMoveEvent
func callbackQColorDialogMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColorDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QColorDialog::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQColorDialogMousePressEvent
func callbackQColorDialogMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColorDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QColorDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQColorDialogMouseReleaseEvent
func callbackQColorDialogMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QColorDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QColorDialog::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQColorDialogTabletEvent
func callbackQColorDialogTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QColorDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QColorDialog::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQColorDialogWheelEvent
func callbackQColorDialogWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QColorDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QColorDialog::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQColorDialogTimerEvent
func callbackQColorDialogTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QColorDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QColorDialog::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQColorDialogChildEvent
func callbackQColorDialogChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColorDialog) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QColorDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QColorDialog) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QColorDialog::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQColorDialogCustomEvent
func callbackQColorDialogCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColorDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
