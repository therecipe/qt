#include "qndefrecord.h"
#include <QModelIndex>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNdefRecord>
#include "_cgo_export.h"

class MyQNdefRecord: public QNdefRecord {
public:
};

void* QNdefRecord_NewQNdefRecord(){
	return new QNdefRecord();
}

void* QNdefRecord_NewQNdefRecord2(void* other){
	return new QNdefRecord(*static_cast<QNdefRecord*>(other));
}

void* QNdefRecord_Id(void* ptr){
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->id());
}

int QNdefRecord_IsEmpty(void* ptr){
	return static_cast<QNdefRecord*>(ptr)->isEmpty();
}

void* QNdefRecord_Payload(void* ptr){
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->payload());
}

void QNdefRecord_SetId(void* ptr, void* id){
	static_cast<QNdefRecord*>(ptr)->setId(*static_cast<QByteArray*>(id));
}

void QNdefRecord_SetPayload(void* ptr, void* payload){
	static_cast<QNdefRecord*>(ptr)->setPayload(*static_cast<QByteArray*>(payload));
}

void QNdefRecord_SetType(void* ptr, void* ty){
	static_cast<QNdefRecord*>(ptr)->setType(*static_cast<QByteArray*>(ty));
}

void QNdefRecord_SetTypeNameFormat(void* ptr, int typeNameFormat){
	static_cast<QNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat));
}

void* QNdefRecord_Type(void* ptr){
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->type());
}

int QNdefRecord_TypeNameFormat(void* ptr){
	return static_cast<QNdefRecord*>(ptr)->typeNameFormat();
}

void QNdefRecord_DestroyQNdefRecord(void* ptr){
	static_cast<QNdefRecord*>(ptr)->~QNdefRecord();
}

