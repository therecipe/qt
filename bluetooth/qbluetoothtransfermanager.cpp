#include "qbluetoothtransfermanager.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QIODevice>
#include <QBluetoothTransferReply>
#include <QBluetoothTransferRequest>
#include <QString>
#include <QBluetoothTransferManager>
#include "_cgo_export.h"

class MyQBluetoothTransferManager: public QBluetoothTransferManager {
public:
void Signal_Finished(QBluetoothTransferReply * reply){callbackQBluetoothTransferManagerFinished(this->objectName().toUtf8().data(), reply);};
};

QtObjectPtr QBluetoothTransferManager_Put(QtObjectPtr ptr, QtObjectPtr request, QtObjectPtr data){
	return static_cast<QBluetoothTransferManager*>(ptr)->put(*static_cast<QBluetoothTransferRequest*>(request), static_cast<QIODevice*>(data));
}

QtObjectPtr QBluetoothTransferManager_NewQBluetoothTransferManager(QtObjectPtr parent){
	return new QBluetoothTransferManager(static_cast<QObject*>(parent));
}

void QBluetoothTransferManager_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));;
}

void QBluetoothTransferManager_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));;
}

void QBluetoothTransferManager_DestroyQBluetoothTransferManager(QtObjectPtr ptr){
	static_cast<QBluetoothTransferManager*>(ptr)->~QBluetoothTransferManager();
}

