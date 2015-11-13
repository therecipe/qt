#include "qdbusunixfiledescriptor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusUnixFileDescriptor>
#include "_cgo_export.h"

class MyQDBusUnixFileDescriptor: public QDBusUnixFileDescriptor {
public:
};

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor(){
	return new QDBusUnixFileDescriptor();
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(void* other){
	return new QDBusUnixFileDescriptor(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor){
	return new QDBusUnixFileDescriptor(fileDescriptor);
}

int QDBusUnixFileDescriptor_FileDescriptor(void* ptr){
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->fileDescriptor();
}

int QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported(){
	return QDBusUnixFileDescriptor::isSupported();
}

int QDBusUnixFileDescriptor_IsValid(void* ptr){
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->isValid();
}

void QDBusUnixFileDescriptor_SetFileDescriptor(void* ptr, int fileDescriptor){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->setFileDescriptor(fileDescriptor);
}

void QDBusUnixFileDescriptor_Swap(void* ptr, void* other){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->swap(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(void* ptr){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->~QDBusUnixFileDescriptor();
}

#include "qdbusvirtualobject.h"
#include <QDBusConnection>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusMessage>
#include <QDBusVirtualObject>
#include "_cgo_export.h"

class MyQDBusVirtualObject: public QDBusVirtualObject {
public:
};

int QDBusVirtualObject_HandleMessage(void* ptr, void* message, void* connection){
	return static_cast<QDBusVirtualObject*>(ptr)->handleMessage(*static_cast<QDBusMessage*>(message), *static_cast<QDBusConnection*>(connection));
}

char* QDBusVirtualObject_Introspect(void* ptr, char* path){
	return static_cast<QDBusVirtualObject*>(ptr)->introspect(QString(path)).toUtf8().data();
}

void QDBusVirtualObject_DestroyQDBusVirtualObject(void* ptr){
	static_cast<QDBusVirtualObject*>(ptr)->~QDBusVirtualObject();
}

#include "qdbuserror.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusError>
#include "_cgo_export.h"

class MyQDBusError: public QDBusError {
public:
};

char* QDBusError_QDBusError_ErrorString(int error){
	return QDBusError::errorString(static_cast<QDBusError::ErrorType>(error)).toUtf8().data();
}

int QDBusError_IsValid(void* ptr){
	return static_cast<QDBusError*>(ptr)->isValid();
}

char* QDBusError_Message(void* ptr){
	return static_cast<QDBusError*>(ptr)->message().toUtf8().data();
}

char* QDBusError_Name(void* ptr){
	return static_cast<QDBusError*>(ptr)->name().toUtf8().data();
}

int QDBusError_Type(void* ptr){
	return static_cast<QDBusError*>(ptr)->type();
}

#include "qdbusmessage.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusMessage>
#include "_cgo_export.h"

class MyQDBusMessage: public QDBusMessage {
public:
};

void* QDBusMessage_NewQDBusMessage(){
	return new QDBusMessage();
}

void* QDBusMessage_NewQDBusMessage2(void* other){
	return new QDBusMessage(*static_cast<QDBusMessage*>(other));
}

int QDBusMessage_AutoStartService(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->autoStartService();
}

char* QDBusMessage_ErrorMessage(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->errorMessage().toUtf8().data();
}

char* QDBusMessage_ErrorName(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->errorName().toUtf8().data();
}

char* QDBusMessage_Interface(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->interface().toUtf8().data();
}

int QDBusMessage_IsDelayedReply(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->isDelayedReply();
}

int QDBusMessage_IsReplyRequired(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->isReplyRequired();
}

char* QDBusMessage_Member(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->member().toUtf8().data();
}

char* QDBusMessage_Path(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->path().toUtf8().data();
}

char* QDBusMessage_Service(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->service().toUtf8().data();
}

void QDBusMessage_SetAutoStartService(void* ptr, int enable){
	static_cast<QDBusMessage*>(ptr)->setAutoStartService(enable != 0);
}

void QDBusMessage_SetDelayedReply(void* ptr, int enable){
	static_cast<QDBusMessage*>(ptr)->setDelayedReply(enable != 0);
}

char* QDBusMessage_Signature(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->signature().toUtf8().data();
}

int QDBusMessage_Type(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->type();
}

void QDBusMessage_DestroyQDBusMessage(void* ptr){
	static_cast<QDBusMessage*>(ptr)->~QDBusMessage();
}

#include "qdbusvariant.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusVariant>
#include "_cgo_export.h"

class MyQDBusVariant: public QDBusVariant {
public:
};

void* QDBusVariant_NewQDBusVariant(){
	return new QDBusVariant();
}

void* QDBusVariant_NewQDBusVariant2(void* variant){
	return new QDBusVariant(*static_cast<QVariant*>(variant));
}

void QDBusVariant_SetVariant(void* ptr, void* variant){
	static_cast<QDBusVariant*>(ptr)->setVariant(*static_cast<QVariant*>(variant));
}

void* QDBusVariant_Variant(void* ptr){
	return new QVariant(static_cast<QDBusVariant*>(ptr)->variant());
}

#include "qdbusargument.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusArgument>
#include "_cgo_export.h"

class MyQDBusArgument: public QDBusArgument {
public:
};

void* QDBusArgument_NewQDBusArgument(){
	return new QDBusArgument();
}

void* QDBusArgument_NewQDBusArgument2(void* other){
	return new QDBusArgument(*static_cast<QDBusArgument*>(other));
}

void* QDBusArgument_AsVariant(void* ptr){
	return new QVariant(static_cast<QDBusArgument*>(ptr)->asVariant());
}

int QDBusArgument_AtEnd(void* ptr){
	return static_cast<QDBusArgument*>(ptr)->atEnd();
}

void QDBusArgument_BeginArray(void* ptr, int id){
	static_cast<QDBusArgument*>(ptr)->beginArray(id);
}

void QDBusArgument_BeginArray2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->beginArray();
}

void QDBusArgument_BeginMap(void* ptr, int kid, int vid){
	static_cast<QDBusArgument*>(ptr)->beginMap(kid, vid);
}

void QDBusArgument_BeginMap2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->beginMap();
}

void QDBusArgument_BeginMapEntry(void* ptr){
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginMapEntry2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginStructure(void* ptr){
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

void QDBusArgument_BeginStructure2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

int QDBusArgument_CurrentType(void* ptr){
	return static_cast<QDBusArgument*>(ptr)->currentType();
}

void QDBusArgument_EndArray(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndArray2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndMap(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMap2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMapEntry(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndMapEntry2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndStructure(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_EndStructure2(void* ptr){
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_DestroyQDBusArgument(void* ptr){
	static_cast<QDBusArgument*>(ptr)->~QDBusArgument();
}

#include "qdbusobjectpath.h"
#include <QLatin1String>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusObjectPath>
#include "_cgo_export.h"

class MyQDBusObjectPath: public QDBusObjectPath {
public:
};

void* QDBusObjectPath_NewQDBusObjectPath(){
	return new QDBusObjectPath();
}

void* QDBusObjectPath_NewQDBusObjectPath3(void* path){
	return new QDBusObjectPath(*static_cast<QLatin1String*>(path));
}

void* QDBusObjectPath_NewQDBusObjectPath4(char* path){
	return new QDBusObjectPath(QString(path));
}

void* QDBusObjectPath_NewQDBusObjectPath2(char* path){
	return new QDBusObjectPath(const_cast<const char*>(path));
}

char* QDBusObjectPath_Path(void* ptr){
	return static_cast<QDBusObjectPath*>(ptr)->path().toUtf8().data();
}

void QDBusObjectPath_SetPath(void* ptr, char* path){
	static_cast<QDBusObjectPath*>(ptr)->setPath(QString(path));
}

#include "qdbusconnection.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QDBusMessage>
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

#include "qdbussignature.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QString>
#include <QDBusSignature>
#include "_cgo_export.h"

class MyQDBusSignature: public QDBusSignature {
public:
};

void* QDBusSignature_NewQDBusSignature(){
	return new QDBusSignature();
}

void* QDBusSignature_NewQDBusSignature3(void* signature){
	return new QDBusSignature(*static_cast<QLatin1String*>(signature));
}

void* QDBusSignature_NewQDBusSignature4(char* signature){
	return new QDBusSignature(QString(signature));
}

void* QDBusSignature_NewQDBusSignature2(char* signature){
	return new QDBusSignature(const_cast<const char*>(signature));
}

void QDBusSignature_SetSignature(void* ptr, char* signature){
	static_cast<QDBusSignature*>(ptr)->setSignature(QString(signature));
}

char* QDBusSignature_Signature(void* ptr){
	return static_cast<QDBusSignature*>(ptr)->signature().toUtf8().data();
}

#include "qdbuspendingcall.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QDBusPendingCall>
#include "_cgo_export.h"

class MyQDBusPendingCall: public QDBusPendingCall {
public:
};

void* QDBusPendingCall_NewQDBusPendingCall(void* other){
	return new QDBusPendingCall(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_Swap(void* ptr, void* other){
	static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr){
	static_cast<QDBusPendingCall*>(ptr)->~QDBusPendingCall();
}

#include "qdbusservicewatcher.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDBusConnection>
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

#include "qdbusabstractadaptor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusAbstractAdaptor>
#include "_cgo_export.h"

class MyQDBusAbstractAdaptor: public QDBusAbstractAdaptor {
public:
};

void QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(void* ptr){
	static_cast<QDBusAbstractAdaptor*>(ptr)->~QDBusAbstractAdaptor();
}

#include "qdbusreply.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include "_cgo_export.h"

#include "qdbuscontext.h"
#include <QUrl>
#include <QModelIndex>
#include <QDBusError>
#include <QString>
#include <QVariant>
#include <QDBusContext>
#include "_cgo_export.h"

class MyQDBusContext: public QDBusContext {
public:
};

void* QDBusContext_NewQDBusContext(){
	return new QDBusContext();
}

int QDBusContext_CalledFromDBus(void* ptr){
	return static_cast<QDBusContext*>(ptr)->calledFromDBus();
}

int QDBusContext_IsDelayedReply(void* ptr){
	return static_cast<QDBusContext*>(ptr)->isDelayedReply();
}

void QDBusContext_SendErrorReply2(void* ptr, int ty, char* msg){
	static_cast<QDBusContext*>(ptr)->sendErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg));
}

void QDBusContext_SendErrorReply(void* ptr, char* name, char* msg){
	static_cast<QDBusContext*>(ptr)->sendErrorReply(QString(name), QString(msg));
}

void QDBusContext_SetDelayedReply(void* ptr, int enable){
	static_cast<QDBusContext*>(ptr)->setDelayedReply(enable != 0);
}

void QDBusContext_DestroyQDBusContext(void* ptr){
	static_cast<QDBusContext*>(ptr)->~QDBusContext();
}

#include "qdbusinterface.h"
#include <QModelIndex>
#include <QDBusConnection>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusInterface>
#include "_cgo_export.h"

class MyQDBusInterface: public QDBusInterface {
public:
};

void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent){
	return new QDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QObject*>(parent));
}

void QDBusInterface_DestroyQDBusInterface(void* ptr){
	static_cast<QDBusInterface*>(ptr)->~QDBusInterface();
}

#include "qdbuspendingcallwatcher.h"
#include <QUrl>
#include <QModelIndex>
#include <QDBusPendingCall>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QDBusPendingCallWatcher>
#include "_cgo_export.h"

class MyQDBusPendingCallWatcher: public QDBusPendingCallWatcher {
public:
void Signal_Finished(QDBusPendingCallWatcher * self){callbackQDBusPendingCallWatcherFinished(this->objectName().toUtf8().data(), self);};
};

void QDBusPendingCallWatcher_WaitForFinished(void* ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->waitForFinished();
}

void* QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(void* call, void* parent){
	return new QDBusPendingCallWatcher(*static_cast<QDBusPendingCall*>(call), static_cast<QObject*>(parent));
}

void QDBusPendingCallWatcher_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));;
}

void QDBusPendingCallWatcher_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));;
}

int QDBusPendingCallWatcher_IsFinished(void* ptr){
	return static_cast<QDBusPendingCallWatcher*>(ptr)->isFinished();
}

void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->~QDBusPendingCallWatcher();
}

#include "qdbuspendingreply.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include "_cgo_export.h"

#include "qdbusabstractinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusAbstractInterface>
#include "_cgo_export.h"

class MyQDBusAbstractInterface: public QDBusAbstractInterface {
public:
};

char* QDBusAbstractInterface_Interface(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->interface().toUtf8().data();
}

int QDBusAbstractInterface_IsValid(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->isValid();
}

char* QDBusAbstractInterface_Path(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->path().toUtf8().data();
}

char* QDBusAbstractInterface_Service(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->service().toUtf8().data();
}

void QDBusAbstractInterface_SetTimeout(void* ptr, int timeout){
	static_cast<QDBusAbstractInterface*>(ptr)->setTimeout(timeout);
}

int QDBusAbstractInterface_Timeout(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->timeout();
}

void QDBusAbstractInterface_DestroyQDBusAbstractInterface(void* ptr){
	static_cast<QDBusAbstractInterface*>(ptr)->~QDBusAbstractInterface();
}

#include "qdbusconnectioninterface.h"
#include <QDBusConnection>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

#include "qdbusserver.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
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

