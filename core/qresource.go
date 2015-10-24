package core

//#include "qresource.h"
import "C"
import (
	"unsafe"
)

type QResource struct {
	ptr unsafe.Pointer
}

type QResourceITF interface {
	QResourcePTR() *QResource
}

func (p *QResource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QResource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQResource(ptr QResourceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QResourcePTR().Pointer()
	}
	return nil
}

func QResourceFromPointer(ptr unsafe.Pointer) *QResource {
	var n = new(QResource)
	n.SetPointer(ptr)
	return n
}

func (ptr *QResource) QResourcePTR() *QResource {
	return ptr
}

func QResource_RegisterResource(rccFileName string, mapRoot string) bool {
	return C.QResource_QResource_RegisterResource(C.CString(rccFileName), C.CString(mapRoot)) != 0
}

func QResource_UnregisterResource(rccFileName string, mapRoot string) bool {
	return C.QResource_QResource_UnregisterResource(C.CString(rccFileName), C.CString(mapRoot)) != 0
}

func NewQResource(file string, locale QLocaleITF) *QResource {
	return QResourceFromPointer(unsafe.Pointer(C.QResource_NewQResource(C.CString(file), C.QtObjectPtr(PointerFromQLocale(locale)))))
}

func (ptr *QResource) AbsoluteFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QResource_AbsoluteFilePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QResource) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QResource_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QResource) IsCompressed() bool {
	if ptr.Pointer() != nil {
		return C.QResource_IsCompressed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QResource) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QResource_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QResource) SetFileName(file string) {
	if ptr.Pointer() != nil {
		C.QResource_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(file))
	}
}

func (ptr *QResource) SetLocale(locale QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QResource_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLocale(locale)))
	}
}

func (ptr *QResource) DestroyQResource() {
	if ptr.Pointer() != nil {
		C.QResource_DestroyQResource(C.QtObjectPtr(ptr.Pointer()))
	}
}
