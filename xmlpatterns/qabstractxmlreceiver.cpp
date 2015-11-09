#include "qabstractxmlreceiver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QXmlName>
#include <QAbstractXmlReceiver>
#include "_cgo_export.h"

class MyQAbstractXmlReceiver: public QAbstractXmlReceiver {
public:
};

void QAbstractXmlReceiver_Attribute(void* ptr, void* name, void* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Characters(void* ptr, void* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Comment(void* ptr, char* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->comment(QString(value));
}

void QAbstractXmlReceiver_EndDocument(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endDocument();
}

void QAbstractXmlReceiver_EndElement(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endElement();
}

void QAbstractXmlReceiver_EndOfSequence(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endOfSequence();
}

void QAbstractXmlReceiver_NamespaceBinding(void* ptr, void* name){
	static_cast<QAbstractXmlReceiver*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_ProcessingInstruction(void* ptr, void* target, char* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->processingInstruction(*static_cast<QXmlName*>(target), QString(value));
}

void QAbstractXmlReceiver_StartDocument(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->startDocument();
}

void QAbstractXmlReceiver_StartElement(void* ptr, void* name){
	static_cast<QAbstractXmlReceiver*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_StartOfSequence(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->startOfSequence();
}

void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->~QAbstractXmlReceiver();
}

