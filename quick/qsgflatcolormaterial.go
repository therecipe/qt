package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func PointerFromQSGFlatColorMaterial(ptr QSGFlatColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGFlatColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) *QSGFlatColorMaterial {
	var n = new(QSGFlatColorMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial {
	return ptr
}

func (ptr *QSGFlatColorMaterial) Color() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGFlatColorMaterial::color")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGFlatColorMaterial_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) SetColor(color gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGFlatColorMaterial::setColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGFlatColorMaterial_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}
