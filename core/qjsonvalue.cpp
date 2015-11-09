#include "qjsonvalue.h"
#include <QModelIndex>
#include <QLatin1String>
#include <QJsonArray>
#include <QJsonObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QJsonValue>
#include "_cgo_export.h"

class MyQJsonValue: public QJsonValue {
public:
};

void* QJsonValue_NewQJsonValue5(void* s){
	return new QJsonValue(*static_cast<QLatin1String*>(s));
}

void* QJsonValue_NewQJsonValue(int ty){
	return new QJsonValue(static_cast<QJsonValue::Type>(ty));
}

void* QJsonValue_NewQJsonValue2(int b){
	return new QJsonValue(b != 0);
}

void* QJsonValue_NewQJsonValue7(void* a){
	return new QJsonValue(*static_cast<QJsonArray*>(a));
}

void* QJsonValue_NewQJsonValue8(void* o){
	return new QJsonValue(*static_cast<QJsonObject*>(o));
}

void* QJsonValue_NewQJsonValue9(void* other){
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

void* QJsonValue_NewQJsonValue4(char* s){
	return new QJsonValue(QString(s));
}

void* QJsonValue_NewQJsonValue6(char* s){
	return new QJsonValue(const_cast<const char*>(s));
}

void* QJsonValue_NewQJsonValue12(int n){
	return new QJsonValue(n);
}

int QJsonValue_IsArray(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isArray();
}

int QJsonValue_IsBool(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isBool();
}

int QJsonValue_IsDouble(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isDouble();
}

int QJsonValue_IsNull(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isNull();
}

int QJsonValue_IsObject(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isObject();
}

int QJsonValue_IsString(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isString();
}

int QJsonValue_IsUndefined(void* ptr){
	return static_cast<QJsonValue*>(ptr)->isUndefined();
}

void* QJsonValue_ToArray2(void* ptr){
	return new QJsonArray(static_cast<QJsonValue*>(ptr)->toArray());
}

void* QJsonValue_ToArray(void* ptr, void* defaultValue){
	return new QJsonArray(static_cast<QJsonValue*>(ptr)->toArray(*static_cast<QJsonArray*>(defaultValue)));
}

int QJsonValue_ToBool(void* ptr, int defaultValue){
	return static_cast<QJsonValue*>(ptr)->toBool(defaultValue != 0);
}

int QJsonValue_ToInt(void* ptr, int defaultValue){
	return static_cast<QJsonValue*>(ptr)->toInt(defaultValue);
}

void* QJsonValue_ToObject2(void* ptr){
	return new QJsonObject(static_cast<QJsonValue*>(ptr)->toObject());
}

void* QJsonValue_ToObject(void* ptr, void* defaultValue){
	return new QJsonObject(static_cast<QJsonValue*>(ptr)->toObject(*static_cast<QJsonObject*>(defaultValue)));
}

char* QJsonValue_ToString(void* ptr, char* defaultValue){
	return static_cast<QJsonValue*>(ptr)->toString(QString(defaultValue)).toUtf8().data();
}

void* QJsonValue_ToVariant(void* ptr){
	return new QVariant(static_cast<QJsonValue*>(ptr)->toVariant());
}

int QJsonValue_Type(void* ptr){
	return static_cast<QJsonValue*>(ptr)->type();
}

void QJsonValue_DestroyQJsonValue(void* ptr){
	static_cast<QJsonValue*>(ptr)->~QJsonValue();
}

