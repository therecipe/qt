#include "qscriptstring.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QScriptString>
#include "_cgo_export.h"

class MyQScriptString: public QScriptString {
public:
};

void* QScriptString_NewQScriptString(){
	return new QScriptString();
}

void* QScriptString_NewQScriptString2(void* other){
	return new QScriptString(*static_cast<QScriptString*>(other));
}

int QScriptString_IsValid(void* ptr){
	return static_cast<QScriptString*>(ptr)->isValid();
}

char* QScriptString_ToString(void* ptr){
	return static_cast<QScriptString*>(ptr)->toString().toUtf8().data();
}

void QScriptString_DestroyQScriptString(void* ptr){
	static_cast<QScriptString*>(ptr)->~QScriptString();
}

