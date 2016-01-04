package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSystemTrayIcon struct {
	core.QObject
}

type QSystemTrayIcon_ITF interface {
	core.QObject_ITF
	QSystemTrayIcon_PTR() *QSystemTrayIcon
}

func PointerFromQSystemTrayIcon(ptr QSystemTrayIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSystemTrayIcon_PTR().Pointer()
	}
	return nil
}

func NewQSystemTrayIconFromPointer(ptr unsafe.Pointer) *QSystemTrayIcon {
	var n = new(QSystemTrayIcon)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSystemTrayIcon_") {
		n.SetObjectName("QSystemTrayIcon_" + qt.Identifier())
	}
	return n
}

func (ptr *QSystemTrayIcon) QSystemTrayIcon_PTR() *QSystemTrayIcon {
	return ptr
}

//QSystemTrayIcon::ActivationReason
type QSystemTrayIcon__ActivationReason int64

const (
	QSystemTrayIcon__Unknown     = QSystemTrayIcon__ActivationReason(0)
	QSystemTrayIcon__Context     = QSystemTrayIcon__ActivationReason(1)
	QSystemTrayIcon__DoubleClick = QSystemTrayIcon__ActivationReason(2)
	QSystemTrayIcon__Trigger     = QSystemTrayIcon__ActivationReason(3)
	QSystemTrayIcon__MiddleClick = QSystemTrayIcon__ActivationReason(4)
)

//QSystemTrayIcon::MessageIcon
type QSystemTrayIcon__MessageIcon int64

const (
	QSystemTrayIcon__NoIcon      = QSystemTrayIcon__MessageIcon(0)
	QSystemTrayIcon__Information = QSystemTrayIcon__MessageIcon(1)
	QSystemTrayIcon__Warning     = QSystemTrayIcon__MessageIcon(2)
	QSystemTrayIcon__Critical    = QSystemTrayIcon__MessageIcon(3)
)

func (ptr *QSystemTrayIcon) Icon() *gui.QIcon {
	defer qt.Recovering("QSystemTrayIcon::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QSystemTrayIcon_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSystemTrayIcon) IsVisible() bool {
	defer qt.Recovering("QSystemTrayIcon::isVisible")

	if ptr.Pointer() != nil {
		return C.QSystemTrayIcon_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSystemTrayIcon) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QSystemTrayIcon::setIcon")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QSystemTrayIcon) SetToolTip(tip string) {
	defer qt.Recovering("QSystemTrayIcon::setToolTip")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetToolTip(ptr.Pointer(), C.CString(tip))
	}
}

func (ptr *QSystemTrayIcon) SetVisible(visible bool) {
	defer qt.Recovering("QSystemTrayIcon::setVisible")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSystemTrayIcon) ShowMessage(title string, message string, icon QSystemTrayIcon__MessageIcon, millisecondsTimeoutHint int) {
	defer qt.Recovering("QSystemTrayIcon::showMessage")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ShowMessage(ptr.Pointer(), C.CString(title), C.CString(message), C.int(icon), C.int(millisecondsTimeoutHint))
	}
}

func (ptr *QSystemTrayIcon) ToolTip() string {
	defer qt.Recovering("QSystemTrayIcon::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemTrayIcon_ToolTip(ptr.Pointer()))
	}
	return ""
}

func NewQSystemTrayIcon(parent core.QObject_ITF) *QSystemTrayIcon {
	defer qt.Recovering("QSystemTrayIcon::QSystemTrayIcon")

	return NewQSystemTrayIconFromPointer(C.QSystemTrayIcon_NewQSystemTrayIcon(core.PointerFromQObject(parent)))
}

func NewQSystemTrayIcon2(icon gui.QIcon_ITF, parent core.QObject_ITF) *QSystemTrayIcon {
	defer qt.Recovering("QSystemTrayIcon::QSystemTrayIcon")

	return NewQSystemTrayIconFromPointer(C.QSystemTrayIcon_NewQSystemTrayIcon2(gui.PointerFromQIcon(icon), core.PointerFromQObject(parent)))
}

