package widgets

//#include "qapplication.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QApplication_" + qt.RandomIdentifier())
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
	C.QApplication_QApplication_Alert(PointerFromQWidget(widget), C.int(msec))
}

func (ptr *QApplication) AutoMaximizeThreshold() int {
	if ptr.Pointer() != nil {
		return int(C.QApplication_AutoMaximizeThreshold(ptr.Pointer()))
	}
	return 0
}

func (ptr *QApplication) AutoSipEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QApplication_AutoSipEnabled(ptr.Pointer()) != 0
	}
	return false
}

func QApplication_Beep() {
	C.QApplication_QApplication_Beep()
}

func QApplication_CursorFlashTime() int {
	return int(C.QApplication_QApplication_CursorFlashTime())
}

func QApplication_DoubleClickInterval() int {
	return int(C.QApplication_QApplication_DoubleClickInterval())
}

func QApplication_IsEffectEnabled(effect core.Qt__UIEffect) bool {
	return C.QApplication_QApplication_IsEffectEnabled(C.int(effect)) != 0
}

func QApplication_KeyboardInputInterval() int {
	return int(C.QApplication_QApplication_KeyboardInputInterval())
}

func QApplication_SetActiveWindow(active QWidget_ITF) {
	C.QApplication_QApplication_SetActiveWindow(PointerFromQWidget(active))
}

func (ptr *QApplication) SetAutoMaximizeThreshold(threshold int) {
	if ptr.Pointer() != nil {
		C.QApplication_SetAutoMaximizeThreshold(ptr.Pointer(), C.int(threshold))
	}
}

func (ptr *QApplication) SetAutoSipEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QApplication_SetAutoSipEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QApplication_SetCursorFlashTime(v int) {
	C.QApplication_QApplication_SetCursorFlashTime(C.int(v))
}

func QApplication_SetDoubleClickInterval(v int) {
	C.QApplication_QApplication_SetDoubleClickInterval(C.int(v))
}

func QApplication_SetEffectEnabled(effect core.Qt__UIEffect, enable bool) {
	C.QApplication_QApplication_SetEffectEnabled(C.int(effect), C.int(qt.GoBoolToInt(enable)))
}

func QApplication_SetGlobalStrut(v core.QSize_ITF) {
	C.QApplication_QApplication_SetGlobalStrut(core.PointerFromQSize(v))
}

func QApplication_SetKeyboardInputInterval(v int) {
	C.QApplication_QApplication_SetKeyboardInputInterval(C.int(v))
}

func QApplication_SetStartDragDistance(l int) {
	C.QApplication_QApplication_SetStartDragDistance(C.int(l))
}

func QApplication_SetStartDragTime(ms int) {
	C.QApplication_QApplication_SetStartDragTime(C.int(ms))
}

func (ptr *QApplication) SetStyleSheet(sheet string) {
	if ptr.Pointer() != nil {
		C.QApplication_SetStyleSheet(ptr.Pointer(), C.CString(sheet))
	}
}

func QApplication_SetWheelScrollLines(v int) {
	C.QApplication_QApplication_SetWheelScrollLines(C.int(v))
}

func QApplication_SetWindowIcon(icon gui.QIcon_ITF) {
	C.QApplication_QApplication_SetWindowIcon(gui.PointerFromQIcon(icon))
}

func QApplication_StartDragDistance() int {
	return int(C.QApplication_QApplication_StartDragDistance())
}

func QApplication_StartDragTime() int {
	return int(C.QApplication_QApplication_StartDragTime())
}

func (ptr *QApplication) StyleSheet() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QApplication_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func QApplication_TopLevelAt(point core.QPoint_ITF) *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_TopLevelAt(core.PointerFromQPoint(point)))
}

func QApplication_WheelScrollLines() int {
	return int(C.QApplication_QApplication_WheelScrollLines())
}

func QApplication_WidgetAt(point core.QPoint_ITF) *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_WidgetAt(core.PointerFromQPoint(point)))
}

func NewQApplication(argc int, argv []string) *QApplication {
	return NewQApplicationFromPointer(C.QApplication_NewQApplication(C.int(argc), C.CString(strings.Join(argv, "|"))))
}

func QApplication_AboutQt() {
	C.QApplication_QApplication_AboutQt()
}

func QApplication_ActiveModalWidget() *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_ActiveModalWidget())
}

func QApplication_ActivePopupWidget() *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_ActivePopupWidget())
}

func QApplication_ActiveWindow() *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_ActiveWindow())
}

func QApplication_CloseAllWindows() {
	C.QApplication_QApplication_CloseAllWindows()
}

func QApplication_ColorSpec() int {
	return int(C.QApplication_QApplication_ColorSpec())
}

func QApplication_Desktop() *QDesktopWidget {
	return NewQDesktopWidgetFromPointer(C.QApplication_QApplication_Desktop())
}

func QApplication_Exec() int {
	return int(C.QApplication_QApplication_Exec())
}

func (ptr *QApplication) ConnectFocusChanged(f func(old *QWidget, now *QWidget)) {
	if ptr.Pointer() != nil {
		C.QApplication_ConnectFocusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusChanged", f)
	}
}

func (ptr *QApplication) DisconnectFocusChanged() {
	if ptr.Pointer() != nil {
		C.QApplication_DisconnectFocusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusChanged")
	}
}

//export callbackQApplicationFocusChanged
func callbackQApplicationFocusChanged(ptrName *C.char, old unsafe.Pointer, now unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "focusChanged").(func(*QWidget, *QWidget))(NewQWidgetFromPointer(old), NewQWidgetFromPointer(now))
}

func QApplication_FocusWidget() *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_FocusWidget())
}

func (ptr *QApplication) Notify(receiver core.QObject_ITF, e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QApplication_Notify(ptr.Pointer(), core.PointerFromQObject(receiver), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func QApplication_SetColorSpec(spec int) {
	C.QApplication_QApplication_SetColorSpec(C.int(spec))
}

func QApplication_SetFont(font gui.QFont_ITF, className string) {
	C.QApplication_QApplication_SetFont(gui.PointerFromQFont(font), C.CString(className))
}

func QApplication_SetPalette(palette gui.QPalette_ITF, className string) {
	C.QApplication_QApplication_SetPalette(gui.PointerFromQPalette(palette), C.CString(className))
}

func QApplication_SetStyle2(style string) *QStyle {
	return NewQStyleFromPointer(C.QApplication_QApplication_SetStyle2(C.CString(style)))
}

func QApplication_SetStyle(style QStyle_ITF) {
	C.QApplication_QApplication_SetStyle(PointerFromQStyle(style))
}

func QApplication_Style() *QStyle {
	return NewQStyleFromPointer(C.QApplication_QApplication_Style())
}

func QApplication_TopLevelAt2(x int, y int) *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_TopLevelAt2(C.int(x), C.int(y)))
}

func QApplication_WidgetAt2(x int, y int) *QWidget {
	return NewQWidgetFromPointer(C.QApplication_QApplication_WidgetAt2(C.int(x), C.int(y)))
}

func (ptr *QApplication) DestroyQApplication() {
	if ptr.Pointer() != nil {
		C.QApplication_DestroyQApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
