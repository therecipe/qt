package websockets

//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMaskGenerator struct {
	core.QObject
}

type QMaskGenerator_ITF interface {
	core.QObject_ITF
	QMaskGenerator_PTR() *QMaskGenerator
}

func PointerFromQMaskGenerator(ptr QMaskGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGenerator_PTR().Pointer()
	}
	return nil
}

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = new(QMaskGenerator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMaskGenerator_") {
		n.SetObjectName("QMaskGenerator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMaskGenerator) QMaskGenerator_PTR() *QMaskGenerator {
	return ptr
}

func (ptr *QMaskGenerator) Seed() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMaskGenerator::seed")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Seed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMaskGenerator::~QMaskGenerator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
