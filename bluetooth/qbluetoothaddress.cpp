#include "qbluetoothaddress.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothAddress>
#include "_cgo_export.h"

class MyQBluetoothAddress: public QBluetoothAddress {
public:
};

void* QBluetoothAddress_NewQBluetoothAddress(){
	return new QBluetoothAddress();
}

void* QBluetoothAddress_NewQBluetoothAddress4(void* other){
	return new QBluetoothAddress(*static_cast<QBluetoothAddress*>(other));
}

void* QBluetoothAddress_NewQBluetoothAddress3(char* address){
	return new QBluetoothAddress(QString(address));
}

void QBluetoothAddress_Clear(void* ptr){
	static_cast<QBluetoothAddress*>(ptr)->clear();
}

int QBluetoothAddress_IsNull(void* ptr){
	return static_cast<QBluetoothAddress*>(ptr)->isNull();
}

char* QBluetoothAddress_ToString(void* ptr){
	return static_cast<QBluetoothAddress*>(ptr)->toString().toUtf8().data();
}

void QBluetoothAddress_DestroyQBluetoothAddress(void* ptr){
	static_cast<QBluetoothAddress*>(ptr)->~QBluetoothAddress();
}

