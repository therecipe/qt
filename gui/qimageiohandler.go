package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QImageIOHandler struct {
	ptr unsafe.Pointer
}

type QImageIOHandler_ITF interface {
	QImageIOHandler_PTR() *QImageIOHandler
}

func (p *QImageIOHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageIOHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageIOHandler(ptr QImageIOHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageIOHandler_PTR().Pointer()
	}
	return nil
}

func NewQImageIOHandlerFromPointer(ptr unsafe.Pointer) *QImageIOHandler {
	var n = new(QImageIOHandler)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageIOHandler) QImageIOHandler_PTR() *QImageIOHandler {
	return ptr
}

//QImageIOHandler::ImageOption
type QImageIOHandler__ImageOption int64

const (
	QImageIOHandler__Size                 = QImageIOHandler__ImageOption(0)
	QImageIOHandler__ClipRect             = QImageIOHandler__ImageOption(1)
	QImageIOHandler__Description          = QImageIOHandler__ImageOption(2)
	QImageIOHandler__ScaledClipRect       = QImageIOHandler__ImageOption(3)
	QImageIOHandler__ScaledSize           = QImageIOHandler__ImageOption(4)
	QImageIOHandler__CompressionRatio     = QImageIOHandler__ImageOption(5)
	QImageIOHandler__Gamma                = QImageIOHandler__ImageOption(6)
	QImageIOHandler__Quality              = QImageIOHandler__ImageOption(7)
	QImageIOHandler__Name                 = QImageIOHandler__ImageOption(8)
	QImageIOHandler__SubType              = QImageIOHandler__ImageOption(9)
	QImageIOHandler__IncrementalReading   = QImageIOHandler__ImageOption(10)
	QImageIOHandler__Endianness           = QImageIOHandler__ImageOption(11)
	QImageIOHandler__Animation            = QImageIOHandler__ImageOption(12)
	QImageIOHandler__BackgroundColor      = QImageIOHandler__ImageOption(13)
	QImageIOHandler__ImageFormat          = QImageIOHandler__ImageOption(14)
	QImageIOHandler__SupportedSubTypes    = QImageIOHandler__ImageOption(15)
	QImageIOHandler__OptimizedWrite       = QImageIOHandler__ImageOption(16)
	QImageIOHandler__ProgressiveScanWrite = QImageIOHandler__ImageOption(17)
	QImageIOHandler__ImageTransformation  = QImageIOHandler__ImageOption(18)
	QImageIOHandler__TransformedByDefault = QImageIOHandler__ImageOption(19)
)

//QImageIOHandler::Transformation
type QImageIOHandler__Transformation int64

const (
	QImageIOHandler__TransformationNone              = QImageIOHandler__Transformation(0)
	QImageIOHandler__TransformationMirror            = QImageIOHandler__Transformation(1)
	QImageIOHandler__TransformationFlip              = QImageIOHandler__Transformation(2)
	QImageIOHandler__TransformationRotate180         = QImageIOHandler__Transformation(QImageIOHandler__TransformationMirror | QImageIOHandler__TransformationFlip)
	QImageIOHandler__TransformationRotate90          = QImageIOHandler__Transformation(4)
	QImageIOHandler__TransformationMirrorAndRotate90 = QImageIOHandler__Transformation(QImageIOHandler__TransformationMirror | QImageIOHandler__TransformationRotate90)
	QImageIOHandler__TransformationFlipAndRotate90   = QImageIOHandler__Transformation(QImageIOHandler__TransformationFlip | QImageIOHandler__TransformationRotate90)
	QImageIOHandler__TransformationRotate270         = QImageIOHandler__Transformation(QImageIOHandler__TransformationRotate180 | QImageIOHandler__TransformationRotate90)
)
