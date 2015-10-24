#include "qdbusconnection.h"
#include <QUrl>
#include <QModelIndex>
#include <QDBusMessage>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QDBusConnection>
#include "_cgo_export.h"

class MyQDBusConnection: public QDBusConnection {
public:
};

QtObjectPtr QDBusConnection_NewQDBusConnection2(QtObjectPtr other){
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

QtObjectPtr QDBusConnection_NewQDBusConnection(char* name){
	return new QDBusConnection(QString(name));
}

char* QDBusConnection_BaseService(QtObjectPtr ptr){
	return static_cast<QDBusConnection*>(ptr)->baseService().toUtf8().data();
}

int QDBusConnection_CallWithCallback(QtObjectPtr ptr, QtObjectPtr message, QtObjectPtr receiver, char* returnMethod, char* errorMethod, int timeout){
	return static_cast<QDBusConnection*>(ptr)->callWithCallback(*static_cast<QDBusMessage*>(message), static_cast<QObject*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod), timeout);
}

int QDBusConnection_Connect(QtObjectPtr ptr, char* service, char* path, char* interfa, char* name, QtObjectPtr receiver, char* slot){
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

int QDBusConnection_Connect2(QtObjectPtr ptr, char* service, char* path, char* interfa, char* name, char* signature, QtObjectPtr receiver, char* slot){
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

int QDBusConnection_Connect3(QtObjectPtr ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, QtObjectPtr receiver, char* slot){
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

int QDBusConnection_ConnectionCapabilities(QtObjectPtr ptr){
	return static_cast<QDBusConnection*>(ptr)->connectionCapabilities();
}

void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name){
	QDBusConnection::disconnectFromBus(QString(name));
}

void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name){
	QDBusConnection::disconnectFromPeer(QString(name));
}

QtObjectPtr QDBusConnection_Interface(QtObjectPtr ptr){
	return static_cast<QDBusConnection*>(ptr)->interface();
}

int QDBusConnection_IsConnected(QtObjectPtr ptr){
	return static_cast<QDBusConnection*>(ptr)->isConnected();
}

char* QDBusConnection_Name(QtObjectPtr ptr){
	return static_cast<QDBusConnection*>(ptr)->name().toUtf8().data();
}

QtObjectPtr QDBusConnection_ObjectRegisteredAt(QtObjectPtr ptr, char* path){
	return static_cast<QDBusConnection*>(ptr)->objectRegisteredAt(QString(path));
}

int QDBusConnection_RegisterObject(QtObjectPtr ptr, char* path, QtObjectPtr object, int options){
	return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
}

int QDBusConnection_RegisterObject2(QtObjectPtr ptr, char* path, char* interfa, QtObjectPtr object, int options){
	return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), QString(interfa), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
}

int QDBusConnection_RegisterService(QtObjectPtr ptr, char* serviceName){
	return static_cast<QDBusConnection*>(ptr)->registerService(QString(serviceName));
}

int QDBusConnection_Send(QtObjectPtr ptr, QtObjectPtr message){
	return static_cast<QDBusConnection*>(ptr)->send(*static_cast<QDBusMessage*>(message));
}

void QDBusConnection_UnregisterObject(QtObjectPtr ptr, char* path, int mode){
	static_cast<QDBusConnection*>(ptr)->unregisterObject(QString(path), static_cast<QDBusConnection::UnregisterMode>(mode));
}

int QDBusConnection_UnregisterService(QtObjectPtr ptr, char* serviceName){
	return static_cast<QDBusConnection*>(ptr)->unregisterService(QString(serviceName));
}

void QDBusConnection_DestroyQDBusConnection(QtObjectPtr ptr){
	static_cast<QDBusConnection*>(ptr)->~QDBusConnection();
}

