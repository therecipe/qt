#include "qdnsdomainnamerecord.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDnsDomainNameRecord>
#include "_cgo_export.h"

class MyQDnsDomainNameRecord: public QDnsDomainNameRecord {
public:
};

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord(){
	return new QDnsDomainNameRecord();
}

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord2(void* other){
	return new QDnsDomainNameRecord(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Name(void* ptr){
	return static_cast<QDnsDomainNameRecord*>(ptr)->name().toUtf8().data();
}

void QDnsDomainNameRecord_Swap(void* ptr, void* other){
	static_cast<QDnsDomainNameRecord*>(ptr)->swap(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Value(void* ptr){
	return static_cast<QDnsDomainNameRecord*>(ptr)->value().toUtf8().data();
}

void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(void* ptr){
	static_cast<QDnsDomainNameRecord*>(ptr)->~QDnsDomainNameRecord();
}

