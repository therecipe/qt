#include "qurlquery.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QChar>
#include <QUrlQuery>
#include "_cgo_export.h"

class MyQUrlQuery: public QUrlQuery {
public:
};

void* QUrlQuery_NewQUrlQuery(){
	return new QUrlQuery();
}

void* QUrlQuery_NewQUrlQuery3(char* queryString){
	return new QUrlQuery(QString(queryString));
}

void* QUrlQuery_NewQUrlQuery2(void* url){
	return new QUrlQuery(*static_cast<QUrl*>(url));
}

void* QUrlQuery_NewQUrlQuery4(void* other){
	return new QUrlQuery(*static_cast<QUrlQuery*>(other));
}

void QUrlQuery_AddQueryItem(void* ptr, char* key, char* value){
	static_cast<QUrlQuery*>(ptr)->addQueryItem(QString(key), QString(value));
}

char* QUrlQuery_AllQueryItemValues(void* ptr, char* key, int encoding){
	return static_cast<QUrlQuery*>(ptr)->allQueryItemValues(QString(key), static_cast<QUrl::ComponentFormattingOption>(encoding)).join("|").toUtf8().data();
}

void QUrlQuery_Clear(void* ptr){
	static_cast<QUrlQuery*>(ptr)->clear();
}

int QUrlQuery_IsEmpty(void* ptr){
	return static_cast<QUrlQuery*>(ptr)->isEmpty();
}

char* QUrlQuery_Query(void* ptr, int encoding){
	return static_cast<QUrlQuery*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(encoding)).toUtf8().data();
}

void QUrlQuery_RemoveAllQueryItems(void* ptr, char* key){
	static_cast<QUrlQuery*>(ptr)->removeAllQueryItems(QString(key));
}

void QUrlQuery_RemoveQueryItem(void* ptr, char* key){
	static_cast<QUrlQuery*>(ptr)->removeQueryItem(QString(key));
}

void QUrlQuery_SetQuery(void* ptr, char* queryString){
	static_cast<QUrlQuery*>(ptr)->setQuery(QString(queryString));
}

void QUrlQuery_SetQueryDelimiters(void* ptr, void* valueDelimiter, void* pairDelimiter){
	static_cast<QUrlQuery*>(ptr)->setQueryDelimiters(*static_cast<QChar*>(valueDelimiter), *static_cast<QChar*>(pairDelimiter));
}

void QUrlQuery_Swap(void* ptr, void* other){
	static_cast<QUrlQuery*>(ptr)->swap(*static_cast<QUrlQuery*>(other));
}

char* QUrlQuery_ToString(void* ptr, int encoding){
	return static_cast<QUrlQuery*>(ptr)->toString(static_cast<QUrl::ComponentFormattingOption>(encoding)).toUtf8().data();
}

void QUrlQuery_DestroyQUrlQuery(void* ptr){
	static_cast<QUrlQuery*>(ptr)->~QUrlQuery();
}

