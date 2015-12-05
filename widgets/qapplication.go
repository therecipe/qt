package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::alert")
		}
	}()

	C.QApplication_QApplication_Alert(PointerFromQWidget(widget), C.int(msec))
}

func (ptr *QApplication) AutoMaximizeThreshold() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::autoMaximizeThreshold")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QApplication_AutoMaximizeThreshold(ptr.Pointer()))
	}
	return 0
}

func (ptr *QApplication) AutoSipEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::autoSipEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QApplication_AutoSipEnabled(ptr.Pointer()) != 0
	}
	return false
}

func QApplication_Beep() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::beep")
		}
	}()

	C.QApplication_QApplication_Beep()
}

func QApplication_CursorFlashTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::cursorFlashTime")
		}
	}()

	return int(C.QApplication_QApplication_CursorFlashTime())
}

func QApplication_DoubleClickInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::doubleClickInterval")
		}
	}()

	return int(C.QApplication_QApplication_DoubleClickInterval())
}

func QApplication_IsEffectEnabled(effect core.Qt__UIEffect) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::isEffectEnabled")
		}
	}()

	return C.QApplication_QApplication_IsEffectEnabled(C.int(effect)) != 0
}

func QApplication_KeyboardInputInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::keyboardInputInterval")
		}
	}()

	return int(C.QApplication_QApplication_KeyboardInputInterval())
}

func QApplication_SetActiveWindow(active QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setActiveWindow")
		}
	}()

	C.QApplication_QApplication_SetActiveWindow(PointerFromQWidget(active))
}

func (ptr *QApplication) SetAutoMaximizeThreshold(threshold int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setAutoMaximizeThreshold")
		}
	}()

	if ptr.Pointer() != nil {
		C.QApplication_SetAutoMaximizeThreshold(ptr.Pointer(), C.int(threshold))
	}
}

func (ptr *QApplication) SetAutoSipEnabled(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setAutoSipEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QApplication_SetAutoSipEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QApplication_SetCursorFlashTime(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setCursorFlashTime")
		}
	}()

	C.QApplication_QApplication_SetCursorFlashTime(C.int(v))
}

func QApplication_SetDoubleClickInterval(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setDoubleClickInterval")
		}
	}()

	C.QApplication_QApplication_SetDoubleClickInterval(C.int(v))
}

func QApplication_SetEffectEnabled(effect core.Qt__UIEffect, enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setEffectEnabled")
		}
	}()

	C.QApplication_QApplication_SetEffectEnabled(C.int(effect), C.int(qt.GoBoolToInt(enable)))
}

func QApplication_SetGlobalStrut(v core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setGlobalStrut")
		}
	}()

	C.QApplication_QApplication_SetGlobalStrut(core.PointerFromQSize(v))
}

func QApplication_SetKeyboardInputInterval(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setKeyboardInputInterval")
		}
	}()

	C.QApplication_QApplication_SetKeyboardInputInterval(C.int(v))
}

func QApplication_SetStartDragDistance(l int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setStartDragDistance")
		}
	}()

	C.QApplication_QApplication_SetStartDragDistance(C.int(l))
}

func QApplication_SetStartDragTime(ms int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setStartDragTime")
		}
	}()

	C.QApplication_QApplication_SetStartDragTime(C.int(ms))
}

func (ptr *QApplication) SetStyleSheet(sheet string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setStyleSheet")
		}
	}()

	if ptr.Pointer() != nil {
		C.QApplication_SetStyleSheet(ptr.Pointer(), C.CString(sheet))
	}
}

func QApplication_SetWheelScrollLines(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setWheelScrollLines")
		}
	}()

	C.QApplication_QApplication_SetWheelScrollLines(C.int(v))
}

func QApplication_SetWindowIcon(icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setWindowIcon")
		}
	}()

	C.QApplication_QApplication_SetWindowIcon(gui.PointerFromQIcon(icon))
}

func QApplication_StartDragDistance() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::startDragDistance")
		}
	}()

	return int(C.QApplication_QApplication_StartDragDistance())
}

func QApplication_StartDragTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::startDragTime")
		}
	}()

	return int(C.QApplication_QApplication_StartDragTime())
}