func (ptr *QSystemTrayIcon) ConnectActivated(f func(reason QSystemTrayIcon__ActivationReason)) {
	defer qt.Recovering("connect QSystemTrayIcon::activated")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectActivated() {
	defer qt.Recovering("disconnect QSystemTrayIcon::activated")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQSystemTrayIconActivated
func callbackQSystemTrayIconActivated(ptr unsafe.Pointer, ptrName *C.char, reason C.int) {
	defer qt.Recovering("callback QSystemTrayIcon::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func(QSystemTrayIcon__ActivationReason))(QSystemTrayIcon__ActivationReason(reason))
	}

}

func (ptr *QSystemTrayIcon) Activated(reason QSystemTrayIcon__ActivationReason) {
	defer qt.Recovering("QSystemTrayIcon::activated")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Activated(ptr.Pointer(), C.int(reason))
	}
}

func (ptr *QSystemTrayIcon) ContextMenu() *QMenu {
	defer qt.Recovering("QSystemTrayIcon::contextMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QSystemTrayIcon_ContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSystemTrayIcon) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSystemTrayIcon::event")

	if ptr.Pointer() != nil {
		return C.QSystemTrayIcon_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSystemTrayIcon) Geometry() *core.QRect {
	defer qt.Recovering("QSystemTrayIcon::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSystemTrayIcon_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSystemTrayIcon) Hide() {
	defer qt.Recovering("QSystemTrayIcon::hide")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Hide(ptr.Pointer())
	}
}

func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	defer qt.Recovering("QSystemTrayIcon::isSystemTrayAvailable")

	return C.QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable() != 0
}

func (ptr *QSystemTrayIcon) ConnectMessageClicked(f func()) {
	defer qt.Recovering("connect QSystemTrayIcon::messageClicked")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ConnectMessageClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageClicked", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectMessageClicked() {
	defer qt.Recovering("disconnect QSystemTrayIcon::messageClicked")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DisconnectMessageClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageClicked")
	}
}

//export callbackQSystemTrayIconMessageClicked
func callbackQSystemTrayIconMessageClicked(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSystemTrayIcon::messageClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageClicked"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSystemTrayIcon) MessageClicked() {
	defer qt.Recovering("QSystemTrayIcon::messageClicked")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_MessageClicked(ptr.Pointer())
	}
}

func (ptr *QSystemTrayIcon) SetContextMenu(menu QMenu_ITF) {
	defer qt.Recovering("QSystemTrayIcon::setContextMenu")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetContextMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QSystemTrayIcon) Show() {
	defer qt.Recovering("QSystemTrayIcon::show")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Show(ptr.Pointer())
	}
}

func QSystemTrayIcon_SupportsMessages() bool {
	defer qt.Recovering("QSystemTrayIcon::supportsMessages")

	return C.QSystemTrayIcon_QSystemTrayIcon_SupportsMessages() != 0
}

func (ptr *QSystemTrayIcon) DestroyQSystemTrayIcon() {
	defer qt.Recovering("QSystemTrayIcon::~QSystemTrayIcon")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DestroyQSystemTrayIcon(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSystemTrayIcon) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSystemTrayIcon::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSystemTrayIcon::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSystemTrayIconTimerEvent
func callbackQSystemTrayIconTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSystemTrayIcon::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSystemTrayIconFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSystemTrayIcon) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSystemTrayIcon::timerEvent")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSystemTrayIcon) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSystemTrayIcon::timerEvent")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSystemTrayIcon) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSystemTrayIcon::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSystemTrayIcon::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSystemTrayIconChildEvent
func callbackQSystemTrayIconChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSystemTrayIcon::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSystemTrayIconFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSystemTrayIcon) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSystemTrayIcon::childEvent")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSystemTrayIcon) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSystemTrayIcon::childEvent")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSystemTrayIcon) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSystemTrayIcon::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSystemTrayIcon::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSystemTrayIconCustomEvent
func callbackQSystemTrayIconCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSystemTrayIcon::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSystemTrayIconFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSystemTrayIcon) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSystemTrayIcon::customEvent")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSystemTrayIcon) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSystemTrayIcon::customEvent")

	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
