package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDialogButtonBox struct {
	QWidget
}

type QDialogButtonBox_ITF interface {
	QWidget_ITF
	QDialogButtonBox_PTR() *QDialogButtonBox
}

func PointerFromQDialogButtonBox(ptr QDialogButtonBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialogButtonBox_PTR().Pointer()
	}
	return nil
}

func NewQDialogButtonBoxFromPointer(ptr unsafe.Pointer) *QDialogButtonBox {
	var n = new(QDialogButtonBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDialogButtonBox_") {
		n.SetObjectName("QDialogButtonBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QDialogButtonBox) QDialogButtonBox_PTR() *QDialogButtonBox {
	return ptr
}

//QDialogButtonBox::ButtonLayout
type QDialogButtonBox__ButtonLayout int64

const (
	QDialogButtonBox__WinLayout   = QDialogButtonBox__ButtonLayout(0)
	QDialogButtonBox__MacLayout   = QDialogButtonBox__ButtonLayout(1)
	QDialogButtonBox__KdeLayout   = QDialogButtonBox__ButtonLayout(2)
	QDialogButtonBox__GnomeLayout = QDialogButtonBox__ButtonLayout(3)
)

//QDialogButtonBox::ButtonRole
type QDialogButtonBox__ButtonRole int64

const (
	QDialogButtonBox__InvalidRole     = QDialogButtonBox__ButtonRole(-1)
	QDialogButtonBox__AcceptRole      = QDialogButtonBox__ButtonRole(0)
	QDialogButtonBox__RejectRole      = QDialogButtonBox__ButtonRole(1)
	QDialogButtonBox__DestructiveRole = QDialogButtonBox__ButtonRole(2)
	QDialogButtonBox__ActionRole      = QDialogButtonBox__ButtonRole(3)
	QDialogButtonBox__HelpRole        = QDialogButtonBox__ButtonRole(4)
	QDialogButtonBox__YesRole         = QDialogButtonBox__ButtonRole(5)
	QDialogButtonBox__NoRole          = QDialogButtonBox__ButtonRole(6)
	QDialogButtonBox__ResetRole       = QDialogButtonBox__ButtonRole(7)
	QDialogButtonBox__ApplyRole       = QDialogButtonBox__ButtonRole(8)
	QDialogButtonBox__NRoles          = QDialogButtonBox__ButtonRole(9)
)

//QDialogButtonBox::StandardButton
type QDialogButtonBox__StandardButton int64

const (
	QDialogButtonBox__NoButton        = QDialogButtonBox__StandardButton(0x00000000)
	QDialogButtonBox__Ok              = QDialogButtonBox__StandardButton(0x00000400)
	QDialogButtonBox__Save            = QDialogButtonBox__StandardButton(0x00000800)
	QDialogButtonBox__SaveAll         = QDialogButtonBox__StandardButton(0x00001000)
	QDialogButtonBox__Open            = QDialogButtonBox__StandardButton(0x00002000)
	QDialogButtonBox__Yes             = QDialogButtonBox__StandardButton(0x00004000)
	QDialogButtonBox__YesToAll        = QDialogButtonBox__StandardButton(0x00008000)
	QDialogButtonBox__No              = QDialogButtonBox__StandardButton(0x00010000)
	QDialogButtonBox__NoToAll         = QDialogButtonBox__StandardButton(0x00020000)
	QDialogButtonBox__Abort           = QDialogButtonBox__StandardButton(0x00040000)
	QDialogButtonBox__Retry           = QDialogButtonBox__StandardButton(0x00080000)
	QDialogButtonBox__Ignore          = QDialogButtonBox__StandardButton(0x00100000)
	QDialogButtonBox__Close           = QDialogButtonBox__StandardButton(0x00200000)
	QDialogButtonBox__Cancel          = QDialogButtonBox__StandardButton(0x00400000)
	QDialogButtonBox__Discard         = QDialogButtonBox__StandardButton(0x00800000)
	QDialogButtonBox__Help            = QDialogButtonBox__StandardButton(0x01000000)
	QDialogButtonBox__Apply           = QDialogButtonBox__StandardButton(0x02000000)
	QDialogButtonBox__Reset           = QDialogButtonBox__StandardButton(0x04000000)
	QDialogButtonBox__RestoreDefaults = QDialogButtonBox__StandardButton(0x08000000)
	QDialogButtonBox__FirstButton     = QDialogButtonBox__StandardButton(QDialogButtonBox__Ok)
	QDialogButtonBox__LastButton      = QDialogButtonBox__StandardButton(QDialogButtonBox__RestoreDefaults)
)

func (ptr *QDialogButtonBox) CenterButtons() bool {
	defer qt.Recovering("QDialogButtonBox::centerButtons")

	if ptr.Pointer() != nil {
		return C.QDialogButtonBox_CenterButtons(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDialogButtonBox) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QDialogButtonBox::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDialogButtonBox_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialogButtonBox) SetCenterButtons(center bool) {
	defer qt.Recovering("QDialogButtonBox::setCenterButtons")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetCenterButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(center)))
	}
}

