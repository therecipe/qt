#include "qndefnfctextrecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefRecord>
#include <QNdefNfcTextRecord>
#include "_cgo_export.h"

class MyQNdefNfcTextRecord: public QNdefNfcTextRecord {
public:
};

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord(){
	return new QNdefNfcTextRecord();
}

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord2(void* other){
	return new QNdefNfcTextRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefNfcTextRecord_Encoding(void* ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->encoding();
}

char* QNdefNfcTextRecord_Locale(void* ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->locale().toUtf8().data();
}

void QNdefNfcTextRecord_SetEncoding(void* ptr, int encoding){
	static_cast<QNdefNfcTextRecord*>(ptr)->setEncoding(static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

void QNdefNfcTextRecord_SetLocale(void* ptr, char* locale){
	static_cast<QNdefNfcTextRecord*>(ptr)->setLocale(QString(locale));
}

void QNdefNfcTextRecord_SetText(void* ptr, char* text){
	static_cast<QNdefNfcTextRecord*>(ptr)->setText(QString(text));
}

char* QNdefNfcTextRecord_Text(void* ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->text().toUtf8().data();
}

