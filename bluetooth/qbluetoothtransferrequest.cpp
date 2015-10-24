#include "qbluetoothtransferrequest.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothAddress>
#include <QBluetoothTransferRequest>
#include "_cgo_export.h"

class MyQBluetoothTransferRequest: public QBluetoothTransferRequest {
public:
};

QtObjectPtr QBluetoothTransferRequest_NewQBluetoothTransferRequest(QtObjectPtr address){
	return new QBluetoothTransferRequest(*static_cast<QBluetoothAddress*>(address));
}

QtObjectPtr QBluetoothTransferRequest_NewQBluetoothTransferRequest2(QtObjectPtr other){
	return new QBluetoothTransferRequest(*static_cast<QBluetoothTransferRequest*>(other));
}

char* QBluetoothTransferRequest_Attribute(QtObjectPtr ptr, int code, char* defaultValue){
	return static_cast<QBluetoothTransferRequest*>(ptr)->attribute(static_cast<QBluetoothTransferRequest::Attribute>(code), QVariant(defaultValue)).toString().toUtf8().data();
}

void QBluetoothTransferRequest_SetAttribute(QtObjectPtr ptr, int code, char* value){
	static_cast<QBluetoothTransferRequest*>(ptr)->setAttribute(static_cast<QBluetoothTransferRequest::Attribute>(code), QVariant(value));
}

void QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(QtObjectPtr ptr){
	static_cast<QBluetoothTransferRequest*>(ptr)->~QBluetoothTransferRequest();
}

