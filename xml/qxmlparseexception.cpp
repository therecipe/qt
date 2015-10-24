#include "qxmlparseexception.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlParseException>
#include "_cgo_export.h"

class MyQXmlParseException: public QXmlParseException {
public:
};

QtObjectPtr QXmlParseException_NewQXmlParseException(char* name, int c, int l, char* p, char* s){
	return new QXmlParseException(QString(name), c, l, QString(p), QString(s));
}

QtObjectPtr QXmlParseException_NewQXmlParseException2(QtObjectPtr other){
	return new QXmlParseException(*static_cast<QXmlParseException*>(other));
}

int QXmlParseException_ColumnNumber(QtObjectPtr ptr){
	return static_cast<QXmlParseException*>(ptr)->columnNumber();
}

int QXmlParseException_LineNumber(QtObjectPtr ptr){
	return static_cast<QXmlParseException*>(ptr)->lineNumber();
}

char* QXmlParseException_Message(QtObjectPtr ptr){
	return static_cast<QXmlParseException*>(ptr)->message().toUtf8().data();
}

char* QXmlParseException_PublicId(QtObjectPtr ptr){
	return static_cast<QXmlParseException*>(ptr)->publicId().toUtf8().data();
}

char* QXmlParseException_SystemId(QtObjectPtr ptr){
	return static_cast<QXmlParseException*>(ptr)->systemId().toUtf8().data();
}

void QXmlParseException_DestroyQXmlParseException(QtObjectPtr ptr){
	static_cast<QXmlParseException*>(ptr)->~QXmlParseException();
}

