package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QPushButton struct {
	QAbstractButton
}

type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func PointerFromQPushButton(ptr QPushButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPushButton_PTR().Pointer()
	}
	return nil
}

func NewQPushButtonFromPointer(ptr unsafe.Pointer) *QPushButton {
	var n = new(QPushButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPushButton_") {
		n.SetObjectName("QPushButton_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton {
	return ptr
}

func (ptr *QPushButton) AutoDefault() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::autoDefault")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPushButton_AutoDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) IsDefault() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::isDefault")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPushButton_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) IsFlat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::isFlat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPushButton_IsFlat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) SetAutoDefault(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::setAutoDefault")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPushButton_SetAutoDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetDefault(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::setDefault")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPushButton_SetDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetFlat(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::setFlat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPushButton_SetFlat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQPushButton(parent QWidget_ITF) *QPushButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::QPushButton")
		}
	}()

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton(PointerFromQWidget(parent)))
}

func NewQPushButton3(icon gui.QIcon_ITF, text string, parent QWidget_ITF) *QPushButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::QPushButton")
		}
	}()

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton3(gui.PointerFromQIcon(icon), C.CString(text), PointerFromQWidget(parent)))
}

func NewQPushButton2(text string, parent QWidget_ITF) *QPushButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::QPushButton")
		}
	}()

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QPushButton) Menu() *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::menu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPushButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) SetMenu(menu QMenu_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::setMenu")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPushButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QPushButton) ShowMenu() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::showMenu")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPushButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QPushButton) DestroyQPushButton() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPushButton::~QPushButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPushButton_DestroyQPushButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
