#include "qjsonobject.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJsonObject>
#include "_cgo_export.h"

class MyQJsonObject: public QJsonObject {
public:
};

int QJsonObject_Contains(QtObjectPtr ptr, char* key){
	return static_cast<QJsonObject*>(ptr)->contains(QString(key));
}

int QJsonObject_Count(QtObjectPtr ptr){
	return static_cast<QJsonObject*>(ptr)->count();
}

int QJsonObject_Empty(QtObjectPtr ptr){
	return static_cast<QJsonObject*>(ptr)->empty();
}

int QJsonObject_IsEmpty(QtObjectPtr ptr){
	return static_cast<QJsonObject*>(ptr)->isEmpty();
}

char* QJsonObject_Keys(QtObjectPtr ptr){
	return static_cast<QJsonObject*>(ptr)->keys().join("|").toUtf8().data();
}

int QJsonObject_Length(QtObjectPtr ptr){
	return static_cast<QJsonObject*>(ptr)->length();
}

int QJsonObject_Size(QtObjectPtr ptr){
	return static_cast<QJsonObject*>(ptr)->size();
}

void QJsonObject_DestroyQJsonObject(QtObjectPtr ptr){
	static_cast<QJsonObject*>(ptr)->~QJsonObject();
}

