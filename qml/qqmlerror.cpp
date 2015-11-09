#include "qqmlerror.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QQmlError>
#include "_cgo_export.h"

class MyQQmlError: public QQmlError {
public:
};

void* QQmlError_NewQQmlError(){
	return new QQmlError();
}

void* QQmlError_NewQQmlError2(void* other){
	return new QQmlError(*static_cast<QQmlError*>(other));
}

int QQmlError_Column(void* ptr){
	return static_cast<QQmlError*>(ptr)->column();
}

char* QQmlError_Description(void* ptr){
	return static_cast<QQmlError*>(ptr)->description().toUtf8().data();
}

int QQmlError_IsValid(void* ptr){
	return static_cast<QQmlError*>(ptr)->isValid();
}

int QQmlError_Line(void* ptr){
	return static_cast<QQmlError*>(ptr)->line();
}

void* QQmlError_Object(void* ptr){
	return static_cast<QQmlError*>(ptr)->object();
}

void QQmlError_SetColumn(void* ptr, int column){
	static_cast<QQmlError*>(ptr)->setColumn(column);
}

void QQmlError_SetDescription(void* ptr, char* description){
	static_cast<QQmlError*>(ptr)->setDescription(QString(description));
}

void QQmlError_SetLine(void* ptr, int line){
	static_cast<QQmlError*>(ptr)->setLine(line);
}

void QQmlError_SetObject(void* ptr, void* object){
	static_cast<QQmlError*>(ptr)->setObject(static_cast<QObject*>(object));
}

void QQmlError_SetUrl(void* ptr, void* url){
	static_cast<QQmlError*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

char* QQmlError_ToString(void* ptr){
	return static_cast<QQmlError*>(ptr)->toString().toUtf8().data();
}

