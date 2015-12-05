package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QDockWidget_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::allowedAreas")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QDockWidget_AllowedAreas(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDockWidget) Features() QDockWidget__DockWidgetFeature {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::features")
		}
	}()

	if ptr.Pointer() != nil {
		return QDockWidget__DockWidgetFeature(C.QDockWidget_Features(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDockWidget) SetAllowedAreas(areas core.Qt__DockWidgetArea) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::setAllowedAreas")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_SetAllowedAreas(ptr.Pointer(), C.int(areas))
	}
}

func (ptr *QDockWidget) SetFeatures(features QDockWidget__DockWidgetFeature) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::setFeatures")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_SetFeatures(ptr.Pointer(), C.int(features))
	}
}

func (ptr *QDockWidget) SetFloating(floating bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::setFloating")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_SetFloating(ptr.Pointer(), C.int(qt.GoBoolToInt(floating)))
	}
}

func NewQDockWidget2(parent QWidget_ITF, flags core.Qt__WindowType) *QDockWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::QDockWidget")
		}
	}()

	return NewQDockWidgetFromPointer(C.QDockWidget_NewQDockWidget2(PointerFromQWidget(parent), C.int(flags)))
}

func NewQDockWidget(title string, parent QWidget_ITF, flags core.Qt__WindowType) *QDockWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::QDockWidget")
		}
	}()

	return NewQDockWidgetFromPointer(C.QDockWidget_NewQDockWidget(C.CString(title), PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QDockWidget) ConnectAllowedAreasChanged(f func(allowedAreas core.Qt__DockWidgetArea)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::allowedAreasChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectAllowedAreasChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "allowedAreasChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectAllowedAreasChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::allowedAreasChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectAllowedAreasChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "allowedAreasChanged")
	}
}

//export callbackQDockWidgetAllowedAreasChanged
func callbackQDockWidgetAllowedAreasChanged(ptrName *C.char, allowedAreas C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::allowedAreasChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged").(func(core.Qt__DockWidgetArea))(core.Qt__DockWidgetArea(allowedAreas))
}

func (ptr *QDockWidget) ConnectDockLocationChanged(f func(area core.Qt__DockWidgetArea)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::dockLocationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectDockLocationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dockLocationChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectDockLocationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::dockLocationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectDockLocationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dockLocationChanged")
	}
}

//export callbackQDockWidgetDockLocationChanged
func callbackQDockWidgetDockLocationChanged(ptrName *C.char, area C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::dockLocationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "dockLocationChanged").(func(core.Qt__DockWidgetArea))(core.Qt__DockWidgetArea(area))
}

func (ptr *QDockWidget) ConnectFeaturesChanged(f func(features QDockWidget__DockWidgetFeature)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::featuresChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectFeaturesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "featuresChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectFeaturesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::featuresChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectFeaturesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "featuresChanged")
	}
}

//export callbackQDockWidgetFeaturesChanged
func callbackQDockWidgetFeaturesChanged(ptrName *C.char, features C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::featuresChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "featuresChanged").(func(QDockWidget__DockWidgetFeature))(QDockWidget__DockWidgetFeature(features))
}

func (ptr *QDockWidget) IsAreaAllowed(area core.Qt__DockWidgetArea) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::isAreaAllowed")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDockWidget_IsAreaAllowed(ptr.Pointer(), C.int(area)) != 0
	}
	return false
}

func (ptr *QDockWidget) IsFloating() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::isFloating")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDockWidget_IsFloating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDockWidget) SetTitleBarWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::setTitleBarWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_SetTitleBarWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDockWidget) SetWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDockWidget) TitleBarWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::titleBarWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDockWidget_TitleBarWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDockWidget) ToggleViewAction() *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::toggleViewAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QDockWidget_ToggleViewAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDockWidget) ConnectTopLevelChanged(f func(topLevel bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::topLevelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectTopLevelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "topLevelChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectTopLevelChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::topLevelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectTopLevelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "topLevelChanged")
	}
}

//export callbackQDockWidgetTopLevelChanged
func callbackQDockWidgetTopLevelChanged(ptrName *C.char, topLevel C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::topLevelChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "topLevelChanged").(func(bool))(int(topLevel) != 0)
}

func (ptr *QDockWidget) ConnectVisibilityChanged(f func(visible bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::visibilityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectVisibilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectVisibilityChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::visibilityChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQDockWidgetVisibilityChanged
func callbackQDockWidgetVisibilityChanged(ptrName *C.char, visible C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::visibilityChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "visibilityChanged").(func(bool))(int(visible) != 0)
}

func (ptr *QDockWidget) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDockWidget_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDockWidget) DestroyQDockWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDockWidget::~QDockWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDockWidget_DestroyQDockWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
