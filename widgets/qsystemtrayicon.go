package widgets

//#include "qsystemtrayicon.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QSystemTrayIcon_" + qt.RandomIdentifier())
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

func (ptr *QSystemTrayIcon) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QSystemTrayIcon_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSystemTrayIcon) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QSystemTrayIcon) SetToolTip(tip string) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetToolTip(ptr.Pointer(), C.CString(tip))
	}
}

func (ptr *QSystemTrayIcon) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSystemTrayIcon) ShowMessage(title string, message string, icon QSystemTrayIcon__MessageIcon, millisecondsTimeoutHint int) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ShowMessage(ptr.Pointer(), C.CString(title), C.CString(message), C.int(icon), C.int(millisecondsTimeoutHint))
	}
}

func (ptr *QSystemTrayIcon) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemTrayIcon_ToolTip(ptr.Pointer()))
	}
	return ""
}

func NewQSystemTrayIcon(parent core.QObject_ITF) *QSystemTrayIcon {
	return NewQSystemTrayIconFromPointer(C.QSystemTrayIcon_NewQSystemTrayIcon(core.PointerFromQObject(parent)))
}

func NewQSystemTrayIcon2(icon gui.QIcon_ITF, parent core.QObject_ITF) *QSystemTrayIcon {
	return NewQSystemTrayIconFromPointer(C.QSystemTrayIcon_NewQSystemTrayIcon2(gui.PointerFromQIcon(icon), core.PointerFromQObject(parent)))
}

func (ptr *QSystemTrayIcon) ConnectActivated(f func(reason QSystemTrayIcon__ActivationReason)) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQSystemTrayIconActivated
func callbackQSystemTrayIconActivated(ptrName *C.char, reason C.int) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(QSystemTrayIcon__ActivationReason))(QSystemTrayIcon__ActivationReason(reason))
}

func (ptr *QSystemTrayIcon) ContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QSystemTrayIcon_ContextMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSystemTrayIcon) Hide() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Hide(ptr.Pointer())
	}
}

func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	return C.QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable() != 0
}

func (ptr *QSystemTrayIcon) ConnectMessageClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ConnectMessageClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageClicked", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectMessageClicked() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DisconnectMessageClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageClicked")
	}
}

//export callbackQSystemTrayIconMessageClicked
func callbackQSystemTrayIconMessageClicked(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "messageClicked").(func())()
}

func (ptr *QSystemTrayIcon) SetContextMenu(menu QMenu_ITF) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetContextMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QSystemTrayIcon) Show() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Show(ptr.Pointer())
	}
}

func QSystemTrayIcon_SupportsMessages() bool {
	return C.QSystemTrayIcon_QSystemTrayIcon_SupportsMessages() != 0
}

func (ptr *QSystemTrayIcon) DestroyQSystemTrayIcon() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DestroyQSystemTrayIcon(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
