#include "qscriptstring.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptString>
#include "_cgo_export.h"

class MyQScriptString: public QScriptString {
public:
};

QtObjectPtr QScriptString_NewQScriptString(){
	return new QScriptString();
}

QtObjectPtr QScriptString_NewQScriptString2(QtObjectPtr other){
	return new QScriptString(*static_cast<QScriptString*>(other));
}

int QScriptString_IsValid(QtObjectPtr ptr){
	return static_cast<QScriptString*>(ptr)->isValid();
}

char* QScriptString_ToString(QtObjectPtr ptr){
	return static_cast<QScriptString*>(ptr)->toString().toUtf8().data();
}

void QScriptString_DestroyQScriptString(QtObjectPtr ptr){
	static_cast<QScriptString*>(ptr)->~QScriptString();
}

