#include "qqmlpropertymap.h"
#include <QObject>
#include <QQmlProperty>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlPropertyMap>
#include "_cgo_export.h"

class MyQQmlPropertyMap: public QQmlPropertyMap {
public:
void Signal_ValueChanged(const QString & key, const QVariant & value){callbackQQmlPropertyMapValueChanged(this->objectName().toUtf8().data(), key.toUtf8().data(), new QVariant(value));};
};

void* QQmlPropertyMap_NewQQmlPropertyMap(void* parent){
	return new QQmlPropertyMap(static_cast<QObject*>(parent));
}

void QQmlPropertyMap_Clear(void* ptr, char* key){
	static_cast<QQmlPropertyMap*>(ptr)->clear(QString(key));
}

int QQmlPropertyMap_Contains(void* ptr, char* key){
	return static_cast<QQmlPropertyMap*>(ptr)->contains(QString(key));
}

int QQmlPropertyMap_Count(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->count();
}

int QQmlPropertyMap_IsEmpty(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->isEmpty();
}

char* QQmlPropertyMap_Keys(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->keys().join("|").toUtf8().data();
}

int QQmlPropertyMap_Size(void* ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->size();
}

void* QQmlPropertyMap_Value(void* ptr, char* key){
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->value(QString(key)));
}

void QQmlPropertyMap_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));;
}

void QQmlPropertyMap_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));;
}

void QQmlPropertyMap_DestroyQQmlPropertyMap(void* ptr){
	static_cast<QQmlPropertyMap*>(ptr)->~QQmlPropertyMap();
}

