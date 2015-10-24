#include "qjsengine.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJSValue>
#include <QJSEngine>
#include "_cgo_export.h"

class MyQJSEngine: public QJSEngine {
public:
};

QtObjectPtr QJSEngine_NewQJSEngine(){
	return new QJSEngine();
}

QtObjectPtr QJSEngine_NewQJSEngine2(QtObjectPtr parent){
	return new QJSEngine(static_cast<QObject*>(parent));
}

void QJSEngine_CollectGarbage(QtObjectPtr ptr){
	static_cast<QJSEngine*>(ptr)->collectGarbage();
}

void QJSEngine_InstallTranslatorFunctions(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QJSEngine*>(ptr)->installTranslatorFunctions(*static_cast<QJSValue*>(object));
}

void QJSEngine_DestroyQJSEngine(QtObjectPtr ptr){
	static_cast<QJSEngine*>(ptr)->~QJSEngine();
}

