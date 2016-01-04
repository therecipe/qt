package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGuiApplication struct {
	core.QCoreApplication
}

type QGuiApplication_ITF interface {
	core.QCoreApplication_ITF
	QGuiApplication_PTR() *QGuiApplication
}

func PointerFromQGuiApplication(ptr QGuiApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGuiApplication_PTR().Pointer()
	}
	return nil
}

func NewQGuiApplicationFromPointer(ptr unsafe.Pointer) *QGuiApplication {
	var n = new(QGuiApplication)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGuiApplication_") {
		n.SetObjectName("QGuiApplication_" + qt.Identifier())
	}
	return n
}

func (ptr *QGuiApplication) QGuiApplication_PTR() *QGuiApplication {
	return ptr
}

func QGuiApplication_ApplicationDisplayName() string {
	defer qt.Recovering("QGuiApplication::applicationDisplayName")

	return C.GoString(C.QGuiApplication_QGuiApplication_ApplicationDisplayName())
}

func QGuiApplication_ApplicationState() core.Qt__ApplicationState {
	defer qt.Recovering("QGuiApplication::applicationState")

	return core.Qt__ApplicationState(C.QGuiApplication_QGuiApplication_ApplicationState())
}

func (ptr *QGuiApplication) IsSavingSession() bool {
	defer qt.Recovering("QGuiApplication::isSavingSession")

	if ptr.Pointer() != nil {
		return C.QGuiApplication_IsSavingSession(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGuiApplication) IsSessionRestored() bool {
	defer qt.Recovering("QGuiApplication::isSessionRestored")

	if ptr.Pointer() != nil {
		return C.QGuiApplication_IsSessionRestored(ptr.Pointer()) != 0
	}
	return false
}

func QGuiApplication_LayoutDirection() core.Qt__LayoutDirection {
	defer qt.Recovering("QGuiApplication::layoutDirection")

	return core.Qt__LayoutDirection(C.QGuiApplication_QGuiApplication_LayoutDirection())
}

func QGuiApplication_OverrideCursor() *QCursor {
	defer qt.Recovering("QGuiApplication::overrideCursor")

	return NewQCursorFromPointer(C.QGuiApplication_QGuiApplication_OverrideCursor())
}

func QGuiApplication_PlatformName() string {
	defer qt.Recovering("QGuiApplication::platformName")

	return C.GoString(C.QGuiApplication_QGuiApplication_PlatformName())
}

func QGuiApplication_QueryKeyboardModifiers() core.Qt__KeyboardModifier {
	defer qt.Recovering("QGuiApplication::queryKeyboardModifiers")

	return core.Qt__KeyboardModifier(C.QGuiApplication_QGuiApplication_QueryKeyboardModifiers())
}

func QGuiApplication_QuitOnLastWindowClosed() bool {
	defer qt.Recovering("QGuiApplication::quitOnLastWindowClosed")

	return C.QGuiApplication_QGuiApplication_QuitOnLastWindowClosed() != 0
}

func QGuiApplication_RestoreOverrideCursor() {
	defer qt.Recovering("QGuiApplication::restoreOverrideCursor")

	C.QGuiApplication_QGuiApplication_RestoreOverrideCursor()
}

func (ptr *QGuiApplication) SessionId() string {
	defer qt.Recovering("QGuiApplication::sessionId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGuiApplication_SessionId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGuiApplication) SessionKey() string {
	defer qt.Recovering("QGuiApplication::sessionKey")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGuiApplication_SessionKey(ptr.Pointer()))
	}
	return ""
}

func QGuiApplication_SetApplicationDisplayName(name string) {
	defer qt.Recovering("QGuiApplication::setApplicationDisplayName")

	C.QGuiApplication_QGuiApplication_SetApplicationDisplayName(C.CString(name))
}

func QGuiApplication_SetLayoutDirection(direction core.Qt__LayoutDirection) {
	defer qt.Recovering("QGuiApplication::setLayoutDirection")

	C.QGuiApplication_QGuiApplication_SetLayoutDirection(C.int(direction))
}

func QGuiApplication_SetOverrideCursor(cursor QCursor_ITF) {
	defer qt.Recovering("QGuiApplication::setOverrideCursor")

	C.QGuiApplication_QGuiApplication_SetOverrideCursor(PointerFromQCursor(cursor))
}

func QGuiApplication_SetQuitOnLastWindowClosed(quit bool) {
	defer qt.Recovering("QGuiApplication::setQuitOnLastWindowClosed")

	C.QGuiApplication_QGuiApplication_SetQuitOnLastWindowClosed(C.int(qt.GoBoolToInt(quit)))
}

func QGuiApplication_SetWindowIcon(icon QIcon_ITF) {
	defer qt.Recovering("QGuiApplication::setWindowIcon")

	C.QGuiApplication_QGuiApplication_SetWindowIcon(PointerFromQIcon(icon))
}

func QGuiApplication_WindowIcon() *QIcon {
	defer qt.Recovering("QGuiApplication::windowIcon")

	return NewQIconFromPointer(C.QGuiApplication_QGuiApplication_WindowIcon())
}

func NewQGuiApplication(argc int, argv []string) *QGuiApplication {
	defer qt.Recovering("QGuiApplication::QGuiApplication")

	return NewQGuiApplicationFromPointer(C.QGuiApplication_NewQGuiApplication(C.int(argc), C.CString(strings.Join(argv, ",,,"))))
}

func (ptr *QGuiApplication) ConnectApplicationStateChanged(f func(state core.Qt__ApplicationState)) {
	defer qt.Recovering("connect QGuiApplication::applicationStateChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectApplicationStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "applicationStateChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectApplicationStateChanged() {
	defer qt.Recovering("disconnect QGuiApplication::applicationStateChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectApplicationStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "applicationStateChanged")
	}
}

//export callbackQGuiApplicationApplicationStateChanged
func callbackQGuiApplicationApplicationStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QGuiApplication::applicationStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "applicationStateChanged"); signal != nil {
		signal.(func(core.Qt__ApplicationState))(core.Qt__ApplicationState(state))
	}

}

func (ptr *QGuiApplication) ApplicationStateChanged(state core.Qt__ApplicationState) {
	defer qt.Recovering("QGuiApplication::applicationStateChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ApplicationStateChanged(ptr.Pointer(), C.int(state))
	}
}

func QGuiApplication_ChangeOverrideCursor(cursor QCursor_ITF) {
	defer qt.Recovering("QGuiApplication::changeOverrideCursor")

	C.QGuiApplication_QGuiApplication_ChangeOverrideCursor(PointerFromQCursor(cursor))
}

func QGuiApplication_Clipboard() *QClipboard {
	defer qt.Recovering("QGuiApplication::clipboard")

	return NewQClipboardFromPointer(C.QGuiApplication_QGuiApplication_Clipboard())
}

func QGuiApplication_DesktopSettingsAware() bool {
	defer qt.Recovering("QGuiApplication::desktopSettingsAware")

	return C.QGuiApplication_QGuiApplication_DesktopSettingsAware() != 0
}

func (ptr *QGuiApplication) DevicePixelRatio() float64 {
	defer qt.Recovering("QGuiApplication::devicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QGuiApplication_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGuiApplication) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QGuiApplication::event")

	if ptr.Pointer() != nil {
		return C.QGuiApplication_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func QGuiApplication_Exec() int {
	defer qt.Recovering("QGuiApplication::exec")

	return int(C.QGuiApplication_QGuiApplication_Exec())
}

func QGuiApplication_FocusObject() *core.QObject {
	defer qt.Recovering("QGuiApplication::focusObject")

	return core.NewQObjectFromPointer(C.QGuiApplication_QGuiApplication_FocusObject())
}

func (ptr *QGuiApplication) ConnectFocusObjectChanged(f func(focusObject *core.QObject)) {
	defer qt.Recovering("connect QGuiApplication::focusObjectChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectFocusObjectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusObjectChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectFocusObjectChanged() {
	defer qt.Recovering("disconnect QGuiApplication::focusObjectChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectFocusObjectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusObjectChanged")
	}
}

//export callbackQGuiApplicationFocusObjectChanged
func callbackQGuiApplicationFocusObjectChanged(ptr unsafe.Pointer, ptrName *C.char, focusObject unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::focusObjectChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusObjectChanged"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(focusObject))
	}

}

func (ptr *QGuiApplication) FocusObjectChanged(focusObject core.QObject_ITF) {
	defer qt.Recovering("QGuiApplication::focusObjectChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_FocusObjectChanged(ptr.Pointer(), core.PointerFromQObject(focusObject))
	}
}

func QGuiApplication_FocusWindow() *QWindow {
	defer qt.Recovering("QGuiApplication::focusWindow")

	return NewQWindowFromPointer(C.QGuiApplication_QGuiApplication_FocusWindow())
}

func (ptr *QGuiApplication) ConnectFocusWindowChanged(f func(focusWindow *QWindow)) {
	defer qt.Recovering("connect QGuiApplication::focusWindowChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectFocusWindowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusWindowChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectFocusWindowChanged() {
	defer qt.Recovering("disconnect QGuiApplication::focusWindowChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectFocusWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusWindowChanged")
	}
}

//export callbackQGuiApplicationFocusWindowChanged
func callbackQGuiApplicationFocusWindowChanged(ptr unsafe.Pointer, ptrName *C.char, focusWindow unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::focusWindowChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusWindowChanged"); signal != nil {
		signal.(func(*QWindow))(NewQWindowFromPointer(focusWindow))
	}

}

func (ptr *QGuiApplication) FocusWindowChanged(focusWindow QWindow_ITF) {
	defer qt.Recovering("QGuiApplication::focusWindowChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_FocusWindowChanged(ptr.Pointer(), PointerFromQWindow(focusWindow))
	}
}

func (ptr *QGuiApplication) ConnectFontDatabaseChanged(f func()) {
	defer qt.Recovering("connect QGuiApplication::fontDatabaseChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectFontDatabaseChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fontDatabaseChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectFontDatabaseChanged() {
	defer qt.Recovering("disconnect QGuiApplication::fontDatabaseChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectFontDatabaseChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fontDatabaseChanged")
	}
}

//export callbackQGuiApplicationFontDatabaseChanged
func callbackQGuiApplicationFontDatabaseChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGuiApplication::fontDatabaseChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fontDatabaseChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGuiApplication) FontDatabaseChanged() {
	defer qt.Recovering("QGuiApplication::fontDatabaseChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_FontDatabaseChanged(ptr.Pointer())
	}
}

