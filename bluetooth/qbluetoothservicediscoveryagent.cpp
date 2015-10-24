#include "qbluetoothservicediscoveryagent.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothAddress>
#include <QBluetoothUuid>
#include <QMetaObject>
#include <QBluetoothServiceDiscoveryAgent>
#include "_cgo_export.h"

class MyQBluetoothServiceDiscoveryAgent: public QBluetoothServiceDiscoveryAgent {
public:
void Signal_Canceled(){callbackQBluetoothServiceDiscoveryAgentCanceled(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQBluetoothServiceDiscoveryAgentFinished(this->objectName().toUtf8().data());};
};

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

QtObjectPtr QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(QtObjectPtr parent){
	return new QBluetoothServiceDiscoveryAgent(static_cast<QObject*>(parent));
}

QtObjectPtr QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(QtObjectPtr deviceAdapter, QtObjectPtr parent){
	return new QBluetoothServiceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

void QBluetoothServiceDiscoveryAgent_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "clear");
}

int QBluetoothServiceDiscoveryAgent_Error(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->error();
}

char* QBluetoothServiceDiscoveryAgent_ErrorString(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothServiceDiscoveryAgent_IsActive(QtObjectPtr ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->isActive();
}

int QBluetoothServiceDiscoveryAgent_SetRemoteAddress(QtObjectPtr ptr, QtObjectPtr address){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setRemoteAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothServiceDiscoveryAgent_SetUuidFilter2(QtObjectPtr ptr, QtObjectPtr uuid){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setUuidFilter(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothServiceDiscoveryAgent_Start(QtObjectPtr ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "start", Q_ARG(QBluetoothServiceDiscoveryAgent::DiscoveryMode, static_cast<QBluetoothServiceDiscoveryAgent::DiscoveryMode>(mode)));
}

void QBluetoothServiceDiscoveryAgent_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(QtObjectPtr ptr){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->~QBluetoothServiceDiscoveryAgent();
}

