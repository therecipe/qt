// +build !minimal

#define protected public
#define private public

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
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>

void* QDBusAbstractAdaptor_NewQDBusAbstractAdaptor(void* obj)
{
	return new QDBusAbstractAdaptor(static_cast<QObject*>(obj));
}

char QDBusAbstractAdaptor_AutoRelaySignals(void* ptr)
{
	return static_cast<QDBusAbstractAdaptor*>(ptr)->autoRelaySignals();
}

void QDBusAbstractAdaptor_SetAutoRelaySignals(void* ptr, char enable)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->setAutoRelaySignals(enable != 0);
}

void QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(void* ptr)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->~QDBusAbstractAdaptor();
}

void QDBusAbstractAdaptor_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractAdaptor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractAdaptor_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractAdaptor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractAdaptor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractAdaptor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractAdaptor_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractAdaptor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractAdaptor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusAbstractAdaptor*>(ptr), "deleteLater");
}

void QDBusAbstractAdaptor_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::deleteLater();
}

void QDBusAbstractAdaptor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractAdaptor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusAbstractAdaptor_Event(void* ptr, void* e)
{
	return static_cast<QDBusAbstractAdaptor*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusAbstractAdaptor_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::event(static_cast<QEvent*>(e));
}

char QDBusAbstractAdaptor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusAbstractAdaptor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusAbstractAdaptor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusAbstractAdaptor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusAbstractAdaptor*>(ptr)->metaObject());
}

void* QDBusAbstractAdaptor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::metaObject());
}

class MyQDBusAbstractInterface: public QDBusAbstractInterface
{
public:
	void timerEvent(QTimerEvent * event) { callbackQDBusAbstractInterface_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDBusAbstractInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusAbstractInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusAbstractInterface_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDBusAbstractInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusAbstractInterface_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusAbstractInterface_MetaObject(const_cast<MyQDBusAbstractInterface*>(this))); };
};

void* QDBusAbstractInterface_AsyncCall(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8)
{
	return new QDBusPendingCall(static_cast<QDBusAbstractInterface*>(ptr)->asyncCall(QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_Call(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8)
{
	return new QDBusMessage(static_cast<QDBusAbstractInterface*>(ptr)->call(QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_Connection(void* ptr)
{
	return new QDBusConnection(static_cast<QDBusAbstractInterface*>(ptr)->connection());
}

char* QDBusAbstractInterface_Interface(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusAbstractInterface*>(ptr)->interface().toUtf8().constData());
}

char QDBusAbstractInterface_IsValid(void* ptr)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->isValid();
}

void* QDBusAbstractInterface_LastError(void* ptr)
{
	return new QDBusError(static_cast<QDBusAbstractInterface*>(ptr)->lastError());
}

char* QDBusAbstractInterface_Path(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusAbstractInterface*>(ptr)->path().toUtf8().constData());
}

char* QDBusAbstractInterface_Service(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusAbstractInterface*>(ptr)->service().toUtf8().constData());
}

void QDBusAbstractInterface_SetTimeout(void* ptr, int timeout)
{
	static_cast<QDBusAbstractInterface*>(ptr)->setTimeout(timeout);
}

int QDBusAbstractInterface_Timeout(void* ptr)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->timeout();
}

void QDBusAbstractInterface_DestroyQDBusAbstractInterface(void* ptr)
{
	static_cast<QDBusAbstractInterface*>(ptr)->~QDBusAbstractInterface();
}

void QDBusAbstractInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusAbstractInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusAbstractInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusAbstractInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusAbstractInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusAbstractInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusAbstractInterface*>(ptr), "deleteLater");
}

void QDBusAbstractInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::deleteLater();
}

void QDBusAbstractInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusAbstractInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusAbstractInterface_Event(void* ptr, void* e)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusAbstractInterface_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::event(static_cast<QEvent*>(e));
}

char QDBusAbstractInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusAbstractInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusAbstractInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusAbstractInterface*>(ptr)->metaObject());
}

void* QDBusAbstractInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::metaObject());
}

void* QDBusArgument_NewQDBusArgument()
{
	return new QDBusArgument();
}

void* QDBusArgument_NewQDBusArgument3(void* other)
{
	return new QDBusArgument(*static_cast<QDBusArgument*>(other));
}

void* QDBusArgument_NewQDBusArgument2(void* other)
{
	return new QDBusArgument(*static_cast<QDBusArgument*>(other));
}

void* QDBusArgument_AsVariant(void* ptr)
{
	return new QVariant(static_cast<QDBusArgument*>(ptr)->asVariant());
}

char QDBusArgument_AtEnd(void* ptr)
{
	return static_cast<QDBusArgument*>(ptr)->atEnd();
}

void QDBusArgument_BeginArray(void* ptr, int id)
{
	static_cast<QDBusArgument*>(ptr)->beginArray(id);
}

void QDBusArgument_BeginArray2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginArray();
}

void QDBusArgument_BeginMap(void* ptr, int kid, int vid)
{
	static_cast<QDBusArgument*>(ptr)->beginMap(kid, vid);
}

