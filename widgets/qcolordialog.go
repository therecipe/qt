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
func callbackQColorDialogChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQColorDialogFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QColorDialog) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QColorDialog) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::changeEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQColorDialogColorSelected(ptr unsafe.Pointer, ptrName *C.char, color unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::colorSelected")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorSelected"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QColorDialog) ColorSelected(color gui.QColor_ITF) {
	defer qt.Recovering("QColorDialog::colorSelected")

	if ptr.Pointer() != nil {
		C.QColorDialog_ColorSelected(ptr.Pointer(), gui.PointerFromQColor(color))
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
func callbackQColorDialogCurrentColorChanged(ptr unsafe.Pointer, ptrName *C.char, color unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::currentColorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentColorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QColorDialog) CurrentColorChanged(color gui.QColor_ITF) {
	defer qt.Recovering("QColorDialog::currentColorChanged")

	if ptr.Pointer() != nil {
		C.QColorDialog_CurrentColorChanged(ptr.Pointer(), gui.PointerFromQColor(color))
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
func callbackQColorDialogDone(ptr unsafe.Pointer, ptrName *C.char, result C.int) {
	defer qt.Recovering("callback QColorDialog::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
	} else {
		NewQColorDialogFromPointer(ptr).DoneDefault(int(result))
	}
}

func (ptr *QColorDialog) Done(result int) {
	defer qt.Recovering("QColorDialog::done")

	if ptr.Pointer() != nil {
		C.QColorDialog_Done(ptr.Pointer(), C.int(result))
	}
}

func (ptr *QColorDialog) DoneDefault(result int) {
	defer qt.Recovering("QColorDialog::done")

	if ptr.Pointer() != nil {
		C.QColorDialog_DoneDefault(ptr.Pointer(), C.int(result))
	}
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
func callbackQColorDialogSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QColorDialog::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQColorDialogFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QColorDialog) SetVisible(visible bool) {
	defer qt.Recovering("QColorDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QColorDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QColorDialog) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QColorDialog::setVisible")

	if ptr.Pointer() != nil {
		C.QColorDialog_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQColorDialogAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QColorDialog::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColorDialog) Accept() {
	defer qt.Recovering("QColorDialog::accept")

	if ptr.Pointer() != nil {
		C.QColorDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QColorDialog) AcceptDefault() {
	defer qt.Recovering("QColorDialog::accept")

	if ptr.Pointer() != nil {
		C.QColorDialog_AcceptDefault(ptr.Pointer())
	}
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
func callbackQColorDialogCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQColorDialogFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QColorDialog) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QColorDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QColorDialog) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QColorDialog::closeEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
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
func callbackQColorDialogContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQColorDialogFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QColorDialog) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QColorDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QColorDialog) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QColorDialog::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
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
func callbackQColorDialogKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQColorDialogFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QColorDialog) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColorDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QColorDialog) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColorDialog::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQColorDialogReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QColorDialog::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColorDialog) Reject() {
	defer qt.Recovering("QColorDialog::reject")

	if ptr.Pointer() != nil {
		C.QColorDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QColorDialog) RejectDefault() {
	defer qt.Recovering("QColorDialog::reject")

	if ptr.Pointer() != nil {
		C.QColorDialog_RejectDefault(ptr.Pointer())
	}
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
func callbackQColorDialogResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQColorDialogFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QColorDialog) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QColorDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QColorDialog) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QColorDialog::resizeEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
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
func callbackQColorDialogShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QColorDialog) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QColorDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QColorDialog) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QColorDialog::showEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQColorDialogActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QColorDialog) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QColorDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QColorDialog) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QColorDialog::actionEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQColorDialogDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QColorDialog) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QColorDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QColorDialog) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QColorDialog::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQColorDialogDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QColorDialog) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QColorDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QColorDialog) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QColorDialog::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQColorDialogDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QColorDialog) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QColorDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QColorDialog) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QColorDialog::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQColorDialogDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QColorDialog) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QColorDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QColorDialog) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QColorDialog::dropEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQColorDialogEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QColorDialog) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColorDialog) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::enterEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQColorDialogFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QColorDialog) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColorDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QColorDialog) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColorDialog::focusInEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQColorDialogFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QColorDialog) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColorDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QColorDialog) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColorDialog::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQColorDialogHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QColorDialog) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QColorDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QColorDialog) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QColorDialog::hideEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQColorDialogLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QColorDialog) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColorDialog) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::leaveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQColorDialogMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QColorDialog) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QColorDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QColorDialog) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QColorDialog::moveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQColorDialogPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QColorDialog) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QColorDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QColorDialog) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QColorDialog::paintEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQColorDialogInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQColorDialogFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QColorDialog) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QColorDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QColorDialog_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QColorDialog) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QColorDialog::initPainter")

	if ptr.Pointer() != nil {
		C.QColorDialog_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQColorDialogInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QColorDialog) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QColorDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QColorDialog) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QColorDialog::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQColorDialogKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QColorDialog) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColorDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QColorDialog) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColorDialog::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQColorDialogMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColorDialog) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColorDialog) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQColorDialogMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColorDialog) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColorDialog) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQColorDialogMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColorDialog) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColorDialog) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQColorDialogMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColorDialog) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColorDialog) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColorDialog::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQColorDialogTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QColorDialog) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QColorDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QColorDialog) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QColorDialog::tabletEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQColorDialogWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QColorDialog) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QColorDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QColorDialog) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QColorDialog::wheelEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQColorDialogTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QColorDialog) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QColorDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QColorDialog) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QColorDialog::timerEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQColorDialogChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QColorDialog) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QColorDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QColorDialog) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QColorDialog::childEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQColorDialogCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColorDialog::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQColorDialogFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QColorDialog) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColorDialog) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QColorDialog::customEvent")

	if ptr.Pointer() != nil {
		C.QColorDialog_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
