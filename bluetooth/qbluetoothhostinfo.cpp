#include "qbluetoothhostinfo.h"
#include <QBluetoothAddress>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothHostInfo>
#include "_cgo_export.h"

class MyQBluetoothHostInfo: public QBluetoothHostInfo {
public:
};

QtObjectPtr QBluetoothHostInfo_NewQBluetoothHostInfo(){
	return new QBluetoothHostInfo();
}

QtObjectPtr QBluetoothHostInfo_NewQBluetoothHostInfo2(QtObjectPtr other){
	return new QBluetoothHostInfo(*static_cast<QBluetoothHostInfo*>(other));
}

char* QBluetoothHostInfo_Name(QtObjectPtr ptr){
	return static_cast<QBluetoothHostInfo*>(ptr)->name().toUtf8().data();
}

void QBluetoothHostInfo_SetAddress(QtObjectPtr ptr, QtObjectPtr address){
	static_cast<QBluetoothHostInfo*>(ptr)->setAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothHostInfo_SetName(QtObjectPtr ptr, char* name){
	static_cast<QBluetoothHostInfo*>(ptr)->setName(QString(name));
}

void QBluetoothHostInfo_DestroyQBluetoothHostInfo(QtObjectPtr ptr){
	static_cast<QBluetoothHostInfo*>(ptr)->~QBluetoothHostInfo();
}

