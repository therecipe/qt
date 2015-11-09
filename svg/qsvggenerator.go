package svg

//#include "qsvggenerator.h"
import "C"
import (
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
	return n
}

func (ptr *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator {
	return ptr
}

func (ptr *QSvgGenerator) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QSvgGenerator_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) Resolution() int {
	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Resolution(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(ptr.Pointer(), core.PointerFromQIODevice(outputDevice))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(ptr.Pointer(), C.int(dpi))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewBox))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewBox))
	}
}

func (ptr *QSvgGenerator) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Title(ptr.Pointer()))
	}
	return ""
}

func NewQSvgGenerator() *QSvgGenerator {
	return NewQSvgGeneratorFromPointer(C.QSvgGenerator_NewQSvgGenerator())
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(ptr.Pointer())
	}
}
