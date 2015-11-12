#include "qjsvalue.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include <QDate>
#include <QLatin1String>
#include <QJSValue>
#include "_cgo_export.h"

class MyQJSValue: public QJSValue {
public:
};

void* QJSValue_NewQJSValue3(void* other){
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue(int value){
	return new QJSValue(static_cast<QJSValue::SpecialValue>(value));
}

void* QJSValue_NewQJSValue4(int value){
	return new QJSValue(value != 0);
}

void* QJSValue_NewQJSValue2(void* other){
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue9(void* value){
	return new QJSValue(*static_cast<QLatin1String*>(value));
}

void* QJSValue_NewQJSValue8(char* value){
	return new QJSValue(QString(value));
}

void* QJSValue_NewQJSValue10(char* value){
	return new QJSValue(const_cast<const char*>(value));
}

void* QJSValue_NewQJSValue5(int value){
	return new QJSValue(value);
}

int QJSValue_DeleteProperty(void* ptr, char* name){
	return static_cast<QJSValue*>(ptr)->deleteProperty(QString(name));
}

int QJSValue_Equals(void* ptr, void* other){
	return static_cast<QJSValue*>(ptr)->equals(*static_cast<QJSValue*>(other));
}

int QJSValue_HasOwnProperty(void* ptr, char* name){
	return static_cast<QJSValue*>(ptr)->hasOwnProperty(QString(name));
}

int QJSValue_HasProperty(void* ptr, char* name){
	return static_cast<QJSValue*>(ptr)->hasProperty(QString(name));
}

int QJSValue_IsArray(void* ptr){
	return static_cast<QJSValue*>(ptr)->isArray();
}

int QJSValue_IsBool(void* ptr){
	return static_cast<QJSValue*>(ptr)->isBool();
}

int QJSValue_IsCallable(void* ptr){
	return static_cast<QJSValue*>(ptr)->isCallable();
}

int QJSValue_IsDate(void* ptr){
	return static_cast<QJSValue*>(ptr)->isDate();
}

int QJSValue_IsError(void* ptr){
	return static_cast<QJSValue*>(ptr)->isError();
}

int QJSValue_IsNull(void* ptr){
	return static_cast<QJSValue*>(ptr)->isNull();
}

int QJSValue_IsNumber(void* ptr){
	return static_cast<QJSValue*>(ptr)->isNumber();
}

int QJSValue_IsObject(void* ptr){
	return static_cast<QJSValue*>(ptr)->isObject();
}

int QJSValue_IsQObject(void* ptr){
	return static_cast<QJSValue*>(ptr)->isQObject();
}

int QJSValue_IsRegExp(void* ptr){
	return static_cast<QJSValue*>(ptr)->isRegExp();
}

int QJSValue_IsString(void* ptr){
	return static_cast<QJSValue*>(ptr)->isString();
}

int QJSValue_IsUndefined(void* ptr){
	return static_cast<QJSValue*>(ptr)->isUndefined();
}

int QJSValue_IsVariant(void* ptr){
	return static_cast<QJSValue*>(ptr)->isVariant();
}

void* QJSValue_Property(void* ptr, char* name){
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(QString(name)));
}

void* QJSValue_Prototype(void* ptr){
	return new QJSValue(static_cast<QJSValue*>(ptr)->prototype());
}

void QJSValue_SetProperty(void* ptr, char* name, void* value){
	static_cast<QJSValue*>(ptr)->setProperty(QString(name), *static_cast<QJSValue*>(value));
}

void QJSValue_SetPrototype(void* ptr, void* prototype){
	static_cast<QJSValue*>(ptr)->setPrototype(*static_cast<QJSValue*>(prototype));
}

int QJSValue_StrictlyEquals(void* ptr, void* other){
	return static_cast<QJSValue*>(ptr)->strictlyEquals(*static_cast<QJSValue*>(other));
}

int QJSValue_ToBool(void* ptr){
	return static_cast<QJSValue*>(ptr)->toBool();
}

void* QJSValue_ToDateTime(void* ptr){
	return new QDateTime(static_cast<QJSValue*>(ptr)->toDateTime());
}

void* QJSValue_ToQObject(void* ptr){
	return static_cast<QJSValue*>(ptr)->toQObject();
}

char* QJSValue_ToString(void* ptr){
	return static_cast<QJSValue*>(ptr)->toString().toUtf8().data();
}

void* QJSValue_ToVariant(void* ptr){
	return new QVariant(static_cast<QJSValue*>(ptr)->toVariant());
}

void QJSValue_DestroyQJSValue(void* ptr){
	static_cast<QJSValue*>(ptr)->~QJSValue();
}

