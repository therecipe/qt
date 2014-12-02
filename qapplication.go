package qt

//#include "qapplication.h"
import "C"

type qapplication struct {
	qguiapplication
}

type QApplication interface {
	QGuiApplication
	StyleSheet() string
	ConnectSlotAutoSipEnabled()
	DisconnectSlotAutoSipEnabled()
	SlotAutoSipEnabled()
}

func (p *qapplication) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qapplication) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQApplication_Int_String(argc int, argv string) QApplication {
	var qapplication = new(qapplication)
	qapplication.SetPointer(C.QApplication_New_Int_String(C.int(argc), C.CString(argv)))
	qapplication.SetObjectName_String("QApplication_" + randomIdentifier())
	return qapplication
}

func (p *qapplication) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QApplication_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qapplication) StyleSheet() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QApplication_StyleSheet(p.Pointer()))
	}
}

func (p *qapplication) ConnectSlotAutoSipEnabled() {
	C.QApplication_ConnectSlotAutoSipEnabled(p.Pointer())
}

func (p *qapplication) DisconnectSlotAutoSipEnabled() {
	C.QApplication_DisconnectSlotAutoSipEnabled(p.Pointer())
}

func (p *qapplication) SlotAutoSipEnabled() {
	if p.Pointer() != nil {
		C.QApplication_AutoSipEnabled(p.Pointer())
	}
}

func QApplication_ActiveModalWidget() QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QApplication_ActiveModalWidget())
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QApplication_ActivePopupWidget() QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QApplication_ActivePopupWidget())
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QApplication_ActiveWindow() QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QApplication_ActiveWindow())
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QApplication_Alert_QWidget_Int(widget QWidget, msec int) {
	var widgetPtr C.QtObjectPtr = nil
	if widget != nil {
		widgetPtr = widget.Pointer()
	}
	C.QApplication_Alert_QWidget_Int(widgetPtr, C.int(msec))
}

func QApplication_ColorSpec() int {
	return int(C.QApplication_ColorSpec())
}

func QApplication_CursorFlashTime() int {
	return int(C.QApplication_CursorFlashTime())
}

func QApplication_DoubleClickInterval() int {
	return int(C.QApplication_DoubleClickInterval())
}

func QApplication_Exec() int {
	return int(C.QApplication_Exec())
}

func QApplication_FocusWidget() QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QApplication_FocusWidget())
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QApplication_IsEffectEnabled_UIEffect(effect UIEffect) bool {
	return C.QApplication_IsEffectEnabled_UIEffect(C.int(effect)) != 0
}

func QApplication_KeyboardInputInterval() int {
	return int(C.QApplication_KeyboardInputInterval())
}

func QApplication_SetActiveWindow_QWidget(active QWidget) {
	var activePtr C.QtObjectPtr = nil
	if active != nil {
		activePtr = active.Pointer()
	}
	C.QApplication_SetActiveWindow_QWidget(activePtr)
}

func QApplication_SetColorSpec_Int(spec int) {
	C.QApplication_SetColorSpec_Int(C.int(spec))
}

func QApplication_SetCursorFlashTime_Int(flashTime int) {
	C.QApplication_SetCursorFlashTime_Int(C.int(flashTime))
}

func QApplication_SetDoubleClickInterval_Int(doubleClickInterval int) {
	C.QApplication_SetDoubleClickInterval_Int(C.int(doubleClickInterval))
}

func QApplication_SetEffectEnabled_UIEffect_Bool(effect UIEffect, enable bool) {
	C.QApplication_SetEffectEnabled_UIEffect_Bool(C.int(effect), goBoolToCInt(enable))
}

func QApplication_SetKeyboardInputInterval_Int(keyboardInputInterval int) {
	C.QApplication_SetKeyboardInputInterval_Int(C.int(keyboardInputInterval))
}

func QApplication_SetStartDragDistance_Int(l int) {
	C.QApplication_SetStartDragDistance_Int(C.int(l))
}

func QApplication_SetStartDragTime_Int(ms int) {
	C.QApplication_SetStartDragTime_Int(C.int(ms))
}

func QApplication_SetWheelScrollLines_Int(wheelScrollLines int) {
	C.QApplication_SetWheelScrollLines_Int(C.int(wheelScrollLines))
}

func QApplication_StartDragDistance() int {
	return int(C.QApplication_StartDragDistance())
}

func QApplication_StartDragTime() int {
	return int(C.QApplication_StartDragTime())
}

func QApplication_TopLevelAt_Int_Int(x int, y int) QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QApplication_TopLevelAt_Int_Int(C.int(x), C.int(y)))
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QApplication_WheelScrollLines() int {
	return int(C.QApplication_WheelScrollLines())
}

func QApplication_WidgetAt_Int_Int(x int, y int) QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QApplication_WidgetAt_Int_Int(C.int(x), C.int(y)))
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
	}
	return qwidget
}
