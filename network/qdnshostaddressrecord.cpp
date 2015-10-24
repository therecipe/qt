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

QtObjectPtr QDnsHostAddressRecord_NewQDnsHostAddressRecord(){
	return new QDnsHostAddressRecord();
}

QtObjectPtr QDnsHostAddressRecord_NewQDnsHostAddressRecord2(QtObjectPtr other){
	return new QDnsHostAddressRecord(*static_cast<QDnsHostAddressRecord*>(other));
}

char* QDnsHostAddressRecord_Name(QtObjectPtr ptr){
	return static_cast<QDnsHostAddressRecord*>(ptr)->name().toUtf8().data();
}

void QDnsHostAddressRecord_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDnsHostAddressRecord*>(ptr)->swap(*static_cast<QDnsHostAddressRecord*>(other));
}

void QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(QtObjectPtr ptr){
	static_cast<QDnsHostAddressRecord*>(ptr)->~QDnsHostAddressRecord();
}

