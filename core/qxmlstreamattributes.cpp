#include "qxmlstreamattributes.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QXmlStreamAttribute>
#include <QString>
#include <QXmlStreamAttributes>
#include "_cgo_export.h"

class MyQXmlStreamAttributes: public QXmlStreamAttributes {
public:
};

QtObjectPtr QXmlStreamAttributes_NewQXmlStreamAttributes(){
	return new QXmlStreamAttributes();
}

void QXmlStreamAttributes_Append(QtObjectPtr ptr, char* namespaceUri, char* name, char* value){
	static_cast<QXmlStreamAttributes*>(ptr)->append(QString(namespaceUri), QString(name), QString(value));
}

void QXmlStreamAttributes_Append2(QtObjectPtr ptr, char* qualifiedName, char* value){
	static_cast<QXmlStreamAttributes*>(ptr)->append(QString(qualifiedName), QString(value));
}

int QXmlStreamAttributes_HasAttribute2(QtObjectPtr ptr, QtObjectPtr qualifiedName){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(*static_cast<QLatin1String*>(qualifiedName));
}

int QXmlStreamAttributes_HasAttribute3(QtObjectPtr ptr, char* namespaceUri, char* name){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(QString(namespaceUri), QString(name));
}

int QXmlStreamAttributes_HasAttribute(QtObjectPtr ptr, char* qualifiedName){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(QString(qualifiedName));
}