func QGuiApplication_InputMethod() *QInputMethod {
	defer qt.Recovering("QGuiApplication::inputMethod")

	return NewQInputMethodFromPointer(C.QGuiApplication_QGuiApplication_InputMethod())
}

func QGuiApplication_IsLeftToRight() bool {
	defer qt.Recovering("QGuiApplication::isLeftToRight")

	return C.QGuiApplication_QGuiApplication_IsLeftToRight() != 0
}

func QGuiApplication_IsRightToLeft() bool {
	defer qt.Recovering("QGuiApplication::isRightToLeft")

	return C.QGuiApplication_QGuiApplication_IsRightToLeft() != 0
}

func QGuiApplication_KeyboardModifiers() core.Qt__KeyboardModifier {
	defer qt.Recovering("QGuiApplication::keyboardModifiers")

	return core.Qt__KeyboardModifier(C.QGuiApplication_QGuiApplication_KeyboardModifiers())
}

func (ptr *QGuiApplication) ConnectLastWindowClosed(f func()) {
	defer qt.Recovering("connect QGuiApplication::lastWindowClosed")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectLastWindowClosed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lastWindowClosed", f)
	}
}

func (ptr *QGuiApplication) DisconnectLastWindowClosed() {
	defer qt.Recovering("disconnect QGuiApplication::lastWindowClosed")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectLastWindowClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lastWindowClosed")
	}
}

