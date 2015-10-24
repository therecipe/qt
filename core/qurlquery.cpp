#include "qurlquery.h"
#include <QChar>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QUrlQuery>
#include "_cgo_export.h"

class MyQUrlQuery: public QUrlQuery {
public:
};

QtObjectPtr QUrlQuery_NewQUrlQuery(){
	return new QUrlQuery();
}

QtObjectPtr QUrlQuery_NewQUrlQuery3(char* queryString){
	return new QUrlQuery(QString(queryString));
}

QtObjectPtr QUrlQuery_NewQUrlQuery2(char* url){
	return new QUrlQuery(QUrl(QString(url)));
}

QtObjectPtr QUrlQuery_NewQUrlQuery4(QtObjectPtr other){
	return new QUrlQuery(*static_cast<QUrlQuery*>(other));
}

void QUrlQuery_AddQueryItem(QtObjectPtr ptr, char* key, char* value){
	static_cast<QUrlQuery*>(ptr)->addQueryItem(QString(key), QString(value));
}

char* QUrlQuery_AllQueryItemValues(QtObjectPtr ptr, char* key, int encoding){
	return static_cast<QUrlQuery*>(ptr)->allQueryItemValues(QString(key), static_cast<QUrl::ComponentFormattingOption>(encoding)).join("|").toUtf8().data();
}

void QUrlQuery_Clear(QtObjectPtr ptr){
	static_cast<QUrlQuery*>(ptr)->clear();
}

int QUrlQuery_IsEmpty(QtObjectPtr ptr){
	return static_cast<QUrlQuery*>(ptr)->isEmpty();
}

char* QUrlQuery_Query(QtObjectPtr ptr, int encoding){
	return static_cast<QUrlQuery*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(encoding)).toUtf8().data();
}

void QUrlQuery_RemoveAllQueryItems(QtObjectPtr ptr, char* key){
	static_cast<QUrlQuery*>(ptr)->removeAllQueryItems(QString(key));
}

void QUrlQuery_RemoveQueryItem(QtObjectPtr ptr, char* key){
	static_cast<QUrlQuery*>(ptr)->removeQueryItem(QString(key));
}

void QUrlQuery_SetQuery(QtObjectPtr ptr, char* queryString){
	static_cast<QUrlQuery*>(ptr)->setQuery(QString(queryString));
}

void QUrlQuery_SetQueryDelimiters(QtObjectPtr ptr, QtObjectPtr valueDelimiter, QtObjectPtr pairDelimiter){
	static_cast<QUrlQuery*>(ptr)->setQueryDelimiters(*static_cast<QChar*>(valueDelimiter), *static_cast<QChar*>(pairDelimiter));
}

void QUrlQuery_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QUrlQuery*>(ptr)->swap(*static_cast<QUrlQuery*>(other));
}

char* QUrlQuery_ToString(QtObjectPtr ptr, int encoding){
	return static_cast<QUrlQuery*>(ptr)->toString(static_cast<QUrl::ComponentFormattingOption>(encoding)).toUtf8().data();
}

void QUrlQuery_DestroyQUrlQuery(QtObjectPtr ptr){
	static_cast<QUrlQuery*>(ptr)->~QUrlQuery();
}

