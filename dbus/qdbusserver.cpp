#include "qdbusserver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDBusServer>
#include "_cgo_export.h"

class MyQDBusServer: public QDBusServer {
public:
};

QtObjectPtr QDBusServer_NewQDBusServer2(QtObjectPtr parent){
	return new QDBusServer(static_cast<QObject*>(parent));
}

QtObjectPtr QDBusServer_NewQDBusServer(char* address, QtObjectPtr parent){
	return new QDBusServer(QString(address), static_cast<QObject*>(parent));
}

char* QDBusServer_Address(QtObjectPtr ptr){
	return static_cast<QDBusServer*>(ptr)->address().toUtf8().data();
}

int QDBusServer_IsAnonymousAuthenticationAllowed(QtObjectPtr ptr){
	return static_cast<QDBusServer*>(ptr)->isAnonymousAuthenticationAllowed();
}

int QDBusServer_IsConnected(QtObjectPtr ptr){
	return static_cast<QDBusServer*>(ptr)->isConnected();
}

void QDBusServer_SetAnonymousAuthenticationAllowed(QtObjectPtr ptr, int value){
	static_cast<QDBusServer*>(ptr)->setAnonymousAuthenticationAllowed(value != 0);
}

void QDBusServer_DestroyQDBusServer(QtObjectPtr ptr){
	static_cast<QDBusServer*>(ptr)->~QDBusServer();
}

