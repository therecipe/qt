package multimediawidgets

//#include "qvideowidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QVideoWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoWidget) QVideoWidget_PTR() *QVideoWidget {
	return ptr
}

func (ptr *QVideoWidget) AspectRatioMode() core.Qt__AspectRatioMode {
	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidget_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Brightness() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Contrast() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) Hue() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) MediaObject() *multimedia.QMediaObject {
	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QVideoWidget_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) Saturation() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidget) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWidget) SetBrightness(brightness int) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidget) SetContrast(contrast int) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidget) SetFullScreen(fullScreen bool) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidget) SetHue(hue int) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidget) SetSaturation(saturation int) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func NewQVideoWidget(parent widgets.QWidget_ITF) *QVideoWidget {
	return NewQVideoWidgetFromPointer(C.QVideoWidget_NewQVideoWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QVideoWidget) ConnectBrightnessChanged(f func(brightness int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectBrightnessChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetBrightnessChanged
func callbackQVideoWidgetBrightnessChanged(ptrName *C.char, brightness C.int) {
	qt.GetSignal(C.GoString(ptrName), "brightnessChanged").(func(int))(int(brightness))
}

func (ptr *QVideoWidget) ConnectContrastChanged(f func(contrast int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectContrastChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetContrastChanged
func callbackQVideoWidgetContrastChanged(ptrName *C.char, contrast C.int) {
	qt.GetSignal(C.GoString(ptrName), "contrastChanged").(func(int))(int(contrast))
}

func (ptr *QVideoWidget) ConnectFullScreenChanged(f func(fullScreen bool)) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectFullScreenChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetFullScreenChanged
func callbackQVideoWidgetFullScreenChanged(ptrName *C.char, fullScreen C.int) {
	qt.GetSignal(C.GoString(ptrName), "fullScreenChanged").(func(bool))(int(fullScreen) != 0)
}

func (ptr *QVideoWidget) ConnectHueChanged(f func(hue int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectHueChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetHueChanged
func callbackQVideoWidgetHueChanged(ptrName *C.char, hue C.int) {
	qt.GetSignal(C.GoString(ptrName), "hueChanged").(func(int))(int(hue))
}

func (ptr *QVideoWidget) IsFullScreen() bool {
	if ptr.Pointer() != nil {
		return C.QVideoWidget_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidget) ConnectSaturationChanged(f func(saturation int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidget) DisconnectSaturationChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetSaturationChanged
func callbackQVideoWidgetSaturationChanged(ptrName *C.char, saturation C.int) {
	qt.GetSignal(C.GoString(ptrName), "saturationChanged").(func(int))(int(saturation))
}

func (ptr *QVideoWidget) DestroyQVideoWidget() {
	if ptr.Pointer() != nil {
		C.QVideoWidget_DestroyQVideoWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
