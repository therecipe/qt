package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QMenuBar struct {
	QWidget
}

type QMenuBar_ITF interface {
	QWidget_ITF
	QMenuBar_PTR() *QMenuBar
}

func PointerFromQMenuBar(ptr QMenuBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuBar_PTR().Pointer()
	}
	return nil
}

func NewQMenuBarFromPointer(ptr unsafe.Pointer) *QMenuBar {
	var n = new(QMenuBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMenuBar_") {
		n.SetObjectName("QMenuBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMenuBar) QMenuBar_PTR() *QMenuBar {
	return ptr
}

func (ptr *QMenuBar) IsDefaultUp() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::isDefaultUp")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenuBar_IsDefaultUp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenuBar) IsNativeMenuBar() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::isNativeMenuBar")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenuBar_IsNativeMenuBar(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenuBar) SetCornerWidget(widget QWidget_ITF, corner core.Qt__Corner) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::setCornerWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(corner))
	}
}

func (ptr *QMenuBar) SetDefaultUp(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::setDefaultUp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_SetDefaultUp(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::setNativeMenuBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_SetNativeMenuBar(ptr.Pointer(), C.int(qt.GoBoolToInt(nativeMenuBar)))
	}
}

func NewQMenuBar(parent QWidget_ITF) *QMenuBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::QMenuBar")
		}
	}()

	return NewQMenuBarFromPointer(C.QMenuBar_NewQMenuBar(PointerFromQWidget(parent)))
}

func (ptr *QMenuBar) ActionAt(pt core.QPoint_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::actionAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_ActionAt(ptr.Pointer(), core.PointerFromQPoint(pt)))
	}
	return nil
}

func (ptr *QMenuBar) ActiveAction() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::activeAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_ActiveAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) AddAction(text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenuBar) AddAction2(text string, receiver core.QObject_ITF, member string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddAction2(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu(menu QMenu_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::addMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::addMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenuBar_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(title)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu2(title string) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::addMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenuBar_AddMenu2(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMenuBar) AddSeparator() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::addSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_Clear(ptr.Pointer())
	}
}

func (ptr *QMenuBar) CornerWidget(corner core.Qt__Corner) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::cornerWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMenuBar_CornerWidget(ptr.Pointer(), C.int(corner)))
	}
	return nil
}

func (ptr *QMenuBar) HeightForWidth(v int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::heightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMenuBar_HeightForWidth(ptr.Pointer(), C.int(v)))
	}
	return 0
}

func (ptr *QMenuBar) ConnectHovered(f func(action *QAction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::hovered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenuBar) DisconnectHovered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::hovered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuBarHovered
func callbackQMenuBarHovered(ptrName *C.char, action unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::hovered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QMenuBar) InsertMenu(before QAction_ITF, menu QMenu_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::insertMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_InsertMenu(ptr.Pointer(), PointerFromQAction(before), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenuBar) InsertSeparator(before QAction_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::insertSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QMenuBar) SetActiveAction(act QAction_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::setActiveAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_SetActiveAction(ptr.Pointer(), PointerFromQAction(act))
	}
}

func (ptr *QMenuBar) SetVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenuBar) ConnectTriggered(f func(action *QAction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::triggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenuBar) DisconnectTriggered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::triggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuBarTriggered
func callbackQMenuBarTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::triggered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QMenuBar) DestroyQMenuBar() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenuBar::~QMenuBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenuBar_DestroyQMenuBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
