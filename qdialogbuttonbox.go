package qt

//#include "qdialogbuttonbox.h"
import "C"

type qdialogbuttonbox struct {
	qwidget
}

type QDialogButtonBox interface {
	QWidget
	AddButton1(button QAbstractButton, role ButtonRole)
	AddButton2(text string, role ButtonRole) QPushButton
	AddButton3(button StandardButton) QPushButton
	Button(which StandardButton) QPushButton
	ButtonRole(button QAbstractButton) ButtonRole
	CenterButtons() bool
	Clear()
	Orientation() Orientation
	RemoveButton(button QAbstractButton)
	SetCenterButtons(center bool)
	SetOrientation(orientation Orientation)
	SetStandardButtons(buttons StandardButton)
	StandardButton(button QAbstractButton) StandardButton
	StandardButtons() StandardButton
	ConnectSignalAccepted(f func())
	DisconnectSignalAccepted()
	SignalAccepted() func()
	ConnectSignalHelpRequested(f func())
	DisconnectSignalHelpRequested()
	SignalHelpRequested() func()
	ConnectSignalRejected(f func())
	DisconnectSignalRejected()
	SignalRejected() func()
}

func (p *qdialogbuttonbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qdialogbuttonbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//ButtonLayout
type ButtonLayout int

var (
	WINLAYOUT   = ButtonLayout(C.QDialogButtonBox_WinLayout())
	MACLAYOUT   = ButtonLayout(C.QDialogButtonBox_MacLayout())
	KDELAYOUT   = ButtonLayout(C.QDialogButtonBox_KdeLayout())
	GNOMELAYOUT = ButtonLayout(C.QDialogButtonBox_GnomeLayout())
)

//ButtonRole
type ButtonRole int

var (
	INVALIDROLE     = ButtonRole(C.QDialogButtonBox_InvalidRole())
	ACCEPTROLE      = ButtonRole(C.QDialogButtonBox_AcceptRole())
	REJECTROLE      = ButtonRole(C.QDialogButtonBox_RejectRole())
	DESTRUCTIVEROLE = ButtonRole(C.QDialogButtonBox_DestructiveRole())
	ACTIONROLE      = ButtonRole(C.QDialogButtonBox_ActionRole())
	HELPROLE        = ButtonRole(C.QDialogButtonBox_HelpRole())
	YESROLE         = ButtonRole(C.QDialogButtonBox_YesRole())
	NOROLE          = ButtonRole(C.QDialogButtonBox_NoRole())
	APPLYROLE       = ButtonRole(C.QDialogButtonBox_ApplyRole())
	RESETROLE       = ButtonRole(C.QDialogButtonBox_ResetRole())
)

//StandardButton
type StandardButton int

var (
	OK              = StandardButton(C.QDialogButtonBox_Ok())
	OPEN            = StandardButton(C.QDialogButtonBox_Open())
	SAVE            = StandardButton(C.QDialogButtonBox_Save())
	CANCEL          = StandardButton(C.QDialogButtonBox_Cancel())
	CLOSE           = StandardButton(C.QDialogButtonBox_Close())
	DISCARD         = StandardButton(C.QDialogButtonBox_Discard())
	APPLY           = StandardButton(C.QDialogButtonBox_Apply())
	RESET           = StandardButton(C.QDialogButtonBox_Reset())
	RESTOREDEFAULTS = StandardButton(C.QDialogButtonBox_RestoreDefaults())
	HELP            = StandardButton(C.QDialogButtonBox_Help())
	SAVEALL         = StandardButton(C.QDialogButtonBox_SaveAll())
	YES             = StandardButton(C.QDialogButtonBox_Yes())
	YESTOALL        = StandardButton(C.QDialogButtonBox_YesToAll())
	NO              = StandardButton(C.QDialogButtonBox_No())
	NOTOALL         = StandardButton(C.QDialogButtonBox_NoToAll())
	ABORT           = StandardButton(C.QDialogButtonBox_Abort())
	RETRY           = StandardButton(C.QDialogButtonBox_Retry())
	IGNORE          = StandardButton(C.QDialogButtonBox_Ignore())
)

func NewQDialogButtonBox1(parent QWidget) QDialogButtonBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdialogbuttonbox = new(qdialogbuttonbox)
	qdialogbuttonbox.SetPointer(C.QDialogButtonBox_New_QWidget(parentPtr))
	qdialogbuttonbox.SetObjectName("QDialogButtonBox_" + randomIdentifier())
	return qdialogbuttonbox
}

