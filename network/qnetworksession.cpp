#include "qnetworksession.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkConfiguration>
#include <QMetaObject>
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

QtObjectPtr QNetworkSession_NewQNetworkSession(QtObjectPtr connectionConfig, QtObjectPtr parent){
	return new QNetworkSession(*static_cast<QNetworkConfiguration*>(connectionConfig), static_cast<QObject*>(parent));
}

void QNetworkSession_Accept(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "accept");
}

void QNetworkSession_Close(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "close");
}

void QNetworkSession_ConnectClosed(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));;
}

void QNetworkSession_DisconnectClosed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::closed), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Closed));;
}

int QNetworkSession_Error(QtObjectPtr ptr){
	return static_cast<QNetworkSession*>(ptr)->error();
}

char* QNetworkSession_ErrorString(QtObjectPtr ptr){
	return static_cast<QNetworkSession*>(ptr)->errorString().toUtf8().data();
}

void QNetworkSession_Ignore(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "ignore");
}

int QNetworkSession_IsOpen(QtObjectPtr ptr){
	return static_cast<QNetworkSession*>(ptr)->isOpen();
}

void QNetworkSession_Migrate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "migrate");
}

void QNetworkSession_ConnectNewConfigurationActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));;
}

void QNetworkSession_DisconnectNewConfigurationActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::newConfigurationActivated), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_NewConfigurationActivated));;
}

void QNetworkSession_Open(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "open");
}

void QNetworkSession_ConnectOpened(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));;
}

void QNetworkSession_DisconnectOpened(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)()>(&QNetworkSession::opened), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)()>(&MyQNetworkSession::Signal_Opened));;
}

void QNetworkSession_Reject(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "reject");
}

char* QNetworkSession_SessionProperty(QtObjectPtr ptr, char* key){
	return static_cast<QNetworkSession*>(ptr)->sessionProperty(QString(key)).toString().toUtf8().data();
}

void QNetworkSession_SetSessionProperty(QtObjectPtr ptr, char* key, char* value){
	static_cast<QNetworkSession*>(ptr)->setSessionProperty(QString(key), QVariant(value));
}

void QNetworkSession_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));;
}

void QNetworkSession_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::State)>(&QNetworkSession::stateChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::State)>(&MyQNetworkSession::Signal_StateChanged));;
}

void QNetworkSession_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkSession*>(ptr), "stop");
}

int QNetworkSession_UsagePolicies(QtObjectPtr ptr){
	return static_cast<QNetworkSession*>(ptr)->usagePolicies();
}

void QNetworkSession_ConnectUsagePoliciesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));;
}

void QNetworkSession_DisconnectUsagePoliciesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkSession*>(ptr), static_cast<void (QNetworkSession::*)(QNetworkSession::UsagePolicies)>(&QNetworkSession::usagePoliciesChanged), static_cast<MyQNetworkSession*>(ptr), static_cast<void (MyQNetworkSession::*)(QNetworkSession::UsagePolicies)>(&MyQNetworkSession::Signal_UsagePoliciesChanged));;
}

int QNetworkSession_WaitForOpened(QtObjectPtr ptr, int msecs){
	return static_cast<QNetworkSession*>(ptr)->waitForOpened(msecs);
}

void QNetworkSession_DestroyQNetworkSession(QtObjectPtr ptr){
	static_cast<QNetworkSession*>(ptr)->~QNetworkSession();
}

