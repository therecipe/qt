package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchQueryWidget_") {
		n.SetObjectName("QHelpSearchQueryWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return ptr
}

func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::isCompactMode")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer()) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchQueryWidget::QHelpSearchQueryWidget")

	return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::collapseExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::expandExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "search")
	}
}

//export callbackQHelpSearchQueryWidgetSearch
func callbackQHelpSearchQueryWidgetSearch(ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::search")

	if signal := qt.GetSignal(C.GoString(ptrName), "search"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	defer qt.Recovering("QHelpSearchQueryWidget::~QHelpSearchQueryWidget")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpSearchQueryWidgetActionEvent
func callbackQHelpSearchQueryWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDragEnterEvent
func callbackQHelpSearchQueryWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDragLeaveEvent
func callbackQHelpSearchQueryWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDragMoveEvent
func callbackQHelpSearchQueryWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDropEvent
func callbackQHelpSearchQueryWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpSearchQueryWidgetEnterEvent
func callbackQHelpSearchQueryWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpSearchQueryWidgetFocusOutEvent
func callbackQHelpSearchQueryWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpSearchQueryWidgetHideEvent
func callbackQHelpSearchQueryWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetLeaveEvent
func callbackQHelpSearchQueryWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMoveEvent
func callbackQHelpSearchQueryWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpSearchQueryWidgetPaintEvent
func callbackQHelpSearchQueryWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpSearchQueryWidgetSetVisible
func callbackQHelpSearchQueryWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpSearchQueryWidgetShowEvent
func callbackQHelpSearchQueryWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpSearchQueryWidgetCloseEvent
func callbackQHelpSearchQueryWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpSearchQueryWidgetContextMenuEvent
func callbackQHelpSearchQueryWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpSearchQueryWidgetInitPainter
func callbackQHelpSearchQueryWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpSearchQueryWidgetInputMethodEvent
func callbackQHelpSearchQueryWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpSearchQueryWidgetKeyPressEvent
func callbackQHelpSearchQueryWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpSearchQueryWidgetKeyReleaseEvent
func callbackQHelpSearchQueryWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMouseDoubleClickEvent
func callbackQHelpSearchQueryWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMouseMoveEvent
func callbackQHelpSearchQueryWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMousePressEvent
func callbackQHelpSearchQueryWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMouseReleaseEvent
func callbackQHelpSearchQueryWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpSearchQueryWidgetResizeEvent
func callbackQHelpSearchQueryWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpSearchQueryWidgetTabletEvent
func callbackQHelpSearchQueryWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpSearchQueryWidgetWheelEvent
func callbackQHelpSearchQueryWidgetWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpSearchQueryWidgetTimerEvent
func callbackQHelpSearchQueryWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpSearchQueryWidgetChildEvent
func callbackQHelpSearchQueryWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpSearchQueryWidgetCustomEvent
func callbackQHelpSearchQueryWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
