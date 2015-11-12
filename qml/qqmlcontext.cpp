#include "qqmlcontext.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlEngine>
#include <QObject>
#include <QQmlContext>
#include "_cgo_export.h"

class MyQQmlContext: public QQmlContext {
public:
};

void* QQmlContext_NewQQmlContext2(void* parentContext, void* parent){
	return new QQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QObject*>(parent));
}

void* QQmlContext_NewQQmlContext(void* engine, void* parent){
	return new QQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlContext_ContextObject(void* ptr){
	return static_cast<QQmlContext*>(ptr)->contextObject();
}

void* QQmlContext_ContextProperty(void* ptr, char* name){
	return new QVariant(static_cast<QQmlContext*>(ptr)->contextProperty(QString(name)));
}

void* QQmlContext_Engine(void* ptr){
	return static_cast<QQmlContext*>(ptr)->engine();
}

int QQmlContext_IsValid(void* ptr){
	return static_cast<QQmlContext*>(ptr)->isValid();
}

char* QQmlContext_NameForObject(void* ptr, void* object){
	return static_cast<QQmlContext*>(ptr)->nameForObject(static_cast<QObject*>(object)).toUtf8().data();
}

void* QQmlContext_ParentContext(void* ptr){
	return static_cast<QQmlContext*>(ptr)->parentContext();
}

void QQmlContext_SetBaseUrl(void* ptr, void* baseUrl){
	static_cast<QQmlContext*>(ptr)->setBaseUrl(*static_cast<QUrl*>(baseUrl));
}

void QQmlContext_SetContextObject(void* ptr, void* object){
	static_cast<QQmlContext*>(ptr)->setContextObject(static_cast<QObject*>(object));
}

void QQmlContext_SetContextProperty(void* ptr, char* name, void* value){
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), static_cast<QObject*>(value));
}

void QQmlContext_SetContextProperty2(void* ptr, char* name, void* value){
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), *static_cast<QVariant*>(value));
}

void QQmlContext_DestroyQQmlContext(void* ptr){
	static_cast<QQmlContext*>(ptr)->~QQmlContext();
}

