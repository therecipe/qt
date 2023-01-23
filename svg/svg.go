// +build !minimal

package svg

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/widgets"
	"strings"
	"unsafe"
)

type QGraphicsSvgItem struct {
	widgets.QGraphicsObject
}

type QGraphicsSvgItem_ITF interface {
	widgets.QGraphicsObject_ITF
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem {
	return ptr
}

func (ptr *QGraphicsSvgItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsObject_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItem_PTR().Pointer()
	}
	return nil
}

func (n *QGraphicsSvgItem) InitFromInternal(ptr uintptr, name string) {
	n.QGraphicsObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGraphicsSvgItem) ClassNameInternalF() string {
	return n.QGraphicsObject_PTR().ClassNameInternalF()
}

func NewQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) (n *QGraphicsSvgItem) {
	n = new(QGraphicsSvgItem)
	n.InitFromInternal(uintptr(ptr), "svg.QGraphicsSvgItem")
	return
}

func (ptr *QGraphicsSvgItem) DestroyQGraphicsSvgItem() {
}

func NewQGraphicsSvgItem(parent widgets.QGraphicsItem_ITF) *QGraphicsSvgItem {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQGraphicsSvgItem", "", parent}).(*QGraphicsSvgItem)
}

func NewQGraphicsSvgItem2(fileName string, parent widgets.QGraphicsItem_ITF) *QGraphicsSvgItem {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQGraphicsSvgItem2", "", fileName, parent}).(*QGraphicsSvgItem)
}

func (ptr *QGraphicsSvgItem) ConnectBoundingRect(f func() *core.QRectF) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBoundingRect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGraphicsSvgItem) DisconnectBoundingRect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBoundingRect"})
}

func (ptr *QGraphicsSvgItem) BoundingRect() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundingRect"}).(*core.QRectF)
}

func (ptr *QGraphicsSvgItem) BoundingRectDefault() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundingRectDefault"}).(*core.QRectF)
}

func (ptr *QGraphicsSvgItem) ElementId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementId"}).(string)
}

func (ptr *QGraphicsSvgItem) MaximumCacheSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumCacheSize"}).(*core.QSize)
}

func (ptr *QGraphicsSvgItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPaint", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGraphicsSvgItem) DisconnectPaint() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPaint"})
}

func (ptr *QGraphicsSvgItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Paint", painter, option, widget})
}

func (ptr *QGraphicsSvgItem) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintDefault", painter, option, widget})
}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Renderer"}).(*QSvgRenderer)
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetElementId", id})
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaximumCacheSize", size})
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSharedRenderer", renderer})
}

func (ptr *QGraphicsSvgItem) TypeDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeDefault"}).(float64))
}

func (ptr *QGraphicsSvgItem) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGraphicsSvgItem) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGraphicsSvgItem) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGraphicsSvgItem) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGraphicsSvgItem) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGraphicsSvgItem) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGraphicsSvgItem) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGraphicsSvgItem) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGraphicsSvgItem) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __childItems_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QGraphicsSvgItem) __childItems_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_setList", i})
}

func (ptr *QGraphicsSvgItem) __childItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__childItems_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __collidingItems_atList(i int) *widgets.QGraphicsItem {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_atList", i}).(*widgets.QGraphicsItem)
}

func (ptr *QGraphicsSvgItem) __collidingItems_setList(i widgets.QGraphicsItem_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_setList", i})
}

func (ptr *QGraphicsSvgItem) __collidingItems_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__collidingItems_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __setTransformations_transformations_atList(i int) *widgets.QGraphicsTransform {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_atList", i}).(*widgets.QGraphicsTransform)
}

func (ptr *QGraphicsSvgItem) __setTransformations_transformations_setList(i widgets.QGraphicsTransform_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_setList", i})
}

func (ptr *QGraphicsSvgItem) __setTransformations_transformations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setTransformations_transformations_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) __transformations_atList(i int) *widgets.QGraphicsTransform {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_atList", i}).(*widgets.QGraphicsTransform)
}

func (ptr *QGraphicsSvgItem) __transformations_setList(i widgets.QGraphicsTransform_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_setList", i})
}

func (ptr *QGraphicsSvgItem) __transformations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__transformations_newList"}).(unsafe.Pointer)
}

