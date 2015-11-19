package quick

//#include "qquickitem.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickItem_" + qt.RandomIdentifier())
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
	return NewQQuickItemFromPointer(C.QQuickItem_NewQQuickItem(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_ActiveFocusOnTab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Antialiasing() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) BaselineOffset() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_BaselineOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Clip() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Clip(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasActiveFocus() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HasActiveFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ImplicitHeight() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Opacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ParentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ResetAntialiasing() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetAntialiasing(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetHeight() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetHeight(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetWidth() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetWidth(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Rotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Rotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Scale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Scale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) SetActiveFocusOnTab(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetActiveFocusOnTab(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetAntialiasing(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAntialiasing(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetBaselineOffset(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetBaselineOffset(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetClip(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetClip(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus2(ptr.Pointer(), C.int(qt.GoBoolToInt(focus)), C.int(reason))
	}
}

func (ptr *QQuickItem) SetHeight(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetHeight(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetImplicitHeight(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitHeight(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetImplicitWidth(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitWidth(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetOpacity(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetOpacity(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetParentItem(ptr.Pointer(), PointerFromQQuickItem(parent))
	}
}

func (ptr *QQuickItem) SetRotation(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetRotation(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetScale(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetSmooth(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetState(v string) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetState(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QQuickItem) SetTransformOrigin(v QQuickItem__TransformOrigin) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickItem) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetWidth(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetWidth(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetX(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetX(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetY(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetY(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetZ(v float64) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetZ(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) Smooth() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Smooth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {
	if ptr.Pointer() != nil {
		return QQuickItem__TransformOrigin(C.QQuickItem_TransformOrigin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) AcceptHoverEvents() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_AcceptHoverEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QQuickItem_AcceptedMouseButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ChildAt(x float64, y float64) *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ChildAt(ptr.Pointer(), C.double(x), C.double(y)))
	}
	return nil
}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_FiltersChildMouseEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {
	if ptr.Pointer() != nil {
		return QQuickItem__Flag(C.QQuickItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ForceActiveFocus() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ForceActiveFocus2(reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus2(ptr.Pointer(), C.int(reason))
	}
}

func (ptr *QQuickItem) GrabMouse() {
	if ptr.Pointer() != nil {
		C.QQuickItem_GrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ImplicitWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickItem) IsFocusScope() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsFocusScope(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepMouseGrab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepMouseGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepTouchGrab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepTouchGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_NextItemInFocusChain(ptr.Pointer(), C.int(qt.GoBoolToInt(forward))))
	}
	return nil
}

func (ptr *QQuickItem) Polish() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Polish(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ScopedFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptHoverEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptedMouseButtons(ptr.Pointer(), C.int(buttons))
	}
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(cursor))
	}
}

func (ptr *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFiltersChildMouseEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(filter)))
	}
}

func (ptr *QQuickItem) SetFlag(flag QQuickItem__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlag(ptr.Pointer(), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetFlags(flags QQuickItem__Flag) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QQuickItem) SetKeepMouseGrab(keep bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepMouseGrab(ptr.Pointer(), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) SetKeepTouchGrab(keep bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepTouchGrab(ptr.Pointer(), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) StackAfter(sibling QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_StackAfter(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) StackBefore(sibling QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_StackBefore(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) UngrabMouse() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UngrabTouchPoints() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabTouchPoints(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Update() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Update(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Window() *QQuickWindow {
	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickItem_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window *QQuickWindow)) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectWindowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectWindowChanged() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowChanged")
	}
}

//export callbackQQuickItemWindowChanged
func callbackQQuickItemWindowChanged(ptrName *C.char, window unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "windowChanged").(func(*QQuickWindow))(NewQQuickWindowFromPointer(window))
}

func (ptr *QQuickItem) DestroyQQuickItem() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
