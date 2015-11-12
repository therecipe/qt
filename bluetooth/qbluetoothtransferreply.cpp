#include "qbluetoothtransferreply.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QBluetoothTransferReply>
#include "_cgo_export.h"

class MyQBluetoothTransferReply: public QBluetoothTransferReply {
public:
void Signal_Finished(QBluetoothTransferReply * reply){callbackQBluetoothTransferReplyFinished(this->objectName().toUtf8().data(), reply);};
};

void QBluetoothTransferReply_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferReply*>(ptr), "abort");
}

int QBluetoothTransferReply_Error(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->error();
}

char* QBluetoothTransferReply_ErrorString(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->errorString().toUtf8().data();
}

void QBluetoothTransferReply_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));;
}

void QBluetoothTransferReply_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));;
}

int QBluetoothTransferReply_IsFinished(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isFinished();
}

int QBluetoothTransferReply_IsRunning(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isRunning();
}

void* QBluetoothTransferReply_Manager(void* ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->manager();
}

void QBluetoothTransferReply_DestroyQBluetoothTransferReply(void* ptr){
	static_cast<QBluetoothTransferReply*>(ptr)->~QBluetoothTransferReply();
}

