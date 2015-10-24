package xmlpatterns

//#include "qxmlschema.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QXmlSchema struct {
	ptr unsafe.Pointer
}

type QXmlSchemaITF interface {
	QXmlSchemaPTR() *QXmlSchema
}

func (p *QXmlSchema) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlSchema) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlSchema(ptr QXmlSchemaITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaPTR().Pointer()
	}
	return nil
}

func QXmlSchemaFromPointer(ptr unsafe.Pointer) *QXmlSchema {
	var n = new(QXmlSchema)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSchema) QXmlSchemaPTR() *QXmlSchema {
	return ptr
}

func NewQXmlSchema() *QXmlSchema {
	return QXmlSchemaFromPointer(unsafe.Pointer(C.QXmlSchema_NewQXmlSchema()))
}

func NewQXmlSchema2(other QXmlSchemaITF) *QXmlSchema {
	return QXmlSchemaFromPointer(unsafe.Pointer(C.QXmlSchema_NewQXmlSchema2(C.QtObjectPtr(PointerFromQXmlSchema(other)))))
}

func (ptr *QXmlSchema) DocumentUri() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSchema_DocumentUri(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlSchema) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load2(source core.QIODeviceITF, documentUri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(source)), C.CString(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load3(data core.QByteArrayITF, documentUri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.CString(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load(source string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(source)) != 0
	}
	return false
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		return QAbstractMessageHandlerFromPointer(unsafe.Pointer(C.QXmlSchema_MessageHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QXmlSchema_NetworkAccessManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetMessageHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractMessageHandler(handler)))
	}
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManagerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetNetworkAccessManager(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQNetworkAccessManager(manager)))
	}
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolverITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetUriResolver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractUriResolver(resolver)))
	}
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		return QAbstractUriResolverFromPointer(unsafe.Pointer(C.QXmlSchema_UriResolver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {
	if ptr.Pointer() != nil {
		C.QXmlSchema_DestroyQXmlSchema(C.QtObjectPtr(ptr.Pointer()))
	}
}