func (ptr *QGraphicsSvgItem) EventDefault(ev core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", ev}).(bool)
}

func (ptr *QGraphicsSvgItem) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QGraphicsSvgItem) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGraphicsSvgItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGraphicsSvgItem) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGraphicsSvgItem) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGraphicsSvgItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGraphicsSvgItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGraphicsSvgItem) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGraphicsSvgItem) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func (ptr *QGraphicsSvgItem) AdvanceDefault(phase int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AdvanceDefault", phase})
}

func (ptr *QGraphicsSvgItem) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollidesWithItemDefault", other, mode}).(bool)
}

func (ptr *QGraphicsSvgItem) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CollidesWithPathDefault", path, mode}).(bool)
}

func (ptr *QGraphicsSvgItem) ContainsDefault(point core.QPointF_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContainsDefault", point}).(bool)
}

func (ptr *QGraphicsSvgItem) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QGraphicsSvgItem) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QGraphicsSvgItem) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QGraphicsSvgItem) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QGraphicsSvgItem) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QGraphicsSvgItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QGraphicsSvgItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QGraphicsSvgItem) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverEnterEventDefault", event})
}

func (ptr *QGraphicsSvgItem) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverLeaveEventDefault", event})
}

func (ptr *QGraphicsSvgItem) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HoverMoveEventDefault", event})
}

func (ptr *QGraphicsSvgItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QGraphicsSvgItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QGraphicsSvgItem) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsObscuredByDefault", item}).(bool)
}

func (ptr *QGraphicsSvgItem) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemChangeDefault", change, value}).(*core.QVariant)
}

func (ptr *QGraphicsSvgItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QGraphicsSvgItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QGraphicsSvgItem) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QGraphicsSvgItem) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QGraphicsSvgItem) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QGraphicsSvgItem) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QGraphicsSvgItem) OpaqueAreaDefault() *gui.QPainterPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpaqueAreaDefault"}).(*gui.QPainterPath)
}

func (ptr *QGraphicsSvgItem) SceneEventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneEventDefault", event}).(bool)
}

func (ptr *QGraphicsSvgItem) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SceneEventFilterDefault", watched, event}).(bool)
}

func (ptr *QGraphicsSvgItem) ShapeDefault() *gui.QPainterPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShapeDefault"}).(*gui.QPainterPath)
}

func (ptr *QGraphicsSvgItem) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

type QSvgGenerator struct {
	gui.QPaintDevice
}

type QSvgGenerator_ITF interface {
	gui.QPaintDevice_ITF
	QSvgGenerator_PTR() *QSvgGenerator
}

func (ptr *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator {
	return ptr
}

func (ptr *QSvgGenerator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgGenerator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQSvgGenerator(ptr QSvgGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgGenerator_PTR().Pointer()
	}
	return nil
}

func (n *QSvgGenerator) InitFromInternal(ptr uintptr, name string) {
	n.QPaintDevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSvgGenerator) ClassNameInternalF() string {
	return n.QPaintDevice_PTR().ClassNameInternalF()
}

func NewQSvgGeneratorFromPointer(ptr unsafe.Pointer) (n *QSvgGenerator) {
	n = new(QSvgGenerator)
	n.InitFromInternal(uintptr(ptr), "svg.QSvgGenerator")
	return
}
func NewQSvgGenerator() *QSvgGenerator {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgGenerator", ""}).(*QSvgGenerator)
}

func (ptr *QSvgGenerator) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QSvgGenerator) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OutputDevice"}).(*core.QIODevice)
}

func (ptr *QSvgGenerator) ConnectPaintEngine(f func() *gui.QPaintEngine) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPaintEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgGenerator) DisconnectPaintEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPaintEngine"})
}

func (ptr *QSvgGenerator) PaintEngine() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngine"}).(*gui.QPaintEngine)
}

func (ptr *QSvgGenerator) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QSvgGenerator) Resolution() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resolution"}).(float64))
}

func (ptr *QSvgGenerator) SetDescription(description string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDescription", description})
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFileName", fileName})
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOutputDevice", outputDevice})
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetResolution", dpi})
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSize", size})
}

func (ptr *QSvgGenerator) SetTitle(title string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTitle", title})
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewBox", viewBox})
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewBox2", viewBox})
}

func (ptr *QSvgGenerator) Size() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(*core.QSize)
}