void QDBusArgument_BeginMap2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginMap();
}

void QDBusArgument_BeginMapEntry(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginMapEntry2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginStructure(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

void QDBusArgument_BeginStructure2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

long long QDBusArgument_CurrentType(void* ptr)
{
	return static_cast<QDBusArgument*>(ptr)->currentType();
}

void QDBusArgument_EndArray(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndArray2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndMap(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMap2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMapEntry(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndMapEntry2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndStructure(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_EndStructure2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_Swap(void* ptr, void* other)
{
	static_cast<QDBusArgument*>(ptr)->swap(*static_cast<QDBusArgument*>(other));
}

void QDBusArgument_DestroyQDBusArgument(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->~QDBusArgument();
}

void* QDBusConnection_QDBusConnection_SessionBus()
{
	return new QDBusConnection(QDBusConnection::sessionBus());
}

void* QDBusConnection_QDBusConnection_SystemBus()
{
	return new QDBusConnection(QDBusConnection::systemBus());
}

void* QDBusConnection_NewQDBusConnection3(void* other)
{
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

void* QDBusConnection_NewQDBusConnection2(void* other)
{
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

void* QDBusConnection_NewQDBusConnection(char* name)
{
	return new QDBusConnection(QString(name));
}

char* QDBusConnection_BaseService(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusConnection*>(ptr)->baseService().toUtf8().constData());
}

void* QDBusConnection_AsyncCall(void* ptr, void* message, int timeout)
{
	return new QDBusPendingCall(static_cast<QDBusConnection*>(ptr)->asyncCall(*static_cast<QDBusMessage*>(message), timeout));
}

char QDBusConnection_CallWithCallback(void* ptr, void* message, void* receiver, char* returnMethod, char* errorMethod, int timeout)
{
	return static_cast<QDBusConnection*>(ptr)->callWithCallback(*static_cast<QDBusMessage*>(message), static_cast<QObject*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod), timeout);
}

char QDBusConnection_Connect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot)
{
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

char QDBusConnection_Connect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot)
{
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

char QDBusConnection_Connect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot)
{
	return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

void* QDBusConnection_QDBusConnection_ConnectToBus(long long ty, char* name)
{
	return new QDBusConnection(QDBusConnection::connectToBus(static_cast<QDBusConnection::BusType>(ty), QString(name)));
}

void* QDBusConnection_QDBusConnection_ConnectToBus2(char* address, char* name)
{
	return new QDBusConnection(QDBusConnection::connectToBus(QString(address), QString(name)));
}

void* QDBusConnection_QDBusConnection_ConnectToPeer(char* address, char* name)
{
	return new QDBusConnection(QDBusConnection::connectToPeer(QString(address), QString(name)));
}

long long QDBusConnection_ConnectionCapabilities(void* ptr)
{
	return static_cast<QDBusConnection*>(ptr)->connectionCapabilities();
}

char QDBusConnection_Disconnect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot)
{
	return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

char QDBusConnection_Disconnect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot)
{
	return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

char QDBusConnection_Disconnect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot)
{
	return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
}

void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name)
{
	QDBusConnection::disconnectFromBus(QString(name));
}

void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name)
{
	QDBusConnection::disconnectFromPeer(QString(name));
}

void* QDBusConnection_Interface(void* ptr)
{
	return static_cast<QDBusConnection*>(ptr)->interface();
}

char QDBusConnection_IsConnected(void* ptr)
{
	return static_cast<QDBusConnection*>(ptr)->isConnected();
}

void* QDBusConnection_LastError(void* ptr)
{
	return new QDBusError(static_cast<QDBusConnection*>(ptr)->lastError());
}

char* QDBusConnection_QDBusConnection_LocalMachineId()
{
	return const_cast<char*>(QDBusConnection::localMachineId().toHex().constData());
}

char* QDBusConnection_Name(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusConnection*>(ptr)->name().toUtf8().constData());
}

void* QDBusConnection_ObjectRegisteredAt(void* ptr, char* path)
{
	return static_cast<QDBusConnection*>(ptr)->objectRegisteredAt(QString(path));
}

char QDBusConnection_RegisterObject(void* ptr, char* path, void* object, long long options)
{
	return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
}

char QDBusConnection_RegisterObject2(void* ptr, char* path, char* interfa, void* object, long long options)
{
	return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), QString(interfa), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
}

char QDBusConnection_RegisterService(void* ptr, char* serviceName)
{
	return static_cast<QDBusConnection*>(ptr)->registerService(QString(serviceName));
}

char QDBusConnection_Send(void* ptr, void* message)
{
	return static_cast<QDBusConnection*>(ptr)->send(*static_cast<QDBusMessage*>(message));
}

void QDBusConnection_Swap(void* ptr, void* other)
{
	static_cast<QDBusConnection*>(ptr)->swap(*static_cast<QDBusConnection*>(other));
}

void QDBusConnection_UnregisterObject(void* ptr, char* path, long long mode)
{
	static_cast<QDBusConnection*>(ptr)->unregisterObject(QString(path), static_cast<QDBusConnection::UnregisterMode>(mode));
}

char QDBusConnection_UnregisterService(void* ptr, char* serviceName)
{
	return static_cast<QDBusConnection*>(ptr)->unregisterService(QString(serviceName));
}

void QDBusConnection_DestroyQDBusConnection(void* ptr)
{
	static_cast<QDBusConnection*>(ptr)->~QDBusConnection();
}

class MyQDBusConnectionInterface: public QDBusConnectionInterface
{
public:
	void Signal_CallWithCallbackFailed(const QDBusError & error, const QDBusMessage & call) { callbackQDBusConnectionInterface_CallWithCallbackFailed(this, const_cast<QDBusError*>(&error), const_cast<QDBusMessage*>(&call)); };
	void Signal_ServiceRegistered(const QString & serviceName) { callbackQDBusConnectionInterface_ServiceRegistered(this, const_cast<char*>(serviceName.toUtf8().constData())); };
	void Signal_ServiceUnregistered(const QString & serviceName) { callbackQDBusConnectionInterface_ServiceUnregistered(this, const_cast<char*>(serviceName.toUtf8().constData())); };
	void timerEvent(QTimerEvent * event) { callbackQDBusConnectionInterface_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDBusConnectionInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusConnectionInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusConnectionInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusConnectionInterface_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusConnectionInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDBusConnectionInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusConnectionInterface_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusConnectionInterface_MetaObject(const_cast<MyQDBusConnectionInterface*>(this))); };
};

