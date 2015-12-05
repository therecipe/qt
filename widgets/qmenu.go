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

type QMenu struct {
	QWidget
}

type QMenu_ITF interface {
	QWidget_ITF
	QMenu_PTR() *QMenu
}

func PointerFromQMenu(ptr QMenu_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenu_PTR().Pointer()
	}
	return nil
}

func NewQMenuFromPointer(ptr unsafe.Pointer) *QMenu {
	var n = new(QMenu)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMenu_") {
		n.SetObjectName("QMenu_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMenu) QMenu_PTR() *QMenu {
	return ptr
}

func (ptr *QMenu) IsTearOffEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::isTearOffEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) SeparatorsCollapsible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::separatorsCollapsible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenu_SeparatorsCollapsible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) SetIcon(icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::setIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMenu) SetSeparatorsCollapsible(collapse bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::setSeparatorsCollapsible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_SetSeparatorsCollapsible(ptr.Pointer(), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QMenu) SetTearOffEnabled(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::setTearOffEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_SetTearOffEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenu) SetTitle(title string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::setTitle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QMenu) SetToolTipsVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::setToolTipsVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_SetToolTipsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenu) Title() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::title")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMenu_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMenu) ToolTipsVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::toolTipsVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenu_ToolTipsVisible(ptr.Pointer()) != 0
	}
	return false
}

func NewQMenu(parent QWidget_ITF) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::QMenu")
		}
	}()

	return NewQMenuFromPointer(C.QMenu_NewQMenu(PointerFromQWidget(parent)))
}

func NewQMenu2(title string, parent QWidget_ITF) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::QMenu")
		}
	}()

	return NewQMenuFromPointer(C.QMenu_NewQMenu2(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QMenu) ConnectAboutToHide(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::aboutToHide")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToHide(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToHide", f)
	}
}

func (ptr *QMenu) DisconnectAboutToHide() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::aboutToHide")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToHide(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToHide")
	}
}

//export callbackQMenuAboutToHide
func callbackQMenuAboutToHide(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::aboutToHide")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "aboutToHide").(func())()
}

func (ptr *QMenu) ConnectAboutToShow(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::aboutToShow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToShow(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToShow", f)
	}
}

func (ptr *QMenu) DisconnectAboutToShow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::aboutToShow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToShow(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToShow")
	}
}

//export callbackQMenuAboutToShow
func callbackQMenuAboutToShow(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::aboutToShow")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "aboutToShow").(func())()
}

func (ptr *QMenu) ActionAt(pt core.QPoint_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::actionAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_ActionAt(ptr.Pointer(), core.PointerFromQPoint(pt)))
	}
	return nil
}

func (ptr *QMenu) ActiveAction() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::activeAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_ActiveAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(receiver), C.CString(member), gui.PointerFromQKeySequence(shortcut)))
	}
	return nil
}

func (ptr *QMenu) AddAction(text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddAction3(text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction3(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member), gui.PointerFromQKeySequence(shortcut)))
	}
	return nil
}

func (ptr *QMenu) AddMenu(menu QMenu_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenu) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenu_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(title)))
	}
	return nil
}

func (ptr *QMenu) AddMenu2(title string) *QMenu {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenu_AddMenu2(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMenu) AddSection2(icon gui.QIcon_ITF, text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addSection")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSection2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddSection(text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addSection")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSection(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddSeparator() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::addSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_Clear(ptr.Pointer())
	}
}

func (ptr *QMenu) Exec() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::exec")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_Exec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Exec2(p core.QPoint_ITF, action QAction_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::exec")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_Exec2(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QMenu) HideTearOffMenu() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::hideTearOffMenu")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_HideTearOffMenu(ptr.Pointer())
	}
}

func (ptr *QMenu) ConnectHovered(f func(action *QAction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::hovered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenu) DisconnectHovered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::hovered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuHovered
func callbackQMenuHovered(ptrName *C.char, action unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::hovered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QMenu) InsertMenu(before QAction_ITF, menu QMenu_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::insertMenu")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertMenu(ptr.Pointer(), PointerFromQAction(before), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenu) InsertSection2(before QAction_ITF, icon gui.QIcon_ITF, text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::insertSection")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSection2(ptr.Pointer(), PointerFromQAction(before), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) InsertSection(before QAction_ITF, text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::insertSection")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSection(ptr.Pointer(), PointerFromQAction(before), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) InsertSeparator(before QAction_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::insertSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QMenu) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenu_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) IsTearOffMenuVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::isTearOffMenuVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffMenuVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) MenuAction() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::menuAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_MenuAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Popup(p core.QPoint_ITF, atAction QAction_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::popup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_Popup(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(atAction))
	}
}

func (ptr *QMenu) SetActiveAction(act QAction_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::setActiveAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_SetActiveAction(ptr.Pointer(), PointerFromQAction(act))
	}
}

func (ptr *QMenu) ConnectTriggered(f func(action *QAction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::triggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenu) DisconnectTriggered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::triggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuTriggered
func callbackQMenuTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::triggered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QMenu) DestroyQMenu() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMenu::~QMenu")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMenu_DestroyQMenu(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
