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
	if n.ObjectName() == "" {
		n.SetObjectName("QMenu_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMenu) QMenu_PTR() *QMenu {
	return ptr
}

func (ptr *QMenu) IsTearOffEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) SeparatorsCollapsible() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_SeparatorsCollapsible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMenu) SetSeparatorsCollapsible(collapse bool) {
	if ptr.Pointer() != nil {
		C.QMenu_SetSeparatorsCollapsible(ptr.Pointer(), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QMenu) SetTearOffEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QMenu_SetTearOffEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenu) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QMenu_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QMenu) SetToolTipsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QMenu_SetToolTipsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenu) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMenu_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMenu) ToolTipsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_ToolTipsVisible(ptr.Pointer()) != 0
	}
	return false
}

func NewQMenu(parent QWidget_ITF) *QMenu {
	return NewQMenuFromPointer(C.QMenu_NewQMenu(PointerFromQWidget(parent)))
}

func NewQMenu2(title string, parent QWidget_ITF) *QMenu {
	return NewQMenuFromPointer(C.QMenu_NewQMenu2(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QMenu) ConnectAboutToHide(f func()) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToHide(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToHide", f)
	}
}

func (ptr *QMenu) DisconnectAboutToHide() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToHide(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToHide")
	}
}

//export callbackQMenuAboutToHide
func callbackQMenuAboutToHide(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToHide").(func())()
}

func (ptr *QMenu) ConnectAboutToShow(f func()) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToShow(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToShow", f)
	}
}

func (ptr *QMenu) DisconnectAboutToShow() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToShow(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToShow")
	}
}

//export callbackQMenuAboutToShow
func callbackQMenuAboutToShow(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToShow").(func())()
}

func (ptr *QMenu) ActionAt(pt core.QPoint_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_ActionAt(ptr.Pointer(), core.PointerFromQPoint(pt)))
	}
	return nil
}

func (ptr *QMenu) ActiveAction() *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_ActiveAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(receiver), C.CString(member), gui.PointerFromQKeySequence(shortcut)))
	}
	return nil
}

func (ptr *QMenu) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddAction3(text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction3(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member), gui.PointerFromQKeySequence(shortcut)))
	}
	return nil
}

func (ptr *QMenu) AddMenu(menu QMenu_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenu) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenu_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(title)))
	}
	return nil
}

func (ptr *QMenu) AddMenu2(title string) *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenu_AddMenu2(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMenu) AddSection2(icon gui.QIcon_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSection2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddSection(text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSection(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Clear() {
	if ptr.Pointer() != nil {
		C.QMenu_Clear(ptr.Pointer())
	}
}

func (ptr *QMenu) Exec() *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_Exec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Exec2(p core.QPoint_ITF, action QAction_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_Exec2(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QMenu) HideTearOffMenu() {
	if ptr.Pointer() != nil {
		C.QMenu_HideTearOffMenu(ptr.Pointer())
	}
}

func (ptr *QMenu) ConnectHovered(f func(action *QAction)) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenu) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuHovered
func callbackQMenuHovered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QMenu) InsertMenu(before QAction_ITF, menu QMenu_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertMenu(ptr.Pointer(), PointerFromQAction(before), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenu) InsertSection2(before QAction_ITF, icon gui.QIcon_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSection2(ptr.Pointer(), PointerFromQAction(before), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) InsertSection(before QAction_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSection(ptr.Pointer(), PointerFromQAction(before), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) InsertSeparator(before QAction_ITF) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QMenu) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) IsTearOffMenuVisible() bool {
	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffMenuVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) MenuAction() *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_MenuAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Popup(p core.QPoint_ITF, atAction QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu_Popup(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(atAction))
	}
}

func (ptr *QMenu) SetActiveAction(act QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu_SetActiveAction(ptr.Pointer(), PointerFromQAction(act))
	}
}

func (ptr *QMenu) ConnectTriggered(f func(action *QAction)) {
	if ptr.Pointer() != nil {
		C.QMenu_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenu) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuTriggered
func callbackQMenuTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QMenu) DestroyQMenu() {
	if ptr.Pointer() != nil {
		C.QMenu_DestroyQMenu(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