void QDBusConnectionInterface_ConnectCallWithCallbackFailed(void* ptr)
{
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&QDBusConnectionInterface::callWithCallbackFailed), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&MyQDBusConnectionInterface::Signal_CallWithCallbackFailed));
}

void QDBusConnectionInterface_DisconnectCallWithCallbackFailed(void* ptr)
{
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&QDBusConnectionInterface::callWithCallbackFailed), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&MyQDBusConnectionInterface::Signal_CallWithCallbackFailed));
}

void QDBusConnectionInterface_CallWithCallbackFailed(void* ptr, void* error, void* call)
{
	static_cast<QDBusConnectionInterface*>(ptr)->callWithCallbackFailed(*static_cast<QDBusError*>(error), *static_cast<QDBusMessage*>(call));
}

void QDBusConnectionInterface_ConnectServiceRegistered(void* ptr)
{
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));
}

void QDBusConnectionInterface_DisconnectServiceRegistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));
}

void QDBusConnectionInterface_ServiceRegistered(void* ptr, char* serviceName)
{
	static_cast<QDBusConnectionInterface*>(ptr)->serviceRegistered(QString(serviceName));
}

void QDBusConnectionInterface_ConnectServiceUnregistered(void* ptr)
{
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));
}

void QDBusConnectionInterface_DisconnectServiceUnregistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));
}

void QDBusConnectionInterface_ServiceUnregistered(void* ptr, char* serviceName)
{
	static_cast<QDBusConnectionInterface*>(ptr)->serviceUnregistered(QString(serviceName));
}

void QDBusConnectionInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusConnectionInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusConnectionInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusConnectionInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusConnectionInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusConnectionInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusConnectionInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusConnectionInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusConnectionInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusConnectionInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusConnectionInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusConnectionInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::customEvent(static_cast<QEvent*>(event));
}

void QDBusConnectionInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusConnectionInterface*>(ptr), "deleteLater");
}

void QDBusConnectionInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::deleteLater();
}

void QDBusConnectionInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusConnectionInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusConnectionInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusConnectionInterface_Event(void* ptr, void* e)
{
	return static_cast<QDBusConnectionInterface*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusConnectionInterface_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::event(static_cast<QEvent*>(e));
}

char QDBusConnectionInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusConnectionInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusConnectionInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusConnectionInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusConnectionInterface*>(ptr)->metaObject());
}

void* QDBusConnectionInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::metaObject());
}

void* QDBusContext_NewQDBusContext()
{
	return new QDBusContext();
}

char QDBusContext_CalledFromDBus(void* ptr)
{
	return static_cast<QDBusContext*>(ptr)->calledFromDBus();
}

void* QDBusContext_Connection(void* ptr)
{
	return new QDBusConnection(static_cast<QDBusContext*>(ptr)->connection());
}

char QDBusContext_IsDelayedReply(void* ptr)
{
	return static_cast<QDBusContext*>(ptr)->isDelayedReply();
}

void* QDBusContext_Message(void* ptr)
{
	return const_cast<QDBusMessage*>(&static_cast<QDBusContext*>(ptr)->message());
}

void QDBusContext_SendErrorReply2(void* ptr, long long ty, char* msg)
{
	static_cast<QDBusContext*>(ptr)->sendErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg));
}

void QDBusContext_SendErrorReply(void* ptr, char* name, char* msg)
{
	static_cast<QDBusContext*>(ptr)->sendErrorReply(QString(name), QString(msg));
}

void QDBusContext_SetDelayedReply(void* ptr, char enable)
{
	static_cast<QDBusContext*>(ptr)->setDelayedReply(enable != 0);
}

void QDBusContext_DestroyQDBusContext(void* ptr)
{
	static_cast<QDBusContext*>(ptr)->~QDBusContext();
}

