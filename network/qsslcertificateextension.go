package network

//#include "qsslcertificateextension.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSslCertificateExtension struct {
	ptr unsafe.Pointer
}

type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (p *QSslCertificateExtension) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCertificateExtension) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtension_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return ptr
}

func NewQSslCertificateExtension() *QSslCertificateExtension {
	return NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension())
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	return NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(PointerFromQSslCertificateExtension(other)))
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsCritical(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Oid() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Oid(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_Swap(ptr.Pointer(), PointerFromQSslCertificateExtension(other))
	}
}

func (ptr *QSslCertificateExtension) Value() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSslCertificateExtension_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(ptr.Pointer())
	}
}
