package multimedia

//#include "qcameraimagecapture.h"
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

type QCameraImageCaptureITF interface {
	core.QObjectITF
	QMediaBindableInterfaceITF
	QCameraImageCapturePTR() *QCameraImageCapture
}

func (p *QCameraImageCapture) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QCameraImageCapture) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QMediaBindableInterfacePTR().SetPointer(ptr)
}

func PointerFromQCameraImageCapture(ptr QCameraImageCaptureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCapturePTR().Pointer()
	}
	return nil
}

func QCameraImageCaptureFromPointer(ptr unsafe.Pointer) *QCameraImageCapture {
	var n = new(QCameraImageCapture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraImageCapture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageCapture) QCameraImageCapturePTR() *QCameraImageCapture {
	return ptr
}

//QCameraImageCapture::CaptureDestination
type QCameraImageCapture__CaptureDestination int

var (
	QCameraImageCapture__CaptureToFile   = QCameraImageCapture__CaptureDestination(0x01)
	QCameraImageCapture__CaptureToBuffer = QCameraImageCapture__CaptureDestination(0x02)
)

//QCameraImageCapture::DriveMode
type QCameraImageCapture__DriveMode int

var (
	QCameraImageCapture__SingleImageCapture = QCameraImageCapture__DriveMode(0)
)

//QCameraImageCapture::Error
type QCameraImageCapture__Error int

var (
	QCameraImageCapture__NoError                  = QCameraImageCapture__Error(0)
	QCameraImageCapture__NotReadyError            = QCameraImageCapture__Error(1)
	QCameraImageCapture__ResourceError            = QCameraImageCapture__Error(2)
	QCameraImageCapture__OutOfSpaceError          = QCameraImageCapture__Error(3)
	QCameraImageCapture__NotSupportedFeatureError = QCameraImageCapture__Error(4)
	QCameraImageCapture__FormatError              = QCameraImageCapture__Error(5)
)

func (ptr *QCameraImageCapture) IsReadyForCapture() bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsReadyForCapture(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQCameraImageCapture(mediaObject QMediaObjectITF, parent core.QObjectITF) *QCameraImageCapture {
	return QCameraImageCaptureFromPointer(unsafe.Pointer(C.QCameraImageCapture_NewQCameraImageCapture(C.QtObjectPtr(PointerFromQMediaObject(mediaObject)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QCameraImageCapture) BufferFormat() QVideoFrame__PixelFormat {
	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraImageCapture_BufferFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectBufferFormatChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectBufferFormatChanged() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectBufferFormatChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraImageCaptureBufferFormatChanged
func callbackQCameraImageCaptureBufferFormatChanged(ptrName *C.char, format C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged").(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
}

func (ptr *QCameraImageCapture) CancelCapture() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CancelCapture(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCameraImageCapture) Capture(file string) int {
	if ptr.Pointer() != nil {
		return int(C.QCameraImageCapture_Capture(C.QtObjectPtr(ptr.Pointer()), C.CString(file)))
	}
	return 0
}

func (ptr *QCameraImageCapture) CaptureDestination() QCameraImageCapture__CaptureDestination {
	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraImageCapture_CaptureDestination(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectCaptureDestinationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectCaptureDestinationChanged() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectCaptureDestinationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraImageCaptureCaptureDestinationChanged
func callbackQCameraImageCaptureCaptureDestinationChanged(ptrName *C.char, destination C.int) {
	qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged").(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
}

func (ptr *QCameraImageCapture) Error() QCameraImageCapture__Error {
	if ptr.Pointer() != nil {
		return QCameraImageCapture__Error(C.QCameraImageCapture_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraImageCapture) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCameraImageCapture) ImageCodecDescription(codec string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ImageCodecDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(codec)))
	}
	return ""
}

func (ptr *QCameraImageCapture) ConnectImageExposed(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageExposed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageExposed() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageExposed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureImageExposed
func callbackQCameraImageCaptureImageExposed(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "imageExposed").(func(int))(int(id))
}

func (ptr *QCameraImageCapture) ConnectImageMetadataAvailable(f func(id int, key string, value string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageMetadataAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageMetadataAvailable() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageMetadataAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureImageMetadataAvailable
func callbackQCameraImageCaptureImageMetadataAvailable(ptrName *C.char, id C.int, key *C.char, value *C.char) {
	qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable").(func(int, string, string))(int(id), C.GoString(key), C.GoString(value))
}

func (ptr *QCameraImageCapture) ConnectImageSaved(f func(id int, fileName string)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageSaved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageSaved() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageSaved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureImageSaved
func callbackQCameraImageCaptureImageSaved(ptrName *C.char, id C.int, fileName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "imageSaved").(func(int, string))(int(id), C.GoString(fileName))
}

func (ptr *QCameraImageCapture) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsCaptureDestinationSupported(C.QtObjectPtr(ptr.Pointer()), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return QMediaObjectFromPointer(unsafe.Pointer(C.QCameraImageCapture_MediaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCameraImageCapture) ConnectReadyForCaptureChanged(f func(ready bool)) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectReadyForCaptureChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectReadyForCaptureChanged() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectReadyForCaptureChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureReadyForCaptureChanged
func callbackQCameraImageCaptureReadyForCaptureChanged(ptrName *C.char, ready C.int) {
	qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged").(func(bool))(int(ready) != 0)
}

func (ptr *QCameraImageCapture) SetBufferFormat(format QVideoFrame__PixelFormat) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetBufferFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QCameraImageCapture) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetCaptureDestination(C.QtObjectPtr(ptr.Pointer()), C.int(destination))
	}
}

func (ptr *QCameraImageCapture) SetEncodingSettings(settings QImageEncoderSettingsITF) {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetEncodingSettings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImageEncoderSettings(settings)))
	}
}

func (ptr *QCameraImageCapture) SupportedImageCodecs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCameraImageCapture_SupportedImageCodecs(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCameraImageCapture) DestroyQCameraImageCapture() {
	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DestroyQCameraImageCapture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