//export callbackQGuiApplicationLastWindowClosed
func callbackQGuiApplicationLastWindowClosed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGuiApplication::lastWindowClosed")

	if signal := qt.GetSignal(C.GoString(ptrName), "lastWindowClosed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGuiApplication) LastWindowClosed() {
	defer qt.Recovering("QGuiApplication::lastWindowClosed")

	if ptr.Pointer() != nil {
		C.QGuiApplication_LastWindowClosed(ptr.Pointer())
	}
}

func (ptr *QGuiApplication) ConnectLayoutDirectionChanged(f func(direction core.Qt__LayoutDirection)) {
	defer qt.Recovering("connect QGuiApplication::layoutDirectionChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectLayoutDirectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "layoutDirectionChanged", f)
	}
}

func (ptr *QGuiApplication) DisconnectLayoutDirectionChanged() {
	defer qt.Recovering("disconnect QGuiApplication::layoutDirectionChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectLayoutDirectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "layoutDirectionChanged")
	}
}

//export callbackQGuiApplicationLayoutDirectionChanged
func callbackQGuiApplicationLayoutDirectionChanged(ptr unsafe.Pointer, ptrName *C.char, direction C.int) {
	defer qt.Recovering("callback QGuiApplication::layoutDirectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "layoutDirectionChanged"); signal != nil {
		signal.(func(core.Qt__LayoutDirection))(core.Qt__LayoutDirection(direction))
	}

}

