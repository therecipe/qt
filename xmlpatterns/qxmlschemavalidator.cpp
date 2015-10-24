#include "qxmlschemavalidator.h"
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QXmlSchema>
#include <QNetworkAccessManager>
#include <QString>
#include <QVariant>
#include <QAbstractUriResolver>
#include <QByteArray>
#include <QAbstractMessageHandler>
#include <QXmlSchemaValidator>
#include "_cgo_export.h"

class MyQXmlSchemaValidator: public QXmlSchemaValidator {
public:
};

QtObjectPtr QXmlSchemaValidator_NewQXmlSchemaValidator(){
	return new QXmlSchemaValidator();
}

QtObjectPtr QXmlSchemaValidator_NewQXmlSchemaValidator2(QtObjectPtr schema){
	return new QXmlSchemaValidator(*static_cast<QXmlSchema*>(schema));
}

QtObjectPtr QXmlSchemaValidator_MessageHandler(QtObjectPtr ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->messageHandler();
}

QtObjectPtr QXmlSchemaValidator_NetworkAccessManager(QtObjectPtr ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->networkAccessManager();
}

void QXmlSchemaValidator_SetMessageHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSchemaValidator*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchemaValidator_SetNetworkAccessManager(QtObjectPtr ptr, QtObjectPtr manager){
	static_cast<QXmlSchemaValidator*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchemaValidator_SetSchema(QtObjectPtr ptr, QtObjectPtr schema){
	static_cast<QXmlSchemaValidator*>(ptr)->setSchema(*static_cast<QXmlSchema*>(schema));
}

void QXmlSchemaValidator_SetUriResolver(QtObjectPtr ptr, QtObjectPtr resolver){
	static_cast<QXmlSchemaValidator*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

QtObjectPtr QXmlSchemaValidator_UriResolver(QtObjectPtr ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchemaValidator*>(ptr)->uriResolver());
}

int QXmlSchemaValidator_Validate2(QtObjectPtr ptr, QtObjectPtr source, char* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(static_cast<QIODevice*>(source), QUrl(QString(documentUri)));
}

int QXmlSchemaValidator_Validate3(QtObjectPtr ptr, QtObjectPtr data, char* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QByteArray*>(data), QUrl(QString(documentUri)));
}

int QXmlSchemaValidator_Validate(QtObjectPtr ptr, char* source){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(QUrl(QString(source)));
}

void QXmlSchemaValidator_DestroyQXmlSchemaValidator(QtObjectPtr ptr){
	static_cast<QXmlSchemaValidator*>(ptr)->~QXmlSchemaValidator();
}

