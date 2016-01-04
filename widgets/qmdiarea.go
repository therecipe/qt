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
func callbackQMdiAreaChildEvent(ptr unsafe.Pointer, ptrName *C.char, childEvent unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(childEvent))
	} else {
		NewQMdiAreaFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(childEvent))
	}
}

func (ptr *QMdiArea) ChildEvent(childEvent core.QChildEvent_ITF) {
	defer qt.Recovering("QMdiArea::childEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(childEvent))
	}
}

func (ptr *QMdiArea) ChildEventDefault(childEvent core.QChildEvent_ITF) {
	defer qt.Recovering("QMdiArea::childEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(childEvent))
	}
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

func (ptr *QMdiArea) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QMdiArea::event")

	if ptr.Pointer() != nil {
		return C.QMdiArea_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMdiArea) EventFilter(object core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMdiArea::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMdiArea_EventFilter(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event)) != 0
	}
	return false
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
func callbackQMdiAreaPaintEvent(ptr unsafe.Pointer, ptrName *C.char, paintEvent unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(paintEvent))
	} else {
		NewQMdiAreaFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(paintEvent))
	}
}

func (ptr *QMdiArea) PaintEvent(paintEvent gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMdiArea::paintEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(paintEvent))
	}
}

func (ptr *QMdiArea) PaintEventDefault(paintEvent gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMdiArea::paintEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(paintEvent))
	}
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
func callbackQMdiAreaResizeEvent(ptr unsafe.Pointer, ptrName *C.char, resizeEvent unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(resizeEvent))
	} else {
		NewQMdiAreaFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(resizeEvent))
	}
}

func (ptr *QMdiArea) ResizeEvent(resizeEvent gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMdiArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(resizeEvent))
	}
}

func (ptr *QMdiArea) ResizeEventDefault(resizeEvent gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMdiArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(resizeEvent))
	}
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
func callbackQMdiAreaScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QMdiArea::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQMdiAreaFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QMdiArea) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QMdiArea::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QMdiArea_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QMdiArea) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QMdiArea::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QMdiArea_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
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
func callbackQMdiAreaSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiArea::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QMdiArea) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QMdiArea::setupViewport")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QMdiArea) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QMdiArea::setupViewport")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
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
func callbackQMdiAreaShowEvent(ptr unsafe.Pointer, ptrName *C.char, showEvent unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(showEvent))
	} else {
		NewQMdiAreaFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(showEvent))
	}
}

func (ptr *QMdiArea) ShowEvent(showEvent gui.QShowEvent_ITF) {
	defer qt.Recovering("QMdiArea::showEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(showEvent))
	}
}

func (ptr *QMdiArea) ShowEventDefault(showEvent gui.QShowEvent_ITF) {
	defer qt.Recovering("QMdiArea::showEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(showEvent))
	}
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
func callbackQMdiAreaSubWindowActivated(ptr unsafe.Pointer, ptrName *C.char, window unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::subWindowActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "subWindowActivated"); signal != nil {
		signal.(func(*QMdiSubWindow))(NewQMdiSubWindowFromPointer(window))
	}

}

func (ptr *QMdiArea) SubWindowActivated(window QMdiSubWindow_ITF) {
	defer qt.Recovering("QMdiArea::subWindowActivated")

	if ptr.Pointer() != nil {
		C.QMdiArea_SubWindowActivated(ptr.Pointer(), PointerFromQMdiSubWindow(window))
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
func callbackQMdiAreaTimerEvent(ptr unsafe.Pointer, ptrName *C.char, timerEvent unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(timerEvent))
	} else {
		NewQMdiAreaFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(timerEvent))
	}
}

func (ptr *QMdiArea) TimerEvent(timerEvent core.QTimerEvent_ITF) {
	defer qt.Recovering("QMdiArea::timerEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(timerEvent))
	}
}

func (ptr *QMdiArea) TimerEventDefault(timerEvent core.QTimerEvent_ITF) {
	defer qt.Recovering("QMdiArea::timerEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(timerEvent))
	}
}

func (ptr *QMdiArea) ViewportEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QMdiArea::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QMdiArea_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
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
func callbackQMdiAreaDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QMdiArea) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMdiArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMdiArea) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMdiArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQMdiAreaDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QMdiArea) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMdiArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMdiArea) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMdiArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQMdiAreaDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QMdiArea) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMdiArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMdiArea) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMdiArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQMdiAreaDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QMdiArea) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMdiArea::dropEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMdiArea) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMdiArea::dropEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQMdiAreaContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QMdiArea) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMdiArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QMdiArea) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMdiArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
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
func callbackQMdiAreaKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QMdiArea) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMdiArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QMdiArea) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMdiArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQMdiAreaMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMdiArea) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMdiArea) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQMdiAreaMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMdiArea) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMdiArea) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQMdiAreaMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMdiArea) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMdiArea) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQMdiAreaMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMdiArea) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMdiArea) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMdiArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQMdiAreaWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQMdiAreaFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QMdiArea) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMdiArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QMdiArea) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMdiArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
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
func callbackQMdiAreaChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQMdiAreaFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QMdiArea) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::changeEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QMdiArea) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::changeEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
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
func callbackQMdiAreaActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QMdiArea) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QMdiArea::actionEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QMdiArea) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QMdiArea::actionEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQMdiAreaEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMdiArea) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::enterEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMdiArea) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::enterEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMdiArea) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMdiArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMdiArea) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMdiArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMdiAreaFocusInEvent
func callbackQMdiAreaFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMdiArea) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMdiArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMdiArea) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMdiArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQMdiAreaFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMdiArea) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMdiArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMdiArea) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMdiArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQMdiAreaHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QMdiArea) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMdiArea::hideEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMdiArea) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMdiArea::hideEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQMdiAreaLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMdiArea) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMdiArea) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQMdiAreaMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QMdiArea) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMdiArea::moveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMdiArea) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMdiArea::moveEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQMdiAreaSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMdiArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMdiArea) SetVisible(visible bool) {
	defer qt.Recovering("QMdiArea::setVisible")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMdiArea) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QMdiArea::setVisible")

	if ptr.Pointer() != nil {
		C.QMdiArea_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQMdiAreaCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QMdiArea) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMdiArea::closeEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMdiArea) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMdiArea::closeEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQMdiAreaInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQMdiAreaFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QMdiArea) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMdiArea::initPainter")

	if ptr.Pointer() != nil {
		C.QMdiArea_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMdiArea) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMdiArea::initPainter")

	if ptr.Pointer() != nil {
		C.QMdiArea_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQMdiAreaInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QMdiArea) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMdiArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMdiArea) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMdiArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQMdiAreaKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMdiArea) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMdiArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMdiArea) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMdiArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQMdiAreaTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QMdiArea) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMdiArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMdiArea) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMdiArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQMdiAreaCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMdiArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMdiAreaFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMdiArea) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::customEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMdiArea) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMdiArea::customEvent")

	if ptr.Pointer() != nil {
		C.QMdiArea_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
