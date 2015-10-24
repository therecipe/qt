#include "qnetworkcookiejar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkCookie>
#include <QObject>
#include <QNetworkCookieJar>
#include "_cgo_export.h"

class MyQNetworkCookieJar: public QNetworkCookieJar {
public:
};

QtObjectPtr QNetworkCookieJar_NewQNetworkCookieJar(QtObjectPtr parent){
	return new QNetworkCookieJar(static_cast<QObject*>(parent));
}

int QNetworkCookieJar_DeleteCookie(QtObjectPtr ptr, QtObjectPtr cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->deleteCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_InsertCookie(QtObjectPtr ptr, QtObjectPtr cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->insertCookie(*static_cast<QNetworkCookie*>(cookie));
}

int QNetworkCookieJar_UpdateCookie(QtObjectPtr ptr, QtObjectPtr cookie){
	return static_cast<QNetworkCookieJar*>(ptr)->updateCookie(*static_cast<QNetworkCookie*>(cookie));
}

void QNetworkCookieJar_DestroyQNetworkCookieJar(QtObjectPtr ptr){
	static_cast<QNetworkCookieJar*>(ptr)->~QNetworkCookieJar();
}

