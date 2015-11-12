#include "qxmlparseexception.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QXmlParseException>
#include "_cgo_export.h"

class MyQXmlParseException: public QXmlParseException {
public:
};

void* QXmlParseException_NewQXmlParseException(char* name, int c, int l, char* p, char* s){
	return new QXmlParseException(QString(name), c, l, QString(p), QString(s));
}

void* QXmlParseException_NewQXmlParseException2(void* other){
	return new QXmlParseException(*static_cast<QXmlParseException*>(other));
}

int QXmlParseException_ColumnNumber(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->columnNumber();
}

int QXmlParseException_LineNumber(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->lineNumber();
}

char* QXmlParseException_Message(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->message().toUtf8().data();
}

char* QXmlParseException_PublicId(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->publicId().toUtf8().data();
}

char* QXmlParseException_SystemId(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->systemId().toUtf8().data();
}

void QXmlParseException_DestroyQXmlParseException(void* ptr){
	static_cast<QXmlParseException*>(ptr)->~QXmlParseException();
}

