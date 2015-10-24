#include "qxmlname.h"
#include <QModelIndex>
#include <QXmlNamePool>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QXmlName>
#include "_cgo_export.h"

class MyQXmlName: public QXmlName {
public:
};

QtObjectPtr QXmlName_NewQXmlName(){
	return new QXmlName();
}

QtObjectPtr QXmlName_NewQXmlName2(QtObjectPtr namePool, char* localName, char* namespaceURI, char* prefix){
	return new QXmlName(*static_cast<QXmlNamePool*>(namePool), QString(localName), QString(namespaceURI), QString(prefix));
}

int QXmlName_QXmlName_IsNCName(char* candidate){
	return QXmlName::isNCName(QString(candidate));
}

int QXmlName_IsNull(QtObjectPtr ptr){
	return static_cast<QXmlName*>(ptr)->isNull();
}

char* QXmlName_LocalName(QtObjectPtr ptr, QtObjectPtr namePool){
	return static_cast<QXmlName*>(ptr)->localName(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_NamespaceUri(QtObjectPtr ptr, QtObjectPtr namePool){
	return static_cast<QXmlName*>(ptr)->namespaceUri(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_Prefix(QtObjectPtr ptr, QtObjectPtr namePool){
	return static_cast<QXmlName*>(ptr)->prefix(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_ToClarkName(QtObjectPtr ptr, QtObjectPtr namePool){
	return static_cast<QXmlName*>(ptr)->toClarkName(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

