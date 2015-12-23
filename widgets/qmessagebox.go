package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMessageBox struct {
	QDialog
}

type QMessageBox_ITF interface {
	QDialog_ITF
	QMessageBox_PTR() *QMessageBox
}

func PointerFromQMessageBox(ptr QMessageBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageBox_PTR().Pointer()
	}
	return nil
}

func NewQMessageBoxFromPointer(ptr unsafe.Pointer) *QMessageBox {
	var n = new(QMessageBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMessageBox_") {
		n.SetObjectName("QMessageBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QMessageBox) QMessageBox_PTR() *QMessageBox {
	return ptr
}

//QMessageBox::ButtonRole
type QMessageBox__ButtonRole int64

const (
	QMessageBox__InvalidRole     = QMessageBox__ButtonRole(-1)
	QMessageBox__AcceptRole      = QMessageBox__ButtonRole(0)
	QMessageBox__RejectRole      = QMessageBox__ButtonRole(1)
	QMessageBox__DestructiveRole = QMessageBox__ButtonRole(2)
	QMessageBox__ActionRole      = QMessageBox__ButtonRole(3)
	QMessageBox__HelpRole        = QMessageBox__ButtonRole(4)
	QMessageBox__YesRole         = QMessageBox__ButtonRole(5)
	QMessageBox__NoRole          = QMessageBox__ButtonRole(6)
	QMessageBox__ResetRole       = QMessageBox__ButtonRole(7)
	QMessageBox__ApplyRole       = QMessageBox__ButtonRole(8)
	QMessageBox__NRoles          = QMessageBox__ButtonRole(9)
)

//QMessageBox::Icon
type QMessageBox__Icon int64

const (
	QMessageBox__NoIcon      = QMessageBox__Icon(0)
	QMessageBox__Information = QMessageBox__Icon(1)
	QMessageBox__Warning     = QMessageBox__Icon(2)
	QMessageBox__Critical    = QMessageBox__Icon(3)
	QMessageBox__Question    = QMessageBox__Icon(4)
)

//QMessageBox::StandardButton
type QMessageBox__StandardButton int64

var (
	QMessageBox__NoButton        = QMessageBox__StandardButton(0x00000000)
	QMessageBox__Ok              = QMessageBox__StandardButton(0x00000400)
	QMessageBox__Save            = QMessageBox__StandardButton(0x00000800)
	QMessageBox__SaveAll         = QMessageBox__StandardButton(0x00001000)
	QMessageBox__Open            = QMessageBox__StandardButton(0x00002000)
	QMessageBox__Yes             = QMessageBox__StandardButton(0x00004000)
	QMessageBox__YesToAll        = QMessageBox__StandardButton(0x00008000)
	QMessageBox__No              = QMessageBox__StandardButton(0x00010000)
	QMessageBox__NoToAll         = QMessageBox__StandardButton(0x00020000)
	QMessageBox__Abort           = QMessageBox__StandardButton(0x00040000)
	QMessageBox__Retry           = QMessageBox__StandardButton(0x00080000)
	QMessageBox__Ignore          = QMessageBox__StandardButton(0x00100000)
	QMessageBox__Close           = QMessageBox__StandardButton(0x00200000)
	QMessageBox__Cancel          = QMessageBox__StandardButton(0x00400000)
	QMessageBox__Discard         = QMessageBox__StandardButton(0x00800000)
	QMessageBox__Help            = QMessageBox__StandardButton(0x01000000)
	QMessageBox__Apply           = QMessageBox__StandardButton(0x02000000)
	QMessageBox__Reset           = QMessageBox__StandardButton(0x04000000)
	QMessageBox__RestoreDefaults = QMessageBox__StandardButton(0x08000000)
	QMessageBox__FirstButton     = QMessageBox__StandardButton(QMessageBox__Ok)
	QMessageBox__LastButton      = QMessageBox__StandardButton(QMessageBox__RestoreDefaults)
	QMessageBox__YesAll          = QMessageBox__StandardButton(QMessageBox__YesToAll)
	QMessageBox__NoAll           = QMessageBox__StandardButton(QMessageBox__NoToAll)
	QMessageBox__Default         = QMessageBox__StandardButton(0x00000100)
	QMessageBox__Escape          = QMessageBox__StandardButton(0x00000200)
	QMessageBox__FlagMask        = QMessageBox__StandardButton(0x00000300)
	QMessageBox__ButtonMask      = QMessageBox__StandardButton(C.QMessageBox_ButtonMask_Type())
)

func (ptr *QMessageBox) DetailedText() string {
	defer qt.Recovering("QMessageBox::detailedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMessageBox_DetailedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMessageBox) Icon() QMessageBox__Icon {
	defer qt.Recovering("QMessageBox::icon")

	if ptr.Pointer() != nil {
		return QMessageBox__Icon(C.QMessageBox_Icon(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMessageBox) InformativeText() string {
	defer qt.Recovering("QMessageBox::informativeText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMessageBox_InformativeText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMessageBox) SetDetailedText(text string) {
	defer qt.Recovering("QMessageBox::setDetailedText")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetDetailedText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMessageBox) SetIcon(v QMessageBox__Icon) {
	defer qt.Recovering("QMessageBox::setIcon")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetIcon(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QMessageBox) SetIconPixmap(pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QMessageBox::setIconPixmap")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetIconPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QMessageBox) SetInformativeText(text string) {
	defer qt.Recovering("QMessageBox::setInformativeText")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetInformativeText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMessageBox) SetStandardButtons(buttons QMessageBox__StandardButton) {
	defer qt.Recovering("QMessageBox::setStandardButtons")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetStandardButtons(ptr.Pointer(), C.int(buttons))
	}
}

func (ptr *QMessageBox) SetText(text string) {
	defer qt.Recovering("QMessageBox::setText")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMessageBox) SetTextFormat(format core.Qt__TextFormat) {
	defer qt.Recovering("QMessageBox::setTextFormat")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetTextFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QMessageBox) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer qt.Recovering("QMessageBox::setTextInteractionFlags")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QMessageBox) StandardButtons() QMessageBox__StandardButton {
	defer qt.Recovering("QMessageBox::standardButtons")

	if ptr.Pointer() != nil {
		return QMessageBox__StandardButton(C.QMessageBox_StandardButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMessageBox) Text() string {
	defer qt.Recovering("QMessageBox::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMessageBox_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMessageBox) TextFormat() core.Qt__TextFormat {
	defer qt.Recovering("QMessageBox::textFormat")

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QMessageBox_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMessageBox) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer qt.Recovering("QMessageBox::textInteractionFlags")

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QMessageBox_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func NewQMessageBox2(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton, parent QWidget_ITF, f core.Qt__WindowType) *QMessageBox {
	defer qt.Recovering("QMessageBox::QMessageBox")

	return NewQMessageBoxFromPointer(C.QMessageBox_NewQMessageBox2(C.int(icon), C.CString(title), C.CString(text), C.int(buttons), PointerFromQWidget(parent), C.int(f)))
}

func NewQMessageBox(parent QWidget_ITF) *QMessageBox {
	defer qt.Recovering("QMessageBox::QMessageBox")

	return NewQMessageBoxFromPointer(C.QMessageBox_NewQMessageBox(PointerFromQWidget(parent)))
}

func QMessageBox_About(parent QWidget_ITF, title string, text string) {
	defer qt.Recovering("QMessageBox::about")

	C.QMessageBox_QMessageBox_About(PointerFromQWidget(parent), C.CString(title), C.CString(text))
}

func QMessageBox_AboutQt(parent QWidget_ITF, title string) {
	defer qt.Recovering("QMessageBox::aboutQt")

	C.QMessageBox_QMessageBox_AboutQt(PointerFromQWidget(parent), C.CString(title))
}

func (ptr *QMessageBox) AddButton3(button QMessageBox__StandardButton) *QPushButton {
	defer qt.Recovering("QMessageBox::addButton")

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QMessageBox_AddButton3(ptr.Pointer(), C.int(button)))
	}
	return nil
}

func (ptr *QMessageBox) AddButton2(text string, role QMessageBox__ButtonRole) *QPushButton {
	defer qt.Recovering("QMessageBox::addButton")

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QMessageBox_AddButton2(ptr.Pointer(), C.CString(text), C.int(role)))
	}
	return nil
}

func (ptr *QMessageBox) AddButton(button QAbstractButton_ITF, role QMessageBox__ButtonRole) {
	defer qt.Recovering("QMessageBox::addButton")

	if ptr.Pointer() != nil {
		C.QMessageBox_AddButton(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(role))
	}
}

func (ptr *QMessageBox) Button(which QMessageBox__StandardButton) *QAbstractButton {
	defer qt.Recovering("QMessageBox::button")

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QMessageBox_Button(ptr.Pointer(), C.int(which)))
	}
	return nil
}

func (ptr *QMessageBox) ConnectButtonClicked(f func(button *QAbstractButton)) {
	defer qt.Recovering("connect QMessageBox::buttonClicked")

	if ptr.Pointer() != nil {
		C.QMessageBox_ConnectButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonClicked", f)
	}
}

func (ptr *QMessageBox) DisconnectButtonClicked() {
	defer qt.Recovering("disconnect QMessageBox::buttonClicked")

	if ptr.Pointer() != nil {
		C.QMessageBox_DisconnectButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonClicked")
	}
}

//export callbackQMessageBoxButtonClicked
func callbackQMessageBoxButtonClicked(ptrName *C.char, button unsafe.Pointer) {
	defer qt.Recovering("callback QMessageBox::buttonClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonClicked"); signal != nil {
		signal.(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
	}

}

func (ptr *QMessageBox) ButtonRole(button QAbstractButton_ITF) QMessageBox__ButtonRole {
	defer qt.Recovering("QMessageBox::buttonRole")

	if ptr.Pointer() != nil {
		return QMessageBox__ButtonRole(C.QMessageBox_ButtonRole(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QMessageBox) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QMessageBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMessageBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMessageBoxChangeEvent
func callbackQMessageBoxChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QMessageBox) CheckBox() *QCheckBox {
	defer qt.Recovering("QMessageBox::checkBox")

	if ptr.Pointer() != nil {
		return NewQCheckBoxFromPointer(C.QMessageBox_CheckBox(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMessageBox) ClickedButton() *QAbstractButton {
	defer qt.Recovering("QMessageBox::clickedButton")

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QMessageBox_ClickedButton(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMessageBox) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMessageBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMessageBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMessageBoxCloseEvent
func callbackQMessageBoxCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func QMessageBox_Critical(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	defer qt.Recovering("QMessageBox::critical")

	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Critical(PointerFromQWidget(parent), C.CString(title), C.CString(text), C.int(buttons), C.int(defaultButton)))
}

func (ptr *QMessageBox) DefaultButton() *QPushButton {
	defer qt.Recovering("QMessageBox::defaultButton")

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QMessageBox_DefaultButton(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMessageBox) EscapeButton() *QAbstractButton {
	defer qt.Recovering("QMessageBox::escapeButton")

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QMessageBox_EscapeButton(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMessageBox) Exec() int {
	defer qt.Recovering("QMessageBox::exec")

	if ptr.Pointer() != nil {
		return int(C.QMessageBox_Exec(ptr.Pointer()))
	}
	return 0
}

func QMessageBox_Information(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	defer qt.Recovering("QMessageBox::information")

	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Information(PointerFromQWidget(parent), C.CString(title), C.CString(text), C.int(buttons), C.int(defaultButton)))
}

func (ptr *QMessageBox) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMessageBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMessageBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMessageBoxKeyPressEvent
func callbackQMessageBoxKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMessageBox) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QMessageBox::open")

	if ptr.Pointer() != nil {
		C.QMessageBox_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func QMessageBox_Question(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	defer qt.Recovering("QMessageBox::question")

	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Question(PointerFromQWidget(parent), C.CString(title), C.CString(text), C.int(buttons), C.int(defaultButton)))
}

func (ptr *QMessageBox) RemoveButton(button QAbstractButton_ITF) {
	defer qt.Recovering("QMessageBox::removeButton")

	if ptr.Pointer() != nil {
		C.QMessageBox_RemoveButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QMessageBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMessageBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMessageBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMessageBoxResizeEvent
func callbackQMessageBoxResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) SetCheckBox(cb QCheckBox_ITF) {
	defer qt.Recovering("QMessageBox::setCheckBox")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetCheckBox(ptr.Pointer(), PointerFromQCheckBox(cb))
	}
}

func (ptr *QMessageBox) SetDefaultButton(button QPushButton_ITF) {
	defer qt.Recovering("QMessageBox::setDefaultButton")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetDefaultButton(ptr.Pointer(), PointerFromQPushButton(button))
	}
}

func (ptr *QMessageBox) SetDefaultButton2(button QMessageBox__StandardButton) {
	defer qt.Recovering("QMessageBox::setDefaultButton")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetDefaultButton2(ptr.Pointer(), C.int(button))
	}
}

func (ptr *QMessageBox) SetEscapeButton(button QAbstractButton_ITF) {
	defer qt.Recovering("QMessageBox::setEscapeButton")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetEscapeButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QMessageBox) SetEscapeButton2(button QMessageBox__StandardButton) {
	defer qt.Recovering("QMessageBox::setEscapeButton")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetEscapeButton2(ptr.Pointer(), C.int(button))
	}
}

func (ptr *QMessageBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMessageBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMessageBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMessageBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMessageBoxSetVisible
func callbackQMessageBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMessageBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMessageBox) SetWindowModality(windowModality core.Qt__WindowModality) {
	defer qt.Recovering("QMessageBox::setWindowModality")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetWindowModality(ptr.Pointer(), C.int(windowModality))
	}
}

func (ptr *QMessageBox) SetWindowTitle(title string) {
	defer qt.Recovering("QMessageBox::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QMessageBox_SetWindowTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QMessageBox) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QMessageBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMessageBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMessageBoxShowEvent
func callbackQMessageBoxShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMessageBox) StandardButton(button QAbstractButton_ITF) QMessageBox__StandardButton {
	defer qt.Recovering("QMessageBox::standardButton")

	if ptr.Pointer() != nil {
		return QMessageBox__StandardButton(C.QMessageBox_StandardButton(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func QMessageBox_Warning(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	defer qt.Recovering("QMessageBox::warning")

	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Warning(PointerFromQWidget(parent), C.CString(title), C.CString(text), C.int(buttons), C.int(defaultButton)))
}

func (ptr *QMessageBox) DestroyQMessageBox() {
	defer qt.Recovering("QMessageBox::~QMessageBox")

	if ptr.Pointer() != nil {
		C.QMessageBox_DestroyQMessageBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMessageBox) ConnectAccept(f func()) {
	defer qt.Recovering("connect QMessageBox::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QMessageBox) DisconnectAccept() {
	defer qt.Recovering("disconnect QMessageBox::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQMessageBoxAccept
func callbackQMessageBoxAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QMessageBox::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMessageBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMessageBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMessageBoxContextMenuEvent
func callbackQMessageBoxContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectDone(f func(r int)) {
	defer qt.Recovering("connect QMessageBox::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QMessageBox) DisconnectDone() {
	defer qt.Recovering("disconnect QMessageBox::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQMessageBoxDone
func callbackQMessageBoxDone(ptrName *C.char, r C.int) bool {
	defer qt.Recovering("callback QMessageBox::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(r))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectOpen(f func()) {
	defer qt.Recovering("connect QMessageBox::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QMessageBox) DisconnectOpen() {
	defer qt.Recovering("disconnect QMessageBox::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQMessageBoxOpen
func callbackQMessageBoxOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QMessageBox::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectReject(f func()) {
	defer qt.Recovering("connect QMessageBox::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QMessageBox) DisconnectReject() {
	defer qt.Recovering("disconnect QMessageBox::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQMessageBoxReject
func callbackQMessageBoxReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QMessageBox::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMessageBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMessageBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMessageBoxActionEvent
func callbackQMessageBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMessageBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMessageBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMessageBoxDragEnterEvent
func callbackQMessageBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMessageBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMessageBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMessageBoxDragLeaveEvent
func callbackQMessageBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMessageBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMessageBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMessageBoxDragMoveEvent
func callbackQMessageBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMessageBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMessageBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMessageBoxDropEvent
func callbackQMessageBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMessageBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMessageBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMessageBoxEnterEvent
func callbackQMessageBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMessageBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMessageBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMessageBoxFocusOutEvent
func callbackQMessageBoxFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMessageBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMessageBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMessageBoxHideEvent
func callbackQMessageBoxHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMessageBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMessageBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMessageBoxLeaveEvent
func callbackQMessageBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMessageBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMessageBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMessageBoxMoveEvent
func callbackQMessageBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMessageBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMessageBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMessageBoxPaintEvent
func callbackQMessageBoxPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMessageBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMessageBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMessageBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMessageBoxInitPainter
func callbackQMessageBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMessageBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMessageBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMessageBoxInputMethodEvent
func callbackQMessageBoxInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMessageBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMessageBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMessageBoxKeyReleaseEvent
func callbackQMessageBoxKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMessageBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMessageBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMessageBoxMouseDoubleClickEvent
func callbackQMessageBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMessageBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMessageBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMessageBoxMouseMoveEvent
func callbackQMessageBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMessageBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMessageBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMessageBoxMousePressEvent
func callbackQMessageBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMessageBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMessageBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMessageBoxMouseReleaseEvent
func callbackQMessageBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMessageBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMessageBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMessageBoxTabletEvent
func callbackQMessageBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMessageBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMessageBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMessageBoxWheelEvent
func callbackQMessageBoxWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMessageBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMessageBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMessageBoxTimerEvent
func callbackQMessageBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMessageBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMessageBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMessageBoxChildEvent
func callbackQMessageBoxChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMessageBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMessageBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMessageBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMessageBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMessageBoxCustomEvent
func callbackQMessageBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMessageBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
