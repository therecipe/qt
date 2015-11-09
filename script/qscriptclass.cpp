#include "qscriptclass.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptValue>
#include <QScriptEngine>
#include <QScriptClass>
#include "_cgo_export.h"

class MyQScriptClass: public QScriptClass {
public:
};

void* QScriptClass_NewQScriptClass(void* engine){
	return new QScriptClass(static_cast<QScriptEngine*>(engine));
}

void* QScriptClass_Engine(void* ptr){
	return static_cast<QScriptClass*>(ptr)->engine();
}

void* QScriptClass_Extension(void* ptr, int extension, void* argument){
	return new QVariant(static_cast<QScriptClass*>(ptr)->extension(static_cast<QScriptClass::Extension>(extension), *static_cast<QVariant*>(argument)));
}

char* QScriptClass_Name(void* ptr){
	return static_cast<QScriptClass*>(ptr)->name().toUtf8().data();
}

void* QScriptClass_NewIterator(void* ptr, void* object){
	return static_cast<QScriptClass*>(ptr)->newIterator(*static_cast<QScriptValue*>(object));
}

void* QScriptClass_Prototype(void* ptr){
	return new QScriptValue(static_cast<QScriptClass*>(ptr)->prototype());
}

int QScriptClass_SupportsExtension(void* ptr, int extension){
	return static_cast<QScriptClass*>(ptr)->supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

void QScriptClass_DestroyQScriptClass(void* ptr){
	static_cast<QScriptClass*>(ptr)->~QScriptClass();
}

