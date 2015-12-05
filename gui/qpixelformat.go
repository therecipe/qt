package gui

//#include "gui.h"
import "C"
import (
	"log"
	"unsafe"
)

type QPixelFormat struct {
	ptr unsafe.Pointer
}

type QPixelFormat_ITF interface {
	QPixelFormat_PTR() *QPixelFormat
}

func (p *QPixelFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPixelFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPixelFormat(ptr QPixelFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixelFormat_PTR().Pointer()
	}
	return nil
}

func NewQPixelFormatFromPointer(ptr unsafe.Pointer) *QPixelFormat {
	var n = new(QPixelFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPixelFormat) QPixelFormat_PTR() *QPixelFormat {
	return ptr
}

//QPixelFormat::AlphaPosition
type QPixelFormat__AlphaPosition int64

const (
	QPixelFormat__AtBeginning = QPixelFormat__AlphaPosition(0)
	QPixelFormat__AtEnd       = QPixelFormat__AlphaPosition(1)
)

//QPixelFormat::AlphaPremultiplied
type QPixelFormat__AlphaPremultiplied int64

const (
	QPixelFormat__NotPremultiplied = QPixelFormat__AlphaPremultiplied(0)
	QPixelFormat__Premultiplied    = QPixelFormat__AlphaPremultiplied(1)
)

//QPixelFormat::AlphaUsage
type QPixelFormat__AlphaUsage int64

const (
	QPixelFormat__UsesAlpha    = QPixelFormat__AlphaUsage(0)
	QPixelFormat__IgnoresAlpha = QPixelFormat__AlphaUsage(1)
)

//QPixelFormat::ByteOrder
type QPixelFormat__ByteOrder int64

const (
	QPixelFormat__LittleEndian        = QPixelFormat__ByteOrder(0)
	QPixelFormat__BigEndian           = QPixelFormat__ByteOrder(1)
	QPixelFormat__CurrentSystemEndian = QPixelFormat__ByteOrder(2)
)

//QPixelFormat::ColorModel
type QPixelFormat__ColorModel int64

const (
	QPixelFormat__RGB       = QPixelFormat__ColorModel(0)
	QPixelFormat__BGR       = QPixelFormat__ColorModel(1)
	QPixelFormat__Indexed   = QPixelFormat__ColorModel(2)
	QPixelFormat__Grayscale = QPixelFormat__ColorModel(3)
	QPixelFormat__CMYK      = QPixelFormat__ColorModel(4)
	QPixelFormat__HSL       = QPixelFormat__ColorModel(5)
	QPixelFormat__HSV       = QPixelFormat__ColorModel(6)
	QPixelFormat__YUV       = QPixelFormat__ColorModel(7)
	QPixelFormat__Alpha     = QPixelFormat__ColorModel(8)
)

//QPixelFormat::TypeInterpretation
type QPixelFormat__TypeInterpretation int64

const (
	QPixelFormat__UnsignedInteger = QPixelFormat__TypeInterpretation(0)
	QPixelFormat__UnsignedShort   = QPixelFormat__TypeInterpretation(1)
	QPixelFormat__UnsignedByte    = QPixelFormat__TypeInterpretation(2)
	QPixelFormat__FloatingPoint   = QPixelFormat__TypeInterpretation(3)
)

//QPixelFormat::YUVLayout
type QPixelFormat__YUVLayout int64

const (
	QPixelFormat__YUV444   = QPixelFormat__YUVLayout(0)
	QPixelFormat__YUV422   = QPixelFormat__YUVLayout(1)
	QPixelFormat__YUV411   = QPixelFormat__YUVLayout(2)
	QPixelFormat__YUV420P  = QPixelFormat__YUVLayout(3)
	QPixelFormat__YUV420SP = QPixelFormat__YUVLayout(4)
	QPixelFormat__YV12     = QPixelFormat__YUVLayout(5)
	QPixelFormat__UYVY     = QPixelFormat__YUVLayout(6)
	QPixelFormat__YUYV     = QPixelFormat__YUVLayout(7)
	QPixelFormat__NV12     = QPixelFormat__YUVLayout(8)
	QPixelFormat__NV21     = QPixelFormat__YUVLayout(9)
	QPixelFormat__IMC1     = QPixelFormat__YUVLayout(10)
	QPixelFormat__IMC2     = QPixelFormat__YUVLayout(11)
	QPixelFormat__IMC3     = QPixelFormat__YUVLayout(12)
	QPixelFormat__IMC4     = QPixelFormat__YUVLayout(13)
	QPixelFormat__Y8       = QPixelFormat__YUVLayout(14)
	QPixelFormat__Y16      = QPixelFormat__YUVLayout(15)
)

func NewQPixelFormat() *QPixelFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::QPixelFormat")
		}
	}()

	return NewQPixelFormatFromPointer(C.QPixelFormat_NewQPixelFormat())
}

func (ptr *QPixelFormat) AlphaPosition() QPixelFormat__AlphaPosition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::alphaPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__AlphaPosition(C.QPixelFormat_AlphaPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixelFormat) AlphaUsage() QPixelFormat__AlphaUsage {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::alphaUsage")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__AlphaUsage(C.QPixelFormat_AlphaUsage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixelFormat) ByteOrder() QPixelFormat__ByteOrder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::byteOrder")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__ByteOrder(C.QPixelFormat_ByteOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixelFormat) ColorModel() QPixelFormat__ColorModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::colorModel")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__ColorModel(C.QPixelFormat_ColorModel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixelFormat) Premultiplied() QPixelFormat__AlphaPremultiplied {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::premultiplied")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__AlphaPremultiplied(C.QPixelFormat_Premultiplied(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixelFormat) TypeInterpretation() QPixelFormat__TypeInterpretation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::typeInterpretation")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__TypeInterpretation(C.QPixelFormat_TypeInterpretation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixelFormat) YuvLayout() QPixelFormat__YUVLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPixelFormat::yuvLayout")
		}
	}()

	if ptr.Pointer() != nil {
		return QPixelFormat__YUVLayout(C.QPixelFormat_YuvLayout(ptr.Pointer()))
	}
	return 0
}
