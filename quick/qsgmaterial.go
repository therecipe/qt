package quick

//#include "qsgmaterial.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSGMaterial struct {
	ptr unsafe.Pointer
}

type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (p *QSGMaterial) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterial) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterial(ptr QSGMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialFromPointer(ptr unsafe.Pointer) *QSGMaterial {
	var n = new(QSGMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterial) QSGMaterial_PTR() *QSGMaterial {
	return ptr
}

//QSGMaterial::Flag
type QSGMaterial__Flag int64

const (
	QSGMaterial__Blending                          = QSGMaterial__Flag(0x0001)
	QSGMaterial__RequiresDeterminant               = QSGMaterial__Flag(0x0002)
	QSGMaterial__RequiresFullMatrixExceptTranslate = QSGMaterial__Flag(0x0004 | QSGMaterial__RequiresDeterminant)
	QSGMaterial__RequiresFullMatrix                = QSGMaterial__Flag(0x0008 | QSGMaterial__RequiresFullMatrixExceptTranslate)
	QSGMaterial__CustomCompileStep                 = QSGMaterial__Flag(0x0010)
)

func (ptr *QSGMaterial) Compare(other QSGMaterial_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSGMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGMaterial) Flags() QSGMaterial__Flag {
	if ptr.Pointer() != nil {
		return QSGMaterial__Flag(C.QSGMaterial_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGMaterial) SetFlag(flags QSGMaterial__Flag, on bool) {
	if ptr.Pointer() != nil {
		C.QSGMaterial_SetFlag(ptr.Pointer(), C.int(flags), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSGMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGMaterial_Type(ptr.Pointer()))
	}
	return nil
}
