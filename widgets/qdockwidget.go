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
