package gui

//#include "qguiapplication.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGuiApplication struct {
	core.QCoreApplication
}

type QGuiApplicationITF interface {
	core.QCoreApplicationITF
	QGuiApplicationPTR() *QGuiApplication
}

func PointerFromQGuiApplication(ptr QGuiApplicationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGuiApplicationPTR().Pointer()
	}
	return nil
}

func QGuiApplicationFromPointer(ptr unsafe.Pointer) *QGuiApplication {
	var n = new(QGuiApplication)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGuiApplication_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGuiApplication) QGuiApplicationPTR() *QGuiApplication {
	return ptr
}

func QGuiApplication_ApplicationDisplayName() string {
	return C.GoString(C.QGuiApplication_QGuiApplication_ApplicationDisplayName())
}

func QGuiApplication_ApplicationState() core.Qt__ApplicationState {
	return core.Qt__ApplicationState(C.QGuiApplication_QGuiApplication_ApplicationState())
}

func (ptr *QGuiApplication) IsSavingSession() bool {
	if ptr.Pointer() != nil {
		return C.QGuiApplication_IsSavingSession(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGuiApplication) IsSessionRestored() bool {
	if ptr.Pointer() != nil {
		return C.QGuiApplication_IsSessionRestored(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QGuiApplication_LayoutDirection() core.Qt__LayoutDirection {
	return core.Qt__LayoutDirection(C.QGuiApplication_QGuiApplication_LayoutDirection())
}

func QGuiApplication_OverrideCursor() *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_OverrideCursor()))
}

func QGuiApplication_PlatformName() string {
	return C.GoString(C.QGuiApplication_QGuiApplication_PlatformName())
}

func QGuiApplication_QueryKeyboardModifiers() core.Qt__KeyboardModifier {
	return core.Qt__KeyboardModifier(C.QGuiApplication_QGuiApplication_QueryKeyboardModifiers())
}

func QGuiApplication_QuitOnLastWindowClosed() bool {
	return C.QGuiApplication_QGuiApplication_QuitOnLastWindowClosed() != 0
}

func QGuiApplication_RestoreOverrideCursor() {
	C.QGuiApplication_QGuiApplication_RestoreOverrideCursor()
}

func (ptr *QGuiApplication) SessionId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGuiApplication_SessionId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGuiApplication) SessionKey() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGuiApplication_SessionKey(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QGuiApplication_SetApplicationDisplayName(name string) {
	C.QGuiApplication_QGuiApplication_SetApplicationDisplayName(C.CString(name))
}

func QGuiApplication_SetLayoutDirection(direction core.Qt__LayoutDirection) {
	C.QGuiApplication_QGuiApplication_SetLayoutDirection(C.int(direction))
}

func QGuiApplication_SetOverrideCursor(cursor QCursorITF) {
	C.QGuiApplication_QGuiApplication_SetOverrideCursor(C.QtObjectPtr(PointerFromQCursor(cursor)))
}

func QGuiApplication_SetQuitOnLastWindowClosed(quit bool) {
	C.QGuiApplication_QGuiApplication_SetQuitOnLastWindowClosed(C.int(qt.GoBoolToInt(quit)))
}

func QGuiApplication_SetWindowIcon(icon QIconITF) {
	C.QGuiApplication_QGuiApplication_SetWindowIcon(C.QtObjectPtr(PointerFromQIcon(icon)))
}

func NewQGuiApplication(argc int, argv string) *QGuiApplication {
	return QGuiApplicationFromPointer(unsafe.Pointer(C.QGuiApplication_NewQGuiApplication(C.int(argc), C.CString(argv))))
}

func (ptr *QGuiApplication) ConnectApplicationStateChanged(f func(state core.Qt__ApplicationState)) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectApplicationStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "applicationStateChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectApplicationStateChanged() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectApplicationStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "applicationStateChanged")
	}
}

//export callbackQGuiApplicationApplicationStateChanged
func callbackQGuiApplicationApplicationStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "applicationStateChanged").(func(core.Qt__ApplicationState))(core.Qt__ApplicationState(state))
}

func QGuiApplication_ChangeOverrideCursor(cursor QCursorITF) {
	C.QGuiApplication_QGuiApplication_ChangeOverrideCursor(C.QtObjectPtr(PointerFromQCursor(cursor)))
}

func QGuiApplication_Clipboard() *QClipboard {
	return QClipboardFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_Clipboard()))
}

func QGuiApplication_DesktopSettingsAware() bool {
	return C.QGuiApplication_QGuiApplication_DesktopSettingsAware() != 0
}

func QGuiApplication_Exec() int {
	return int(C.QGuiApplication_QGuiApplication_Exec())
}

func QGuiApplication_FocusObject() *core.QObject {
	return core.QObjectFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_FocusObject()))
}

func (ptr *QGuiApplication) ConnectFocusObjectChanged(f func(focusObject core.QObjectITF)) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectFocusObjectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusObjectChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectFocusObjectChanged() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectFocusObjectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusObjectChanged")
	}
}

//export callbackQGuiApplicationFocusObjectChanged
func callbackQGuiApplicationFocusObjectChanged(ptrName *C.char, focusObject unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "focusObjectChanged").(func(*core.QObject))(core.QObjectFromPointer(focusObject))
}

