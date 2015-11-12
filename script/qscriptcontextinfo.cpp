#include "qscriptcontextinfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QScriptContext>
#include <QString>
#include <QVariant>
#include <QScriptContextInfo>
#include "_cgo_export.h"

class MyQScriptContextInfo: public QScriptContextInfo {
public:
};

void* QScriptContextInfo_NewQScriptContextInfo3(){
	return new QScriptContextInfo();
}

void* QScriptContextInfo_NewQScriptContextInfo(void* context){
	return new QScriptContextInfo(static_cast<QScriptContext*>(context));
}

void* QScriptContextInfo_NewQScriptContextInfo2(void* other){
	return new QScriptContextInfo(*static_cast<QScriptContextInfo*>(other));
}

char* QScriptContextInfo_FileName(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->fileName().toUtf8().data();
}

int QScriptContextInfo_FunctionEndLineNumber(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionEndLineNumber();
}

int QScriptContextInfo_FunctionMetaIndex(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionMetaIndex();
}

char* QScriptContextInfo_FunctionName(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionName().toUtf8().data();
}

char* QScriptContextInfo_FunctionParameterNames(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionParameterNames().join("|").toUtf8().data();
}

int QScriptContextInfo_FunctionStartLineNumber(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionStartLineNumber();
}

int QScriptContextInfo_FunctionType(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->functionType();
}

int QScriptContextInfo_IsNull(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->isNull();
}

int QScriptContextInfo_LineNumber(void* ptr){
	return static_cast<QScriptContextInfo*>(ptr)->lineNumber();
}

void QScriptContextInfo_DestroyQScriptContextInfo(void* ptr){
	static_cast<QScriptContextInfo*>(ptr)->~QScriptContextInfo();
}

