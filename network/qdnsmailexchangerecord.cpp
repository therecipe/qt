#include "qdnsmailexchangerecord.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDnsMailExchangeRecord>
#include "_cgo_export.h"

class MyQDnsMailExchangeRecord: public QDnsMailExchangeRecord {
public:
};

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord(){
	return new QDnsMailExchangeRecord();
}

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(void* other){
	return new QDnsMailExchangeRecord(*static_cast<QDnsMailExchangeRecord*>(other));
}

char* QDnsMailExchangeRecord_Exchange(void* ptr){
	return static_cast<QDnsMailExchangeRecord*>(ptr)->exchange().toUtf8().data();
}

char* QDnsMailExchangeRecord_Name(void* ptr){
	return static_cast<QDnsMailExchangeRecord*>(ptr)->name().toUtf8().data();
}

void QDnsMailExchangeRecord_Swap(void* ptr, void* other){
	static_cast<QDnsMailExchangeRecord*>(ptr)->swap(*static_cast<QDnsMailExchangeRecord*>(other));
}

void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(void* ptr){
	static_cast<QDnsMailExchangeRecord*>(ptr)->~QDnsMailExchangeRecord();
}

