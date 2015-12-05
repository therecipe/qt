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
		n.SetObjectName("QToolBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QToolBar) QToolBar_PTR() *QToolBar {
	return ptr
}

func (ptr *QToolBar) AllowedAreas() core.Qt__ToolBarArea {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::allowedAreas")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ToolBarArea(C.QToolBar_AllowedAreas(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBar) IsFloatable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::isFloatable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QToolBar_IsFloatable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolBar) IsFloating() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::isFloating")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QToolBar_IsFloating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolBar) IsMovable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::isMovable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QToolBar_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolBar) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QToolBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBar) SetAllowedAreas(areas core.Qt__ToolBarArea) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::setAllowedAreas")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_SetAllowedAreas(ptr.Pointer(), C.int(areas))
	}
}

func (ptr *QToolBar) SetFloatable(floatable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::setFloatable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_SetFloatable(ptr.Pointer(), C.int(qt.GoBoolToInt(floatable)))
	}
}

func (ptr *QToolBar) SetIconSize(iconSize core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::setIconSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_SetIconSize(ptr.Pointer(), core.PointerFromQSize(iconSize))
	}
}

func (ptr *QToolBar) SetMovable(movable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::setMovable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QToolBar) SetOrientation(orientation core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QToolBar) SetToolButtonStyle(toolButtonStyle core.Qt__ToolButtonStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::setToolButtonStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_SetToolButtonStyle(ptr.Pointer(), C.int(toolButtonStyle))
	}
}

func (ptr *QToolBar) ToolButtonStyle() core.Qt__ToolButtonStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::toolButtonStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolBar_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func NewQToolBar2(parent QWidget_ITF) *QToolBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::QToolBar")
		}
	}()

	return NewQToolBarFromPointer(C.QToolBar_NewQToolBar2(PointerFromQWidget(parent)))
}

func NewQToolBar(title string, parent QWidget_ITF) *QToolBar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::QToolBar")
		}
	}()

	return NewQToolBarFromPointer(C.QToolBar_NewQToolBar(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QToolBar) ActionAt(p core.QPoint_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::actionAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_ActionAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QToolBar) ActionAt2(x int, y int) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::actionAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_ActionAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QToolBar) ConnectActionTriggered(f func(action *QAction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::actionTriggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectActionTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QToolBar) DisconnectActionTriggered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::actionTriggered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectActionTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQToolBarActionTriggered
func callbackQToolBarActionTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::actionTriggered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "actionTriggered").(func(*QAction))(NewQActionFromPointer(action))
}

func (ptr *QToolBar) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QToolBar) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QToolBar) AddAction(text string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QToolBar) AddAction3(text string, receiver core.QObject_ITF, member string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddAction3(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QToolBar) AddSeparator() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::addSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBar) AddWidget(widget QWidget_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QToolBar) ConnectAllowedAreasChanged(f func(allowedAreas core.Qt__ToolBarArea)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::allowedAreasChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectAllowedAreasChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "allowedAreasChanged", f)
	}
}

func (ptr *QToolBar) DisconnectAllowedAreasChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::allowedAreasChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectAllowedAreasChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "allowedAreasChanged")
	}
}

//export callbackQToolBarAllowedAreasChanged
func callbackQToolBarAllowedAreasChanged(ptrName *C.char, allowedAreas C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::allowedAreasChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged").(func(core.Qt__ToolBarArea))(core.Qt__ToolBarArea(allowedAreas))
}

func (ptr *QToolBar) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_Clear(ptr.Pointer())
	}
}

func (ptr *QToolBar) InsertSeparator(before QAction_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::insertSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QToolBar) InsertWidget(before QAction_ITF, widget QWidget_ITF) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::insertWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_InsertWidget(ptr.Pointer(), PointerFromQAction(before), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QToolBar) IsAreaAllowed(area core.Qt__ToolBarArea) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::isAreaAllowed")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QToolBar_IsAreaAllowed(ptr.Pointer(), C.int(area)) != 0
	}
	return false
}

func (ptr *QToolBar) ConnectMovableChanged(f func(movable bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::movableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectMovableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "movableChanged", f)
	}
}

func (ptr *QToolBar) DisconnectMovableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::movableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectMovableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "movableChanged")
	}
}

//export callbackQToolBarMovableChanged
func callbackQToolBarMovableChanged(ptrName *C.char, movable C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::movableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "movableChanged").(func(bool))(int(movable) != 0)
}

func (ptr *QToolBar) ConnectOrientationChanged(f func(orientation core.Qt__Orientation)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::orientationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "orientationChanged", f)
	}
}

func (ptr *QToolBar) DisconnectOrientationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::orientationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "orientationChanged")
	}
}

//export callbackQToolBarOrientationChanged
func callbackQToolBarOrientationChanged(ptrName *C.char, orientation C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::orientationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "orientationChanged").(func(core.Qt__Orientation))(core.Qt__Orientation(orientation))
}

func (ptr *QToolBar) ToggleViewAction() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::toggleViewAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QToolBar_ToggleViewAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBar) ConnectToolButtonStyleChanged(f func(toolButtonStyle core.Qt__ToolButtonStyle)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::toolButtonStyleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectToolButtonStyleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toolButtonStyleChanged", f)
	}
}

func (ptr *QToolBar) DisconnectToolButtonStyleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::toolButtonStyleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectToolButtonStyleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toolButtonStyleChanged")
	}
}

//export callbackQToolBarToolButtonStyleChanged
func callbackQToolBarToolButtonStyleChanged(ptrName *C.char, toolButtonStyle C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::toolButtonStyleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged").(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
}

func (ptr *QToolBar) ConnectTopLevelChanged(f func(topLevel bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::topLevelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectTopLevelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "topLevelChanged", f)
	}
}

func (ptr *QToolBar) DisconnectTopLevelChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::topLevelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectTopLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "topLevelChanged")
	}
}

//export callbackQToolBarTopLevelChanged
func callbackQToolBarTopLevelChanged(ptrName *C.char, topLevel C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::topLevelChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "topLevelChanged").(func(bool))(int(topLevel) != 0)
}

func (ptr *QToolBar) ConnectVisibilityChanged(f func(visible bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::visibilityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_ConnectVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QToolBar) DisconnectVisibilityChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::visibilityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQToolBarVisibilityChanged
func callbackQToolBarVisibilityChanged(ptrName *C.char, visible C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::visibilityChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "visibilityChanged").(func(bool))(int(visible) != 0)
}

func (ptr *QToolBar) WidgetForAction(action QAction_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::widgetForAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBar_WidgetForAction(ptr.Pointer(), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QToolBar) DestroyQToolBar() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QToolBar::~QToolBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QToolBar_DestroyQToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
