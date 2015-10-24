package network

//#include "qsslcertificateextension.h"
import "C"
import (
	"unsafe"
)

type QSslCertificateExtension struct {
	ptr unsafe.Pointer
}

type QSslCertificateExtensionITF interface {
	QSslCertificateExtensionPTR() *QSslCertificateExtension
}

func (p *QSslCertificateExtension) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCertificateExtension) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtensionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtensionPTR().Pointer()
	}
	return nil
}

func QSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslCertificateExtension) QSslCertificateExtensionPTR() *QSslCertificateExtension {
	return ptr
}

func NewQSslCertificateExtension() *QSslCertificateExtension {
	return QSslCertificateExtensionFromPointer(unsafe.Pointer(C.QSslCertificateExtension_NewQSslCertificateExtension()))
}

func NewQSslCertificateExtension2(other QSslCertificateExtensionITF) *QSslCertificateExtension {
	return QSslCertificateExtensionFromPointer(unsafe.Pointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(C.QtObjectPtr(PointerFromQSslCertificateExtension(other)))))
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsCritical(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsSupported(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Oid() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Oid(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtensionITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslCertificateExtension(other)))
	}
}

func (ptr *QSslCertificateExtension) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(C.QtObjectPtr(ptr.Pointer()))
	}
}
