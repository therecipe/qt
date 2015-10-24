#include "qscriptprogram.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptProgram>
#include "_cgo_export.h"

class MyQScriptProgram: public QScriptProgram {
public:
};

QtObjectPtr QScriptProgram_NewQScriptProgram(){
	return new QScriptProgram();
}

QtObjectPtr QScriptProgram_NewQScriptProgram3(QtObjectPtr other){
	return new QScriptProgram(*static_cast<QScriptProgram*>(other));
}

QtObjectPtr QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber){
	return new QScriptProgram(QString(sourceCode), QString(fileName), firstLineNumber);
}

char* QScriptProgram_FileName(QtObjectPtr ptr){
	return static_cast<QScriptProgram*>(ptr)->fileName().toUtf8().data();
}

int QScriptProgram_FirstLineNumber(QtObjectPtr ptr){
	return static_cast<QScriptProgram*>(ptr)->firstLineNumber();
}

int QScriptProgram_IsNull(QtObjectPtr ptr){
	return static_cast<QScriptProgram*>(ptr)->isNull();
}

char* QScriptProgram_SourceCode(QtObjectPtr ptr){
	return static_cast<QScriptProgram*>(ptr)->sourceCode().toUtf8().data();
}

void QScriptProgram_DestroyQScriptProgram(QtObjectPtr ptr){
	static_cast<QScriptProgram*>(ptr)->~QScriptProgram();
}

