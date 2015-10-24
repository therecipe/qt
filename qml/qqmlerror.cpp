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

QtObjectPtr QQmlError_NewQQmlError(){
	return new QQmlError();
}

QtObjectPtr QQmlError_NewQQmlError2(QtObjectPtr other){
	return new QQmlError(*static_cast<QQmlError*>(other));
}

int QQmlError_Column(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->column();
}

char* QQmlError_Description(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->description().toUtf8().data();
}

int QQmlError_IsValid(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->isValid();
}

int QQmlError_Line(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->line();
}

QtObjectPtr QQmlError_Object(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->object();
}

void QQmlError_SetColumn(QtObjectPtr ptr, int column){
	static_cast<QQmlError*>(ptr)->setColumn(column);
}

void QQmlError_SetDescription(QtObjectPtr ptr, char* description){
	static_cast<QQmlError*>(ptr)->setDescription(QString(description));
}

void QQmlError_SetLine(QtObjectPtr ptr, int line){
	static_cast<QQmlError*>(ptr)->setLine(line);
}

void QQmlError_SetObject(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QQmlError*>(ptr)->setObject(static_cast<QObject*>(object));
}

void QQmlError_SetUrl(QtObjectPtr ptr, char* url){
	static_cast<QQmlError*>(ptr)->setUrl(QUrl(QString(url)));
}

char* QQmlError_ToString(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->toString().toUtf8().data();
}

char* QQmlError_Url(QtObjectPtr ptr){
	return static_cast<QQmlError*>(ptr)->url().toString().toUtf8().data();
}

