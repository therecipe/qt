package widgets

//#include "qapplication.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QApplication struct {
	gui.QGuiApplication
}

type QApplicationITF interface {
	gui.QGuiApplicationITF
	QApplicationPTR() *QApplication
}

func PointerFromQApplication(ptr QApplicationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplicationPTR().Pointer()
	}
	return nil
}

func QApplicationFromPointer(ptr unsafe.Pointer) *QApplication {
	var n = new(QApplication)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QApplication_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QApplication) QApplicationPTR() *QApplication {
	return ptr
}

//QApplication::ColorSpec
type QApplication__ColorSpec int

var (
	QApplication__NormalColor = QApplication__ColorSpec(0)
	QApplication__CustomColor = QApplication__ColorSpec(1)
	QApplication__ManyColor   = QApplication__ColorSpec(2)
)

func QApplication_Alert(widget QWidgetITF, msec int) {
	C.QApplication_QApplication_Alert(C.QtObjectPtr(PointerFromQWidget(widget)), C.int(msec))
}

func (ptr *QApplication) AutoMaximizeThreshold() int {
	if ptr.Pointer() != nil {
		return int(C.QApplication_AutoMaximizeThreshold(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QApplication) AutoSipEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QApplication_AutoSipEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
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

func QApplication_SetActiveWindow(active QWidgetITF) {
	C.QApplication_QApplication_SetActiveWindow(C.QtObjectPtr(PointerFromQWidget(active)))
}

func (ptr *QApplication) SetAutoMaximizeThreshold(threshold int) {
	if ptr.Pointer() != nil {
		C.QApplication_SetAutoMaximizeThreshold(C.QtObjectPtr(ptr.Pointer()), C.int(threshold))
	}
}

func (ptr *QApplication) SetAutoSipEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QApplication_SetAutoSipEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
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

func QApplication_SetGlobalStrut(v core.QSizeITF) {
	C.QApplication_QApplication_SetGlobalStrut(C.QtObjectPtr(core.PointerFromQSize(v)))
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
		C.QApplication_SetStyleSheet(C.QtObjectPtr(ptr.Pointer()), C.CString(sheet))
	}
}

func QApplication_SetWheelScrollLines(v int) {
	C.QApplication_QApplication_SetWheelScrollLines(C.int(v))
}

func QApplication_SetWindowIcon(icon gui.QIconITF) {
	C.QApplication_QApplication_SetWindowIcon(C.QtObjectPtr(gui.PointerFromQIcon(icon)))
}

func QApplication_StartDragDistance() int {
	return int(C.QApplication_QApplication_StartDragDistance())
}

func QApplication_StartDragTime() int {
	return int(C.QApplication_QApplication_StartDragTime())
}

func (ptr *QApplication) StyleSheet() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QApplication_StyleSheet(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QApplication_TopLevelAt(point core.QPointITF) *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_TopLevelAt(C.QtObjectPtr(core.PointerFromQPoint(point)))))
}

func QApplication_WheelScrollLines() int {
	return int(C.QApplication_QApplication_WheelScrollLines())
}

func QApplication_WidgetAt(point core.QPointITF) *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_WidgetAt(C.QtObjectPtr(core.PointerFromQPoint(point)))))
}

func NewQApplication(argc int, argv string) *QApplication {
	return QApplicationFromPointer(unsafe.Pointer(C.QApplication_NewQApplication(C.int(argc), C.CString(argv))))
}

func QApplication_AboutQt() {
	C.QApplication_QApplication_AboutQt()
}

func QApplication_ActiveModalWidget() *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_ActiveModalWidget()))
}

func QApplication_ActivePopupWidget() *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_ActivePopupWidget()))
}

func QApplication_ActiveWindow() *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_ActiveWindow()))
}

func QApplication_CloseAllWindows() {
	C.QApplication_QApplication_CloseAllWindows()
}

func QApplication_ColorSpec() int {
	return int(C.QApplication_QApplication_ColorSpec())
}

func QApplication_Desktop() *QDesktopWidget {
	return QDesktopWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_Desktop()))
}

func QApplication_Exec() int {
	return int(C.QApplication_QApplication_Exec())
}

func (ptr *QApplication) ConnectFocusChanged(f func(old QWidgetITF, now QWidgetITF)) {
	if ptr.Pointer() != nil {
		C.QApplication_ConnectFocusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusChanged", f)
	}
}

func (ptr *QApplication) DisconnectFocusChanged() {
	if ptr.Pointer() != nil {
		C.QApplication_DisconnectFocusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusChanged")
	}
}

//export callbackQApplicationFocusChanged
func callbackQApplicationFocusChanged(ptrName *C.char, old unsafe.Pointer, now unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "focusChanged").(func(*QWidget, *QWidget))(QWidgetFromPointer(old), QWidgetFromPointer(now))
}

func QApplication_FocusWidget() *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_FocusWidget()))
}

func (ptr *QApplication) Notify(receiver core.QObjectITF, e core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QApplication_Notify(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.QtObjectPtr(core.PointerFromQEvent(e))) != 0
	}
	return false
}

func QApplication_SetColorSpec(spec int) {
	C.QApplication_QApplication_SetColorSpec(C.int(spec))
}

func QApplication_SetFont(font gui.QFontITF, className string) {
	C.QApplication_QApplication_SetFont(C.QtObjectPtr(gui.PointerFromQFont(font)), C.CString(className))
}

func QApplication_SetPalette(palette gui.QPaletteITF, className string) {
	C.QApplication_QApplication_SetPalette(C.QtObjectPtr(gui.PointerFromQPalette(palette)), C.CString(className))
}

func QApplication_SetStyle2(style string) *QStyle {
	return QStyleFromPointer(unsafe.Pointer(C.QApplication_QApplication_SetStyle2(C.CString(style))))
}

func QApplication_SetStyle(style QStyleITF) {
	C.QApplication_QApplication_SetStyle(C.QtObjectPtr(PointerFromQStyle(style)))
}

func QApplication_Style() *QStyle {
	return QStyleFromPointer(unsafe.Pointer(C.QApplication_QApplication_Style()))
}

func QApplication_TopLevelAt2(x int, y int) *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_TopLevelAt2(C.int(x), C.int(y))))
}

func QApplication_WidgetAt2(x int, y int) *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QApplication_QApplication_WidgetAt2(C.int(x), C.int(y))))
}

func (ptr *QApplication) DestroyQApplication() {
	if ptr.Pointer() != nil {
		C.QApplication_DestroyQApplication(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
