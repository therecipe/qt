package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
	"unsafe"
)

type QXmlSchemaValidator struct {
	ptr unsafe.Pointer
}

type QXmlSchemaValidator_ITF interface {
	QXmlSchemaValidator_PTR() *QXmlSchemaValidator
}

func (p *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlSchemaValidator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlSchemaValidator(ptr QXmlSchemaValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaValidator_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) *QXmlSchemaValidator {
	var n = new(QXmlSchemaValidator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSchemaValidator) QXmlSchemaValidator_PTR() *QXmlSchemaValidator {
	return ptr
}

func NewQXmlSchemaValidator() *QXmlSchemaValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::QXmlSchemaValidator")
		}
	}()

	return NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator())
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::QXmlSchemaValidator")
		}
	}()

	return NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(PointerFromQXmlSchema(schema)))
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::messageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractMessageHandlerFromPointer(C.QXmlSchemaValidator_MessageHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::networkAccessManager")
		}
	}()

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QXmlSchemaValidator_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::setMessageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::setNetworkAccessManager")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchema_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::setSchema")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetSchema(ptr.Pointer(), PointerFromQXmlSchema(schema))
	}
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::setUriResolver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::uriResolver")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractUriResolverFromPointer(C.QXmlSchemaValidator_UriResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::validate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::validate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::validate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate(ptr.Pointer(), core.PointerFromQUrl(source)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlSchemaValidator::~QXmlSchemaValidator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_DestroyQXmlSchemaValidator(ptr.Pointer())
	}
}
