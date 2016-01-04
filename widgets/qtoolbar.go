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

func (ptr *QToolBar) IconSize() *core.QSize {
	defer qt.Recovering("QToolBar::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QToolBar_IconSize(ptr.Pointer()))
	}
	return nil
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
func callbackQToolBarActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QToolBar) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QToolBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QToolBar) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QToolBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQToolBarActionTriggered(ptr unsafe.Pointer, ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::actionTriggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionTriggered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QToolBar) ActionTriggered(action QAction_ITF) {
	defer qt.Recovering("QToolBar::actionTriggered")

	if ptr.Pointer() != nil {
		C.QToolBar_ActionTriggered(ptr.Pointer(), PointerFromQAction(action))
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
func callbackQToolBarAllowedAreasChanged(ptr unsafe.Pointer, ptrName *C.char, allowedAreas C.int) {
	defer qt.Recovering("callback QToolBar::allowedAreasChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged"); signal != nil {
		signal.(func(core.Qt__ToolBarArea))(core.Qt__ToolBarArea(allowedAreas))
	}

}

func (ptr *QToolBar) AllowedAreasChanged(allowedAreas core.Qt__ToolBarArea) {
	defer qt.Recovering("QToolBar::allowedAreasChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_AllowedAreasChanged(ptr.Pointer(), C.int(allowedAreas))
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
func callbackQToolBarChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBar) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) Clear() {
	defer qt.Recovering("QToolBar::clear")

	if ptr.Pointer() != nil {
		C.QToolBar_Clear(ptr.Pointer())
	}
}

func (ptr *QToolBar) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QToolBar::event")

	if ptr.Pointer() != nil {
		return C.QToolBar_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QToolBar) ConnectIconSizeChanged(f func(iconSize *core.QSize)) {
	defer qt.Recovering("connect QToolBar::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectIconSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconSizeChanged", f)
	}
}

func (ptr *QToolBar) DisconnectIconSizeChanged() {
	defer qt.Recovering("disconnect QToolBar::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectIconSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconSizeChanged")
	}
}

//export callbackQToolBarIconSizeChanged
func callbackQToolBarIconSizeChanged(ptr unsafe.Pointer, ptrName *C.char, iconSize unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::iconSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(iconSize))
	}

}

