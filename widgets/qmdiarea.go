package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMdiArea struct {
	QAbstractScrollArea
}

type QMdiArea_ITF interface {
	QAbstractScrollArea_ITF
	QMdiArea_PTR() *QMdiArea
}

func PointerFromQMdiArea(ptr QMdiArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiArea_PTR().Pointer()
	}
	return nil
}

func NewQMdiAreaFromPointer(ptr unsafe.Pointer) *QMdiArea {
	var n = new(QMdiArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMdiArea_") {
		n.SetObjectName("QMdiArea_" + qt.Identifier())
	}
	return n
}

func (ptr *QMdiArea) QMdiArea_PTR() *QMdiArea {
	return ptr
}

//QMdiArea::AreaOption
type QMdiArea__AreaOption int64

const (
	QMdiArea__DontMaximizeSubWindowOnActivation = QMdiArea__AreaOption(0x1)
)

//QMdiArea::ViewMode
type QMdiArea__ViewMode int64

const (
	QMdiArea__SubWindowView = QMdiArea__ViewMode(0)
	QMdiArea__TabbedView    = QMdiArea__ViewMode(1)
)

//QMdiArea::WindowOrder
type QMdiArea__WindowOrder int64

const (
	QMdiArea__CreationOrder          = QMdiArea__WindowOrder(0)
	QMdiArea__StackingOrder          = QMdiArea__WindowOrder(1)
	QMdiArea__ActivationHistoryOrder = QMdiArea__WindowOrder(2)
)

