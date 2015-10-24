#include "qxmlreader.h"
#include <QXmlLexicalHandler>
#include <QXmlEntityResolver>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QXmlContentHandler>
#include <QXmlDTDHandler>
#include <QUrl>
#include <QXmlErrorHandler>
#include <QXmlInputSource>
#include <QXmlDeclHandler>
#include <QXmlReader>
#include "_cgo_export.h"

class MyQXmlReader: public QXmlReader {
public:
};

QtObjectPtr QXmlReader_DTDHandler(QtObjectPtr ptr){
	return static_cast<QXmlReader*>(ptr)->DTDHandler();
}

QtObjectPtr QXmlReader_ContentHandler(QtObjectPtr ptr){
	return static_cast<QXmlReader*>(ptr)->contentHandler();
}

QtObjectPtr QXmlReader_DeclHandler(QtObjectPtr ptr){
	return static_cast<QXmlReader*>(ptr)->declHandler();
}

QtObjectPtr QXmlReader_EntityResolver(QtObjectPtr ptr){
	return static_cast<QXmlReader*>(ptr)->entityResolver();
}

QtObjectPtr QXmlReader_ErrorHandler(QtObjectPtr ptr){
	return static_cast<QXmlReader*>(ptr)->errorHandler();
}

int QXmlReader_Feature(QtObjectPtr ptr, char* name, int ok){
	return static_cast<QXmlReader*>(ptr)->feature(QString(name), NULL);
}

int QXmlReader_HasFeature(QtObjectPtr ptr, char* name){
	return static_cast<QXmlReader*>(ptr)->hasFeature(QString(name));
}

int QXmlReader_HasProperty(QtObjectPtr ptr, char* name){
	return static_cast<QXmlReader*>(ptr)->hasProperty(QString(name));
}

QtObjectPtr QXmlReader_LexicalHandler(QtObjectPtr ptr){
	return static_cast<QXmlReader*>(ptr)->lexicalHandler();
}

int QXmlReader_Parse(QtObjectPtr ptr, QtObjectPtr input){
	return static_cast<QXmlReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input));
}

void QXmlReader_Property(QtObjectPtr ptr, char* name, int ok){
	static_cast<QXmlReader*>(ptr)->property(QString(name), NULL);
}

void QXmlReader_SetContentHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlReader_SetDTDHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlReader_SetDeclHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlReader_SetEntityResolver(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlReader_SetErrorHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlReader_SetFeature(QtObjectPtr ptr, char* name, int value){
	static_cast<QXmlReader*>(ptr)->setFeature(QString(name), value != 0);
}

void QXmlReader_SetLexicalHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlReader_DestroyQXmlReader(QtObjectPtr ptr){
	static_cast<QXmlReader*>(ptr)->~QXmlReader();
}

