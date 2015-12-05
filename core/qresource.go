package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QResource struct {
	ptr unsafe.Pointer
}

type QResource_ITF interface {
	QResource_PTR() *QResource
}

func (p *QResource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QResource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQResource(ptr QResource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QResource_PTR().Pointer()
	}
	return nil
}

func NewQResourceFromPointer(ptr unsafe.Pointer) *QResource {
	var n = new(QResource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QResource) QResource_PTR() *QResource {
	return ptr
}

func QResource_RegisterResource(rccFileName string, mapRoot string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::registerResource")
		}
	}()

	return C.QResource_QResource_RegisterResource(C.CString(rccFileName), C.CString(mapRoot)) != 0
}

func QResource_UnregisterResource(rccFileName string, mapRoot string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::unregisterResource")
		}
	}()

	return C.QResource_QResource_UnregisterResource(C.CString(rccFileName), C.CString(mapRoot)) != 0
}

func NewQResource(file string, locale QLocale_ITF) *QResource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::QResource")
		}
	}()

	return NewQResourceFromPointer(C.QResource_NewQResource(C.CString(file), PointerFromQLocale(locale)))
}

func (ptr *QResource) AbsoluteFilePath() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::absoluteFilePath")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QResource_AbsoluteFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QResource) FileName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QResource_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QResource) IsCompressed() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::isCompressed")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QResource_IsCompressed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QResource) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QResource_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QResource) SetFileName(file string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::setFileName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QResource_SetFileName(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QResource) SetLocale(locale QLocale_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::setLocale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QResource_SetLocale(ptr.Pointer(), PointerFromQLocale(locale))
	}
}

func (ptr *QResource) DestroyQResource() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QResource::~QResource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QResource_DestroyQResource(ptr.Pointer())
	}
}
