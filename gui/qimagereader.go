package gui

//#include "gui.h"
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

type QImageReader_ITF interface {
	QImageReader_PTR() *QImageReader
}

func (p *QImageReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageReader(ptr QImageReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageReader_PTR().Pointer()
	}
	return nil
}

func NewQImageReaderFromPointer(ptr unsafe.Pointer) *QImageReader {
	var n = new(QImageReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageReader) QImageReader_PTR() *QImageReader {
	return ptr
}

//QImageReader::ImageReaderError
type QImageReader__ImageReaderError int64

const (
	QImageReader__UnknownError           = QImageReader__ImageReaderError(0)
	QImageReader__FileNotFoundError      = QImageReader__ImageReaderError(1)
	QImageReader__DeviceError            = QImageReader__ImageReaderError(2)
	QImageReader__UnsupportedFormatError = QImageReader__ImageReaderError(3)
	QImageReader__InvalidDataError       = QImageReader__ImageReaderError(4)
)

func NewQImageReader() *QImageReader {
	defer qt.Recovering("QImageReader::QImageReader")

	return NewQImageReaderFromPointer(C.QImageReader_NewQImageReader())
}

func NewQImageReader2(device core.QIODevice_ITF, format core.QByteArray_ITF) *QImageReader {
	defer qt.Recovering("QImageReader::QImageReader")

	return NewQImageReaderFromPointer(C.QImageReader_NewQImageReader2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format)))
}

func NewQImageReader3(fileName string, format core.QByteArray_ITF) *QImageReader {
	defer qt.Recovering("QImageReader::QImageReader")

	return NewQImageReaderFromPointer(C.QImageReader_NewQImageReader3(C.CString(fileName), core.PointerFromQByteArray(format)))
}

