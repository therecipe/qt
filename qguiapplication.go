package qt

//#include "qguiapplication.h"
import "C"

type qguiapplication struct {
	qcoreapplication
}

type QGuiApplication interface {
	QCoreApplication
	IsSavingSession() bool
	IsSessionRestored() bool
	SessionId() string
	SessionKey() string
	ConnectSignalApplicationStateChanged(f func())
	DisconnectSignalApplicationStateChanged()
	SignalApplicationStateChanged() func()
	ConnectSignalFocusObjectChanged(f func())
	DisconnectSignalFocusObjectChanged()
	SignalFocusObjectChanged() func()
	ConnectSignalFontDatabaseChanged(f func())
	DisconnectSignalFontDatabaseChanged()
	SignalFontDatabaseChanged() func()
	ConnectSignalLastWindowClosed(f func())
	DisconnectSignalLastWindowClosed()
	SignalLastWindowClosed() func()
}

func (p *qguiapplication) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qguiapplication) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQGuiApplication_Int_String(argc int, argv string) QGuiApplication {
	var qguiapplication = new(qguiapplication)
	qguiapplication.SetPointer(C.QGuiApplication_New_Int_String(C.int(argc), C.CString(argv)))
	qguiapplication.SetObjectName_String("QGuiApplication_" + randomIdentifier())
	return qguiapplication
}

func (p *qguiapplication) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QGuiApplication_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qguiapplication) IsSavingSession() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QGuiApplication_IsSavingSession(p.Pointer()) != 0
	}
}

func (p *qguiapplication) IsSessionRestored() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QGuiApplication_IsSessionRestored(p.Pointer()) != 0
	}
}

func (p *qguiapplication) SessionId() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QGuiApplication_SessionId(p.Pointer()))
	}
}

func (p *qguiapplication) SessionKey() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QGuiApplication_SessionKey(p.Pointer()))
	}
}

func (p *qguiapplication) ConnectSignalApplicationStateChanged(f func()) {
	C.QGuiApplication_ConnectSignalApplicationStateChanged(p.Pointer())
	connectSignal(p.ObjectName(), "applicationStateChanged", f)
}

func (p *qguiapplication) DisconnectSignalApplicationStateChanged() {
	C.QGuiApplication_DisconnectSignalApplicationStateChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "applicationStateChanged")
}

func (p *qguiapplication) SignalApplicationStateChanged() func() {
	return getSignal(p.ObjectName(), "applicationStateChanged")
}

func (p *qguiapplication) ConnectSignalFocusObjectChanged(f func()) {
	C.QGuiApplication_ConnectSignalFocusObjectChanged(p.Pointer())
	connectSignal(p.ObjectName(), "focusObjectChanged", f)
}

func (p *qguiapplication) DisconnectSignalFocusObjectChanged() {
	C.QGuiApplication_DisconnectSignalFocusObjectChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "focusObjectChanged")
}

func (p *qguiapplication) SignalFocusObjectChanged() func() {
	return getSignal(p.ObjectName(), "focusObjectChanged")
}

func (p *qguiapplication) ConnectSignalFontDatabaseChanged(f func()) {
	C.QGuiApplication_ConnectSignalFontDatabaseChanged(p.Pointer())
	connectSignal(p.ObjectName(), "fontDatabaseChanged", f)
}

func (p *qguiapplication) DisconnectSignalFontDatabaseChanged() {
	C.QGuiApplication_DisconnectSignalFontDatabaseChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "fontDatabaseChanged")
}

func (p *qguiapplication) SignalFontDatabaseChanged() func() {
	return getSignal(p.ObjectName(), "fontDatabaseChanged")
}

func (p *qguiapplication) ConnectSignalLastWindowClosed(f func()) {
	C.QGuiApplication_ConnectSignalLastWindowClosed(p.Pointer())
	connectSignal(p.ObjectName(), "lastWindowClosed", f)
}

func (p *qguiapplication) DisconnectSignalLastWindowClosed() {
	C.QGuiApplication_DisconnectSignalLastWindowClosed(p.Pointer())
	disconnectSignal(p.ObjectName(), "lastWindowClosed")
}

func (p *qguiapplication) SignalLastWindowClosed() func() {
	return getSignal(p.ObjectName(), "lastWindowClosed")
}

func QGuiApplication_ApplicationDisplayName() string {
	return C.GoString(C.QGuiApplication_ApplicationDisplayName())
}

func QGuiApplication_ApplicationState() ApplicationState {
	return ApplicationState(C.QGuiApplication_ApplicationState())
}

func QGuiApplication_DesktopSettingsAware() bool {
	return C.QGuiApplication_DesktopSettingsAware() != 0
}

func QGuiApplication_FocusObject() QObject {
	var qobject = new(qobject)
	qobject.SetPointer(C.QGuiApplication_FocusObject())
	if qobject.ObjectName() == "" {
		qobject.SetObjectName_String("QObject_" + randomIdentifier())
	}
	return qobject
}

func QGuiApplication_IsLeftToRight() bool {
	return C.QGuiApplication_IsLeftToRight() != 0
}

func QGuiApplication_IsRightToLeft() bool {
	return C.QGuiApplication_IsRightToLeft() != 0
}

func QGuiApplication_KeyboardModifiers() KeyboardModifier {
	return KeyboardModifier(C.QGuiApplication_KeyboardModifiers())
}

func QGuiApplication_LayoutDirection() LayoutDirection {
	return LayoutDirection(C.QGuiApplication_LayoutDirection())
}

func QGuiApplication_MouseButtons() MouseButton {
	return MouseButton(C.QGuiApplication_MouseButtons())
}

func QGuiApplication_PlatformName() string {
	return C.GoString(C.QGuiApplication_PlatformName())
}

func QGuiApplication_QueryKeyboardModifiers() KeyboardModifier {
	return KeyboardModifier(C.QGuiApplication_QueryKeyboardModifiers())
}

func QGuiApplication_QuitOnLastWindowClosed() bool {
	return C.QGuiApplication_QuitOnLastWindowClosed() != 0
}

func QGuiApplication_SetApplicationDisplayName_String(name string) {
	C.QGuiApplication_SetApplicationDisplayName_String(C.CString(name))
}

func QGuiApplication_SetDesktopSettingsAware_Bool(on bool) {
	C.QGuiApplication_SetDesktopSettingsAware_Bool(goBoolToCInt(on))
}

func QGuiApplication_SetLayoutDirection_LayoutDirection(direction LayoutDirection) {
	C.QGuiApplication_SetLayoutDirection_LayoutDirection(C.int(direction))
}

func QGuiApplication_SetQuitOnLastWindowClosed_Bool(quit bool) {
	C.QGuiApplication_SetQuitOnLastWindowClosed_Bool(goBoolToCInt(quit))
}

//export callbackQGuiApplication
func callbackQGuiApplication(ptr C.QtObjectPtr, msg *C.char) {
	var qguiapplication = new(qguiapplication)
	qguiapplication.SetPointer(ptr)
	getSignal(qguiapplication.ObjectName(), C.GoString(msg))()
}
