#include "qjsvalue.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QJSValue>
#include "_cgo_export.h"

class MyQJSValue: public QJSValue {
public:
};

QtObjectPtr QJSValue_NewQJSValue3(QtObjectPtr other){
	return new QJSValue(*static_cast<QJSValue*>(other));
}

QtObjectPtr QJSValue_NewQJSValue(int value){
	return new QJSValue(static_cast<QJSValue::SpecialValue>(value));
}

QtObjectPtr QJSValue_NewQJSValue4(int value){
	return new QJSValue(value != 0);
}

QtObjectPtr QJSValue_NewQJSValue2(QtObjectPtr other){
	return new QJSValue(*static_cast<QJSValue*>(other));
}

QtObjectPtr QJSValue_NewQJSValue9(QtObjectPtr value){
	return new QJSValue(*static_cast<QLatin1String*>(value));
}

QtObjectPtr QJSValue_NewQJSValue8(char* value){
	return new QJSValue(QString(value));
}

QtObjectPtr QJSValue_NewQJSValue10(char* value){
	return new QJSValue(const_cast<const char*>(value));
}

QtObjectPtr QJSValue_NewQJSValue5(int value){
	return new QJSValue(value);
}

int QJSValue_DeleteProperty(QtObjectPtr ptr, char* name){
	return static_cast<QJSValue*>(ptr)->deleteProperty(QString(name));
}

int QJSValue_Equals(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QJSValue*>(ptr)->equals(*static_cast<QJSValue*>(other));
}

int QJSValue_HasOwnProperty(QtObjectPtr ptr, char* name){
	return static_cast<QJSValue*>(ptr)->hasOwnProperty(QString(name));
}

int QJSValue_HasProperty(QtObjectPtr ptr, char* name){
	return static_cast<QJSValue*>(ptr)->hasProperty(QString(name));
}

int QJSValue_IsArray(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isArray();
}

int QJSValue_IsBool(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isBool();
}

int QJSValue_IsCallable(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isCallable();
}

int QJSValue_IsDate(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isDate();
}

int QJSValue_IsError(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isError();
}

int QJSValue_IsNull(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isNull();
}

int QJSValue_IsNumber(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isNumber();
}

int QJSValue_IsObject(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isObject();
}

int QJSValue_IsQObject(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isQObject();
}

int QJSValue_IsRegExp(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isRegExp();
}

int QJSValue_IsString(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isString();
}

int QJSValue_IsUndefined(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isUndefined();
}

int QJSValue_IsVariant(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->isVariant();
}

void QJSValue_SetProperty(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QJSValue*>(ptr)->setProperty(QString(name), *static_cast<QJSValue*>(value));
}

void QJSValue_SetPrototype(QtObjectPtr ptr, QtObjectPtr prototype){
	static_cast<QJSValue*>(ptr)->setPrototype(*static_cast<QJSValue*>(prototype));
}

int QJSValue_StrictlyEquals(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QJSValue*>(ptr)->strictlyEquals(*static_cast<QJSValue*>(other));
}

int QJSValue_ToBool(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->toBool();
}

QtObjectPtr QJSValue_ToQObject(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->toQObject();
}

char* QJSValue_ToString(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->toString().toUtf8().data();
}

char* QJSValue_ToVariant(QtObjectPtr ptr){
	return static_cast<QJSValue*>(ptr)->toVariant().toString().toUtf8().data();
}

void QJSValue_DestroyQJSValue(QtObjectPtr ptr){
	static_cast<QJSValue*>(ptr)->~QJSValue();
}

