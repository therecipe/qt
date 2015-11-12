#include "qnetworkcookiejar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNetworkCookie>
#include <QNetworkCookieJar>
#include "_cgo_export.h"

class MyQNetworkCookieJar: public QNetworkCookieJar {
public:
};

void* QNetworkCookieJar_NewQNetworkCookieJar(void* parent){
	return new QNetworkCookieJar(static_cast<QObject*>(parent));
}

int QNetworkCookieJar_DeleteCookie(void* ptr, void* cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_InsertCookie(void* ptr, void* cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_UpdateCookie(void* ptr, void* cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

void QNetworkCookieJar_DestroyQNetworkCookieJar(void* ptr){
	static_cast<QNetworkCookieJar*>(ptr)->~QNetworkCookieJar();
}

