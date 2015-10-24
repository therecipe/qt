#include "qqmlcontext.h"
#include <QQmlEngine>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlContext>
#include "_cgo_export.h"

class MyQQmlContext: public QQmlContext {
public:
};

QtObjectPtr QQmlContext_NewQQmlContext2(QtObjectPtr parentContext, QtObjectPtr parent){
	return new QQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlContext_NewQQmlContext(QtObjectPtr engine, QtObjectPtr parent){
	return new QQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

char* QQmlContext_BaseUrl(QtObjectPtr ptr){
	return static_cast<QQmlContext*>(ptr)->baseUrl().toString().toUtf8().data();
}

QtObjectPtr QQmlContext_ContextObject(QtObjectPtr ptr){
	return static_cast<QQmlContext*>(ptr)->contextObject();
}

char* QQmlContext_ContextProperty(QtObjectPtr ptr, char* name){
	return static_cast<QQmlContext*>(ptr)->contextProperty(QString(name)).toString().toUtf8().data();
}

QtObjectPtr QQmlContext_Engine(QtObjectPtr ptr){
	return static_cast<QQmlContext*>(ptr)->engine();
}

int QQmlContext_IsValid(QtObjectPtr ptr){
	return static_cast<QQmlContext*>(ptr)->isValid();
}

char* QQmlContext_NameForObject(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QQmlContext*>(ptr)->nameForObject(static_cast<QObject*>(object)).toUtf8().data();
}

QtObjectPtr QQmlContext_ParentContext(QtObjectPtr ptr){
	return static_cast<QQmlContext*>(ptr)->parentContext();
}

char* QQmlContext_ResolvedUrl(QtObjectPtr ptr, char* src){
	return static_cast<QQmlContext*>(ptr)->resolvedUrl(QUrl(QString(src))).toString().toUtf8().data();
}

void QQmlContext_SetBaseUrl(QtObjectPtr ptr, char* baseUrl){
	static_cast<QQmlContext*>(ptr)->setBaseUrl(QUrl(QString(baseUrl)));
}

void QQmlContext_SetContextObject(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QQmlContext*>(ptr)->setContextObject(static_cast<QObject*>(object));
}

void QQmlContext_SetContextProperty(QtObjectPtr ptr, char* name, QtObjectPtr value){
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), static_cast<QObject*>(value));
}

void QQmlContext_SetContextProperty2(QtObjectPtr ptr, char* name, char* value){
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString(name), QVariant(value));
}

void QQmlContext_DestroyQQmlContext(QtObjectPtr ptr){
	static_cast<QQmlContext*>(ptr)->~QQmlContext();
}

