#include "qscriptextensionplugin.h"
#include <QScriptEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptValue>
#include <QScriptExtensionPlugin>
#include "_cgo_export.h"

class MyQScriptExtensionPlugin: public QScriptExtensionPlugin {
public:
};

void QScriptExtensionPlugin_Initialize(void* ptr, char* key, void* engine){
	static_cast<QScriptExtensionPlugin*>(ptr)->initialize(QString(key), static_cast<QScriptEngine*>(engine));
}

char* QScriptExtensionPlugin_Keys(void* ptr){
	return static_cast<QScriptExtensionPlugin*>(ptr)->keys().join("|").toUtf8().data();
}

void* QScriptExtensionPlugin_SetupPackage(void* ptr, char* key, void* engine){
	return new QScriptValue(static_cast<QScriptExtensionPlugin*>(ptr)->setupPackage(QString(key), static_cast<QScriptEngine*>(engine)));
}

void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(void* ptr){
	static_cast<QScriptExtensionPlugin*>(ptr)->~QScriptExtensionPlugin();
}