func (ptr *QMdiArea) ActivationOrder() QMdiArea__WindowOrder {
	defer qt.Recovering("QMdiArea::activationOrder")

	if ptr.Pointer() != nil {
		return QMdiArea__WindowOrder(C.QMdiArea_ActivationOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiArea) Background() *gui.QBrush {
	defer qt.Recovering("QMdiArea::background")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QMdiArea_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) DocumentMode() bool {
	defer qt.Recovering("QMdiArea::documentMode")

	if ptr.Pointer() != nil {
		return C.QMdiArea_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiArea) SetActivationOrder(order QMdiArea__WindowOrder) {
	defer qt.Recovering("QMdiArea::setActivationOrder")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetActivationOrder(ptr.Pointer(), C.int(order))
	}
}

func (ptr *QMdiArea) SetBackground(background gui.QBrush_ITF) {
	defer qt.Recovering("QMdiArea::setBackground")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetBackground(ptr.Pointer(), gui.PointerFromQBrush(background))
	}
}

func (ptr *QMdiArea) SetDocumentMode(enabled bool) {
	defer qt.Recovering("QMdiArea::setDocumentMode")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMdiArea) SetTabPosition(position QTabWidget__TabPosition) {
	defer qt.Recovering("QMdiArea::setTabPosition")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabPosition(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QMdiArea) SetTabShape(shape QTabWidget__TabShape) {
	defer qt.Recovering("QMdiArea::setTabShape")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QMdiArea) SetTabsClosable(closable bool) {
	defer qt.Recovering("QMdiArea::setTabsClosable")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QMdiArea) SetTabsMovable(movable bool) {
	defer qt.Recovering("QMdiArea::setTabsMovable")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabsMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QMdiArea) SetViewMode(mode QMdiArea__ViewMode) {
	defer qt.Recovering("QMdiArea::setViewMode")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QMdiArea) TabPosition() QTabWidget__TabPosition {
	defer qt.Recovering("QMdiArea::tabPosition")

	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QMdiArea_TabPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiArea) TabShape() QTabWidget__TabShape {
	defer qt.Recovering("QMdiArea::tabShape")

	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QMdiArea_TabShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiArea) TabsClosable() bool {
	defer qt.Recovering("QMdiArea::tabsClosable")

	if ptr.Pointer() != nil {
		return C.QMdiArea_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiArea) TabsMovable() bool {
	defer qt.Recovering("QMdiArea::tabsMovable")

	if ptr.Pointer() != nil {
		return C.QMdiArea_TabsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiArea) ViewMode() QMdiArea__ViewMode {
	defer qt.Recovering("QMdiArea::viewMode")

	if ptr.Pointer() != nil {
		return QMdiArea__ViewMode(C.QMdiArea_ViewMode(ptr.Pointer()))
	}
	return 0
}

func NewQMdiArea(parent QWidget_ITF) *QMdiArea {
	defer qt.Recovering("QMdiArea::QMdiArea")

	return NewQMdiAreaFromPointer(C.QMdiArea_NewQMdiArea(PointerFromQWidget(parent)))
}

func (ptr *QMdiArea) ActivateNextSubWindow() {
	defer qt.Recovering("QMdiArea::activateNextSubWindow")

	if ptr.Pointer() != nil {
		C.QMdiArea_ActivateNextSubWindow(ptr.Pointer())
	}
}

func (ptr *QMdiArea) ActivatePreviousSubWindow() {
	defer qt.Recovering("QMdiArea::activatePreviousSubWindow")

	if ptr.Pointer() != nil {
		C.QMdiArea_ActivatePreviousSubWindow(ptr.Pointer())
	}
}

func (ptr *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	defer qt.Recovering("QMdiArea::activeSubWindow")

	if ptr.Pointer() != nil {
		return NewQMdiSubWindowFromPointer(C.QMdiArea_ActiveSubWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) AddSubWindow(widget QWidget_ITF, windowFlags core.Qt__WindowType) *QMdiSubWindow {
	defer qt.Recovering("QMdiArea::addSubWindow")

	if ptr.Pointer() != nil {
		return NewQMdiSubWindowFromPointer(C.QMdiArea_AddSubWindow(ptr.Pointer(), PointerFromQWidget(widget), C.int(windowFlags)))
	}
	return nil
}

func (ptr *QMdiArea) CascadeSubWindows() {
	defer qt.Recovering("QMdiArea::cascadeSubWindows")

	if ptr.Pointer() != nil {
		C.QMdiArea_CascadeSubWindows(ptr.Pointer())
	}
}

func (ptr *QMdiArea) ConnectChildEvent(f func(childEvent *core.QChildEvent)) {
	defer qt.Recovering("connect QMdiArea::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMdiArea::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMdiAreaChildEvent
func callbackQMdiAreaChildEvent(ptrName *C.char, childEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(childEvent))
		return true
	}
	return false

}

func (ptr *QMdiArea) CloseActiveSubWindow() {
	defer qt.Recovering("QMdiArea::closeActiveSubWindow")

	if ptr.Pointer() != nil {
		C.QMdiArea_CloseActiveSubWindow(ptr.Pointer())
	}
}

func (ptr *QMdiArea) CloseAllSubWindows() {
	defer qt.Recovering("QMdiArea::closeAllSubWindows")

	if ptr.Pointer() != nil {
		C.QMdiArea_CloseAllSubWindows(ptr.Pointer())
	}
}

func (ptr *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	defer qt.Recovering("QMdiArea::currentSubWindow")

	if ptr.Pointer() != nil {
		return NewQMdiSubWindowFromPointer(C.QMdiArea_CurrentSubWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QMdiArea::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMdiArea_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) ConnectPaintEvent(f func(paintEvent *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMdiArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMdiArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMdiAreaPaintEvent
func callbackQMdiAreaPaintEvent(ptrName *C.char, paintEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(paintEvent))
		return true
	}
	return false

}

func (ptr *QMdiArea) RemoveSubWindow(widget QWidget_ITF) {
	defer qt.Recovering("QMdiArea::removeSubWindow")

	if ptr.Pointer() != nil {
		C.QMdiArea_RemoveSubWindow(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMdiArea) ConnectResizeEvent(f func(resizeEvent *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMdiArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMdiArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMdiAreaResizeEvent
func callbackQMdiAreaResizeEvent(ptrName *C.char, resizeEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(resizeEvent))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QMdiArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QMdiArea) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QMdiArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQMdiAreaScrollContentsBy
func callbackQMdiAreaScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QMdiArea::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QMdiArea) SetActiveSubWindow(window QMdiSubWindow_ITF) {
	defer qt.Recovering("QMdiArea::setActiveSubWindow")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetActiveSubWindow(ptr.Pointer(), PointerFromQMdiSubWindow(window))
	}
}

func (ptr *QMdiArea) SetOption(option QMdiArea__AreaOption, on bool) {
	defer qt.Recovering("QMdiArea::setOption")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiArea) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QMdiArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QMdiArea) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QMdiArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQMdiAreaSetupViewport
func callbackQMdiAreaSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectShowEvent(f func(showEvent *gui.QShowEvent)) {
	defer qt.Recovering("connect QMdiArea::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMdiArea::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMdiAreaShowEvent
func callbackQMdiAreaShowEvent(ptrName *C.char, showEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(showEvent))
		return true
	}
	return false

}

func (ptr *QMdiArea) SizeHint() *core.QSize {
	defer qt.Recovering("QMdiArea::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMdiArea_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) ConnectSubWindowActivated(f func(window *QMdiSubWindow)) {
	defer qt.Recovering("connect QMdiArea::subWindowActivated")

	if ptr.Pointer() != nil {
		C.QMdiArea_ConnectSubWindowActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "subWindowActivated", f)
	}
}

func (ptr *QMdiArea) DisconnectSubWindowActivated() {
	defer qt.Recovering("disconnect QMdiArea::subWindowActivated")

	if ptr.Pointer() != nil {
		C.QMdiArea_DisconnectSubWindowActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "subWindowActivated")
	}
}

//export callbackQMdiAreaSubWindowActivated
func callbackQMdiAreaSubWindowActivated(ptrName *C.char, window unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::subWindowActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "subWindowActivated"); signal != nil {
		signal.(func(*QMdiSubWindow))(NewQMdiSubWindowFromPointer(window))
	}

}

func (ptr *QMdiArea) TestOption(option QMdiArea__AreaOption) bool {
	defer qt.Recovering("QMdiArea::testOption")

	if ptr.Pointer() != nil {
		return C.QMdiArea_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiArea) TileSubWindows() {
	defer qt.Recovering("QMdiArea::tileSubWindows")

	if ptr.Pointer() != nil {
		C.QMdiArea_TileSubWindows(ptr.Pointer())
	}
}

func (ptr *QMdiArea) ConnectTimerEvent(f func(timerEvent *core.QTimerEvent)) {
	defer qt.Recovering("connect QMdiArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMdiArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMdiAreaTimerEvent
func callbackQMdiAreaTimerEvent(ptrName *C.char, timerEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(timerEvent))
		return true
	}
	return false

}

func (ptr *QMdiArea) DestroyQMdiArea() {
	defer qt.Recovering("QMdiArea::~QMdiArea")

	if ptr.Pointer() != nil {
		C.QMdiArea_DestroyQMdiArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMdiArea) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMdiArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMdiArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMdiAreaDragEnterEvent
func callbackQMdiAreaDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMdiArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMdiArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMdiAreaDragLeaveEvent
func callbackQMdiAreaDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMdiArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMdiArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMdiAreaDragMoveEvent
func callbackQMdiAreaDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMdiArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMdiArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMdiAreaDropEvent
func callbackQMdiAreaDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMdiArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMdiArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMdiAreaContextMenuEvent
func callbackQMdiAreaContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMdiArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMdiArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMdiAreaKeyPressEvent
func callbackQMdiAreaKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMdiArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMdiAreaMouseDoubleClickEvent
func callbackQMdiAreaMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMdiArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMdiAreaMouseMoveEvent
func callbackQMdiAreaMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMdiArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMdiAreaMousePressEvent
func callbackQMdiAreaMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMdiArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMdiAreaMouseReleaseEvent
func callbackQMdiAreaMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMdiArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMdiArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMdiAreaWheelEvent
func callbackQMdiAreaWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QMdiArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMdiArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMdiAreaChangeEvent
func callbackQMdiAreaChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMdiArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMdiArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMdiAreaActionEvent
func callbackQMdiAreaActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMdiArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMdiArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMdiAreaEnterEvent
func callbackQMdiAreaEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMdiArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMdiArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMdiAreaFocusOutEvent
func callbackQMdiAreaFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMdiArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMdiArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMdiAreaHideEvent
func callbackQMdiAreaHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMdiArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMdiArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMdiAreaLeaveEvent
func callbackQMdiAreaLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMdiArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMdiArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMdiAreaMoveEvent
func callbackQMdiAreaMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMdiArea::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMdiArea) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMdiArea::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMdiAreaSetVisible
func callbackQMdiAreaSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMdiArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMdiArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMdiArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMdiAreaCloseEvent
func callbackQMdiAreaCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMdiArea::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMdiArea) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMdiArea::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMdiAreaInitPainter
func callbackQMdiAreaInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMdiArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMdiArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMdiAreaInputMethodEvent
func callbackQMdiAreaInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMdiArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMdiArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMdiAreaKeyReleaseEvent
func callbackQMdiAreaKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMdiArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMdiArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMdiAreaTabletEvent
func callbackQMdiAreaTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiArea) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMdiArea::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMdiArea::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMdiAreaCustomEvent
func callbackQMdiAreaCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
