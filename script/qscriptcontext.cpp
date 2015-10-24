#include "qscriptcontext.h"
#include <QModelIndex>
#include <QScriptValue>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QScriptContext>
#include "_cgo_export.h"

class MyQScriptContext: public QScriptContext {
public:
};

int QScriptContext_ArgumentCount(QtObjectPtr ptr){
	return static_cast<QScriptContext*>(ptr)->argumentCount();
}

char* QScriptContext_Backtrace(QtObjectPtr ptr){
	return static_cast<QScriptContext*>(ptr)->backtrace().join("|").toUtf8().data();
}

QtObjectPtr QScriptContext_Engine(QtObjectPtr ptr){
	return static_cast<QScriptContext*>(ptr)->engine();
}

int QScriptContext_IsCalledAsConstructor(QtObjectPtr ptr){
	return static_cast<QScriptContext*>(ptr)->isCalledAsConstructor();
}

QtObjectPtr QScriptContext_ParentContext(QtObjectPtr ptr){
	return static_cast<QScriptContext*>(ptr)->parentContext();
}

void QScriptContext_SetActivationObject(QtObjectPtr ptr, QtObjectPtr activation){
	static_cast<QScriptContext*>(ptr)->setActivationObject(*static_cast<QScriptValue*>(activation));
}

void QScriptContext_SetThisObject(QtObjectPtr ptr, QtObjectPtr thisObject){
	static_cast<QScriptContext*>(ptr)->setThisObject(*static_cast<QScriptValue*>(thisObject));
}

char* QScriptContext_ToString(QtObjectPtr ptr){
	return static_cast<QScriptContext*>(ptr)->toString().toUtf8().data();
}

void QScriptContext_DestroyQScriptContext(QtObjectPtr ptr){
	static_cast<QScriptContext*>(ptr)->~QScriptContext();
}

