package widgets

//#include "qdialogbuttonbox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDialogButtonBox struct {
	QWidget
}

type QDialogButtonBoxITF interface {
	QWidgetITF
	QDialogButtonBoxPTR() *QDialogButtonBox
}

func PointerFromQDialogButtonBox(ptr QDialogButtonBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialogButtonBoxPTR().Pointer()
	}
	return nil
}

func QDialogButtonBoxFromPointer(ptr unsafe.Pointer) *QDialogButtonBox {
	var n = new(QDialogButtonBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDialogButtonBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDialogButtonBox) QDialogButtonBoxPTR() *QDialogButtonBox {
	return ptr
}

//QDialogButtonBox::ButtonLayout
type QDialogButtonBox__ButtonLayout int

var (
	QDialogButtonBox__WinLayout   = QDialogButtonBox__ButtonLayout(0)
	QDialogButtonBox__MacLayout   = QDialogButtonBox__ButtonLayout(1)
	QDialogButtonBox__KdeLayout   = QDialogButtonBox__ButtonLayout(2)
	QDialogButtonBox__GnomeLayout = QDialogButtonBox__ButtonLayout(3)
)

//QDialogButtonBox::ButtonRole
type QDialogButtonBox__ButtonRole int

var (
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
type QDialogButtonBox__StandardButton int

var (
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
	if ptr.Pointer() != nil {
		return C.QDialogButtonBox_CenterButtons(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDialogButtonBox) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDialogButtonBox_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDialogButtonBox) SetCenterButtons(center bool) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetCenterButtons(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(center)))
	}
}

func (ptr *QDialogButtonBox) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}

func (ptr *QDialogButtonBox) SetStandardButtons(buttons QDialogButtonBox__StandardButton) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetStandardButtons(C.QtObjectPtr(ptr.Pointer()), C.int(buttons))
	}
}

func (ptr *QDialogButtonBox) StandardButtons() QDialogButtonBox__StandardButton {
	if ptr.Pointer() != nil {
		return QDialogButtonBox__StandardButton(C.QDialogButtonBox_StandardButtons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQDialogButtonBox(parent QWidgetITF) *QDialogButtonBox {
	return QDialogButtonBoxFromPointer(unsafe.Pointer(C.QDialogButtonBox_NewQDialogButtonBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQDialogButtonBox2(orientation core.Qt__Orientation, parent QWidgetITF) *QDialogButtonBox {
	return QDialogButtonBoxFromPointer(unsafe.Pointer(C.QDialogButtonBox_NewQDialogButtonBox2(C.int(orientation), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQDialogButtonBox3(buttons QDialogButtonBox__StandardButton, parent QWidgetITF) *QDialogButtonBox {
	return QDialogButtonBoxFromPointer(unsafe.Pointer(C.QDialogButtonBox_NewQDialogButtonBox3(C.int(buttons), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQDialogButtonBox4(buttons QDialogButtonBox__StandardButton, orientation core.Qt__Orientation, parent QWidgetITF) *QDialogButtonBox {
	return QDialogButtonBoxFromPointer(unsafe.Pointer(C.QDialogButtonBox_NewQDialogButtonBox4(C.int(buttons), C.int(orientation), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QDialogButtonBox) ConnectAccepted(f func()) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectAccepted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "accepted", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectAccepted() {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectAccepted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "accepted")
	}
}

//export callbackQDialogButtonBoxAccepted
func callbackQDialogButtonBoxAccepted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "accepted").(func())()
}

func (ptr *QDialogButtonBox) AddButton3(button QDialogButtonBox__StandardButton) *QPushButton {
	if ptr.Pointer() != nil {
		return QPushButtonFromPointer(unsafe.Pointer(C.QDialogButtonBox_AddButton3(C.QtObjectPtr(ptr.Pointer()), C.int(button))))
	}
	return nil
}

func (ptr *QDialogButtonBox) AddButton2(text string, role QDialogButtonBox__ButtonRole) *QPushButton {
	if ptr.Pointer() != nil {
		return QPushButtonFromPointer(unsafe.Pointer(C.QDialogButtonBox_AddButton2(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.int(role))))
	}
	return nil
}

func (ptr *QDialogButtonBox) AddButton(button QAbstractButtonITF, role QDialogButtonBox__ButtonRole) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_AddButton(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button)), C.int(role))
	}
}

func (ptr *QDialogButtonBox) Button(which QDialogButtonBox__StandardButton) *QPushButton {
	if ptr.Pointer() != nil {
		return QPushButtonFromPointer(unsafe.Pointer(C.QDialogButtonBox_Button(C.QtObjectPtr(ptr.Pointer()), C.int(which))))
	}
	return nil
}

func (ptr *QDialogButtonBox) ButtonRole(button QAbstractButtonITF) QDialogButtonBox__ButtonRole {
	if ptr.Pointer() != nil {
		return QDialogButtonBox__ButtonRole(C.QDialogButtonBox_ButtonRole(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button))))
	}
	return 0
}

func (ptr *QDialogButtonBox) Clear() {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDialogButtonBox) ConnectClicked(f func(button QAbstractButtonITF)) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQDialogButtonBoxClicked
func callbackQDialogButtonBoxClicked(ptrName *C.char, button unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(*QAbstractButton))(QAbstractButtonFromPointer(button))
}

func (ptr *QDialogButtonBox) ConnectHelpRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectHelpRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectHelpRequested() {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectHelpRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQDialogButtonBoxHelpRequested
func callbackQDialogButtonBoxHelpRequested(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "helpRequested").(func())()
}

func (ptr *QDialogButtonBox) ConnectRejected(f func()) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectRejected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "rejected", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectRejected() {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectRejected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "rejected")
	}
}

//export callbackQDialogButtonBoxRejected
func callbackQDialogButtonBoxRejected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "rejected").(func())()
}

func (ptr *QDialogButtonBox) RemoveButton(button QAbstractButtonITF) {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_RemoveButton(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button)))
	}
}

func (ptr *QDialogButtonBox) StandardButton(button QAbstractButtonITF) QDialogButtonBox__StandardButton {
	if ptr.Pointer() != nil {
		return QDialogButtonBox__StandardButton(C.QDialogButtonBox_StandardButton(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractButton(button))))
	}
	return 0
}

func (ptr *QDialogButtonBox) DestroyQDialogButtonBox() {
	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DestroyQDialogButtonBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