func (ptr *QDialogButtonBox) SetOrientation(orientation core.Qt__Orientation) {
	defer qt.Recovering("QDialogButtonBox::setOrientation")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QDialogButtonBox) SetStandardButtons(buttons QDialogButtonBox__StandardButton) {
	defer qt.Recovering("QDialogButtonBox::setStandardButtons")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetStandardButtons(ptr.Pointer(), C.int(buttons))
	}
}

func (ptr *QDialogButtonBox) StandardButtons() QDialogButtonBox__StandardButton {
	defer qt.Recovering("QDialogButtonBox::standardButtons")

	if ptr.Pointer() != nil {
		return QDialogButtonBox__StandardButton(C.QDialogButtonBox_StandardButtons(ptr.Pointer()))
	}
	return 0
}

func NewQDialogButtonBox(parent QWidget_ITF) *QDialogButtonBox {
	defer qt.Recovering("QDialogButtonBox::QDialogButtonBox")

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox(PointerFromQWidget(parent)))
}

func NewQDialogButtonBox2(orientation core.Qt__Orientation, parent QWidget_ITF) *QDialogButtonBox {
	defer qt.Recovering("QDialogButtonBox::QDialogButtonBox")

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox2(C.int(orientation), PointerFromQWidget(parent)))
}

func NewQDialogButtonBox3(buttons QDialogButtonBox__StandardButton, parent QWidget_ITF) *QDialogButtonBox {
	defer qt.Recovering("QDialogButtonBox::QDialogButtonBox")

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox3(C.int(buttons), PointerFromQWidget(parent)))
}

func NewQDialogButtonBox4(buttons QDialogButtonBox__StandardButton, orientation core.Qt__Orientation, parent QWidget_ITF) *QDialogButtonBox {
	defer qt.Recovering("QDialogButtonBox::QDialogButtonBox")

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox4(C.int(buttons), C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QDialogButtonBox) ConnectAccepted(f func()) {
	defer qt.Recovering("connect QDialogButtonBox::accepted")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectAccepted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accepted", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectAccepted() {
	defer qt.Recovering("disconnect QDialogButtonBox::accepted")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accepted")
	}
}

//export callbackQDialogButtonBoxAccepted
func callbackQDialogButtonBoxAccepted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDialogButtonBox::accepted")

	if signal := qt.GetSignal(C.GoString(ptrName), "accepted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialogButtonBox) Accepted() {
	defer qt.Recovering("QDialogButtonBox::accepted")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_Accepted(ptr.Pointer())
	}
}

func (ptr *QDialogButtonBox) AddButton3(button QDialogButtonBox__StandardButton) *QPushButton {
	defer qt.Recovering("QDialogButtonBox::addButton")

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QDialogButtonBox_AddButton3(ptr.Pointer(), C.int(button)))
	}
	return nil
}

func (ptr *QDialogButtonBox) AddButton2(text string, role QDialogButtonBox__ButtonRole) *QPushButton {
	defer qt.Recovering("QDialogButtonBox::addButton")

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QDialogButtonBox_AddButton2(ptr.Pointer(), C.CString(text), C.int(role)))
	}
	return nil
}

func (ptr *QDialogButtonBox) AddButton(button QAbstractButton_ITF, role QDialogButtonBox__ButtonRole) {
	defer qt.Recovering("QDialogButtonBox::addButton")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_AddButton(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(role))
	}
}

func (ptr *QDialogButtonBox) Button(which QDialogButtonBox__StandardButton) *QPushButton {
	defer qt.Recovering("QDialogButtonBox::button")

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QDialogButtonBox_Button(ptr.Pointer(), C.int(which)))
	}
	return nil
}

