#include "qqmlndefrecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QNdefRecord>
#include <QQmlNdefRecord>
#include "_cgo_export.h"

class MyQQmlNdefRecord: public QQmlNdefRecord {
public:
void Signal_RecordChanged(){callbackQQmlNdefRecordRecordChanged(this->objectName().toUtf8().data());};
void Signal_TypeChanged(){callbackQQmlNdefRecordTypeChanged(this->objectName().toUtf8().data());};
void Signal_TypeNameFormatChanged(){callbackQQmlNdefRecordTypeNameFormatChanged(this->objectName().toUtf8().data());};
};

int QQmlNdefRecord_TypeNameFormat(QtObjectPtr ptr){
	return static_cast<QQmlNdefRecord*>(ptr)->typeNameFormat();
}

QtObjectPtr QQmlNdefRecord_NewQQmlNdefRecord(QtObjectPtr parent){
	return new QQmlNdefRecord(static_cast<QObject*>(parent));
}

QtObjectPtr QQmlNdefRecord_NewQQmlNdefRecord2(QtObjectPtr record, QtObjectPtr parent){
	return new QQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QObject*>(parent));
}

void QQmlNdefRecord_ConnectRecordChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));;
}

void QQmlNdefRecord_DisconnectRecordChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));;
}

void QQmlNdefRecord_SetRecord(QtObjectPtr ptr, QtObjectPtr record){
	static_cast<QQmlNdefRecord*>(ptr)->setRecord(*static_cast<QNdefRecord*>(record));
}

void QQmlNdefRecord_SetType(QtObjectPtr ptr, char* newtype){
	static_cast<QQmlNdefRecord*>(ptr)->setType(QString(newtype));
}

void QQmlNdefRecord_SetTypeNameFormat(QtObjectPtr ptr, int newTypeNameFormat){
	static_cast<QQmlNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QQmlNdefRecord::TypeNameFormat>(newTypeNameFormat));
}

char* QQmlNdefRecord_Type(QtObjectPtr ptr){
	return static_cast<QQmlNdefRecord*>(ptr)->type().toUtf8().data();
}

void QQmlNdefRecord_ConnectTypeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));;
}

void QQmlNdefRecord_DisconnectTypeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));;
}

void QQmlNdefRecord_ConnectTypeNameFormatChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

void QQmlNdefRecord_DisconnectTypeNameFormatChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

