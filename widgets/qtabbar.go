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

type QTabBar_ITF interface {
	QWidget_ITF
	QTabBar_PTR() *QTabBar
}

func PointerFromQTabBar(ptr QTabBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabBar_PTR().Pointer()
	}
	return nil
}

func NewQTabBarFromPointer(ptr unsafe.Pointer) *QTabBar {
	var n = new(QTabBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTabBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTabBar) QTabBar_PTR() *QTabBar {
	return ptr
}

//QTabBar::ButtonPosition
type QTabBar__ButtonPosition int64

const (
	QTabBar__LeftSide  = QTabBar__ButtonPosition(0)
	QTabBar__RightSide = QTabBar__ButtonPosition(1)
)

//QTabBar::SelectionBehavior
type QTabBar__SelectionBehavior int64

const (
	QTabBar__SelectLeftTab     = QTabBar__SelectionBehavior(0)
	QTabBar__SelectRightTab    = QTabBar__SelectionBehavior(1)
	QTabBar__SelectPreviousTab = QTabBar__SelectionBehavior(2)
)

//QTabBar::Shape
type QTabBar__Shape int64

const (
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
		return C.QTabBar_AutoHide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) ChangeCurrentOnDrag() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_ChangeCurrentOnDrag(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) DrawBase() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_DrawBase(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) ElideMode() core.Qt__TextElideMode {
	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabBar_ElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) Expanding() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_Expanding(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) IsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) SelectionBehaviorOnRemove() QTabBar__SelectionBehavior {
	if ptr.Pointer() != nil {
		return QTabBar__SelectionBehavior(C.QTabBar_SelectionBehaviorOnRemove(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) SetAutoHide(hide bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetAutoHide(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTabBar) SetChangeCurrentOnDrag(change bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetChangeCurrentOnDrag(ptr.Pointer(), C.int(qt.GoBoolToInt(change)))
	}
}

func (ptr *QTabBar) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) SetDocumentMode(set bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabBar) SetDrawBase(drawTheBase bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetDrawBase(ptr.Pointer(), C.int(qt.GoBoolToInt(drawTheBase)))
	}
}

func (ptr *QTabBar) SetElideMode(v core.Qt__TextElideMode) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetElideMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabBar) SetExpanding(enabled bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetExpanding(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetIconSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QTabBar) SetMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabBar) SetSelectionBehaviorOnRemove(behavior QTabBar__SelectionBehavior) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetSelectionBehaviorOnRemove(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QTabBar) SetShape(shape QTabBar__Shape) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QTabBar) SetTabsClosable(closable bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QTabBar) SetUsesScrollButtons(useButtons bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetUsesScrollButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabBar) Shape() QTabBar__Shape {
	if ptr.Pointer() != nil {
		return QTabBar__Shape(C.QTabBar_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) TabsClosable() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) UsesScrollButtons() bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_UsesScrollButtons(ptr.Pointer()) != 0
	}
	return false
}

func NewQTabBar(parent QWidget_ITF) *QTabBar {
	return NewQTabBarFromPointer(C.QTabBar_NewQTabBar(PointerFromQWidget(parent)))
}

func (ptr *QTabBar) AddTab2(icon gui.QIcon_ITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) AddTab(text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab(ptr.Pointer(), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabBar) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabBarCurrentChanged
func callbackQTabBarCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QTabBar) InsertTab2(index int, icon gui.QIcon_ITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab2(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) InsertTab(index int, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab(ptr.Pointer(), C.int(index), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) IsTabEnabled(index int) bool {
	if ptr.Pointer() != nil {
		return C.QTabBar_IsTabEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabBar) MoveTab(from int, to int) {
	if ptr.Pointer() != nil {
		C.QTabBar_MoveTab(ptr.Pointer(), C.int(from), C.int(to))
	}
}

func (ptr *QTabBar) RemoveTab(index int) {
	if ptr.Pointer() != nil {
		C.QTabBar_RemoveTab(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) SetTabButton(index int, position QTabBar__ButtonPosition, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabButton(ptr.Pointer(), C.int(index), C.int(position), PointerFromQWidget(widget))
	}
}

func (ptr *QTabBar) SetTabData(index int, data core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabData(ptr.Pointer(), C.int(index), core.PointerFromQVariant(data))
	}
}

func (ptr *QTabBar) SetTabEnabled(index int, enabled bool) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetTabIcon(index int, icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTabBar) SetTabText(index int, text string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) SetTabTextColor(index int, color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabTextColor(ptr.Pointer(), C.int(index), gui.PointerFromQColor(color))
	}
}

func (ptr *QTabBar) SetTabToolTip(index int, tip string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabToolTip(ptr.Pointer(), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabBar) SetTabWhatsThis(index int, text string) {
	if ptr.Pointer() != nil {
		C.QTabBar_SetTabWhatsThis(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) TabAt(position core.QPoint_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTabBar_TabAt(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return 0
}

func (ptr *QTabBar) ConnectTabBarClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarClicked() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabBarTabBarClicked
func callbackQTabBarTabBarClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarClicked").(func(int))(int(index))
}

func (ptr *QTabBar) ConnectTabBarDoubleClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabBarTabBarDoubleClicked
func callbackQTabBarTabBarDoubleClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked").(func(int))(int(index))
}

func (ptr *QTabBar) TabButton(index int, position QTabBar__ButtonPosition) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabBar_TabButton(ptr.Pointer(), C.int(index), C.int(position)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabCloseRequested(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabBar) DisconnectTabCloseRequested() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabBarTabCloseRequested
func callbackQTabBarTabCloseRequested(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabCloseRequested").(func(int))(int(index))
}

func (ptr *QTabBar) TabData(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTabBar_TabData(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabMoved(f func(from int, to int)) {
	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabMoved", f)
	}
}

func (ptr *QTabBar) DisconnectTabMoved() {
	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabMoved")
	}
}

//export callbackQTabBarTabMoved
func callbackQTabBarTabMoved(ptrName *C.char, from C.int, to C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabMoved").(func(int, int))(int(from), int(to))
}

func (ptr *QTabBar) TabText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabTextColor(index int) *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTabBar_TabTextColor(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) TabToolTip(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabWhatsThis(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabWhatsThis(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) DestroyQTabBar() {
	if ptr.Pointer() != nil {
		C.QTabBar_DestroyQTabBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
