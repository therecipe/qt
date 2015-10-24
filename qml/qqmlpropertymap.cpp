#include "qqmlpropertymap.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QQmlProperty>
#include <QString>
#include <QQmlPropertyMap>
#include "_cgo_export.h"

class MyQQmlPropertyMap: public QQmlPropertyMap {
public:
void Signal_ValueChanged(const QString & key, const QVariant & value){callbackQQmlPropertyMapValueChanged(this->objectName().toUtf8().data(), key.toUtf8().data(), value.toString().toUtf8().data());};
};

QtObjectPtr QQmlPropertyMap_NewQQmlPropertyMap(QtObjectPtr parent){
	return new QQmlPropertyMap(static_cast<QObject*>(parent));
}

void QQmlPropertyMap_Clear(QtObjectPtr ptr, char* key){
	static_cast<QQmlPropertyMap*>(ptr)->clear(QString(key));
}

int QQmlPropertyMap_Contains(QtObjectPtr ptr, char* key){
	return static_cast<QQmlPropertyMap*>(ptr)->contains(QString(key));
}

int QQmlPropertyMap_Count(QtObjectPtr ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->count();
}

int QQmlPropertyMap_IsEmpty(QtObjectPtr ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->isEmpty();
}

char* QQmlPropertyMap_Keys(QtObjectPtr ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->keys().join("|").toUtf8().data();
}

int QQmlPropertyMap_Size(QtObjectPtr ptr){
	return static_cast<QQmlPropertyMap*>(ptr)->size();
}

char* QQmlPropertyMap_Value(QtObjectPtr ptr, char* key){
	return static_cast<QQmlPropertyMap*>(ptr)->value(QString(key)).toString().toUtf8().data();
}

void QQmlPropertyMap_ConnectValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));;
}

void QQmlPropertyMap_DisconnectValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));;
}

void QQmlPropertyMap_DestroyQQmlPropertyMap(QtObjectPtr ptr){
	static_cast<QQmlPropertyMap*>(ptr)->~QQmlPropertyMap();
}

