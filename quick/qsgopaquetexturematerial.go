package quick

//#include "quick.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSGOpaqueTextureMaterial struct {
	QSGMaterial
}

type QSGOpaqueTextureMaterial_ITF interface {
	QSGMaterial_ITF
	QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial
}

func PointerFromQSGOpaqueTextureMaterial(ptr QSGOpaqueTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGOpaqueTextureMaterial {
	var n = new(QSGOpaqueTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial {
	return ptr
}

func (ptr *QSGOpaqueTextureMaterial) Filtering() QSGTexture__Filtering {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::filtering")
		}
	}()

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) HorizontalWrapMode() QSGTexture__WrapMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::horizontalWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) MipmapFiltering() QSGTexture__Filtering {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::mipmapFiltering")
		}
	}()

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetFiltering(filtering QSGTexture__Filtering) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::setFiltering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetFiltering(ptr.Pointer(), C.int(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetHorizontalWrapMode(mode QSGTexture__WrapMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::setHorizontalWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetHorizontalWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetMipmapFiltering(filtering QSGTexture__Filtering) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::setMipmapFiltering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetMipmapFiltering(ptr.Pointer(), C.int(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetTexture(texture QSGTexture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::setTexture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetVerticalWrapMode(mode QSGTexture__WrapMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::setVerticalWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetVerticalWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) Texture() *QSGTexture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::texture")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) VerticalWrapMode() QSGTexture__WrapMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpaqueTextureMaterial::verticalWrapMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}
