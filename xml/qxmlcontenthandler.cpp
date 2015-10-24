#include "qxmlcontenthandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlAttributes>
#include <QXmlLocator>
#include <QXmlContentHandler>
#include "_cgo_export.h"

class MyQXmlContentHandler: public QXmlContentHandler {
public:
};

int QXmlContentHandler_Characters(QtObjectPtr ptr, char* ch){
	return static_cast<QXmlContentHandler*>(ptr)->characters(QString(ch));
}

int QXmlContentHandler_EndDocument(QtObjectPtr ptr){
	return static_cast<QXmlContentHandler*>(ptr)->endDocument();
}

int QXmlContentHandler_EndElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName){
	return static_cast<QXmlContentHandler*>(ptr)->endElement(QString(namespaceURI), QString(localName), QString(qName));
}

int QXmlContentHandler_EndPrefixMapping(QtObjectPtr ptr, char* prefix){
	return static_cast<QXmlContentHandler*>(ptr)->endPrefixMapping(QString(prefix));
}

char* QXmlContentHandler_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlContentHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlContentHandler_IgnorableWhitespace(QtObjectPtr ptr, char* ch){
	return static_cast<QXmlContentHandler*>(ptr)->ignorableWhitespace(QString(ch));
}

int QXmlContentHandler_ProcessingInstruction(QtObjectPtr ptr, char* target, char* data){
	return static_cast<QXmlContentHandler*>(ptr)->processingInstruction(QString(target), QString(data));
}

void QXmlContentHandler_SetDocumentLocator(QtObjectPtr ptr, QtObjectPtr locator){
	static_cast<QXmlContentHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

int QXmlContentHandler_SkippedEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlContentHandler*>(ptr)->skippedEntity(QString(name));
}

int QXmlContentHandler_StartDocument(QtObjectPtr ptr){
	return static_cast<QXmlContentHandler*>(ptr)->startDocument();
}

int QXmlContentHandler_StartElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName, QtObjectPtr atts){
	return static_cast<QXmlContentHandler*>(ptr)->startElement(QString(namespaceURI), QString(localName), QString(qName), *static_cast<QXmlAttributes*>(atts));
}

int QXmlContentHandler_StartPrefixMapping(QtObjectPtr ptr, char* prefix, char* uri){
	return static_cast<QXmlContentHandler*>(ptr)->startPrefixMapping(QString(prefix), QString(uri));
}

void QXmlContentHandler_DestroyQXmlContentHandler(QtObjectPtr ptr){
	static_cast<QXmlContentHandler*>(ptr)->~QXmlContentHandler();
}

