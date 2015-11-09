#include "qdnshostaddressrecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDnsHostAddressRecord>
#include "_cgo_export.h"

class MyQDnsHostAddressRecord: public QDnsHostAddressRecord {
public:
};

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord(){
	return new QDnsHostAddressRecord();
}

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord2(void* other){
	return new QDnsHostAddressRecord(*static_cast<QDnsHostAddressRecord*>(other));
}

char* QDnsHostAddressRecord_Name(void* ptr){
	return static_cast<QDnsHostAddressRecord*>(ptr)->name().toUtf8().data();
}

void QDnsHostAddressRecord_Swap(void* ptr, void* other){
	static_cast<QDnsHostAddressRecord*>(ptr)->swap(*static_cast<QDnsHostAddressRecord*>(other));
}

void QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(void* ptr){
	static_cast<QDnsHostAddressRecord*>(ptr)->~QDnsHostAddressRecord();
}

