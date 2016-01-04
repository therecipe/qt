package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScreen struct {
	core.QObject
}

type QScreen_ITF interface {
	core.QObject_ITF
	QScreen_PTR() *QScreen
}

func PointerFromQScreen(ptr QScreen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScreen_PTR().Pointer()
	}
	return nil
}

func NewQScreenFromPointer(ptr unsafe.Pointer) *QScreen {
	var n = new(QScreen)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScreen_") {
		n.SetObjectName("QScreen_" + qt.Identifier())
	}
	return n
}

func (ptr *QScreen) QScreen_PTR() *QScreen {
	return ptr
}

func (ptr *QScreen) AvailableGeometry() *core.QRect {
	defer qt.Recovering("QScreen::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QScreen_AvailableGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) AvailableSize() *core.QSize {
	defer qt.Recovering("QScreen::availableSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScreen_AvailableSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) AvailableVirtualGeometry() *core.QRect {
	defer qt.Recovering("QScreen::availableVirtualGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QScreen_AvailableVirtualGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) AvailableVirtualSize() *core.QSize {
	defer qt.Recovering("QScreen::availableVirtualSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScreen_AvailableVirtualSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) Depth() int {
	defer qt.Recovering("QScreen::depth")

	if ptr.Pointer() != nil {
		return int(C.QScreen_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) DevicePixelRatio() float64 {
	defer qt.Recovering("QScreen::devicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Geometry() *core.QRect {
	defer qt.Recovering("QScreen::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QScreen_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) LogicalDotsPerInch() float64 {
	defer qt.Recovering("QScreen::logicalDotsPerInch")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_LogicalDotsPerInch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) LogicalDotsPerInchX() float64 {
	defer qt.Recovering("QScreen::logicalDotsPerInchX")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_LogicalDotsPerInchX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) LogicalDotsPerInchY() float64 {
	defer qt.Recovering("QScreen::logicalDotsPerInchY")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_LogicalDotsPerInchY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Name() string {
	defer qt.Recovering("QScreen::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScreen_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScreen) NativeOrientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::nativeOrientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_NativeOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Orientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PhysicalDotsPerInch() float64 {
	defer qt.Recovering("QScreen::physicalDotsPerInch")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_PhysicalDotsPerInch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PhysicalDotsPerInchX() float64 {
	defer qt.Recovering("QScreen::physicalDotsPerInchX")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_PhysicalDotsPerInchX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PhysicalDotsPerInchY() float64 {
	defer qt.Recovering("QScreen::physicalDotsPerInchY")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_PhysicalDotsPerInchY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PrimaryOrientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::primaryOrientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_PrimaryOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) RefreshRate() float64 {
	defer qt.Recovering("QScreen::refreshRate")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_RefreshRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Size() *core.QSize {
	defer qt.Recovering("QScreen::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScreen_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) VirtualGeometry() *core.QRect {
	defer qt.Recovering("QScreen::virtualGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QScreen_VirtualGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) VirtualSize() *core.QSize {
	defer qt.Recovering("QScreen::virtualSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScreen_VirtualSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScreen) AngleBetween(a core.Qt__ScreenOrientation, b core.Qt__ScreenOrientation) int {
	defer qt.Recovering("QScreen::angleBetween")

	if ptr.Pointer() != nil {
		return int(C.QScreen_AngleBetween(ptr.Pointer(), C.int(a), C.int(b)))
	}
	return 0
}

func (ptr *QScreen) ConnectAvailableGeometryChanged(f func(geometry *core.QRect)) {
	defer qt.Recovering("connect QScreen::availableGeometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectAvailableGeometryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableGeometryChanged", f)
	}
}

func (ptr *QScreen) DisconnectAvailableGeometryChanged() {
	defer qt.Recovering("disconnect QScreen::availableGeometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectAvailableGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableGeometryChanged")
	}
}

//export callbackQScreenAvailableGeometryChanged
func callbackQScreenAvailableGeometryChanged(ptr unsafe.Pointer, ptrName *C.char, geometry unsafe.Pointer) {
	defer qt.Recovering("callback QScreen::availableGeometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableGeometryChanged"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(geometry))
	}

}

func (ptr *QScreen) AvailableGeometryChanged(geometry core.QRect_ITF) {
	defer qt.Recovering("QScreen::availableGeometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_AvailableGeometryChanged(ptr.Pointer(), core.PointerFromQRect(geometry))
	}
}

func (ptr *QScreen) ConnectGeometryChanged(f func(geometry *core.QRect)) {
	defer qt.Recovering("connect QScreen::geometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectGeometryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QScreen) DisconnectGeometryChanged() {
	defer qt.Recovering("disconnect QScreen::geometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQScreenGeometryChanged
func callbackQScreenGeometryChanged(ptr unsafe.Pointer, ptrName *C.char, geometry unsafe.Pointer) {
	defer qt.Recovering("callback QScreen::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(geometry))
	}

}

func (ptr *QScreen) GeometryChanged(geometry core.QRect_ITF) {
	defer qt.Recovering("QScreen::geometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_GeometryChanged(ptr.Pointer(), core.PointerFromQRect(geometry))
	}
}

func (ptr *QScreen) IsLandscape(o core.Qt__ScreenOrientation) bool {
	defer qt.Recovering("QScreen::isLandscape")

	if ptr.Pointer() != nil {
		return C.QScreen_IsLandscape(ptr.Pointer(), C.int(o)) != 0
	}
	return false
}

func (ptr *QScreen) IsPortrait(o core.Qt__ScreenOrientation) bool {
	defer qt.Recovering("QScreen::isPortrait")

	if ptr.Pointer() != nil {
		return C.QScreen_IsPortrait(ptr.Pointer(), C.int(o)) != 0
	}
	return false
}

func (ptr *QScreen) ConnectLogicalDotsPerInchChanged(f func(dpi float64)) {
	defer qt.Recovering("connect QScreen::logicalDotsPerInchChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectLogicalDotsPerInchChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "logicalDotsPerInchChanged", f)
	}
}

func (ptr *QScreen) DisconnectLogicalDotsPerInchChanged() {
	defer qt.Recovering("disconnect QScreen::logicalDotsPerInchChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectLogicalDotsPerInchChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "logicalDotsPerInchChanged")
	}
}

//export callbackQScreenLogicalDotsPerInchChanged
func callbackQScreenLogicalDotsPerInchChanged(ptr unsafe.Pointer, ptrName *C.char, dpi C.double) {
	defer qt.Recovering("callback QScreen::logicalDotsPerInchChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "logicalDotsPerInchChanged"); signal != nil {
		signal.(func(float64))(float64(dpi))
	}

}

func (ptr *QScreen) LogicalDotsPerInchChanged(dpi float64) {
	defer qt.Recovering("QScreen::logicalDotsPerInchChanged")

	if ptr.Pointer() != nil {
		C.QScreen_LogicalDotsPerInchChanged(ptr.Pointer(), C.double(dpi))
	}
}

func (ptr *QScreen) MapBetween(a core.Qt__ScreenOrientation, b core.Qt__ScreenOrientation, rect core.QRect_ITF) *core.QRect {
	defer qt.Recovering("QScreen::mapBetween")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QScreen_MapBetween(ptr.Pointer(), C.int(a), C.int(b), core.PointerFromQRect(rect)))
	}
	return nil
}

func (ptr *QScreen) ConnectOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	defer qt.Recovering("connect QScreen::orientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "orientationChanged", f)
	}
}

func (ptr *QScreen) DisconnectOrientationChanged() {
	defer qt.Recovering("disconnect QScreen::orientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "orientationChanged")
	}
}

//export callbackQScreenOrientationChanged
func callbackQScreenOrientationChanged(ptr unsafe.Pointer, ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QScreen::orientationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "orientationChanged"); signal != nil {
		signal.(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
	}

}

func (ptr *QScreen) OrientationChanged(orientation core.Qt__ScreenOrientation) {
	defer qt.Recovering("QScreen::orientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_OrientationChanged(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QScreen) OrientationUpdateMask() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::orientationUpdateMask")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_OrientationUpdateMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) ConnectPhysicalDotsPerInchChanged(f func(dpi float64)) {
	defer qt.Recovering("connect QScreen::physicalDotsPerInchChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectPhysicalDotsPerInchChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "physicalDotsPerInchChanged", f)
	}
}

func (ptr *QScreen) DisconnectPhysicalDotsPerInchChanged() {
	defer qt.Recovering("disconnect QScreen::physicalDotsPerInchChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectPhysicalDotsPerInchChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "physicalDotsPerInchChanged")
	}
}

//export callbackQScreenPhysicalDotsPerInchChanged
func callbackQScreenPhysicalDotsPerInchChanged(ptr unsafe.Pointer, ptrName *C.char, dpi C.double) {
	defer qt.Recovering("callback QScreen::physicalDotsPerInchChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "physicalDotsPerInchChanged"); signal != nil {
		signal.(func(float64))(float64(dpi))
	}

}

func (ptr *QScreen) PhysicalDotsPerInchChanged(dpi float64) {
	defer qt.Recovering("QScreen::physicalDotsPerInchChanged")

	if ptr.Pointer() != nil {
		C.QScreen_PhysicalDotsPerInchChanged(ptr.Pointer(), C.double(dpi))
	}
}

func (ptr *QScreen) ConnectPrimaryOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	defer qt.Recovering("connect QScreen::primaryOrientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectPrimaryOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "primaryOrientationChanged", f)
	}
}

func (ptr *QScreen) DisconnectPrimaryOrientationChanged() {
	defer qt.Recovering("disconnect QScreen::primaryOrientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectPrimaryOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "primaryOrientationChanged")
	}
}

//export callbackQScreenPrimaryOrientationChanged
func callbackQScreenPrimaryOrientationChanged(ptr unsafe.Pointer, ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QScreen::primaryOrientationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "primaryOrientationChanged"); signal != nil {
		signal.(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
	}

}

func (ptr *QScreen) PrimaryOrientationChanged(orientation core.Qt__ScreenOrientation) {
	defer qt.Recovering("QScreen::primaryOrientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_PrimaryOrientationChanged(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QScreen) ConnectRefreshRateChanged(f func(refreshRate float64)) {
	defer qt.Recovering("connect QScreen::refreshRateChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectRefreshRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "refreshRateChanged", f)
	}
}

func (ptr *QScreen) DisconnectRefreshRateChanged() {
	defer qt.Recovering("disconnect QScreen::refreshRateChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectRefreshRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "refreshRateChanged")
	}
}

//export callbackQScreenRefreshRateChanged
func callbackQScreenRefreshRateChanged(ptr unsafe.Pointer, ptrName *C.char, refreshRate C.double) {
	defer qt.Recovering("callback QScreen::refreshRateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "refreshRateChanged"); signal != nil {
		signal.(func(float64))(float64(refreshRate))
	}

}

func (ptr *QScreen) RefreshRateChanged(refreshRate float64) {
	defer qt.Recovering("QScreen::refreshRateChanged")

	if ptr.Pointer() != nil {
		C.QScreen_RefreshRateChanged(ptr.Pointer(), C.double(refreshRate))
	}
}

func (ptr *QScreen) SetOrientationUpdateMask(mask core.Qt__ScreenOrientation) {
	defer qt.Recovering("QScreen::setOrientationUpdateMask")

	if ptr.Pointer() != nil {
		C.QScreen_SetOrientationUpdateMask(ptr.Pointer(), C.int(mask))
	}
}

func (ptr *QScreen) ConnectVirtualGeometryChanged(f func(rect *core.QRect)) {
	defer qt.Recovering("connect QScreen::virtualGeometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectVirtualGeometryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "virtualGeometryChanged", f)
	}
}

func (ptr *QScreen) DisconnectVirtualGeometryChanged() {
	defer qt.Recovering("disconnect QScreen::virtualGeometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectVirtualGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "virtualGeometryChanged")
	}
}

//export callbackQScreenVirtualGeometryChanged
func callbackQScreenVirtualGeometryChanged(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer) {
	defer qt.Recovering("callback QScreen::virtualGeometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "virtualGeometryChanged"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(rect))
	}

}

func (ptr *QScreen) VirtualGeometryChanged(rect core.QRect_ITF) {
	defer qt.Recovering("QScreen::virtualGeometryChanged")

	if ptr.Pointer() != nil {
		C.QScreen_VirtualGeometryChanged(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QScreen) DestroyQScreen() {
	defer qt.Recovering("QScreen::~QScreen")

	if ptr.Pointer() != nil {
		C.QScreen_DestroyQScreen(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScreen) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScreen::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScreen) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScreen::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQScreenTimerEvent
func callbackQScreenTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScreen::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScreenFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScreen) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScreen::timerEvent")

	if ptr.Pointer() != nil {
		C.QScreen_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScreen) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScreen::timerEvent")

	if ptr.Pointer() != nil {
		C.QScreen_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScreen) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScreen::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScreen) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScreen::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQScreenChildEvent
func callbackQScreenChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScreen::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScreenFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScreen) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScreen::childEvent")

	if ptr.Pointer() != nil {
		C.QScreen_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScreen) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScreen::childEvent")

	if ptr.Pointer() != nil {
		C.QScreen_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScreen) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScreen::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScreen) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScreen::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQScreenCustomEvent
func callbackQScreenCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScreen::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScreenFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScreen) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScreen::customEvent")

	if ptr.Pointer() != nil {
		C.QScreen_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScreen) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScreen::customEvent")

	if ptr.Pointer() != nil {
		C.QScreen_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
