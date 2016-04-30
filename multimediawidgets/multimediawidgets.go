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

type QCameraViewfinder struct {
	QVideoWidget
}

type QCameraViewfinder_ITF interface {
	QVideoWidget_ITF
	QCameraViewfinder_PTR() *QCameraViewfinder
}

func (p *QCameraViewfinder) QCameraViewfinder_PTR() *QCameraViewfinder {
	return p
}

func (p *QCameraViewfinder) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QVideoWidget_PTR().Pointer()
	}
	return nil
}

func (p *QCameraViewfinder) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QVideoWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQCameraViewfinder(ptr QCameraViewfinder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinder_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderFromPointer(ptr unsafe.Pointer) *QCameraViewfinder {
	var n = new(QCameraViewfinder)
	n.SetPointer(ptr)
	return n
}

func newQCameraViewfinderFromPointer(ptr unsafe.Pointer) *QCameraViewfinder {
	var n = NewQCameraViewfinderFromPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinder_") {
		n.SetObjectName("QCameraViewfinder_" + qt.Identifier())
	}
	return n
}

func NewQCameraViewfinder(parent widgets.QWidget_ITF) *QCameraViewfinder {
	defer qt.Recovering("QCameraViewfinder::QCameraViewfinder")

	return newQCameraViewfinderFromPointer(C.QCameraViewfinder_NewQCameraViewfinder(widgets.PointerFromQWidget(parent)))
}

//export callbackQCameraViewfinder_MediaObject
func callbackQCameraViewfinder_MediaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCameraViewfinder::mediaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaObject"); signal != nil {
		return multimedia.PointerFromQMediaObject(signal.(func() *multimedia.QMediaObject)())
	}

	return multimedia.PointerFromQMediaObject(NewQCameraViewfinderFromPointer(ptr).MediaObjectDefault())
}

func (ptr *QCameraViewfinder) ConnectMediaObject(f func() *multimedia.QMediaObject) {
	defer qt.Recovering("connect QCameraViewfinder::mediaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mediaObject", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMediaObject() {
	defer qt.Recovering("disconnect QCameraViewfinder::mediaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mediaObject")
	}
}

func (ptr *QCameraViewfinder) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QCameraViewfinder::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QCameraViewfinder_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) MediaObjectDefault() *multimedia.QMediaObject {
	defer qt.Recovering("QCameraViewfinder::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QCameraViewfinder_MediaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCameraViewfinder_SetMediaObject
func callbackQCameraViewfinder_SetMediaObject(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCameraViewfinder::setMediaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMediaObject"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*multimedia.QMediaObject) bool)(multimedia.NewQMediaObjectFromPointer(object))))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).SetMediaObjectDefault(multimedia.NewQMediaObjectFromPointer(object))))
}

func (ptr *QCameraViewfinder) ConnectSetMediaObject(f func(object *multimedia.QMediaObject) bool) {
	defer qt.Recovering("connect QCameraViewfinder::setMediaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMediaObject", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetMediaObject() {
	defer qt.Recovering("disconnect QCameraViewfinder::setMediaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMediaObject")
	}
}

func (ptr *QCameraViewfinder) SetMediaObject(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_SetMediaObject(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) SetMediaObjectDefault(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_SetMediaObjectDefault(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) DestroyQCameraViewfinder() {
	defer qt.Recovering("QCameraViewfinder::~QCameraViewfinder")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DestroyQCameraViewfinder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQCameraViewfinder_SetAspectRatioMode
func callbackQCameraViewfinder_SetAspectRatioMode(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setAspectRatioMode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setAspectRatioMode"); signal != nil {
		signal.(func(core.Qt__AspectRatioMode))(core.Qt__AspectRatioMode(mode))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetAspectRatioModeDefault(core.Qt__AspectRatioMode(mode))
	}
}

func (ptr *QCameraViewfinder) ConnectSetAspectRatioMode(f func(mode core.Qt__AspectRatioMode)) {
	defer qt.Recovering("connect QCameraViewfinder::setAspectRatioMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setAspectRatioMode", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetAspectRatioMode() {
	defer qt.Recovering("disconnect QCameraViewfinder::setAspectRatioMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setAspectRatioMode")
	}
}

func (ptr *QCameraViewfinder) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QCameraViewfinder::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraViewfinder) SetAspectRatioModeDefault(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QCameraViewfinder::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetAspectRatioModeDefault(ptr.Pointer(), C.int(mode))
	}
}

//export callbackQCameraViewfinder_SetBrightness
func callbackQCameraViewfinder_SetBrightness(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setBrightness")

	if signal := qt.GetSignal(C.GoString(ptrName), "setBrightness"); signal != nil {
		signal.(func(int))(int(brightness))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetBrightnessDefault(int(brightness))
	}
}

func (ptr *QCameraViewfinder) ConnectSetBrightness(f func(brightness int)) {
	defer qt.Recovering("connect QCameraViewfinder::setBrightness")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setBrightness", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetBrightness() {
	defer qt.Recovering("disconnect QCameraViewfinder::setBrightness")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setBrightness")
	}
}

func (ptr *QCameraViewfinder) SetBrightness(brightness int) {
	defer qt.Recovering("QCameraViewfinder::setBrightness")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QCameraViewfinder) SetBrightnessDefault(brightness int) {
	defer qt.Recovering("QCameraViewfinder::setBrightness")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetBrightnessDefault(ptr.Pointer(), C.int(brightness))
	}
}

//export callbackQCameraViewfinder_SetContrast
func callbackQCameraViewfinder_SetContrast(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setContrast")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContrast"); signal != nil {
		signal.(func(int))(int(contrast))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetContrastDefault(int(contrast))
	}
}

func (ptr *QCameraViewfinder) ConnectSetContrast(f func(contrast int)) {
	defer qt.Recovering("connect QCameraViewfinder::setContrast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setContrast", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetContrast() {
	defer qt.Recovering("disconnect QCameraViewfinder::setContrast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setContrast")
	}
}

func (ptr *QCameraViewfinder) SetContrast(contrast int) {
	defer qt.Recovering("QCameraViewfinder::setContrast")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QCameraViewfinder) SetContrastDefault(contrast int) {
	defer qt.Recovering("QCameraViewfinder::setContrast")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetContrastDefault(ptr.Pointer(), C.int(contrast))
	}
}

//export callbackQCameraViewfinder_SetFullScreen
func callbackQCameraViewfinder_SetFullScreen(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFullScreen"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetFullScreenDefault(int(fullScreen) != 0)
	}
}

func (ptr *QCameraViewfinder) ConnectSetFullScreen(f func(fullScreen bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFullScreen", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetFullScreen() {
	defer qt.Recovering("disconnect QCameraViewfinder::setFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFullScreen")
	}
}

func (ptr *QCameraViewfinder) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QCameraViewfinder::setFullScreen")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QCameraViewfinder) SetFullScreenDefault(fullScreen bool) {
	defer qt.Recovering("QCameraViewfinder::setFullScreen")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetFullScreenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

//export callbackQCameraViewfinder_SetHue
func callbackQCameraViewfinder_SetHue(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setHue")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHue"); signal != nil {
		signal.(func(int))(int(hue))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetHueDefault(int(hue))
	}
}

func (ptr *QCameraViewfinder) ConnectSetHue(f func(hue int)) {
	defer qt.Recovering("connect QCameraViewfinder::setHue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHue", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetHue() {
	defer qt.Recovering("disconnect QCameraViewfinder::setHue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHue")
	}
}

func (ptr *QCameraViewfinder) SetHue(hue int) {
	defer qt.Recovering("QCameraViewfinder::setHue")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QCameraViewfinder) SetHueDefault(hue int) {
	defer qt.Recovering("QCameraViewfinder::setHue")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetHueDefault(ptr.Pointer(), C.int(hue))
	}
}

//export callbackQCameraViewfinder_SetSaturation
func callbackQCameraViewfinder_SetSaturation(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setSaturation")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSaturation"); signal != nil {
		signal.(func(int))(int(saturation))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetSaturationDefault(int(saturation))
	}
}

func (ptr *QCameraViewfinder) ConnectSetSaturation(f func(saturation int)) {
	defer qt.Recovering("connect QCameraViewfinder::setSaturation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSaturation", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetSaturation() {
	defer qt.Recovering("disconnect QCameraViewfinder::setSaturation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSaturation")
	}
}

func (ptr *QCameraViewfinder) SetSaturation(saturation int) {
	defer qt.Recovering("QCameraViewfinder::setSaturation")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QCameraViewfinder) SetSaturationDefault(saturation int) {
	defer qt.Recovering("QCameraViewfinder::setSaturation")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetSaturationDefault(ptr.Pointer(), C.int(saturation))
	}
}

//export callbackQCameraViewfinder_Event
func callbackQCameraViewfinder_Event(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCameraViewfinder::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event))))
}

func (ptr *QCameraViewfinder) ConnectEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QCameraViewfinder::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QCameraViewfinder) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::event")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) EventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::event")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCameraViewfinder_HideEvent
func callbackQCameraViewfinder_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *QCameraViewfinder) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QCameraViewfinder) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQCameraViewfinder_MoveEvent
func callbackQCameraViewfinder_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *QCameraViewfinder) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QCameraViewfinder) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQCameraViewfinder_PaintEvent
func callbackQCameraViewfinder_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

func (ptr *QCameraViewfinder) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QCameraViewfinder) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQCameraViewfinder_ResizeEvent
func callbackQCameraViewfinder_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *QCameraViewfinder) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QCameraViewfinder) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQCameraViewfinder_ShowEvent
func callbackQCameraViewfinder_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *QCameraViewfinder) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QCameraViewfinder) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQCameraViewfinder_SizeHint
func callbackQCameraViewfinder_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCameraViewfinder::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQCameraViewfinderFromPointer(ptr).SizeHintDefault())
}

