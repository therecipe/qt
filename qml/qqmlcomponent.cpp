#include "qqmlcomponent.h"
#include <QString>
#include <QQmlContext>
#include <QMetaObject>
#include <QQmlIncubator>
#include <QObject>
#include <QByteArray>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlEngine>
#include <QQmlComponent>
#include "_cgo_export.h"

class MyQQmlComponent: public QQmlComponent {
public:
void Signal_StatusChanged(QQmlComponent::Status status){callbackQQmlComponentStatusChanged(this->objectName().toUtf8().data(), status);};
};

double QQmlComponent_Progress(void* ptr){
	return static_cast<double>(static_cast<QQmlComponent*>(ptr)->progress());
}

int QQmlComponent_Status(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->status();
}

void* QQmlComponent_NewQQmlComponent(void* engine, void* parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent4(void* engine, char* fileName, int mode, void* parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent3(void* engine, char* fileName, void* parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), QString(fileName), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent6(void* engine, void* url, int mode, void* parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
}

void* QQmlComponent_NewQQmlComponent5(void* engine, void* url, void* parent){
	return new QQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QObject*>(parent));
}

void* QQmlComponent_BeginCreate(void* ptr, void* publicContext){
	return static_cast<QQmlComponent*>(ptr)->beginCreate(static_cast<QQmlContext*>(publicContext));
}

void QQmlComponent_CompleteCreate(void* ptr){
	static_cast<QQmlComponent*>(ptr)->completeCreate();
}

void* QQmlComponent_Create(void* ptr, void* context){
	return static_cast<QQmlComponent*>(ptr)->create(static_cast<QQmlContext*>(context));
}

void QQmlComponent_Create2(void* ptr, void* incubator, void* context, void* forContext){
	static_cast<QQmlComponent*>(ptr)->create(*static_cast<QQmlIncubator*>(incubator), static_cast<QQmlContext*>(context), static_cast<QQmlContext*>(forContext));
}

void* QQmlComponent_CreationContext(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->creationContext();
}

int QQmlComponent_IsError(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isError();
}

int QQmlComponent_IsLoading(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isLoading();
}

int QQmlComponent_IsNull(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isNull();
}

int QQmlComponent_IsReady(void* ptr){
	return static_cast<QQmlComponent*>(ptr)->isReady();
}

void QQmlComponent_LoadUrl(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_LoadUrl2(void* ptr, void* url, int mode){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(QUrl, *static_cast<QUrl*>(url)), Q_ARG(QQmlComponent::CompilationMode, static_cast<QQmlComponent::CompilationMode>(mode)));
}

void QQmlComponent_SetData(void* ptr, void* data, void* url){
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "setData", Q_ARG(QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));;
}

void QQmlComponent_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));;
}

void QQmlComponent_DestroyQQmlComponent(void* ptr){
	static_cast<QQmlComponent*>(ptr)->~QQmlComponent();
}

