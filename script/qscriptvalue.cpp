#include "qscriptvalue.h"
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QScriptString>
#include <QMetaObject>
#include <QScriptClass>
#include <QString>
#include <QVariant>
#include <QObject>
#include <QScriptValue>
#include "_cgo_export.h"

class MyQScriptValue: public QScriptValue {
public:
};

QtObjectPtr QScriptValue_NewQScriptValue(){
	return new QScriptValue();
}

QtObjectPtr QScriptValue_NewQScriptValue10(int value){
	return new QScriptValue(static_cast<QScriptValue::SpecialValue>(value));
}

QtObjectPtr QScriptValue_NewQScriptValue11(int value){
	return new QScriptValue(value != 0);
}

QtObjectPtr QScriptValue_NewQScriptValue16(QtObjectPtr value){
	return new QScriptValue(*static_cast<QLatin1String*>(value));
}

QtObjectPtr QScriptValue_NewQScriptValue2(QtObjectPtr other){
	return new QScriptValue(*static_cast<QScriptValue*>(other));
}

QtObjectPtr QScriptValue_NewQScriptValue15(char* value){
	return new QScriptValue(QString(value));
}

QtObjectPtr QScriptValue_NewQScriptValue17(char* value){
	return new QScriptValue(const_cast<const char*>(value));
}

QtObjectPtr QScriptValue_NewQScriptValue12(int value){
	return new QScriptValue(value);
}

QtObjectPtr QScriptValue_Engine(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->engine();
}

int QScriptValue_Equals(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QScriptValue*>(ptr)->equals(*static_cast<QScriptValue*>(other));
}

int QScriptValue_InstanceOf(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QScriptValue*>(ptr)->instanceOf(*static_cast<QScriptValue*>(other));
}

int QScriptValue_IsArray(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isArray();
}

int QScriptValue_IsBool(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isBool();
}

int QScriptValue_IsDate(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isDate();
}

int QScriptValue_IsError(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isError();
}

int QScriptValue_IsFunction(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isFunction();
}

int QScriptValue_IsNull(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isNull();
}

int QScriptValue_IsNumber(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isNumber();
}

int QScriptValue_IsObject(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isObject();
}

int QScriptValue_IsQMetaObject(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isQMetaObject();
}

int QScriptValue_IsQObject(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isQObject();
}

int QScriptValue_IsRegExp(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isRegExp();
}

int QScriptValue_IsString(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isString();
}

int QScriptValue_IsUndefined(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isUndefined();
}

int QScriptValue_IsValid(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isValid();
}

int QScriptValue_IsVariant(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->isVariant();
}

int QScriptValue_LessThan(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QScriptValue*>(ptr)->lessThan(*static_cast<QScriptValue*>(other));
}

int QScriptValue_PropertyFlags2(QtObjectPtr ptr, QtObjectPtr name, int mode){
	return static_cast<QScriptValue*>(ptr)->propertyFlags(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

int QScriptValue_PropertyFlags(QtObjectPtr ptr, char* name, int mode){
	return static_cast<QScriptValue*>(ptr)->propertyFlags(QString(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

QtObjectPtr QScriptValue_ScriptClass(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->scriptClass();
}

void QScriptValue_SetData(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QScriptValue*>(ptr)->setData(*static_cast<QScriptValue*>(data));
}

void QScriptValue_SetProperty2(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value, int flags){
	static_cast<QScriptValue*>(ptr)->setProperty(*static_cast<QScriptString*>(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetProperty(QtObjectPtr ptr, char* name, QtObjectPtr value, int flags){
	static_cast<QScriptValue*>(ptr)->setProperty(QString(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetPrototype(QtObjectPtr ptr, QtObjectPtr prototype){
	static_cast<QScriptValue*>(ptr)->setPrototype(*static_cast<QScriptValue*>(prototype));
}

void QScriptValue_SetScriptClass(QtObjectPtr ptr, QtObjectPtr scriptClass){
	static_cast<QScriptValue*>(ptr)->setScriptClass(static_cast<QScriptClass*>(scriptClass));
}

int QScriptValue_StrictlyEquals(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QScriptValue*>(ptr)->strictlyEquals(*static_cast<QScriptValue*>(other));
}

int QScriptValue_ToBool(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->toBool();
}

QtObjectPtr QScriptValue_ToQMetaObject(QtObjectPtr ptr){
	return const_cast<QMetaObject*>(static_cast<QScriptValue*>(ptr)->toQMetaObject());
}

QtObjectPtr QScriptValue_ToQObject(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->toQObject();
}

char* QScriptValue_ToString(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->toString().toUtf8().data();
}

char* QScriptValue_ToVariant(QtObjectPtr ptr){
	return static_cast<QScriptValue*>(ptr)->toVariant().toString().toUtf8().data();
}

void QScriptValue_DestroyQScriptValue(QtObjectPtr ptr){
	static_cast<QScriptValue*>(ptr)->~QScriptValue();
}

