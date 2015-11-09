#include "qnetworkaddressentry.h"
#include <QUrl>
#include <QModelIndex>
#include <QHostAddress>
#include <QString>
#include <QVariant>
#include <QNetworkAddressEntry>
#include "_cgo_export.h"

class MyQNetworkAddressEntry: public QNetworkAddressEntry {
public:
};

void* QNetworkAddressEntry_NewQNetworkAddressEntry(){
	return new QNetworkAddressEntry();
}

void* QNetworkAddressEntry_NewQNetworkAddressEntry2(void* other){
	return new QNetworkAddressEntry(*static_cast<QNetworkAddressEntry*>(other));
}

int QNetworkAddressEntry_PrefixLength(void* ptr){
	return static_cast<QNetworkAddressEntry*>(ptr)->prefixLength();
}

void QNetworkAddressEntry_SetBroadcast(void* ptr, void* newBroadcast){
	static_cast<QNetworkAddressEntry*>(ptr)->setBroadcast(*static_cast<QHostAddress*>(newBroadcast));
}

void QNetworkAddressEntry_SetIp(void* ptr, void* newIp){
	static_cast<QNetworkAddressEntry*>(ptr)->setIp(*static_cast<QHostAddress*>(newIp));
}

void QNetworkAddressEntry_SetNetmask(void* ptr, void* newNetmask){
	static_cast<QNetworkAddressEntry*>(ptr)->setNetmask(*static_cast<QHostAddress*>(newNetmask));
}

void QNetworkAddressEntry_SetPrefixLength(void* ptr, int length){
	static_cast<QNetworkAddressEntry*>(ptr)->setPrefixLength(length);
}

void QNetworkAddressEntry_Swap(void* ptr, void* other){
	static_cast<QNetworkAddressEntry*>(ptr)->swap(*static_cast<QNetworkAddressEntry*>(other));
}

void QNetworkAddressEntry_DestroyQNetworkAddressEntry(void* ptr){
	static_cast<QNetworkAddressEntry*>(ptr)->~QNetworkAddressEntry();
}

