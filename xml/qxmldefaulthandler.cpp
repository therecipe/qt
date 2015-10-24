#include "qxmldefaulthandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlLocator>
#include <QXmlParseException>
#include <QXmlAttributes>
#include <QXmlDefaultHandler>
#include "_cgo_export.h"

class MyQXmlDefaultHandler: public QXmlDefaultHandler {
public:
};

QtObjectPtr QXmlDefaultHandler_NewQXmlDefaultHandler(){
	return new QXmlDefaultHandler();
}

void QXmlDefaultHandler_DestroyQXmlDefaultHandler(QtObjectPtr ptr){
	static_cast<QXmlDefaultHandler*>(ptr)->~QXmlDefaultHandler();
}

int QXmlDefaultHandler_AttributeDecl(QtObjectPtr ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value){
	return static_cast<QXmlDefaultHandler*>(ptr)->attributeDecl(QString(eName), QString(aName), QString(ty), QString(valueDefault), QString(value));
}

int QXmlDefaultHandler_Characters(QtObjectPtr ptr, char* ch){
	return static_cast<QXmlDefaultHandler*>(ptr)->characters(QString(ch));
}

int QXmlDefaultHandler_Comment(QtObjectPtr ptr, char* ch){
	return static_cast<QXmlDefaultHandler*>(ptr)->comment(QString(ch));
}

int QXmlDefaultHandler_EndCDATA(QtObjectPtr ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->endCDATA();
}

int QXmlDefaultHandler_EndDTD(QtObjectPtr ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->endDTD();
}

int QXmlDefaultHandler_EndDocument(QtObjectPtr ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->endDocument();
}

int QXmlDefaultHandler_EndElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName){
	return static_cast<QXmlDefaultHandler*>(ptr)->endElement(QString(namespaceURI), QString(localName), QString(qName));
}

int QXmlDefaultHandler_EndEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlDefaultHandler*>(ptr)->endEntity(QString(name));
}

int QXmlDefaultHandler_EndPrefixMapping(QtObjectPtr ptr, char* prefix){
	return static_cast<QXmlDefaultHandler*>(ptr)->endPrefixMapping(QString(prefix));
}

int QXmlDefaultHandler_Error(QtObjectPtr ptr, QtObjectPtr exception){
	return static_cast<QXmlDefaultHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

char* QXmlDefaultHandler_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlDefaultHandler_ExternalEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDefaultHandler*>(ptr)->externalEntityDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDefaultHandler_FatalError(QtObjectPtr ptr, QtObjectPtr exception){
	return static_cast<QXmlDefaultHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

int QXmlDefaultHandler_IgnorableWhitespace(QtObjectPtr ptr, char* ch){
	return static_cast<QXmlDefaultHandler*>(ptr)->ignorableWhitespace(QString(ch));
}

int QXmlDefaultHandler_InternalEntityDecl(QtObjectPtr ptr, char* name, char* value){
	return static_cast<QXmlDefaultHandler*>(ptr)->internalEntityDecl(QString(name), QString(value));
}

int QXmlDefaultHandler_NotationDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDefaultHandler*>(ptr)->notationDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDefaultHandler_ProcessingInstruction(QtObjectPtr ptr, char* target, char* data){
	return static_cast<QXmlDefaultHandler*>(ptr)->processingInstruction(QString(target), QString(data));
}

void QXmlDefaultHandler_SetDocumentLocator(QtObjectPtr ptr, QtObjectPtr locator){
	static_cast<QXmlDefaultHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

int QXmlDefaultHandler_SkippedEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlDefaultHandler*>(ptr)->skippedEntity(QString(name));
}

int QXmlDefaultHandler_StartCDATA(QtObjectPtr ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->startCDATA();
}

int QXmlDefaultHandler_StartDTD(QtObjectPtr ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDefaultHandler*>(ptr)->startDTD(QString(name), QString(publicId), QString(systemId));
}

int QXmlDefaultHandler_StartDocument(QtObjectPtr ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->startDocument();
}

int QXmlDefaultHandler_StartElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName, QtObjectPtr atts){
	return static_cast<QXmlDefaultHandler*>(ptr)->startElement(QString(namespaceURI), QString(localName), QString(qName), *static_cast<QXmlAttributes*>(atts));
}

int QXmlDefaultHandler_StartEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlDefaultHandler*>(ptr)->startEntity(QString(name));
}

int QXmlDefaultHandler_StartPrefixMapping(QtObjectPtr ptr, char* prefix, char* uri){
	return static_cast<QXmlDefaultHandler*>(ptr)->startPrefixMapping(QString(prefix), QString(uri));
}

int QXmlDefaultHandler_UnparsedEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId, char* notationName){
	return static_cast<QXmlDefaultHandler*>(ptr)->unparsedEntityDecl(QString(name), QString(publicId), QString(systemId), QString(notationName));
}

int QXmlDefaultHandler_Warning(QtObjectPtr ptr, QtObjectPtr exception){
	return static_cast<QXmlDefaultHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