func (ptr *QSvgGenerator) Title() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Title"}).(string)
}

func (ptr *QSvgGenerator) ViewBox() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewBox"}).(*core.QRect)
}

func (ptr *QSvgGenerator) ViewBoxF() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewBoxF"}).(*core.QRectF)
}

func (ptr *QSvgGenerator) ConnectDestroyQSvgGenerator(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSvgGenerator", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgGenerator) DisconnectDestroyQSvgGenerator() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSvgGenerator"})
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSvgGenerator"})
}

func (ptr *QSvgGenerator) DestroyQSvgGeneratorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSvgGeneratorDefault"})
}

type QSvgIOHandler struct {
	gui.QImageIOHandler
}

type QSvgIOHandler_ITF interface {
	gui.QImageIOHandler_ITF
	QSvgIOHandler_PTR() *QSvgIOHandler
}

func (ptr *QSvgIOHandler) QSvgIOHandler_PTR() *QSvgIOHandler {
	return ptr
}

func (ptr *QSvgIOHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageIOHandler_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgIOHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QImageIOHandler_PTR().SetPointer(p)
	}
}

func PointerFromQSvgIOHandler(ptr QSvgIOHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgIOHandler_PTR().Pointer()
	}
	return nil
}

func (n *QSvgIOHandler) InitFromInternal(ptr uintptr, name string) {
	n.QImageIOHandler_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSvgIOHandler) ClassNameInternalF() string {
	return n.QImageIOHandler_PTR().ClassNameInternalF()
}

func NewQSvgIOHandlerFromPointer(ptr unsafe.Pointer) (n *QSvgIOHandler) {
	n = new(QSvgIOHandler)
	n.InitFromInternal(uintptr(ptr), "svg.QSvgIOHandler")
	return
}

func (ptr *QSvgIOHandler) DestroyQSvgIOHandler() {
}

type QSvgIconEngine struct {
	gui.QIconEngine
}

type QSvgIconEngine_ITF interface {
	gui.QIconEngine_ITF
	QSvgIconEngine_PTR() *QSvgIconEngine
}

func (ptr *QSvgIconEngine) QSvgIconEngine_PTR() *QSvgIconEngine {
	return ptr
}

func (ptr *QSvgIconEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgIconEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIconEngine_PTR().SetPointer(p)
	}
}

func PointerFromQSvgIconEngine(ptr QSvgIconEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgIconEngine_PTR().Pointer()
	}
	return nil
}

func (n *QSvgIconEngine) InitFromInternal(ptr uintptr, name string) {
	n.QIconEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSvgIconEngine) ClassNameInternalF() string {
	return n.QIconEngine_PTR().ClassNameInternalF()
}

func NewQSvgIconEngineFromPointer(ptr unsafe.Pointer) (n *QSvgIconEngine) {
	n = new(QSvgIconEngine)
	n.InitFromInternal(uintptr(ptr), "svg.QSvgIconEngine")
	return
}

func (ptr *QSvgIconEngine) DestroyQSvgIconEngine() {
}

type QSvgRenderer struct {
	core.QObject
}

type QSvgRenderer_ITF interface {
	core.QObject_ITF
	QSvgRenderer_PTR() *QSvgRenderer
}

func (ptr *QSvgRenderer) QSvgRenderer_PTR() *QSvgRenderer {
	return ptr
}

func (ptr *QSvgRenderer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgRenderer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSvgRenderer(ptr QSvgRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgRenderer_PTR().Pointer()
	}
	return nil
}

func (n *QSvgRenderer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSvgRenderer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSvgRendererFromPointer(ptr unsafe.Pointer) (n *QSvgRenderer) {
	n = new(QSvgRenderer)
	n.InitFromInternal(uintptr(ptr), "svg.QSvgRenderer")
	return
}
func NewQSvgRenderer(parent core.QObject_ITF) *QSvgRenderer {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgRenderer", "", parent}).(*QSvgRenderer)
}

func NewQSvgRenderer2(filename string, parent core.QObject_ITF) *QSvgRenderer {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgRenderer2", "", filename, parent}).(*QSvgRenderer)
}

func NewQSvgRenderer3(contents core.QByteArray_ITF, parent core.QObject_ITF) *QSvgRenderer {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgRenderer3", "", contents, parent}).(*QSvgRenderer)
}

func NewQSvgRenderer4(contents core.QXmlStreamReader_ITF, parent core.QObject_ITF) *QSvgRenderer {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgRenderer4", "", contents, parent}).(*QSvgRenderer)
}

func (ptr *QSvgRenderer) Animated() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Animated"}).(bool)
}