func (ptr *QCameraViewfinder) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QCameraViewfinder::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QCameraViewfinder::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QCameraViewfinder) SizeHint() *core.QSize {
	defer qt.Recovering("QCameraViewfinder::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCameraViewfinder_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QCameraViewfinder::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCameraViewfinder_SizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCameraViewfinder_ActionEvent
func callbackQCameraViewfinder_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

func (ptr *QCameraViewfinder) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QCameraViewfinder) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQCameraViewfinder_DragEnterEvent
func callbackQCameraViewfinder_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *QCameraViewfinder) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QCameraViewfinder) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQCameraViewfinder_DragLeaveEvent
func callbackQCameraViewfinder_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *QCameraViewfinder) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QCameraViewfinder) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQCameraViewfinder_DragMoveEvent
func callbackQCameraViewfinder_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *QCameraViewfinder) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QCameraViewfinder) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQCameraViewfinder_DropEvent
func callbackQCameraViewfinder_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *QCameraViewfinder) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QCameraViewfinder) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQCameraViewfinder_EnterEvent
func callbackQCameraViewfinder_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

func (ptr *QCameraViewfinder) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCameraViewfinder_FocusInEvent
func callbackQCameraViewfinder_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *QCameraViewfinder) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCameraViewfinder) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQCameraViewfinder_FocusOutEvent
func callbackQCameraViewfinder_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *QCameraViewfinder) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QCameraViewfinder) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQCameraViewfinder_LeaveEvent
func callbackQCameraViewfinder_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

func (ptr *QCameraViewfinder) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCameraViewfinder_Metric
func callbackQCameraViewfinder_Metric(ptr unsafe.Pointer, ptrName *C.char, m C.int) C.int {
	defer qt.Recovering("callback QCameraViewfinder::metric")

	if signal := qt.GetSignal(C.GoString(ptrName), "metric"); signal != nil {
		return C.int(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m)))
	}

	return C.int(NewQCameraViewfinderFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m)))
}

func (ptr *QCameraViewfinder) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	defer qt.Recovering("connect QCameraViewfinder::metric")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metric", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMetric() {
	defer qt.Recovering("disconnect QCameraViewfinder::metric")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metric")
	}
}

func (ptr *QCameraViewfinder) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QCameraViewfinder::metric")

	if ptr.Pointer() != nil {
		return int(C.QCameraViewfinder_Metric(ptr.Pointer(), C.int(m)))
	}
	return 0
}

func (ptr *QCameraViewfinder) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QCameraViewfinder::metric")

	if ptr.Pointer() != nil {
		return int(C.QCameraViewfinder_MetricDefault(ptr.Pointer(), C.int(m)))
	}
	return 0
}

//export callbackQCameraViewfinder_MinimumSizeHint
func callbackQCameraViewfinder_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCameraViewfinder::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQCameraViewfinderFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QCameraViewfinder) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QCameraViewfinder::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QCameraViewfinder::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *QCameraViewfinder) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QCameraViewfinder::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCameraViewfinder_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QCameraViewfinder::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCameraViewfinder_MinimumSizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCameraViewfinder_PaintEngine
func callbackQCameraViewfinder_PaintEngine(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCameraViewfinder::paintEngine")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQCameraViewfinderFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QCameraViewfinder) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	defer qt.Recovering("connect QCameraViewfinder::paintEngine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEngine", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectPaintEngine() {
	defer qt.Recovering("disconnect QCameraViewfinder::paintEngine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEngine")
	}
}

func (ptr *QCameraViewfinder) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QCameraViewfinder::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QCameraViewfinder_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) PaintEngineDefault() *gui.QPaintEngine {
	defer qt.Recovering("QCameraViewfinder::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QCameraViewfinder_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQCameraViewfinder_SetEnabled
func callbackQCameraViewfinder_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *QCameraViewfinder) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QCameraViewfinder::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *QCameraViewfinder) SetEnabled(vbo bool) {
	defer qt.Recovering("QCameraViewfinder::setEnabled")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QCameraViewfinder) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QCameraViewfinder::setEnabled")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQCameraViewfinder_SetStyleSheet
func callbackQCameraViewfinder_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QCameraViewfinder) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QCameraViewfinder::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QCameraViewfinder::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *QCameraViewfinder) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QCameraViewfinder::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QCameraViewfinder) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QCameraViewfinder::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetStyleSheetDefault(ptr.Pointer(), C.CString(styleSheet))
	}
}

//export callbackQCameraViewfinder_SetVisible
func callbackQCameraViewfinder_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QCameraViewfinder) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *QCameraViewfinder) SetVisible(visible bool) {
	defer qt.Recovering("QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QCameraViewfinder) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQCameraViewfinder_SetWindowModified
func callbackQCameraViewfinder_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *QCameraViewfinder) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QCameraViewfinder::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *QCameraViewfinder) SetWindowModified(vbo bool) {
	defer qt.Recovering("QCameraViewfinder::setWindowModified")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QCameraViewfinder) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QCameraViewfinder::setWindowModified")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQCameraViewfinder_SetWindowTitle
func callbackQCameraViewfinder_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QCameraViewfinder) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QCameraViewfinder::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QCameraViewfinder::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *QCameraViewfinder) SetWindowTitle(vqs string) {
	defer qt.Recovering("QCameraViewfinder::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetWindowTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QCameraViewfinder) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QCameraViewfinder::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetWindowTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQCameraViewfinder_ChangeEvent
func callbackQCameraViewfinder_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *QCameraViewfinder) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCameraViewfinder_Close
func callbackQCameraViewfinder_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QCameraViewfinder::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).CloseDefault()))
}

