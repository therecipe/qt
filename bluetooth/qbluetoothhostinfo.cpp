#include "qbluetoothhostinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothAddress>
#include <QBluetoothHostInfo>
#include "_cgo_export.h"

class MyQBluetoothHostInfo: public QBluetoothHostInfo {
public:
};

void* QBluetoothHostInfo_NewQBluetoothHostInfo(){
	return new QBluetoothHostInfo();
}

void* QBluetoothHostInfo_NewQBluetoothHostInfo2(void* other){
	return new QBluetoothHostInfo(*static_cast<QBluetoothHostInfo*>(other));
}

char* QBluetoothHostInfo_Name(void* ptr){
	return static_cast<QBluetoothHostInfo*>(ptr)->name().toUtf8().data();
}

void QBluetoothHostInfo_SetAddress(void* ptr, void* address){
	static_cast<QBluetoothHostInfo*>(ptr)->setAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothHostInfo_SetName(void* ptr, char* name){
	static_cast<QBluetoothHostInfo*>(ptr)->setName(QString(name));
}

void QBluetoothHostInfo_DestroyQBluetoothHostInfo(void* ptr){
	static_cast<QBluetoothHostInfo*>(ptr)->~QBluetoothHostInfo();
}

