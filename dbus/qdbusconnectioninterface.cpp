#include "qdbusconnectioninterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusConnection>
#include <QObject>
#include <QDBusConnectionInterface>
#include "_cgo_export.h"

class MyQDBusConnectionInterface: public QDBusConnectionInterface {
public:
void Signal_ServiceRegistered(const QString & serviceName){callbackQDBusConnectionInterfaceServiceRegistered(this->objectName().toUtf8().data(), serviceName.toUtf8().data());};
void Signal_ServiceUnregistered(const QString & serviceName){callbackQDBusConnectionInterfaceServiceUnregistered(this->objectName().toUtf8().data(), serviceName.toUtf8().data());};
};

void QDBusConnectionInterface_ConnectServiceRegistered(void* ptr){
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));;
}

void QDBusConnectionInterface_DisconnectServiceRegistered(void* ptr){
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));;
}

void QDBusConnectionInterface_ConnectServiceUnregistered(void* ptr){
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));;
}

void QDBusConnectionInterface_DisconnectServiceUnregistered(void* ptr){
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));;
}