void* QDBusError_NewQDBusError(void* other)
{
	return new QDBusError(*static_cast<QDBusError*>(other));
}

char* QDBusError_QDBusError_ErrorString(long long error)
{
	return const_cast<char*>(QDBusError::errorString(static_cast<QDBusError::ErrorType>(error)).toUtf8().constData());
}

char QDBusError_IsValid(void* ptr)
{
	return static_cast<QDBusError*>(ptr)->isValid();
}

char* QDBusError_Message(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusError*>(ptr)->message().toUtf8().constData());
}

char* QDBusError_Name(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusError*>(ptr)->name().toUtf8().constData());
}

void QDBusError_Swap(void* ptr, void* other)
{
	static_cast<QDBusError*>(ptr)->swap(*static_cast<QDBusError*>(other));
}

long long QDBusError_Type(void* ptr)
{
	return static_cast<QDBusError*>(ptr)->type();
}

void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent)
{
	return new QDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QObject*>(parent));
}

void QDBusInterface_DestroyQDBusInterface(void* ptr)
{
	static_cast<QDBusInterface*>(ptr)->~QDBusInterface();
}

void QDBusInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::customEvent(static_cast<QEvent*>(event));
}

void QDBusInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusInterface*>(ptr), "deleteLater");
}

void QDBusInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::deleteLater();
}

void QDBusInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusInterface*>(ptr)->QDBusInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusInterface_Event(void* ptr, void* e)
{
	return static_cast<QDBusInterface*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusInterface_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusInterface*>(ptr)->QDBusInterface::event(static_cast<QEvent*>(e));
}

char QDBusInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusInterface*>(ptr)->QDBusInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusInterface*>(ptr)->metaObject());
}

void* QDBusInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusInterface*>(ptr)->QDBusInterface::metaObject());
}

void* QDBusMessage_CreateErrorReply3(void* ptr, long long ty, char* msg)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg)));
}

void* QDBusMessage_NewQDBusMessage()
{
	return new QDBusMessage();
}

void* QDBusMessage_NewQDBusMessage2(void* other)
{
	return new QDBusMessage(*static_cast<QDBusMessage*>(other));
}

char QDBusMessage_AutoStartService(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->autoStartService();
}

void* QDBusMessage_QDBusMessage_CreateError3(long long ty, char* msg)
{
	return new QDBusMessage(QDBusMessage::createError(static_cast<QDBusError::ErrorType>(ty), QString(msg)));
}

void* QDBusMessage_QDBusMessage_CreateError2(void* error)
{
	return new QDBusMessage(QDBusMessage::createError(*static_cast<QDBusError*>(error)));
}

void* QDBusMessage_QDBusMessage_CreateError(char* name, char* msg)
{
	return new QDBusMessage(QDBusMessage::createError(QString(name), QString(msg)));
}

void* QDBusMessage_CreateErrorReply2(void* ptr, void* error)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(*static_cast<QDBusError*>(error)));
}

void* QDBusMessage_CreateErrorReply(void* ptr, char* name, char* msg)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(QString(name), QString(msg)));
}

void* QDBusMessage_QDBusMessage_CreateMethodCall(char* service, char* path, char* interfa, char* method)
{
	return new QDBusMessage(QDBusMessage::createMethodCall(QString(service), QString(path), QString(interfa), QString(method)));
}

void* QDBusMessage_CreateReply2(void* ptr, void* argument)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createReply(*static_cast<QVariant*>(argument)));
}

void* QDBusMessage_QDBusMessage_CreateSignal(char* path, char* interfa, char* name)
{
	return new QDBusMessage(QDBusMessage::createSignal(QString(path), QString(interfa), QString(name)));
}

void* QDBusMessage_QDBusMessage_CreateTargetedSignal(char* service, char* path, char* interfa, char* name)
{
	return new QDBusMessage(QDBusMessage::createTargetedSignal(QString(service), QString(path), QString(interfa), QString(name)));
}

char* QDBusMessage_ErrorMessage(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->errorMessage().toUtf8().constData());
}

char* QDBusMessage_ErrorName(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->errorName().toUtf8().constData());
}

char* QDBusMessage_Interface(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->interface().toUtf8().constData());
}

char QDBusMessage_IsDelayedReply(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->isDelayedReply();
}

char QDBusMessage_IsReplyRequired(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->isReplyRequired();
}

char* QDBusMessage_Member(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->member().toUtf8().constData());
}

char* QDBusMessage_Path(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->path().toUtf8().constData());
}

char* QDBusMessage_Service(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->service().toUtf8().constData());
}

void QDBusMessage_SetAutoStartService(void* ptr, char enable)
{
	static_cast<QDBusMessage*>(ptr)->setAutoStartService(enable != 0);
}

void QDBusMessage_SetDelayedReply(void* ptr, char enable)
{
	static_cast<QDBusMessage*>(ptr)->setDelayedReply(enable != 0);
}

char* QDBusMessage_Signature(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusMessage*>(ptr)->signature().toUtf8().constData());
}

