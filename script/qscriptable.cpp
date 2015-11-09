#include "qscriptable.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptValue>
#include <QScriptable>
#include "_cgo_export.h"

class MyQScriptable: public QScriptable {
public:
};

void* QScriptable_Argument(void* ptr, int index){
	return new QScriptValue(static_cast<QScriptable*>(ptr)->argument(index));
}

int QScriptable_ArgumentCount(void* ptr){
	return static_cast<QScriptable*>(ptr)->argumentCount();
}

void* QScriptable_Context(void* ptr){
	return static_cast<QScriptable*>(ptr)->context();
}

void* QScriptable_Engine(void* ptr){
	return static_cast<QScriptable*>(ptr)->engine();
}

void* QScriptable_ThisObject(void* ptr){
	return new QScriptValue(static_cast<QScriptable*>(ptr)->thisObject());
}

