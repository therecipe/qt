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

QtObjectPtr QDnsServiceRecord_NewQDnsServiceRecord(){
	return new QDnsServiceRecord();
}

QtObjectPtr QDnsServiceRecord_NewQDnsServiceRecord2(QtObjectPtr other){
	return new QDnsServiceRecord(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Name(QtObjectPtr ptr){
	return static_cast<QDnsServiceRecord*>(ptr)->name().toUtf8().data();
}

void QDnsServiceRecord_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDnsServiceRecord*>(ptr)->swap(*static_cast<QDnsServiceRecord*>(other));
}

char* QDnsServiceRecord_Target(QtObjectPtr ptr){
	return static_cast<QDnsServiceRecord*>(ptr)->target().toUtf8().data();
}

void QDnsServiceRecord_DestroyQDnsServiceRecord(QtObjectPtr ptr){
	static_cast<QDnsServiceRecord*>(ptr)->~QDnsServiceRecord();
}