func NewQDialogButtonBox2(orientation Orientation, parent QWidget) QDialogButtonBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdialogbuttonbox = new(qdialogbuttonbox)
	qdialogbuttonbox.SetPointer(C.QDialogButtonBox_New_Orientation_QWidget(C.int(orientation), parentPtr))
	qdialogbuttonbox.SetObjectName("QDialogButtonBox_" + randomIdentifier())
	return qdialogbuttonbox
}

func NewQDialogButtonBox3(buttons StandardButton, parent QWidget) QDialogButtonBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdialogbuttonbox = new(qdialogbuttonbox)
	qdialogbuttonbox.SetPointer(C.QDialogButtonBox_New_StandardButton_QWidget(C.int(buttons), parentPtr))
	qdialogbuttonbox.SetObjectName("QDialogButtonBox_" + randomIdentifier())
	return qdialogbuttonbox
}

func NewQDialogButtonBox4(buttons StandardButton, orientation Orientation, parent QWidget) QDialogButtonBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qdialogbuttonbox = new(qdialogbuttonbox)
	qdialogbuttonbox.SetPointer(C.QDialogButtonBox_New_StandardButton_Orientation_QWidget(C.int(buttons), C.int(orientation), parentPtr))
	qdialogbuttonbox.SetObjectName("QDialogButtonBox_" + randomIdentifier())
	return qdialogbuttonbox
}

func (p *qdialogbuttonbox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QDialogButtonBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qdialogbuttonbox) AddButton1(button QAbstractButton, role ButtonRole) {
	if p.Pointer() != nil {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QDialogButtonBox_AddButton_QAbstractButton_ButtonRole(p.Pointer(), buttonPtr, C.int(role))
	}
}

func (p *qdialogbuttonbox) AddButton2(text string, role ButtonRole) QPushButton {
	if p.Pointer() == nil {
		return nil
	} else {
		var qpushbutton = new(qpushbutton)
		qpushbutton.SetPointer(C.QDialogButtonBox_AddButton_String_ButtonRole(p.Pointer(), C.CString(text), C.int(role)))
		if qpushbutton.ObjectName() == "" {
			qpushbutton.SetObjectName("QPushButton_" + randomIdentifier())
		}
		return qpushbutton
	}
}

func (p *qdialogbuttonbox) AddButton3(button StandardButton) QPushButton {
	if p.Pointer() == nil {
		return nil
	} else {
		var qpushbutton = new(qpushbutton)
		qpushbutton.SetPointer(C.QDialogButtonBox_AddButton_StandardButton(p.Pointer(), C.int(button)))
		if qpushbutton.ObjectName() == "" {
			qpushbutton.SetObjectName("QPushButton_" + randomIdentifier())
		}
		return qpushbutton
	}
}

func (p *qdialogbuttonbox) Button(which StandardButton) QPushButton {
	if p.Pointer() == nil {
		return nil
	} else {
		var qpushbutton = new(qpushbutton)
		qpushbutton.SetPointer(C.QDialogButtonBox_Button_StandardButton(p.Pointer(), C.int(which)))
		if qpushbutton.ObjectName() == "" {
			qpushbutton.SetObjectName("QPushButton_" + randomIdentifier())
		}
		return qpushbutton
	}
}

func (p *qdialogbuttonbox) ButtonRole(button QAbstractButton) ButtonRole {
	if p.Pointer() == nil {
		return 0
	} else {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		return ButtonRole(C.QDialogButtonBox_ButtonRole_QAbstractButton(p.Pointer(), buttonPtr))
	}
}

func (p *qdialogbuttonbox) CenterButtons() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QDialogButtonBox_CenterButtons(p.Pointer()) != 0
}

