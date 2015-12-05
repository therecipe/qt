package svg

//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::description")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::outputDevice")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QSvgGenerator_OutputDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) Resolution() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::resolution")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSvgGenerator_Resolution(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setDescription")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setFileName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setOutputDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(ptr.Pointer(), core.PointerFromQIODevice(outputDevice))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(ptr.Pointer(), C.int(dpi))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setTitle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setViewBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewBox))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::setViewBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewBox))
	}
}

func (ptr *QSvgGenerator) Title() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::title")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Title(ptr.Pointer()))
	}
	return ""
}

func NewQSvgGenerator() *QSvgGenerator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::QSvgGenerator")
		}
	}()

	return NewQSvgGeneratorFromPointer(C.QSvgGenerator_NewQSvgGenerator())
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSvgGenerator::~QSvgGenerator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(ptr.Pointer())
	}
}
