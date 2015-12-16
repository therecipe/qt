package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QResource::registerResource")

	return C.QResource_QResource_RegisterResource(C.CString(rccFileName), C.CString(mapRoot)) != 0
}

func QResource_UnregisterResource(rccFileName string, mapRoot string) bool {
	defer qt.Recovering("QResource::unregisterResource")

	return C.QResource_QResource_UnregisterResource(C.CString(rccFileName), C.CString(mapRoot)) != 0
}

func NewQResource(file string, locale QLocale_ITF) *QResource {
	defer qt.Recovering("QResource::QResource")

	return NewQResourceFromPointer(C.QResource_NewQResource(C.CString(file), PointerFromQLocale(locale)))
}

func (ptr *QResource) AbsoluteFilePath() string {
	defer qt.Recovering("QResource::absoluteFilePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QResource_AbsoluteFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QResource) FileName() string {
	defer qt.Recovering("QResource::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QResource_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QResource) IsCompressed() bool {
	defer qt.Recovering("QResource::isCompressed")

	if ptr.Pointer() != nil {
		return C.QResource_IsCompressed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QResource) IsValid() bool {
	defer qt.Recovering("QResource::isValid")

	if ptr.Pointer() != nil {
		return C.QResource_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QResource) SetFileName(file string) {
	defer qt.Recovering("QResource::setFileName")

	if ptr.Pointer() != nil {
		C.QResource_SetFileName(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QResource) SetLocale(locale QLocale_ITF) {
	defer qt.Recovering("QResource::setLocale")

	if ptr.Pointer() != nil {
		C.QResource_SetLocale(ptr.Pointer(), PointerFromQLocale(locale))
	}
}

func (ptr *QResource) Size() int64 {
	defer qt.Recovering("QResource::size")

	if ptr.Pointer() != nil {
		return int64(C.QResource_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QResource) DestroyQResource() {
	defer qt.Recovering("QResource::~QResource")

	if ptr.Pointer() != nil {
		C.QResource_DestroyQResource(ptr.Pointer())
	}
}
