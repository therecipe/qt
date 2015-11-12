#include "qdnstextrecord.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QDnsTextRecord>
#include "_cgo_export.h"

class MyQDnsTextRecord: public QDnsTextRecord {
public:
};

void* QDnsTextRecord_NewQDnsTextRecord(){
	return new QDnsTextRecord();
}

void* QDnsTextRecord_NewQDnsTextRecord2(void* other){
	return new QDnsTextRecord(*static_cast<QDnsTextRecord*>(other));
}

char* QDnsTextRecord_Name(void* ptr){
	return static_cast<QDnsTextRecord*>(ptr)->name().toUtf8().data();
}

void QDnsTextRecord_Swap(void* ptr, void* other){
	static_cast<QDnsTextRecord*>(ptr)->swap(*static_cast<QDnsTextRecord*>(other));
}

void QDnsTextRecord_DestroyQDnsTextRecord(void* ptr){
	static_cast<QDnsTextRecord*>(ptr)->~QDnsTextRecord();
}