func (ptr *QGuiApplication) LayoutDirectionChanged(direction core.Qt__LayoutDirection) {
	defer qt.Recovering("QGuiApplication::layoutDirectionChanged")

	if ptr.Pointer() != nil {
		C.QGuiApplication_LayoutDirectionChanged(ptr.Pointer(), C.int(direction))
	}
}

func QGuiApplication_ModalWindow() *QWindow {
	defer qt.Recovering("QGuiApplication::modalWindow")

	return NewQWindowFromPointer(C.QGuiApplication_QGuiApplication_ModalWindow())
}

func QGuiApplication_MouseButtons() core.Qt__MouseButton {
	defer qt.Recovering("QGuiApplication::mouseButtons")

	return core.Qt__MouseButton(C.QGuiApplication_QGuiApplication_MouseButtons())
}

func (ptr *QGuiApplication) Notify(object core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGuiApplication::notify")

	if ptr.Pointer() != nil {
		return C.QGuiApplication_Notify(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func QGuiApplication_PrimaryScreen() *QScreen {
	defer qt.Recovering("QGuiApplication::primaryScreen")

	return NewQScreenFromPointer(C.QGuiApplication_QGuiApplication_PrimaryScreen())
}

func (ptr *QGuiApplication) ConnectScreenAdded(f func(screen *QScreen)) {
	defer qt.Recovering("connect QGuiApplication::screenAdded")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectScreenAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenAdded", f)
	}
}

func (ptr *QGuiApplication) DisconnectScreenAdded() {
	defer qt.Recovering("disconnect QGuiApplication::screenAdded")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectScreenAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenAdded")
	}
}

//export callbackQGuiApplicationScreenAdded
func callbackQGuiApplicationScreenAdded(ptr unsafe.Pointer, ptrName *C.char, screen unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::screenAdded")

	if signal := qt.GetSignal(C.GoString(ptrName), "screenAdded"); signal != nil {
		signal.(func(*QScreen))(NewQScreenFromPointer(screen))
	}

}

func (ptr *QGuiApplication) ScreenAdded(screen QScreen_ITF) {
	defer qt.Recovering("QGuiApplication::screenAdded")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ScreenAdded(ptr.Pointer(), PointerFromQScreen(screen))
	}
}

