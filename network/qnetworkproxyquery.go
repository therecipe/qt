package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkProxyQuery struct {
	ptr unsafe.Pointer
}

type QNetworkProxyQuery_ITF interface {
	QNetworkProxyQuery_PTR() *QNetworkProxyQuery
}

func (p *QNetworkProxyQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxyQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxyQuery(ptr QNetworkProxyQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyQuery_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) *QNetworkProxyQuery {
	var n = new(QNetworkProxyQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkProxyQuery) QNetworkProxyQuery_PTR() *QNetworkProxyQuery {
	return ptr
}

//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__TcpServer  = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest = QNetworkProxyQuery__QueryType(101)
)

func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery())
}

func NewQNetworkProxyQuery7(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery7(PointerFromQNetworkConfiguration(networkConfiguration), C.CString(hostname), C.int(port), C.CString(protocolTag), C.int(queryType)))
}

func NewQNetworkProxyQuery6(networkConfiguration QNetworkConfiguration_ITF, requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery6(PointerFromQNetworkConfiguration(networkConfiguration), core.PointerFromQUrl(requestUrl), C.int(queryType)))
}

func NewQNetworkProxyQuery5(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery5(PointerFromQNetworkProxyQuery(other)))
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery3(C.CString(hostname), C.int(port), C.CString(protocolTag), C.int(queryType)))
}

func NewQNetworkProxyQuery2(requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery2(core.PointerFromQUrl(requestUrl), C.int(queryType)))
}

func (ptr *QNetworkProxyQuery) LocalPort() int {
	defer qt.Recovering("QNetworkProxyQuery::localPort")

	if ptr.Pointer() != nil {
		return int(C.QNetworkProxyQuery_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) PeerHostName() string {
	defer qt.Recovering("QNetworkProxyQuery::peerHostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_PeerHostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) PeerPort() int {
	defer qt.Recovering("QNetworkProxyQuery::peerPort")

	if ptr.Pointer() != nil {
		return int(C.QNetworkProxyQuery_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) ProtocolTag() string {
	defer qt.Recovering("QNetworkProxyQuery::protocolTag")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_ProtocolTag(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) QueryType() QNetworkProxyQuery__QueryType {
	defer qt.Recovering("QNetworkProxyQuery::queryType")

	if ptr.Pointer() != nil {
		return QNetworkProxyQuery__QueryType(C.QNetworkProxyQuery_QueryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) SetLocalPort(port int) {
	defer qt.Recovering("QNetworkProxyQuery::setLocalPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetLocalPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::setNetworkConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetNetworkConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(networkConfiguration))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	defer qt.Recovering("QNetworkProxyQuery::setPeerHostName")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerHostName(ptr.Pointer(), C.CString(hostname))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {
	defer qt.Recovering("QNetworkProxyQuery::setPeerPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	defer qt.Recovering("QNetworkProxyQuery::setProtocolTag")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetProtocolTag(ptr.Pointer(), C.CString(protocolTag))
	}
}

func (ptr *QNetworkProxyQuery) SetQueryType(ty QNetworkProxyQuery__QueryType) {
	defer qt.Recovering("QNetworkProxyQuery::setQueryType")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetQueryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QNetworkProxyQuery) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkProxyQuery) Swap(other QNetworkProxyQuery_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::swap")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_Swap(ptr.Pointer(), PointerFromQNetworkProxyQuery(other))
	}
}

func (ptr *QNetworkProxyQuery) Url() *core.QUrl {
	defer qt.Recovering("QNetworkProxyQuery::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNetworkProxyQuery_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {
	defer qt.Recovering("QNetworkProxyQuery::~QNetworkProxyQuery")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_DestroyQNetworkProxyQuery(ptr.Pointer())
	}
}