func (ptr *QSvgRenderer) BoundsOnElement(id string) *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundsOnElement", id}).(*core.QRectF)
}

func (ptr *QSvgRenderer) DefaultSize() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultSize"}).(*core.QSize)
}

func (ptr *QSvgRenderer) ElementExists(id string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementExists", id}).(bool)
}

func (ptr *QSvgRenderer) FramesPerSecond() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FramesPerSecond"}).(float64))
}

func (ptr *QSvgRenderer) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSvgRenderer) ConnectLoad(f func(filename string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectLoad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad"})
}

func (ptr *QSvgRenderer) Load(filename string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", filename}).(bool)
}

func (ptr *QSvgRenderer) LoadDefault(filename string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadDefault", filename}).(bool)
}

func (ptr *QSvgRenderer) ConnectLoad2(f func(contents *core.QByteArray) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectLoad2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad2"})
}

func (ptr *QSvgRenderer) Load2(contents core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2", contents}).(bool)
}

func (ptr *QSvgRenderer) Load2Default(contents core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2Default", contents}).(bool)
}

func (ptr *QSvgRenderer) ConnectLoad3(f func(contents *core.QXmlStreamReader) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad3", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectLoad3() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad3"})
}

func (ptr *QSvgRenderer) Load3(contents core.QXmlStreamReader_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load3", contents}).(bool)
}

func (ptr *QSvgRenderer) Load3Default(contents core.QXmlStreamReader_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load3Default", contents}).(bool)
}

func (ptr *QSvgRenderer) MatrixForElement(id string) *gui.QMatrix {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MatrixForElement", id}).(*gui.QMatrix)
}

func (ptr *QSvgRenderer) ConnectRender(f func(painter *gui.QPainter)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRender", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectRender() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRender"})
}

func (ptr *QSvgRenderer) Render(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Render", painter})
}

func (ptr *QSvgRenderer) RenderDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RenderDefault", painter})
}

func (ptr *QSvgRenderer) ConnectRender2(f func(painter *gui.QPainter, bounds *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRender2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectRender2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRender2"})
}

func (ptr *QSvgRenderer) Render2(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Render2", painter, bounds})
}

func (ptr *QSvgRenderer) Render2Default(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Render2Default", painter, bounds})
}

func (ptr *QSvgRenderer) ConnectRender3(f func(painter *gui.QPainter, elementId string, bounds *core.QRectF)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRender3", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectRender3() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRender3"})
}

func (ptr *QSvgRenderer) Render3(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Render3", painter, elementId, bounds})
}

func (ptr *QSvgRenderer) Render3Default(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Render3Default", painter, elementId, bounds})
}

func (ptr *QSvgRenderer) ConnectRepaintNeeded(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRepaintNeeded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectRepaintNeeded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRepaintNeeded"})
}

func (ptr *QSvgRenderer) RepaintNeeded() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintNeeded"})
}

func (ptr *QSvgRenderer) SetFramesPerSecond(num int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFramesPerSecond", num})
}

func (ptr *QSvgRenderer) SetViewBox(viewbox core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewBox", viewbox})
}

func (ptr *QSvgRenderer) SetViewBox2(viewbox core.QRectF_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetViewBox2", viewbox})
}

func (ptr *QSvgRenderer) ViewBox() *core.QRect {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewBox"}).(*core.QRect)
}

func (ptr *QSvgRenderer) ViewBoxF() *core.QRectF {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ViewBoxF"}).(*core.QRectF)
}

func (ptr *QSvgRenderer) ConnectDestroyQSvgRenderer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSvgRenderer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgRenderer) DisconnectDestroyQSvgRenderer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSvgRenderer"})
}

func (ptr *QSvgRenderer) DestroyQSvgRenderer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSvgRenderer"})
}

func (ptr *QSvgRenderer) DestroyQSvgRendererDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSvgRendererDefault"})
}

