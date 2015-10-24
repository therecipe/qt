package network

//#include "qsslerror.h"
import "C"
import (
	"unsafe"
)

type QSslError struct {
	ptr unsafe.Pointer
}

type QSslErrorITF interface {
	QSslErrorPTR() *QSslError
}

func (p *QSslError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslError(ptr QSslErrorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslErrorPTR().Pointer()
	}
	return nil
}

func QSslErrorFromPointer(ptr unsafe.Pointer) *QSslError {
	var n = new(QSslError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslError) QSslErrorPTR() *QSslError {
	return ptr
}

//QSslError::SslError
type QSslError__SslError int

var (
	QSslError__NoError                             = QSslError__SslError(0)
	QSslError__UnableToGetIssuerCertificate        = QSslError__SslError(1)
	QSslError__UnableToDecryptCertificateSignature = QSslError__SslError(2)
	QSslError__UnableToDecodeIssuerPublicKey       = QSslError__SslError(3)
	QSslError__CertificateSignatureFailed          = QSslError__SslError(4)
	QSslError__CertificateNotYetValid              = QSslError__SslError(5)
	QSslError__CertificateExpired                  = QSslError__SslError(6)
	QSslError__InvalidNotBeforeField               = QSslError__SslError(7)
	QSslError__InvalidNotAfterField                = QSslError__SslError(8)
	QSslError__SelfSignedCertificate               = QSslError__SslError(9)
	QSslError__SelfSignedCertificateInChain        = QSslError__SslError(10)
	QSslError__UnableToGetLocalIssuerCertificate   = QSslError__SslError(11)
	QSslError__UnableToVerifyFirstCertificate      = QSslError__SslError(12)
	QSslError__CertificateRevoked                  = QSslError__SslError(13)
	QSslError__InvalidCaCertificate                = QSslError__SslError(14)
	QSslError__PathLengthExceeded                  = QSslError__SslError(15)
	QSslError__InvalidPurpose                      = QSslError__SslError(16)
	QSslError__CertificateUntrusted                = QSslError__SslError(17)
	QSslError__CertificateRejected                 = QSslError__SslError(18)
	QSslError__SubjectIssuerMismatch               = QSslError__SslError(19)
	QSslError__AuthorityIssuerSerialNumberMismatch = QSslError__SslError(20)
	QSslError__NoPeerCertificate                   = QSslError__SslError(21)
	QSslError__HostNameMismatch                    = QSslError__SslError(22)
	QSslError__NoSslSupport                        = QSslError__SslError(23)
	QSslError__CertificateBlacklisted              = QSslError__SslError(24)
	QSslError__UnspecifiedError                    = QSslError__SslError(-1)
)

func NewQSslError() *QSslError {
	return QSslErrorFromPointer(unsafe.Pointer(C.QSslError_NewQSslError()))
}

func NewQSslError2(error QSslError__SslError) *QSslError {
	return QSslErrorFromPointer(unsafe.Pointer(C.QSslError_NewQSslError2(C.int(error))))
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificateITF) *QSslError {
	return QSslErrorFromPointer(unsafe.Pointer(C.QSslError_NewQSslError3(C.int(error), C.QtObjectPtr(PointerFromQSslCertificate(certificate)))))
}

func NewQSslError4(other QSslErrorITF) *QSslError {
	return QSslErrorFromPointer(unsafe.Pointer(C.QSslError_NewQSslError4(C.QtObjectPtr(PointerFromQSslError(other)))))
}

func (ptr *QSslError) Error() QSslError__SslError {
	if ptr.Pointer() != nil {
		return QSslError__SslError(C.QSslError_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslError) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslError_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslError) Swap(other QSslErrorITF) {
	if ptr.Pointer() != nil {
		C.QSslError_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslError(other)))
	}
}

func (ptr *QSslError) DestroyQSslError() {
	if ptr.Pointer() != nil {
		C.QSslError_DestroyQSslError(C.QtObjectPtr(ptr.Pointer()))
	}
}
