#include "qxmlserializer.h"
#include <QXmlName>
#include <QString>
#include <QTextCodec>
#include <QStringRef>
#include <QXmlQuery>
#include <QIODevice>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlSerializer>
#include "_cgo_export.h"

class MyQXmlSerializer: public QXmlSerializer {
public:
};

void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice){
	return new QXmlSerializer(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlSerializer_Attribute(void* ptr, void* name, void* value){
	static_cast<QXmlSerializer*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlSerializer_Characters(void* ptr, void* value){
	static_cast<QXmlSerializer*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlSerializer_Comment(void* ptr, char* value){
	static_cast<QXmlSerializer*>(ptr)->comment(QString(value));
}

void QXmlSerializer_EndDocument(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->endDocument();
}

void QXmlSerializer_EndElement(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->endElement();
}

void* QXmlSerializer_Codec(void* ptr){
	return const_cast<QTextCodec*>(static_cast<QXmlSerializer*>(ptr)->codec());
}

void QXmlSerializer_EndOfSequence(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->endOfSequence();
}

void QXmlSerializer_NamespaceBinding(void* ptr, void* nb){
	static_cast<QXmlSerializer*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(nb));
}

void* QXmlSerializer_OutputDevice(void* ptr){
	return static_cast<QXmlSerializer*>(ptr)->outputDevice();
}

void QXmlSerializer_ProcessingInstruction(void* ptr, void* name, char* value){
	static_cast<QXmlSerializer*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString(value));
}

void QXmlSerializer_SetCodec(void* ptr, void* outputCodec){
	static_cast<QXmlSerializer*>(ptr)->setCodec(static_cast<QTextCodec*>(outputCodec));
}

void QXmlSerializer_StartDocument(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->startDocument();
}

void QXmlSerializer_StartElement(void* ptr, void* name){
	static_cast<QXmlSerializer*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlSerializer_StartOfSequence(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->startOfSequence();
}

