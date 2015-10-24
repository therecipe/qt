package xmlpatterns

//#include "qxmlquery.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QXmlQuery struct {
	ptr unsafe.Pointer
}

type QXmlQueryITF interface {
	QXmlQueryPTR() *QXmlQuery
}

func (p *QXmlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlQuery(ptr QXmlQueryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlQueryPTR().Pointer()
	}
	return nil
}

func QXmlQueryFromPointer(ptr unsafe.Pointer) *QXmlQuery {
	var n = new(QXmlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlQuery) QXmlQueryPTR() *QXmlQuery {
	return ptr
}

//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int

var (
	QXmlQuery__XQuery10                              = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               = QXmlQuery__QueryLanguage(4096)
)

func NewQXmlQuery() *QXmlQuery {
	return QXmlQueryFromPointer(unsafe.Pointer(C.QXmlQuery_NewQXmlQuery()))
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePoolITF) *QXmlQuery {
	return QXmlQueryFromPointer(unsafe.Pointer(C.QXmlQuery_NewQXmlQuery4(C.int(queryLanguage), C.QtObjectPtr(PointerFromQXmlNamePool(np)))))
}

func NewQXmlQuery3(np QXmlNamePoolITF) *QXmlQuery {
	return QXmlQueryFromPointer(unsafe.Pointer(C.QXmlQuery_NewQXmlQuery3(C.QtObjectPtr(PointerFromQXmlNamePool(np)))))
}

func NewQXmlQuery2(other QXmlQueryITF) *QXmlQuery {
	return QXmlQueryFromPointer(unsafe.Pointer(C.QXmlQuery_NewQXmlQuery2(C.QtObjectPtr(PointerFromQXmlQuery(other)))))
}

func (ptr *QXmlQuery) BindVariable5(localName string, device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(C.QtObjectPtr(ptr.Pointer()), C.CString(localName), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QXmlQuery) BindVariable4(localName string, value QXmlItemITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable4(C.QtObjectPtr(ptr.Pointer()), C.CString(localName), C.QtObjectPtr(PointerFromQXmlItem(value)))
	}
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQueryITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable6(C.QtObjectPtr(ptr.Pointer()), C.CString(localName), C.QtObjectPtr(PointerFromQXmlQuery(query)))
	}
}

func (ptr *QXmlQuery) BindVariable2(name QXmlNameITF, device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QXmlQuery) BindVariable(name QXmlNameITF, value QXmlItemITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.QtObjectPtr(PointerFromQXmlItem(value)))
	}
}

func (ptr *QXmlQuery) BindVariable3(name QXmlNameITF, query QXmlQueryITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)), C.QtObjectPtr(PointerFromQXmlQuery(query)))
	}
}

func (ptr *QXmlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		return QAbstractMessageHandlerFromPointer(unsafe.Pointer(C.QXmlQuery_MessageHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QXmlQuery_NetworkAccessManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {
	if ptr.Pointer() != nil {
		return QXmlQuery__QueryLanguage(C.QXmlQuery_QueryLanguage(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODeviceITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(document))) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus4(C.QtObjectPtr(ptr.Pointer()), C.CString(focus)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus2(documentURI string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus2(C.QtObjectPtr(ptr.Pointer()), C.CString(documentURI)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus(item QXmlItemITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlItem(item)))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName2(C.QtObjectPtr(ptr.Pointer()), C.CString(localName))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlNameITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlName(name)))
	}
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetMessageHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractMessageHandler(aMessageHandler)))
	}
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManagerITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetNetworkAccessManager(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQNetworkAccessManager(newManager)))
	}
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODeviceITF, documentURI string) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(sourceCode)), C.CString(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery3(sourceCode string, documentURI string) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery3(C.QtObjectPtr(ptr.Pointer()), C.CString(sourceCode), C.CString(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery2(queryURI string, baseURI string) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery2(C.QtObjectPtr(ptr.Pointer()), C.CString(queryURI), C.CString(baseURI))
	}
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolverITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetUriResolver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractUriResolver(resolver)))
	}
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		return QAbstractUriResolverFromPointer(unsafe.Pointer(C.QXmlQuery_UriResolver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {
	if ptr.Pointer() != nil {
		C.QXmlQuery_DestroyQXmlQuery(C.QtObjectPtr(ptr.Pointer()))
	}
}
