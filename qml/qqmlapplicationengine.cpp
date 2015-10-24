#include "qqmlapplicationengine.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QMetaObject>
#include <QObject>
#include <QQmlApplicationEngine>
#include "_cgo_export.h"

class MyQQmlApplicationEngine: public QQmlApplicationEngine {
public:
void Signal_ObjectCreated(QObject * object, const QUrl & url){callbackQQmlApplicationEngineObjectCreated(this->objectName().toUtf8().data(), object, url.toString().toUtf8().data());};
};

QtObjectPtr QQmlApplicationEngine_NewQQmlApplicationEngine(QtObjectPtr parent){
	return new QQmlApplicationEngine(static_cast<QObject*>(parent));
}

QtObjectPtr QQmlApplicationEngine_NewQQmlApplicationEngine3(char* filePath, QtObjectPtr parent){
	return new QQmlApplicationEngine(QString(filePath), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlApplicationEngine_NewQQmlApplicationEngine2(char* url, QtObjectPtr parent){
	return new QQmlApplicationEngine(QUrl(QString(url)), static_cast<QObject*>(parent));
}

void QQmlApplicationEngine_Load2(QtObjectPtr ptr, char* filePath){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QString, QString(filePath)));
}

void QQmlApplicationEngine_Load(QtObjectPtr ptr, char* url){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(QUrl, QUrl(QString(url))));
}

void QQmlApplicationEngine_LoadData(QtObjectPtr ptr, QtObjectPtr data, char* url){
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "loadData", Q_ARG(QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(QUrl, QUrl(QString(url))));
}

void QQmlApplicationEngine_ConnectObjectCreated(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));;
}

void QQmlApplicationEngine_DisconnectObjectCreated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));;
}

void QQmlApplicationEngine_DestroyQQmlApplicationEngine(QtObjectPtr ptr){
	static_cast<QQmlApplicationEngine*>(ptr)->~QQmlApplicationEngine();
}

