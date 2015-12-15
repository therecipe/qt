package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QVideoFrame struct {
	ptr unsafe.Pointer
}

type QVideoFrame_ITF interface {
	QVideoFrame_PTR() *QVideoFrame
}

func (p *QVideoFrame) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoFrame) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoFrame(ptr QVideoFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoFrame_PTR().Pointer()
	}
	return nil
}

func NewQVideoFrameFromPointer(ptr unsafe.Pointer) *QVideoFrame {
	var n = new(QVideoFrame)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoFrame) QVideoFrame_PTR() *QVideoFrame {
	return ptr
}

//QVideoFrame::FieldType
type QVideoFrame__FieldType int64

const (
	QVideoFrame__ProgressiveFrame = QVideoFrame__FieldType(0)
	QVideoFrame__TopField         = QVideoFrame__FieldType(1)
	QVideoFrame__BottomField      = QVideoFrame__FieldType(2)
	QVideoFrame__InterlacedFrame  = QVideoFrame__FieldType(3)
)

//QVideoFrame::PixelFormat
type QVideoFrame__PixelFormat int64

const (
	QVideoFrame__Format_Invalid                = QVideoFrame__PixelFormat(0)
	QVideoFrame__Format_ARGB32                 = QVideoFrame__PixelFormat(1)
	QVideoFrame__Format_ARGB32_Premultiplied   = QVideoFrame__PixelFormat(2)
	QVideoFrame__Format_RGB32                  = QVideoFrame__PixelFormat(3)
	QVideoFrame__Format_RGB24                  = QVideoFrame__PixelFormat(4)
	QVideoFrame__Format_RGB565                 = QVideoFrame__PixelFormat(5)
	QVideoFrame__Format_RGB555                 = QVideoFrame__PixelFormat(6)
	QVideoFrame__Format_ARGB8565_Premultiplied = QVideoFrame__PixelFormat(7)
	QVideoFrame__Format_BGRA32                 = QVideoFrame__PixelFormat(8)
	QVideoFrame__Format_BGRA32_Premultiplied   = QVideoFrame__PixelFormat(9)
	QVideoFrame__Format_BGR32                  = QVideoFrame__PixelFormat(10)
	QVideoFrame__Format_BGR24                  = QVideoFrame__PixelFormat(11)
	QVideoFrame__Format_BGR565                 = QVideoFrame__PixelFormat(12)
	QVideoFrame__Format_BGR555                 = QVideoFrame__PixelFormat(13)
	QVideoFrame__Format_BGRA5658_Premultiplied = QVideoFrame__PixelFormat(14)
	QVideoFrame__Format_AYUV444                = QVideoFrame__PixelFormat(15)
	QVideoFrame__Format_AYUV444_Premultiplied  = QVideoFrame__PixelFormat(16)
	QVideoFrame__Format_YUV444                 = QVideoFrame__PixelFormat(17)
	QVideoFrame__Format_YUV420P                = QVideoFrame__PixelFormat(18)
	QVideoFrame__Format_YV12                   = QVideoFrame__PixelFormat(19)
	QVideoFrame__Format_UYVY                   = QVideoFrame__PixelFormat(20)
	QVideoFrame__Format_YUYV                   = QVideoFrame__PixelFormat(21)
	QVideoFrame__Format_NV12                   = QVideoFrame__PixelFormat(22)
	QVideoFrame__Format_NV21                   = QVideoFrame__PixelFormat(23)
	QVideoFrame__Format_IMC1                   = QVideoFrame__PixelFormat(24)
	QVideoFrame__Format_IMC2                   = QVideoFrame__PixelFormat(25)
	QVideoFrame__Format_IMC3                   = QVideoFrame__PixelFormat(26)
	QVideoFrame__Format_IMC4                   = QVideoFrame__PixelFormat(27)
	QVideoFrame__Format_Y8                     = QVideoFrame__PixelFormat(28)
	QVideoFrame__Format_Y16                    = QVideoFrame__PixelFormat(29)
	QVideoFrame__Format_Jpeg                   = QVideoFrame__PixelFormat(30)
	QVideoFrame__Format_CameraRaw              = QVideoFrame__PixelFormat(31)
	QVideoFrame__Format_AdobeDng               = QVideoFrame__PixelFormat(32)
	QVideoFrame__Format_User                   = QVideoFrame__PixelFormat(1000)
)

func NewQVideoFrame() *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return NewQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame())
}

func NewQVideoFrame2(buffer QAbstractVideoBuffer_ITF, size core.QSize_ITF, format QVideoFrame__PixelFormat) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return NewQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame2(PointerFromQAbstractVideoBuffer(buffer), core.PointerFromQSize(size), C.int(format)))
}