func (ptr *QCameraViewfinder) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QCameraViewfinder::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectClose() {
	defer qt.Recovering("disconnect QCameraViewfinder::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QCameraViewfinder) Close() bool {
	defer qt.Recovering("QCameraViewfinder::close")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) CloseDefault() bool {
	defer qt.Recovering("QCameraViewfinder::close")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCameraViewfinder_CloseEvent
func callbackQCameraViewfinder_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *QCameraViewfinder) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QCameraViewfinder) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQCameraViewfinder_ContextMenuEvent
func callbackQCameraViewfinder_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *QCameraViewfinder) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QCameraViewfinder) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQCameraViewfinder_FocusNextPrevChild
func callbackQCameraViewfinder_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QCameraViewfinder::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QCameraViewfinder) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QCameraViewfinder::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QCameraViewfinder) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QCameraViewfinder::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QCameraViewfinder::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQCameraViewfinder_HasHeightForWidth
func callbackQCameraViewfinder_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QCameraViewfinder::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *QCameraViewfinder) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QCameraViewfinder::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QCameraViewfinder::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *QCameraViewfinder) HasHeightForWidth() bool {
	defer qt.Recovering("QCameraViewfinder::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QCameraViewfinder::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCameraViewfinder_HeightForWidth
func callbackQCameraViewfinder_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback QCameraViewfinder::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewQCameraViewfinderFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *QCameraViewfinder) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QCameraViewfinder::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QCameraViewfinder::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *QCameraViewfinder) HeightForWidth(w int) int {
	defer qt.Recovering("QCameraViewfinder::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QCameraViewfinder_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QCameraViewfinder) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QCameraViewfinder::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QCameraViewfinder_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackQCameraViewfinder_Hide
func callbackQCameraViewfinder_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).HideDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectHide(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectHide() {
	defer qt.Recovering("disconnect QCameraViewfinder::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QCameraViewfinder) Hide() {
	defer qt.Recovering("QCameraViewfinder::hide")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_Hide(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) HideDefault() {
	defer qt.Recovering("QCameraViewfinder::hide")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_HideDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_InitPainter
func callbackQCameraViewfinder_InitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQCameraViewfinderFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QCameraViewfinder) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

func (ptr *QCameraViewfinder) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QCameraViewfinder) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQCameraViewfinder_InputMethodEvent
func callbackQCameraViewfinder_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *QCameraViewfinder) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QCameraViewfinder) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQCameraViewfinder_InputMethodQuery
func callbackQCameraViewfinder_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QCameraViewfinder::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQCameraViewfinderFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QCameraViewfinder) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QCameraViewfinder::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QCameraViewfinder::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QCameraViewfinder) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QCameraViewfinder::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraViewfinder_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QCameraViewfinder) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QCameraViewfinder::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraViewfinder_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQCameraViewfinder_KeyPressEvent
func callbackQCameraViewfinder_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *QCameraViewfinder) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCameraViewfinder) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQCameraViewfinder_KeyReleaseEvent
func callbackQCameraViewfinder_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *QCameraViewfinder) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QCameraViewfinder) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQCameraViewfinder_Lower
func callbackQCameraViewfinder_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectLower(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectLower() {
	defer qt.Recovering("disconnect QCameraViewfinder::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QCameraViewfinder) Lower() {
	defer qt.Recovering("QCameraViewfinder::lower")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_Lower(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) LowerDefault() {
	defer qt.Recovering("QCameraViewfinder::lower")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_LowerDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_MouseDoubleClickEvent
func callbackQCameraViewfinder_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *QCameraViewfinder) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQCameraViewfinder_MouseMoveEvent
func callbackQCameraViewfinder_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *QCameraViewfinder) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQCameraViewfinder_MousePressEvent
func callbackQCameraViewfinder_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *QCameraViewfinder) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQCameraViewfinder_MouseReleaseEvent
func callbackQCameraViewfinder_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *QCameraViewfinder) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QCameraViewfinder) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQCameraViewfinder_NativeEvent
func callbackQCameraViewfinder_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QCameraViewfinder::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QCameraViewfinder) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QCameraViewfinder::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QCameraViewfinder) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QCameraViewfinder::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QCameraViewfinder::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQCameraViewfinder_Raise
func callbackQCameraViewfinder_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectRaise(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectRaise() {
	defer qt.Recovering("disconnect QCameraViewfinder::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QCameraViewfinder) Raise() {
	defer qt.Recovering("QCameraViewfinder::raise")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_Raise(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) RaiseDefault() {
	defer qt.Recovering("QCameraViewfinder::raise")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_Repaint
func callbackQCameraViewfinder_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectRepaint() {
	defer qt.Recovering("disconnect QCameraViewfinder::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *QCameraViewfinder) Repaint() {
	defer qt.Recovering("QCameraViewfinder::repaint")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_Repaint(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) RepaintDefault() {
	defer qt.Recovering("QCameraViewfinder::repaint")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_SetDisabled
func callbackQCameraViewfinder_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *QCameraViewfinder) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QCameraViewfinder::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *QCameraViewfinder) SetDisabled(disable bool) {
	defer qt.Recovering("QCameraViewfinder::setDisabled")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QCameraViewfinder) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QCameraViewfinder::setDisabled")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackQCameraViewfinder_SetFocus2
func callbackQCameraViewfinder_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QCameraViewfinder) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QCameraViewfinder::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *QCameraViewfinder) SetFocus2() {
	defer qt.Recovering("QCameraViewfinder::setFocus")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) SetFocus2Default() {
	defer qt.Recovering("QCameraViewfinder::setFocus")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_SetHidden
func callbackQCameraViewfinder_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback QCameraViewfinder::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewQCameraViewfinderFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *QCameraViewfinder) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QCameraViewfinder::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *QCameraViewfinder) SetHidden(hidden bool) {
	defer qt.Recovering("QCameraViewfinder::setHidden")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QCameraViewfinder) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QCameraViewfinder::setHidden")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackQCameraViewfinder_Show
func callbackQCameraViewfinder_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectShow(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShow() {
	defer qt.Recovering("disconnect QCameraViewfinder::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QCameraViewfinder) Show() {
	defer qt.Recovering("QCameraViewfinder::show")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_Show(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) ShowDefault() {
	defer qt.Recovering("QCameraViewfinder::show")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_ShowFullScreen
func callbackQCameraViewfinder_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QCameraViewfinder::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QCameraViewfinder) ShowFullScreen() {
	defer qt.Recovering("QCameraViewfinder::showFullScreen")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) ShowFullScreenDefault() {
	defer qt.Recovering("QCameraViewfinder::showFullScreen")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_ShowMaximized
func callbackQCameraViewfinder_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QCameraViewfinder::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QCameraViewfinder) ShowMaximized() {
	defer qt.Recovering("QCameraViewfinder::showMaximized")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) ShowMaximizedDefault() {
	defer qt.Recovering("QCameraViewfinder::showMaximized")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_ShowMinimized
func callbackQCameraViewfinder_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QCameraViewfinder::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QCameraViewfinder) ShowMinimized() {
	defer qt.Recovering("QCameraViewfinder::showMinimized")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) ShowMinimizedDefault() {
	defer qt.Recovering("QCameraViewfinder::showMinimized")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_ShowNormal
func callbackQCameraViewfinder_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QCameraViewfinder::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QCameraViewfinder) ShowNormal() {
	defer qt.Recovering("QCameraViewfinder::showNormal")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) ShowNormalDefault() {
	defer qt.Recovering("QCameraViewfinder::showNormal")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_TabletEvent
func callbackQCameraViewfinder_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

func (ptr *QCameraViewfinder) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QCameraViewfinder) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQCameraViewfinder_Update
func callbackQCameraViewfinder_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectUpdate() {
	defer qt.Recovering("disconnect QCameraViewfinder::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QCameraViewfinder) Update() {
	defer qt.Recovering("QCameraViewfinder::update")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_Update(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) UpdateDefault() {
	defer qt.Recovering("QCameraViewfinder::update")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_UpdateMicroFocus
func callbackQCameraViewfinder_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QCameraViewfinder::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QCameraViewfinder) UpdateMicroFocus() {
	defer qt.Recovering("QCameraViewfinder::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QCameraViewfinder) UpdateMicroFocusDefault() {
	defer qt.Recovering("QCameraViewfinder::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQCameraViewfinder_WheelEvent
func callbackQCameraViewfinder_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *QCameraViewfinder) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QCameraViewfinder) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQCameraViewfinder_TimerEvent
func callbackQCameraViewfinder_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QCameraViewfinder) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinder) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCameraViewfinder_ChildEvent
func callbackQCameraViewfinder_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QCameraViewfinder) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinder) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCameraViewfinder_ConnectNotify
func callbackQCameraViewfinder_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCameraViewfinderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCameraViewfinder) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QCameraViewfinder::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QCameraViewfinder::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QCameraViewfinder) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCameraViewfinder::connectNotify")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCameraViewfinder) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCameraViewfinder::connectNotify")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCameraViewfinder_CustomEvent
func callbackQCameraViewfinder_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QCameraViewfinder) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinder) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCameraViewfinder_DeleteLater
func callbackQCameraViewfinder_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCameraViewfinder::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCameraViewfinderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCameraViewfinder) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QCameraViewfinder::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QCameraViewfinder::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QCameraViewfinder) DeleteLater() {
	defer qt.Recovering("QCameraViewfinder::deleteLater")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinder) DeleteLaterDefault() {
	defer qt.Recovering("QCameraViewfinder::deleteLater")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQCameraViewfinder_DisconnectNotify
func callbackQCameraViewfinder_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinder::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCameraViewfinderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCameraViewfinder) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QCameraViewfinder::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QCameraViewfinder::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QCameraViewfinder) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCameraViewfinder::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCameraViewfinder) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCameraViewfinder::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCameraViewfinder_EventFilter
func callbackQCameraViewfinder_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCameraViewfinder::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQCameraViewfinderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QCameraViewfinder) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QCameraViewfinder::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QCameraViewfinder::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QCameraViewfinder) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCameraViewfinder) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCameraViewfinder::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCameraViewfinder_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCameraViewfinder_MetaObject
func callbackQCameraViewfinder_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCameraViewfinder::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCameraViewfinderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCameraViewfinder) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QCameraViewfinder::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QCameraViewfinder::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QCameraViewfinder) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QCameraViewfinder::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCameraViewfinder_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QCameraViewfinder::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCameraViewfinder_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QGraphicsVideoItem struct {
	multimedia.QMediaBindableInterface
	widgets.QGraphicsObject
}

type QGraphicsVideoItem_ITF interface {
	multimedia.QMediaBindableInterface_ITF
	widgets.QGraphicsObject_ITF
	QGraphicsVideoItem_PTR() *QGraphicsVideoItem
}

func (p *QGraphicsVideoItem) QGraphicsVideoItem_PTR() *QGraphicsVideoItem {
	return p
}

func (p *QGraphicsVideoItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QMediaBindableInterface_PTR().Pointer()
	}
	return nil
}

func (p *QGraphicsVideoItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QMediaBindableInterface_PTR().SetPointer(ptr)
		p.QGraphicsObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQGraphicsVideoItem(ptr QGraphicsVideoItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsVideoItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsVideoItemFromPointer(ptr unsafe.Pointer) *QGraphicsVideoItem {
	var n = new(QGraphicsVideoItem)
	n.SetPointer(ptr)
	return n
}

func newQGraphicsVideoItemFromPointer(ptr unsafe.Pointer) *QGraphicsVideoItem {
	var n = NewQGraphicsVideoItemFromPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsVideoItem_") {
		n.SetObjectName("QGraphicsVideoItem_" + qt.Identifier())
	}
	return n
}

//export callbackQGraphicsVideoItem_NativeSizeChanged
func callbackQGraphicsVideoItem_NativeSizeChanged(ptr unsafe.Pointer, ptrName *C.char, size unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::nativeSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeSizeChanged"); signal != nil {
		signal.(func(*core.QSizeF))(core.NewQSizeFFromPointer(size))
	}

}

func (ptr *QGraphicsVideoItem) ConnectNativeSizeChanged(f func(size *core.QSizeF)) {
	defer qt.Recovering("connect QGraphicsVideoItem::nativeSizeChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ConnectNativeSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nativeSizeChanged", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectNativeSizeChanged() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::nativeSizeChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DisconnectNativeSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nativeSizeChanged")
	}
}

