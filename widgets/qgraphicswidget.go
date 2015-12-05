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

type QGraphicsWidget struct {
	QGraphicsObject
	QGraphicsLayoutItem
}

type QGraphicsWidget_ITF interface {
	QGraphicsObject_ITF
	QGraphicsLayoutItem_ITF
	QGraphicsWidget_PTR() *QGraphicsWidget
}

func (p *QGraphicsWidget) Pointer() unsafe.Pointer {
	return p.QGraphicsObject_PTR().Pointer()
}

func (p *QGraphicsWidget) SetPointer(ptr unsafe.Pointer) {
	p.QGraphicsObject_PTR().SetPointer(ptr)
	p.QGraphicsLayoutItem_PTR().SetPointer(ptr)
}

func PointerFromQGraphicsWidget(ptr QGraphicsWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsWidgetFromPointer(ptr unsafe.Pointer) *QGraphicsWidget {
	var n = new(QGraphicsWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsWidget_") {
		n.SetObjectName("QGraphicsWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsWidget) QGraphicsWidget_PTR() *QGraphicsWidget {
	return ptr
}

func (ptr *QGraphicsWidget) AutoFillBackground() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::autoFillBackground")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_AutoFillBackground(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) FocusPolicy() core.Qt__FocusPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::focusPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__FocusPolicy(C.QGraphicsWidget_FocusPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) LayoutDirection() core.Qt__LayoutDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::layoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QGraphicsWidget_LayoutDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) Resize(size core.QSizeF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::resize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setAutoFillBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetAutoFillBackground(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetFocusPolicy(policy core.Qt__FocusPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setFocusPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFocusPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGraphicsWidget) SetFont(font gui.QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsWidget) SetGeometry(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsWidget) SetLayout(layout QGraphicsLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayout(ptr.Pointer(), PointerFromQGraphicsLayout(layout))
	}
}

func (ptr *QGraphicsWidget) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setLayoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayoutDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QGraphicsWidget) SetPalette(palette gui.QPalette_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setPalette")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(palette))
	}
}

func (ptr *QGraphicsWidget) SetWindowFlags(wFlags core.Qt__WindowType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setWindowFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowFlags(ptr.Pointer(), C.int(wFlags))
	}
}

func (ptr *QGraphicsWidget) SetWindowTitle(title string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setWindowTitle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QGraphicsWidget) UnsetLayoutDirection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::unsetLayoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UnsetLayoutDirection(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) WindowFlags() core.Qt__WindowType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::windowFlags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QGraphicsWidget_WindowFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) WindowTitle() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::windowTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func NewQGraphicsWidget(parent QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QGraphicsWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::QGraphicsWidget")
		}
	}()

	return NewQGraphicsWidgetFromPointer(C.QGraphicsWidget_NewQGraphicsWidget(PointerFromQGraphicsItem(parent), C.int(wFlags)))
}

func (ptr *QGraphicsWidget) AddAction(action QAction_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::addAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AddAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) AdjustSize() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::adjustSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) Close() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::close")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) FocusWidget() *QGraphicsWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::focusWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsWidgetFromPointer(C.QGraphicsWidget_FocusWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) ConnectGeometryChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::geometryChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ConnectGeometryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectGeometryChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::geometryChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_DisconnectGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQGraphicsWidgetGeometryChanged
func callbackQGraphicsWidgetGeometryChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::geometryChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "geometryChanged").(func())()
}

func (ptr *QGraphicsWidget) GrabShortcut(sequence gui.QKeySequence_ITF, context core.Qt__ShortcutContext) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::grabShortcut")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsWidget_GrabShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(sequence), C.int(context)))
	}
	return 0
}

func (ptr *QGraphicsWidget) InsertAction(before QAction_ITF, action QAction_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::insertAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_InsertAction(ptr.Pointer(), PointerFromQAction(before), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) IsActiveWindow() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::isActiveWindow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_IsActiveWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) Layout() *QGraphicsLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::layout")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutFromPointer(C.QGraphicsWidget_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsWidget) PaintWindowFrame(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::paintWindowFrame")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PaintWindowFrame(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsWidget) ReleaseShortcut(id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::releaseShortcut")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ReleaseShortcut(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QGraphicsWidget) RemoveAction(action QAction_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::removeAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) Resize2(w float64, h float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::resize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize2(ptr.Pointer(), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsWidget) SetAttribute(attribute core.Qt__WidgetAttribute, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetAttribute(ptr.Pointer(), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setContentsMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetContentsMargins(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

func (ptr *QGraphicsWidget) SetGeometry2(x float64, y float64, w float64, h float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetGeometry2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setShortcutAutoRepeat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetShortcutAutoRepeat(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setShortcutEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetShortcutEnabled(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetStyle(style QStyle_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func QGraphicsWidget_SetTabOrder(first QGraphicsWidget_ITF, second QGraphicsWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setTabOrder")
		}
	}()

	C.QGraphicsWidget_QGraphicsWidget_SetTabOrder(PointerFromQGraphicsWidget(first), PointerFromQGraphicsWidget(second))
}

func (ptr *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::setWindowFrameMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowFrameMargins(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

func (ptr *QGraphicsWidget) Style() *QStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::style")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QGraphicsWidget_Style(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) TestAttribute(attribute core.Qt__WidgetAttribute) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::testAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_TestAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsWidget_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) UnsetWindowFrameMargins() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::unsetWindowFrameMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UnsetWindowFrameMargins(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) WindowType() core.Qt__WindowType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::windowType")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QGraphicsWidget_WindowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) DestroyQGraphicsWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsWidget::~QGraphicsWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_DestroyQGraphicsWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
