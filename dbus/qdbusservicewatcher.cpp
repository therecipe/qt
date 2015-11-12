#include "qdbusservicewatcher.h"
#include <QModelIndex>
#include <QDBusConnection>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusServiceWatcher>
#include "_cgo_export.h"

class MyQDBusServiceWatcher: public QDBusServiceWatcher {
public:
void Signal_ServiceOwnerChanged(const QString & serviceName, const QString & oldOwner, const QString & newOwner){callbackQDBusServiceWatcherServiceOwnerChanged(this->objectName().toUtf8().data(), serviceName.toUtf8().data(), oldOwner.toUtf8().data(), newOwner.toUtf8().data());};
void Signal_ServiceRegistered(const QString & serviceName){callbackQDBusServiceWatcherServiceRegistered(this->objectName().toUtf8().data(), serviceName.toUtf8().data());};
void Signal_ServiceUnregistered(const QString & serviceName){callbackQDBusServiceWatcherServiceUnregistered(this->objectName().toUtf8().data(), serviceName.toUtf8().data());};
};

void QDBusServiceWatcher_SetWatchMode(void* ptr, int mode){
	static_cast<QDBusServiceWatcher*>(ptr)->setWatchMode(static_cast<QDBusServiceWatcher::WatchModeFlag>(mode));
}

int QDBusServiceWatcher_WatchMode(void* ptr){
	return static_cast<QDBusServiceWatcher*>(ptr)->watchMode();
}

void* QDBusServiceWatcher_NewQDBusServiceWatcher(void* parent){
	return new QDBusServiceWatcher(static_cast<QObject*>(parent));
}

void* QDBusServiceWatcher_NewQDBusServiceWatcher2(char* service, void* connection, int watchMode, void* parent){
	return new QDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QObject*>(parent));
}

void QDBusServiceWatcher_AddWatchedService(void* ptr, char* newService){
	static_cast<QDBusServiceWatcher*>(ptr)->addWatchedService(QString(newService));
}

int QDBusServiceWatcher_RemoveWatchedService(void* ptr, char* service){
	return static_cast<QDBusServiceWatcher*>(ptr)->removeWatchedService(QString(service));
}

void QDBusServiceWatcher_ConnectServiceOwnerChanged(void* ptr){
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&QDBusServiceWatcher::serviceOwnerChanged), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceOwnerChanged));;
}

void QDBusServiceWatcher_DisconnectServiceOwnerChanged(void* ptr){
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&QDBusServiceWatcher::serviceOwnerChanged), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceOwnerChanged));;
}

void QDBusServiceWatcher_ConnectServiceRegistered(void* ptr){
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));;
}

void QDBusServiceWatcher_DisconnectServiceRegistered(void* ptr){
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));;
}

void QDBusServiceWatcher_ConnectServiceUnregistered(void* ptr){
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));;
}

void QDBusServiceWatcher_DisconnectServiceUnregistered(void* ptr){
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));;
}

void QDBusServiceWatcher_SetConnection(void* ptr, void* connection){
	static_cast<QDBusServiceWatcher*>(ptr)->setConnection(*static_cast<QDBusConnection*>(connection));
}

void QDBusServiceWatcher_SetWatchedServices(void* ptr, char* services){
	static_cast<QDBusServiceWatcher*>(ptr)->setWatchedServices(QString(services).split("|", QString::SkipEmptyParts));
}

char* QDBusServiceWatcher_WatchedServices(void* ptr){
	return static_cast<QDBusServiceWatcher*>(ptr)->watchedServices().join("|").toUtf8().data();
}

void QDBusServiceWatcher_DestroyQDBusServiceWatcher(void* ptr){
	static_cast<QDBusServiceWatcher*>(ptr)->~QDBusServiceWatcher();
}

