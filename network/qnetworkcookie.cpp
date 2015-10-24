#include "qnetworkcookie.h"
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QByteArray>
#include <QDateTime>
#include <QString>
#include <QVariant>
#include <QNetworkCookie>
#include "_cgo_export.h"

class MyQNetworkCookie: public QNetworkCookie {
public:
};

QtObjectPtr QNetworkCookie_NewQNetworkCookie(QtObjectPtr name, QtObjectPtr value){
	return new QNetworkCookie(*static_cast<QByteArray*>(name), *static_cast<QByteArray*>(value));
}

QtObjectPtr QNetworkCookie_NewQNetworkCookie2(QtObjectPtr other){
	return new QNetworkCookie(*static_cast<QNetworkCookie*>(other));
}

char* QNetworkCookie_Domain(QtObjectPtr ptr){
	return static_cast<QNetworkCookie*>(ptr)->domain().toUtf8().data();
}

int QNetworkCookie_HasSameIdentifier(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QNetworkCookie*>(ptr)->hasSameIdentifier(*static_cast<QNetworkCookie*>(other));
}

int QNetworkCookie_IsHttpOnly(QtObjectPtr ptr){
	return static_cast<QNetworkCookie*>(ptr)->isHttpOnly();
}

int QNetworkCookie_IsSecure(QtObjectPtr ptr){
	return static_cast<QNetworkCookie*>(ptr)->isSecure();
}

int QNetworkCookie_IsSessionCookie(QtObjectPtr ptr){
	return static_cast<QNetworkCookie*>(ptr)->isSessionCookie();
}

void QNetworkCookie_Normalize(QtObjectPtr ptr, char* url){
	static_cast<QNetworkCookie*>(ptr)->normalize(QUrl(QString(url)));
}

char* QNetworkCookie_Path(QtObjectPtr ptr){
	return static_cast<QNetworkCookie*>(ptr)->path().toUtf8().data();
}

void QNetworkCookie_SetDomain(QtObjectPtr ptr, char* domain){
	static_cast<QNetworkCookie*>(ptr)->setDomain(QString(domain));
}

void QNetworkCookie_SetExpirationDate(QtObjectPtr ptr, QtObjectPtr date){
	static_cast<QNetworkCookie*>(ptr)->setExpirationDate(*static_cast<QDateTime*>(date));
}

void QNetworkCookie_SetHttpOnly(QtObjectPtr ptr, int enable){
	static_cast<QNetworkCookie*>(ptr)->setHttpOnly(enable != 0);
}

void QNetworkCookie_SetName(QtObjectPtr ptr, QtObjectPtr cookieName){
	static_cast<QNetworkCookie*>(ptr)->setName(*static_cast<QByteArray*>(cookieName));
}

void QNetworkCookie_SetPath(QtObjectPtr ptr, char* path){
	static_cast<QNetworkCookie*>(ptr)->setPath(QString(path));
}

void QNetworkCookie_SetSecure(QtObjectPtr ptr, int enable){
	static_cast<QNetworkCookie*>(ptr)->setSecure(enable != 0);
}

void QNetworkCookie_SetValue(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QNetworkCookie*>(ptr)->setValue(*static_cast<QByteArray*>(value));
}

void QNetworkCookie_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkCookie*>(ptr)->swap(*static_cast<QNetworkCookie*>(other));
}

void QNetworkCookie_DestroyQNetworkCookie(QtObjectPtr ptr){
	static_cast<QNetworkCookie*>(ptr)->~QNetworkCookie();
}

