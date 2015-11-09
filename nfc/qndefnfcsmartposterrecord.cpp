#include "qndefnfcsmartposterrecord.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefNfcUriRecord>
#include <QNdefRecord>
#include <QNdefNfcTextRecord>
#include <QByteArray>
#include <QString>
#include <QNdefNfcSmartPosterRecord>
#include "_cgo_export.h"

class MyQNdefNfcSmartPosterRecord: public QNdefNfcSmartPosterRecord {
public:
};

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord(){
	return new QNdefNfcSmartPosterRecord();
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(void* other){
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefNfcSmartPosterRecord*>(other));
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(void* other){
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefNfcSmartPosterRecord_Action(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->action();
}

void QNdefNfcSmartPosterRecord_AddIcon2(void* ptr, void* ty, void* data){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addIcon(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(data));
}

int QNdefNfcSmartPosterRecord_AddTitle(void* ptr, void* text){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

int QNdefNfcSmartPosterRecord_AddTitle2(void* ptr, char* text, char* locale, int encoding){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(QString(text), QString(locale), static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

int QNdefNfcSmartPosterRecord_HasAction(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasAction();
}

int QNdefNfcSmartPosterRecord_HasIcon(void* ptr, void* mimetype){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasIcon(*static_cast<QByteArray*>(mimetype));
}

int QNdefNfcSmartPosterRecord_HasSize(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasSize();
}

int QNdefNfcSmartPosterRecord_HasTitle(void* ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTitle(QString(locale));
}

int QNdefNfcSmartPosterRecord_HasTypeInfo(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTypeInfo();
}

void* QNdefNfcSmartPosterRecord_Icon(void* ptr, void* mimetype){
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->icon(*static_cast<QByteArray*>(mimetype)));
}

int QNdefNfcSmartPosterRecord_IconCount(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->iconCount();
}

int QNdefNfcSmartPosterRecord_RemoveIcon2(void* ptr, void* ty){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeIcon(*static_cast<QByteArray*>(ty));
}

int QNdefNfcSmartPosterRecord_RemoveTitle(void* ptr, void* text){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

int QNdefNfcSmartPosterRecord_RemoveTitle2(void* ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(QString(locale));
}

void QNdefNfcSmartPosterRecord_SetAction(void* ptr, int act){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setAction(static_cast<QNdefNfcSmartPosterRecord::Action>(act));
}

void QNdefNfcSmartPosterRecord_SetTypeInfo(void* ptr, void* ty){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setTypeInfo(*static_cast<QByteArray*>(ty));
}

void QNdefNfcSmartPosterRecord_SetUri(void* ptr, void* url){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QNdefNfcUriRecord*>(url));
}

void QNdefNfcSmartPosterRecord_SetUri2(void* ptr, void* url){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QUrl*>(url));
}

char* QNdefNfcSmartPosterRecord_Title(void* ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->title(QString(locale)).toUtf8().data();
}

int QNdefNfcSmartPosterRecord_TitleCount(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleCount();
}

void* QNdefNfcSmartPosterRecord_TypeInfo(void* ptr){
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->typeInfo());
}

void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->~QNdefNfcSmartPosterRecord();
}

