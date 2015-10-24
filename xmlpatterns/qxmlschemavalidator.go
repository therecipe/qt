package xmlpatterns

//#include "qxmlschemavalidator.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QXmlSchemaValidator struct {
	ptr unsafe.Pointer
}

type QXmlSchemaValidatorITF interface {
	QXmlSchemaValidatorPTR() *QXmlSchemaValidator
}

func (p *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlSchemaValidator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlSchemaValidator(ptr QXmlSchemaValidatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaValidatorPTR().Pointer()
	}
	return nil
}

func QXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) *QXmlSchemaValidator {
	var n = new(QXmlSchemaValidator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlSchemaValidator) QXmlSchemaValidatorPTR() *QXmlSchemaValidator {
	return ptr
}

func NewQXmlSchemaValidator() *QXmlSchemaValidator {
	return QXmlSchemaValidatorFromPointer(unsafe.Pointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator()))
}

func NewQXmlSchemaValidator2(schema QXmlSchemaITF) *QXmlSchemaValidator {
	return QXmlSchemaValidatorFromPointer(unsafe.Pointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(C.QtObjectPtr(PointerFromQXmlSchema(schema)))))
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		return QAbstractMessageHandlerFromPointer(unsafe.Pointer(C.QXmlSchemaValidator_MessageHandler(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QXmlSchemaValidator_NetworkAccessManager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandlerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetMessageHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractMessageHandler(handler)))
	}
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManagerITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetNetworkAccessManager(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQNetworkAccessManager(manager)))
	}
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchemaITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetSchema(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlSchema(schema)))
	}
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolverITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetUriResolver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractUriResolver(resolver)))
	}
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		return QAbstractUriResolverFromPointer(unsafe.Pointer(C.QXmlSchemaValidator_UriResolver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODeviceITF, documentUri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(source)), C.CString(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate3(data core.QByteArrayITF, documentUri string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.CString(documentUri)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate(source string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlSchemaValidator_Validate(C.QtObjectPtr(ptr.Pointer()), C.CString(source)) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_DestroyQXmlSchemaValidator(C.QtObjectPtr(ptr.Pointer()))
	}
}