func (ptr *QGraphicsVideoItem) NativeSizeChanged(size core.QSizeF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::nativeSizeChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_NativeSizeChanged(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func NewQGraphicsVideoItem(parent widgets.QGraphicsItem_ITF) *QGraphicsVideoItem {
	defer qt.Recovering("QGraphicsVideoItem::QGraphicsVideoItem")

	return newQGraphicsVideoItemFromPointer(C.QGraphicsVideoItem_NewQGraphicsVideoItem(widgets.PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsVideoItem) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QGraphicsVideoItem::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QGraphicsVideoItem_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsVideoItem_BoundingRect
func callbackQGraphicsVideoItem_BoundingRect(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::boundingRect")

	if signal := qt.GetSignal(C.GoString(ptrName), "boundingRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQGraphicsVideoItemFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsVideoItem) ConnectBoundingRect(f func() *core.QRectF) {
	defer qt.Recovering("connect QGraphicsVideoItem::boundingRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "boundingRect", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectBoundingRect() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::boundingRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "boundingRect")
	}
}

func (ptr *QGraphicsVideoItem) BoundingRect() *core.QRectF {
	defer qt.Recovering("QGraphicsVideoItem::boundingRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QGraphicsVideoItem_BoundingRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) BoundingRectDefault() *core.QRectF {
	defer qt.Recovering("QGraphicsVideoItem::boundingRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QGraphicsVideoItem_BoundingRectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsVideoItem_MediaObject
func callbackQGraphicsVideoItem_MediaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::mediaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaObject"); signal != nil {
		return multimedia.PointerFromQMediaObject(signal.(func() *multimedia.QMediaObject)())
	}

	return multimedia.PointerFromQMediaObject(NewQGraphicsVideoItemFromPointer(ptr).MediaObjectDefault())
}

func (ptr *QGraphicsVideoItem) ConnectMediaObject(f func() *multimedia.QMediaObject) {
	defer qt.Recovering("connect QGraphicsVideoItem::mediaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mediaObject", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectMediaObject() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::mediaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mediaObject")
	}
}

func (ptr *QGraphicsVideoItem) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QGraphicsVideoItem::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QGraphicsVideoItem_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) MediaObjectDefault() *multimedia.QMediaObject {
	defer qt.Recovering("QGraphicsVideoItem::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QGraphicsVideoItem_MediaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) NativeSize() *core.QSizeF {
	defer qt.Recovering("QGraphicsVideoItem::nativeSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFFromPointer(C.QGraphicsVideoItem_NativeSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) Offset() *core.QPointF {
	defer qt.Recovering("QGraphicsVideoItem::offset")

	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QGraphicsVideoItem_Offset(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsVideoItem_Paint
func callbackQGraphicsVideoItem_Paint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsVideoItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

func (ptr *QGraphicsVideoItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsVideoItem) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsVideoItem) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsVideoItem::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsVideoItem) SetOffset(offset core.QPointF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetOffset(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QGraphicsVideoItem) SetSize(size core.QSizeF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::setSize")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsVideoItem) Size() *core.QSizeF {
	defer qt.Recovering("QGraphicsVideoItem::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFFromPointer(C.QGraphicsVideoItem_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) DestroyQGraphicsVideoItem() {
	defer qt.Recovering("QGraphicsVideoItem::~QGraphicsVideoItem")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DestroyQGraphicsVideoItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsVideoItem_SetMediaObject
func callbackQGraphicsVideoItem_SetMediaObject(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::setMediaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMediaObject"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*multimedia.QMediaObject) bool)(multimedia.NewQMediaObjectFromPointer(object))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).SetMediaObjectDefault(multimedia.NewQMediaObjectFromPointer(object))))
}

func (ptr *QGraphicsVideoItem) ConnectSetMediaObject(f func(object *multimedia.QMediaObject) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::setMediaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMediaObject", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectSetMediaObject() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::setMediaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMediaObject")
	}
}

func (ptr *QGraphicsVideoItem) SetMediaObject(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_SetMediaObject(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) SetMediaObjectDefault(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_SetMediaObjectDefault(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_Event
func callbackQGraphicsVideoItem_Event(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev))))
}

func (ptr *QGraphicsVideoItem) ConnectEvent(f func(ev *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QGraphicsVideoItem) Event(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) EventDefault(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_UpdateMicroFocus
func callbackQGraphicsVideoItem_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsVideoItem::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QGraphicsVideoItem) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QGraphicsVideoItem::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QGraphicsVideoItem) UpdateMicroFocus() {
	defer qt.Recovering("QGraphicsVideoItem::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QGraphicsVideoItem) UpdateMicroFocusDefault() {
	defer qt.Recovering("QGraphicsVideoItem::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsVideoItem_Advance
func callbackQGraphicsVideoItem_Advance(ptr unsafe.Pointer, ptrName *C.char, phase C.int) {
	defer qt.Recovering("callback QGraphicsVideoItem::advance")

	if signal := qt.GetSignal(C.GoString(ptrName), "advance"); signal != nil {
		signal.(func(int))(int(phase))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).AdvanceDefault(int(phase))
	}
}

func (ptr *QGraphicsVideoItem) ConnectAdvance(f func(phase int)) {
	defer qt.Recovering("connect QGraphicsVideoItem::advance")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "advance", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectAdvance() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::advance")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "advance")
	}
}

func (ptr *QGraphicsVideoItem) Advance(phase int) {
	defer qt.Recovering("QGraphicsVideoItem::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_Advance(ptr.Pointer(), C.int(phase))
	}
}

func (ptr *QGraphicsVideoItem) AdvanceDefault(phase int) {
	defer qt.Recovering("QGraphicsVideoItem::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_AdvanceDefault(ptr.Pointer(), C.int(phase))
	}
}

//export callbackQGraphicsVideoItem_CollidesWithItem
func callbackQGraphicsVideoItem_CollidesWithItem(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer, mode C.int) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::collidesWithItem")

	if signal := qt.GetSignal(C.GoString(ptrName), "collidesWithItem"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode))))
}

func (ptr *QGraphicsVideoItem) ConnectCollidesWithItem(f func(other *widgets.QGraphicsItem, mode core.Qt__ItemSelectionMode) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::collidesWithItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "collidesWithItem", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectCollidesWithItem() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::collidesWithItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "collidesWithItem")
	}
}

func (ptr *QGraphicsVideoItem) CollidesWithItem(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsVideoItem::collidesWithItem")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_CollidesWithItem(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.int(mode)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsVideoItem::collidesWithItem")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.int(mode)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_CollidesWithPath
func callbackQGraphicsVideoItem_CollidesWithPath(ptr unsafe.Pointer, ptrName *C.char, path unsafe.Pointer, mode C.int) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::collidesWithPath")

	if signal := qt.GetSignal(C.GoString(ptrName), "collidesWithPath"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode))))
}

func (ptr *QGraphicsVideoItem) ConnectCollidesWithPath(f func(path *gui.QPainterPath, mode core.Qt__ItemSelectionMode) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::collidesWithPath")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "collidesWithPath", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectCollidesWithPath() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::collidesWithPath")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "collidesWithPath")
	}
}

func (ptr *QGraphicsVideoItem) CollidesWithPath(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsVideoItem::collidesWithPath")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_CollidesWithPath(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(mode)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsVideoItem::collidesWithPath")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(mode)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_Contains
func callbackQGraphicsVideoItem_Contains(ptr unsafe.Pointer, ptrName *C.char, point unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::contains")

	if signal := qt.GetSignal(C.GoString(ptrName), "contains"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point))))
}

func (ptr *QGraphicsVideoItem) ConnectContains(f func(point *core.QPointF) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::contains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contains", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectContains() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::contains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contains")
	}
}

func (ptr *QGraphicsVideoItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) ContainsDefault(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_ContextMenuEvent
func callbackQGraphicsVideoItem_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneContextMenuEvent))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectContextMenuEvent(f func(event *widgets.QGraphicsSceneContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *QGraphicsVideoItem) ContextMenuEvent(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ContextMenuEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

//export callbackQGraphicsVideoItem_DragEnterEvent
func callbackQGraphicsVideoItem_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectDragEnterEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *QGraphicsVideoItem) DragEnterEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DragEnterEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsVideoItem_DragLeaveEvent
func callbackQGraphicsVideoItem_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectDragLeaveEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *QGraphicsVideoItem) DragLeaveEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DragLeaveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsVideoItem_DragMoveEvent
func callbackQGraphicsVideoItem_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectDragMoveEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *QGraphicsVideoItem) DragMoveEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DragMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsVideoItem_DropEvent
func callbackQGraphicsVideoItem_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectDropEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *QGraphicsVideoItem) DropEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DropEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsVideoItem_FocusInEvent
func callbackQGraphicsVideoItem_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *QGraphicsVideoItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQGraphicsVideoItem_FocusOutEvent
func callbackQGraphicsVideoItem_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *QGraphicsVideoItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQGraphicsVideoItem_HoverEnterEvent
func callbackQGraphicsVideoItem_HoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectHoverEnterEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

func (ptr *QGraphicsVideoItem) HoverEnterEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_HoverEnterEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsVideoItem_HoverLeaveEvent
func callbackQGraphicsVideoItem_HoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectHoverLeaveEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

func (ptr *QGraphicsVideoItem) HoverLeaveEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_HoverLeaveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsVideoItem_HoverMoveEvent
func callbackQGraphicsVideoItem_HoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectHoverMoveEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

func (ptr *QGraphicsVideoItem) HoverMoveEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_HoverMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsVideoItem_InputMethodEvent
func callbackQGraphicsVideoItem_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *QGraphicsVideoItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQGraphicsVideoItem_InputMethodQuery
func callbackQGraphicsVideoItem_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQGraphicsVideoItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QGraphicsVideoItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QGraphicsVideoItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QGraphicsVideoItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsVideoItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsVideoItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsVideoItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsVideoItem_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQGraphicsVideoItem_IsObscuredBy
func callbackQGraphicsVideoItem_IsObscuredBy(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::isObscuredBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "isObscuredBy"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem) bool)(widgets.NewQGraphicsItemFromPointer(item))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item))))
}

func (ptr *QGraphicsVideoItem) ConnectIsObscuredBy(f func(item *widgets.QGraphicsItem) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::isObscuredBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isObscuredBy", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectIsObscuredBy() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::isObscuredBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isObscuredBy")
	}
}

