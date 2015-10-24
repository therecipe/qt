#include "qscriptcontextinfo.h"
#include <QModelIndex>
#include <QScriptContext>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QScriptContextInfo>
#include "_cgo_export.h"

class MyQScriptContextInfo: public QScriptContextInfo {
public:
};

QtObjectPtr QScriptContextInfo_NewQScriptContextInfo3(){
	return new QScriptContextInfo();
}

QtObjectPtr QScriptContextInfo_NewQScriptContextInfo(QtObjectPtr context){
	return new QScriptContextInfo(static_cast<QScriptContext*>(context));
}

QtObjectPtr QScriptContextInfo_NewQScriptContextInfo2(QtObjectPtr other){
	return new QScriptContextInfo(*static_cast<QScriptContextInfo*>(other));
}

char* QScriptContextInfo_FileName(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->fileName().toUtf8().data();
}

int QScriptContextInfo_FunctionEndLineNumber(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionEndLineNumber();
}

int QScriptContextInfo_FunctionMetaIndex(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionMetaIndex();
}

char* QScriptContextInfo_FunctionName(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionName().toUtf8().data();
}

char* QScriptContextInfo_FunctionParameterNames(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionParameterNames().join("|").toUtf8().data();
}

int QScriptContextInfo_FunctionStartLineNumber(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionStartLineNumber();
}

int QScriptContextInfo_FunctionType(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionType();
}

int QScriptContextInfo_IsNull(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->isNull();
}

int QScriptContextInfo_LineNumber(QtObjectPtr ptr){
	return static_cast<QScriptContextInfo*>(ptr)->lineNumber();
}

void QScriptContextInfo_DestroyQScriptContextInfo(QtObjectPtr ptr){
	static_cast<QScriptContextInfo*>(ptr)->~QScriptContextInfo();
}

