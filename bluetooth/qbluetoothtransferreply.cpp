#include "qbluetoothtransferreply.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QBluetoothTransferReply>
#include "_cgo_export.h"

class MyQBluetoothTransferReply: public QBluetoothTransferReply {
public:
void Signal_Finished(QBluetoothTransferReply * reply){callbackQBluetoothTransferReplyFinished(this->objectName().toUtf8().data(), reply);};
};

void QBluetoothTransferReply_Abort(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferReply*>(ptr), "abort");
}

int QBluetoothTransferReply_Error(QtObjectPtr ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->error();
}

char* QBluetoothTransferReply_ErrorString(QtObjectPtr ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->errorString().toUtf8().data();
}

void QBluetoothTransferReply_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));;
}

void QBluetoothTransferReply_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));;
}

int QBluetoothTransferReply_IsFinished(QtObjectPtr ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isFinished();
}

int QBluetoothTransferReply_IsRunning(QtObjectPtr ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->isRunning();
}

QtObjectPtr QBluetoothTransferReply_Manager(QtObjectPtr ptr){
	return static_cast<QBluetoothTransferReply*>(ptr)->manager();
}

void QBluetoothTransferReply_DestroyQBluetoothTransferReply(QtObjectPtr ptr){
	static_cast<QBluetoothTransferReply*>(ptr)->~QBluetoothTransferReply();
}

