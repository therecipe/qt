#include "qxmlstreamreader.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QIODevice>
#include <QXmlStreamEntityResolver>
#include <QXmlStreamNamespaceDeclaration>
#include <QString>
#include <QXmlStreamReader>
#include "_cgo_export.h"

class MyQXmlStreamReader: public QXmlStreamReader {
public:
};

int QXmlStreamReader_NamespaceProcessing(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->namespaceProcessing();
}

void QXmlStreamReader_SetNamespaceProcessing(QtObjectPtr ptr, int v){
	static_cast<QXmlStreamReader*>(ptr)->setNamespaceProcessing(v != 0);
}

QtObjectPtr QXmlStreamReader_NewQXmlStreamReader(){
	return new QXmlStreamReader();
}

QtObjectPtr QXmlStreamReader_NewQXmlStreamReader2(QtObjectPtr device){
	return new QXmlStreamReader(static_cast<QIODevice*>(device));
}

QtObjectPtr QXmlStreamReader_NewQXmlStreamReader3(QtObjectPtr data){
	return new QXmlStreamReader(*static_cast<QByteArray*>(data));
}

QtObjectPtr QXmlStreamReader_NewQXmlStreamReader4(char* data){
	return new QXmlStreamReader(QString(data));
}

QtObjectPtr QXmlStreamReader_NewQXmlStreamReader5(char* data){
	return new QXmlStreamReader(const_cast<const char*>(data));
}

void QXmlStreamReader_AddData(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QXmlStreamReader*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QXmlStreamReader_AddData2(QtObjectPtr ptr, char* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(QString(data));
}

void QXmlStreamReader_AddData3(QtObjectPtr ptr, char* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(const_cast<const char*>(data));
}

void QXmlStreamReader_AddExtraNamespaceDeclaration(QtObjectPtr ptr, QtObjectPtr extraNamespaceDeclaration){
	static_cast<QXmlStreamReader*>(ptr)->addExtraNamespaceDeclaration(*static_cast<QXmlStreamNamespaceDeclaration*>(extraNamespaceDeclaration));
}

int QXmlStreamReader_AtEnd(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->atEnd();
}

void QXmlStreamReader_Clear(QtObjectPtr ptr){
	static_cast<QXmlStreamReader*>(ptr)->clear();
}

QtObjectPtr QXmlStreamReader_Device(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->device();
}

QtObjectPtr QXmlStreamReader_EntityResolver(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->entityResolver();
}

int QXmlStreamReader_Error(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->error();
}

char* QXmlStreamReader_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->errorString().toUtf8().data();
}

int QXmlStreamReader_HasError(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->hasError();
}

int QXmlStreamReader_IsCDATA(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isCDATA();
}

int QXmlStreamReader_IsCharacters(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isCharacters();
}

int QXmlStreamReader_IsComment(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isComment();
}

int QXmlStreamReader_IsDTD(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isDTD();
}

int QXmlStreamReader_IsEndDocument(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEndDocument();
}

int QXmlStreamReader_IsEndElement(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEndElement();
}

int QXmlStreamReader_IsEntityReference(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEntityReference();
}

int QXmlStreamReader_IsProcessingInstruction(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isProcessingInstruction();
}

int QXmlStreamReader_IsStandaloneDocument(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStandaloneDocument();
}

int QXmlStreamReader_IsStartDocument(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStartDocument();
}

int QXmlStreamReader_IsStartElement(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStartElement();
}

int QXmlStreamReader_IsWhitespace(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isWhitespace();
}

void QXmlStreamReader_RaiseError(QtObjectPtr ptr, char* message){
	static_cast<QXmlStreamReader*>(ptr)->raiseError(QString(message));
}

char* QXmlStreamReader_ReadElementText(QtObjectPtr ptr, int behaviour){
	return static_cast<QXmlStreamReader*>(ptr)->readElementText(static_cast<QXmlStreamReader::ReadElementTextBehaviour>(behaviour)).toUtf8().data();
}

int QXmlStreamReader_ReadNext(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->readNext();
}

int QXmlStreamReader_ReadNextStartElement(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->readNextStartElement();
}

void QXmlStreamReader_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QXmlStreamReader*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QXmlStreamReader_SetEntityResolver(QtObjectPtr ptr, QtObjectPtr resolver){
	static_cast<QXmlStreamReader*>(ptr)->setEntityResolver(static_cast<QXmlStreamEntityResolver*>(resolver));
}

void QXmlStreamReader_SkipCurrentElement(QtObjectPtr ptr){
	static_cast<QXmlStreamReader*>(ptr)->skipCurrentElement();
}

char* QXmlStreamReader_TokenString(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->tokenString().toUtf8().data();
}

int QXmlStreamReader_TokenType(QtObjectPtr ptr){
	return static_cast<QXmlStreamReader*>(ptr)->tokenType();
}

void QXmlStreamReader_DestroyQXmlStreamReader(QtObjectPtr ptr){
	static_cast<QXmlStreamReader*>(ptr)->~QXmlStreamReader();
}

