package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSslError struct {
	ptr unsafe.Pointer
}

type QSslError_ITF interface {
	QSslError_PTR() *QSslError
}

func (p *QSslError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslError(ptr QSslError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslError_PTR().Pointer()
	}
	return nil
}

func NewQSslErrorFromPointer(ptr unsafe.Pointer) *QSslError {
	var n = new(QSslError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslError) QSslError_PTR() *QSslError {
	return ptr
}

//QSslError::SslError
type QSslError__SslError int64

const (
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
	defer qt.Recovering("QSslError::QSslError")

	return NewQSslErrorFromPointer(C.QSslError_NewQSslError())
}

func NewQSslError2(error QSslError__SslError) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return NewQSslErrorFromPointer(C.QSslError_NewQSslError2(C.int(error)))
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificate_ITF) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return NewQSslErrorFromPointer(C.QSslError_NewQSslError3(C.int(error), PointerFromQSslCertificate(certificate)))
}

func NewQSslError4(other QSslError_ITF) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return NewQSslErrorFromPointer(C.QSslError_NewQSslError4(PointerFromQSslError(other)))
}

func (ptr *QSslError) Error() QSslError__SslError {
	defer qt.Recovering("QSslError::error")

	if ptr.Pointer() != nil {
		return QSslError__SslError(C.QSslError_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslError) ErrorString() string {
	defer qt.Recovering("QSslError::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslError_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslError) Swap(other QSslError_ITF) {
	defer qt.Recovering("QSslError::swap")

	if ptr.Pointer() != nil {
		C.QSslError_Swap(ptr.Pointer(), PointerFromQSslError(other))
	}
}

func (ptr *QSslError) DestroyQSslError() {
	defer qt.Recovering("QSslError::~QSslError")

	if ptr.Pointer() != nil {
		C.QSslError_DestroyQSslError(ptr.Pointer())
	}
}
