#include "quuid.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QUuid>
#include "_cgo_export.h"

class MyQUuid: public QUuid {
public:
};

int QUuid_Variant(void* ptr){
	return static_cast<QUuid*>(ptr)->variant();
}

int QUuid_Version(void* ptr){
	return static_cast<QUuid*>(ptr)->version();
}

void* QUuid_NewQUuid(){
	return new QUuid();
}

void* QUuid_NewQUuid5(void* text){
	return new QUuid(*static_cast<QByteArray*>(text));
}

void* QUuid_NewQUuid3(char* text){
	return new QUuid(QString(text));
}

int QUuid_IsNull(void* ptr){
	return static_cast<QUuid*>(ptr)->isNull();
}

void* QUuid_ToByteArray(void* ptr){
	return new QByteArray(static_cast<QUuid*>(ptr)->toByteArray());
}

void* QUuid_ToRfc4122(void* ptr){
	return new QByteArray(static_cast<QUuid*>(ptr)->toRfc4122());
}

char* QUuid_ToString(void* ptr){
	return static_cast<QUuid*>(ptr)->toString().toUtf8().data();
}

