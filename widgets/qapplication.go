package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

type QApplication struct {
	gui.QGuiApplication
}

type QApplication_ITF interface {
	gui.QGuiApplication_ITF
	QApplication_PTR() *QApplication
}

func PointerFromQApplication(ptr QApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplication_PTR().Pointer()
	}
	return nil
}

func NewQApplicationFromPointer(ptr unsafe.Pointer) *QApplication {
	var n = new(QApplication)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QApplication_") {
		n.SetObjectName("QApplication_" + qt.Identifier())
	}
	return n
}

func (ptr *QApplication) QApplication_PTR() *QApplication {
	return ptr
}

//QApplication::ColorSpec
type QApplication__ColorSpec int64

const (
	QApplication__NormalColor = QApplication__ColorSpec(0)
	QApplication__CustomColor = QApplication__ColorSpec(1)
	QApplication__ManyColor   = QApplication__ColorSpec(2)
)

func QApplication_Alert(widget QWidget_ITF, msec int) {
	defer qt.Recovering("QApplication::alert")

	C.QApplication_QApplication_Alert(PointerFromQWidget(widget), C.int(msec))
}

func (ptr *QApplication) AutoMaximizeThreshold() int {
	defer qt.Recovering("QApplication::autoMaximizeThreshold")

	if ptr.Pointer() != nil {
		return int(C.QApplication_AutoMaximizeThreshold(ptr.Pointer()))
	}
	return 0
}

func (ptr *QApplication) AutoSipEnabled() bool {
	defer qt.Recovering("QApplication::autoSipEnabled")

	if ptr.Pointer() != nil {
		return C.QApplication_AutoSipEnabled(ptr.Pointer()) != 0
	}
	return false
}

func QApplication_Beep() {
	defer qt.Recovering("QApplication::beep")

	C.QApplication_QApplication_Beep()
}

func QApplication_CursorFlashTime() int {
	defer qt.Recovering("QApplication::cursorFlashTime")

	return int(C.QApplication_QApplication_CursorFlashTime())
}

func QApplication_DoubleClickInterval() int {
	defer qt.Recovering("QApplication::doubleClickInterval")

	return int(C.QApplication_QApplication_DoubleClickInterval())
}

func QApplication_GlobalStrut() *core.QSize {
	defer qt.Recovering("QApplication::globalStrut")

	return core.NewQSizeFromPointer(C.QApplication_QApplication_GlobalStrut())
}

func QApplication_IsEffectEnabled(effect core.Qt__UIEffect) bool {
	defer qt.Recovering("QApplication::isEffectEnabled")

	return C.QApplication_QApplication_IsEffectEnabled(C.int(effect)) != 0
}

func QApplication_KeyboardInputInterval() int {
	defer qt.Recovering("QApplication::keyboardInputInterval")

	return int(C.QApplication_QApplication_KeyboardInputInterval())
}

func QApplication_SetActiveWindow(active QWidget_ITF) {
	defer qt.Recovering("QApplication::setActiveWindow")

	C.QApplication_QApplication_SetActiveWindow(PointerFromQWidget(active))
}

func (ptr *QApplication) SetAutoMaximizeThreshold(threshold int) {
	defer qt.Recovering("QApplication::setAutoMaximizeThreshold")

	if ptr.Pointer() != nil {
		C.QApplication_SetAutoMaximizeThreshold(ptr.Pointer(), C.int(threshold))
	}
}

