#include "qxmlstreamreader.h"
#include <QXmlStreamEntityResolver>
#include <QXmlStreamNamespaceDeclaration>
#include <QByteArray>
#include <QIODevice>
#include <QString>
#include <QUrl>
#include <QStringRef>
#include <QVariant>
#include <QModelIndex>
#include <QXmlStreamReader>
#include "_cgo_export.h"

class MyQXmlStreamReader: public QXmlStreamReader {
public:
};

int QXmlStreamReader_NamespaceProcessing(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->namespaceProcessing();
}

void QXmlStreamReader_SetNamespaceProcessing(void* ptr, int v){
	static_cast<QXmlStreamReader*>(ptr)->setNamespaceProcessing(v != 0);
}

void* QXmlStreamReader_NewQXmlStreamReader(){
	return new QXmlStreamReader();
}

void* QXmlStreamReader_NewQXmlStreamReader2(void* device){
	return new QXmlStreamReader(static_cast<QIODevice*>(device));
}

void* QXmlStreamReader_NewQXmlStreamReader3(void* data){
	return new QXmlStreamReader(*static_cast<QByteArray*>(data));
}

void* QXmlStreamReader_NewQXmlStreamReader4(char* data){
	return new QXmlStreamReader(QString(data));
}

void* QXmlStreamReader_NewQXmlStreamReader5(char* data){
	return new QXmlStreamReader(const_cast<const char*>(data));
}

void QXmlStreamReader_AddData(void* ptr, void* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QXmlStreamReader_AddData2(void* ptr, char* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(QString(data));
}

void QXmlStreamReader_AddData3(void* ptr, char* data){
	static_cast<QXmlStreamReader*>(ptr)->addData(const_cast<const char*>(data));
}

void QXmlStreamReader_AddExtraNamespaceDeclaration(void* ptr, void* extraNamespaceDeclaration){
	static_cast<QXmlStreamReader*>(ptr)->addExtraNamespaceDeclaration(*static_cast<QXmlStreamNamespaceDeclaration*>(extraNamespaceDeclaration));
}

int QXmlStreamReader_AtEnd(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->atEnd();
}

void QXmlStreamReader_Clear(void* ptr){
	static_cast<QXmlStreamReader*>(ptr)->clear();
}

void* QXmlStreamReader_Device(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->device();
}

void* QXmlStreamReader_DocumentEncoding(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->documentEncoding());
}

void* QXmlStreamReader_DocumentVersion(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->documentVersion());
}

void* QXmlStreamReader_DtdName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->dtdName());
}

void* QXmlStreamReader_DtdPublicId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->dtdPublicId());
}

void* QXmlStreamReader_DtdSystemId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->dtdSystemId());
}

void* QXmlStreamReader_EntityResolver(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->entityResolver();
}

int QXmlStreamReader_Error(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->error();
}

char* QXmlStreamReader_ErrorString(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->errorString().toUtf8().data();
}

int QXmlStreamReader_HasError(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->hasError();
}

int QXmlStreamReader_IsCDATA(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isCDATA();
}

int QXmlStreamReader_IsCharacters(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isCharacters();
}

int QXmlStreamReader_IsComment(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isComment();
}

int QXmlStreamReader_IsDTD(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isDTD();
}

int QXmlStreamReader_IsEndDocument(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEndDocument();
}

int QXmlStreamReader_IsEndElement(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEndElement();
}

int QXmlStreamReader_IsEntityReference(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isEntityReference();
}

int QXmlStreamReader_IsProcessingInstruction(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isProcessingInstruction();
}

int QXmlStreamReader_IsStandaloneDocument(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStandaloneDocument();
}

int QXmlStreamReader_IsStartDocument(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStartDocument();
}

int QXmlStreamReader_IsStartElement(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isStartElement();
}

int QXmlStreamReader_IsWhitespace(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->isWhitespace();
}

void* QXmlStreamReader_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->name());
}

void* QXmlStreamReader_NamespaceUri(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->namespaceUri());
}

void* QXmlStreamReader_Prefix(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->prefix());
}

void* QXmlStreamReader_ProcessingInstructionData(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->processingInstructionData());
}

void* QXmlStreamReader_ProcessingInstructionTarget(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->processingInstructionTarget());
}

void* QXmlStreamReader_QualifiedName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->qualifiedName());
}

void QXmlStreamReader_RaiseError(void* ptr, char* message){
	static_cast<QXmlStreamReader*>(ptr)->raiseError(QString(message));
}

char* QXmlStreamReader_ReadElementText(void* ptr, int behaviour){
	return static_cast<QXmlStreamReader*>(ptr)->readElementText(static_cast<QXmlStreamReader::ReadElementTextBehaviour>(behaviour)).toUtf8().data();
}

int QXmlStreamReader_ReadNext(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->readNext();
}

int QXmlStreamReader_ReadNextStartElement(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->readNextStartElement();
}

void QXmlStreamReader_SetDevice(void* ptr, void* device){
	static_cast<QXmlStreamReader*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QXmlStreamReader_SetEntityResolver(void* ptr, void* resolver){
	static_cast<QXmlStreamReader*>(ptr)->setEntityResolver(static_cast<QXmlStreamEntityResolver*>(resolver));
}

void QXmlStreamReader_SkipCurrentElement(void* ptr){
	static_cast<QXmlStreamReader*>(ptr)->skipCurrentElement();
}

void* QXmlStreamReader_Text(void* ptr){
	return new QStringRef(static_cast<QXmlStreamReader*>(ptr)->text());
}

char* QXmlStreamReader_TokenString(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->tokenString().toUtf8().data();
}

int QXmlStreamReader_TokenType(void* ptr){
	return static_cast<QXmlStreamReader*>(ptr)->tokenType();
}

void QXmlStreamReader_DestroyQXmlStreamReader(void* ptr){
	static_cast<QXmlStreamReader*>(ptr)->~QXmlStreamReader();
}