void QDBusMessage_Swap(void* ptr, void* other)
{
	static_cast<QDBusMessage*>(ptr)->swap(*static_cast<QDBusMessage*>(other));
}

long long QDBusMessage_Type(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->type();
}

void QDBusMessage_DestroyQDBusMessage(void* ptr)
{
	static_cast<QDBusMessage*>(ptr)->~QDBusMessage();
}

void* QDBusObjectPath_NewQDBusObjectPath()
{
	return new QDBusObjectPath();
}

void* QDBusObjectPath_NewQDBusObjectPath3(void* path)
{
	return new QDBusObjectPath(*static_cast<QLatin1String*>(path));
}

void* QDBusObjectPath_NewQDBusObjectPath5(char* p)
{
	return new QDBusObjectPath(*(new QString(p)));
}

void* QDBusObjectPath_NewQDBusObjectPath4(char* path)
{
	return new QDBusObjectPath(QString(path));
}

void* QDBusObjectPath_NewQDBusObjectPath2(char* path)
{
	return new QDBusObjectPath(const_cast<const char*>(path));
}

char* QDBusObjectPath_Path(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusObjectPath*>(ptr)->path().toUtf8().constData());
}

void QDBusObjectPath_SetPath(void* ptr, char* path)
{
	static_cast<QDBusObjectPath*>(ptr)->setPath(QString(path));
}

void QDBusObjectPath_Swap(void* ptr, void* other)
{
	static_cast<QDBusObjectPath*>(ptr)->swap(*static_cast<QDBusObjectPath*>(other));
}

void* QDBusPendingCall_NewQDBusPendingCall(void* other)
{
	return new QDBusPendingCall(*static_cast<QDBusPendingCall*>(other));
}

void* QDBusPendingCall_QDBusPendingCall_FromCompletedCall(void* msg)
{
	return new QDBusPendingCall(QDBusPendingCall::fromCompletedCall(*static_cast<QDBusMessage*>(msg)));
}

void* QDBusPendingCall_QDBusPendingCall_FromError(void* error)
{
	return new QDBusPendingCall(QDBusPendingCall::fromError(*static_cast<QDBusError*>(error)));
}

void QDBusPendingCall_Swap(void* ptr, void* other)
{
	static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr)
{
	static_cast<QDBusPendingCall*>(ptr)->~QDBusPendingCall();
}

class MyQDBusPendingCallWatcher: public QDBusPendingCallWatcher
{
public:
	MyQDBusPendingCallWatcher(const QDBusPendingCall &call, QObject *parent) : QDBusPendingCallWatcher(call, parent) {};
	void Signal_Finished(QDBusPendingCallWatcher * self) { callbackQDBusPendingCallWatcher_Finished(this, self); };
	void timerEvent(QTimerEvent * event) { callbackQDBusPendingCallWatcher_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDBusPendingCallWatcher_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusPendingCallWatcher_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusPendingCallWatcher_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusPendingCallWatcher_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusPendingCallWatcher_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDBusPendingCallWatcher_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusPendingCallWatcher_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusPendingCallWatcher_MetaObject(const_cast<MyQDBusPendingCallWatcher*>(this))); };
};

void QDBusPendingCallWatcher_WaitForFinished(void* ptr)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->waitForFinished();
}

void* QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(void* call, void* parent)
{
	return new MyQDBusPendingCallWatcher(*static_cast<QDBusPendingCall*>(call), static_cast<QObject*>(parent));
}

void QDBusPendingCallWatcher_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));
}

void QDBusPendingCallWatcher_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));
}

void QDBusPendingCallWatcher_Finished(void* ptr, void* self)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->finished(static_cast<QDBusPendingCallWatcher*>(self));
}

char QDBusPendingCallWatcher_IsFinished(void* ptr)
{
	return static_cast<QDBusPendingCallWatcher*>(ptr)->isFinished();
}

void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->~QDBusPendingCallWatcher();
}

void QDBusPendingCallWatcher_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusPendingCallWatcher_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusPendingCallWatcher_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusPendingCallWatcher_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusPendingCallWatcher_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusPendingCallWatcher_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::customEvent(static_cast<QEvent*>(event));
}

void QDBusPendingCallWatcher_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusPendingCallWatcher*>(ptr), "deleteLater");
}

void QDBusPendingCallWatcher_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::deleteLater();
}

void QDBusPendingCallWatcher_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusPendingCallWatcher_Event(void* ptr, void* e)
{
	return static_cast<QDBusPendingCallWatcher*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusPendingCallWatcher_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::event(static_cast<QEvent*>(e));
}

char QDBusPendingCallWatcher_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusPendingCallWatcher*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusPendingCallWatcher_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusPendingCallWatcher_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusPendingCallWatcher*>(ptr)->metaObject());
}

void* QDBusPendingCallWatcher_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::metaObject());
}

class MyQDBusServer: public QDBusServer
{
public:
	MyQDBusServer(QObject *parent) : QDBusServer(parent) {};
	MyQDBusServer(const QString &address, QObject *parent) : QDBusServer(address, parent) {};
	void Signal_NewConnection(const QDBusConnection & connection) { callbackQDBusServer_NewConnection(this, const_cast<QDBusConnection*>(&connection)); };
	void timerEvent(QTimerEvent * event) { callbackQDBusServer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDBusServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusServer_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusServer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDBusServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusServer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusServer_MetaObject(const_cast<MyQDBusServer*>(this))); };
};

