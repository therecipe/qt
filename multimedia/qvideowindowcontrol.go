package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QVideoWindowControl struct {
	QMediaControl
}

type QVideoWindowControl_ITF interface {
	QMediaControl_ITF
	QVideoWindowControl_PTR() *QVideoWindowControl
}

func PointerFromQVideoWindowControl(ptr QVideoWindowControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWindowControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoWindowControlFromPointer(ptr unsafe.Pointer) *QVideoWindowControl {
	var n = new(QVideoWindowControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWindowControl_") {
		n.SetObjectName("QVideoWindowControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoWindowControl) QVideoWindowControl_PTR() *QVideoWindowControl {
	return ptr
}

func (ptr *QVideoWindowControl) AspectRatioMode() core.Qt__AspectRatioMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::aspectRatioMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWindowControl_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) Brightness() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::brightness")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectBrightnessChanged(f func(brightness int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::brightnessChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectBrightnessChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::brightnessChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWindowControlBrightnessChanged
func callbackQVideoWindowControlBrightnessChanged(ptrName *C.char, brightness C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::brightnessChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "brightnessChanged").(func(int))(int(brightness))
}

func (ptr *QVideoWindowControl) Contrast() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::contrast")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectContrastChanged(f func(contrast int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::contrastChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectContrastChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::contrastChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWindowControlContrastChanged
func callbackQVideoWindowControlContrastChanged(ptrName *C.char, contrast C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::contrastChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contrastChanged").(func(int))(int(contrast))
}

func (ptr *QVideoWindowControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::fullScreenChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectFullScreenChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::fullScreenChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWindowControlFullScreenChanged
func callbackQVideoWindowControlFullScreenChanged(ptrName *C.char, fullScreen C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::fullScreenChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "fullScreenChanged").(func(bool))(int(fullScreen) != 0)
}

func (ptr *QVideoWindowControl) Hue() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::hue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectHueChanged(f func(hue int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::hueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectHueChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::hueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWindowControlHueChanged
func callbackQVideoWindowControlHueChanged(ptrName *C.char, hue C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::hueChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "hueChanged").(func(int))(int(hue))
}

func (ptr *QVideoWindowControl) IsFullScreen() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::isFullScreen")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QVideoWindowControl_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWindowControl) ConnectNativeSizeChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::nativeSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectNativeSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nativeSizeChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectNativeSizeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::nativeSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectNativeSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nativeSizeChanged")
	}
}

//export callbackQVideoWindowControlNativeSizeChanged
func callbackQVideoWindowControlNativeSizeChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::nativeSizeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "nativeSizeChanged").(func())()
}

func (ptr *QVideoWindowControl) Repaint() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::repaint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_Repaint(ptr.Pointer())
	}
}

func (ptr *QVideoWindowControl) Saturation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::saturation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWindowControl_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWindowControl) ConnectSaturationChanged(f func(saturation int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::saturationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWindowControl) DisconnectSaturationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::saturationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWindowControlSaturationChanged
func callbackQVideoWindowControlSaturationChanged(ptrName *C.char, saturation C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::saturationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "saturationChanged").(func(int))(int(saturation))
}

func (ptr *QVideoWindowControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setAspectRatioMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWindowControl) SetBrightness(brightness int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setBrightness")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWindowControl) SetContrast(contrast int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setContrast")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWindowControl) SetDisplayRect(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setDisplayRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetDisplayRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QVideoWindowControl) SetFullScreen(fullScreen bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setFullScreen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWindowControl) SetHue(hue int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setHue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWindowControl) SetSaturation(saturation int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::setSaturation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWindowControl) DestroyQVideoWindowControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWindowControl::~QVideoWindowControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWindowControl_DestroyQVideoWindowControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
