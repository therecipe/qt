package multimediawidgets

//#include "multimediawidgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QVideoWidget struct {
	multimedia.QMediaBindableInterface
	widgets.QWidget
}

type QVideoWidget_ITF interface {
	multimedia.QMediaBindableInterface_ITF
	widgets.QWidget_ITF
	QVideoWidget_PTR() *QVideoWidget
}

func (p *QVideoWidget) Pointer() unsafe.Pointer {
	return p.QMediaBindableInterface_PTR().Pointer()
}

func (p *QVideoWidget) SetPointer(ptr unsafe.Pointer) {
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
	p.QWidget_PTR().SetPointer(ptr)
}

func PointerFromQVideoWidget(ptr QVideoWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidget_PTR().Pointer()
	}
	return nil
}

func NewQVideoWidgetFromPointer(ptr unsafe.Pointer) *QVideoWidget {
	var n = new(QVideoWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidget_") {
		n.SetObjectName("QVideoWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoWidget) QVideoWidget_PTR() *QVideoWidget {
	return ptr
}

func (ptr *QVideoWidget) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QVideoWidget::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidget_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Brightness() int {
	defer qt.Recovering("QVideoWidget::brightness")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Contrast() int {
	defer qt.Recovering("QVideoWidget::contrast")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Hue() int {
	defer qt.Recovering("QVideoWidget::hue")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QVideoWidget::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QVideoWidget_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) Saturation() int {
	defer qt.Recovering("QVideoWidget::saturation")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWidget::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWidget) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWidget::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidget) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWidget::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidget) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWidget::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidget) SetHue(hue int) {
	defer qt.Recovering("QVideoWidget::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidget) SetSaturation(saturation int) {
	defer qt.Recovering("QVideoWidget::setSaturation")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func NewQVideoWidget(parent widgets.QWidget_ITF) *QVideoWidget {
	defer qt.Recovering("QVideoWidget::QVideoWidget")

	return NewQVideoWidgetFromPointer(C.QVideoWidget_NewQVideoWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QVideoWidget) ConnectBrightnessChanged(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectBrightnessChanged() {
	defer qt.Recovering("disconnect QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetBrightnessChanged
func callbackQVideoWidgetBrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidget::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidget) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidget) ConnectContrastChanged(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectContrastChanged() {
	defer qt.Recovering("disconnect QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetContrastChanged
func callbackQVideoWidgetContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidget::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidget) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidget) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidget::event")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QVideoWidget) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectFullScreenChanged() {
	defer qt.Recovering("disconnect QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetFullScreenChanged
func callbackQVideoWidgetFullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidget::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidget) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQVideoWidgetHideEvent
func callbackQVideoWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QVideoWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QVideoWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectHueChanged(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidget::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectHueChanged() {
	defer qt.Recovering("disconnect QVideoWidget::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetHueChanged
func callbackQVideoWidgetHueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidget::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidget) HueChanged(hue int) {
	defer qt.Recovering("QVideoWidget::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HueChanged(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidget) IsFullScreen() bool {
	defer qt.Recovering("QVideoWidget::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQVideoWidgetMoveEvent
func callbackQVideoWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QVideoWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQVideoWidgetPaintEvent
func callbackQVideoWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QVideoWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QVideoWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQVideoWidgetResizeEvent
func callbackQVideoWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QVideoWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QVideoWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectSaturationChanged(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectSaturationChanged() {
	defer qt.Recovering("disconnect QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetSaturationChanged
func callbackQVideoWidgetSaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidget::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidget) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QVideoWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QVideoWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQVideoWidgetShowEvent
func callbackQVideoWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QVideoWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QVideoWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QVideoWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QVideoWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QVideoWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) DestroyQVideoWidget() {
	defer qt.Recovering("QVideoWidget::~QVideoWidget")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DestroyQVideoWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQVideoWidgetActionEvent
func callbackQVideoWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QVideoWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QVideoWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQVideoWidgetDragEnterEvent
func callbackQVideoWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QVideoWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQVideoWidgetDragLeaveEvent
func callbackQVideoWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QVideoWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQVideoWidgetDragMoveEvent
func callbackQVideoWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QVideoWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQVideoWidgetDropEvent
func callbackQVideoWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QVideoWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QVideoWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQVideoWidgetEnterEvent
func callbackQVideoWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQVideoWidgetFocusInEvent
func callbackQVideoWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQVideoWidgetFocusOutEvent
func callbackQVideoWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QVideoWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQVideoWidgetLeaveEvent
func callbackQVideoWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QVideoWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QVideoWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQVideoWidgetSetVisible
func callbackQVideoWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QVideoWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QVideoWidget) SetVisible(visible bool) {
	defer qt.Recovering("QVideoWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QVideoWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QVideoWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QVideoWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQVideoWidgetChangeEvent
func callbackQVideoWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQVideoWidgetCloseEvent
func callbackQVideoWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QVideoWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQVideoWidgetContextMenuEvent
func callbackQVideoWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QVideoWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QVideoWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QVideoWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QVideoWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QVideoWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQVideoWidgetInitPainter
func callbackQVideoWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQVideoWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QVideoWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QVideoWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QVideoWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QVideoWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QVideoWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQVideoWidgetInputMethodEvent
func callbackQVideoWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QVideoWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QVideoWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQVideoWidgetKeyPressEvent
func callbackQVideoWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQVideoWidgetKeyReleaseEvent
func callbackQVideoWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QVideoWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQVideoWidgetMouseDoubleClickEvent
func callbackQVideoWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQVideoWidgetMouseMoveEvent
func callbackQVideoWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQVideoWidgetMousePressEvent
func callbackQVideoWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQVideoWidgetMouseReleaseEvent
func callbackQVideoWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QVideoWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQVideoWidgetTabletEvent
func callbackQVideoWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QVideoWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QVideoWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQVideoWidgetWheelEvent
func callbackQVideoWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QVideoWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QVideoWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoWidgetTimerEvent
func callbackQVideoWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoWidgetChildEvent
func callbackQVideoWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoWidgetCustomEvent
func callbackQVideoWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