func (ptr *QApplication) SetAutoSipEnabled(enabled bool) {
	defer qt.Recovering("QApplication::setAutoSipEnabled")

	if ptr.Pointer() != nil {
		C.QApplication_SetAutoSipEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QApplication_SetCursorFlashTime(v int) {
	defer qt.Recovering("QApplication::setCursorFlashTime")

	C.QApplication_QApplication_SetCursorFlashTime(C.int(v))
}

func QApplication_SetDoubleClickInterval(v int) {
	defer qt.Recovering("QApplication::setDoubleClickInterval")

	C.QApplication_QApplication_SetDoubleClickInterval(C.int(v))
}

func QApplication_SetEffectEnabled(effect core.Qt__UIEffect, enable bool) {
	defer qt.Recovering("QApplication::setEffectEnabled")

	C.QApplication_QApplication_SetEffectEnabled(C.int(effect), C.int(qt.GoBoolToInt(enable)))
}

func QApplication_SetGlobalStrut(v core.QSize_ITF) {
	defer qt.Recovering("QApplication::setGlobalStrut")

	C.QApplication_QApplication_SetGlobalStrut(core.PointerFromQSize(v))
}

func QApplication_SetKeyboardInputInterval(v int) {
	defer qt.Recovering("QApplication::setKeyboardInputInterval")

	C.QApplication_QApplication_SetKeyboardInputInterval(C.int(v))
}

func QApplication_SetStartDragDistance(l int) {
	defer qt.Recovering("QApplication::setStartDragDistance")

	C.QApplication_QApplication_SetStartDragDistance(C.int(l))
}

func QApplication_SetStartDragTime(ms int) {
	defer qt.Recovering("QApplication::setStartDragTime")

	C.QApplication_QApplication_SetStartDragTime(C.int(ms))
}

func (ptr *QApplication) SetStyleSheet(sheet string) {
	defer qt.Recovering("QApplication::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QApplication_SetStyleSheet(ptr.Pointer(), C.CString(sheet))
	}
}

func QApplication_SetWheelScrollLines(v int) {
	defer qt.Recovering("QApplication::setWheelScrollLines")

	C.QApplication_QApplication_SetWheelScrollLines(C.int(v))
}

func QApplication_SetWindowIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QApplication::setWindowIcon")

	C.QApplication_QApplication_SetWindowIcon(gui.PointerFromQIcon(icon))
}

func QApplication_StartDragDistance() int {
	defer qt.Recovering("QApplication::startDragDistance")

	return int(C.QApplication_QApplication_StartDragDistance())
}

func QApplication_StartDragTime() int {
	defer qt.Recovering("QApplication::startDragTime")

	return int(C.QApplication_QApplication_StartDragTime())
}

func (ptr *QApplication) StyleSheet() string {
	defer qt.Recovering("QApplication::styleSheet")

	if ptr.Pointer() != nil {
		return C.GoString(C.QApplication_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func QApplication_TopLevelAt(point core.QPoint_ITF) *QWidget {
	defer qt.Recovering("QApplication::topLevelAt")

	return NewQWidgetFromPointer(C.QApplication_QApplication_TopLevelAt(core.PointerFromQPoint(point)))
}

func QApplication_WheelScrollLines() int {
	defer qt.Recovering("QApplication::wheelScrollLines")

	return int(C.QApplication_QApplication_WheelScrollLines())
}

func QApplication_WidgetAt(point core.QPoint_ITF) *QWidget {
	defer qt.Recovering("QApplication::widgetAt")

	return NewQWidgetFromPointer(C.QApplication_QApplication_WidgetAt(core.PointerFromQPoint(point)))
}

func QApplication_WindowIcon() *gui.QIcon {
	defer qt.Recovering("QApplication::windowIcon")

	return gui.NewQIconFromPointer(C.QApplication_QApplication_WindowIcon())
}

func NewQApplication(argc int, argv []string) *QApplication {
	defer qt.Recovering("QApplication::QApplication")

	return NewQApplicationFromPointer(C.QApplication_NewQApplication(C.int(argc), C.CString(strings.Join(argv, ",,,"))))
}

func QApplication_AboutQt() {
	defer qt.Recovering("QApplication::aboutQt")

	C.QApplication_QApplication_AboutQt()
}

func QApplication_ActiveModalWidget() *QWidget {
	defer qt.Recovering("QApplication::activeModalWidget")

	return NewQWidgetFromPointer(C.QApplication_QApplication_ActiveModalWidget())
}

func QApplication_ActivePopupWidget() *QWidget {
	defer qt.Recovering("QApplication::activePopupWidget")

	return NewQWidgetFromPointer(C.QApplication_QApplication_ActivePopupWidget())
}

func QApplication_ActiveWindow() *QWidget {
	defer qt.Recovering("QApplication::activeWindow")

	return NewQWidgetFromPointer(C.QApplication_QApplication_ActiveWindow())
}

func QApplication_CloseAllWindows() {
	defer qt.Recovering("QApplication::closeAllWindows")

	C.QApplication_QApplication_CloseAllWindows()
}

func QApplication_ColorSpec() int {
	defer qt.Recovering("QApplication::colorSpec")

	return int(C.QApplication_QApplication_ColorSpec())
}

func QApplication_Desktop() *QDesktopWidget {
	defer qt.Recovering("QApplication::desktop")

	return NewQDesktopWidgetFromPointer(C.QApplication_QApplication_Desktop())
}

func (ptr *QApplication) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QApplication::event")

	if ptr.Pointer() != nil {
		return C.QApplication_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func QApplication_Exec() int {
	defer qt.Recovering("QApplication::exec")

	return int(C.QApplication_QApplication_Exec())
}

func (ptr *QApplication) ConnectFocusChanged(f func(old *QWidget, now *QWidget)) {
	defer qt.Recovering("connect QApplication::focusChanged")

	if ptr.Pointer() != nil {
		C.QApplication_ConnectFocusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusChanged", f)
	}
}

func (ptr *QApplication) DisconnectFocusChanged() {
	defer qt.Recovering("disconnect QApplication::focusChanged")

	if ptr.Pointer() != nil {
		C.QApplication_DisconnectFocusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusChanged")
	}
}

//export callbackQApplicationFocusChanged
func callbackQApplicationFocusChanged(ptr unsafe.Pointer, ptrName *C.char, old unsafe.Pointer, now unsafe.Pointer) {
	defer qt.Recovering("callback QApplication::focusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusChanged"); signal != nil {
		signal.(func(*QWidget, *QWidget))(NewQWidgetFromPointer(old), NewQWidgetFromPointer(now))
	}

}

func (ptr *QApplication) FocusChanged(old QWidget_ITF, now QWidget_ITF) {
	defer qt.Recovering("QApplication::focusChanged")

	if ptr.Pointer() != nil {
		C.QApplication_FocusChanged(ptr.Pointer(), PointerFromQWidget(old), PointerFromQWidget(now))
	}
}

func QApplication_FocusWidget() *QWidget {
	defer qt.Recovering("QApplication::focusWidget")

	return NewQWidgetFromPointer(C.QApplication_QApplication_FocusWidget())
}

func (ptr *QApplication) Notify(receiver core.QObject_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QApplication::notify")

	if ptr.Pointer() != nil {
		return C.QApplication_Notify(ptr.Pointer(), core.PointerFromQObject(receiver), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func QApplication_SetColorSpec(spec int) {
	defer qt.Recovering("QApplication::setColorSpec")

	C.QApplication_QApplication_SetColorSpec(C.int(spec))
}

func QApplication_SetFont(font gui.QFont_ITF, className string) {
	defer qt.Recovering("QApplication::setFont")

	C.QApplication_QApplication_SetFont(gui.PointerFromQFont(font), C.CString(className))
}

func QApplication_SetPalette(palette gui.QPalette_ITF, className string) {
	defer qt.Recovering("QApplication::setPalette")

	C.QApplication_QApplication_SetPalette(gui.PointerFromQPalette(palette), C.CString(className))
}

func QApplication_SetStyle2(style string) *QStyle {
	defer qt.Recovering("QApplication::setStyle")

	return NewQStyleFromPointer(C.QApplication_QApplication_SetStyle2(C.CString(style)))
}

func QApplication_SetStyle(style QStyle_ITF) {
	defer qt.Recovering("QApplication::setStyle")

	C.QApplication_QApplication_SetStyle(PointerFromQStyle(style))
}

func QApplication_Style() *QStyle {
	defer qt.Recovering("QApplication::style")

	return NewQStyleFromPointer(C.QApplication_QApplication_Style())
}

func QApplication_TopLevelAt2(x int, y int) *QWidget {
	defer qt.Recovering("QApplication::topLevelAt")

	return NewQWidgetFromPointer(C.QApplication_QApplication_TopLevelAt2(C.int(x), C.int(y)))
}

func QApplication_WidgetAt2(x int, y int) *QWidget {
	defer qt.Recovering("QApplication::widgetAt")

	return NewQWidgetFromPointer(C.QApplication_QApplication_WidgetAt2(C.int(x), C.int(y)))
}

func (ptr *QApplication) DestroyQApplication() {
	defer qt.Recovering("QApplication::~QApplication")

	if ptr.Pointer() != nil {
		C.QApplication_DestroyQApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QApplication) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QApplication::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QApplication) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QApplication::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQApplicationTimerEvent
func callbackQApplicationTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QApplication::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQApplicationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QApplication) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QApplication::timerEvent")

	if ptr.Pointer() != nil {
		C.QApplication_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QApplication) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QApplication::timerEvent")

	if ptr.Pointer() != nil {
		C.QApplication_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QApplication) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QApplication::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QApplication) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QApplication::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQApplicationChildEvent
func callbackQApplicationChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QApplication::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQApplicationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QApplication) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QApplication::childEvent")

	if ptr.Pointer() != nil {
		C.QApplication_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QApplication) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QApplication::childEvent")

	if ptr.Pointer() != nil {
		C.QApplication_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QApplication) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QApplication::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QApplication) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QApplication::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQApplicationCustomEvent
func callbackQApplicationCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QApplication::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQApplicationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QApplication) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QApplication::customEvent")

	if ptr.Pointer() != nil {
		C.QApplication_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QApplication) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QApplication::customEvent")

	if ptr.Pointer() != nil {
		C.QApplication_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
