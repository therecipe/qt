#include "qbluetoothservicediscoveryagent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothUuid>
#include <QObject>
#include <QBluetoothAddress>
#include <QMetaObject>
#include <QBluetoothServiceDiscoveryAgent>
#include "_cgo_export.h"

class MyQBluetoothServiceDiscoveryAgent: public QBluetoothServiceDiscoveryAgent {
public:
void Signal_Canceled(){callbackQBluetoothServiceDiscoveryAgentCanceled(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQBluetoothServiceDiscoveryAgentFinished(this->objectName().toUtf8().data());};
};

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothServiceDiscoveryAgent_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothServiceDiscoveryAgent_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));;
}

void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(void* parent){
	return new QBluetoothServiceDiscoveryAgent(static_cast<QObject*>(parent));
}

void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(void* deviceAdapter, void* parent){
	return new QBluetoothServiceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

void QBluetoothServiceDiscoveryAgent_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "clear");
}

int QBluetoothServiceDiscoveryAgent_Error(void* ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->error();
}

char* QBluetoothServiceDiscoveryAgent_ErrorString(void* ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothServiceDiscoveryAgent_IsActive(void* ptr){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->isActive();
}

int QBluetoothServiceDiscoveryAgent_SetRemoteAddress(void* ptr, void* address){
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setRemoteAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothServiceDiscoveryAgent_SetUuidFilter2(void* ptr, void* uuid){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setUuidFilter(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothServiceDiscoveryAgent_Start(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "start", Q_ARG(QBluetoothServiceDiscoveryAgent::DiscoveryMode, static_cast<QBluetoothServiceDiscoveryAgent::DiscoveryMode>(mode)));
}

void QBluetoothServiceDiscoveryAgent_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(void* ptr){
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->~QBluetoothServiceDiscoveryAgent();
}

