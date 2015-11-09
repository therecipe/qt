#include "qxmlcontenthandler.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlAttributes>
#include <QXmlLocator>
#include <QString>
#include <QVariant>
#include <QXmlContentHandler>
#include "_cgo_export.h"

class MyQXmlContentHandler: public QXmlContentHandler {
public:
};

int QXmlContentHandler_Characters(void* ptr, char* ch){
	return static_cast<QXmlContentHandler*>(ptr)->characters(QString(ch));
}

int QXmlContentHandler_EndDocument(void* ptr){
	return static_cast<QXmlContentHandler*>(ptr)->endDocument();
}

int QXmlContentHandler_EndElement(void* ptr, char* namespaceURI, char* localName, char* qName){
	return static_cast<QXmlContentHandler*>(ptr)->endElement(QString(namespaceURI), QString(localName), QString(qName));
}

int QXmlContentHandler_EndPrefixMapping(void* ptr, char* prefix){
	return static_cast<QXmlContentHandler*>(ptr)->endPrefixMapping(QString(prefix));
}

char* QXmlContentHandler_ErrorString(void* ptr){
	return static_cast<QXmlContentHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlContentHandler_IgnorableWhitespace(void* ptr, char* ch){
	return static_cast<QXmlContentHandler*>(ptr)->ignorableWhitespace(QString(ch));
}

int QXmlContentHandler_ProcessingInstruction(void* ptr, char* target, char* data){
	return static_cast<QXmlContentHandler*>(ptr)->processingInstruction(QString(target), QString(data));
}

void QXmlContentHandler_SetDocumentLocator(void* ptr, void* locator){
	static_cast<QXmlContentHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

int QXmlContentHandler_SkippedEntity(void* ptr, char* name){
	return static_cast<QXmlContentHandler*>(ptr)->skippedEntity(QString(name));
}

int QXmlContentHandler_StartDocument(void* ptr){
	return static_cast<QXmlContentHandler*>(ptr)->startDocument();
}

int QXmlContentHandler_StartElement(void* ptr, char* namespaceURI, char* localName, char* qName, void* atts){
	return static_cast<QXmlContentHandler*>(ptr)->startElement(QString(namespaceURI), QString(localName), QString(qName), *static_cast<QXmlAttributes*>(atts));
}

int QXmlContentHandler_StartPrefixMapping(void* ptr, char* prefix, char* uri){
	return static_cast<QXmlContentHandler*>(ptr)->startPrefixMapping(QString(prefix), QString(uri));
}

void QXmlContentHandler_DestroyQXmlContentHandler(void* ptr){
	static_cast<QXmlContentHandler*>(ptr)->~QXmlContentHandler();
}