func (ptr *QApplication) StyleSheet() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::styleSheet")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QApplication_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func QApplication_TopLevelAt(point core.QPoint_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::topLevelAt")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_TopLevelAt(core.PointerFromQPoint(point)))
}

func QApplication_WheelScrollLines() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::wheelScrollLines")
		}
	}()

	return int(C.QApplication_QApplication_WheelScrollLines())
}

func QApplication_WidgetAt(point core.QPoint_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::widgetAt")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_WidgetAt(core.PointerFromQPoint(point)))
}

func NewQApplication(argc int, argv []string) *QApplication {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::QApplication")
		}
	}()

	return NewQApplicationFromPointer(C.QApplication_NewQApplication(C.int(argc), C.CString(strings.Join(argv, ",,,"))))
}

func QApplication_AboutQt() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::aboutQt")
		}
	}()

	C.QApplication_QApplication_AboutQt()
}

func QApplication_ActiveModalWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::activeModalWidget")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_ActiveModalWidget())
}

func QApplication_ActivePopupWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::activePopupWidget")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_ActivePopupWidget())
}

func QApplication_ActiveWindow() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::activeWindow")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_ActiveWindow())
}

func QApplication_CloseAllWindows() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::closeAllWindows")
		}
	}()

	C.QApplication_QApplication_CloseAllWindows()
}

func QApplication_ColorSpec() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::colorSpec")
		}
	}()

	return int(C.QApplication_QApplication_ColorSpec())
}

func QApplication_Desktop() *QDesktopWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::desktop")
		}
	}()

	return NewQDesktopWidgetFromPointer(C.QApplication_QApplication_Desktop())
}

func QApplication_Exec() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::exec")
		}
	}()

	return int(C.QApplication_QApplication_Exec())
}

func (ptr *QApplication) ConnectFocusChanged(f func(old *QWidget, now *QWidget)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::focusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QApplication_ConnectFocusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusChanged", f)
	}
}

func (ptr *QApplication) DisconnectFocusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::focusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QApplication_DisconnectFocusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusChanged")
	}
}

//export callbackQApplicationFocusChanged
func callbackQApplicationFocusChanged(ptrName *C.char, old unsafe.Pointer, now unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::focusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "focusChanged").(func(*QWidget, *QWidget))(NewQWidgetFromPointer(old), NewQWidgetFromPointer(now))
}

func QApplication_FocusWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::focusWidget")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_FocusWidget())
}

func (ptr *QApplication) Notify(receiver core.QObject_ITF, e core.QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::notify")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QApplication_Notify(ptr.Pointer(), core.PointerFromQObject(receiver), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func QApplication_SetColorSpec(spec int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setColorSpec")
		}
	}()

	C.QApplication_QApplication_SetColorSpec(C.int(spec))
}

func QApplication_SetFont(font gui.QFont_ITF, className string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setFont")
		}
	}()

	C.QApplication_QApplication_SetFont(gui.PointerFromQFont(font), C.CString(className))
}

func QApplication_SetPalette(palette gui.QPalette_ITF, className string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setPalette")
		}
	}()

	C.QApplication_QApplication_SetPalette(gui.PointerFromQPalette(palette), C.CString(className))
}

func QApplication_SetStyle2(style string) *QStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setStyle")
		}
	}()

	return NewQStyleFromPointer(C.QApplication_QApplication_SetStyle2(C.CString(style)))
}

func QApplication_SetStyle(style QStyle_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::setStyle")
		}
	}()

	C.QApplication_QApplication_SetStyle(PointerFromQStyle(style))
}

func QApplication_Style() *QStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::style")
		}
	}()

	return NewQStyleFromPointer(C.QApplication_QApplication_Style())
}

func QApplication_TopLevelAt2(x int, y int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::topLevelAt")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_TopLevelAt2(C.int(x), C.int(y)))
}

func QApplication_WidgetAt2(x int, y int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::widgetAt")
		}
	}()

	return NewQWidgetFromPointer(C.QApplication_QApplication_WidgetAt2(C.int(x), C.int(y)))
}

func (ptr *QApplication) DestroyQApplication() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QApplication::~QApplication")
		}
	}()

	if ptr.Pointer() != nil {
		C.QApplication_DestroyQApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
