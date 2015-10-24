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

type QQuickItemITF interface {
	core.QObjectITF
	qml.QQmlParserStatusITF
	QQuickItemPTR() *QQuickItem
}

func (p *QQuickItem) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QQuickItem) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QQmlParserStatusPTR().SetPointer(ptr)
}

func PointerFromQQuickItem(ptr QQuickItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemPTR().Pointer()
	}
	return nil
}

func QQuickItemFromPointer(ptr unsafe.Pointer) *QQuickItem {
	var n = new(QQuickItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickItem) QQuickItemPTR() *QQuickItem {
	return ptr
}

//QQuickItem::Flag
type QQuickItem__Flag int

var (
	QQuickItem__ItemClipsChildrenToShape = QQuickItem__Flag(0x01)
	QQuickItem__ItemAcceptsInputMethod   = QQuickItem__Flag(0x02)
	QQuickItem__ItemIsFocusScope         = QQuickItem__Flag(0x04)
	QQuickItem__ItemHasContents          = QQuickItem__Flag(0x08)
	QQuickItem__ItemAcceptsDrops         = QQuickItem__Flag(0x10)
)

//QQuickItem::ItemChange
type QQuickItem__ItemChange int

var (
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
type QQuickItem__TransformOrigin int

var (
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

func NewQQuickItem(parent QQuickItemITF) *QQuickItem {
	return QQuickItemFromPointer(unsafe.Pointer(C.QQuickItem_NewQQuickItem(C.QtObjectPtr(PointerFromQQuickItem(parent)))))
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_ActiveFocusOnTab(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) Antialiasing() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Antialiasing(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) Clip() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Clip(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) HasActiveFocus() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HasActiveFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_HasFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProvider(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickItem_ParentItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickItem) ResetAntialiasing() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetAntialiasing(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) ResetHeight() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetHeight(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) ResetWidth() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ResetWidth(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) SetActiveFocusOnTab(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetActiveFocusOnTab(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetAntialiasing(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAntialiasing(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetClip(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetClip(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus2(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(focus)), C.int(reason))
	}
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItemITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetParentItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQuickItem(parent)))
	}
}

func (ptr *QQuickItem) SetSmooth(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetState(v string) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetState(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QQuickItem) SetTransformOrigin(v QQuickItem__TransformOrigin) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QQuickItem) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) Smooth() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Smooth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		return QSGTextureProviderFromPointer(unsafe.Pointer(C.QQuickItem_TextureProvider(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {
	if ptr.Pointer() != nil {
		return QQuickItem__TransformOrigin(C.QQuickItem_TransformOrigin(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickItem) AcceptHoverEvents() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_AcceptHoverEvents(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QQuickItem_AcceptedMouseButtons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_FiltersChildMouseEvents(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {
	if ptr.Pointer() != nil {
		return QQuickItem__Flag(C.QQuickItem_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickItem) ForceActiveFocus() {
	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) ForceActiveFocus2(reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus2(C.QtObjectPtr(ptr.Pointer()), C.int(reason))
	}
}

func (ptr *QQuickItem) GrabMouse() {
	if ptr.Pointer() != nil {
		C.QQuickItem_GrabMouse(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickItem_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QQuickItem) IsFocusScope() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_IsFocusScope(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepMouseGrab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepMouseGrab(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepTouchGrab() bool {
	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepTouchGrab(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickItem_NextItemInFocusChain(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(forward)))))
	}
	return nil
}

func (ptr *QQuickItem) Polish() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Polish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {
	if ptr.Pointer() != nil {
		return QQuickItemFromPointer(unsafe.Pointer(C.QQuickItem_ScopedFocusItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptHoverEvents(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptedMouseButtons(C.QtObjectPtr(ptr.Pointer()), C.int(buttons))
	}
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursorITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQCursor(cursor)))
	}
}

func (ptr *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFiltersChildMouseEvents(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(filter)))
	}
}

func (ptr *QQuickItem) SetFlag(flag QQuickItem__Flag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlag(C.QtObjectPtr(ptr.Pointer()), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetFlags(flags QQuickItem__Flag) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QQuickItem) SetKeepMouseGrab(keep bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepMouseGrab(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) SetKeepTouchGrab(keep bool) {
	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepTouchGrab(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) StackAfter(sibling QQuickItemITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_StackAfter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQuickItem(sibling)))
	}
}

func (ptr *QQuickItem) StackBefore(sibling QQuickItemITF) {
	if ptr.Pointer() != nil {
		C.QQuickItem_StackBefore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQuickItem(sibling)))
	}
}

func (ptr *QQuickItem) UngrabMouse() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabMouse(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) UngrabTouchPoints() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabTouchPoints(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QQuickItem_UnsetCursor(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) Update() {
	if ptr.Pointer() != nil {
		C.QQuickItem_Update(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickItem) Window() *QQuickWindow {
	if ptr.Pointer() != nil {
		return QQuickWindowFromPointer(unsafe.Pointer(C.QQuickItem_Window(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window QQuickWindowITF)) {
	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectWindowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "windowChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectWindowChanged() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectWindowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "windowChanged")
	}
}

//export callbackQQuickItemWindowChanged
func callbackQQuickItemWindowChanged(ptrName *C.char, window unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "windowChanged").(func(*QQuickWindow))(QQuickWindowFromPointer(window))
}

func (ptr *QQuickItem) DestroyQQuickItem() {
	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItem(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
