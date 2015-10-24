#include "qnetworkaddressentry.h"
#include <QModelIndex>
#include <QHostAddress>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNetworkAddressEntry>
#include "_cgo_export.h"

class MyQNetworkAddressEntry: public QNetworkAddressEntry {
public:
};

QtObjectPtr QNetworkAddressEntry_NewQNetworkAddressEntry(){
	return new QNetworkAddressEntry();
}

QtObjectPtr QNetworkAddressEntry_NewQNetworkAddressEntry2(QtObjectPtr other){
	return new QNetworkAddressEntry(*static_cast<QNetworkAddressEntry*>(other));
}

int QNetworkAddressEntry_PrefixLength(QtObjectPtr ptr){
	return static_cast<QNetworkAddressEntry*>(ptr)->prefixLength();
}

void QNetworkAddressEntry_SetBroadcast(QtObjectPtr ptr, QtObjectPtr newBroadcast){
	static_cast<QNetworkAddressEntry*>(ptr)->setBroadcast(*static_cast<QHostAddress*>(newBroadcast));
}

void QNetworkAddressEntry_SetIp(QtObjectPtr ptr, QtObjectPtr newIp){
	static_cast<QNetworkAddressEntry*>(ptr)->setIp(*static_cast<QHostAddress*>(newIp));
}

void QNetworkAddressEntry_SetNetmask(QtObjectPtr ptr, QtObjectPtr newNetmask){
	static_cast<QNetworkAddressEntry*>(ptr)->setNetmask(*static_cast<QHostAddress*>(newNetmask));
}

void QNetworkAddressEntry_SetPrefixLength(QtObjectPtr ptr, int length){
	static_cast<QNetworkAddressEntry*>(ptr)->setPrefixLength(length);
}

void QNetworkAddressEntry_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkAddressEntry*>(ptr)->swap(*static_cast<QNetworkAddressEntry*>(other));
}

void QNetworkAddressEntry_DestroyQNetworkAddressEntry(QtObjectPtr ptr){
	static_cast<QNetworkAddressEntry*>(ptr)->~QNetworkAddressEntry();
}

