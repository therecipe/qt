package quick

//#include "qsgflatcolormaterial.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterialITF interface {
	QSGMaterialITF
	QSGFlatColorMaterialPTR() *QSGFlatColorMaterial
}

func PointerFromQSGFlatColorMaterial(ptr QSGFlatColorMaterialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGFlatColorMaterialPTR().Pointer()
	}
	return nil
}

func QSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) *QSGFlatColorMaterial {
	var n = new(QSGFlatColorMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGFlatColorMaterial) QSGFlatColorMaterialPTR() *QSGFlatColorMaterial {
	return ptr
}

func (ptr *QSGFlatColorMaterial) SetColor(color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QSGFlatColorMaterial_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}
