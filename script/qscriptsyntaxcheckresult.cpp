#include "qscriptsyntaxcheckresult.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QScriptSyntaxCheckResult>
#include "_cgo_export.h"

class MyQScriptSyntaxCheckResult: public QScriptSyntaxCheckResult {
public:
};

QtObjectPtr QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(QtObjectPtr other){
	return new QScriptSyntaxCheckResult(*static_cast<QScriptSyntaxCheckResult*>(other));
}

int QScriptSyntaxCheckResult_ErrorColumnNumber(QtObjectPtr ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorColumnNumber();
}

int QScriptSyntaxCheckResult_ErrorLineNumber(QtObjectPtr ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorLineNumber();
}

char* QScriptSyntaxCheckResult_ErrorMessage(QtObjectPtr ptr){
	return static_cast<QScriptSyntaxCheckResult*>(ptr)->errorMessage().toUtf8().data();
}

void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(QtObjectPtr ptr){
	static_cast<QScriptSyntaxCheckResult*>(ptr)->~QScriptSyntaxCheckResult();
}

