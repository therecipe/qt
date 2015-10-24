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

QtObjectPtr QNdefNfcTextRecord_NewQNdefNfcTextRecord(){
	return new QNdefNfcTextRecord();
}

QtObjectPtr QNdefNfcTextRecord_NewQNdefNfcTextRecord2(QtObjectPtr other){
	return new QNdefNfcTextRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefNfcTextRecord_Encoding(QtObjectPtr ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->encoding();
}

char* QNdefNfcTextRecord_Locale(QtObjectPtr ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->locale().toUtf8().data();
}

void QNdefNfcTextRecord_SetEncoding(QtObjectPtr ptr, int encoding){
	static_cast<QNdefNfcTextRecord*>(ptr)->setEncoding(static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

void QNdefNfcTextRecord_SetLocale(QtObjectPtr ptr, char* locale){
	static_cast<QNdefNfcTextRecord*>(ptr)->setLocale(QString(locale));
}

void QNdefNfcTextRecord_SetText(QtObjectPtr ptr, char* text){
	static_cast<QNdefNfcTextRecord*>(ptr)->setText(QString(text));
}

char* QNdefNfcTextRecord_Text(QtObjectPtr ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->text().toUtf8().data();
}

