package widgets

//#include "qtoolbar.h"
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

type QToolBarITF interface {
	QWidgetITF
	QToolBarPTR() *QToolBar
}

func PointerFromQToolBar(ptr QToolBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolBarPTR().Pointer()
	}
	return nil
}

func QToolBarFromPointer(ptr unsafe.Pointer) *QToolBar {
	var n = new(QToolBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QToolBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QToolBar) QToolBarPTR() *QToolBar {
	return ptr
}

func (ptr *QToolBar) AllowedAreas() core.Qt__ToolBarArea {
	if ptr.Pointer() != nil {
		return core.Qt__ToolBarArea(C.QToolBar_AllowedAreas(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QToolBar) IsFloatable() bool {
	if ptr.Pointer() != nil {
		return C.QToolBar_IsFloatable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QToolBar) IsFloating() bool {
	if ptr.Pointer() != nil {
		return C.QToolBar_IsFloating(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QToolBar) IsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QToolBar_IsMovable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QToolBar) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QToolBar_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QToolBar) SetAllowedAreas(areas core.Qt__ToolBarArea) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetAllowedAreas(C.QtObjectPtr(ptr.Pointer()), C.int(areas))
	}
}

func (ptr *QToolBar) SetFloatable(floatable bool) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetFloatable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(floatable)))
	}
}

func (ptr *QToolBar) SetIconSize(iconSize core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(iconSize)))
	}
}

func (ptr *QToolBar) SetMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetMovable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QToolBar) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}

func (ptr *QToolBar) SetToolButtonStyle(toolButtonStyle core.Qt__ToolButtonStyle) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetToolButtonStyle(C.QtObjectPtr(ptr.Pointer()), C.int(toolButtonStyle))
	}
}

func (ptr *QToolBar) ToolButtonStyle() core.Qt__ToolButtonStyle {
	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolBar_ToolButtonStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQToolBar2(parent QWidgetITF) *QToolBar {
	return QToolBarFromPointer(unsafe.Pointer(C.QToolBar_NewQToolBar2(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQToolBar(title string, parent QWidgetITF) *QToolBar {
	return QToolBarFromPointer(unsafe.Pointer(C.QToolBar_NewQToolBar(C.CString(title), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QToolBar) ActionAt(p core.QPointITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_ActionAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)))))
	}
	return nil
}

func (ptr *QToolBar) ActionAt2(x int, y int) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_ActionAt2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))))
	}
	return nil
}

func (ptr *QToolBar) ConnectActionTriggered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectActionTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QToolBar) DisconnectActionTriggered() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectActionTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQToolBarActionTriggered
func callbackQToolBarActionTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "actionTriggered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QToolBar) AddAction2(icon gui.QIconITF, text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_AddAction2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QToolBar) AddAction4(icon gui.QIconITF, text string, receiver core.QObjectITF, member string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_AddAction4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))))
	}
	return nil
}

func (ptr *QToolBar) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_AddAction(C.QtObjectPtr(ptr.Pointer()), C.CString(text))))
	}
	return nil
}

func (ptr *QToolBar) AddAction3(text string, receiver core.QObjectITF, member string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_AddAction3(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))))
	}
	return nil
}

func (ptr *QToolBar) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_AddSeparator(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QToolBar) AddWidget(widget QWidgetITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))))
	}
	return nil
}

func (ptr *QToolBar) ConnectAllowedAreasChanged(f func(allowedAreas core.Qt__ToolBarArea)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectAllowedAreasChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "allowedAreasChanged", f)
	}
}

func (ptr *QToolBar) DisconnectAllowedAreasChanged() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectAllowedAreasChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "allowedAreasChanged")
	}
}

//export callbackQToolBarAllowedAreasChanged
func callbackQToolBarAllowedAreasChanged(ptrName *C.char, allowedAreas C.int) {
	qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged").(func(core.Qt__ToolBarArea))(core.Qt__ToolBarArea(allowedAreas))
}

func (ptr *QToolBar) Clear() {
	if ptr.Pointer() != nil {
		C.QToolBar_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QToolBar) InsertSeparator(before QActionITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_InsertSeparator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)))))
	}
	return nil
}

func (ptr *QToolBar) InsertWidget(before QActionITF, widget QWidgetITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_InsertWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.QtObjectPtr(PointerFromQWidget(widget)))))
	}
	return nil
}

func (ptr *QToolBar) IsAreaAllowed(area core.Qt__ToolBarArea) bool {
	if ptr.Pointer() != nil {
		return C.QToolBar_IsAreaAllowed(C.QtObjectPtr(ptr.Pointer()), C.int(area)) != 0
	}
	return false
}

func (ptr *QToolBar) ConnectMovableChanged(f func(movable bool)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectMovableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "movableChanged", f)
	}
}

func (ptr *QToolBar) DisconnectMovableChanged() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectMovableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "movableChanged")
	}
}

//export callbackQToolBarMovableChanged
func callbackQToolBarMovableChanged(ptrName *C.char, movable C.int) {
	qt.GetSignal(C.GoString(ptrName), "movableChanged").(func(bool))(int(movable) != 0)
}

func (ptr *QToolBar) ConnectOrientationChanged(f func(orientation core.Qt__Orientation)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "orientationChanged", f)
	}
}

func (ptr *QToolBar) DisconnectOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectOrientationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "orientationChanged")
	}
}

//export callbackQToolBarOrientationChanged
func callbackQToolBarOrientationChanged(ptrName *C.char, orientation C.int) {
	qt.GetSignal(C.GoString(ptrName), "orientationChanged").(func(core.Qt__Orientation))(core.Qt__Orientation(orientation))
}

func (ptr *QToolBar) ToggleViewAction() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QToolBar_ToggleViewAction(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QToolBar) ConnectToolButtonStyleChanged(f func(toolButtonStyle core.Qt__ToolButtonStyle)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectToolButtonStyleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "toolButtonStyleChanged", f)
	}
}

func (ptr *QToolBar) DisconnectToolButtonStyleChanged() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectToolButtonStyleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "toolButtonStyleChanged")
	}
}

//export callbackQToolBarToolButtonStyleChanged
func callbackQToolBarToolButtonStyleChanged(ptrName *C.char, toolButtonStyle C.int) {
	qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged").(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
}

func (ptr *QToolBar) ConnectTopLevelChanged(f func(topLevel bool)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectTopLevelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "topLevelChanged", f)
	}
}

func (ptr *QToolBar) DisconnectTopLevelChanged() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectTopLevelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "topLevelChanged")
	}
}

//export callbackQToolBarTopLevelChanged
func callbackQToolBarTopLevelChanged(ptrName *C.char, topLevel C.int) {
	qt.GetSignal(C.GoString(ptrName), "topLevelChanged").(func(bool))(int(topLevel) != 0)
}

func (ptr *QToolBar) ConnectVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QToolBar_ConnectVisibilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QToolBar) DisconnectVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QToolBar_DisconnectVisibilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQToolBarVisibilityChanged
func callbackQToolBarVisibilityChanged(ptrName *C.char, visible C.int) {
	qt.GetSignal(C.GoString(ptrName), "visibilityChanged").(func(bool))(int(visible) != 0)
}

func (ptr *QToolBar) WidgetForAction(action QActionITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QToolBar_WidgetForAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))))
	}
	return nil
}

func (ptr *QToolBar) DestroyQToolBar() {
	if ptr.Pointer() != nil {
		C.QToolBar_DestroyQToolBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
