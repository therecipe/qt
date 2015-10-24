#include "qxmlschema.h"
#include <QAbstractMessageHandler>
#include <QNetworkAccessManager>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractUriResolver>
#include <QByteArray>
#include <QXmlSchema>
#include "_cgo_export.h"

class MyQXmlSchema: public QXmlSchema {
public:
};

QtObjectPtr QXmlSchema_NewQXmlSchema(){
	return new QXmlSchema();
}

QtObjectPtr QXmlSchema_NewQXmlSchema2(QtObjectPtr other){
	return new QXmlSchema(*static_cast<QXmlSchema*>(other));
}

char* QXmlSchema_DocumentUri(QtObjectPtr ptr){
	return static_cast<QXmlSchema*>(ptr)->documentUri().toString().toUtf8().data();
}

int QXmlSchema_IsValid(QtObjectPtr ptr){
	return static_cast<QXmlSchema*>(ptr)->isValid();
}

int QXmlSchema_Load2(QtObjectPtr ptr, QtObjectPtr source, char* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(static_cast<QIODevice*>(source), QUrl(QString(documentUri)));
}

int QXmlSchema_Load3(QtObjectPtr ptr, QtObjectPtr data, char* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QByteArray*>(data), QUrl(QString(documentUri)));
}

int QXmlSchema_Load(QtObjectPtr ptr, char* source){
	return static_cast<QXmlSchema*>(ptr)->load(QUrl(QString(source)));
}

QtObjectPtr QXmlSchema_MessageHandler(QtObjectPtr ptr){
	return static_cast<QXmlSchema*>(ptr)->messageHandler();
}

QtObjectPtr QXmlSchema_NetworkAccessManager(QtObjectPtr ptr){
	return static_cast<QXmlSchema*>(ptr)->networkAccessManager();
}

void QXmlSchema_SetMessageHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSchema*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchema_SetNetworkAccessManager(QtObjectPtr ptr, QtObjectPtr manager){
	static_cast<QXmlSchema*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchema_SetUriResolver(QtObjectPtr ptr, QtObjectPtr resolver){
	static_cast<QXmlSchema*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

QtObjectPtr QXmlSchema_UriResolver(QtObjectPtr ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchema*>(ptr)->uriResolver());
}

void QXmlSchema_DestroyQXmlSchema(QtObjectPtr ptr){
	static_cast<QXmlSchema*>(ptr)->~QXmlSchema();
}

