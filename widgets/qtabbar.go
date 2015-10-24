package widgets

//#include "qtabbar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QTabBar struct {
	QWidget
}

type QTabBarITF interface {
	QWidgetITF
	QTabBarPTR() *QTabBar
}

func PointerFromQTabBar(ptr QTabBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabBarPTR().Pointer()
	}
	return nil
}

func QTabBarFromPointer(ptr unsafe.Pointer) *QTabBar {
	var n = new(QTabBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTabBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTabBar) QTabBarPTR() *QTabBar {
	return ptr
}

//QTabBar::ButtonPosition
type QTabBar__ButtonPosition int

var (
	QTabBar__LeftSide  = QTabBar__ButtonPosition(0)
	QTabBar__RightSide = QTabBar__ButtonPosition(1)
)

//QTabBar::SelectionBehavior
type QTabBar__SelectionBehavior int

var (
	QTabBar__SelectLeftTab     = QTabBar__SelectionBehavior(0)
	QTabBar__SelectRightTab    = QTabBar__SelectionBehavior(1)
	QTabBar__SelectPreviousTab = QTabBar__SelectionBehavior(2)
)

//QTabBar::Shape
type QTabBar__Shape int

var (
	QTabBar__RoundedNorth    = QTabBar__Shape(0)
	QTabBar__RoundedSouth    = QTabBar__Shape(1)
	QTabBar__RoundedWest     = QTabBar__Shape(2)
	QTabBar__RoundedEast     = QTabBar__Shape(3)
	QTabBar__TriangularNorth = QTabBar__Shape(4)
	QTabBar__TriangularSouth = QTabBar__Shape(5)
	QTabBar__TriangularWest  = QTabBar__Shape(6)
	QTabBar__TriangularEast  = QTabBar__Shape(7)
)

func (ptr *QTabBar) AutoHide() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_AutoHide(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) ChangeCurrentOnDrag() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_ChangeCurrentOnDrag(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabBar) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabBar) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_DocumentMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) DrawBase() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_DrawBase(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) ElideMode() core.Qt__TextElideMode {
	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabBar_ElideMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabBar) Expanding() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_Expanding(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) IsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_IsMovable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) SelectionBehaviorOnRemove() QTabBar__SelectionBehavior {
	if ptr.Pointer() != nil {
		return QTabBar__SelectionBehavior(C.QTabBar_SelectionBehaviorOnRemove(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabBar) SetAutoHide(hide bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetAutoHide(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTabBar) SetChangeCurrentOnDrag(change bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetChangeCurrentOnDrag(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(change)))
	}
}

func (ptr *QTabBar) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QTabBar) SetDocumentMode(set bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetDocumentMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabBar) SetDrawBase(drawTheBase bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetDrawBase(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(drawTheBase)))
	}
}

func (ptr *QTabBar) SetElideMode(v core.Qt__TextElideMode) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetElideMode(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QTabBar) SetExpanding(enabled bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetExpanding(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetIconSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QTabBar) SetMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetMovable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabBar) SetSelectionBehaviorOnRemove(behavior QTabBar__SelectionBehavior) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetSelectionBehaviorOnRemove(C.QtObjectPtr(ptr.Pointer()), C.int(behavior))
	}
}

func (ptr *QTabBar) SetShape(shape QTabBar__Shape) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetShape(C.QtObjectPtr(ptr.Pointer()), C.int(shape))
	}
}

func (ptr *QTabBar) SetTabsClosable(closable bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabsClosable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QTabBar) SetUsesScrollButtons(useButtons bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetUsesScrollButtons(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabBar) Shape() QTabBar__Shape {
	if ptr.Pointer() != nil {
		return QTabBar__Shape(C.QTabBar_Shape(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabBar) TabsClosable() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_TabsClosable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabBar) UsesScrollButtons() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_UsesScrollButtons(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQTabBar(parent QWidgetITF) *QTabBar {
	return QTabBarFromPointer(unsafe.Pointer(C.QTabBar_NewQTabBar(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTabBar) AddTab2(icon gui.QIconITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) AddTab(text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab(C.QtObjectPtr(ptr.Pointer()), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabBar) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabBarCurrentChanged
func callbackQTabBarCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QTabBar) InsertTab2(index int, icon gui.QIconITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab2(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) InsertTab(index int, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) IsTabEnabled(index int) bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_IsTabEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabBar) MoveTab(from int, to int) {
	if ptr.Pointer() != nil {
		C.QTabBar_MoveTab(C.QtObjectPtr(ptr.Pointer()), C.int(from), C.int(to))
	}
}

func (ptr *QTabBar) RemoveTab(index int) {
	if ptr.Pointer() != nil {
		C.QTabBar_RemoveTab(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QTabBar) SetTabButton(index int, position QTabBar__ButtonPosition, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabButton(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(position), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QTabBar) SetTabData(index int, data string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabData(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(data))
	}
}

func (ptr *QTabBar) SetTabEnabled(index int, enabled bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetTabIcon(index int, icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabIcon(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QTabBar) SetTabText(index int, text string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabText(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) SetTabTextColor(index int, color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabTextColor(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func (ptr *QTabBar) SetTabToolTip(index int, tip string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabBar) SetTabWhatsThis(index int, text string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) TabAt(position core.QPointITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_TabAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position))))
	}
	return 0
}

func (ptr *QTabBar) ConnectTabBarClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarClicked() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabBarTabBarClicked
func callbackQTabBarTabBarClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarClicked").(func(int))(int(index))
}

func (ptr *QTabBar) ConnectTabBarDoubleClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabBarTabBarDoubleClicked
func callbackQTabBarTabBarDoubleClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked").(func(int))(int(index))
}

func (ptr *QTabBar) TabButton(index int, position QTabBar__ButtonPosition) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QTabBar_TabButton(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(position))))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabCloseRequested(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabCloseRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabBar) DisconnectTabCloseRequested() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabCloseRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabBarTabCloseRequested
func callbackQTabBarTabCloseRequested(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabCloseRequested").(func(int))(int(index))
}

func (ptr *QTabBar) TabData(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabData(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) ConnectTabMoved(f func(from int, to int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabMoved", f)
	}
}

func (ptr *QTabBar) DisconnectTabMoved() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabMoved")
	}
}

//export callbackQTabBarTabMoved
func callbackQTabBarTabMoved(ptrName *C.char, from C.int, to C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabMoved").(func(int, int))(int(from), int(to))
}

func (ptr *QTabBar) TabText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabText(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabToolTip(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabWhatsThis(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) DestroyQTabBar() {
	if ptr.Pointer() != nil {
		C.QTabBar_DestroyQTabBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
