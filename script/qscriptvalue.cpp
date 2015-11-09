#include "qscriptvalue.h"
#include <QDate>
#include <QObject>
#include <QLatin1String>
#include <QVariant>
#include <QUrl>
#include <QScriptString>
#include <QRegExp>
#include <QDateTime>
#include <QMetaObject>
#include <QString>
#include <QModelIndex>
#include <QScriptClass>
#include <QScriptValue>
#include "_cgo_export.h"

class MyQScriptValue: public QScriptValue {
public:
};

void* QScriptValue_NewQScriptValue(){
	return new QScriptValue();
}

void* QScriptValue_NewQScriptValue10(int value){
	return new QScriptValue(static_cast<QScriptValue::SpecialValue>(value));
}

void* QScriptValue_NewQScriptValue11(int value){
	return new QScriptValue(value != 0);
}

void* QScriptValue_NewQScriptValue16(void* value){
	return new QScriptValue(*static_cast<QLatin1String*>(value));
}

void* QScriptValue_NewQScriptValue2(void* other){
	return new QScriptValue(*static_cast<QScriptValue*>(other));
}

void* QScriptValue_NewQScriptValue15(char* value){
	return new QScriptValue(QString(value));
}

void* QScriptValue_NewQScriptValue17(char* value){
	return new QScriptValue(const_cast<const char*>(value));
}

void* QScriptValue_NewQScriptValue12(int value){
	return new QScriptValue(value);
}

void* QScriptValue_Call2(void* ptr, void* thisObject, void* arguments){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->call(*static_cast<QScriptValue*>(thisObject), *static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_Construct2(void* ptr, void* arguments){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->construct(*static_cast<QScriptValue*>(arguments)));
}

void* QScriptValue_Data(void* ptr){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->data());
}

void* QScriptValue_Engine(void* ptr){
	return static_cast<QScriptValue*>(ptr)->engine();
}

int QScriptValue_Equals(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->equals(*static_cast<QScriptValue*>(other));
}

int QScriptValue_InstanceOf(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->instanceOf(*static_cast<QScriptValue*>(other));
}

int QScriptValue_IsArray(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isArray();
}

int QScriptValue_IsBool(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isBool();
}

int QScriptValue_IsDate(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isDate();
}

int QScriptValue_IsError(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isError();
}

int QScriptValue_IsFunction(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isFunction();
}

int QScriptValue_IsNull(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isNull();
}

int QScriptValue_IsNumber(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isNumber();
}

int QScriptValue_IsObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isObject();
}

int QScriptValue_IsQMetaObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isQMetaObject();
}

int QScriptValue_IsQObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isQObject();
}

int QScriptValue_IsRegExp(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isRegExp();
}

int QScriptValue_IsString(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isString();
}

int QScriptValue_IsUndefined(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isUndefined();
}

int QScriptValue_IsValid(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isValid();
}

int QScriptValue_IsVariant(void* ptr){
	return static_cast<QScriptValue*>(ptr)->isVariant();
}

int QScriptValue_LessThan(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->lessThan(*static_cast<QScriptValue*>(other));
}

void* QScriptValue_Property2(void* ptr, void* name, int mode){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

void* QScriptValue_Property(void* ptr, char* name, int mode){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->property(QString(name), static_cast<QScriptValue::ResolveFlag>(mode)));
}

int QScriptValue_PropertyFlags2(void* ptr, void* name, int mode){
	return static_cast<QScriptValue*>(ptr)->propertyFlags(*static_cast<QScriptString*>(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

int QScriptValue_PropertyFlags(void* ptr, char* name, int mode){
	return static_cast<QScriptValue*>(ptr)->propertyFlags(QString(name), static_cast<QScriptValue::ResolveFlag>(mode));
}

void* QScriptValue_Prototype(void* ptr){
	return new QScriptValue(static_cast<QScriptValue*>(ptr)->prototype());
}

void* QScriptValue_ScriptClass(void* ptr){
	return static_cast<QScriptValue*>(ptr)->scriptClass();
}

void QScriptValue_SetData(void* ptr, void* data){
	static_cast<QScriptValue*>(ptr)->setData(*static_cast<QScriptValue*>(data));
}

void QScriptValue_SetProperty2(void* ptr, void* name, void* value, int flags){
	static_cast<QScriptValue*>(ptr)->setProperty(*static_cast<QScriptString*>(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetProperty(void* ptr, char* name, void* value, int flags){
	static_cast<QScriptValue*>(ptr)->setProperty(QString(name), *static_cast<QScriptValue*>(value), static_cast<QScriptValue::PropertyFlag>(flags));
}

void QScriptValue_SetPrototype(void* ptr, void* prototype){
	static_cast<QScriptValue*>(ptr)->setPrototype(*static_cast<QScriptValue*>(prototype));
}

void QScriptValue_SetScriptClass(void* ptr, void* scriptClass){
	static_cast<QScriptValue*>(ptr)->setScriptClass(static_cast<QScriptClass*>(scriptClass));
}

int QScriptValue_StrictlyEquals(void* ptr, void* other){
	return static_cast<QScriptValue*>(ptr)->strictlyEquals(*static_cast<QScriptValue*>(other));
}

int QScriptValue_ToBool(void* ptr){
	return static_cast<QScriptValue*>(ptr)->toBool();
}

void* QScriptValue_ToDateTime(void* ptr){
	return new QDateTime(static_cast<QScriptValue*>(ptr)->toDateTime());
}

void* QScriptValue_ToQMetaObject(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QScriptValue*>(ptr)->toQMetaObject());
}

void* QScriptValue_ToQObject(void* ptr){
	return static_cast<QScriptValue*>(ptr)->toQObject();
}

void* QScriptValue_ToRegExp(void* ptr){
	return new QRegExp(static_cast<QScriptValue*>(ptr)->toRegExp());
}

char* QScriptValue_ToString(void* ptr){
	return static_cast<QScriptValue*>(ptr)->toString().toUtf8().data();
}

void* QScriptValue_ToVariant(void* ptr){
	return new QVariant(static_cast<QScriptValue*>(ptr)->toVariant());
}

void QScriptValue_DestroyQScriptValue(void* ptr){
	static_cast<QScriptValue*>(ptr)->~QScriptValue();
}

