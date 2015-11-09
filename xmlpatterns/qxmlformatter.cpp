#include "qxmlformatter.h"
#include <QStringRef>
#include <QXmlQuery>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QIODevice>
#include <QXmlFormatter>
#include "_cgo_export.h"

class MyQXmlFormatter: public QXmlFormatter {
public:
};

void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice){
	return new QXmlFormatter(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlFormatter_Attribute(void* ptr, void* name, void* value){
	static_cast<QXmlFormatter*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlFormatter_Characters(void* ptr, void* value){
	static_cast<QXmlFormatter*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlFormatter_Comment(void* ptr, char* value){
	static_cast<QXmlFormatter*>(ptr)->comment(QString(value));
}

void QXmlFormatter_EndDocument(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->endDocument();
}

void QXmlFormatter_EndElement(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->endElement();
}

void QXmlFormatter_EndOfSequence(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->endOfSequence();
}

int QXmlFormatter_IndentationDepth(void* ptr){
	return static_cast<QXmlFormatter*>(ptr)->indentationDepth();
}

void QXmlFormatter_ProcessingInstruction(void* ptr, void* name, char* value){
	static_cast<QXmlFormatter*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString(value));
}

void QXmlFormatter_SetIndentationDepth(void* ptr, int depth){
	static_cast<QXmlFormatter*>(ptr)->setIndentationDepth(depth);
}

void QXmlFormatter_StartDocument(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->startDocument();
}

void QXmlFormatter_StartElement(void* ptr, void* name){
	static_cast<QXmlFormatter*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlFormatter_StartOfSequence(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->startOfSequence();
}

