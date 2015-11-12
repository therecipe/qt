#include "qtextlistformat.h"
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextList>
#include <QList>
#include <QTextListFormat>
#include "_cgo_export.h"

class MyQTextListFormat: public QTextListFormat {
public:
};

void* QTextListFormat_NewQTextListFormat(){
	return new QTextListFormat();
}

int QTextListFormat_Indent(void* ptr){
	return static_cast<QTextListFormat*>(ptr)->indent();
}

int QTextListFormat_IsValid(void* ptr){
	return static_cast<QTextListFormat*>(ptr)->isValid();
}

char* QTextListFormat_NumberPrefix(void* ptr){
	return static_cast<QTextListFormat*>(ptr)->numberPrefix().toUtf8().data();
}

char* QTextListFormat_NumberSuffix(void* ptr){
	return static_cast<QTextListFormat*>(ptr)->numberSuffix().toUtf8().data();
}

void QTextListFormat_SetIndent(void* ptr, int indentation){
	static_cast<QTextListFormat*>(ptr)->setIndent(indentation);
}

void QTextListFormat_SetNumberPrefix(void* ptr, char* numberPrefix){
	static_cast<QTextListFormat*>(ptr)->setNumberPrefix(QString(numberPrefix));
}

void QTextListFormat_SetNumberSuffix(void* ptr, char* numberSuffix){
	static_cast<QTextListFormat*>(ptr)->setNumberSuffix(QString(numberSuffix));
}

void QTextListFormat_SetStyle(void* ptr, int style){
	static_cast<QTextListFormat*>(ptr)->setStyle(static_cast<QTextListFormat::Style>(style));
}

int QTextListFormat_Style(void* ptr){
	return static_cast<QTextListFormat*>(ptr)->style();
}