func (ptr *QImageReader) AutoDetectImageFormat() bool {
	defer qt.Recovering("QImageReader::autoDetectImageFormat")

	if ptr.Pointer() != nil {
		return C.QImageReader_AutoDetectImageFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) AutoTransform() bool {
	defer qt.Recovering("QImageReader::autoTransform")

	if ptr.Pointer() != nil {
		return C.QImageReader_AutoTransform(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) BackgroundColor() *QColor {
	defer qt.Recovering("QImageReader::backgroundColor")

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QImageReader_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) CanRead() bool {
	defer qt.Recovering("QImageReader::canRead")

	if ptr.Pointer() != nil {
		return C.QImageReader_CanRead(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) ClipRect() *core.QRect {
	defer qt.Recovering("QImageReader::clipRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QImageReader_ClipRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) CurrentImageNumber() int {
	defer qt.Recovering("QImageReader::currentImageNumber")

	if ptr.Pointer() != nil {
		return int(C.QImageReader_CurrentImageNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) CurrentImageRect() *core.QRect {
	defer qt.Recovering("QImageReader::currentImageRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QImageReader_CurrentImageRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) DecideFormatFromContent() bool {
	defer qt.Recovering("QImageReader::decideFormatFromContent")

	if ptr.Pointer() != nil {
		return C.QImageReader_DecideFormatFromContent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) Device() *core.QIODevice {
	defer qt.Recovering("QImageReader::device")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QImageReader_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) Error() QImageReader__ImageReaderError {
	defer qt.Recovering("QImageReader::error")

	if ptr.Pointer() != nil {
		return QImageReader__ImageReaderError(C.QImageReader_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) ErrorString() string {
	defer qt.Recovering("QImageReader::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageReader) FileName() string {
	defer qt.Recovering("QImageReader::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageReader) Format() *core.QByteArray {
	defer qt.Recovering("QImageReader::format")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QImageReader_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) ImageCount() int {
	defer qt.Recovering("QImageReader::imageCount")

	if ptr.Pointer() != nil {
		return int(C.QImageReader_ImageCount(ptr.Pointer()))
	}
	return 0
}

func QImageReader_ImageFormat3(device core.QIODevice_ITF) *core.QByteArray {
	defer qt.Recovering("QImageReader::imageFormat")

	return core.NewQByteArrayFromPointer(C.QImageReader_QImageReader_ImageFormat3(core.PointerFromQIODevice(device)))
}

func QImageReader_ImageFormat2(fileName string) *core.QByteArray {
	defer qt.Recovering("QImageReader::imageFormat")

	return core.NewQByteArrayFromPointer(C.QImageReader_QImageReader_ImageFormat2(C.CString(fileName)))
}

func (ptr *QImageReader) ImageFormat() QImage__Format {
	defer qt.Recovering("QImageReader::imageFormat")

	if ptr.Pointer() != nil {
		return QImage__Format(C.QImageReader_ImageFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) JumpToImage(imageNumber int) bool {
	defer qt.Recovering("QImageReader::jumpToImage")

	if ptr.Pointer() != nil {
		return C.QImageReader_JumpToImage(ptr.Pointer(), C.int(imageNumber)) != 0
	}
	return false
}

func (ptr *QImageReader) JumpToNextImage() bool {
	defer qt.Recovering("QImageReader::jumpToNextImage")

	if ptr.Pointer() != nil {
		return C.QImageReader_JumpToNextImage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) LoopCount() int {
	defer qt.Recovering("QImageReader::loopCount")

	if ptr.Pointer() != nil {
		return int(C.QImageReader_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) NextImageDelay() int {
	defer qt.Recovering("QImageReader::nextImageDelay")

	if ptr.Pointer() != nil {
		return int(C.QImageReader_NextImageDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) Quality() int {
	defer qt.Recovering("QImageReader::quality")

	if ptr.Pointer() != nil {
		return int(C.QImageReader_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) Read2(image QImage_ITF) bool {
	defer qt.Recovering("QImageReader::read")

	if ptr.Pointer() != nil {
		return C.QImageReader_Read2(ptr.Pointer(), PointerFromQImage(image)) != 0
	}
	return false
}

func (ptr *QImageReader) ScaledClipRect() *core.QRect {
	defer qt.Recovering("QImageReader::scaledClipRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QImageReader_ScaledClipRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) ScaledSize() *core.QSize {
	defer qt.Recovering("QImageReader::scaledSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QImageReader_ScaledSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	defer qt.Recovering("QImageReader::setAutoDetectImageFormat")

	if ptr.Pointer() != nil {
		C.QImageReader_SetAutoDetectImageFormat(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QImageReader) SetAutoTransform(enabled bool) {
	defer qt.Recovering("QImageReader::setAutoTransform")

	if ptr.Pointer() != nil {
		C.QImageReader_SetAutoTransform(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QImageReader) SetBackgroundColor(color QColor_ITF) {
	defer qt.Recovering("QImageReader::setBackgroundColor")

	if ptr.Pointer() != nil {
		C.QImageReader_SetBackgroundColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QImageReader) SetClipRect(rect core.QRect_ITF) {
	defer qt.Recovering("QImageReader::setClipRect")

	if ptr.Pointer() != nil {
		C.QImageReader_SetClipRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QImageReader) SetDecideFormatFromContent(ignored bool) {
	defer qt.Recovering("QImageReader::setDecideFormatFromContent")

	if ptr.Pointer() != nil {
		C.QImageReader_SetDecideFormatFromContent(ptr.Pointer(), C.int(qt.GoBoolToInt(ignored)))
	}
}

func (ptr *QImageReader) SetDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QImageReader::setDevice")

	if ptr.Pointer() != nil {
		C.QImageReader_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QImageReader) SetFileName(fileName string) {
	defer qt.Recovering("QImageReader::setFileName")

	if ptr.Pointer() != nil {
		C.QImageReader_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QImageReader) SetFormat(format core.QByteArray_ITF) {
	defer qt.Recovering("QImageReader::setFormat")

	if ptr.Pointer() != nil {
		C.QImageReader_SetFormat(ptr.Pointer(), core.PointerFromQByteArray(format))
	}
}

func (ptr *QImageReader) SetQuality(quality int) {
	defer qt.Recovering("QImageReader::setQuality")

	if ptr.Pointer() != nil {
		C.QImageReader_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QImageReader) SetScaledClipRect(rect core.QRect_ITF) {
	defer qt.Recovering("QImageReader::setScaledClipRect")

	if ptr.Pointer() != nil {
		C.QImageReader_SetScaledClipRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QImageReader) SetScaledSize(size core.QSize_ITF) {
	defer qt.Recovering("QImageReader::setScaledSize")

	if ptr.Pointer() != nil {
		C.QImageReader_SetScaledSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QImageReader) Size() *core.QSize {
	defer qt.Recovering("QImageReader::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QImageReader_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) SubType() *core.QByteArray {
	defer qt.Recovering("QImageReader::subType")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QImageReader_SubType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageReader) SupportsAnimation() bool {
	defer qt.Recovering("QImageReader::supportsAnimation")

	if ptr.Pointer() != nil {
		return C.QImageReader_SupportsAnimation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageReader) SupportsOption(option QImageIOHandler__ImageOption) bool {
	defer qt.Recovering("QImageReader::supportsOption")

	if ptr.Pointer() != nil {
		return C.QImageReader_SupportsOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QImageReader) Text(key string) string {
	defer qt.Recovering("QImageReader::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageReader_Text(ptr.Pointer(), C.CString(key)))
	}
	return ""
}

func (ptr *QImageReader) TextKeys() []string {
	defer qt.Recovering("QImageReader::textKeys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageReader_TextKeys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QImageReader) Transformation() QImageIOHandler__Transformation {
	defer qt.Recovering("QImageReader::transformation")

	if ptr.Pointer() != nil {
		return QImageIOHandler__Transformation(C.QImageReader_Transformation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageReader) DestroyQImageReader() {
	defer qt.Recovering("QImageReader::~QImageReader")

	if ptr.Pointer() != nil {
		C.QImageReader_DestroyQImageReader(ptr.Pointer())
	}
}
