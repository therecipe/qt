#include "qscriptextensionplugin.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptEngine>
#include <QScriptExtensionPlugin>
#include "_cgo_export.h"

class MyQScriptExtensionPlugin: public QScriptExtensionPlugin {
public:
};

void QScriptExtensionPlugin_Initialize(QtObjectPtr ptr, char* key, QtObjectPtr engine){
	static_cast<QScriptExtensionPlugin*>(ptr)->initialize(QString(key), static_cast<QScriptEngine*>(engine));
}

char* QScriptExtensionPlugin_Keys(QtObjectPtr ptr){
	return static_cast<QScriptExtensionPlugin*>(ptr)->keys().join("|").toUtf8().data();
}

void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(QtObjectPtr ptr){
	static_cast<QScriptExtensionPlugin*>(ptr)->~QScriptExtensionPlugin();
}

