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

type QToolButton_ITF interface {
	QAbstractButton_ITF
	QToolButton_PTR() *QToolButton
}

func PointerFromQToolButton(ptr QToolButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolButton_PTR().Pointer()
	}
	return nil
}

func NewQToolButtonFromPointer(ptr unsafe.Pointer) *QToolButton {
	var n = new(QToolButton)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QToolButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QToolButton) QToolButton_PTR() *QToolButton {
	return ptr
}

//QToolButton::ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int64

const (
	QToolButton__DelayedPopup    = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    = QToolButton__ToolButtonPopupMode(2)
)

func (ptr *QToolButton) ArrowType() core.Qt__ArrowType {
	if ptr.Pointer() != nil {
		return core.Qt__ArrowType(C.QToolButton_ArrowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) AutoRaise() bool {
	if ptr.Pointer() != nil {
		return C.QToolButton_AutoRaise(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolButton) PopupMode() QToolButton__ToolButtonPopupMode {
	if ptr.Pointer() != nil {
		return QToolButton__ToolButtonPopupMode(C.QToolButton_PopupMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) SetArrowType(ty core.Qt__ArrowType) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetArrowType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QToolButton) SetAutoRaise(enable bool) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetAutoRaise(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QToolButton) SetPopupMode(mode QToolButton__ToolButtonPopupMode) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetPopupMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QToolButton) SetToolButtonStyle(style core.Qt__ToolButtonStyle) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetToolButtonStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QToolButton) ToolButtonStyle() core.Qt__ToolButtonStyle {
	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolButton_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func NewQToolButton(parent QWidget_ITF) *QToolButton {
	return NewQToolButtonFromPointer(C.QToolButton_NewQToolButton(PointerFromQWidget(parent)))
}

func (ptr *QToolButton) Menu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QToolButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) SetMenu(menu QMenu_ITF) {
	if ptr.Pointer() != nil {
		C.QToolButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QToolButton) ShowMenu() {
	if ptr.Pointer() != nil {
		C.QToolButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QToolButton) ConnectTriggered(f func(action *QAction)) {
	if ptr.Pointer() != nil {
		C.QToolButton_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QToolButton) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QToolButton_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQToolButtonTriggered
func callbackQToolButtonTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QToolButton) DestroyQToolButton() {
	if ptr.Pointer() != nil {
		C.QToolButton_DestroyQToolButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
