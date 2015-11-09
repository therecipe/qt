#include "qnetworkconfigurationmanager.h"
#include <QModelIndex>
#include <QNetworkConfiguration>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNetworkConfigurationManager>
#include "_cgo_export.h"

class MyQNetworkConfigurationManager: public QNetworkConfigurationManager {
public:
void Signal_OnlineStateChanged(bool isOnline){callbackQNetworkConfigurationManagerOnlineStateChanged(this->objectName().toUtf8().data(), isOnline);};
void Signal_UpdateCompleted(){callbackQNetworkConfigurationManagerUpdateCompleted(this->objectName().toUtf8().data());};
};

void* QNetworkConfigurationManager_NewQNetworkConfigurationManager(void* parent){
	return new QNetworkConfigurationManager(static_cast<QObject*>(parent));
}

int QNetworkConfigurationManager_Capabilities(void* ptr){
	return static_cast<QNetworkConfigurationManager*>(ptr)->capabilities();
}

int QNetworkConfigurationManager_IsOnline(void* ptr){
	return static_cast<QNetworkConfigurationManager*>(ptr)->isOnline();
}

void QNetworkConfigurationManager_ConnectOnlineStateChanged(void* ptr){
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));;
}

void QNetworkConfigurationManager_DisconnectOnlineStateChanged(void* ptr){
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));;
}

void QNetworkConfigurationManager_ConnectUpdateCompleted(void* ptr){
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));;
}

void QNetworkConfigurationManager_DisconnectUpdateCompleted(void* ptr){
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));;
}

void QNetworkConfigurationManager_UpdateConfigurations(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkConfigurationManager*>(ptr), "updateConfigurations");
}

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(void* ptr){
	static_cast<QNetworkConfigurationManager*>(ptr)->~QNetworkConfigurationManager();
}

