#include "qnetworksession.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QNetworkConfiguration>
#include <QObject>
#include <QNetworkSession>
#include "_cgo_export.h"

class MyQNetworkSession: public QNetworkSession {
public:
void Signal_Closed(){callbackQNetworkSessionClosed(this->objectName().toUtf8().data());};
void Signal_NewConfigurationActivated(){callbackQNetworkSessionNewConfigurationActivated(this->objectName().toUtf8().data());};
void Signal_Opened(){callbackQNetworkSessionOpened(this->objectName().toUtf8().data());};
void Signal_StateChanged(QNetworkSession::State state){callbackQNetworkSessionStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_UsagePoliciesChanged(QNetworkSession::UsagePolicies usagePolicies){callbackQNetworkSessionUsagePoliciesChanged(this->objectName().toUtf8().data(), usagePolicies);};
};

void* QNetworkSession_NewQNetworkSession(void* connectionConfig, void* parent){
	return new QNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QObject*>(parent));
}

void QNetworkSession_Accept(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "accept");
}

void QNetworkSession_Close(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "close");
}

void QNetworkSession_ConnectClosed(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));;
}

void QNetworkSession_DisconnectClosed(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));;
}

int QNetworkSession_Error(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->error();
}

char* QNetworkSession_ErrorString(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->errorString().toUtf8().data();
}

void QNetworkSession_Ignore(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "ignore");
}

int QNetworkSession_IsOpen(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->isOpen();
}

void QNetworkSession_Migrate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "migrate");
}

void QNetworkSession_ConnectNewConfigurationActivated(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));;
}

void QNetworkSession_DisconnectNewConfigurationActivated(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));;
}

void QNetworkSession_Open(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "open");
}

void QNetworkSession_ConnectOpened(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));;
}

void QNetworkSession_DisconnectOpened(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));;
}

void QNetworkSession_Reject(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "reject");
}

void* QNetworkSession_SessionProperty(void* ptr, char* key){
	return new QVariant(static_cast<QNetworkSession*>(ptr)->sessionProperty(QString(key)));
}

void QNetworkSession_SetSessionProperty(void* ptr, char* key, void* value){
	static_cast<QNetworkSession*>(ptr)->setSessionProperty(QString(key), *static_cast<QVariant*>(value));
}

void QNetworkSession_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));;
}

void QNetworkSession_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));;
}

void QNetworkSession_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "stop");
}

int QNetworkSession_UsagePolicies(void* ptr){
	return static_cast<QNetworkSession*>(ptr)->usagePolicies();
}

void QNetworkSession_ConnectUsagePoliciesChanged(void* ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));;
}

void QNetworkSession_DisconnectUsagePoliciesChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));;
}

int QNetworkSession_WaitForOpened(void* ptr, int msecs){
	return static_cast<QNetworkSession*>(ptr)->waitForOpened(msecs);
}

void QNetworkSession_DestroyQNetworkSession(void* ptr){
	static_cast<QNetworkSession*>(ptr)->~QNetworkSession();
}

