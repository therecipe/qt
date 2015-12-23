package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"unsafe"
)

type QQuickItem struct {
	core.QObject
	qml.QQmlParserStatus
}

type QQuickItem_ITF interface {
	core.QObject_ITF
	qml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (p *QQuickItem) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QQuickItem) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QQmlParserStatus_PTR().SetPointer(ptr)
}

func PointerFromQQuickItem(ptr QQuickItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemFromPointer(ptr unsafe.Pointer) *QQuickItem {
	var n = new(QQuickItem)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickItem_") {
		n.SetObjectName("QQuickItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickItem) QQuickItem_PTR() *QQuickItem {
	return ptr
}

//QQuickItem::Flag
type QQuickItem__Flag int64

const (
	QQuickItem__ItemClipsChildrenToShape = QQuickItem__Flag(0x01)
	QQuickItem__ItemAcceptsInputMethod   = QQuickItem__Flag(0x02)
	QQuickItem__ItemIsFocusScope         = QQuickItem__Flag(0x04)
	QQuickItem__ItemHasContents          = QQuickItem__Flag(0x08)
	QQuickItem__ItemAcceptsDrops         = QQuickItem__Flag(0x10)
)

//QQuickItem::ItemChange
type QQuickItem__ItemChange int64

const (
	QQuickItem__ItemChildAddedChange       = QQuickItem__ItemChange(0)
	QQuickItem__ItemChildRemovedChange     = QQuickItem__ItemChange(1)
	QQuickItem__ItemSceneChange            = QQuickItem__ItemChange(2)
	QQuickItem__ItemVisibleHasChanged      = QQuickItem__ItemChange(3)
	QQuickItem__ItemParentHasChanged       = QQuickItem__ItemChange(4)
	QQuickItem__ItemOpacityHasChanged      = QQuickItem__ItemChange(5)
	QQuickItem__ItemActiveFocusHasChanged  = QQuickItem__ItemChange(6)
	QQuickItem__ItemRotationHasChanged     = QQuickItem__ItemChange(7)
	QQuickItem__ItemAntialiasingHasChanged = QQuickItem__ItemChange(8)
)

//QQuickItem::TransformOrigin
type QQuickItem__TransformOrigin int64

const (
	QQuickItem__TopLeft     = QQuickItem__TransformOrigin(0)
	QQuickItem__Top         = QQuickItem__TransformOrigin(1)
	QQuickItem__TopRight    = QQuickItem__TransformOrigin(2)
	QQuickItem__Left        = QQuickItem__TransformOrigin(3)
	QQuickItem__Center      = QQuickItem__TransformOrigin(4)
	QQuickItem__Right       = QQuickItem__TransformOrigin(5)
	QQuickItem__BottomLeft  = QQuickItem__TransformOrigin(6)
	QQuickItem__Bottom      = QQuickItem__TransformOrigin(7)
	QQuickItem__BottomRight = QQuickItem__TransformOrigin(8)
)

func NewQQuickItem(parent QQuickItem_ITF) *QQuickItem {
	defer qt.Recovering("QQuickItem::QQuickItem")

	return NewQQuickItemFromPointer(C.QQuickItem_NewQQuickItem(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {
	defer qt.Recovering("QQuickItem::activeFocusOnTab")

	if ptr.Pointer() != nil {
		return C.QQuickItem_ActiveFocusOnTab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Antialiasing() bool {
	defer qt.Recovering("QQuickItem::antialiasing")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) BaselineOffset() float64 {
	defer qt.Recovering("QQuickItem::baselineOffset")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_BaselineOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Clip() bool {
	defer qt.Recovering("QQuickItem::clip")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Clip(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasActiveFocus() bool {
	defer qt.Recovering("QQuickItem::hasActiveFocus")

	if ptr.Pointer() != nil {
		return C.QQuickItem_HasActiveFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasFocus() bool {
	defer qt.Recovering("QQuickItem::hasFocus")

	if ptr.Pointer() != nil {
		return C.QQuickItem_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Height() float64 {
	defer qt.Recovering("QQuickItem::height")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ImplicitHeight() float64 {
	defer qt.Recovering("QQuickItem::implicitHeight")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) IsEnabled() bool {
	defer qt.Recovering("QQuickItem::isEnabled")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	defer qt.Recovering("QQuickItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsVisible() bool {
	defer qt.Recovering("QQuickItem::isVisible")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Opacity() float64 {
	defer qt.Recovering("QQuickItem::opacity")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {
	defer qt.Recovering("QQuickItem::parentItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ParentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ResetAntialiasing() {
	defer qt.Recovering("QQuickItem::resetAntialiasing")

	if ptr.Pointer() != nil {
		C.QQuickItem_ResetAntialiasing(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetHeight() {
	defer qt.Recovering("QQuickItem::resetHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_ResetHeight(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetWidth() {
	defer qt.Recovering("QQuickItem::resetWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_ResetWidth(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Rotation() float64 {
	defer qt.Recovering("QQuickItem::rotation")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Rotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Scale() float64 {
	defer qt.Recovering("QQuickItem::scale")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Scale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) SetActiveFocusOnTab(v bool) {
	defer qt.Recovering("QQuickItem::setActiveFocusOnTab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetActiveFocusOnTab(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetAntialiasing(v bool) {
	defer qt.Recovering("QQuickItem::setAntialiasing")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAntialiasing(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetBaselineOffset(v float64) {
	defer qt.Recovering("QQuickItem::setBaselineOffset")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetBaselineOffset(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetClip(v bool) {
	defer qt.Recovering("QQuickItem::setClip")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetClip(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetEnabled(v bool) {
	defer qt.Recovering("QQuickItem::setEnabled")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus(v bool) {
	defer qt.Recovering("QQuickItem::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {
	defer qt.Recovering("QQuickItem::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus2(ptr.Pointer(), C.int(qt.GoBoolToInt(focus)), C.int(reason))
	}
}

func (ptr *QQuickItem) SetHeight(v float64) {
	defer qt.Recovering("QQuickItem::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetHeight(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetImplicitHeight(v float64) {
	defer qt.Recovering("QQuickItem::setImplicitHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitHeight(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetImplicitWidth(v float64) {
	defer qt.Recovering("QQuickItem::setImplicitWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitWidth(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetOpacity(v float64) {
	defer qt.Recovering("QQuickItem::setOpacity")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetOpacity(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::setParentItem")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetParentItem(ptr.Pointer(), PointerFromQQuickItem(parent))
	}
}

func (ptr *QQuickItem) SetRotation(v float64) {
	defer qt.Recovering("QQuickItem::setRotation")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetRotation(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetScale(v float64) {
	defer qt.Recovering("QQuickItem::setScale")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetSmooth(v bool) {
	defer qt.Recovering("QQuickItem::setSmooth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetState(v string) {
	defer qt.Recovering("QQuickItem::setState")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetState(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QQuickItem) SetTransformOrigin(v QQuickItem__TransformOrigin) {
	defer qt.Recovering("QQuickItem::setTransformOrigin")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickItem) SetVisible(v bool) {
	defer qt.Recovering("QQuickItem::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetWidth(v float64) {
	defer qt.Recovering("QQuickItem::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetWidth(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetX(v float64) {
	defer qt.Recovering("QQuickItem::setX")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetX(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetY(v float64) {
	defer qt.Recovering("QQuickItem::setY")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetY(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetZ(v float64) {
	defer qt.Recovering("QQuickItem::setZ")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetZ(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) Smooth() bool {
	defer qt.Recovering("QQuickItem::smooth")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Smooth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) State() string {
	defer qt.Recovering("QQuickItem::state")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickItem_State(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {
	defer qt.Recovering("QQuickItem::transformOrigin")

	if ptr.Pointer() != nil {
		return QQuickItem__TransformOrigin(C.QQuickItem_TransformOrigin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Width() float64 {
	defer qt.Recovering("QQuickItem::width")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) X() float64 {
	defer qt.Recovering("QQuickItem::x")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Y() float64 {
	defer qt.Recovering("QQuickItem::y")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Z() float64 {
	defer qt.Recovering("QQuickItem::z")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) AcceptHoverEvents() bool {
	defer qt.Recovering("QQuickItem::acceptHoverEvents")

	if ptr.Pointer() != nil {
		return C.QQuickItem_AcceptHoverEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {
	defer qt.Recovering("QQuickItem::acceptedMouseButtons")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QQuickItem_AcceptedMouseButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ChildAt(x float64, y float64) *QQuickItem {
	defer qt.Recovering("QQuickItem::childAt")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ChildAt(ptr.Pointer(), C.double(x), C.double(y)))
	}
	return nil
}

func (ptr *QQuickItem) ConnectClassBegin(f func()) {
	defer qt.Recovering("connect QQuickItem::classBegin")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "classBegin", f)
	}
}

func (ptr *QQuickItem) DisconnectClassBegin() {
	defer qt.Recovering("disconnect QQuickItem::classBegin")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "classBegin")
	}
}

//export callbackQQuickItemClassBegin
func callbackQQuickItemClassBegin(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickItem::classBegin")

	if signal := qt.GetSignal(C.GoString(ptrName), "classBegin"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectComponentComplete(f func()) {
	defer qt.Recovering("connect QQuickItem::componentComplete")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "componentComplete", f)
	}
}

func (ptr *QQuickItem) DisconnectComponentComplete() {
	defer qt.Recovering("disconnect QQuickItem::componentComplete")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "componentComplete")
	}
}

//export callbackQQuickItemComponentComplete
func callbackQQuickItemComponentComplete(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickItem::componentComplete")

	if signal := qt.GetSignal(C.GoString(ptrName), "componentComplete"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickItem::contains")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QQuickItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QQuickItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQQuickItemDragEnterEvent
func callbackQQuickItemDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QQuickItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QQuickItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQQuickItemDragLeaveEvent
func callbackQQuickItemDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QQuickItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QQuickItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQQuickItemDragMoveEvent
func callbackQQuickItemDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QQuickItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QQuickItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQQuickItemDropEvent
func callbackQQuickItemDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {
	defer qt.Recovering("QQuickItem::filtersChildMouseEvents")

	if ptr.Pointer() != nil {
		return C.QQuickItem_FiltersChildMouseEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {
	defer qt.Recovering("QQuickItem::flags")

	if ptr.Pointer() != nil {
		return QQuickItem__Flag(C.QQuickItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickItemFocusInEvent
func callbackQQuickItemFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickItemFocusOutEvent
func callbackQQuickItemFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ForceActiveFocus() {
	defer qt.Recovering("QQuickItem::forceActiveFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ForceActiveFocus2(reason core.Qt__FocusReason) {
	defer qt.Recovering("QQuickItem::forceActiveFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus2(ptr.Pointer(), C.int(reason))
	}
}

func (ptr *QQuickItem) GrabMouse() {
	defer qt.Recovering("QQuickItem::grabMouse")

	if ptr.Pointer() != nil {
		C.QQuickItem_GrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QQuickItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

//export callbackQQuickItemHoverEnterEvent
func callbackQQuickItemHoverEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QQuickItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQQuickItemHoverLeaveEvent
func callbackQQuickItemHoverLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QQuickItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQQuickItemHoverMoveEvent
func callbackQQuickItemHoverMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ImplicitWidth() float64 {
	defer qt.Recovering("QQuickItem::implicitWidth")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QQuickItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QQuickItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQQuickItemInputMethodEvent
func callbackQQuickItemInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickItem) IsFocusScope() bool {
	defer qt.Recovering("QQuickItem::isFocusScope")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsFocusScope(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepMouseGrab() bool {
	defer qt.Recovering("QQuickItem::keepMouseGrab")

	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepMouseGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepTouchGrab() bool {
	defer qt.Recovering("QQuickItem::keepTouchGrab")

	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepTouchGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickItemKeyPressEvent
func callbackQQuickItemKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickItemKeyReleaseEvent
func callbackQQuickItemKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickItemMouseDoubleClickEvent
func callbackQQuickItemMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickItemMouseMoveEvent
func callbackQQuickItemMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickItemMousePressEvent
func callbackQQuickItemMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickItemMouseReleaseEvent
func callbackQQuickItemMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectMouseUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseUngrabEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseUngrabEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseUngrabEvent")
	}
}

//export callbackQQuickItemMouseUngrabEvent
func callbackQQuickItemMouseUngrabEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickItem::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {
	defer qt.Recovering("QQuickItem::nextItemInFocusChain")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_NextItemInFocusChain(ptr.Pointer(), C.int(qt.GoBoolToInt(forward))))
	}
	return nil
}

func (ptr *QQuickItem) Polish() {
	defer qt.Recovering("QQuickItem::polish")

	if ptr.Pointer() != nil {
		C.QQuickItem_Polish(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickItem::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickItem) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickItem::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
	}
}

//export callbackQQuickItemReleaseResources
func callbackQQuickItemReleaseResources(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickItem::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {
	defer qt.Recovering("QQuickItem::scopedFocusItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ScopedFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	defer qt.Recovering("QQuickItem::setAcceptHoverEvents")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptHoverEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {
	defer qt.Recovering("QQuickItem::setAcceptedMouseButtons")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptedMouseButtons(ptr.Pointer(), C.int(buttons))
	}
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursor_ITF) {
	defer qt.Recovering("QQuickItem::setCursor")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(cursor))
	}
}

func (ptr *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	defer qt.Recovering("QQuickItem::setFiltersChildMouseEvents")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFiltersChildMouseEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(filter)))
	}
}

func (ptr *QQuickItem) SetFlag(flag QQuickItem__Flag, enabled bool) {
	defer qt.Recovering("QQuickItem::setFlag")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlag(ptr.Pointer(), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetFlags(flags QQuickItem__Flag) {
	defer qt.Recovering("QQuickItem::setFlags")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QQuickItem) SetKeepMouseGrab(keep bool) {
	defer qt.Recovering("QQuickItem::setKeepMouseGrab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepMouseGrab(ptr.Pointer(), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) SetKeepTouchGrab(keep bool) {
	defer qt.Recovering("QQuickItem::setKeepTouchGrab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepTouchGrab(ptr.Pointer(), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) StackAfter(sibling QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::stackAfter")

	if ptr.Pointer() != nil {
		C.QQuickItem_StackAfter(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) StackBefore(sibling QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::stackBefore")

	if ptr.Pointer() != nil {
		C.QQuickItem_StackBefore(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickItem::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickItem::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickItemTouchEvent
func callbackQQuickItemTouchEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectTouchUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickItem::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchUngrabEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTouchUngrabEvent() {
	defer qt.Recovering("disconnect QQuickItem::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchUngrabEvent")
	}
}

//export callbackQQuickItemTouchUngrabEvent
func callbackQQuickItemTouchUngrabEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickItem::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickItem) UngrabMouse() {
	defer qt.Recovering("QQuickItem::ungrabMouse")

	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UngrabTouchPoints() {
	defer qt.Recovering("QQuickItem::ungrabTouchPoints")

	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabTouchPoints(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UnsetCursor() {
	defer qt.Recovering("QQuickItem::unsetCursor")

	if ptr.Pointer() != nil {
		C.QQuickItem_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Update() {
	defer qt.Recovering("QQuickItem::update")

	if ptr.Pointer() != nil {
		C.QQuickItem_Update(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectUpdatePolish(f func()) {
	defer qt.Recovering("connect QQuickItem::updatePolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updatePolish", f)
	}
}

func (ptr *QQuickItem) DisconnectUpdatePolish() {
	defer qt.Recovering("disconnect QQuickItem::updatePolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updatePolish")
	}
}

//export callbackQQuickItemUpdatePolish
func callbackQQuickItemUpdatePolish(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickItem::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickItemWheelEvent
func callbackQQuickItemWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) Window() *QQuickWindow {
	defer qt.Recovering("QQuickItem::window")

	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickItem_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window *QQuickWindow)) {
	defer qt.Recovering("connect QQuickItem::windowChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectWindowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectWindowChanged() {
	defer qt.Recovering("disconnect QQuickItem::windowChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowChanged")
	}
}

//export callbackQQuickItemWindowChanged
func callbackQQuickItemWindowChanged(ptrName *C.char, window unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::windowChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowChanged"); signal != nil {
		signal.(func(*QQuickWindow))(NewQQuickWindowFromPointer(window))
	}

}

func (ptr *QQuickItem) DestroyQQuickItem() {
	defer qt.Recovering("QQuickItem::~QQuickItem")

	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickItemTimerEvent
func callbackQQuickItemTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickItemChildEvent
func callbackQQuickItemChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickItemCustomEvent
func callbackQQuickItemCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
