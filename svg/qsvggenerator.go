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

type QSvgGeneratorITF interface {
	gui.QPaintDeviceITF
	QSvgGeneratorPTR() *QSvgGenerator
}

func PointerFromQSvgGenerator(ptr QSvgGeneratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgGeneratorPTR().Pointer()
	}
	return nil
}

func QSvgGeneratorFromPointer(ptr unsafe.Pointer) *QSvgGenerator {
	var n = new(QSvgGenerator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSvgGenerator) QSvgGeneratorPTR() *QSvgGenerator {
	return ptr
}

func (ptr *QSvgGenerator) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Description(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QSvgGenerator_OutputDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSvgGenerator) Resolution() int {
	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Resolution(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(outputDevice)))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(C.QtObjectPtr(ptr.Pointer()), C.int(dpi))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(viewBox)))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(viewBox)))
	}
}

func (ptr *QSvgGenerator) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQSvgGenerator() *QSvgGenerator {
	return QSvgGeneratorFromPointer(unsafe.Pointer(C.QSvgGenerator_NewQSvgGenerator()))
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(C.QtObjectPtr(ptr.Pointer()))
	}
}
