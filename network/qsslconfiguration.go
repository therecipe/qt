package network

//#include "qsslconfiguration.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSslConfiguration struct {
	ptr unsafe.Pointer
}

type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (p *QSslConfiguration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslConfiguration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslConfiguration(ptr QSslConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQSslConfigurationFromPointer(ptr unsafe.Pointer) *QSslConfiguration {
	var n = new(QSslConfiguration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration {
	return ptr
}

//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int64

const (
	QSslConfiguration__NextProtocolNegotiationNone        = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

func NewQSslConfiguration() *QSslConfiguration {
	return NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration())
}

func NewQSslConfiguration2(other QSslConfiguration_ITF) *QSslConfiguration {
	return NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration2(PointerFromQSslConfiguration(other)))
}

func (ptr *QSslConfiguration) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslConfiguration_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslConfiguration) NextNegotiatedProtocol() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslConfiguration_NextNegotiatedProtocol(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) NextProtocolNegotiationStatus() QSslConfiguration__NextProtocolNegotiationStatus {
	if ptr.Pointer() != nil {
		return QSslConfiguration__NextProtocolNegotiationStatus(C.QSslConfiguration_NextProtocolNegotiationStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QSslConfiguration_PeerVerifyDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslConfiguration_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) SessionTicket() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslConfiguration_SessionTicket(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) SessionTicketLifeTimeHint() int {
	if ptr.Pointer() != nil {
		return int(C.QSslConfiguration_SessionTicketLifeTimeHint(ptr.Pointer()))
	}
	return 0
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslConfiguration) SetSessionTicket(sessionTicket core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetSessionTicket(ptr.Pointer(), core.PointerFromQByteArray(sessionTicket))
	}
}

func (ptr *QSslConfiguration) Swap(other QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_Swap(ptr.Pointer(), PointerFromQSslConfiguration(other))
	}
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_DestroyQSslConfiguration(ptr.Pointer())
	}
}
