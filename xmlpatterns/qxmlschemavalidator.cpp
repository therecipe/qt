#include "qxmlschemavalidator.h"
#include <QAbstractMessageHandler>
#include <QVariant>
#include <QAbstractUriResolver>
#include <QModelIndex>
#include <QXmlSchema>
#include <QByteArray>
#include <QNetworkAccessManager>
#include <QIODevice>
#include <QString>
#include <QUrl>
#include <QXmlSchemaValidator>
#include "_cgo_export.h"

class MyQXmlSchemaValidator: public QXmlSchemaValidator {
public:
};

void* QXmlSchemaValidator_NewQXmlSchemaValidator(){
	return new QXmlSchemaValidator();
}

void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema){
	return new QXmlSchemaValidator(*static_cast<QXmlSchema*>(schema));
}

void* QXmlSchemaValidator_MessageHandler(void* ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->messageHandler();
}

void* QXmlSchemaValidator_NetworkAccessManager(void* ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->networkAccessManager();
}

void QXmlSchemaValidator_SetMessageHandler(void* ptr, void* handler){
	static_cast<QXmlSchemaValidator*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchemaValidator_SetNetworkAccessManager(void* ptr, void* manager){
	static_cast<QXmlSchemaValidator*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchemaValidator_SetSchema(void* ptr, void* schema){
	static_cast<QXmlSchemaValidator*>(ptr)->setSchema(*static_cast<QXmlSchema*>(schema));
}

void QXmlSchemaValidator_SetUriResolver(void* ptr, void* resolver){
	static_cast<QXmlSchemaValidator*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void* QXmlSchemaValidator_UriResolver(void* ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchemaValidator*>(ptr)->uriResolver());
}

int QXmlSchemaValidator_Validate2(void* ptr, void* source, void* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

int QXmlSchemaValidator_Validate3(void* ptr, void* data, void* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(documentUri));
}

int QXmlSchemaValidator_Validate(void* ptr, void* source){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QUrl*>(source));
}

void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr){
	static_cast<QXmlSchemaValidator*>(ptr)->~QXmlSchemaValidator();
}

