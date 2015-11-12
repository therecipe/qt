#include "qscriptcontext.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptValue>
#include <QScriptContext>
#include "_cgo_export.h"

class MyQScriptContext: public QScriptContext {
public:
};

void* QScriptContext_ActivationObject(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->activationObject());
}

void* QScriptContext_Argument(void* ptr, int index){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argument(index));
}

int QScriptContext_ArgumentCount(void* ptr){
	return static_cast<QScriptContext*>(ptr)->argumentCount();
}

void* QScriptContext_ArgumentsObject(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->argumentsObject());
}

char* QScriptContext_Backtrace(void* ptr){
	return static_cast<QScriptContext*>(ptr)->backtrace().join("|").toUtf8().data();
}

void* QScriptContext_Callee(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->callee());
}

void* QScriptContext_Engine(void* ptr){
	return static_cast<QScriptContext*>(ptr)->engine();
}

int QScriptContext_IsCalledAsConstructor(void* ptr){
	return static_cast<QScriptContext*>(ptr)->isCalledAsConstructor();
}

void* QScriptContext_ParentContext(void* ptr){
	return static_cast<QScriptContext*>(ptr)->parentContext();
}

void QScriptContext_SetActivationObject(void* ptr, void* activation){
	static_cast<QScriptContext*>(ptr)->setActivationObject(*static_cast<QScriptValue*>(activation));
}

void QScriptContext_SetThisObject(void* ptr, void* thisObject){
	static_cast<QScriptContext*>(ptr)->setThisObject(*static_cast<QScriptValue*>(thisObject));
}

void* QScriptContext_ThisObject(void* ptr){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->thisObject());
}

void* QScriptContext_ThrowError(void* ptr, int error, char* text){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(static_cast<QScriptContext::Error>(error), QString(text)));
}

void* QScriptContext_ThrowError2(void* ptr, char* text){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwError(QString(text)));
}

void* QScriptContext_ThrowValue(void* ptr, void* value){
	return new QScriptValue(static_cast<QScriptContext*>(ptr)->throwValue(*static_cast<QScriptValue*>(value)));
}

char* QScriptContext_ToString(void* ptr){
	return static_cast<QScriptContext*>(ptr)->toString().toUtf8().data();
}

void QScriptContext_DestroyQScriptContext(void* ptr){
	static_cast<QScriptContext*>(ptr)->~QScriptContext();
}

