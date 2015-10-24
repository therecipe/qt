package multimediawidgets

//#include "qvideowidgetcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QVideoWidgetControl struct {
	multimedia.QMediaControl
}

type QVideoWidgetControlITF interface {
	multimedia.QMediaControlITF
	QVideoWidgetControlPTR() *QVideoWidgetControl
}

func PointerFromQVideoWidgetControl(ptr QVideoWidgetControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidgetControlPTR().Pointer()
	}
	return nil
}

func QVideoWidgetControlFromPointer(ptr unsafe.Pointer) *QVideoWidgetControl {
	var n = new(QVideoWidgetControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVideoWidgetControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoWidgetControl) QVideoWidgetControlPTR() *QVideoWidgetControl {
	return ptr
}

func (ptr *QVideoWidgetControl) AspectRatioMode() core.Qt__AspectRatioMode {
	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidgetControl_AspectRatioMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWidgetControl) Brightness() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Brightness(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectBrightnessChanged(f func(brightness int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectBrightnessChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectBrightnessChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectBrightnessChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetControlBrightnessChanged
func callbackQVideoWidgetControlBrightnessChanged(ptrName *C.char, brightness C.int) {
	qt.GetSignal(C.GoString(ptrName), "brightnessChanged").(func(int))(int(brightness))
}

func (ptr *QVideoWidgetControl) Contrast() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Contrast(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectContrastChanged(f func(contrast int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectContrastChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectContrastChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectContrastChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetControlContrastChanged
func callbackQVideoWidgetControlContrastChanged(ptrName *C.char, contrast C.int) {
	qt.GetSignal(C.GoString(ptrName), "contrastChanged").(func(int))(int(contrast))
}

func (ptr *QVideoWidgetControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectFullScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectFullScreenChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectFullScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetControlFullScreenChanged
func callbackQVideoWidgetControlFullScreenChanged(ptrName *C.char, fullScreen C.int) {
	qt.GetSignal(C.GoString(ptrName), "fullScreenChanged").(func(bool))(int(fullScreen) != 0)
}

func (ptr *QVideoWidgetControl) Hue() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Hue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectHueChanged(f func(hue int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectHueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectHueChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectHueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetControlHueChanged
func callbackQVideoWidgetControlHueChanged(ptrName *C.char, hue C.int) {
	qt.GetSignal(C.GoString(ptrName), "hueChanged").(func(int))(int(hue))
}

func (ptr *QVideoWidgetControl) IsFullScreen() bool {
	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_IsFullScreen(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoWidgetControl) Saturation() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Saturation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectSaturationChanged(f func(saturation int)) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectSaturationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSaturationChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectSaturationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetControlSaturationChanged
func callbackQVideoWidgetControlSaturationChanged(ptrName *C.char, saturation C.int) {
	qt.GetSignal(C.GoString(ptrName), "saturationChanged").(func(int))(int(saturation))
}

func (ptr *QVideoWidgetControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetAspectRatioMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QVideoWidgetControl) SetBrightness(brightness int) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetBrightness(C.QtObjectPtr(ptr.Pointer()), C.int(brightness))
	}
}

func (ptr *QVideoWidgetControl) SetContrast(contrast int) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetContrast(C.QtObjectPtr(ptr.Pointer()), C.int(contrast))
	}
}

func (ptr *QVideoWidgetControl) SetFullScreen(fullScreen bool) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetFullScreen(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidgetControl) SetHue(hue int) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetHue(C.QtObjectPtr(ptr.Pointer()), C.int(hue))
	}
}

func (ptr *QVideoWidgetControl) SetSaturation(saturation int) {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetSaturation(C.QtObjectPtr(ptr.Pointer()), C.int(saturation))
	}
}

func (ptr *QVideoWidgetControl) VideoWidget() *widgets.QWidget {
	if ptr.Pointer() != nil {
		return widgets.QWidgetFromPointer(unsafe.Pointer(C.QVideoWidgetControl_VideoWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QVideoWidgetControl) DestroyQVideoWidgetControl() {
	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DestroyQVideoWidgetControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
