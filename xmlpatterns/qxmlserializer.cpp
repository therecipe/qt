#include "qxmlserializer.h"
#include <QVariant>
#include <QTextCodec>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QIODevice>
#include <QXmlQuery>
#include <QStringRef>
#include <QXmlSerializer>
#include "_cgo_export.h"

class MyQXmlSerializer: public QXmlSerializer {
public:
};

QtObjectPtr QXmlSerializer_NewQXmlSerializer(QtObjectPtr query, QtObjectPtr outputDevice){
	return new QXmlSerializer(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlSerializer_Attribute(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value){
	static_cast<QXmlSerializer*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlSerializer_Characters(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QXmlSerializer*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlSerializer_Comment(QtObjectPtr ptr, char* value){
	static_cast<QXmlSerializer*>(ptr)->comment(QString(value));
}

void QXmlSerializer_EndDocument(QtObjectPtr ptr){
	static_cast<QXmlSerializer*>(ptr)->endDocument();
}

void QXmlSerializer_EndElement(QtObjectPtr ptr){
	static_cast<QXmlSerializer*>(ptr)->endElement();
}

QtObjectPtr QXmlSerializer_Codec(QtObjectPtr ptr){
	return const_cast<QTextCodec*>(static_cast<QXmlSerializer*>(ptr)->codec());
}

void QXmlSerializer_EndOfSequence(QtObjectPtr ptr){
	static_cast<QXmlSerializer*>(ptr)->endOfSequence();
}

void QXmlSerializer_NamespaceBinding(QtObjectPtr ptr, QtObjectPtr nb){
	static_cast<QXmlSerializer*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(nb));
}

QtObjectPtr QXmlSerializer_OutputDevice(QtObjectPtr ptr){
	return static_cast<QXmlSerializer*>(ptr)->outputDevice();
}

void QXmlSerializer_ProcessingInstruction(QtObjectPtr ptr, QtObjectPtr name, char* value){
	static_cast<QXmlSerializer*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString(value));
}

void QXmlSerializer_SetCodec(QtObjectPtr ptr, QtObjectPtr outputCodec){
	static_cast<QXmlSerializer*>(ptr)->setCodec(static_cast<QTextCodec*>(outputCodec));
}

void QXmlSerializer_StartDocument(QtObjectPtr ptr){
	static_cast<QXmlSerializer*>(ptr)->startDocument();
}

void QXmlSerializer_StartElement(QtObjectPtr ptr, QtObjectPtr name){
	static_cast<QXmlSerializer*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlSerializer_StartOfSequence(QtObjectPtr ptr){
	static_cast<QXmlSerializer*>(ptr)->startOfSequence();
}

