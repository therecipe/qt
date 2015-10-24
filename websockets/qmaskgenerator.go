package websockets

//#include "qmaskgenerator.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMaskGenerator struct {
	core.QObject
}

type QMaskGeneratorITF interface {
	core.QObjectITF
	QMaskGeneratorPTR() *QMaskGenerator
}

func PointerFromQMaskGenerator(ptr QMaskGeneratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGeneratorPTR().Pointer()
	}
	return nil
}

func QMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = new(QMaskGenerator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMaskGenerator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMaskGenerator) QMaskGeneratorPTR() *QMaskGenerator {
	return ptr
}

func (ptr *QMaskGenerator) Seed() bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Seed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGenerator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
