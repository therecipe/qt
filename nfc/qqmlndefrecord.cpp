#include "qqmlndefrecord.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNdefRecord>
#include <QString>
#include <QQmlNdefRecord>
#include "_cgo_export.h"

class MyQQmlNdefRecord: public QQmlNdefRecord {
public:
void Signal_RecordChanged(){callbackQQmlNdefRecordRecordChanged(this->objectName().toUtf8().data());};
void Signal_TypeChanged(){callbackQQmlNdefRecordTypeChanged(this->objectName().toUtf8().data());};
void Signal_TypeNameFormatChanged(){callbackQQmlNdefRecordTypeNameFormatChanged(this->objectName().toUtf8().data());};
};

int QQmlNdefRecord_TypeNameFormat(void* ptr){
	return static_cast<QQmlNdefRecord*>(ptr)->typeNameFormat();
}

void* QQmlNdefRecord_NewQQmlNdefRecord(void* parent){
	return new QQmlNdefRecord(static_cast<QObject*>(parent));
}

void* QQmlNdefRecord_NewQQmlNdefRecord2(void* record, void* parent){
	return new QQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QObject*>(parent));
}

void QQmlNdefRecord_ConnectRecordChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));;
}

void QQmlNdefRecord_DisconnectRecordChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));;
}

void QQmlNdefRecord_SetRecord(void* ptr, void* record){
	static_cast<QQmlNdefRecord*>(ptr)->setRecord(*static_cast<QNdefRecord*>(record));
}

void QQmlNdefRecord_SetType(void* ptr, char* newtype){
	static_cast<QQmlNdefRecord*>(ptr)->setType(QString(newtype));
}

void QQmlNdefRecord_SetTypeNameFormat(void* ptr, int newTypeNameFormat){
	static_cast<QQmlNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QQmlNdefRecord::TypeNameFormat>(newTypeNameFormat));
}

char* QQmlNdefRecord_Type(void* ptr){
	return static_cast<QQmlNdefRecord*>(ptr)->type().toUtf8().data();
}

void QQmlNdefRecord_ConnectTypeChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));;
}

void QQmlNdefRecord_DisconnectTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));;
}

void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

