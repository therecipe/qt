#include "qxmlstreamwriter.h"
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QXmlStreamReader>
#include <QString>
#include <QVariant>
#include <QTextCodec>
#include <QXmlStreamAttributes>
#include <QXmlStreamAttribute>
#include <QXmlStreamWriter>
#include "_cgo_export.h"

class MyQXmlStreamWriter: public QXmlStreamWriter {
public:
};

int QXmlStreamWriter_AutoFormattingIndent(QtObjectPtr ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->autoFormattingIndent();
}

void QXmlStreamWriter_SetAutoFormattingIndent(QtObjectPtr ptr, int spacesOrTabs){
	static_cast<QXmlStreamWriter*>(ptr)->setAutoFormattingIndent(spacesOrTabs);
}

int QXmlStreamWriter_AutoFormatting(QtObjectPtr ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->autoFormatting();
}

QtObjectPtr QXmlStreamWriter_Codec(QtObjectPtr ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->codec();
}

QtObjectPtr QXmlStreamWriter_Device(QtObjectPtr ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->device();
}

int QXmlStreamWriter_HasError(QtObjectPtr ptr){
	return static_cast<QXmlStreamWriter*>(ptr)->hasError();
}

void QXmlStreamWriter_SetAutoFormatting(QtObjectPtr ptr, int enable){
	static_cast<QXmlStreamWriter*>(ptr)->setAutoFormatting(enable != 0);
}

void QXmlStreamWriter_SetCodec(QtObjectPtr ptr, QtObjectPtr codec){
	static_cast<QXmlStreamWriter*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QXmlStreamWriter_SetCodec2(QtObjectPtr ptr, char* codecName){
	static_cast<QXmlStreamWriter*>(ptr)->setCodec(const_cast<const char*>(codecName));
}

void QXmlStreamWriter_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QXmlStreamWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QXmlStreamWriter_WriteAttribute(QtObjectPtr ptr, char* namespaceUri, char* name, char* value){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(QString(namespaceUri), QString(name), QString(value));
}

void QXmlStreamWriter_WriteAttribute2(QtObjectPtr ptr, char* qualifiedName, char* value){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(QString(qualifiedName), QString(value));
}

void QXmlStreamWriter_WriteAttribute3(QtObjectPtr ptr, QtObjectPtr attribute){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttribute(*static_cast<QXmlStreamAttribute*>(attribute));
}

void QXmlStreamWriter_WriteAttributes(QtObjectPtr ptr, QtObjectPtr attributes){
	static_cast<QXmlStreamWriter*>(ptr)->writeAttributes(*static_cast<QXmlStreamAttributes*>(attributes));
}

void QXmlStreamWriter_WriteCDATA(QtObjectPtr ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeCDATA(QString(text));
}

void QXmlStreamWriter_WriteCharacters(QtObjectPtr ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeCharacters(QString(text));
}

void QXmlStreamWriter_WriteComment(QtObjectPtr ptr, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeComment(QString(text));
}

void QXmlStreamWriter_WriteCurrentToken(QtObjectPtr ptr, QtObjectPtr reader){
	static_cast<QXmlStreamWriter*>(ptr)->writeCurrentToken(*static_cast<QXmlStreamReader*>(reader));
}

void QXmlStreamWriter_WriteDTD(QtObjectPtr ptr, char* dtd){
	static_cast<QXmlStreamWriter*>(ptr)->writeDTD(QString(dtd));
}

void QXmlStreamWriter_WriteDefaultNamespace(QtObjectPtr ptr, char* namespaceUri){
	static_cast<QXmlStreamWriter*>(ptr)->writeDefaultNamespace(QString(namespaceUri));
}

void QXmlStreamWriter_WriteEmptyElement(QtObjectPtr ptr, char* namespaceUri, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeEmptyElement(QString(namespaceUri), QString(name));
}

void QXmlStreamWriter_WriteEmptyElement2(QtObjectPtr ptr, char* qualifiedName){
	static_cast<QXmlStreamWriter*>(ptr)->writeEmptyElement(QString(qualifiedName));
}

void QXmlStreamWriter_WriteEndDocument(QtObjectPtr ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeEndDocument();
}

void QXmlStreamWriter_WriteEndElement(QtObjectPtr ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeEndElement();
}

void QXmlStreamWriter_WriteEntityReference(QtObjectPtr ptr, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeEntityReference(QString(name));
}

void QXmlStreamWriter_WriteNamespace(QtObjectPtr ptr, char* namespaceUri, char* prefix){
	static_cast<QXmlStreamWriter*>(ptr)->writeNamespace(QString(namespaceUri), QString(prefix));
}

void QXmlStreamWriter_WriteProcessingInstruction(QtObjectPtr ptr, char* target, char* data){
	static_cast<QXmlStreamWriter*>(ptr)->writeProcessingInstruction(QString(target), QString(data));
}

void QXmlStreamWriter_WriteStartDocument3(QtObjectPtr ptr){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument();
}

void QXmlStreamWriter_WriteStartDocument(QtObjectPtr ptr, char* version){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument(QString(version));
}

void QXmlStreamWriter_WriteStartDocument2(QtObjectPtr ptr, char* version, int standalone){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartDocument(QString(version), standalone != 0);
}

void QXmlStreamWriter_WriteStartElement(QtObjectPtr ptr, char* namespaceUri, char* name){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartElement(QString(namespaceUri), QString(name));
}

void QXmlStreamWriter_WriteStartElement2(QtObjectPtr ptr, char* qualifiedName){
	static_cast<QXmlStreamWriter*>(ptr)->writeStartElement(QString(qualifiedName));
}

void QXmlStreamWriter_WriteTextElement(QtObjectPtr ptr, char* namespaceUri, char* name, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeTextElement(QString(namespaceUri), QString(name), QString(text));
}

void QXmlStreamWriter_WriteTextElement2(QtObjectPtr ptr, char* qualifiedName, char* text){
	static_cast<QXmlStreamWriter*>(ptr)->writeTextElement(QString(qualifiedName), QString(text));
}

void QXmlStreamWriter_DestroyQXmlStreamWriter(QtObjectPtr ptr){
	static_cast<QXmlStreamWriter*>(ptr)->~QXmlStreamWriter();
}