func (ptr *QGuiApplication) ConnectScreenRemoved(f func(screen *QScreen)) {
	defer qt.Recovering("connect QGuiApplication::screenRemoved")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectScreenRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenRemoved", f)
	}
}

func (ptr *QGuiApplication) DisconnectScreenRemoved() {
	defer qt.Recovering("disconnect QGuiApplication::screenRemoved")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectScreenRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenRemoved")
	}
}

//export callbackQGuiApplicationScreenRemoved
func callbackQGuiApplicationScreenRemoved(ptr unsafe.Pointer, ptrName *C.char, screen unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::screenRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "screenRemoved"); signal != nil {
		signal.(func(*QScreen))(NewQScreenFromPointer(screen))
	}

}

func (ptr *QGuiApplication) ScreenRemoved(screen QScreen_ITF) {
	defer qt.Recovering("QGuiApplication::screenRemoved")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ScreenRemoved(ptr.Pointer(), PointerFromQScreen(screen))
	}
}

func QGuiApplication_SetDesktopSettingsAware(on bool) {
	defer qt.Recovering("QGuiApplication::setDesktopSettingsAware")

	C.QGuiApplication_QGuiApplication_SetDesktopSettingsAware(C.int(qt.GoBoolToInt(on)))
}

func QGuiApplication_SetFont(font QFont_ITF) {
	defer qt.Recovering("QGuiApplication::setFont")

	C.QGuiApplication_QGuiApplication_SetFont(PointerFromQFont(font))
}

func QGuiApplication_SetPalette(pal QPalette_ITF) {
	defer qt.Recovering("QGuiApplication::setPalette")

	C.QGuiApplication_QGuiApplication_SetPalette(PointerFromQPalette(pal))
}

func QGuiApplication_StyleHints() *QStyleHints {
	defer qt.Recovering("QGuiApplication::styleHints")

	return NewQStyleHintsFromPointer(C.QGuiApplication_QGuiApplication_StyleHints())
}

func QGuiApplication_Sync() {
	defer qt.Recovering("QGuiApplication::sync")

	C.QGuiApplication_QGuiApplication_Sync()
}

func QGuiApplication_TopLevelAt(pos core.QPoint_ITF) *QWindow {
	defer qt.Recovering("QGuiApplication::topLevelAt")

	return NewQWindowFromPointer(C.QGuiApplication_QGuiApplication_TopLevelAt(core.PointerFromQPoint(pos)))
}

func (ptr *QGuiApplication) DestroyQGuiApplication() {
	defer qt.Recovering("QGuiApplication::~QGuiApplication")

	if ptr.Pointer() != nil {
		C.QGuiApplication_DestroyQGuiApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGuiApplication) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGuiApplication::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGuiApplication) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGuiApplication::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGuiApplicationTimerEvent
func callbackQGuiApplicationTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGuiApplicationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGuiApplication) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGuiApplication::timerEvent")

	if ptr.Pointer() != nil {
		C.QGuiApplication_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGuiApplication) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGuiApplication::timerEvent")

	if ptr.Pointer() != nil {
		C.QGuiApplication_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGuiApplication) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGuiApplication::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGuiApplication) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGuiApplication::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGuiApplicationChildEvent
func callbackQGuiApplicationChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGuiApplicationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGuiApplication) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGuiApplication::childEvent")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGuiApplication) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGuiApplication::childEvent")

	if ptr.Pointer() != nil {
		C.QGuiApplication_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGuiApplication) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGuiApplication::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGuiApplication) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGuiApplication::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGuiApplicationCustomEvent
func callbackQGuiApplicationCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGuiApplication::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGuiApplicationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGuiApplication) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGuiApplication::customEvent")

	if ptr.Pointer() != nil {
		C.QGuiApplication_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGuiApplication) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGuiApplication::customEvent")

	if ptr.Pointer() != nil {
		C.QGuiApplication_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
