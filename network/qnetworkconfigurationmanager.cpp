#include "qnetworkconfigurationmanager.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNetworkConfiguration>
#include <QMetaObject>
#include <QString>
#include <QNetworkConfigurationManager>
#include "_cgo_export.h"

class MyQNetworkConfigurationManager: public QNetworkConfigurationManager {
public:
void Signal_OnlineStateChanged(bool isOnline){callbackQNetworkConfigurationManagerOnlineStateChanged(this->objectName().toUtf8().data(), isOnline);};
void Signal_UpdateCompleted(){callbackQNetworkConfigurationManagerUpdateCompleted(this->objectName().toUtf8().data());};
};

QtObjectPtr QNetworkConfigurationManager_NewQNetworkConfigurationManager(QtObjectPtr parent){
	return new QNetworkConfigurationManager(static_cast<QObject*>(parent));
}

int QNetworkConfigurationManager_Capabilities(QtObjectPtr ptr){
	return static_cast<QNetworkConfigurationManager*>(ptr)->capabilities();
}

int QNetworkConfigurationManager_IsOnline(QtObjectPtr ptr){
	return static_cast<QNetworkConfigurationManager*>(ptr)->isOnline();
}

void QNetworkConfigurationManager_ConnectOnlineStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));;
}

void QNetworkConfigurationManager_DisconnectOnlineStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)(bool)>(&QNetworkConfigurationManager::onlineStateChanged), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)(bool)>(&MyQNetworkConfigurationManager::Signal_OnlineStateChanged));;
}

void QNetworkConfigurationManager_ConnectUpdateCompleted(QtObjectPtr ptr){
	QObject::connect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));;
}

void QNetworkConfigurationManager_DisconnectUpdateCompleted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNetworkConfigurationManager*>(ptr), static_cast<void (QNetworkConfigurationManager::*)()>(&QNetworkConfigurationManager::updateCompleted), static_cast<MyQNetworkConfigurationManager*>(ptr), static_cast<void (MyQNetworkConfigurationManager::*)()>(&MyQNetworkConfigurationManager::Signal_UpdateCompleted));;
}

void QNetworkConfigurationManager_UpdateConfigurations(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QNetworkConfigurationManager*>(ptr), "updateConfigurations");
}

void QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(QtObjectPtr ptr){
	static_cast<QNetworkConfigurationManager*>(ptr)->~QNetworkConfigurationManager();
}