void* QDBusServer_NewQDBusServer2(void* parent)
{
	return new MyQDBusServer(static_cast<QObject*>(parent));
}

void* QDBusServer_NewQDBusServer(char* address, void* parent)
{
	return new MyQDBusServer(QString(address), static_cast<QObject*>(parent));
}

char* QDBusServer_Address(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusServer*>(ptr)->address().toUtf8().constData());
}

char QDBusServer_IsAnonymousAuthenticationAllowed(void* ptr)
{
	return static_cast<QDBusServer*>(ptr)->isAnonymousAuthenticationAllowed();
}

char QDBusServer_IsConnected(void* ptr)
{
	return static_cast<QDBusServer*>(ptr)->isConnected();
}

void* QDBusServer_LastError(void* ptr)
{
	return new QDBusError(static_cast<QDBusServer*>(ptr)->lastError());
}

void QDBusServer_ConnectNewConnection(void* ptr)
{
	QObject::connect(static_cast<QDBusServer*>(ptr), static_cast<void (QDBusServer::*)(const QDBusConnection &)>(&QDBusServer::newConnection), static_cast<MyQDBusServer*>(ptr), static_cast<void (MyQDBusServer::*)(const QDBusConnection &)>(&MyQDBusServer::Signal_NewConnection));
}

void QDBusServer_DisconnectNewConnection(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServer*>(ptr), static_cast<void (QDBusServer::*)(const QDBusConnection &)>(&QDBusServer::newConnection), static_cast<MyQDBusServer*>(ptr), static_cast<void (MyQDBusServer::*)(const QDBusConnection &)>(&MyQDBusServer::Signal_NewConnection));
}

void QDBusServer_NewConnection(void* ptr, void* connection)
{
	static_cast<QDBusServer*>(ptr)->newConnection(*static_cast<QDBusConnection*>(connection));
}

void QDBusServer_SetAnonymousAuthenticationAllowed(void* ptr, char value)
{
	static_cast<QDBusServer*>(ptr)->setAnonymousAuthenticationAllowed(value != 0);
}

void QDBusServer_DestroyQDBusServer(void* ptr)
{
	static_cast<QDBusServer*>(ptr)->~QDBusServer();
}

void QDBusServer_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusServer*>(ptr)->QDBusServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServer_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusServer*>(ptr)->QDBusServer::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusServer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusServer*>(ptr)->QDBusServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServer_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusServer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusServer*>(ptr)->QDBusServer::customEvent(static_cast<QEvent*>(event));
}

void QDBusServer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusServer*>(ptr), "deleteLater");
}

void QDBusServer_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusServer*>(ptr)->QDBusServer::deleteLater();
}

void QDBusServer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusServer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusServer*>(ptr)->QDBusServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusServer_Event(void* ptr, void* e)
{
	return static_cast<QDBusServer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusServer*>(ptr)->QDBusServer::event(static_cast<QEvent*>(e));
}

char QDBusServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusServer*>(ptr)->QDBusServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusServer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusServer*>(ptr)->metaObject());
}

void* QDBusServer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusServer*>(ptr)->QDBusServer::metaObject());
}

class MyQDBusServiceWatcher: public QDBusServiceWatcher
{
public:
	MyQDBusServiceWatcher(QObject *parent) : QDBusServiceWatcher(parent) {};
	MyQDBusServiceWatcher(const QString &service, const QDBusConnection &connection, WatchMode watchMode, QObject *parent) : QDBusServiceWatcher(service, connection, watchMode, parent) {};
	void Signal_ServiceOwnerChanged(const QString & serviceName, const QString & oldOwner, const QString & newOwner) { callbackQDBusServiceWatcher_ServiceOwnerChanged(this, const_cast<char*>(serviceName.toUtf8().constData()), const_cast<char*>(oldOwner.toUtf8().constData()), const_cast<char*>(newOwner.toUtf8().constData())); };
	void Signal_ServiceRegistered(const QString & serviceName) { callbackQDBusServiceWatcher_ServiceRegistered(this, const_cast<char*>(serviceName.toUtf8().constData())); };
	void Signal_ServiceUnregistered(const QString & serviceName) { callbackQDBusServiceWatcher_ServiceUnregistered(this, const_cast<char*>(serviceName.toUtf8().constData())); };
	void timerEvent(QTimerEvent * event) { callbackQDBusServiceWatcher_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDBusServiceWatcher_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusServiceWatcher_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusServiceWatcher_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusServiceWatcher_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusServiceWatcher_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDBusServiceWatcher_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusServiceWatcher_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusServiceWatcher_MetaObject(const_cast<MyQDBusServiceWatcher*>(this))); };
};

void QDBusServiceWatcher_SetWatchMode(void* ptr, long long mode)
{
	static_cast<QDBusServiceWatcher*>(ptr)->setWatchMode(static_cast<QDBusServiceWatcher::WatchModeFlag>(mode));
}

