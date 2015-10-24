package widgets

//#include "qdockwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDockWidget struct {
	QWidget
}

type QDockWidgetITF interface {
	QWidgetITF
	QDockWidgetPTR() *QDockWidget
}

func PointerFromQDockWidget(ptr QDockWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDockWidgetPTR().Pointer()
	}
	return nil
}

func QDockWidgetFromPointer(ptr unsafe.Pointer) *QDockWidget {
	var n = new(QDockWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDockWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDockWidget) QDockWidgetPTR() *QDockWidget {
	return ptr
}

//QDockWidget::DockWidgetFeature
type QDockWidget__DockWidgetFeature int

var (
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
	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QDockWidget_AllowedAreas(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDockWidget) Features() QDockWidget__DockWidgetFeature {
	if ptr.Pointer() != nil {
		return QDockWidget__DockWidgetFeature(C.QDockWidget_Features(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDockWidget) SetAllowedAreas(areas core.Qt__DockWidgetArea) {
	if ptr.Pointer() != nil {
		C.QDockWidget_SetAllowedAreas(C.QtObjectPtr(ptr.Pointer()), C.int(areas))
	}
}

func (ptr *QDockWidget) SetFeatures(features QDockWidget__DockWidgetFeature) {
	if ptr.Pointer() != nil {
		C.QDockWidget_SetFeatures(C.QtObjectPtr(ptr.Pointer()), C.int(features))
	}
}

func (ptr *QDockWidget) SetFloating(floating bool) {
	if ptr.Pointer() != nil {
		C.QDockWidget_SetFloating(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(floating)))
	}
}

func NewQDockWidget2(parent QWidgetITF, flags core.Qt__WindowType) *QDockWidget {
	return QDockWidgetFromPointer(unsafe.Pointer(C.QDockWidget_NewQDockWidget2(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func NewQDockWidget(title string, parent QWidgetITF, flags core.Qt__WindowType) *QDockWidget {
	return QDockWidgetFromPointer(unsafe.Pointer(C.QDockWidget_NewQDockWidget(C.CString(title), C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func (ptr *QDockWidget) ConnectAllowedAreasChanged(f func(allowedAreas core.Qt__DockWidgetArea)) {
	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectAllowedAreasChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "allowedAreasChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectAllowedAreasChanged() {
	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectAllowedAreasChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "allowedAreasChanged")
	}
}

//export callbackQDockWidgetAllowedAreasChanged
func callbackQDockWidgetAllowedAreasChanged(ptrName *C.char, allowedAreas C.int) {
	qt.GetSignal(C.GoString(ptrName), "allowedAreasChanged").(func(core.Qt__DockWidgetArea))(core.Qt__DockWidgetArea(allowedAreas))
}

func (ptr *QDockWidget) ConnectDockLocationChanged(f func(area core.Qt__DockWidgetArea)) {
	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectDockLocationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "dockLocationChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectDockLocationChanged() {
	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectDockLocationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "dockLocationChanged")
	}
}

//export callbackQDockWidgetDockLocationChanged
func callbackQDockWidgetDockLocationChanged(ptrName *C.char, area C.int) {
	qt.GetSignal(C.GoString(ptrName), "dockLocationChanged").(func(core.Qt__DockWidgetArea))(core.Qt__DockWidgetArea(area))
}

func (ptr *QDockWidget) ConnectFeaturesChanged(f func(features QDockWidget__DockWidgetFeature)) {
	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectFeaturesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "featuresChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectFeaturesChanged() {
	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectFeaturesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "featuresChanged")
	}
}

//export callbackQDockWidgetFeaturesChanged
func callbackQDockWidgetFeaturesChanged(ptrName *C.char, features C.int) {
	qt.GetSignal(C.GoString(ptrName), "featuresChanged").(func(QDockWidget__DockWidgetFeature))(QDockWidget__DockWidgetFeature(features))
}

func (ptr *QDockWidget) IsAreaAllowed(area core.Qt__DockWidgetArea) bool {
	if ptr.Pointer() != nil {
		return C.QDockWidget_IsAreaAllowed(C.QtObjectPtr(ptr.Pointer()), C.int(area)) != 0
	}
	return false
}

func (ptr *QDockWidget) IsFloating() bool {
	if ptr.Pointer() != nil {
		return C.QDockWidget_IsFloating(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDockWidget) SetTitleBarWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QDockWidget_SetTitleBarWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QDockWidget) SetWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QDockWidget_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QDockWidget) TitleBarWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QDockWidget_TitleBarWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDockWidget) ToggleViewAction() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QDockWidget_ToggleViewAction(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDockWidget) ConnectTopLevelChanged(f func(topLevel bool)) {
	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectTopLevelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "topLevelChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectTopLevelChanged() {
	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectTopLevelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "topLevelChanged")
	}
}

//export callbackQDockWidgetTopLevelChanged
func callbackQDockWidgetTopLevelChanged(ptrName *C.char, topLevel C.int) {
	qt.GetSignal(C.GoString(ptrName), "topLevelChanged").(func(bool))(int(topLevel) != 0)
}

func (ptr *QDockWidget) ConnectVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {
		C.QDockWidget_ConnectVisibilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "visibilityChanged", f)
	}
}

func (ptr *QDockWidget) DisconnectVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QDockWidget_DisconnectVisibilityChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "visibilityChanged")
	}
}

//export callbackQDockWidgetVisibilityChanged
func callbackQDockWidgetVisibilityChanged(ptrName *C.char, visible C.int) {
	qt.GetSignal(C.GoString(ptrName), "visibilityChanged").(func(bool))(int(visible) != 0)
}

func (ptr *QDockWidget) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QDockWidget_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDockWidget) DestroyQDockWidget() {
	if ptr.Pointer() != nil {
		C.QDockWidget_DestroyQDockWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
