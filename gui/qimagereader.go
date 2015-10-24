package gui

//#include "qimagereader.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QImageReader struct {
	ptr unsafe.Pointer
}

type QImageReaderITF interface {
	QImageReaderPTR() *QImageReader
}

func (p *QImageReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageReader(ptr QImageReaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageReaderPTR().Pointer()
	}
	return nil
}

func QImageReaderFromPointer(ptr unsafe.Pointer) *QImageReader {
	var n = new(QImageReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageReader) QImageReaderPTR() *QImageReader {
	return ptr
}

//QImageReader::ImageReaderError
type QImageReader__ImageReaderError int

var (
	QImageReader__UnknownError           = QImageReader__ImageReaderError(0)
	QImageReader__FileNotFoundError      = QImageReader__ImageReaderError(1)
	QImageReader__DeviceError            = QImageReader__ImageReaderError(2)
	QImageReader__UnsupportedFormatError = QImageReader__ImageReaderError(3)
	QImageReader__InvalidDataError       = QImageReader__ImageReaderError(4)
)

func NewQImageReader() *QImageReader {
	return QImageReaderFromPointer(unsafe.Pointer(C.QImageReader_NewQImageReader()))
}

func NewQImageReader2(device core.QIODeviceITF, format core.QByteArrayITF) *QImageReader {
	return QImageReaderFromPointer(unsafe.Pointer(C.QImageReader_NewQImageReader2(C.QtObjectPtr(core.PointerFromQIODevice(device)), C.QtObjectPtr(core.PointerFromQByteArray(format)))))
}

func NewQImageReader3(fileName string, format core.QByteArrayITF) *QImageReader {
	return QImageReaderFromPointer(unsafe.Pointer(C.QImageReader_NewQImageReader3(C.CString(fileName), C.QtObjectPtr(core.PointerFromQByteArray(format)))))
}

func (ptr *QImageReader) AutoDetectImageFormat() bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_AutoDetectImageFormat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageReader) AutoTransform() bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_AutoTransform(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageReader) CanRead() bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_CanRead(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageReader) CurrentImageNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QImageReader_CurrentImageNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) DecideFormatFromContent() bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_DecideFormatFromContent(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageReader) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QImageReader_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QImageReader) Error() QImageReader__ImageReaderError {
	if ptr.Pointer() != nil {
		return QImageReader__ImageReaderError(C.QImageReader_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QImageReader) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QImageReader) ImageCount() int {
	if ptr.Pointer() != nil {
		return int(C.QImageReader_ImageCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) ImageFormat() QImage__Format {
	if ptr.Pointer() != nil {
		return QImage__Format(C.QImageReader_ImageFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) JumpToImage(imageNumber int) bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_JumpToImage(C.QtObjectPtr(ptr.Pointer()), C.int(imageNumber)) != 0
	}
	return false
}

func (ptr *QImageReader) JumpToNextImage() bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_JumpToNextImage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageReader) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QImageReader_LoopCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) NextImageDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QImageReader_NextImageDelay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) Quality() int {
	if ptr.Pointer() != nil {
		return int(C.QImageReader_Quality(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) Read2(image QImageITF) bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_Read2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image))) != 0
	}
	return false
}

func (ptr *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetAutoDetectImageFormat(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QImageReader) SetAutoTransform(enabled bool) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetAutoTransform(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QImageReader) SetBackgroundColor(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetBackgroundColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QImageReader) SetClipRect(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetClipRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QImageReader) SetDecideFormatFromContent(ignored bool) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetDecideFormatFromContent(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ignored)))
	}
}

func (ptr *QImageReader) SetDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QImageReader) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QImageReader) SetFormat(format core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(format)))
	}
}

func (ptr *QImageReader) SetQuality(quality int) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetQuality(C.QtObjectPtr(ptr.Pointer()), C.int(quality))
	}
}

func (ptr *QImageReader) SetScaledClipRect(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetScaledClipRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QImageReader) SetScaledSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QImageReader_SetScaledSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QImageReader) SupportsAnimation() bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_SupportsAnimation(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageReader) SupportsOption(option QImageIOHandler__ImageOption) bool {
	if ptr.Pointer() != nil {
		return C.QImageReader_SupportsOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QImageReader) Text(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_Text(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QImageReader) TextKeys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageReader_TextKeys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QImageReader) Transformation() QImageIOHandler__Transformation {
	if ptr.Pointer() != nil {
		return QImageIOHandler__Transformation(C.QImageReader_Transformation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageReader) DestroyQImageReader() {
	if ptr.Pointer() != nil {
		C.QImageReader_DestroyQImageReader(C.QtObjectPtr(ptr.Pointer()))
	}
}
