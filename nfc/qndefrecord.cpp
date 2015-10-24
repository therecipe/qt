#include "qndefrecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QNdefRecord>
#include "_cgo_export.h"

class MyQNdefRecord: public QNdefRecord {
public:
};

QtObjectPtr QNdefRecord_NewQNdefRecord(){
	return new QNdefRecord();
}

QtObjectPtr QNdefRecord_NewQNdefRecord2(QtObjectPtr other){
	return new QNdefRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefRecord_IsEmpty(QtObjectPtr ptr){
	return static_cast<QNdefRecord*>(ptr)->isEmpty();
}

void QNdefRecord_SetId(QtObjectPtr ptr, QtObjectPtr id){
	static_cast<QNdefRecord*>(ptr)->setId(*static_cast<QByteArray*>(id));
}

void QNdefRecord_SetPayload(QtObjectPtr ptr, QtObjectPtr payload){
	static_cast<QNdefRecord*>(ptr)->setPayload(*static_cast<QByteArray*>(payload));
}

void QNdefRecord_SetType(QtObjectPtr ptr, QtObjectPtr ty){
	static_cast<QNdefRecord*>(ptr)->setType(*static_cast<QByteArray*>(ty));
}

void QNdefRecord_SetTypeNameFormat(QtObjectPtr ptr, int typeNameFormat){
	static_cast<QNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat));
}

int QNdefRecord_TypeNameFormat(QtObjectPtr ptr){
	return static_cast<QNdefRecord*>(ptr)->typeNameFormat();
}

void QNdefRecord_DestroyQNdefRecord(QtObjectPtr ptr){
	static_cast<QNdefRecord*>(ptr)->~QNdefRecord();
}

