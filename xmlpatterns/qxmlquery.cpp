#include "qxmlquery.h"
#include <QXmlNamePool>
#include <QXmlItem>
#include <QUrl>
#include <QNetworkAccessManager>
#include <QXmlName>
#include <QIODevice>
#include <QAbstractUriResolver>
#include <QAbstractMessageHandler>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QXmlQuery>
#include "_cgo_export.h"

class MyQXmlQuery: public QXmlQuery {
public:
};

void* QXmlQuery_NewQXmlQuery(){
	return new QXmlQuery();
}

void* QXmlQuery_NewQXmlQuery4(int queryLanguage, void* np){
	return new QXmlQuery(static_cast<QXmlQuery::QueryLanguage>(queryLanguage), *static_cast<QXmlNamePool*>(np));
}

void* QXmlQuery_NewQXmlQuery3(void* np){
	return new QXmlQuery(*static_cast<QXmlNamePool*>(np));
}

void* QXmlQuery_NewQXmlQuery2(void* other){
	return new QXmlQuery(*static_cast<QXmlQuery*>(other));
}

void QXmlQuery_BindVariable5(void* ptr, char* localName, void* device){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable4(void* ptr, char* localName, void* value){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable6(void* ptr, char* localName, void* query){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), *static_cast<QXmlQuery*>(query));
}

void QXmlQuery_BindVariable2(void* ptr, void* name, void* device){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable(void* ptr, void* name, void* value){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable3(void* ptr, void* name, void* query){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlQuery*>(query));
}

int QXmlQuery_IsValid(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->isValid();
}

void* QXmlQuery_MessageHandler(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->messageHandler();
}

void* QXmlQuery_NetworkAccessManager(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->networkAccessManager();
}

int QXmlQuery_QueryLanguage(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->queryLanguage();
}

int QXmlQuery_SetFocus3(void* ptr, void* document){
	return static_cast<QXmlQuery*>(ptr)->setFocus(static_cast<QIODevice*>(document));
}

int QXmlQuery_SetFocus4(void* ptr, char* focus){
	return static_cast<QXmlQuery*>(ptr)->setFocus(QString(focus));
}

int QXmlQuery_SetFocus2(void* ptr, void* documentURI){
	return static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetFocus(void* ptr, void* item){
	static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QXmlItem*>(item));
}

void QXmlQuery_SetInitialTemplateName2(void* ptr, char* localName){
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(QString(localName));
}

void QXmlQuery_SetInitialTemplateName(void* ptr, void* name){
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(*static_cast<QXmlName*>(name));
}

void QXmlQuery_SetMessageHandler(void* ptr, void* aMessageHandler){
	static_cast<QXmlQuery*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(aMessageHandler));
}

void QXmlQuery_SetNetworkAccessManager(void* ptr, void* newManager){
	static_cast<QXmlQuery*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(newManager));
}

void QXmlQuery_SetQuery(void* ptr, void* sourceCode, void* documentURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(static_cast<QIODevice*>(sourceCode), *static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetQuery3(void* ptr, char* sourceCode, void* documentURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(QString(sourceCode), *static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetQuery2(void* ptr, void* queryURI, void* baseURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(*static_cast<QUrl*>(queryURI), *static_cast<QUrl*>(baseURI));
}

void QXmlQuery_SetUriResolver(void* ptr, void* resolver){
	static_cast<QXmlQuery*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void* QXmlQuery_UriResolver(void* ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlQuery*>(ptr)->uriResolver());
}

void QXmlQuery_DestroyQXmlQuery(void* ptr){
	static_cast<QXmlQuery*>(ptr)->~QXmlQuery();
}

