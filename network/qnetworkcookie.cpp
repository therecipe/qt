#include "qnetworkcookie.h"
#include <QDate>
#include <QByteArray>
#include <QDateTime>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkCookie>
#include "_cgo_export.h"

class MyQNetworkCookie: public QNetworkCookie {
public:
};

void* QNetworkCookie_NewQNetworkCookie(void* name, void* value){
	return new QNetworkCookie(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

void* QNetworkCookie_NewQNetworkCookie2(void* other){
	return new QNetworkCookie(*static_cast<QNetworkCookie*>(other));
}

char* QNetworkCookie_Domain(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->domain().toUtf8().data();
}

void* QNetworkCookie_ExpirationDate(void* ptr){
	return new QDateTime(static_cast<QNetworkCookie*>(ptr)->expirationDate());
}

int QNetworkCookie_HasSameIdentifier(void* ptr, void* other){
	return static_cast<QNetworkCookie*>(ptr)->hasSameIdentifier(*static_cast<QNetworkCookie*>(other));
}

int QNetworkCookie_IsHttpOnly(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->isHttpOnly();
}

int QNetworkCookie_IsSecure(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->isSecure();
}

int QNetworkCookie_IsSessionCookie(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->isSessionCookie();
}

void* QNetworkCookie_Name(void* ptr){
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->name());
}

void QNetworkCookie_Normalize(void* ptr, void* url){
	static_cast<QNetworkCookie*>(ptr)->normalize(*static_cast<QUrl*>(url));
}

char* QNetworkCookie_Path(void* ptr){
	return static_cast<QNetworkCookie*>(ptr)->path().toUtf8().data();
}

void QNetworkCookie_SetDomain(void* ptr, char* domain){
	static_cast<QNetworkCookie*>(ptr)->setDomain(QString(domain));
}

void QNetworkCookie_SetExpirationDate(void* ptr, void* date){
	static_cast<QNetworkCookie*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(date));
}

void QNetworkCookie_SetHttpOnly(void* ptr, int enable){
	static_cast<QNetworkCookie*>(ptr)->setHttpOnly(enable != 0);
}

void QNetworkCookie_SetName(void* ptr, void* cookieName){
	static_cast<QNetworkCookie*>(ptr)->setName(*static_cast<QByteArray*>(cookieName));
}

void QNetworkCookie_SetPath(void* ptr, char* path){
	static_cast<QNetworkCookie*>(ptr)->setPath(QString(path));
}

void QNetworkCookie_SetSecure(void* ptr, int enable){
	static_cast<QNetworkCookie*>(ptr)->setSecure(enable != 0);
}

void QNetworkCookie_SetValue(void* ptr, void* value){
	static_cast<QNetworkCookie*>(ptr)->setValue(*static_cast<QByteArray*>(value));
}

void QNetworkCookie_Swap(void* ptr, void* other){
	static_cast<QNetworkCookie*>(ptr)->swap(*static_cast<QNetworkCookie*>(other));
}

void* QNetworkCookie_ToRawForm(void* ptr, int form){
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->toRawForm(static_cast<QNetworkCookie::RawForm>(form)));
}

void* QNetworkCookie_Value(void* ptr){
	return new QByteArray(static_cast<QNetworkCookie*>(ptr)->value());
}

void QNetworkCookie_DestroyQNetworkCookie(void* ptr){
	static_cast<QNetworkCookie*>(ptr)->~QNetworkCookie();
}

