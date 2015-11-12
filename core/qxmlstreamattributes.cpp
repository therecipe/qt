#include "qxmlstreamattributes.h"
#include <QModelIndex>
#include <QLatin1String>
#include <QXmlStreamAttribute>
#include <QStringRef>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QXmlStreamAttributes>
#include "_cgo_export.h"

class MyQXmlStreamAttributes: public QXmlStreamAttributes {
public:
};

void* QXmlStreamAttributes_NewQXmlStreamAttributes(){
	return new QXmlStreamAttributes();
}

void QXmlStreamAttributes_Append(void* ptr, char* namespaceUri, char* name, char* value){
	static_cast<QXmlStreamAttributes*>(ptr)->append(QString(namespaceUri), QString(name), QString(value));
}

void QXmlStreamAttributes_Append2(void* ptr, char* qualifiedName, char* value){
	static_cast<QXmlStreamAttributes*>(ptr)->append(QString(qualifiedName), QString(value));
}

int QXmlStreamAttributes_HasAttribute2(void* ptr, void* qualifiedName){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(*static_cast<QLatin1String*>(qualifiedName));
}

int QXmlStreamAttributes_HasAttribute3(void* ptr, char* namespaceUri, char* name){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(QString(namespaceUri), QString(name));
}

int QXmlStreamAttributes_HasAttribute(void* ptr, char* qualifiedName){
	return static_cast<QXmlStreamAttributes*>(ptr)->hasAttribute(QString(qualifiedName));
}

void* QXmlStreamAttributes_Value3(void* ptr, void* namespaceUri, void* name){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(*static_cast<QLatin1String*>(namespaceUri), *static_cast<QLatin1String*>(name)));
}

void* QXmlStreamAttributes_Value5(void* ptr, void* qualifiedName){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(*static_cast<QLatin1String*>(qualifiedName)));
}

void* QXmlStreamAttributes_Value2(void* ptr, char* namespaceUri, void* name){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(QString(namespaceUri), *static_cast<QLatin1String*>(name)));
}

void* QXmlStreamAttributes_Value(void* ptr, char* namespaceUri, char* name){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(QString(namespaceUri), QString(name)));
}

void* QXmlStreamAttributes_Value4(void* ptr, char* qualifiedName){
	return new QStringRef(static_cast<QXmlStreamAttributes*>(ptr)->value(QString(qualifiedName)));
}