long long QDBusServiceWatcher_WatchMode(void* ptr)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->watchMode();
}

void* QDBusServiceWatcher_NewQDBusServiceWatcher(void* parent)
{
	return new MyQDBusServiceWatcher(static_cast<QObject*>(parent));
}

void* QDBusServiceWatcher_NewQDBusServiceWatcher2(char* service, void* connection, long long watchMode, void* parent)
{
	return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QObject*>(parent));
}

void QDBusServiceWatcher_AddWatchedService(void* ptr, char* newService)
{
	static_cast<QDBusServiceWatcher*>(ptr)->addWatchedService(QString(newService));
}

void* QDBusServiceWatcher_Connection(void* ptr)
{
	return new QDBusConnection(static_cast<QDBusServiceWatcher*>(ptr)->connection());
}

char QDBusServiceWatcher_RemoveWatchedService(void* ptr, char* service)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->removeWatchedService(QString(service));
}

void QDBusServiceWatcher_ConnectServiceOwnerChanged(void* ptr)
{
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&QDBusServiceWatcher::serviceOwnerChanged), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceOwnerChanged));
}

void QDBusServiceWatcher_DisconnectServiceOwnerChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&QDBusServiceWatcher::serviceOwnerChanged), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceOwnerChanged));
}

void QDBusServiceWatcher_ServiceOwnerChanged(void* ptr, char* serviceName, char* oldOwner, char* newOwner)
{
	static_cast<QDBusServiceWatcher*>(ptr)->serviceOwnerChanged(QString(serviceName), QString(oldOwner), QString(newOwner));
}

void QDBusServiceWatcher_ConnectServiceRegistered(void* ptr)
{
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));
}

void QDBusServiceWatcher_DisconnectServiceRegistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));
}

void QDBusServiceWatcher_ServiceRegistered(void* ptr, char* serviceName)
{
	static_cast<QDBusServiceWatcher*>(ptr)->serviceRegistered(QString(serviceName));
}

void QDBusServiceWatcher_ConnectServiceUnregistered(void* ptr)
{
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));
}

void QDBusServiceWatcher_DisconnectServiceUnregistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));
}

void QDBusServiceWatcher_ServiceUnregistered(void* ptr, char* serviceName)
{
	static_cast<QDBusServiceWatcher*>(ptr)->serviceUnregistered(QString(serviceName));
}

void QDBusServiceWatcher_SetConnection(void* ptr, void* connection)
{
	static_cast<QDBusServiceWatcher*>(ptr)->setConnection(*static_cast<QDBusConnection*>(connection));
}

void QDBusServiceWatcher_SetWatchedServices(void* ptr, char* services)
{
	static_cast<QDBusServiceWatcher*>(ptr)->setWatchedServices(QString(services).split("|", QString::SkipEmptyParts));
}

char* QDBusServiceWatcher_WatchedServices(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusServiceWatcher*>(ptr)->watchedServices().join("|").toUtf8().constData());
}

void QDBusServiceWatcher_DestroyQDBusServiceWatcher(void* ptr)
{
	static_cast<QDBusServiceWatcher*>(ptr)->~QDBusServiceWatcher();
}

void QDBusServiceWatcher_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusServiceWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServiceWatcher_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusServiceWatcher_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusServiceWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServiceWatcher_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServiceWatcher_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusServiceWatcher*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServiceWatcher_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServiceWatcher_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusServiceWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusServiceWatcher_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::customEvent(static_cast<QEvent*>(event));
}

void QDBusServiceWatcher_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusServiceWatcher*>(ptr), "deleteLater");
}

void QDBusServiceWatcher_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::deleteLater();
}

void QDBusServiceWatcher_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusServiceWatcher*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServiceWatcher_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusServiceWatcher_Event(void* ptr, void* e)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusServiceWatcher_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::event(static_cast<QEvent*>(e));
}

char QDBusServiceWatcher_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusServiceWatcher_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusServiceWatcher_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusServiceWatcher*>(ptr)->metaObject());
}

void* QDBusServiceWatcher_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::metaObject());
}

void* QDBusSignature_NewQDBusSignature()
{
	return new QDBusSignature();
}

void* QDBusSignature_NewQDBusSignature3(void* signature)
{
	return new QDBusSignature(*static_cast<QLatin1String*>(signature));
}

void* QDBusSignature_NewQDBusSignature5(char* sig)
{
	return new QDBusSignature(*(new QString(sig)));
}

void* QDBusSignature_NewQDBusSignature4(char* signature)
{
	return new QDBusSignature(QString(signature));
}

void* QDBusSignature_NewQDBusSignature2(char* signature)
{
	return new QDBusSignature(const_cast<const char*>(signature));
}

void QDBusSignature_SetSignature(void* ptr, char* signature)
{
	static_cast<QDBusSignature*>(ptr)->setSignature(QString(signature));
}

char* QDBusSignature_Signature(void* ptr)
{
	return const_cast<char*>(static_cast<QDBusSignature*>(ptr)->signature().toUtf8().constData());
}

