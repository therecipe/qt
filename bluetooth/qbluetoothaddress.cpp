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

QtObjectPtr QBluetoothAddress_NewQBluetoothAddress(){
	return new QBluetoothAddress();
}

QtObjectPtr QBluetoothAddress_NewQBluetoothAddress4(QtObjectPtr other){
	return new QBluetoothAddress(*static_cast<QBluetoothAddress*>(other));
}

QtObjectPtr QBluetoothAddress_NewQBluetoothAddress3(char* address){
	return new QBluetoothAddress(QString(address));
}

void QBluetoothAddress_Clear(QtObjectPtr ptr){
	static_cast<QBluetoothAddress*>(ptr)->clear();
}

int QBluetoothAddress_IsNull(QtObjectPtr ptr){
	return static_cast<QBluetoothAddress*>(ptr)->isNull();
}

char* QBluetoothAddress_ToString(QtObjectPtr ptr){
	return static_cast<QBluetoothAddress*>(ptr)->toString().toUtf8().data();
}

void QBluetoothAddress_DestroyQBluetoothAddress(QtObjectPtr ptr){
	static_cast<QBluetoothAddress*>(ptr)->~QBluetoothAddress();
}

