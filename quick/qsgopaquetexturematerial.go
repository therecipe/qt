package quick

//#include "qsgopaquetexturematerial.h"
import "C"
import (
	"unsafe"
)

type QSGOpaqueTextureMaterial struct {
	QSGMaterial
}

type QSGOpaqueTextureMaterialITF interface {
	QSGMaterialITF
	QSGOpaqueTextureMaterialPTR() *QSGOpaqueTextureMaterial
}

func PointerFromQSGOpaqueTextureMaterial(ptr QSGOpaqueTextureMaterialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterialPTR().Pointer()
	}
	return nil
}

func QSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGOpaqueTextureMaterial {
	var n = new(QSGOpaqueTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterialPTR() *QSGOpaqueTextureMaterial {
	return ptr
}

func (ptr *QSGOpaqueTextureMaterial) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_Filtering(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_HorizontalWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_MipmapFiltering(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetFiltering(C.QtObjectPtr(ptr.Pointer()), C.int(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetHorizontalWrapMode(mode QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetHorizontalWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetMipmapFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetMipmapFiltering(C.QtObjectPtr(ptr.Pointer()), C.int(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetTexture(texture QSGTextureITF) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetTexture(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGTexture(texture)))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetVerticalWrapMode(mode QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetVerticalWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QSGOpaqueTextureMaterial_Texture(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_VerticalWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
