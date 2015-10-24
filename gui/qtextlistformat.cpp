#include "qtextlistformat.h"
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QList>
#include <QTextList>
#include <QTextListFormat>
#include "_cgo_export.h"

class MyQTextListFormat: public QTextListFormat {
public:
};

QtObjectPtr QTextListFormat_NewQTextListFormat(){
	return new QTextListFormat();
}

int QTextListFormat_Indent(QtObjectPtr ptr){
	return static_cast<QTextListFormat*>(ptr)->indent();
}

int QTextListFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextListFormat*>(ptr)->isValid();
}

char* QTextListFormat_NumberPrefix(QtObjectPtr ptr){
	return static_cast<QTextListFormat*>(ptr)->numberPrefix().toUtf8().data();
}

char* QTextListFormat_NumberSuffix(QtObjectPtr ptr){
	return static_cast<QTextListFormat*>(ptr)->numberSuffix().toUtf8().data();
}

void QTextListFormat_SetIndent(QtObjectPtr ptr, int indentation){
	static_cast<QTextListFormat*>(ptr)->setIndent(indentation);
}

void QTextListFormat_SetNumberPrefix(QtObjectPtr ptr, char* numberPrefix){
	static_cast<QTextListFormat*>(ptr)->setNumberPrefix(QString(numberPrefix));
}

void QTextListFormat_SetNumberSuffix(QtObjectPtr ptr, char* numberSuffix){
	static_cast<QTextListFormat*>(ptr)->setNumberSuffix(QString(numberSuffix));
}

void QTextListFormat_SetStyle(QtObjectPtr ptr, int style){
	static_cast<QTextListFormat*>(ptr)->setStyle(static_cast<QTextListFormat::Style>(style));
}

int QTextListFormat_Style(QtObjectPtr ptr){
	return static_cast<QTextListFormat*>(ptr)->style();
}

