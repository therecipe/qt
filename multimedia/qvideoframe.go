package multimedia

//#include "qvideoframe.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QVideoFrame struct {
	ptr unsafe.Pointer
}

type QVideoFrameITF interface {
	QVideoFramePTR() *QVideoFrame
}

func (p *QVideoFrame) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoFrame) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoFrame(ptr QVideoFrameITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoFramePTR().Pointer()
	}
	return nil
}

func QVideoFrameFromPointer(ptr unsafe.Pointer) *QVideoFrame {
	var n = new(QVideoFrame)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoFrame) QVideoFramePTR() *QVideoFrame {
	return ptr
}

//QVideoFrame::FieldType
type QVideoFrame__FieldType int

var (
	QVideoFrame__ProgressiveFrame = QVideoFrame__FieldType(0)
	QVideoFrame__TopField         = QVideoFrame__FieldType(1)
	QVideoFrame__BottomField      = QVideoFrame__FieldType(2)
	QVideoFrame__InterlacedFrame  = QVideoFrame__FieldType(3)
)

//QVideoFrame::PixelFormat
type QVideoFrame__PixelFormat int

var (
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
	return QVideoFrameFromPointer(unsafe.Pointer(C.QVideoFrame_NewQVideoFrame()))
}

func NewQVideoFrame2(buffer QAbstractVideoBufferITF, size core.QSizeITF, format QVideoFrame__PixelFormat) *QVideoFrame {
	return QVideoFrameFromPointer(unsafe.Pointer(C.QVideoFrame_NewQVideoFrame2(C.QtObjectPtr(PointerFromQAbstractVideoBuffer(buffer)), C.QtObjectPtr(core.PointerFromQSize(size)), C.int(format))))
}

func NewQVideoFrame4(image gui.QImageITF) *QVideoFrame {
	return QVideoFrameFromPointer(unsafe.Pointer(C.QVideoFrame_NewQVideoFrame4(C.QtObjectPtr(gui.PointerFromQImage(image)))))
}

func NewQVideoFrame5(other QVideoFrameITF) *QVideoFrame {
	return QVideoFrameFromPointer(unsafe.Pointer(C.QVideoFrame_NewQVideoFrame5(C.QtObjectPtr(PointerFromQVideoFrame(other)))))
}

func NewQVideoFrame3(bytes int, size core.QSizeITF, bytesPerLine int, format QVideoFrame__PixelFormat) *QVideoFrame {
	return QVideoFrameFromPointer(unsafe.Pointer(C.QVideoFrame_NewQVideoFrame3(C.int(bytes), C.QtObjectPtr(core.PointerFromQSize(size)), C.int(bytesPerLine), C.int(format))))
}

func (ptr *QVideoFrame) BytesPerLine() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_BytesPerLine(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) BytesPerLine2(plane int) int {
	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_BytesPerLine2(C.QtObjectPtr(ptr.Pointer()), C.int(plane)))
	}
	return 0
}

func (ptr *QVideoFrame) FieldType() QVideoFrame__FieldType {
	if ptr.Pointer() != nil {
		return QVideoFrame__FieldType(C.QVideoFrame_FieldType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) Handle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoFrame_Handle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QVideoFrame) HandleType() QAbstractVideoBuffer__HandleType {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QVideoFrame_HandleType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QVideoFrame_ImageFormatFromPixelFormat(format QVideoFrame__PixelFormat) gui.QImage__Format {
	return gui.QImage__Format(C.QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(C.int(format)))
}

func (ptr *QVideoFrame) IsMapped() bool {
	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsMapped(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsReadable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoFrame) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QVideoFrame_IsWritable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoFrame) Map(mode QAbstractVideoBuffer__MapMode) bool {
	if ptr.Pointer() != nil {
		return C.QVideoFrame_Map(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QVideoFrame) MapMode() QAbstractVideoBuffer__MapMode {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QVideoFrame_MapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) MappedBytes() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_MappedBytes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) MetaData(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoFrame_MetaData(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QVideoFrame) PixelFormat() QVideoFrame__PixelFormat {
	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QVideoFrame_PixelFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QVideoFrame_PixelFormatFromImageFormat(format gui.QImage__Format) QVideoFrame__PixelFormat {
	return QVideoFrame__PixelFormat(C.QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(C.int(format)))
}

func (ptr *QVideoFrame) PlaneCount() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_PlaneCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) SetFieldType(field QVideoFrame__FieldType) {
	if ptr.Pointer() != nil {
		C.QVideoFrame_SetFieldType(C.QtObjectPtr(ptr.Pointer()), C.int(field))
	}
}

func (ptr *QVideoFrame) SetMetaData(key string, value string) {
	if ptr.Pointer() != nil {
		C.QVideoFrame_SetMetaData(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(value))
	}
}

func (ptr *QVideoFrame) Unmap() {
	if ptr.Pointer() != nil {
		C.QVideoFrame_Unmap(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QVideoFrame) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoFrame_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoFrame) DestroyQVideoFrame() {
	if ptr.Pointer() != nil {
		C.QVideoFrame_DestroyQVideoFrame(C.QtObjectPtr(ptr.Pointer()))
	}
}
