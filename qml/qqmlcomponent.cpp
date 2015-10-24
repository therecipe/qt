#include "qqmlcomponent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlContext>
#include <QQmlIncubator>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <QQmlEngine>
#include <QQmlComponent>
#include "_cgo_export.h"

class MyQQmlComponent: public QQmlComponent {
public:
void Signal_StatusChanged(QQmlComponent::Status status){callbackQQmlComponentStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QQmlComponent_Status(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->status();
}

char* QQmlComponent_Url(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->url().toString().toUtf8().data();
}

QtObjectPtr QQmlComponent_NewQQmlComponent(QtObjectPtr engine, QtObjectPtr parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlComponent_NewQQmlComponent4(QtObjectPtr engine, char* fileName, int mode, QtObjectPtr parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlComponent_NewQQmlComponent3(QtObjectPtr engine, char* fileName, QtObjectPtr parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlComponent_NewQQmlComponent6(QtObjectPtr engine, char* url, int mode, QtObjectPtr parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), QUrl(QString(url)), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlComponent_NewQQmlComponent5(QtObjectPtr engine, char* url, QtObjectPtr parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), QUrl(QString(url)), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlComponent_BeginCreate(QtObjectPtr ptr, QtObjectPtr publicContext){
	return static_cast<QQmlComponent*>(ptr)->beginCreate(static_cast<QQmlContext*>(publicContext));
}

void QQmlComponent_CompleteCreate(QtObjectPtr ptr){
	static_cast<QQmlComponent*>(ptr)->completeCreate();
}

QtObjectPtr QQmlComponent_Create(QtObjectPtr ptr, QtObjectPtr context){
	return static_cast<QQmlComponent*>(ptr)->create(static_cast<QQmlContext*>(context));
}

void QQmlComponent_Create2(QtObjectPtr ptr, QtObjectPtr incubator, QtObjectPtr context, QtObjectPtr forContext){
	static_cast<QQmlComponent*>(ptr)->create(*static_cast<QQmlIncubator*>(incubator), static_cast<QQmlContext*>(context), static_cast<QQmlContext*>(forContext));
}

QtObjectPtr QQmlComponent_CreationContext(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->creationContext();
}

int QQmlComponent_IsError(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->isError();
}

int QQmlComponent_IsLoading(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->isLoading();
}

int QQmlComponent_IsNull(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->isNull();
}

int QQmlComponent_IsReady(QtObjectPtr ptr){
	return static_cast<QQmlComponent*>(ptr)->isReady();
}

void QQmlComponent_LoadUrl(QtObjectPtr ptr, char* url){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, QUrl(QString(url))));
}

void QQmlComponent_LoadUrl2(QtObjectPtr ptr, char* url, int mode){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, QUrl(QString(url))), Q_ARG(QQmlComponent::CompilationMode, static_cast<QQmlComponent::CompilationMode>(mode)));
}

void QQmlComponent_SetData(QtObjectPtr ptr, QtObjectPtr data, char* url){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "setData", Q_ARG(QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(QUrl, QUrl(QString(url))));
}

void QQmlComponent_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));;
}

void QQmlComponent_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));;
}

void QQmlComponent_DestroyQQmlComponent(QtObjectPtr ptr){
	static_cast<QQmlComponent*>(ptr)->~QQmlComponent();
}