func (p *qdialogbuttonbox) Clear() {
	if p.Pointer() != nil {
		C.QDialogButtonBox_Clear(p.Pointer())
	}
}

func (p *qdialogbuttonbox) Orientation() Orientation {
	if p.Pointer() == nil {
		return 0
	}
	return Orientation(C.QDialogButtonBox_Orientation(p.Pointer()))
}

func (p *qdialogbuttonbox) RemoveButton(button QAbstractButton) {
	if p.Pointer() != nil {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		C.QDialogButtonBox_RemoveButton_QAbstractButton(p.Pointer(), buttonPtr)
	}
}

func (p *qdialogbuttonbox) SetCenterButtons(center bool) {
	if p.Pointer() != nil {
		C.QDialogButtonBox_SetCenterButtons_Bool(p.Pointer(), goBoolToCInt(center))
	}
}

func (p *qdialogbuttonbox) SetOrientation(orientation Orientation) {
	if p.Pointer() != nil {
		C.QDialogButtonBox_SetOrientation_Orientation(p.Pointer(), C.int(orientation))
	}
}

func (p *qdialogbuttonbox) SetStandardButtons(buttons StandardButton) {
	if p.Pointer() != nil {
		C.QDialogButtonBox_SetStandardButtons_StandardButton(p.Pointer(), C.int(buttons))
	}
}

func (p *qdialogbuttonbox) StandardButton(button QAbstractButton) StandardButton {
	if p.Pointer() == nil {
		return 0
	} else {
		var buttonPtr C.QtObjectPtr
		if button != nil {
			buttonPtr = button.Pointer()
		}
		return StandardButton(C.QDialogButtonBox_StandardButton_QAbstractButton(p.Pointer(), buttonPtr))
	}
}

func (p *qdialogbuttonbox) StandardButtons() StandardButton {
	if p.Pointer() == nil {
		return 0
	}
	return StandardButton(C.QDialogButtonBox_StandardButtons(p.Pointer()))
}

func (p *qdialogbuttonbox) ConnectSignalAccepted(f func()) {
	C.QDialogButtonBox_ConnectSignalAccepted(p.Pointer())
	connectSignal(p.ObjectName(), "accepted", f)
}

func (p *qdialogbuttonbox) DisconnectSignalAccepted() {
	C.QDialogButtonBox_DisconnectSignalAccepted(p.Pointer())
	disconnectSignal(p.ObjectName(), "accepted")
}

func (p *qdialogbuttonbox) SignalAccepted() func() {
	return getSignal(p.ObjectName(), "accepted")
}

func (p *qdialogbuttonbox) ConnectSignalHelpRequested(f func()) {
	C.QDialogButtonBox_ConnectSignalHelpRequested(p.Pointer())
	connectSignal(p.ObjectName(), "helpRequested", f)
}

func (p *qdialogbuttonbox) DisconnectSignalHelpRequested() {
	C.QDialogButtonBox_DisconnectSignalHelpRequested(p.Pointer())
	disconnectSignal(p.ObjectName(), "helpRequested")
}

func (p *qdialogbuttonbox) SignalHelpRequested() func() {
	return getSignal(p.ObjectName(), "helpRequested")
}

func (p *qdialogbuttonbox) ConnectSignalRejected(f func()) {
	C.QDialogButtonBox_ConnectSignalRejected(p.Pointer())
	connectSignal(p.ObjectName(), "rejected", f)
}

func (p *qdialogbuttonbox) DisconnectSignalRejected() {
	C.QDialogButtonBox_DisconnectSignalRejected(p.Pointer())
	disconnectSignal(p.ObjectName(), "rejected")
}

func (p *qdialogbuttonbox) SignalRejected() func() {
	return getSignal(p.ObjectName(), "rejected")
}

//export callbackQDialogButtonBox
func callbackQDialogButtonBox(ptr C.QtObjectPtr, msg *C.char) {
	var qdialogbuttonbox = new(qdialogbuttonbox)
	qdialogbuttonbox.SetPointer(ptr)
	getSignal(qdialogbuttonbox.ObjectName(), C.GoString(msg))()
}
