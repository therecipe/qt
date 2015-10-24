#include "qxmlformatter.h"
#include <QXmlName>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlQuery>
#include <QStringRef>
#include <QXmlFormatter>
#include "_cgo_export.h"

class MyQXmlFormatter: public QXmlFormatter {
public:
};

QtObjectPtr QXmlFormatter_NewQXmlFormatter(QtObjectPtr query, QtObjectPtr outputDevice){
	return new QXmlFormatter(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlFormatter_Attribute(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value){
	static_cast<QXmlFormatter*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlFormatter_Characters(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QXmlFormatter*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlFormatter_Comment(QtObjectPtr ptr, char* value){
	static_cast<QXmlFormatter*>(ptr)->comment(QString(value));
}

void QXmlFormatter_EndDocument(QtObjectPtr ptr){
	static_cast<QXmlFormatter*>(ptr)->endDocument();
}

void QXmlFormatter_EndElement(QtObjectPtr ptr){
	static_cast<QXmlFormatter*>(ptr)->endElement();
}

void QXmlFormatter_EndOfSequence(QtObjectPtr ptr){
	static_cast<QXmlFormatter*>(ptr)->endOfSequence();
}

int QXmlFormatter_IndentationDepth(QtObjectPtr ptr){
	return static_cast<QXmlFormatter*>(ptr)->indentationDepth();
}

void QXmlFormatter_ProcessingInstruction(QtObjectPtr ptr, QtObjectPtr name, char* value){
	static_cast<QXmlFormatter*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString(value));
}

void QXmlFormatter_SetIndentationDepth(QtObjectPtr ptr, int depth){
	static_cast<QXmlFormatter*>(ptr)->setIndentationDepth(depth);
}

void QXmlFormatter_StartDocument(QtObjectPtr ptr){
	static_cast<QXmlFormatter*>(ptr)->startDocument();
}

void QXmlFormatter_StartElement(QtObjectPtr ptr, QtObjectPtr name){
	static_cast<QXmlFormatter*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlFormatter_StartOfSequence(QtObjectPtr ptr){
	static_cast<QXmlFormatter*>(ptr)->startOfSequence();
}

