package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QCameraImageCapture struct {
	core.QObject
	QMediaBindableInterface
}

type QCameraImageCapture_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QCameraImageCapture_PTR() *QCameraImageCapture
}

func (p *QCameraImageCapture) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QCameraImageCapture) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQCameraImageCapture(ptr QCameraImageCapture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCapture_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageCaptureFromPointer(ptr unsafe.Pointer) *QCameraImageCapture {
	var n = new(QCameraImageCapture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageCapture_") {
		n.SetObjectName("QCameraImageCapture_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageCapture) QCameraImageCapture_PTR() *QCameraImageCapture {
	return ptr
}

//QCameraImageCapture::CaptureDestination
type QCameraImageCapture__CaptureDestination int64

const (
	QCameraImageCapture__CaptureToFile   = QCameraImageCapture__CaptureDestination(0x01)
	QCameraImageCapture__CaptureToBuffer = QCameraImageCapture__CaptureDestination(0x02)
)

//QCameraImageCapture::DriveMode
type QCameraImageCapture__DriveMode int64

const (
	QCameraImageCapture__SingleImageCapture = QCameraImageCapture__DriveMode(0)
)

//QCameraImageCapture::Error
type QCameraImageCapture__Error int64

const (
	QCameraImageCapture__NoError                  = QCameraImageCapture__Error(0)
	QCameraImageCapture__NotReadyError            = QCameraImageCapture__Error(1)
	QCameraImageCapture__ResourceError            = QCameraImageCapture__Error(2)
	QCameraImageCapture__OutOfSpaceError          = QCameraImageCapture__Error(3)
	QCameraImageCapture__NotSupportedFeatureError = QCameraImageCapture__Error(4)
	QCameraImageCapture__FormatError              = QCameraImageCapture__Error(5)
)

func (ptr *QCameraImageCapture) IsReadyForCapture() bool {
	defer qt.Recovering("QCameraImageCapture::isReadyForCapture")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsReadyForCapture(ptr.Pointer()) != 0
	}
	return false
}

func NewQCameraImageCapture(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QCameraImageCapture {
	defer qt.Recovering("QCameraImageCapture::QCameraImageCapture")

	return NewQCameraImageCaptureFromPointer(C.QCameraImageCapture_NewQCameraImageCapture(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QCameraImageCapture) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QCameraImageCapture::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QCameraImageCapture_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) BufferFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QCameraImageCapture::bufferFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraImageCapture_BufferFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	defer qt.Recovering("connect QCameraImageCapture::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectBufferFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectBufferFormatChanged() {
	defer qt.Recovering("disconnect QCameraImageCapture::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectBufferFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraImageCaptureBufferFormatChanged
func callbackQCameraImageCaptureBufferFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format C.int) {
	defer qt.Recovering("callback QCameraImageCapture::bufferFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged"); signal != nil {
		signal.(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
	}

}

func (ptr *QCameraImageCapture) BufferFormatChanged(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraImageCapture::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_BufferFormatChanged(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraImageCapture) CancelCapture() {
	defer qt.Recovering("QCameraImageCapture::cancelCapture")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CancelCapture(ptr.Pointer())
	}
}

func (ptr *QCameraImageCapture) Capture(file string) int {
	defer qt.Recovering("QCameraImageCapture::capture")

	if ptr.Pointer() != nil {
		return int(C.QCameraImageCapture_Capture(ptr.Pointer(), C.CString(file)))
	}
	return 0
}

func (ptr *QCameraImageCapture) CaptureDestination() QCameraImageCapture__CaptureDestination {
	defer qt.Recovering("QCameraImageCapture::captureDestination")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraImageCapture_CaptureDestination(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	defer qt.Recovering("connect QCameraImageCapture::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectCaptureDestinationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectCaptureDestinationChanged() {
	defer qt.Recovering("disconnect QCameraImageCapture::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectCaptureDestinationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraImageCaptureCaptureDestinationChanged
func callbackQCameraImageCaptureCaptureDestinationChanged(ptr unsafe.Pointer, ptrName *C.char, destination C.int) {
	defer qt.Recovering("callback QCameraImageCapture::captureDestinationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged"); signal != nil {
		signal.(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
	}

}

func (ptr *QCameraImageCapture) CaptureDestinationChanged(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraImageCapture::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CaptureDestinationChanged(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraImageCapture) ConnectError2(f func(id int, error QCameraImageCapture__Error, errorString string)) {
	defer qt.Recovering("connect QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectError2() {
	defer qt.Recovering("disconnect QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQCameraImageCaptureError2
func callbackQCameraImageCaptureError2(ptr unsafe.Pointer, ptrName *C.char, id C.int, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QCameraImageCapture::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(int, QCameraImageCapture__Error, string))(int(id), QCameraImageCapture__Error(error), C.GoString(errorString))
	}

}

func (ptr *QCameraImageCapture) Error2(id int, error QCameraImageCapture__Error, errorString string) {
	defer qt.Recovering("QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_Error2(ptr.Pointer(), C.int(id), C.int(error), C.CString(errorString))
	}
}

func (ptr *QCameraImageCapture) Error() QCameraImageCapture__Error {
	defer qt.Recovering("QCameraImageCapture::error")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__Error(C.QCameraImageCapture_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ErrorString() string {
	defer qt.Recovering("QCameraImageCapture::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraImageCapture) ImageCodecDescription(codec string) string {
	defer qt.Recovering("QCameraImageCapture::imageCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ImageCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QCameraImageCapture) ConnectImageExposed(f func(id int)) {
	defer qt.Recovering("connect QCameraImageCapture::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageExposed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageExposed() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageExposed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureImageExposed
func callbackQCameraImageCaptureImageExposed(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QCameraImageCapture::imageExposed")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageExposed"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QCameraImageCapture) ImageExposed(id int) {
	defer qt.Recovering("QCameraImageCapture::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageExposed(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QCameraImageCapture) ConnectImageMetadataAvailable(f func(id int, key string, value *core.QVariant)) {
	defer qt.Recovering("connect QCameraImageCapture::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageMetadataAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageMetadataAvailable() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageMetadataAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureImageMetadataAvailable
func callbackQCameraImageCaptureImageMetadataAvailable(ptr unsafe.Pointer, ptrName *C.char, id C.int, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::imageMetadataAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable"); signal != nil {
		signal.(func(int, string, *core.QVariant))(int(id), C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QCameraImageCapture) ImageMetadataAvailable(id int, key string, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraImageCapture::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageMetadataAvailable(ptr.Pointer(), C.int(id), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageCapture) ConnectImageSaved(f func(id int, fileName string)) {
	defer qt.Recovering("connect QCameraImageCapture::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageSaved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageSaved() {
	defer qt.Recovering("disconnect QCameraImageCapture::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageSaved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureImageSaved
func callbackQCameraImageCaptureImageSaved(ptr unsafe.Pointer, ptrName *C.char, id C.int, fileName *C.char) {
	defer qt.Recovering("callback QCameraImageCapture::imageSaved")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageSaved"); signal != nil {
		signal.(func(int, string))(int(id), C.GoString(fileName))
	}

}

func (ptr *QCameraImageCapture) ImageSaved(id int, fileName string) {
	defer qt.Recovering("QCameraImageCapture::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ImageSaved(ptr.Pointer(), C.int(id), C.CString(fileName))
	}
}

func (ptr *QCameraImageCapture) IsAvailable() bool {
	defer qt.Recovering("QCameraImageCapture::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	defer qt.Recovering("QCameraImageCapture::isCaptureDestinationSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsCaptureDestinationSupported(ptr.Pointer(), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) MediaObject() *QMediaObject {
	defer qt.Recovering("QCameraImageCapture::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QCameraImageCapture_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraImageCapture) ConnectReadyForCaptureChanged(f func(ready bool)) {
	defer qt.Recovering("connect QCameraImageCapture::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectReadyForCaptureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectReadyForCaptureChanged() {
	defer qt.Recovering("disconnect QCameraImageCapture::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectReadyForCaptureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureReadyForCaptureChanged
func callbackQCameraImageCaptureReadyForCaptureChanged(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraImageCapture::readyForCaptureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraImageCapture) ReadyForCaptureChanged(ready bool) {
	defer qt.Recovering("QCameraImageCapture::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ReadyForCaptureChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraImageCapture) SetBufferFormat(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraImageCapture::setBufferFormat")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetBufferFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraImageCapture) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraImageCapture::setCaptureDestination")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetCaptureDestination(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraImageCapture) SetEncodingSettings(settings QImageEncoderSettings_ITF) {
	defer qt.Recovering("QCameraImageCapture::setEncodingSettings")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetEncodingSettings(ptr.Pointer(), PointerFromQImageEncoderSettings(settings))
	}
}

func (ptr *QCameraImageCapture) SetMediaObject(mediaObject QMediaObject_ITF) bool {
	defer qt.Recovering("QCameraImageCapture::setMediaObject")

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_SetMediaObject(ptr.Pointer(), PointerFromQMediaObject(mediaObject)) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) SupportedImageCodecs() []string {
	defer qt.Recovering("QCameraImageCapture::supportedImageCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCameraImageCapture_SupportedImageCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCameraImageCapture) DestroyQCameraImageCapture() {
	defer qt.Recovering("QCameraImageCapture::~QCameraImageCapture")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DestroyQCameraImageCapture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraImageCapture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageCaptureTimerEvent
func callbackQCameraImageCaptureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageCaptureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageCapture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCapture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCapture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageCaptureChildEvent
func callbackQCameraImageCaptureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageCaptureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageCapture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCapture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCapture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageCaptureCustomEvent
func callbackQCameraImageCaptureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCapture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageCaptureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageCapture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageCapture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCapture::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