func NewQVideoFrame4(image gui.QImage_ITF) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return NewQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame4(gui.PointerFromQImage(image)))
}

func NewQVideoFrame5(other QVideoFrame_ITF) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return NewQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame5(PointerFromQVideoFrame(other)))
}

func NewQVideoFrame3(bytes int, size core.QSize_ITF, bytesPerLine int, format QVideoFrame__PixelFormat) *QVideoFrame {
	defer qt.Recovering("QVideoFrame::QVideoFrame")

	return NewQVideoFrameFromPointer(C.QVideoFrame_NewQVideoFrame3(C.int(bytes), core.PointerFromQSize(size), C.int(bytesPerLine), C.int(format)))
}

func (ptr *QVideoFrame) BytesPerLine() int {
	defer qt.Recovering("QVideoFrame::bytesPerLine")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_BytesPerLine(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) BytesPerLine2(plane int) int {
	defer qt.Recovering("QVideoFrame::bytesPerLine")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_BytesPerLine2(ptr.Pointer(), C.int(plane)))
	}
	return 0
}

func (ptr *QVideoFrame) FieldType() QVideoFrame__FieldType {
	defer qt.Recovering("QVideoFrame::fieldType")

	if ptr.Pointer() != nil {
		return QVideoFrame__FieldType(C.QVideoFrame_FieldType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) Handle() *core.QVariant {
	defer qt.Recovering("QVideoFrame::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoFrame_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoFrame) HandleType() QAbstractVideoBuffer__HandleType {
	defer qt.Recovering("QVideoFrame::handleType")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QVideoFrame_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) Height() int {
	defer qt.Recovering("QVideoFrame::height")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_Height(ptr.Pointer()))
	}
	return 0
}

func QVideoFrame_ImageFormatFromPixelFormat(format QVideoFrame__PixelFormat) gui.QImage__Format {
	defer qt.Recovering("QVideoFrame::imageFormatFromPixelFormat")

	return gui.QImage__Format(C.QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(C.int(format)))
}

func (ptr *QVideoFrame) IsMapped() bool {
	defer qt.Recovering("QVideoFrame::isMapped")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsMapped(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsReadable() bool {
	defer qt.Recovering("QVideoFrame::isReadable")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsValid() bool {
	defer qt.Recovering("QVideoFrame::isValid")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsWritable() bool {
	defer qt.Recovering("QVideoFrame::isWritable")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoFrame) Map(mode QAbstractVideoBuffer__MapMode) bool {
	defer qt.Recovering("QVideoFrame::map")

	if ptr.Pointer() != nil {
		return C.QVideoFrame_Map(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QVideoFrame) MapMode() QAbstractVideoBuffer__MapMode {
	defer qt.Recovering("QVideoFrame::mapMode")

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QVideoFrame_MapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) MappedBytes() int {
	defer qt.Recovering("QVideoFrame::mappedBytes")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_MappedBytes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QVideoFrame::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoFrame_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QVideoFrame) PixelFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QVideoFrame::pixelFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QVideoFrame_PixelFormat(ptr.Pointer()))
	}
	return 0
}

func QVideoFrame_PixelFormatFromImageFormat(format gui.QImage__Format) QVideoFrame__PixelFormat {
	defer qt.Recovering("QVideoFrame::pixelFormatFromImageFormat")

	return QVideoFrame__PixelFormat(C.QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(C.int(format)))
}

func (ptr *QVideoFrame) PlaneCount() int {
	defer qt.Recovering("QVideoFrame::planeCount")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_PlaneCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) SetFieldType(field QVideoFrame__FieldType) {
	defer qt.Recovering("QVideoFrame::setFieldType")

	if ptr.Pointer() != nil {
		C.QVideoFrame_SetFieldType(ptr.Pointer(), C.int(field))
	}
}

func (ptr *QVideoFrame) SetMetaData(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QVideoFrame::setMetaData")

	if ptr.Pointer() != nil {
		C.QVideoFrame_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QVideoFrame) Unmap() {
	defer qt.Recovering("QVideoFrame::unmap")

	if ptr.Pointer() != nil {
		C.QVideoFrame_Unmap(ptr.Pointer())
	}
}

func (ptr *QVideoFrame) Width() int {
	defer qt.Recovering("QVideoFrame::width")

	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoFrame) DestroyQVideoFrame() {
	defer qt.Recovering("QVideoFrame::~QVideoFrame")

	if ptr.Pointer() != nil {
		C.QVideoFrame_DestroyQVideoFrame(ptr.Pointer())
	}
}
