package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QCameraImageCapture_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::isReadyForCapture")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsReadyForCapture(ptr.Pointer()) != 0
	}
	return false
}

func NewQCameraImageCapture(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QCameraImageCapture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::QCameraImageCapture")
		}
	}()

	return NewQCameraImageCaptureFromPointer(C.QCameraImageCapture_NewQCameraImageCapture(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QCameraImageCapture) BufferFormat() QVideoFrame__PixelFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::bufferFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraImageCapture_BufferFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::bufferFormatChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectBufferFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectBufferFormatChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::bufferFormatChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectBufferFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraImageCaptureBufferFormatChanged
func callbackQCameraImageCaptureBufferFormatChanged(ptrName *C.char, format C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::bufferFormatChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged").(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
}

func (ptr *QCameraImageCapture) CancelCapture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::cancelCapture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_CancelCapture(ptr.Pointer())
	}
}

func (ptr *QCameraImageCapture) Capture(file string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::capture")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCameraImageCapture_Capture(ptr.Pointer(), C.CString(file)))
	}
	return 0
}

func (ptr *QCameraImageCapture) CaptureDestination() QCameraImageCapture__CaptureDestination {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::captureDestination")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraImageCapture_CaptureDestination(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::captureDestinationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectCaptureDestinationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectCaptureDestinationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::captureDestinationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectCaptureDestinationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraImageCaptureCaptureDestinationChanged
func callbackQCameraImageCaptureCaptureDestinationChanged(ptrName *C.char, destination C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::captureDestinationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged").(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
}

func (ptr *QCameraImageCapture) Error() QCameraImageCapture__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraImageCapture__Error(C.QCameraImageCapture_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCapture) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraImageCapture) ImageCodecDescription(codec string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageCodecDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraImageCapture_ImageCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QCameraImageCapture) ConnectImageExposed(f func(id int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageExposed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageExposed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageExposed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageExposed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageExposed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureImageExposed
func callbackQCameraImageCaptureImageExposed(ptrName *C.char, id C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageExposed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "imageExposed").(func(int))(int(id))
}

func (ptr *QCameraImageCapture) ConnectImageMetadataAvailable(f func(id int, key string, value *core.QVariant)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageMetadataAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageMetadataAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageMetadataAvailable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageMetadataAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageMetadataAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureImageMetadataAvailable
func callbackQCameraImageCaptureImageMetadataAvailable(ptrName *C.char, id C.int, key *C.char, value unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageMetadataAvailable")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable").(func(int, string, *core.QVariant))(int(id), C.GoString(key), core.NewQVariantFromPointer(value))
}

func (ptr *QCameraImageCapture) ConnectImageSaved(f func(id int, fileName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageSaved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectImageSaved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectImageSaved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageSaved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectImageSaved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureImageSaved
func callbackQCameraImageCaptureImageSaved(ptrName *C.char, id C.int, fileName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::imageSaved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "imageSaved").(func(int, string))(int(id), C.GoString(fileName))
}

func (ptr *QCameraImageCapture) IsAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::isAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::isCaptureDestinationSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraImageCapture_IsCaptureDestinationSupported(ptr.Pointer(), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraImageCapture) MediaObject() *QMediaObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::mediaObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QCameraImageCapture_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraImageCapture) ConnectReadyForCaptureChanged(f func(ready bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::readyForCaptureChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_ConnectReadyForCaptureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCapture) DisconnectReadyForCaptureChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::readyForCaptureChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DisconnectReadyForCaptureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureReadyForCaptureChanged
func callbackQCameraImageCaptureReadyForCaptureChanged(ptrName *C.char, ready C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::readyForCaptureChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged").(func(bool))(int(ready) != 0)
}

func (ptr *QCameraImageCapture) SetBufferFormat(format QVideoFrame__PixelFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::setBufferFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetBufferFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraImageCapture) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::setCaptureDestination")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetCaptureDestination(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraImageCapture) SetEncodingSettings(settings QImageEncoderSettings_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::setEncodingSettings")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_SetEncodingSettings(ptr.Pointer(), PointerFromQImageEncoderSettings(settings))
	}
}

func (ptr *QCameraImageCapture) SupportedImageCodecs() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::supportedImageCodecs")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCameraImageCapture_SupportedImageCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCameraImageCapture) DestroyQCameraImageCapture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageCapture::~QCameraImageCapture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageCapture_DestroyQCameraImageCapture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