func (ptr *QGraphicsVideoItem) IsObscuredBy(item widgets.QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_IsObscuredBy(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_ItemChange
func callbackQGraphicsVideoItem_ItemChange(ptr unsafe.Pointer, ptrName *C.char, change C.int, value unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::itemChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemChange"); signal != nil {
		return core.PointerFromQVariant(signal.(func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQGraphicsVideoItemFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QGraphicsVideoItem) ConnectItemChange(f func(change widgets.QGraphicsItem__GraphicsItemChange, value *core.QVariant) *core.QVariant) {
	defer qt.Recovering("connect QGraphicsVideoItem::itemChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "itemChange", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectItemChange() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::itemChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "itemChange")
	}
}

func (ptr *QGraphicsVideoItem) ItemChange(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsVideoItem::itemChange")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsVideoItem_ItemChange(ptr.Pointer(), C.int(change), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsVideoItem::itemChange")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsVideoItem_ItemChangeDefault(ptr.Pointer(), C.int(change), core.PointerFromQVariant(value)))
	}
	return nil
}

//export callbackQGraphicsVideoItem_KeyPressEvent
func callbackQGraphicsVideoItem_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *QGraphicsVideoItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsVideoItem_KeyReleaseEvent
func callbackQGraphicsVideoItem_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *QGraphicsVideoItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsVideoItem_MouseDoubleClickEvent
func callbackQGraphicsVideoItem_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectMouseDoubleClickEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *QGraphicsVideoItem) MouseDoubleClickEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MouseDoubleClickEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsVideoItem_MouseMoveEvent
func callbackQGraphicsVideoItem_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectMouseMoveEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *QGraphicsVideoItem) MouseMoveEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MouseMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsVideoItem_MousePressEvent
func callbackQGraphicsVideoItem_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectMousePressEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *QGraphicsVideoItem) MousePressEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MousePressEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsVideoItem_MouseReleaseEvent
func callbackQGraphicsVideoItem_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectMouseReleaseEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *QGraphicsVideoItem) MouseReleaseEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MouseReleaseEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsVideoItem_OpaqueArea
func callbackQGraphicsVideoItem_OpaqueArea(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::opaqueArea")

	if signal := qt.GetSignal(C.GoString(ptrName), "opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsVideoItemFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QGraphicsVideoItem) ConnectOpaqueArea(f func() *gui.QPainterPath) {
	defer qt.Recovering("connect QGraphicsVideoItem::opaqueArea")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "opaqueArea", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectOpaqueArea() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::opaqueArea")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "opaqueArea")
	}
}

func (ptr *QGraphicsVideoItem) OpaqueArea() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsVideoItem::opaqueArea")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsVideoItem_OpaqueArea(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) OpaqueAreaDefault() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsVideoItem::opaqueArea")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsVideoItem_OpaqueAreaDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsVideoItem_SceneEvent
func callbackQGraphicsVideoItem_SceneEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::sceneEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsVideoItem) ConnectSceneEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::sceneEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sceneEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectSceneEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::sceneEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sceneEvent")
	}
}

func (ptr *QGraphicsVideoItem) SceneEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_SceneEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) SceneEventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_SceneEventFilter
func callbackQGraphicsVideoItem_SceneEventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::sceneEventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneEventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, *core.QEvent) bool)(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsVideoItem) ConnectSceneEventFilter(f func(watched *widgets.QGraphicsItem, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::sceneEventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sceneEventFilter", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectSceneEventFilter() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::sceneEventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sceneEventFilter")
	}
}

func (ptr *QGraphicsVideoItem) SceneEventFilter(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::sceneEventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_SceneEventFilter(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::sceneEventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_Shape
func callbackQGraphicsVideoItem_Shape(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::shape")

	if signal := qt.GetSignal(C.GoString(ptrName), "shape"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsVideoItemFromPointer(ptr).ShapeDefault())
}

func (ptr *QGraphicsVideoItem) ConnectShape(f func() *gui.QPainterPath) {
	defer qt.Recovering("connect QGraphicsVideoItem::shape")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "shape", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectShape() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::shape")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "shape")
	}
}

func (ptr *QGraphicsVideoItem) Shape() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsVideoItem::shape")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsVideoItem_Shape(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) ShapeDefault() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsVideoItem::shape")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsVideoItem_ShapeDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsVideoItem_Type
func callbackQGraphicsVideoItem_Type(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(NewQGraphicsVideoItemFromPointer(ptr).TypeDefault())
}

func (ptr *QGraphicsVideoItem) ConnectType(f func() int) {
	defer qt.Recovering("connect QGraphicsVideoItem::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "type", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectType() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "type")
	}
}

func (ptr *QGraphicsVideoItem) Type() int {
	defer qt.Recovering("QGraphicsVideoItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsVideoItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsVideoItem) TypeDefault() int {
	defer qt.Recovering("QGraphicsVideoItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsVideoItem_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsVideoItem_WheelEvent
func callbackQGraphicsVideoItem_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneWheelEvent))(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectWheelEvent(f func(event *widgets.QGraphicsSceneWheelEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *QGraphicsVideoItem) WheelEvent(event widgets.QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_WheelEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

//export callbackQGraphicsVideoItem_TimerEvent
func callbackQGraphicsVideoItem_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QGraphicsVideoItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGraphicsVideoItem_ChildEvent
func callbackQGraphicsVideoItem_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QGraphicsVideoItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsVideoItem_ConnectNotify
func callbackQGraphicsVideoItem_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsVideoItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGraphicsVideoItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QGraphicsVideoItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsVideoItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsVideoItem_CustomEvent
func callbackQGraphicsVideoItem_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsVideoItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QGraphicsVideoItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsVideoItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsVideoItem_DeleteLater
func callbackQGraphicsVideoItem_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsVideoItem::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsVideoItem) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QGraphicsVideoItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QGraphicsVideoItem) DeleteLater() {
	defer qt.Recovering("QGraphicsVideoItem::deleteLater")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsVideoItem) DeleteLaterDefault() {
	defer qt.Recovering("QGraphicsVideoItem::deleteLater")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsVideoItem_DisconnectNotify
func callbackQGraphicsVideoItem_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsVideoItem::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsVideoItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsVideoItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGraphicsVideoItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QGraphicsVideoItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsVideoItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsVideoItem_EventFilter
func callbackQGraphicsVideoItem_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsVideoItem::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsVideoItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsVideoItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsVideoItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QGraphicsVideoItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsVideoItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsVideoItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsVideoItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsVideoItem_MetaObject
func callbackQGraphicsVideoItem_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsVideoItem::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGraphicsVideoItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGraphicsVideoItem) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QGraphicsVideoItem::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QGraphicsVideoItem) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QGraphicsVideoItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsVideoItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QGraphicsVideoItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsVideoItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QVideoWidget struct {
	multimedia.QMediaBindableInterface
	widgets.QWidget
}

type QVideoWidget_ITF interface {
	multimedia.QMediaBindableInterface_ITF
	widgets.QWidget_ITF
	QVideoWidget_PTR() *QVideoWidget
}

func (p *QVideoWidget) QVideoWidget_PTR() *QVideoWidget {
	return p
}

func (p *QVideoWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QMediaBindableInterface_PTR().Pointer()
	}
	return nil
}

func (p *QVideoWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QMediaBindableInterface_PTR().SetPointer(ptr)
		p.QWidget_PTR().SetPointer(ptr)
	}
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
	return n
}

func newQVideoWidgetFromPointer(ptr unsafe.Pointer) *QVideoWidget {
	var n = NewQVideoWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidget_") {
		n.SetObjectName("QVideoWidget_" + qt.Identifier())
	}
	return n
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

//export callbackQVideoWidget_MediaObject
func callbackQVideoWidget_MediaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidget::mediaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaObject"); signal != nil {
		return multimedia.PointerFromQMediaObject(signal.(func() *multimedia.QMediaObject)())
	}

	return multimedia.PointerFromQMediaObject(NewQVideoWidgetFromPointer(ptr).MediaObjectDefault())
}

func (ptr *QVideoWidget) ConnectMediaObject(f func() *multimedia.QMediaObject) {
	defer qt.Recovering("connect QVideoWidget::mediaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mediaObject", f)
	}
}

func (ptr *QVideoWidget) DisconnectMediaObject() {
	defer qt.Recovering("disconnect QVideoWidget::mediaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mediaObject")
	}
}

func (ptr *QVideoWidget) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QVideoWidget::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QVideoWidget_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) MediaObjectDefault() *multimedia.QMediaObject {
	defer qt.Recovering("QVideoWidget::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QVideoWidget_MediaObjectDefault(ptr.Pointer()))
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

//export callbackQVideoWidget_SetAspectRatioMode
func callbackQVideoWidget_SetAspectRatioMode(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QVideoWidget::setAspectRatioMode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setAspectRatioMode"); signal != nil {
		signal.(func(core.Qt__AspectRatioMode))(core.Qt__AspectRatioMode(mode))
	}

}

func (ptr *QVideoWidget) ConnectSetAspectRatioMode(f func(mode core.Qt__AspectRatioMode)) {
	defer qt.Recovering("connect QVideoWidget::setAspectRatioMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setAspectRatioMode", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("disconnect QVideoWidget::setAspectRatioMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setAspectRatioMode")
	}
}

func (ptr *QVideoWidget) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWidget::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

//export callbackQVideoWidget_SetBrightness
func callbackQVideoWidget_SetBrightness(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidget::setBrightness")

	if signal := qt.GetSignal(C.GoString(ptrName), "setBrightness"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidget) ConnectSetBrightness(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidget::setBrightness")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setBrightness", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetBrightness(brightness int) {
	defer qt.Recovering("disconnect QVideoWidget::setBrightness")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setBrightness")
	}
}

func (ptr *QVideoWidget) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWidget::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

//export callbackQVideoWidget_SetContrast
func callbackQVideoWidget_SetContrast(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidget::setContrast")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContrast"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidget) ConnectSetContrast(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidget::setContrast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setContrast", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetContrast(contrast int) {
	defer qt.Recovering("disconnect QVideoWidget::setContrast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setContrast")
	}
}

func (ptr *QVideoWidget) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWidget::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

//export callbackQVideoWidget_SetFullScreen
func callbackQVideoWidget_SetFullScreen(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidget::setFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFullScreen"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidget) ConnectSetFullScreen(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidget::setFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFullScreen", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetFullScreen(fullScreen bool) {
	defer qt.Recovering("disconnect QVideoWidget::setFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFullScreen")
	}
}

func (ptr *QVideoWidget) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWidget::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

//export callbackQVideoWidget_SetHue
func callbackQVideoWidget_SetHue(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidget::setHue")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHue"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidget) ConnectSetHue(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidget::setHue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHue", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetHue(hue int) {
	defer qt.Recovering("disconnect QVideoWidget::setHue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHue")
	}
}

