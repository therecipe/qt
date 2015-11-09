#include "qxmlstreamwriter.h"
#include <QString>
#include <QXmlStreamAttribute>
#include <QTextCodec>
#include <QXmlStreamAttributes>
#include <QIODevice>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlStreamReader>
#include <QXmlStreamWriter>
#include "_cgo_export.h"

class MyQXmlStreamWriter: public QXmlStreamWriter {
public:
};

int QXmlStreamWriter_AutoFormattingIndent(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->autoFormattingIndent();
}

void QXmlStreamWriter_SetAutoFormattingIndent(void* ptr, int spacesOrTabs){
	static_cast<QXmlStreamWriter*>(ptr)->setAutoFormattingIndent(spacesOrTabs);
}

int QXmlStreamWriter_AutoFormatting(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->autoFormatting();
}

void* QXmlStreamWriter_Codec(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->codec();
}

void* QXmlStreamWriter_Device(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->device();
}

int QXmlStreamWriter_HasError(void* ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->hasError();
}

void QXmlStreamWriter_SetAutoFormatting(void* ptr, int enable){
	static_cast<QXmlStreamWriter*>(ptr)->setAutoFormatting(enable != 0);
}

void QXmlStreamWriter_SetCodec(void* ptr, void* codec){
	static_cast<QXmlStreamWriter*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QXmlStreamWriter_SetCodec2(void* ptr, char* codecName){
	static_cast<QXmlStreamWriter*>(ptr)->setCodec(const_cast<const char*>(codecName));
}

void QXmlStreamWriter_SetDevice(void* ptr, void* device){
	static_cast<QXmlStreamWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QXmlStreamWriter_WriteAttribute(void* ptr, char* namespaceUri, char* name, char* value){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(QString(namespaceUri), QString(name), QString(value));
}

void QXmlStreamWriter_WriteAttribute2(void* ptr, char* qualifiedName, char* value){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(QString(qualifiedName), QString(value));
}

void QXmlStreamWriter_WriteAttribute3(void* ptr, void* attribute){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(*static_cast<QXmlStreamAttribute*>(attribute));
}

void QXmlStreamWriter_WriteAttributes(void* ptr, void* attributes){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttributes(*static_cast<QXmlStreamAttributes*>(attributes));
}

void QXmlStreamWriter_WriteCDATA(void* ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeCDATA(QString(text));
}

void QXmlStreamWriter_WriteCharacters(void* ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeCharacters(QString(text));
}

void QXmlStreamWriter_WriteComment(void* ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeComment(QString(text));
}

void QXmlStreamWriter_WriteCurrentToken(void* ptr, void* reader){
	static_cast<QXmlStreamWriter*>(ptr)->writeCurrentToken(*static_cast<QXmlStreamReader*>(reader));
}

void QXmlStreamWriter_WriteDTD(void* ptr, char* dtd){
	static_cast<QXmlStreamWriter*>(ptr)->writeDTD(QString(dtd));
}

void QXmlStreamWriter_WriteDefaultNamespace(void* ptr, char* namespaceUri){
	static_cast<QXmlStreamWriter*>(ptr)->writeDefaultNamespace(QString(namespaceUri));
}

void QXmlStreamWriter_WriteEmptyElement(void* ptr, char* namespaceUri, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeEmptyElement(QString(namespaceUri), QString(name));
}

void QXmlStreamWriter_WriteEmptyElement2(void* ptr, char* qualifiedName){
	static_cast<QXmlStreamWriter*>(ptr)->writeEmptyElement(QString(qualifiedName));
}

void QXmlStreamWriter_WriteEndDocument(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeEndDocument();
}

void QXmlStreamWriter_WriteEndElement(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeEndElement();
}

void QXmlStreamWriter_WriteEntityReference(void* ptr, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeEntityReference(QString(name));
}

void QXmlStreamWriter_WriteNamespace(void* ptr, char* namespaceUri, char* prefix){
	static_cast<QXmlStreamWriter*>(ptr)->writeNamespace(QString(namespaceUri), QString(prefix));
}

void QXmlStreamWriter_WriteProcessingInstruction(void* ptr, char* target, char* data){
	static_cast<QXmlStreamWriter*>(ptr)->writeProcessingInstruction(QString(target), QString(data));
}

void QXmlStreamWriter_WriteStartDocument3(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument();
}

void QXmlStreamWriter_WriteStartDocument(void* ptr, char* version){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument(QString(version));
}

void QXmlStreamWriter_WriteStartDocument2(void* ptr, char* version, int standalone){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument(QString(version), standalone != 0);
}

void QXmlStreamWriter_WriteStartElement(void* ptr, char* namespaceUri, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartElement(QString(namespaceUri), QString(name));
}

void QXmlStreamWriter_WriteStartElement2(void* ptr, char* qualifiedName){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartElement(QString(qualifiedName));
}

void QXmlStreamWriter_WriteTextElement(void* ptr, char* namespaceUri, char* name, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeTextElement(QString(namespaceUri), QString(name), QString(text));
}

void QXmlStreamWriter_WriteTextElement2(void* ptr, char* qualifiedName, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeTextElement(QString(qualifiedName), QString(text));
}

void QXmlStreamWriter_DestroyQXmlStreamWriter(void* ptr){
	static_cast<QXmlStreamWriter*>(ptr)->~QXmlStreamWriter();
}

