#include "qdnsmailexchangerecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDnsMailExchangeRecord>
#include "_cgo_export.h"

class MyQDnsMailExchangeRecord: public QDnsMailExchangeRecord {
public:
};

QtObjectPtr QDnsMailExchangeRecord_NewQDnsMailExchangeRecord(){
	return new QDnsMailExchangeRecord();
}

QtObjectPtr QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(QtObjectPtr other){
	return new QDnsMailExchangeRecord(*static_cast<QDnsMailExchangeRecord*>(other));
}

char* QDnsMailExchangeRecord_Exchange(QtObjectPtr ptr){
	return static_cast<QDnsMailExchangeRecord*>(ptr)->exchange().toUtf8().data();
}

char* QDnsMailExchangeRecord_Name(QtObjectPtr ptr){
	return static_cast<QDnsMailExchangeRecord*>(ptr)->name().toUtf8().data();
}

void QDnsMailExchangeRecord_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDnsMailExchangeRecord*>(ptr)->swap(*static_cast<QDnsMailExchangeRecord*>(other));
}

void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(QtObjectPtr ptr){
	static_cast<QDnsMailExchangeRecord*>(ptr)->~QDnsMailExchangeRecord();
}

