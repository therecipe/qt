#include "qdbusserver.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QDBusServer>
#include "_cgo_export.h"

class MyQDBusServer: public QDBusServer {
public:
};

void* QDBusServer_NewQDBusServer2(void* parent){
	return new QDBusServer(static_cast<QObject*>(parent));
}

void* QDBusServer_NewQDBusServer(char* address, void* parent){
	return new QDBusServer(QString(address), static_cast<QObject*>(parent));
}

char* QDBusServer_Address(void* ptr){
	return static_cast<QDBusServer*>(ptr)->address().toUtf8().data();
}

int QDBusServer_IsAnonymousAuthenticationAllowed(void* ptr){
	return static_cast<QDBusServer*>(ptr)->isAnonymousAuthenticationAllowed();
}

int QDBusServer_IsConnected(void* ptr){
	return static_cast<QDBusServer*>(ptr)->isConnected();
}

void QDBusServer_SetAnonymousAuthenticationAllowed(void* ptr, int value){
	static_cast<QDBusServer*>(ptr)->setAnonymousAuthenticationAllowed(value != 0);
}

void QDBusServer_DestroyQDBusServer(void* ptr){
	static_cast<QDBusServer*>(ptr)->~QDBusServer();
}

