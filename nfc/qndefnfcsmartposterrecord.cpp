#include "qndefnfcsmartposterrecord.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefNfcUriRecord>
#include <QNdefNfcTextRecord>
#include <QByteArray>
#include <QNdefRecord>
#include <QString>
#include <QNdefNfcSmartPosterRecord>
#include "_cgo_export.h"

class MyQNdefNfcSmartPosterRecord: public QNdefNfcSmartPosterRecord {
public:
};

QtObjectPtr QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord(){
	return new QNdefNfcSmartPosterRecord();
}

QtObjectPtr QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(QtObjectPtr other){
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefNfcSmartPosterRecord*>(other));
}

QtObjectPtr QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(QtObjectPtr other){
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefNfcSmartPosterRecord_Action(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->action();
}

void QNdefNfcSmartPosterRecord_AddIcon2(QtObjectPtr ptr, QtObjectPtr ty, QtObjectPtr data){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addIcon(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(data));
}

int QNdefNfcSmartPosterRecord_AddTitle(QtObjectPtr ptr, QtObjectPtr text){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

int QNdefNfcSmartPosterRecord_AddTitle2(QtObjectPtr ptr, char* text, char* locale, int encoding){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(QString(text), QString(locale), static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

int QNdefNfcSmartPosterRecord_HasAction(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasAction();
}

int QNdefNfcSmartPosterRecord_HasIcon(QtObjectPtr ptr, QtObjectPtr mimetype){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasIcon(*static_cast<QByteArray*>(mimetype));
}

int QNdefNfcSmartPosterRecord_HasSize(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasSize();
}

int QNdefNfcSmartPosterRecord_HasTitle(QtObjectPtr ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTitle(QString(locale));
}

int QNdefNfcSmartPosterRecord_HasTypeInfo(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTypeInfo();
}

int QNdefNfcSmartPosterRecord_IconCount(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->iconCount();
}

int QNdefNfcSmartPosterRecord_RemoveIcon2(QtObjectPtr ptr, QtObjectPtr ty){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeIcon(*static_cast<QByteArray*>(ty));
}

int QNdefNfcSmartPosterRecord_RemoveTitle(QtObjectPtr ptr, QtObjectPtr text){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

int QNdefNfcSmartPosterRecord_RemoveTitle2(QtObjectPtr ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(QString(locale));
}

void QNdefNfcSmartPosterRecord_SetAction(QtObjectPtr ptr, int act){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setAction(static_cast<QNdefNfcSmartPosterRecord::Action>(act));
}

void QNdefNfcSmartPosterRecord_SetTypeInfo(QtObjectPtr ptr, QtObjectPtr ty){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setTypeInfo(*static_cast<QByteArray*>(ty));
}

void QNdefNfcSmartPosterRecord_SetUri(QtObjectPtr ptr, QtObjectPtr url){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QNdefNfcUriRecord*>(url));
}

void QNdefNfcSmartPosterRecord_SetUri2(QtObjectPtr ptr, char* url){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(QUrl(QString(url)));
}

char* QNdefNfcSmartPosterRecord_Title(QtObjectPtr ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->title(QString(locale)).toUtf8().data();
}

int QNdefNfcSmartPosterRecord_TitleCount(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleCount();
}

char* QNdefNfcSmartPosterRecord_Uri(QtObjectPtr ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->uri().toString().toUtf8().data();
}

void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(QtObjectPtr ptr){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->~QNdefNfcSmartPosterRecord();
}

