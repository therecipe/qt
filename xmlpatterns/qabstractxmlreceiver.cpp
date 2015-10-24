#include "qabstractxmlreceiver.h"
#include <QModelIndex>
#include <QStringRef>
#include <QXmlName>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractXmlReceiver>
#include "_cgo_export.h"

class MyQAbstractXmlReceiver: public QAbstractXmlReceiver {
public:
};

void QAbstractXmlReceiver_Attribute(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value){
	static_cast<QAbstractXmlReceiver*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Characters(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QAbstractXmlReceiver*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Comment(QtObjectPtr ptr, char* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->comment(QString(value));
}

void QAbstractXmlReceiver_EndDocument(QtObjectPtr ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endDocument();
}

void QAbstractXmlReceiver_EndElement(QtObjectPtr ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endElement();
}

void QAbstractXmlReceiver_EndOfSequence(QtObjectPtr ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endOfSequence();
}

void QAbstractXmlReceiver_NamespaceBinding(QtObjectPtr ptr, QtObjectPtr name){
	static_cast<QAbstractXmlReceiver*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_ProcessingInstruction(QtObjectPtr ptr, QtObjectPtr target, char* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->processingInstruction(*static_cast<QXmlName*>(target), QString(value));
}

void QAbstractXmlReceiver_StartDocument(QtObjectPtr ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->startDocument();
}

void QAbstractXmlReceiver_StartElement(QtObjectPtr ptr, QtObjectPtr name){
	static_cast<QAbstractXmlReceiver*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_StartOfSequence(QtObjectPtr ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->startOfSequence();
}

void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(QtObjectPtr ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->~QAbstractXmlReceiver();
}

