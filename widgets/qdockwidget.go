package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDockWidget struct {
	QWidget
}

type QDockWidget_ITF interface {
	QWidget_ITF
	QDockWidget_PTR() *QDockWidget
}

func PointerFromQDockWidget(ptr QDockWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDockWidget_PTR().Pointer()
	}
	return nil
}

func NewQDockWidgetFromPointer(ptr unsafe.Pointer) *QDockWidget {
	var n = new(QDockWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDockWidget_") {
		n.SetObjectName("QDockWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QDockWidget) QDockWidget_PTR() *QDockWidget {
	return ptr
}

//QDockWidget::DockWidgetFeature
type QDockWidget__DockWidgetFeature int64

const (
	QDockWidget__DockWidgetClosable         = QDockWidget__DockWidgetFeature(0x01)
	QDockWidget__DockWidgetMovable          = QDockWidget__DockWidgetFeature(0x02)
	QDockWidget__DockWidgetFloatable        = QDockWidget__DockWidgetFeature(0x04)
	QDockWidget__DockWidgetVerticalTitleBar = QDockWidget__DockWidgetFeature(0x08)
	QDockWidget__DockWidgetFeatureMask      = QDockWidget__DockWidgetFeature(0x0f)
	QDockWidget__AllDockWidgetFeatures      = QDockWidget__DockWidgetFeature(QDockWidget__DockWidgetClosable | QDockWidget__DockWidgetMovable | QDockWidget__DockWidgetFloatable)
	QDockWidget__NoDockWidgetFeatures       = QDockWidget__DockWidgetFeature(0x00)
	QDockWidget__Reserved                   = QDockWidget__DockWidgetFeature(0xff)
)

func (ptr *QDockWidget) AllowedAreas() core.Qt__DockWidgetArea {
	defer qt.Recovering("QDockWidget::allowedAreas")

	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QDockWidget_AllowedAreas(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDockWidget) Features() QDockWidget__DockWidgetFeature {
	defer qt.Recovering("QDockWidget::features")

	if ptr.Pointer() != nil {
		return QDockWidget__DockWidgetFeature(C.QDockWidget_Features(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDockWidget) SetAllowedAreas(areas core.Qt__DockWidgetArea) {
	defer qt.Recovering("QDockWidget::setAllowedAreas")

	if ptr.Pointer() != nil {
		C.QDockWidget_SetAllowedAreas(ptr.Pointer(), C.int(areas))
	}
}

func (ptr *QDockWidget) SetFeatures(features QDockWidget__DockWidgetFeature) {
	defer qt.Recovering("QDockWidget::setFeatures")

	if ptr.Pointer() != nil {
		C.QDockWidget_SetFeatures(ptr.Pointer(), C.int(features))
	}
}

func (ptr *QDockWidget) SetFloating(floating bool) {
	defer qt.Recovering("QDockWidget::setFloating")

	if ptr.Pointer() != nil {
		C.QDockWidget_SetFloating(ptr.Pointer(), C.int(qt.GoBoolToInt(floating)))
	}
}

func NewQDockWidget2(parent QWidget_ITF, flags core.Qt__WindowType) *QDockWidget {
	defer qt.Recovering("QDockWidget::QDockWidget")

	return NewQDockWidgetFromPointer(C.QDockWidget_NewQDockWidget2(PointerFromQWidget(parent), C.int(flags)))
}

func NewQDockWidget(title string, parent QWidget_ITF, flags core.Qt__WindowType) *QDockWidget {
	defer qt.Recovering("QDockWidget::QDockWidget")

	return NewQDockWidgetFromPointer(C.QDockWidget_NewQDockWidget(C.CString(title), PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QDockWidget) ConnectAllowedAreasChanged(f func(allowedAreas core.Qt__DockWidgetArea)) {
	defer qt.Recovering("connect QDockWidget::allowedAreasChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectAllowedAreasChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "allowedAreasChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectAllowedAreasChanged() {
	defer qt.Recovering("disconnect QDockWidget::allowedAreasChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectAllowedAreasChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "allowedAreasChanged")
	}
}

//export callbackQDockWidgetAllowedAreasChanged
func callbackQDockWidgetAllowedAreasChanged(ptrName *C.char, allowedAreas C.int) {
	defer qt.Recovering("callback QDockWidget::allowedAreasChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged"); signal != nil {
		signal.(func(core.Qt__DockWidgetArea))(core.Qt__DockWidgetArea(allowedAreas))
	}

}

func (ptr *QDockWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDockWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDockWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDockWidgetChangeEvent
func callbackQDockWidgetChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDockWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDockWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDockWidgetCloseEvent
func callbackQDockWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectDockLocationChanged(f func(area core.Qt__DockWidgetArea)) {
	defer qt.Recovering("connect QDockWidget::dockLocationChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectDockLocationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dockLocationChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectDockLocationChanged() {
	defer qt.Recovering("disconnect QDockWidget::dockLocationChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectDockLocationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dockLocationChanged")
	}
}

//export callbackQDockWidgetDockLocationChanged
func callbackQDockWidgetDockLocationChanged(ptrName *C.char, area C.int) {
	defer qt.Recovering("callback QDockWidget::dockLocationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "dockLocationChanged"); signal != nil {
		signal.(func(core.Qt__DockWidgetArea))(core.Qt__DockWidgetArea(area))
	}

}

func (ptr *QDockWidget) ConnectFeaturesChanged(f func(features QDockWidget__DockWidgetFeature)) {
	defer qt.Recovering("connect QDockWidget::featuresChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectFeaturesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "featuresChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectFeaturesChanged() {
	defer qt.Recovering("disconnect QDockWidget::featuresChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectFeaturesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "featuresChanged")
	}
}

//export callbackQDockWidgetFeaturesChanged
func callbackQDockWidgetFeaturesChanged(ptrName *C.char, features C.int) {
	defer qt.Recovering("callback QDockWidget::featuresChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "featuresChanged"); signal != nil {
		signal.(func(QDockWidget__DockWidgetFeature))(QDockWidget__DockWidgetFeature(features))
	}

}

func (ptr *QDockWidget) IsAreaAllowed(area core.Qt__DockWidgetArea) bool {
	defer qt.Recovering("QDockWidget::isAreaAllowed")

	if ptr.Pointer() != nil {
		return C.QDockWidget_IsAreaAllowed(ptr.Pointer(), C.int(area)) != 0
	}
	return false
}

func (ptr *QDockWidget) IsFloating() bool {
	defer qt.Recovering("QDockWidget::isFloating")

	if ptr.Pointer() != nil {
		return C.QDockWidget_IsFloating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDockWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDockWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDockWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDockWidgetPaintEvent
func callbackQDockWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) SetTitleBarWidget(widget QWidget_ITF) {
	defer qt.Recovering("QDockWidget::setTitleBarWidget")

	if ptr.Pointer() != nil {
		C.QDockWidget_SetTitleBarWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDockWidget) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QDockWidget::setWidget")

	if ptr.Pointer() != nil {
		C.QDockWidget_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDockWidget) TitleBarWidget() *QWidget {
	defer qt.Recovering("QDockWidget::titleBarWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDockWidget_TitleBarWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDockWidget) ToggleViewAction() *QAction {
	defer qt.Recovering("QDockWidget::toggleViewAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QDockWidget_ToggleViewAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDockWidget) ConnectTopLevelChanged(f func(topLevel bool)) {
	defer qt.Recovering("connect QDockWidget::topLevelChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectTopLevelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "topLevelChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectTopLevelChanged() {
	defer qt.Recovering("disconnect QDockWidget::topLevelChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectTopLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "topLevelChanged")
	}
}

//export callbackQDockWidgetTopLevelChanged
func callbackQDockWidgetTopLevelChanged(ptrName *C.char, topLevel C.int) {
	defer qt.Recovering("callback QDockWidget::topLevelChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "topLevelChanged"); signal != nil {
		signal.(func(bool))(int(topLevel) != 0)
	}

}

func (ptr *QDockWidget) ConnectVisibilityChanged(f func(visible bool)) {
	defer qt.Recovering("connect QDockWidget::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectVisibilityChanged() {
	defer qt.Recovering("disconnect QDockWidget::visibilityChanged")

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQDockWidgetVisibilityChanged
func callbackQDockWidgetVisibilityChanged(ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QDockWidget::visibilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "visibilityChanged"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	}

}

func (ptr *QDockWidget) Widget() *QWidget {
	defer qt.Recovering("QDockWidget::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDockWidget_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDockWidget) DestroyQDockWidget() {
	defer qt.Recovering("QDockWidget::~QDockWidget")

	if ptr.Pointer() != nil {
		C.QDockWidget_DestroyQDockWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDockWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDockWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDockWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDockWidgetActionEvent
func callbackQDockWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDockWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDockWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDockWidgetDragEnterEvent
func callbackQDockWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDockWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDockWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDockWidgetDragLeaveEvent
func callbackQDockWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDockWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDockWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDockWidgetDragMoveEvent
func callbackQDockWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDockWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDockWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDockWidgetDropEvent
func callbackQDockWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDockWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDockWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDockWidgetEnterEvent
func callbackQDockWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDockWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDockWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDockWidgetFocusInEvent
func callbackQDockWidgetFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDockWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDockWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDockWidgetFocusOutEvent
func callbackQDockWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDockWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDockWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDockWidgetHideEvent
func callbackQDockWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDockWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDockWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDockWidgetLeaveEvent
func callbackQDockWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDockWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDockWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDockWidgetMoveEvent
func callbackQDockWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDockWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDockWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDockWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDockWidgetSetVisible
func callbackQDockWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDockWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDockWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDockWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDockWidgetShowEvent
func callbackQDockWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDockWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDockWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDockWidgetContextMenuEvent
func callbackQDockWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDockWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDockWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDockWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDockWidgetInitPainter
func callbackQDockWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDockWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDockWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDockWidgetInputMethodEvent
func callbackQDockWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDockWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDockWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDockWidgetKeyPressEvent
func callbackQDockWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDockWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDockWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDockWidgetKeyReleaseEvent
func callbackQDockWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDockWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDockWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDockWidgetMouseDoubleClickEvent
func callbackQDockWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDockWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDockWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDockWidgetMouseMoveEvent
func callbackQDockWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDockWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDockWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDockWidgetMousePressEvent
func callbackQDockWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDockWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDockWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDockWidgetMouseReleaseEvent
func callbackQDockWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDockWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDockWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDockWidgetResizeEvent
func callbackQDockWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDockWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDockWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDockWidgetTabletEvent
func callbackQDockWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDockWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDockWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDockWidgetWheelEvent
func callbackQDockWidgetWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDockWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDockWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDockWidgetTimerEvent
func callbackQDockWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDockWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDockWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDockWidgetChildEvent
func callbackQDockWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDockWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDockWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDockWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDockWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDockWidgetCustomEvent
func callbackQDockWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDockWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
