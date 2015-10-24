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

type QSslConfigurationITF interface {
	QSslConfigurationPTR() *QSslConfiguration
}

func (p *QSslConfiguration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslConfiguration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslConfiguration(ptr QSslConfigurationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslConfigurationPTR().Pointer()
	}
	return nil
}

func QSslConfigurationFromPointer(ptr unsafe.Pointer) *QSslConfiguration {
	var n = new(QSslConfiguration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslConfiguration) QSslConfigurationPTR() *QSslConfiguration {
	return ptr
}

//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int

var (
	QSslConfiguration__NextProtocolNegotiationNone        = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

func NewQSslConfiguration() *QSslConfiguration {
	return QSslConfigurationFromPointer(unsafe.Pointer(C.QSslConfiguration_NewQSslConfiguration()))
}

func NewQSslConfiguration2(other QSslConfigurationITF) *QSslConfiguration {
	return QSslConfigurationFromPointer(unsafe.Pointer(C.QSslConfiguration_NewQSslConfiguration2(C.QtObjectPtr(PointerFromQSslConfiguration(other)))))
}

func (ptr *QSslConfiguration) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslConfiguration_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslConfiguration) NextProtocolNegotiationStatus() QSslConfiguration__NextProtocolNegotiationStatus {
	if ptr.Pointer() != nil {
		return QSslConfiguration__NextProtocolNegotiationStatus(C.QSslConfiguration_NextProtocolNegotiationStatus(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QSslConfiguration_PeerVerifyDepth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslConfiguration_PeerVerifyMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslConfiguration) SessionTicketLifeTimeHint() int {
	if ptr.Pointer() != nil {
		return int(C.QSslConfiguration_SessionTicketLifeTimeHint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfigurationITF) {
	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(C.QtObjectPtr(PointerFromQSslConfiguration(configuration)))
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificateITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslCertificate(certificate)))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyDepth(C.QtObjectPtr(ptr.Pointer()), C.int(depth))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QSslConfiguration) SetPrivateKey(key QSslKeyITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPrivateKey(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslKey(key)))
	}
}

func (ptr *QSslConfiguration) SetSessionTicket(sessionTicket core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetSessionTicket(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(sessionTicket)))
	}
}

func (ptr *QSslConfiguration) Swap(other QSslConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslConfiguration(other)))
	}
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {
	if ptr.Pointer() != nil {
		C.QSslConfiguration_DestroyQSslConfiguration(C.QtObjectPtr(ptr.Pointer()))
	}
}
