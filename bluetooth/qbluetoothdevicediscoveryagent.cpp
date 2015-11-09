#include "qbluetoothdevicediscoveryagent.h"
#include <QBluetoothAddress>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothDeviceDiscoveryAgent>
#include "_cgo_export.h"

class MyQBluetoothDeviceDiscoveryAgent: public QBluetoothDeviceDiscoveryAgent {
public:
void Signal_Canceled(){callbackQBluetoothDeviceDiscoveryAgentCanceled(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQBluetoothDeviceDiscoveryAgentFinished(this->objectName().toUtf8().data());};
};

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(void* parent){
	return new QBluetoothDeviceDiscoveryAgent(static_cast<QObject*>(parent));
}

void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(void* deviceAdapter, void* parent){
	return new QBluetoothDeviceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

int QBluetoothDeviceDiscoveryAgent_Error(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->error();
}

char* QBluetoothDeviceDiscoveryAgent_ErrorString(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothDeviceDiscoveryAgent_InquiryType(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->inquiryType();
}

int QBluetoothDeviceDiscoveryAgent_IsActive(void* ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->isActive();
}

void QBluetoothDeviceDiscoveryAgent_SetInquiryType(void* ptr, int ty){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->setInquiryType(static_cast<QBluetoothDeviceDiscoveryAgent::InquiryType>(ty));
}

void QBluetoothDeviceDiscoveryAgent_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "start");
}

void QBluetoothDeviceDiscoveryAgent_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(void* ptr){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->~QBluetoothDeviceDiscoveryAgent();
}

