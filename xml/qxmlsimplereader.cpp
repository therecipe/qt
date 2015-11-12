#include "qxmlsimplereader.h"
#include <QModelIndex>
#include <QXmlContentHandler>
#include <QXmlInputSource>
#include <QXmlErrorHandler>
#include <QXmlDeclHandler>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QXmlLexicalHandler>
#include <QXmlDTDHandler>
#include <QXmlEntityResolver>
#include <QXmlSimpleReader>
#include "_cgo_export.h"

class MyQXmlSimpleReader: public QXmlSimpleReader {
public:
};

void* QXmlSimpleReader_DTDHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->DTDHandler();
}

void* QXmlSimpleReader_NewQXmlSimpleReader(){
	return new QXmlSimpleReader();
}

void* QXmlSimpleReader_ContentHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->contentHandler();
}

void* QXmlSimpleReader_DeclHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->declHandler();
}

void* QXmlSimpleReader_EntityResolver(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->entityResolver();
}

void* QXmlSimpleReader_ErrorHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->errorHandler();
}

int QXmlSimpleReader_Feature(void* ptr, char* name, int ok){
	return static_cast<QXmlSimpleReader*>(ptr)->feature(QString(name), NULL);
}

int QXmlSimpleReader_HasFeature(void* ptr, char* name){
	return static_cast<QXmlSimpleReader*>(ptr)->hasFeature(QString(name));
}

int QXmlSimpleReader_HasProperty(void* ptr, char* name){
	return static_cast<QXmlSimpleReader*>(ptr)->hasProperty(QString(name));
}

void* QXmlSimpleReader_LexicalHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->lexicalHandler();
}

int QXmlSimpleReader_Parse(void* ptr, void* input){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(*static_cast<QXmlInputSource*>(input));
}

int QXmlSimpleReader_Parse2(void* ptr, void* input){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input));
}

int QXmlSimpleReader_Parse3(void* ptr, void* input, int incremental){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input), incremental != 0);
}

int QXmlSimpleReader_ParseContinue(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->parseContinue();
}

void* QXmlSimpleReader_Property(void* ptr, char* name, int ok){
	return static_cast<QXmlSimpleReader*>(ptr)->property(QString(name), NULL);
}

void QXmlSimpleReader_SetContentHandler(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlSimpleReader_SetDTDHandler(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlSimpleReader_SetDeclHandler(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlSimpleReader_SetEntityResolver(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlSimpleReader_SetErrorHandler(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlSimpleReader_SetFeature(void* ptr, char* name, int enable){
	static_cast<QXmlSimpleReader*>(ptr)->setFeature(QString(name), enable != 0);
}

void QXmlSimpleReader_SetLexicalHandler(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlSimpleReader_DestroyQXmlSimpleReader(void* ptr){
	static_cast<QXmlSimpleReader*>(ptr)->~QXmlSimpleReader();
}

