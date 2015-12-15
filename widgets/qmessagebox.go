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

	var signal = qt.GetSignal(C.GoString(ptrName), "buttonClicked")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "closeEvent")
	if signal != nil {
		defer signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
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

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
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
