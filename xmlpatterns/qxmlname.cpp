#include "qxmlname.h"
#include <QXmlNamePool>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include "_cgo_export.h"

class MyQXmlName: public QXmlName {
public:
};

void* QXmlName_NewQXmlName(){
	return new QXmlName();
}

void* QXmlName_NewQXmlName2(void* namePool, char* localName, char* namespaceURI, char* prefix){
	return new QXmlName(*static_cast<QXmlNamePool*>(namePool), QString(localName), QString(namespaceURI), QString(prefix));
}

int QXmlName_QXmlName_IsNCName(char* candidate){
	return QXmlName::isNCName(QString(candidate));
}

int QXmlName_IsNull(void* ptr){
	return static_cast<QXmlName*>(ptr)->isNull();
}

char* QXmlName_LocalName(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->localName(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_NamespaceUri(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->namespaceUri(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_Prefix(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->prefix(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_ToClarkName(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->toClarkName(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

