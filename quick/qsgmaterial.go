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

type QSGMaterialITF interface {
	QSGMaterialPTR() *QSGMaterial
}

func (p *QSGMaterial) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterial) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterial(ptr QSGMaterialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialPTR().Pointer()
	}
	return nil
}

func QSGMaterialFromPointer(ptr unsafe.Pointer) *QSGMaterial {
	var n = new(QSGMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterial) QSGMaterialPTR() *QSGMaterial {
	return ptr
}

//QSGMaterial::Flag
type QSGMaterial__Flag int

var (
	QSGMaterial__Blending                          = QSGMaterial__Flag(0x0001)
	QSGMaterial__RequiresDeterminant               = QSGMaterial__Flag(0x0002)
	QSGMaterial__RequiresFullMatrixExceptTranslate = QSGMaterial__Flag(0x0004 | QSGMaterial__RequiresDeterminant)
	QSGMaterial__RequiresFullMatrix                = QSGMaterial__Flag(0x0008 | QSGMaterial__RequiresFullMatrixExceptTranslate)
	QSGMaterial__CustomCompileStep                 = QSGMaterial__Flag(0x0010)
)

func (ptr *QSGMaterial) Compare(other QSGMaterialITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSGMaterial_Compare(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGMaterial(other))))
	}
	return 0
}

func (ptr *QSGMaterial) CreateShader() *QSGMaterialShader {
	if ptr.Pointer() != nil {
		return QSGMaterialShaderFromPointer(unsafe.Pointer(C.QSGMaterial_CreateShader(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGMaterial) Flags() QSGMaterial__Flag {
	if ptr.Pointer() != nil {
		return QSGMaterial__Flag(C.QSGMaterial_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGMaterial) SetFlag(flags QSGMaterial__Flag, on bool) {
	if ptr.Pointer() != nil {
		C.QSGMaterial_SetFlag(C.QtObjectPtr(ptr.Pointer()), C.int(flags), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSGMaterial) Type() *QSGMaterialType {
	if ptr.Pointer() != nil {
		return QSGMaterialTypeFromPointer(unsafe.Pointer(C.QSGMaterial_Type(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
