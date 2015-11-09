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

void QQmlExtensionPlugin_InitializeEngine(void* ptr, void* engine, char* uri){
	static_cast<QQmlExtensionPlugin*>(ptr)->initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

void QQmlExtensionPlugin_RegisterTypes(void* ptr, char* uri){
	static_cast<QQmlExtensionPlugin*>(ptr)->registerTypes(const_cast<const char*>(uri));
}

