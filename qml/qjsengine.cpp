#include "qjsengine.h"
#include <QObject>
#include <QJSValue>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJSEngine>
#include "_cgo_export.h"

class MyQJSEngine: public QJSEngine {
public:
};

void* QJSEngine_NewQJSEngine(){
	return new QJSEngine();
}

void* QJSEngine_NewQJSEngine2(void* parent){
	return new QJSEngine(static_cast<QObject*>(parent));
}

void QJSEngine_CollectGarbage(void* ptr){
	static_cast<QJSEngine*>(ptr)->collectGarbage();
}

void* QJSEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->evaluate(QString(program), QString(fileName), lineNumber));
}

void* QJSEngine_GlobalObject(void* ptr){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->globalObject());
}

void QJSEngine_InstallTranslatorFunctions(void* ptr, void* object){
	static_cast<QJSEngine*>(ptr)->installTranslatorFunctions(*static_cast<QJSValue*>(object));
}

void* QJSEngine_NewObject(void* ptr){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newObject());
}

void* QJSEngine_NewQObject(void* ptr, void* object){
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newQObject(static_cast<QObject*>(object)));
}

void QJSEngine_DestroyQJSEngine(void* ptr){
	static_cast<QJSEngine*>(ptr)->~QJSEngine();
}

