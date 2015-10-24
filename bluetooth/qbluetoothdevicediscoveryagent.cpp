#include "qbluetoothdevicediscoveryagent.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothAddress>
#include <QBluetoothDeviceDiscoveryAgent>
#include "_cgo_export.h"

class MyQBluetoothDeviceDiscoveryAgent: public QBluetoothDeviceDiscoveryAgent {
public:
void Signal_Canceled(){callbackQBluetoothDeviceDiscoveryAgentCanceled(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQBluetoothDeviceDiscoveryAgentFinished(this->objectName().toUtf8().data());};
};

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));;
}

void QBluetoothDeviceDiscoveryAgent_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));;
}

QtObjectPtr QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(QtObjectPtr parent){
	return new QBluetoothDeviceDiscoveryAgent(static_cast<QObject*>(parent));
}

QtObjectPtr QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(QtObjectPtr deviceAdapter, QtObjectPtr parent){
	return new QBluetoothDeviceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

int QBluetoothDeviceDiscoveryAgent_Error(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->error();
}

char* QBluetoothDeviceDiscoveryAgent_ErrorString(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->errorString().toUtf8().data();
}

int QBluetoothDeviceDiscoveryAgent_InquiryType(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->inquiryType();
}

int QBluetoothDeviceDiscoveryAgent_IsActive(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->isActive();
}

void QBluetoothDeviceDiscoveryAgent_SetInquiryType(QtObjectPtr ptr, int ty){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->setInquiryType(static_cast<QBluetoothDeviceDiscoveryAgent::InquiryType>(ty));
}

void QBluetoothDeviceDiscoveryAgent_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "start");
}

void QBluetoothDeviceDiscoveryAgent_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(QtObjectPtr ptr){
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->~QBluetoothDeviceDiscoveryAgent();
}

