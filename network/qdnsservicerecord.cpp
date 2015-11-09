#include "qdnsservicerecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDnsServiceRecord>
#include "_cgo_export.h"

class MyQDnsServiceRecord: public QDnsServiceRecord {
public:
};

void* QDnsServiceRecord_NewQDnsServiceRecord(){
	return new QDnsServiceRecord();
}

void* QDnsServiceRecord_NewQDnsServiceRecord2(void* other){
	return new QDnsServiceRecord(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Name(void* ptr){
	return static_cast<QDnsServiceRecord*>(ptr)->name().toUtf8().data();
}

void QDnsServiceRecord_Swap(void* ptr, void* other){
	static_cast<QDnsServiceRecord*>(ptr)->swap(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Target(void* ptr){
	return static_cast<QDnsServiceRecord*>(ptr)->target().toUtf8().data();
}

void QDnsServiceRecord_DestroyQDnsServiceRecord(void* ptr){
	static_cast<QDnsServiceRecord*>(ptr)->~QDnsServiceRecord();
}