func (ptr *QVideoWidget) SetHue(hue int) {
	defer qt.Recovering("QVideoWidget::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetHue(ptr.Pointer(), C.int(hue))
	}
}

//export callbackQVideoWidget_SetSaturation
func callbackQVideoWidget_SetSaturation(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidget::setSaturation")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSaturation"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidget) ConnectSetSaturation(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidget::setSaturation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSaturation", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetSaturation(saturation int) {
	defer qt.Recovering("disconnect QVideoWidget::setSaturation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSaturation")
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

	return newQVideoWidgetFromPointer(C.QVideoWidget_NewQVideoWidget(widgets.PointerFromQWidget(parent)))
}

//export callbackQVideoWidget_BrightnessChanged
func callbackQVideoWidget_BrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidget::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

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

func (ptr *QVideoWidget) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWidget::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

//export callbackQVideoWidget_ContrastChanged
func callbackQVideoWidget_ContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidget::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
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

func (ptr *QVideoWidget) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWidget::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

//export callbackQVideoWidget_Event
func callbackQVideoWidget_Event(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QVideoWidget::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event))))
}

func (ptr *QVideoWidget) ConnectEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QVideoWidget::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QVideoWidget) DisconnectEvent() {
	defer qt.Recovering("disconnect QVideoWidget::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QVideoWidget) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidget::event")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QVideoWidget) EventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidget::event")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQVideoWidget_FullScreenChanged
func callbackQVideoWidget_FullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidget::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
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

func (ptr *QVideoWidget) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWidget::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

//export callbackQVideoWidget_HideEvent
func callbackQVideoWidget_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
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

//export callbackQVideoWidget_HueChanged
func callbackQVideoWidget_HueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidget::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
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

//export callbackQVideoWidget_MoveEvent
func callbackQVideoWidget_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
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

//export callbackQVideoWidget_PaintEvent
func callbackQVideoWidget_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
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

//export callbackQVideoWidget_ResizeEvent
func callbackQVideoWidget_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
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

//export callbackQVideoWidget_SaturationChanged
func callbackQVideoWidget_SaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidget::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
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

func (ptr *QVideoWidget) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWidget::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

//export callbackQVideoWidget_ShowEvent
func callbackQVideoWidget_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
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

//export callbackQVideoWidget_SizeHint
func callbackQVideoWidget_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidget::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQVideoWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QVideoWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QVideoWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QVideoWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QVideoWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QVideoWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QVideoWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QVideoWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWidget_SizeHintDefault(ptr.Pointer()))
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

//export callbackQVideoWidget_SetMediaObject
func callbackQVideoWidget_SetMediaObject(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) C.int {
	defer qt.Recovering("callback QVideoWidget::setMediaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMediaObject"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*multimedia.QMediaObject) bool)(multimedia.NewQMediaObjectFromPointer(object))))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).SetMediaObjectDefault(multimedia.NewQMediaObjectFromPointer(object))))
}

func (ptr *QVideoWidget) ConnectSetMediaObject(f func(object *multimedia.QMediaObject) bool) {
	defer qt.Recovering("connect QVideoWidget::setMediaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMediaObject", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetMediaObject() {
	defer qt.Recovering("disconnect QVideoWidget::setMediaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMediaObject")
	}
}

func (ptr *QVideoWidget) SetMediaObject(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QVideoWidget::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_SetMediaObject(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

func (ptr *QVideoWidget) SetMediaObjectDefault(object multimedia.QMediaObject_ITF) bool {
	defer qt.Recovering("QVideoWidget::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_SetMediaObjectDefault(ptr.Pointer(), multimedia.PointerFromQMediaObject(object)) != 0
	}
	return false
}

//export callbackQVideoWidget_ActionEvent
func callbackQVideoWidget_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
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

//export callbackQVideoWidget_DragEnterEvent
func callbackQVideoWidget_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
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

//export callbackQVideoWidget_DragLeaveEvent
func callbackQVideoWidget_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
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

//export callbackQVideoWidget_DragMoveEvent
func callbackQVideoWidget_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
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

//export callbackQVideoWidget_DropEvent
func callbackQVideoWidget_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
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

//export callbackQVideoWidget_EnterEvent
func callbackQVideoWidget_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQVideoWidget_FocusInEvent
func callbackQVideoWidget_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQVideoWidget_FocusOutEvent
func callbackQVideoWidget_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQVideoWidget_LeaveEvent
func callbackQVideoWidget_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQVideoWidget_Metric
func callbackQVideoWidget_Metric(ptr unsafe.Pointer, ptrName *C.char, m C.int) C.int {
	defer qt.Recovering("callback QVideoWidget::metric")

	if signal := qt.GetSignal(C.GoString(ptrName), "metric"); signal != nil {
		return C.int(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m)))
	}

	return C.int(NewQVideoWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m)))
}

func (ptr *QVideoWidget) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	defer qt.Recovering("connect QVideoWidget::metric")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metric", f)
	}
}

func (ptr *QVideoWidget) DisconnectMetric() {
	defer qt.Recovering("disconnect QVideoWidget::metric")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metric")
	}
}

func (ptr *QVideoWidget) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QVideoWidget::metric")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_Metric(ptr.Pointer(), C.int(m)))
	}
	return 0
}

func (ptr *QVideoWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QVideoWidget::metric")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_MetricDefault(ptr.Pointer(), C.int(m)))
	}
	return 0
}

//export callbackQVideoWidget_MinimumSizeHint
func callbackQVideoWidget_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidget::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQVideoWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QVideoWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QVideoWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *QVideoWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QVideoWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *QVideoWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QVideoWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWidget_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QVideoWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QVideoWidget_MinimumSizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVideoWidget_PaintEngine
func callbackQVideoWidget_PaintEngine(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidget::paintEngine")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQVideoWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QVideoWidget) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	defer qt.Recovering("connect QVideoWidget::paintEngine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEngine", f)
	}
}

func (ptr *QVideoWidget) DisconnectPaintEngine() {
	defer qt.Recovering("disconnect QVideoWidget::paintEngine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEngine")
	}
}

func (ptr *QVideoWidget) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QVideoWidget::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QVideoWidget_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) PaintEngineDefault() *gui.QPaintEngine {
	defer qt.Recovering("QVideoWidget::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QVideoWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVideoWidget_SetEnabled
func callbackQVideoWidget_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QVideoWidget::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQVideoWidgetFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *QVideoWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QVideoWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QVideoWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *QVideoWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QVideoWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QVideoWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QVideoWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQVideoWidget_SetStyleSheet
func callbackQVideoWidget_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback QVideoWidget::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQVideoWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QVideoWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QVideoWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QVideoWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *QVideoWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QVideoWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QVideoWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QVideoWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetStyleSheetDefault(ptr.Pointer(), C.CString(styleSheet))
	}
}

//export callbackQVideoWidget_SetVisible
func callbackQVideoWidget_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QVideoWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQVideoWidgetFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
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

//export callbackQVideoWidget_SetWindowModified
func callbackQVideoWidget_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QVideoWidget::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQVideoWidgetFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *QVideoWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QVideoWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QVideoWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *QVideoWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QVideoWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QVideoWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QVideoWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQVideoWidget_SetWindowTitle
func callbackQVideoWidget_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QVideoWidget::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQVideoWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QVideoWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QVideoWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QVideoWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *QVideoWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QVideoWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetWindowTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QVideoWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QVideoWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetWindowTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQVideoWidget_ChangeEvent
func callbackQVideoWidget_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQVideoWidget_Close
func callbackQVideoWidget_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidget::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).CloseDefault()))
}

