package qt

//#include "qtoolbar.h"
import "C"

type qtoolbar struct {
	qwidget
}

type QToolBar interface {
	QWidget
	AllowedAreas() ToolBarArea
	Clear()
	IsAreaAllowed(area ToolBarArea) bool
	IsFloatable() bool
	IsFloating() bool
	IsMovable() bool
	Orientation() Orientation
	SetAllowedAreas(areas ToolBarArea)
	SetFloatable(floatable bool)
	SetMovable(movable bool)
	SetOrientation(orientation Orientation)
	ToolButtonStyle() ToolButtonStyle
	ConnectSignalActionTriggered(f func())
	DisconnectSignalActionTriggered()
	SignalActionTriggered() func()
	ConnectSignalAllowedAreasChanged(f func())
	DisconnectSignalAllowedAreasChanged()
	SignalAllowedAreasChanged() func()
	ConnectSignalMovableChanged(f func())
	DisconnectSignalMovableChanged()
	SignalMovableChanged() func()
	ConnectSignalOrientationChanged(f func())
	DisconnectSignalOrientationChanged()
	SignalOrientationChanged() func()
	ConnectSignalToolButtonStyleChanged(f func())
	DisconnectSignalToolButtonStyleChanged()
	SignalToolButtonStyleChanged() func()
	ConnectSignalTopLevelChanged(f func())
	DisconnectSignalTopLevelChanged()
	SignalTopLevelChanged() func()
	ConnectSignalVisibilityChanged(f func())
	DisconnectSignalVisibilityChanged()
	SignalVisibilityChanged() func()
}

func (p *qtoolbar) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtoolbar) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQToolBar1(title string, parent QWidget) QToolBar {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtoolbar = new(qtoolbar)
	qtoolbar.SetPointer(C.QToolBar_New_String_QWidget(C.CString(title), parentPtr))
	qtoolbar.SetObjectName("QToolBar_" + randomIdentifier())
	return qtoolbar
}

func NewQToolBar2(parent QWidget) QToolBar {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtoolbar = new(qtoolbar)
	qtoolbar.SetPointer(C.QToolBar_New_QWidget(parentPtr))
	qtoolbar.SetObjectName("QToolBar_" + randomIdentifier())
	return qtoolbar
}

func (p *qtoolbar) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QToolBar_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtoolbar) AllowedAreas() ToolBarArea {
	if p.Pointer() == nil {
		return 0
	}
	return ToolBarArea(C.QToolBar_AllowedAreas(p.Pointer()))
}

func (p *qtoolbar) Clear() {
	if p.Pointer() != nil {
		C.QToolBar_Clear(p.Pointer())
	}
}

func (p *qtoolbar) IsAreaAllowed(area ToolBarArea) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QToolBar_IsAreaAllowed_ToolBarArea(p.Pointer(), C.int(area)) != 0
}

func (p *qtoolbar) IsFloatable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QToolBar_IsFloatable(p.Pointer()) != 0
}

func (p *qtoolbar) IsFloating() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QToolBar_IsFloating(p.Pointer()) != 0
}

func (p *qtoolbar) IsMovable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QToolBar_IsMovable(p.Pointer()) != 0
}

func (p *qtoolbar) Orientation() Orientation {
	if p.Pointer() == nil {
		return 0
	}
	return Orientation(C.QToolBar_Orientation(p.Pointer()))
}

func (p *qtoolbar) SetAllowedAreas(areas ToolBarArea) {
	if p.Pointer() != nil {
		C.QToolBar_SetAllowedAreas_ToolBarArea(p.Pointer(), C.int(areas))
	}
}

func (p *qtoolbar) SetFloatable(floatable bool) {
	if p.Pointer() != nil {
		C.QToolBar_SetFloatable_Bool(p.Pointer(), goBoolToCInt(floatable))
	}
}

func (p *qtoolbar) SetMovable(movable bool) {
	if p.Pointer() != nil {
		C.QToolBar_SetMovable_Bool(p.Pointer(), goBoolToCInt(movable))
	}
}

func (p *qtoolbar) SetOrientation(orientation Orientation) {
	if p.Pointer() != nil {
		C.QToolBar_SetOrientation_Orientation(p.Pointer(), C.int(orientation))
	}
}

