#include "qxmlschema.h"
#include <QUrl>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QNetworkAccessManager>
#include <QAbstractMessageHandler>
#include <QByteArray>
#include <QModelIndex>
#include <QAbstractUriResolver>
#include <QXmlSchema>
#include "_cgo_export.h"

class MyQXmlSchema: public QXmlSchema {
public:
};

void* QXmlSchema_NewQXmlSchema(){
	return new QXmlSchema();
}

void* QXmlSchema_NewQXmlSchema2(void* other){
	return new QXmlSchema(*static_cast<QXmlSchema*>(other));
}

int QXmlSchema_IsValid(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->isValid();
}

int QXmlSchema_Load2(void* ptr, void* source, void* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

int QXmlSchema_Load3(void* ptr, void* data, void* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(documentUri));
}

int QXmlSchema_Load(void* ptr, void* source){
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QUrl*>(source));
}

void* QXmlSchema_MessageHandler(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->messageHandler();
}

void* QXmlSchema_NetworkAccessManager(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->networkAccessManager();
}

void QXmlSchema_SetMessageHandler(void* ptr, void* handler){
	static_cast<QXmlSchema*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchema_SetNetworkAccessManager(void* ptr, void* manager){
	static_cast<QXmlSchema*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchema_SetUriResolver(void* ptr, void* resolver){
	static_cast<QXmlSchema*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void* QXmlSchema_UriResolver(void* ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchema*>(ptr)->uriResolver());
}

void QXmlSchema_DestroyQXmlSchema(void* ptr){
	static_cast<QXmlSchema*>(ptr)->~QXmlSchema();
}

