package widgets

//#include "qpushbutton.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QPushButton struct {
	QAbstractButton
}

type QPushButtonITF interface {
	QAbstractButtonITF
	QPushButtonPTR() *QPushButton
}

func PointerFromQPushButton(ptr QPushButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPushButtonPTR().Pointer()
	}
	return nil
}

func QPushButtonFromPointer(ptr unsafe.Pointer) *QPushButton {
	var n = new(QPushButton)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPushButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPushButton) QPushButtonPTR() *QPushButton {
	return ptr
}

func (ptr *QPushButton) AutoDefault() bool {
	if ptr.Pointer() != nil {
		return C.QPushButton_AutoDefault(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPushButton) IsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QPushButton_IsDefault(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPushButton) IsFlat() bool {
	if ptr.Pointer() != nil {
		return C.QPushButton_IsFlat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPushButton) SetAutoDefault(v bool) {
	if ptr.Pointer() != nil {
		C.QPushButton_SetAutoDefault(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetDefault(v bool) {
	if ptr.Pointer() != nil {
		C.QPushButton_SetDefault(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetFlat(v bool) {
	if ptr.Pointer() != nil {
		C.QPushButton_SetFlat(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQPushButton(parent QWidgetITF) *QPushButton {
	return QPushButtonFromPointer(unsafe.Pointer(C.QPushButton_NewQPushButton(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQPushButton3(icon gui.QIconITF, text string, parent QWidgetITF) *QPushButton {
	return QPushButtonFromPointer(unsafe.Pointer(C.QPushButton_NewQPushButton3(C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQPushButton2(text string, parent QWidgetITF) *QPushButton {
	return QPushButtonFromPointer(unsafe.Pointer(C.QPushButton_NewQPushButton2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QPushButton) Menu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QPushButton_Menu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPushButton) SetMenu(menu QMenuITF) {
	if ptr.Pointer() != nil {
		C.QPushButton_SetMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(menu)))
	}
}

func (ptr *QPushButton) ShowMenu() {
	if ptr.Pointer() != nil {
		C.QPushButton_ShowMenu(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPushButton) DestroyQPushButton() {
	if ptr.Pointer() != nil {
		C.QPushButton_DestroyQPushButton(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