func (p *qtoolbar) ToolButtonStyle() ToolButtonStyle {
	if p.Pointer() == nil {
		return 0
	}
	return ToolButtonStyle(C.QToolBar_ToolButtonStyle(p.Pointer()))
}

func (p *qtoolbar) ConnectSignalActionTriggered(f func()) {
	C.QToolBar_ConnectSignalActionTriggered(p.Pointer())
	connectSignal(p.ObjectName(), "actionTriggered", f)
}

func (p *qtoolbar) DisconnectSignalActionTriggered() {
	C.QToolBar_DisconnectSignalActionTriggered(p.Pointer())
	disconnectSignal(p.ObjectName(), "actionTriggered")
}

func (p *qtoolbar) SignalActionTriggered() func() {
	return getSignal(p.ObjectName(), "actionTriggered")
}

func (p *qtoolbar) ConnectSignalAllowedAreasChanged(f func()) {
	C.QToolBar_ConnectSignalAllowedAreasChanged(p.Pointer())
	connectSignal(p.ObjectName(), "allowedAreasChanged", f)
}

func (p *qtoolbar) DisconnectSignalAllowedAreasChanged() {
	C.QToolBar_DisconnectSignalAllowedAreasChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "allowedAreasChanged")
}

func (p *qtoolbar) SignalAllowedAreasChanged() func() {
	return getSignal(p.ObjectName(), "allowedAreasChanged")
}

func (p *qtoolbar) ConnectSignalMovableChanged(f func()) {
	C.QToolBar_ConnectSignalMovableChanged(p.Pointer())
	connectSignal(p.ObjectName(), "movableChanged", f)
}

func (p *qtoolbar) DisconnectSignalMovableChanged() {
	C.QToolBar_DisconnectSignalMovableChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "movableChanged")
}

func (p *qtoolbar) SignalMovableChanged() func() {
	return getSignal(p.ObjectName(), "movableChanged")
}

func (p *qtoolbar) ConnectSignalOrientationChanged(f func()) {
	C.QToolBar_ConnectSignalOrientationChanged(p.Pointer())
	connectSignal(p.ObjectName(), "orientationChanged", f)
}

func (p *qtoolbar) DisconnectSignalOrientationChanged() {
	C.QToolBar_DisconnectSignalOrientationChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "orientationChanged")
}

func (p *qtoolbar) SignalOrientationChanged() func() {
	return getSignal(p.ObjectName(), "orientationChanged")
}

func (p *qtoolbar) ConnectSignalToolButtonStyleChanged(f func()) {
	C.QToolBar_ConnectSignalToolButtonStyleChanged(p.Pointer())
	connectSignal(p.ObjectName(), "toolButtonStyleChanged", f)
}

func (p *qtoolbar) DisconnectSignalToolButtonStyleChanged() {
	C.QToolBar_DisconnectSignalToolButtonStyleChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "toolButtonStyleChanged")
}

func (p *qtoolbar) SignalToolButtonStyleChanged() func() {
	return getSignal(p.ObjectName(), "toolButtonStyleChanged")
}

func (p *qtoolbar) ConnectSignalTopLevelChanged(f func()) {
	C.QToolBar_ConnectSignalTopLevelChanged(p.Pointer())
	connectSignal(p.ObjectName(), "topLevelChanged", f)
}

func (p *qtoolbar) DisconnectSignalTopLevelChanged() {
	C.QToolBar_DisconnectSignalTopLevelChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "topLevelChanged")
}

func (p *qtoolbar) SignalTopLevelChanged() func() {
	return getSignal(p.ObjectName(), "topLevelChanged")
}

func (p *qtoolbar) ConnectSignalVisibilityChanged(f func()) {
	C.QToolBar_ConnectSignalVisibilityChanged(p.Pointer())
	connectSignal(p.ObjectName(), "visibilityChanged", f)
}

func (p *qtoolbar) DisconnectSignalVisibilityChanged() {
	C.QToolBar_DisconnectSignalVisibilityChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "visibilityChanged")
}

func (p *qtoolbar) SignalVisibilityChanged() func() {
	return getSignal(p.ObjectName(), "visibilityChanged")
}

//export callbackQToolBar
func callbackQToolBar(ptr C.QtObjectPtr, msg *C.char) {
	var qtoolbar = new(qtoolbar)
	qtoolbar.SetPointer(ptr)
	getSignal(qtoolbar.ObjectName(), C.GoString(msg))()
}
