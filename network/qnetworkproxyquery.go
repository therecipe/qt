package network

//#include "qnetworkproxyquery.h"
import "C"
import (
	"unsafe"
)

type QNetworkProxyQuery struct {
	ptr unsafe.Pointer
}

type QNetworkProxyQueryITF interface {
	QNetworkProxyQueryPTR() *QNetworkProxyQuery
}

func (p *QNetworkProxyQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxyQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxyQuery(ptr QNetworkProxyQueryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyQueryPTR().Pointer()
	}
	return nil
}

func QNetworkProxyQueryFromPointer(ptr unsafe.Pointer) *QNetworkProxyQuery {
	var n = new(QNetworkProxyQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkProxyQuery) QNetworkProxyQueryPTR() *QNetworkProxyQuery {
	return ptr
}

//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int

var (
	QNetworkProxyQuery__TcpSocket  = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__TcpServer  = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest = QNetworkProxyQuery__QueryType(101)
)

func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	return QNetworkProxyQueryFromPointer(unsafe.Pointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery()))
}

func NewQNetworkProxyQuery7(networkConfiguration QNetworkConfigurationITF, hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	return QNetworkProxyQueryFromPointer(unsafe.Pointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery7(C.QtObjectPtr(PointerFromQNetworkConfiguration(networkConfiguration)), C.CString(hostname), C.int(port), C.CString(protocolTag), C.int(queryType))))
}

func NewQNetworkProxyQuery6(networkConfiguration QNetworkConfigurationITF, requestUrl string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	return QNetworkProxyQueryFromPointer(unsafe.Pointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery6(C.QtObjectPtr(PointerFromQNetworkConfiguration(networkConfiguration)), C.CString(requestUrl), C.int(queryType))))
}

func NewQNetworkProxyQuery5(other QNetworkProxyQueryITF) *QNetworkProxyQuery {
	return QNetworkProxyQueryFromPointer(unsafe.Pointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery5(C.QtObjectPtr(PointerFromQNetworkProxyQuery(other)))))
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	return QNetworkProxyQueryFromPointer(unsafe.Pointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery3(C.CString(hostname), C.int(port), C.CString(protocolTag), C.int(queryType))))
}

func NewQNetworkProxyQuery2(requestUrl string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	return QNetworkProxyQueryFromPointer(unsafe.Pointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery2(C.CString(requestUrl), C.int(queryType))))
}

func (ptr *QNetworkProxyQuery) LocalPort() int {
	if ptr.Pointer() != nil {
		return int(C.QNetworkProxyQuery_LocalPort(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) PeerHostName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_PeerHostName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) PeerPort() int {
	if ptr.Pointer() != nil {
		return int(C.QNetworkProxyQuery_PeerPort(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) ProtocolTag() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_ProtocolTag(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) QueryType() QNetworkProxyQuery__QueryType {
	if ptr.Pointer() != nil {
		return QNetworkProxyQuery__QueryType(C.QNetworkProxyQuery_QueryType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) SetLocalPort(port int) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetLocalPort(C.QtObjectPtr(ptr.Pointer()), C.int(port))
	}
}

func (ptr *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration QNetworkConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetNetworkConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkConfiguration(networkConfiguration)))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerHostName(C.QtObjectPtr(ptr.Pointer()), C.CString(hostname))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerPort(C.QtObjectPtr(ptr.Pointer()), C.int(port))
	}
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetProtocolTag(C.QtObjectPtr(ptr.Pointer()), C.CString(protocolTag))
	}
}

func (ptr *QNetworkProxyQuery) SetQueryType(ty QNetworkProxyQuery__QueryType) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetQueryType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QNetworkProxyQuery) SetUrl(url string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QNetworkProxyQuery) Swap(other QNetworkProxyQueryITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkProxyQuery(other)))
	}
}

func (ptr *QNetworkProxyQuery) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {
	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_DestroyQNetworkProxyQuery(C.QtObjectPtr(ptr.Pointer()))
	}
}