func (ptr *QToolBar) IconSizeChanged(iconSize core.QSize_ITF) {
	defer qt.Recovering("QToolBar::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_IconSizeChanged(ptr.Pointer(), core.PointerFromQSize(iconSize))
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
func callbackQToolBarMovableChanged(ptr unsafe.Pointer, ptrName *C.char, movable C.int) {
	defer qt.Recovering("callback QToolBar::movableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "movableChanged"); signal != nil {
		signal.(func(bool))(int(movable) != 0)
	}

}

func (ptr *QToolBar) MovableChanged(movable bool) {
	defer qt.Recovering("QToolBar::movableChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_MovableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
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
func callbackQToolBarOrientationChanged(ptr unsafe.Pointer, ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QToolBar::orientationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "orientationChanged"); signal != nil {
		signal.(func(core.Qt__Orientation))(core.Qt__Orientation(orientation))
	}

}

func (ptr *QToolBar) OrientationChanged(orientation core.Qt__Orientation) {
	defer qt.Recovering("QToolBar::orientationChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_OrientationChanged(ptr.Pointer(), C.int(orientation))
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
func callbackQToolBarPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QToolBar) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QToolBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QToolBar) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QToolBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQToolBarToolButtonStyleChanged(ptr unsafe.Pointer, ptrName *C.char, toolButtonStyle C.int) {
	defer qt.Recovering("callback QToolBar::toolButtonStyleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged"); signal != nil {
		signal.(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

func (ptr *QToolBar) ToolButtonStyleChanged(toolButtonStyle core.Qt__ToolButtonStyle) {
	defer qt.Recovering("QToolBar::toolButtonStyleChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_ToolButtonStyleChanged(ptr.Pointer(), C.int(toolButtonStyle))
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
func callbackQToolBarTopLevelChanged(ptr unsafe.Pointer, ptrName *C.char, topLevel C.int) {
	defer qt.Recovering("callback QToolBar::topLevelChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "topLevelChanged"); signal != nil {
		signal.(func(bool))(int(topLevel) != 0)
	}

}

func (ptr *QToolBar) TopLevelChanged(topLevel bool) {
	defer qt.Recovering("QToolBar::topLevelChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_TopLevelChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(topLevel)))
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
func callbackQToolBarVisibilityChanged(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QToolBar::visibilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "visibilityChanged"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	}

}

func (ptr *QToolBar) VisibilityChanged(visible bool) {
	defer qt.Recovering("QToolBar::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QToolBar_VisibilityChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
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

func (ptr *QToolBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QToolBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QToolBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QToolBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQToolBarDragEnterEvent
func callbackQToolBarDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QToolBar) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QToolBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QToolBar) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QToolBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QToolBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QToolBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QToolBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QToolBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQToolBarDragLeaveEvent
func callbackQToolBarDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QToolBar) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QToolBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QToolBar) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QToolBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QToolBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QToolBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QToolBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QToolBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQToolBarDragMoveEvent
func callbackQToolBarDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QToolBar) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QToolBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QToolBar) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QToolBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QToolBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QToolBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QToolBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QToolBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQToolBarDropEvent
func callbackQToolBarDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QToolBar) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QToolBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QToolBar) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QToolBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QToolBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QToolBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QToolBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQToolBarEnterEvent
func callbackQToolBarEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBar) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QToolBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QToolBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQToolBarFocusInEvent
func callbackQToolBarFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QToolBar) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBar) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBar) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QToolBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QToolBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQToolBarFocusOutEvent
func callbackQToolBarFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QToolBar) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBar) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBar) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QToolBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QToolBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QToolBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQToolBarHideEvent
func callbackQToolBarHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QToolBar) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QToolBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QToolBar) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QToolBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QToolBar) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QToolBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QToolBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQToolBarLeaveEvent
func callbackQToolBarLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBar) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QToolBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QToolBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QToolBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQToolBarMoveEvent
func callbackQToolBarMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QToolBar) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QToolBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QToolBar) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QToolBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QToolBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QToolBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QToolBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QToolBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQToolBarSetVisible
func callbackQToolBarSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QToolBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QToolBar) SetVisible(visible bool) {
	defer qt.Recovering("QToolBar::setVisible")

	if ptr.Pointer() != nil {
		C.QToolBar_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QToolBar) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QToolBar::setVisible")

	if ptr.Pointer() != nil {
		C.QToolBar_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QToolBar) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QToolBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QToolBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QToolBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQToolBarShowEvent
func callbackQToolBarShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QToolBar) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QToolBar::showEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QToolBar) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QToolBar::showEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QToolBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QToolBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QToolBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QToolBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQToolBarCloseEvent
func callbackQToolBarCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QToolBar) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QToolBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QToolBar) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QToolBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QToolBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QToolBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QToolBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QToolBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQToolBarContextMenuEvent
func callbackQToolBarContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QToolBar) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QToolBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QToolBar) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QToolBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QToolBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QToolBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QToolBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QToolBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQToolBarInitPainter
func callbackQToolBarInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQToolBarFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QToolBar) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QToolBar::initPainter")

	if ptr.Pointer() != nil {
		C.QToolBar_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QToolBar) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QToolBar::initPainter")

	if ptr.Pointer() != nil {
		C.QToolBar_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QToolBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QToolBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QToolBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QToolBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQToolBarInputMethodEvent
func callbackQToolBarInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QToolBar) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QToolBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QToolBar) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QToolBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QToolBar) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QToolBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QToolBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQToolBarKeyPressEvent
func callbackQToolBarKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QToolBar) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBar) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QToolBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QToolBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQToolBarKeyReleaseEvent
func callbackQToolBarKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QToolBar) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBar) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QToolBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QToolBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQToolBarMouseDoubleClickEvent
func callbackQToolBarMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBar) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QToolBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QToolBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQToolBarMouseMoveEvent
func callbackQToolBarMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBar) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QToolBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QToolBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQToolBarMousePressEvent
func callbackQToolBarMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBar) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QToolBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QToolBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQToolBarMouseReleaseEvent
func callbackQToolBarMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBar) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBar) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QToolBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QToolBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QToolBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQToolBarResizeEvent
func callbackQToolBarResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QToolBar) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QToolBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QToolBar) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QToolBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QToolBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QToolBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QToolBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QToolBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQToolBarTabletEvent
func callbackQToolBarTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QToolBar) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QToolBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QToolBar) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QToolBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QToolBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QToolBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QToolBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QToolBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQToolBarWheelEvent
func callbackQToolBarWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QToolBar) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QToolBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QToolBar) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QToolBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QToolBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QToolBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QToolBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QToolBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQToolBarTimerEvent
func callbackQToolBarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QToolBar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QToolBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QToolBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QToolBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QToolBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QToolBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QToolBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQToolBarChildEvent
func callbackQToolBarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QToolBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QToolBar::childEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QToolBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QToolBar::childEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QToolBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QToolBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QToolBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQToolBarCustomEvent
func callbackQToolBarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::customEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBar::customEvent")

	if ptr.Pointer() != nil {
		C.QToolBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