func (ptr *QDialogButtonBox) ButtonRole(button QAbstractButton_ITF) QDialogButtonBox__ButtonRole {
	defer qt.Recovering("QDialogButtonBox::buttonRole")

	if ptr.Pointer() != nil {
		return QDialogButtonBox__ButtonRole(C.QDialogButtonBox_ButtonRole(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QDialogButtonBox) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDialogButtonBoxChangeEvent
func callbackQDialogButtonBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) Clear() {
	defer qt.Recovering("QDialogButtonBox::clear")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_Clear(ptr.Pointer())
	}
}

func (ptr *QDialogButtonBox) ConnectClicked(f func(button *QAbstractButton)) {
	defer qt.Recovering("connect QDialogButtonBox::clicked")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectClicked() {
	defer qt.Recovering("disconnect QDialogButtonBox::clicked")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQDialogButtonBoxClicked
func callbackQDialogButtonBoxClicked(ptr unsafe.Pointer, ptrName *C.char, button unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
	}

}

func (ptr *QDialogButtonBox) Clicked(button QAbstractButton_ITF) {
	defer qt.Recovering("QDialogButtonBox::clicked")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_Clicked(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QDialogButtonBox) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QDialogButtonBox::event")

	if ptr.Pointer() != nil {
		return C.QDialogButtonBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDialogButtonBox) ConnectHelpRequested(f func()) {
	defer qt.Recovering("connect QDialogButtonBox::helpRequested")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectHelpRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectHelpRequested() {
	defer qt.Recovering("disconnect QDialogButtonBox::helpRequested")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectHelpRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQDialogButtonBoxHelpRequested
func callbackQDialogButtonBoxHelpRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDialogButtonBox::helpRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "helpRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialogButtonBox) HelpRequested() {
	defer qt.Recovering("QDialogButtonBox::helpRequested")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_HelpRequested(ptr.Pointer())
	}
}

func (ptr *QDialogButtonBox) ConnectRejected(f func()) {
	defer qt.Recovering("connect QDialogButtonBox::rejected")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectRejected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rejected", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectRejected() {
	defer qt.Recovering("disconnect QDialogButtonBox::rejected")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectRejected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rejected")
	}
}

//export callbackQDialogButtonBoxRejected
func callbackQDialogButtonBoxRejected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDialogButtonBox::rejected")

	if signal := qt.GetSignal(C.GoString(ptrName), "rejected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDialogButtonBox) Rejected() {
	defer qt.Recovering("QDialogButtonBox::rejected")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_Rejected(ptr.Pointer())
	}
}

func (ptr *QDialogButtonBox) RemoveButton(button QAbstractButton_ITF) {
	defer qt.Recovering("QDialogButtonBox::removeButton")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_RemoveButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QDialogButtonBox) StandardButton(button QAbstractButton_ITF) QDialogButtonBox__StandardButton {
	defer qt.Recovering("QDialogButtonBox::standardButton")

	if ptr.Pointer() != nil {
		return QDialogButtonBox__StandardButton(C.QDialogButtonBox_StandardButton(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QDialogButtonBox) DestroyQDialogButtonBox() {
	defer qt.Recovering("QDialogButtonBox::~QDialogButtonBox")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DestroyQDialogButtonBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDialogButtonBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDialogButtonBoxActionEvent
func callbackQDialogButtonBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDialogButtonBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDialogButtonBoxDragEnterEvent
func callbackQDialogButtonBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDialogButtonBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDialogButtonBoxDragLeaveEvent
func callbackQDialogButtonBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDialogButtonBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDialogButtonBoxDragMoveEvent
func callbackQDialogButtonBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDialogButtonBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDialogButtonBoxDropEvent
func callbackQDialogButtonBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDialogButtonBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDialogButtonBoxEnterEvent
func callbackQDialogButtonBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDialogButtonBoxFocusInEvent
func callbackQDialogButtonBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialogButtonBox) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDialogButtonBoxFocusOutEvent
func callbackQDialogButtonBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialogButtonBox) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDialogButtonBoxHideEvent
func callbackQDialogButtonBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDialogButtonBox) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDialogButtonBoxLeaveEvent
func callbackQDialogButtonBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDialogButtonBoxMoveEvent
func callbackQDialogButtonBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDialogButtonBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDialogButtonBoxPaintEvent
func callbackQDialogButtonBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDialogButtonBox) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDialogButtonBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDialogButtonBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDialogButtonBoxSetVisible
func callbackQDialogButtonBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDialogButtonBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDialogButtonBox) SetVisible(visible bool) {
	defer qt.Recovering("QDialogButtonBox::setVisible")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDialogButtonBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDialogButtonBox::setVisible")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDialogButtonBox) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDialogButtonBoxShowEvent
func callbackQDialogButtonBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::showEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDialogButtonBox) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::showEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDialogButtonBoxCloseEvent
func callbackQDialogButtonBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDialogButtonBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDialogButtonBoxContextMenuEvent
func callbackQDialogButtonBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDialogButtonBox) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDialogButtonBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDialogButtonBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDialogButtonBoxInitPainter
func callbackQDialogButtonBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QDialogButtonBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDialogButtonBox::initPainter")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDialogButtonBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDialogButtonBox::initPainter")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDialogButtonBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDialogButtonBoxInputMethodEvent
func callbackQDialogButtonBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDialogButtonBox) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDialogButtonBoxKeyPressEvent
func callbackQDialogButtonBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDialogButtonBox) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDialogButtonBoxKeyReleaseEvent
func callbackQDialogButtonBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDialogButtonBox) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDialogButtonBoxMouseDoubleClickEvent
func callbackQDialogButtonBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDialogButtonBoxMouseMoveEvent
func callbackQDialogButtonBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDialogButtonBoxMousePressEvent
func callbackQDialogButtonBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDialogButtonBoxMouseReleaseEvent
func callbackQDialogButtonBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDialogButtonBoxResizeEvent
func callbackQDialogButtonBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDialogButtonBox) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDialogButtonBoxTabletEvent
func callbackQDialogButtonBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDialogButtonBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDialogButtonBoxWheelEvent
func callbackQDialogButtonBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDialogButtonBox) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDialogButtonBoxTimerEvent
func callbackQDialogButtonBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDialogButtonBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDialogButtonBoxChildEvent
func callbackQDialogButtonBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::childEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDialogButtonBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::childEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDialogButtonBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDialogButtonBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDialogButtonBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDialogButtonBoxCustomEvent
func callbackQDialogButtonBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDialogButtonBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialogButtonBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDialogButtonBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::customEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDialogButtonBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDialogButtonBox::customEvent")

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