func (ptr *QSvgRenderer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSvgRenderer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSvgRenderer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgRenderer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSvgRenderer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSvgRenderer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgRenderer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSvgRenderer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSvgRenderer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgRenderer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSvgRenderer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSvgRenderer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSvgRenderer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSvgRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSvgRenderer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSvgRenderer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSvgRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSvgRenderer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSvgRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSvgRenderer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSvgRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSvgWidget struct {
	widgets.QWidget
}

type QSvgWidget_ITF interface {
	widgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func (ptr *QSvgWidget) QSvgWidget_PTR() *QSvgWidget {
	return ptr
}

func (ptr *QSvgWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QSvgWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQSvgWidget(ptr QSvgWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidget_PTR().Pointer()
	}
	return nil
}

func (n *QSvgWidget) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSvgWidget) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQSvgWidgetFromPointer(ptr unsafe.Pointer) (n *QSvgWidget) {
	n = new(QSvgWidget)
	n.InitFromInternal(uintptr(ptr), "svg.QSvgWidget")
	return
}
func NewQSvgWidget(parent widgets.QWidget_ITF) *QSvgWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgWidget", "", parent}).(*QSvgWidget)
}

func NewQSvgWidget2(file string, parent widgets.QWidget_ITF) *QSvgWidget {

	return internal.CallLocalFunction([]interface{}{"", "", "svg.NewQSvgWidget2", "", file, parent}).(*QSvgWidget)
}

func (ptr *QSvgWidget) ConnectLoad(f func(file string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgWidget) DisconnectLoad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad"})
}

func (ptr *QSvgWidget) Load(file string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", file})
}

func (ptr *QSvgWidget) LoadDefault(file string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadDefault", file})
}

func (ptr *QSvgWidget) ConnectLoad2(f func(contents *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgWidget) DisconnectLoad2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad2"})
}

func (ptr *QSvgWidget) Load2(contents core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2", contents})
}

func (ptr *QSvgWidget) Load2Default(contents core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load2Default", contents})
}

func (ptr *QSvgWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Renderer"}).(*QSvgRenderer)
}

func (ptr *QSvgWidget) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QSvgWidget) ConnectDestroyQSvgWidget(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSvgWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSvgWidget) DisconnectDestroyQSvgWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSvgWidget"})
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSvgWidget"})
}

func (ptr *QSvgWidget) DestroyQSvgWidgetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSvgWidgetDefault"})
}

func (ptr *QSvgWidget) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QSvgWidget) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QSvgWidget) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QSvgWidget) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QSvgWidget) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QSvgWidget) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QSvgWidget) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSvgWidget) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSvgWidget) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSvgWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSvgWidget) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSvgWidget) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSvgWidget) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSvgWidget) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSvgWidget) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSvgWidget) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QSvgWidget) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QSvgWidget) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QSvgWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QSvgWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QSvgWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QSvgWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QSvgWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QSvgWidget) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QSvgWidget) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QSvgWidget) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QSvgWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QSvgWidget) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QSvgWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QSvgWidget) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QSvgWidget) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QSvgWidget) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QSvgWidget) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QSvgWidget) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QSvgWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QSvgWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QSvgWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QSvgWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QSvgWidget) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QSvgWidget) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QSvgWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QSvgWidget) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QSvgWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QSvgWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QSvgWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QSvgWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QSvgWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QSvgWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QSvgWidget) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QSvgWidget) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QSvgWidget) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QSvgWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QSvgWidget) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QSvgWidget) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QSvgWidget) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QSvgWidget) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QSvgWidget) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QSvgWidget) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QSvgWidget) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QSvgWidget) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QSvgWidget) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QSvgWidget) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QSvgWidget) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QSvgWidget) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QSvgWidget) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QSvgWidget) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QSvgWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QSvgWidget) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QSvgWidget) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QSvgWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QSvgWidget) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSvgWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSvgWidget) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSvgWidget) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSvgWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSvgWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSvgWidget) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSvgWidget) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["svg.QGraphicsSvgItem"] = NewQGraphicsSvgItemFromPointer
	internal.ConstructorTable["svg.QSvgGenerator"] = NewQSvgGeneratorFromPointer
	internal.ConstructorTable["svg.QSvgRenderer"] = NewQSvgRendererFromPointer
	internal.ConstructorTable["svg.QSvgWidget"] = NewQSvgWidgetFromPointer
}
