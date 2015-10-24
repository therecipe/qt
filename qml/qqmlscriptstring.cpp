#include "qqmlscriptstring.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlScriptString>
#include "_cgo_export.h"

class MyQQmlScriptString: public QQmlScriptString {
public:
};

QtObjectPtr QQmlScriptString_NewQQmlScriptString(){
	return new QQmlScriptString();
}

QtObjectPtr QQmlScriptString_NewQQmlScriptString2(QtObjectPtr other){
	return new QQmlScriptString(*static_cast<QQmlScriptString*>(other));
}

int QQmlScriptString_BooleanLiteral(QtObjectPtr ptr, int ok){
	return static_cast<QQmlScriptString*>(ptr)->booleanLiteral(NULL);
}

int QQmlScriptString_IsEmpty(QtObjectPtr ptr){
	return static_cast<QQmlScriptString*>(ptr)->isEmpty();
}

int QQmlScriptString_IsNullLiteral(QtObjectPtr ptr){
	return static_cast<QQmlScriptString*>(ptr)->isNullLiteral();
}

int QQmlScriptString_IsUndefinedLiteral(QtObjectPtr ptr){
	return static_cast<QQmlScriptString*>(ptr)->isUndefinedLiteral();
}

char* QQmlScriptString_StringLiteral(QtObjectPtr ptr){
	return static_cast<QQmlScriptString*>(ptr)->stringLiteral().toUtf8().data();
}

