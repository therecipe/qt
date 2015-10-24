#include "qscriptclass.h"
#include <QScriptEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptValue>
#include <QScriptClass>
#include "_cgo_export.h"

class MyQScriptClass: public QScriptClass {
public:
};

QtObjectPtr QScriptClass_NewQScriptClass(QtObjectPtr engine){
	return new QScriptClass(static_cast<QScriptEngine*>(engine));
}

QtObjectPtr QScriptClass_Engine(QtObjectPtr ptr){
	return static_cast<QScriptClass*>(ptr)->engine();
}

char* QScriptClass_Extension(QtObjectPtr ptr, int extension, char* argument){
	return static_cast<QScriptClass*>(ptr)->extension(static_cast<QScriptClass::Extension>(extension), QVariant(argument)).toString().toUtf8().data();
}

char* QScriptClass_Name(QtObjectPtr ptr){
	return static_cast<QScriptClass*>(ptr)->name().toUtf8().data();
}

QtObjectPtr QScriptClass_NewIterator(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QScriptClass*>(ptr)->newIterator(*static_cast<QScriptValue*>(object));
}

int QScriptClass_SupportsExtension(QtObjectPtr ptr, int extension){
	return static_cast<QScriptClass*>(ptr)->supportsExtension(static_cast<QScriptClass::Extension>(extension));
}

void QScriptClass_DestroyQScriptClass(QtObjectPtr ptr){
	static_cast<QScriptClass*>(ptr)->~QScriptClass();
}

