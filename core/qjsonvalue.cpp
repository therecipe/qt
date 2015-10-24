#include "qjsonvalue.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJsonObject>
#include <QJsonArray>
#include <QLatin1String>
#include <QString>
#include <QJsonValue>
#include "_cgo_export.h"

class MyQJsonValue: public QJsonValue {
public:
};

QtObjectPtr QJsonValue_NewQJsonValue5(QtObjectPtr s){
	return new QJsonValue(*static_cast<QLatin1String*>(s));
}

QtObjectPtr QJsonValue_NewQJsonValue(int ty){
	return new QJsonValue(static_cast<QJsonValue::Type>(ty));
}

QtObjectPtr QJsonValue_NewQJsonValue2(int b){
	return new QJsonValue(b != 0);
}

QtObjectPtr QJsonValue_NewQJsonValue7(QtObjectPtr a){
	return new QJsonValue(*static_cast<QJsonArray*>(a));
}

QtObjectPtr QJsonValue_NewQJsonValue8(QtObjectPtr o){
	return new QJsonValue(*static_cast<QJsonObject*>(o));
}

QtObjectPtr QJsonValue_NewQJsonValue9(QtObjectPtr other){
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

QtObjectPtr QJsonValue_NewQJsonValue4(char* s){
	return new QJsonValue(QString(s));
}

QtObjectPtr QJsonValue_NewQJsonValue6(char* s){
	return new QJsonValue(const_cast<const char*>(s));
}

QtObjectPtr QJsonValue_NewQJsonValue12(int n){
	return new QJsonValue(n);
}

int QJsonValue_IsArray(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isArray();
}

int QJsonValue_IsBool(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isBool();
}

int QJsonValue_IsDouble(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isDouble();
}

int QJsonValue_IsNull(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isNull();
}

int QJsonValue_IsObject(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isObject();
}

int QJsonValue_IsString(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isString();
}

int QJsonValue_IsUndefined(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->isUndefined();
}

int QJsonValue_ToBool(QtObjectPtr ptr, int defaultValue){
	return static_cast<QJsonValue*>(ptr)->toBool(defaultValue != 0);
}

int QJsonValue_ToInt(QtObjectPtr ptr, int defaultValue){
	return static_cast<QJsonValue*>(ptr)->toInt(defaultValue);
}

char* QJsonValue_ToString(QtObjectPtr ptr, char* defaultValue){
	return static_cast<QJsonValue*>(ptr)->toString(QString(defaultValue)).toUtf8().data();
}

char* QJsonValue_ToVariant(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->toVariant().toString().toUtf8().data();
}

int QJsonValue_Type(QtObjectPtr ptr){
	return static_cast<QJsonValue*>(ptr)->type();
}

void QJsonValue_DestroyQJsonValue(QtObjectPtr ptr){
	static_cast<QJsonValue*>(ptr)->~QJsonValue();
}

