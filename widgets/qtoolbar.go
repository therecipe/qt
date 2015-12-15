package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolBar struct {
	QWidget
}

type QToolBar_ITF interface {
	QWidget_ITF
	QToolBar_PTR() *QToolBar
}

func PointerFromQToolBar(ptr QToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolBar_PTR().Pointer()
	}
	return nil
}

func NewQToolBarFromPointer(ptr unsafe.Pointer) *QToolBar {
	var n = new(QToolBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QToolBar_") {
		n.SetObjectName("QToolBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QToolBar) QToolBar_PTR() *QToolBar {
	return ptr
}

func (ptr *QToolBar) AllowedAreas() core.Qt__ToolBarArea {
	defer qt.Recovering("QToolBar::allowedAreas")

	if ptr.Pointer() != nil {
		return core.Qt__ToolBarArea(C.QToolBar_AllowedAreas(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBar) IsFloatable() bool {
	defer qt.Recovering("QToolBar::isFloatable")

	if ptr.Pointer() != nil {
		return C.QToolBar_IsFloatable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolBar) IsFloating() bool {
	defer qt.Recovering("QToolBar::isFloating")

	if ptr.Pointer() != nil {
		return C.QToolBar_IsFloating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolBar) IsMovable() bool {
	defer qt.Recovering("QToolBar::isMovable")

	if ptr.Pointer() != nil {
		return C.QToolBar_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolBar) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QToolBar::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QToolBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBar) SetAllowedAreas(areas core.Qt__ToolBarArea) {
	defer qt.Recovering("QToolBar::setAllowedAreas")

	if ptr.Pointer() != nil {
		C.QToolBar_SetAllowedAreas(ptr.Pointer(), C.int(areas))
	}
}

func (ptr *QToolBar) SetFloatable(floatable bool) {
	defer qt.Recovering("QToolBar::setFloatable")

	if ptr.Pointer() != nil {
		C.QToolBar_SetFloatable(ptr.Pointer(), C.int(qt.GoBoolToInt(floatable)))
	}
}

func (ptr *QToolBar) SetIconSize(iconSize core.QSize_ITF) {
	defer qt.Recovering("QToolBar::setIconSize")

	if ptr.Pointer() != nil {
		C.QToolBar_SetIconSize(ptr.Pointer(), core.PointerFromQSize(iconSize))
	}
}

func (ptr *QToolBar) SetMovable(movable bool) {
	defer qt.Recovering("QToolBar::setMovable")

	if ptr.Pointer() != nil {
		C.QToolBar_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QToolBar) SetOrientation(orientation core.Qt__Orientation) {
	defer qt.Recovering("QToolBar::setOrientation")

	if ptr.Pointer() != nil {
		C.QToolBar_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QToolBar) SetToolButtonStyle(toolButtonStyle core.Qt__ToolButtonStyle) {
	defer qt.Recovering("QToolBar::setToolButtonStyle")

	if ptr.Pointer() != nil {
		C.QToolBar_SetToolButtonStyle(ptr.Pointer(), C.int(toolButtonStyle))
	}
}

func (ptr *QToolBar) ToolButtonStyle() core.Qt__ToolButtonStyle {
	defer qt.Recovering("QToolBar::toolButtonStyle")

	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolBar_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func NewQToolBar2(parent QWidget_ITF) *QToolBar {
	defer qt.Recovering("QToolBar::QToolBar")

	return NewQToolBarFromPointer(C.QToolBar_NewQToolBar2(PointerFromQWidget(parent)))
}

func NewQToolBar(title string, parent QWidget_ITF) *QToolBar {
	defer qt.Recovering("QToolBar::QToolBar")

	return NewQToolBarFromPointer(C.QToolBar_NewQToolBar(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QToolBar) ActionAt(p core.QPoint_ITF) *QAction {
	defer qt.Recovering("QToolBar::actionAt")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_ActionAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QToolBar) ActionAt2(x int, y int) *QAction {
	defer qt.Recovering("QToolBar::actionAt")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_ActionAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QToolBar) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QToolBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QToolBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QToolBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQToolBarActionEvent
func callbackQToolBarActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBar::actionEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "actionEvent")
	if signal != nil {
		defer signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolBar) ConnectActionTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QToolBar::actionTriggered")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectActionTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QToolBar) DisconnectActionTriggered() {
	defer qt.Recovering("disconnect QToolBar::actionTriggered")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectActionTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQToolBarActionTriggered
func callbackQToolBarActionTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::actionTriggered")

	var signal = qt.GetSignal(C.GoString(ptrName), "actionTriggered")
	if signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QToolBar) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	defer qt.Recovering("QToolBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QToolBar) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string) *QAction {
	defer qt.Recovering("QToolBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QToolBar) AddAction(text string) *QAction {
	defer qt.Recovering("QToolBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QToolBar) AddAction3(text string, receiver core.QObject_ITF, member string) *QAction {
	defer qt.Recovering("QToolBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction3(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QToolBar) AddSeparator() *QAction {
	defer qt.Recovering("QToolBar::addSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBar) AddWidget(widget QWidget_ITF) *QAction {
	defer qt.Recovering("QToolBar::addWidget")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QToolBar) ConnectAllowedAreasChanged(f func(allowedAreas core.Qt__ToolBarArea)) {
	defer qt.Recovering("connect QToolBar::allowedAreasChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectAllowedAreasChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "allowedAreasChanged", f)
	}
}

func (ptr *QToolBar) DisconnectAllowedAreasChanged() {
	defer qt.Recovering("disconnect QToolBar::allowedAreasChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectAllowedAreasChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "allowedAreasChanged")
	}
}

//export callbackQToolBarAllowedAreasChanged
func callbackQToolBarAllowedAreasChanged(ptrName *C.char, allowedAreas C.int) {
	defer qt.Recovering("callback QToolBar::allowedAreasChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged")
	if signal != nil {
		signal.(func(core.Qt__ToolBarArea))(core.Qt__ToolBarArea(allowedAreas))
	}

}

func (ptr *QToolBar) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QToolBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QToolBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQToolBarChangeEvent
func callbackQToolBarChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBar::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolBar) Clear() {
	defer qt.Recovering("QToolBar::clear")

	if ptr.Pointer() != nil {
		C.QToolBar_Clear(ptr.Pointer())
	}
}

func (ptr *QToolBar) InsertSeparator(before QAction_ITF) *QAction {
	defer qt.Recovering("QToolBar::insertSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QToolBar) InsertWidget(before QAction_ITF, widget QWidget_ITF) *QAction {
	defer qt.Recovering("QToolBar::insertWidget")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_InsertWidget(ptr.Pointer(), PointerFromQAction(before), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QToolBar) IsAreaAllowed(area core.Qt__ToolBarArea) bool {
	defer qt.Recovering("QToolBar::isAreaAllowed")

	if ptr.Pointer() != nil {
		return C.QToolBar_IsAreaAllowed(ptr.Pointer(), C.int(area)) != 0
	}
	return false
}

func (ptr *QToolBar) ConnectMovableChanged(f func(movable bool)) {
	defer qt.Recovering("connect QToolBar::movableChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectMovableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "movableChanged", f)
	}
}

func (ptr *QToolBar) DisconnectMovableChanged() {
	defer qt.Recovering("disconnect QToolBar::movableChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectMovableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "movableChanged")
	}
}

//export callbackQToolBarMovableChanged
func callbackQToolBarMovableChanged(ptrName *C.char, movable C.int) {
	defer qt.Recovering("callback QToolBar::movableChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "movableChanged")
	if signal != nil {
		signal.(func(bool))(int(movable) != 0)
	}

}

func (ptr *QToolBar) ConnectOrientationChanged(f func(orientation core.Qt__Orientation)) {
	defer qt.Recovering("connect QToolBar::orientationChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "orientationChanged", f)
	}
}

func (ptr *QToolBar) DisconnectOrientationChanged() {
	defer qt.Recovering("disconnect QToolBar::orientationChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "orientationChanged")
	}
}

//export callbackQToolBarOrientationChanged
func callbackQToolBarOrientationChanged(ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QToolBar::orientationChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "orientationChanged")
	if signal != nil {
		signal.(func(core.Qt__Orientation))(core.Qt__Orientation(orientation))
	}

}

func (ptr *QToolBar) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QToolBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QToolBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QToolBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQToolBarPaintEvent
func callbackQToolBarPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBar::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolBar) ToggleViewAction() *QAction {
	defer qt.Recovering("QToolBar::toggleViewAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_ToggleViewAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBar) ConnectToolButtonStyleChanged(f func(toolButtonStyle core.Qt__ToolButtonStyle)) {
	defer qt.Recovering("connect QToolBar::toolButtonStyleChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectToolButtonStyleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toolButtonStyleChanged", f)
	}
}

func (ptr *QToolBar) DisconnectToolButtonStyleChanged() {
	defer qt.Recovering("disconnect QToolBar::toolButtonStyleChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectToolButtonStyleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toolButtonStyleChanged")
	}
}

//export callbackQToolBarToolButtonStyleChanged
func callbackQToolBarToolButtonStyleChanged(ptrName *C.char, toolButtonStyle C.int) {
	defer qt.Recovering("callback QToolBar::toolButtonStyleChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged")
	if signal != nil {
		signal.(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

func (ptr *QToolBar) ConnectTopLevelChanged(f func(topLevel bool)) {
	defer qt.Recovering("connect QToolBar::topLevelChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectTopLevelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "topLevelChanged", f)
	}
}

func (ptr *QToolBar) DisconnectTopLevelChanged() {
	defer qt.Recovering("disconnect QToolBar::topLevelChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectTopLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "topLevelChanged")
	}
}

//export callbackQToolBarTopLevelChanged
func callbackQToolBarTopLevelChanged(ptrName *C.char, topLevel C.int) {
	defer qt.Recovering("callback QToolBar::topLevelChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "topLevelChanged")
	if signal != nil {
		signal.(func(bool))(int(topLevel) != 0)
	}

}

func (ptr *QToolBar) ConnectVisibilityChanged(f func(visible bool)) {
	defer qt.Recovering("connect QToolBar::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QToolBar) DisconnectVisibilityChanged() {
	defer qt.Recovering("disconnect QToolBar::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQToolBarVisibilityChanged
func callbackQToolBarVisibilityChanged(ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QToolBar::visibilityChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "visibilityChanged")
	if signal != nil {
		signal.(func(bool))(int(visible) != 0)
	}

}

func (ptr *QToolBar) WidgetForAction(action QAction_ITF) *QWidget {
	defer qt.Recovering("QToolBar::widgetForAction")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBar_WidgetForAction(ptr.Pointer(), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QToolBar) DestroyQToolBar() {
	defer qt.Recovering("QToolBar::~QToolBar")

	if ptr.Pointer() != nil {
		C.QToolBar_DestroyQToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
