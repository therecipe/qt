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

type QSystemTrayIconITF interface {
	core.QObjectITF
	QSystemTrayIconPTR() *QSystemTrayIcon
}

func PointerFromQSystemTrayIcon(ptr QSystemTrayIconITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSystemTrayIconPTR().Pointer()
	}
	return nil
}

func QSystemTrayIconFromPointer(ptr unsafe.Pointer) *QSystemTrayIcon {
	var n = new(QSystemTrayIcon)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSystemTrayIcon_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSystemTrayIcon) QSystemTrayIconPTR() *QSystemTrayIcon {
	return ptr
}

//QSystemTrayIcon::ActivationReason
type QSystemTrayIcon__ActivationReason int

var (
	QSystemTrayIcon__Unknown     = QSystemTrayIcon__ActivationReason(0)
	QSystemTrayIcon__Context     = QSystemTrayIcon__ActivationReason(1)
	QSystemTrayIcon__DoubleClick = QSystemTrayIcon__ActivationReason(2)
	QSystemTrayIcon__Trigger     = QSystemTrayIcon__ActivationReason(3)
	QSystemTrayIcon__MiddleClick = QSystemTrayIcon__ActivationReason(4)
)

//QSystemTrayIcon::MessageIcon
type QSystemTrayIcon__MessageIcon int

var (
	QSystemTrayIcon__NoIcon      = QSystemTrayIcon__MessageIcon(0)
	QSystemTrayIcon__Information = QSystemTrayIcon__MessageIcon(1)
	QSystemTrayIcon__Warning     = QSystemTrayIcon__MessageIcon(2)
	QSystemTrayIcon__Critical    = QSystemTrayIcon__MessageIcon(3)
)

func (ptr *QSystemTrayIcon) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QSystemTrayIcon_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSystemTrayIcon) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QSystemTrayIcon) SetToolTip(tip string) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(tip))
	}
}

func (ptr *QSystemTrayIcon) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSystemTrayIcon) ShowMessage(title string, message string, icon QSystemTrayIcon__MessageIcon, millisecondsTimeoutHint int) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ShowMessage(C.QtObjectPtr(ptr.Pointer()), C.CString(title), C.CString(message), C.int(icon), C.int(millisecondsTimeoutHint))
	}
}

func (ptr *QSystemTrayIcon) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemTrayIcon_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQSystemTrayIcon(parent core.QObjectITF) *QSystemTrayIcon {
	return QSystemTrayIconFromPointer(unsafe.Pointer(C.QSystemTrayIcon_NewQSystemTrayIcon(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQSystemTrayIcon2(icon gui.QIconITF, parent core.QObjectITF) *QSystemTrayIcon {
	return QSystemTrayIconFromPointer(unsafe.Pointer(C.QSystemTrayIcon_NewQSystemTrayIcon2(C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSystemTrayIcon) ConnectActivated(f func(reason QSystemTrayIcon__ActivationReason)) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQSystemTrayIconActivated
func callbackQSystemTrayIconActivated(ptrName *C.char, reason C.int) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(QSystemTrayIcon__ActivationReason))(QSystemTrayIcon__ActivationReason(reason))
}

func (ptr *QSystemTrayIcon) ContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QSystemTrayIcon_ContextMenu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSystemTrayIcon) Hide() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Hide(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	return C.QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable() != 0
}

func (ptr *QSystemTrayIcon) ConnectMessageClicked(f func()) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_ConnectMessageClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "messageClicked", f)
	}
}

func (ptr *QSystemTrayIcon) DisconnectMessageClicked() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DisconnectMessageClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "messageClicked")
	}
}

//export callbackQSystemTrayIconMessageClicked
func callbackQSystemTrayIconMessageClicked(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "messageClicked").(func())()
}

func (ptr *QSystemTrayIcon) SetContextMenu(menu QMenuITF) {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_SetContextMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(menu)))
	}
}

func (ptr *QSystemTrayIcon) Show() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_Show(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QSystemTrayIcon_SupportsMessages() bool {
	return C.QSystemTrayIcon_QSystemTrayIcon_SupportsMessages() != 0
}

func (ptr *QSystemTrayIcon) DestroyQSystemTrayIcon() {
	if ptr.Pointer() != nil {
		C.QSystemTrayIcon_DestroyQSystemTrayIcon(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
