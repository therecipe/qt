package gui

//#include "qimage.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QImage struct {
	QPaintDevice
}

type QImageITF interface {
	QPaintDeviceITF
	QImagePTR() *QImage
}

func PointerFromQImage(ptr QImageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImagePTR().Pointer()
	}
	return nil
}

func QImageFromPointer(ptr unsafe.Pointer) *QImage {
	var n = new(QImage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImage) QImagePTR() *QImage {
	return ptr
}

//QImage::Format
type QImage__Format int

var (
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
type QImage__InvertMode int

var (
	QImage__InvertRgb  = QImage__InvertMode(0)
	QImage__InvertRgba = QImage__InvertMode(1)
)

func (ptr *QImage) ColorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_ColorCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) Fill2(color core.Qt__GlobalColor) {
	if ptr.Pointer() != nil {
		C.QImage_Fill2(C.QtObjectPtr(ptr.Pointer()), C.int(color))
	}
}

func (ptr *QImage) Fill3(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QImage_Fill3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QImage) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QImage_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImage) SetOffset(offset core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QImage_SetOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(offset)))
	}
}

func (ptr *QImage) SetText(key string, text string) {
	if ptr.Pointer() != nil {
		C.QImage_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(text))
	}
}

func (ptr *QImage) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) AllGray() bool {
	if ptr.Pointer() != nil {
		return C.QImage_AllGray(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImage) BitPlaneCount() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_BitPlaneCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) ByteCount() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_ByteCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) BytesPerLine() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_BytesPerLine(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_Depth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) DotsPerMeterX() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_DotsPerMeterX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) DotsPerMeterY() int {
	if ptr.Pointer() != nil {
		return int(C.QImage_DotsPerMeterY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) Format() QImage__Format {
	if ptr.Pointer() != nil {
		return QImage__Format(C.QImage_Format(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImage) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QImage_HasAlphaChannel(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImage) InvertPixels(mode QImage__InvertMode) {
	if ptr.Pointer() != nil {
		C.QImage_InvertPixels(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QImage) IsGrayscale() bool {
	if ptr.Pointer() != nil {
		return C.QImage_IsGrayscale(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImage) Load2(device core.QIODeviceITF, format string) bool {
	if ptr.Pointer() != nil {
		return C.QImage_Load2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)), C.CString(format)) != 0
	}
	return false
}

func (ptr *QImage) Load(fileName string, format string) bool {
	if ptr.Pointer() != nil {
		return C.QImage_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.CString(format)) != 0
	}
	return false
}

func (ptr *QImage) LoadFromData2(data core.QByteArrayITF, format string) bool {
	if ptr.Pointer() != nil {
		return C.QImage_LoadFromData2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.CString(format)) != 0
	}
	return false
}

func (ptr *QImage) PixelIndex(position core.QPointITF) int {
	if ptr.Pointer() != nil {
		return int(C.QImage_PixelIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position))))
	}
	return 0
}

func (ptr *QImage) PixelIndex2(x int, y int) int {
	if ptr.Pointer() != nil {
		return int(C.QImage_PixelIndex2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y)))
	}
	return 0
}

func (ptr *QImage) Save2(device core.QIODeviceITF, format string, quality int) bool {
	if ptr.Pointer() != nil {
		return C.QImage_Save2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QImage) Save(fileName string, format string, quality int) bool {
	if ptr.Pointer() != nil {
		return C.QImage_Save(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QImage) SetColorCount(colorCount int) {
	if ptr.Pointer() != nil {
		C.QImage_SetColorCount(C.QtObjectPtr(ptr.Pointer()), C.int(colorCount))
	}
}

func (ptr *QImage) SetDotsPerMeterX(x int) {
	if ptr.Pointer() != nil {
		C.QImage_SetDotsPerMeterX(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QImage) SetDotsPerMeterY(y int) {
	if ptr.Pointer() != nil {
		C.QImage_SetDotsPerMeterY(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QImage) Swap(other QImageITF) {
	if ptr.Pointer() != nil {
		C.QImage_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(other)))
	}
}

func (ptr *QImage) Text(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImage_Text(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QImage) TextKeys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImage_TextKeys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func QImage_ToImageFormat(format QPixelFormatITF) QImage__Format {
	return QImage__Format(C.QImage_QImage_ToImageFormat(C.QtObjectPtr(PointerFromQPixelFormat(format))))
}

func (ptr *QImage) Valid(pos core.QPointITF) bool {
	if ptr.Pointer() != nil {
		return C.QImage_Valid(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos))) != 0
	}
	return false
}

func (ptr *QImage) Valid2(x int, y int) bool {
	if ptr.Pointer() != nil {
		return C.QImage_Valid2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y)) != 0
	}
	return false
}

func (ptr *QImage) DestroyQImage() {
	if ptr.Pointer() != nil {
		C.QImage_DestroyQImage(C.QtObjectPtr(ptr.Pointer()))
	}
}
