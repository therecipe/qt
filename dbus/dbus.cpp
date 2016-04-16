#define protected public

#include "dbus.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QDBusAbstractAdaptor>
#include <QDBusAbstractInterface>
#include <QDBusArgument>
#include <QDBusConnection>
#include <QDBusConnectionInterface>
#include <QDBusContext>
#include <QDBusError>
#include <QDBusInterface>
#include <QDBusMessage>
#include <QDBusObjectPath>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDBusServer>
#include <QDBusServiceWatcher>
#include <QDBusSignature>
#include <QDBusUnixFileDescriptor>
#include <QDBusVariant>
#include <QDBusVirtualObject>
#include <QEvent>
#include <QLatin1String>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>

void QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(void* ptr){
	static_cast<QDBusAbstractAdaptor*>(ptr)->~QDBusAbstractAdaptor();
}

void QDBusAbstractAdaptor_TimerEvent(void* ptr, void* event){
	static_cast<QDBusAbstractAdaptor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractAdaptor_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractAdaptor_ChildEvent(void* ptr, void* event){
	static_cast<QDBusAbstractAdaptor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractAdaptor_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractAdaptor_CustomEvent(void* ptr, void* event){
	static_cast<QDBusAbstractAdaptor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractAdaptor_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::customEvent(static_cast<QEvent*>(event));
}

class MyQDBusAbstractInterface: public QDBusAbstractInterface {
public:
	void timerEvent(QTimerEvent * event) { callbackQDBusAbstractInterfaceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDBusAbstractInterfaceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDBusAbstractInterfaceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QDBusAbstractInterface_AsyncCall(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8){
	return new QDBusPendingCall(static_cast<QDBusAbstractInterface*>(ptr)->asyncCall(QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_Call(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8){
	return new QDBusMessage(static_cast<QDBusAbstractInterface*>(ptr)->call(QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_Connection(void* ptr){
	return new QDBusConnection(static_cast<QDBusAbstractInterface*>(ptr)->connection());
}

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

void QDBusAbstractInterface_TimerEvent(void* ptr, void* event){
	static_cast<MyQDBusAbstractInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractInterface_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractInterface_ChildEvent(void* ptr, void* event){
	static_cast<MyQDBusAbstractInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractInterface_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractInterface_CustomEvent(void* ptr, void* event){
	static_cast<MyQDBusAbstractInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractInterface_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::customEvent(static_cast<QEvent*>(event));
}

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

void* QDBusConnection_QDBusConnection_SessionBus(){
	return new QDBusConnection(QDBusConnection::sessionBus());
}

void* QDBusConnection_QDBusConnection_SystemBus(){
	return new QDBusConnection(QDBusConnection::systemBus());
}

void* QDBusConnection_NewQDBusConnection2(void* other){
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

void* QDBusConnection_NewQDBusConnection(char* name){
	return new QDBusConnection(QString(name));
}

char* QDBusConnection_BaseService(void* ptr){
	return static_cast<QDBusConnection*>(ptr)->baseService().toUtf8().data();
}

void* QDBusConnection_QDBusConnection_ConnectToBus(int ty, char* name){
	return new QDBusConnection(QDBusConnection::connectToBus(static_cast<QDBusConnection::BusType>(ty), QString(name)));
}

void* QDBusConnection_QDBusConnection_ConnectToBus2(char* address, char* name){
	return new QDBusConnection(QDBusConnection::connectToBus(QString(address), QString(name)));
}

void* QDBusConnection_AsyncCall(void* ptr, void* message, int timeout){
	return new QDBusPendingCall(static_cast<QDBusConnection*>(ptr)->asyncCall(*static_cast<QDBusMessage*>(message), timeout));
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

void* QDBusConnection_QDBusConnection_ConnectToPeer(char* address, char* name){
	return new QDBusConnection(QDBusConnection::connectToPeer(QString(address), QString(name)));
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

char* QDBusConnection_QDBusConnection_LocalMachineId(){
	return QString(QDBusConnection::localMachineId()).toUtf8().data();
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

class MyQDBusConnectionInterface: public QDBusConnectionInterface {
public:
	void Signal_ServiceRegistered(const QString & serviceName) { callbackQDBusConnectionInterfaceServiceRegistered(this, this->objectName().toUtf8().data(), serviceName.toUtf8().data()); };
	void Signal_ServiceUnregistered(const QString & serviceName) { callbackQDBusConnectionInterfaceServiceUnregistered(this, this->objectName().toUtf8().data(), serviceName.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQDBusConnectionInterfaceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDBusConnectionInterfaceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDBusConnectionInterfaceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QDBusConnectionInterface_ConnectServiceRegistered(void* ptr){
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));;
}

void QDBusConnectionInterface_DisconnectServiceRegistered(void* ptr){
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));;
}

void QDBusConnectionInterface_ServiceRegistered(void* ptr, char* serviceName){
	static_cast<QDBusConnectionInterface*>(ptr)->serviceRegistered(QString(serviceName));
}

void QDBusConnectionInterface_ConnectServiceUnregistered(void* ptr){
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));;
}

void QDBusConnectionInterface_DisconnectServiceUnregistered(void* ptr){
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));;
}

void QDBusConnectionInterface_ServiceUnregistered(void* ptr, char* serviceName){
	static_cast<QDBusConnectionInterface*>(ptr)->serviceUnregistered(QString(serviceName));
}

void QDBusConnectionInterface_TimerEvent(void* ptr, void* event){
	static_cast<MyQDBusConnectionInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusConnectionInterface_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusConnectionInterface_ChildEvent(void* ptr, void* event){
	static_cast<MyQDBusConnectionInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusConnectionInterface_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusConnectionInterface_CustomEvent(void* ptr, void* event){
	static_cast<MyQDBusConnectionInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusConnectionInterface_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::customEvent(static_cast<QEvent*>(event));
}

void* QDBusContext_NewQDBusContext(){
	return new QDBusContext();
}

int QDBusContext_CalledFromDBus(void* ptr){
	return static_cast<QDBusContext*>(ptr)->calledFromDBus();
}

void* QDBusContext_Connection(void* ptr){
	return new QDBusConnection(static_cast<QDBusContext*>(ptr)->connection());
}

int QDBusContext_IsDelayedReply(void* ptr){
	return static_cast<QDBusContext*>(ptr)->isDelayedReply();
}

void* QDBusContext_Message(void* ptr){
	return new QDBusMessage(static_cast<QDBusContext*>(ptr)->message());
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

void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent){
	return new QDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QObject*>(parent));
}

void QDBusInterface_DestroyQDBusInterface(void* ptr){
	static_cast<QDBusInterface*>(ptr)->~QDBusInterface();
}

void QDBusInterface_TimerEvent(void* ptr, void* event){
	static_cast<QDBusInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusInterface_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusInterface_ChildEvent(void* ptr, void* event){
	static_cast<QDBusInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusInterface_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusInterface_CustomEvent(void* ptr, void* event){
	static_cast<QDBusInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusInterface_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::customEvent(static_cast<QEvent*>(event));
}

void* QDBusMessage_CreateErrorReply3(void* ptr, int ty, char* msg){
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg)));
}

void* QDBusMessage_NewQDBusMessage(){
	return new QDBusMessage();
}

void* QDBusMessage_NewQDBusMessage2(void* other){
	return new QDBusMessage(*static_cast<QDBusMessage*>(other));
}

int QDBusMessage_AutoStartService(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->autoStartService();
}

void* QDBusMessage_QDBusMessage_CreateError3(int ty, char* msg){
	return new QDBusMessage(QDBusMessage::createError(static_cast<QDBusError::ErrorType>(ty), QString(msg)));
}

void* QDBusMessage_QDBusMessage_CreateError2(void* error){
	return new QDBusMessage(QDBusMessage::createError(*static_cast<QDBusError*>(error)));
}

void* QDBusMessage_QDBusMessage_CreateError(char* name, char* msg){
	return new QDBusMessage(QDBusMessage::createError(QString(name), QString(msg)));
}

void* QDBusMessage_CreateErrorReply2(void* ptr, void* error){
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(*static_cast<QDBusError*>(error)));
}

void* QDBusMessage_CreateErrorReply(void* ptr, char* name, char* msg){
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(QString(name), QString(msg)));
}

void* QDBusMessage_QDBusMessage_CreateMethodCall(char* service, char* path, char* interfa, char* method){
	return new QDBusMessage(QDBusMessage::createMethodCall(QString(service), QString(path), QString(interfa), QString(method)));
}

void* QDBusMessage_CreateReply2(void* ptr, void* argument){
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createReply(*static_cast<QVariant*>(argument)));
}

void* QDBusMessage_QDBusMessage_CreateSignal(char* path, char* interfa, char* name){
	return new QDBusMessage(QDBusMessage::createSignal(QString(path), QString(interfa), QString(name)));
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

void* QDBusPendingCall_NewQDBusPendingCall(void* other){
	return new QDBusPendingCall(*static_cast<QDBusPendingCall*>(other));
}

void* QDBusPendingCall_QDBusPendingCall_FromCompletedCall(void* msg){
	return new QDBusPendingCall(QDBusPendingCall::fromCompletedCall(*static_cast<QDBusMessage*>(msg)));
}

void* QDBusPendingCall_QDBusPendingCall_FromError(void* error){
	return new QDBusPendingCall(QDBusPendingCall::fromError(*static_cast<QDBusError*>(error)));
}

void QDBusPendingCall_Swap(void* ptr, void* other){
	static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr){
	static_cast<QDBusPendingCall*>(ptr)->~QDBusPendingCall();
}

class MyQDBusPendingCallWatcher: public QDBusPendingCallWatcher {
public:
	void Signal_Finished(QDBusPendingCallWatcher * self) { callbackQDBusPendingCallWatcherFinished(this, this->objectName().toUtf8().data(), self); };
	void timerEvent(QTimerEvent * event) { callbackQDBusPendingCallWatcherTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDBusPendingCallWatcherChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDBusPendingCallWatcherCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QDBusPendingCallWatcher_Finished(void* ptr, void* self){
	static_cast<QDBusPendingCallWatcher*>(ptr)->finished(static_cast<QDBusPendingCallWatcher*>(self));
}

int QDBusPendingCallWatcher_IsFinished(void* ptr){
	return static_cast<QDBusPendingCallWatcher*>(ptr)->isFinished();
}

void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->~QDBusPendingCallWatcher();
}

void QDBusPendingCallWatcher_TimerEvent(void* ptr, void* event){
	static_cast<MyQDBusPendingCallWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusPendingCallWatcher_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusPendingCallWatcher_ChildEvent(void* ptr, void* event){
	static_cast<MyQDBusPendingCallWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusPendingCallWatcher_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusPendingCallWatcher_CustomEvent(void* ptr, void* event){
	static_cast<MyQDBusPendingCallWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusPendingCallWatcher_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::customEvent(static_cast<QEvent*>(event));
}

class MyQDBusServer: public QDBusServer {
public:
	MyQDBusServer(QObject *parent) : QDBusServer(parent) {};
	MyQDBusServer(const QString &address, QObject *parent) : QDBusServer(address, parent) {};
	void Signal_NewConnection(const QDBusConnection & connection) { callbackQDBusServerNewConnection(this, this->objectName().toUtf8().data(), new QDBusConnection(connection)); };
	void timerEvent(QTimerEvent * event) { callbackQDBusServerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDBusServerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDBusServerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QDBusServer_NewQDBusServer2(void* parent){
	return new MyQDBusServer(static_cast<QObject*>(parent));
}

void* QDBusServer_NewQDBusServer(char* address, void* parent){
	return new MyQDBusServer(QString(address), static_cast<QObject*>(parent));
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

void QDBusServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QDBusServer*>(ptr), static_cast<void (QDBusServer::*)(const QDBusConnection &)>(&QDBusServer::newConnection), static_cast<MyQDBusServer*>(ptr), static_cast<void (MyQDBusServer::*)(const QDBusConnection &)>(&MyQDBusServer::Signal_NewConnection));;
}

void QDBusServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QDBusServer*>(ptr), static_cast<void (QDBusServer::*)(const QDBusConnection &)>(&QDBusServer::newConnection), static_cast<MyQDBusServer*>(ptr), static_cast<void (MyQDBusServer::*)(const QDBusConnection &)>(&MyQDBusServer::Signal_NewConnection));;
}

void QDBusServer_NewConnection(void* ptr, void* connection){
	static_cast<QDBusServer*>(ptr)->newConnection(*static_cast<QDBusConnection*>(connection));
}

void QDBusServer_SetAnonymousAuthenticationAllowed(void* ptr, int value){
	static_cast<QDBusServer*>(ptr)->setAnonymousAuthenticationAllowed(value != 0);
}

void QDBusServer_DestroyQDBusServer(void* ptr){
	static_cast<QDBusServer*>(ptr)->~QDBusServer();
}

void QDBusServer_TimerEvent(void* ptr, void* event){
	static_cast<MyQDBusServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServer_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusServer*>(ptr)->QDBusServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServer_ChildEvent(void* ptr, void* event){
	static_cast<MyQDBusServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServer_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusServer*>(ptr)->QDBusServer::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServer_CustomEvent(void* ptr, void* event){
	static_cast<MyQDBusServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusServer_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusServer*>(ptr)->QDBusServer::customEvent(static_cast<QEvent*>(event));
}

class MyQDBusServiceWatcher: public QDBusServiceWatcher {
public:
	void Signal_ServiceOwnerChanged(const QString & serviceName, const QString & oldOwner, const QString & newOwner) { callbackQDBusServiceWatcherServiceOwnerChanged(this, this->objectName().toUtf8().data(), serviceName.toUtf8().data(), oldOwner.toUtf8().data(), newOwner.toUtf8().data()); };
	void Signal_ServiceRegistered(const QString & serviceName) { callbackQDBusServiceWatcherServiceRegistered(this, this->objectName().toUtf8().data(), serviceName.toUtf8().data()); };
	void Signal_ServiceUnregistered(const QString & serviceName) { callbackQDBusServiceWatcherServiceUnregistered(this, this->objectName().toUtf8().data(), serviceName.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQDBusServiceWatcherTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDBusServiceWatcherChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDBusServiceWatcherCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void* QDBusServiceWatcher_Connection(void* ptr){
	return new QDBusConnection(static_cast<QDBusServiceWatcher*>(ptr)->connection());
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

void QDBusServiceWatcher_ServiceOwnerChanged(void* ptr, char* serviceName, char* oldOwner, char* newOwner){
	static_cast<QDBusServiceWatcher*>(ptr)->serviceOwnerChanged(QString(serviceName), QString(oldOwner), QString(newOwner));
}

void QDBusServiceWatcher_ConnectServiceRegistered(void* ptr){
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));;
}

void QDBusServiceWatcher_DisconnectServiceRegistered(void* ptr){
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));;
}

void QDBusServiceWatcher_ServiceRegistered(void* ptr, char* serviceName){
	static_cast<QDBusServiceWatcher*>(ptr)->serviceRegistered(QString(serviceName));
}

void QDBusServiceWatcher_ConnectServiceUnregistered(void* ptr){
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));;
}

void QDBusServiceWatcher_DisconnectServiceUnregistered(void* ptr){
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));;
}

void QDBusServiceWatcher_ServiceUnregistered(void* ptr, char* serviceName){
	static_cast<QDBusServiceWatcher*>(ptr)->serviceUnregistered(QString(serviceName));
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

void QDBusServiceWatcher_TimerEvent(void* ptr, void* event){
	static_cast<MyQDBusServiceWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServiceWatcher_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServiceWatcher_ChildEvent(void* ptr, void* event){
	static_cast<MyQDBusServiceWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServiceWatcher_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServiceWatcher_CustomEvent(void* ptr, void* event){
	static_cast<MyQDBusServiceWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusServiceWatcher_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::customEvent(static_cast<QEvent*>(event));
}

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

class MyQDBusVirtualObject: public QDBusVirtualObject {
public:
	void timerEvent(QTimerEvent * event) { callbackQDBusVirtualObjectTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDBusVirtualObjectChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQDBusVirtualObjectCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QDBusVirtualObject_TimerEvent(void* ptr, void* event){
	static_cast<MyQDBusVirtualObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusVirtualObject_TimerEventDefault(void* ptr, void* event){
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusVirtualObject_ChildEvent(void* ptr, void* event){
	static_cast<MyQDBusVirtualObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusVirtualObject_ChildEventDefault(void* ptr, void* event){
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusVirtualObject_CustomEvent(void* ptr, void* event){
	static_cast<MyQDBusVirtualObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusVirtualObject_CustomEventDefault(void* ptr, void* event){
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::customEvent(static_cast<QEvent*>(event));
}

