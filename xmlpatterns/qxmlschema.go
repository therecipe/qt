package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
	"unsafe"
)

type QXmlSchema struct {
	ptr unsafe.Pointer
}

type QXmlSchema_ITF interface {
	QXmlSchema_PTR() *QXmlSchema
}

func (p *QXmlSchema) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlSchema) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlSchema(ptr QXmlSchema_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchema_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaFromPointer(ptr unsafe.Pointer) *QXmlSchema {
	var n = new(QXmlSchema)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSchema) QXmlSchema_PTR() *QXmlSchema {
	return ptr
}

func NewQXmlSchema() *QXmlSchema {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::QXmlSchema")
		}
	}()

	return NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema())
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::QXmlSchema")
		}
	}()

	return NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema2(PointerFromQXmlSchema(other)))
}

func (ptr *QXmlSchema) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchema_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchema_Load(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::messageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlSchema_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::networkAccessManager")
		}
	}()

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlSchema_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::setMessageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::setNetworkAccessManager")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::setUriResolver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchema_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::uriResolver")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlSchema_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchema::~QXmlSchema")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchema_DestroyQXmlSchema(ptr.Pointer())
	}
}
