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
	for len(n.ObjectName()) < len("QTabBar_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::autoHide")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_AutoHide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) ChangeCurrentOnDrag() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::changeCurrentOnDrag")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_ChangeCurrentOnDrag(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) DocumentMode() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::documentMode")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) DrawBase() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::drawBase")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_DrawBase(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) ElideMode() core.Qt__TextElideMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::elideMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabBar_ElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) Expanding() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::expanding")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_Expanding(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) IsMovable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::isMovable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) SelectionBehaviorOnRemove() QTabBar__SelectionBehavior {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::selectionBehaviorOnRemove")
		}
	}()

	if ptr.Pointer() != nil {
		return QTabBar__SelectionBehavior(C.QTabBar_SelectionBehaviorOnRemove(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) SetAutoHide(hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setAutoHide")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetAutoHide(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTabBar) SetChangeCurrentOnDrag(change bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setChangeCurrentOnDrag")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetChangeCurrentOnDrag(ptr.Pointer(), C.int(qt.GoBoolToInt(change)))
	}
}

func (ptr *QTabBar) SetCurrentIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) SetDocumentMode(set bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setDocumentMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabBar) SetDrawBase(drawTheBase bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setDrawBase")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetDrawBase(ptr.Pointer(), C.int(qt.GoBoolToInt(drawTheBase)))
	}
}

func (ptr *QTabBar) SetElideMode(v core.Qt__TextElideMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setElideMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetElideMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabBar) SetExpanding(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setExpanding")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetExpanding(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetIconSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setIconSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QTabBar) SetMovable(movable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setMovable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabBar) SetSelectionBehaviorOnRemove(behavior QTabBar__SelectionBehavior) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setSelectionBehaviorOnRemove")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetSelectionBehaviorOnRemove(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QTabBar) SetShape(shape QTabBar__Shape) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setShape")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QTabBar) SetTabsClosable(closable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabsClosable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QTabBar) SetUsesScrollButtons(useButtons bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setUsesScrollButtons")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetUsesScrollButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabBar) Shape() QTabBar__Shape {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::shape")
		}
	}()

	if ptr.Pointer() != nil {
		return QTabBar__Shape(C.QTabBar_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) TabsClosable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabsClosable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) UsesScrollButtons() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::usesScrollButtons")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_UsesScrollButtons(ptr.Pointer()) != 0
	}
	return false
}

func NewQTabBar(parent QWidget_ITF) *QTabBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::QTabBar")
		}
	}()

	return NewQTabBarFromPointer(C.QTabBar_NewQTabBar(PointerFromQWidget(parent)))
}

func (ptr *QTabBar) AddTab2(icon gui.QIcon_ITF, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::addTab")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) AddTab(text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::addTab")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab(ptr.Pointer(), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) ConnectCurrentChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabBar) DisconnectCurrentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabBarCurrentChanged
func callbackQTabBarCurrentChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::currentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QTabBar) InsertTab2(index int, icon gui.QIcon_ITF, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::insertTab")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab2(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) InsertTab(index int, text string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::insertTab")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab(ptr.Pointer(), C.int(index), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) IsTabEnabled(index int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::isTabEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTabBar_IsTabEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabBar) MoveTab(from int, to int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::moveTab")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_MoveTab(ptr.Pointer(), C.int(from), C.int(to))
	}
}

func (ptr *QTabBar) RemoveTab(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::removeTab")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_RemoveTab(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) SetTabButton(index int, position QTabBar__ButtonPosition, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabButton(ptr.Pointer(), C.int(index), C.int(position), PointerFromQWidget(widget))
	}
}

func (ptr *QTabBar) SetTabData(index int, data core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabData(ptr.Pointer(), C.int(index), core.PointerFromQVariant(data))
	}
}

func (ptr *QTabBar) SetTabEnabled(index int, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetTabIcon(index int, icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTabBar) SetTabText(index int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) SetTabTextColor(index int, color gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabTextColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabTextColor(ptr.Pointer(), C.int(index), gui.PointerFromQColor(color))
	}
}

func (ptr *QTabBar) SetTabToolTip(index int, tip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabToolTip(ptr.Pointer(), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabBar) SetTabWhatsThis(index int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::setTabWhatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabWhatsThis(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) TabAt(position core.QPoint_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabBar_TabAt(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return 0
}

func (ptr *QTabBar) ConnectTabBarClicked(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabBarClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabBarClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabBarTabBarClicked
func callbackQTabBarTabBarClicked(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabBarClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "tabBarClicked").(func(int))(int(index))
}

func (ptr *QTabBar) ConnectTabBarDoubleClicked(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabBarDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarDoubleClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabBarDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabBarTabBarDoubleClicked
func callbackQTabBarTabBarDoubleClicked(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabBarDoubleClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked").(func(int))(int(index))
}

func (ptr *QTabBar) TabButton(index int, position QTabBar__ButtonPosition) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabButton")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabBar_TabButton(ptr.Pointer(), C.int(index), C.int(position)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabCloseRequested(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabCloseRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabBar) DisconnectTabCloseRequested() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabCloseRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabBarTabCloseRequested
func callbackQTabBarTabCloseRequested(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabCloseRequested")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "tabCloseRequested").(func(int))(int(index))
}

func (ptr *QTabBar) TabData(index int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTabBar_TabData(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabMoved(f func(from int, to int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabMoved", f)
	}
}

func (ptr *QTabBar) DisconnectTabMoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabMoved")
	}
}

//export callbackQTabBarTabMoved
func callbackQTabBarTabMoved(ptrName *C.char, from C.int, to C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabMoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "tabMoved").(func(int, int))(int(from), int(to))
}

func (ptr *QTabBar) TabText(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabTextColor(index int) *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabTextColor")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTabBar_TabTextColor(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) TabToolTip(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabWhatsThis(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::tabWhatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabWhatsThis(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) DestroyQTabBar() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabBar::~QTabBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTabBar_DestroyQTabBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
