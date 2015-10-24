#include "qqmlextensionplugin.h"
#include <QQmlEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlExtensionPlugin>
#include "_cgo_export.h"

class MyQQmlExtensionPlugin: public QQmlExtensionPlugin {
public:
};

void QQmlExtensionPlugin_InitializeEngine(QtObjectPtr ptr, QtObjectPtr engine, char* uri){
	static_cast<QQmlExtensionPlugin*>(ptr)->initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

char* QQmlExtensionPlugin_BaseUrl(QtObjectPtr ptr){
	return static_cast<QQmlExtensionPlugin*>(ptr)->baseUrl().toString().toUtf8().data();
}

void QQmlExtensionPlugin_RegisterTypes(QtObjectPtr ptr, char* uri){
	static_cast<QQmlExtensionPlugin*>(ptr)->registerTypes(const_cast<const char*>(uri));
}

