package multimedia

//#include "qvideowindowcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVideoWindowControl struct {
	QMediaControl
}

type QVideoWindowControlITF interface {
	QMediaControlITF
	QVideoWindowControlPTR() *QVideoWindowControl
}

func PointerFromQVideoWindowControl(ptr QVideoWindowControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWindowControlPTR().Pointer()
	}
	return nil
}

func QVideoWindowControlFromPointer(ptr unsafe.Pointer) *QVideoWindowControl {
	var n = new(QVideoWindowControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVideoWindowControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoWindowControl) QVideoWindowControlPTR() *QVideoWindowControl {
	return ptr
}

func (ptr *QVideoWindowControl) AspectRatioMode() core.Qt__AspectRatioMode {
	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWindowControl_AspectRatioMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWindowControl) Brightness() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Brightness(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectBrightnessChanged(f func(brightness int)) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectBrightnessChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectBrightnessChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectBrightnessChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWindowControlBrightnessChanged
func callbackQVideoWindowControlBrightnessChanged(ptrName *C.char, brightness C.int) {
	qt.GetSignal(C.GoString(ptrName), "brightnessChanged").(func(int))(int(brightness))
}

func (ptr *QVideoWindowControl) Contrast() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Contrast(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectContrastChanged(f func(contrast int)) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectContrastChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectContrastChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectContrastChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWindowControlContrastChanged
func callbackQVideoWindowControlContrastChanged(ptrName *C.char, contrast C.int) {
	qt.GetSignal(C.GoString(ptrName), "contrastChanged").(func(int))(int(contrast))
}

func (ptr *QVideoWindowControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectFullScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectFullScreenChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectFullScreenChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWindowControlFullScreenChanged
func callbackQVideoWindowControlFullScreenChanged(ptrName *C.char, fullScreen C.int) {
	qt.GetSignal(C.GoString(ptrName), "fullScreenChanged").(func(bool))(int(fullScreen) != 0)
}

func (ptr *QVideoWindowControl) Hue() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Hue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectHueChanged(f func(hue int)) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectHueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectHueChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectHueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWindowControlHueChanged
func callbackQVideoWindowControlHueChanged(ptrName *C.char, hue C.int) {
	qt.GetSignal(C.GoString(ptrName), "hueChanged").(func(int))(int(hue))
}

func (ptr *QVideoWindowControl) IsFullScreen() bool {
	if ptr.Pointer() != nil {
		return C.QVideoWindowControl_IsFullScreen(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoWindowControl) ConnectNativeSizeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectNativeSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "nativeSizeChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectNativeSizeChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectNativeSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "nativeSizeChanged")
	}
}

//export callbackQVideoWindowControlNativeSizeChanged
func callbackQVideoWindowControlNativeSizeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "nativeSizeChanged").(func())()
}

func (ptr *QVideoWindowControl) Repaint() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_Repaint(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QVideoWindowControl) Saturation() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Saturation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectSaturationChanged(f func(saturation int)) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectSaturationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectSaturationChanged() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectSaturationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWindowControlSaturationChanged
func callbackQVideoWindowControlSaturationChanged(ptrName *C.char, saturation C.int) {
	qt.GetSignal(C.GoString(ptrName), "saturationChanged").(func(int))(int(saturation))
}

func (ptr *QVideoWindowControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetAspectRatioMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QVideoWindowControl) SetBrightness(brightness int) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetBrightness(C.QtObjectPtr(ptr.Pointer()), C.int(brightness))
	}
}

func (ptr *QVideoWindowControl) SetContrast(contrast int) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetContrast(C.QtObjectPtr(ptr.Pointer()), C.int(contrast))
	}
}

func (ptr *QVideoWindowControl) SetDisplayRect(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetDisplayRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QVideoWindowControl) SetFullScreen(fullScreen bool) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetFullScreen(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWindowControl) SetHue(hue int) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetHue(C.QtObjectPtr(ptr.Pointer()), C.int(hue))
	}
}

func (ptr *QVideoWindowControl) SetSaturation(saturation int) {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetSaturation(C.QtObjectPtr(ptr.Pointer()), C.int(saturation))
	}
}

func (ptr *QVideoWindowControl) DestroyQVideoWindowControl() {
	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DestroyQVideoWindowControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
