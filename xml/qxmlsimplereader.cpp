#include "qxmlsimplereader.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlErrorHandler>
#include <QXmlEntityResolver>
#include <QXmlDTDHandler>
#include <QXmlContentHandler>
#include <QString>
#include <QXmlDeclHandler>
#include <QXmlLexicalHandler>
#include <QXmlInputSource>
#include <QXmlSimpleReader>
#include "_cgo_export.h"

class MyQXmlSimpleReader: public QXmlSimpleReader {
public:
};

QtObjectPtr QXmlSimpleReader_DTDHandler(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->DTDHandler();
}

QtObjectPtr QXmlSimpleReader_NewQXmlSimpleReader(){
	return new QXmlSimpleReader();
}

QtObjectPtr QXmlSimpleReader_ContentHandler(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->contentHandler();
}

QtObjectPtr QXmlSimpleReader_DeclHandler(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->declHandler();
}

QtObjectPtr QXmlSimpleReader_EntityResolver(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->entityResolver();
}

QtObjectPtr QXmlSimpleReader_ErrorHandler(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->errorHandler();
}

int QXmlSimpleReader_Feature(QtObjectPtr ptr, char* name, int ok){
	return static_cast<QXmlSimpleReader*>(ptr)->feature(QString(name), NULL);
}

int QXmlSimpleReader_HasFeature(QtObjectPtr ptr, char* name){
	return static_cast<QXmlSimpleReader*>(ptr)->hasFeature(QString(name));
}

int QXmlSimpleReader_HasProperty(QtObjectPtr ptr, char* name){
	return static_cast<QXmlSimpleReader*>(ptr)->hasProperty(QString(name));
}

QtObjectPtr QXmlSimpleReader_LexicalHandler(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->lexicalHandler();
}

int QXmlSimpleReader_Parse(QtObjectPtr ptr, QtObjectPtr input){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(*static_cast<QXmlInputSource*>(input));
}

int QXmlSimpleReader_Parse2(QtObjectPtr ptr, QtObjectPtr input){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input));
}

int QXmlSimpleReader_Parse3(QtObjectPtr ptr, QtObjectPtr input, int incremental){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input), incremental != 0);
}

int QXmlSimpleReader_ParseContinue(QtObjectPtr ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->parseContinue();
}

void QXmlSimpleReader_Property(QtObjectPtr ptr, char* name, int ok){
	static_cast<QXmlSimpleReader*>(ptr)->property(QString(name), NULL);
}

void QXmlSimpleReader_SetContentHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSimpleReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlSimpleReader_SetDTDHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSimpleReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlSimpleReader_SetDeclHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSimpleReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlSimpleReader_SetEntityResolver(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSimpleReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlSimpleReader_SetErrorHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSimpleReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlSimpleReader_SetFeature(QtObjectPtr ptr, char* name, int enable){
	static_cast<QXmlSimpleReader*>(ptr)->setFeature(QString(name), enable != 0);
}

void QXmlSimpleReader_SetLexicalHandler(QtObjectPtr ptr, QtObjectPtr handler){
	static_cast<QXmlSimpleReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlSimpleReader_DestroyQXmlSimpleReader(QtObjectPtr ptr){
	static_cast<QXmlSimpleReader*>(ptr)->~QXmlSimpleReader();
}

