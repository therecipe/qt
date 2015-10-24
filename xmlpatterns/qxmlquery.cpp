#include "qxmlquery.h"
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QAbstractUriResolver>
#include <QAbstractMessageHandler>
#include <QXmlItem>
#include <QXmlName>
#include <QUrl>
#include <QIODevice>
#include <QXmlNamePool>
#include <QNetworkAccessManager>
#include <QXmlQuery>
#include "_cgo_export.h"

class MyQXmlQuery: public QXmlQuery {
public:
};

QtObjectPtr QXmlQuery_NewQXmlQuery(){
	return new QXmlQuery();
}

QtObjectPtr QXmlQuery_NewQXmlQuery4(int queryLanguage, QtObjectPtr np){
	return new QXmlQuery(static_cast<QXmlQuery::QueryLanguage>(queryLanguage), *static_cast<QXmlNamePool*>(np));
}

QtObjectPtr QXmlQuery_NewQXmlQuery3(QtObjectPtr np){
	return new QXmlQuery(*static_cast<QXmlNamePool*>(np));
}

QtObjectPtr QXmlQuery_NewQXmlQuery2(QtObjectPtr other){
	return new QXmlQuery(*static_cast<QXmlQuery*>(other));
}

void QXmlQuery_BindVariable5(QtObjectPtr ptr, char* localName, QtObjectPtr device){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable4(QtObjectPtr ptr, char* localName, QtObjectPtr value){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable6(QtObjectPtr ptr, char* localName, QtObjectPtr query){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), *static_cast<QXmlQuery*>(query));
}

void QXmlQuery_BindVariable2(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr device){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable3(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr query){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlQuery*>(query));
}

int QXmlQuery_IsValid(QtObjectPtr ptr){
	return static_cast<QXmlQuery*>(ptr)->isValid();
}

QtObjectPtr QXmlQuery_MessageHandler(QtObjectPtr ptr){
	return static_cast<QXmlQuery*>(ptr)->messageHandler();
}

QtObjectPtr QXmlQuery_NetworkAccessManager(QtObjectPtr ptr){
	return static_cast<QXmlQuery*>(ptr)->networkAccessManager();
}

int QXmlQuery_QueryLanguage(QtObjectPtr ptr){
	return static_cast<QXmlQuery*>(ptr)->queryLanguage();
}

int QXmlQuery_SetFocus3(QtObjectPtr ptr, QtObjectPtr document){
	return static_cast<QXmlQuery*>(ptr)->setFocus(static_cast<QIODevice*>(document));
}

int QXmlQuery_SetFocus4(QtObjectPtr ptr, char* focus){
	return static_cast<QXmlQuery*>(ptr)->setFocus(QString(focus));
}

int QXmlQuery_SetFocus2(QtObjectPtr ptr, char* documentURI){
	return static_cast<QXmlQuery*>(ptr)->setFocus(QUrl(QString(documentURI)));
}

void QXmlQuery_SetFocus(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QXmlItem*>(item));
}

void QXmlQuery_SetInitialTemplateName2(QtObjectPtr ptr, char* localName){
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(QString(localName));
}

void QXmlQuery_SetInitialTemplateName(QtObjectPtr ptr, QtObjectPtr name){
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(*static_cast<QXmlName*>(name));
}

void QXmlQuery_SetMessageHandler(QtObjectPtr ptr, QtObjectPtr aMessageHandler){
	static_cast<QXmlQuery*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(aMessageHandler));
}

void QXmlQuery_SetNetworkAccessManager(QtObjectPtr ptr, QtObjectPtr newManager){
	static_cast<QXmlQuery*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(newManager));
}

void QXmlQuery_SetQuery(QtObjectPtr ptr, QtObjectPtr sourceCode, char* documentURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(static_cast<QIODevice*>(sourceCode), QUrl(QString(documentURI)));
}

void QXmlQuery_SetQuery3(QtObjectPtr ptr, char* sourceCode, char* documentURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(QString(sourceCode), QUrl(QString(documentURI)));
}

void QXmlQuery_SetQuery2(QtObjectPtr ptr, char* queryURI, char* baseURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(QUrl(QString(queryURI)), QUrl(QString(baseURI)));
}

void QXmlQuery_SetUriResolver(QtObjectPtr ptr, QtObjectPtr resolver){
	static_cast<QXmlQuery*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

QtObjectPtr QXmlQuery_UriResolver(QtObjectPtr ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlQuery*>(ptr)->uriResolver());
}

void QXmlQuery_DestroyQXmlQuery(QtObjectPtr ptr){
	static_cast<QXmlQuery*>(ptr)->~QXmlQuery();
}

