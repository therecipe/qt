#include "qscriptprogram.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QScriptProgram>
#include "_cgo_export.h"

class MyQScriptProgram: public QScriptProgram {
public:
};

void* QScriptProgram_NewQScriptProgram(){
	return new QScriptProgram();
}

void* QScriptProgram_NewQScriptProgram3(void* other){
	return new QScriptProgram(*static_cast<QScriptProgram*>(other));
}

void* QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber){
	return new QScriptProgram(QString(sourceCode), QString(fileName), firstLineNumber);
}

char* QScriptProgram_FileName(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->fileName().toUtf8().data();
}

int QScriptProgram_FirstLineNumber(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->firstLineNumber();
}

int QScriptProgram_IsNull(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->isNull();
}

char* QScriptProgram_SourceCode(void* ptr){
	return static_cast<QScriptProgram*>(ptr)->sourceCode().toUtf8().data();
}

void QScriptProgram_DestroyQScriptProgram(void* ptr){
	static_cast<QScriptProgram*>(ptr)->~QScriptProgram();
}

