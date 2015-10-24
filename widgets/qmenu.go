package widgets

//#include "qmenu.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMenu struct {
	QWidget
}

type QMenuITF interface {
	QWidgetITF
	QMenuPTR() *QMenu
}

func PointerFromQMenu(ptr QMenuITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuPTR().Pointer()
	}
	return nil
}

func QMenuFromPointer(ptr unsafe.Pointer) *QMenu {
	var n = new(QMenu)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMenu_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMenu) QMenuPTR() *QMenu {
	return ptr
}

func (ptr *QMenu) IsTearOffEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenu) SeparatorsCollapsible() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_SeparatorsCollapsible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenu) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QMenu_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QMenu) SetSeparatorsCollapsible(collapse bool) {
	if ptr.Pointer() != nil {
		C.QMenu_SetSeparatorsCollapsible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QMenu) SetTearOffEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QMenu_SetTearOffEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenu) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QMenu_SetTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QMenu) SetToolTipsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QMenu_SetToolTipsVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenu) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMenu_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMenu) ToolTipsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_ToolTipsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQMenu(parent QWidgetITF) *QMenu {
	return QMenuFromPointer(unsafe.Pointer(C.QMenu_NewQMenu(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQMenu2(title string, parent QWidgetITF) *QMenu {
	return QMenuFromPointer(unsafe.Pointer(C.QMenu_NewQMenu2(C.CString(title), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QMenu) ConnectAboutToHide(f func()) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToHide(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToHide", f)
	}
}

func (ptr *QMenu) DisconnectAboutToHide() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToHide(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToHide")
	}
}

//export callbackQMenuAboutToHide
func callbackQMenuAboutToHide(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToHide").(func())()
}

func (ptr *QMenu) ConnectAboutToShow(f func()) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToShow(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToShow", f)
	}
}

func (ptr *QMenu) DisconnectAboutToShow() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToShow(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToShow")
	}
}

//export callbackQMenuAboutToShow
func callbackQMenuAboutToShow(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToShow").(func())()
}

func (ptr *QMenu) ActionAt(pt core.QPointITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_ActionAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pt)))))
	}
	return nil
}

func (ptr *QMenu) ActiveAction() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_ActiveAction(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMenu) AddAction2(icon gui.QIconITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddAction2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QMenu) AddAction4(icon gui.QIconITF, text string, receiver core.QObjectITF, member string, shortcut gui.QKeySequenceITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddAction4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member), C.QtObjectPtr(gui.PointerFromQKeySequence(shortcut)))))
	}
	return nil
}

func (ptr *QMenu) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddAction(C.QtObjectPtr(ptr.Pointer()), C.CString(text))))
	}
	return nil
}

func (ptr *QMenu) AddAction3(text string, receiver core.QObjectITF, member string, shortcut gui.QKeySequenceITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddAction3(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member), C.QtObjectPtr(gui.PointerFromQKeySequence(shortcut)))))
	}
	return nil
}

func (ptr *QMenu) AddMenu(menu QMenuITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(menu)))))
	}
	return nil
}

func (ptr *QMenu) AddMenu3(icon gui.QIconITF, title string) *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QMenu_AddMenu3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(title))))
	}
	return nil
}

func (ptr *QMenu) AddMenu2(title string) *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QMenu_AddMenu2(C.QtObjectPtr(ptr.Pointer()), C.CString(title))))
	}
	return nil
}

func (ptr *QMenu) AddSection2(icon gui.QIconITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddSection2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QMenu) AddSection(text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddSection(C.QtObjectPtr(ptr.Pointer()), C.CString(text))))
	}
	return nil
}

func (ptr *QMenu) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_AddSeparator(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMenu) Clear() {
	if ptr.Pointer() != nil {
		C.QMenu_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMenu) Exec() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_Exec(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMenu) Exec2(p core.QPointITF, action QActionITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_Exec2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)), C.QtObjectPtr(PointerFromQAction(action)))))
	}
	return nil
}

func (ptr *QMenu) HideTearOffMenu() {
	if ptr.Pointer() != nil {
		C.QMenu_HideTearOffMenu(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMenu) ConnectHovered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenu) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuHovered
func callbackQMenuHovered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QMenu) InsertMenu(before QActionITF, menu QMenuITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_InsertMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.QtObjectPtr(PointerFromQMenu(menu)))))
	}
	return nil
}

func (ptr *QMenu) InsertSection2(before QActionITF, icon gui.QIconITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_InsertSection2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QMenu) InsertSection(before QActionITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_InsertSection(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.CString(text))))
	}
	return nil
}

func (ptr *QMenu) InsertSeparator(before QActionITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_InsertSeparator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)))))
	}
	return nil
}

func (ptr *QMenu) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenu) IsTearOffMenuVisible() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffMenuVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenu) MenuAction() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenu_MenuAction(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMenu) Popup(p core.QPointITF, atAction QActionITF) {
	if ptr.Pointer() != nil {
		C.QMenu_Popup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)), C.QtObjectPtr(PointerFromQAction(atAction)))
	}
}

func (ptr *QMenu) SetActiveAction(act QActionITF) {
	if ptr.Pointer() != nil {
		C.QMenu_SetActiveAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(act)))
	}
}

func (ptr *QMenu) ConnectTriggered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenu) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuTriggered
func callbackQMenuTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QMenu) DestroyQMenu() {
	if ptr.Pointer() != nil {
		C.QMenu_DestroyQMenu(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
