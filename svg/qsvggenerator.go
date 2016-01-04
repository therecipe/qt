package svg

//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSvgGenerator struct {
	gui.QPaintDevice
}

type QSvgGenerator_ITF interface {
	gui.QPaintDevice_ITF
	QSvgGenerator_PTR() *QSvgGenerator
}

func PointerFromQSvgGenerator(ptr QSvgGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgGenerator_PTR().Pointer()
	}
	return nil
}

func NewQSvgGeneratorFromPointer(ptr unsafe.Pointer) *QSvgGenerator {
	var n = new(QSvgGenerator)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSvgGenerator_") {
		n.SetObjectNameAbs("QSvgGenerator_" + qt.Identifier())
	}
	return n
}

func (ptr *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator {
	return ptr
}

func (ptr *QSvgGenerator) Description() string {
	defer qt.Recovering("QSvgGenerator::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	defer qt.Recovering("QSvgGenerator::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	defer qt.Recovering("QSvgGenerator::outputDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QSvgGenerator_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) Resolution() int {
	defer qt.Recovering("QSvgGenerator::resolution")

	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Resolution(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	defer qt.Recovering("QSvgGenerator::setDescription")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	defer qt.Recovering("QSvgGenerator::setFileName")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {
	defer qt.Recovering("QSvgGenerator::setOutputDevice")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(ptr.Pointer(), core.PointerFromQIODevice(outputDevice))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	defer qt.Recovering("QSvgGenerator::setResolution")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(ptr.Pointer(), C.int(dpi))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {
	defer qt.Recovering("QSvgGenerator::setSize")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	defer qt.Recovering("QSvgGenerator::setTitle")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {
	defer qt.Recovering("QSvgGenerator::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewBox))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {
	defer qt.Recovering("QSvgGenerator::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewBox))
	}
}

func (ptr *QSvgGenerator) Size() *core.QSize {
	defer qt.Recovering("QSvgGenerator::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSvgGenerator_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) Title() string {
	defer qt.Recovering("QSvgGenerator::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Title(ptr.Pointer()))
	}
	return ""
}

func NewQSvgGenerator() *QSvgGenerator {
	defer qt.Recovering("QSvgGenerator::QSvgGenerator")

	return NewQSvgGeneratorFromPointer(C.QSvgGenerator_NewQSvgGenerator())
}

func (ptr *QSvgGenerator) Metric(metric gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QSvgGenerator::metric")

	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Metric(ptr.Pointer(), C.int(metric)))
	}
	return 0
}

func (ptr *QSvgGenerator) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QSvgGenerator::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QSvgGenerator_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) ViewBox() *core.QRect {
	defer qt.Recovering("QSvgGenerator::viewBox")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSvgGenerator_ViewBox(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	defer qt.Recovering("QSvgGenerator::~QSvgGenerator")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(ptr.Pointer())
	}
}

func (ptr *QSvgGenerator) ObjectNameAbs() string {
	defer qt.Recovering("QSvgGenerator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSvgGenerator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
