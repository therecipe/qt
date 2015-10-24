package network

//#include "qsslcertificate.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSslCertificate struct {
	ptr unsafe.Pointer
}

type QSslCertificateITF interface {
	QSslCertificatePTR() *QSslCertificate
}

func (p *QSslCertificate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCertificate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCertificate(ptr QSslCertificateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificatePTR().Pointer()
	}
	return nil
}

func QSslCertificateFromPointer(ptr unsafe.Pointer) *QSslCertificate {
	var n = new(QSslCertificate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslCertificate) QSslCertificatePTR() *QSslCertificate {
	return ptr
}

//QSslCertificate::SubjectInfo
type QSslCertificate__SubjectInfo int

var (
	QSslCertificate__Organization               = QSslCertificate__SubjectInfo(0)
	QSslCertificate__CommonName                 = QSslCertificate__SubjectInfo(1)
	QSslCertificate__LocalityName               = QSslCertificate__SubjectInfo(2)
	QSslCertificate__OrganizationalUnitName     = QSslCertificate__SubjectInfo(3)
	QSslCertificate__CountryName                = QSslCertificate__SubjectInfo(4)
	QSslCertificate__StateOrProvinceName        = QSslCertificate__SubjectInfo(5)
	QSslCertificate__DistinguishedNameQualifier = QSslCertificate__SubjectInfo(6)
	QSslCertificate__SerialNumber               = QSslCertificate__SubjectInfo(7)
	QSslCertificate__EmailAddress               = QSslCertificate__SubjectInfo(8)
)

func NewQSslCertificate3(other QSslCertificateITF) *QSslCertificate {
	return QSslCertificateFromPointer(unsafe.Pointer(C.QSslCertificate_NewQSslCertificate3(C.QtObjectPtr(PointerFromQSslCertificate(other)))))
}

func (ptr *QSslCertificate) Clear() {
	if ptr.Pointer() != nil {
		C.QSslCertificate_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslCertificate) IsBlacklisted() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsBlacklisted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificate) Swap(other QSslCertificateITF) {
	if ptr.Pointer() != nil {
		C.QSslCertificate_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslCertificate(other)))
	}
}

func (ptr *QSslCertificate) DestroyQSslCertificate() {
	if ptr.Pointer() != nil {
		C.QSslCertificate_DestroyQSslCertificate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslCertificate) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificate) IsSelfSigned() bool {
	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsSelfSigned(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCertificate) IssuerInfo(subject QSslCertificate__SubjectInfo) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_IssuerInfo(C.QtObjectPtr(ptr.Pointer()), C.int(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfo2(attribute core.QByteArrayITF) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_IssuerInfo2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(attribute)))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfo(subject QSslCertificate__SubjectInfo) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_SubjectInfo(C.QtObjectPtr(ptr.Pointer()), C.int(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfo2(attribute core.QByteArrayITF) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_SubjectInfo2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(attribute)))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) ToText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificate_ToText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
