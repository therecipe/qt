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

type QXmlQuery_ITF interface {
	QXmlQuery_PTR() *QXmlQuery
}

func (p *QXmlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlQuery(ptr QXmlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlQuery_PTR().Pointer()
	}
	return nil
}

func NewQXmlQueryFromPointer(ptr unsafe.Pointer) *QXmlQuery {
	var n = new(QXmlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlQuery) QXmlQuery_PTR() *QXmlQuery {
	return ptr
}

//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int64

const (
	QXmlQuery__XQuery10                              = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               = QXmlQuery__QueryLanguage(4096)
)

func NewQXmlQuery() *QXmlQuery {
	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery())
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {
	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery4(C.int(queryLanguage), PointerFromQXmlNamePool(np)))
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {
	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery3(PointerFromQXmlNamePool(np)))
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {
	return NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery2(PointerFromQXmlQuery(other)))
}

func (ptr *QXmlQuery) BindVariable5(localName string, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(ptr.Pointer(), C.CString(localName), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable4(localName string, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable4(ptr.Pointer(), C.CString(localName), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable6(ptr.Pointer(), C.CString(localName), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) BindVariable2(name QXmlName_ITF, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable2(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlQuery_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlQuery_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {
	if ptr.Pointer() != nil {
		return QXmlQuery__QueryLanguage(C.QXmlQuery_QueryLanguage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODevice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus3(ptr.Pointer(), core.PointerFromQIODevice(document)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus4(ptr.Pointer(), C.CString(focus)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlQuery_SetFocus2(ptr.Pointer(), core.PointerFromQUrl(documentURI)) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(ptr.Pointer(), PointerFromQXmlItem(item))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName2(ptr.Pointer(), C.CString(localName))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(aMessageHandler))
	}
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(newManager))
	}
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODevice_ITF, documentURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery(ptr.Pointer(), core.PointerFromQIODevice(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery3(sourceCode string, documentURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery3(ptr.Pointer(), C.CString(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery2(queryURI core.QUrl_ITF, baseURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery2(ptr.Pointer(), core.PointerFromQUrl(queryURI), core.PointerFromQUrl(baseURI))
	}
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlQuery_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {
	if ptr.Pointer() != nil {
		C.QXmlQuery_DestroyQXmlQuery(ptr.Pointer())
	}
}
