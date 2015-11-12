#include "qbluetoothtransfermanager.h"
#include <QObject>
#include <QBluetoothTransferReply>
#include <QIODevice>
#include <QBluetoothTransferRequest>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothTransferManager>
#include "_cgo_export.h"

class MyQBluetoothTransferManager: public QBluetoothTransferManager {
public:
void Signal_Finished(QBluetoothTransferReply * reply){callbackQBluetoothTransferManagerFinished(this->objectName().toUtf8().data(), reply);};
};

void* QBluetoothTransferManager_Put(void* ptr, void* request, void* data){
	return static_cast<QBluetoothTransferManager*>(ptr)->put(*static_cast<QBluetoothTransferRequest*>(request), static_cast<QIODevice*>(data));
}

void* QBluetoothTransferManager_NewQBluetoothTransferManager(void* parent){
	return new QBluetoothTransferManager(static_cast<QObject*>(parent));
}

void QBluetoothTransferManager_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));;
}

void QBluetoothTransferManager_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));;
}

void QBluetoothTransferManager_DestroyQBluetoothTransferManager(void* ptr){
	static_cast<QBluetoothTransferManager*>(ptr)->~QBluetoothTransferManager();
}

