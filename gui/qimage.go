package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QImage struct {
	QPaintDevice
}

type QImage_ITF interface {
	QPaintDevice_ITF
	QImage_PTR() *QImage
}

func PointerFromQImage(ptr QImage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImage_PTR().Pointer()
	}
	return nil
}

func NewQImageFromPointer(ptr unsafe.Pointer) *QImage {
	var n = new(QImage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImage) QImage_PTR() *QImage {
	return ptr
}

//QImage::Format
type QImage__Format int64

const (
	QImage__Format_Invalid                = QImage__Format(0)
	QImage__Format_Mono                   = QImage__Format(1)
	QImage__Format_MonoLSB                = QImage__Format(2)
	QImage__Format_Indexed8               = QImage__Format(3)
	QImage__Format_RGB32                  = QImage__Format(4)
	QImage__Format_ARGB32                 = QImage__Format(5)
	QImage__Format_ARGB32_Premultiplied   = QImage__Format(6)
	QImage__Format_RGB16                  = QImage__Format(7)
	QImage__Format_ARGB8565_Premultiplied = QImage__Format(8)
	QImage__Format_RGB666                 = QImage__Format(9)
	QImage__Format_ARGB6666_Premultiplied = QImage__Format(10)
	QImage__Format_RGB555                 = QImage__Format(11)
	QImage__Format_ARGB8555_Premultiplied = QImage__Format(12)
	QImage__Format_RGB888                 = QImage__Format(13)
	QImage__Format_RGB444                 = QImage__Format(14)
	QImage__Format_ARGB4444_Premultiplied = QImage__Format(15)
	QImage__Format_RGBX8888               = QImage__Format(16)
	QImage__Format_RGBA8888               = QImage__Format(17)
	QImage__Format_RGBA8888_Premultiplied = QImage__Format(18)
	QImage__Format_BGR30                  = QImage__Format(19)
	QImage__Format_A2BGR30_Premultiplied  = QImage__Format(20)
	QImage__Format_RGB30                  = QImage__Format(21)
	QImage__Format_A2RGB30_Premultiplied  = QImage__Format(22)
	QImage__Format_Alpha8                 = QImage__Format(23)
	QImage__Format_Grayscale8             = QImage__Format(24)
)

//QImage::InvertMode
type QImage__InvertMode int64

const (
	QImage__InvertRgb  = QImage__InvertMode(0)
	QImage__InvertRgba = QImage__InvertMode(1)
)

func (ptr *QImage) ColorCount() int {
	defer qt.Recovering("QImage::colorCount")

	if ptr.Pointer() != nil {
		return int(C.QImage_ColorCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) Fill2(color core.Qt__GlobalColor) {
	defer qt.Recovering("QImage::fill")

	if ptr.Pointer() != nil {
		C.QImage_Fill2(ptr.Pointer(), C.int(color))
	}
}

func (ptr *QImage) Fill3(color QColor_ITF) {
	defer qt.Recovering("QImage::fill")

	if ptr.Pointer() != nil {
		C.QImage_Fill3(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QImage) Height() int {
	defer qt.Recovering("QImage::height")

	if ptr.Pointer() != nil {
		return int(C.QImage_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) IsNull() bool {
	defer qt.Recovering("QImage::isNull")

	if ptr.Pointer() != nil {
		return C.QImage_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImage) Offset() *core.QPoint {
	defer qt.Recovering("QImage::offset")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QImage_Offset(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImage) Rect() *core.QRect {
	defer qt.Recovering("QImage::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QImage_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImage) SetOffset(offset core.QPoint_ITF) {
	defer qt.Recovering("QImage::setOffset")

	if ptr.Pointer() != nil {
		C.QImage_SetOffset(ptr.Pointer(), core.PointerFromQPoint(offset))
	}
}

func (ptr *QImage) SetText(key string, text string) {
	defer qt.Recovering("QImage::setText")

	if ptr.Pointer() != nil {
		C.QImage_SetText(ptr.Pointer(), C.CString(key), C.CString(text))
	}
}

func (ptr *QImage) Size() *core.QSize {
	defer qt.Recovering("QImage::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QImage_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImage) Width() int {
	defer qt.Recovering("QImage::width")

	if ptr.Pointer() != nil {
		return int(C.QImage_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) AllGray() bool {
	defer qt.Recovering("QImage::allGray")

	if ptr.Pointer() != nil {
		return C.QImage_AllGray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImage) BitPlaneCount() int {
	defer qt.Recovering("QImage::bitPlaneCount")

	if ptr.Pointer() != nil {
		return int(C.QImage_BitPlaneCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) ByteCount() int {
	defer qt.Recovering("QImage::byteCount")

	if ptr.Pointer() != nil {
		return int(C.QImage_ByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) BytesPerLine() int {
	defer qt.Recovering("QImage::bytesPerLine")

	if ptr.Pointer() != nil {
		return int(C.QImage_BytesPerLine(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) CacheKey() int64 {
	defer qt.Recovering("QImage::cacheKey")

	if ptr.Pointer() != nil {
		return int64(C.QImage_CacheKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) Depth() int {
	defer qt.Recovering("QImage::depth")

	if ptr.Pointer() != nil {
		return int(C.QImage_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) DevicePixelRatio() float64 {
	defer qt.Recovering("QImage::devicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QImage_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) DotsPerMeterX() int {
	defer qt.Recovering("QImage::dotsPerMeterX")

	if ptr.Pointer() != nil {
		return int(C.QImage_DotsPerMeterX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) DotsPerMeterY() int {
	defer qt.Recovering("QImage::dotsPerMeterY")

	if ptr.Pointer() != nil {
		return int(C.QImage_DotsPerMeterY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) Format() QImage__Format {
	defer qt.Recovering("QImage::format")

	if ptr.Pointer() != nil {
		return QImage__Format(C.QImage_Format(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImage) HasAlphaChannel() bool {
	defer qt.Recovering("QImage::hasAlphaChannel")

	if ptr.Pointer() != nil {
		return C.QImage_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImage) InvertPixels(mode QImage__InvertMode) {
	defer qt.Recovering("QImage::invertPixels")

	if ptr.Pointer() != nil {
		C.QImage_InvertPixels(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QImage) IsGrayscale() bool {
	defer qt.Recovering("QImage::isGrayscale")

	if ptr.Pointer() != nil {
		return C.QImage_IsGrayscale(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImage) Load2(device core.QIODevice_ITF, format string) bool {
	defer qt.Recovering("QImage::load")

	if ptr.Pointer() != nil {
		return C.QImage_Load2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format)) != 0
	}
	return false
}

func (ptr *QImage) Load(fileName string, format string) bool {
	defer qt.Recovering("QImage::load")

	if ptr.Pointer() != nil {
		return C.QImage_Load(ptr.Pointer(), C.CString(fileName), C.CString(format)) != 0
	}
	return false
}

func (ptr *QImage) LoadFromData2(data core.QByteArray_ITF, format string) bool {
	defer qt.Recovering("QImage::loadFromData")

	if ptr.Pointer() != nil {
		return C.QImage_LoadFromData2(ptr.Pointer(), core.PointerFromQByteArray(data), C.CString(format)) != 0
	}
	return false
}

func (ptr *QImage) PixelIndex(position core.QPoint_ITF) int {
	defer qt.Recovering("QImage::pixelIndex")

	if ptr.Pointer() != nil {
		return int(C.QImage_PixelIndex(ptr.Pointer(), core.PointerFromQPoint(position)))
	}
	return 0
}

func (ptr *QImage) PixelIndex2(x int, y int) int {
	defer qt.Recovering("QImage::pixelIndex")

	if ptr.Pointer() != nil {
		return int(C.QImage_PixelIndex2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return 0
}

func (ptr *QImage) Save2(device core.QIODevice_ITF, format string, quality int) bool {
	defer qt.Recovering("QImage::save")

	if ptr.Pointer() != nil {
		return C.QImage_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QImage) Save(fileName string, format string, quality int) bool {
	defer qt.Recovering("QImage::save")

	if ptr.Pointer() != nil {
		return C.QImage_Save(ptr.Pointer(), C.CString(fileName), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QImage) SetColorCount(colorCount int) {
	defer qt.Recovering("QImage::setColorCount")

	if ptr.Pointer() != nil {
		C.QImage_SetColorCount(ptr.Pointer(), C.int(colorCount))
	}
}

func (ptr *QImage) SetDevicePixelRatio(scaleFactor float64) {
	defer qt.Recovering("QImage::setDevicePixelRatio")

	if ptr.Pointer() != nil {
		C.QImage_SetDevicePixelRatio(ptr.Pointer(), C.double(scaleFactor))
	}
}

func (ptr *QImage) SetDotsPerMeterX(x int) {
	defer qt.Recovering("QImage::setDotsPerMeterX")

	if ptr.Pointer() != nil {
		C.QImage_SetDotsPerMeterX(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QImage) SetDotsPerMeterY(y int) {
	defer qt.Recovering("QImage::setDotsPerMeterY")

	if ptr.Pointer() != nil {
		C.QImage_SetDotsPerMeterY(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QImage) Swap(other QImage_ITF) {
	defer qt.Recovering("QImage::swap")

	if ptr.Pointer() != nil {
		C.QImage_Swap(ptr.Pointer(), PointerFromQImage(other))
	}
}

func (ptr *QImage) Text(key string) string {
	defer qt.Recovering("QImage::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImage_Text(ptr.Pointer(), C.CString(key)))
	}
	return ""
}

func (ptr *QImage) TextKeys() []string {
	defer qt.Recovering("QImage::textKeys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImage_TextKeys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func QImage_ToImageFormat(format QPixelFormat_ITF) QImage__Format {
	defer qt.Recovering("QImage::toImageFormat")

	return QImage__Format(C.QImage_QImage_ToImageFormat(PointerFromQPixelFormat(format)))
}

func (ptr *QImage) Valid(pos core.QPoint_ITF) bool {
	defer qt.Recovering("QImage::valid")

	if ptr.Pointer() != nil {
		return C.QImage_Valid(ptr.Pointer(), core.PointerFromQPoint(pos)) != 0
	}
	return false
}

func (ptr *QImage) Valid2(x int, y int) bool {
	defer qt.Recovering("QImage::valid")

	if ptr.Pointer() != nil {
		return C.QImage_Valid2(ptr.Pointer(), C.int(x), C.int(y)) != 0
	}
	return false
}

func (ptr *QImage) DestroyQImage() {
	defer qt.Recovering("QImage::~QImage")

	if ptr.Pointer() != nil {
		C.QImage_DestroyQImage(ptr.Pointer())
	}
}
