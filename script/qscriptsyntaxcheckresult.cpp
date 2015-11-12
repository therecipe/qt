#include "qscriptsyntaxcheckresult.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QScriptSyntaxCheckResult>
#include "_cgo_export.h"

class MyQScriptSyntaxCheckResult: public QScriptSyntaxCheckResult {
public:
};

void* QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(void* other){
	return new QScriptSyntaxCheckResult(*static_cast<QScriptSyntaxCheckResult*>(other));
}

int QScriptSyntaxCheckResult_ErrorColumnNumber(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorColumnNumber();
}

int QScriptSyntaxCheckResult_ErrorLineNumber(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorLineNumber();
}

char* QScriptSyntaxCheckResult_ErrorMessage(void* ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorMessage().toUtf8().data();
}

void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(void* ptr){
	static_cast<QScriptSyntaxCheckResult*>(ptr)->~QScriptSyntaxCheckResult();
}

