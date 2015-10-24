#include "qtextimageformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextImageFormat>
#include "_cgo_export.h"

class MyQTextImageFormat: public QTextImageFormat {
public:
};

QtObjectPtr QTextImageFormat_NewQTextImageFormat(){
	return new QTextImageFormat();
}

int QTextImageFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextImageFormat*>(ptr)->isValid();
}

char* QTextImageFormat_Name(QtObjectPtr ptr){
	return static_cast<QTextImageFormat*>(ptr)->name().toUtf8().data();
}

void QTextImageFormat_SetName(QtObjectPtr ptr, char* name){
	static_cast<QTextImageFormat*>(ptr)->setName(QString(name));
}

