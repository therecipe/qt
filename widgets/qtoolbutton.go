package widgets

//#include "qtoolbutton.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QToolButton struct {
	QAbstractButton
}

type QToolButtonITF interface {
	QAbstractButtonITF
	QToolButtonPTR() *QToolButton
}

func PointerFromQToolButton(ptr QToolButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolButtonPTR().Pointer()
	}
	return nil
}

func QToolButtonFromPointer(ptr unsafe.Pointer) *QToolButton {
	var n = new(QToolButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QToolButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QToolButton) QToolButtonPTR() *QToolButton {
	return ptr
}

//QToolButton::ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int

var (
	QToolButton__DelayedPopup    = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    = QToolButton__ToolButtonPopupMode(2)
)

func (ptr *QToolButton) ArrowType() core.Qt__ArrowType {
	if ptr.Pointer() != nil {
		return core.Qt__ArrowType(C.QToolButton_ArrowType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QToolButton) AutoRaise() bool {
	if ptr.Pointer() != nil {
		return C.QToolButton_AutoRaise(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QToolButton) PopupMode() QToolButton__ToolButtonPopupMode {
	if ptr.Pointer() != nil {
		return QToolButton__ToolButtonPopupMode(C.QToolButton_PopupMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QToolButton) SetArrowType(ty core.Qt__ArrowType) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetArrowType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QToolButton) SetAutoRaise(enable bool) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetAutoRaise(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QToolButton) SetPopupMode(mode QToolButton__ToolButtonPopupMode) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetPopupMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QToolButton) SetToolButtonStyle(style core.Qt__ToolButtonStyle) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetToolButtonStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QToolButton) ToolButtonStyle() core.Qt__ToolButtonStyle {
	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolButton_ToolButtonStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQToolButton(parent QWidgetITF) *QToolButton {
	return QToolButtonFromPointer(unsafe.Pointer(C.QToolButton_NewQToolButton(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QToolButton) Menu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QToolButton_Menu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QToolButton) SetMenu(menu QMenuITF) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(menu)))
	}
}

func (ptr *QToolButton) ShowMenu() {
	if ptr.Pointer() != nil {
		C.QToolButton_ShowMenu(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QToolButton) ConnectTriggered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QToolButton_ConnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QToolButton) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QToolButton_DisconnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQToolButtonTriggered
func callbackQToolButtonTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QToolButton) DestroyQToolButton() {
	if ptr.Pointer() != nil {
		C.QToolButton_DestroyQToolButton(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