void QDBusSignature_Swap(void* ptr, void* other)
{
	static_cast<QDBusSignature*>(ptr)->swap(*static_cast<QDBusSignature*>(other));
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor()
{
	return new QDBusUnixFileDescriptor();
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(void* other)
{
	return new QDBusUnixFileDescriptor(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor)
{
	return new QDBusUnixFileDescriptor(fileDescriptor);
}

int QDBusUnixFileDescriptor_FileDescriptor(void* ptr)
{
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->fileDescriptor();
}

char QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported()
{
	return QDBusUnixFileDescriptor::isSupported();
}

char QDBusUnixFileDescriptor_IsValid(void* ptr)
{
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->isValid();
}

void QDBusUnixFileDescriptor_SetFileDescriptor(void* ptr, int fileDescriptor)
{
	static_cast<QDBusUnixFileDescriptor*>(ptr)->setFileDescriptor(fileDescriptor);
}

void QDBusUnixFileDescriptor_Swap(void* ptr, void* other)
{
	static_cast<QDBusUnixFileDescriptor*>(ptr)->swap(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(void* ptr)
{
	static_cast<QDBusUnixFileDescriptor*>(ptr)->~QDBusUnixFileDescriptor();
}

void* QDBusVariant_NewQDBusVariant()
{
	return new QDBusVariant();
}

void* QDBusVariant_NewQDBusVariant3(void* v)
{
	return new QDBusVariant(*static_cast<QVariant*>(v));
}

void* QDBusVariant_NewQDBusVariant2(void* variant)
{
	return new QDBusVariant(*static_cast<QVariant*>(variant));
}

void QDBusVariant_SetVariant(void* ptr, void* variant)
{
	static_cast<QDBusVariant*>(ptr)->setVariant(*static_cast<QVariant*>(variant));
}

void QDBusVariant_Swap(void* ptr, void* other)
{
	static_cast<QDBusVariant*>(ptr)->swap(*static_cast<QDBusVariant*>(other));
}

void* QDBusVariant_Variant(void* ptr)
{
	return new QVariant(static_cast<QDBusVariant*>(ptr)->variant());
}

class MyQDBusVirtualObject: public QDBusVirtualObject
{
public:
	MyQDBusVirtualObject(QObject *parent) : QDBusVirtualObject(parent) {};
	bool handleMessage(const QDBusMessage & message, const QDBusConnection & connection) { return callbackQDBusVirtualObject_HandleMessage(this, const_cast<QDBusMessage*>(&message), const_cast<QDBusConnection*>(&connection)) != 0; };
	QString introspect(const QString & path) const { return QString(callbackQDBusVirtualObject_Introspect(const_cast<MyQDBusVirtualObject*>(this), const_cast<char*>(path.toUtf8().constData()))); };
	void timerEvent(QTimerEvent * event) { callbackQDBusVirtualObject_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQDBusVirtualObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusVirtualObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusVirtualObject_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusVirtualObject_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusVirtualObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDBusVirtualObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusVirtualObject_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusVirtualObject_MetaObject(const_cast<MyQDBusVirtualObject*>(this))); };
};

void* QDBusVirtualObject_NewQDBusVirtualObject(void* parent)
{
	return new MyQDBusVirtualObject(static_cast<QObject*>(parent));
}

char QDBusVirtualObject_HandleMessage(void* ptr, void* message, void* connection)
{
	return static_cast<QDBusVirtualObject*>(ptr)->handleMessage(*static_cast<QDBusMessage*>(message), *static_cast<QDBusConnection*>(connection));
}

char* QDBusVirtualObject_Introspect(void* ptr, char* path)
{
	return const_cast<char*>(static_cast<QDBusVirtualObject*>(ptr)->introspect(QString(path)).toUtf8().constData());
}

void QDBusVirtualObject_DestroyQDBusVirtualObject(void* ptr)
{
	static_cast<QDBusVirtualObject*>(ptr)->~QDBusVirtualObject();
}

void QDBusVirtualObject_TimerEvent(void* ptr, void* event)
{
	static_cast<QDBusVirtualObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusVirtualObject_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusVirtualObject_ChildEvent(void* ptr, void* event)
{
	static_cast<QDBusVirtualObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusVirtualObject_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusVirtualObject_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusVirtualObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusVirtualObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusVirtualObject_CustomEvent(void* ptr, void* event)
{
	static_cast<QDBusVirtualObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusVirtualObject_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::customEvent(static_cast<QEvent*>(event));
}

void QDBusVirtualObject_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDBusVirtualObject*>(ptr), "deleteLater");
}

void QDBusVirtualObject_DeleteLaterDefault(void* ptr)
{
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::deleteLater();
}

void QDBusVirtualObject_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDBusVirtualObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusVirtualObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDBusVirtualObject_Event(void* ptr, void* e)
{
	return static_cast<QDBusVirtualObject*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusVirtualObject_EventDefault(void* ptr, void* e)
{
	return static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::event(static_cast<QEvent*>(e));
}

char QDBusVirtualObject_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusVirtualObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QDBusVirtualObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDBusVirtualObject_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusVirtualObject*>(ptr)->metaObject());
}

void* QDBusVirtualObject_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::metaObject());
}

