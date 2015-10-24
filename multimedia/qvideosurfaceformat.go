package multimedia

//#include "qvideosurfaceformat.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVideoSurfaceFormat struct {
	ptr unsafe.Pointer
}

type QVideoSurfaceFormatITF interface {
	QVideoSurfaceFormatPTR() *QVideoSurfaceFormat
}

func (p *QVideoSurfaceFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoSurfaceFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoSurfaceFormat(ptr QVideoSurfaceFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoSurfaceFormatPTR().Pointer()
	}
	return nil
}

func QVideoSurfaceFormatFromPointer(ptr unsafe.Pointer) *QVideoSurfaceFormat {
	var n = new(QVideoSurfaceFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoSurfaceFormat) QVideoSurfaceFormatPTR() *QVideoSurfaceFormat {
	return ptr
}

//QVideoSurfaceFormat::Direction
type QVideoSurfaceFormat__Direction int

var (
	QVideoSurfaceFormat__TopToBottom = QVideoSurfaceFormat__Direction(0)
	QVideoSurfaceFormat__BottomToTop = QVideoSurfaceFormat__Direction(1)
)

//QVideoSurfaceFormat::YCbCrColorSpace
type QVideoSurfaceFormat__YCbCrColorSpace int

var (
	QVideoSurfaceFormat__YCbCr_Undefined = QVideoSurfaceFormat__YCbCrColorSpace(0)
	QVideoSurfaceFormat__YCbCr_BT601     = QVideoSurfaceFormat__YCbCrColorSpace(1)
	QVideoSurfaceFormat__YCbCr_BT709     = QVideoSurfaceFormat__YCbCrColorSpace(2)
	QVideoSurfaceFormat__YCbCr_xvYCC601  = QVideoSurfaceFormat__YCbCrColorSpace(3)
	QVideoSurfaceFormat__YCbCr_xvYCC709  = QVideoSurfaceFormat__YCbCrColorSpace(4)
	QVideoSurfaceFormat__YCbCr_JPEG      = QVideoSurfaceFormat__YCbCrColorSpace(5)
)

func NewQVideoSurfaceFormat() *QVideoSurfaceFormat {
	return QVideoSurfaceFormatFromPointer(unsafe.Pointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat()))
}

func NewQVideoSurfaceFormat2(size core.QSizeITF, format QVideoFrame__PixelFormat, ty QAbstractVideoBuffer__HandleType) *QVideoSurfaceFormat {
	return QVideoSurfaceFormatFromPointer(unsafe.Pointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat2(C.QtObjectPtr(core.PointerFromQSize(size)), C.int(format), C.int(ty))))
}

func NewQVideoSurfaceFormat3(other QVideoSurfaceFormatITF) *QVideoSurfaceFormat {
	return QVideoSurfaceFormatFromPointer(unsafe.Pointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat3(C.QtObjectPtr(PointerFromQVideoSurfaceFormat(other)))))
}

func (ptr *QVideoSurfaceFormat) FrameHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoSurfaceFormat_FrameHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) FrameWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoSurfaceFormat_FrameWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) HandleType() QAbstractVideoBuffer__HandleType {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QVideoSurfaceFormat_HandleType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QVideoSurfaceFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVideoSurfaceFormat) PixelFormat() QVideoFrame__PixelFormat {
	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QVideoSurfaceFormat_PixelFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) Property(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoSurfaceFormat_Property(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QVideoSurfaceFormat) ScanLineDirection() QVideoSurfaceFormat__Direction {
	if ptr.Pointer() != nil {
		return QVideoSurfaceFormat__Direction(C.QVideoSurfaceFormat_ScanLineDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) SetFrameSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QVideoSurfaceFormat) SetFrameSize2(width int, height int) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameSize2(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height))
	}
}

func (ptr *QVideoSurfaceFormat) SetPixelAspectRatio(ratio core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetPixelAspectRatio(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(ratio)))
	}
}

func (ptr *QVideoSurfaceFormat) SetPixelAspectRatio2(horizontal int, vertical int) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetPixelAspectRatio2(C.QtObjectPtr(ptr.Pointer()), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QVideoSurfaceFormat) SetProperty(name string, value string) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value))
	}
}

func (ptr *QVideoSurfaceFormat) SetScanLineDirection(direction QVideoSurfaceFormat__Direction) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetScanLineDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QVideoSurfaceFormat) SetViewport(viewport core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetViewport(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(viewport)))
	}
}

func (ptr *QVideoSurfaceFormat) SetYCbCrColorSpace(space QVideoSurfaceFormat__YCbCrColorSpace) {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetYCbCrColorSpace(C.QtObjectPtr(ptr.Pointer()), C.int(space))
	}
}

func (ptr *QVideoSurfaceFormat) YCbCrColorSpace() QVideoSurfaceFormat__YCbCrColorSpace {
	if ptr.Pointer() != nil {
		return QVideoSurfaceFormat__YCbCrColorSpace(C.QVideoSurfaceFormat_YCbCrColorSpace(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) DestroyQVideoSurfaceFormat() {
	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(C.QtObjectPtr(ptr.Pointer()))
	}
}