func QGuiApplication_FocusWindow() *QWindow {
	return QWindowFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_FocusWindow()))
}

func (ptr *QGuiApplication) ConnectFocusWindowChanged(f func(focusWindow QWindowITF)) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectFocusWindowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusWindowChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectFocusWindowChanged() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectFocusWindowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusWindowChanged")
	}
}

//export callbackQGuiApplicationFocusWindowChanged
func callbackQGuiApplicationFocusWindowChanged(ptrName *C.char, focusWindow unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "focusWindowChanged").(func(*QWindow))(QWindowFromPointer(focusWindow))
}

func (ptr *QGuiApplication) ConnectFontDatabaseChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectFontDatabaseChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fontDatabaseChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectFontDatabaseChanged() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectFontDatabaseChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fontDatabaseChanged")
	}
}

//export callbackQGuiApplicationFontDatabaseChanged
func callbackQGuiApplicationFontDatabaseChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fontDatabaseChanged").(func())()
}

func QGuiApplication_InputMethod() *QInputMethod {
	return QInputMethodFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_InputMethod()))
}

func QGuiApplication_IsLeftToRight() bool {
	return C.QGuiApplication_QGuiApplication_IsLeftToRight() != 0
}

func QGuiApplication_IsRightToLeft() bool {
	return C.QGuiApplication_QGuiApplication_IsRightToLeft() != 0
}

func QGuiApplication_KeyboardModifiers() core.Qt__KeyboardModifier {
	return core.Qt__KeyboardModifier(C.QGuiApplication_QGuiApplication_KeyboardModifiers())
}

func (ptr *QGuiApplication) ConnectLastWindowClosed(f func()) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectLastWindowClosed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "lastWindowClosed", f)
	}
}

func (ptr *QGuiApplication) DisconnectLastWindowClosed() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectLastWindowClosed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "lastWindowClosed")
	}
}

//export callbackQGuiApplicationLastWindowClosed
func callbackQGuiApplicationLastWindowClosed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "lastWindowClosed").(func())()
}

func (ptr *QGuiApplication) ConnectLayoutDirectionChanged(f func(direction core.Qt__LayoutDirection)) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectLayoutDirectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "layoutDirectionChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectLayoutDirectionChanged() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectLayoutDirectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "layoutDirectionChanged")
	}
}

//export callbackQGuiApplicationLayoutDirectionChanged
func callbackQGuiApplicationLayoutDirectionChanged(ptrName *C.char, direction C.int) {
	qt.GetSignal(C.GoString(ptrName), "layoutDirectionChanged").(func(core.Qt__LayoutDirection))(core.Qt__LayoutDirection(direction))
}

func QGuiApplication_ModalWindow() *QWindow {
	return QWindowFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_ModalWindow()))
}

func QGuiApplication_MouseButtons() core.Qt__MouseButton {
	return core.Qt__MouseButton(C.QGuiApplication_QGuiApplication_MouseButtons())
}

func (ptr *QGuiApplication) Notify(object core.QObjectITF, event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QGuiApplication_Notify(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func QGuiApplication_PrimaryScreen() *QScreen {
	return QScreenFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_PrimaryScreen()))
}

func (ptr *QGuiApplication) ConnectScreenAdded(f func(screen QScreenITF)) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectScreenAdded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "screenAdded", f)
	}
}

func (ptr *QGuiApplication) DisconnectScreenAdded() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectScreenAdded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "screenAdded")
	}
}

//export callbackQGuiApplicationScreenAdded
func callbackQGuiApplicationScreenAdded(ptrName *C.char, screen unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "screenAdded").(func(*QScreen))(QScreenFromPointer(screen))
}

func (ptr *QGuiApplication) ConnectScreenRemoved(f func(screen QScreenITF)) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectScreenRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "screenRemoved", f)
	}
}

func (ptr *QGuiApplication) DisconnectScreenRemoved() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectScreenRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "screenRemoved")
	}
}

//export callbackQGuiApplicationScreenRemoved
func callbackQGuiApplicationScreenRemoved(ptrName *C.char, screen unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "screenRemoved").(func(*QScreen))(QScreenFromPointer(screen))
}

func QGuiApplication_SetDesktopSettingsAware(on bool) {
	C.QGuiApplication_QGuiApplication_SetDesktopSettingsAware(C.int(qt.GoBoolToInt(on)))
}

func QGuiApplication_SetFont(font QFontITF) {
	C.QGuiApplication_QGuiApplication_SetFont(C.QtObjectPtr(PointerFromQFont(font)))
}

func QGuiApplication_SetPalette(pal QPaletteITF) {
	C.QGuiApplication_QGuiApplication_SetPalette(C.QtObjectPtr(PointerFromQPalette(pal)))
}

func QGuiApplication_StyleHints() *QStyleHints {
	return QStyleHintsFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_StyleHints()))
}

func QGuiApplication_Sync() {
	C.QGuiApplication_QGuiApplication_Sync()
}

func QGuiApplication_TopLevelAt(pos core.QPointITF) *QWindow {
	return QWindowFromPointer(unsafe.Pointer(C.QGuiApplication_QGuiApplication_TopLevelAt(C.QtObjectPtr(core.PointerFromQPoint(pos)))))
}

func (ptr *QGuiApplication) DestroyQGuiApplication() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DestroyQGuiApplication(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
