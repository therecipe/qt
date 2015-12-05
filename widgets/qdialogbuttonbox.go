package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QDialogButtonBox_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::centerButtons")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDialogButtonBox_CenterButtons(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDialogButtonBox) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QDialogButtonBox_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialogButtonBox) SetCenterButtons(center bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::setCenterButtons")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetCenterButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(center)))
	}
}

func (ptr *QDialogButtonBox) SetOrientation(orientation core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QDialogButtonBox) SetStandardButtons(buttons QDialogButtonBox__StandardButton) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::setStandardButtons")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_SetStandardButtons(ptr.Pointer(), C.int(buttons))
	}
}

func (ptr *QDialogButtonBox) StandardButtons() QDialogButtonBox__StandardButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::standardButtons")
		}
	}()

	if ptr.Pointer() != nil {
		return QDialogButtonBox__StandardButton(C.QDialogButtonBox_StandardButtons(ptr.Pointer()))
	}
	return 0
}

func NewQDialogButtonBox(parent QWidget_ITF) *QDialogButtonBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::QDialogButtonBox")
		}
	}()

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox(PointerFromQWidget(parent)))
}

func NewQDialogButtonBox2(orientation core.Qt__Orientation, parent QWidget_ITF) *QDialogButtonBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::QDialogButtonBox")
		}
	}()

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox2(C.int(orientation), PointerFromQWidget(parent)))
}

func NewQDialogButtonBox3(buttons QDialogButtonBox__StandardButton, parent QWidget_ITF) *QDialogButtonBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::QDialogButtonBox")
		}
	}()

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox3(C.int(buttons), PointerFromQWidget(parent)))
}

func NewQDialogButtonBox4(buttons QDialogButtonBox__StandardButton, orientation core.Qt__Orientation, parent QWidget_ITF) *QDialogButtonBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::QDialogButtonBox")
		}
	}()

	return NewQDialogButtonBoxFromPointer(C.QDialogButtonBox_NewQDialogButtonBox4(C.int(buttons), C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QDialogButtonBox) ConnectAccepted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::accepted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectAccepted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accepted", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectAccepted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::accepted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accepted")
	}
}

//export callbackQDialogButtonBoxAccepted
func callbackQDialogButtonBoxAccepted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::accepted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "accepted").(func())()
}

func (ptr *QDialogButtonBox) AddButton3(button QDialogButtonBox__StandardButton) *QPushButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::addButton")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QDialogButtonBox_AddButton3(ptr.Pointer(), C.int(button)))
	}
	return nil
}

func (ptr *QDialogButtonBox) AddButton2(text string, role QDialogButtonBox__ButtonRole) *QPushButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::addButton")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QDialogButtonBox_AddButton2(ptr.Pointer(), C.CString(text), C.int(role)))
	}
	return nil
}

func (ptr *QDialogButtonBox) AddButton(button QAbstractButton_ITF, role QDialogButtonBox__ButtonRole) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::addButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_AddButton(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(role))
	}
}

func (ptr *QDialogButtonBox) Button(which QDialogButtonBox__StandardButton) *QPushButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::button")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPushButtonFromPointer(C.QDialogButtonBox_Button(ptr.Pointer(), C.int(which)))
	}
	return nil
}

func (ptr *QDialogButtonBox) ButtonRole(button QAbstractButton_ITF) QDialogButtonBox__ButtonRole {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::buttonRole")
		}
	}()

	if ptr.Pointer() != nil {
		return QDialogButtonBox__ButtonRole(C.QDialogButtonBox_ButtonRole(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QDialogButtonBox) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_Clear(ptr.Pointer())
	}
}

func (ptr *QDialogButtonBox) ConnectClicked(f func(button *QAbstractButton)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::clicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::clicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQDialogButtonBoxClicked
func callbackQDialogButtonBoxClicked(ptrName *C.char, button unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::clicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "clicked").(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
}

func (ptr *QDialogButtonBox) ConnectHelpRequested(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::helpRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectHelpRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectHelpRequested() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::helpRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectHelpRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQDialogButtonBoxHelpRequested
func callbackQDialogButtonBoxHelpRequested(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::helpRequested")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "helpRequested").(func())()
}

func (ptr *QDialogButtonBox) ConnectRejected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::rejected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_ConnectRejected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rejected", f)
	}
}

func (ptr *QDialogButtonBox) DisconnectRejected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::rejected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DisconnectRejected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rejected")
	}
}

//export callbackQDialogButtonBoxRejected
func callbackQDialogButtonBoxRejected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::rejected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "rejected").(func())()
}

func (ptr *QDialogButtonBox) RemoveButton(button QAbstractButton_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::removeButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_RemoveButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QDialogButtonBox) StandardButton(button QAbstractButton_ITF) QDialogButtonBox__StandardButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::standardButton")
		}
	}()

	if ptr.Pointer() != nil {
		return QDialogButtonBox__StandardButton(C.QDialogButtonBox_StandardButton(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QDialogButtonBox) DestroyQDialogButtonBox() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialogButtonBox::~QDialogButtonBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialogButtonBox_DestroyQDialogButtonBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
