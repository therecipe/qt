#include "qbluetoothtransferrequest.h"
#include <QModelIndex>
#include <QBluetoothAddress>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QBluetoothTransferRequest>
#include "_cgo_export.h"

class MyQBluetoothTransferRequest: public QBluetoothTransferRequest {
public:
};

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest(void* address){
	return new QBluetoothTransferRequest(*static_cast<QBluetoothAddress*>(address));
}

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest2(void* other){
	return new QBluetoothTransferRequest(*static_cast<QBluetoothTransferRequest*>(other));
}

void* QBluetoothTransferRequest_Attribute(void* ptr, int code, void* defaultValue){
	return new QVariant(static_cast<QBluetoothTransferRequest*>(ptr)->attribute(static_cast<QBluetoothTransferRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

void QBluetoothTransferRequest_SetAttribute(void* ptr, int code, void* value){
	static_cast<QBluetoothTransferRequest*>(ptr)->setAttribute(static_cast<QBluetoothTransferRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(void* ptr){
	static_cast<QBluetoothTransferRequest*>(ptr)->~QBluetoothTransferRequest();
}

