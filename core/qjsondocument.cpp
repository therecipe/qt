#include "qjsondocument.h"
#include <QModelIndex>
#include <QJsonArray>
#include <QJsonObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QJsonDocument>
#include "_cgo_export.h"

class MyQJsonDocument: public QJsonDocument {
public:
};

QtObjectPtr QJsonDocument_NewQJsonDocument(){
	return new QJsonDocument();
}

QtObjectPtr QJsonDocument_NewQJsonDocument3(QtObjectPtr array){
	return new QJsonDocument(*static_cast<QJsonArray*>(array));
}

QtObjectPtr QJsonDocument_NewQJsonDocument4(QtObjectPtr other){
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

QtObjectPtr QJsonDocument_NewQJsonDocument2(QtObjectPtr object){
	return new QJsonDocument(*static_cast<QJsonObject*>(object));
}

int QJsonDocument_IsArray(QtObjectPtr ptr){
	return static_cast<QJsonDocument*>(ptr)->isArray();
}

int QJsonDocument_IsEmpty(QtObjectPtr ptr){
	return static_cast<QJsonDocument*>(ptr)->isEmpty();
}

int QJsonDocument_IsNull(QtObjectPtr ptr){
	return static_cast<QJsonDocument*>(ptr)->isNull();
}

int QJsonDocument_IsObject(QtObjectPtr ptr){
	return static_cast<QJsonDocument*>(ptr)->isObject();
}

void QJsonDocument_SetArray(QtObjectPtr ptr, QtObjectPtr array){
	static_cast<QJsonDocument*>(ptr)->setArray(*static_cast<QJsonArray*>(array));
}

void QJsonDocument_SetObject(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QJsonDocument*>(ptr)->setObject(*static_cast<QJsonObject*>(object));
}

char* QJsonDocument_ToVariant(QtObjectPtr ptr){
	return static_cast<QJsonDocument*>(ptr)->toVariant().toString().toUtf8().data();
}

void QJsonDocument_DestroyQJsonDocument(QtObjectPtr ptr){
	static_cast<QJsonDocument*>(ptr)->~QJsonDocument();
}

