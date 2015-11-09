#include "qxmlreader.h"
#include <QXmlContentHandler>
#include <QModelIndex>
#include <QXmlDTDHandler>
#include <QXmlLexicalHandler>
#include <QXmlDeclHandler>
#include <QXmlInputSource>
#include <QXmlErrorHandler>
#include <QXmlEntityResolver>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QXmlReader>
#include "_cgo_export.h"

class MyQXmlReader: public QXmlReader {
public:
};

void* QXmlReader_DTDHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->DTDHandler();
}

void* QXmlReader_ContentHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->contentHandler();
}

void* QXmlReader_DeclHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->declHandler();
}

void* QXmlReader_EntityResolver(void* ptr){
	return static_cast<QXmlReader*>(ptr)->entityResolver();
}

void* QXmlReader_ErrorHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->errorHandler();
}

int QXmlReader_Feature(void* ptr, char* name, int ok){
	return static_cast<QXmlReader*>(ptr)->feature(QString(name), NULL);
}

int QXmlReader_HasFeature(void* ptr, char* name){
	return static_cast<QXmlReader*>(ptr)->hasFeature(QString(name));
}

int QXmlReader_HasProperty(void* ptr, char* name){
	return static_cast<QXmlReader*>(ptr)->hasProperty(QString(name));
}

void* QXmlReader_LexicalHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->lexicalHandler();
}

int QXmlReader_Parse(void* ptr, void* input){
	return static_cast<QXmlReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input));
}

void* QXmlReader_Property(void* ptr, char* name, int ok){
	return static_cast<QXmlReader*>(ptr)->property(QString(name), NULL);
}

void QXmlReader_SetContentHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlReader_SetDTDHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlReader_SetDeclHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlReader_SetEntityResolver(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlReader_SetErrorHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlReader_SetFeature(void* ptr, char* name, int value){
	static_cast<QXmlReader*>(ptr)->setFeature(QString(name), value != 0);
}

void QXmlReader_SetLexicalHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlReader_DestroyQXmlReader(void* ptr){
	static_cast<QXmlReader*>(ptr)->~QXmlReader();
}

