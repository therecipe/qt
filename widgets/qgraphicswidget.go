package widgets

//#include "qgraphicswidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsWidget struct {
	QGraphicsObject
	QGraphicsLayoutItem
}

type QGraphicsWidgetITF interface {
	QGraphicsObjectITF
	QGraphicsLayoutItemITF
	QGraphicsWidgetPTR() *QGraphicsWidget
}

func (p *QGraphicsWidget) Pointer() unsafe.Pointer {
	return p.QGraphicsObjectPTR().Pointer()
}

func (p *QGraphicsWidget) SetPointer(ptr unsafe.Pointer) {
	p.QGraphicsObjectPTR().SetPointer(ptr)
	p.QGraphicsLayoutItemPTR().SetPointer(ptr)
}

func PointerFromQGraphicsWidget(ptr QGraphicsWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidgetPTR().Pointer()
	}
	return nil
}

func QGraphicsWidgetFromPointer(ptr unsafe.Pointer) *QGraphicsWidget {
	var n = new(QGraphicsWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsWidget) QGraphicsWidgetPTR() *QGraphicsWidget {
	return ptr
}

func (ptr *QGraphicsWidget) AutoFillBackground() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_AutoFillBackground(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) FocusPolicy() core.Qt__FocusPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__FocusPolicy(C.QGraphicsWidget_FocusPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsWidget) LayoutDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QGraphicsWidget_LayoutDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsWidget) Resize(size core.QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSizeF(size)))
	}
}

func (ptr *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetAutoFillBackground(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetFocusPolicy(policy core.Qt__FocusPolicy) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFocusPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QGraphicsWidget) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QGraphicsWidget) SetGeometry(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsWidget) SetLayout(layout QGraphicsLayoutITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayout(layout)))
	}
}

func (ptr *QGraphicsWidget) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayoutDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QGraphicsWidget) SetPalette(palette gui.QPaletteITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetPalette(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPalette(palette)))
	}
}

func (ptr *QGraphicsWidget) SetWindowFlags(wFlags core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowFlags(C.QtObjectPtr(ptr.Pointer()), C.int(wFlags))
	}
}

func (ptr *QGraphicsWidget) SetWindowTitle(title string) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QGraphicsWidget) UnsetLayoutDirection() {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UnsetLayoutDirection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsWidget) WindowFlags() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QGraphicsWidget_WindowFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsWidget_WindowTitle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQGraphicsWidget(parent QGraphicsItemITF, wFlags core.Qt__WindowType) *QGraphicsWidget {
	return QGraphicsWidgetFromPointer(unsafe.Pointer(C.QGraphicsWidget_NewQGraphicsWidget(C.QtObjectPtr(PointerFromQGraphicsItem(parent)), C.int(wFlags))))
}

func (ptr *QGraphicsWidget) AddAction(action QActionITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AddAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QGraphicsWidget) AdjustSize() {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AdjustSize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_Close(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) FocusWidget() *QGraphicsWidget {
	if ptr.Pointer() != nil {
		return QGraphicsWidgetFromPointer(unsafe.Pointer(C.QGraphicsWidget_FocusWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsWidget) ConnectGeometryChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ConnectGeometryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectGeometryChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_DisconnectGeometryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQGraphicsWidgetGeometryChanged
func callbackQGraphicsWidgetGeometryChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "geometryChanged").(func())()
}

func (ptr *QGraphicsWidget) GrabShortcut(sequence gui.QKeySequenceITF, context core.Qt__ShortcutContext) int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsWidget_GrabShortcut(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQKeySequence(sequence)), C.int(context)))
	}
	return 0
}

func (ptr *QGraphicsWidget) InsertAction(before QActionITF, action QActionITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_InsertAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QGraphicsWidget) IsActiveWindow() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_IsActiveWindow(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) Layout() *QGraphicsLayout {
	if ptr.Pointer() != nil {
		return QGraphicsLayoutFromPointer(unsafe.Pointer(C.QGraphicsWidget_Layout(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsWidget) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsWidget) PaintWindowFrame(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PaintWindowFrame(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsWidget) ReleaseShortcut(id int) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ReleaseShortcut(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QGraphicsWidget) RemoveAction(action QActionITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_RemoveAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QGraphicsWidget) SetAttribute(attribute core.Qt__WidgetAttribute, on bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetShortcutAutoRepeat(C.QtObjectPtr(ptr.Pointer()), C.int(id), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetShortcutEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(id), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetStyle(style QStyleITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStyle(style)))
	}
}

func QGraphicsWidget_SetTabOrder(first QGraphicsWidgetITF, second QGraphicsWidgetITF) {
	C.QGraphicsWidget_QGraphicsWidget_SetTabOrder(C.QtObjectPtr(PointerFromQGraphicsWidget(first)), C.QtObjectPtr(PointerFromQGraphicsWidget(second)))
}

func (ptr *QGraphicsWidget) Style() *QStyle {
	if ptr.Pointer() != nil {
		return QStyleFromPointer(unsafe.Pointer(C.QGraphicsWidget_Style(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsWidget) TestAttribute(attribute core.Qt__WidgetAttribute) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_TestAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsWidget_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsWidget) UnsetWindowFrameMargins() {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UnsetWindowFrameMargins(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsWidget) WindowType() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QGraphicsWidget_WindowType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsWidget) DestroyQGraphicsWidget() {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_DestroyQGraphicsWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
