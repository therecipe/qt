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
func callbackQVideoWidgetBrightnessChanged(ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidget::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
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
func callbackQVideoWidgetContrastChanged(ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidget::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

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
func callbackQVideoWidgetFullScreenChanged(ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidget::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
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
func callbackQVideoWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetHueChanged(ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidget::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
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
func callbackQVideoWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetSaturationChanged(ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidget::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
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
func callbackQVideoWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QVideoWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQVideoWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQVideoWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQVideoWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVideoWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
