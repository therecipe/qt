#include "qqmlapplicationengine.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QByteArray>
#include <QObject>
#include <QQmlApplicationEngine>
#include "_cgo_export.h"

class MyQQmlApplicationEngine: public QQmlApplicationEngine {
public:
};

void* QQmlApplicationEngine_NewQQmlApplicationEngine(void* parent){
	return new QQmlApplicationEngine(static_cast<QObject*>(parent));
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine3(char* filePath, void* parent){
	return new QQmlApplicationEngine(QString(filePath), static_cast<QObject*>(parent));
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine2(void* url, void* parent){
	return new QQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QObject*>(parent));
}

void QQmlApplicationEngine_Load2(void* ptr, char* filePath){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QString, QString(filePath)));
}

void QQmlApplicationEngine_Load(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_LoadData(void* ptr, void* data, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "loadData", Q_ARG(QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_DestroyQQmlApplicationEngine(void* ptr){
	static_cast<QQmlApplicationEngine*>(ptr)->~QQmlApplicationEngine();
}

