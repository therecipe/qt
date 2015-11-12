#include "qjsondocument.h"
#include <QByteArray>
#include <QJsonArray>
#include <QJsonObject>
#include <QJsonParseError>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJsonDocument>
#include "_cgo_export.h"

class MyQJsonDocument: public QJsonDocument {
public:
};

void* QJsonDocument_NewQJsonDocument(){
	return new QJsonDocument();
}

void* QJsonDocument_NewQJsonDocument3(void* array){
	return new QJsonDocument(*static_cast<QJsonArray*>(array));
}

void* QJsonDocument_NewQJsonDocument4(void* other){
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_NewQJsonDocument2(void* object){
	return new QJsonDocument(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_Array(void* ptr){
	return new QJsonArray(static_cast<QJsonDocument*>(ptr)->array());
}

void* QJsonDocument_QJsonDocument_FromBinaryData(void* data, int validation){
	return new QJsonDocument(QJsonDocument::fromBinaryData(*static_cast<QByteArray*>(data), static_cast<QJsonDocument::DataValidation>(validation)));
}

void* QJsonDocument_QJsonDocument_FromJson(void* json, void* error){
	return new QJsonDocument(QJsonDocument::fromJson(*static_cast<QByteArray*>(json), static_cast<QJsonParseError*>(error)));
}

void* QJsonDocument_QJsonDocument_FromRawData(char* data, int size, int validation){
	return new QJsonDocument(QJsonDocument::fromRawData(const_cast<const char*>(data), size, static_cast<QJsonDocument::DataValidation>(validation)));
}

void* QJsonDocument_QJsonDocument_FromVariant(void* variant){
	return new QJsonDocument(QJsonDocument::fromVariant(*static_cast<QVariant*>(variant)));
}

int QJsonDocument_IsArray(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isArray();
}

int QJsonDocument_IsEmpty(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isEmpty();
}

int QJsonDocument_IsNull(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isNull();
}

int QJsonDocument_IsObject(void* ptr){
	return static_cast<QJsonDocument*>(ptr)->isObject();
}

void* QJsonDocument_Object(void* ptr){
	return new QJsonObject(static_cast<QJsonDocument*>(ptr)->object());
}

void QJsonDocument_SetArray(void* ptr, void* array){
	static_cast<QJsonDocument*>(ptr)->setArray(*static_cast<QJsonArray*>(array));
}

void QJsonDocument_SetObject(void* ptr, void* object){
	static_cast<QJsonDocument*>(ptr)->setObject(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_ToBinaryData(void* ptr){
	return new QByteArray(static_cast<QJsonDocument*>(ptr)->toBinaryData());
}

void* QJsonDocument_ToJson(void* ptr, int format){
	return new QByteArray(static_cast<QJsonDocument*>(ptr)->toJson(static_cast<QJsonDocument::JsonFormat>(format)));
}

void* QJsonDocument_ToVariant(void* ptr){
	return new QVariant(static_cast<QJsonDocument*>(ptr)->toVariant());
}

void QJsonDocument_DestroyQJsonDocument(void* ptr){
	static_cast<QJsonDocument*>(ptr)->~QJsonDocument();
}

