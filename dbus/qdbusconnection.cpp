#include "qdbusconnection.h"
#include <QDBusMessage>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDBusConnection>
#include "_cgo_export.h"

class MyQDBusConnection: public QDBusConnection {
public:
};

void* QDBusConnection_NewQDBusConnection2(void* other){
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

void* QDBusConnection_NewQDBusConnection(char* name){
	return new QDBusConnection(QString(name));
}

char* QDBusConnection_BaseService(void* ptr){
	return static_cast<QDBusConnection*>(ptr)->baseService().toUtf8().data();
}

int QDBusConnection_CallWithCallback(void* ptr, void* message, void* receiver, char* returnMethod, char* errorMethod, int timeout){
	return static_cast<QDBusConnection*>(ptr)->callWithCallback(*static_cast<QDBusMessage*>(message), static_cast<QObject*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod), timeout);
}

int QDBusConnection_Connect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot){
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

int QDBusConnection_Connect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot){
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

int QDBusConnection_Connect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot){
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

int QDBusConnection_ConnectionCapabilities(void* ptr){
	return static_cast<QDBusConnection*>(ptr)->connectionCapabilities();
}

void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name){
	QDBusConnection::disconnectFromBus(QString(name));
}

void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name){
	QDBusConnection::disconnectFromPeer(QString(name));
}

void* QDBusConnection_Interface(void* ptr){
	return static_cast<QDBusConnection*>(ptr)->interface();
}

int QDBusConnection_IsConnected(void* ptr){
	return static_cast<QDBusConnection*>(ptr)->isConnected();
}

void* QDBusConnection_QDBusConnection_LocalMachineId(){
	return new QByteArray(QDBusConnection::localMachineId());
}

char* QDBusConnection_Name(void* ptr){
	return static_cast<QDBusConnection*>(ptr)->name().toUtf8().data();
}

void* QDBusConnection_ObjectRegisteredAt(void* ptr, char* path){
	return static_cast<QDBusConnection*>(ptr)->objectRegisteredAt(QString(path));
}

int QDBusConnection_RegisterObject(void* ptr, char* path, void* object, int options){
	return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
}

int QDBusConnection_RegisterObject2(void* ptr, char* path, char* interfa, void* object, int options){
	return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), QString(interfa), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
}

int QDBusConnection_RegisterService(void* ptr, char* serviceName){
	return static_cast<QDBusConnection*>(ptr)->registerService(QString(serviceName));
}

int QDBusConnection_Send(void* ptr, void* message){
	return static_cast<QDBusConnection*>(ptr)->send(*static_cast<QDBusMessage*>(message));
}

void QDBusConnection_UnregisterObject(void* ptr, char* path, int mode){
	static_cast<QDBusConnection*>(ptr)->unregisterObject(QString(path), static_cast<QDBusConnection::UnregisterMode>(mode));
}

int QDBusConnection_UnregisterService(void* ptr, char* serviceName){
	return static_cast<QDBusConnection*>(ptr)->unregisterService(QString(serviceName));
}

void QDBusConnection_DestroyQDBusConnection(void* ptr){
	static_cast<QDBusConnection*>(ptr)->~QDBusConnection();
}

