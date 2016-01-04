package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QTabBar_") {
		n.SetObjectName("QTabBar_" + qt.Identifier())
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
	defer qt.Recovering("QTabBar::autoHide")

	if ptr.Pointer() != nil {
		return C.QTabBar_AutoHide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) ChangeCurrentOnDrag() bool {
	defer qt.Recovering("QTabBar::changeCurrentOnDrag")

	if ptr.Pointer() != nil {
		return C.QTabBar_ChangeCurrentOnDrag(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) Count() int {
	defer qt.Recovering("QTabBar::count")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) CurrentIndex() int {
	defer qt.Recovering("QTabBar::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) DocumentMode() bool {
	defer qt.Recovering("QTabBar::documentMode")

	if ptr.Pointer() != nil {
		return C.QTabBar_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) DrawBase() bool {
	defer qt.Recovering("QTabBar::drawBase")

	if ptr.Pointer() != nil {
		return C.QTabBar_DrawBase(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) ElideMode() core.Qt__TextElideMode {
	defer qt.Recovering("QTabBar::elideMode")

	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabBar_ElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) Expanding() bool {
	defer qt.Recovering("QTabBar::expanding")

	if ptr.Pointer() != nil {
		return C.QTabBar_Expanding(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) IconSize() *core.QSize {
	defer qt.Recovering("QTabBar::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabBar_IconSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabBar) IsMovable() bool {
	defer qt.Recovering("QTabBar::isMovable")

	if ptr.Pointer() != nil {
		return C.QTabBar_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) SelectionBehaviorOnRemove() QTabBar__SelectionBehavior {
	defer qt.Recovering("QTabBar::selectionBehaviorOnRemove")

	if ptr.Pointer() != nil {
		return QTabBar__SelectionBehavior(C.QTabBar_SelectionBehaviorOnRemove(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) SetAutoHide(hide bool) {
	defer qt.Recovering("QTabBar::setAutoHide")

	if ptr.Pointer() != nil {
		C.QTabBar_SetAutoHide(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTabBar) SetChangeCurrentOnDrag(change bool) {
	defer qt.Recovering("QTabBar::setChangeCurrentOnDrag")

	if ptr.Pointer() != nil {
		C.QTabBar_SetChangeCurrentOnDrag(ptr.Pointer(), C.int(qt.GoBoolToInt(change)))
	}
}

func (ptr *QTabBar) SetCurrentIndex(index int) {
	defer qt.Recovering("QTabBar::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QTabBar_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) SetDocumentMode(set bool) {
	defer qt.Recovering("QTabBar::setDocumentMode")

	if ptr.Pointer() != nil {
		C.QTabBar_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabBar) SetDrawBase(drawTheBase bool) {
	defer qt.Recovering("QTabBar::setDrawBase")

	if ptr.Pointer() != nil {
		C.QTabBar_SetDrawBase(ptr.Pointer(), C.int(qt.GoBoolToInt(drawTheBase)))
	}
}

func (ptr *QTabBar) SetElideMode(v core.Qt__TextElideMode) {
	defer qt.Recovering("QTabBar::setElideMode")

	if ptr.Pointer() != nil {
		C.QTabBar_SetElideMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabBar) SetExpanding(enabled bool) {
	defer qt.Recovering("QTabBar::setExpanding")

	if ptr.Pointer() != nil {
		C.QTabBar_SetExpanding(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetIconSize(size core.QSize_ITF) {
	defer qt.Recovering("QTabBar::setIconSize")

	if ptr.Pointer() != nil {
		C.QTabBar_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QTabBar) SetMovable(movable bool) {
	defer qt.Recovering("QTabBar::setMovable")

	if ptr.Pointer() != nil {
		C.QTabBar_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabBar) SetSelectionBehaviorOnRemove(behavior QTabBar__SelectionBehavior) {
	defer qt.Recovering("QTabBar::setSelectionBehaviorOnRemove")

	if ptr.Pointer() != nil {
		C.QTabBar_SetSelectionBehaviorOnRemove(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QTabBar) SetShape(shape QTabBar__Shape) {
	defer qt.Recovering("QTabBar::setShape")

	if ptr.Pointer() != nil {
		C.QTabBar_SetShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QTabBar) SetTabsClosable(closable bool) {
	defer qt.Recovering("QTabBar::setTabsClosable")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QTabBar) SetUsesScrollButtons(useButtons bool) {
	defer qt.Recovering("QTabBar::setUsesScrollButtons")

	if ptr.Pointer() != nil {
		C.QTabBar_SetUsesScrollButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabBar) Shape() QTabBar__Shape {
	defer qt.Recovering("QTabBar::shape")

	if ptr.Pointer() != nil {
		return QTabBar__Shape(C.QTabBar_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabBar) TabsClosable() bool {
	defer qt.Recovering("QTabBar::tabsClosable")

	if ptr.Pointer() != nil {
		return C.QTabBar_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabBar) UsesScrollButtons() bool {
	defer qt.Recovering("QTabBar::usesScrollButtons")

	if ptr.Pointer() != nil {
		return C.QTabBar_UsesScrollButtons(ptr.Pointer()) != 0
	}
	return false
}

func NewQTabBar(parent QWidget_ITF) *QTabBar {
	defer qt.Recovering("QTabBar::QTabBar")

	return NewQTabBarFromPointer(C.QTabBar_NewQTabBar(PointerFromQWidget(parent)))
}

func (ptr *QTabBar) AddTab2(icon gui.QIcon_ITF, text string) int {
	defer qt.Recovering("QTabBar::addTab")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) AddTab(text string) int {
	defer qt.Recovering("QTabBar::addTab")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_AddTab(ptr.Pointer(), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTabBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTabBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTabBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTabBarChangeEvent
func callbackQTabBarChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTabBar) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) ConnectCurrentChanged(f func(index int)) {
	defer qt.Recovering("connect QTabBar::currentChanged")

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabBar) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QTabBar::currentChanged")

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabBarCurrentChanged
func callbackQTabBarCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabBar::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabBar) CurrentChanged(index int) {
	defer qt.Recovering("QTabBar::currentChanged")

	if ptr.Pointer() != nil {
		C.QTabBar_CurrentChanged(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QTabBar::event")

	if ptr.Pointer() != nil {
		return C.QTabBar_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTabBar) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QTabBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTabBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTabBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTabBarHideEvent
func callbackQTabBarHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQTabBarFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QTabBar) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QTabBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QTabBar) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QTabBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QTabBar) InsertTab2(index int, icon gui.QIcon_ITF, text string) int {
	defer qt.Recovering("QTabBar::insertTab")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab2(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) InsertTab(index int, text string) int {
	defer qt.Recovering("QTabBar::insertTab")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_InsertTab(ptr.Pointer(), C.int(index), C.CString(text)))
	}
	return 0
}

func (ptr *QTabBar) IsTabEnabled(index int) bool {
	defer qt.Recovering("QTabBar::isTabEnabled")

	if ptr.Pointer() != nil {
		return C.QTabBar_IsTabEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabBar) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTabBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTabBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTabBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTabBarKeyPressEvent
func callbackQTabBarKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QTabBar) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTabBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTabBar) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTabBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTabBar) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QTabBar::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabBar_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabBar) MinimumTabSizeHint(index int) *core.QSize {
	defer qt.Recovering("QTabBar::minimumTabSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabBar_MinimumTabSizeHint(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTabBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTabBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTabBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTabBarMouseMoveEvent
func callbackQTabBarMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTabBar) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTabBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTabBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTabBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTabBarMousePressEvent
func callbackQTabBarMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTabBar) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTabBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTabBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTabBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTabBarMouseReleaseEvent
func callbackQTabBarMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTabBar) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) MoveTab(from int, to int) {
	defer qt.Recovering("QTabBar::moveTab")

	if ptr.Pointer() != nil {
		C.QTabBar_MoveTab(ptr.Pointer(), C.int(from), C.int(to))
	}
}

func (ptr *QTabBar) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTabBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTabBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTabBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTabBarPaintEvent
func callbackQTabBarPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQTabBarFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QTabBar) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTabBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QTabBar) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTabBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QTabBar) RemoveTab(index int) {
	defer qt.Recovering("QTabBar::removeTab")

	if ptr.Pointer() != nil {
		C.QTabBar_RemoveTab(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTabBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTabBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTabBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTabBarResizeEvent
func callbackQTabBarResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQTabBarFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QTabBar) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTabBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QTabBar) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTabBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QTabBar) SetTabButton(index int, position QTabBar__ButtonPosition, widget QWidget_ITF) {
	defer qt.Recovering("QTabBar::setTabButton")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabButton(ptr.Pointer(), C.int(index), C.int(position), PointerFromQWidget(widget))
	}
}

func (ptr *QTabBar) SetTabData(index int, data core.QVariant_ITF) {
	defer qt.Recovering("QTabBar::setTabData")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabData(ptr.Pointer(), C.int(index), core.PointerFromQVariant(data))
	}
}

func (ptr *QTabBar) SetTabEnabled(index int, enabled bool) {
	defer qt.Recovering("QTabBar::setTabEnabled")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabBar) SetTabIcon(index int, icon gui.QIcon_ITF) {
	defer qt.Recovering("QTabBar::setTabIcon")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTabBar) SetTabText(index int, text string) {
	defer qt.Recovering("QTabBar::setTabText")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) SetTabTextColor(index int, color gui.QColor_ITF) {
	defer qt.Recovering("QTabBar::setTabTextColor")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabTextColor(ptr.Pointer(), C.int(index), gui.PointerFromQColor(color))
	}
}

func (ptr *QTabBar) SetTabToolTip(index int, tip string) {
	defer qt.Recovering("QTabBar::setTabToolTip")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabToolTip(ptr.Pointer(), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabBar) SetTabWhatsThis(index int, text string) {
	defer qt.Recovering("QTabBar::setTabWhatsThis")

	if ptr.Pointer() != nil {
		C.QTabBar_SetTabWhatsThis(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabBar) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QTabBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTabBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTabBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTabBarShowEvent
func callbackQTabBarShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQTabBarFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QTabBar) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QTabBar::showEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QTabBar) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QTabBar::showEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QTabBar) SizeHint() *core.QSize {
	defer qt.Recovering("QTabBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabBar) TabAt(position core.QPoint_ITF) int {
	defer qt.Recovering("QTabBar::tabAt")

	if ptr.Pointer() != nil {
		return int(C.QTabBar_TabAt(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return 0
}

func (ptr *QTabBar) ConnectTabBarClicked(f func(index int)) {
	defer qt.Recovering("connect QTabBar::tabBarClicked")

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarClicked() {
	defer qt.Recovering("disconnect QTabBar::tabBarClicked")

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabBarTabBarClicked
func callbackQTabBarTabBarClicked(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabBar::tabBarClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabBarClicked"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabBar) TabBarClicked(index int) {
	defer qt.Recovering("QTabBar::tabBarClicked")

	if ptr.Pointer() != nil {
		C.QTabBar_TabBarClicked(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) ConnectTabBarDoubleClicked(f func(index int)) {
	defer qt.Recovering("connect QTabBar::tabBarDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabBarDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabBar) DisconnectTabBarDoubleClicked() {
	defer qt.Recovering("disconnect QTabBar::tabBarDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabBarDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabBarTabBarDoubleClicked
func callbackQTabBarTabBarDoubleClicked(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabBar::tabBarDoubleClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabBar) TabBarDoubleClicked(index int) {
	defer qt.Recovering("QTabBar::tabBarDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTabBar_TabBarDoubleClicked(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) TabButton(index int, position QTabBar__ButtonPosition) *QWidget {
	defer qt.Recovering("QTabBar::tabButton")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabBar_TabButton(ptr.Pointer(), C.int(index), C.int(position)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabCloseRequested(f func(index int)) {
	defer qt.Recovering("connect QTabBar::tabCloseRequested")

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabBar) DisconnectTabCloseRequested() {
	defer qt.Recovering("disconnect QTabBar::tabCloseRequested")

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabBarTabCloseRequested
func callbackQTabBarTabCloseRequested(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabBar::tabCloseRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabCloseRequested"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabBar) TabCloseRequested(index int) {
	defer qt.Recovering("QTabBar::tabCloseRequested")

	if ptr.Pointer() != nil {
		C.QTabBar_TabCloseRequested(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) TabData(index int) *core.QVariant {
	defer qt.Recovering("QTabBar::tabData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTabBar_TabData(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) TabIcon(index int) *gui.QIcon {
	defer qt.Recovering("QTabBar::tabIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QTabBar_TabIcon(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabInserted(f func(index int)) {
	defer qt.Recovering("connect QTabBar::tabInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabInserted", f)
	}
}

func (ptr *QTabBar) DisconnectTabInserted() {
	defer qt.Recovering("disconnect QTabBar::tabInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabInserted")
	}
}

//export callbackQTabBarTabInserted
func callbackQTabBarTabInserted(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabBar::tabInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabInserted"); signal != nil {
		signal.(func(int))(int(index))
	} else {
		NewQTabBarFromPointer(ptr).TabInsertedDefault(int(index))
	}
}

func (ptr *QTabBar) TabInserted(index int) {
	defer qt.Recovering("QTabBar::tabInserted")

	if ptr.Pointer() != nil {
		C.QTabBar_TabInserted(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) TabInsertedDefault(index int) {
	defer qt.Recovering("QTabBar::tabInserted")

	if ptr.Pointer() != nil {
		C.QTabBar_TabInsertedDefault(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) ConnectTabLayoutChange(f func()) {
	defer qt.Recovering("connect QTabBar::tabLayoutChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabLayoutChange", f)
	}
}

func (ptr *QTabBar) DisconnectTabLayoutChange() {
	defer qt.Recovering("disconnect QTabBar::tabLayoutChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabLayoutChange")
	}
}

//export callbackQTabBarTabLayoutChange
func callbackQTabBarTabLayoutChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTabBar::tabLayoutChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabLayoutChange"); signal != nil {
		signal.(func())()
	} else {
		NewQTabBarFromPointer(ptr).TabLayoutChangeDefault()
	}
}

func (ptr *QTabBar) TabLayoutChange() {
	defer qt.Recovering("QTabBar::tabLayoutChange")

	if ptr.Pointer() != nil {
		C.QTabBar_TabLayoutChange(ptr.Pointer())
	}
}

func (ptr *QTabBar) TabLayoutChangeDefault() {
	defer qt.Recovering("QTabBar::tabLayoutChange")

	if ptr.Pointer() != nil {
		C.QTabBar_TabLayoutChangeDefault(ptr.Pointer())
	}
}

func (ptr *QTabBar) ConnectTabMoved(f func(from int, to int)) {
	defer qt.Recovering("connect QTabBar::tabMoved")

	if ptr.Pointer() != nil {
		C.QTabBar_ConnectTabMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabMoved", f)
	}
}

func (ptr *QTabBar) DisconnectTabMoved() {
	defer qt.Recovering("disconnect QTabBar::tabMoved")

	if ptr.Pointer() != nil {
		C.QTabBar_DisconnectTabMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabMoved")
	}
}

//export callbackQTabBarTabMoved
func callbackQTabBarTabMoved(ptr unsafe.Pointer, ptrName *C.char, from C.int, to C.int) {
	defer qt.Recovering("callback QTabBar::tabMoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabMoved"); signal != nil {
		signal.(func(int, int))(int(from), int(to))
	}

}

func (ptr *QTabBar) TabMoved(from int, to int) {
	defer qt.Recovering("QTabBar::tabMoved")

	if ptr.Pointer() != nil {
		C.QTabBar_TabMoved(ptr.Pointer(), C.int(from), C.int(to))
	}
}

func (ptr *QTabBar) TabRect(index int) *core.QRect {
	defer qt.Recovering("QTabBar::tabRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTabBar_TabRect(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) ConnectTabRemoved(f func(index int)) {
	defer qt.Recovering("connect QTabBar::tabRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabRemoved", f)
	}
}

func (ptr *QTabBar) DisconnectTabRemoved() {
	defer qt.Recovering("disconnect QTabBar::tabRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabRemoved")
	}
}

//export callbackQTabBarTabRemoved
func callbackQTabBarTabRemoved(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabBar::tabRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabRemoved"); signal != nil {
		signal.(func(int))(int(index))
	} else {
		NewQTabBarFromPointer(ptr).TabRemovedDefault(int(index))
	}
}

func (ptr *QTabBar) TabRemoved(index int) {
	defer qt.Recovering("QTabBar::tabRemoved")

	if ptr.Pointer() != nil {
		C.QTabBar_TabRemoved(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) TabRemovedDefault(index int) {
	defer qt.Recovering("QTabBar::tabRemoved")

	if ptr.Pointer() != nil {
		C.QTabBar_TabRemovedDefault(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabBar) TabSizeHint(index int) *core.QSize {
	defer qt.Recovering("QTabBar::tabSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabBar_TabSizeHint(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) TabText(index int) string {
	defer qt.Recovering("QTabBar::tabText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabTextColor(index int) *gui.QColor {
	defer qt.Recovering("QTabBar::tabTextColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QTabBar_TabTextColor(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabBar) TabToolTip(index int) string {
	defer qt.Recovering("QTabBar::tabToolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) TabWhatsThis(index int) string {
	defer qt.Recovering("QTabBar::tabWhatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabBar_TabWhatsThis(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTabBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTabBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTabBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTabBarTimerEvent
func callbackQTabBarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTabBar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTabBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTabBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTabBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTabBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTabBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTabBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTabBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTabBarWheelEvent
func callbackQTabBarWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QTabBar) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTabBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QTabBar) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTabBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QTabBar) DestroyQTabBar() {
	defer qt.Recovering("QTabBar::~QTabBar")

	if ptr.Pointer() != nil {
		C.QTabBar_DestroyQTabBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTabBar) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTabBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTabBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTabBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTabBarActionEvent
func callbackQTabBarActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QTabBar) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTabBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTabBar) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTabBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTabBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTabBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTabBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTabBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTabBarDragEnterEvent
func callbackQTabBarDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QTabBar) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTabBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QTabBar) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTabBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QTabBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTabBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTabBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTabBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTabBarDragLeaveEvent
func callbackQTabBarDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QTabBar) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTabBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QTabBar) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTabBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QTabBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTabBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTabBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTabBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTabBarDragMoveEvent
func callbackQTabBarDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QTabBar) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTabBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QTabBar) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTabBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QTabBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QTabBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTabBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTabBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTabBarDropEvent
func callbackQTabBarDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QTabBar) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QTabBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QTabBar) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QTabBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QTabBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTabBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTabBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTabBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTabBarEnterEvent
func callbackQTabBarEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTabBar) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTabBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTabBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTabBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTabBarFocusInEvent
func callbackQTabBarFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QTabBar) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTabBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTabBar) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTabBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTabBar) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTabBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTabBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTabBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTabBarFocusOutEvent
func callbackQTabBarFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QTabBar) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTabBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTabBar) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTabBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTabBar) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTabBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTabBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTabBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTabBarLeaveEvent
func callbackQTabBarLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTabBar) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTabBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTabBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTabBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTabBarMoveEvent
func callbackQTabBarMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QTabBar) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTabBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTabBar) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTabBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTabBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTabBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTabBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTabBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTabBarSetVisible
func callbackQTabBarSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTabBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTabBar) SetVisible(visible bool) {
	defer qt.Recovering("QTabBar::setVisible")

	if ptr.Pointer() != nil {
		C.QTabBar_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTabBar) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QTabBar::setVisible")

	if ptr.Pointer() != nil {
		C.QTabBar_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTabBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTabBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTabBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTabBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTabBarCloseEvent
func callbackQTabBarCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QTabBar) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTabBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTabBar) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTabBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTabBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTabBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTabBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTabBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTabBarContextMenuEvent
func callbackQTabBarContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QTabBar) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTabBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTabBar) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTabBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QTabBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTabBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTabBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTabBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTabBarInitPainter
func callbackQTabBarInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQTabBarFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QTabBar) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTabBar::initPainter")

	if ptr.Pointer() != nil {
		C.QTabBar_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTabBar) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTabBar::initPainter")

	if ptr.Pointer() != nil {
		C.QTabBar_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTabBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTabBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTabBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTabBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTabBarInputMethodEvent
func callbackQTabBarInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QTabBar) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTabBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QTabBar) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTabBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QTabBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTabBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTabBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTabBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTabBarKeyReleaseEvent
func callbackQTabBarKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QTabBar) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTabBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTabBar) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTabBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTabBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTabBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTabBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTabBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTabBarMouseDoubleClickEvent
func callbackQTabBarMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTabBar) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTabBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTabBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTabBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTabBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTabBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTabBarTabletEvent
func callbackQTabBarTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QTabBar) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTabBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTabBar) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTabBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTabBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTabBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTabBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTabBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTabBarChildEvent
func callbackQTabBarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTabBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTabBar::childEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTabBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTabBar::childEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTabBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTabBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTabBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTabBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTabBarCustomEvent
func callbackQTabBarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTabBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTabBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTabBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::customEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTabBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTabBar::customEvent")

	if ptr.Pointer() != nil {
		C.QTabBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
