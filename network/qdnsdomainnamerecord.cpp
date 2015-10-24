#include "qdnsdomainnamerecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDnsDomainNameRecord>
#include "_cgo_export.h"

class MyQDnsDomainNameRecord: public QDnsDomainNameRecord {
public:
};

QtObjectPtr QDnsDomainNameRecord_NewQDnsDomainNameRecord(){
	return new QDnsDomainNameRecord();
}

QtObjectPtr QDnsDomainNameRecord_NewQDnsDomainNameRecord2(QtObjectPtr other){
	return new QDnsDomainNameRecord(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Name(QtObjectPtr ptr){
	return static_cast<QDnsDomainNameRecord*>(ptr)->name().toUtf8().data();
}

void QDnsDomainNameRecord_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDnsDomainNameRecord*>(ptr)->swap(*static_cast<QDnsDomainNameRecord*>(other));
}

char* QDnsDomainNameRecord_Value(QtObjectPtr ptr){
	return static_cast<QDnsDomainNameRecord*>(ptr)->value().toUtf8().data();
}

void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(QtObjectPtr ptr){
	static_cast<QDnsDomainNameRecord*>(ptr)->~QDnsDomainNameRecord();
}

