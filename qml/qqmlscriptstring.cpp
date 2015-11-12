#include "qqmlscriptstring.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QQmlScriptString>
#include "_cgo_export.h"

class MyQQmlScriptString: public QQmlScriptString {
public:
};

void* QQmlScriptString_NewQQmlScriptString(){
	return new QQmlScriptString();
}

void* QQmlScriptString_NewQQmlScriptString2(void* other){
	return new QQmlScriptString(*static_cast<QQmlScriptString*>(other));
}

int QQmlScriptString_BooleanLiteral(void* ptr, int ok){
	return static_cast<QQmlScriptString*>(ptr)->booleanLiteral(NULL);
}

int QQmlScriptString_IsEmpty(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->isEmpty();
}

int QQmlScriptString_IsNullLiteral(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->isNullLiteral();
}

int QQmlScriptString_IsUndefinedLiteral(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->isUndefinedLiteral();
}

double QQmlScriptString_NumberLiteral(void* ptr, int ok){
	return static_cast<double>(static_cast<QQmlScriptString*>(ptr)->numberLiteral(NULL));
}

char* QQmlScriptString_StringLiteral(void* ptr){
	return static_cast<QQmlScriptString*>(ptr)->stringLiteral().toUtf8().data();
}