func (ptr *QVideoWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QVideoWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QVideoWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QVideoWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QVideoWidget) Close() bool {
	defer qt.Recovering("QVideoWidget::close")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidget) CloseDefault() bool {
	defer qt.Recovering("QVideoWidget::close")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQVideoWidget_CloseEvent
func callbackQVideoWidget_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
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

//export callbackQVideoWidget_ContextMenuEvent
func callbackQVideoWidget_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
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

//export callbackQVideoWidget_FocusNextPrevChild
func callbackQVideoWidget_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QVideoWidget::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QVideoWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QVideoWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QVideoWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QVideoWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QVideoWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QVideoWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QVideoWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QVideoWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQVideoWidget_HasHeightForWidth
func callbackQVideoWidget_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidget::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *QVideoWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QVideoWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *QVideoWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QVideoWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *QVideoWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QVideoWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QVideoWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQVideoWidget_HeightForWidth
func callbackQVideoWidget_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback QVideoWidget::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewQVideoWidgetFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *QVideoWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QVideoWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *QVideoWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QVideoWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *QVideoWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QVideoWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QVideoWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QVideoWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidget_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackQVideoWidget_Hide
func callbackQVideoWidget_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QVideoWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QVideoWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QVideoWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QVideoWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QVideoWidget) Hide() {
	defer qt.Recovering("QVideoWidget::hide")

	if ptr.Pointer() != nil {
		C.QVideoWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) HideDefault() {
	defer qt.Recovering("QVideoWidget::hide")

	if ptr.Pointer() != nil {
		C.QVideoWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_InitPainter
func callbackQVideoWidget_InitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQVideoWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
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

//export callbackQVideoWidget_InputMethodEvent
func callbackQVideoWidget_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
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

//export callbackQVideoWidget_InputMethodQuery
func callbackQVideoWidget_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidget::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQVideoWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QVideoWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QVideoWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QVideoWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QVideoWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QVideoWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QVideoWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoWidget_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QVideoWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QVideoWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoWidget_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQVideoWidget_KeyPressEvent
func callbackQVideoWidget_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQVideoWidget_KeyReleaseEvent
func callbackQVideoWidget_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQVideoWidget_Lower
func callbackQVideoWidget_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QVideoWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QVideoWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QVideoWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QVideoWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QVideoWidget) Lower() {
	defer qt.Recovering("QVideoWidget::lower")

	if ptr.Pointer() != nil {
		C.QVideoWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) LowerDefault() {
	defer qt.Recovering("QVideoWidget::lower")

	if ptr.Pointer() != nil {
		C.QVideoWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_MouseDoubleClickEvent
func callbackQVideoWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQVideoWidget_MouseMoveEvent
func callbackQVideoWidget_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQVideoWidget_MousePressEvent
func callbackQVideoWidget_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQVideoWidget_MouseReleaseEvent
func callbackQVideoWidget_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQVideoWidget_NativeEvent
func callbackQVideoWidget_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QVideoWidget::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QVideoWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QVideoWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QVideoWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QVideoWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QVideoWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QVideoWidget::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QVideoWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QVideoWidget::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQVideoWidget_Raise
func callbackQVideoWidget_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QVideoWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QVideoWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QVideoWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QVideoWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QVideoWidget) Raise() {
	defer qt.Recovering("QVideoWidget::raise")

	if ptr.Pointer() != nil {
		C.QVideoWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) RaiseDefault() {
	defer qt.Recovering("QVideoWidget::raise")

	if ptr.Pointer() != nil {
		C.QVideoWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_Repaint
func callbackQVideoWidget_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QVideoWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QVideoWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *QVideoWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QVideoWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *QVideoWidget) Repaint() {
	defer qt.Recovering("QVideoWidget::repaint")

	if ptr.Pointer() != nil {
		C.QVideoWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) RepaintDefault() {
	defer qt.Recovering("QVideoWidget::repaint")

	if ptr.Pointer() != nil {
		C.QVideoWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_SetDisabled
func callbackQVideoWidget_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback QVideoWidget::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewQVideoWidgetFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *QVideoWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QVideoWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QVideoWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *QVideoWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QVideoWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QVideoWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QVideoWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackQVideoWidget_SetFocus2
func callbackQVideoWidget_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QVideoWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QVideoWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QVideoWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *QVideoWidget) SetFocus2() {
	defer qt.Recovering("QVideoWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) SetFocus2Default() {
	defer qt.Recovering("QVideoWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQVideoWidget_SetHidden
func callbackQVideoWidget_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback QVideoWidget::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewQVideoWidgetFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *QVideoWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QVideoWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *QVideoWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QVideoWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *QVideoWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QVideoWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QVideoWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QVideoWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QVideoWidget_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackQVideoWidget_Show
func callbackQVideoWidget_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QVideoWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QVideoWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QVideoWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QVideoWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QVideoWidget) Show() {
	defer qt.Recovering("QVideoWidget::show")

	if ptr.Pointer() != nil {
		C.QVideoWidget_Show(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) ShowDefault() {
	defer qt.Recovering("QVideoWidget::show")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_ShowFullScreen
func callbackQVideoWidget_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QVideoWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QVideoWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QVideoWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QVideoWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QVideoWidget) ShowFullScreen() {
	defer qt.Recovering("QVideoWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QVideoWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_ShowMaximized
func callbackQVideoWidget_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QVideoWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QVideoWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QVideoWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QVideoWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QVideoWidget) ShowMaximized() {
	defer qt.Recovering("QVideoWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QVideoWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_ShowMinimized
func callbackQVideoWidget_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QVideoWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QVideoWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QVideoWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QVideoWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QVideoWidget) ShowMinimized() {
	defer qt.Recovering("QVideoWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QVideoWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_ShowNormal
func callbackQVideoWidget_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QVideoWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QVideoWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QVideoWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QVideoWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QVideoWidget) ShowNormal() {
	defer qt.Recovering("QVideoWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) ShowNormalDefault() {
	defer qt.Recovering("QVideoWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_TabletEvent
func callbackQVideoWidget_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
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

//export callbackQVideoWidget_Update
func callbackQVideoWidget_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QVideoWidget) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QVideoWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QVideoWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QVideoWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QVideoWidget) Update() {
	defer qt.Recovering("QVideoWidget::update")

	if ptr.Pointer() != nil {
		C.QVideoWidget_Update(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) UpdateDefault() {
	defer qt.Recovering("QVideoWidget::update")

	if ptr.Pointer() != nil {
		C.QVideoWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_UpdateMicroFocus
func callbackQVideoWidget_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QVideoWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QVideoWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QVideoWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QVideoWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QVideoWidget) UpdateMicroFocus() {
	defer qt.Recovering("QVideoWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QVideoWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QVideoWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QVideoWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QVideoWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQVideoWidget_WheelEvent
func callbackQVideoWidget_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
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

//export callbackQVideoWidget_TimerEvent
func callbackQVideoWidget_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQVideoWidget_ChildEvent
func callbackQVideoWidget_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQVideoWidget_ConnectNotify
func callbackQVideoWidget_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVideoWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVideoWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QVideoWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QVideoWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QVideoWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QVideoWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QVideoWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVideoWidget_CustomEvent
func callbackQVideoWidget_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQVideoWidget_DeleteLater
func callbackQVideoWidget_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidget::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVideoWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QVideoWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QVideoWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QVideoWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QVideoWidget) DeleteLater() {
	defer qt.Recovering("QVideoWidget::deleteLater")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWidget) DeleteLaterDefault() {
	defer qt.Recovering("QVideoWidget::deleteLater")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQVideoWidget_DisconnectNotify
func callbackQVideoWidget_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidget::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVideoWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVideoWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QVideoWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QVideoWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QVideoWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QVideoWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QVideoWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVideoWidget_EventFilter
func callbackQVideoWidget_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QVideoWidget::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QVideoWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QVideoWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QVideoWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QVideoWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QVideoWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QVideoWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QVideoWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQVideoWidget_MetaObject
func callbackQVideoWidget_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidget::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQVideoWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVideoWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QVideoWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QVideoWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QVideoWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QVideoWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QVideoWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVideoWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QVideoWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVideoWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QVideoWidgetControl struct {
	multimedia.QMediaControl
}

type QVideoWidgetControl_ITF interface {
	multimedia.QMediaControl_ITF
	QVideoWidgetControl_PTR() *QVideoWidgetControl
}

func (p *QVideoWidgetControl) QVideoWidgetControl_PTR() *QVideoWidgetControl {
	return p
}

func (p *QVideoWidgetControl) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QMediaControl_PTR().Pointer()
	}
	return nil
}

func (p *QVideoWidgetControl) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QMediaControl_PTR().SetPointer(ptr)
	}
}

func PointerFromQVideoWidgetControl(ptr QVideoWidgetControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidgetControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoWidgetControlFromPointer(ptr unsafe.Pointer) *QVideoWidgetControl {
	var n = new(QVideoWidgetControl)
	n.SetPointer(ptr)
	return n
}

func newQVideoWidgetControlFromPointer(ptr unsafe.Pointer) *QVideoWidgetControl {
	var n = NewQVideoWidgetControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidgetControl_") {
		n.SetObjectName("QVideoWidgetControl_" + qt.Identifier())
	}
	return n
}

func NewQVideoWidgetControl(parent core.QObject_ITF) *QVideoWidgetControl {
	defer qt.Recovering("QVideoWidgetControl::QVideoWidgetControl")

	return newQVideoWidgetControlFromPointer(C.QVideoWidgetControl_NewQVideoWidgetControl(core.PointerFromQObject(parent)))
}

//export callbackQVideoWidgetControl_AspectRatioMode
func callbackQVideoWidgetControl_AspectRatioMode(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::aspectRatioMode")

	if signal := qt.GetSignal(C.GoString(ptrName), "aspectRatioMode"); signal != nil {
		return C.int(signal.(func() core.Qt__AspectRatioMode)())
	}

	return C.int(0)
}

func (ptr *QVideoWidgetControl) ConnectAspectRatioMode(f func() core.Qt__AspectRatioMode) {
	defer qt.Recovering("connect QVideoWidgetControl::aspectRatioMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "aspectRatioMode", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectAspectRatioMode() {
	defer qt.Recovering("disconnect QVideoWidgetControl::aspectRatioMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "aspectRatioMode")
	}
}

func (ptr *QVideoWidgetControl) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QVideoWidgetControl::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidgetControl_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQVideoWidgetControl_Brightness
func callbackQVideoWidgetControl_Brightness(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::brightness")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightness"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QVideoWidgetControl) ConnectBrightness(f func() int) {
	defer qt.Recovering("connect QVideoWidgetControl::brightness")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "brightness", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectBrightness() {
	defer qt.Recovering("disconnect QVideoWidgetControl::brightness")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "brightness")
	}
}

func (ptr *QVideoWidgetControl) Brightness() int {
	defer qt.Recovering("QVideoWidgetControl::brightness")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Brightness(ptr.Pointer()))
	}
	return 0
}

//export callbackQVideoWidgetControl_BrightnessChanged
func callbackQVideoWidgetControl_BrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidgetControl) ConnectBrightnessChanged(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectBrightnessChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

func (ptr *QVideoWidgetControl) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

//export callbackQVideoWidgetControl_Contrast
func callbackQVideoWidgetControl_Contrast(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::contrast")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrast"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QVideoWidgetControl) ConnectContrast(f func() int) {
	defer qt.Recovering("connect QVideoWidgetControl::contrast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contrast", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectContrast() {
	defer qt.Recovering("disconnect QVideoWidgetControl::contrast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contrast")
	}
}

func (ptr *QVideoWidgetControl) Contrast() int {
	defer qt.Recovering("QVideoWidgetControl::contrast")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Contrast(ptr.Pointer()))
	}
	return 0
}

//export callbackQVideoWidgetControl_ContrastChanged
func callbackQVideoWidgetControl_ContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidgetControl) ConnectContrastChanged(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectContrastChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

func (ptr *QVideoWidgetControl) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

//export callbackQVideoWidgetControl_FullScreenChanged
func callbackQVideoWidgetControl_FullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidgetControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectFullScreenChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

func (ptr *QVideoWidgetControl) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

//export callbackQVideoWidgetControl_Hue
func callbackQVideoWidgetControl_Hue(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::hue")

	if signal := qt.GetSignal(C.GoString(ptrName), "hue"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QVideoWidgetControl) ConnectHue(f func() int) {
	defer qt.Recovering("connect QVideoWidgetControl::hue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hue", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectHue() {
	defer qt.Recovering("disconnect QVideoWidgetControl::hue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hue")
	}
}

func (ptr *QVideoWidgetControl) Hue() int {
	defer qt.Recovering("QVideoWidgetControl::hue")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Hue(ptr.Pointer()))
	}
	return 0
}

//export callbackQVideoWidgetControl_HueChanged
func callbackQVideoWidgetControl_HueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidgetControl) ConnectHueChanged(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectHueChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

func (ptr *QVideoWidgetControl) HueChanged(hue int) {
	defer qt.Recovering("QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_HueChanged(ptr.Pointer(), C.int(hue))
	}
}

//export callbackQVideoWidgetControl_IsFullScreen
func callbackQVideoWidgetControl_IsFullScreen(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::isFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "isFullScreen"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QVideoWidgetControl) ConnectIsFullScreen(f func() bool) {
	defer qt.Recovering("connect QVideoWidgetControl::isFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isFullScreen", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectIsFullScreen() {
	defer qt.Recovering("disconnect QVideoWidgetControl::isFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isFullScreen")
	}
}

func (ptr *QVideoWidgetControl) IsFullScreen() bool {
	defer qt.Recovering("QVideoWidgetControl::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQVideoWidgetControl_Saturation
func callbackQVideoWidgetControl_Saturation(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::saturation")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturation"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QVideoWidgetControl) ConnectSaturation(f func() int) {
	defer qt.Recovering("connect QVideoWidgetControl::saturation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "saturation", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSaturation() {
	defer qt.Recovering("disconnect QVideoWidgetControl::saturation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "saturation")
	}
}

func (ptr *QVideoWidgetControl) Saturation() int {
	defer qt.Recovering("QVideoWidgetControl::saturation")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Saturation(ptr.Pointer()))
	}
	return 0
}

//export callbackQVideoWidgetControl_SaturationChanged
func callbackQVideoWidgetControl_SaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidgetControl) ConnectSaturationChanged(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSaturationChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

func (ptr *QVideoWidgetControl) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

//export callbackQVideoWidgetControl_SetAspectRatioMode
func callbackQVideoWidgetControl_SetAspectRatioMode(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::setAspectRatioMode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setAspectRatioMode"); signal != nil {
		signal.(func(core.Qt__AspectRatioMode))(core.Qt__AspectRatioMode(mode))
	}

}

func (ptr *QVideoWidgetControl) ConnectSetAspectRatioMode(f func(mode core.Qt__AspectRatioMode)) {
	defer qt.Recovering("connect QVideoWidgetControl::setAspectRatioMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setAspectRatioMode", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("disconnect QVideoWidgetControl::setAspectRatioMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setAspectRatioMode")
	}
}

func (ptr *QVideoWidgetControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWidgetControl::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

//export callbackQVideoWidgetControl_SetBrightness
func callbackQVideoWidgetControl_SetBrightness(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::setBrightness")

	if signal := qt.GetSignal(C.GoString(ptrName), "setBrightness"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidgetControl) ConnectSetBrightness(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidgetControl::setBrightness")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setBrightness", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSetBrightness(brightness int) {
	defer qt.Recovering("disconnect QVideoWidgetControl::setBrightness")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setBrightness")
	}
}

func (ptr *QVideoWidgetControl) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWidgetControl::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

//export callbackQVideoWidgetControl_SetContrast
func callbackQVideoWidgetControl_SetContrast(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::setContrast")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContrast"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidgetControl) ConnectSetContrast(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidgetControl::setContrast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setContrast", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSetContrast(contrast int) {
	defer qt.Recovering("disconnect QVideoWidgetControl::setContrast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setContrast")
	}
}

func (ptr *QVideoWidgetControl) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWidgetControl::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

//export callbackQVideoWidgetControl_SetFullScreen
func callbackQVideoWidgetControl_SetFullScreen(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::setFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFullScreen"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidgetControl) ConnectSetFullScreen(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidgetControl::setFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFullScreen", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSetFullScreen(fullScreen bool) {
	defer qt.Recovering("disconnect QVideoWidgetControl::setFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFullScreen")
	}
}

func (ptr *QVideoWidgetControl) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWidgetControl::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

//export callbackQVideoWidgetControl_SetHue
func callbackQVideoWidgetControl_SetHue(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::setHue")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHue"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidgetControl) ConnectSetHue(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidgetControl::setHue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHue", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSetHue(hue int) {
	defer qt.Recovering("disconnect QVideoWidgetControl::setHue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHue")
	}
}

func (ptr *QVideoWidgetControl) SetHue(hue int) {
	defer qt.Recovering("QVideoWidgetControl::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetHue(ptr.Pointer(), C.int(hue))
	}
}

//export callbackQVideoWidgetControl_SetSaturation
func callbackQVideoWidgetControl_SetSaturation(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::setSaturation")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSaturation"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidgetControl) ConnectSetSaturation(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidgetControl::setSaturation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSaturation", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSetSaturation(saturation int) {
	defer qt.Recovering("disconnect QVideoWidgetControl::setSaturation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSaturation")
	}
}

func (ptr *QVideoWidgetControl) SetSaturation(saturation int) {
	defer qt.Recovering("QVideoWidgetControl::setSaturation")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

//export callbackQVideoWidgetControl_VideoWidget
func callbackQVideoWidgetControl_VideoWidget(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidgetControl::videoWidget")

	if signal := qt.GetSignal(C.GoString(ptrName), "videoWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func() *widgets.QWidget)())
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QVideoWidgetControl) ConnectVideoWidget(f func() *widgets.QWidget) {
	defer qt.Recovering("connect QVideoWidgetControl::videoWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "videoWidget", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectVideoWidget() {
	defer qt.Recovering("disconnect QVideoWidgetControl::videoWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "videoWidget")
	}
}

func (ptr *QVideoWidgetControl) VideoWidget() *widgets.QWidget {
	defer qt.Recovering("QVideoWidgetControl::videoWidget")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QVideoWidgetControl_VideoWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidgetControl) DestroyQVideoWidgetControl() {
	defer qt.Recovering("QVideoWidgetControl::~QVideoWidgetControl")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DestroyQVideoWidgetControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQVideoWidgetControl_TimerEvent
func callbackQVideoWidgetControl_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QVideoWidgetControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidgetControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQVideoWidgetControl_ChildEvent
func callbackQVideoWidgetControl_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QVideoWidgetControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQVideoWidgetControl_ConnectNotify
func callbackQVideoWidgetControl_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVideoWidgetControl) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QVideoWidgetControl::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QVideoWidgetControl::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QVideoWidgetControl) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidgetControl::connectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QVideoWidgetControl) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidgetControl::connectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVideoWidgetControl_CustomEvent
func callbackQVideoWidgetControl_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QVideoWidgetControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidgetControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQVideoWidgetControl_DeleteLater
func callbackQVideoWidgetControl_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoWidgetControl::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQVideoWidgetControlFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVideoWidgetControl) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QVideoWidgetControl::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QVideoWidgetControl::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QVideoWidgetControl) DeleteLater() {
	defer qt.Recovering("QVideoWidgetControl::deleteLater")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWidgetControl) DeleteLaterDefault() {
	defer qt.Recovering("QVideoWidgetControl::deleteLater")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQVideoWidgetControl_DisconnectNotify
func callbackQVideoWidgetControl_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVideoWidgetControl) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QVideoWidgetControl::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QVideoWidgetControl::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QVideoWidgetControl) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidgetControl::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QVideoWidgetControl) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QVideoWidgetControl::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVideoWidgetControl_Event
func callbackQVideoWidgetControl_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetControlFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QVideoWidgetControl) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QVideoWidgetControl::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QVideoWidgetControl) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidgetControl::event")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QVideoWidgetControl) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidgetControl::event")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQVideoWidgetControl_EventFilter
func callbackQVideoWidgetControl_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QVideoWidgetControl::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQVideoWidgetControlFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QVideoWidgetControl) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QVideoWidgetControl::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QVideoWidgetControl::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QVideoWidgetControl) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidgetControl::eventFilter")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QVideoWidgetControl) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QVideoWidgetControl::eventFilter")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQVideoWidgetControl_MetaObject
func callbackQVideoWidgetControl_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QVideoWidgetControl::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQVideoWidgetControlFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QVideoWidgetControl) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QVideoWidgetControl::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QVideoWidgetControl::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QVideoWidgetControl) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QVideoWidgetControl::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVideoWidgetControl_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidgetControl) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QVideoWidgetControl::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QVideoWidgetControl_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
