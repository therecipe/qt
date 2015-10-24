#include "qdnstextrecord.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QDnsTextRecord>
#include "_cgo_export.h"

class MyQDnsTextRecord: public QDnsTextRecord {
public:
};

QtObjectPtr QDnsTextRecord_NewQDnsTextRecord(){
	return new QDnsTextRecord();
}

QtObjectPtr QDnsTextRecord_NewQDnsTextRecord2(QtObjectPtr other){
	return new QDnsTextRecord(*static_cast<QDnsTextRecord*>(other));
}

char* QDnsTextRecord_Name(QtObjectPtr ptr){
	return static_cast<QDnsTextRecord*>(ptr)->name().toUtf8().data();
}

void QDnsTextRecord_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDnsTextRecord*>(ptr)->swap(*static_cast<QDnsTextRecord*>(other));
}

void QDnsTextRecord_DestroyQDnsTextRecord(QtObjectPtr ptr){
	static_cast<QDnsTextRecord*>(ptr)->~QDnsTextRecord();
}

