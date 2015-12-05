package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QVideoSurfaceFormat struct {
	ptr unsafe.Pointer
}

type QVideoSurfaceFormat_ITF interface {
	QVideoSurfaceFormat_PTR() *QVideoSurfaceFormat
}

func (p *QVideoSurfaceFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVideoSurfaceFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVideoSurfaceFormat(ptr QVideoSurfaceFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoSurfaceFormat_PTR().Pointer()
	}
	return nil
}

func NewQVideoSurfaceFormatFromPointer(ptr unsafe.Pointer) *QVideoSurfaceFormat {
	var n = new(QVideoSurfaceFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVideoSurfaceFormat) QVideoSurfaceFormat_PTR() *QVideoSurfaceFormat {
	return ptr
}

//QVideoSurfaceFormat::Direction
type QVideoSurfaceFormat__Direction int64

const (
	QVideoSurfaceFormat__TopToBottom = QVideoSurfaceFormat__Direction(0)
	QVideoSurfaceFormat__BottomToTop = QVideoSurfaceFormat__Direction(1)
)

//QVideoSurfaceFormat::YCbCrColorSpace
type QVideoSurfaceFormat__YCbCrColorSpace int64

const (
	QVideoSurfaceFormat__YCbCr_Undefined = QVideoSurfaceFormat__YCbCrColorSpace(0)
	QVideoSurfaceFormat__YCbCr_BT601     = QVideoSurfaceFormat__YCbCrColorSpace(1)
	QVideoSurfaceFormat__YCbCr_BT709     = QVideoSurfaceFormat__YCbCrColorSpace(2)
	QVideoSurfaceFormat__YCbCr_xvYCC601  = QVideoSurfaceFormat__YCbCrColorSpace(3)
	QVideoSurfaceFormat__YCbCr_xvYCC709  = QVideoSurfaceFormat__YCbCrColorSpace(4)
	QVideoSurfaceFormat__YCbCr_JPEG      = QVideoSurfaceFormat__YCbCrColorSpace(5)
)

func NewQVideoSurfaceFormat() *QVideoSurfaceFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::QVideoSurfaceFormat")
		}
	}()

	return NewQVideoSurfaceFormatFromPointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat())
}

func NewQVideoSurfaceFormat2(size core.QSize_ITF, format QVideoFrame__PixelFormat, ty QAbstractVideoBuffer__HandleType) *QVideoSurfaceFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::QVideoSurfaceFormat")
		}
	}()

	return NewQVideoSurfaceFormatFromPointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat2(core.PointerFromQSize(size), C.int(format), C.int(ty)))
}

func NewQVideoSurfaceFormat3(other QVideoSurfaceFormat_ITF) *QVideoSurfaceFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::QVideoSurfaceFormat")
		}
	}()

	return NewQVideoSurfaceFormatFromPointer(C.QVideoSurfaceFormat_NewQVideoSurfaceFormat3(PointerFromQVideoSurfaceFormat(other)))
}

func (ptr *QVideoSurfaceFormat) FrameHeight() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::frameHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoSurfaceFormat_FrameHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) FrameRate() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::frameRate")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QVideoSurfaceFormat_FrameRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) FrameWidth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::frameWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoSurfaceFormat_FrameWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) HandleType() QAbstractVideoBuffer__HandleType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::handleType")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QVideoSurfaceFormat_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QVideoSurfaceFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoSurfaceFormat) PixelFormat() QVideoFrame__PixelFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::pixelFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QVideoSurfaceFormat_PixelFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) Property(name string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::property")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QVideoSurfaceFormat_Property(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QVideoSurfaceFormat) ScanLineDirection() QVideoSurfaceFormat__Direction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::scanLineDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return QVideoSurfaceFormat__Direction(C.QVideoSurfaceFormat_ScanLineDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) SetFrameRate(rate float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setFrameRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QVideoSurfaceFormat) SetFrameSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setFrameSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QVideoSurfaceFormat) SetFrameSize2(width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setFrameSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetFrameSize2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QVideoSurfaceFormat) SetPixelAspectRatio(ratio core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setPixelAspectRatio")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetPixelAspectRatio(ptr.Pointer(), core.PointerFromQSize(ratio))
	}
}

func (ptr *QVideoSurfaceFormat) SetPixelAspectRatio2(horizontal int, vertical int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setPixelAspectRatio")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetPixelAspectRatio2(ptr.Pointer(), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QVideoSurfaceFormat) SetProperty(name string, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetProperty(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QVideoSurfaceFormat) SetScanLineDirection(direction QVideoSurfaceFormat__Direction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setScanLineDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetScanLineDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QVideoSurfaceFormat) SetViewport(viewport core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setViewport")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetViewport(ptr.Pointer(), core.PointerFromQRect(viewport))
	}
}

func (ptr *QVideoSurfaceFormat) SetYCbCrColorSpace(space QVideoSurfaceFormat__YCbCrColorSpace) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::setYCbCrColorSpace")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_SetYCbCrColorSpace(ptr.Pointer(), C.int(space))
	}
}

func (ptr *QVideoSurfaceFormat) YCbCrColorSpace() QVideoSurfaceFormat__YCbCrColorSpace {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::yCbCrColorSpace")
		}
	}()

	if ptr.Pointer() != nil {
		return QVideoSurfaceFormat__YCbCrColorSpace(C.QVideoSurfaceFormat_YCbCrColorSpace(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoSurfaceFormat) DestroyQVideoSurfaceFormat() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoSurfaceFormat::~QVideoSurfaceFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(ptr.Pointer())
	}
}
